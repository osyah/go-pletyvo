// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const shutdownTimeout = time.Second * 5

type Controller func(fiber.Router)

type Server struct {
	notify chan error
	config Config
	server *fiber.App
}

func NewServer(config Config) *Server {
	return &Server{
		notify: make(chan error, 1),
		config: config,
		server: fiber.New(config.Fiber.Wrap()),
	}
}

func (s Server) Listen() {
	s.notify <- s.server.Listen(s.config.Address)
}

func (s Server) Register(c Controller) { c(s.server) }

func (s Server) Notify() <-chan error { return s.notify }

func (s Server) Shutdown() error {
	return s.server.ShutdownWithTimeout(shutdownTimeout)
}
