// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package database

import "encoding/json"

func DecoderFuncJSON[T any]() DecoderFunc[T] {
	return func(b []byte) *T {
		var value T

		if err := json.Unmarshal(b, &value); err != nil {
			return nil
		}

		return &value
	}
}

func EncoderFuncJSON[T any]() EncoderFunc[T] {
	return func(value *T) []byte {
		b, err := json.Marshal(value)
		if err != nil {
			return nil
		}

		return b
	}
}
