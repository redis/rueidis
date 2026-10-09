package rueidiscompatrate

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidiscompat"
	"github.com/redis/rueidis/rueidislimiter"
)

const redisPrefix = "rate:"

var errNilClient = errors.New("rueidiscompatrate: client is nil")

type Limit struct {
	Rate   int
	Burst  int
	Period time.Duration
}

func (l Limit) String() string {
	return fmt.Sprintf("%d req/%s (burst %d)", l.Rate, fmtDur(l.Period), l.Burst)
}

func (l Limit) IsZero() bool {
	return l == Limit{}
}

func fmtDur(d time.Duration) string {
	switch d {
	case time.Second:
		return "s"
	case time.Minute:
		return "m"
	case time.Hour:
		return "h"
	}
	return d.String()
}

func PerSecond(rate int) Limit {
	return Limit{
		Rate:   rate,
		Period: time.Second,
		Burst:  rate,
	}
}

func PerMinute(rate int) Limit {
	return Limit{
		Rate:   rate,
		Period: time.Minute,
		Burst:  rate,
	}
}

func PerHour(rate int) Limit {
	return Limit{
		Rate:   rate,
		Period: time.Hour,
		Burst:  rate,
	}
}

type Result struct {
	Limit      Limit
	Allowed    int
	Remaining  int
	RetryAfter time.Duration
	ResetAfter time.Duration
}

type Limiter struct {
	client rueidis.Client
}

func NewLimiter(rdb rueidiscompat.Cmdable) *Limiter {
	if rdb == nil {
		return &Limiter{}
	}
	return &Limiter{
		client: rdb.Client(),
	}
}

func NewLimiterFromClient(client rueidis.Client) *Limiter {
	return &Limiter{
		client: client,
	}
}

func (l Limiter) Allow(ctx context.Context, key string, limit Limit) (*Result, error) {
	return l.AllowN(ctx, key, limit, 1)
}

func (l Limiter) AllowN(ctx context.Context, key string, limit Limit, n int) (*Result, error) {
	if l.client == nil {
		return nil, errNilClient
	}
	burst := limit.Burst
	if burst <= 0 {
		burst = limit.Rate
	}
	if burst <= 0 {
		burst = 1
	}

	res, err := rueidislimiter.ExecGCRAAllowN(
		ctx,
		l.client,
		redisPrefix+key,
		int64(burst),
		int64(limit.Rate),
		limit.Period,
		int64(n),
	)
	if err != nil {
		return nil, err
	}

	return &Result{
		Limit:      limit,
		Allowed:    int(res.Allowed),
		Remaining:  int(res.Remaining),
		RetryAfter: res.RetryAfter,
		ResetAfter: res.ResetAfter,
	}, nil
}

func (l Limiter) AllowAtMost(ctx context.Context, key string, limit Limit, n int) (*Result, error) {
	if l.client == nil {
		return nil, errNilClient
	}
	burst := limit.Burst
	if burst <= 0 {
		burst = limit.Rate
	}
	if burst <= 0 {
		burst = 1
	}

	res, err := rueidislimiter.ExecGCRAAllowAtMost(
		ctx,
		l.client,
		redisPrefix+key,
		int64(burst),
		int64(limit.Rate),
		limit.Period,
		int64(n),
	)
	if err != nil {
		return nil, err
	}

	return &Result{
		Limit:      limit,
		Allowed:    int(res.Allowed),
		Remaining:  int(res.Remaining),
		RetryAfter: res.RetryAfter,
		ResetAfter: res.ResetAfter,
	}, nil
}

func (l *Limiter) Reset(ctx context.Context, key string) error {
	if l.client == nil {
		return errNilClient
	}
	return l.client.Do(ctx, l.client.B().Del().Key(redisPrefix+key).Build()).Error()
}

