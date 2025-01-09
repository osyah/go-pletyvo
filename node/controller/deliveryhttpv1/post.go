// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package deliveryhttpv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo/node/pkg/net/http"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Post struct{ service delivery.PostQuery }

func NewPost(service delivery.PostQuery) *Post {
	return &Post{service: service}
}

func (p Post) RegisterRoutes(router fiber.Router) {
	post := router.Group("/channel/:channel/posts")
	{
		post.Get("/", p.getHandler)
		post.Get("/:post", p.getByIDHandler)
	}
}

func (p Post) getHandler(ctx *fiber.Ctx) error {
	channel, err := uuid.Parse(ctx.Params("channel"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	option, err := http.QueryOption(ctx)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	posts, err := p.service.Get(ctx.Context(), channel, option)
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.JSON(posts)
}

func (p Post) getByIDHandler(ctx *fiber.Ctx) error {
	channel, err := uuid.Parse(ctx.Params("channel"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	id, err := uuid.Parse(ctx.Params("post"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	post, err := p.service.GetByID(ctx.Context(), channel, id)
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.JSON(post)
}
