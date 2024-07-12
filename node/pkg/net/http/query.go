// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
)

func QueryOption(ctx *fiber.Ctx) (*pletyvo.QueryOption, error) {
	var (
		opt pletyvo.QueryOption
		err error
	)

	if limit := ctx.QueryInt("limit"); limit != 0 {
		opt.Limit = limit
	}

	if sort := ctx.Query("order"); sort == "asc" {
		opt.Order = true
	}

	if after := ctx.Query("after"); len(after) != 0 {
		opt.After, err = uuid.Parse(after)
		if err != nil {
			return nil, err
		}
	}

	if before := ctx.Query("before"); len(before) != 0 {
		opt.Before, err = uuid.Parse(before)
		if err != nil {
			return nil, err
		}
	}

	return &opt, nil
}
