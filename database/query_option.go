// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package database

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type QueryOption struct {
	Limit  int
	Order  bool
	After  []byte
	Before []byte
}

func (qo QueryOption) String() string {
	var (
		buf   strings.Builder
		token byte = '?'
	)

	if qo.Limit != 0 {
		buf.WriteByte(token)
		buf.WriteString(fmt.Sprintf("limit=%d", qo.Limit))

		token = '&'
	}

	if qo.Order {
		buf.WriteByte(token)
		buf.WriteString("order=asc")

		token = '&'
	}

	if qo.After != nil {
		buf.WriteByte(token)
		buf.WriteString("after=")
		buf.WriteString(
			base64.RawURLEncoding.EncodeToString(qo.After),
		)

		token = '&'
	}

	if qo.Before != nil {
		buf.WriteByte(token)
		buf.WriteString("before=")
		buf.WriteString(
			base64.RawURLEncoding.EncodeToString(qo.Before),
		)

		token = '&'
	}

	return buf.String()
}
