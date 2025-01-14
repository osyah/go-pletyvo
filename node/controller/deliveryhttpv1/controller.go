// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package deliveryhttpv1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type Controller struct{ service *delivery.Service }

func New(service *delivery.Service) *Controller {
	return &Controller{service: service}
}

func (c Controller) RegisterRoutes(router fiber.Router) {
	v1 := router.Group("/v1")
	{
		NewChannel(c.service.Channel).RegisterRoutes(v1)
		NewMessage(c.service.Message).RegisterRoutes(v1)
		NewPost(c.service.Post).RegisterRoutes(v1)
	}
}
