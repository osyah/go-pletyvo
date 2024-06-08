// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapphttpv1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Controller struct{ service *dapp.Service }

func New(service *dapp.Service) *Controller {
	return &Controller{service: service}
}

func (c Controller) RegisterRoutes(router fiber.Router) {
	v1 := router.Group("/v1")
	{
		NewEvent(c.service.Event).RegisterRoutes(v1)
	}
}
