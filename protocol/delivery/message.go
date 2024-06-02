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
	Get(context.Context, uuid.UUID, ...*pletyvo.QueryOption) ([]*Message, error)
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
	Get(ctx context.Context, ns, ch uuid.UUID, option *pletyvo.QueryOption) ([]*Message, error)
	GetByID(ctx context.Context, ns, ch, id uuid.UUID) (*Message, error)
	Create(context.Context, uuid.UUID, *Message) error
	Update(context.Context, uuid.UUID, *MessageUpdateInput) error
}

type MessageService interface {
	MessageQuery
	Create(context.Context, *MessageCreateInput) (*dapp.EventResponse, error)
	Update(context.Context, *MessageUpdateInput) (*dapp.EventResponse, error)
}
