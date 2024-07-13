// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registry

import (
	"context"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

type NetworkExecutor struct{ repos NetworkRepository }

func NewNetworkExecutor(repos NetworkRepository) *NetworkExecutor {
	return &NetworkExecutor{repos: repos}
}

func (ne NetworkExecutor) Create(ctx context.Context, base dapp.EventBase[NetworkCreateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	return ne.repos.Create(ctx, &Network{
		EventHeader:  base.EventHeader,
		NetworkInput: base.Input.NetworkInput,
	})
}

func (ne NetworkExecutor) Update(ctx context.Context, base dapp.EventBase[NetworkUpdateInput]) error {
	if err := base.Input.Validate(); err != nil {
		return err
	}

	network, err := ne.repos.Get(ctx)
	if err != nil {
		return err
	}

	if network.Author != base.Author {
		return pletyvo.CodePermissionDenied
	}

	return ne.repos.Update(ctx, base.Input)
}
