// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package config

import (
	"github.com/osyah/go-pletyvo/node/relay"
	"github.com/osyah/go-pletyvo/node/service"
	"github.com/osyah/go-pletyvo/node/store"
	"github.com/osyah/go-pletyvo/node/transport"
)

type Node struct {
	Store     store.Config     `cfg:"store"`
	Protocol  Protocol         `cfg:"protocol"`
	Relay     relay.Config     `cfg:"relay"`
	Service   service.Config   `cfg:"service"`
	Transport transport.Config `cfg:"transport"`
}

type Protocol struct {
	Delivery ProtocolBase `cfg:"delivery"`
}

type ProtocolBase struct {
	Repos string `cfg:"repos"`
}
