// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"github.com/VictoriaMetrics/easyproto"
	"github.com/cockroachdb/pebble"
	"github.com/osyah/hryzun/container"
)

var mp easyproto.MarshalerPool

type Config struct {
	Dirname string `cfg:"dirname"`
}

func New(config Config) (*pebble.DB, error) {
	return pebble.Open(config.Dirname, &pebble.Options{})
}

func Register(base *container.Base, config Config) {
	base.RegisterHandler("pebble", func(base *container.Base) any {
		db, err := New(config)
		if err != nil {
			panic("go-pletyvo/store/pebble: " + err.Error())
		}

		base.RegisterCloser("pebble", db.Close)

		return db
	})

	base.RegisterHandler("dapp_pebble", func(base *container.Base) any {
		return NewDApp(container.Get[*pebble.DB](base, "pebble"))
	})
	base.RegisterHandler("delivery_pebble", func(base *container.Base) any {
		return NewDelivery(container.Get[*pebble.DB](base, "pebble"))
	})
}
