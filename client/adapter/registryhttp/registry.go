// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registryhttp

import (
	"github.com/osyah/go-pletyvo/client/engine"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/dapp/crypto"
	"github.com/osyah/go-pletyvo/protocol/registry"
)

func New(engine engine.HTTP, signer crypto.Signer, event dapp.EventService) *registry.Service {
	return &registry.Service{Network: NewNetwork(engine, signer, event)}
}
