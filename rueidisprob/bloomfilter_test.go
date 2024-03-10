package rueidisprob

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"testing"

	"github.com/redis/rueidis"
)

var address = []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"}

func cleanup(client rueidis.Client, keys ...string) error {
	cmds := make([]rueidis.Completed, 0, len(keys))
	for _, key := range keys {
		cmds = append(cmds, client.B().Del().Key(key).Build())
	}

	resps := client.DoMulti(context.Background(), cmds...)
	for _, resp := range resps {
		if resp.Error() != nil {
			if !rueidis.IsRedisNil(resp.Error()) {
				return resp.Error()
			}
		}
	}
	return nil
}

func TestNewBloomFilter(t *testing.T) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		t.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 100, 0.01)
	if err != nil {
		t.Error(err)
	}

	if bf == nil {
		t.Error("Bloom filter is nil")
	}
	if bf.(*bloomFilter).client == nil {
		t.Error("Client is nil")
	}
	if bf.(*bloomFilter).name != "{test}" {
		t.Error("Name is not {test}")
	}
	if bf.(*bloomFilter).counter != "{test}:c" {
		t.Error("Counter is not test:c")
	}
	if bf.(*bloomFilter).hashIterations != 7 {
		t.Error("Hash iterations is not 7")
	}
	if bf.(*bloomFilter).client != client {
		t.Error("Client is not equal")
	}
}

func TestNewBloomFilterError(t *testing.T) {
	t.Run("EmptyName", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		_, err = NewBloomFilter(client, "", 100, 0.01)
		if !errors.Is(err, ErrEmptyName) {
			t.Error("Error is not ErrEmptyName")
		}
	})

	t.Run("NegativeFalsePositiveRate", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		_, err = NewBloomFilter(client, "test", 100, -0.01)
		if !errors.Is(err, ErrFalsePositiveRateLessThanEqualZero) {
			t.Error("Error is not ErrFalsePositiveRateNegative")
		}
	})

	t.Run("GreaterThanOneFalsePositiveRate", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		_, err = NewBloomFilter(client, "test", 100, 1.01)
		if !errors.Is(err, ErrFalsePositiveRateGreaterThanOne) {
			t.Error("Error is not ErrFalsePositiveRateGreaterThanOne")
		}
	})

	t.Run("BitsSizeZero", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		_, err = NewBloomFilter(client, "test", 0, 0.01)
		if !errors.Is(err, ErrBitsSizeZero) {
			t.Error("Error is not ErrBitsSizeZero")
		}
	})

	t.Run("BitsSizeTooLarge", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		_, err = NewBloomFilter(client, "test", 1<<32, 0.01)
		if !errors.Is(err, ErrBitsSizeTooLarge) {
			t.Error("Error is not ErrBitsSizeTooLarge")
		}
	})
}

func TestBloomFilterAdd(t *testing.T) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		t.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

	// cleanup
	err = cleanup(client, "{test}", "{test}:c")
	if err != nil {
		t.Error(err)
	}
}

func TestBloomFilterAddError(t *testing.T) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		t.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

func TestBloomFilterAddMulti(t *testing.T) {
	t.Run("add multiple items", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("add empty items", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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
		if count != 3 {
			t.Error("Count is not 3")
		}

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("add duplicate items", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
		if err != nil {
			t.Error(err)
		}

		keys := []string{"1", "2", "3", "1", "2", "3"}
		err = bf.AddMulti(context.Background(), keys)
		if err != nil {
			t.Error(err)
		}

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

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})
}

func TestBloomFilterAddMultiError(t *testing.T) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		t.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

func TestBloomFilterExists(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})
}

func TestBloomFilterExistsError(t *testing.T) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		t.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

func TestBloomFilterExistsMulti(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("empty keys", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

func TestBloomFilterExistsMultiError(t *testing.T) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		t.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

func TestBloomFilterReset(t *testing.T) {
	t.Run("reset exists", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
		if err != nil {
			t.Error(err)
		}

		err = bf.Add(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}

		err = bf.Reset(context.Background())
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

		count, err := bf.Count(context.Background())
		if err != nil {
			t.Error(err)
		}
		if count != 0 {
			t.Error("Count is not 0")
		}

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("reset does not exist", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
		if err != nil {
			t.Error(err)
		}

		err = bf.Reset(context.Background())
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
}

func TestBloomFilterResetError(t *testing.T) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		t.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

func TestBloomFilterDelete(t *testing.T) {
	t.Run("delete exists", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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
				Key("{test}").
				Build(),
		)
		if !rueidis.IsRedisNil(resp.Error()) {
			t.Error("Error is not rueidis.ErrNil")
		}

		resp = client.Do(
			context.Background(),
			client.B().
				Get().
				Key("{test}:c").
				Build(),
		)
		if !rueidis.IsRedisNil(resp.Error()) {
			t.Error("Error is not rueidis.ErrNil")
		}
	})

	t.Run("delete does not exist", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
		if err != nil {
			t.Error(err)
		}

		err = bf.Delete(context.Background())
		if err != nil {
			t.Error(err)
		}
	})
}

func TestBloomFilterDeleteError(t *testing.T) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		t.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

func TestBloomFilterCount(t *testing.T) {
	t.Run("count exists", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("count does not exist", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("add multiple items", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

		err = cleanup(client, "{test}", "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})
}

func TestBloomFilterCountError(t *testing.T) {
	t.Run("count error", func(t *testing.T) {
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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
		client, err := rueidis.NewClient(
			rueidis.ClientOption{InitAddress: address},
		)
		if err != nil {
			t.Error(err)
		}
		bf, err := NewBloomFilter(client, "test", 100, 0.01)
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

		err = cleanup(client, "{test}:c")
		if err != nil {
			t.Error(err)
		}
	})
}

func BenchmarkBloomFilterAddMultiBigSize(b *testing.B) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		b.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 100000000, 0.01)
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

	err = cleanup(client, "{test}", "{test}:c")
	if err != nil {
		b.Error(err)
	}
}

func BenchmarkBloomFilterAddMultiLowRate(b *testing.B) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		b.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 1000000, 0.0000000001)
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

	err = cleanup(client, "{test}", "{test}:c")
	if err != nil {
		b.Error(err)
	}
}

func BenchmarkBloomFilterAddMultiManyKeys(b *testing.B) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		b.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 1000000, 0.01)
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

	err = cleanup(client, "{test}", "{test}:c")
	if err != nil {
		b.Error(err)
	}
}

func BenchmarkBloomFilterExistsMultiBigSize(b *testing.B) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		b.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 100000000, 0.01)
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

	err = cleanup(client, "{test}", "{test}:c")
	if err != nil {
		b.Error(err)
	}
}

func BenchmarkBloomFilterExistsMultiLowRate(b *testing.B) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		b.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 1000000, 0.0000000001)
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

	err = cleanup(client, "{test}", "{test}:c")
	if err != nil {
		b.Error(err)
	}
}

func BenchmarkBloomFilterExistsMultiManyKeys(b *testing.B) {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{InitAddress: address},
	)
	if err != nil {
		b.Error(err)
	}
	bf, err := NewBloomFilter(client, "test", 1000000, 0.01)
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

	err = cleanup(client, "{test}", "{test}:c")
	if err != nil {
		b.Error(err)
	}
}
