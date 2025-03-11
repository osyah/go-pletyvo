// Copyright (c) 2024-2025 Osyah
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

func NewNetwork(zone uint16, payload uint32) Network {
	var network Network

	network[0] = byte(zone >> 8)
	network[1] = byte(zone)

	network[2] = byte(payload >> 24)
	network[3] = byte(payload >> 16)
	network[4] = byte(payload >> 8)
	network[5] = byte(payload)

	return network
}

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
