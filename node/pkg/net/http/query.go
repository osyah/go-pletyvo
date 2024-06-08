// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
)

func QueryOption(ctx *fiber.Ctx) (*pletyvo.QueryOption, error) {
	var option pletyvo.QueryOption

	if limit := ctx.QueryInt("limit"); limit != 0 {
		option.Limit = limit
	}

	if sort := ctx.Query("order"); sort == "asc" {
		option.Order = true
	}

	if after := ctx.Query("after"); after != "" {
		id, err := uuid.Parse(after)
		if err != nil {
			return nil, err
		}

		option.After = id
	}

	if before := ctx.Query("before"); before != "" {
		id, err := uuid.Parse(before)
		if err != nil {
			return nil, err
		}

		option.Before = id
	}

	return &option, nil
}
