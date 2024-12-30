// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"context"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type DeliveryMessage struct{ db *pebble.DB }

func NewDeliveryMessage(db *pebble.DB) *DeliveryMessage {
	return &DeliveryMessage{db: db}
}

func (DeliveryMessage) key(network pletyvo.Network, channel, id *uuid.UUID, sufix byte) []byte {
	var key []byte

	if id.Version() != 0 {
		key = make([]byte, 37)
		copy(key[21:], id[:])
	} else {
		key = make([]byte, 22)
		key[21] = sufix
	}

	key[4] = DeliveryMessagePrefix

	copy(key[0:4], network[2:6])
	copy(key[5:21], channel[:])

	return key
}

func (dm DeliveryMessage) Get(ctx context.Context, ch uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Message, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	messages := make([]*delivery.Message, 0, option.Limit)

	iter, err := dm.db.NewIterWithContext(ctx, &pebble.IterOptions{
		LowerBound: dm.key(network, &ch, &option.After, 0),
		UpperBound: dm.key(network, &ch, &option.Before, 255),
	})
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	var next func() bool

	if option.Order {
		if !iter.First() {
			return nil, pletyvo.CodeNotFound
		}

		next = iter.Next
	} else {
		if !iter.Last() {
			return nil, pletyvo.CodeNotFound
		}

		next = iter.Prev
	}

	if option.After.Version() != 0 {
		if !next() {
			return nil, pletyvo.CodeNotFound
		}
	}

	for i := 0; i < option.Limit; i++ {
		var message delivery.Message

		if err := dm.unmarshal(iter.Value(), &message); err != nil {
			return nil, err
		}

		message.ID = uuid.UUID(iter.Key()[21:37])
		message.Channel = ch

		messages = append(messages, &message)

		if !next() {
			break
		}
	}

	return messages, nil
}

func (dm DeliveryMessage) GetByID(ctx context.Context, ch, id uuid.UUID) (*delivery.Message, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	b, closer, err := dm.db.Get(dm.key(network, &ch, &id, 0))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var message delivery.Message

	if err := dm.unmarshal(b, &message); err != nil {
		return nil, err
	}

	message.ID = id
	message.Channel = ch

	return &message, nil
}

func (dm DeliveryMessage) Create(ctx context.Context, event *dapp.SystemEvent, input *delivery.MessageInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	channel, err := getAggregate(dm.db, network, &input.Channel)
	if err != nil {
		return err
	}

	batch := dm.db.NewBatchWithSize(3)
	defer batch.Close()

	batch.Set(dm.key(network, &channel, &event.ID, 0), dm.marshal(event, input), nil)

	saveEventAndHash(batch, network, event, &event.ID)

	return batch.Commit(pebble.Sync)
}

func (dm DeliveryMessage) Update(ctx context.Context, event *dapp.SystemEvent, input *delivery.MessageUpdateInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	channel, err := getAggregate(dm.db, network, &input.Channel)
	if err != nil {
		return err
	}

	id, err := getAggregate(dm.db, network, &input.Message)
	if err != nil {
		return err
	}

	key := dm.key(network, &channel, &id, 0)

	b, closer, err := dm.db.Get(key)
	if err != nil {
		return err
	}
	defer closer.Close()

	var message delivery.Message

	if err := dm.unmarshal(b, &message); err != nil {
		return err
	}

	if message.Hash != input.Message {
		return pletyvo.CodeConflict
	}

	if message.Author != event.Author {
		return pletyvo.CodePermissionDenied
	}

	batch := dm.db.NewBatchWithSize(3)
	defer batch.Close()

	batch.Set(key, dm.marshal(event, input.MessageInput), nil)

	saveEventAndHash(batch, network, event, &id)

	return batch.Commit(pebble.Sync)
}
