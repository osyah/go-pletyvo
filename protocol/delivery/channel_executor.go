// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type ChannelExecutor struct{ repos ChannelRepository }

func NewChannelExecutor(repos ChannelRepository) *ChannelExecutor {
	return &ChannelExecutor{repos: repos}
}

func (ce ChannelExecutor) Create(ctx context.Context, base dapp.EventBase[ChannelCreateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	return ce.repos.Create(ctx, &Channel{
		EventHeader:  base.EventHeader,
		ChannelInput: base.Input.ChannelInput,
	})
}

func (ce ChannelExecutor) Update(ctx context.Context, base dapp.EventBase[ChannelUpdateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	channel, err := ce.repos.GetByID(ctx, base.Input.Channel)
	if err != nil {
		return err
	}

	if channel.Author != base.Author {
		return pletyvo.CodePermissionDenied
	}

	return ce.repos.Update(ctx, base.Input)
}
