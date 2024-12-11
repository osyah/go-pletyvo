// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapp

import (
	"context"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
)

type Event struct {
	*EventHeader
	*EventInput
}

type EventHeader struct {
	ID   uuid.UUID `json:"id"`
	Hash Hash      `json:"hash"`
}

type EventQuery interface {
	Get(context.Context, *pletyvo.QueryOption) ([]*Event, error)
	GetByID(context.Context, uuid.UUID) (*Event, error)
}

type EventRepository interface {
	EventQuery
	Create(context.Context, *Event) error
}

type EventResponse struct {
	ID uuid.UUID `json:"id"`
}

type EventService interface {
	EventQuery
	Create(context.Context, *EventInput) (*EventResponse, error)
}

type EventInput struct {
	Body EventBody  `json:"body"`
	Auth AuthHeader `json:"auth"`
}

type EventInputVerifierFunc func(*EventInput) bool

func (ei *EventInput) Verify(f EventInputVerifierFunc) bool { return f(ei) }
