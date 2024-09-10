// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package httpapi

import (
	"strings"

	"github.com/osyah/go-pletyvo/node/pkg/net/http"
)

type Config struct {
	Server http.Config `cfg:"server"`
	CORS   CORS        `cfg:"cors"`
}

type CORS struct {
	Origins     []string `cfg:"origins"`
	Methods     []string `cfg:"methods"`
	Headers     []string `cfg:"headers"`
	Credentials bool     `cfg:"credentials"`
}

func (c CORS) AllowOrigins() string {
	return strings.Join(c.Origins, ",")
}

func (c CORS) AllowMethods() string {
	return strings.Join(c.Methods, ",")
}

func (c CORS) AllowHeaders() string {
	return strings.Join(c.Headers, ",")
}
