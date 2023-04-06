package util

import (
	"math"
	"testing"
)

func TestToFloat64(t *testing.T) {
	for _, c := range []struct {
		Str string
		Val float64
	}{
		{Str: "1", Val: 1},
		{Str: "-1", Val: -1},
		{Str: "nan", Val: math.NaN()},
		{Str: "-nan", Val: math.NaN()},
	} {
		if v, _ := ToFloat64(c.Str); v != c.Val && !(math.IsNaN(v) && math.IsNaN(c.Val)) {
			t.Fail()
		}
	}
}

func TestToFloat32(t *testing.T) {
	for _, c := range []struct {
		Str string
		Val float32
	}{
		{Str: "1", Val: 1},
		{Str: "-1", Val: -1},
		{Str: "nan", Val: float32(math.NaN())},
		{Str: "-nan", Val: float32(math.NaN())},
	} {
		if v, _ := ToFloat32(c.Str); v != c.Val && !(math.IsNaN(float64(v)) && math.IsNaN(float64(c.Val))) {
			t.Fail()
		}
	}
}
