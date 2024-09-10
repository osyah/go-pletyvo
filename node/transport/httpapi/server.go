// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package httpapi

import "github.com/osyah/go-pletyvo/node/pkg/net/http"

func (b Builder) Build() *http.Server {
	server := http.NewServer(b.config.Server)

	server.Register(b.register)

	return server
}
