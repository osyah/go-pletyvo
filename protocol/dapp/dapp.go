// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapp

type AuthHeader struct {
	Schema    byte   `json:"sch"`
	PublicKey []byte `json:"pub"`
	Signature []byte `json:"sig"`
}

type Repository struct{ Event EventRepository }

type Service struct{ Event EventService }
