// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package deliveryhttp

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/client/engine"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

const (
	messagesPath = channelPath + "/messages"
	messagePath  = messagesPath + "/%s"
)

type Message struct{ engine engine.HTTP }

func NewMessage(engine engine.HTTP) *Message {
	return &Message{engine: engine}
}

func (m Message) Get(ctx context.Context, channel uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Message, error) {
	messages := make([]*delivery.Message, option.Limit)

	if err := m.engine.Get(ctx, (fmt.Sprintf(messagesPath, channel) + option.String()), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func (m Message) GetByID(ctx context.Context, channel uuid.UUID, id uuid.UUID) (*delivery.Message, error) {
	var message delivery.Message

	if err := m.engine.Get(ctx, fmt.Sprintf(postPath, channel, id), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (m Message) Send(ctx context.Context, message *delivery.Message) error {
	return m.engine.Post(ctx, "/delivery/v1/channel/send", message, nil)
}
