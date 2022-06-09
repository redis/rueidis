package rueidis

import (
	"encoding/json"
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

type recursive struct {
	R *recursive
}
