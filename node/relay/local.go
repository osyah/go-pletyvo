// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package relay

import (
	"context"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Local struct {
	handler dapp.Handler
	repos   dapp.EventRepository
}

func NewLocal(handler dapp.Handler, repos dapp.EventRepository) *Local {
	return &Local{handler: handler, repos: repos}
}

func (l Local) OnEvent(ctx context.Context, event *dapp.Event) error {
	if err := l.handler.Handle(ctx, event); err != nil {
		return err
	}

	return l.repos.Create(ctx, event)
}
