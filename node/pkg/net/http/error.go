// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	switch t := err.(type) {
	case status.Status:
		return ctx.Status(WrapStatus(t.Code)).
			JSON(&ErrorResponse{Message: t.Message})
	case status.Code:
		return ctx.SendStatus(WrapStatus(t))
	default:
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
}

func WrapStatus(code status.Code) int {
	switch code {
	case pletyvo.CodeInternal:
		return fiber.StatusInternalServerError
	case pletyvo.CodeNotFound:
		return fiber.StatusNotFound
	case pletyvo.CodePermissionDenied:
		return fiber.StatusForbidden
	case pletyvo.CodeInvalidArgument:
		return fiber.StatusBadRequest
	case pletyvo.CodeUnauthorized:
		return fiber.StatusUnauthorized
	default:
		return fiber.StatusBadRequest
	}
}
