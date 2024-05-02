// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pletyvo

import "github.com/osyah/hryzun/status"

const (
	CodeNil status.Code = iota
	CodeInvalidArgument
)

type QueryOption struct {
	Limit  int
	Order  bool
	After  []byte
	Before []byte
}
