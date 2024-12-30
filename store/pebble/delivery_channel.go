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

type DeliveryChannel struct{ db *pebble.DB }

func NewDeliveryChannel(db *pebble.DB) *DeliveryChannel {
	return &DeliveryChannel{db: db}
}

func (DeliveryChannel) key(network pletyvo.Network, id *uuid.UUID) []byte {
	key := make([]byte, 21)
	key[4] = DeliveryChannelPrefix

	copy(key[0:4], network[2:6])
	copy(key[6:], id[:])

	return key
}

func (dc DeliveryChannel) GetByID(ctx context.Context, id uuid.UUID) (*delivery.Channel, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	b, closer, err := dc.db.Get(dc.key(network, &id))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var channel delivery.Channel

	if err := dc.unmarshal(b, &channel); err != nil {
		return nil, err
	}

	channel.ID = id

	return &channel, nil
}

func (dc DeliveryChannel) Create(ctx context.Context, event *dapp.SystemEvent, input *delivery.ChannelInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	batch := dc.db.NewBatchWithSize(3)
	defer batch.Close()

	batch.Set(dc.key(network, &event.ID), dc.marshal(event, input), nil)

	saveEventAndHash(batch, network, event, &event.ID)

	return batch.Commit(pebble.Sync)
}

func (dc DeliveryChannel) Update(ctx context.Context, event *dapp.SystemEvent, input *delivery.ChannelUpdateInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	id, err := getAggregate(dc.db, network, &input.Channel)
	if err != nil {
		return err
	}

	key := dc.key(network, &id)

	b, closer, err := dc.db.Get(key)
	if err != nil {
		return err
	}
	defer closer.Close()

	var channel delivery.Channel

	if err := dc.unmarshal(b, &channel); err != nil {
		return err
	}

	if channel.Hash != input.Channel {
		return pletyvo.CodeConflict
	}

	if channel.Author != event.Author {
		return pletyvo.CodePermissionDenied
	}

	batch := dc.db.NewBatchWithSize(3)
	defer batch.Close()

	batch.Set(key, dc.marshal(event, input.ChannelInput), nil)

	saveEventAndHash(batch, network, event, &id)

	return batch.Commit(pebble.Sync)
}
