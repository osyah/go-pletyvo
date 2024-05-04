// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

const ChannelAggregate = 1

const (
	ChannelCreate = iota
	ChannelUpdate
)

var (
	ChannelCreateEventType = dapp.NewEventType(ChannelCreate, ChannelAggregate, Version, Protocol)
	ChannelUpdateEventType = dapp.NewEventType(ChannelUpdate, ChannelAggregate, Version, Protocol)
)

type Channel struct {
	ID    uuid.UUID    `json:"id"`
	Name  string       `json:"name"`
	Owner dapp.Address `json:"owner"`
}

type ChannelQuery interface {
	GetByID(context.Context, uuid.UUID) (*Channel, error)
}

type ChannelCreateInput struct {
	Name string `json:"name"`
}

type ChannelUpdateInput struct {
	ChannelCreateInput
	Channel uuid.UUID `json:"channel"`
}

type ChannelRepository interface {
	ChannelQuery
	Create(context.Context, *Channel) error
	Update(context.Context, *ChannelUpdateInput) error
}

type ChannelService interface {
	ChannelQuery
	Create(context.Context, *ChannelCreateInput) (*dapp.EventResponse, error)
	Update(context.Context, *ChannelUpdateInput) (*dapp.EventResponse, error)
}
