// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package dapp

import (
	"context"
	"crypto/ed25519"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
)

type Event struct {
	ID   uuid.UUID  `json:"id"`
	Body EventBody  `json:"body"`
	Auth AuthHeader `json:"auth"`
}

type SystemEvent struct {
	*EventHeader
	*EventInput
	Author Hash
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
	Create(context.Context, *SystemEvent) error
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

func EventInputVerifier(input *EventInput) bool {
	switch input.Auth.Schema {
	case SchemaED25519:
		if len(input.Auth.PublicKey) != ed25519.PublicKeySize {
			return false
		}

		return ed25519.Verify(input.Auth.PublicKey, input.Body, input.Auth.Signature)
	default:
		return false
	}
}
