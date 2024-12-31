// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package httpapi

import (
	"github.com/gofiber/fiber/v2"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/node/pkg/net/http"
)

func NetworkMiddleware(ctx *fiber.Ctx) error {
	s := ctx.Get("Network")

	if len(s) != 0 {
		network, err := pletyvo.NetworkFromString(s)
		if err != nil {
			return http.ErrorHandler(ctx, err)
		}

		ctx.Context().SetUserValue(pletyvo.ContextKeyNetwork, network)
	}

	return ctx.Next()
}
