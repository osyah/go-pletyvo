// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapphttpv1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/osyah/go-pletyvo/node/pkg/net/http"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Hash struct{ service dapp.HashQuery }

func NewHash(service dapp.HashQuery) *Hash {
	return &Hash{service: service}
}

func (h Hash) RegisterRoutes(router fiber.Router) {
	router.Get("/hash/:id", h.getByIDHandler)
}

func (h Hash) getByIDHandler(ctx *fiber.Ctx) error {
	id, err := dapp.HashFromString(ctx.Params("id"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	header, err := h.service.GetByID(ctx.Context(), id)
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.JSON(header)
}
