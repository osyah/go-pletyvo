// Copyright (c) 2025 Osyah
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

func (dm DeliveryMessage) key(network pletyvo.Network, channel, id *uuid.UUID, sufix byte) []byte {
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

func (dm DeliveryMessage) Get(ctx context.Context, ch uuid.UUID, opt *pletyvo.QueryOption) ([]*delivery.Message, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	messages := make([]*delivery.Message, 0, opt.Limit)

	iter, err := dm.db.NewIterWithContext(ctx, &pebble.IterOptions{
		LowerBound: dm.key(network, &ch, &opt.After, 0),
		UpperBound: dm.key(network, &ch, &opt.Before, 255),
	})
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	var next func() bool

	if opt.Order {
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

	if opt.After.Version() != 0 {
		if !next() {
			return nil, pletyvo.CodeNotFound
		}
	}

	for i := 0; i < opt.Limit; i++ {
		var message delivery.Message

		if err := dm.unmarshal(iter.Value(), &message); err != nil {
			return nil, err
		}

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

	return &message, nil
}

func (dm DeliveryMessage) Create(ctx context.Context, message *dapp.EventInput, input *delivery.MessageInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	channel, err := getAggregate(dm.db, network, &input.Channel)
	if err != nil {
		return err
	}

	key := dm.key(network, &channel, &input.ID, 0)

	_, closer, err := dm.db.Get(key)
	if err != nil {
		if err == pebble.ErrNotFound {
			return dm.db.Set(key, dm.marshal(message), pebble.Sync)
		}

		return err
	}
	defer closer.Close()

	return pletyvo.CodeConflict
}
