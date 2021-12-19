package mock

import (
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

type Wire struct {
	DoFn      func(cmd cmds.Completed) proto.Result
	DoCacheFn func(cmd cmds.Cacheable, ttl time.Duration) proto.Result
	DoMultiFn func(multi ...cmds.Completed) []proto.Result
	InfoFn    func() map[string]proto.Message
	ErrorFn   func() error
	CloseFn   func()
}

func (m *Wire) Do(cmd cmds.Completed) proto.Result {
	if m.DoFn != nil {
		return m.DoFn(cmd)
	}
	return proto.Result{}
}

func (m *Wire) DoCache(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
	if m.DoCacheFn != nil {
		return m.DoCacheFn(cmd, ttl)
	}
	return proto.Result{}
}

func (m *Wire) DoMulti(multi ...cmds.Completed) []proto.Result {
	if m.DoMultiFn != nil {
		return m.DoMultiFn(multi...)
	}
	return nil
}

func (m *Wire) Info() map[string]proto.Message {
	if m.InfoFn != nil {
		return m.InfoFn()
	}
	return nil
}

func (m *Wire) Error() error {
	if m.ErrorFn != nil {
		return m.ErrorFn()
	}
	return nil
}

func (m *Wire) Close() {
	if m.CloseFn != nil {
		m.CloseFn()
	}
}
