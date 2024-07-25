// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Message struct{ query delivery.MessageQuery }

func NewMessage(query delivery.MessageQuery) *Message {
	return &Message{query: query}
}

func (m Message) Get(ctx context.Context, id uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Message, error) {
	option.Prepare()

	return m.query.Get(ctx, id, option)
}

func (m Message) GetByID(ctx context.Context, channel, id uuid.UUID) (*delivery.Message, error) {
	return m.query.GetByID(ctx, channel, id)
}
