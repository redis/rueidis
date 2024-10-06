package util

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	if len(arr) != 10 {
		t.Errorf("Expected array length 10, got %d", len(arr))
	}
}

func TestFastRand(t *testing.T) {
	n := 10
	res := FastRand(n)
	if res < 0 || res >= n {
		t.Errorf("Expected result between 0 and %d, got %d", n-1, res)
	}
}

func TestRandomBytes(t *testing.T) {
	val := RandomBytes()
	if len(val) != 24 {
		t.Errorf("Expected length 24, got %d", len(val))
	}
}
