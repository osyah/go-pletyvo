// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Handler struct {
	Channel *ChannelHandler
	Message *MessageHandler
}

func NewHandler(executor *Executor) *Handler {
	return &Handler{
		Channel: NewChannelHandler(executor.Channel),
		Message: NewMessageHandler(executor.Message),
	}
}

func (h Handler) Handle(ctx context.Context, event *dapp.Event) error {
	if t := event.Body.Type(); t.Version() == Version {
		switch t.Aggregate() {
		case ChannelAggregate:
			return h.Channel.Handle(ctx, event)
		case MessageAggregate:
			return h.Message.Handle(ctx, event)
		}
	}

	return dapp.ErrInvalidEventType
}
