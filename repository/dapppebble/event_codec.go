// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package dapppebble

import (
	"github.com/VictoriaMetrics/easyproto"
	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
)

func (Event) marshal(event *dapp.SystemEvent) []byte {
	m := mp.Get()

	mm := m.MessageMarshaler()
	mm.AppendBytes(1, event.Body)

	am := mm.AppendMessage(2)
	am.AppendUint32(1, uint32(event.Auth.Schema))
	am.AppendBytes(2, event.Auth.PublicKey)
	am.AppendBytes(3, event.Auth.Signature)

	dst := m.Marshal(nil)

	mp.Put(m)

	return dst
}

func (e Event) unmarshal(src []byte, event *dapp.Event) (err error) {
	var fc easyproto.FieldContext

	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return status.New(
				pletyvo.CodeInternal,
				"go-pletyvo/repository/dapppebble: cannot read next field",
			)
		}

		switch fc.FieldNum {
		case 1:
			body, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/dapppebble: cannot read body",
				)
			}

			event.Body = append(event.Body, body...)
		case 2:
			data, ok := fc.MessageData()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/dapppebble: cannot read auth",
				)
			}

			if err := e.unmarshalAuth(data, &event.Auth); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e Event) unmarshalAuth(src []byte, auth *dapp.AuthHeader) (err error) {
	var fc easyproto.FieldContext

	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return status.New(
				pletyvo.CodeInternal,
				"go-pletyvo/repository/dapppebble: cannot read next auth field",
			)
		}

		switch fc.FieldNum {
		case 1:
			schema, ok := fc.Uint32()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/dapppebble: cannot read auth.schema",
				)
			}

			auth.Schema = byte(schema)
		case 2:
			publicKey, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/dapppebble: cannot read auth.public_key",
				)
			}

			auth.PublicKey = append(auth.PublicKey, publicKey...)
		case 3:
			signature, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/dapppebble: cannot read auth.signature",
				)
			}

			auth.Signature = append(auth.Signature, signature...)
		}
	}

	return nil
}
