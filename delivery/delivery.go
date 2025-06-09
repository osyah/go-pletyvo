// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"github.com/istyna/go-pletyvo"
	"github.com/istyna/go-pletyvo/dapp"
)

type Repository struct {
	Channel ChannelRepository
	Message MessageRepository
	Post    PostRepository
}

type Service struct {
	Channel ChannelService
	Message MessageService
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

func NewClient(engine pletyvo.DefaultEngine, signer dapp.Signer, event dapp.EventService) *Service {
	return &Service{
		Channel: NewChannelClient(engine, signer, event),
		Message: NewMessageClient(engine),
		Post:    NewPostClient(engine, signer, event),
	}
}
