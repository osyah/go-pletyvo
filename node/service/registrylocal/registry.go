// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registrylocal

import (
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/protocol/registry"
)

type Config struct {
	Repos string `cfg:"repos"`
}

func NewQuery(repos *registry.Repository) *registry.Query {
	return &registry.Query{Network: NewNetwork(repos.Network)}
}

func RegisterQuery(config Config) container.Handler {
	return func(base *container.Base) any {
		return NewQuery(
			container.Get[*registry.Repository](base, config.Repos),
		)
	}
}
