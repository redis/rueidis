package cache

import (
	"container/list"
	"sync"
	"time"
	"unsafe"

	"github.com/rueian/rueidis/internal/proto"
)

const (
	entrySize   = int(unsafe.Sizeof(Entry{})) + int(unsafe.Sizeof(&Entry{}))
	elementSize = int(unsafe.Sizeof(list.Element{})) + int(unsafe.Sizeof(&list.Element{}))
	stringSSize = int(unsafe.Sizeof(""))

	EntryMinSize = entrySize + elementSize + stringSSize*2 + proto.MessageStructSize
)

type Entry struct {
	val  proto.Message
	key  string
	cmd  string
	ch   chan struct{}
	size int
}

func (e *Entry) Wait() proto.Message {
	<-e.ch
	return e.val
}

type keyCache struct {
	cache map[string]*list.Element
	ttl   time.Time
}

type LRU struct {
	mu sync.Mutex

	store map[string]keyCache
	list  *list.List

	size int
	max  int
}

func NewLRU(max int) *LRU {
	return &LRU{
		max:   max,
		store: make(map[string]keyCache),
		list:  list.New(),
	}
}

func (c *LRU) GetOrPrepare(key, cmd string, ttl time.Duration) (v proto.Message, entry *Entry) {
	c.mu.Lock()
	store, ok := c.store[key]
	if !ok {
		store = keyCache{cache: make(map[string]*list.Element), ttl: time.Now().Add(ttl)}
		c.store[key] = store
	}
	if ele, ok := store.cache[cmd]; ok {
		if entry = ele.Value.(*Entry); store.ttl.After(time.Now()) {
			v = entry.val
			c.list.MoveToBack(ele)
		} else {
			entry = nil
			c.list.Remove(ele)
			store.ttl = time.Now().Add(ttl)
		}
	}
	if entry == nil && c.list != nil {
		c.list.PushBack(&Entry{
			key: key,
			cmd: cmd,
			ch:  make(chan struct{}, 1),
		})
		store.cache[cmd] = c.list.Back()
	}
	c.mu.Unlock()
	return v, entry
}

func (c *LRU) Update(key, cmd string, value proto.Message, pttl int64) {
	var ch chan struct{}
	c.mu.Lock()
	if store, ok := c.store[key]; ok {
		if ele, ok := store.cache[cmd]; ok {
			if e := ele.Value.(*Entry); e.val.Type == 0 {
				e.val = value
				e.size = entrySize + elementSize + 2*(stringSSize+len(key)+stringSSize+len(cmd)) + value.ApproximateSize()
				c.size += e.size
				ch = e.ch
			}

			ele = c.list.Front()
			for c.size > c.max && ele != nil {
				e := ele.Value.(*Entry)
				if e.val.Type != 0 { // do not delete pending entries
					delete(c.store[e.key].cache, e.cmd)
					c.list.Remove(ele)
					c.size -= e.size
				}
				ele = ele.Next()
			}
		}
		if pttl == -2 {
			store.ttl = time.Time{}
		} else {
			if pttl != -1 {
				if ttl := time.Now().Add(time.Duration(pttl) * time.Millisecond); ttl.Before(store.ttl) {
					store.ttl = ttl
				}
			}
		}
	}
	c.mu.Unlock()
	if ch != nil {
		close(ch)
	}
}

func (c *LRU) Delete(keys []proto.Message) {
	c.mu.Lock()
	for _, k := range keys {
		e, ok := c.store[k.String]
		if ok {
			delete(c.store, k.String)
			for _, ele := range e.cache {
				c.list.Remove(ele)
			}
		}
	}
	c.mu.Unlock()
}

func (c *LRU) DeleteAll() {
	c.mu.Lock()
	c.store = make(map[string]keyCache)
	c.list = nil
	c.mu.Unlock()
}
