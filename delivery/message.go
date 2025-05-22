// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/google/uuid"
	"github.com/osyah/hryzun"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/dapp"
)

const MessageCreate = 768

var ErrInvalidMessageTime = hryzun.NewStatus(
	pletyvo.CodeInvalidArgument, "invalid message time",
)

type MessageInput struct {
	ID      uuid.UUID `json:"id"`
	Channel dapp.Hash `json:"channel"`
	Content string    `json:"content"`
}

func (mi MessageInput) Validate() error {
	if len(mi.Content) > 2048 || len(mi.Content) == 0 {
		return hryzun.NewStatus(pletyvo.CodeInvalidArgument, "invalid content length")
	}

	return nil
}

type MessageQuery interface {
	Get(context.Context, uuid.UUID, *pletyvo.QueryOption) ([]*dapp.Event, error)
	GetByID(ctx context.Context, channel uuid.UUID, id uuid.UUID) (*dapp.Event, error)
}

type MessageRepository interface {
	MessageQuery
	Create(context.Context, *dapp.EventInput, *MessageInput) error
}

type MessageService interface {
	MessageQuery
	Send(context.Context, *dapp.EventInput) error
}
