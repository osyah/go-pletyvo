// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package dapp

import (
	"context"

	"github.com/osyah/go-pletyvo"
)

type HashClient struct{ engine pletyvo.DefaultEngine }

func NewHashClient(engine pletyvo.DefaultEngine) *HashClient {
	return &HashClient{engine: engine}
}

func (hc HashClient) GetByID(ctx context.Context, id Hash) (*EventResponse, error) {
	var response EventResponse

	if err := hc.engine.Get(ctx, ("/dapp/v1/hash/" + id.String()), &response); err != nil {
		return nil, err
	}

	return &response, nil
}
