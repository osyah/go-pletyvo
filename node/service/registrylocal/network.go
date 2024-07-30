// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registrylocal

import (
	"context"

	"github.com/osyah/go-pletyvo/protocol/registry"
)

type Network struct{ query registry.NetworkQuery }

func NewNetwork(query registry.NetworkQuery) *Network {
	return &Network{query: query}
}

func (n Network) Get(ctx context.Context) (*registry.Network, error) {
	return n.query.Get(ctx)
}
