package rueidis

import (
	"context"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

//gocyclo:ignore
func TestRing(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	t.Run("PutOne", func(t *testing.T) {
		ring := newRing(DefaultRingScale)
		size := 5000
		fixture := make(map[string]struct{}, size)
		for i := 0; i < size; i++ {
			fixture[strconv.Itoa(i)] = struct{}{}
		}

		for cmd := range fixture {
			go ring.PutOne(context.Background(), cmds.NewCompleted([]string{cmd}))
		}

		for len(fixture) != 0 {
			cmd1, _, _ := ring.NextWriteCmd()
			if cmd1.IsEmpty() {
				runtime.Gosched()
				continue
			}
			c, done := ring.NextResultCh()
			cmd2 := c.one
			ch := c.ch
			done()
			if cmd1.Commands()[0] != cmd2.Commands()[0] {
				t.Fatalf("cmds read by NextWriteCmd and NextResultCh is not the same one")
			}
			if ch == nil || len(ch) != 0 {
				t.Fatalf("channel from NextResultCh is broken")
			}
			delete(fixture, cmd1.Commands()[0])
		}
	})

	t.Run("PutMulti", func(t *testing.T) {
		ring := newRing(DefaultRingScale)
		size := 5000
		fixture := make(map[string]struct{}, size)
		for i := 0; i < size; i++ {
			fixture[strconv.Itoa(i)] = struct{}{}
		}

		base := [][]string{{"a"}, {"b"}, {"c"}, {"d"}}
		for cmd := range fixture {
			go ring.PutMulti(context.Background(), cmds.NewMultiCompleted(append([][]string{{cmd}}, base...)), nil)
		}

		for len(fixture) != 0 {
			_, cmd1, _ := ring.NextWriteCmd()
			if cmd1 == nil {
				runtime.Gosched()
				continue
			}
			c, done := ring.NextResultCh()
			cmd2 := c.multi
			ch := c.ch
			done()
			for j := 0; j < len(cmd1); j++ {
				if cmd1[j].Commands()[0] != cmd2[j].Commands()[0] {
					t.Fatalf("cmds read by NextWriteCmd and NextResultCh is not the same one")
				}
			}
			if ch == nil || len(ch) != 0 {
				t.Fatalf("channel from NextResultCh is broken")
			}
			delete(fixture, cmd1[0].Commands()[0])
		}
	})

	t.Run("NextWriteCmd & NextResultCh", func(t *testing.T) {
		ring := newRing(DefaultRingScale)
		if one, multi, _ := ring.NextWriteCmd(); !one.IsEmpty() || multi != nil {
			t.Fatalf("NextWriteCmd should returns nil if empty")
		}
		c, done := ring.NextResultCh()
		one, multi, ch := c.one, c.multi, c.ch
		if !one.IsEmpty() || multi != nil || ch != nil {
			t.Fatalf("NextResultCh should returns nil if not NextWriteCmd yet")
		} else {
			done()
		}

		ring.PutOne(context.Background(), cmds.NewCompleted([]string{"0"}))
		if one, _, _ := ring.NextWriteCmd(); len(one.Commands()) == 0 || one.Commands()[0] != "0" {
			t.Fatalf("NextWriteCmd should returns next cmd")
		}
		c, done = ring.NextResultCh()
		one, multi, ch = c.one, c.multi, c.ch
		if len(one.Commands()) == 0 || one.Commands()[0] != "0" || ch == nil {
			t.Fatalf("NextResultCh should returns next cmd after NextWriteCmd")
		} else {
			done()
		}

		ring.PutMulti(context.Background(), cmds.NewMultiCompleted([][]string{{"0"}}), nil)
		if _, multi, _ := ring.NextWriteCmd(); len(multi) == 0 || multi[0].Commands()[0] != "0" {
			t.Fatalf("NextWriteCmd should returns next cmd")
		}
		c, done = ring.NextResultCh()
		multi, ch = c.multi, c.ch
		if len(multi) == 0 || multi[0].Commands()[0] != "0" || ch == nil {
			t.Fatalf("NextResultCh should returns next cmd after NextWriteCmd")
		} else {
			done()
		}
	})

	t.Run("PutOne Wakeup WaitForWrite", func(t *testing.T) {
		ring := newRing(DefaultRingScale)
		if one, _, ch := ring.NextWriteCmd(); ch == nil {
			go func() {
				time.Sleep(time.Millisecond * 100)
				ring.PutOne(context.Background(), cmds.PingCmd)
			}()
			if one, _, ch = ring.WaitForWrite(); ch != nil && one.Commands()[0] == cmds.PingCmd.Commands()[0] {
				return
			}
		}
		t.Fatal("Should sleep")
	})

	t.Run("PutMulti Wakeup WaitForWrite", func(t *testing.T) {
		ring := newRing(DefaultRingScale)
		if _, _, ch := ring.NextWriteCmd(); ch == nil {
			go func() {
				time.Sleep(time.Millisecond * 100)
				ring.PutMulti(context.Background(), []Completed{cmds.PingCmd}, nil)
			}()
			if _, multi, ch := ring.WaitForWrite(); ch != nil && multi[0].Commands()[0] == cmds.PingCmd.Commands()[0] {
				return
			}
		}
		t.Fatal("Should sleep")
	})
}
