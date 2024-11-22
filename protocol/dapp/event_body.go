// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapp

import (
	"encoding/base64"
	"encoding/json"

	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
)

const (
	EventBodyBasic = 1 + iota
	EventBodyLinked
	maxEventBodyVersion

	JSONDataType = 1
)

var (
	ErrInvalidEventBodyVersion = status.New(
		pletyvo.CodeInvalidArgument, "invalid event body version",
	)
	ErrInvalidEventBodyDataType = status.New(
		pletyvo.CodeInvalidArgument, "invalid event body data type",
	)
	ErrInvalidEventType = status.New(
		pletyvo.CodeInvalidArgument, "unsupported event type",
	)

	EventBodyMetaSize = [3]int{EventBodyBasic: 4, EventBodyLinked: 36}
)

type EventBody []byte

func NewEventBody(version, dataType byte, eventType uint16, value any) EventBody {
	if dataType != JSONDataType {
		panic("go-pletyvo/protocol/dapp: unsupported event body data type")
	}

	data, err := json.Marshal(value)
	if err != nil {
		panic("go-pletyvo/protocol/dapp: " + err.Error())
	}

	body := make(EventBody, (len(data) + EventBodyMetaSize[version]))

	body[0] = version
	body[1] = dataType
	body[2] = byte(eventType >> 8)
	body[3] = byte(eventType)

	copy(body[EventBodyMetaSize[version]:], data[:])

	return body
}

func (eb EventBody) Version() byte { return eb[0] }

func (eb EventBody) DataType() byte { return eb[1] }

func (eb EventBody) Type() uint16 {
	return uint16(eb[3]) | (uint16(eb[2]) << 8)
}

func (eb EventBody) Data() []byte {
	return eb[EventBodyMetaSize[eb.Version()]:]
}

func (eb EventBody) Bytes() []byte { return eb[:] }

func (eb EventBody) String() string {
	return base64.RawURLEncoding.EncodeToString(eb[:])
}

func (eb EventBody) MarshalJSON() ([]byte, error) {
	size := base64.RawURLEncoding.EncodedLen(len(eb)) + 2

	b := make([]byte, size)
	b[0], b[size-1] = '"', '"'

	base64.RawURLEncoding.Encode(b[1:size-1], eb[:])

	return b, nil
}

func (eb *EventBody) UnmarshalJSON(b []byte) error {
	*eb = make(EventBody, base64.RawURLEncoding.DecodedLen((len(b) - 2)))
	_, err := base64.RawURLEncoding.Decode(*eb, b[1:len(b)-1])

	return err
}

func (eb EventBody) Unmarshal(v any) error {
	if eb.Version() >= maxEventBodyVersion {
		return ErrInvalidEventBodyVersion
	}

	if eb.DataType() != JSONDataType {
		return ErrInvalidEventBodyDataType
	}

	return json.Unmarshal(eb.Data(), v)
}

func (eb EventBody) Parent() Hash { return Hash(eb[4:36]) }

func (eb EventBody) SetParent(hash Hash) {
	if eb.Version() != EventBodyLinked {
		panic("go-pletyvo/protocol/dapp: event body dont support linked version")
	}

	copy(eb[4:36], hash[:])
}
