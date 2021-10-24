package queue

import (
	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

type Queue interface {
	PutOne(m cmds.Completed) chan proto.Result
	PutMulti(m []cmds.Completed) chan proto.Result
	NextCmd() (cmds.Completed, []cmds.Completed)
	NextResultCh() (cmds.Completed, []cmds.Completed, chan proto.Result)
}
