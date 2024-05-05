// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

const (
	contentTypeKey  = "Content-Type"
	contentTypeJSON = "application/json"
)

type Config struct {
	URL string `cfg:"url"`
}

type Engine struct {
	url  string
	doer *http.Client
}

func New(cfg Config) *Engine {
	return &Engine{url: cfg.URL, doer: http.DefaultClient}
}

func (e Engine) Get(ctx context.Context, endpoint string, value any) error {
	req, err := e.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := e.doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(value)
}

func (e Engine) Post(ctx context.Context, endpoint string, body, value any) error {
	req, err := e.NewRequest(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return err
	}

	req.Header.Set(contentTypeKey, contentTypeJSON)

	resp, err := e.doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(value)
}

func (e Engine) NewRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	var r io.Reader

	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		r = bytes.NewReader(b)
	}

	return http.NewRequestWithContext(ctx, method, (e.url + path), r)
}
