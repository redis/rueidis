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

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Commands", func() {
	ginkgo.It("Setter & Getter", func() {
		err := errors.New("any")
		{
			cmd := &Cmd{}
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
			_, e := cmd.Int()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Int64()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Uint64()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Float32()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Float64()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Text()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Bool()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Slice()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Int64Slice()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Uint64Slice()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Float32Slice()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.Float64Slice()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.BoolSlice()
			gomega.Expect(e).To(gomega.Equal(err))
			_, e = cmd.StringSlice()
			gomega.Expect(e).To(gomega.Equal(err))
		}
		{
			cmd := &Cmd{}
			cmd.SetVal(int64(1))
			gomega.Expect(cmd.Err()).To(gomega.BeNil())
			gomega.Expect(cmd.Val()).To(gomega.Equal(int64(1)))
			i, _ := cmd.Int()
			gomega.Expect(i).To(gomega.Equal(1))
			i64, _ := cmd.Int64()
			gomega.Expect(i64).To(gomega.Equal(int64(1)))
			u64, _ := cmd.Uint64()
			gomega.Expect(u64).To(gomega.Equal(uint64(1)))
			f32, _ := cmd.Float32()
			gomega.Expect(f32).To(gomega.Equal(float32(1)))
			f64, _ := cmd.Float64()
			gomega.Expect(f64).To(gomega.Equal(float64(1)))
			_, e := cmd.Text()
			gomega.Expect(e).NotTo(gomega.BeNil())
			b, _ := cmd.Bool()
			gomega.Expect(b).To(gomega.BeTrue())
			_, e = cmd.Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Int64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Uint64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float32Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.BoolSlice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.StringSlice()
			gomega.Expect(e).NotTo(gomega.BeNil())
		}
		{
			cmd := &Cmd{}
			cmd.SetVal("1")
			gomega.Expect(cmd.Err()).To(gomega.BeNil())
			i, _ := cmd.Int()
			gomega.Expect(i).To(gomega.Equal(1))
			i64, _ := cmd.Int64()
			gomega.Expect(i64).To(gomega.Equal(int64(1)))
			u64, _ := cmd.Uint64()
			gomega.Expect(u64).To(gomega.Equal(uint64(1)))
			f32, _ := cmd.Float32()
			gomega.Expect(f32).To(gomega.Equal(float32(1)))
			f64, _ := cmd.Float64()
			gomega.Expect(f64).To(gomega.Equal(float64(1)))
			t, _ := cmd.Text()
			gomega.Expect(t).To(gomega.Equal("1"))
			b, _ := cmd.Bool()
			gomega.Expect(b).To(gomega.BeTrue())
			_, e := cmd.Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Int64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Uint64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float32Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.BoolSlice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.StringSlice()
			gomega.Expect(e).NotTo(gomega.BeNil())
		}
		{
			cmd := &Cmd{}
			cmd.SetVal([]any{"1"})
			gomega.Expect(cmd.Err()).To(gomega.BeNil())
			_, e := cmd.Int()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Int64()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Uint64()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float32()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float64()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Text()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Bool()
			gomega.Expect(e).NotTo(gomega.BeNil())
			s, _ := cmd.Slice()
			gomega.Expect(s).To(gomega.Equal([]any{"1"}))
			si64, _ := cmd.Int64Slice()
			gomega.Expect(si64).To(gomega.Equal([]int64{1}))
			su64, _ := cmd.Uint64Slice()
			gomega.Expect(su64).To(gomega.Equal([]uint64{1}))
			sf32, _ := cmd.Float32Slice()
			gomega.Expect(sf32).To(gomega.Equal([]float32{1}))
			sf64, _ := cmd.Float64Slice()
			gomega.Expect(sf64).To(gomega.Equal([]float64{1}))
			bs, _ := cmd.BoolSlice()
			gomega.Expect(bs).To(gomega.Equal([]bool{true}))
			ss, _ := cmd.StringSlice()
			gomega.Expect(ss).To(gomega.Equal([]string{"1"}))
		}
		{
			cmd := &Cmd{}
			cmd.SetVal("Text")
			_, e := cmd.Int64()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Uint64()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float32()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float64()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Bool()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Int64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Uint64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float32Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.BoolSlice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.StringSlice()
			gomega.Expect(e).NotTo(gomega.BeNil())
		}
		{
			cmd := &Cmd{}
			cmd.SetVal([]any{1})
			_, e := cmd.StringSlice()
			gomega.Expect(e).NotTo(gomega.BeNil())
		}
		{
			cmd := &Cmd{}
			cmd.SetVal([]any{"Text"})
			_, e := cmd.Int64()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Uint64()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float32()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float64()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Bool()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Int64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Uint64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float32Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.Float64Slice()
			gomega.Expect(e).NotTo(gomega.BeNil())
			_, e = cmd.BoolSlice()
			gomega.Expect(e).NotTo(gomega.BeNil())
		}
		{
			cmd := &StringCmd{}
			cmd.SetVal("xxx")
			_, err := cmd.Float32()
			gomega.Expect(err).NotTo(gomega.BeNil())
			_, err = cmd.Float64()
			gomega.Expect(err).NotTo(gomega.BeNil())
		}
		{
			cmd := &StringCmd{}
			cmd.SetVal("1")
			gomega.Expect(cmd.Val()).To(gomega.Equal("1"))

			bs, _ := cmd.Bytes()
			gomega.Expect(bs).To(gomega.Equal([]byte("1")))

			bv, _ := cmd.Bool()
			gomega.Expect(bv).To(gomega.BeTrue())

			i, _ := cmd.Int()
			gomega.Expect(i).To(gomega.Equal(1))

			i64, _ := cmd.Int64()
			gomega.Expect(i64).To(gomega.Equal(int64(1)))

			u64, _ := cmd.Uint64()
			gomega.Expect(u64).To(gomega.Equal(uint64(1)))

			f32, _ := cmd.Float32()
			gomega.Expect(f32).To(gomega.Equal(float32(1)))

			f64, _ := cmd.Float64()
			gomega.Expect(f64).To(gomega.Equal(float64(1)))

			gomega.Expect(cmd.String()).To(gomega.Equal("1"))

			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))

			_, e := cmd.Bytes()
			gomega.Expect(e).To(gomega.Equal(err))

			_, e = cmd.Int()
			gomega.Expect(e).To(gomega.Equal(err))

			_, e = cmd.Int64()
			gomega.Expect(e).To(gomega.Equal(err))

			_, e = cmd.Uint64()
			gomega.Expect(e).To(gomega.Equal(err))

			_, e = cmd.Float32()
			gomega.Expect(e).To(gomega.Equal(err))

			_, e = cmd.Float64()
			gomega.Expect(e).To(gomega.Equal(err))

			_, e = cmd.Bool()
			gomega.Expect(e).To(gomega.Equal(err))

			_, e = cmd.Time()
			gomega.Expect(e).To(gomega.Equal(err))
		}
		{
			cmd := &BoolCmd{}
			cmd.SetVal(true)
			gomega.Expect(cmd.Val()).To(gomega.Equal(true))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &IntCmd{}
			cmd.SetVal(1)
			gomega.Expect(cmd.Val()).To(gomega.Equal(int64(1)))
			v, _ := cmd.Uint64()
			gomega.Expect(v).To(gomega.Equal(uint64(1)))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &DurationCmd{}
			cmd.SetVal(1)
			gomega.Expect(cmd.Val()).To(gomega.Equal(time.Duration(1)))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &StatusCmd{}
			cmd.SetVal("ok")
			gomega.Expect(cmd.Val()).To(gomega.Equal("ok"))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &SliceCmd{}
			cmd.SetVal([]any{1})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]any{1}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &StringSliceCmd{}
			cmd.SetVal([]string{"1"})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]string{"1"}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &IntSliceCmd{}
			cmd.SetVal([]int64{1})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]int64{1}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &BoolSliceCmd{}
			cmd.SetVal([]bool{true})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]bool{true}))
			ret, _ := cmd.Result()
			gomega.Expect(ret).To(gomega.Equal([]bool{true}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &FloatSliceCmd{}
			cmd.SetVal([]float64{1})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]float64{1}))
			ret, _ := cmd.Result()
			gomega.Expect(ret).To(gomega.Equal([]float64{1}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &ScanCmd{}
			cmd.SetVal([]string{"1"}, 1)
			keys, cursor := cmd.Val()
			gomega.Expect(keys).To(gomega.Equal([]string{"1"}))
			gomega.Expect(cursor).To(gomega.Equal(uint64(1)))
			gomega.Expect(cmd.Err()).To(gomega.BeNil())
		}
		{
			cmd := &ZSliceCmd{}
			cmd.SetVal([]Z{{Score: 1}})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]Z{{Score: 1}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &FloatCmd{}
			cmd.SetVal(1)
			gomega.Expect(cmd.Val()).To(gomega.Equal(float64(1)))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &StringStringMapCmd{}
			cmd.SetVal(map[string]string{"a": "b"})
			gomega.Expect(cmd.Val()).To(gomega.Equal(map[string]string{"a": "b"}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &StringIntMapCmd{}
			cmd.SetVal(map[string]int64{"a": 1})
			gomega.Expect(cmd.Val()).To(gomega.Equal(map[string]int64{"a": 1}))
			m, _ := cmd.Result()
			gomega.Expect(m).To(gomega.Equal(map[string]int64{"a": 1}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &StringStructMapCmd{}
			cmd.SetVal(map[string]struct{}{"a": {}})
			gomega.Expect(cmd.Val()).To(gomega.Equal(map[string]struct{}{"a": {}}))
			m, _ := cmd.Result()
			gomega.Expect(m).To(gomega.Equal(map[string]struct{}{"a": {}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &XMessageSliceCmd{}
			cmd.SetVal([]XMessage{{ID: "a"}})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]XMessage{{ID: "a"}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &XStreamSliceCmd{}
			cmd.SetVal([]XStream{{Stream: "a"}})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]XStream{{Stream: "a"}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &XPendingCmd{}
			cmd.SetVal(XPending{Count: 1})
			gomega.Expect(cmd.Val()).To(gomega.Equal(XPending{Count: 1}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &XPendingExtCmd{}
			cmd.SetVal([]XPendingExt{{ID: "a"}})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]XPendingExt{{ID: "a"}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &XInfoGroupsCmd{}
			cmd.SetVal([]XInfoGroup{{Name: "a"}})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]XInfoGroup{{Name: "a"}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &XInfoStreamCmd{}
			cmd.SetVal(XInfoStream{Length: 1})
			gomega.Expect(cmd.Val()).To(gomega.Equal(XInfoStream{Length: 1}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &XInfoStreamFullCmd{}
			cmd.SetVal(XInfoStreamFull{Length: 1})
			gomega.Expect(cmd.Val()).To(gomega.Equal(XInfoStreamFull{Length: 1}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &XInfoConsumersCmd{}
			cmd.SetVal([]XInfoConsumer{{Name: "a"}})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]XInfoConsumer{{Name: "a"}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &ZWithKeyCmd{}
			cmd.SetVal(ZWithKey{Key: "a"})
			gomega.Expect(cmd.Val()).To(gomega.Equal(ZWithKey{Key: "a"}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &TimeCmd{}
			cmd.SetVal(time.Time{})
			gomega.Expect(cmd.Val()).To(gomega.Equal(time.Time{}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &ClusterSlotsCmd{}
			cmd.SetVal([]ClusterSlot{{Start: 1}})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]ClusterSlot{{Start: 1}}))
			v, _ := cmd.Result()
			gomega.Expect(v).To(gomega.Equal([]ClusterSlot{{Start: 1}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &ClusterShardsCmd{}
			cmd.SetVal([]ClusterShard{{Slots: []SlotRange{{Start: 1}}}})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]ClusterShard{{Slots: []SlotRange{{Start: 1}}}}))
			v, _ := cmd.Result()
			gomega.Expect(v).To(gomega.Equal([]ClusterShard{{Slots: []SlotRange{{Start: 1}}}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &GeoPosCmd{}
			cmd.SetVal([]*GeoPos{nil})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]*GeoPos{nil}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &GeoLocationCmd{}
			cmd.SetVal([]GeoLocation{{Name: "a"}})
			gomega.Expect(cmd.Val()).To(gomega.Equal([]GeoLocation{{Name: "a"}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &CommandsInfoCmd{}
			cmd.SetVal(map[string]CommandInfo{"a": {}})
			gomega.Expect(cmd.Val()).To(gomega.Equal(map[string]CommandInfo{"a": {}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &XAutoClaimCmd{}
			cmd.SetVal([]XMessage{{ID: "1", Values: map[string]any{}}}, "0")
			v, s := cmd.Val()
			gomega.Expect(v).To(gomega.Equal([]XMessage{{ID: "1", Values: map[string]any{}}}))
			gomega.Expect(s).To(gomega.Equal("0"))
			gomega.Expect(cmd.Err()).To(gomega.BeNil())
		}
		{
			cmd := &XAutoClaimJustIDCmd{}
			cmd.SetVal([]string{"1"}, "0")
			v, s := cmd.Val()
			gomega.Expect(v).To(gomega.Equal([]string{"1"}))
			gomega.Expect(s).To(gomega.Equal("0"))
			gomega.Expect(cmd.Err()).To(gomega.BeNil())
		}
		{
			cmd := &KeyValueSliceCmd{}
			cmd.SetVal([]KeyValue{{Key: "1", Value: "2"}})
			v := cmd.Val()
			gomega.Expect(v).To(gomega.Equal([]KeyValue{{Key: "1", Value: "2"}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &KeyValuesCmd{}
			cmd.SetVal("k", []string{"1"})
			k, v := cmd.Val()
			gomega.Expect(k).To(gomega.Equal("k"))
			gomega.Expect(v).To(gomega.Equal([]string{"1"}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &ZSliceWithKeyCmd{}
			cmd.SetVal("k", []Z{{Member: "1", Score: 1}})
			k, v := cmd.Val()
			gomega.Expect(k).To(gomega.Equal("k"))
			gomega.Expect(v).To(gomega.Equal([]Z{{Member: "1", Score: 1}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &KeyFlagsCmd{}
			cmd.SetVal([]KeyFlags{{Key: "k", Flags: []string{"1", "2"}}})
			v := cmd.Val()
			gomega.Expect(v).To(gomega.Equal([]KeyFlags{{Key: "k", Flags: []string{"1", "2"}}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &RankWithScoreCmd{}
			cmd.SetVal(RankScore{Rank: 1, Score: 2})
			v := cmd.Val()
			gomega.Expect(v).To(gomega.Equal(RankScore{Rank: 1, Score: 2}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
		{
			cmd := &FunctionListCmd{}
			cmd.SetVal([]Library{{Name: "a"}})
			v := cmd.Val()
			gomega.Expect(v).To(gomega.Equal([]Library{{Name: "a"}}))
			cmd.SetErr(err)
			gomega.Expect(cmd.Err()).To(gomega.Equal(err))
		}
	})
})

var _ = ginkgo.Describe("RESP3 Cmd", func() {
	testCmd(true)
})

var _ = ginkgo.Describe("RESP2 Cmd", func() {
	testCmd(false)
})

func testCmd(resp3 bool) {
	var adapter Cmdable

	ginkgo.BeforeEach(func() {
		if resp3 {
			adapter = adapterresp3
		} else {
			adapter = adapterresp2
		}
		gomega.Expect(adapter.FlushAll(ctx).Err()).NotTo(gomega.HaveOccurred())
	})

	ginkgo.It("has val/err", func() {
		set := adapter.Set(ctx, "key", "hello", 0)
		gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
		gomega.Expect(set.Val()).To(gomega.Equal("OK"))

		get := adapter.Get(ctx, "key")
		gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
		gomega.Expect(get.Val()).To(gomega.Equal("hello"))

		gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
		gomega.Expect(set.Val()).To(gomega.Equal("OK"))
	})

	ginkgo.It("has helpers", func() {
		set := adapter.Set(ctx, "key", "10", 0)
		gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())

		n, err := adapter.Get(ctx, "key").Int64()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(n).To(gomega.Equal(int64(10)))

		un, err := adapter.Get(ctx, "key").Uint64()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(un).To(gomega.Equal(uint64(10)))

		f, err := adapter.Get(ctx, "key").Float64()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(f).To(gomega.Equal(float64(10)))
	})

	ginkgo.It("supports float32", func() {
		f := float32(66.97)

		err := adapter.Set(ctx, "float_key", f, 0).Err()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		val, err := adapter.Get(ctx, "float_key").Float32()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(val).To(gomega.Equal(f))
	})

	ginkgo.It("supports time.Time", func() {
		tm := time.Date(2019, 1, 1, 9, 45, 10, 222125, time.UTC)

		err := adapter.Set(ctx, "time_key", tm, 0).Err()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		s, err := adapter.Get(ctx, "time_key").Result()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(s).To(gomega.Equal("2019-01-01T09:45:10.000222125Z"))

		tm2, err := adapter.Get(ctx, "time_key").Time()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(tm2).To(gomega.BeTemporally("==", tm))
	})
}
