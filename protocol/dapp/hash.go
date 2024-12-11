// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapp

import (
	"context"
	"encoding/base64"

	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
)

const (
	HashSize = 32
	HashLen  = 43
)

var (
	HashNil Hash

	ErrInvalidHashSize = status.New(
		pletyvo.CodeInvalidArgument, "invalid hash size",
	)
	ErrInvalidHashLen = status.New(
		pletyvo.CodeInvalidArgument, "invalid hash length",
	)
)

type HashQuery interface {
	GetByID(context.Context, Hash) (*EventResponse, error)
}

type HashRepository interface {
	HashQuery
	Create(context.Context, *EventHeader) error
}

type HashService interface{ HashQuery }

type Hash [HashSize]byte

func (h Hash) String() string {
	return base64.RawURLEncoding.EncodeToString(h[:])
}

func (h Hash) MarshalJSON() ([]byte, error) {
	b := make([]byte, (HashLen + 2))

	b[0], b[44] = '"', '"'
	base64.RawURLEncoding.Encode(b[1:44], h[:])

	return b, nil
}

func (h *Hash) UnmarshalJSON(b []byte) error {
	if len(b) != (HashLen + 2) {
		return ErrInvalidHashLen
	}

	return DecodeHash(h[:], b[1:44])
}

func HashFromString(s string) (hash Hash, err error) {
	if len(s) != HashLen {
		err = ErrInvalidHashLen
		return
	}

	err = DecodeHash(hash[:], []byte(s))

	return
}

func DecodeHash(dst []byte, src []byte) error {
	n, err := base64.RawURLEncoding.Decode(dst, src)
	if err != nil {
		return err
	}

	if n != HashSize {
		return ErrInvalidHashSize
	}

	return nil
}
