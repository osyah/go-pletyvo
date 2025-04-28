// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/dapp"
)

const (
	messagesPath = channelPath + "/messages"
	messagePath  = messagesPath + "/%s"
)

type MessageClient struct{ engine pletyvo.DefaultEngine }

func NewMessageClient(engine pletyvo.DefaultEngine) *MessageClient {
	return &MessageClient{engine: engine}
}

func (mc MessageClient) Get(ctx context.Context, channel uuid.UUID, option *pletyvo.QueryOption) ([]*Message, error) {
	messages := make([]*Message, option.Limit)

	if err := mc.engine.Get(ctx, (fmt.Sprintf(messagesPath, channel) + option.String()), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func (mc MessageClient) GetByID(ctx context.Context, channel uuid.UUID, id uuid.UUID) (*Message, error) {
	var message Message

	if err := mc.engine.Get(ctx, fmt.Sprintf(postPath, channel, id), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (mc MessageClient) Send(ctx context.Context, message *dapp.EventInput) error {
	return mc.engine.Post(ctx, "/delivery/v1/channel/send", message, nil)
}
