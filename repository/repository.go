// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package repository

import (
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/repository/dapppebble"
	"github.com/osyah/go-pletyvo/repository/deliverypebble"
)

func Register(base *container.Base) {
	base.RegisterHandler("dapp_pebble", dapppebble.Register())
	base.RegisterHandler("delivery_pebble", deliverypebble.Register())
}
