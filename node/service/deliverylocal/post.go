// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package deliverylocal

import (
	"context"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Post struct{ query delivery.PostQuery }

func NewPost(query delivery.PostQuery) *Post {
	return &Post{query: query}
}

func (p Post) Get(ctx context.Context, id uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Post, error) {
	option.Prepare()

	return p.query.Get(ctx, id, option)
}

func (p Post) GetByID(ctx context.Context, channel, id uuid.UUID) (*delivery.Post, error) {
	return p.query.GetByID(ctx, channel, id)
}

func (p Post) Create(context.Context, *delivery.PostCreateInput) (*dapp.EventResponse, error) {
	return nil, pletyvo.CodeNotImplemented
}

func (p Post) Update(context.Context, *delivery.PostUpdateInput) (*dapp.EventResponse, error) {
	return nil, pletyvo.CodeNotImplemented
}
