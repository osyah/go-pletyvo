// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package main

import (
	"flag"

	"github.com/rs/zerolog/log"

	"github.com/osyah/go-pletyvo/node"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "",
		"The path to the configuration file. "+
			"This flag is required to run the application.",
	)
	flag.Parse()
}

func main() {
	if configPath == "" {
		log.Fatal().Msg(
			"You need to specify the path to the configuration file. " +
				"This can be done by adding a '-config' flag.",
		)
	}

	node.Run(configPath)
}
