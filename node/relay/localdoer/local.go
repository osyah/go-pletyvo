// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package localdoer

import (
	"context"

	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Relay struct{ handler *dapp.Handler }

func New(handler *dapp.Handler) *Relay {
	return &Relay{handler: handler}
}

func (r Relay) OnEvent(ctx context.Context, event *dapp.SystemEvent) error {
	return r.handler.Handle(ctx, event)
}

func Register() container.Handler {
	return func(base *container.Base) any {
		return New(container.Get[*dapp.Handler](base, "handler"))
	}
}
