// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"github.com/cockroachdb/pebble"
	"github.com/osyah/hryzun/container"
)

type Config struct {
	Dirname string `cfg:"dirname"`
}

func New(config Config) (*pebble.DB, error) {
	return pebble.Open(config.Dirname, &pebble.Options{})
}

func Register(config Config) container.Handler {
	return func(base *container.Base) any {
		db, err := New(config)
		if err != nil {
			panic("go-pletyvo/store/pebble: " + err.Error())
		}

		base.RegisterCloser("pebble", db.Close)

		return db
	}
}
