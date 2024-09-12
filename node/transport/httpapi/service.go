// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package httpapi

import (
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
	"github.com/osyah/go-pletyvo/protocol/registry"
)

type Service struct {
	DApp     *dapp.Service
	Delivery *delivery.Query
	Registry *registry.Query
}