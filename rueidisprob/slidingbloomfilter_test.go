package rueidisprob

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/redis/rueidis"
)

func TestNewSlidingBloomFilter(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		if bf == nil {
			t.Error("Bloom filter is nil")
		}
		sbf := bf.(*slidingBloomFilter)
		if sbf.client == nil {
			t.Error("Client is nil")
		}
		if sbf.name != "{test}" {
			t.Error("Name is not {test}")
		}
		if sbf.hashIterations != 4 {
			t.Error("Hash iterations is not 4")
		}
		if sbf.window != time.Minute {
			t.Error("Window size is not 1 minute")
		}
	})

	t.Run("with read operation enabled", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute, WithReadOnlyExists(true))
		if err != nil {
			t.Error(err)
		}

		if bf == nil {
			t.Error("Bloom filter is nil")
		}
	})
}

func TestNewSlidingBloomFilterError(t *testing.T) {
	t.Run("EmptyName", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		_, err = NewSlidingBloomFilter(client, "", 100, 0.05, time.Minute)
		if !errors.Is(err, ErrEmptyName) {
			t.Error("Error is not ErrEmptyName")
		}
	})

	t.Run("NegativeFalsePositiveRate", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		_, err = NewSlidingBloomFilter(client, "test", 100, -0.01, time.Minute)
		if !errors.Is(err, ErrFalsePositiveRateLessThanEqualZero) {
			t.Error("Error is not ErrFalsePositiveRateLessThanEqualZero")
		}
	})

	t.Run("WindowSizeLessThanOneSecond", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		_, err = NewSlidingBloomFilter(client, "test", 100, 0.05, time.Millisecond)
		if !errors.Is(err, ErrWindowSizeLessThanOneSecond) {
			t.Error("Error is not ErrWindowSizeLessThanOneSecond")
		}
	})

	t.Run("ZeroSize", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		_, err = NewSlidingBloomFilter(client, "test", 0, 0.05, time.Minute)
		if !errors.Is(err, ErrBitsSizeZero) {
			t.Error("Error is not ErrBitsSizeZero")
		}
	})

	t.Run("TooLargeSize", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		_, err = NewSlidingBloomFilter(client, "test", 1<<33, 0.05, time.Minute)
		if !errors.Is(err, ErrBitsSizeTooLarge) {
			t.Error("Error is not ErrBitsSizeTooLarge")
		}
	})
}

func TestSlidingBloomFilterAdd(t *testing.T) {
	client, flushAllAndClose, err := setupRedis7Cluster()
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err := flushAllAndClose()
		if err != nil {
			t.Error(err)
		}
	}()

	bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
	if err != nil {
		t.Error(err)
	}

	err = bf.Add(context.Background(), "1")
	if err != nil {
		t.Error(err)
	}

	exists, err := bf.Exists(context.Background(), "1")
	if err != nil {
		t.Error(err)
	}
	if !exists {
		t.Error("Key test does not exist")
	}
}

func TestSlidingBloomFilterAddMulti(t *testing.T) {
	t.Run("add multiple items", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		keys := []string{"1", "2", "3"}
		err = bf.AddMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}

		for _, key := range keys {
			exists, err := bf.Exists(context.Background(), key)
			if err != nil {
				t.Error(err)
			}
			if !exists {
				t.Errorf("Key %s does not exist", key)
			}
		}
	})

	t.Run("add empty items", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		err = bf.AddMulti(context.Background(), []string{})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestSlidingBloomFilterRotation(t *testing.T) {
	t.Run("rotation after window", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()
		windowDuration := time.Second
		rotationInterval := time.Millisecond * 600

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, windowDuration)
		if err != nil {
			t.Error(err)
		}

		err = bf.Add(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}

		exists, err := bf.Exists(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
		if !exists {
			t.Error("Key should exist before rotation")
		}

		// Wait for rotation
		time.Sleep(rotationInterval)

		// Check that item is still in the filter
		exists, err = bf.Exists(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
		if !exists {
			t.Error("Key should exist in the filter")
		}

		// Wait for another rotation
		time.Sleep(rotationInterval)

		// Item should be gone after second rotation
		exists, err = bf.Exists(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
		if exists {
			t.Error("Key should not exist after second rotation")
		}
	})
}

func TestSlidingBloomFilterDelete(t *testing.T) {
	t.Run("delete exists", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		err = bf.Add(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}

		err = bf.Delete(context.Background())
		if err != nil {
			t.Error(err)
		}

		// Verify all keys are deleted
		for _, key := range []string{"{test}", "{test}:n", "{test}:c", "{test}:nc", "{test}:lr"} {
			resp := client.Do(context.Background(), client.B().Get().Key(key).Build())
			if !rueidis.IsRedisNil(resp.Error()) {
				t.Errorf("Key %s still exists", key)
			}
		}
	})
}

func BenchmarkSlidingBloomFilterAddMultiBigSize(b *testing.B) {
	client, flushAllAndClose, err := setupRedis7Cluster()
	if err != nil {
		b.Error(err)
	}
	defer func() {
		err := flushAllAndClose()
		if err != nil {
			b.Error(err)
		}
	}()

	bf, err := NewSlidingBloomFilter(client, "test", 100000000, 0.01, time.Minute)
	if err != nil {
		b.Error(err)
	}

	keys := make([]string, 10)
	for i := 0; i < 10; i++ {
		keys[i] = strconv.Itoa(i)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := bf.AddMulti(context.Background(), keys)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSlidingBloomFilterAddMultiLowRate(b *testing.B) {
	client, flushAllAndClose, err := setupRedis7Cluster()
	if err != nil {
		b.Error(err)
	}
	defer func() {
		err := flushAllAndClose()
		if err != nil {
			b.Error(err)
		}
	}()

	bf, err := NewSlidingBloomFilter(client, "test", 1000000, 0.0000000001, time.Minute)
	if err != nil {
		b.Error(err)
	}

	keys := make([]string, 10)
	for i := 0; i < 10; i++ {
		keys[i] = strconv.Itoa(i)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := bf.AddMulti(context.Background(), keys)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSlidingBloomFilterAddMultiManyKeys(b *testing.B) {
	client, flushAllAndClose, err := setupRedis7Cluster()
	if err != nil {
		b.Error(err)
	}
	defer func() {
		err := flushAllAndClose()
		if err != nil {
			b.Error(err)
		}
	}()

	bf, err := NewSlidingBloomFilter(client, "test", 1000000, 0.01, time.Minute)
	if err != nil {
		b.Error(err)
	}

	keys := make([]string, 200)
	for i := 0; i < 200; i++ {
		keys[i] = strconv.Itoa(i)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := bf.AddMulti(context.Background(), keys)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSlidingBloomFilterExistsMultiBigSize(b *testing.B) {
	client, flushAllAndClose, err := setupRedis7Cluster()
	if err != nil {
		b.Error(err)
	}
	defer func() {
		err := flushAllAndClose()
		if err != nil {
			b.Error(err)
		}
	}()

	bf, err := NewSlidingBloomFilter(client, "test", 100000000, 0.01, time.Minute)
	if err != nil {
		b.Error(err)
	}

	keys := make([]string, 10)
	for i := 0; i < 10; i++ {
		keys[i] = strconv.Itoa(i)
	}
	err = bf.AddMulti(context.Background(), keys)
	if err != nil {
		b.Error(err)
	}

	var benchKeys []string
	for i := 0; i < 10; i++ {
		key := strconv.Itoa(rand.Intn(b.N))
		benchKeys = append(benchKeys, key)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := bf.ExistsMulti(context.Background(), benchKeys)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSlidingBloomFilterExistsMultiLowRate(b *testing.B) {
	client, flushAllAndClose, err := setupRedis7Cluster()
	if err != nil {
		b.Error(err)
	}
	defer func() {
		err := flushAllAndClose()
		if err != nil {
			b.Error(err)
		}
	}()

	bf, err := NewSlidingBloomFilter(client, "test", 1000000, 0.0000000001, time.Minute)
	if err != nil {
		b.Error(err)
	}

	keys := make([]string, 10)
	for i := 0; i < 10; i++ {
		keys[i] = strconv.Itoa(i)
	}
	err = bf.AddMulti(context.Background(), keys)
	if err != nil {
		b.Error(err)
	}

	var benchKeys []string
	for i := 0; i < 10; i++ {
		key := strconv.Itoa(rand.Intn(b.N))
		benchKeys = append(benchKeys, key)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := bf.ExistsMulti(context.Background(), benchKeys)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSlidingBloomFilterExistsMultiManyKeys(b *testing.B) {
	client, flushAllAndClose, err := setupRedis7Cluster()
	if err != nil {
		b.Error(err)
	}
	defer func() {
		err := flushAllAndClose()
		if err != nil {
			b.Error(err)
		}
	}()

	bf, err := NewSlidingBloomFilter(client, "test", 1000000, 0.01, time.Minute)
	if err != nil {
		b.Error(err)
	}

	keys := make([]string, 200)
	for i := 0; i < 200; i++ {
		keys[i] = strconv.Itoa(i)
	}
	err = bf.AddMulti(context.Background(), keys)
	if err != nil {
		b.Error(err)
	}

	var benchKeys []string
	for i := 0; i < 200; i++ {
		key := strconv.Itoa(rand.Intn(b.N))
		benchKeys = append(benchKeys, key)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := bf.ExistsMulti(context.Background(), benchKeys)
		if err != nil {
			b.Error(err)
		}
	}
}

func TestSlidingBloomFilterConcurrentRotation(t *testing.T) {
	client, flushAllAndClose, err := setupRedis7Cluster()
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err := flushAllAndClose()
		if err != nil {
			t.Error(err)
		}
	}()

	windowDuration := time.Second
	numClients := 10

	// Create multiple bloom filter instances
	filters := make([]BloomFilter, numClients)
	for i := 0; i < numClients; i++ {
		bf, err := NewSlidingBloomFilter(client, "test", 1000, 0.01, windowDuration)
		if err != nil {
			t.Fatal(err)
		}
		filters[i] = bf
	}
	startTime := time.Now()

	// Add some initial items that should stay during first half window
	initialKeys := []string{"initial1", "initial2", "initial3"}
	err = filters[0].AddMulti(context.Background(), initialKeys)
	if err != nil {
		t.Fatal(err)
	}

	// Create context with timeout for the entire test
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Run concurrent operations
	var wg sync.WaitGroup
	errChan := make(chan error, numClients)

	// Run client operations
	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			bf := filters[clientID]

			for {
				select {
				case <-ctx.Done():
					return
				default:
					key := fmt.Sprintf("client%d-key%d", clientID, time.Now().UnixNano())

					// Randomly choose between Add and Exists operations
					if rand.Float32() < 0.5 {
						err := bf.Add(ctx, key)
						if err != nil && !errors.Is(err, context.Canceled) {
							errChan <- fmt.Errorf("client %d add error: %w", clientID, err)
							return
						}
					} else {
						_, err := bf.Exists(ctx, key)
						if err != nil && !errors.Is(err, context.Canceled) {
							errChan <- fmt.Errorf("client %d exists error: %w", clientID, err)
							return
						}
					}

					time.Sleep(time.Millisecond)
				}
			}
		}(i)
	}

	newKeys := []string{"new1", "new2", "new3"}
	go func() {
		time.Sleep(windowDuration / 2)
		err = filters[0].AddMulti(ctx, newKeys)
		if err != nil {
			panic(err)
		}
	}()

	// Wait for initial keys to disappear and verify it took at least windowDuration
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	allGone := false
	for !allGone {
		select {
		case <-ctx.Done():
			t.Fatal("context deadline exceeded before keys disappeared")
		case <-ticker.C:
			exists, err := filters[0].ExistsMulti(ctx, initialKeys)
			if errors.Is(err, context.DeadlineExceeded) {
				break
			}
			if err != nil {
				t.Fatal(err)
			}

			allGone = true
			for _, exists := range exists {
				if exists {
					allGone = false
					break
				}
			}
		}
	}

	// Verify new keys are still present
	exists, err := filters[0].ExistsMulti(ctx, newKeys)
	if err != nil {
		t.Fatal(err)
	}
	for i, exists := range exists {
		if !exists {
			t.Errorf("new key %s not present", newKeys[i])
		}
	}

	// Cancel context and wait for all operations to complete
	cancel()
	wg.Wait()

	// Check for any errors from goroutines
	close(errChan)
	for err := range errChan {
		t.Error(err)
	}

	t.Logf("Test ran for %v with %d clients", time.Since(startTime), numClients)
}

func TestSlidingBloomFilterCount(t *testing.T) {
	t.Run("count exists", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		err = bf.Add(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}

		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 1 {
			t.Error("Count is not 1")
		}
	})

	t.Run("count does not exist", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 0 {
			t.Error("Count is not 0")
		}
	})

	t.Run("add multiple items", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		keys := []string{"1", "2", "3"}
		err = bf.AddMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}

		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 3 {
			t.Error("Count is not 3")
		}
	})
}

func TestSlidingBloomFilterCountError(t *testing.T) {
	t.Run("count error", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err = bf.Count(ctx)
		if !errors.Is(err, context.Canceled) {
			t.Error("Error is not context.Canceled")
		}
	})

	t.Run("counter key is corrupted", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		resp := client.Do(
			context.Background(),
			client.B().
				Set().
				Key("{test}:c").
				Value("not a number").
				Build(),
		)
		if resp.Error() != nil {
			t.Error(resp.Error())
		}

		_, err = bf.Count(context.Background())
		if !errors.Is(err, strconv.ErrSyntax) {
			t.Error("Error is not strconv.ErrSyntax")
		}
	})
}

func TestSlidingBloomFilterReset(t *testing.T) {
	t.Run("reset exists", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		// Reload scripts
		sbf := bf.(*slidingBloomFilter)
		if err := sbf.initialize(); err != nil {
			t.Fatalf("failed to initialize scripts: %v", err)
		}

		// Add an element first
		err = bf.Add(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}

		// Verify the element was added
		exists, err := bf.Exists(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
		if !exists {
			t.Error("Key should exist before reset")
		}

		// Wait a bit before resetting
		time.Sleep(100 * time.Millisecond)

		// Clean up first
		err = bf.Delete(context.Background())
		if err != nil {
			t.Error(err)
		}

		// Wait a bit after resetting
		time.Sleep(100 * time.Millisecond)

		// Create a new filter (same name)
		bf, err = NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		// Verify the element does not exist after reset
		exists, err = bf.Exists(context.Background(), "1")
		if err != nil && !rueidis.IsRedisNil(err) {
			t.Error(err)
		}
		if exists {
			t.Error("Key should not exist after reset")
		}

		// Verify the count is 0 after reset
		count, err := bf.Count(context.Background())
		if err != nil && !rueidis.IsRedisNil(err) {
			t.Error(err)
		}
		if count != 0 {
			t.Errorf("Count should be 0 after reset, got %d", count)
		}
	})

	t.Run("reset does not exist", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		// Reload scripts
		err = bf.(*slidingBloomFilter).initialize()
		if err != nil {
			t.Fatalf("failed to initialize: %v", err)
		}

		// Wait a bit before resetting
		time.Sleep(100 * time.Millisecond)

		// Reset on empty filter
		err = bf.Reset(context.Background())
		if err != nil && !rueidis.IsRedisNil(err) {
			t.Error(err)
		}

		// Wait a bit after resetting
		time.Sleep(100 * time.Millisecond)

		// Verify the count is 0 after reset
		count, err := bf.Count(context.Background())
		if err != nil && !rueidis.IsRedisNil(err) {
			t.Error(err)
		}
		if count != 0 {
			t.Errorf("Count should be 0 after reset, got %d", count)
		}
	})
}

func TestSlidingBloomFilterResetError(t *testing.T) {
	client, flushAllAndClose, err := setupRedis7Cluster()
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err := flushAllAndClose()
		if err != nil {
			t.Error(err)
		}
	}()

	bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = bf.Reset(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Error("Error is not context.Canceled")
	}
}

func TestSlidingBloomFilterExistsMulti(t *testing.T) {
	t.Run("error case", func(t *testing.T) {
		client, flushAllAndClose, err := setupRedis7Cluster()
		if err != nil {
			t.Error(err)
		}
		defer func() {
			err := flushAllAndClose()
			if err != nil {
				t.Error(err)
			}
		}()

		bf, err := NewSlidingBloomFilter(client, "test", 100, 0.05, time.Minute)
		if err != nil {
			t.Error(err)
		}

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, err = bf.ExistsMulti(ctx, []string{"1", "2", "3"})
		if !errors.Is(err, context.Canceled) {
			t.Error("Expected context.Canceled error")
		}
	})
}
