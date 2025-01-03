// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"strings"

	"github.com/VictoriaMetrics/easyproto"
	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

func (DeliveryMessage) marshal(event *dapp.SystemEvent, input *delivery.MessageInput) []byte {
	mr := mp.Get()

	mm := mr.MessageMarshaler()
	mm.AppendBytes(1, event.Hash[:])
	mm.AppendBytes(2, event.Author[:])
	mm.AppendString(3, input.Content)

	dst := mr.Marshal(nil)

	mp.Put(mr)

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
			hash, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/store/pebble: cannot read DeliveryMessage hash",
				)
			}

			message.Hash = dapp.Hash(hash)
		case 2:
			author, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/store/pebble: cannot read DeliveryMessage author",
				)
			}

			message.Author = dapp.Hash(author)
		case 3:
			content, ok := fc.String()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/store/pebble: cannot read DeliveryMessage content",
				)
			}

			message.Content = strings.Clone(content)
		}
	}

	return nil
}
