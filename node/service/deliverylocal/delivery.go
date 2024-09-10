// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package deliverylocal

import (
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Config struct {
	Repos string `cfg:"repos"`
}

func NewQuery(repos *delivery.Repository) *delivery.Query {
	return &delivery.Query{
		Channel: NewChannel(repos.Channel),
		Message: NewMessage(repos.Message),
	}
}

func RegisterQuery(config Config) container.Handler {
	return func(base *container.Base) any {
		return NewQuery(
			container.Get[*delivery.Repository](base, config.Repos),
		)
	}
}
