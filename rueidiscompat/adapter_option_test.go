// Copyright (c) 2013 The github.com/go-redis/redis Authors.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package rueidiscompat

import (
	"context"
	"errors"
	"runtime"
	"sync/atomic"
	"testing"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"go.uber.org/mock/gomock"
)

func TestWithNodeScaleoutLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewClient(ctrl)

	t.Run("default maxp is runtime.GOMAXPROCS(0)", func(t *testing.T) {
		adapter := NewAdapter(client)
		compat := adapter.(*Compat)
		if compat.maxp != runtime.GOMAXPROCS(0) {
			t.Errorf("expected maxp to be %d, got %d", runtime.GOMAXPROCS(0), compat.maxp)
		}
	})

	t.Run("WithNodeScaleoutLimit sets maxp", func(t *testing.T) {
		adapter := NewAdapter(client, WithNodeScaleoutLimit(4))
		compat := adapter.(*Compat)
		if compat.maxp != 4 {
			t.Errorf("expected maxp to be 4, got %d", compat.maxp)
		}
	})

	t.Run("multiple options are applied in order", func(t *testing.T) {
		adapter := NewAdapter(client, WithNodeScaleoutLimit(2), WithNodeScaleoutLimit(8))
		compat := adapter.(*Compat)
		if compat.maxp != 8 {
			t.Errorf("expected maxp to be 8, got %d", compat.maxp)
		}
	})

	t.Run("WithNodeScaleoutLimit with value less than 1 defaults to 1", func(t *testing.T) {
		adapter := NewAdapter(client, WithNodeScaleoutLimit(0))
		compat := adapter.(*Compat)
		if compat.maxp != 1 {
			t.Errorf("expected maxp to be 1, got %d", compat.maxp)
		}

		adapter = NewAdapter(client, WithNodeScaleoutLimit(-5))
		compat = adapter.(*Compat)
		if compat.maxp != 1 {
			t.Errorf("expected maxp to be 1, got %d", compat.maxp)
		}
	})
}

func TestForEachMaster(t *testing.T) {
	ctx := context.Background()

	t.Run("calls fn for each master node", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		masterNode := mock.NewClient(ctrl)
		replicaNode := mock.NewClient(ctrl)
		clusterClient := mock.NewClient(ctrl)

		clusterClient.EXPECT().Nodes().Return(map[string]rueidis.Client{
			"master:6379":  masterNode,
			"replica:6380": replicaNode,
		})

		masterNode.EXPECT().Do(gomock.Any(), mock.Match("ROLE")).Return(
			mock.Result(mock.RedisArray(mock.RedisString("master"))),
		)
		replicaNode.EXPECT().Do(gomock.Any(), mock.Match("ROLE")).Return(
			mock.Result(mock.RedisArray(mock.RedisString("slave"))),
		)

		var called atomic.Int32
		adapter := NewAdapter(clusterClient)
		err := adapter.ForEachMaster(ctx, func(ctx context.Context, client Cmdable) error {
			called.Add(1)
			return nil
		})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if called.Load() != 1 {
			t.Errorf("expected fn to be called 1 time, got %d", called.Load())
		}
	})

	t.Run("returns first error from fn", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		masterNode := mock.NewClient(ctrl)
		clusterClient := mock.NewClient(ctrl)

		clusterClient.EXPECT().Nodes().Return(map[string]rueidis.Client{
			"master:6379": masterNode,
		})

		masterNode.EXPECT().Do(gomock.Any(), mock.Match("ROLE")).Return(
			mock.Result(mock.RedisArray(mock.RedisString("master"))),
		)

		expectedErr := errors.New("test error")
		adapter := NewAdapter(clusterClient)
		err := adapter.ForEachMaster(ctx, func(ctx context.Context, client Cmdable) error {
			return expectedErr
		})
		if err == nil {
			t.Error("expected an error, got nil")
		} else if err.Error() != expectedErr.Error() {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("does not call fn when no master nodes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		replicaNode := mock.NewClient(ctrl)
		clusterClient := mock.NewClient(ctrl)

		clusterClient.EXPECT().Nodes().Return(map[string]rueidis.Client{
			"replica:6380": replicaNode,
		})

		replicaNode.EXPECT().Do(gomock.Any(), mock.Match("ROLE")).Return(
			mock.Result(mock.RedisArray(mock.RedisString("slave"))),
		)

		var called atomic.Int32
		adapter := NewAdapter(clusterClient)
		err := adapter.ForEachMaster(ctx, func(ctx context.Context, client Cmdable) error {
			called.Add(1)
			return nil
		})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if called.Load() != 0 {
			t.Errorf("expected fn to not be called, got %d", called.Load())
		}
	})

	t.Run("returns error when ROLE command fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		node := mock.NewClient(ctrl)
		clusterClient := mock.NewClient(ctrl)

		clusterClient.EXPECT().Nodes().Return(map[string]rueidis.Client{
			"node:6379": node,
		})

		roleErr := errors.New("connection refused")
		node.EXPECT().Do(gomock.Any(), mock.Match("ROLE")).Return(
			mock.ErrorResult(roleErr),
		)

		adapter := NewAdapter(clusterClient)
		err := adapter.ForEachMaster(ctx, func(ctx context.Context, client Cmdable) error {
			t.Error("fn should not be called when ROLE fails")
			return nil
		})
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("calls fn with multiple master nodes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		master1 := mock.NewClient(ctrl)
		master2 := mock.NewClient(ctrl)
		clusterClient := mock.NewClient(ctrl)

		clusterClient.EXPECT().Nodes().Return(map[string]rueidis.Client{
			"master1:6379": master1,
			"master2:6380": master2,
		})

		master1.EXPECT().Do(gomock.Any(), mock.Match("ROLE")).Return(
			mock.Result(mock.RedisArray(mock.RedisString("master"))),
		)
		master2.EXPECT().Do(gomock.Any(), mock.Match("ROLE")).Return(
			mock.Result(mock.RedisArray(mock.RedisString("master"))),
		)

		var called atomic.Int32
		adapter := NewAdapter(clusterClient)
		err := adapter.ForEachMaster(ctx, func(ctx context.Context, client Cmdable) error {
			called.Add(1)
			return nil
		})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if called.Load() != 2 {
			t.Errorf("expected fn to be called 2 times, got %d", called.Load())
		}
	})

	t.Run("works with empty nodes map", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		clusterClient := mock.NewClient(ctrl)

		clusterClient.EXPECT().Nodes().Return(map[string]rueidis.Client{})

		var called atomic.Int32
		adapter := NewAdapter(clusterClient)
		err := adapter.ForEachMaster(ctx, func(ctx context.Context, client Cmdable) error {
			called.Add(1)
			return nil
		})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if called.Load() != 0 {
			t.Errorf("expected fn to not be called, got %d", called.Load())
		}
	})

	t.Run("provides a working Cmdable to fn", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		masterNode := mock.NewClient(ctrl)
		clusterClient := mock.NewClient(ctrl)

		clusterClient.EXPECT().Nodes().Return(map[string]rueidis.Client{
			"master:6379": masterNode,
		})

		masterNode.EXPECT().Do(gomock.Any(), mock.Match("ROLE")).Return(
			mock.Result(mock.RedisArray(mock.RedisString("master"))),
		)

		adapter := NewAdapter(clusterClient)
		err := adapter.ForEachMaster(ctx, func(ctx context.Context, client Cmdable) error {
			if client == nil {
				t.Error("expected non-nil client")
			}
			if client.Client() != masterNode {
				t.Error("expected client to wrap the master node")
			}
			return nil
		})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
