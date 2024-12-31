// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pletyvo

import (
	"encoding/base64"

	"github.com/osyah/hryzun/status"
)

const (
	NetworkSize = 6
	NetworkLen  = 8
)

var (
	DefaultNetwork, NetworkNil Network

	ErrInvalidNetworkFormat = status.New(
		CodeInvalidArgument, "invalid network format",
	)
)

type Network [NetworkSize]byte

func (n Network) String() string {
	return base64.RawURLEncoding.EncodeToString(n[:])
}

func NetworkFromString(s string) (Network, error) {
	var network Network

	if len(s) != NetworkLen {
		return NetworkNil, ErrInvalidNetworkFormat
	}

	return network, DecodeNetwork(network[:], []byte(s))
}

func DecodeNetwork(dst []byte, src []byte) error {
	n, err := base64.RawURLEncoding.Decode(dst, src)
	if err != nil {
		return err
	}

	if n != NetworkSize {
		return ErrInvalidNetworkFormat
	}

	return nil
}
