// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package dapp

import (
	"context"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
)

type EventClient struct{ engine pletyvo.DefaultEngine }

func NewEventClient(engine pletyvo.DefaultEngine) *EventClient {
	return &EventClient{engine: engine}
}

func (ec EventClient) Get(ctx context.Context, option *pletyvo.QueryOption) ([]*Event, error) {
	events := make([]*Event, option.Limit)

	if err := ec.engine.Get(ctx, ("/dapp/v1/events" + option.String()), &events); err != nil {
		return nil, err
	}

	return events, nil
}

func (ec EventClient) GetByID(ctx context.Context, id uuid.UUID) (*Event, error) {
	var event Event

	if err := ec.engine.Get(ctx, ("/dapp/v1/events/" + id.String()), &event); err != nil {
		return nil, err
	}

	return &event, nil
}

func (ec EventClient) Create(ctx context.Context, input *EventInput) (*EventResponse, error) {
	var response EventResponse

	if err := ec.engine.Post(ctx, "/dapp/v1/events", input, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
