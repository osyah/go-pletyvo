// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registry

import (
	"context"

	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
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
	*dapp.EventHeader
	*NetworkInput
}

type NetworkInput struct {
	Name string `json:"name"`
}

func (ni NetworkInput) Validate() error {
	if len(ni.Name) > 40 {
		return status.New(pletyvo.CodeInvalidArgument, "invalid name length")
	}

	return nil
}

type NetworkQuery interface {
	Get(context.Context) (*Network, error)
}

type NetworkCreateInput struct{ *NetworkInput }

func (nci NetworkCreateInput) Validate() error {
	return nci.NetworkInput.Validate()
}

type NetworkUpdateInput struct{ *NetworkInput }

func (nui NetworkUpdateInput) Validate() error {
	return nui.NetworkInput.Validate()
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
