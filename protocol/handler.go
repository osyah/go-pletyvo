// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package protocol

import (
	"context"

	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
	"github.com/osyah/go-pletyvo/protocol/registry"
)

type Handler struct{ Registry, Delivery dapp.Handler }

func NewHandler(executor *Executor) *Handler {
	return &Handler{
		Registry: registry.NewHandler(executor.Registry),
		Delivery: delivery.NewHandler(executor.Delivery),
	}
}

func (h Handler) Handle(ctx context.Context, event *dapp.Event) error {
	switch event.Body.Type().Protocol() {
	case registry.Protocol:
		return h.Registry.Handle(ctx, event)
	case delivery.Protocol:
		return h.Delivery.Handle(ctx, event)
	default:
		return dapp.ErrInvalidEventType
	}
}
