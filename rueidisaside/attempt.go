package rueidisaside

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/cmds"
)

const attemptGuardPrefix = "rueidisaside:attempt:"

const (
	acquireRevoked int64 = iota
	acquireSucceeded
	acquireOccupied
)

type attempt struct {
	guard string
	token string
}

func (c *Client) populate(
	ctx context.Context,
	ttl time.Duration,
	key, id string,
	fn func(ctx context.Context, key string) (val string, err error),
	f *flight,
) (val string, err error) {
	var cleanup bool
	a := c.newAttempt(key, id)
	defer c.finishFlight(key, id, f)
	defer func() {
		if cleanup {
			// Cleanup must finish before the flight is released. If Redis cannot
			// confirm it, rotate the client id so a late command cannot match a
			// subsequent flight from this client.
			cleanupCtx, cancel := context.WithTimeout(context.Background(), c.ttl)
			cleanupErr := revokeAttempt.Exec(cleanupCtx, c.client, []string{key, a.guard}, []string{id, a.token}).Error()
			cancel()
			if cleanupErr != nil {
				c.abandonGeneration(id)
				if err == nil {
					err = cleanupErr
				}
			}
		}
	}()

	deadline, _ := ctx.Deadline()
	guardTTL := time.Until(deadline)
	if guardTTL <= 0 {
		if err = ctx.Err(); err == nil {
			err = context.DeadlineExceeded
		}
		return "", err
	}
	if guardTTL < time.Millisecond {
		guardTTL = time.Millisecond
	}
	if err = c.client.Do(ctx, c.client.B().Set().Key(a.guard).Value(a.token).Nx().Px(guardTTL).Build()).Error(); err != nil {
		return "", err
	}
	cleanup = true

	var acquired bool
	if val, acquired, err = c.acquire(ctx, key, id, a); err != nil {
		return val, err
	}
	if !acquired {
		cleanup = false
		return val, nil
	}

	ctx = context.WithValue(ctx, ttlKey, &ttl)
	if val, err = fn(ctx, key); err != nil {
		return val, err
	}
	var status int64
	status, err = guardedSet.Exec(
		ctx,
		c.client,
		[]string{key, a.guard},
		[]string{id, a.token, val, strconv.FormatInt(ttl.Milliseconds(), 10)},
	).AsInt64()
	if err == nil && status == 1 {
		cleanup = false
	}
	return val, err
}

func (c *Client) acquire(ctx context.Context, key, id string, a attempt) (val string, acquired bool, err error) {
	script := guardedAcquire
	if c.useLuaLock {
		script = guardedAcquireLegacy
	}
	resp, err := script.Exec(
		ctx,
		c.client,
		[]string{key, a.guard},
		[]string{id, a.token},
	).ToArray()
	if err != nil {
		return "", false, err
	}
	if len(resp) == 0 {
		return "", false, fmt.Errorf("rueidisaside: empty acquire response")
	}
	status, err := resp[0].AsInt64()
	if err != nil {
		return "", false, fmt.Errorf("rueidisaside: invalid acquire response: %w", err)
	}
	switch status {
	case acquireRevoked:
		if err = ctx.Err(); err == nil {
			err = context.DeadlineExceeded
		}
		return "", false, err
	case acquireSucceeded:
		return "", true, nil
	case acquireOccupied:
		if len(resp) != 2 {
			return "", false, fmt.Errorf("rueidisaside: occupied acquire response has %d elements", len(resp))
		}
		val, err = resp[1].ToString()
		return val, false, err
	default:
		return "", false, fmt.Errorf("rueidisaside: unknown acquire status %d", status)
	}
}

func (c *Client) newAttempt(key, id string) attempt {
	n := c.attempts.Add(1)
	token := id + ":" + strconv.FormatUint(n, 10)
	return attempt{
		guard: attemptGuardKey(key, token),
		token: token,
	}
}

func (c *Client) abandonGeneration(id string) {
	c.mu.Lock()
	if c.id == id {
		c.id = ""
	}
	c.mu.Unlock()
}

var (
	guardedAcquire = rueidis.NewLuaScript(`
if redis.call("GET", KEYS[2]) ~= ARGV[2] then
	return {0}
end
local ttl = redis.call("PTTL", KEYS[2])
if ttl <= 0 then
	return {0}
end
local current = redis.call("SET", KEYS[1], ARGV[1], "NX", "GET", "PX", ttl)
if current then
	redis.call("DEL", KEYS[2])
	return {2, current}
end
return {1}
`)
	guardedAcquireLegacy = rueidis.NewLuaScript(`
if redis.call("GET", KEYS[2]) ~= ARGV[2] then
	return {0}
end
local ttl = redis.call("PTTL", KEYS[2])
if ttl <= 0 then
	return {0}
end
local current = redis.call("GET", KEYS[1])
if current then
	redis.call("DEL", KEYS[2])
	return {2, current}
end
redis.call("SET", KEYS[1], ARGV[1], "PX", ttl)
return {1}
`)
	guardedSet = rueidis.NewLuaScript(`
if redis.call("GET", KEYS[2]) ~= ARGV[2] then
	return 0
end
if redis.call("GET", KEYS[1]) == ARGV[1] then
	redis.call("SET", KEYS[1], ARGV[3], "PX", ARGV[4])
	redis.call("DEL", KEYS[2])
	return 1
end
redis.call("DEL", KEYS[2])
return 0
`)
	revokeAttempt = rueidis.NewLuaScriptRetryable(`
if redis.call("GET", KEYS[2]) == ARGV[2] then
	redis.call("DEL", KEYS[2])
end
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
end
return 0
`)
)

const slotTagAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

var (
	slotTags     [16384][3]byte
	slotTagsOnce sync.Once
)

func attemptGuardKey(key, token string) string {
	slotTagsOnce.Do(initSlotTags)
	tag := slotTags[cmds.Slot(key)]
	return attemptGuardPrefix + "{" + string(tag[:]) + "}:" + token
}

func initSlotTags() {
	// Lua requires the cache key and its attempt guard to share a cluster slot.
	// Build one printable hash tag for every possible slot.
	remaining := len(slotTags)
	for i := 0; i < len(slotTagAlphabet) && remaining != 0; i++ {
		for j := 0; j < len(slotTagAlphabet) && remaining != 0; j++ {
			for k := 0; k < len(slotTagAlphabet) && remaining != 0; k++ {
				tag := [3]byte{slotTagAlphabet[i], slotTagAlphabet[j], slotTagAlphabet[k]}
				slot := cmds.Slot(string(tag[:]))
				if slotTags[slot][0] == 0 {
					slotTags[slot] = tag
					remaining--
				}
			}
		}
	}
	if remaining != 0 {
		panic("rueidisaside: unable to generate Redis slot tags")
	}
}
