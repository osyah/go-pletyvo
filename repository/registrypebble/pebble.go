// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registrypebble

import (
	"github.com/VictoriaMetrics/easyproto"
	"github.com/cockroachdb/pebble"

	"github.com/osyah/go-pletyvo/protocol/registry"
)

var mp easyproto.MarshalerPool

func New(db *pebble.DB) *registry.Repository {
	return &registry.Repository{Network: NewNetwork(db)}
}
