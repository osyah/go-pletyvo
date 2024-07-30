// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapplocal

import (
	"context"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/node/relay"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/dapp/crypto"
)

type Event struct {
	query dapp.EventQuery
	relay relay.Relay
}

func NewEvent(query dapp.EventQuery, relay relay.Relay) *Event {
	return &Event{query: query, relay: relay}
}

func (e Event) Get(ctx context.Context, option *pletyvo.QueryOption) ([]*dapp.Event, error) {
	option.Prepare()

	return e.query.Get(ctx, option)
}

func (e Event) GetByID(ctx context.Context, id uuid.UUID) (*dapp.Event, error) {
	return e.query.GetByID(ctx, id)
}

func (e Event) Create(ctx context.Context, input *dapp.EventInput) (*dapp.EventResponse, error) {
	var (
		header dapp.EventHeader
		err    error
	)

	header.ID, err = uuid.NewV7()
	if err != nil {
		return nil, pletyvo.CodeInternal
	}

	if !input.Verify(crypto.EventInputVerifier) {
		return nil, pletyvo.CodeUnauthorized
	}

	header.Author = crypto.NewAddress(input.Auth.Schema, input.Auth.PublicKey)

	err = e.relay.OnEvent(ctx, &dapp.Event{
		EventHeader: &header,
		EventInput:  input,
	})
	if err != nil {
		return nil, err
	}

	return &dapp.EventResponse{ID: header.ID}, nil
}
