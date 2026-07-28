package rueidisaside

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/rueidis"
)

type flight struct {
	done chan struct{}
}

func (c *Client) beginFlight(key, id string) (f *flight, leader bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if f = c.flights[key]; f != nil {
		return f, false
	}
	if c.id != id {
		return nil, false
	}
	f = &flight{done: make(chan struct{})}
	c.flights[key] = f
	return f, true
}

func (c *Client) finishFlight(key string, f *flight) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.flights[key] == f {
		delete(c.flights, key)
		close(f.done)
	}
}

func (c *Client) populate(
	ctx context.Context,
	ttl time.Duration,
	key, id string,
	fn func(ctx context.Context, key string) (val string, err error),
	f *flight,
) (val string, err error) {
	cleanup := true
	defer c.finishFlight(key, f)
	defer func() {
		if cleanup {
			delkey.Exec(context.Background(), c.client, []string{key}, []string{id})
		}
	}()

	if c.useLuaLock {
		val, err = acquireLock.Exec(ctx, c.client, []string{key}, []string{id, strconv.FormatInt(ttl.Milliseconds(), 10)}).ToString()
	} else {
		val, err = c.client.Do(ctx, c.client.B().Set().Key(key).Value(id).Nx().Get().Px(ttl).Build()).ToString()
	}
	if err == nil {
		cleanup = false
		return val, nil
	}
	if !rueidis.IsRedisNil(err) {
		return val, err
	}

	ctx = context.WithValue(ctx, ttlKey, &ttl)
	if val, err = fn(ctx, key); err == nil {
		err = setkey.Exec(ctx, c.client, []string{key}, []string{id, val, strconv.FormatInt(ttl.Milliseconds(), 10)}).Error()
	}
	if err == nil {
		cleanup = false
	}
	return val, err
}
