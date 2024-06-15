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

func (m Message) marshal(message *delivery.Message) []byte {
	mr := mp.Get()

	mm := mr.MessageMarshaler()
	mm.AppendBytes(1, message.Author[:])
	mm.AppendString(2, message.Content)

	dst := mr.Marshal(nil)

	mp.Put(mr)

	return dst
}

func (m Message) unmarshal(src []byte, message *delivery.Message) (err error) {
	var fc easyproto.FieldContext

	message.EventHeader = &dapp.EventHeader{}
	message.MessageInput = &delivery.MessageInput{}

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
			author, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/deliverypebble: cannot read author",
				)
			}

			message.Author = dapp.Address(author)
		case 2:
			content, ok := fc.String()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/deliverypebble: cannot read content",
				)
			}

			message.Content = strings.Clone(content)
		}
	}

	return nil
}
