// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import "github.com/osyah/go-pletyvo/protocol/dapp"

type Query struct {
	Channel ChannelQuery
	Post    PostQuery
}

type Repository struct {
	Channel ChannelRepository
	Post    PostRepository
}

type Service struct {
	Channel ChannelService
	Post    PostService
}

type Executor struct {
	Channel *ChannelExecutor
	Post    *PostExecutor
}

func NewExecutor(repos *Repository) *Executor {
	return &Executor{
		Channel: NewChannelExecutor(repos.Channel),
		Post:    NewPostExecutor(repos.Post, repos.Channel),
	}
}

func (e Executor) Register(handler *dapp.Handler) {
	e.Channel.Register(handler)
	e.Post.Register(handler)
}
