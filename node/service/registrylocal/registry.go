// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registrylocal

import "github.com/osyah/go-pletyvo/protocol/registry"

func NewQuery(repos *registry.Repository) *registry.Query {
	return &registry.Query{Network: NewNetwork(repos.Network)}
}
