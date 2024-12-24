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

func (DeliveryChannel) key(ns, id *uuid.UUID) []byte {
	key := make([]byte, 33)
	key[0] = 0 // TODO: use new format

	copy(key[1:], ns[:])
	copy(key[17:], id[:])

	return key
}

func (dc DeliveryChannel) GetByID(ctx context.Context, id uuid.UUID) (*delivery.Channel, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
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
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	return dc.db.Set(dc.key(network, &event.ID), dc.marshal(event, input), pebble.Sync)
}

func (dc DeliveryChannel) Update(ctx context.Context, event *dapp.SystemEvent, input *delivery.ChannelUpdateInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	key := dc.key(network, nil) // TODO: use new format

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

	return dc.db.Set(key, dc.marshal(event, input.ChannelInput), pebble.Sync)
}
