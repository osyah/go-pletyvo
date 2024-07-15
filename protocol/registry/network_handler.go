// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registry

import (
	"context"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type NetworkHandler struct{ service *NetworkExecutor }

func NewNetworkHandler(service *NetworkExecutor) *NetworkHandler {
	return &NetworkHandler{service: service}
}

func (nh NetworkHandler) Handle(ctx context.Context, event *dapp.Event) error {
	switch event.Body.Type().Event() {
	case NetworkCreate:
		var input NetworkCreateInput

		if err := event.Body.Unmarshal(&input); err != nil {
			return err
		}

		id, ok := ctx.Value(pletyvo.ContextKeyNetwork).(*uuid.UUID)
		if ok {
			copy(id[:], event.ID[:])
		}

		return nh.service.Create(ctx, dapp.EventBase[NetworkCreateInput]{
			EventHeader: event.EventHeader,
			Input:       &input,
		})
	case NetworkUpdate:
		var input NetworkUpdateInput

		if err := event.Body.Unmarshal(&input); err != nil {
			return err
		}

		return nh.service.Update(ctx, dapp.EventBase[NetworkUpdateInput]{
			EventHeader: event.EventHeader,
			Input:       &input,
		})
	default:
		return dapp.ErrInvalidEventType
	}
}
