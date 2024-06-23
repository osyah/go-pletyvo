// Copyright (c) 2024 Osyah
// SPDX-License-Identifier: MIT

package registrypebble

import (
	"strings"

	"github.com/VictoriaMetrics/easyproto"
	"github.com/osyah/hryzun/status"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/registry"
)

func (n Network) marshal(network *registry.Network) []byte {
	mr := mp.Get()

	mm := mr.MessageMarshaler()
	mm.AppendBytes(1, network.Author[:])
	mm.AppendString(2, network.Name)

	dst := mr.Marshal(nil)

	mp.Put(mr)

	return dst
}

func (n Network) unmarshal(src []byte, network *registry.Network) (err error) {
	var fc easyproto.FieldContext

	network.EventHeader = &dapp.EventHeader{}
	network.NetworkInput = &registry.NetworkInput{}

	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return status.New(
				pletyvo.CodeInternal,
				"go-pletyvo/repository/registrypebble: cannot read next field",
			)
		}

		switch fc.FieldNum {
		case 1:
			author, ok := fc.Bytes()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/registrypebble: cannot read author",
				)
			}

			network.Author = dapp.Address(author)
		case 2:
			name, ok := fc.String()
			if !ok {
				return status.New(
					pletyvo.CodeInternal,
					"go-pletyvo/repository/registrypebble: cannot read name",
				)
			}

			network.Name = strings.Clone(name)
		}
	}

	return nil
}
