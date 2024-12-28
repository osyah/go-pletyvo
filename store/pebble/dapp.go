// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"github.com/osyah/go-pletyvo/protocol/dapp"

	"github.com/cockroachdb/pebble"
)

const (
	DAppEventPrefix = 1
	DAppHashPrefix  = 2
)

func NewDApp(db *pebble.DB) *dapp.Repository {
	return &dapp.Repository{
		Event: NewDAppEvent(db),
		Hash:  NewDAppHash(db),
	}
}
