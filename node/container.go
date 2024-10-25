// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package node

import (
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/node/config"
	"github.com/osyah/go-pletyvo/node/relay"
	"github.com/osyah/go-pletyvo/node/service"
	"github.com/osyah/go-pletyvo/protocol"
	"github.com/osyah/go-pletyvo/protocol/delivery"
	"github.com/osyah/go-pletyvo/protocol/registry"
	"github.com/osyah/go-pletyvo/repository"
	"github.com/osyah/go-pletyvo/store"
)

func InitContainer(base *container.Base, config config.Node) {
	if base == nil {
		base = container.New()
	}

	store.Register(base, config.Store)
	repository.Register(base)
	relay.Register(base, config.Relay)
	service.Register(base, config.Service)

	base.RegisterHandler("executor", func(base *container.Base) any {
		return &protocol.Executor{
			Registry: registry.NewExecutor(
				container.Get[*registry.Repository](base, config.Protocol.Registry.Repos),
			),
			Delivery: delivery.NewExecutor(
				container.Get[*delivery.Repository](base, config.Protocol.Delivery.Repos),
			),
		}
	})
}
