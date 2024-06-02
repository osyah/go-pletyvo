// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/google/uuid"
	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
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

func (cci ChannelCreateInput) Validate() error {
	if len(cci.Name) > 40 {
		return status.New(pletyvo.CodeInvalidArgument, "invalid name length")
	}

	return nil
}

type ChannelUpdateInput struct {
	ChannelCreateInput
	Channel uuid.UUID `json:"channel"`
}

func (cui ChannelUpdateInput) Validate() error {
	return cui.ChannelCreateInput.Validate()
}

type ChannelRepository interface {
	GetByID(ctx context.Context, ns, id uuid.UUID) (*Channel, error)
	Create(context.Context, uuid.UUID, *Channel) error
	Update(context.Context, uuid.UUID, *ChannelUpdateInput) error
}

type ChannelService interface {
	ChannelQuery
	Create(context.Context, *ChannelCreateInput) (*dapp.EventResponse, error)
	Update(context.Context, *ChannelUpdateInput) (*dapp.EventResponse, error)
}

type ChannelExecutor struct{ repos ChannelRepository }

func NewChannelExecutor(repos ChannelRepository) *ChannelExecutor {
	return &ChannelExecutor{repos: repos}
}

func (ce ChannelExecutor) Create(ctx context.Context, base dapp.EventBase[ChannelCreateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	return ce.repos.Create(ctx, base.Network, &Channel{
		ID:    base.ID,
		Name:  base.Input.Name,
		Owner: base.Address,
	})
}

func (ce ChannelExecutor) Update(ctx context.Context, base dapp.EventBase[ChannelUpdateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	channel, err := ce.repos.GetByID(ctx, base.Network, base.Input.Channel)
	if err != nil {
		return err
	}

	if channel.Owner != base.Address {
		return pletyvo.CodePermissionDenied
	}

	return ce.repos.Update(ctx, base.Network, base.Input)
}
