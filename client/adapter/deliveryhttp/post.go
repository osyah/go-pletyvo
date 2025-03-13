// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package deliveryhttp

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/client/engine"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/dapp/crypto"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

const (
	postsPath = channelPath + "/posts"
	postPath  = postsPath + "/%s"
)

type Post struct {
	engine engine.HTTP
	signer crypto.Signer
	event  dapp.EventService
}

func NewPost(engine engine.HTTP, signer crypto.Signer, event dapp.EventService) *Post {
	return &Post{engine: engine, signer: signer, event: event}
}

func (p Post) Get(ctx context.Context, ch uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Post, error) {
	posts := make([]*delivery.Post, option.Limit)

	if err := p.engine.Get(ctx, (fmt.Sprintf(postsPath, ch) + option.String()), &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p Post) GetByID(ctx context.Context, ch uuid.UUID, id uuid.UUID) (*delivery.Post, error) {
	var post delivery.Post

	if err := p.engine.Get(ctx, fmt.Sprintf(postPath, ch, id), &post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (p Post) Create(ctx context.Context, input *delivery.PostCreateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBody(
		dapp.EventBodyBasic, dapp.JSONDataType, delivery.PostCreate, input,
	)

	return p.event.Create(ctx, &dapp.EventInput{
		Body: body, Auth: p.signer.Auth(body),
	})
}

func (p Post) Update(ctx context.Context, input *delivery.PostUpdateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBody(
		dapp.EventBodyBasic, dapp.JSONDataType, delivery.PostUpdate, input,
	)

	return p.event.Create(ctx, &dapp.EventInput{
		Body: body, Auth: p.signer.Auth(body),
	})
}
