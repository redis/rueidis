package rueidis

import "testing"

func TestSubs_Publish(t *testing.T) {
	t.Run("without subs", func(t *testing.T) {
		s := newSubs()
		s.Publish("aa", PubSubMessage{}) // just no block
	})

	t.Run("with multiple subs", func(t *testing.T) {
		s := newSubs()
		id1, ch1 := s.Subscribe([]string{"a"})
		id2, ch2 := s.Subscribe([]string{"a"})
		id3, ch3 := s.Subscribe([]string{"b"})
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
			s.Remove(id1)
		}
		for m := range ch2 {
			if m != m1 {
				t.Fatalf("unexpected msg %v", m)
			}
			s.Remove(id2)
		}
		for m := range ch3 {
			if m != m2 {
				t.Fatalf("unexpected msg %v", m)
			}
			s.Remove(id3)
		}
	})
}

func TestSubs_Unsubscribe(t *testing.T) {
	s := newSubs()
	_, ch := s.Subscribe([]string{"1", "2"})
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
