// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapp

import (
	"encoding/base64"
	"encoding/json"
)

const (
	EventTypeSize = 4

	EventBodyJSON = 0
)

type EventType [EventTypeSize]byte

func NewEventType(event, aggregate, version, protocol byte) EventType {
	return EventType{event, aggregate, version, protocol}
}

func (e EventType) Event() byte     { return e[0] }
func (e EventType) Aggregate() byte { return e[1] }
func (e EventType) Version() byte   { return e[2] }
func (e EventType) Protocol() byte  { return e[3] }

type EventBody []byte

func NewEventBody(data []byte, version byte, et EventType) EventBody {
	body := make(EventBody, (len(data) + 5))
	copy(body[5:], data[:])

	body.SetVersion(version)
	body.SetType(et)

	return body
}

func NewEventBodyJSON(value any, et EventType) EventBody {
	b, err := json.Marshal(value)
	if err != nil {
		panic("go-pletyvo/protocol/dapp: " + err.Error())
	}

	return NewEventBody(b, EventBodyJSON, et)
}

func (eb EventBody) Version() byte { return eb[0] }

func (eb EventBody) SetVersion(version byte) { eb[0] = version }

func (eb EventBody) Type() EventType {
	return EventType{eb[1], eb[2], eb[3], eb[4]}
}

func (eb EventBody) SetType(et EventType) {
	eb[1] = et.Event()
	eb[2] = et.Aggregate()
	eb[3] = et.Version()
	eb[4] = et.Protocol()
}

func (eb EventBody) Data() []byte { return eb[5:] }

func (eb EventBody) VerifyData() bool { return json.Valid(eb[5:]) }

func (eb EventBody) Bytes() []byte { return eb[:] }

func (eb EventBody) String() string {
	return base64.StdEncoding.EncodeToString(eb[:])
}
