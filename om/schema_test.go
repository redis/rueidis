package om

import (
	"reflect"
	"strings"
	"testing"
)

type s1 struct {
	A int `redis:",key"`
}

type s2 struct {
	A string `redis:",ver"`
}

type s3 struct {
	A       string `json:"-" redis:",key"`
	B       int64  `redis:",ver"`
	private int64
}

type s4 struct {
	A       string `redis:",key"`
	B       int64  `redis:",ver"`
	private int64
}

type s5 struct {
	A       string `redis:",key"`
	private int64
}

type s6 struct {
	A string `redis:",key"`
	B int64  `redis:",ver"`
	C int64  `redis:",exat"`
}

func TestSchema(t *testing.T) {
	t.Run("non struct", func(t *testing.T) {
		if v := recovered(func() {
			newSchema(reflect.TypeOf(map[string]string{}))
		}); !strings.Contains(v, "should be a struct") {
			t.Fatalf("unexpected msg %v", v)
		}
	})
	t.Run("non string `redis:\",key\"`", func(t *testing.T) {
		if v := recovered(func() {
			newSchema(reflect.TypeOf(s1{}))
		}); !strings.Contains(v, "should be a string") {
			t.Fatalf("unexpected msg %v", v)
		}
	})
	t.Run("non int64 `redis:\",ver\"`", func(t *testing.T) {
		if v := recovered(func() {
			newSchema(reflect.TypeOf(s2{}))
		}); !strings.Contains(v, "should be a int64") {
			t.Fatalf("unexpected msg %v", v)
		}
	})
	t.Run("missing `redis:\",key\"`", func(t *testing.T) {
		if v := recovered(func() {
			newSchema(reflect.TypeOf(s3{}))
		}); !strings.Contains(v, "should have one field with `redis:\",key\"` tag") {
			t.Fatalf("unexpected msg %v", v)
		}
	})
	t.Run("ver is not verless", func(t *testing.T) {
		v := newSchema(reflect.TypeOf(s4{}))
		if v.verless {
			t.Fatal("schema should not be verless")
		}
	})
	t.Run("missing `redis:\",ver\"` should be verless", func(t *testing.T) {
		v := newSchema(reflect.TypeOf(s5{}))
		if !v.verless {
			t.Fatal("schema should be verless")
		}
	})
	t.Run("non time.Time `redis:\",exat\"`", func(t *testing.T) {
		if v := recovered(func() {
			newSchema(reflect.TypeOf(s6{}))
		}); !strings.Contains(v, "should be a time.Time") {
			t.Fatalf("unexpected msg %v", v)
		}
	})
}

func recovered(fn func()) (msg string) {
	defer func() {
		msg = recover().(string)
	}()
	fn()
	return msg
}
