// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapplocal

import (
	"github.com/osyah/go-pletyvo/node/relay"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

func New(repos *dapp.Repository, relay relay.Relay) *dapp.Service {
	return &dapp.Service{Event: NewEvent(repos.Event, relay)}
}
