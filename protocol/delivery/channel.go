// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/google/uuid"
	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

const (
	ChannelCreateEventType = 3
	ChannelUpdateEventType = 4
)

type Channel struct {
	ID     uuid.UUID `json:"id"`
	Hash   dapp.Hash `json:"hash"`
	Author dapp.Hash `json:"author"`
	Name   string    `json:"name"`
}

type ChannelInput struct {
	Name string `json:"name"`
}

func (ci ChannelInput) Validate() error {
	if len(ci.Name) > 40 || len(ci.Name) == 0 {
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
	Channel dapp.Hash `json:"channel"`
}

func (cui ChannelUpdateInput) Validate() error {
	return cui.ChannelInput.Validate()
}

type ChannelRepository interface {
	ChannelQuery
	GetByHash(context.Context, dapp.Hash) (*Channel, error)
	Create(context.Context, *dapp.SystemEvent, *ChannelInput) error
	Update(context.Context, *dapp.SystemEvent, *ChannelUpdateInput) error
}

type ChannelService interface {
	ChannelQuery
	Create(context.Context, *ChannelCreateInput) (*dapp.EventResponse, error)
	Update(context.Context, *ChannelUpdateInput) (*dapp.EventResponse, error)
}
