package rueidis

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestIsRedisNil(t *testing.T) {
	err := &RedisError{typ: '_'}
	if !IsRedisNil(err) {
		t.Fatal("IsRedisNil fail")
	}
	if IsRedisNil(errors.New("other")) {
		t.Fatal("IsRedisNil fail")
	}
	if err.Error() != "redis nil message" {
		t.Fatal("IsRedisNil fail")
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

	t.Run("DecodeJSON", func(t *testing.T) {
		v := map[string]string{}
		if err := (RedisResult{err: errors.New("other")}).DecodeJSON(&v); err == nil {
			t.Fatal("AsReader not failed as expected")
		}
		if err := (RedisResult{val: RedisMessage{typ: '-'}}).DecodeJSON(&v); err == nil {
			t.Fatal("AsReader not failed as expected")
		}
		if _ = (RedisResult{val: RedisMessage{typ: '+', string: `{"k":"v"}`}}).DecodeJSON(&v); v["k"] != "v" {
			t.Fatalf("AsReader not get value as expected %v", v)
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
	})

	t.Run("AsFloat64", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsFloat64(); err == nil {
			t.Fatal("AsFloat64 not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsFloat64(); err == nil {
			t.Fatal("AsFloat64 not failed as expected")
		}
		if v, _ := (RedisResult{val: RedisMessage{typ: '+', string: "1"}}).AsFloat64(); v != 1 {
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

	t.Run("AsXRange", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}
		values := []RedisMessage{{string: "id", typ: '+'}, {typ: '*', values: []RedisMessage{{typ: '+', string: "a"}, {typ: '+', string: "b"}}}}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: values}}).AsXRange(); !reflect.DeepEqual(XRange{
			ID:          "id",
			FieldValues: map[string]string{"a": "b"},
		}, ret) {
			t.Fatal("AsXRange not get value as expected")
		}
	})

	t.Run("AsXRangeSlice", func(t *testing.T) {
		if _, err := (RedisResult{err: errors.New("other")}).AsXRangeSlice(); err == nil {
			t.Fatal("AsXRangeSlice not failed as expected")
		}
		if _, err := (RedisResult{val: RedisMessage{typ: '-'}}).AsXRangeSlice(); err == nil {
			t.Fatal("AsXRangeSlice not failed as expected")
		}
		values := []RedisMessage{{string: "id", typ: '+'}, {typ: '*', values: []RedisMessage{{typ: '+', string: "a"}, {typ: '+', string: "b"}}}}
		if ret, _ := (RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '*', values: values}}}}).AsXRangeSlice(); !reflect.DeepEqual([]XRange{{
			ID:          "id",
			FieldValues: map[string]string{"a": "b"},
		}}, ret) {
			t.Fatal("AsXRangeSlice not get value as expected")
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
}

//gocyclo:ignore
func TestRedisMessage(t *testing.T) {
	t.Run("IsNil", func(t *testing.T) {
		if !(&RedisMessage{typ: '_'}).IsNil() {
			t.Fatal("IsNil fail")
		}
	})
	t.Run("ToInt64", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).ToInt64(); err == nil {
			t.Fatal("ToInt64 not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a int64") {
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
			if !strings.Contains(recover().(string), "redis message type t is not a bool") {
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
			if !strings.Contains(recover().(string), "redis message type t is not a float64") {
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
			if !strings.Contains(recover().(string), "redis message type : is not a string") {
				t.Fatal("AsInt64 not panic as expected")
			}
		}()
		(&RedisMessage{typ: ':'}).AsInt64()
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

	t.Run("AsMap", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsMap(); err == nil {
			t.Fatal("AsMap not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a array") {
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

	t.Run("ToMap", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).ToMap(); err == nil {
			t.Fatal("ToMap not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a map") {
				t.Fatal("ToString not panic as expected")
			}
		}()
		(&RedisMessage{typ: 't'}).ToMap()
	})

	t.Run("AsXRange - no range id", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}

		if _, err := (&RedisMessage{typ: '*'}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}

		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{{typ: '_'}, {typ: '%'}}}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type : is not a string") {
				t.Fatal("AsXRange not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*', values: []RedisMessage{{typ: ':'}, {typ: '%'}}}).AsXRange()
	})

	t.Run("AsXRange - no range field values", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}

		if _, err := (&RedisMessage{typ: '*'}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}

		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{{typ: '+'}, {typ: '_'}}}).AsXRange(); err == nil {
			t.Fatal("AsXRange not failed as expected")
		}

		defer func() {
			if !strings.Contains(recover().(string), "redis message type t is not a map/array/set") {
				t.Fatal("AsXRange not panic as expected")
			}
		}()
		(&RedisMessage{typ: '*', values: []RedisMessage{{typ: '+'}, {typ: 't'}}}).AsXRange()
	})

	t.Run("AsXRangeSlice", func(t *testing.T) {
		if _, err := (&RedisMessage{typ: '_'}).AsXRangeSlice(); err == nil {
			t.Fatal("AsXRangeSlice not failed as expected")
		}

		if _, err := (&RedisMessage{typ: '*', values: []RedisMessage{{typ: '_'}}}).AsXRangeSlice(); err == nil {
			t.Fatal("AsXRangeSlice not failed as expected")
		}
	})

	t.Run("ToMap with non string key", func(t *testing.T) {
		defer func() {
			if !strings.Contains(recover().(string), "redis message type : as map key is not supported") {
				t.Fatal("ToString not panic as expected")
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
}
