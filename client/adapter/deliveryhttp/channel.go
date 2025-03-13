// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package deliveryhttp

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo/client/engine"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/dapp/crypto"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

const channelPath = "/delivery/v1/channel/%s"

type Channel struct {
	engine engine.HTTP
	signer crypto.Signer
	event  dapp.EventService
}

func NewChannel(engine engine.HTTP, signer crypto.Signer, event dapp.EventService) *Channel {
	return &Channel{engine: engine, signer: signer, event: event}
}

func (c Channel) GetByID(ctx context.Context, id uuid.UUID) (*delivery.Channel, error) {
	var channel delivery.Channel

	if err := c.engine.Get(context.Background(), fmt.Sprintf(channelPath, id), &channel); err != nil {
		return nil, err
	}

	return &channel, nil
}

func (c Channel) Create(ctx context.Context, input *delivery.ChannelCreateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBody(
		dapp.EventBodyBasic, dapp.JSONDataType, delivery.ChannelCreate, input,
	)

	return c.event.Create(ctx, &dapp.EventInput{
		Body: body, Auth: c.signer.Auth(body),
	})
}

func (c Channel) Update(ctx context.Context, input *delivery.ChannelUpdateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBody(
		dapp.EventBodyBasic, dapp.JSONDataType, delivery.ChannelUpdate, input,
	)

	return c.event.Create(ctx, &dapp.EventInput{
		Body: body, Auth: c.signer.Auth(body),
	})
}
