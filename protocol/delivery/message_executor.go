// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type MessageExecutor struct{ repos MessageRepository }

func NewMessageExecutor(repos MessageRepository) *MessageExecutor {
	return &MessageExecutor{repos: repos}
}

func (me MessageExecutor) Create(ctx context.Context, base dapp.EventBase[MessageCreateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	return me.repos.Create(ctx, &Message{
		EventHeader:  base.EventHeader,
		MessageInput: base.Input.MessageInput,
	})
}

func (me MessageExecutor) Update(ctx context.Context, base dapp.EventBase[MessageUpdateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	message, err := me.repos.GetByID(ctx, base.Input.Channel, base.Input.Message)
	if err != nil {
		return err
	}

	if message.Author != base.Author {
		return pletyvo.CodePermissionDenied
	}

	return me.repos.Update(ctx, base.Input)
}
