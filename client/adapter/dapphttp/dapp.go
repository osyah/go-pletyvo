// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapphttp

import (
	"github.com/osyah/go-pletyvo/client/engine"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

func New(engine engine.HTTP) *dapp.Service {
	return &dapp.Service{Event: NewEvent(engine), Hash: NewHash(engine)}
}
