// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import "github.com/osyah/go-pletyvo/protocol/dapp"

type Query struct {
	Channel ChannelQuery
	Message MessageQuery
}

type Repository struct {
	Channel ChannelRepository
	Message MessageRepository
}

type Service struct {
	Channel ChannelService
	Message MessageService
}

type Executor struct {
	Channel *ChannelExecutor
	Message *MessageExecutor
}

func NewExecutor(repos *Repository) *Executor {
	return &Executor{
		Channel: NewChannelExecutor(repos.Channel),
		Message: NewMessageExecutor(repos.Message, repos.Channel),
	}
}

func (e Executor) Register(handler *dapp.Handler) {
	e.Channel.Register(handler)
	e.Message.Register(handler)
}
