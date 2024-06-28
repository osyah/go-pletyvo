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

type Message struct{ db *pebble.DB }

func NewMessage(db *pebble.DB) *Message {
	return &Message{db: db}
}

func (Message) key(ns, ch, id *uuid.UUID, sufix byte) []byte {
	var key []byte

	if id.Version() != 0 {
		key = make([]byte, 49)
		copy(key[33:], id[:])
	} else {
		key = make([]byte, 34)
		key[33] = sufix
	}

	key[0] = delivery.Protocol

	copy(key[1:], ns[:])
	copy(key[17:], ch[:])

	return key
}

func (m Message) Get(ctx context.Context, ch uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Message, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	messages := make([]*delivery.Message, 0, option.Limit)

	iter, err := m.db.NewIterWithContext(ctx, &pebble.IterOptions{
		LowerBound: m.key(network, &ch, &option.After, 0),
		UpperBound: m.key(network, &ch, &option.Before, 255),
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

		if err := m.unmarshal(iter.Value(), &message); err != nil {
			return nil, err
		}

		message.ID = uuid.UUID(iter.Key()[33:49])
		message.Channel = ch

		messages = append(messages, &message)

		if !next() {
			break
		}
	}

	return messages, nil
}

func (m Message) GetByID(ctx context.Context, ch, id uuid.UUID) (*delivery.Message, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	b, closer, err := m.db.Get(m.key(network, &ch, &id, 0))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var message delivery.Message

	if err := m.unmarshal(b, &message); err != nil {
		return nil, err
	}

	message.ID = id
	message.Channel = ch

	return &message, nil
}

func (m Message) Create(ctx context.Context, message *delivery.Message) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	key := m.key(network, &message.Channel, &message.ID, 0)

	return m.db.Set(key, m.marshal(message), pebble.Sync)
}

func (m Message) Update(ctx context.Context, input *delivery.MessageUpdateInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	key := m.key(network, &input.Channel, &input.Message, 0)

	b, closer, err := m.db.Get(key)
	if err != nil {
		return err
	}
	defer closer.Close()

	var message delivery.Message

	if err := m.unmarshal(b, &message); err != nil {
		return err
	}

	message.Content = input.Content

	return m.db.Set(key, m.marshal(&message), pebble.Sync)
}
