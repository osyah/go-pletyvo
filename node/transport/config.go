// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package transport

import "github.com/osyah/go-pletyvo/node/transport/httpapi"

type Config struct {
	HTTPAPI httpapi.Config `cfg:"http_api"`
}
