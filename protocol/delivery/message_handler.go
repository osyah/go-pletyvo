// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type MessageHandler struct{ service *MessageExecutor }

func NewMessageHandler(service *MessageExecutor) *MessageHandler {
	return &MessageHandler{service: service}
}

func (mh MessageHandler) Handle(ctx context.Context, event *dapp.Event) error {
	switch event.Body.Type().Event() {
	case MessageCreate:
		var input MessageCreateInput

		if err := event.Body.Unmarshal(&input); err != nil {
			return err
		}

		return mh.service.Create(ctx, dapp.EventBase[MessageCreateInput]{
			EventHeader: event.EventHeader,
			Input:       &input,
		})
	case MessageUpdate:
		var input MessageUpdateInput

		if err := event.Body.Unmarshal(&input); err != nil {
			return err
		}

		return mh.service.Update(ctx, dapp.EventBase[MessageUpdateInput]{
			EventHeader: event.EventHeader,
			Input:       &input,
		})
	default:
		return dapp.ErrInvalidEventType
	}
}
