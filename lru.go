package rueidis

import (
	"container/list"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	entrySize   = int(unsafe.Sizeof(entry{})) + int(unsafe.Sizeof(&entry{}))
	elementSize = int(unsafe.Sizeof(list.Element{})) + int(unsafe.Sizeof(&list.Element{}))
	stringSSize = int(unsafe.Sizeof(""))

	entryMinSize = entrySize + elementSize + stringSSize*2 + messageStructSize

	moveThreshold = uint64(1024 - 1)
)

type cache interface {
	GetOrPrepare(key, cmd string, ttl time.Duration) (v RedisMessage, entry *entry)
	Update(key, cmd string, value RedisMessage, pttl int64)
	Cancel(key, cmd string, value RedisMessage, err error)
	Delete(keys []RedisMessage)
	GetTTL(key string) time.Duration
	FreeAndClose(notice RedisMessage)
}

type entry struct {
	ch   chan struct{}
	key  string
	cmd  string
	val  RedisMessage
	err  error
	size int
}

func (e *entry) Wait() (RedisMessage, error) {
	<-e.ch
	return e.val, e.err
}

type keyCache struct {
	hits  uint64
	miss  uint64
	cache map[string]*list.Element
	ttl   time.Time
}

var _ cache = (*lru)(nil)

type lru struct {
	store map[string]*keyCache
	list  *list.List
	mu    sync.RWMutex
	size  int
	max   int
}

func newLRU(max int) *lru {
	return &lru{
		max:   max,
		store: make(map[string]*keyCache),
		list:  list.New(),
	}
}

func (c *lru) GetOrPrepare(key, cmd string, ttl time.Duration) (v RedisMessage, e *entry) {
	var ok bool
	var store *keyCache
	var now = time.Now()
	var storeTTL time.Time
	var ele, back *list.Element

	c.mu.RLock()
	if store, ok = c.store[key]; ok {
		storeTTL = store.ttl
		if ele, ok = store.cache[cmd]; ok {
			e = ele.Value.(*entry)
			v = e.val
			back = c.list.Back()
		}
	}
	c.mu.RUnlock()

	if e != nil && (v.typ == 0 || storeTTL.After(now)) {
		hits := atomic.AddUint64(&store.hits, 1)
		if ele != back && hits&moveThreshold == 0 {
			c.mu.Lock()
			c.list.MoveToBack(ele)
			c.mu.Unlock()
		}
		return v, e
	}

	v = RedisMessage{}
	e = nil

	c.mu.Lock()
	if store == nil {
		if store, ok = c.store[key]; !ok {
			store = &keyCache{cache: make(map[string]*list.Element), ttl: now.Add(ttl)}
			c.store[key] = store
		}
	}
	if ele, ok = store.cache[cmd]; ok {
		if e = ele.Value.(*entry); e.val.typ == 0 || store.ttl.After(now) {
			atomic.AddUint64(&store.hits, 1)
			v = e.val
			c.list.MoveToBack(ele)
		} else {
			c.list.Remove(ele)
			c.size -= e.size
			e = nil
		}
	}
	if e == nil && c.list != nil {
		atomic.AddUint64(&store.miss, 1)
		c.list.PushBack(&entry{
			key: key,
			cmd: cmd,
			ch:  make(chan struct{}, 1),
		})
		store.ttl = now.Add(ttl)
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
					store := c.store[e.key]
					if delete(store.cache, e.cmd); len(store.cache) == 0 {
						delete(c.store, e.key)
					}
					c.list.Remove(ele)
					c.size -= e.size
				}
				ele = ele.Next()
			}
		}
		if pttl >= 0 {
			// server side ttl should only shorten client side ttl
			if ttl := time.Now().Add(time.Duration(pttl) * time.Millisecond); ttl.Before(store.ttl) {
				store.ttl = ttl
			}
		}
	}
	c.mu.Unlock()
	if ch != nil {
		close(ch)
	}
}

func (c *lru) Cancel(key, cmd string, val RedisMessage, err error) {
	var ch chan struct{}
	c.mu.Lock()
	if store, ok := c.store[key]; ok {
		if ele, ok := store.cache[cmd]; ok {
			if e := ele.Value.(*entry); e.val.typ == 0 {
				e.val = val
				e.err = err
				ch = e.ch
				store := c.store[key]
				if delete(store.cache, cmd); len(store.cache) == 0 {
					delete(c.store, key)
				}
				c.list.Remove(ele)
			}
		}
	}
	c.mu.Unlock()
	if ch != nil {
		close(ch)
	}
}

func (c *lru) GetTTL(key string) (ttl time.Duration) {
	c.mu.Lock()
	if store, ok := c.store[key]; ok && len(store.cache) != 0 {
		ttl = store.ttl.Sub(time.Now())
	}
	if ttl <= 0 {
		ttl = -2
	}
	c.mu.Unlock()
	return
}

func (c *lru) purge(key string, store *keyCache) {
	if store != nil {
		for cmd, ele := range store.cache {
			if e := ele.Value.(*entry); e.val.typ != 0 { // do not delete pending entries
				if delete(store.cache, cmd); len(store.cache) == 0 {
					delete(c.store, key)
				}
				c.list.Remove(ele)
				c.size -= e.size
			}
		}
	}
}

func (c *lru) Delete(keys []RedisMessage) {
	c.mu.Lock()
	if keys == nil {
		for key, store := range c.store {
			c.purge(key, store)
		}
	} else {
		for _, k := range keys {
			c.purge(k.string, c.store[k.string])
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
