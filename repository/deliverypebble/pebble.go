// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package deliverypebble

import (
	"github.com/VictoriaMetrics/easyproto"
	"github.com/cockroachdb/pebble"
	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/protocol/delivery"
)

var mp easyproto.MarshalerPool

func New(db *pebble.DB) *delivery.Repository {
	return &delivery.Repository{
		Channel: NewChannel(db),
		Message: NewMessage(db),
	}
}

func Register() container.Handler {
	return func(base *container.Base) any {
		return New(
			container.Get[*pebble.DB](base, "pebble"),
		)
	}
}
