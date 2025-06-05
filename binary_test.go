package rueidis

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestBinaryString(t *testing.T) {
	if str := []byte{0, 1, 2, 3, 4}; string(str) != BinaryString(str) {
		t.Fatalf("BinaryString mismatch")
	}
}

func TestJSON(t *testing.T) {
	if v := JSON("a"); v != `"a"` {
		t.Fatalf("unexpected JSON result")
	}
}

func TestJSONPanic(t *testing.T) {
	defer func() {
		if m := recover().(*json.UnsupportedValueError); !strings.Contains(m.Error(), "encountered a cycle") {
			t.Fatalf("should panic")
		}
	}()
	a := &recursive{}
	a.R = a
	JSON(a)
}

func TestVectorString32(t *testing.T) {
	for _, test := range [][]float32{
		{},
		{0, 0, 0, 0},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{.9, .9, .9, .9, .9, .9, .9, .9, .9, .9, .9},
		{-.1, -.1, -.1, -.1, -.1, -.1, -.1, -.1, -.1, -.1},
		{.1, -.1, .1, -.1, .1, -.1, .1, -.1, .1, -.1},
	} {
		if !reflect.DeepEqual(test, ToVector32(VectorString32(test))) {
			t.Fatalf("fail to convert %v", test)
		}
	}
}

func TestVectorString64(t *testing.T) {
	for _, test := range [][]float64{
		{},
		{0, 0, 0, 0},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{.9, .9, .9, .9, .9, .9, .9, .9, .9, .9, .9},
		{-.1, -.1, -.1, -.1, -.1, -.1, -.1, -.1, -.1, -.1},
		{.1, -.1, .1, -.1, .1, -.1, .1, -.1, .1, -.1},
	} {
		if !reflect.DeepEqual(test, ToVector64(VectorString64(test))) {
			t.Fatalf("fail to convert %v", test)
		}
	}
}

type recursive struct {
	R *recursive
}
