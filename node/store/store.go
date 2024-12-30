// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package store

import (
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/store/pebble"
)

type Config struct {
	Pebble *pebble.Config `cfg:"pebble"`
}

func Register(base *container.Base, config Config) {
	if config.Pebble == nil {
		panic("go-pletyvo/node/store: 'pebble' config not found")
	} else {
		pebble.Register(base, *config.Pebble)
	}
}
