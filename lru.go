package rueidis

import (
	"container/list"
	"sync"
	"time"
	"unsafe"
)

const (
	entrySize   = int(unsafe.Sizeof(entry{})) + int(unsafe.Sizeof(&entry{}))
	elementSize = int(unsafe.Sizeof(list.Element{})) + int(unsafe.Sizeof(&list.Element{}))
	stringSSize = int(unsafe.Sizeof(""))

	entryMinSize = entrySize + elementSize + stringSSize*2 + messageStructSize
)

type cache interface {
	GetOrPrepare(key, cmd string, ttl time.Duration) (v RedisMessage, entry *entry)
	Update(key, cmd string, value RedisMessage, pttl int64)
	Delete(keys []RedisMessage)
	FreeAndClose(notice RedisMessage)
}

type entry struct {
	val  RedisMessage
	key  string
	cmd  string
	ch   chan struct{}
	size int
}

func (e *entry) Wait() RedisMessage {
	<-e.ch
	return e.val
}

type keyCache struct {
	cache map[string]*list.Element
	ttl   time.Time
}

var _ cache = (*lru)(nil)

type lru struct {
	mu sync.Mutex

	store map[string]*keyCache
	list  *list.List

	size int
	max  int
}

func newLRU(max int) *lru {
	return &lru{
		max:   max,
		store: make(map[string]*keyCache),
		list:  list.New(),
	}
}

func (c *lru) GetOrPrepare(key, cmd string, ttl time.Duration) (v RedisMessage, e *entry) {
	c.mu.Lock()
	store, ok := c.store[key]
	if !ok {
		store = &keyCache{cache: make(map[string]*list.Element), ttl: time.Now().Add(ttl)}
		c.store[key] = store
	}
	if ele, ok := store.cache[cmd]; ok {
		if e = ele.Value.(*entry); e.val.typ == 0 || store.ttl.After(time.Now()) {
			v = e.val
			c.list.MoveToBack(ele)
		} else {
			e = nil
			c.list.Remove(ele)
			store.ttl = time.Now().Add(ttl)
		}
	}
	if e == nil && c.list != nil {
		c.list.PushBack(&entry{
			key: key,
			cmd: cmd,
			ch:  make(chan struct{}, 1),
		})
		store.cache[cmd] = c.list.Back()
	}
	c.mu.Unlock()
	return v, e
}

func (c *lru) Update(key, cmd string, value RedisMessage, pttl int64) {
	var ch chan struct{}
	c.mu.Lock()
	if store, ok := c.store[key]; ok {
		if ele, ok := store.cache[cmd]; ok {
			if e := ele.Value.(*entry); e.val.typ == 0 {
				e.val = value
				e.size = entrySize + elementSize + 2*(stringSSize+len(key)+stringSSize+len(cmd)) + value.approximateSize()
				c.size += e.size
				ch = e.ch
			}

			ele = c.list.Front()
			for c.size > c.max && ele != nil {
				if e := ele.Value.(*entry); e.val.typ != 0 { // do not delete pending entries
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
				ttl := time.Now().Add(time.Duration(pttl) * time.Millisecond)
				if ttl.Before(store.ttl) {
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

func (c *lru) Delete(keys []RedisMessage) {
	c.mu.Lock()
	for _, k := range keys {
		if store, ok := c.store[k.string]; ok {
			for cmd, ele := range store.cache {
				if e := ele.Value.(*entry); e.val.typ != 0 { // do not delete pending entries
					delete(store.cache, cmd)
					c.list.Remove(ele)
					c.size -= e.size
				}
			}
		}
	}
	c.mu.Unlock()
}

func (c *lru) FreeAndClose(notice RedisMessage) {
	c.mu.Lock()
	for _, store := range c.store {
		for _, ele := range store.cache {
			if e := ele.Value.(*entry); e.val.typ == 0 {
				e.val = notice
				close(e.ch)
			}
		}
	}
	c.store = make(map[string]*keyCache)
	c.list = nil
	c.mu.Unlock()
}
