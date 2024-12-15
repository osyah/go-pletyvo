// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package service

import (
	"github.com/osyah/go-pletyvo/node/service/dapplocal"
	"github.com/osyah/go-pletyvo/node/service/deliverylocal"
)

type Config struct {
	DAppLocal     *dapplocal.Config     `cfg:"dapp_local"`
	DeliveryLocal *deliverylocal.Config `cfg:"delivery_local"`
}
