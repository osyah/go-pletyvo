// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package crypto

import (
	"github.com/zeebo/blake3"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Signer interface {
	Sign([]byte) []byte
	Public() []byte
	Address() dapp.Address
}

func NewAddress(schema byte, publicKey []byte) dapp.Address {
	tmp := []byte{schema}
	tmp = append(tmp, publicKey...)

	return dapp.Address(blake3.Sum256(tmp))
}
