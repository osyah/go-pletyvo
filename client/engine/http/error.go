// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package http

import (
	"encoding/json"
	"net/http"

	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
)

func ErrorHandler(resp *http.Response) error {
	if resp.Header.Get(contentTypeKey) == contentTypeJSON {
		value := status.Status{Code: WrapStatus(resp.StatusCode)}

		if err := json.NewDecoder(resp.Body).Decode(&value); err != nil {
			return err
		}

		return &value
	}

	return WrapStatus(resp.StatusCode)
}

func WrapStatus(status int) status.Code {
	switch status {
	case http.StatusInternalServerError:
		return pletyvo.CodeInternal
	case http.StatusNotFound:
		return pletyvo.CodeNotFound
	case http.StatusForbidden:
		return pletyvo.CodePermissionDenied
	case http.StatusBadRequest:
		return pletyvo.CodeInvalidArgument
	case http.StatusUnauthorized:
		return pletyvo.CodeUnauthorized
	case http.StatusConflict:
		return pletyvo.CodeConflict
	default:
		return pletyvo.CodeInternal
	}
}
