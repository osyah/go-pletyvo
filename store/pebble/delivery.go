// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"github.com/osyah/go-pletyvo/protocol/delivery"

	"github.com/cockroachdb/pebble"
)

const (
	DeliveryChannelPrefix = 4
	DeliveryMessagePrefix = 5
)

func NewDelivery(db *pebble.DB) *delivery.Repository {
	return &delivery.Repository{
		Channel: NewDeliveryChannel(db),
		Message: NewDeliveryMessage(db),
	}
}
