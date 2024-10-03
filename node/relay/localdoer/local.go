// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package localdoer

import (
	"context"

	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/protocol"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Config struct {
	Repos string `cfg:"repos"`
}

type Relay struct {
	handler dapp.Handler
	repos   dapp.EventRepository
}

func New(handler dapp.Handler, repos dapp.EventRepository) *Relay {
	return &Relay{handler: handler, repos: repos}
}

func (r Relay) OnEvent(ctx context.Context, event *dapp.Event) error {
	if err := r.handler.Handle(ctx, event); err != nil {
		return err
	}

	return r.repos.Create(ctx, event)
}

func Register(config Config) container.Handler {
	return func(base *container.Base) any {
		return New(
			protocol.NewHandler(
				container.Get[*protocol.Executor](base, "executor"),
			),
			container.Get[*dapp.Repository](base, config.Repos).Event,
		)
	}
}
