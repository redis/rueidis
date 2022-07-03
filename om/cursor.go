package om

import (
	"context"
	"errors"

	"github.com/rueian/rueidis"
)

var EndOfCursor = errors.New("end of cursor")

func newAggregateCursor(idx string, client rueidis.Client, resp []rueidis.RedisMessage) *AggregateCursor {
	c := &AggregateCursor{client: client, idx: idx}
	c.n, c.first, c.id = readAggregateResponse(resp)
	return c
}

// AggregateCursor unifies the response of FT.AGGREGATE with or without WITHCURSOR
type AggregateCursor struct {
	client rueidis.Client
	idx    string
	first  []map[string]string
	id     int64
	n      int64
}

// Total return the total numbers of record of the initial FT.AGGREGATE result
func (c *AggregateCursor) Total() int64 {
	return c.n
}

// Read return the partial result from the initial FT.AGGREGATE
// This may invoke FT.CURSOR READ to retrieve further result
func (c *AggregateCursor) Read(ctx context.Context) (partial []map[string]string, err error) {
	if first := c.first; first != nil {
		c.first = nil
		return first, nil
	}
	if c.id == 0 {
		return nil, EndOfCursor
	}
	resp, err := c.client.Do(ctx, c.client.B().FtCursorRead().Index(c.idx).CursorId(c.id).Build()).ToArray()
	if err != nil {
		return nil, err
	}
	_, partial, c.id = readAggregateResponse(resp)
	return
}

// Del uses FT.CURSOR DEL to destroy the cursor
func (c *AggregateCursor) Del(ctx context.Context) (err error) {
	if c.id == 0 {
		return nil
	}
	return c.client.Do(ctx, c.client.B().FtCursorDel().Index(c.idx).CursorId(c.id).Build()).Error()
}

func readAggregateResponse(resp []rueidis.RedisMessage) (n int64, partial []map[string]string, cursor int64) {
	var results []rueidis.RedisMessage
	if resp[0].IsArray() {
		results, _ = resp[0].ToArray()
		cursor, _ = resp[1].ToInt64()
	} else {
		results = resp
	}
	n, _ = results[0].ToInt64()
	partial = make([]map[string]string, len(results[1:]))
	for i, record := range results[1:] {
		partial[i], _ = record.AsStrMap()
	}
	return
}
