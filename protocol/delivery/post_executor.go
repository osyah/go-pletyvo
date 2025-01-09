// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type PostExecutor struct {
	repos   PostRepository
	channel ChannelRepository
}

func NewPostExecutor(repos PostRepository, channel ChannelRepository) *PostExecutor {
	return &PostExecutor{repos: repos, channel: channel}
}

func (pe PostExecutor) Register(handler *dapp.Handler) {
	handler.Register(PostCreateEventType, pe.Create)
	handler.Register(PostUpdateEventType, pe.Update)
}

func (pe PostExecutor) Create(ctx context.Context, event *dapp.SystemEvent) error {
	var input PostInput

	if err := event.Body.Unmarshal(&input); err != nil {
		return err
	}

	if err := input.Validate(); err != nil {
		return err
	}

	channel, err := pe.channel.GetByHash(ctx, input.Channel)
	if err != nil {
		return err
	}

	if channel.Author != event.Author {
		return pletyvo.CodePermissionDenied
	}

	return pe.repos.Create(ctx, event, &input)
}

func (pe PostExecutor) Update(ctx context.Context, event *dapp.SystemEvent) error {
	var input PostUpdateInput

	if err := event.Body.Unmarshal(&input); err != nil {
		return err
	}

	if err := input.Validate(); err != nil {
		return err
	}

	channel, err := pe.channel.GetByHash(ctx, input.Channel)
	if err != nil {
		return err
	}

	if channel.Author != event.Author {
		return pletyvo.CodePermissionDenied
	}

	return pe.repos.Update(ctx, event, &input)
}
