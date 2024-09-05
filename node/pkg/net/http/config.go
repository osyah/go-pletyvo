// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package http

import "github.com/gofiber/fiber/v2"

type Config struct {
	Address string      `cfg:"address"`
	Fiber   FiberConfig `cfg:"fiber"`
}

type FiberConfig struct {
	Prefork     bool   `cfg:"prefork"`
	BodyLimit   int    `cfg:"body_limit"`
	Concurrency int    `cfg:"concurrency"`
	AppName     string `cfg:"app_name"`
}

func (fc FiberConfig) Wrap() fiber.Config {
	return fiber.Config{
		Prefork:     fc.Prefork,
		BodyLimit:   fc.BodyLimit,
		Concurrency: fc.Concurrency,
		AppName:     fc.AppName,
	}
}
