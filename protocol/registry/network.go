// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registry

import (
	"context"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

const NetworkAggregate = 1

const (
	NetworkCreate = iota
	NetworkUpdate
)

var (
	NetworkCreateEventType = dapp.NewEventType(NetworkCreate, NetworkAggregate, Version, Protocol)
	NetworkUpdateEventType = dapp.NewEventType(NetworkUpdate, NetworkAggregate, Version, Protocol)
)

type Network struct {
	ID    uuid.UUID    `json:"id"`
	Name  string       `json:"name"`
	Owner dapp.Address `json:"owner"`
}

type NetworkQuery interface {
	GetByID(context.Context, uuid.UUID) (*Network, error)
}

type NetworkCreateInput struct {
	Name string `json:"name"`
}

type NetworkUpdateInput struct {
	NetworkCreateInput
	Network uuid.UUID `json:"network"`
}

type NetworkRepository interface {
	NetworkQuery
	Create(context.Context, *Network) error
	Update(context.Context, *NetworkUpdateInput) error
}

type NetworkService interface {
	NetworkQuery
	Create(context.Context, *NetworkCreateInput) (*dapp.EventResponse, error)
	Update(context.Context, *NetworkUpdateInput) (*dapp.EventResponse, error)
}
