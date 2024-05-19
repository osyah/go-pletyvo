// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapppebble

import (
	"github.com/VictoriaMetrics/easyproto"
	"github.com/cockroachdb/pebble"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

var mp easyproto.MarshalerPool

func New(db *pebble.DB) *dapp.Repository {
	return &dapp.Repository{Event: NewEvent(db)}
}
