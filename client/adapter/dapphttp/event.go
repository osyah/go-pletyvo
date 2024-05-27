// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapphttp

import (
	"context"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/client/engine"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Event struct{ engine engine.HTTP }

func NewEvent(engine engine.HTTP) *Event {
	return &Event{engine: engine}
}

func (e Event) Get(ctx context.Context, opts ...*pletyvo.QueryOption) ([]*dapp.Event, error) {
	var (
		events []*dapp.Event
		option string
	)

	if opts != nil {
		events = make([]*dapp.Event, opts[0].Limit)
		option = opts[0].String()
	}

	if err := e.engine.Get(ctx, ("/dapp/v1/events" + option), &events); err != nil {
		return nil, err
	}

	return events, nil
}

func (e Event) GetByID(ctx context.Context, id uuid.UUID) (*dapp.Event, error) {
	var event dapp.Event

	if err := e.engine.Get(ctx, ("/dapp/v1/events/" + id.String()), &event); err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Create(ctx context.Context, input *dapp.EventInput) (*dapp.EventResponse, error) {
	var response dapp.EventResponse

	if err := e.engine.Post(ctx, "/dapp/v1/events", input, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
