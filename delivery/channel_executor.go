// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/osyah/go-pletyvo/dapp"
)

type ChannelExecutor struct{ repos ChannelRepository }

func NewChannelExecutor(repos ChannelRepository) *ChannelExecutor {
	return &ChannelExecutor{repos: repos}
}

func (ce ChannelExecutor) Register(handler *dapp.Handler) {
	handler.Register(ChannelCreate, ce.Create)
	handler.Register(ChannelUpdate, ce.Update)
}

func (ce ChannelExecutor) Create(ctx context.Context, event *dapp.SystemEvent) error {
	var input ChannelInput

	if err := event.Body.Unmarshal(&input); err != nil {
		return err
	}

	if err := input.Validate(); err != nil {
		return err
	}

	return ce.repos.Create(ctx, event, &input)
}

func (ce ChannelExecutor) Update(ctx context.Context, event *dapp.SystemEvent) error {
	var input ChannelUpdateInput

	err := event.Body.Unmarshal(&input)
	if err != nil {
		return err
	}

	if err = input.Validate(); err != nil {
		return err
	}

	return ce.repos.Update(ctx, event, &input)
}
