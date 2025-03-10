package rueidislimiter

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/redis/rueidis"
)

var (
	ErrInvalidTokens   = errors.New("number of tokens must be non-negative")
	ErrInvalidResponse = errors.New("invalid response from Redis")
	ErrInvalidLimit    = errors.New("limit must be positive")
	ErrInvalidWindow   = errors.New("window must be positive")
	ErrNilBuilder      = errors.New("client builder is required")
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
	Limit() int
}

const (
	PlaceholderPrefix = "rueidislimiter"
	keyDelimOpen      = ":{"
	keyDelimClose     = "}"
)

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
	if option.Window <= 0 {
		return nil, ErrInvalidWindow
	}
	if option.Limit <= 0 {
		return nil, ErrInvalidLimit
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

	bufs := rateBuffersPool.Get(0, 128)
	defer rateBuffersPool.Put(bufs)

	now := time.Now().UTC()

	offset := len(bufs.keyBuf)
	bufs.keyBuf = append(bufs.keyBuf, l.keyPrefix...)
	bufs.keyBuf = append(bufs.keyBuf, keyDelimOpen...)
	bufs.keyBuf = append(bufs.keyBuf, identifier...)
	bufs.keyBuf = append(bufs.keyBuf, keyDelimClose...)
	key := rueidis.BinaryString(bufs.keyBuf[offset:])

	offset = len(bufs.keyBuf)
	bufs.keyBuf = strconv.AppendInt(bufs.keyBuf, n, 10)
	arg1 := rueidis.BinaryString(bufs.keyBuf[offset:])

	offset = len(bufs.keyBuf)
	bufs.keyBuf = strconv.AppendInt(bufs.keyBuf, now.Add(rl.window).UnixMilli(), 10)
	arg2 := rueidis.BinaryString(bufs.keyBuf[offset:])

	offset = len(bufs.keyBuf)
	bufs.keyBuf = strconv.AppendInt(bufs.keyBuf, now.UnixMilli(), 10)
	arg3 := rueidis.BinaryString(bufs.keyBuf[offset:])

	resp := rateLimitScript.Exec(ctx, l.client, []string{key}, []string{arg1, arg2, arg3})
	if err := resp.Error(); err != nil {
		return Result{}, err
	}

	arr, err := resp.ToArray()
	if err != nil || len(arr) != 2 {
		return Result{}, ErrInvalidResponse
	}

	current, err := arr[0].ToInt64()
	if err != nil {
		return Result{}, ErrInvalidResponse
	}

	resetAt, err := arr[1].ToInt64()
	if err != nil {
		return Result{}, ErrInvalidResponse
	}

	remaining := max(rl.limit-current, 0)
	allowed := current <= rl.limit && (n > 0 || current < rl.limit)

	return Result{
		Allowed:   allowed,
		Remaining: remaining,
		ResetAtMs: resetAt,
	}, nil
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
