// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/dapp"
)

const channelPath = "/delivery/v1/channel/%s"

type ChannelClient struct {
	engine pletyvo.DefaultEngine
	signer dapp.Signer
	event  dapp.EventService
}

func NewChannelClient(engine pletyvo.DefaultEngine, signer dapp.Signer, event dapp.EventService) *ChannelClient {
	return &ChannelClient{engine: engine, signer: signer, event: event}
}

func (cc ChannelClient) GetByID(ctx context.Context, id uuid.UUID) (*Channel, error) {
	var channel Channel

	if err := cc.engine.Get(context.Background(), fmt.Sprintf(channelPath, id), &channel); err != nil {
		return nil, err
	}

	return &channel, nil
}

func (cc ChannelClient) Create(ctx context.Context, input *ChannelCreateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBody(
		dapp.EventBodyBasic, dapp.JSONDataType, ChannelCreate, input,
	)

	return cc.event.Create(ctx, &dapp.EventInput{
		Body: body, Auth: cc.signer.Auth(body),
	})
}

func (cc ChannelClient) Update(ctx context.Context, input *ChannelUpdateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBody(
		dapp.EventBodyBasic, dapp.JSONDataType, ChannelUpdate, input,
	)

	return cc.event.Create(ctx, &dapp.EventInput{
		Body: body, Auth: cc.signer.Auth(body),
	})
}
