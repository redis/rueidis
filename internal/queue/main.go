package queue

import "github.com/rueian/rueidis/internal/proto"

type Queue interface {
	PutOne(m []string) chan proto.Result
	PutMulti(m [][]string) chan proto.Result
	NextCmd() [][]string
	NextResultCh() ([][]string, chan proto.Result)
}
