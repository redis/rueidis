package rueidis

import (
	"errors"
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
}

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
}
