// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapppebble

import (
	"github.com/VictoriaMetrics/easyproto"
	"github.com/cockroachdb/pebble"
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

var mp easyproto.MarshalerPool

func New(db *pebble.DB) *dapp.Repository {
	return &dapp.Repository{Event: NewEvent(db)}
}

func Register() container.Handler {
	return func(base *container.Base) any {
		return New(
			container.Get[*pebble.DB](base, "pebble"),
		)
	}
}
