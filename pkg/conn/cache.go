package conn

import (
	"container/list"
	"sync"
	"time"
	"unsafe"

	"github.com/rueian/rueidis/pkg/proto"
)

const (
	EntrySize        = int(unsafe.Sizeof(entry{})) + int(unsafe.Sizeof(&entry{}))
	ElementSize      = int(unsafe.Sizeof(list.Element{})) + int(unsafe.Sizeof(&list.Element{}))
	StringStructSize = int(unsafe.Sizeof(""))
)

type entry struct {
	val  proto.Message
	ttl  time.Time
	key  string
	size int
}

type cache struct {
	mu sync.Mutex

	store map[string]*list.Element
	list  *list.List

	size int
	max  int
}

func NewCache(max int) *cache {
	return &cache{
		max:   max,
		store: make(map[string]*list.Element),
		list:  list.New(),
	}
}

func (c *cache) GetOrPrepare(key string, ttl time.Duration) (v proto.Message) {
	c.mu.Lock()
	ele, ok := c.store[key]
	if ok {
		e := ele.Value.(*entry)
		if e.ttl.After(time.Now()) {
			v = e.val
			c.list.MoveToBack(ele)
		} else {
			delete(c.store, key)
			c.list.Remove(ele)
		}
	} else if c.list != nil {
		c.list.PushBack(&entry{
			key: key,
			ttl: time.Now().Add(ttl),
		})
		c.store[key] = c.list.Back()
	}
	c.mu.Unlock()
	return v
}

func (c *cache) Update(key string, value proto.Message) {
	c.mu.Lock()
	ele, ok := c.store[key]
	if ok {
		e := ele.Value.(*entry)
		e.val = value
		e.size = EntrySize + ElementSize + 2*(StringStructSize+len(key)) + value.Size()

		c.size += e.size
		for c.size > c.max {
			if ele = c.list.Front(); ele != nil {
				e = ele.Value.(*entry)
				delete(c.store, e.key)
				c.list.Remove(ele)
				c.size -= e.size
			}
		}
	}
	c.mu.Unlock()
}

func (c *cache) Delete(keys []proto.Message) {
	c.mu.Lock()
	for _, k := range keys {
		e, ok := c.store[k.String]
		if ok {
			delete(c.store, k.String)
			c.list.Remove(e)
		}
	}
	c.mu.Unlock()
}

func (c *cache) DeleteAll() {
	c.mu.Lock()
	c.store = make(map[string]*list.Element)
	c.list = nil
	c.mu.Unlock()
}
