// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registry

import (
	"context"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Handler struct{ Network *NetworkHandler }

func NewHandler(executor *Executor) *Handler {
	return &Handler{Network: NewNetworkHandler(executor.Network)}
}

func (h Handler) Handle(ctx context.Context, event *dapp.Event) error {
	if t := event.Body.Type(); t.Version() == Version {
		if t.Aggregate() == NetworkAggregate {
			return h.Network.Handle(ctx, event)
		}
	}

	return dapp.ErrInvalidEventType
}
