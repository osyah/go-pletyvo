// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package config

import (
	"github.com/osyah/go-pletyvo/node/relay"
	"github.com/osyah/go-pletyvo/node/service"
	"github.com/osyah/go-pletyvo/node/transport"
	"github.com/osyah/go-pletyvo/store"
)

type Node struct {
	Store     store.Config     `cfg:"store"`
	Protocol  ProtocolConfig   `cfg:"protocol"`
	Relay     relay.Config     `cfg:"relay"`
	Service   service.Config   `cfg:"service"`
	Transport transport.Config `cfg:"transport"`
}

type ProtocolConfig struct {
	Delivery ProtocolBase `cfg:"delivery"`
}

type ProtocolBase struct {
	Repos string `cfg:"repos"`
}
