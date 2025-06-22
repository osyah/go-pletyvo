// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package database

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/osyah/go-pletyvo"
)

func (s Space[T]) GetRecords(ctx context.Context, query *QueryOption) ([]*T, error) {
	url := fmt.Sprintf("%s/records%s", s.URL, query.String())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(TokenHeader, s.Token)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, pletyvo.ConvertClientStatus(resp.StatusCode)
	}

	buf := make([]byte, resp.ContentLength)

	n, err := io.ReadFull(resp.Body, buf)
	if err != nil {
		return nil, err
	}

	values := make([]*T, 0, query.Limit)

	for i, size := 2, 0; n > i; i = size + 2 {
		size = (int(buf[i-1]) | (int(buf[i-2]) << 8)) + i
		values = append(values, s.Decoder(buf[i:size]))
	}

	return values, nil
}

func (s Space[T]) GetRecord(ctx context.Context, id []byte) (*T, error) {
	url := fmt.Sprintf(
		"%s/records/%s", s.URL, base64.RawURLEncoding.EncodeToString(id),
	)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(TokenHeader, s.Token)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, pletyvo.ConvertClientStatus(resp.StatusCode)
	}

	buf := make([]byte, resp.ContentLength)

	_, err = io.ReadFull(resp.Body, buf)
	if err != nil {
		return nil, err
	}

	return s.Decoder(buf), nil
}

func (s Space[T]) SetRecord(ctx context.Context, id []byte, value *T) error {
	url := fmt.Sprintf(
		"%s/records/%s", s.URL, base64.RawURLEncoding.EncodeToString(id),
	)
	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewReader(s.Encoder(value)),
	)
	if err != nil {
		return err
	}

	req.Header.Set(TokenHeader, s.Token)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return pletyvo.ConvertClientStatus(resp.StatusCode)
	}

	return nil
}

func (s Space[T]) DeleteRecord(ctx context.Context, id []byte) error {
	url := fmt.Sprintf(
		"%s/records/%s", s.URL, base64.RawURLEncoding.EncodeToString(id),
	)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set(TokenHeader, s.Token)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return pletyvo.ConvertClientStatus(resp.StatusCode)
	}

	return nil
}
