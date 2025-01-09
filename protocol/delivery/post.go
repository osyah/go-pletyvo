// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"context"

	"github.com/google/uuid"
	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

const (
	PostCreateEventType = 5
	PostUpdateEventType = 6
)

type Post struct {
	ID      uuid.UUID `json:"id"`
	Hash    dapp.Hash `json:"hash"`
	Author  dapp.Hash `json:"author"`
	Channel uuid.UUID `json:"channel"`
	Content string    `json:"content"`
}

type PostInput struct {
	Channel dapp.Hash `json:"channel"`
	Content string    `json:"content"`
}

func (pi PostInput) Validate() error {
	if len(pi.Content) > 2048 {
		return status.New(pletyvo.CodeInvalidArgument, "invalid content length")
	}

	return nil
}

type PostQuery interface {
	Get(context.Context, uuid.UUID, *pletyvo.QueryOption) ([]*Post, error)
	GetByID(ctx context.Context, channel uuid.UUID, id uuid.UUID) (*Post, error)
}

type PostCreateInput struct{ *PostInput }

func (pci PostCreateInput) Validate() error {
	return pci.PostInput.Validate()
}

type PostUpdateInput struct {
	*PostInput
	Post dapp.Hash `json:"post"`
}

func (pui PostUpdateInput) Validate() error {
	return pui.PostInput.Validate()
}

type PostRepository interface {
	PostQuery
	Create(context.Context, *dapp.SystemEvent, *PostInput) error
	Update(context.Context, *dapp.SystemEvent, *PostUpdateInput) error
}

type PostService interface {
	PostQuery
	Create(context.Context, *PostCreateInput) (*dapp.EventResponse, error)
	Update(context.Context, *PostUpdateInput) (*dapp.EventResponse, error)
}
