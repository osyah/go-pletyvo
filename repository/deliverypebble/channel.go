// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package deliverypebble

import (
	"context"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Channel struct{ db *pebble.DB }

func NewChannel(db *pebble.DB) *Channel {
	return &Channel{db: db}
}

func (Channel) key(ns, id *uuid.UUID) []byte {
	key := make([]byte, 33)
	key[0] = 0 // TODO: use new format

	copy(key[1:], ns[:])
	copy(key[17:], id[:])

	return key
}

func (c Channel) GetByID(ctx context.Context, id uuid.UUID) (*delivery.Channel, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	b, closer, err := c.db.Get(c.key(network, &id))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var channel delivery.Channel

	if err := c.unmarshal(b, &channel); err != nil {
		return nil, err
	}

	channel.ID = id

	return &channel, nil
}

func (c Channel) Create(ctx context.Context, event *dapp.SystemEvent, input *delivery.ChannelInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	return c.db.Set(c.key(network, &event.ID), c.marshal(event, input), pebble.Sync)
}

func (c Channel) Update(ctx context.Context, event *dapp.SystemEvent, input *delivery.ChannelUpdateInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	key := c.key(network, nil) // TODO: use new format

	b, closer, err := c.db.Get(key)
	if err != nil {
		return err
	}
	defer closer.Close()

	var channel delivery.Channel

	if err := c.unmarshal(b, &channel); err != nil {
		return err
	}

	if channel.Hash != input.Channel {
		return pletyvo.CodeConflict
	}

	if channel.Author != event.Author {
		return pletyvo.CodePermissionDenied
	}

	return c.db.Set(key, c.marshal(event, input.ChannelInput), pebble.Sync)
}
