// Copyright (c) 2025 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"github.com/VictoriaMetrics/easyproto"
	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

func (DeliveryMessage) marshal(message *dapp.EventInput) []byte {
	m := mp.Get()

	mm := m.MessageMarshaler()
	mm.AppendBytes(1, message.Body)

	am := mm.AppendMessage(2)
	am.AppendUint32(1, uint32(message.Auth.Schema))
	am.AppendBytes(2, message.Auth.PublicKey)
	am.AppendBytes(3, message.Auth.Signature)

	dst := m.Marshal(nil)

	mp.Put(m)

	return dst
}

func (DeliveryMessage) unmarshal(src []byte, message *delivery.Message) (err error) {
	var fc easyproto.FieldContext

	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return status.New(
				pletyvo.CodeInternal,
				"go-pletyvo/store/pebble: cannot read next DeliveryMessage field",
			)
		}

		switch fc.FieldNum {
		case 1:
			body, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/store/pebble: cannot read DeliveryMessage body",
				)
			}

			message.Body = append(message.Body, body...)
		case 2:
			data, ok := fc.MessageData()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/store/pebble: cannot read DeliveryMessage auth",
				)
			}

			if err = unmarshalDAppAuth(data, &message.Auth); err != nil {
				return err
			}
		}
	}

	return
}
