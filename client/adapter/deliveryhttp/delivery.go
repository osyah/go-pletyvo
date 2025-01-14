// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package deliveryhttp

import (
	"github.com/osyah/go-pletyvo/client/engine"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/dapp/crypto"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

func New(engine engine.HTTP, signer crypto.Signer, event dapp.EventService) *delivery.Service {
	return &delivery.Service{
		Channel: NewChannel(engine, signer, event),
		Message: NewMessage(engine),
		Post:    NewPost(engine, signer, event),
	}
}
