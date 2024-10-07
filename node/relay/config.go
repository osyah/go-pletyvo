// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package relay

import "github.com/osyah/go-pletyvo/node/relay/localdoer"

type Config struct {
	LocalDoer *localdoer.Config `cfg:"local_doer"`
}
