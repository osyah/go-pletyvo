// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package deliverypebble

import (
	"context"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Channel struct{ db *pebble.DB }

func NewChannel(db *pebble.DB) *Channel {
	return &Channel{db: db}
}

func (Channel) key(ns, id uuid.UUID) []byte {
	key := make([]byte, 33)
	key[0] = delivery.Protocol

	copy(key[1:], ns[:])
	copy(key[17:], id[:])

	return key
}

func (c Channel) GetByID(ctx context.Context, ns, id uuid.UUID) (*delivery.Channel, error) {
	channel, err := c.getByID(c.key(ns, id))
	if err != nil {
		return nil, err
	}

	channel.ID = id

	return channel, nil
}

func (c Channel) getByID(key []byte) (*delivery.Channel, error) {
	b, closer, err := c.db.Get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var channel delivery.Channel

	if err := c.unmarshal(b, &channel); err != nil {
		return nil, err
	}

	return &channel, nil
}

func (c Channel) Create(ctx context.Context, ns uuid.UUID, channel *delivery.Channel) error {
	return c.db.Set(c.key(ns, channel.ID), c.marshal(channel), pebble.Sync)
}

func (c Channel) Update(ctx context.Context, ns uuid.UUID, input *delivery.ChannelUpdateInput) error {
	key := c.key(ns, input.Channel)

	channel, err := c.getByID(key)
	if err != nil {
		return err
	}

	channel.Name = input.Name

	return c.db.Set(key, c.marshal(channel), pebble.Sync)
}
