package queue

import (
	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

type Queue interface {
	PutOne(m cmds.Completed) chan proto.Result
	PutMulti(m []cmds.Completed) chan proto.Result
	NextWriteCmd() (cmds.Completed, []cmds.Completed, chan proto.Result)
	NextResultCh() (cmds.Completed, []cmds.Completed, chan proto.Result)
}
