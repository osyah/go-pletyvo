// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

const (
	Protocol = 2
	Version  = 0
)

type Repository struct {
	Channel ChannelRepository
	Message MessageRepository
}

type Service struct {
	Channel ChannelService
	Message MessageService
}
