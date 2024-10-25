// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package relay

import (
	"context"

	"github.com/osyah/hryzun/container"

	"github.com/osyah/go-pletyvo/node/relay/localdoer"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type Relay interface {
	OnEvent(context.Context, *dapp.Event) error
}

func Register(base *container.Base, config Config) {
	if config.LocalDoer == nil {
		panic("go-pletyvo/node/relay: 'local_doer' config not found")
	} else {
		base.RegisterHandler("local_doer", localdoer.Register(*config.LocalDoer))
	}
}
