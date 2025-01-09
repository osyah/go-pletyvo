// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package deliveryhttpv1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Controller struct{ query *delivery.Query }

func New(query *delivery.Query) *Controller {
	return &Controller{query: query}
}

func (c Controller) RegisterRoutes(router fiber.Router) {
	v1 := router.Group("/v1")
	{
		NewChannel(c.query.Channel).RegisterRoutes(v1)
		NewPost(c.query.Post).RegisterRoutes(v1)
	}
}
