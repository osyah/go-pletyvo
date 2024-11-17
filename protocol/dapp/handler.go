// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapp

import "context"

type HandlerFunc func(context.Context, *Event) error

type Handler struct{ buf [1][9]HandlerFunc }

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Register(eventType uint16, handlerFunc HandlerFunc) {
	h.buf[byte((eventType >> 8))][byte(eventType)] = handlerFunc
}

func (h Handler) Handle(ctx context.Context, event *Event) error {
	f := h.buf[event.Body[2]][event.Body[3]]
	if f == nil {
		return ErrInvalidEventType
	}

	return f(ctx, event)
}
