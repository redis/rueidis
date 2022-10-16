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
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
			_, e := cmd.Int()
			Expect(e).To(Equal(err))
			_, e = cmd.Int64()
			Expect(e).To(Equal(err))
			_, e = cmd.Uint64()
			Expect(e).To(Equal(err))
			_, e = cmd.Float32()
			Expect(e).To(Equal(err))
			_, e = cmd.Float64()
			Expect(e).To(Equal(err))
			_, e = cmd.Text()
			Expect(e).To(Equal(err))
			_, e = cmd.Bool()
			Expect(e).To(Equal(err))
			_, e = cmd.Slice()
			Expect(e).To(Equal(err))
			_, e = cmd.Int64Slice()
			Expect(e).To(Equal(err))
			_, e = cmd.Uint64Slice()
			Expect(e).To(Equal(err))
			_, e = cmd.Float32Slice()
			Expect(e).To(Equal(err))
			_, e = cmd.Float64Slice()
			Expect(e).To(Equal(err))
			_, e = cmd.BoolSlice()
			Expect(e).To(Equal(err))
			_, e = cmd.StringSlice()
			Expect(e).To(Equal(err))
		}
		{
			cmd := &Cmd{}
			cmd.SetVal(int64(1))
			Expect(cmd.Err()).To(BeNil())
			Expect(cmd.Val()).To(Equal(int64(1)))
			i, _ := cmd.Int()
			Expect(i).To(Equal(1))
			i64, _ := cmd.Int64()
			Expect(i64).To(Equal(int64(1)))
			u64, _ := cmd.Uint64()
			Expect(u64).To(Equal(uint64(1)))
			f32, _ := cmd.Float32()
			Expect(f32).To(Equal(float32(1)))
			f64, _ := cmd.Float64()
			Expect(f64).To(Equal(float64(1)))
			_, e := cmd.Text()
			Expect(e).NotTo(BeNil())
			b, _ := cmd.Bool()
			Expect(b).To(BeTrue())
			_, e = cmd.Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Int64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Uint64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float32Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.BoolSlice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.StringSlice()
			Expect(e).NotTo(BeNil())
		}
		{
			cmd := &Cmd{}
			cmd.SetVal("1")
			Expect(cmd.Err()).To(BeNil())
			i, _ := cmd.Int()
			Expect(i).To(Equal(1))
			i64, _ := cmd.Int64()
			Expect(i64).To(Equal(int64(1)))
			u64, _ := cmd.Uint64()
			Expect(u64).To(Equal(uint64(1)))
			f32, _ := cmd.Float32()
			Expect(f32).To(Equal(float32(1)))
			f64, _ := cmd.Float64()
			Expect(f64).To(Equal(float64(1)))
			t, _ := cmd.Text()
			Expect(t).To(Equal("1"))
			b, _ := cmd.Bool()
			Expect(b).To(BeTrue())
			_, e := cmd.Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Int64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Uint64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float32Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.BoolSlice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.StringSlice()
			Expect(e).NotTo(BeNil())
		}
		{
			cmd := &Cmd{}
			cmd.SetVal([]interface{}{"1"})
			Expect(cmd.Err()).To(BeNil())
			_, e := cmd.Int()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Int64()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Uint64()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float32()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float64()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Text()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Bool()
			Expect(e).NotTo(BeNil())
			s, _ := cmd.Slice()
			Expect(s).To(Equal([]interface{}{"1"}))
			si64, _ := cmd.Int64Slice()
			Expect(si64).To(Equal([]int64{1}))
			su64, _ := cmd.Uint64Slice()
			Expect(su64).To(Equal([]uint64{1}))
			sf32, _ := cmd.Float32Slice()
			Expect(sf32).To(Equal([]float32{1}))
			sf64, _ := cmd.Float64Slice()
			Expect(sf64).To(Equal([]float64{1}))
			bs, _ := cmd.BoolSlice()
			Expect(bs).To(Equal([]bool{true}))
			ss, _ := cmd.StringSlice()
			Expect(ss).To(Equal([]string{"1"}))
		}
		{
			cmd := &Cmd{}
			cmd.SetVal("Text")
			_, e := cmd.Int64()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Uint64()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float32()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float64()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Bool()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Int64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Uint64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float32Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.BoolSlice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.StringSlice()
			Expect(e).NotTo(BeNil())
		}
		{
			cmd := &Cmd{}
			cmd.SetVal([]interface{}{1})
			_, e := cmd.StringSlice()
			Expect(e).NotTo(BeNil())
		}
		{
			cmd := &Cmd{}
			cmd.SetVal([]interface{}{"Text"})
			_, e := cmd.Int64()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Uint64()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float32()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float64()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Bool()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Int64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Uint64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float32Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.Float64Slice()
			Expect(e).NotTo(BeNil())
			_, e = cmd.BoolSlice()
			Expect(e).NotTo(BeNil())
		}
		{
			cmd := &StringCmd{}
			cmd.SetVal("xxx")
			_, err := cmd.Float32()
			Expect(err).NotTo(BeNil())
			_, err = cmd.Float64()
			Expect(err).NotTo(BeNil())
		}
		{
			cmd := &StringCmd{}
			cmd.SetVal("1")
			Expect(cmd.Val()).To(Equal("1"))

			bs, _ := cmd.Bytes()
			Expect(bs).To(Equal([]byte("1")))

			bv, _ := cmd.Bool()
			Expect(bv).To(BeTrue())

			i, _ := cmd.Int()
			Expect(i).To(Equal(1))

			i64, _ := cmd.Int64()
			Expect(i64).To(Equal(int64(1)))

			u64, _ := cmd.Uint64()
			Expect(u64).To(Equal(uint64(1)))

			f32, _ := cmd.Float32()
			Expect(f32).To(Equal(float32(1)))

			f64, _ := cmd.Float64()
			Expect(f64).To(Equal(float64(1)))

			Expect(cmd.String()).To(Equal("1"))

			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))

			_, e := cmd.Bytes()
			Expect(e).To(Equal(err))

			_, e = cmd.Int()
			Expect(e).To(Equal(err))

			_, e = cmd.Int64()
			Expect(e).To(Equal(err))

			_, e = cmd.Uint64()
			Expect(e).To(Equal(err))

			_, e = cmd.Float32()
			Expect(e).To(Equal(err))

			_, e = cmd.Float64()
			Expect(e).To(Equal(err))

			_, e = cmd.Bool()
			Expect(e).To(Equal(err))

			_, e = cmd.Time()
			Expect(e).To(Equal(err))
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
			v, _ := cmd.Uint64()
			Expect(v).To(Equal(uint64(1)))
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
			ret, _ := cmd.Result()
			Expect(ret).To(Equal([]bool{true}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &FloatSliceCmd{}
			cmd.SetVal([]float64{1})
			Expect(cmd.Val()).To(Equal([]float64{1}))
			ret, _ := cmd.Result()
			Expect(ret).To(Equal([]float64{1}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &ScanCmd{}
			cmd.SetVal([]string{"1"}, 1)
			keys, cursor := cmd.Val()
			Expect(keys).To(Equal([]string{"1"}))
			Expect(cursor).To(Equal(uint64(1)))
			Expect(cmd.Err()).To(BeNil())
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
			m, _ := cmd.Result()
			Expect(m).To(Equal(map[string]int64{"a": 1}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &StringStructMapCmd{}
			cmd.SetVal(map[string]struct{}{"a": {}})
			Expect(cmd.Val()).To(Equal(map[string]struct{}{"a": {}}))
			m, _ := cmd.Result()
			Expect(m).To(Equal(map[string]struct{}{"a": {}}))
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
			v, _ := cmd.Result()
			Expect(v).To(Equal([]ClusterSlot{{Start: 1}}))
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
		{
			cmd := &XAutoClaimCmd{}
			cmd.SetVal([]XMessage{{ID: "1", Values: map[string]interface{}{}}}, "0")
			v, s := cmd.Val()
			Expect(v).To(Equal([]XMessage{{ID: "1", Values: map[string]interface{}{}}}))
			Expect(s).To(Equal("0"))
			Expect(cmd.Err()).To(BeNil())
		}
		{
			cmd := &XAutoClaimJustIDCmd{}
			cmd.SetVal([]string{"1"}, "0")
			v, s := cmd.Val()
			Expect(v).To(Equal([]string{"1"}))
			Expect(s).To(Equal("0"))
			Expect(cmd.Err()).To(BeNil())
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
