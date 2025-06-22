// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package database

import "net/http"

const TokenHeader = "Token"

type (
	DecoderFunc[T any] func([]byte) *T
	EncoderFunc[T any] func(*T) []byte

	SpaceConfig[T any] struct {
		URL, Token string

		Decoder DecoderFunc[T]
		Encoder EncoderFunc[T]
	}
)

type Space[T any] struct {
	*SpaceConfig[T]
	client *http.Client
}

func NewSpace[T any](config *SpaceConfig[T]) *Space[T] {
	if config.Decoder == nil {
		config.Decoder = DecoderFuncJSON[T]()
	}
	if config.Encoder == nil {
		config.Encoder = EncoderFuncJSON[T]()
	}

	return &Space[T]{SpaceConfig: config, client: http.DefaultClient}
}
