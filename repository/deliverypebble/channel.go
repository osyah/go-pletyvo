// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package deliverypebble

import (
	"context"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
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

func (c Channel) GetByID(ctx context.Context, id uuid.UUID) (*delivery.Channel, error) {
	network := ctx.Value(pletyvo.ContextKeyNetwork).(uuid.UUID)

	channel, err := c.getByID(c.key(network, id))
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

func (c Channel) Create(ctx context.Context, channel *delivery.Channel) error {
	network := ctx.Value(pletyvo.ContextKeyNetwork).(uuid.UUID)
	return c.db.Set(c.key(network, channel.ID), c.marshal(channel), pebble.Sync)
}

func (c Channel) Update(ctx context.Context, input *delivery.ChannelUpdateInput) error {
	key := c.key(ctx.Value(pletyvo.ContextKeyNetwork).(uuid.UUID), input.Channel)

	channel, err := c.getByID(key)
	if err != nil {
		return err
	}

	channel.Name = input.Name

	return c.db.Set(key, c.marshal(channel), pebble.Sync)
}
