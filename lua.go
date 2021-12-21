package rueidis

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/rueian/rueidis/internal/proto"
)

type evalFn func(script string, keys []string, args []string) proto.Result

type Lua struct {
	script string
	sha1   string

	eval    evalFn
	evalSha evalFn
}

func (s *Lua) Exec(keys, args []string) proto.Result {
	r := s.evalSha(s.sha1, keys, args)
	if err := r.RedisError(); err != nil && err.IsNoScript() {
		r = s.eval(s.script, keys, args)
	}
	return r
}

func newLuaScript(body string, eval, evalSha evalFn) *Lua {
	sum := sha1.Sum([]byte(body))
	return &Lua{
		script:  body,
		sha1:    hex.EncodeToString(sum[:]),
		eval:    eval,
		evalSha: evalSha,
	}
}
