package rueidis

import (
	"bytes"
	"errors"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

type wrapped struct {
	msg string
	err error
}

func (e wrapped) Error() string { return e.msg }
func (e wrapped) Unwrap() error { return e.err }

func TestIsRedisNil(t *testing.T) {
	err := Nil
	if !IsRedisNil(err) {
		t.Fatal("IsRedisNil fail")
	}
	if IsRedisNil(errors.New("other")) {
		t.Fatal("IsRedisNil fail")
	}
	if err.Error() != "redis nil message" {
		t.Fatal("IsRedisNil fail")
	}
	wrappedErr := wrapped{msg: "wrapped", err: Nil}
	if IsRedisNil(wrappedErr) {
		t.Fatal("IsRedisNil fail : wrapped error")
	}
}

func TestIsRedisErr(t *testing.T) {
	err := Nil
	if ret, ok := IsRedisErr(err); ok || ret != Nil {
		t.Fatal("TestIsRedisErr fail")
	}
	if ret, ok := IsRedisErr(nil); ok || ret != nil {
		t.Fatal("TestIsRedisErr fail")
	}
	if ret, ok := IsRedisErr(errors.New("other")); ok || ret != nil {
		t.Fatal("TestIsRedisErr fail")
	}
	if ret, ok := IsRedisErr(&RedisError{typ: '-'}); !ok || ret.typ != '-' {
		t.Fatal("TestIsRedisErr fail")
	}
	wrappedErr := wrapped{msg: "wrapped", err: Nil}
	if ret, ok := IsRedisErr(wrappedErr); ok || ret == Nil {
		t.Fatal("TestIsRedisErr fail : wrapped error")
	}
}

//gocyclo:ignore
func TestRedisResult(t *testing.T) {
	t.Run("ToInt64", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).ToInt64(); err == nil {
			t.Fatal("ToInt64 not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).ToInt64(); err == nil {
			t.Fatal("ToInt64 not failed as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: ':', integer: 1}}).ToInt64(); v != 1 {
			t.Fatal("ToInt64 not get value as expected")
		}
	})

	t.Run("ToBool", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).ToBool(); err == nil {
			t.Fatal("ToBool not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).ToBool(); err == nil {
			t.Fatal("ToBool not failed as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: '#', integer: 1}}).ToBool(); !v {
			t.Fatal("ToBool not get value as expected")
		}
	})

	t.Run("AsBool", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsBool(); err == nil {
			t.Fatal("ToBool not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsBool(); err == nil {
			t.Fatal("ToBool not failed as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: '#', integer: 1}}).AsBool(); !v {
			t.Fatal("ToBool not get value as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: ':', integer: 1}}).AsBool(); !v {
			t.Fatal("ToBool not get value as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: '+', string: "OK"}}).AsBool(); !v {
			t.Fatal("ToBool not get value as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: '$', string: "OK"}}).AsBool(); !v {
			t.Fatal("ToBool not get value as expected")
		}
	})

	t.Run("ToFloat64", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).ToFloat64(); err == nil {
			t.Fatal("ToFloat64 not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).ToFloat64(); err == nil {
			t.Fatal("ToFloat64 not failed as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: ',', string: "0.1"}}).ToFloat64(); v != 0.1 {
			t.Fatal("ToFloat64 not get value as expected")
		}
	})

	t.Run("ToString", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).ToString(); err == nil {
			t.Fatal("ToString not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).ToString(); err == nil {
			t.Fatal("ToString not failed as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: '+', string: "0.1"}}).ToString(); v != "0.1" {
			t.Fatal("ToString not get value as expected")
		}
	})

	t.Run("AsReader", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsReader(); err == nil {
			t.Fatal("AsReader not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsReader(); err == nil {
			t.Fatal("AsReader not failed as expected")
		}
		r, _ := (RedisResult{val: RedisMessage{typ: '+', string: "0.1"}}).AsReader()
		bs, _ := io.ReadAll(r)
		if !bytes.Equal(bs, []byte("0.1")) {
			t.Fatalf("AsReader not get value as expected %v", bs)
		}
	})

	t.Run("AsBytes", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsBytes(); err == nil {
			t.Fatal("AsBytes not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsBytes(); err == nil {
			t.Fatal("AsBytes not failed as expected")
		}
		bs, _ := (RedisResult{val: RedisMessage{typ: '+', string: "0.1"}}).AsBytes()
		if !bytes.Equal(bs, []byte("0.1")) {
			t.Fatalf("AsBytes not get value as expected %v", bs)
		}
	})

	t.Run("DecodeJSON", func(t *testing.T) {
		v := map[string]string{}
		if err := (RedisResult{err: errors.New("other")}).DecodeJSON(&v); err == nil {
			t.Fatal("DecodeJSON not failed as expected")
		}
		if err := (RedisResult{val: RedisMessage{typ: '-'}}).DecodeJSON(&v); err == nil {
			t.Fatal("DecodeJSON not failed as expected")
		}
		if _ = (RedisResult{val: RedisMessage{typ: '+', string: `{"k":"v"}`}}).DecodeJSON(&v); v["k"] != "v" {
			t.Fatalf("DecodeJSON not get value as expected %v", v)
		}
	})

	t.Run("AsInt64", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsInt64(); err == nil {
			t.Fatal("AsInt64 not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsInt64(); err == nil {
			t.Fatal("AsInt64 not failed as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: '+', string: "1"}}).AsInt64(); v != 1 {
			t.Fatal("AsInt64 not get value as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: ':', integer: 2}}).AsInt64(); v != 2 {
			t.Fatal("AsInt64 not get value as expected")
		}
	})

	t.Run("AsUint64", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsUint64(); err == nil {
			t.Fatal("AsUint64 not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsUint64(); err == nil {
			t.Fatal("AsUint64 not failed as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: '+', string: "1"}}).AsUint64(); v != 1 {
			t.Fatal("AsUint64 not get value as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: ':', integer: 2}}).AsUint64(); v != 2 {
			t.Fatal("AsUint64 not get value as expected")
		}
	})

	t.Run("AsFloat64", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsFloat64(); err == nil {
			t.Fatal("AsFloat64 not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsFloat64(); err == nil {
			t.Fatal("AsFloat64 not failed as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: '+', string: "1.1"}}).AsFloat64(); v != 1.1 {
			t.Fatal("AsFloat64 not get value as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: ',', string: "2.2"}}).AsFloat64(); v != 2.2 {
			t.Fatal("AsFloat64 not get value as expected")
		}
	})

	t.Run("ToArray", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).ToArray(); err == nil {
			t.Fatal("ToArray not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).ToArray(); err == nil {
			t.Fatal("ToArray not failed as expected")
		}
		values := []RedisMessage{{string: "item", typ: '+'}}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: values}}).ToArray(); !reflect.DeepEqual(ret, values) {
			t.Fatal("ToArray not get value as expected")
		}
	})

	t.Run("AsStrSlice", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsStrSlice(); err == nil {
			t.Fatal("AsStrSlice not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsStrSlice(); err == nil {
			t.Fatal("AsStrSlice not failed as expected")
		}
		values := []RedisMessage{{string: "item", typ: '+'}}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: values}}).AsStrSlice(); !reflect.DeepEqual(ret, []string{"item"}) {
			t.Fatal("AsStrSlice not get value as expected")
		}
	})

	t.Run("AsIntSlice", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsIntSlice(); err == nil {
			t.Fatal("AsIntSlice not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsIntSlice(); err == nil {
			t.Fatal("AsIntSlice not failed as expected")
		}
		values := []RedisMessage{{integer: 2, typ: ':'}}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: values}}).AsIntSlice(); !reflect.DeepEqual(ret, []int64{2}) {
			t.Fatal("AsIntSlice not get value as expected")
		}
	})

	t.Run("AsFloatSlice", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsFloatSlice(); err == nil {
			t.Fatal("AsFloatSlice not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsFloatSlice(); err == nil {
			t.Fatal("AsFloatSlice not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{string: "fff", typ: ','}}}}).AsFloatSlice(); err == nil {
			t.Fatal("AsFloatSlice not failed as expected")
		}
		values := []RedisMessage{{integer: 1, typ: ':'}, {string: "2", typ: '+'}, {string: "3", typ: '$'}, {string: "4", typ: ','}}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: values}}).AsFloatSlice(); !reflect.DeepEqual(ret, []float64{1, 2, 3, 4}) {
			t.Fatal("AsFloatSlice not get value as expected")
		}
	})

	t.Run("AsMap", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsMap(); err == nil {
			t.Fatal("AsMap not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsMap(); err == nil {
			t.Fatal("AsMap not failed as expected")
		}
		values := []RedisMessage{{string: "key", typ: '+'}, {string: "value", typ: '+'}}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: values}}).AsMap(); !reflect.DeepEqual(map[string]RedisMessage{
			values[0].string: values[1],
		}, ret) {
			t.Fatal("AsMap not get value as expected")
		}
	})

	t.Run("AsStrMap", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsStrMap(); err == nil {
			t.Fatal("AsStrMap not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsStrMap(); err == nil {
			t.Fatal("AsStrMap not failed as expected")
		}
		values := []RedisMessage{{string: "key", typ: '+'}, {string: "value", typ: '+'}}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: values}}).AsStrMap(); !reflect.DeepEqual(map[string]string{
			values[0].string: values[1].string,
		}, ret) {
			t.Fatal("AsStrMap not get value as expected")
		}
	})

	t.Run("AsIntMap", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsIntMap(); err == nil {
			t.Fatal("AsIntMap not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsIntMap(); err == nil {
			t.Fatal("AsIntMap not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{string: "key", typ: '+'}, {string: "value", typ: '+'}}}}).AsIntMap(); err == nil {
			t.Fatal("AsIntMap not failed as expected")
		}
		values := []RedisMessage{{string: "k1", typ: '+'}, {string: "1", typ: '+'}, {string: "k2", typ: '+'}, {integer: 2, typ: ':'}}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: values}}).AsIntMap(); !reflect.DeepEqual(map[string]int64{
			"k1": 1,
			"k2": 2,
		}, ret) {
			t.Fatal("AsIntMap not get value as expected")
		}
	})

	t.Run("ToMap", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).ToMap(); err == nil {
			t.Fatal("ToMap not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).ToMap(); err == nil {
			t.Fatal("ToMap not failed as expected")
		}
		values := []RedisMessage{{string: "key", typ: '+'}, {string: "value", typ: '+'}}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '%', values: values}}).ToMap(); !reflect.DeepEqual(map[string]RedisMessage{
			values[0].string: values[1],
		}, ret) {
			t.Fatal("ToMap not get value as expected")
		}
	})

	t.Run("ToAny", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).ToAny(); err == nil {
			t.Fatal("ToAny not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).ToAny(); err == nil {
			t.Fatal("ToAny not failed as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '%', values: []RedisMessage{{typ: '+', string: "key"}, {typ: ':', integer: 1}}},
			{typ: '%', values: []RedisMessage{{typ: '+', string: "nil"}, {typ: '_'}}},
			{typ: '%', values: []RedisMessage{{typ: '+', string: "err"}, {typ: '-', string: "err"}}},
			{typ: ',', string: "1.2"},
			{typ: '+', string: "str"},
			{typ: '#', integer: 0},
			{typ: '-', string: "err"},
			{typ: '_'},
		}}}).ToAny(); !reflect.DeepEqual([]any{
			map[string]any{"key": int64(1)},
			map[string]any{"nil": nil},
			map[string]any{"err": &RedisError{typ: '-', string: "err"}},
			1.2,
			"str",
			false,
			&RedisError{typ: '-', string: "err"},
			nil,
		}, ret) {
			t.Fatal("ToAny not get value as expected")
		}
	})

	t.Run("AsXRangeEntry", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry not failed as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{string: "id", typ: '+'}, {typ: '*', values: []RedisMessage{{typ: '+', string: "a"}, {typ: '+', string: "b"}}}}}}).AsXRangeEntry(); !reflect.DeepEqual(XRangeEntry{
			ID:          "id",
			FieldValues: map[string]string{"a": "b"},
		}, ret) {
			t.Fatal("AsXRangeEntry not get value as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{string: "id", typ: '+'}, {typ: '_'}}}}).AsXRangeEntry(); !reflect.DeepEqual(XRangeEntry{
			ID:          "id",
			FieldValues: nil,
		}, ret) {
			t.Fatal("AsXRangeEntry not get value as expected")
		}
	})

	t.Run("AsXRange", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{{string: "id1", typ: '+'}, {typ: '*', values: []RedisMessage{{typ: '+', string: "a"}, {typ: '+', string: "b"}}}}},
			{typ: '*', values: []RedisMessage{{string: "id2", typ: '+'}, {typ: '_'}}},
		}}}).AsXRange(); !reflect.DeepEqual([]XRangeEntry{{
			ID:          "id1",
			FieldValues: map[string]string{"a": "b"},
		}, {
			ID:          "id2",
			FieldValues: nil,
		}}, ret) {
			t.Fatal("AsXRange not get value as expected")
		}
	})

	t.Run("AsXRead", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsXRead(); err == nil {
			t.Fatal("AsXRead not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsXRead(); err == nil {
			t.Fatal("AsXRead not failed as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '%', values: []RedisMessage{
			{typ: '+', string: "stream1"},
			{typ: '*', values: []RedisMessage{
				{typ: '*', values: []RedisMessage{{string: "id1", typ: '+'}, {typ: '*', values: []RedisMessage{{typ: '+', string: "a"}, {typ: '+', string: "b"}}}}},
				{typ: '*', values: []RedisMessage{{string: "id2", typ: '+'}, {typ: '_'}}},
			}},
			{typ: '+', string: "stream2"},
			{typ: '*', values: []RedisMessage{
				{typ: '*', values: []RedisMessage{{string: "id3", typ: '+'}, {typ: '*', values: []RedisMessage{{typ: '+', string: "c"}, {typ: '+', string: "d"}}}}},
			}},
		}}}).AsXRead(); !reflect.DeepEqual(map[string][]XRangeEntry{
			"stream1": {
				{ID: "id1", FieldValues: map[string]string{"a": "b"}},
				{ID: "id2", FieldValues: nil}},
			"stream2": {
				{ID: "id3", FieldValues: map[string]string{"c": "d"}},
			},
		}, ret) {
			t.Fatal("AsXRead not get value as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "stream1"},
				{typ: '*', values: []RedisMessage{
					{typ: '*', values: []RedisMessage{{string: "id1", typ: '+'}, {typ: '*', values: []RedisMessage{{typ: '+', string: "a"}, {typ: '+', string: "b"}}}}},
					{typ: '*', values: []RedisMessage{{string: "id2", typ: '+'}, {typ: '_'}}},
				}},
			}},
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "stream2"},
				{typ: '*', values: []RedisMessage{
					{typ: '*', values: []RedisMessage{{string: "id3", typ: '+'}, {typ: '*', values: []RedisMessage{{typ: '+', string: "c"}, {typ: '+', string: "d"}}}}},
				}},
			}},
		}}}).AsXRead(); !reflect.DeepEqual(map[string][]XRangeEntry{
			"stream1": {
				{ID: "id1", FieldValues: map[string]string{"a": "b"}},
				{ID: "id2", FieldValues: nil}},
			"stream2": {
				{ID: "id3", FieldValues: map[string]string{"c": "d"}},
			},
		}, ret) {
			t.Fatal("AsXRead not get value as expected")
		}
	})

	t.Run("AsZScore", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsZScore(); err == nil {
			t.Fatal("AsZScore not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsZScore(); err == nil {
			t.Fatal("AsZScore not failed as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '+', string: "m1"},
			{typ: '+', string: "1"},
		}}}).AsZScore(); !reflect.DeepEqual(ZScore{Member: "m1", Score: 1}, ret) {
			t.Fatal("AsZScore not get value as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '+', string: "m1"},
			{typ: ',', string: "1"},
		}}}).AsZScore(); !reflect.DeepEqual(ZScore{Member: "m1", Score: 1}, ret) {
			t.Fatal("AsZScore not get value as expected")
		}
	})

	t.Run("AsZScores", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsZScores(); err == nil {
			t.Fatal("AsZScores not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsZScores(); err == nil {
			t.Fatal("AsZScores not failed as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '+', string: "m1"},
			{typ: '+', string: "1"},
			{typ: '+', string: "m2"},
			{typ: '+', string: "2"},
		}}}).AsZScores(); !reflect.DeepEqual([]ZScore{
			{Member: "m1", Score: 1},
			{Member: "m2", Score: 2},
		}, ret) {
			t.Fatal("AsZScores not get value as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "m1"},
				{typ: ',', string: "1"},
			}},
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "m2"},
				{typ: ',', string: "2"},
			}},
		}}}).AsZScores(); !reflect.DeepEqual([]ZScore{
			{Member: "m1", Score: 1},
			{Member: "m2", Score: 2},
		}, ret) {
			t.Fatal("AsZScores not get value as expected")
		}
	})

	t.Run("AsLMPop", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsLMPop(); err == nil {
			t.Fatal("AsLMPop not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsLMPop(); err == nil {
			t.Fatal("AsLMPop not failed as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '+', string: "k"},
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "1"},
				{typ: '+', string: "2"},
			}},
		}}}).AsLMPop(); !reflect.DeepEqual(KeyValues{
			Key:    "k",
			Values: []string{"1", "2"},
		}, ret) {
			t.Fatal("AsZScores not get value as expected")
		}
	})

	t.Run("AsZMPop", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsZMPop(); err == nil {
			t.Fatal("AsZMPop not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsZMPop(); err == nil {
			t.Fatal("AsZMPop not failed as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '+', string: "k"},
			{typ: '*', values: []RedisMessage{
				{typ: '*', values: []RedisMessage{
					{typ: '+', string: "1"},
					{typ: ',', string: "1"},
				}},
				{typ: '*', values: []RedisMessage{
					{typ: '+', string: "2"},
					{typ: ',', string: "2"},
				}},
			}},
		}}}).AsZMPop(); !reflect.DeepEqual(KeyZScores{
			Key: "k",
			Values: []ZScore{
				{Member: "1", Score: 1},
				{Member: "2", Score: 2},
			},
		}, ret) {
			t.Fatal("AsZMPop not get value as expected")
		}
	})

	t.Run("AsFtSearch", func(t *testing.T) {
		if _, _, err := (RedisResult{err: errors.New("other")}).AsFtSearch(); err == nil {
			t.Fatal("AsFtSearch not failed as expected")
		}
		if _, _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsFtSearch(); err == nil {
			t.Fatal("AsFtSearch not failed as expected")
		}
		if n, ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: ':', integer: 3},
			{typ: '+', string: "1"},
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "k1"},
				{typ: '+', string: "v1"},
				{typ: '+', string: "kk"},
				{typ: '+', string: "vv"},
			}},
			{typ: '+', string: "2"},
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "k2"},
				{typ: '+', string: "v2"},
				{typ: '+', string: "kk"},
				{typ: '+', string: "vv"},
			}},
		}}}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "1", Doc: map[string]string{"k1": "v1", "kk": "vv"}},
			{Key: "2", Doc: map[string]string{"k2": "v2", "kk": "vv"}},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: ':', integer: 3},
			{typ: '+', string: "1"},
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "k1"},
				{typ: '+', string: "v1"},
				{typ: '+', string: "kk"},
				{typ: '+', string: "vv"},
			}},
		}}}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "1", Doc: map[string]string{"k1": "v1", "kk": "vv"}},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: ':', integer: 3},
			{typ: '+', string: "1"},
			{typ: '+', string: "2"},
		}}}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "1", Doc: nil},
			{Key: "2", Doc: nil},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: ':', integer: 3},
			{typ: '+', string: "1"},
		}}}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "1", Doc: nil},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: ':', integer: 3},
		}}}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
	})

	t.Run("asGeosearch", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsGeosearch(); err == nil {
			t.Fatal("asGeosearch not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsGeosearch(); err == nil {
			t.Fatal("asGeosearch not failed as expected")
		}
		//WithDist, WithHash, WithCoord
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k1"},
				{typ: '$', string: "2.5"},
				{typ: ':', integer: 1},
				{typ: '*', values: []RedisMessage{
					{typ: ',', string: "28.0473"},
					{typ: ',', string: "26.2041"},
				}},
			}},
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k2"},
				{typ: '$', string: "4.5"},
				{typ: ':', integer: 4},
				{typ: '*', values: []RedisMessage{
					{typ: ',', string: "72.4973"},
					{typ: ',', string: "13.2263"},
				}},
			}},
		}}}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", Dist: 2.5, GeoHash: 1, Longitude: 28.0473, Latitude: 26.2041},
			{Name: "k2", Dist: 4.5, GeoHash: 4, Longitude: 72.4973, Latitude: 13.2263},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//WithHash, WithCoord
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k1"},
				{typ: ':', integer: 1},
				{typ: '*', values: []RedisMessage{
					{typ: ',', string: "84.3877"},
					{typ: ',', string: "33.7488"},
				}},
			}},
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k2"},
				{typ: ':', integer: 4},
				{typ: '*', values: []RedisMessage{
					{typ: ',', string: "115.8613"},
					{typ: ',', string: "31.9523"},
				}},
			}},
		}}}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", GeoHash: 1, Longitude: 84.3877, Latitude: 33.7488},
			{Name: "k2", GeoHash: 4, Longitude: 115.8613, Latitude: 31.9523},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//WithDist, WithCoord
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k1"},
				{typ: '$', string: "2.50076"},
				{typ: '*', values: []RedisMessage{
					{typ: ',', string: "84.3877"},
					{typ: ',', string: "33.7488"},
				}},
			}},
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k2"},
				{typ: '$', string: "1024.96"},
				{typ: '*', values: []RedisMessage{
					{typ: ',', string: "115.8613"},
					{typ: ',', string: "31.9523"},
				}},
			}},
		}}}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", Dist: 2.50076, Longitude: 84.3877, Latitude: 33.7488},
			{Name: "k2", Dist: 1024.96, Longitude: 115.8613, Latitude: 31.9523},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//WithCoord
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k1"},
				{typ: '*', values: []RedisMessage{
					{typ: ',', string: "122.4194"},
					{typ: ',', string: "37.7749"},
				}},
			}},
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k2"},
				{typ: '*', values: []RedisMessage{
					{typ: ',', string: "35.6762"},
					{typ: ',', string: "139.6503"},
				}},
			}},
		}}}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", Longitude: 122.4194, Latitude: 37.7749},
			{Name: "k2", Longitude: 35.6762, Latitude: 139.6503},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//WithDist
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k1"},
				{typ: ',', string: "2.50076"},
			}},
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k2"},
				{typ: ',', string: strconv.FormatFloat(math.MaxFloat64, 'E', -1, 64)},
			}},
		}}}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", Dist: 2.50076},
			{Name: "k2", Dist: math.MaxFloat64},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//WithHash
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k1"},
				{typ: ':', integer: math.MaxInt64},
			}},
			{typ: '*', values: []RedisMessage{
				{typ: '$', string: "k2"},
				{typ: ':', integer: 22296},
			}},
		}}}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", GeoHash: math.MaxInt64},
			{Name: "k2", GeoHash: 22296},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//With no additional options
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '$', string: "k1"},
			{typ: '$', string: "k2"},
		}}}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1"},
			{Name: "k2"},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}

	})

	t.Run("IsCacheHit", func(t *testing.T) {
		if (RedisResult{err: errors.New("other")}).IsCacheHit() {
			t.Fatal("IsCacheHit not as expected")
		}
		if !(RedisResult{val: RedisMessage{attrs: cacheMark}}).IsCacheHit() {
			t.Fatal("IsCacheHit not as expected")
		}
	})

	t.Run("CacheTTL", func(t *testing.T) {
		if (RedisResult{err: errors.New("other")}).CacheTTL() != -1 {
			t.Fatal("CacheTTL != -1")
		}
		m := RedisMessage{}
		m.setExpireAt(time.Now().Add(time.Millisecond * 100).UnixMilli())
		if (RedisResult{val: m}).CacheTTL() <= 0 {
			t.Fatal("CacheTTL <= 0")
		}
		time.Sleep(150 * time.Millisecond)
		if (RedisResult{val: m}).CacheTTL() != 0 {
			t.Fatal("CacheTTL != 0")
		}
	})

	t.Run("CachePTTL", func(t *testing.T) {
		if (RedisResult{err: errors.New("other")}).CachePTTL() != -1 {
			t.Fatal("CachePTTL != -1")
		}
		m := RedisMessage{}
		m.setExpireAt(time.Now().Add(time.Millisecond * 100).UnixMilli())
		if (RedisResult{val: m}).CachePTTL() <= 0 {
			t.Fatal("CachePTTL <= 0")
		}
		time.Sleep(150 * time.Millisecond)
		if (RedisResult{val: m}).CachePTTL() != 0 {
			t.Fatal("CachePTTL != 0")
		}
	})

	t.Run("CachePXAT", func(t *testing.T) {
		if (RedisResult{err: errors.New("other")}).CachePXAT() != -1 {
			t.Fatal("CachePTTL != -1")
		}
		m := RedisMessage{}
		m.setExpireAt(time.Now().Add(time.Millisecond * 100).UnixMilli())
		if (RedisResult{val: m}).CachePXAT() <= 0 {
			t.Fatal("CachePXAT <= 0")
		}
	})
}

//gocyclo:ignore
func TestRedisMessage(t *testing.T) {
	t.Run("IsNil", func(t *testing.T) {
		if !(&RedisMessage{typ: '_'}).IsNil() {
			t.Fatal("IsNil fail")
		}
	})
	t.Run("Trim ERR prefix", func(t *testing.T) {
		// kvrocks: https://github.com/redis/rueidis/issues/152#issuecomment-1333923750
		if (&RedisMessage{typ: '-', string: "ERR no_prefix"}).Error().Error() != "no_prefix" {
			t.Fatal("fail to trim ERR")
		}
	})
	t.Run("ToInt64", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).ToInt64(); err == nil {
			t.Fatal("ToInt64 not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a RESP3 int64") {
				t.Fatal("ToInt64 not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).ToInt64()
	})

	t.Run("ToBool", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).ToBool(); err == nil {
			t.Fatal("ToBool not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a RESP3 bool") {
				t.Fatal("ToBool not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).ToBool()
	})

	t.Run("AsBool", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsBool(); err == nil {
			t.Fatal("AsBool not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a int, string or bool") {
				t.Fatal("AsBool not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).AsBool()
	})

	t.Run("ToFloat64", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).ToFloat64(); err == nil {
			t.Fatal("ToFloat64 not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a RESP3 float64") {
				t.Fatal("ToFloat64 not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).ToFloat64()
	})

	t.Run("ToString", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).ToString(); err == nil {
			t.Fatal("ToString not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type : is not a string") {
				t.Fatal("ToString not panic as expected")
			}
		}()
		(&RedisMessage{typ: ':'}).ToString()
	})

	t.Run("AsReader", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsReader(); err == nil {
			t.Fatal("AsReader not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type : is not a string") {
				t.Fatal("AsReader not panic as expected")
			}
		}()
		(&RedisMessage{typ: ':'}).AsReader()
	})

	t.Run("AsBytes", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsBytes(); err == nil {
			t.Fatal("AsBytes not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type : is not a string") {
				t.Fatal("AsBytes not panic as expected")
			}
		}()
		(&RedisMessage{typ: ':'}).AsBytes()
	})

	t.Run("DecodeJSON", func(t *testing.T) {
		if err := (&RedisMessage{typ: '_'}).DecodeJSON(nil); err == nil {
			t.Fatal("DecodeJSON not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type : is not a string") {
				t.Fatal("DecodeJSON not panic as expected")
			}
		}()
		(&RedisMessage{typ: ':'}).DecodeJSON(nil)
	})

	t.Run("AsInt64", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsInt64(); err == nil {
			t.Fatal("AsInt64 not failed as expected")
		}
		defer func() {
			if !strings.Contains(recover().(string), "redis message type * is not a string") {
				t.Fatal("AsInt64 not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*', values: []RedisMessage{{}}}).AsInt64()
	})

	t.Run("AsUint64", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsUint64(); err == nil {
			t.Fatal("AsUint64 not failed as expected")
		}
		defer func() {
			if !strings.Contains(recover().(string), "redis message type * is not a string") {
				t.Fatal("AsUint64 not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*', values: []RedisMessage{{}}}).AsUint64()
	})

	t.Run("AsFloat64", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsFloat64(); err == nil {
			t.Fatal("AsFloat64 not failed as expected")
		}
		defer func() {
			if !strings.Contains(recover().(string), "redis message type : is not a string") {
				t.Fatal("AsFloat64 not panic as expected")
			}
		}()
		(&RedisMessage{typ: ':'}).AsFloat64()
	})

	t.Run("ToArray", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).ToArray(); err == nil {
			t.Fatal("ToArray not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a array") {
				t.Fatal("ToArray not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).ToArray()
	})

	t.Run("AsStrSlice", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsStrSlice(); err == nil {
			t.Fatal("AsStrSlice not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a array") {
				t.Fatal("AsStrSlice not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).AsStrSlice()
	})

	t.Run("AsIntSlice", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsIntSlice(); err == nil {
			t.Fatal("AsIntSlice not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a array") {
				t.Fatal("AsIntSlice not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).AsIntSlice()
	})

	t.Run("AsFloatSlice", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsFloatSlice(); err == nil {
			t.Fatal("AsFloatSlice not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a array") {
				t.Fatal("AsFloatSlice not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).AsFloatSlice()
	})

	t.Run("AsMap", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsMap(); err == nil {
			t.Fatal("AsMap not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a map/array/set or its length is not even") {
				t.Fatal("AsMap not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).AsMap()
	})

	t.Run("AsStrMap", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsStrMap(); err == nil {
			t.Fatal("AsStrMap not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a map/array/set") {
				t.Fatal("AsMap not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).AsStrMap()
	})

	t.Run("AsIntMap", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsIntMap(); err == nil {
			t.Fatal("AsIntMap not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a map/array/set") {
				t.Fatal("AsMap not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).AsIntMap()
	})

	t.Run("ToMap", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).ToMap(); err == nil {
			t.Fatal("ToMap not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a RESP3 map") {
				t.Fatal("ToMap not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).ToMap()
	})

	t.Run("ToAny", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).ToAny(); err == nil {
			t.Fatal("ToAny not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a supported in ToAny") {
				t.Fatal("ToAny not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).ToAny()
	})

	t.Run("AsXRangeEntry - no range id", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry not failed as expected")
		}

		if _, err := (&RedisMessage{typ: '*'}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry not failed as expected")
		}

		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{{typ: '_'}, {typ: '%'}}}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type : is not a string") {
				t.Fatal("AsXRangeEntry not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*', values: []RedisMessage{{typ: ':'}, {typ: '%'}}}).AsXRangeEntry()
	})

	t.Run("AsXRangeEntry - no range field values", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry not failed as expected")
		}

		if _, err := (&RedisMessage{typ: '*'}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry not failed as expected")
		}

		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{{typ: '+'}, {typ: '-'}}}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a map/array/set") {
				t.Fatal("AsXRangeEntry not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*', values: []RedisMessage{{typ: '+'}, {typ: 't'}}}).AsXRangeEntry()
	})

	t.Run("AsXRange", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}

		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{{typ: '_'}}}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}
	})

	t.Run("AsXRead", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRead(); err == nil {
			t.Fatal("AsXRead not failed as expected")
		}
		if _, err := (&RedisMessage{typ: '%', values: []RedisMessage{
			{typ: '+', string: "stream1"},
			{typ: '*', values: []RedisMessage{{typ: '*', values: []RedisMessage{{string: "id1", typ: '+'}}}}},
		}}).AsXRead(); err == nil {
			t.Fatal("AsXRead not failed as expected")
		}
		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "stream1"},
			}},
		}}).AsXRead(); err == nil {
			t.Fatal("AsXRead not failed as expected")
		}
		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "stream1"},
				{typ: '*', values: []RedisMessage{{typ: '*', values: []RedisMessage{{string: "id1", typ: '+'}}}}},
			}},
		}}).AsXRead(); err == nil {
			t.Fatal("AsXRead not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a map/array/set") {
				t.Fatal("AsXRangeEntry not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).AsXRead()
	})

	t.Run("AsZScore", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsZScore(); err == nil {
			t.Fatal("AsZScore not failed as expected")
		}
		defer func() {
			if !strings.Contains(recover().(string), "redis message is not a map/array/set or its length is not 2") {
				t.Fatal("AsZScore not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*'}).AsZScore()
	})

	t.Run("AsZScores", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsZScores(); err == nil {
			t.Fatal("AsZScore not failed as expected")
		}
		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '+', string: "m1"},
			{typ: '+', string: "m2"},
		}}).AsZScores(); err == nil {
			t.Fatal("AsZScores not fails as expected")
		}
		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '*', values: []RedisMessage{
				{typ: '+', string: "m1"},
				{typ: '+', string: "m2"},
			}},
		}}).AsZScores(); err == nil {
			t.Fatal("AsZScores not fails as expected")
		}
	})

	t.Run("AsLMPop", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsLMPop(); err == nil {
			t.Fatal("AsLMPop not failed as expected")
		}
		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '+', string: "k"},
			{typ: '_'},
		}}).AsLMPop(); err == nil {
			t.Fatal("AsLMPop not fails as expected")
		}
		defer func() {
			if !strings.Contains(recover().(string), "redis message type * is not a LMPOP response") {
				t.Fatal("AsLMPop not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '+', string: "k"},
		}}).AsLMPop()
	})

	t.Run("AsZMPop", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsZMPop(); err == nil {
			t.Fatal("AsZMPop not failed as expected")
		}
		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '+', string: "k"},
			{typ: '_'},
		}}).AsZMPop(); err == nil {
			t.Fatal("AsZMPop not fails as expected")
		}
		defer func() {
			if !strings.Contains(recover().(string), "redis message type * is not a ZMPOP response") {
				t.Fatal("AsZMPop not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*', values: []RedisMessage{
			{typ: '+', string: "k"},
		}}).AsZMPop()
	})

	t.Run("AsFtSearch", func(t *testing.T) {
		if _, _, err := (&RedisMessage{typ: '_'}).AsFtSearch(); err == nil {
			t.Fatal("AsFtSearch not failed as expected")
		}
		defer func() {
			if !strings.Contains(recover().(string), "redis message type * is not a FT.SEARCH response") {
				t.Fatal("AsFtSearch not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*', values: []RedisMessage{}}).AsFtSearch()
	})

	t.Run("AsScanEntry", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsScanEntry(); err == nil {
			t.Fatal("AsScanEntry not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsScanEntry(); err == nil {
			t.Fatal("AsScanEntry not failed as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{string: "1", typ: '+'}, {typ: '*', values: []RedisMessage{{typ: '+', string: "a"}, {typ: '+', string: "b"}}}}}}).AsScanEntry(); !reflect.DeepEqual(ScanEntry{
			Cursor:   1,
			Elements: []string{"a", "b"},
		}, ret) {
			t.Fatal("AsScanEntry not get value as expected")
		}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{string: "0", typ: '+'}, {typ: '_'}}}}).AsScanEntry(); !reflect.DeepEqual(ScanEntry{}, ret) {
			t.Fatal("AsScanEntry not get value as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type * is not a scan response or its length is not at least 2") {
				t.Fatal("AsScanEntry not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*', values: []RedisMessage{{typ: ':'}}}).AsScanEntry()
	})

	t.Run("ToMap with non string key", func(t *testing.T) {
		defer func() {
			if !strings.Contains(recover().(string), "redis message type : as map key is not supported") {
				t.Fatal("ToMap not panic as expected")
			}
		}()
		(&RedisMessage{typ: '%', values: []RedisMessage{{typ: ':'}, {typ: ':'}}}).ToMap()
	})

	t.Run("IsCacheHit", func(t *testing.T) {
		if (&RedisMessage{typ: '_'}).IsCacheHit() {
			t.Fatal("IsCacheHit not as expected")
		}
		if !(&RedisMessage{typ: '_', attrs: cacheMark}).IsCacheHit() {
			t.Fatal("IsCacheHit not as expected")
		}
	})

	t.Run("CacheTTL", func(t *testing.T) {
		if (&RedisMessage{typ: '_'}).CacheTTL() != -1 {
			t.Fatal("CacheTTL != -1")
		}
		m := &RedisMessage{typ: '_'}
		m.setExpireAt(time.Now().Add(time.Millisecond * 100).UnixMilli())
		if m.CacheTTL() <= 0 {
			t.Fatal("CacheTTL <= 0")
		}
		time.Sleep(100 * time.Millisecond)
		if m.CachePTTL() > 0 {
			t.Fatal("CachePTTL > 0")
		}
	})

	t.Run("CachePTTL", func(t *testing.T) {
		if (&RedisMessage{typ: '_'}).CachePTTL() != -1 {
			t.Fatal("CachePTTL != -1")
		}
		m := &RedisMessage{typ: '_'}
		m.setExpireAt(time.Now().Add(time.Millisecond * 100).UnixMilli())
		if m.CachePTTL() <= 0 {
			t.Fatal("CachePTTL <= 0")
		}
		time.Sleep(100 * time.Millisecond)
		if m.CachePTTL() > 0 {
			t.Fatal("CachePTTL > 0")
		}
	})

	t.Run("CachePXAT", func(t *testing.T) {
		if (&RedisMessage{typ: '_'}).CachePXAT() != -1 {
			t.Fatal("CachePXAT != -1")
		}
		m := &RedisMessage{typ: '_'}
		m.setExpireAt(time.Now().Add(time.Millisecond * 100).UnixMilli())
		if m.CachePXAT() <= 0 {
			t.Fatal("CachePXAT <= 0")
		}
	})
}
