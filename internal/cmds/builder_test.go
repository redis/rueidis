package cmds

import "testing"

func TestPut(t *testing.T) {
retry:
	cs1 := get()
	cs1.s = append(cs1.s, "1", "1", "1", "1", "1")
	Put(cs1)
	cs2 := get()
	if cs1 != cs2 {
		goto retry
	}
	if len(cs2.s) != 0 {
		t.Fatalf("Put doesn't clean the CommandSlice")
	}
}
