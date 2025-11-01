package rueidis

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/redis/rueidis/internal/util"
)

// LuaOption is a functional option for configuring Lua script behavior.
type LuaOption func(*Lua)

// WithLoadSha1 enables loading of SHA-1 from Redis via SCRIPT LOAD.
// When enabled, the SHA-1 hash is not calculated client-side (maintaining FIPS compliance).
// Instead, on first execution, SCRIPT LOAD is called to obtain the SHA-1 from Redis,
// which is then used for EVALSHA commands in subsequent executions.
// This option is only effective when used with NewLuaScriptNoSha or NewLuaScriptReadOnlyNoSha.
// The loading of sha1 is thread-safe and ensures SCRIPT LOAD is called exactly once.
func WithLoadSha1() LuaOption {
	return func(l *Lua) {
		l.lazyLoadSha1 = true
	}
}

// NewLuaScript creates a Lua instance whose Lua.Exec uses EVALSHA and EVAL.
func NewLuaScript(script string) *Lua {
	return newLuaScript(script, false, false)
}

// NewLuaScriptReadOnly creates a Lua instance whose Lua.Exec uses EVALSHA_RO and EVAL_RO.
func NewLuaScriptReadOnly(script string) *Lua {
	return newLuaScript(script, true, false)
}

// NewLuaScriptNoSha creates a Lua instance whose Lua.Exec uses EVAL.
// Sha1 is not calculated, SCRIPT LOAD is not used, no EVALSHA is used.
// The main motivation is to be FIPS compliant, also avoid the tiny chance of SHA-1 collisions.
// This comes with a performance cost as the script is sent to a server every time.
// Use WithLoadSha1() option to enable loading of SHA-1 from Redis for better performance.
func NewLuaScriptNoSha(script string, opts ...LuaOption) *Lua {
	return newLuaScript(script, false, true, opts...)
}

// NewLuaScriptReadOnlyNoSha creates a Lua instance whose Lua.Exec uses EVAL_RO.
// Sha1 is not calculated, SCRIPT LOAD is not used, no EVALSHA_RO is used.
// The main motivation is to be FIPS compliant, also avoid the tiny chance of SHA-1 collisions.
// This comes with a performance cost as the script is sent to a server every time.
// Use WithLoadSha1() option to enable loading of SHA-1 from Redis for better performance.
func NewLuaScriptReadOnlyNoSha(script string, opts ...LuaOption) *Lua {
	return newLuaScript(script, true, true, opts...)
}

func newLuaScript(script string, readonly bool, noSha1 bool, opts ...LuaOption) *Lua {
	var sha1Hex string
	if !noSha1 {
		// It's important to avoid calling sha1 methods since Go will panic in FIPS mode.
		sum := sha1.Sum([]byte(script))
		sha1Hex = hex.EncodeToString(sum[:])
	}
	l := &Lua{
		script:   script,
		sha1:     sha1Hex,
		maxp:     runtime.GOMAXPROCS(0),
		readonly: readonly,
		nosha1:   noSha1,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

// Lua represents a redis lua script. It should be created from the NewLuaScript() or NewLuaScriptReadOnly().
type Lua struct {
	script       string
	sha1         string
	sha1Mu       sync.Mutex
	sha1Call     call
	maxp         int
	readonly     bool
	nosha1       bool
	lazyLoadSha1 bool
}

// Exec the script to the given Client.
// It will first try with the EVALSHA/EVALSHA_RO and then EVAL/EVAL_RO if the first try failed.
// If Lua is initialized with disabled SHA1, it will use EVAL/EVAL_RO without the EVALSHA/EVALSHA_RO attempt.
// If Lua is initialized with lazy SHA-1 loading, it will call SCRIPT LOAD once to obtain the SHA-1 from Redis.
// Cross-slot keys are prohibited if the Client is a cluster client.
func (s *Lua) Exec(ctx context.Context, c Client, keys, args []string) (resp RedisResult) {
	var isNoScript bool
	var scriptSha1 string

	// Determine which SHA-1 to use.
	if s.nosha1 && s.lazyLoadSha1 {
		// Check if SHA-1 is already loaded.
		s.sha1Mu.Lock()
		scriptSha1 = s.sha1
		s.sha1Mu.Unlock()

		// If not loaded yet, use singleflight to load it.
		if scriptSha1 == "" {
			err := s.sha1Call.Do(ctx, func() error {
				result := c.Do(ctx, c.B().ScriptLoad().Script(s.script).Build())
				if shaStr, err := result.ToString(); err == nil {
					s.sha1Mu.Lock()
					s.sha1 = shaStr
					s.sha1Mu.Unlock()
					return nil
				}
				return result.Error()
			})
			if err != nil {
				return newErrResult(err)
			}
			// Reload scriptSha1 after singleflight completes.
			s.sha1Mu.Lock()
			scriptSha1 = s.sha1
			s.sha1Mu.Unlock()
		}
	} else {
		scriptSha1 = s.sha1
	}

	if !s.nosha1 || (s.nosha1 && s.lazyLoadSha1 && scriptSha1 != "") {
		if s.readonly {
			resp = c.Do(ctx, c.B().EvalshaRo().Sha1(scriptSha1).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
		} else {
			resp = c.Do(ctx, c.B().Evalsha().Sha1(scriptSha1).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
		}
		err, isErr := IsRedisErr(resp.Error())
		isNoScript = isErr && err.IsNoScript()
	}
	if (s.nosha1 && (!s.lazyLoadSha1 || scriptSha1 == "")) || isNoScript {
		if s.readonly {
			resp = c.Do(ctx, c.B().EvalRo().Script(s.script).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
		} else {
			resp = c.Do(ctx, c.B().Eval().Script(s.script).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
		}
	}
	return resp
}

// LuaExec is a single execution unit of Lua.ExecMulti.
type LuaExec struct {
	Keys []string
	Args []string
}

// ExecMulti exec the script multiple times by the provided LuaExec to the given Client.
// It will first try SCRIPT LOAD the script to all redis nodes and then exec it with the EVALSHA/EVALSHA_RO.
// If Lua is initialized with disabled SHA1, it will use EVAL/EVAL_RO and no script loading.
// If Lua is initialized with lazy SHA-1 loading, it will obtain the SHA-1 from SCRIPT LOAD response.
// Cross-slot keys within the single LuaExec are prohibited if the Client is a cluster client.
func (s *Lua) ExecMulti(ctx context.Context, c Client, multi ...LuaExec) (resp []RedisResult) {
	// Handle script loading for both regular and lazy load modes.
	if !s.nosha1 || (s.nosha1 && s.lazyLoadSha1) {
		var e atomic.Value
		var sha1Result atomic.Value
		util.ParallelVals(s.maxp, c.Nodes(), func(n Client) {
			result := n.Do(ctx, n.B().ScriptLoad().Script(s.script).Build())
			if err := result.Error(); err != nil {
				e.CompareAndSwap(nil, &errs{error: err})
			} else if s.nosha1 && s.lazyLoadSha1 {
				// Store the first successful SHA-1 result for lazy loading.
				if sha, err := result.ToString(); err == nil {
					sha1Result.CompareAndSwap(nil, sha)
				}
			}
		})
		if err := e.Load(); err != nil {
			resp = make([]RedisResult, len(multi))
			for i := 0; i < len(resp); i++ {
				resp[i] = newErrResult(err.(*errs).error)
			}
			return
		}
		// Set SHA-1 from Redis if lazy loading is enabled.
		if s.nosha1 && s.lazyLoadSha1 {
			if sha := sha1Result.Load(); sha != nil {
				s.sha1Mu.Lock()
				if s.sha1 == "" {
					s.sha1 = sha.(string)
				}
				s.sha1Mu.Unlock()
			}
		}
	}

	s.sha1Mu.Lock()
	scriptSha1 := s.sha1
	s.sha1Mu.Unlock()

	cmds := make(Commands, 0, len(multi))
	for _, m := range multi {
		if !s.nosha1 || (s.nosha1 && s.lazyLoadSha1 && scriptSha1 != "") {
			if s.readonly {
				cmds = append(cmds, c.B().EvalshaRo().Sha1(scriptSha1).Numkeys(int64(len(m.Keys))).Key(m.Keys...).Arg(m.Args...).Build())
			} else {
				cmds = append(cmds, c.B().Evalsha().Sha1(scriptSha1).Numkeys(int64(len(m.Keys))).Key(m.Keys...).Arg(m.Args...).Build())
			}
		} else {
			if s.readonly {
				cmds = append(cmds, c.B().EvalRo().Script(s.script).Numkeys(int64(len(m.Keys))).Key(m.Keys...).Arg(m.Args...).Build())
			} else {
				cmds = append(cmds, c.B().Eval().Script(s.script).Numkeys(int64(len(m.Keys))).Key(m.Keys...).Arg(m.Args...).Build())
			}
		}
	}
	return c.DoMulti(ctx, cmds...)
}
