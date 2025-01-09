// Copyright (c) 2024-2025 Osyah
// SPDX-License-Identifier: MIT

package pebble

import (
	"context"

	"github.com/cockroachdb/pebble"
	"github.com/google/uuid"

	"github.com/osyah/go-pletyvo"
	"github.com/osyah/go-pletyvo/protocol/dapp"
	"github.com/osyah/go-pletyvo/protocol/delivery"
)

type DeliveryPost struct{ db *pebble.DB }

func NewDeliveryPost(db *pebble.DB) *DeliveryPost {
	return &DeliveryPost{db: db}
}

func (DeliveryPost) key(network pletyvo.Network, channel, id *uuid.UUID, sufix byte) []byte {
	var key []byte

	if id.Version() != 0 {
		key = make([]byte, 37)
		copy(key[21:], id[:])
	} else {
		key = make([]byte, 22)
		key[21] = sufix
	}

	key[4] = DeliveryPostPrefix

	copy(key[0:4], network[2:6])
	copy(key[5:21], channel[:])

	return key
}

func (dp DeliveryPost) Get(ctx context.Context, ch uuid.UUID, option *pletyvo.QueryOption) ([]*delivery.Post, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	posts := make([]*delivery.Post, 0, option.Limit)

	iter, err := dp.db.NewIterWithContext(ctx, &pebble.IterOptions{
		LowerBound: dp.key(network, &ch, &option.After, 0),
		UpperBound: dp.key(network, &ch, &option.Before, 255),
	})
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	var next func() bool

	if option.Order {
		if !iter.First() {
			return nil, pletyvo.CodeNotFound
		}

		next = iter.Next
	} else {
		if !iter.Last() {
			return nil, pletyvo.CodeNotFound
		}

		next = iter.Prev
	}

	if option.After.Version() != 0 {
		if !next() {
			return nil, pletyvo.CodeNotFound
		}
	}

	for i := 0; i < option.Limit; i++ {
		var post delivery.Post

		if err := dp.unmarshal(iter.Value(), &post); err != nil {
			return nil, err
		}

		post.ID = uuid.UUID(iter.Key()[21:37])
		post.Channel = ch

		posts = append(posts, &post)

		if !next() {
			break
		}
	}

	return posts, nil
}

func (dp DeliveryPost) GetByID(ctx context.Context, ch, id uuid.UUID) (*delivery.Post, error) {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	b, closer, err := dp.db.Get(dp.key(network, &ch, &id, 0))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	var post delivery.Post

	if err := dp.unmarshal(b, &post); err != nil {
		return nil, err
	}

	post.ID = id
	post.Channel = ch

	return &post, nil
}

func (dp DeliveryPost) Create(ctx context.Context, event *dapp.SystemEvent, input *delivery.PostInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	channel, err := getAggregate(dp.db, network, &input.Channel)
	if err != nil {
		return err
	}

	batch := dp.db.NewBatchWithSize(3)
	defer batch.Close()

	batch.Set(dp.key(network, &channel, &event.ID, 0), dp.marshal(event, input), nil)

	saveEventAndHash(batch, network, event, &event.ID)

	return batch.Commit(pebble.Sync)
}

func (dp DeliveryPost) Update(ctx context.Context, event *dapp.SystemEvent, input *delivery.PostUpdateInput) error {
	network, ok := ctx.Value(pletyvo.ContextKeyNetwork).(pletyvo.Network)
	if !ok {
		network = pletyvo.DefaultNetwork
	}

	channel, err := getAggregate(dp.db, network, &input.Channel)
	if err != nil {
		return err
	}

	id, err := getAggregate(dp.db, network, &input.Post)
	if err != nil {
		return err
	}

	key := dp.key(network, &channel, &id, 0)

	b, closer, err := dp.db.Get(key)
	if err != nil {
		return err
	}
	defer closer.Close()

	var post delivery.Post

	if err := dp.unmarshal(b, &post); err != nil {
		return err
	}

	if post.Hash != input.Post {
		return pletyvo.CodeConflict
	}

	batch := dp.db.NewBatchWithSize(3)
	defer batch.Close()

	batch.Set(key, dp.marshal(event, input.PostInput), nil)

	saveEventAndHash(batch, network, event, &id)

	return batch.Commit(pebble.Sync)
}
