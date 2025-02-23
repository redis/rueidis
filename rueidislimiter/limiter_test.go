package rueidislimiter_test

import (
	"context"
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidislimiter"
)

func failingClientBuilder(option rueidis.ClientOption) (rueidis.Client, error) {
	return nil, errors.New("failed to create Redis client")
}

func TestNewRateLimiterFailure(t *testing.T) {
	_, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
		ClientBuilder: failingClientBuilder,
	})
	if err == nil {
		t.Fatal("Expected error when failing to create Redis client, but got nil")
	}
}

func TestLimitMethod(t *testing.T) {
	client := setup(t)
	t.Cleanup(client.Close)

	limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			return client, nil
		},
		Limit:  5,
		Window: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	if limiter.Limit() != 5 {
		t.Fatalf("Expected Limit() to return 5, got %d", limiter.Limit())
	}
}

type mockRedisClient struct{
	responseErr error
	responseData []int64
}

func (m *mockRedisClient) Do(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	if m.responseErr != nil {
		return rueidis.RedisResult{Err: m.responseErr}
	}
	return rueidis.RedisResult{Data: m.responseData}
}

func TestRedisErrorHandling(t *testing.T) {
	mockClient := &mockRedisClient{responseErr: errors.New("redis error")}
	limiter := &rueidislimiter.rateLimiter{client: mockClient}

	_, err := limiter.AllowN(context.Background(), "test-key", 1)
	if err == nil {
		t.Fatal("Expected error when Redis fails, but got nil")
	}
}

func TestInvalidRedisResponse(t *testing.T) {
	mockClient := &mockRedisClient{responseData: []int64{1}}
	limiter := &rueidislimiter.rateLimiter{client: mockClient}

	_, err := limiter.AllowN(context.Background(), "test-key", 1)
	if err == nil || err != rueidislimiter.ErrInvalidResponse {
		t.Fatalf("Expected ErrInvalidResponse, got %v", err)
	}
}
