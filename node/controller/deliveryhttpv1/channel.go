// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package deliveryhttpv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo/node/pkg/net/http"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Channel struct{ service delivery.ChannelQuery }

func NewChannel(service delivery.ChannelQuery) *Channel {
	return &Channel{service: service}
}

func (c Channel) RegisterRoutes(router fiber.Router) {
	channel := router.Group("/channels/:channel")
	{
		channel.Get("/", c.getByIDHandler)
	}
}

func (c Channel) getByIDHandler(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("channel"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	channel, err := c.service.GetByID(ctx.Context(), id)
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.JSON(channel)
}
