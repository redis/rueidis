package om

import (
	"context"
	"errors"

	"github.com/redis/rueidis"
)

var EndOfCursor = errors.New("end of cursor")

func newAggregateCursor(idx string, client rueidis.Client, first []map[string]string, cursor, total int64) *AggregateCursor {
	return &AggregateCursor{client: client, idx: idx, first: first, id: cursor, n: total}
}

// AggregateCursor unifies the response of FT.AGGREGATE with or without WITHCURSOR
type AggregateCursor struct {
	client rueidis.Client
	idx    string
	first  []map[string]string
	id     int64
	n      int64
}

// Total return the total numbers of records of the initial FT.AGGREGATE result
func (c *AggregateCursor) Total() int64 {
	return c.n
}

// Read return the partial result from the initial FT.AGGREGATE
// This may invoke FT.CURSOR READ to retrieve a further result
func (c *AggregateCursor) Read(ctx context.Context) (partial []map[string]string, err error) {
	if first := c.first; first != nil {
		c.first = nil
		return first, nil
	}
	if c.id == 0 {
		return nil, EndOfCursor
	}
	c.id, _, partial, err = c.client.Do(ctx, c.client.B().FtCursorRead().Index(c.idx).CursorId(c.id).Build()).AsFtAggregateCursor()
	return
}

// Del uses FT.CURSOR DEL to destroy the cursor
func (c *AggregateCursor) Del(ctx context.Context) (err error) {
	if c.id == 0 {
		return nil
	}
	return c.client.Do(ctx, c.client.B().FtCursorDel().Index(c.idx).CursorId(c.id).Build()).Error()
}
