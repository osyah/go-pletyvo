// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package service

import (
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/node/service/dapplocal"
	"github.com/osyah/go-pletyvo/node/service/deliverylocal"
	"github.com/osyah/go-pletyvo/node/service/registrylocal"
)

func Register(base *container.Base, config Config) {
	if config.DAppLocal == nil {
		panic("go-pletyvo/node/service: 'dapp_local' config not found")
	} else {
		base.RegisterHandler("dapp_local", dapplocal.Register(*config.DAppLocal))
	}

	if config.RegistryLocal == nil {
		panic("go-pletyvo/node/service: 'registry_local' config not found")
	} else {
		base.RegisterHandler("registry_local", registrylocal.RegisterQuery(*config.RegistryLocal))
	}

	if config.DeliveryLocal == nil {
		panic("go-pletyvo/node/service: 'delivery_local' config not found")
	} else {
		base.RegisterHandler("delivery_local", deliverylocal.RegisterQuery(*config.DeliveryLocal))
	}
}
