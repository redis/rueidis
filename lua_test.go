package rueidis

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

func TestNewLuaScriptOnePass(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	c := &client{
		BFn: func() Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd Completed) (resp RedisResult) {
			if reflect.DeepEqual(cmd.Commands(), []string{"EVALSHA", sha, "2", "1", "2", "3", "4"}) {
				return newResult(strmsg('+', "OK"), nil)
			}
			return newResult(strmsg('+', "unexpected"), nil)
		},
	}

	script := NewLuaScript(body)

	if v, err := script.Exec(context.Background(), c, k, a).ToString(); err != nil || v != "OK" {
		t.Fatalf("ret mistmatch")
	}
}

func TestNewLuaScript(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	eval := false

	c := &client{
		BFn: func() Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd Completed) (resp RedisResult) {
			if reflect.DeepEqual(cmd.Commands(), []string{"EVALSHA", sha, "2", "1", "2", "3", "4"}) {
				eval = true
				return newResult(strmsg('-', "NOSCRIPT"), nil)
			}
			if eval && reflect.DeepEqual(cmd.Commands(), []string{"EVAL", body, "2", "1", "2", "3", "4"}) {
				return newResult(RedisMessage{typ: '_'}, nil)
			}
			return newResult(strmsg('+', "unexpected"), nil)
		},
	}

	script := NewLuaScript(body)

	if err, ok := IsRedisErr(script.Exec(context.Background(), c, k, a).Error()); ok && !err.IsNil() {
		t.Fatalf("ret mistmatch")
	}
}

func TestNewLuaScriptReadOnly(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	eval := false

	c := &client{
		BFn: func() Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd Completed) (resp RedisResult) {
			if reflect.DeepEqual(cmd.Commands(), []string{"EVALSHA_RO", sha, "2", "1", "2", "3", "4"}) {
				eval = true
				return newResult(strmsg('-', "NOSCRIPT"), nil)
			}
			if eval && reflect.DeepEqual(cmd.Commands(), []string{"EVAL_RO", body, "2", "1", "2", "3", "4"}) {
				return newResult(RedisMessage{typ: '_'}, nil)
			}
			return newResult(strmsg('+', "unexpected"), nil)
		},
	}

	script := NewLuaScriptReadOnly(body)

	if err, ok := IsRedisErr(script.Exec(context.Background(), c, k, a).Error()); ok && !err.IsNil() {
		t.Fatalf("ret mistmatch")
	}
}

func TestNewLuaScriptExecMultiError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	c := &client{
		BFn: func() Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd Completed) (resp RedisResult) {
			return newResult(strmsg('-', "ANY ERR"), nil)
		},
	}

	script := NewLuaScript(body)
	if script.ExecMulti(context.Background(), c, LuaExec{Keys: k, Args: a})[0].Error().Error() != "ANY ERR" {
		t.Fatalf("ret mistmatch")
	}
}

func TestNewLuaScriptExecMulti(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	c := &client{
		BFn: func() Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd Completed) (resp RedisResult) {
			return newResult(strmsg('+', "OK"), nil)
		},
		DoMultiFn: func(ctx context.Context, multi ...Completed) (resp []RedisResult) {
			for _, cmd := range multi {
				if reflect.DeepEqual(cmd.Commands(), []string{"EVALSHA", sha, "2", "1", "2", "3", "4"}) {
					resp = append(resp, newResult(strmsg('+', "OK"), nil))
				}
			}
			return resp
		},
	}

	script := NewLuaScript(body)
	if v, err := script.ExecMulti(context.Background(), c, LuaExec{Keys: k, Args: a})[0].ToString(); err != nil || v != "OK" {
		t.Fatalf("ret mistmatch")
	}
}

func TestNewLuaScriptExecMultiRo(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	c := &client{
		BFn: func() Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd Completed) (resp RedisResult) {
			return newResult(strmsg('+', "OK"), nil)
		},
		DoMultiFn: func(ctx context.Context, multi ...Completed) (resp []RedisResult) {
			for _, cmd := range multi {
				if reflect.DeepEqual(cmd.Commands(), []string{"EVALSHA_RO", sha, "2", "1", "2", "3", "4"}) {
					resp = append(resp, newResult(strmsg('+', "OK"), nil))
				}
			}
			return resp
		},
	}

	script := NewLuaScriptReadOnly(body)
	if v, err := script.ExecMulti(context.Background(), c, LuaExec{Keys: k, Args: a})[0].ToString(); err != nil || v != "OK" {
		t.Fatalf("ret mistmatch")
	}
}

type client struct {
	BFn            func() Builder
	DoFn           func(ctx context.Context, cmd Completed) (resp RedisResult)
	DoMultiFn      func(ctx context.Context, cmd ...Completed) (resp []RedisResult)
	DoCacheFn      func(ctx context.Context, cmd Cacheable, ttl time.Duration) (resp RedisResult)
	DoMultiCacheFn func(ctx context.Context, cmd ...CacheableTTL) (resp []RedisResult)
	DedicatedFn    func(fn func(DedicatedClient) error) (err error)
	DedicateFn     func() (DedicatedClient, func())
	CloseFn        func()
	ModeFn         func() ClientMode
}

func (c *client) Receive(ctx context.Context, subscribe Completed, fn func(msg PubSubMessage)) error {
	return nil
}

func (c *client) B() Builder {
	if c.BFn != nil {
		return c.BFn()
	}
	return Builder{}
}

func (c *client) Do(ctx context.Context, cmd Completed) (resp RedisResult) {
	if c.DoFn != nil {
		return c.DoFn(ctx, cmd)
	}
	return RedisResult{}
}

func (c *client) DoMulti(ctx context.Context, cmd ...Completed) (resp []RedisResult) {
	if c.DoMultiFn != nil {
		return c.DoMultiFn(ctx, cmd...)
	}
	return nil
}

func (c *client) DoStream(ctx context.Context, cmd Completed) (resp RedisResultStream) {
	return RedisResultStream{}
}

func (c *client) DoMultiStream(ctx context.Context, cmd ...Completed) (resp MultiRedisResultStream) {
	return MultiRedisResultStream{}
}

func (c *client) DoMultiCache(ctx context.Context, cmd ...CacheableTTL) (resp []RedisResult) {
	if c.DoMultiCacheFn != nil {
		return c.DoMultiCacheFn(ctx, cmd...)
	}
	return nil
}

func (c *client) DoCache(ctx context.Context, cmd Cacheable, ttl time.Duration) (resp RedisResult) {
	if c.DoCacheFn != nil {
		return c.DoCacheFn(ctx, cmd, ttl)
	}
	return RedisResult{}
}

func (c *client) Dedicated(fn func(DedicatedClient) error) (err error) {
	if c.DedicatedFn != nil {
		return c.DedicatedFn(fn)
	}
	return nil
}

func (c *client) Dedicate() (DedicatedClient, func()) {
	if c.DedicateFn != nil {
		return c.DedicateFn()
	}
	return nil, nil
}

func (c *client) Nodes() map[string]Client {
	return map[string]Client{"addr": c}
}

func (c *client) Mode() ClientMode {
	return c.ModeFn()
}

func (c *client) Close() {
	if c.CloseFn != nil {
		c.CloseFn()
	}
}

func ExampleLua_exec() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	script := NewLuaScript("return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}")

	script.Exec(ctx, client, []string{"k1", "k2"}, []string{"a1", "a2"}).ToArray()
}
