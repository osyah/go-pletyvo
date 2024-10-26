// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

const (
	contentTypeKey  = "Content-Type"
	contentTypeJSON = "application/json"

	networkIdentifyKey = "Network"
)

type Config struct {
	URL     string `cfg:"url"`
	Network string `cfg:"network"`
}

type Engine struct {
	Config
	doer *http.Client
}

func New(cfg Config) *Engine {
	return &Engine{Config: cfg, doer: http.DefaultClient}
}

func (e Engine) Get(ctx context.Context, endpoint string, value any) error {
	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, (e.URL + endpoint), nil,
	)
	if err != nil {
		return err
	}

	if len(e.Network) != 0 {
		req.Header.Set(networkIdentifyKey, e.Network)
	}

	resp, err := e.doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrorHandler(resp)
	}

	return json.NewDecoder(resp.Body).Decode(value)
}

func (e Engine) Post(ctx context.Context, endpoint string, body, value any) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		(e.URL + endpoint),
		bytes.NewReader(b),
	)
	if err != nil {
		return err
	}

	req.Header.Set(contentTypeKey, contentTypeJSON)

	if len(e.Network) != 0 {
		req.Header.Set(networkIdentifyKey, e.Network)
	}

	resp, err := e.doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrorHandler(resp)
	}

	return json.NewDecoder(resp.Body).Decode(value)
}
