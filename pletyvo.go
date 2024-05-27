// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pletyvo

import (
	"github.com/google/uuid"

	"github.com/osyah/hryzun/status"
)

const (
	CodeNil status.Code = iota
	CodeInternal
	CodeNotFound
	CodeInvalidArgument
)

type QueryOption struct {
	Limit  int
	Order  bool
	After  uuid.UUID
	Before uuid.UUID
}
