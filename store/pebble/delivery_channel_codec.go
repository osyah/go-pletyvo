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

func (DeliveryChannel) marshal(event *dapp.SystemEvent, input *delivery.ChannelInput) []byte {
	m := mp.Get()

	mm := m.MessageMarshaler()
	mm.AppendBytes(1, event.Hash[:])
	mm.AppendBytes(2, event.Author[:])
	mm.AppendString(3, input.Name)

	dst := m.Marshal(nil)

	mp.Put(m)

	return dst
}

func (DeliveryChannel) unmarshal(src []byte, channel *delivery.Channel) (err error) {
	var fc easyproto.FieldContext

	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return status.New(
				pletyvo.CodeInternal,
				"go-pletyvo/store/pebble: cannot read next DeliveryChannel field",
			)
		}

		switch fc.FieldNum {
		case 1:
			hash, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/store/pebble: cannot read DeliveryChannel hash",
				)
			}

			channel.Hash = dapp.Hash(hash)
		case 2:
			author, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/store/pebble: cannot read DeliveryChannel author",
				)
			}

			channel.Author = dapp.Hash(author)
		case 3:
			name, ok := fc.String()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/store/pebble: cannot read DeliveryChannel name",
				)
			}

			channel.Name = strings.Clone(name)
		}
	}

	return nil
}
