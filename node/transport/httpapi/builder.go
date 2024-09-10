// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package httpapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/osyah/go-pletyvo/node/controller/dapphttpv1"
	"github.com/osyah/go-pletyvo/node/controller/deliveryhttpv1"
	"github.com/osyah/go-pletyvo/node/controller/registryhttpv1"
)

type Builder struct {
	config  Config
	service *Service
}

func New(config Config, service *Service) *Builder {
	return &Builder{config: config, service: service}
}

func (b Builder) register(router fiber.Router) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     b.config.CORS.AllowOrigins(),
		AllowMethods:     b.config.CORS.AllowMethods(),
		AllowHeaders:     b.config.CORS.AllowHeaders(),
		AllowCredentials: b.config.CORS.Credentials,
	}))

	router.Use(NetworkMiddleware)

	b.registerAPIRoutes(router.Group("/api"))
}

func (b Builder) registerAPIRoutes(router fiber.Router) {
	dapp := router.Group("/dapp")
	{
		dapphttpv1.New(b.service.DApp).RegisterRoutes(dapp)
	}

	delivery := router.Group("/delivery")
	{
		deliveryhttpv1.New(b.service.Delivery).RegisterRoutes(delivery)
	}

	registry := router.Group("/registry")
	{
		registryhttpv1.New(b.service.Registry).RegisterRoutes(registry)
	}
}
