// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"strings"

	"github.com/osyah/hryzun"

	"github.com/osyah/go-pletyvo"
)

var ErrEmptyContent = hryzun.NewStatus(
	pletyvo.CodeInvalidArgument, "empty content",
)

func PrepareContent(s string) string {
	return strings.Trim(s, " \n\r\t")
}
