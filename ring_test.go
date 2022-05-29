package rueidis

import (
	"runtime"
	"strconv"
	"testing"

	"github.com/rueian/rueidis/internal/cmds"
)

//gocyclo:ignore
func TestRing(t *testing.T) {
	t.Run("PutOne", func(t *testing.T) {
		ring := newRing()
		size := 5000
		fixture := make(map[string]struct{}, size)
		for i := 0; i < size; i++ {
			fixture[strconv.Itoa(i)] = struct{}{}
		}

		for cmd := range fixture {
			go ring.PutOne(cmds.NewCompleted([]string{cmd}))
		}

		for len(fixture) != 0 {
			cmd1, _, _ := ring.NextWriteCmd()
			if cmd1.IsEmpty() {
				runtime.Gosched()
				continue
			}
			cmd2, _, ch, cond := ring.NextResultCh()
			cond.L.Unlock()
			cond.Signal()
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
		ring := newRing()
		size := 5000
		fixture := make(map[string]struct{}, size)
		for i := 0; i < size; i++ {
			fixture[strconv.Itoa(i)] = struct{}{}
		}

		base := [][]string{{"a"}, {"b"}, {"c"}, {"d"}}
		for cmd := range fixture {
			go ring.PutMulti(cmds.NewMultiCompleted(append([][]string{{cmd}}, base...)))
		}

		for len(fixture) != 0 {
			_, cmd1, _ := ring.NextWriteCmd()
			if cmd1 == nil {
				runtime.Gosched()
				continue
			}
			_, cmd2, ch, cond := ring.NextResultCh()
			cond.L.Unlock()
			cond.Signal()
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
		ring := newRing()
		if one, multi, _ := ring.NextWriteCmd(); !one.IsEmpty() || multi != nil {
			t.Fatalf("NextWriteCmd should returns nil if empty")
		}
		if one, multi, ch, cond := ring.NextResultCh(); !one.IsEmpty() || multi != nil || ch != nil {
			t.Fatalf("NextResultCh should returns nil if not NextWriteCmd yet")
		} else {
			cond.L.Unlock()
			cond.Signal()
		}

		ring.PutOne(cmds.NewCompleted([]string{"0"}))
		if one, _, _ := ring.NextWriteCmd(); len(one.Commands()) == 0 || one.Commands()[0] != "0" {
			t.Fatalf("NextWriteCmd should returns next cmd")
		}
		if one, _, ch, cond := ring.NextResultCh(); len(one.Commands()) == 0 || one.Commands()[0] != "0" || ch == nil {
			t.Fatalf("NextResultCh should returns next cmd after NextWriteCmd")
		} else {
			cond.L.Unlock()
			cond.Signal()
		}

		ring.PutMulti(cmds.NewMultiCompleted([][]string{{"0"}}))
		if _, multi, _ := ring.NextWriteCmd(); len(multi) == 0 || multi[0].Commands()[0] != "0" {
			t.Fatalf("NextWriteCmd should returns next cmd")
		}
		if _, multi, ch, cond := ring.NextResultCh(); len(multi) == 0 || multi[0].Commands()[0] != "0" || ch == nil {
			t.Fatalf("NextResultCh should returns next cmd after NextWriteCmd")
		} else {
			cond.L.Unlock()
			cond.Signal()
		}
	})

	t.Run("NextWriteCmd & CleanNoReply No Effect", func(t *testing.T) {
		ring := newRing()
		ring.CleanNoReply() // no effect
		ring.PutOne(cmds.NewCompleted([]string{"0"}))
		if one, _, _ := ring.NextWriteCmd(); len(one.Commands()) == 0 || one.Commands()[0] != "0" {
			t.Fatalf("NextWriteCmd should returns next cmd")
		}
		ring.CleanNoReply() // no effect
		if one, _, ch, cond := ring.NextResultCh(); len(one.Commands()) == 0 || one.Commands()[0] != "0" || ch == nil {
			t.Fatalf("NextResultCh should returns next cmd after NextWriteCmd")
		} else {
			cond.L.Unlock()
			cond.Signal()
		}
	})

	t.Run("NextWriteCmd & CleanNoReply One", func(t *testing.T) {
		ring := newRing()
		ring.PutOne(cmds.UnsubscribeCmd)
		if one, _, _ := ring.NextWriteCmd(); len(one.Commands()) == 0 || one.Commands()[0] != "UNSUBSCRIBE" {
			t.Fatalf("NextWriteCmd should returns next cmd")
		}
		ring.CleanNoReply()
		if _, _, ch, _ := ring.NextResultCh(); ch != nil {
			t.Fatalf("NextResultCh should returns nothing after CleanNoReply")
		}
	})

	t.Run("NextWriteCmd & CleanNoReply Multi", func(t *testing.T) {
		ring := newRing()
		ring.PutMulti([]cmds.Completed{cmds.UnsubscribeCmd, cmds.SUnsubscribeCmd})
		if _, multi, _ := ring.NextWriteCmd(); len(multi) != 2 || multi[0].Commands()[0] != "UNSUBSCRIBE" || multi[1].Commands()[0] != "SUNSUBSCRIBE" {
			t.Fatalf("NextWriteCmd should returns next cmd")
		}
		ring.CleanNoReply()
		if _, _, ch, _ := ring.NextResultCh(); ch != nil {
			t.Fatalf("NextResultCh should returns nothing after CleanNoReply")
		}
	})
}
