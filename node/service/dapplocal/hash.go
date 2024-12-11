// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapplocal

import (
	"context"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Hash struct{ query dapp.HashQuery }

func NewHash(query dapp.HashQuery) *Hash {
	return &Hash{query: query}
}

func (h Hash) GetByID(ctx context.Context, id dapp.Hash) (*dapp.EventResponse, error) {
	return h.query.GetByID(ctx, id)
}
