// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package config

import (
	"github.com/osyah/go-pletyvo/node/relay"
	"github.com/osyah/go-pletyvo/node/service"
	"github.com/osyah/go-pletyvo/node/transport"
	"github.com/osyah/go-pletyvo/protocol"
)

type Node struct {
	Protocol  protocol.Config  `cfg:"protocol"`
	Relay     relay.Config     `cfg:"relay"`
	Service   service.Config   `cfg:"service"`
	Transport transport.Config `cfg:"transport"`
}
