package rueidis

import "testing"

func TestBinaryString(t *testing.T) {
	if str := []byte{0, 1, 2, 3, 4}; string(str) != BinaryString(str) {
		t.Fatalf("BinaryString mismatch")
	}
}
