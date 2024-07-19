// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package relay

import (
	"context"

	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Relay interface {
	OnEvent(context.Context, *dapp.Event) error
}
