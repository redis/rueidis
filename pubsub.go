package rueidis

import "sync"

// PubSubMessage represent a pubsub message from redis
type PubSubMessage struct {
	// Pattern is only available with pmessage.
	Pattern string
	// Channel is the channel the message belongs to
	Channel string
	// Message is the message content
	Message string
}

// PubSubSubscription represent a pubsub "subscribe", "unsubscribe", "psubscribe" or "punsubscribe" event.
type PubSubSubscription struct {
	// Kind is "subscribe", "unsubscribe", "psubscribe" or "punsubscribe"
	Kind string
	// Channel is the event subject.
	Channel string
	// Count is the current number of subscriptions for connection.
	Count int64
}

// PubSubHooks can be registered into DedicatedClient to process pubsub messages without using Client.Receive
type PubSubHooks struct {
	// OnMessage will be called when receiving "message" and "pmessage" event.
	OnMessage func(m PubSubMessage)
	// OnSubscription will be called when receiving "subscribe", "unsubscribe", "psubscribe" and "punsubscribe" event.
	OnSubscription func(s PubSubSubscription)
}

func (h *PubSubHooks) isZero() bool {
	return h.OnMessage == nil && h.OnSubscription == nil
}

func newSubs() *subs {
	return &subs{chs: make(map[string]chs), sub: make(map[int]*sub)}
}

type subs struct {
	mu  sync.RWMutex
	chs map[string]chs
	sub map[int]*sub
	cnt int
}

type chs struct {
	sub map[int]*sub
	cnf bool
}

type sub struct {
	cs []string
	ch chan PubSubMessage
}

func (s *subs) Publish(channel string, msg PubSubMessage) {
	s.mu.RLock()
	if s.chs != nil {
		for _, sb := range s.chs[channel].sub {
			sb.ch <- msg
		}
	}
	s.mu.RUnlock()
}

func (s *subs) Subscribe(channels []string) (ch chan PubSubMessage, cancel func()) {
	s.mu.Lock()
	if s.chs != nil {
		s.cnt++
		ch = make(chan PubSubMessage, 16)
		sb := &sub{cs: channels, ch: ch}
		id := s.cnt
		s.sub[id] = sb
		for _, channel := range channels {
			c := s.chs[channel].sub
			if c == nil {
				c = make(map[int]*sub)
				s.chs[channel] = chs{sub: c, cnf: false}
			}
			c[id] = sb
		}
		cancel = func() {
			go func() {
				for range ch {
				}
			}()
			s.mu.Lock()
			if s.chs != nil {
				s.remove(id)
			}
			s.mu.Unlock()
		}
	}
	s.mu.Unlock()
	return ch, cancel
}

func (s *subs) remove(id int) {
	if sb := s.sub[id]; sb != nil {
		for _, channel := range sb.cs {
			if c := s.chs[channel].sub; c != nil {
				delete(c, id)
			}
		}
		close(sb.ch)
		delete(s.sub, id)
	}
}

func (s *subs) Confirm(channel string) {
	s.mu.Lock()
	if s.chs != nil {
		c := s.chs[channel]
		c.cnf = true
		s.chs[channel] = c
	}
	s.mu.Unlock()
}

func (s *subs) Confirmed() (count int) {
	s.mu.RLock()
	if s.chs != nil {
		for _, c := range s.chs {
			if c.cnf {
				count++
			}
		}
	}
	s.mu.RUnlock()
	return
}

func (s *subs) Unsubscribe(channel string) {
	s.mu.Lock()
	if s.chs != nil {
		for id := range s.chs[channel].sub {
			s.remove(id)
		}
		delete(s.chs, channel)
	}
	s.mu.Unlock()
}

func (s *subs) Close() {
	var sbs map[int]*sub
	s.mu.Lock()
	sbs = s.sub
	s.chs = nil
	s.sub = nil
	s.mu.Unlock()
	for _, sb := range sbs {
		close(sb.ch)
	}
}
