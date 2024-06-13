// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/google/uuid"
	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

const MessageAggregate = 2

const (
	MessageCreate = iota
	MessageUpdate
)

var (
	MessageCreateEventType = dapp.NewEventType(MessageCreate, MessageAggregate, Version, Protocol)
	MessageUpdateEventType = dapp.NewEventType(MessageUpdate, MessageAggregate, Version, Protocol)
)

type Message struct {
	ID      uuid.UUID    `json:"id"`
	Channel uuid.UUID    `json:"channel"`
	Author  dapp.Address `json:"author"`
	Content string       `json:"content"`
}

type MessageQuery interface {
	Get(context.Context, uuid.UUID, *pletyvo.QueryOption) ([]*Message, error)
	GetByID(ctx context.Context, channel uuid.UUID, id uuid.UUID) (*Message, error)
}

type MessageCreateInput struct {
	Channel uuid.UUID `json:"channel"`
	Content string    `json:"content"`
}

func (mci MessageCreateInput) Validate() error {
	if len(mci.Content) > 2048 {
		return status.New(pletyvo.CodeInvalidArgument, "invalid content length")
	}

	return nil
}

type MessageUpdateInput struct {
	MessageCreateInput
	Message uuid.UUID `json:"message"`
}

func (mui MessageUpdateInput) Validate() error {
	return mui.MessageCreateInput.Validate()
}

type MessageRepository interface {
	MessageQuery
	Create(context.Context, *Message) error
	Update(context.Context, *MessageUpdateInput) error
}

type MessageService interface {
	MessageQuery
	Create(context.Context, *MessageCreateInput) (*dapp.EventResponse, error)
	Update(context.Context, *MessageUpdateInput) (*dapp.EventResponse, error)
}

type MessageExecutor struct{ repos MessageRepository }

func NewMessageExecutor(repos MessageRepository) *MessageExecutor {
	return &MessageExecutor{repos: repos}
}

func (me MessageExecutor) Create(ctx context.Context, base dapp.EventBase[MessageCreateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	return me.repos.Create(ctx, &Message{
		ID:      base.ID,
		Channel: base.Input.Channel,
		Author:  base.Author,
		Content: base.Input.Content,
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
