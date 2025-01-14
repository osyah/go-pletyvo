// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package deliverylocal

import (
	"context"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Channel struct{ query delivery.ChannelQuery }

func NewChannel(query delivery.ChannelQuery) *Channel {
	return &Channel{query: query}
}

func (c Channel) GetByID(ctx context.Context, id uuid.UUID) (*delivery.Channel, error) {
	return c.query.GetByID(ctx, id)
}

func (c Channel) Create(context.Context, *delivery.ChannelCreateInput) (*dapp.EventResponse, error) {
	return nil, pletyvo.CodeNotImplemented
}

func (c Channel) Update(context.Context, *delivery.ChannelUpdateInput) (*dapp.EventResponse, error) {
	return nil, pletyvo.CodeNotImplemented
}
