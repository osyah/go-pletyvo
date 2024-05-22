// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package deliverypebble

import (
	"strings"

	"github.com/VictoriaMetrics/easyproto"
	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

func (Channel) marshal(channel *delivery.Channel) []byte {
	m := mp.Get()

	mm := m.MessageMarshaler()
	mm.AppendString(1, channel.Name)
	mm.AppendBytes(2, channel.Owner[:])

	dst := m.Marshal(nil)

	mp.Put(m)

	return dst
}

func (Channel) unmarshal(src []byte, channel *delivery.Channel) (err error) {
	var fc easyproto.FieldContext

	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return status.New(
				pletyvo.CodeInternal,
				"go-pletyvo/repository/deliverypebble: cannot read next field",
			)
		}

		switch fc.FieldNum {
		case 1:
			name, ok := fc.String()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/deliverypebble: cannot read name",
				)
			}

			channel.Name = strings.Clone(name)
		case 2:
			owner, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/deliverypebble: cannot read owner",
				)
			}

			channel.Owner = dapp.Address(owner)
		}
	}

	return nil
}
