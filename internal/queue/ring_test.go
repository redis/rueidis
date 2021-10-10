package queue

import (
	"strconv"
	"testing"
)

func TestRing(t *testing.T) {
	t.Run("Write & Read", func(t *testing.T) {
		ring := NewRing()
		size := 8000
		fixture := make(map[string]struct{}, size)
		for i := 0; i < size; i++ {
			fixture[strconv.Itoa(i)] = struct{}{}
		}

		for cmd := range fixture {
			go ring.PutOne([]string{cmd})
		}

		for i := 0; i < size; i++ {
			cmd1 := ring.NextCmd()
			cmd2, ch := ring.NextResultCh()
			if cmd1[0][0] != cmd2[0][0] {
				t.Fatalf("cmds read by NextCmd and NextResultCh is not the same one")
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
			cmd1 := ring.NextCmd()
			cmd2, ch := ring.NextResultCh()
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

	t.Run("TryNextCmd", func(t *testing.T) {
		ring := NewRing()
		if ring.TryNextCmd() != nil {
			t.Fatalf("TryNextCmd should returns nil if empty")
		}
		ring.PutOne([]string{"0"})
		if cmd := ring.TryNextCmd(); len(cmd) == 0 || cmd[0][0] != "0" {
			t.Fatalf("TryNextCmd should returns next cmd")
		}
	})

	t.Run("NextResultCh", func(t *testing.T) {
		ring := NewRing()
		ring.PutOne([]string{"0"})

		if cmd, ch := ring.NextResultCh(); cmd != nil || ch != nil {
			t.Fatalf("NextResultCh should returns nil if not NextCmd yet")
		}

		if cmd := ring.TryNextCmd(); len(cmd) == 0 || cmd[0][0] != "0" {
			t.Fatalf("TryNextCmd should returns next cmd")
		}

		if cmd, ch := ring.NextResultCh(); len(cmd) == 0 || cmd[0][0] != "0" || ch == nil {
			t.Fatalf("NextResultCh should returns next cmd after NextCmd")
		}
	})
}
