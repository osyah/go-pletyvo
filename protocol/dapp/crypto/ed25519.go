// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

const SchemaED25519 = 1

type ED25519 struct{ privateKey ed25519.PrivateKey }

func NewED25519(seed []byte) *ED25519 {
	return &ED25519{privateKey: ed25519.NewKeyFromSeed(seed)}
}

func GenerateED25519(r io.Reader) *ED25519 {
	if r == nil {
		r = rand.Reader
	}

	seed := make([]byte, ed25519.SeedSize)

	if _, err := io.ReadFull(r, seed); err != nil {
		panic("go-pletyvo/protocol/dapp/crypto: " + err.Error())
	}

	return NewED25519(seed)
}

func (ED25519) Schema() byte { return SchemaED25519 }

func (ed ED25519) Sign(msg []byte) []byte {
	return ed25519.Sign(ed.privateKey, msg)
}

func (ed ED25519) Public() []byte {
	publicKey := make([]byte, ed25519.PublicKeySize)
	copy(publicKey, ed.privateKey[ed25519.PublicKeySize:])

	return publicKey
}

func (ed ED25519) Address() dapp.Hash {
	return NewHash(SchemaED25519, ed.Public())
}

func (ed ED25519) Auth(msg []byte) dapp.AuthHeader {
	return dapp.AuthHeader{
		Schema:    SchemaED25519,
		PublicKey: ed.Public(),
		Signature: ed.Sign(msg),
	}
}
