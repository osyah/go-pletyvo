// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package config

import "github.com/osyah/hryzun/config"

type Root struct {
	Node Node `cfg:"node"`
}

func New(filename string) (*Root, error) {
	return config.New[Root](filename)
}
