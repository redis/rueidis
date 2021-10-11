package queue

import (
	"runtime"
	"strconv"
	"testing"
)

func TestRing(t *testing.T) {
	t.Run("PutOne", func(t *testing.T) {
		ring := NewRing()
		size := 8000
		fixture := make(map[string]struct{}, size)
		for i := 0; i < size; i++ {
			fixture[strconv.Itoa(i)] = struct{}{}
		}

		for cmd := range fixture {
			go ring.PutOne([]string{cmd})
		}

		for len(fixture) != 0 {
			cmd1, _ := ring.NextCmd()
			if cmd1 == nil {
				runtime.Gosched()
				continue
			}
			cmd2, _, ch := ring.NextResultCh()
			if cmd1[0] != cmd2[0] {
				t.Fatalf("cmds read by NextCmd and NextResultCh is not the same one")
			}
			if ch == nil || len(ch) != 0 {
				t.Fatalf("channel from NextResultCh is broken")
			}
			delete(fixture, cmd1[0])
		}
	})

	t.Run("PutMulti", func(t *testing.T) {
		ring := NewRing()
		size := 8000
		fixture := make(map[string]struct{}, size)
		for i := 0; i < size; i++ {
			fixture[strconv.Itoa(i)] = struct{}{}
		}

		base := [][]string{{"a"}, {"b"}, {"c"}, {"d"}}
		for cmd := range fixture {
			go ring.PutMulti(append([][]string{{cmd}}, base...))
		}

		for i := 0; i < size; i++ {
			_, cmd1 := ring.NextCmd()
			if cmd1 == nil {
				runtime.Gosched()
				continue
			}
			_, cmd2, ch := ring.NextResultCh()
			for j := 0; j < len(cmd1); j++ {
				if cmd1[j][0] != cmd2[j][0] {
					t.Fatalf("cmds read by NextCmd and NextResultCh is not the same one")
				}
			}
			if ch == nil || len(ch) != 0 {
				t.Fatalf("channel from NextResultCh is broken")
			}
			delete(fixture, cmd1[0][0])
		}

		if len(fixture) != 0 {
			t.Fatalf("not all cmds are read by NextCmd and NextResultCh")
		}
	})

	t.Run("NextCmd & NextResultCh", func(t *testing.T) {
		ring := NewRing()
		if one, multi := ring.NextCmd(); one != nil || multi != nil {
			t.Fatalf("NextCmd should returns nil if empty")
		}
		if one, multi, ch := ring.NextResultCh(); one != nil || multi != nil || ch != nil {
			t.Fatalf("NextResultCh should returns nil if not NextCmd yet")
		}

		ring.PutOne([]string{"0"})
		if one, _ := ring.NextCmd(); len(one) == 0 || one[0] != "0" {
			t.Fatalf("NextCmd should returns next cmd")
		}
		if one, _, ch := ring.NextResultCh(); len(one) == 0 || one[0] != "0" || ch == nil {
			t.Fatalf("NextResultCh should returns next cmd after NextCmd")
		}

		ring.PutMulti([][]string{{"0"}})
		if _, multi := ring.NextCmd(); len(multi) == 0 || multi[0][0] != "0" {
			t.Fatalf("NextCmd should returns next cmd")
		}
		if _, multi, ch := ring.NextResultCh(); len(multi) == 0 || multi[0][0] != "0" || ch == nil {
			t.Fatalf("NextResultCh should returns next cmd after NextCmd")
		}
	})
}
