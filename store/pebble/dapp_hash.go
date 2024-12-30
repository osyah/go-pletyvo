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

type DAppHash struct{ db *pebble.DB }

func NewDAppHash(db *pebble.DB) *DAppHash {
	return &DAppHash{db: db}
}

func dAppHashKey(network pletyvo.Network, hash *dapp.Hash) []byte {
	key := make([]byte, 37)
	key[4] = DAppHashPrefix

	copy(key[0:4], network[2:6])
	copy(key[5:], hash[:])

	return key
}

func (dah DAppHash) GetByID(ctx context.Context, hash dapp.Hash) (*dapp.EventResponse, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	b, closer, err := dah.db.Get(dAppHashKey(network, &hash))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	return &dapp.EventResponse{ID: uuid.UUID(b[:16])}, nil
}

func (dah DAppHash) Create(ctx context.Context, header *dapp.EventHeader) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	data := make([]byte, 32)
	copy(data[:16], header.ID[:])

	return dah.db.Set(dAppHashKey(network, &header.Hash), data, pebble.Sync)
}

func getAggregate(db *pebble.DB, network pletyvo.Network, hash *dapp.Hash) (uuid.UUID, error) {
	b, closer, err := db.Get(dAppHashKey(network, hash))
	if err != nil {
		return uuid.Nil, err
	}
	defer closer.Close()

	return uuid.UUID(b[16:]), nil
}
