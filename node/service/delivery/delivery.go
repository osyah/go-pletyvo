// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package delivery

import "github.com/osyah/go-pletyvo/protocol/delivery"

func NewQuery(repos *delivery.Query) *delivery.Query {
	return &delivery.Query{
		Channel: NewChannel(repos.Channel),
		Message: NewMessage(repos.Message),
	}
}
