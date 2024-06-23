// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registryhttpv1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/osyah/go-pletyvo/node/pkg/net/http"
	"github.com/osyah/go-pletyvo/protocol/registry"
)

type Network struct{ service registry.NetworkQuery }

func NewNetwork(service registry.NetworkQuery) *Network {
	return &Network{service: service}
}

func (n Network) RegisterRoutes(router fiber.Router) {
	network := router.Group("/network")
	{
		network.Get("/", n.getHandler)
	}
}

func (n Network) getHandler(ctx *fiber.Ctx) error {
	channel, err := n.service.Get(ctx.Context())
	if err != nil {
		return http.ErrorHandler(ctx, err)
	}

	return ctx.JSON(channel)
}
