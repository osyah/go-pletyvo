// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapp

import (
	"encoding/hex"
	"encoding/json"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/hryzun/status"
)

const (
	AddressSize = 32
	AddressLen  = (AddressSize * 2) + 2
)

var (
	AddressNil Address

	ErrInvalidAddressSize = status.New(
		pletyvo.CodeInvalidArgument, "invalid address size",
	)
	ErrInvalidAddressLen = status.New(
		pletyvo.CodeInvalidArgument, "invalid address length",
	)
)

type Address [AddressSize]byte

func AddressFromString(s string) (Address, error) {
	if len(s) != AddressLen {
		return AddressNil, ErrInvalidAddressLen
	}

	var addr Address

	if err := DecodeAddress(addr[:], []byte(s[2:])); err != nil {
		return AddressNil, err
	}

	return addr, nil
}

func DecodeAddress(dst, src []byte) error {
	n, err := hex.Decode(dst[:], src[:])
	if err != nil {
		return err
	}

	if n != AddressSize {
		return ErrInvalidAddressSize
	}

	return nil
}

func (a Address) String() string {
	return "0x" + hex.EncodeToString(a[:])
}

func (a Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a *Address) UnmarshalJSON(b []byte) error {
	if len(b) != (AddressLen + 2) {
		return ErrInvalidAddressSize
	}

	return DecodeAddress(a[:], b[3:67])
}
