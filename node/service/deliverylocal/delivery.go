// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package deliverylocal

import (
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Config struct {
	Repos string `cfg:"repos"`
}

func New(repos *delivery.Repository) *delivery.Service {
	return &delivery.Service{
		Channel: NewChannel(repos.Channel),
		Message: NewMessage(repos.Message),
		Post:    NewPost(repos.Post),
	}
}

func Register(config Config) container.Handler {
	return func(base *container.Base) any {
		return New(container.Get[*delivery.Repository](base, config.Repos))
	}
}
