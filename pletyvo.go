// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pletyvo

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/osyah/hryzun/status"
)

type ContextKey uint8

const ContextKeyNetwork ContextKey = iota

const (
	CodeNil status.Code = iota
	CodeInternal
	CodeNotFound
	CodePermissionDenied
	CodeInvalidArgument
	CodeUnauthorized
)

type QueryOption struct {
	Limit  int
	Order  bool
	After  uuid.UUID
	Before uuid.UUID
}

func (qo *QueryOption) Prepare() {
	if qo.Limit < 1 || qo.Limit > 50 {
		qo.Limit = 25
	}
}

func (qo QueryOption) String() string {
	var (
		buf   strings.Builder
		token byte = '?'
	)

	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			if qo.Limit != 0 {
				buf.WriteByte(token)
				buf.WriteString(fmt.Sprintf("limit=%d", qo.Limit))
			} else {
				continue
			}
		case 1:
			if qo.Order {
				buf.WriteByte(token)
				buf.WriteString("order=asc")
			} else {
				continue
			}
		case 2:
			if qo.After.Version() != 0 {
				buf.WriteByte(token)
				buf.WriteString("after=")
				buf.WriteString(qo.After.String())
			} else {
				continue
			}
		case 3:
			if qo.Before.Version() != 0 {
				buf.WriteByte(token)
				buf.WriteString("before=")
				buf.WriteString(qo.Before.String())
			} else {
				continue
			}
		}

		token = '&'
	}

	return buf.String()
}
