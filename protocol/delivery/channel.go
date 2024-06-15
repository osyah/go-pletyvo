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
	*dapp.EventHeader
	*ChannelInput
}

type ChannelInput struct {
	Name string `json:"name"`
}

func (ci ChannelInput) Validate() error {
	if len(ci.Name) > 40 {
		return status.New(pletyvo.CodeInvalidArgument, "invalid name length")
	}

	return nil
}

type ChannelQuery interface {
	GetByID(context.Context, uuid.UUID) (*Channel, error)
}

type ChannelCreateInput struct{ *ChannelInput }

func (cci ChannelCreateInput) Validate() error {
	return cci.ChannelInput.Validate()
}

type ChannelUpdateInput struct {
	*ChannelInput
	Channel uuid.UUID `json:"channel"`
}

func (cui ChannelUpdateInput) Validate() error {
	return cui.ChannelInput.Validate()
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

type ChannelExecutor struct{ repos ChannelRepository }

func NewChannelExecutor(repos ChannelRepository) *ChannelExecutor {
	return &ChannelExecutor{repos: repos}
}

func (ce ChannelExecutor) Create(ctx context.Context, base dapp.EventBase[ChannelCreateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	return ce.repos.Create(ctx, &Channel{
		EventHeader:  base.EventHeader,
		ChannelInput: base.Input.ChannelInput,
	})
}

func (ce ChannelExecutor) Update(ctx context.Context, base dapp.EventBase[ChannelUpdateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	channel, err := ce.repos.GetByID(ctx, base.Input.Channel)
	if err != nil {
		return err
	}

	if channel.Author != base.Author {
		return pletyvo.CodePermissionDenied
	}

	return ce.repos.Update(ctx, base.Input)
}
