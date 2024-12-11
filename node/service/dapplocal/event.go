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
	hash  dapp.HashQuery
	relay relay.Relay
}

func NewEvent(repos *dapp.Repository, relay relay.Relay) *Event {
	return &Event{query: repos.Event, hash: repos.Hash, relay: relay}
}

func (e Event) Get(ctx context.Context, option *pletyvo.QueryOption) ([]*dapp.Event, error) {
	option.Prepare()

	return e.query.Get(ctx, option)
}

func (e Event) GetByID(ctx context.Context, id uuid.UUID) (*dapp.Event, error) {
	return e.query.GetByID(ctx, id)
}

func (e Event) Create(ctx context.Context, input *dapp.EventInput) (*dapp.EventResponse, error) {
	header := &dapp.EventHeader{
		Hash: crypto.NewHash(input.Auth.Schema, input.Auth.Signature),
	}

	response, err := e.hash.GetByID(ctx, header.Hash)
	if err == nil {
		return response, nil
	}

	if !input.Verify(crypto.EventInputVerifier) {
		return nil, pletyvo.CodeUnauthorized
	}

	header.ID, err = uuid.NewV7()
	if err != nil {
		return nil, pletyvo.CodeInternal
	}

	event := &dapp.Event{EventHeader: header, EventInput: input}

	if err = e.relay.OnEvent(ctx, event); err != nil {
		return nil, pletyvo.CodeInternal
	}

	return &dapp.EventResponse{ID: event.ID}, nil
}
