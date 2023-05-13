package mock

import (
	"errors"
	"reflect"
	"testing"

	"github.com/redis/rueidis"
)

func TestRedisString(t *testing.T) {
	m := RedisString("s")
	if v, err := m.ToString(); err != nil || v != "s" {
		t.Fatalf("unexpected value %v", v)
	}
}

func TestRedisError(t *testing.T) {
	m := RedisError("s")
	if err := m.Error(); err == nil || err.Error() != "s" {
		t.Fatalf("unexpected value %v", err)
	}
}

func TestRedisInt64(t *testing.T) {
	m := RedisInt64(1)
	if v, err := m.ToInt64(); err != nil || v != int64(1) {
		t.Fatalf("unexpected value %v", v)
	}
	if v, err := m.AsInt64(); err != nil || v != int64(1) {
		t.Fatalf("unexpected value %v", v)
	}
}

func TestRedisFloat64(t *testing.T) {
	m := RedisFloat64(1)
	if v, err := m.ToFloat64(); err != nil || v != float64(1) {
		t.Fatalf("unexpected value %v", v)
	}
	if v, err := m.AsFloat64(); err != nil || v != float64(1) {
		t.Fatalf("unexpected value %v", v)
	}
}

func TestRedisBool(t *testing.T) {
	m := RedisBool(true)
	if v, err := m.ToBool(); err != nil || v != true {
		t.Fatalf("unexpected value %v", v)
	}
	if v, err := m.AsBool(); err != nil || v != true {
		t.Fatalf("unexpected value %v", v)
	}
}

func TestRedisNil(t *testing.T) {
	m := RedisNil()
	if err := m.Error(); err == nil {
		t.Fatalf("unexpected value %v", err)
	}
	if v := m.IsNil(); v != true {
		t.Fatalf("unexpected value %v", v)
	}
}

func TestRedisArray(t *testing.T) {
	m := RedisArray(RedisString("0"), RedisString("1"), RedisString("2"))
	if arr, err := m.AsStrSlice(); err != nil || !reflect.DeepEqual(arr, []string{"0", "1", "2"}) {
		t.Fatalf("unexpected value %v", err)
	}
}

func TestRedisMap(t *testing.T) {
	m := RedisMap(map[string]rueidis.RedisMessage{
		"a": RedisString("0"),
		"b": RedisString("1"),
	})
	if arr, err := m.AsStrMap(); err != nil || !reflect.DeepEqual(arr, map[string]string{
		"a": "0",
		"b": "1",
	}) {
		t.Fatalf("unexpected value %v", err)
	}
	if arr, err := m.ToMap(); err != nil || !reflect.DeepEqual(arr, map[string]rueidis.RedisMessage{
		"a": RedisString("0"),
		"b": RedisString("1"),
	}) {
		t.Fatalf("unexpected value %v", err)
	}
}

func TestRedisResult(t *testing.T) {
	r := Result(RedisNil())
	if err := r.Error(); !rueidis.IsRedisNil(err) {
		t.Fatalf("unexpected value %v", err)
	}
}

func TestErrorResult(t *testing.T) {
	r := ErrorResult(errors.New("any"))
	if err := r.Error(); err.Error() != "any" {
		t.Fatalf("unexpected value %v", err)
	}
}
