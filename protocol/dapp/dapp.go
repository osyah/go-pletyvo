// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapp

type AuthHeader struct {
	Schema    byte   `json:"sch"`
	PublicKey []byte `json:"pub"`
	Signature []byte `json:"sig"`
}

type Query struct {
	Event EventQuery
	Hash  HashQuery
}

type Repository struct {
	Event EventRepository
	Hash  HashRepository
}

type Service struct {
	Event EventService
	Hash  HashService
}
