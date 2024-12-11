// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapphttp

import (
	"context"

	"github.com/osyah/go-pletyvo/client/engine"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Hash struct{ engine engine.HTTP }

func NewHash(engine engine.HTTP) *Hash {
	return &Hash{engine: engine}
}

func (h Hash) GetByID(ctx context.Context, id dapp.Hash) (*dapp.EventResponse, error) {
	var response dapp.EventResponse

	if err := h.engine.Get(ctx, ("/dapp/v1/hash/" + id.String()), &response); err != nil {
		return nil, err
	}

	return &response, nil
}
