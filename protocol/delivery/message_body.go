// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package delivery

import (
	"encoding/base64"
	"encoding/json"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

const MessageCreate = 1 + iota

type MessageBody []byte

func NewMessageBody(messageType, dataType byte, value any) MessageBody {
	if dataType != dapp.JSONDataType {
		panic("go-pletyvo/protocol/delivery: unsupported message body data type")
	}

	data, err := json.Marshal(value)
	if err != nil {
		panic("go-pletyvo/protocol/delivery: " + err.Error())
	}

	body := make(MessageBody, (len(data) + 2))
	body[0] = messageType
	body[1] = dataType

	copy(body[2:], data[:])

	return body
}

func (mb MessageBody) Type() byte { return mb[0] }

func (mb MessageBody) DataType() byte { return mb[1] }

func (mb MessageBody) Data() []byte { return mb[2:] }

func (mb MessageBody) Bytes() []byte { return mb[:] }

func (mb MessageBody) String() string {
	return base64.RawURLEncoding.EncodeToString(mb[:])
}

func (mb MessageBody) MarshalJSON() ([]byte, error) {
	size := base64.RawURLEncoding.EncodedLen(len(mb)) + 2

	b := make([]byte, size)
	b[0], b[size-1] = '"', '"'

	base64.RawURLEncoding.Encode(b[1:size-1], mb[:])

	return b, nil
}

func (mb *MessageBody) UnmarshalJSON(b []byte) error {
	*mb = make(MessageBody, base64.RawURLEncoding.DecodedLen((len(b) - 2)))
	_, err := base64.RawURLEncoding.Decode(*mb, b[1:len(b)-1])

	return err
}

func (mb MessageBody) Unmarshal(v any) error {
	if mb.DataType() != dapp.JSONDataType {
		return pletyvo.CodeInvalidArgument
	}

	return json.Unmarshal(mb.Data(), v)
}
