package rueidis

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"reflect"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

func TestNewLuaScriptOnePass(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
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
		t.Fatalf("ret mismatch")
	}
}

func TestNewLuaScript(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
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
		t.Fatalf("ret mismatch")
	}
}

func TestNewLuaScriptNoSha(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
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
				t.Fatal("EVALSHA must not be called")
			}
			if reflect.DeepEqual(cmd.Commands(), []string{"EVAL", body, "2", "1", "2", "3", "4"}) {
				return newResult(RedisMessage{typ: '_'}, nil)
			}
			return newResult(strmsg('+', "unexpected"), nil)
		},
	}

	script := NewLuaScriptNoSha(body)

	if err, ok := IsRedisErr(script.Exec(context.Background(), c, k, a).Error()); ok && !err.IsNil() {
		t.Fatalf("ret mismatch")
	}
}

func TestNewLuaScriptReadOnly(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
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
		t.Fatalf("ret mismatch")
	}
}

func TestNewLuaScriptReadOnlyNoSha(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
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
			if reflect.DeepEqual(cmd.Commands(), []string{"EVALSHA_RO", sha, "2", "1", "2", "3", "4"}) {
				t.Fatal("EVALSHA_RO must not be called")
			}
			if reflect.DeepEqual(cmd.Commands(), []string{"EVAL_RO", body, "2", "1", "2", "3", "4"}) {
				return newResult(RedisMessage{typ: '_'}, nil)
			}
			return newResult(strmsg('+', "unexpected"), nil)
		},
	}

	script := NewLuaScriptReadOnlyNoSha(body)

	if err, ok := IsRedisErr(script.Exec(context.Background(), c, k, a).Error()); ok && !err.IsNil() {
		t.Fatalf("ret mismatch")
	}
}

func TestNewLuaScriptExecMultiError(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
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
		t.Fatalf("ret mismatch")
	}
}

func TestNewLuaScriptExecMulti(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
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
		t.Fatalf("ret mismatch")
	}
}

func TestNewLuaScriptExecMultiNoSha(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())

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
				if reflect.DeepEqual(cmd.Commands(), []string{"EVAL", body, "2", "1", "2", "3", "4"}) {
					resp = append(resp, newResult(strmsg('+', "OK"), nil))
				}
			}
			return resp
		},
	}

	script := NewLuaScriptNoSha(body)
	if v, err := script.ExecMulti(context.Background(), c, LuaExec{Keys: k, Args: a})[0].ToString(); err != nil || v != "OK" {
		t.Fatalf("ret mismatch")
	}
}

func TestNewLuaScriptExecMultiRo(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
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
		t.Fatalf("ret mismatch")
	}
}

func TestNewLuaScriptExecMultiRoNoSha(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())

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
				if reflect.DeepEqual(cmd.Commands(), []string{"EVAL_RO", body, "2", "1", "2", "3", "4"}) {
					resp = append(resp, newResult(strmsg('+', "OK"), nil))
				}
			}
			return resp
		},
	}

	script := NewLuaScriptReadOnlyNoSha(body)
	if v, err := script.ExecMulti(context.Background(), c, LuaExec{Keys: k, Args: a})[0].ToString(); err != nil || v != "OK" {
		t.Fatalf("ret mismatch")
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

func (c *client) Receive(_ context.Context, _ Completed, _ func(msg PubSubMessage)) error {
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

func (c *client) DoStream(_ context.Context, _ Completed) (resp RedisResultStream) {
	return RedisResultStream{}
}

func (c *client) DoMultiStream(_ context.Context, _ ...Completed) (resp MultiRedisResultStream) {
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

func TestNewLuaScriptWithLoadSha1(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	scriptLoadCalled := false
	evalshaInvoked := false

	c := &client{
		BFn: func() Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd Completed) (resp RedisResult) {
			commands := cmd.Commands()
			if reflect.DeepEqual(commands, []string{"SCRIPT", "LOAD", body}) {
				scriptLoadCalled = true
				return newResult(strmsg('+', sha), nil)
			}
			if reflect.DeepEqual(commands, []string{"EVALSHA", sha, "2", "1", "2", "3", "4"}) {
				evalshaInvoked = true
				return newResult(strmsg('+', "OK"), nil)
			}
			if reflect.DeepEqual(commands, []string{"EVAL", body, "2", "1", "2", "3", "4"}) {
				t.Fatal("EVAL must not be called when load succeeds")
			}
			return newResult(strmsg('+', "unexpected"), nil)
		},
	}

	script := NewLuaScript(body, WithLoadSHA1(true))

	if v, err := script.Exec(context.Background(), c, k, a).ToString(); err != nil || v != "OK" {
		t.Fatalf("ret mismatch")
	}

	if !scriptLoadCalled {
		t.Fatal("SCRIPT LOAD should have been called")
	}
	if !evalshaInvoked {
		t.Fatal("EVALSHA should have been called")
	}
}

func TestNewLuaScriptReadOnlyWithLoadSha1(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	scriptLoadCalled := false
	evalshaRoInvoked := false

	c := &client{
		BFn: func() Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd Completed) (resp RedisResult) {
			commands := cmd.Commands()
			if reflect.DeepEqual(commands, []string{"SCRIPT", "LOAD", body}) {
				scriptLoadCalled = true
				return newResult(strmsg('+', sha), nil)
			}
			if reflect.DeepEqual(commands, []string{"EVALSHA_RO", sha, "2", "1", "2", "3", "4"}) {
				evalshaRoInvoked = true
				return newResult(strmsg('+', "OK"), nil)
			}
			if reflect.DeepEqual(commands, []string{"EVAL_RO", body, "2", "1", "2", "3", "4"}) {
				t.Fatal("EVAL_RO must not be called when load succeeds")
			}
			return newResult(strmsg('+', "unexpected"), nil)
		},
	}

	script := NewLuaScriptReadOnly(body, WithLoadSHA1(true))

	if v, err := script.Exec(context.Background(), c, k, a).ToString(); err != nil || v != "OK" {
		t.Fatalf("ret mismatch")
	}

	if !scriptLoadCalled {
		t.Fatal("SCRIPT LOAD should have been called")
	}
	if !evalshaRoInvoked {
		t.Fatal("EVALSHA_RO should have been called")
	}
}

func TestNewLuaScriptWithLoadSha1Concurrent(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	var scriptLoadCount atomic.Int64

	c := &client{
		BFn: func() Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd Completed) (resp RedisResult) {
			commands := cmd.Commands()
			if reflect.DeepEqual(commands, []string{"SCRIPT", "LOAD", body}) {
				scriptLoadCount.Add(1)
				return newResult(strmsg('+', sha), nil)
			}
			if reflect.DeepEqual(commands, []string{"EVALSHA", sha, "2", "1", "2", "3", "4"}) {
				return newResult(strmsg('+', "OK"), nil)
			}
			return newResult(strmsg('+', "unexpected"), nil)
		},
	}

	script := NewLuaScript(body, WithLoadSHA1(true))

	// Execute concurrently to verify singleflight works correctly
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			script.Exec(context.Background(), c, k, a)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	// SCRIPT LOAD should only be called once due to singleflight
	if count := scriptLoadCount.Load(); count != 1 {
		t.Fatalf("SCRIPT LOAD should be called exactly once, but was called %d times", count)
	}
}

func TestNewLuaScriptWithLoadSha1ExecMulti(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	scriptLoadCalled := false

	c := &client{
		BFn: func() Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd Completed) (resp RedisResult) {
			commands := cmd.Commands()
			if reflect.DeepEqual(commands, []string{"SCRIPT", "LOAD", body}) {
				scriptLoadCalled = true
				return newResult(strmsg('+', sha), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		},
		DoMultiFn: func(ctx context.Context, multi ...Completed) (resp []RedisResult) {
			for _, cmd := range multi {
				if reflect.DeepEqual(cmd.Commands(), []string{"EVALSHA", sha, "2", "1", "2", "3", "4"}) {
					resp = append(resp, newResult(strmsg('+', "OK"), nil))
				} else if reflect.DeepEqual(cmd.Commands(), []string{"EVAL", body, "2", "1", "2", "3", "4"}) {
					t.Fatal("EVAL should not be called when load succeeds")
				}
			}
			return resp
		},
	}

	script := NewLuaScript(body, WithLoadSHA1(true))
	if v, err := script.ExecMulti(context.Background(), c, LuaExec{Keys: k, Args: a})[0].ToString(); err != nil || v != "OK" {
		t.Fatalf("ret mismatch")
	}

	if !scriptLoadCalled {
		t.Fatal("SCRIPT LOAD should have been called")
	}
}

func BenchmarkLuaScript_Exec(b *testing.B) {
	ctx := context.Background()
	c, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		b.Fatal(err)
	}
	defer c.Close()

	// Benchmark configuration
	const (
		numKeys   = 5
		keyPrefix = "BenchmarkLuaScript_key"
	)

	// More complex Lua script with math operations, conversions, and multiple keys
	script := `
		local sum = 0
		local results = {}

		-- Read multiple keys and perform calculations
		for i = 1, #KEYS do
			local val = redis.call('GET', KEYS[i])
			if val then
				local num = tonumber(val)
				if num then
					sum = sum + num
					table.insert(results, tostring(num * 2))
				else
					table.insert(results, val)
				end
			end
		end

		-- Apply some math operations
		local multiplier = tonumber(ARGV[1]) or 1
		local threshold = tonumber(ARGV[2]) or 100
		local final_result = math.floor(sum * multiplier)

		-- Conditional logic
		if final_result > threshold then
			table.insert(results, 'high:' .. tostring(final_result))
		else
			table.insert(results, 'low:' .. tostring(final_result))
		end

		return results
	`

	// Setup: create test keys with numeric values
	k := make([]string, numKeys)
	for i := 0; i < numKeys; i++ {
		key := keyPrefix + strconv.Itoa(i+1)
		k[i] = key
		if err := c.Do(ctx, c.B().Set().Key(key).Value(strconv.Itoa((i+1)*10)).Build()).Error(); err != nil {
			b.Fatal(err)
		}
	}
	a := []string{"1.5", "200"}

	// Table-driven sub-benchmarks
	cases := []struct {
		name   string
		script *Lua
	}{
		{"Default", NewLuaScript(script)},
		{"LoadSHA1", NewLuaScript(script, WithLoadSHA1(true))},
		{"NoSha", NewLuaScriptNoSha(script)},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			b.SetParallelism(128)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if err := tc.script.Exec(ctx, c, k, a).Error(); err != nil {
						b.Errorf("unexpected error: %v", err)
					}
				}
			})
		})
	}
}

func BenchmarkLuaScript_ExecMulti(b *testing.B) {
	ctx := context.Background()
	c, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		b.Fatal(err)
	}
	defer c.Close()

	// Benchmark configuration
	const (
		numExecs    = 5 // Number of executions in the batch
		keysPerExec = 3 // Number of keys per execution
		keyPrefix   = "BenchmarkLuaScriptExecMulti_key"
	)
	totalKeys := numExecs * keysPerExec

	// More complex Lua script with math operations, conversions, and multiple keys
	script := `
		local sum = 0
		local results = {}

		-- Read multiple keys and perform calculations
		for i = 1, #KEYS do
			local val = redis.call('GET', KEYS[i])
			if val then
				local num = tonumber(val)
				if num then
					sum = sum + num
					table.insert(results, tostring(num * 2))
				else
					table.insert(results, val)
				end
			end
		end

		-- Apply some math operations
		local multiplier = tonumber(ARGV[1]) or 1
		local threshold = tonumber(ARGV[2]) or 100
		local final_result = math.floor(sum * multiplier)

		-- Conditional logic
		if final_result > threshold then
			table.insert(results, 'high:' .. tostring(final_result))
		else
			table.insert(results, 'low:' .. tostring(final_result))
		end

		return results
	`

	// Setup: create test keys with numeric values
	for i := 1; i <= totalKeys; i++ {
		key := keyPrefix + strconv.Itoa(i)
		if err := c.Do(ctx, c.B().Set().Key(key).Value(strconv.Itoa(i*10)).Build()).Error(); err != nil {
			b.Fatal(err)
		}
	}

	// Programmatically build execution batch
	execs := make([]LuaExec, 0, numExecs)
	multipliers := []string{"1.2", "2.0", "0.5", "1.8", "3.0"}
	thresholds := []string{"150", "200", "50", "300", "500"}

	for i := 0; i < numExecs; i++ {
		keys := make([]string, keysPerExec)
		for j := 0; j < keysPerExec; j++ {
			keyNum := i*keysPerExec + j + 1
			keys[j] = keyPrefix + strconv.Itoa(keyNum)
		}
		execs = append(execs, LuaExec{
			Keys: keys,
			Args: []string{multipliers[i%len(multipliers)], thresholds[i%len(thresholds)]},
		})
	}

	// Table-driven sub-benchmarks
	cases := []struct {
		name   string
		script *Lua
	}{
		{"Default", NewLuaScript(script)},
		{"LoadSHA1", NewLuaScript(script, WithLoadSHA1(true))},
		{"NoSha", NewLuaScriptNoSha(script)},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			b.SetParallelism(128)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				results := tc.script.ExecMulti(ctx, c, execs...)
				for _, r := range results {
					if err := r.Error(); err != nil {
						b.Errorf("unexpected error: %v", err)
					}
				}
			}
		})
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

	_, _ = script.Exec(ctx, client, []string{"k1", "k2"}, []string{"a1", "a2"}).ToArray()
}
