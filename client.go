package rueidis

import (
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/om"
)

type SingleClientOption struct {
	Address    string
	ConnOption ConnOption
}

type SingleClient struct {
	Cmd  *cmds.Builder
	conn conn
}

func newSingleClient(option SingleClientOption, connFn connFn) (*SingleClient, error) {
	c := connFn(option.Address, option.ConnOption)
	if err := c.Dial(); err != nil {
		return nil, err
	}
	return &SingleClient{Cmd: cmds.NewBuilder(), conn: c}, nil
}

func (c *SingleClient) Info() map[string]proto.Message {
	return c.conn.Info()
}

func (c *SingleClient) Do(cmd cmds.Completed) (resp proto.Result) {
	resp = c.conn.Do(cmd)
	c.Cmd.Put(cmd.Commands())
	return resp
}

func (c *SingleClient) DoCache(cmd cmds.Cacheable, ttl time.Duration) (resp proto.Result) {
	resp = c.conn.DoCache(cmd, ttl)
	c.Cmd.Put(cmd.Commands())
	return resp
}

func (c *SingleClient) DedicatedWire(fn func(*DedicatedSingleClient) error) (err error) {
	wire := c.conn.Acquire()
	err = fn(&DedicatedSingleClient{cmd: c.Cmd, wire: wire})
	c.conn.Store(wire)
	return err
}

func (c *SingleClient) NewLuaScript(body string) *Lua {
	return newLuaScript(body, c.eval, c.evalSha)
}

func (c *SingleClient) NewLuaScriptReadOnly(body string) *Lua {
	return newLuaScript(body, c.evalRo, c.evalShaRo)
}

func (c *SingleClient) eval(body string, keys, args []string) proto.Result {
	return c.Do(c.Cmd.Eval().Script(body).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
}

func (c *SingleClient) evalSha(sha string, keys, args []string) proto.Result {
	return c.Do(c.Cmd.Evalsha().Sha1(sha).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
}

func (c *SingleClient) evalRo(body string, keys, args []string) proto.Result {
	return c.Do(c.Cmd.EvalRo().Script(body).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
}

func (c *SingleClient) evalShaRo(sha string, keys, args []string) proto.Result {
	return c.Do(c.Cmd.EvalshaRo().Sha1(sha).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
}

func (c *SingleClient) NewHashRepository(prefix string, schema interface{}) *om.HashRepository {
	return om.NewHashRepository(prefix, schema, &hashObjectSingleClientAdapter{c: c}, func(script string) om.ExecFn {
		return c.NewLuaScript(script).Exec
	})
}

func (c *SingleClient) Close() {
	c.conn.Close()
}

type DedicatedSingleClient struct {
	cmd  *cmds.Builder
	wire wire
}

func (c *DedicatedSingleClient) Do(cmd cmds.Completed) (resp proto.Result) {
	resp = c.wire.Do(cmd)
	c.cmd.Put(cmd.Commands())
	return resp
}

func (c *DedicatedSingleClient) DoMulti(multi ...cmds.Completed) (resp []proto.Result) {
	if len(multi) == 0 {
		return nil
	}
	resp = c.wire.DoMulti(multi...)
	for _, cmd := range multi {
		c.cmd.Put(cmd.Commands())
	}
	return resp
}

type hashObjectSingleClientAdapter struct {
	c *SingleClient
}

func (h *hashObjectSingleClientAdapter) Save(key string, fields map[string]string) error {
	cmd := h.c.Cmd.Hset().Key(key).FieldValue()
	for f, v := range fields {
		cmd = cmd.FieldValue(f, v)
	}
	return h.c.Do(cmd.Build()).Error()
}

func (h *hashObjectSingleClientAdapter) Fetch(key string) (map[string]proto.Message, error) {
	return h.c.Do(h.c.Cmd.Hgetall().Key(key).Build()).ToMap()
}

func (h *hashObjectSingleClientAdapter) FetchCache(key string, ttl time.Duration) (map[string]proto.Message, error) {
	return h.c.DoCache(h.c.Cmd.Hgetall().Key(key).Cache(), ttl).ToMap()
}

func (h *hashObjectSingleClientAdapter) Remove(key string) error {
	return h.c.Do(h.c.Cmd.Del().Key(key).Build()).Error()
}
