// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package localdoer

import (
	"context"

	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Config struct {
	Repos string `cfg:"repos"`
}

type Relay struct {
	handler *dapp.Handler
	repos   dapp.EventRepository
	hash    dapp.HashRepository
}

func New(handler *dapp.Handler, repos *dapp.Repository) *Relay {
	return &Relay{handler: handler, repos: repos.Event, hash: repos.Hash}
}

func (r Relay) OnEvent(ctx context.Context, event *dapp.SystemEvent) error {
	err := r.handler.Handle(ctx, event)
	if err != nil {
		return err
	}

	if err = r.repos.Create(ctx, event); err != nil {
		return err
	}

	return r.hash.Create(ctx, event.EventHeader)
}

func Register(config Config) container.Handler {
	return func(base *container.Base) any {
		return New(
			container.Get[*dapp.Handler](base, "handler"),
			container.Get[*dapp.Repository](base, config.Repos),
		)
	}
}
