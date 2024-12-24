// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"context"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type DAppEvent struct{ db *pebble.DB }

func NewDAppEvent(db *pebble.DB) *DAppEvent {
	return &DAppEvent{db: db}
}

func (DAppEvent) key(ns, id *uuid.UUID, sufix byte) []byte {
	var key []byte

	if id.Version() != 0 {
		key = make([]byte, 33)
		copy(key[17:], id[:])
	} else {
		key = make([]byte, 18)
		key[17] = sufix
	}

	key[0] = 0 // TODO: use new format
	copy(key[1:], ns[:])

	return key
}

func (dae DAppEvent) Get(ctx context.Context, option *pletyvo.QueryOption) ([]*dapp.Event, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	events := make([]*dapp.Event, 0, option.Limit)

	iter, err := dae.db.NewIterWithContext(ctx, &pebble.IterOptions{
		LowerBound: dae.key(network, &option.After, 0),
		UpperBound: dae.key(network, &option.Before, 255),
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
		var event dapp.Event

		if err := dae.unmarshal(iter.Value(), &event); err != nil {
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

func (dae DAppEvent) GetByID(ctx context.Context, id uuid.UUID) (*dapp.Event, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	b, closer, err := dae.db.Get(dae.key(network, &id, 0))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var event dapp.Event

	if err := dae.unmarshal(b, &event); err != nil {
		return nil, err
	}

	event.ID = id

	return &event, nil
}

func (dae DAppEvent) Create(ctx context.Context, event *dapp.SystemEvent) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
	if !ok {
		network = &uuid.Nil
	}

	return dae.db.Set(dae.key(network, &event.ID, 0), dae.marshal(event), pebble.Sync)
}
