package rueidis

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"
)

type wrapped struct {
	err error
	msg string
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

func TestIsParseErr(t *testing.T) {
	err := errParse
	if !IsParseErr(err) {
		t.Fatal("IsParseErr fail")
	}
	if IsParseErr(errors.New("other")) {
		t.Fatal("IsParseErr fail")
	}
	if err.Error() != "rueidis: parse error" {
		t.Fatal("IsRedisNil fail")
	}
	wrappedErr := wrapped{msg: "wrapped", err: errParse}
	wrappedNonParseErr := wrapped{msg: "wrapped", err: errors.New("other")}
	if !IsParseErr(wrappedErr) || IsParseErr(wrappedNonParseErr) {
		t.Fatal("IsParseErr fail : wrapped error")
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

func TestRedisErrorIsMoved(t *testing.T) {
	for _, c := range []struct {
		err  string
		addr string
	}{
		{err: "MOVED 1 127.0.0.1:1", addr: "127.0.0.1:1"},
		{err: "MOVED 1 [::1]:1", addr: "[::1]:1"},
		{err: "MOVED 1 ::1:1", addr: "[::1]:1"},
	} {
		e := RedisError(strmsg('-', c.err))
		if addr, ok := e.IsMoved(); !ok || addr != c.addr {
			t.Fail()
		}
	}
}

func TestRedisErrorIsAsk(t *testing.T) {
	for _, c := range []struct {
		err  string
		addr string
	}{
		{err: "ASK 1 127.0.0.1:1", addr: "127.0.0.1:1"},
		{err: "ASK 1 [::1]:1", addr: "[::1]:1"},
		{err: "ASK 1 ::1:1", addr: "[::1]:1"},
	} {
		e := RedisError(strmsg('-', c.err))
		if addr, ok := e.IsAsk(); !ok || addr != c.addr {
			t.Fail()
		}
	}
}

func TestRedisErrorIsRedirect(t *testing.T) {
	for _, c := range []struct {
		err  string
		addr string
	}{
		{err: "REDIRECT 127.0.0.1:6380", addr: "127.0.0.1:6380"},
		{err: "REDIRECT [::1]:6380", addr: "[::1]:6380"},
		{err: "REDIRECT ::1:6380", addr: "[::1]:6380"},
	} {
		e := RedisError(strmsg('-', c.err))
		if addr, ok := e.IsRedirect(); !ok || addr != c.addr {
			t.Fail()
		}
	}
}

func TestIsRedisRedirect(t *testing.T) {
	err := errors.New("other")
	if ret, yes := IsRedisErr(err); yes {
		if addr, ok := ret.IsRedirect(); ok {
			t.Fatalf("TestIsRedisRedirect fail: expected false, got addr=%s", addr)
		}
	}

	redisErr := RedisError(strmsg('-', "REDIRECT 127.0.0.1:6380"))
	err = &redisErr
	if ret, yes := IsRedisErr(err); yes {
		if addr, ok := ret.IsRedirect(); !ok || addr != "127.0.0.1:6380" {
			t.Fatalf("TestIsRedisRedirect fail: expected addr=127.0.0.1:6380, got addr=%s, ok=%t", addr, ok)
		}
	} else {
		t.Fatal("TestIsRedisRedirect fail: expected RedisError")
	}
}

func TestIsRedisBusyGroup(t *testing.T) {
	err := errors.New("other")
	if IsRedisBusyGroup(err) {
		t.Fatal("TestIsRedisBusyGroup fail")
	}

	redisErr := RedisError(strmsg('-', "BUSYGROUP Consumer Group name already exists"))
	err = &redisErr
	if !IsRedisBusyGroup(err) {
		t.Fatal("TestIsRedisBusyGroup fail")
	}
}

//gocyclo:ignore
func TestRedisResult(t *testing.T) {
	//Add erroneous type
	typeNames['t'] = "t"

	t.Run("ToInt64", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).ToInt64(); err == nil {
			t.Fatal("ToInt64 not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).ToInt64(); err == nil {
			t.Fatal("ToInt64 not failed as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: ':', intlen: 1}}).ToInt64(); v != 1 {
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
		if v, _ := (RedisResult{val: RedisMessage{typ: '#', intlen: 1}}).ToBool(); !v {
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
		if v, _ := (RedisResult{val: RedisMessage{typ: '#', intlen: 1}}).AsBool(); !v {
			t.Fatal("ToBool not get value as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: ':', intlen: 1}}).AsBool(); !v {
			t.Fatal("ToBool not get value as expected")
		}
		if v, _ := (RedisResult{val: strmsg('+', "OK")}).AsBool(); !v {
			t.Fatal("ToBool not get value as expected")
		}
		if v, _ := (RedisResult{val: strmsg('$', "OK")}).AsBool(); !v {
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
		if v, _ := (RedisResult{val: strmsg(',', "0.1")}).ToFloat64(); v != 0.1 {
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
		if v, _ := (RedisResult{val: strmsg('+', "0.1")}).ToString(); v != "0.1" {
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
		r, _ := (RedisResult{val: strmsg('+', "0.1")}).AsReader()
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
		bs, _ := (RedisResult{val: strmsg('+', "0.1")}).AsBytes()
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
		if _ = (RedisResult{val: strmsg('+', `{"k":"v"}`)}).DecodeJSON(&v); v["k"] != "v" {
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
		if v, _ := (RedisResult{val: strmsg('+', "1")}).AsInt64(); v != 1 {
			t.Fatal("AsInt64 not get value as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: ':', intlen: 2}}).AsInt64(); v != 2 {
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
		if v, _ := (RedisResult{val: strmsg('+', "1")}).AsUint64(); v != 1 {
			t.Fatal("AsUint64 not get value as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: ':', intlen: 2}}).AsUint64(); v != 2 {
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
		if v, _ := (RedisResult{val: strmsg('+', "1.1")}).AsFloat64(); v != 1.1 {
			t.Fatal("AsFloat64 not get value as expected")
		}
		if v, _ := (RedisResult{val: strmsg(',', "2.2")}).AsFloat64(); v != 2.2 {
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
		values := []RedisMessage{strmsg('+', "item")}
		if ret, _ := (RedisResult{val: slicemsg('*', values)}).ToArray(); !reflect.DeepEqual(ret, values) {
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
		values := []RedisMessage{strmsg('+', "item")}
		if ret, _ := (RedisResult{val: slicemsg('*', values)}).AsStrSlice(); !reflect.DeepEqual(ret, []string{"item"}) {
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
		values := []RedisMessage{{intlen: 2, typ: ':'}}
		if ret, _ := (RedisResult{val: slicemsg('*', values)}).AsIntSlice(); !reflect.DeepEqual(ret, []int64{2}) {
			t.Fatal("AsIntSlice not get value as expected")
		}
		values = []RedisMessage{strmsg('+', "3")}
		if ret, _ := (RedisResult{val: slicemsg('*', values)}).AsIntSlice(); !reflect.DeepEqual(ret, []int64{3}) {
			t.Fatal("AsIntSlice not get value as expected")
		}
		values = []RedisMessage{strmsg('+', "ab")}
		if _, err := (RedisResult{val: slicemsg('*', values)}).AsIntSlice(); err == nil {
			t.Fatal("AsIntSlice not failed as expected")
		}
	})

	t.Run("AsFloatSlice", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsFloatSlice(); err == nil {
			t.Fatal("AsFloatSlice not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsFloatSlice(); err == nil {
			t.Fatal("AsFloatSlice not failed as expected")
		}
		if _, err := (RedisResult{val: slicemsg('*', []RedisMessage{strmsg(',', "fff")})}).AsFloatSlice(); err == nil {
			t.Fatal("AsFloatSlice not failed as expected")
		}
		values := []RedisMessage{{intlen: 1, typ: ':'}, strmsg('+', "2"), strmsg('$', "3"), strmsg(',', "4")}
		if ret, _ := (RedisResult{val: slicemsg('*', values)}).AsFloatSlice(); !reflect.DeepEqual(ret, []float64{1, 2, 3, 4}) {
			t.Fatal("AsFloatSlice not get value as expected")
		}
	})

	t.Run("AsBoolSlice", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsBoolSlice(); err == nil {
			t.Fatal("AsBoolSlice not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsBoolSlice(); err == nil {
			t.Fatal("AsBoolSlice not failed as expected")
		}
		values := []RedisMessage{{intlen: 1, typ: ':'}, strmsg('+', "0"), {intlen: 1, typ: typeBool}}
		if ret, _ := (RedisResult{val: slicemsg('*', values)}).AsBoolSlice(); !reflect.DeepEqual(ret, []bool{true, false, true}) {
			t.Fatal("AsBoolSlice not get value as expected")
		}
	})

	t.Run("AsMap", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsMap(); err == nil {
			t.Fatal("AsMap not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsMap(); err == nil {
			t.Fatal("AsMap not failed as expected")
		}
		values := []RedisMessage{strmsg('+', "key"), strmsg('+', "value")}
		if ret, _ := (RedisResult{val: slicemsg('*', values)}).AsMap(); !reflect.DeepEqual(map[string]RedisMessage{
			values[0].string(): values[1],
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
		values := []RedisMessage{strmsg('+', "key"), strmsg('+', "value")}
		if ret, _ := (RedisResult{val: slicemsg('*', values)}).AsStrMap(); !reflect.DeepEqual(map[string]string{
			values[0].string(): values[1].string(),
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
		if _, err := (RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "key"), strmsg('+', "value")})}).AsIntMap(); err == nil {
			t.Fatal("AsIntMap not failed as expected")
		}
		values := []RedisMessage{strmsg('+', "k1"), strmsg('+', "1"), strmsg('+', "k2"), {intlen: 2, typ: ':'}}
		if ret, _ := (RedisResult{val: slicemsg('*', values)}).AsIntMap(); !reflect.DeepEqual(map[string]int64{
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
		values := []RedisMessage{strmsg('+', "key"), strmsg('+', "value")}
		if ret, _ := (RedisResult{val: slicemsg('%', values)}).ToMap(); !reflect.DeepEqual(map[string]RedisMessage{
			values[0].string(): values[1],
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
		redisErr := RedisError(strmsg('-', "err"))
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('%', []RedisMessage{strmsg('+', "key"), {typ: ':', intlen: 1}}),
			slicemsg('%', []RedisMessage{strmsg('+', "nil"), {typ: '_'}}),
			slicemsg('%', []RedisMessage{strmsg('+', "err"), strmsg('-', "err")}),
			strmsg(',', "1.2"),
			strmsg('+', "str"),
			{typ: '#', intlen: 0},
			strmsg('-', "err"),
			{typ: '_'},
		})}).ToAny(); !reflect.DeepEqual([]any{
			map[string]any{"key": int64(1)},
			map[string]any{"nil": nil},
			map[string]any{"err": &redisErr},
			1.2,
			"str",
			false,
			&redisErr,
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
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "id"), slicemsg('*', []RedisMessage{strmsg('+', "a"), strmsg('+', "b")})})}).AsXRangeEntry(); !reflect.DeepEqual(XRangeEntry{
			ID:          "id",
			FieldValues: map[string]string{"a": "b"},
		}, ret) {
			t.Fatal("AsXRangeEntry not get value as expected")
		}
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "id"), {typ: '_'}})}).AsXRangeEntry(); !reflect.DeepEqual(XRangeEntry{
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
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{strmsg('+', "id1"), slicemsg('*', []RedisMessage{strmsg('+', "a"), strmsg('+', "b")})}),
			slicemsg('*', []RedisMessage{strmsg('+', "id2"), {typ: '_'}}),
		})}).AsXRange(); !reflect.DeepEqual([]XRangeEntry{{
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
		if ret, _ := (RedisResult{val: slicemsg('%', []RedisMessage{
			strmsg('+', "stream1"),
			slicemsg('*', []RedisMessage{
				slicemsg('*', []RedisMessage{strmsg('+', "id1"), slicemsg('*', []RedisMessage{strmsg('+', "a"), strmsg('+', "b")})}),
				slicemsg('*', []RedisMessage{strmsg('+', "id2"), {typ: '_'}}),
			}),
			strmsg('+', "stream2"),
			slicemsg('*', []RedisMessage{
				slicemsg('*', []RedisMessage{strmsg('+', "id3"), slicemsg('*', []RedisMessage{strmsg('+', "c"), strmsg('+', "d")})}),
			}),
		})}).AsXRead(); !reflect.DeepEqual(map[string][]XRangeEntry{
			"stream1": {
				{ID: "id1", FieldValues: map[string]string{"a": "b"}},
				{ID: "id2", FieldValues: nil}},
			"stream2": {
				{ID: "id3", FieldValues: map[string]string{"c": "d"}},
			},
		}, ret) {
			t.Fatal("AsXRead not get value as expected")
		}
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "stream1"),
				slicemsg('*', []RedisMessage{
					slicemsg('*', []RedisMessage{strmsg('+', "id1"), slicemsg('*', []RedisMessage{strmsg('+', "a"), strmsg('+', "b")})}),
					slicemsg('*', []RedisMessage{strmsg('+', "id2"), {typ: '_'}}),
				}),
			}),
			slicemsg('*', []RedisMessage{
				strmsg('+', "stream2"),
				slicemsg('*', []RedisMessage{
					slicemsg('*', []RedisMessage{strmsg('+', "id3"), slicemsg('*', []RedisMessage{strmsg('+', "c"), strmsg('+', "d")})}),
				}),
			}),
		})}).AsXRead(); !reflect.DeepEqual(map[string][]XRangeEntry{
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
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			strmsg('+', "m1"),
			strmsg('+', "1"),
		})}).AsZScore(); !reflect.DeepEqual(ZScore{Member: "m1", Score: 1}, ret) {
			t.Fatal("AsZScore not get value as expected")
		}
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			strmsg('+', "m1"),
			strmsg(',', "1"),
		})}).AsZScore(); !reflect.DeepEqual(ZScore{Member: "m1", Score: 1}, ret) {
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
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			strmsg('+', "m1"),
			strmsg('+', "1"),
			strmsg('+', "m2"),
			strmsg('+', "2"),
		})}).AsZScores(); !reflect.DeepEqual([]ZScore{
			{Member: "m1", Score: 1},
			{Member: "m2", Score: 2},
		}, ret) {
			t.Fatal("AsZScores not get value as expected")
		}
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "m1"),
				strmsg(',', "1"),
			}),
			slicemsg('*', []RedisMessage{
				strmsg('+', "m2"),
				strmsg(',', "2"),
			}),
		})}).AsZScores(); !reflect.DeepEqual([]ZScore{
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
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			strmsg('+', "k"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "1"),
				strmsg('+', "2"),
			}),
		})}).AsLMPop(); !reflect.DeepEqual(KeyValues{
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
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			strmsg('+', "k"),
			slicemsg('*', []RedisMessage{
				slicemsg('*', []RedisMessage{
					strmsg('+', "1"),
					strmsg(',', "1"),
				}),
				slicemsg('*', []RedisMessage{
					strmsg('+', "2"),
					strmsg(',', "2"),
				}),
			}),
		})}).AsZMPop(); !reflect.DeepEqual(KeyZScores{
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
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
			strmsg('+', "a"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "k1"),
				strmsg('+', "v1"),
				strmsg('+', "kk"),
				strmsg('+', "vv"),
			}),
			strmsg('+', "b"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "k2"),
				strmsg('+', "v2"),
				strmsg('+', "kk"),
				strmsg('+', "vv"),
			}),
		})}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "a", Doc: map[string]string{"k1": "v1", "kk": "vv"}},
			{Key: "b", Doc: map[string]string{"k2": "v2", "kk": "vv"}},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
			strmsg('+', "a"),
			strmsg('+', "1"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "k1"),
				strmsg('+', "v1"),
			}),
			strmsg('+', "b"),
			strmsg('+', "2"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "k2"),
				strmsg('+', "v2"),
			}),
		})}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "a", Doc: map[string]string{"k1": "v1"}, Score: 1},
			{Key: "b", Doc: map[string]string{"k2": "v2"}, Score: 2},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
			strmsg('+', "a"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "k1"),
				strmsg('+', "v1"),
				strmsg('+', "kk"),
				strmsg('+', "vv"),
			}),
		})}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "a", Doc: map[string]string{"k1": "v1", "kk": "vv"}},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
			strmsg('+', "a"),
			strmsg('+', "b"),
		})}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "a", Doc: nil},
			{Key: "b", Doc: nil},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
			strmsg('+', "a"),
			strmsg('+', "1"),
			strmsg('+', "b"),
			strmsg('+', "2"),
		})}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "a", Doc: nil, Score: 1},
			{Key: "b", Doc: nil, Score: 2},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
			strmsg('+', "1"),
			strmsg('+', "2"),
		})}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "1", Doc: nil},
			{Key: "2", Doc: nil},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
			strmsg('+', "a"),
		})}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "a", Doc: nil},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
		})}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
	})

	t.Run("AsFtSearch RESP3", func(t *testing.T) {
		if n, ret, _ := (RedisResult{val: slicemsg('%', []RedisMessage{
			strmsg('+', "total_results"),
			{typ: ':', intlen: 3},
			strmsg('+', "results"),
			slicemsg('*', []RedisMessage{
				slicemsg('%', []RedisMessage{
					strmsg('+', "id"),
					strmsg('+', "1"),
					strmsg('+', "extra_attributes"),
					slicemsg('%', []RedisMessage{
						strmsg('+', "$"),
						strmsg('+', "1"),
					}),
					strmsg('+', "score"),
					strmsg(',', "1"),
				}),
				slicemsg('%', []RedisMessage{
					strmsg('+', "id"),
					strmsg('+', "2"),
					strmsg('+', "extra_attributes"),
					slicemsg('%', []RedisMessage{
						strmsg('+', "$"),
						strmsg('+', "2"),
					}),
					strmsg('+', "score"),
					strmsg(',', "2"),
				}),
			}),
			strmsg('+', "error"),
			slicemsg('*', []RedisMessage{}),
		})}).AsFtSearch(); n != 3 || !reflect.DeepEqual([]FtSearchDoc{
			{Key: "1", Doc: map[string]string{"$": "1"}, Score: 1},
			{Key: "2", Doc: map[string]string{"$": "2"}, Score: 2},
		}, ret) {
			t.Fatal("AsFtSearch not get value as expected")
		}
		if _, _, err := (RedisResult{val: slicemsg('%', []RedisMessage{
			strmsg('+', "total_results"),
			{typ: ':', intlen: 3},
			strmsg('+', "results"),
			slicemsg('*', []RedisMessage{
				slicemsg('%', []RedisMessage{
					strmsg('+', "id"),
					strmsg('+', "1"),
					strmsg('+', "extra_attributes"),
					slicemsg('%', []RedisMessage{
						strmsg('+', "$"),
						strmsg('+', "1"),
					}),
					strmsg('+', "score"),
					strmsg(',', "1"),
				}),
			}),
			strmsg('+', "error"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "mytimeout"),
			}),
		})}).AsFtSearch(); err == nil || err.Error() != "mytimeout" {
			t.Fatal("AsFtSearch not get value as expected")
		}
	})

	t.Run("AsFtAggregate", func(t *testing.T) {
		if _, _, err := (RedisResult{err: errors.New("other")}).AsFtAggregate(); err == nil {
			t.Fatal("AsFtAggregate not failed as expected")
		}
		if _, _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsFtAggregate(); err == nil {
			t.Fatal("AsFtAggregate not failed as expected")
		}
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
			slicemsg('*', []RedisMessage{
				strmsg('+', "k1"),
				strmsg('+', "v1"),
				strmsg('+', "kk"),
				strmsg('+', "vv"),
			}),
			slicemsg('*', []RedisMessage{
				strmsg('+', "k2"),
				strmsg('+', "v2"),
				strmsg('+', "kk"),
				strmsg('+', "vv"),
			}),
		})}).AsFtAggregate(); n != 3 || !reflect.DeepEqual([]map[string]string{
			{"k1": "v1", "kk": "vv"},
			{"k2": "v2", "kk": "vv"},
		}, ret) {
			t.Fatal("AsFtAggregate not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
			slicemsg('*', []RedisMessage{
				strmsg('+', "k1"),
				strmsg('+', "v1"),
				strmsg('+', "kk"),
				strmsg('+', "vv"),
			}),
		})}).AsFtAggregate(); n != 3 || !reflect.DeepEqual([]map[string]string{
			{"k1": "v1", "kk": "vv"},
		}, ret) {
			t.Fatal("AsFtAggregate not get value as expected")
		}
		if n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			{typ: ':', intlen: 3},
		})}).AsFtAggregate(); n != 3 || !reflect.DeepEqual([]map[string]string{}, ret) {
			t.Fatal("AsFtAggregate not get value as expected")
		}
	})

	t.Run("AsFtAggregate RESP3", func(t *testing.T) {
		if n, ret, _ := (RedisResult{val: slicemsg('%', []RedisMessage{
			strmsg('+', "total_results"),
			{typ: ':', intlen: 3},
			strmsg('+', "results"),
			slicemsg('*', []RedisMessage{
				slicemsg('%', []RedisMessage{
					strmsg('+', "extra_attributes"),
					slicemsg('%', []RedisMessage{
						strmsg('+', "$"),
						strmsg('+', "1"),
					}),
				}),
				slicemsg('%', []RedisMessage{
					strmsg('+', "extra_attributes"),
					slicemsg('%', []RedisMessage{
						strmsg('+', "$"),
						strmsg('+', "2"),
					}),
				}),
			}),
			strmsg('+', "error"),
			slicemsg('*', []RedisMessage{}),
		})}).AsFtAggregate(); n != 3 || !reflect.DeepEqual([]map[string]string{
			{"$": "1"},
			{"$": "2"},
		}, ret) {
			t.Fatal("AsFtAggregate not get value as expected")
		}
		if _, _, err := (RedisResult{val: slicemsg('%', []RedisMessage{
			strmsg('+', "total_results"),
			{typ: ':', intlen: 3},
			strmsg('+', "results"),
			slicemsg('*', []RedisMessage{
				slicemsg('%', []RedisMessage{
					strmsg('+', "extra_attributes"),
					slicemsg('%', []RedisMessage{
						strmsg('+', "$"),
						strmsg('+', "1"),
					}),
				}),
			}),
			strmsg('+', "error"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "mytimeout"),
			}),
		})}).AsFtAggregate(); err == nil || err.Error() != "mytimeout" {
			t.Fatal("AsFtAggregate not get value as expected")
		}
	})

	t.Run("AsFtAggregate Cursor", func(t *testing.T) {
		if _, _, _, err := (RedisResult{err: errors.New("other")}).AsFtAggregateCursor(); err == nil {
			t.Fatal("AsFtAggregate not failed as expected")
		}
		if _, _, _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsFtAggregateCursor(); err == nil {
			t.Fatal("AsFtAggregate not failed as expected")
		}
		if c, n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				{typ: ':', intlen: 3},
				slicemsg('*', []RedisMessage{
					strmsg('+', "k1"),
					strmsg('+', "v1"),
					strmsg('+', "kk"),
					strmsg('+', "vv"),
				}),
				slicemsg('*', []RedisMessage{
					strmsg('+', "k2"),
					strmsg('+', "v2"),
					strmsg('+', "kk"),
					strmsg('+', "vv"),
				}),
			}),
			{typ: ':', intlen: 1},
		})}).AsFtAggregateCursor(); c != 1 || n != 3 || !reflect.DeepEqual([]map[string]string{
			{"k1": "v1", "kk": "vv"},
			{"k2": "v2", "kk": "vv"},
		}, ret) {
			t.Fatal("AsFtAggregate not get value as expected")
		}
		if c, n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				{typ: ':', intlen: 3},
				slicemsg('*', []RedisMessage{
					strmsg('+', "k1"),
					strmsg('+', "v1"),
					strmsg('+', "kk"),
					strmsg('+', "vv"),
				}),
			}),
			{typ: ':', intlen: 1},
		})}).AsFtAggregateCursor(); c != 1 || n != 3 || !reflect.DeepEqual([]map[string]string{
			{"k1": "v1", "kk": "vv"},
		}, ret) {
			t.Fatal("AsFtAggregate not get value as expected")
		}
		if c, n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				{typ: ':', intlen: 3},
			}),
			{typ: ':', intlen: 1},
		})}).AsFtAggregateCursor(); c != 1 || n != 3 || !reflect.DeepEqual([]map[string]string{}, ret) {
			t.Fatal("AsFtAggregate not get value as expected")
		}
	})

	t.Run("AsFtAggregate Cursor RESP3", func(t *testing.T) {
		if c, n, ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('%', []RedisMessage{
				strmsg('+', "total_results"),
				{typ: ':', intlen: 3},
				strmsg('+', "results"),
				slicemsg('*', []RedisMessage{
					slicemsg('%', []RedisMessage{
						strmsg('+', "extra_attributes"),
						slicemsg('%', []RedisMessage{
							strmsg('+', "$"),
							strmsg('+', "1"),
						}),
					}),
					slicemsg('%', []RedisMessage{
						strmsg('+', "extra_attributes"),
						slicemsg('%', []RedisMessage{
							strmsg('+', "$"),
							strmsg('+', "2"),
						}),
					}),
				}),
				strmsg('+', "error"),
				slicemsg('*', []RedisMessage{}),
			}),
			{typ: ':', intlen: 1},
		})}).AsFtAggregateCursor(); c != 1 || n != 3 || !reflect.DeepEqual([]map[string]string{
			{"$": "1"},
			{"$": "2"},
		}, ret) {
			t.Fatal("AsFtAggregate not get value as expected")
		}
		if _, _, _, err := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('%', []RedisMessage{
				strmsg('+', "total_results"),
				{typ: ':', intlen: 3},
				strmsg('+', "results"),
				slicemsg('*', []RedisMessage{
					slicemsg('%', []RedisMessage{
						strmsg('+', "extra_attributes"),
						slicemsg('%', []RedisMessage{
							strmsg('+', "$"),
							strmsg('+', "1"),
						}),
					}),
				}),
				strmsg('+', "error"),
				slicemsg('*', []RedisMessage{
					strmsg('+', "mytimeout"),
				}),
			}),
			{typ: ':', intlen: 1},
		})}).AsFtAggregateCursor(); err == nil || err.Error() != "mytimeout" {
			t.Fatal("AsFtAggregate not get value as expected")
		}
	})

	t.Run("AsGeosearch", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsGeosearch(); err == nil {
			t.Fatal("AsGeosearch not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsGeosearch(); err == nil {
			t.Fatal("AsGeosearch not failed as expected")
		}
		//WithDist, WithHash, WithCoord
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('$', "k1"),
				strmsg(',', "2.5"),
				{typ: ':', intlen: 1},
				slicemsg('*', []RedisMessage{
					strmsg(',', "28.0473"),
					strmsg(',', "26.2041"),
				}),
			}),
			slicemsg('*', []RedisMessage{
				strmsg('$', "k2"),
				strmsg(',', "4.5"),
				{typ: ':', intlen: 4},
				slicemsg('*', []RedisMessage{
					strmsg(',', "72.4973"),
					strmsg(',', "13.2263"),
				}),
			}),
		})}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", Dist: 2.5, GeoHash: 1, Longitude: 28.0473, Latitude: 26.2041},
			{Name: "k2", Dist: 4.5, GeoHash: 4, Longitude: 72.4973, Latitude: 13.2263},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//WithHash, WithCoord
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('$', "k1"),
				{typ: ':', intlen: 1},
				slicemsg('*', []RedisMessage{
					strmsg(',', "84.3877"),
					strmsg(',', "33.7488"),
				}),
			}),
			slicemsg('*', []RedisMessage{
				strmsg('$', "k2"),
				{typ: ':', intlen: 4},
				slicemsg('*', []RedisMessage{
					strmsg(',', "115.8613"),
					strmsg(',', "31.9523"),
				}),
			}),
		})}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", GeoHash: 1, Longitude: 84.3877, Latitude: 33.7488},
			{Name: "k2", GeoHash: 4, Longitude: 115.8613, Latitude: 31.9523},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//WithDist, WithCoord
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('$', "k1"),
				strmsg(',', "2.50076"),
				slicemsg('*', []RedisMessage{
					strmsg(',', "84.3877"),
					strmsg(',', "33.7488"),
				}),
			}),
			slicemsg('*', []RedisMessage{
				strmsg('$', "k2"),
				strmsg(',', "1024.96"),
				slicemsg('*', []RedisMessage{
					strmsg(',', "115.8613"),
					strmsg(',', "31.9523"),
				}),
			}),
		})}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", Dist: 2.50076, Longitude: 84.3877, Latitude: 33.7488},
			{Name: "k2", Dist: 1024.96, Longitude: 115.8613, Latitude: 31.9523},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//WithCoord
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('$', "k1"),
				slicemsg('*', []RedisMessage{
					strmsg(',', "122.4194"),
					strmsg(',', "37.7749"),
				}),
			}),
			slicemsg('*', []RedisMessage{
				strmsg('$', "k2"),
				slicemsg('*', []RedisMessage{
					strmsg(',', "35.6762"),
					strmsg(',', "139.6503"),
				}),
			}),
		})}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", Longitude: 122.4194, Latitude: 37.7749},
			{Name: "k2", Longitude: 35.6762, Latitude: 139.6503},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//WithDist
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('$', "k1"),
				strmsg(',', "2.50076"),
			}),
			slicemsg('*', []RedisMessage{
				strmsg('$', "k2"),
				strmsg(',', strconv.FormatFloat(math.MaxFloat64, 'E', -1, 64)),
			}),
		})}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", Dist: 2.50076},
			{Name: "k2", Dist: math.MaxFloat64},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//WithHash
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('$', "k1"),
				{typ: ':', intlen: math.MaxInt64},
			}),
			slicemsg('*', []RedisMessage{
				strmsg('$', "k2"),
				{typ: ':', intlen: 22296},
			}),
		})}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1", GeoHash: math.MaxInt64},
			{Name: "k2", GeoHash: 22296},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//With no additional options
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{
			strmsg('$', "k1"),
			strmsg('$', "k2"),
		})}).AsGeosearch(); !reflect.DeepEqual([]GeoLocation{
			{Name: "k1"},
			{Name: "k2"},
		}, ret) {
			t.Fatal("AsGeosearch not get value as expected")
		}
		//With wrong distance
		if _, err := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('$', "k1"),
				strmsg(',', "wrong distance"),
			}),
		})}).AsGeosearch(); err == nil {
			t.Fatal("AsGeosearch not failed as expected")
		}
		//With wrong coordinates
		if _, err := (RedisResult{val: slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('$', "k2"),
				slicemsg('*', []RedisMessage{
					strmsg(',', "35.6762"),
				}),
			}),
		})}).AsGeosearch(); err == nil {
			t.Fatal("AsGeosearch not failed as expected")
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

	t.Run("Stringer", func(t *testing.T) {
		tests := []struct {
			expected string
			input    RedisResult
		}{
			{
				input: RedisResult{
					val: slicemsg('*', []RedisMessage{
						slicemsg('*', []RedisMessage{
							{typ: ':', intlen: 0},
							{typ: ':', intlen: 0},
							slicemsg('*', []RedisMessage{ // master
								strmsg('+', "127.0.3.1"),
								{typ: ':', intlen: 3},
								strmsg('+', ""),
							}),
						}),
					}),
				},
				expected: `{"Message":{"Value":[{"Value":[{"Value":0,"Type":"int64"},{"Value":0,"Type":"int64"},{"Value":[{"Value":"127.0.3.1","Type":"simple string"},{"Value":3,"Type":"int64"},{"Value":"","Type":"simple string"}],"Type":"array"}],"Type":"array"}],"Type":"array"}}`,
			},
			{
				input:    RedisResult{err: errors.New("foo")},
				expected: `{"Error":"foo"}`,
			},
		}
		for _, test := range tests {
			msg := test.input.String()
			if msg != test.expected {
				t.Fatalf("unexpected string. got %v expected %v", msg, test.expected)
			}
		}
	})
}

//gocyclo:ignore
func TestRedisMessage(t *testing.T) {
	//Add erroneous type
	typeNames['t'] = "t"

	t.Run("IsNil", func(t *testing.T) {
		if !(&RedisMessage{typ: '_'}).IsNil() {
			t.Fatal("IsNil fail")
		}
	})
	t.Run("Trim ERR prefix", func(t *testing.T) {
		// kvrocks: https://github.com/redis/rueidis/issues/152#issuecomment-1333923750
		redisMessageError := strmsg('-', "ERR no_prefix")
		if (&redisMessageError).Error().Error() != "no_prefix" {
			t.Fatal("fail to trim ERR")
		}
	})
	t.Run("ToInt64", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 int64
		if val, err := (&RedisMessage{typ: '_'}).ToInt64(); err == nil {
			t.Fatal("ToInt64 did not fail as expected")
		} else if val != 0 {
			t.Fatalf("expected 0, got %d", val)
		}

		// Test case where the message type is 't', which is not a RESP3 int64
		if val, err := (&RedisMessage{typ: 't'}).ToInt64(); err == nil {
			t.Fatal("ToInt64 did not fail as expected")
		} else if val != 0 {
			t.Fatalf("expected 0, got %d", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a RESP3 int64") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("ToBool", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 bool
		if val, err := (&RedisMessage{typ: '_'}).ToBool(); err == nil {
			t.Fatal("ToBool did not fail as expected")
		} else if val != false {
			t.Fatalf("expected false, got %v", val)
		}

		// Test case where the message type is 't', which is not a RESP3 bool
		if val, err := (&RedisMessage{typ: 't'}).ToBool(); err == nil {
			t.Fatal("ToBool did not fail as expected")
		} else if val != false {
			t.Fatalf("expected false, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a RESP3 bool") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsBool", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 int, string, or bool
		if val, err := (&RedisMessage{typ: '_'}).AsBool(); err == nil {
			t.Fatal("AsBool did not fail as expected")
		} else if val != false {
			t.Fatalf("expected false, got %v", val)
		}

		// Test case where the message type is 't', which is not a RESP3 int, string, or bool
		if val, err := (&RedisMessage{typ: 't'}).AsBool(); err == nil {
			t.Fatal("AsBool did not fail as expected")
		} else if val != false {
			t.Fatalf("expected false, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a int, string or bool") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("ToFloat64", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 float64
		if val, err := (&RedisMessage{typ: '_'}).ToFloat64(); err == nil {
			t.Fatal("ToFloat64 did not fail as expected")
		} else if val != 0 {
			t.Fatalf("expected 0, got %f", val)
		}

		// Test case where the message type is 't', which is not a RESP3 float64
		if val, err := (&RedisMessage{typ: 't'}).ToFloat64(); err == nil {
			t.Fatal("ToFloat64 did not fail as expected")
		} else if val != 0 {
			t.Fatalf("expected 0, got %f", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a RESP3 float64") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("ToString", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 string
		if val, err := (&RedisMessage{typ: '_'}).ToString(); err == nil {
			t.Fatal("ToString did not fail as expected")
		} else if val != "" {
			t.Fatalf("expected empty string, got %v", val)
		}

		// Test case where the message type is ':', which is not a RESP3 string
		if val, err := (&RedisMessage{typ: ':'}).ToString(); err == nil {
			t.Fatal("ToString did not fail as expected")
		} else if val != "" {
			t.Fatalf("expected empty string, got %v", val)
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a string", typeNames[':'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsReader", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 string
		if val, err := (&RedisMessage{typ: '_'}).AsReader(); err == nil {
			t.Fatal("AsReader did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		// Test case where the message type is ':', which is not a RESP3 string
		if val, err := (&RedisMessage{typ: ':'}).AsReader(); err == nil {
			t.Fatal("AsReader did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a string", typeNames[':'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsBytes", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 string
		if val, err := (&RedisMessage{typ: '_'}).AsBytes(); err == nil {
			t.Fatal("AsBytes did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		// Test case where the message type is ':', which is not a RESP3 string
		if val, err := (&RedisMessage{typ: ':'}).AsBytes(); err == nil {
			t.Fatal("AsBytes did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a string", typeNames[':'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("DecodeJSON", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 string
		if err := (&RedisMessage{typ: '_'}).DecodeJSON(nil); err == nil {
			t.Fatal("DecodeJSON did not fail as expected")
		}

		// Test case where the message type is ':', which is not a RESP3 string
		if err := (&RedisMessage{typ: ':'}).DecodeJSON(nil); err == nil {
			t.Fatal("DecodeJSON did not fail as expected")
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a string", typeNames[':'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsInt64", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 string
		if val, err := (&RedisMessage{typ: '_'}).AsInt64(); err == nil {
			t.Fatal("AsInt64 did not fail as expected")
		} else if val != 0 {
			t.Fatalf("expected 0, got %d", val)
		}

		// Test case where the message type is '*', which is not a RESP3 string
		redisMessageArrayWithEmptyMessage := slicemsg('*', []RedisMessage{{}})
		if val, err := (&redisMessageArrayWithEmptyMessage).AsInt64(); err == nil {
			t.Fatal("AsInt64 did not fail as expected")
		} else if val != 0 {
			t.Fatalf("expected 0, got %d", val)
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a string", typeNames['*'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsUint64", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 string
		if val, err := (&RedisMessage{typ: '_'}).AsUint64(); err == nil {
			t.Fatal("AsUint64 did not fail as expected")
		} else if val != 0 {
			t.Fatalf("expected 0, got %d", val)
		}

		// Test case where the message type is '*', which is not a RESP3 string
		redisMessageArrayWithEmptyMessage := slicemsg('*', []RedisMessage{{}})
		if val, err := (&redisMessageArrayWithEmptyMessage).AsUint64(); err == nil {
			t.Fatal("AsUint64 did not fail as expected")
		} else if val != 0 {
			t.Fatalf("expected 0, got %d", val)
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a string", typeNames['*'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsFloat64", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 string
		if val, err := (&RedisMessage{typ: '_'}).AsFloat64(); err == nil {
			t.Fatal("AsFloat64 did not fail as expected")
		} else if val != 0 {
			t.Fatalf("expected 0, got %f", val)
		}

		// Test case where the message type is ':', which is not a RESP3 string
		if val, err := (&RedisMessage{typ: ':'}).AsFloat64(); err == nil {
			t.Fatal("AsFloat64 did not fail as expected")
		} else if val != 0 {
			t.Fatalf("expected 0, got %f", val)
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a string", typeNames[':'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("ToArray", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 array
		if val, err := (&RedisMessage{typ: '_'}).ToArray(); err == nil {
			t.Fatal("ToArray did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		// Test case where the message type is 't', which is not a RESP3 array
		if val, err := (&RedisMessage{typ: 't'}).ToArray(); err == nil {
			t.Fatal("ToArray did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a array") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsStrSlice", func(t *testing.T) {
		// Test case where the message type is '_', which is not a RESP3 array
		if val, err := (&RedisMessage{typ: '_'}).AsStrSlice(); err == nil {
			t.Fatal("AsStrSlice did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		// Test case where the message type is 't', which is not a RESP3 array
		if val, err := (&RedisMessage{typ: 't'}).AsStrSlice(); err == nil {
			t.Fatal("AsStrSlice did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a array") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsIntSlice", func(t *testing.T) {
		if val, err := (&RedisMessage{typ: '_'}).AsIntSlice(); err == nil {
			t.Fatal("AsIntSlice did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		if val, err := (&RedisMessage{typ: 't'}).AsIntSlice(); err == nil {
			t.Fatal("AsIntSlice did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a array") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsFloatSlice", func(t *testing.T) {
		if val, err := (&RedisMessage{typ: '_'}).AsFloatSlice(); err == nil {
			t.Fatal("AsFloatSlice did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		if val, err := (&RedisMessage{typ: 't'}).AsFloatSlice(); err == nil {
			t.Fatal("AsFloatSlice did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a array") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsBoolSlice", func(t *testing.T) {
		if val, err := (&RedisMessage{typ: '_'}).AsBoolSlice(); err == nil {
			t.Fatal("AsBoolSlice did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		if val, err := (&RedisMessage{typ: 't'}).AsBoolSlice(); err == nil {
			t.Fatal("AsBoolSlice did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a array") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsMap", func(t *testing.T) {
		if val, err := (&RedisMessage{typ: '_'}).AsMap(); err == nil {
			t.Fatal("AsMap did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		if val, err := (&RedisMessage{typ: 't'}).AsMap(); err == nil {
			t.Fatal("AsMap did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a map/array/set or its length is not even") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsStrMap", func(t *testing.T) {
		if val, err := (&RedisMessage{typ: '_'}).AsStrMap(); err == nil {
			t.Fatal("AsStrMap did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		if val, err := (&RedisMessage{typ: 't'}).AsStrMap(); err == nil {
			t.Fatal("AsStrMap did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a map/array/set") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsIntMap", func(t *testing.T) {
		if val, err := (&RedisMessage{typ: '_'}).AsIntMap(); err == nil {
			t.Fatal("AsIntMap did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		if val, err := (&RedisMessage{typ: 't'}).AsIntMap(); err == nil {
			t.Fatal("AsIntMap did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a map/array/set") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("ToMap", func(t *testing.T) {
		if val, err := (&RedisMessage{typ: '_'}).ToMap(); err == nil {
			t.Fatal("ToMap did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		if val, err := (&RedisMessage{typ: 't'}).ToMap(); err == nil {
			t.Fatal("ToMap did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a RESP3 map") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("ToAny", func(t *testing.T) {
		if val, err := (&RedisMessage{typ: '_'}).ToAny(); err == nil {
			t.Fatal("ToAny did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		}

		if val, err := (&RedisMessage{typ: 't'}).ToAny(); err == nil {
			t.Fatal("ToAny did not fail as expected")
		} else if val != nil {
			t.Fatalf("expected nil, got %v", val)
		} else if !strings.Contains(err.Error(), "redis message type t is not a supported in ToAny") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsXRangeEntry - no range id", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry did not fail as expected")
		}

		if _, err := (&RedisMessage{typ: '*'}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry did not fail as expected")
		}

		redisMessageNullAndMap := slicemsg('*', []RedisMessage{{typ: '_'}, {typ: '%'}})
		if _, err := (&redisMessageNullAndMap).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry did not fail as expected")
		}

		redisMessageIntAndMap := slicemsg('*', []RedisMessage{{typ: ':'}, {typ: '%'}})
		if _, err := (&redisMessageIntAndMap).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry did not fail as expected")
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a string", typeNames[':'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsXRangeEntry - no range field values", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry did not fail as expected")
		}

		if _, err := (&RedisMessage{typ: '*'}).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry did not fail as expected")
		}

		redisMessageStringAndErr := slicemsg('*', []RedisMessage{{typ: '+'}, {typ: '-'}})
		if _, err := (&redisMessageStringAndErr).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry did not fail as expected")
		}

		redisMessageStringAndUnknown := slicemsg('*', []RedisMessage{{typ: '+'}, {typ: 't'}})
		if _, err := (&redisMessageStringAndUnknown).AsXRangeEntry(); err == nil {
			t.Fatal("AsXRangeEntry did not fail as expected")
		} else if !strings.Contains(err.Error(), "redis message type t is not a map/array/set") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsXRange", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}

		redisMessageArrayWithNull := slicemsg('*', []RedisMessage{{typ: '_'}})
		if _, err := (&redisMessageArrayWithNull).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}
	})

	t.Run("AsXRead", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRead(); err == nil {
			t.Fatal("AsXRead did not fail as expected")
		}

		redisMessageMapIncorrectLen := slicemsg('%', []RedisMessage{
			strmsg('+', "stream1"),
			slicemsg('*', []RedisMessage{slicemsg('*', []RedisMessage{strmsg('+', "id1")})}),
		})
		if _, err := (&redisMessageMapIncorrectLen).AsXRead(); err == nil {
			t.Fatal("AsXRead did not fail as expected")
		}

		redisMessageArrayIncorrectLen := slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "stream1"),
			}),
		})
		if _, err := (&redisMessageArrayIncorrectLen).AsXRead(); err == nil {
			t.Fatal("AsXRead did not fail as expected")
		}

		redisMessageArrayIncorrectLen2 := slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "stream1"),
				slicemsg('*', []RedisMessage{slicemsg('*', []RedisMessage{strmsg('+', "id1")})}),
			}),
		})
		if _, err := (&redisMessageArrayIncorrectLen2).AsXRead(); err == nil {
			t.Fatal("AsXRead did not fail as expected")
		}

		if _, err := (&RedisMessage{typ: 't'}).AsXRead(); err == nil {
			t.Fatal("AsXRead did not fail as expected")
		} else if !strings.Contains(err.Error(), "redis message type t is not a map/array/set") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsZScore", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsZScore(); err == nil {
			t.Fatal("AsZScore did not fail as expected")
		}

		if _, err := (&RedisMessage{typ: '*'}).AsZScore(); err == nil {
			t.Fatal("AsZScore did not fail as expected")
		} else if !strings.Contains(err.Error(), "redis message is not a map/array/set or its length is not 2") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsZScores", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsZScores(); err == nil {
			t.Fatal("AsZScore not failed as expected")
		}
		redisMessageStringArray := slicemsg('*', []RedisMessage{
			strmsg('+', "m1"),
			strmsg('+', "m2"),
		})
		if _, err := (&redisMessageStringArray).AsZScores(); err == nil {
			t.Fatal("AsZScores not fails as expected")
		}
		redisMessageNestedStringArray := slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "m1"),
				strmsg('+', "m2"),
			}),
		})
		if _, err := (&redisMessageNestedStringArray).AsZScores(); err == nil {
			t.Fatal("AsZScores not fails as expected")
		}
	})

	t.Run("AsLMPop", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsLMPop(); err == nil {
			t.Fatal("AsLMPop did not fail as expected")
		}

		redisMessageStringAndNull := slicemsg('*', []RedisMessage{
			strmsg('+', "k"),
			{typ: '_'},
		})
		if _, err := (&redisMessageStringAndNull).AsLMPop(); err == nil {
			t.Fatal("AsLMPop did not fail as expected")
		}

		redisMessageStringArray := slicemsg('*', []RedisMessage{
			strmsg('+', "k"),
		})
		if _, err := (&redisMessageStringArray).AsLMPop(); err == nil {
			t.Fatal("AsLMPop did not fail as expected")
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a LMPOP response", typeNames['*'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsZMPop", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsZMPop(); err == nil {
			t.Fatal("AsZMPop did not fail as expected")
		}

		redisMessageStringAndNull := slicemsg('*', []RedisMessage{
			strmsg('+', "k"),
			{typ: '_'},
		})
		if _, err := (&redisMessageStringAndNull).AsZMPop(); err == nil {
			t.Fatal("AsZMPop did not fail as expected")
		}

		redisMessageStringArray := slicemsg('*', []RedisMessage{
			strmsg('+', "k"),
		})
		if _, err := (&redisMessageStringArray).AsZMPop(); err == nil {
			t.Fatal("AsZMPop did not fail as expected")
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a ZMPOP response", typeNames['*'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsFtSearch", func(t *testing.T) {
		if _, _, err := (&RedisMessage{typ: '_'}).AsFtSearch(); err == nil {
			t.Fatal("AsFtSearch did not fail as expected")
		}

		if _, _, err := (&RedisMessage{typ: '*'}).AsFtSearch(); err == nil {
			t.Fatal("AsFtSearch did not fail as expected")
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a FT.SEARCH response", typeNames['*'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsFtAggregate", func(t *testing.T) {
		if _, _, err := (&RedisMessage{typ: '_'}).AsFtAggregate(); err == nil {
			t.Fatal("AsFtAggregate did not fail as expected")
		}

		if _, _, err := (&RedisMessage{typ: '*'}).AsFtAggregate(); err == nil {
			t.Fatal("AsFtAggregate did not fail as expected")
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a FT.AGGREGATE response", typeNames['*'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsFtAggregateCursor", func(t *testing.T) {
		if _, _, _, err := (&RedisMessage{typ: '_'}).AsFtAggregateCursor(); err == nil {
			t.Fatal("AsFtAggregateCursor did not fail as expected")
		}

		if _, _, _, err := (&RedisMessage{typ: '*'}).AsFtAggregateCursor(); err == nil {
			t.Fatal("AsFtAggregateCursor did not fail as expected")
		} else if !strings.Contains(err.Error(), fmt.Sprintf("redis message type %s is not a FT.AGGREGATE response", typeNames['*'])) {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("AsScanEntry", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsScanEntry(); err == nil {
			t.Fatal("AsScanEntry not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsScanEntry(); err == nil {
			t.Fatal("AsScanEntry not failed as expected")
		}
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "1"), slicemsg('*', []RedisMessage{strmsg('+', "a"), strmsg('+', "b")})})}).AsScanEntry(); !reflect.DeepEqual(ScanEntry{
			Cursor:   1,
			Elements: []string{"a", "b"},
		}, ret) {
			t.Fatal("AsScanEntry not get value as expected")
		}
		if ret, _ := (RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "0"), {typ: '_'}})}).AsScanEntry(); !reflect.DeepEqual(ScanEntry{}, ret) {
			t.Fatal("AsScanEntry not get value as expected")
		}
		if _, err := (RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "0")})}).AsScanEntry(); err == nil || !strings.Contains(err.Error(), "a scan response or its length is not at least 2") {
			t.Fatal("AsScanEntry not get value as expected")
		}
	})

	t.Run("ToMap with non-string key", func(t *testing.T) {
		redisMessageSet := slicemsg('~', []RedisMessage{{typ: ':'}, {typ: ':'}})
		_, err := (&redisMessageSet).ToMap()
		if err == nil {
			t.Fatal("ToMap did not fail as expected")
		}
		if !strings.Contains(err.Error(), "redis message type set is not a RESP3 map") {
			t.Fatalf("ToMap failed with unexpected error: %v", err)
		}
		redisMessageMap := slicemsg('%', []RedisMessage{{typ: ':'}, {typ: ':'}})
		_, err = (&redisMessageMap).ToMap()
		if err == nil {
			t.Fatal("ToMap did not fail as expected")
		}
		if !strings.Contains(err.Error(), "int64 as map key is not supported") {
			t.Fatalf("ToMap failed with unexpected error: %v", err)
		}
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

	t.Run("Stringer", func(t *testing.T) {
		tests := []struct {
			expected string
			input    RedisMessage
		}{
			{
				input: slicemsg('*', []RedisMessage{
					slicemsg('*', []RedisMessage{
						{typ: ':', intlen: 0},
						{typ: ':', intlen: 0},
						slicemsg('*', []RedisMessage{
							strmsg('+', "127.0.3.1"),
							{typ: ':', intlen: 3},
							strmsg('+', ""),
						}),
					}),
				}),
				expected: `{"Value":[{"Value":[{"Value":0,"Type":"int64"},{"Value":0,"Type":"int64"},{"Value":[{"Value":"127.0.3.1","Type":"simple string"},{"Value":3,"Type":"int64"},{"Value":"","Type":"simple string"}],"Type":"array"}],"Type":"array"}],"Type":"array"}`,
			},
			{
				input: RedisMessage{
					typ:    '+',
					bytes:  unsafe.StringData("127.0.3.1"),
					intlen: int64(len("127.0.3.1")),
					ttl:    [7]byte{97, 77, 74, 61, 138, 1, 0},
				},
				expected: `{"Value":"127.0.3.1","Type":"simple string","TTL":"2023-08-28 17:56:34.273 +0000 UTC"}`,
			},
			{
				input:    RedisMessage{typ: '0'},
				expected: `{"Type":"unknown"}`,
			},
			{
				input:    RedisMessage{typ: typeBool, intlen: 1},
				expected: `{"Value":true,"Type":"boolean"}`,
			},
			{
				input:    RedisMessage{typ: typeNull},
				expected: `{"Type":"null","Error":"redis nil message"}`,
			},
			{
				input:    strmsg(typeSimpleErr, "ERR foo"),
				expected: `{"Type":"simple error","Error":"foo"}`,
			},
			{
				input:    strmsg(typeBlobErr, "ERR foo"),
				expected: `{"Type":"blob error","Error":"foo"}`,
			},
		}
		for _, test := range tests {
			msg := test.input.String()
			if msg != test.expected {
				t.Fatalf("unexpected string. got %v expected %v", msg, test.expected)
			}
		}
	})

	t.Run("CacheMarshal and CacheUnmarshalView", func(t *testing.T) {
		m1 := RedisMessage{typ: '_'}
		m2 := strmsg('+', "random")
		m3 := RedisMessage{typ: '#', intlen: 1}
		m4 := RedisMessage{typ: ':', intlen: -1234}
		m5 := strmsg(',', "-1.5")
		m6 := slicemsg('%', nil)
		m7 := slicemsg('~', []RedisMessage{m1, m2, m3, m4, m5, m6})
		m8 := slicemsg('*', []RedisMessage{m1, m2, m3, m4, m5, m6, m7})
		msgs := []RedisMessage{m1, m2, m3, m4, m5, m6, m7, m8}
		now := time.Now()
		for i := range msgs {
			msgs[i].setExpireAt(now.Add(time.Second * time.Duration(i)).UnixMilli())
		}
		for i, m1 := range msgs {
			siz := m1.CacheSize()
			bs1 := m1.CacheMarshal(nil)
			if len(bs1) != siz {
				t.Fatal("size not match")
			}
			bs2 := m1.CacheMarshal(bs1)
			if !bytes.Equal(bs2[:siz], bs2[siz:]) {
				t.Fatal("byte not match")
			}
			var m2 RedisMessage
			if err := m2.CacheUnmarshalView(bs1); err != nil {
				t.Fatal(err)
			}
			if m1.String() != m2.String() {
				t.Fatal("content not match")
			}
			if !m2.IsCacheHit() {
				t.Fatal("should be cache hit")
			}
			if m2.CachePXAT() != now.Add(time.Second*time.Duration(i)).UnixMilli() {
				t.Fatal("should have the same ttl")
			}
			for l := 0; l < siz; l++ {
				var m3 RedisMessage
				if err := m3.CacheUnmarshalView(bs2[:l]); err != ErrCacheUnmarshal {
					t.Fatal("should fail as expected")
				}
			}
		}
	})
}

func TestRedisMessage_AsXRangeSlice(t *testing.T) {
	t.Run("normal XRange entry with field-value pairs", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			strmsg('+', "1234567890-0"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "field1"),
				strmsg('+', "value1"),
				strmsg('+', "field2"),
				strmsg('+', "value2"),
			}),
		})

		want := XRangeSlice{
			ID: "1234567890-0",
			FieldValues: []XRangeFieldValue{
				{Field: "field1", Value: "value1"},
				{Field: "field2", Value: "value2"},
			},
		}

		got, err := message.AsXRangeSlice()
		if err != nil {
			t.Fatalf("AsXRangeSlice() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXRangeSlice() = %v, want %v", got, want)
		}
	})

	t.Run("XRange entry with duplicate fields (preserves order and duplicates)", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			strmsg('+', "1747784186966-0"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "foo"),
				strmsg('+', "1"),
				strmsg('+', "foo"),
				strmsg('+', "2"),
				strmsg('+', "bar"),
				strmsg('+', "3"),
				strmsg('+', "bar"),
				strmsg('+', "4"),
			}),
		})

		want := XRangeSlice{
			ID: "1747784186966-0",
			FieldValues: []XRangeFieldValue{
				{Field: "foo", Value: "1"},
				{Field: "foo", Value: "2"},
				{Field: "bar", Value: "3"},
				{Field: "bar", Value: "4"},
			},
		}

		got, err := message.AsXRangeSlice()
		if err != nil {
			t.Fatalf("AsXRangeSlice() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXRangeSlice() = %v, want %v", got, want)
		}
	})

	t.Run("XRange entry with nil field-values", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			strmsg('+', "1234567890-2"),
			{typ: '_'},
		})

		want := XRangeSlice{
			ID:          "1234567890-2",
			FieldValues: nil,
		}

		got, err := message.AsXRangeSlice()
		if err != nil {
			t.Fatalf("AsXRangeSlice() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXRangeSlice() = %v, want %v", got, want)
		}
	})

	t.Run("XRange entry with empty field-values array", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			strmsg('+', "1234567890-3"),
			slicemsg('*', []RedisMessage{}),
		})

		want := XRangeSlice{
			ID:          "1234567890-3",
			FieldValues: []XRangeFieldValue{},
		}

		got, err := message.AsXRangeSlice()
		if err != nil {
			t.Fatalf("AsXRangeSlice() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXRangeSlice() = %v, want %v", got, want)
		}
	})

	t.Run("XRange entry with odd number of field-values (handles gracefully)", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			strmsg('+', "1234567890-4"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "field1"),
				strmsg('+', "value1"),
				strmsg('+', "field2"),
			}),
		})

		want := XRangeSlice{
			ID: "1234567890-4",
			FieldValues: []XRangeFieldValue{
				{Field: "field1", Value: "value1"},
			},
		}

		got, err := message.AsXRangeSlice()
		if err != nil {
			t.Fatalf("AsXRangeSlice() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXRangeSlice() = %v, want %v", got, want)
		}
	})

	t.Run("invalid array length", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			strmsg('+', "1234567890-0"),
		})

		_, err := message.AsXRangeSlice()
		if err == nil {
			t.Fatal("AsXRangeSlice() expected error but got none")
		}
		if !strings.Contains(err.Error(), "got 1, wanted 2") {
			t.Errorf("AsXRangeSlice() error = %v, want error containing 'got 1, wanted 2'", err)
		}
	})

	t.Run("not an array", func(t *testing.T) {
		message := strmsg('+', "not-an-array")

		_, err := message.AsXRangeSlice()
		if err == nil {
			t.Fatal("AsXRangeSlice() expected error but got none")
		}
	})

	t.Run("error response", func(t *testing.T) {
		message := RedisMessage{typ: '_'}

		_, err := message.AsXRangeSlice()
		if err == nil {
			t.Fatal("AsXRangeSlice() expected error but got none")
		}
	})

	t.Run("error in ID parsing", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			{typ: '_'}, // ID that will cause error
			slicemsg('*', []RedisMessage{}),
		})

		_, err := message.AsXRangeSlice()
		if err == nil {
			t.Fatal("AsXRangeSlice() expected error but got none")
		}
	})

	t.Run("field-values array parsing error", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			strmsg('+', "1234567890-0"),
			{typ: '-'}, // Error type for field-values
		})

		_, err := message.AsXRangeSlice()
		if err == nil {
			t.Fatal("AsXRangeSlice() expected error but got none")
		}
	})
}

func TestRedisMessage_AsXRangeSlices(t *testing.T) {
	t.Run("multiple XRange entries", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "1234567890-0"),
				slicemsg('*', []RedisMessage{
					strmsg('+', "field1"),
					strmsg('+', "value1"),
				}),
			}),
			slicemsg('*', []RedisMessage{
				strmsg('+', "1234567890-1"),
				slicemsg('*', []RedisMessage{
					strmsg('+', "field2"),
					strmsg('+', "value2"),
				}),
			}),
		})

		want := []XRangeSlice{
			{
				ID: "1234567890-0",
				FieldValues: []XRangeFieldValue{
					{Field: "field1", Value: "value1"},
				},
			},
			{
				ID: "1234567890-1",
				FieldValues: []XRangeFieldValue{
					{Field: "field2", Value: "value2"},
				},
			},
		}

		got, err := message.AsXRangeSlices()
		if err != nil {
			t.Fatalf("AsXRangeSlices() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXRangeSlices() = %v, want %v", got, want)
		}
	})

	t.Run("empty array", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{})

		want := []XRangeSlice{}
		got, err := message.AsXRangeSlices()
		if err != nil {
			t.Fatalf("AsXRangeSlices() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXRangeSlices() = %v, want %v", got, want)
		}
	})

	t.Run("not an array", func(t *testing.T) {
		message := strmsg('+', "not-an-array")

		_, err := message.AsXRangeSlices()
		if err == nil {
			t.Fatal("AsXRangeSlices() expected error but got none")
		}
	})

	t.Run("invalid entry in array", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			strmsg('+', "invalid-entry"),
		})

		_, err := message.AsXRangeSlices()
		if err == nil {
			t.Fatal("AsXRangeSlices() expected error but got none")
		}
	})

	t.Run("error response", func(t *testing.T) {
		message := RedisMessage{typ: '_'}

		_, err := message.AsXRangeSlices()
		if err == nil {
			t.Fatal("AsXRangeSlices() expected error but got none")
		}
	})
}

func TestRedisMessage_AsXReadSlices(t *testing.T) {
	t.Run("XREAD response with map format", func(t *testing.T) {
		message := slicemsg('%', []RedisMessage{
			strmsg('+', "stream1"),
			slicemsg('*', []RedisMessage{
				slicemsg('*', []RedisMessage{
					strmsg('+', "1234567890-0"),
					slicemsg('*', []RedisMessage{
						strmsg('+', "field1"),
						strmsg('+', "value1"),
					}),
				}),
			}),
			strmsg('+', "stream2"),
			slicemsg('*', []RedisMessage{
				slicemsg('*', []RedisMessage{
					strmsg('+', "1234567890-1"),
					slicemsg('*', []RedisMessage{
						strmsg('+', "field2"),
						strmsg('+', "value2"),
					}),
				}),
			}),
		})

		want := map[string][]XRangeSlice{
			"stream1": {{
				ID: "1234567890-0",
				FieldValues: []XRangeFieldValue{
					{Field: "field1", Value: "value1"},
				},
			}},
			"stream2": {{
				ID: "1234567890-1",
				FieldValues: []XRangeFieldValue{
					{Field: "field2", Value: "value2"},
				},
			}},
		}

		got, err := message.AsXReadSlices()
		if err != nil {
			t.Fatalf("AsXReadSlices() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXReadSlices() = %v, want %v", got, want)
		}
	})

	t.Run("XREAD response with array format", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "stream1"),
				slicemsg('*', []RedisMessage{
					slicemsg('*', []RedisMessage{
						strmsg('+', "1234567890-0"),
						slicemsg('*', []RedisMessage{
							strmsg('+', "field1"),
							strmsg('+', "value1"),
						}),
					}),
				}),
			}),
		})

		want := map[string][]XRangeSlice{
			"stream1": {{
				ID: "1234567890-0",
				FieldValues: []XRangeFieldValue{
					{Field: "field1", Value: "value1"},
				},
			}},
		}

		got, err := message.AsXReadSlices()
		if err != nil {
			t.Fatalf("AsXReadSlices() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXReadSlices() = %v, want %v", got, want)
		}
	})

	t.Run("error response", func(t *testing.T) {
		message := strmsg('-', "ERR some error")

		_, err := message.AsXReadSlices()
		if err == nil {
			t.Fatal("AsXReadSlices() expected error but got none")
		}
	})

	t.Run("invalid type", func(t *testing.T) {
		message := strmsg('+', "invalid")

		_, err := message.AsXReadSlices()
		if err == nil {
			t.Fatal("AsXReadSlices() expected error but got none")
		}
		if !strings.Contains(err.Error(), "is not a map/array/set") {
			t.Errorf("AsXReadSlices() error = %v, want error containing 'is not a map/array/set'", err)
		}
	})

	t.Run("invalid array entry length", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "stream1"),
			}),
		})

		_, err := message.AsXReadSlices()
		if err == nil {
			t.Fatal("AsXReadSlices() expected error but got none")
		}
		if !strings.Contains(err.Error(), "got 1, wanted 2") {
			t.Errorf("AsXReadSlices() error = %v, want error containing 'got 1, wanted 2'", err)
		}
	})

	t.Run("map format with AsXRangeSlices error", func(t *testing.T) {
		message := slicemsg('%', []RedisMessage{
			strmsg('+', "stream1"),
			strmsg('+', "invalid-range-data"), // This will cause AsXRangeSlices to fail
		})

		_, err := message.AsXReadSlices()
		if err == nil {
			t.Fatal("AsXReadSlices() expected error but got none")
		}
	})

	t.Run("array format with AsXRangeSlices error", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "stream1"),
				strmsg('+', "invalid-range-data"), // This will cause AsXRangeSlices to fail
			}),
		})

		_, err := message.AsXReadSlices()
		if err == nil {
			t.Fatal("AsXReadSlices() expected error but got none")
		}
	})

	t.Run("array format with non-array entry", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			strmsg('+', "not-an-array-entry"),
		})

		_, err := message.AsXReadSlices()
		if err == nil {
			t.Fatal("AsXReadSlices() expected error but got none")
		}
		if !strings.Contains(err.Error(), "got 0, wanted 2") {
			t.Errorf("AsXReadSlices() error = %v, want error containing 'got 0, wanted 2'", err)
		}
	})
}
func TestRedisResult_XRangeSlice_Methods(t *testing.T) {
	t.Run("AsXRangeSlice with error", func(t *testing.T) {
		result := RedisResult{err: errors.New("network error")}

		_, err := result.AsXRangeSlice()
		if err == nil {
			t.Fatal("AsXRangeSlice() expected error but got none")
		}
		if err.Error() != "network error" {
			t.Errorf("AsXRangeSlice() error = %v, want 'network error'", err)
		}
	})

	t.Run("AsXRangeSlice success", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			strmsg('+', "1234567890-0"),
			slicemsg('*', []RedisMessage{
				strmsg('+', "field1"),
				strmsg('+', "value1"),
			}),
		})
		result := RedisResult{val: message}

		want := XRangeSlice{
			ID: "1234567890-0",
			FieldValues: []XRangeFieldValue{
				{Field: "field1", Value: "value1"},
			},
		}

		got, err := result.AsXRangeSlice()
		if err != nil {
			t.Fatalf("AsXRangeSlice() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXRangeSlice() = %v, want %v", got, want)
		}
	})

	t.Run("AsXRangeSlices with error", func(t *testing.T) {
		result := RedisResult{err: errors.New("network error")}

		_, err := result.AsXRangeSlices()
		if err == nil {
			t.Fatal("AsXRangeSlices() expected error but got none")
		}
		if err.Error() != "network error" {
			t.Errorf("AsXRangeSlices() error = %v, want 'network error'", err)
		}
	})

	t.Run("AsXRangeSlices success", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "1234567890-0"),
				slicemsg('*', []RedisMessage{
					strmsg('+', "field1"),
					strmsg('+', "value1"),
				}),
			}),
		})
		result := RedisResult{val: message}

		want := []XRangeSlice{{
			ID: "1234567890-0",
			FieldValues: []XRangeFieldValue{
				{Field: "field1", Value: "value1"},
			},
		}}

		got, err := result.AsXRangeSlices()
		if err != nil {
			t.Fatalf("AsXRangeSlices() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXRangeSlices() = %v, want %v", got, want)
		}
	})

	t.Run("AsXReadSlices with error", func(t *testing.T) {
		result := RedisResult{err: errors.New("network error")}

		_, err := result.AsXReadSlices()
		if err == nil {
			t.Fatal("AsXReadSlices() expected error but got none")
		}
		if err.Error() != "network error" {
			t.Errorf("AsXReadSlices() error = %v, want 'network error'", err)
		}
	})

	t.Run("AsXReadSlices success", func(t *testing.T) {
		message := slicemsg('*', []RedisMessage{
			slicemsg('*', []RedisMessage{
				strmsg('+', "stream1"),
				slicemsg('*', []RedisMessage{
					slicemsg('*', []RedisMessage{
						strmsg('+', "1234567890-0"),
						slicemsg('*', []RedisMessage{
							strmsg('+', "field1"),
							strmsg('+', "value1"),
						}),
					}),
				}),
			}),
		})
		result := RedisResult{val: message}

		want := map[string][]XRangeSlice{
			"stream1": {{
				ID: "1234567890-0",
				FieldValues: []XRangeFieldValue{
					{Field: "field1", Value: "value1"},
				},
			}},
		}

		got, err := result.AsXReadSlices()
		if err != nil {
			t.Fatalf("AsXReadSlices() error = %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AsXReadSlices() = %v, want %v", got, want)
		}
	})
}

// Test to verify order preservation and duplicate handling (the core issue)
func TestXRangeSlice_OrderAndDuplicates(t *testing.T) {
	// This test specifically verifies the key features mentioned in the issue
	message := slicemsg('*', []RedisMessage{
		strmsg('+', "1747784186966-0"),
		slicemsg('*', []RedisMessage{
			strmsg('+', "foo"),
			strmsg('+', "1"),
			strmsg('+', "foo"),
			strmsg('+', "2"),
			strmsg('+', "bar"),
			strmsg('+', "3"),
			strmsg('+', "bar"),
			strmsg('+', "4"),
		}),
	})

	result, err := message.AsXRangeSlice()
	if err != nil {
		t.Fatalf("AsXRangeSlice() error = %v", err)
	}

	// Verify order is preserved
	expectedOrder := []XRangeFieldValue{
		{Field: "foo", Value: "1"},
		{Field: "foo", Value: "2"},
		{Field: "bar", Value: "3"},
		{Field: "bar", Value: "4"},
	}

	if !reflect.DeepEqual(result.FieldValues, expectedOrder) {
		t.Errorf("Order not preserved. Got %v, want %v", result.FieldValues, expectedOrder)
	}

	// Verify duplicates are preserved
	fooCount := 0
	barCount := 0
	for _, fv := range result.FieldValues {
		switch fv.Field {
		case "foo":
			fooCount++
		case "bar":
			barCount++
		}
	}

	if fooCount != 2 {
		t.Errorf("Expected 2 'foo' entries, got %d", fooCount)
	}
	if barCount != 2 {
		t.Errorf("Expected 2 'bar' entries, got %d", barCount)
	}

	// Show that converting to map loses information (for comparison)
	oldStyleMap := map[string]string{
		"foo": "2", // Only keeps the last value
		"bar": "4", // Only keeps the last value
	}

	// Convert new style back to map should match old behavior for last values
	newStyleAsMap := make(map[string]string)
	for _, fv := range result.FieldValues {
		newStyleAsMap[fv.Field] = fv.Value // This overwrites, just like the old map behavior
	}

	if !reflect.DeepEqual(newStyleAsMap, oldStyleMap) {
		t.Errorf("Map conversion doesn't match expected behavior. Got %v, want %v", newStyleAsMap, oldStyleMap)
	}
}
