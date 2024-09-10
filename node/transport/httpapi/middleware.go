// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package httpapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
)

func NetworkMiddleware(ctx *fiber.Ctx) error {
	s := ctx.Get("Network")

	if len(s) != 0 {
		id, err := uuid.Parse(s)
		if err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		ctx.Context().SetUserValue(pletyvo.ContextKeyNetwork, &id)
	} else {
		ctx.Context().SetUserValue(pletyvo.ContextKeyNetwork, &uuid.UUID{})
	}

	return ctx.Next()
}
