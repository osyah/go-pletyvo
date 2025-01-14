// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package service

import (
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/node/service/dapplocal"
	"github.com/osyah/go-pletyvo/node/service/deliverylocal"
)

func Register(base *container.Base, config Config) {
	if config.DAppLocal == nil {
		panic("go-pletyvo/node/service: 'dapp_local' config not found")
	} else {
		base.RegisterHandler("dapp_local", dapplocal.Register(*config.DAppLocal))
	}

	if config.DeliveryLocal == nil {
		panic("go-pletyvo/node/service: 'delivery_local' config not found")
	} else {
		base.RegisterHandler("delivery_local", deliverylocal.Register(*config.DeliveryLocal))
	}
}
