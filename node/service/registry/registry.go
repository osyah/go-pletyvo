// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registry

import "github.com/osyah/go-pletyvo/protocol/registry"

func NewQuery(query *registry.Query) *registry.Query {
	return &registry.Query{Network: NewNetwork(query.Network)}
}
