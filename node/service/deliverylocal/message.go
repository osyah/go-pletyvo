// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package deliverylocal

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Message struct{ repos delivery.MessageRepository }

func NewMessage(repos delivery.MessageRepository) *Message {
	return &Message{repos: repos}
}

func (m Message) Get(ctx context.Context, id uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Message, error) {
	option.Prepare()

	return m.repos.Get(ctx, id, option)
}

func (m Message) GetByID(ctx context.Context, channel uuid.UUID, id uuid.UUID) (*delivery.Message, error) {
	return m.repos.GetByID(ctx, channel, id)
}

func (m Message) Send(ctx context.Context, message *delivery.Message) error {
	if message.Body.Version() != dapp.EventBodyBasic {
		return dapp.ErrInvalidEventBodyVersion
	}

	if message.Body.Type() != delivery.MessageCreate {
		return dapp.ErrInvalidEventType
	}

	var input delivery.MessageInput

	err := message.Body.Unmarshal(&input)
	if err != nil {
		return pletyvo.CodeInvalidArgument
	}

	if err = input.Validate(); err != nil {
		return err
	}

	if input.ID.Version() != 7 {
		return pletyvo.CodeInvalidArgument
	}

	sec, _ := input.ID.Time().UnixTime()
	interval := time.Now().Unix() - sec
	if interval > 5 || interval < -5 {
		return delivery.ErrInvalidMessageTime
	}

	return m.repos.Create(ctx, message, &input)
}
