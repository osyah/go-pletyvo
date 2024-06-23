// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registrypebble

import (
	"context"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/registry"
)

type Network struct{ db *pebble.DB }

func NewNetwork(db *pebble.DB) *Network {
	return &Network{db: db}
}

func (Network) key(id uuid.UUID) []byte {
	key := make([]byte, 17)
	key[0] = registry.Protocol

	copy(key[1:], id[:])

	return key
}

func (n Network) Get(ctx context.Context) (*registry.Network, error) {
	id := ctx.Value(pletyvo.ContextKeyNetwork).(uuid.UUID)

	network, err := n.get(n.key(id))
	if err != nil {
		return nil, err
	}

	network.ID = id

	return network, nil
}

func (n Network) get(key []byte) (*registry.Network, error) {
	b, closer, err := n.db.Get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var network registry.Network

	if err := n.unmarshal(b, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

func (n Network) Create(ctx context.Context, network *registry.Network) error {
	return n.db.Set(n.key(network.ID), n.marshal(network), pebble.Sync)
}

func (n Network) Update(ctx context.Context, input *registry.NetworkUpdateInput) error {
	key := n.key(ctx.Value(pletyvo.ContextKeyNetwork).(uuid.UUID))

	network, err := n.get(key)
	if err != nil {
		return err
	}

	network.Name = input.Name

	return n.db.Set(key, n.marshal(network), pebble.Sync)
}
