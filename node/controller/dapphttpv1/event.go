// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapphttpv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo/node/pkg/net/http"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Event struct{ service dapp.EventService }

func NewEvent(service dapp.EventService) *Event {
	return &Event{service: service}
}

func (e Event) RegisterRoutes(router fiber.Router) {
	event := router.Group("/events")
	{
		event.Get("/", e.getHandler)
		event.Post("/", e.createHandler)
		event.Get("/:id", e.getByIDHandler)
	}
}

func (e Event) getHandler(ctx *fiber.Ctx) error {
	option, err := http.QueryOption(ctx)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	events, err := e.service.Get(ctx.Context(), option)
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.JSON(events)
}

func (e Event) createHandler(ctx *fiber.Ctx) error {
	var input dapp.EventInput

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	response, err := e.service.Create(ctx.Context(), &input)
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.JSON(response)
}

func (e Event) getByIDHandler(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	event, err := e.service.GetByID(ctx.Context(), id)
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.JSON(event)
}
