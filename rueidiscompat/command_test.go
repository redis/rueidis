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
	"fmt"
	"github.com/redis/rueidis"
	"github.com/stretchr/testify/assert"
	"testing"
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
			cmd.SetVal([]any{"1"})
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
			Expect(s).To(Equal([]any{"1"}))
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
			cmd.SetVal([]any{1})
			_, e := cmd.StringSlice()
			Expect(e).NotTo(BeNil())
		}
		{
			cmd := &Cmd{}
			cmd.SetVal([]any{"Text"})
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
			cmd.SetVal([]any{1})
			Expect(cmd.Val()).To(Equal([]any{1}))
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
			cmd := &ClusterShardsCmd{}
			cmd.SetVal([]ClusterShard{{Slots: []SlotRange{{Start: 1}}}})
			Expect(cmd.Val()).To(Equal([]ClusterShard{{Slots: []SlotRange{{Start: 1}}}}))
			v, _ := cmd.Result()
			Expect(v).To(Equal([]ClusterShard{{Slots: []SlotRange{{Start: 1}}}}))
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
			cmd.SetVal([]XMessage{{ID: "1", Values: map[string]any{}}}, "0")
			v, s := cmd.Val()
			Expect(v).To(Equal([]XMessage{{ID: "1", Values: map[string]any{}}}))
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
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &KeyValueSliceCmd{}
			cmd.SetVal([]KeyValue{{Key: "1", Value: "2"}})
			v := cmd.Val()
			Expect(v).To(Equal([]KeyValue{{Key: "1", Value: "2"}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &KeyValuesCmd{}
			cmd.SetVal("k", []string{"1"})
			k, v := cmd.Val()
			Expect(k).To(Equal("k"))
			Expect(v).To(Equal([]string{"1"}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &ZSliceWithKeyCmd{}
			cmd.SetVal("k", []Z{{Member: "1", Score: 1}})
			k, v := cmd.Val()
			Expect(k).To(Equal("k"))
			Expect(v).To(Equal([]Z{{Member: "1", Score: 1}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &KeyFlagsCmd{}
			cmd.SetVal([]KeyFlags{{Key: "k", Flags: []string{"1", "2"}}})
			v := cmd.Val()
			Expect(v).To(Equal([]KeyFlags{{Key: "k", Flags: []string{"1", "2"}}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &RankWithScoreCmd{}
			cmd.SetVal(RankScore{Rank: 1, Score: 2})
			v := cmd.Val()
			Expect(v).To(Equal(RankScore{Rank: 1, Score: 2}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &FunctionListCmd{}
			cmd.SetVal([]Library{{Name: "a"}})
			v := cmd.Val()
			Expect(v).To(Equal([]Library{{Name: "a"}}))
			cmd.SetErr(err)
			Expect(cmd.Err()).To(Equal(err))
		}
		{
			cmd := &JSONCmd{}
			cmd.SetVal("a")
			Expect(cmd.Val()).To(Equal("a"))
			v, _ := cmd.Result()
			Expect(v).To(Equal("a"))
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
func TestGeoSearchQueryArgs(t *testing.T) {
	tests := []struct {
		name     string
		query    GeoSearchQuery
		expected []string
	}{
		{
			name: "Radius with default unit",
			query: GeoSearchQuery{
				Radius: 5.0,
			},
			expected: []string{"FROMLONLAT", "0", "0", "BYRADIUS", "5", "KM"},
		},
		{
			name: "Radius with specified unit",
			query: GeoSearchQuery{
				Radius:     5.0,
				RadiusUnit: "M",
			},
			expected: []string{"FROMLONLAT", "0", "0", "BYRADIUS", "5", "M"},
		},
		{
			name: "Box with default unit",
			query: GeoSearchQuery{
				BoxWidth:  10.0,
				BoxHeight: 20.0,
			},
			expected: []string{"FROMLONLAT", "0", "0", "BYBOX", "10", "20", "KM"},
		},
		{
			name: "Box with specified unit",
			query: GeoSearchQuery{
				BoxWidth:  10.0,
				BoxHeight: 20.0,
				BoxUnit:   "M",
			},
			expected: []string{"FROMLONLAT", "0", "0", "BYBOX", "10", "20", "M"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.query.args()
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %v args, got %v", len(tt.expected), len(result))
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("Expected %v at position %d, got %v", tt.expected[i], i, v)
				}
			}
		})
	}
}

func TestSetErr(t *testing.T) {
	// Create a new XAutoClaimCmd instance
	cmd := &XAutoClaimCmd{}

	// Set an error using the SetErr method
	errMsg := "test error"
	cmd.SetErr(fmt.Errorf(errMsg))

	// Check if the error is properly set in the command object
	if cmd.Err() == nil {
		t.Error("expected non-nil error, got nil")
	}

	// Check if the error message matches the expected error message
	if cmd.Err().Error() != errMsg {
		t.Errorf("expected error message: %s, got: %s", errMsg, cmd.Err().Error())
	}
}
func TestStringInvalid(t *testing.T) {
	agg := Invalid
	expected := ""
	if result := agg.String(); result != expected {
		t.Errorf("Invalid: expected %s, got %s", expected, result)
	}
}

func TestStringAvg(t *testing.T) {
	agg := Avg
	expected := "AVG"
	if result := agg.String(); result != expected {
		t.Errorf("Avg: expected %s, got %s", expected, result)
	}
}

func TestStringSum(t *testing.T) {
	agg := Sum
	expected := "SUM"
	if result := agg.String(); result != expected {
		t.Errorf("Sum: expected %s, got %s", expected, result)
	}
}

func TestStringMin(t *testing.T) {
	agg := Min
	expected := "MIN"
	if result := agg.String(); result != expected {
		t.Errorf("Min: expected %s, got %s", expected, result)
	}
}

func TestStringMax(t *testing.T) {
	agg := Max
	expected := "MAX"
	if result := agg.String(); result != expected {
		t.Errorf("MAX: expected %s, got %s", expected, result)
	}
}

func TestStringRange(t *testing.T) {
	agg := Range
	expected := "RANGE"
	if result := agg.String(); result != expected {
		t.Errorf("Range: expected %s, got %s", expected, result)
	}
}

func TestStringCount(t *testing.T) {
	agg := Count
	expected := "COUNT"
	if result := agg.String(); result != expected {
		t.Errorf("Count: expected %s, got %s", expected, result)
	}
}

func TestStringFirst(t *testing.T) {
	agg := First
	expected := "FIRST"
	if result := agg.String(); result != expected {
		t.Errorf("First: expected %s, got %s", expected, result)
	}
}

func TestStringLast(t *testing.T) {
	agg := Last
	expected := "LAST"
	if result := agg.String(); result != expected {
		t.Errorf("Last: expected %s, got %s", expected, result)
	}
}

func TestStringStdP(t *testing.T) {
	agg := StdP
	expected := "STD.P"
	if result := agg.String(); result != expected {
		t.Errorf("StdP: expected %s, got %s", expected, result)
	}
}

func TestStringStdS(t *testing.T) {
	agg := StdS
	expected := "STD.S"
	if result := agg.String(); result != expected {
		t.Errorf("StdS: expected %s, got %s", expected, result)
	}
}

func TestStringVarP(t *testing.T) {
	agg := VarP
	expected := "VAR.P"
	if result := agg.String(); result != expected {
		t.Errorf("VarP: expected %s, got %s", expected, result)
	}
}

func TestStringVarS(t *testing.T) {
	agg := VarS
	expected := "VAR.S"
	if result := agg.String(); result != expected {
		t.Errorf("VarS: expected %s, got %s", expected, result)
	}
}

func TestStringTwa(t *testing.T) {
	agg := Aggregator(100)
	expected := ""
	if result := agg.String(); result != expected {
		t.Errorf("Empty string: expected %s, got %s", expected, result)
	}
}

func TestStringDefault(t *testing.T) {
	agg := Twa
	expected := "TWA"
	if result := agg.String(); result != expected {
		t.Errorf("Twa: expected %s, got %s", expected, result)
	}
}

func TestFormatMs(t *testing.T) {
	// Test case 1: Duration greater than 0 and less than 1 millisecond
	dur1 := time.Microsecond / 2 // Half a microsecond
	expected1 := int64(1)
	if result1 := formatMs(dur1); result1 != expected1 {
		t.Errorf("Test case 1 failed: Expected %d, got %d", expected1, result1)
	}

	// Test case 2: Duration equal to 1 millisecond
	dur2 := time.Millisecond
	expected2 := int64(1)
	if result2 := formatMs(dur2); result2 != expected2 {
		t.Errorf("Test case 2 failed: Expected %d, got %d", expected2, result2)
	}
}

func TestInitialErrorInRedisResult(t *testing.T) {
	t.Run("Initial error in RedisResult", func(t *testing.T) {
		mockRes := rueidis.NewMockResult(rueidis.RedisMessage{}, errors.New("Initial error"))
		cmd := newJSONSliceCmd(mockRes)
		assert.NotNil(t, cmd.Err())
		assert.EqualError(t, cmd.Err(), "Initial error")
		cmd2 := newIntPointerSliceCmd(mockRes)
		assert.NotNil(t, cmd2.Err())
		assert.EqualError(t, cmd2.Err(), "Initial error")
		cmd3 := newMapStringSliceInterfaceCmd(mockRes)
		assert.NotNil(t, cmd3.Err())
		assert.EqualError(t, cmd3.Err(), "Initial error")
		cmd4 := newTSTimestampValueSliceCmd(mockRes)
		assert.NotNil(t, cmd4.Err())
		assert.EqualError(t, cmd4.Err(), "Initial error")
		cmd5 := newMapStringInterfaceCmd(mockRes)
		assert.NotNil(t, cmd5.Err())
		assert.EqualError(t, cmd5.Err(), "Initial error")
		cmd6 := newTSTimestampValueCmd(mockRes)
		assert.NotNil(t, cmd6.Err())
		assert.EqualError(t, cmd6.Err(), "Initial error")

		cmd7 := newTDigestInfoCmd(mockRes)
		assert.NotNil(t, cmd7.Err())
		assert.EqualError(t, cmd7.Err(), "Initial error")

		cmd8 := newMapStringIntCmd(mockRes)
		assert.NotNil(t, cmd8.Err())
		assert.EqualError(t, cmd8.Err(), "Initial error")

		cmd9 := newTopKInfoCmd(mockRes)
		assert.NotNil(t, cmd9.Err())
		assert.EqualError(t, cmd9.Err(), "Initial error")

		cmd10 := newCMSInfoCmd(mockRes)
		assert.NotNil(t, cmd10.Err())
		assert.EqualError(t, cmd10.Err(), "Initial error")

		cmd11 := newCFInfoCmd(mockRes)
		assert.NotNil(t, cmd11.Err())
		assert.EqualError(t, cmd11.Err(), "Initial error")

		cmd12 := newScanDumpCmd(mockRes)
		assert.NotNil(t, cmd12.Err())
		assert.EqualError(t, cmd12.Err(), "Initial error")

		cmd13 := newBFInfoCmd(mockRes)
		assert.NotNil(t, cmd13.Err())
		assert.EqualError(t, cmd13.Err(), "Initial error")

		cmd14 := newMapStringInterfaceSliceCmd(mockRes)
		assert.NotNil(t, cmd14.Err())
		assert.EqualError(t, cmd14.Err(), "Initial error")

		cmd15 := newFunctionListCmd(mockRes)
		assert.NotNil(t, cmd15.Err())
		assert.EqualError(t, cmd15.Err(), "Initial error")

		cmd16 := newCommandsInfoCmd(mockRes)
		assert.NotNil(t, cmd16.Err())
		assert.EqualError(t, cmd16.Err(), "Initial error")

		cmd17 := newGeoPosCmd(mockRes)
		assert.NotNil(t, cmd17.Err())
		assert.EqualError(t, cmd17.Err(), "Initial error")

		cmd18 := newClusterShardsCmd(mockRes)
		assert.NotNil(t, cmd18.Err())
		assert.EqualError(t, cmd18.Err(), "Initial error")

		cmd19 := newClusterSlotsCmd(mockRes)
		assert.NotNil(t, cmd19.Err())
		assert.EqualError(t, cmd19.Err(), "Initial error")

		cmd20 := newTimeCmd(mockRes)
		assert.NotNil(t, cmd20.Err())
		assert.EqualError(t, cmd20.Err(), "Initial error")

		cmd21 := newXInfoConsumersCmd(mockRes)
		assert.NotNil(t, cmd21.Err())
		assert.EqualError(t, cmd21.Err(), "Initial error")

		cmd22 := newXInfoStreamFullCmd(mockRes)
		assert.NotNil(t, cmd22.Err())
		assert.EqualError(t, cmd22.Err(), "Initial error")

		cmd23 := newXInfoStreamCmd(mockRes)
		assert.NotNil(t, cmd23.Err())
		assert.EqualError(t, cmd23.Err(), "Initial error")

		cmd24 := newXInfoGroupsCmd(mockRes)
		assert.NotNil(t, cmd24.Err())
		assert.EqualError(t, cmd24.Err(), "Initial error")

		cmd25 := newXAutoClaimCmd(mockRes)
		assert.NotNil(t, cmd25.Err())
		assert.EqualError(t, cmd25.Err(), "Initial error")

		cmd26 := newXPendingExtCmd(mockRes)
		assert.NotNil(t, cmd26.Err())
		assert.EqualError(t, cmd26.Err(), "Initial error")

		cmd27 := newXPendingCmd(mockRes)
		assert.NotNil(t, cmd27.Err())
		assert.EqualError(t, cmd27.Err(), "Initial error")

		cmd28 := newStringStructMapCmd(mockRes)
		assert.NotNil(t, cmd28.Err())
		assert.EqualError(t, cmd28.Err(), "Initial error")

		cmd29 := newZSliceSingleCmd(mockRes)
		assert.NotNil(t, cmd29.Err())
		assert.EqualError(t, cmd29.Err(), "Initial error")

		cmd30 := newZSliceCmd(mockRes)
		assert.NotNil(t, cmd30.Err())
		assert.EqualError(t, cmd30.Err(), "Initial error")
	})
}
