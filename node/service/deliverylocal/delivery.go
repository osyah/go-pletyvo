// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package deliverylocal

import "github.com/osyah/go-pletyvo/protocol/delivery"

func NewQuery(repos *delivery.Repository) *delivery.Query {
	return &delivery.Query{
		Channel: NewChannel(repos.Channel),
		Message: NewMessage(repos.Message),
	}
}
