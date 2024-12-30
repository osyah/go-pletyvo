// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package node

import (
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/node/config"
	"github.com/osyah/go-pletyvo/node/relay"
	"github.com/osyah/go-pletyvo/node/service"
	"github.com/osyah/go-pletyvo/node/store"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

func InitContainer(base *container.Base, config config.Node) {
	if base == nil {
		base = container.New()
	}

	store.Register(base, config.Store)
	relay.Register(base, config.Relay)
	service.Register(base, config.Service)

	base.RegisterHandler("handler", func(base *container.Base) any {
		handler := dapp.NewHandler()

		delivery.NewExecutor(
			container.Get[*delivery.Repository](base, config.Protocol.Delivery.Repos),
		).Register(handler)

		return handler
	})
}
