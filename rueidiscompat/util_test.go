package rueidiscompat

import (
	"fmt"
	"math"
	"testing"
)

func TestToLower(t *testing.T) {
	tests := []struct {
		input string
		exp   string
	}{{"HELLO", "hello"}, {"Hello", "hello"}, {"123!@#AbC!@#", "123!@#abc!@#"}, {"", ""}}
	for _, test := range tests {
		res := ToLower(test.input)
		if res != test.exp {
			t.Errorf("ToLower(%q) = %q; want %q", test.input, res, test.exp)
		}
	}
}

func TestStringToBytes(t *testing.T) {
	exp := []byte("hello")
	input := "hello"
	res := StringToBytes(input)
	if string(res) != string(exp) {
		t.Errorf("StringToBytes(%q) = %q; want %q", input, res, exp)
	}
}

func TestReplaceSpaces(t *testing.T) {
	tests := []struct {
		input string
		exp   string
	}{{"one space", "one-space"}, {"multiple   spaces", "multiple---spaces"}, {"", ""}}
	for _, test := range tests {
		res := ReplaceSpaces(test.input)
		if res != test.exp {
			t.Errorf("ReplaceSpaces(%q)= %q; want %q", test.input, res, test.exp)
		}
	}
}

func TestGetAddr(t *testing.T) {
	tests := []struct {
		input string
		exp   string
	}{{"192.168.1.1:8080", "192.168.1.1:8080"}, {"[2001:db8::1]:443", "[2001:db8::1]:443"}, {"localhost:3000", "localhost:3000"}, {"12345", ""}, {"", ""}}
	for _, test := range tests {
		res := GetAddr(test.input)
		if res != test.exp {
			t.Errorf("GetAddr(%q)= %q; want %q", test.input, res, test.exp)
		}
	}
}

func TestToInteger(t *testing.T) {
	tests := []struct {
		input interface{}
		exp   int
	}{{123, 123}, {int64(123), 123}, {"123", 123}, {"abc", 0}, {nil, 0}, {float64(123), 0}}
	for _, test := range tests {
		res := ToInteger(test.input)
		if res != test.exp {
			t.Errorf("ToInteger(%v) = %d, want %d", test.input, res, test.exp)
		}
	}
}

func TestToFloat(t *testing.T) {
	tests := []struct {
		input interface{}
		exp   float64
	}{
		{123.45, 123.45},
		{int64(123), 0.0},
		{123, 0.0},
		{"123.45", 123.45},
		{"abc", 0},
		{nil, 0},
	}
	for _, test := range tests {
		res := ToFloat(test.input)
		if math.Abs(res-test.exp) > 0.001 {
			t.Errorf("Testing ToFloat(%v): got %.1f, expected %.1f",
				test.input, res, test.exp)
		}
	}
}

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		input interface{}
		exp   []string
	}{{[]interface{}{"abc", "def"}, []string{"abc", "def"}}, {[]interface{}{1, "abc", 2}, []string{"", "abc", ""}}, {[]interface{}{nil, "abc"}, []string{"", "abc"}}, {[]interface{}{1.2, true, 3.5}, []string{"", "", ""}}, {"abc", nil}}
	for _, test := range tests {
		res := ToStringSlice(test.input)
		if fmt.Sprintf("%v", res) != fmt.Sprintf("%v", test.exp) {
			t.Errorf("For %v, expected %v, but got %v", test.input, test.exp, res)
		}
	}
}
