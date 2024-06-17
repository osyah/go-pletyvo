// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registry

const (
	Protocol = 3
	Version  = 0
)

type Query struct{ Network NetworkQuery }

type Repository struct{ Network NetworkRepository }

type Service struct{ Network NetworkService }
