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
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Commands", func() {
	It("Setter & Getter", func() {
		err := errors.New("any")
		{
			cmd := &Cmd{}
			cmd.SetVal(1)
			Expect(cmd.Val()).To(Equal(1))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &StringCmd{}
			cmd.SetVal("1")
			Expect(cmd.Val()).To(Equal("1"))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &BoolCmd{}
			cmd.SetVal(true)
			Expect(cmd.Val()).To(Equal(true))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &IntCmd{}
			cmd.SetVal(1)
			Expect(cmd.Val()).To(Equal(int64(1)))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &DurationCmd{}
			cmd.SetVal(1)
			Expect(cmd.Val()).To(Equal(time.Duration(1)))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &StatusCmd{}
			cmd.SetVal("ok")
			Expect(cmd.Val()).To(Equal("ok"))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &SliceCmd{}
			cmd.SetVal([]interface{}{1})
			Expect(cmd.Val()).To(Equal([]interface{}{1}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &StringSliceCmd{}
			cmd.SetVal([]string{"1"})
			Expect(cmd.Val()).To(Equal([]string{"1"}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &IntSliceCmd{}
			cmd.SetVal([]int64{1})
			Expect(cmd.Val()).To(Equal([]int64{1}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &BoolSliceCmd{}
			cmd.SetVal([]bool{true})
			Expect(cmd.Val()).To(Equal([]bool{true}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &FloatSliceCmd{}
			cmd.SetVal([]float64{1})
			Expect(cmd.Val()).To(Equal([]float64{1}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &ZSliceCmd{}
			cmd.SetVal([]Z{{Score: 1}})
			Expect(cmd.Val()).To(Equal([]Z{{Score: 1}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &FloatCmd{}
			cmd.SetVal(1)
			Expect(cmd.Val()).To(Equal(float64(1)))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &StringStringMapCmd{}
			cmd.SetVal(map[string]string{"a": "b"})
			Expect(cmd.Val()).To(Equal(map[string]string{"a": "b"}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &StringIntMapCmd{}
			cmd.SetVal(map[string]int64{"a": 1})
			Expect(cmd.Val()).To(Equal(map[string]int64{"a": 1}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &StringStructMapCmd{}
			cmd.SetVal(map[string]struct{}{"a": {}})
			Expect(cmd.Val()).To(Equal(map[string]struct{}{"a": {}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &XMessageSliceCmd{}
			cmd.SetVal([]XMessage{{ID: "a"}})
			Expect(cmd.Val()).To(Equal([]XMessage{{ID: "a"}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &XStreamSliceCmd{}
			cmd.SetVal([]XStream{{Stream: "a"}})
			Expect(cmd.Val()).To(Equal([]XStream{{Stream: "a"}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &XPendingCmd{}
			cmd.SetVal(XPending{Count: 1})
			Expect(cmd.Val()).To(Equal(XPending{Count: 1}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &XPendingExtCmd{}
			cmd.SetVal([]XPendingExt{{ID: "a"}})
			Expect(cmd.Val()).To(Equal([]XPendingExt{{ID: "a"}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &XInfoGroupsCmd{}
			cmd.SetVal([]XInfoGroup{{Name: "a"}})
			Expect(cmd.Val()).To(Equal([]XInfoGroup{{Name: "a"}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &XInfoStreamCmd{}
			cmd.SetVal(XInfoStream{Length: 1})
			Expect(cmd.Val()).To(Equal(XInfoStream{Length: 1}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &XInfoStreamFullCmd{}
			cmd.SetVal(XInfoStreamFull{Length: 1})
			Expect(cmd.Val()).To(Equal(XInfoStreamFull{Length: 1}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &XInfoConsumersCmd{}
			cmd.SetVal([]XInfoConsumer{{Name: "a"}})
			Expect(cmd.Val()).To(Equal([]XInfoConsumer{{Name: "a"}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &ZWithKeyCmd{}
			cmd.SetVal(ZWithKey{Key: "a"})
			Expect(cmd.Val()).To(Equal(ZWithKey{Key: "a"}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &TimeCmd{}
			cmd.SetVal(time.Time{})
			Expect(cmd.Val()).To(Equal(time.Time{}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &ClusterSlotsCmd{}
			cmd.SetVal([]ClusterSlot{{Start: 1}})
			Expect(cmd.Val()).To(Equal([]ClusterSlot{{Start: 1}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &GeoPosCmd{}
			cmd.SetVal([]*GeoPos{nil})
			Expect(cmd.Val()).To(Equal([]*GeoPos{nil}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &GeoLocationCmd{}
			cmd.SetVal([]GeoLocation{{Name: "a"}})
			Expect(cmd.Val()).To(Equal([]GeoLocation{{Name: "a"}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &CommandsInfoCmd{}
			cmd.SetVal(map[string]CommandInfo{"a": {}})
			Expect(cmd.Val()).To(Equal(map[string]CommandInfo{"a": {}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
	})
})

var _ = Describe("RESP3 Cmd", func() {
	testCmd(true)
})

var _ = Describe("RESP2 Cmd", func() {
	testCmd(false)
})

func testCmd(resp3 bool) {
	var adapter Cmdable

	BeforeEach(func() {
		if resp3 {
			adapter = adapterresp3
		} else {
			adapter = adapterresp2
		}
		Expect(adapter.FlushAll(ctx).Err()).NotTo(HaveOccurred())
	})

	It("has val/err", func() {
		set := adapter.Set(ctx, "key", "hello", 0)
		Expect(set.Err()).NotTo(HaveOccurred())
		Expect(set.Val()).To(Equal("OK"))

		get := adapter.Get(ctx, "key")
		Expect(get.Err()).NotTo(HaveOccurred())
		Expect(get.Val()).To(Equal("hello"))

		Expect(set.Err()).NotTo(HaveOccurred())
		Expect(set.Val()).To(Equal("OK"))
	})

	It("has helpers", func() {
		set := adapter.Set(ctx, "key", "10", 0)
		Expect(set.Err()).NotTo(HaveOccurred())

		n, err := adapter.Get(ctx, "key").Int64()
		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(Equal(int64(10)))

		un, err := adapter.Get(ctx, "key").Uint64()
		Expect(err).NotTo(HaveOccurred())
		Expect(un).To(Equal(uint64(10)))

		f, err := adapter.Get(ctx, "key").Float64()
		Expect(err).NotTo(HaveOccurred())
		Expect(f).To(Equal(float64(10)))
	})

	It("supports float32", func() {
		f := float32(66.97)

		err := adapter.Set(ctx, "float_key", f, 0).Err()
		Expect(err).NotTo(HaveOccurred())

		val, err := adapter.Get(ctx, "float_key").Float32()
		Expect(err).NotTo(HaveOccurred())
		Expect(val).To(Equal(f))
	})

	It("supports time.Time", func() {
		tm := time.Date(2019, 1, 1, 9, 45, 10, 222125, time.UTC)

		err := adapter.Set(ctx, "time_key", tm, 0).Err()
		Expect(err).NotTo(HaveOccurred())

		s, err := adapter.Get(ctx, "time_key").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(s).To(Equal("2019-01-01T09:45:10.000222125Z"))

		tm2, err := adapter.Get(ctx, "time_key").Time()
		Expect(err).NotTo(HaveOccurred())
		Expect(tm2).To(BeTemporally("==", tm))
	})
}
