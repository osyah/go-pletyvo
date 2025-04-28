// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package dapp

import "github.com/zeebo/blake3"

type Signer interface {
	Schema() byte
	Sign([]byte) []byte
	Public() []byte
	Address() Hash
	Auth([]byte) AuthHeader
}

func NewHash(schema byte, data []byte) Hash {
	tmp := make([]byte, (len(data) + 1))

	tmp[0] = schema
	copy(tmp[1:], data)

	return blake3.Sum256(tmp)
}
