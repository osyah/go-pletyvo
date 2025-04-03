// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"errors"
	"strings"
)

var ErrEmptyContent = errors.New("empty content")

func PrepareContent(s string) string {
	return strings.Trim(s, " \n\r\t")
}
