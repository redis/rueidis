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
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RESP3 Pipeline Commands", func() {
	testAdapterPipeline(true)
})

var _ = Describe("RESP2 Pipeline Commands", func() {
	testAdapterPipeline(false)
})

func testAdapterPipeline(resp3 bool) {
	var adapter Cmdable

	BeforeEach(func() {
		if resp3 {
			adapter = adapterresp3
		} else {
			adapter = adapterresp2
		}
		Expect(adapter.FlushDB(ctx).Err()).NotTo(HaveOccurred())
		Expect(adapter.FlushAll(ctx).Err()).NotTo(HaveOccurred())
	})

	It("should Pipelined", func() {
		var echo, ping *StringCmd
		rets, err := adapter.Pipelined(ctx, func(pipe Pipeliner) error {
			echo = pipe.Echo(ctx, "hello")
			ping = pipe.Ping(ctx)
			Expect(echo.Err()).To(MatchError(placeholder.err))
			Expect(ping.Err()).To(MatchError(placeholder.err))
			return nil
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(rets).To(HaveLen(2))
		Expect(rets[0]).To(Equal(echo))
		Expect(rets[1]).To(Equal(ping))
		Expect(echo.Err()).NotTo(HaveOccurred())
		Expect(echo.Val()).To(Equal("hello"))
		Expect(ping.Err()).NotTo(HaveOccurred())
		Expect(ping.Val()).To(Equal("PONG"))
	})

	It("should Pipeline", func() {
		pipe := adapter.Pipeline()
		echo := pipe.Echo(ctx, "hello")
		ping := pipe.Ping(ctx)
		Expect(echo.Err()).To(MatchError(placeholder.err))
		Expect(ping.Err()).To(MatchError(placeholder.err))
		Expect(pipe.Len()).To(Equal(2))

		rets, err := pipe.Exec(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(pipe.Len()).To(Equal(0))
		Expect(rets).To(HaveLen(2))
		Expect(rets[0]).To(Equal(echo))
		Expect(rets[1]).To(Equal(ping))
		Expect(echo.Err()).NotTo(HaveOccurred())
		Expect(echo.Val()).To(Equal("hello"))
		Expect(ping.Err()).NotTo(HaveOccurred())
		Expect(ping.Val()).To(Equal("PONG"))
	})

	It("should Discard", func() {
		pipe := adapter.Pipeline()
		echo := pipe.Echo(ctx, "hello")
		ping := pipe.Ping(ctx)

		pipe.Discard()
		Expect(pipe.Len()).To(Equal(0))

		rets, err := pipe.Exec(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(rets).To(HaveLen(0))

		Expect(echo.Err()).To(MatchError(placeholder.err))
		Expect(ping.Err()).To(MatchError(placeholder.err))
	})
}
