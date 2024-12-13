package cmds

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestSlot(t *testing.T) {
	t.Run("use tag", func(t *testing.T) {
		for i := 0; i < 10000; i++ {
			key1 := strconv.Itoa(rand.Int())
			key2 := fmt.Sprintf("%s{%s}%s", strconv.Itoa(rand.Int()), key1, strconv.Itoa(rand.Int()))
			if slot(key1) != slot(key2) {
				t.Fatalf("%v and %v should be in the same slot", key1, key2)
			}
		}
	})
	t.Run("not use tag", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			key1 := strconv.Itoa(i)
			key2 := fmt.Sprintf("%s{}", key1)
			if slot(key1) == slot(key2) {
				t.Fatalf("%v and %v should not be in the same slot", key1, key2)
			}
		}
	})
}

func TestCRC16(t *testing.T) {
	t.Run("123456789", func(t *testing.T) {
		if v := crc16("123456789"); v != 0x31C3 {
			t.Fatalf("crc16(123456789) should be 0x31C3, but got %v", v)
		}
	})
}
