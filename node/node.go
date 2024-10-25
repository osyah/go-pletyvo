// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package node

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/osyah/hryzun/container"
	"github.com/rs/zerolog/log"

	"github.com/osyah/go-pletyvo/node/config"
	"github.com/osyah/go-pletyvo/node/transport/httpapi"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
	"github.com/osyah/go-pletyvo/protocol/registry"
)

func Run(configPath string) {
	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	base := container.New()

	InitContainer(base, cfg.Node)

	server := httpapi.New(cfg.Node.Transport.HTTPAPI, &httpapi.Service{
		DApp:     container.Get[*dapp.Service](base, "dapp_local"),
		Registry: container.Get[*registry.Query](base, "registry_local"),
		Delivery: container.Get[*delivery.Query](base, "delivery_local"),
	}).Build()

	go server.Listen()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Str("signal", s.String()).Send()
	case err := <-server.Notify():
		log.Error().Err(err).Send()
	}

	if err := server.Shutdown(); err != nil {
		log.Fatal().Err(err).Send()
	}

	if err := base.Close(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
