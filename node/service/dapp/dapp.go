// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapp

import (
	"github.com/osyah/go-pletyvo/node/relay"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

func New(query *dapp.Query, relay relay.Relay) *dapp.Service {
	return &dapp.Service{Event: NewEvent(query.Event, relay)}
}
