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

func dAppEventKey(network pletyvo.Network, id *uuid.UUID, sufix byte) []byte {
	var key []byte

	if id.Version() != 0 {
		key = make([]byte, 21)
		copy(key[5:], id[:])
	} else {
		key = make([]byte, 6)
		key[5] = sufix
	}

	key[4] = DAppEventPrefix
	copy(key[0:4], network[2:6])

	return key
}

func (dae DAppEvent) Get(ctx context.Context, option *pletyvo.QueryOption) ([]*dapp.Event, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	events := make([]*dapp.Event, 0, option.Limit)

	iter, err := dae.db.NewIterWithContext(ctx, &pebble.IterOptions{
		LowerBound: dAppEventKey(network, &option.After, 0),
		UpperBound: dAppEventKey(network, &option.Before, 255),
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

		event.ID = uuid.UUID(iter.Key()[5:21])
		events = append(events, &event)

		if !next() {
			break
		}
	}

	return events, nil
}

func (dae DAppEvent) GetByID(ctx context.Context, id uuid.UUID) (*dapp.Event, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	b, closer, err := dae.db.Get(dAppEventKey(network, &id, 0))
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
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	batch := dae.db.NewBatchWithSize(2)
	defer batch.Close()

	saveEventAndHash(batch, network, event, nil)

	return batch.Commit(pebble.Sync)
}

func saveEventAndHash(batch *pebble.Batch, network pletyvo.Network, event *dapp.SystemEvent, aggregate *uuid.UUID) {
	batch.Set(dAppEventKey(network, &event.ID, 0), marshalDAppEvent(event), nil)

	data := make([]byte, 32)
	copy(data[:16], event.ID[:])

	if aggregate != nil {
		copy(data[16:], aggregate[:])
	}

	batch.Set(dAppHashKey(network, &event.Hash), data, nil)
}
