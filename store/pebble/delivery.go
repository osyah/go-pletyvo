// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"github.com/cockroachdb/pebble"

	"github.com/osyah/go-pletyvo/protocol/delivery"
)

const (
	DeliveryChannelPrefix = 4
	DeliveryPostPrefix    = 5
	DeliveryMessagePrefix = 6
)

func NewDelivery(db *pebble.DB) *delivery.Repository {
	return &delivery.Repository{
		Channel: NewDeliveryChannel(db),
		Message: NewDeliveryMessage(db),
		Post:    NewDeliveryPost(db),
	}
}
