// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registryhttp

import (
	"context"

	"github.com/osyah/go-pletyvo/client/engine"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/dapp/crypto"
	"github.com/osyah/go-pletyvo/protocol/registry"
)

type Network struct {
	engine engine.HTTP
	signer crypto.Signer
	event  dapp.EventService
}

func NewNetwork(engine engine.HTTP, signer crypto.Signer, event dapp.EventService) *Network {
	return &Network{engine: engine, signer: signer, event: event}
}

func (n Network) Get(ctx context.Context) (*registry.Network, error) {
	var network registry.Network

	if err := n.engine.Get(ctx, "/registry/v1/network", &network); err != nil {
		return nil, err
	}

	return &network, nil
}

func (n Network) Create(ctx context.Context, input *registry.NetworkCreateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBodyJSON(input, registry.NetworkCreateEventType)

	return n.event.Create(ctx, &dapp.EventInput{
		Body: body,
		Auth: n.signer.Auth(body),
	})
}

func (n Network) Update(ctx context.Context, input *registry.NetworkUpdateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBodyJSON(input, registry.NetworkUpdateEventType)

	return n.event.Create(ctx, &dapp.EventInput{
		Body: body,
		Auth: n.signer.Auth(body),
	})
}
