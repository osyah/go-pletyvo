// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package deliveryhttpv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo/node/pkg/net/http"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Message struct{ service delivery.MessageService }

func NewMessage(service delivery.MessageService) *Message {
	return &Message{service: service}
}

func (m Message) RegisterRoutes(router fiber.Router) {
	channel := router.Group("/channel")
	{
		channel.Post("/send", m.sendHandler)
	}

	message := channel.Group("/:channel/messages")
	{
		message.Get("/", m.getHandler)
		message.Get("/:message", m.getByIDHandler)
	}
}

func (m Message) sendHandler(ctx *fiber.Ctx) error {
	var message delivery.Message

	err := ctx.BodyParser(&message)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = m.service.Send(ctx.Context(), &message)
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (m Message) getHandler(ctx *fiber.Ctx) error {
	channel, err := uuid.Parse(ctx.Params("channel"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	option, err := http.QueryOption(ctx)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	posts, err := m.service.Get(ctx.Context(), channel, option)
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.JSON(posts)
}

func (m Message) getByIDHandler(ctx *fiber.Ctx) error {
	channel, err := uuid.Parse(ctx.Params("channel"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(ctx.Params("message"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	post, err := m.service.GetByID(ctx.Context(), channel, id)
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.JSON(post)
}
