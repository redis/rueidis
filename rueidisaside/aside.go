package rueidisaside

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/redis/rueidis"
)

type ClientOption struct {
	// ClientBuilder can be used to modify rueidis.Client used by Locker
	ClientBuilder func(option rueidis.ClientOption) (rueidis.Client, error)
	ClientOption  rueidis.ClientOption
	ClientTTL     time.Duration // TTL for the client marker, refreshed every 1/2 TTL. Defaults to 10s. The marker allows other clients to know if this client is still alive.
	UseLuaLock    bool
}

type CacheAsideClient interface {
	Get(ctx context.Context, ttl time.Duration, key string, fn func(ctx context.Context, key string) (val string, err error)) (val string, err error)
	Del(ctx context.Context, key string) error
	Client() rueidis.Client
	Close()
}

func NewClient(option ClientOption) (cc CacheAsideClient, err error) {
	if option.ClientTTL <= 0 {
		option.ClientTTL = 10 * time.Second
	}
	ca := &Client{
		waits:      make(map[string]chan struct{}),
		ttl:        option.ClientTTL,
		useLuaLock: option.UseLuaLock,
	}
	option.ClientOption.OnInvalidations = ca.onInvalidation
	if option.ClientBuilder != nil {
		ca.client, err = option.ClientBuilder(option.ClientOption)
	} else {
		ca.client, err = rueidis.NewClient(option.ClientOption)
	}
	if err != nil {
		return nil, err
	}
	ca.ctx, ca.cancel = context.WithCancel(context.Background())
	return ca, nil
}

type Client struct {
	client     rueidis.Client
	ctx        context.Context
	waits      map[string]chan struct{}
	cancel     context.CancelFunc
	id         string
	ttl        time.Duration
	mu         sync.Mutex
	useLuaLock bool
}

func (c *Client) onInvalidation(messages []rueidis.RedisMessage) {
	var id string
	c.mu.Lock()
	if messages == nil {
		id = c.id
		c.id = ""
		for _, ch := range c.waits {
			close(ch)
		}
		c.waits = make(map[string]chan struct{})
	} else {
		for _, m := range messages {
			key, _ := m.ToString()
			if ch := c.waits[key]; ch != nil {
				close(ch)
				delete(c.waits, key)
			}
		}
	}
	c.mu.Unlock()
	if id != "" {
		c.client.Do(context.Background(), c.client.B().Del().Key(id).Build())
	}
}

func (c *Client) register(key string) (ch chan struct{}) {
	c.mu.Lock()
	if ch = c.waits[key]; ch == nil {
		ch = make(chan struct{})
		c.waits[key] = ch
	}
	c.mu.Unlock()
	return
}

func (c *Client) refresh(id string) {
	for interval := c.ttl / 2; ; {
		select {
		case <-time.After(interval):
			c.mu.Lock()
			id2 := c.id
			c.mu.Unlock()
			if id2 != id {
				return // client id has changed, abort this goroutine
			}
			c.client.Do(c.ctx, c.client.B().Set().Key(id).Value("").Px(c.ttl).Build())
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Client) keepalive() (id string, err error) {
	c.mu.Lock()
	id = c.id
	c.mu.Unlock()
	if id == "" {
		id = PlaceholderPrefix + randStr()
		if err = c.client.Do(c.ctx, c.client.B().Set().Key(id).Value("").Px(c.ttl).Build()).Error(); err == nil {
			c.mu.Lock()
			if c.id == "" {
				c.id = id
				go c.refresh(id)
			} else {
				id = c.id
			}
			c.mu.Unlock()
		}
	}
	return id, err
}

// randStr generates a 24-byte long, random string.
func randStr() string {
	b := make([]byte, 24)
	binary.LittleEndian.PutUint64(b[12:], rand.Uint64())
	binary.LittleEndian.PutUint32(b[20:], rand.Uint32())
	hex.Encode(b, b[12:])

	return unsafe.String(unsafe.SliceData(b), len(b))
}

func (c *Client) Get(ctx context.Context, ttl time.Duration, key string, fn func(ctx context.Context, key string) (val string, err error)) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, ttl)
	defer cancel()

retry:
	wait := c.register(key)
	resp := c.client.DoCache(ctx, c.client.B().Get().Key(key).Cache(), ttl)
	val, err := resp.ToString()

	if rueidis.IsRedisNil(err) && fn != nil { // cache miss, prepare to populate the value by fn()
		var id string
		if id, err = c.keepalive(); err == nil { // acquire client id
			if c.useLuaLock {
				val, err = acquireLock.Exec(ctx, c.client, []string{key}, []string{id, strconv.FormatInt(ttl.Milliseconds(), 10)}).ToString()
			} else {
				val, err = c.client.Do(ctx, c.client.B().Set().Key(key).Value(id).Nx().Get().Px(ttl).Build()).ToString()
			}

			if rueidis.IsRedisNil(err) { // successfully set client id on the key as a lock
				if val, err = fn(ctx, key); err == nil {
					err = setkey.Exec(ctx, c.client, []string{key}, []string{id, val, strconv.FormatInt(ttl.Milliseconds(), 10)}).Error()
				}
				if err != nil { // failed to populate the value, release the lock.
					delkey.Exec(context.Background(), c.client, []string{key}, []string{id})
				}
			}
		}
	}

	if err != nil {
		return val, err
	}

	if strings.HasPrefix(val, PlaceholderPrefix) {
		ph := c.register(val)
		err = c.client.DoCache(ctx, c.client.B().Get().Key(val).Cache(), c.ttl).Error()
		if rueidis.IsRedisNil(err) {
			// the client who held the lock has gone, release the lock.
			delkey.Exec(context.Background(), c.client, []string{key}, []string{val})
			goto retry
		}
		val = ""
		if err == nil {
			select {
			case <-ph:
			case <-wait:
			case <-ctx.Done():
				return "", ctx.Err()
			}
			goto retry
		}
	}

	return val, err
}

func (c *Client) Del(ctx context.Context, key string) error {
	return c.client.Do(ctx, c.client.B().Del().Key(key).Build()).Error()
}

// Client exports the underlying rueidis.Client
func (c *Client) Client() rueidis.Client {
	return c.client
}

func (c *Client) Close() {
	c.cancel()
	c.mu.Lock()
	id := c.id
	c.mu.Unlock()
	if id != "" {
		c.client.Do(context.Background(), c.client.B().Del().Key(c.id).Build())
	}
	c.client.Close()
}

const PlaceholderPrefix = "rueidisid:"

var (
	delkey      = rueidis.NewLuaScript(`if redis.call("GET",KEYS[1]) == ARGV[1] then return redis.call("DEL",KEYS[1]) else return 0 end`)
	setkey      = rueidis.NewLuaScript(`if redis.call("GET",KEYS[1]) == ARGV[1] then return redis.call("SET",KEYS[1],ARGV[2],"PX",ARGV[3]) else return 0 end`)
	acquireLock = rueidis.NewLuaScript(`if redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2]) then return nil else return redis.call("GET", KEYS[1]) end`)
)
