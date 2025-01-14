package rueidisprob

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/redis/rueidis"
)

func TestNewCountingBloomFilter(t *testing.T) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
	if err != nil {
		t.Error(err)
	}

	if bf == nil {
		t.Error("Bloom filter is nil")
	}
	if bf.(*countingBloomFilter).client == nil {
		t.Error("Client is nil")
	}
	if bf.(*countingBloomFilter).name != "{test}:cbf" {
		t.Error("Name is not {test}:cbf")
	}
	if bf.(*countingBloomFilter).counter != "{test}:cbf:c" {
		t.Error("Counter is not test:cbf:c")
	}
	if bf.(*countingBloomFilter).hashIterations != 4 {
		t.Error("Hash iterations is not 4")
	}
	if bf.(*countingBloomFilter).client != client {
		t.Error("Client is not equal")
	}
}

func TestNewCountingBloomFilterError(t *testing.T) {
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

		_, err = NewCountingBloomFilter(client, "", 100, 0.05)
		if !errors.Is(err, ErrEmptyCountingBloomFilterName) {
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

		_, err = NewCountingBloomFilter(client, "test", 100, -0.01)
		if !errors.Is(err, ErrCountingBloomFilterFalsePositiveRateLessThanEqualZero) {
			t.Error("Error is not ErrFalsePositiveRateNegative")
		}
	})

	t.Run("GreaterThanOneFalsePositiveRate", func(t *testing.T) {
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

		_, err = NewCountingBloomFilter(client, "test", 100, 1.01)
		if !errors.Is(err, ErrCountingBloomFilterFalsePositiveRateGreaterThanOne) {
			t.Error("Error is not ErrFalsePositiveRateGreaterThanOne")
		}
	})

	t.Run("BitsSizeZero", func(t *testing.T) {
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

		_, err = NewCountingBloomFilter(client, "test", 0, 0.01)
		if !errors.Is(err, ErrCountingBloomFilterBitsSizeZero) {
			t.Error("Error is not ErrBitsSizeZero")
		}
	})
}

func TestCountingBloomFilterAdd(t *testing.T) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
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

	count, err := bf.Count(context.Background())
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("Count is not 1")
	}
}

func TestCountingBloomFilterAddError(t *testing.T) {
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

	bf, err := NewBloomFilter(client, "test", 100, 0.05)
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = bf.Add(ctx, "1")
	if !errors.Is(err, context.Canceled) {
		t.Error("Error is not context.Canceled")
	}
}

func TestCountingBloomFilterAddMulti(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
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

		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 3 {
			t.Error("Count is not 3")
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.AddMulti(context.Background(), []string{})
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

	t.Run("add already exists items", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.Add(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
		exist, err := bf.Exists(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
		if !exist {
			t.Error("Key 1 does not exist")
		}

		err = bf.AddMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}

		for _, key := range []string{"1", "2", "3"} {
			exists, err := bf.Exists(context.Background(), key)
			if err != nil {
				t.Error(err)
			}
			if !exists {
				t.Errorf("Key %s does not exist", key)
			}
		}

		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 4 {
			t.Error("Count is not 4")
		}

		buf := make([]byte, 0)
		existingIndexes := bf.(*countingBloomFilter).indexes([]string{"1"}, &buf)
		resp := client.Do(
			context.Background(),
			client.B().
				Hmget().
				Key("{test}:cbf").
				Field(existingIndexes...).
				Build(),
		)
		if resp.Error() != nil {
			t.Error(resp.Error())
		}

		arr, err := resp.AsIntSlice()
		if err != nil {
			t.Error(err)
		}
		for _, v := range arr {
			if v != 2 {
				t.Error("Value is not 2")
			}
		}
	})

	t.Run("add duplicate items", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		keys := []string{"1", "2", "3", "1", "2", "3"}
		err = bf.AddMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		exists, err := bf.ExistsMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if !e {
				t.Error("Key does not exist")
			}
		}

		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 6 {
			t.Error("Count is not 6")
		}
	})
}

func TestCountingBloomFilterAddMultiError(t *testing.T) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = bf.AddMulti(ctx, []string{"1", "2", "3"})
	if !errors.Is(err, context.Canceled) {
		t.Error("Error is not context.Canceled")
	}
}

func TestCountingBloomFilterExists(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
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
	})

	t.Run("does not exist", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		exists, err := bf.Exists(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
		if exists {
			t.Error("Key test exists")
		}
	})
}

func TestCountingBloomFilterExistsError(t *testing.T) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = bf.Exists(ctx, "1")
	if !errors.Is(err, context.Canceled) {
		t.Error("Error is not context.Canceled")
	}
}

func TestCountingBloomFilterExistsMulti(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.AddMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}

		exists, err := bf.ExistsMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if !e {
				t.Error("Key test does not exist")
			}
		}
	})

	t.Run("does not exist", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		exists, err := bf.ExistsMulti(context.Background(), []string{"4", "5", "6"})
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if e {
				t.Error("Key test exists")
			}
		}
	})

	t.Run("duplicated saved items exist", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.AddMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}
		err = bf.AddMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}

		err = bf.RemoveMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}

		exists, err := bf.ExistsMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if !e {
				t.Error("Key test does not exist")
			}
		}
		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 3 {
			t.Error("Count is not 3")
		}
	})

	t.Run("empty keys", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		exists, err := bf.ExistsMulti(context.Background(), []string{})
		if err != nil {
			t.Error(err)
		}
		if len(exists) != 0 {
			t.Error("Exists is not empty")
		}
	})
}

func TestCountingBloomFilterExistsMultiError(t *testing.T) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = bf.ExistsMulti(ctx, []string{"1", "2", "3"})
	if !errors.Is(err, context.Canceled) {
		t.Error("Error is not context.Canceled")
	}
}

func TestCountingBloomFilterRemove(t *testing.T) {
	t.Run("remove", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
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

		err = bf.Remove(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}

		exists, err = bf.Exists(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
		if exists {
			t.Error("Key test exists")
		}

		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 0 {
			t.Error("Count is not 0")
		}
	})

	t.Run("does not exist", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.Remove(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
	})
}

func TestCountingBloomFilterRemoveError(t *testing.T) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = bf.Remove(ctx, "1")
	if !errors.Is(err, context.Canceled) {
		t.Error("Error is not context.Canceled")
	}
}

func TestCountingBloomFilterRemoveMulti(t *testing.T) {
	t.Run("remove multiple items", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		keys := []string{"1", "2", "3"}
		err = bf.AddMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		exists, err := bf.ExistsMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if !e {
				t.Error("Key does not exist")
			}
		}

		err = bf.RemoveMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		exists, err = bf.ExistsMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if e {
				t.Error("Key exists")
			}
		}

		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 0 {
			t.Error("Count is not 0")
		}
	})

	t.Run("remove empty items", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.RemoveMulti(context.Background(), []string{})
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

	t.Run("remove not exist items", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.RemoveMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}

		exists, err := bf.ExistsMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if e {
				t.Error("Key exists")
			}
		}
	})

	t.Run("remove already deleted items", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.AddMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}
		exists, err := bf.ExistsMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if !e {
				t.Error("Key does not exist")
			}
		}

		err = bf.RemoveMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}
		exists, err = bf.ExistsMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if e {
				t.Error("Key exists")
			}
		}
		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 0 {
			t.Error("Count is not 0")
		}

		err = bf.RemoveMulti(context.Background(), []string{"1", "2", "3"})
		if err != nil {
			t.Error(err)
		}

		buf := make([]byte, 0)
		removedIndexes := bf.(*countingBloomFilter).indexes([]string{"1", "2", "3"}, &buf)
		resp := client.Do(
			context.Background(),
			client.B().
				Hmget().
				Key("{test}:cbf").
				Field(removedIndexes...).
				Build(),
		)
		if resp.Error() != nil {
			t.Error(resp.Error())
		}

		arr, err := resp.AsIntSlice()
		if err != nil {
			t.Error(err)
		}
		for _, v := range arr {
			if v != 0 {
				t.Error("Value is not 0")
			}
		}
	})

	t.Run("remove duplicate items", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		keys := []string{"1", "1", "1", "2", "2", "2"}
		err = bf.AddMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		exists, err := bf.ExistsMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if !e {
				t.Error("Key does not exist")
			}
		}

		err = bf.RemoveMulti(context.Background(), []string{"1", "1", "2", "2"})
		if err != nil {
			t.Error(err)
		}
		exists, err = bf.ExistsMulti(context.Background(), []string{"1", "2"})
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if !e {
				t.Error("Key does not exist")
			}
		}
		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 2 {
			t.Error("Count is not 2")
		}

		buf := make([]byte, 0)
		removedIndexes := bf.(*countingBloomFilter).indexes([]string{"1", "2"}, &buf)
		resp := client.Do(
			context.Background(),
			client.B().
				Hmget().
				Key("{test}:cbf").
				Field(removedIndexes...).
				Build(),
		)
		if resp.Error() != nil {
			t.Error(resp.Error())
		}

		arr, err := resp.AsIntSlice()
		if err != nil {
			t.Error(err)
		}
		for _, v := range arr {
			if v != 1 {
				t.Error("Value is not 1")
			}
		}
	})

	t.Run("remove more than count", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		keys := []string{"1", "1"}
		err = bf.AddMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		exists, err := bf.ExistsMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if !e {
				t.Error("Key does not exist")
			}
		}

		err = bf.RemoveMulti(context.Background(), []string{"1", "1", "1"})
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

		buf := make([]byte, 0)
		removedIndexes := bf.(*countingBloomFilter).indexes([]string{"1"}, &buf)
		resp := client.Do(
			context.Background(),
			client.B().
				Hmget().
				Key("{test}:cbf").
				Field(removedIndexes...).
				Build(),
		)
		if resp.Error() != nil {
			t.Error(resp.Error())
		}

		arr, err := resp.AsIntSlice()
		if err != nil {
			t.Error(err)
		}
		for _, v := range arr {
			if v != 0 {
				t.Error("Value is not 0")
			}
		}
	})

	t.Run("remove very large items", func(t *testing.T) {
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

		// Above `LUAI_MAXCSTACK`(8000) limit
		keys := make([]string, 8001)
		for i := 0; i < 8001; i++ {
			keys[i] = fmt.Sprintf("%d", i)
		}

		bf, err := NewCountingBloomFilter(client, "test", 10000, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.AddMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}

		err = bf.RemoveMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}

		exists, err := bf.ExistsMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		for _, e := range exists {
			if e {
				t.Error("Key exists")
			}
		}
	})
}

func TestCountingBloomFilterDelete(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
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

		err = bf.Delete(context.Background())
		if err != nil {
			t.Error(err)
		}

		resp := client.Do(
			context.Background(),
			client.B().
				Get().
				Key("{test}:cbf").
				Build(),
		)
		if !rueidis.IsRedisNil(resp.Error()) {
			t.Error("Error is not rueidis.ErrNil")
		}

		resp = client.Do(
			context.Background(),
			client.B().
				Get().
				Key("{test}:cbf:c").
				Build(),
		)
		if !rueidis.IsRedisNil(resp.Error()) {
			t.Error("Error is not rueidis.ErrNil")
		}
	})

	t.Run("delete does not exist", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.Delete(context.Background())
		if err != nil {
			t.Error(err)
		}
	})
}

func TestCountingBloomFilterDeleteError(t *testing.T) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = bf.Delete(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Error("Error is not context.Canceled")
	}
}

func TestCountingBloomFilterItemMintCount(t *testing.T) {
	t.Run("item exists", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		err = bf.AddMulti(context.Background(), []string{"1", "1"})
		if err != nil {
			t.Error(err)
		}

		count, err := bf.ItemMinCount(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
		if count != 2 {
			t.Error("Count is not 2")
		}
	})

	t.Run("item does not exist", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		count, err := bf.ItemMinCount(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}
		if count != 0 {
			t.Error("Count is not 0")
		}
	})
}

func TestCountingBloomFilterItemMinCountError(t *testing.T) {
	t.Run("min item count error", func(t *testing.T) {
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

		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		_, err = bf.ItemMinCount(ctx, "1")
		if !errors.Is(err, context.Canceled) {
			t.Error("Error is not context.Canceled")
		}
	})
}

func TestCountingBloomFilterItemMinCountMulti(t *testing.T) {
	t.Run("item exists", func(t *testing.T) {
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

		keys := []string{"1", "2", "3"}
		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}
		err = bf.AddMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}

		counts, err := bf.ItemMinCountMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		for _, c := range counts {
			if c != 1 {
				t.Error("Count is not 1")
			}
		}
	})

	t.Run("item does not exist", func(t *testing.T) {
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

		keys := []string{"1", "2", "3"}
		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		counts, err := bf.ItemMinCountMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		for _, c := range counts {
			if c != 0 {
				t.Error("Count is not 0")
			}
		}
	})

	t.Run("empty keys", func(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		counts, err := bf.ItemMinCountMulti(context.Background(), []string{})
		if err != nil {
			t.Error(err)
		}
		if len(counts) != 0 {
			t.Error("Counts is not empty")
		}
	})

	t.Run("zero count items", func(t *testing.T) {
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

		keys := []string{"1", "2", "3"}
		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}
		err = bf.AddMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		err = bf.RemoveMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}

		counts, err := bf.ItemMinCountMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}
		for _, c := range counts {
			if c != 0 {
				t.Error("Count is not 0")
			}
		}
	})
}

func TestCountingBloomFilterCount(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
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

func TestCountingBloomFilterCountError(t *testing.T) {
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
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

		bf, err := NewCountingBloomFilter(client, "test", 100, 0.05)
		if err != nil {
			t.Error(err)
		}

		resp := client.Do(
			context.Background(),
			client.B().
				Set().
				Key("{test}:cbf:c").
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

func BenchmarkCountingBloomFilterAddMultiBigSize(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100000000, 0.01)
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

func BenchmarkCountingBloomFilterAddMultiLowRate(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 1000000, 0.0000000001)
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

func BenchmarkCountingBloomFilterAddMultiManyKeys(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 1000000, 0.01)
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

func BenchmarkCountingBloomFilterExistsMultiBigSize(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100000000, 0.01)
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

func BenchmarkCountingBloomFilterExistsMultiLowRate(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 1000000, 0.0000000001)
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

func BenchmarkCountingBloomFilterExistsMultiManyKeys(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 1000000, 0.01)
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

func BenchmarkCountingBloomFilterRemoveMultiBigSize(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100000000, 0.01)
	if err != nil {
		b.Error(err)
	}

	keys := make([]string, 10)
	for i := 0; i < 10; i++ {
		keys[i] = strconv.Itoa(i)
	}
	for i := 0; i < b.N; i++ {
		err := bf.AddMulti(context.Background(), keys)
		if err != nil {
			b.Error(err)
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := bf.RemoveMulti(context.Background(), keys)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCountingBloomFilterRemoveMultiLowRate(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 1000000, 0.0000000001)
	if err != nil {
		b.Error(err)
	}

	keys := make([]string, 10)
	for i := 0; i < 10; i++ {
		keys[i] = strconv.Itoa(i)
	}
	for i := 0; i < b.N; i++ {
		err := bf.AddMulti(context.Background(), keys)
		if err != nil {
			b.Error(err)
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := bf.RemoveMulti(context.Background(), keys)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCountingBloomFilterRemoveMultiManyKeys(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 1000000, 0.01)
	if err != nil {
		b.Error(err)
	}

	keys := make([]string, 200)
	for i := 0; i < 200; i++ {
		keys[i] = strconv.Itoa(i)
	}
	for i := 0; i < b.N; i++ {
		err := bf.AddMulti(context.Background(), keys)
		if err != nil {
			b.Error(err)
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := bf.RemoveMulti(context.Background(), keys)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCountingBloomFilterItemMinCountMultiBigSize(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 100000000, 0.01)
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
		_, err := bf.ItemMinCountMulti(context.Background(), benchKeys)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCountingBloomFilterItemMinCountMultiLowRate(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 1000000, 0.0000000001)
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
		_, err := bf.ItemMinCountMulti(context.Background(), benchKeys)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCountingBloomFilterItemMinCountMultiManyKeys(b *testing.B) {
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

	bf, err := NewCountingBloomFilter(client, "test", 1000000, 0.01)
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
		_, err := bf.ItemMinCountMulti(context.Background(), benchKeys)
		if err != nil {
			b.Error(err)
		}
	}
}
