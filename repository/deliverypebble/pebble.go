// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package deliverypebble

import (
	"github.com/VictoriaMetrics/easyproto"
	"github.com/cockroachdb/pebble"

	"github.com/osyah/go-pletyvo/protocol/delivery"
)

var mp easyproto.MarshalerPool

func New(db *pebble.DB) *delivery.Repository {
	return &delivery.Repository{
		Channel: NewChannel(db),
		Message: NewMessage(db),
	}
}
