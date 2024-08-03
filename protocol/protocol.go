// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package protocol

import (
	"github.com/osyah/go-pletyvo/protocol/delivery"
	"github.com/osyah/go-pletyvo/protocol/registry"
)

type Executor struct {
	Registry *registry.Executor
	Delivery *delivery.Executor
}
