// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/istyna/go-pletyvo"
	"github.com/istyna/go-pletyvo/dapp"
)

const (
	messagesPath = channelPath + "/messages"
	messagePath  = messagesPath + "/%s"
)

type MessageClient struct{ engine pletyvo.DefaultEngine }

func NewMessageClient(engine pletyvo.DefaultEngine) *MessageClient {
	return &MessageClient{engine: engine}
}

func (mc MessageClient) Get(ctx context.Context, channel uuid.UUID, option *pletyvo.QueryOption) ([]*dapp.Event, error) {
	events := make([]*dapp.Event, option.Limit)

	if err := mc.engine.Get(ctx, (fmt.Sprintf(messagesPath, channel) + option.String()), &events); err != nil {
		return nil, err
	}

	return events, nil
}

func (mc MessageClient) GetByID(ctx context.Context, channel uuid.UUID, id uuid.UUID) (*dapp.Event, error) {
	var event dapp.Event

	if err := mc.engine.Get(ctx, fmt.Sprintf(postPath, channel, id), &event); err != nil {
		return nil, err
	}

	return &event, nil
}

func (mc MessageClient) Send(ctx context.Context, message *dapp.EventInput) error {
	return mc.engine.Post(ctx, "/delivery/v1/channel/send", message, nil)
}
