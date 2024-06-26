// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

const (
	Protocol = 2
	Version  = 0
)

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
		Message: NewMessageExecutor(repos.Message),
	}
}
