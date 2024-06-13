// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapppebble

import (
	"context"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Event struct{ db *pebble.DB }

func NewEvent(db *pebble.DB) *Event {
	return &Event{db: db}
}

func (Event) key(ns, id uuid.UUID, sufix ...byte) []byte {
	var key []byte

	if id != uuid.Nil {
		key = make([]byte, 33)
		copy(key[17:], id[:])
	} else {
		key = make([]byte, 18)
		key[17] = sufix[0]
	}

	key[0] = dapp.Protocol
	copy(key[1:], ns[:])

	return key
}

func (e Event) Get(ctx context.Context, option *pletyvo.QueryOption) ([]*dapp.Event, error) {
	network := ctx.Value(pletyvo.ContextKeyNetwork).(uuid.UUID)
	events := make([]*dapp.Event, 0, option.Limit)

	iter, err := e.db.NewIterWithContext(ctx, &pebble.IterOptions{
		LowerBound: e.key(network, option.After, 0),
		UpperBound: e.key(network, option.Before, 255),
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

	if option.After != uuid.Nil {
		if !next() {
			return nil, pletyvo.CodeNotFound
		}
	}

	for i := 0; i < option.Limit; i++ {
		var event dapp.Event

		if err := e.unmarshal(iter.Value(), &event); err != nil {
			return nil, err
		}

		event.ID = uuid.UUID(iter.Key()[17:33])
		events = append(events, &event)

		if !next() {
			break
		}
	}

	return events, nil
}

func (e Event) GetByID(ctx context.Context, id uuid.UUID) (*dapp.Event, error) {
	network := ctx.Value(pletyvo.ContextKeyNetwork).(uuid.UUID)

	b, closer, err := e.db.Get(e.key(network, id))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var event dapp.Event

	if err := e.unmarshal(b, &event); err != nil {
		return nil, err
	}

	event.ID = id

	return &event, nil
}

func (e Event) Create(ctx context.Context, event *dapp.Event) error {
	network := ctx.Value(pletyvo.ContextKeyNetwork).(uuid.UUID)
	return e.db.Set(e.key(network, event.ID), e.marshal(event), pebble.Sync)
}
