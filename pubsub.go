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

func newSubs() *subs {
	return &subs{chs: make(map[string]map[int]*sub), sub: make(map[int]*sub)}
}

type subs struct {
	mu  sync.RWMutex
	chs map[string]map[int]*sub
	sub map[int]*sub
	cnt int
}

type sub struct {
	cs []string
	ch chan PubSubMessage
}

func (s *subs) Publish(channel string, msg PubSubMessage) {
	s.mu.RLock()
	if s.chs != nil {
		for _, sb := range s.chs[channel] {
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
			c := s.chs[channel]
			if c == nil {
				c = make(map[int]*sub)
				s.chs[channel] = c
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
			if c := s.chs[channel]; c != nil {
				delete(c, id)
				if len(c) == 0 {
					delete(s.chs, channel)
				}
			}
		}
		close(sb.ch)
		delete(s.sub, id)
	}
}

func (s *subs) Unsubscribe(channel string) {
	s.mu.Lock()
	if s.chs != nil {
		for id := range s.chs[channel] {
			s.remove(id)
		}
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
