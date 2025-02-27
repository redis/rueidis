package rueidislimiter

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/dannotripp/rueidis"
)

var (
	ErrInvalidTokens   = errors.New("number of tokens must be non-negative")
	ErrInvalidResponse = errors.New("invalid response from Redis")
)

type Result struct {
	Allowed   bool
	Remaining int64
	ResetAtMs int64
}

type RateLimiterClient interface {
	Check(ctx context.Context, identifier string, options ...RateLimitOption) (Result, error)
	Allow(ctx context.Context, identifier string, options ...RateLimitOption) (Result, error)
	AllowN(ctx context.Context, identifier string, n int64, options ...RateLimitOption) (Result, error)
}

const PlaceholderPrefix = "rueidislimiter"

type rateLimiter struct {
	client           rueidis.Client
	keyPrefix        string
	defaultRateLimit RateLimitOption
}

type RateLimiterOption struct {
	ClientBuilder func(option rueidis.ClientOption) (rueidis.Client, error)
	KeyPrefix     string
	ClientOption  rueidis.ClientOption
	Limit         int
	Window        time.Duration
}

func NewRateLimiter(option RateLimiterOption) (RateLimiterClient, error) {
	if option.Window < time.Millisecond {
		option.Window = time.Millisecond
	}
	if option.Limit <= 0 {
		option.Limit = 1
	}
	if option.KeyPrefix == "" {
		option.KeyPrefix = PlaceholderPrefix
	}

	rl := &rateLimiter{
		defaultRateLimit: RateLimitOption{
			limit:  int64(option.Limit),
			window: option.Window,
		},
	}

	var err error
	if option.ClientBuilder != nil {
		rl.client, err = option.ClientBuilder(option.ClientOption)
	} else {
		rl.client, err = rueidis.NewClient(option.ClientOption)
	}
	if err != nil {
		return nil, err
	}
	rl.keyPrefix = option.KeyPrefix
	return rl, nil
}

func (l *rateLimiter) Limit() int {
	return int(l.defaultRateLimit.limit)
}

func (l *rateLimiter) Check(ctx context.Context, identifier string, options ...RateLimitOption) (Result, error) {
	return l.AllowN(ctx, identifier, 0, options...)
}

func (l *rateLimiter) Allow(ctx context.Context, identifier string, options ...RateLimitOption) (Result, error) {
	return l.AllowN(ctx, identifier, 1, options...)
}

func (l *rateLimiter) AllowN(ctx context.Context, identifier string, n int64, options ...RateLimitOption) (Result, error) {
	if n < 0 {
		return Result{}, ErrInvalidTokens
	}
	rl := l.defaultRateLimit
	if len(options) > 0 {
		rl = options[len(options)-1]
	}

	now := time.Now().UTC()
	keys := []string{l.getKey(identifier)}
	args := []string{
		strconv.FormatInt(n, 10),
		strconv.FormatInt(now.Add(rl.window).UnixMilli(), 10),
		strconv.FormatInt(now.UnixMilli(), 10),
	}

	resp := rateLimitScript.Exec(ctx, l.client, keys, args)
	if err := resp.Error(); err != nil {
		return Result{}, err
	}

	data, err := resp.AsIntSlice()
	if err != nil || len(data) != 2 {
		return Result{}, ErrInvalidResponse
	}

	current := data[0]
	remaining := rl.limit - current
	if remaining < 0 {
		remaining = 0
	}

	allowed := current <= rl.limit
	if n == 0 {
		allowed = current < rl.limit
	}

	return Result{
		Allowed:   allowed,
		Remaining: remaining,
		ResetAtMs: data[1],
	}, nil
}

func (l *rateLimiter) getKey(identifier string) string {
	sb := strings.Builder{}
	sb.Grow(len(l.keyPrefix) + len(identifier) + 3)
	sb.WriteString(l.keyPrefix)
	sb.WriteString(":{")
	sb.WriteString(identifier)
	sb.WriteString("}")
	return sb.String()
}

var rateLimitScript = rueidis.NewLuaScript(`
local rate_limit_key = KEYS[1]
local increment_amount = tonumber(ARGV[1])
local next_expires_at = tonumber(ARGV[2])
local current_time = tonumber(ARGV[3])
local expires_at_key = rate_limit_key .. ":ex"
local expires_at = tonumber(redis.call("get", expires_at_key))
if not expires_at or expires_at < current_time then
  redis.call("set", rate_limit_key, 0, "pxat", next_expires_at + 1000)
  redis.call("set", expires_at_key, next_expires_at, "pxat", next_expires_at + 1000)
  expires_at = next_expires_at
end
local current = redis.call("incrby", rate_limit_key, increment_amount)
return { current, expires_at }
`)
