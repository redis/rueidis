package rueidis

import (
	"testing"
	"time"
)

func TestSubs_Publish(t *testing.T) {
	t.Run("without subs", func(t *testing.T) {
		s := newSubs()
		s.Publish("aa", PubSubMessage{}) // just no block
	})

	t.Run("with multiple subs", func(t *testing.T) {
		s := newSubs()
		ch1, cancel1 := s.Subscribe([]string{"a"})
		ch2, cancel2 := s.Subscribe([]string{"a"})
		ch3, cancel3 := s.Subscribe([]string{"b"})
		m1 := PubSubMessage{Pattern: "1", Channel: "2", Message: "3"}
		m2 := PubSubMessage{Pattern: "11", Channel: "22", Message: "33"}
		go func() {
			s.Publish("a", m1)
			s.Publish("b", m2)
		}()
		for m := range ch1 {
			if m != m1 {
				t.Fatalf("unexpected msg %v", m)
			}
			cancel1()
		}
		for m := range ch2 {
			if m != m1 {
				t.Fatalf("unexpected msg %v", m)
			}
			cancel2()
		}
		for m := range ch3 {
			if m != m2 {
				t.Fatalf("unexpected msg %v", m)
			}
			cancel3()
		}
	})

	t.Run("drain ch", func(t *testing.T) {
		s := newSubs()
		ch, cancel := s.Subscribe([]string{"a"})
		s.Publish("a", PubSubMessage{})
		if len(ch) != 1 {
			t.Fatalf("unexpected ch len %v", len(ch))
		}
		cancel()
		for ; len(ch) != 0; time.Sleep(time.Millisecond * 100) {
			t.Log("wait ch to be drain")
		}
	})
}

func TestSubs_Unsubscribe(t *testing.T) {
	s := newSubs()
	ch, _ := s.Subscribe([]string{"1", "2"})
	go func() {
		s.Publish("1", PubSubMessage{})
	}()
	_, ok := <-ch
	if !ok {
		t.Fatalf("unexpected ch closed")
	}
	s.Unsubscribe("1")
	_, ok = <-ch
	if ok {
		t.Fatalf("unexpected ch unclosed")
	}
}

func TestSubs_Confirm(t *testing.T) {
	s := newSubs()
	_, remove := s.Subscribe([]string{"1", "2"})

	if v := s.Confirmed(); v != 0 {
		t.Fatalf("unexpected confirmed count %v", v)
	}

	s.Confirm("1")
	if v := s.Confirmed(); v != 1 {
		t.Fatalf("unexpected confirmed count %v", v)
	}

	remove() // confirmed count should not be affected by remove()
	if v := s.Confirmed(); v != 1 {
		t.Fatalf("unexpected confirmed count %v", v)
	}

	s.Confirm("1")
	s.Confirm("2")
	if v := s.Confirmed(); v != 2 {
		t.Fatalf("unexpected confirmed count %v", v)
	}

	s.Confirm("3")
	if v := s.Confirmed(); v != 3 {
		t.Fatalf("unexpected confirmed count %v", v)
	}

	s.Unsubscribe("1")
	if v := s.Confirmed(); v != 2 {
		t.Fatalf("unexpected confirmed count %v", v)
	}
	s.Unsubscribe("1")
	if v := s.Confirmed(); v != 2 {
		t.Fatalf("unexpected confirmed count %v", v)
	}
	s.Unsubscribe("2")
	if v := s.Confirmed(); v != 1 {
		t.Fatalf("unexpected confirmed count %v", v)
	}
	s.Unsubscribe("3")
	if v := s.Confirmed(); v != 0 {
		t.Fatalf("unexpected confirmed count %v", v)
	}
}
