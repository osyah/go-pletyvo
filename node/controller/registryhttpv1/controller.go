// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registryhttpv1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/osyah/go-pletyvo/protocol/registry"
)

type Controller struct{ query *registry.Query }

func New(query *registry.Query) *Controller {
	return &Controller{query: query}
}

func (c Controller) RegisterRoutes(router fiber.Router) {
	v1 := router.Group("/v1")
	{
		NewNetwork(c.query.Network).RegisterRoutes(v1)
	}
}
