package rueidislimiter_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"github.com/redis/rueidis/rueidislimiter"
	"go.uber.org/mock/gomock"
)

func TestNewRateLimiter(t *testing.T) {
	tests := []struct {
		name    string
		opt     rueidislimiter.RateLimiterOption
		wantErr error
	}{
		{
			name: "default values",
			opt: rueidislimiter.RateLimiterOption{
				ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
					return mock.NewClient(gomock.NewController(t)), nil
				},
				Limit:  1,
				Window: time.Second,
			},
		},
		{
			name: "custom values",
			opt: rueidislimiter.RateLimiterOption{
				ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
					return mock.NewClient(gomock.NewController(t)), nil
				},
				Limit:     100,
				Window:    time.Second,
				KeyPrefix: "test:",
			},
		},
		{
			name: "invalid window",
			opt: rueidislimiter.RateLimiterOption{
				ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
					return mock.NewClient(gomock.NewController(t)), nil
				},
				Limit:  1,
				Window: -time.Second,
			},
			wantErr: rueidislimiter.ErrInvalidWindow,
		},
		{
			name: "invalid limit",
			opt: rueidislimiter.RateLimiterOption{
				ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
					return mock.NewClient(gomock.NewController(t)), nil
				},
				Limit:  -1,
				Window: time.Second,
			},
			wantErr: rueidislimiter.ErrInvalidLimit,
		},
		{
			name: "empty key prefix",
			opt: rueidislimiter.RateLimiterOption{
				ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
					return mock.NewClient(gomock.NewController(t)), nil
				},
				Limit:  1,
				Window: time.Second,
			},
		},
		{
			name: "nil client builder",
			opt: rueidislimiter.RateLimiterOption{
				ClientOption: rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}},
				Limit:        1,
				Window:       time.Second,
			},
		},
		{
			name: "new client error",
			opt: rueidislimiter.RateLimiterOption{
				ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
					return nil, errors.New("client error")
				},
				Limit:  1,
				Window: time.Second,
			},
			wantErr: errors.New("client error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := rueidislimiter.NewRateLimiter(tt.opt)
			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("NewRateLimiter() error = nil, wantErr %v", tt.wantErr)
				}
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("NewRateLimiter() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("NewRateLimiter() error = %v, wantErr nil", err)
			}
		})
	}
}

func TestRateLimiter_AllowN(t *testing.T) {
	now := time.Now()
	resetTime := now.Add(time.Second).UnixMilli()

	tests := []struct {
		name       string
		mockResp   rueidis.RedisResult
		n          int64
		customOpt  *rueidislimiter.RateLimitOption
		wantResult rueidislimiter.Result
		wantErr    bool
		setupMock  bool
	}{
		{
			name:    "negative tokens",
			n:       -1,
			wantErr: true,
		},
		{
			name: "success with default limit",
			mockResp: mock.Result(mock.RedisArray(
				mock.RedisInt64(1),
				mock.RedisInt64(resetTime),
			)),
			n:         1,
			setupMock: true,
			wantResult: rueidislimiter.Result{
				Allowed:   true,
				Remaining: 9,
				ResetAtMs: resetTime,
			},
		},
		{
			name: "success with custom limit",
			mockResp: mock.Result(mock.RedisArray(
				mock.RedisInt64(5),
				mock.RedisInt64(resetTime),
			)),
			n:         1,
			setupMock: true,
			customOpt: func() *rueidislimiter.RateLimitOption {
				opt := rueidislimiter.WithCustomRateLimit(20, time.Second*2)
				return &opt
			}(),
			wantResult: rueidislimiter.Result{
				Allowed:   true,
				Remaining: 15,
				ResetAtMs: resetTime,
			},
		},
		{
			name: "limit exceeded",
			mockResp: mock.Result(mock.RedisArray(
				mock.RedisInt64(11),
				mock.RedisInt64(resetTime),
			)),
			n:         1,
			setupMock: true,
			wantResult: rueidislimiter.Result{
				Allowed:   false,
				Remaining: 0,
				ResetAtMs: resetTime,
			},
		},
		{
			name:      "redis error",
			mockResp:  mock.ErrorResult(errors.New("redis error")),
			n:         1,
			setupMock: true,
			wantErr:   true,
		},
		{
			name:      "invalid response type",
			mockResp:  mock.Result(mock.RedisString("invalid")),
			n:         1,
			setupMock: true,
			wantErr:   true,
		},
		{
			name:      "invalid array length",
			mockResp:  mock.Result(mock.RedisArray(mock.RedisInt64(1))),
			n:         1,
			setupMock: true,
			wantErr:   true,
		},
		{
			name: "invalid first element",
			mockResp: mock.Result(mock.RedisArray(
				mock.RedisString("invalid"),
				mock.RedisInt64(1),
			)),
			n:         1,
			setupMock: true,
			wantErr:   true,
		},
		{
			name: "invalid second element",
			mockResp: mock.Result(mock.RedisArray(
				mock.RedisInt64(1),
				mock.RedisString("invalid"),
			)),
			n:         1,
			setupMock: true,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			client := mock.NewClient(ctrl)
			if tt.setupMock {
				client.EXPECT().Do(gomock.Any(), gomock.Any()).Return(tt.mockResp).Times(1)
			}

			limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
				ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
					return client, nil
				},
				Limit:  10,
				Window: time.Second,
			})
			if err != nil {
				t.Fatal(err)
			}

			var got rueidislimiter.Result
			if tt.customOpt != nil {
				got, err = limiter.AllowN(context.Background(), "test", tt.n, *tt.customOpt)
			} else {
				got, err = limiter.AllowN(context.Background(), "test", tt.n)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("AllowN() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if got != tt.wantResult {
				t.Fatalf("AllowN() = %+v, want %+v", got, tt.wantResult)
			}
		})
	}
}

func TestRateLimiter_Check(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	resetTime := now.Add(time.Second).UnixMilli()

	client := mock.NewClient(ctrl)
	client.EXPECT().Do(gomock.Any(), gomock.Any()).Return(mock.Result(mock.RedisArray(
		mock.RedisInt64(5),
		mock.RedisInt64(resetTime),
	))).Times(1)

	limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			return client, nil
		},
		Limit:  10,
		Window: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := limiter.Check(context.Background(), "test")
	if err != nil {
		t.Fatalf("Check() error = %v", err)
	}

	want := rueidislimiter.Result{
		Allowed:   true,
		Remaining: 5,
		ResetAtMs: resetTime,
	}
	if got != want {
		t.Fatalf("Check() = %+v, want %+v", got, want)
	}
}

func TestRateLimiter_Allow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	resetTime := now.Add(time.Second).UnixMilli()

	client := mock.NewClient(ctrl)
	client.EXPECT().Do(gomock.Any(), gomock.Any()).Return(mock.Result(mock.RedisArray(
		mock.RedisInt64(1),
		mock.RedisInt64(resetTime),
	))).Times(1)

	limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			return client, nil
		},
		Limit:  10,
		Window: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := limiter.Allow(context.Background(), "test")
	if err != nil {
		t.Fatalf("Allow() error = %v", err)
	}

	want := rueidislimiter.Result{
		Allowed:   true,
		Remaining: 9,
		ResetAtMs: resetTime,
	}
	if got != want {
		t.Fatalf("Allow() = %+v, want %+v", got, want)
	}
}

func TestRateLimiter_Limit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewClient(ctrl)
	limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			return client, nil
		},
		Limit:  42,
		Window: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	if got := limiter.Limit(); got != 42 {
		t.Fatalf("Limit() = %v, want %v", got, 42)
	}
}

func BenchmarkAllowN(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	now := time.Now()
	resetTime := now.Add(time.Second).UnixMilli()

	client := mock.NewClient(ctrl)
	client.EXPECT().Do(gomock.Any(), gomock.Any()).Return(mock.Result(mock.RedisArray(
		mock.RedisInt64(1),
		mock.RedisInt64(resetTime),
	))).Times(b.N)

	limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			return client, nil
		},
		Limit:  1000,
		Window: time.Second,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := limiter.AllowN(context.Background(), "test", 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}
