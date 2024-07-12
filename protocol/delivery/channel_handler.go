// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type ChannelHandler struct{ service *ChannelExecutor }

func NewChannelHandler(service *ChannelExecutor) *ChannelHandler {
	return &ChannelHandler{service: service}
}

func (ch ChannelHandler) Handle(ctx context.Context, event *dapp.Event) error {
	switch event.Body.Type().Event() {
	case ChannelCreate:
		var input ChannelCreateInput

		if err := event.Body.Unmarshal(&input); err != nil {
			return err
		}

		return ch.service.Create(ctx, dapp.EventBase[ChannelCreateInput]{
			EventHeader: event.EventHeader,
			Input:       &input,
		})
	case ChannelUpdate:
		var input ChannelUpdateInput

		if err := event.Body.Unmarshal(&input); err != nil {
			return err
		}

		return ch.service.Update(ctx, dapp.EventBase[ChannelUpdateInput]{
			EventHeader: event.EventHeader,
			Input:       &input,
		})
	default:
		return status.New(pletyvo.CodeInvalidArgument, "unsupported body type")
	}
}
