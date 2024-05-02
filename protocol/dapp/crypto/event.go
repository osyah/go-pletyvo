// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package crypto

import (
	"crypto/ed25519"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

func EventInputVerifier(input *dapp.EventInput) bool {
	switch input.Auth.Schema {
	case SchemaED25519:
		if len(input.Auth.PublicKey) != ed25519.PublicKeySize {
			return false
		}

		return ed25519.Verify(input.Auth.PublicKey, input.Body, input.Auth.Signature)
	default:
		return false
	}
}
