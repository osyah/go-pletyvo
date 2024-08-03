// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapp

import "context"

const Protocol = 0

type Handler interface {
	Handle(context.Context, *Event) error
}

type AuthHeader struct {
	Schema    byte   `json:"sch"`
	PublicKey []byte `json:"pub"`
	Signature []byte `json:"sig"`
}

type Query struct{ Event EventQuery }

type Repository struct{ Event EventRepository }

type Service struct{ Event EventService }
