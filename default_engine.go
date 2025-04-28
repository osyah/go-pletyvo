// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package pletyvo

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/osyah/hryzun"
)

type DefaultEngine interface {
	Get(ctx context.Context, endpoint string, value any) error
	Post(ctx context.Context, endpoint string, body, value any) error
}

const (
	contentTypeKey  = "Content-Type"
	contentTypeJSON = "application/json"

	networkIdentifyKey = "Network"
)

type EngineConfig struct {
	URL     string `cfg:"url"`
	Network string `cfg:"network"`
}

type Engine struct {
	EngineConfig
	doer *http.Client
}

func NewEngine(cfg EngineConfig) *Engine {
	return &Engine{EngineConfig: cfg, doer: http.DefaultClient}
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
		return e.ErrorHandler(resp)
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
		return e.ErrorHandler(resp)
	}

	if value == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(value)
}

func (e Engine) ErrorHandler(resp *http.Response) error {
	if resp.Header.Get(contentTypeKey) == contentTypeJSON {
		value := hryzun.Status{Code: e.WrapStatus(resp.StatusCode)}

		if err := json.NewDecoder(resp.Body).Decode(&value); err != nil {
			return err
		}

		return &value
	}

	return e.WrapStatus(resp.StatusCode)
}

func (Engine) WrapStatus(status int) hryzun.Code {
	switch status {
	case http.StatusInternalServerError:
		return CodeInternal
	case http.StatusNotFound:
		return CodeNotFound
	case http.StatusForbidden:
		return CodePermissionDenied
	case http.StatusBadRequest:
		return CodeInvalidArgument
	case http.StatusUnauthorized:
		return CodeUnauthorized
	case http.StatusConflict:
		return CodeConflict
	case http.StatusNotImplemented:
		return CodeNotImplemented
	default:
		return CodeInternal
	}
}
