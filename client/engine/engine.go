// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package engine

import "context"

type HTTP interface {
	Get(ctx context.Context, endpoint string, value any) error
	Post(ctx context.Context, endpoint string, body, value any) error
}
