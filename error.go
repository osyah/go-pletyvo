// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package pletyvo

import "net/http"

type Code uint8

const (
	CodeInternal Code = iota
	CodeInvalidArgument
	CodeNotFound
	CodeUnauthorized
	CodePermissionDenied
	maxCodeSize
)

var CodeToString = [maxCodeSize]string{
	CodeInternal:         "CODE_INTERNAL",
	CodeInvalidArgument:  "CODE_INVALID_ARGUMENT",
	CodeNotFound:         "CODE_NOT_FOUND",
	CodeUnauthorized:     "CODE_UNAUTHORIZED",
	CodePermissionDenied: "CODE_PERMISSION_DENIED",
}

func (c Code) Error() string { return CodeToString[c] }

func ConvertClientStatus(code int) Code {
	switch code {
	case http.StatusInternalServerError:
		return CodeInternal
	case http.StatusBadRequest:
		return CodeInvalidArgument
	case http.StatusForbidden:
		return CodePermissionDenied
	case http.StatusNotFound:
		return CodeNotFound
	case http.StatusUnauthorized:
		return CodeUnauthorized
	default:
		return CodeInternal
	}
}
