// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/dapp"
)

const (
	postsPath = channelPath + "/posts"
	postPath  = postsPath + "/%s"
)

type PostClient struct {
	engine pletyvo.DefaultEngine
	signer dapp.Signer
	event  dapp.EventService
}

func NewPostClient(engine pletyvo.DefaultEngine, signer dapp.Signer, event dapp.EventService) *PostClient {
	return &PostClient{engine: engine, signer: signer, event: event}
}

func (pc PostClient) Get(ctx context.Context, ch uuid.UUID, option *pletyvo.QueryOption) ([]*Post, error) {
	posts := make([]*Post, option.Limit)

	if err := pc.engine.Get(ctx, (fmt.Sprintf(postsPath, ch) + option.String()), &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (pc PostClient) GetByID(ctx context.Context, ch uuid.UUID, id uuid.UUID) (*Post, error) {
	var post Post

	if err := pc.engine.Get(ctx, fmt.Sprintf(postPath, ch, id), &post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (pc PostClient) Create(ctx context.Context, input *PostCreateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBody(
		dapp.EventBodyBasic, dapp.JSONDataType, PostCreate, input,
	)

	return pc.event.Create(ctx, &dapp.EventInput{
		Body: body, Auth: pc.signer.Auth(body),
	})
}

func (pc PostClient) Update(ctx context.Context, input *PostUpdateInput) (*dapp.EventResponse, error) {
	body := dapp.NewEventBody(
		dapp.EventBodyBasic, dapp.JSONDataType, PostUpdate, input,
	)

	return pc.event.Create(ctx, &dapp.EventInput{
		Body: body, Auth: pc.signer.Auth(body),
	})
}
