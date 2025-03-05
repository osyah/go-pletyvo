// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package crypto

import (
	"github.com/zeebo/blake3"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Signer interface {
	Schema() byte
	Sign([]byte) []byte
	Public() []byte
	Address() dapp.Hash
	Auth([]byte) dapp.AuthHeader
}

func NewHash(schema byte, data []byte) dapp.Hash {
	tmp := make([]byte, (len(data) + 1))

	tmp[0] = schema
	copy(tmp[1:], data)

	return blake3.Sum256(tmp)
}
