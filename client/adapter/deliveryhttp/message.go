// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package deliveryhttp

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/client/engine"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/dapp/crypto"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

const (
	messagesPath = channelPath + "/messages"
	messagePath  = messagesPath + "/%s"
)

type Message struct {
	engine engine.HTTP
	signer crypto.Signer
	event  dapp.EventService
}

func NewMessage(engine engine.HTTP, signer crypto.Signer, event dapp.EventService) *Message {
	return &Message{
		engine: engine,
		signer: signer,
		event:  event,
	}
}

func (m Message) Get(ctx context.Context, ch uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Message, error) {
	messages := make([]*delivery.Message, option.Limit)

	if err := m.engine.Get(ctx, (fmt.Sprintf(messagesPath, ch) + option.String()), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func (m Message) GetByID(ctx context.Context, ch uuid.UUID, id uuid.UUID) (*delivery.Message, error) {
	var message delivery.Message

	if err := m.engine.Get(ctx, fmt.Sprintf(messagePath, ch, id), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (m Message) Create(ctx context.Context, input *delivery.MessageCreateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBodyJSON(input, delivery.MessageCreateEventType)

	return m.event.Create(ctx, &dapp.EventInput{
		Body: body,
		Auth: m.signer.Auth(body),
	})
}

func (m Message) Update(ctx context.Context, input *delivery.MessageUpdateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBodyJSON(input, delivery.MessageUpdateEventType)

	return m.event.Create(ctx, &dapp.EventInput{
		Body: body,
		Auth: m.signer.Auth(body),
	})
}
