// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapplocal

import (
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/node/relay"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Config struct {
	Repos string `cfg:"repos"`
	Relay string `cfg:"relay"`
}

func New(repos *dapp.Repository, relay relay.Relay) *dapp.Service {
	return &dapp.Service{Event: NewEvent(repos.Event, relay)}
}

func Register(config Config) container.Handler {
	return func(base *container.Base) any {
		return New(
			container.Get[*dapp.Repository](base, config.Repos),
			container.Get[relay.Relay](base, config.Relay),
		)
	}
}
