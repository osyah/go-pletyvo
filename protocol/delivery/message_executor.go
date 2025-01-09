// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type MessageExecutor struct {
	repos   MessageRepository
	channel ChannelRepository
}

func NewMessageExecutor(repos MessageRepository, channel ChannelRepository) *MessageExecutor {
	return &MessageExecutor{repos: repos, channel: channel}
}

func (me MessageExecutor) Register(handler *dapp.Handler) {
	handler.Register(MessageCreateEventType, me.Create)
	handler.Register(MessageUpdateEventType, me.Update)
}

func (me MessageExecutor) Create(ctx context.Context, event *dapp.SystemEvent) error {
	var input MessageInput

	if err := event.Body.Unmarshal(&input); err != nil {
		return err
	}

	if err := input.Validate(); err != nil {
		return err
	}

	channel, err := me.channel.GetByHash(ctx, input.Channel)
	if err != nil {
		return err
	}

	if channel.Author != event.Author {
		return pletyvo.CodePermissionDenied
	}

	return me.repos.Create(ctx, event, &input)
}

func (me MessageExecutor) Update(ctx context.Context, event *dapp.SystemEvent) error {
	var input MessageUpdateInput

	if err := event.Body.Unmarshal(&input); err != nil {
		return err
	}

	if err := input.Validate(); err != nil {
		return err
	}

	channel, err := me.channel.GetByHash(ctx, input.Channel)
	if err != nil {
		return err
	}

	if channel.Author != event.Author {
		return pletyvo.CodePermissionDenied
	}

	return me.repos.Update(ctx, event, &input)
}
