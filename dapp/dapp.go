// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package dapp

import "github.com/osyah/go-pletyvo"

type AuthHeader struct {
	Schema    byte   `json:"sch"`
	PublicKey []byte `json:"pub"`
	Signature []byte `json:"sig"`
}

type Repository struct {
	Event EventRepository
	Hash  HashRepository
}

type Service struct {
	Event EventService
	Hash  HashService
}

func NewClient(engine pletyvo.DefaultEngine) *Service {
	return &Service{Event: NewEventClient(engine), Hash: NewHashClient(engine)}
}
