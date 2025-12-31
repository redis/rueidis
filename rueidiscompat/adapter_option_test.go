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
	"runtime"
	"testing"

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
