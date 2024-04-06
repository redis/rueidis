package rueidisprob

import (
	"context"
	"errors"
	"github.com/redis/rueidis"
	"math"
	"strconv"
)

var (
	ErrEmptyCountingBloomFilterName                          = errors.New("name cannot be empty")
	ErrCountingBloomFilterFalsePositiveRateLessThanEqualZero = errors.New("false positive rate cannot be less than or equal to zero")
	ErrCountingBloomFilterFalsePositiveRateGreaterThanOne    = errors.New("false positive rate cannot be greater than 1")
	ErrCountingBloomFilterBitsSizeZero                       = errors.New("bits size cannot be zero")
)

const (
	countingBloomFilterAddMultiScript = `
local itemCount = tonumber(ARGV[1])
local numElements = tonumber(#ARGV) - 1
local filterKey = KEYS[1]
local counterKey = KEYS[2]

for i=2, numElements+1 do
    redis.call('HINCRBY', filterKey, ARGV[i], 1)
end

return redis.call('INCRBY', counterKey, itemCount)
`

	countingBloomFilterRemoveMultiScript = `
local function MergeTables(t1, t2)
	for i=1, #t2 do
		table.insert(t1, t2[i])
	end

	return t1
end

local hashIterations = tonumber(ARGV[1])
local numElements = tonumber(#ARGV) - 1
local filterKey = KEYS[1]
local counterKey = KEYS[2]

local hmgetArgs = {}
for i=2, numElements+1 do
    table.insert(hmgetArgs, ARGV[i])
end

local counts = redis.call('HMGET', filterKey, unpack(hmgetArgs))

local existingItemIndexes = {}
local temp = {}
local deleteItemCount = 0
local isExistingItem = true
for i=1, #counts do
	table.insert(temp, ARGV[i+1])

    if (not counts[i]) or (tonumber(counts[i]) == 0)  then
        isExistingItem = false
    end

    if (i % hashIterations == 0) then
        if isExistingItem then
            deleteItemCount = deleteItemCount - 1
			
			existingItemIndexes = MergeTables(existingItemIndexes, temp)
        end

		temp = {}
        isExistingItem = true
    end
end

for i=1, #existingItemIndexes do
    redis.call('HINCRBY', filterKey, existingItemIndexes[i], -1)
end

return redis.call('INCRBY', counterKey, deleteItemCount)
`

	countingBloomFilterDeleteScript = `
local filterKey = KEYS[1]
local counterKey = KEYS[2]

redis.call('DEL', filterKey)
redis.call('DEL', counterKey)

return 1
`
)

// CountingBloomFilter based on Hashes.
// CountingBloomFilter uses 128-bit murmur3 hash function.
type CountingBloomFilter interface {
	// Add adds an item to the Counting Bloom Filter.
	Add(ctx context.Context, key string) error

	// AddMulti adds one or more items to the Counting Bloom Filter.
	// NOTE: If keys are too many, it can block the Redis server for a long time.
	AddMulti(ctx context.Context, keys []string) error

	// Exists checks if an item is in the Counting Bloom Filter.
	Exists(ctx context.Context, key string) (bool, error)

	// ExistsMulti checks if one or more items are in the Counting Bloom Filter.
	// Returns a slice of bool values where each bool indicates
	// whether the corresponding key was found.
	ExistsMulti(ctx context.Context, keys []string) ([]bool, error)

	// Remove removes an item from the Counting Bloom Filter.
	Remove(ctx context.Context, key string) error

	// RemoveMulti removes one or more items from the Counting Bloom Filter.
	// If there are duplicate keys, they are deduplicated.
	// NOTE: If keys are too many, it can block the Redis server for a long time.
	RemoveMulti(ctx context.Context, keys []string) error

	// Delete deletes the Counting Bloom Filter.
	Delete(ctx context.Context) error

	// ItemMinCount returns the minimum count of item in the Counting Bloom Filter.
	// If the item is not in the Counting Bloom Filter, it returns a zero value.
	// Minimum count is not always accurate because of the hash collisions.
	ItemMinCount(ctx context.Context, key string) (uint, error)

	// ItemMinCountMulti returns the minimum count of items in the Counting Bloom Filter.
	// If the item is not in the Counting Bloom Filter, it returns a zero value.
	// Minimum count is not always accurate because of the hash collisions.
	ItemMinCountMulti(ctx context.Context, keys []string) ([]uint, error)

	// Count returns count of items in Counting Bloom Filter.
	Count(ctx context.Context) (uint, error)
}

type countingBloomFilter struct {
	client rueidis.Client

	// name is the name of the Counting Bloom Filter.
	// It is used as a key in the Redis.
	name string

	// counter is the name of the counter.
	counter string

	// hashIterations is the number of hash functions to use.
	hashIterations      uint
	hashIterationString string

	// size is the number of bits to use.
	size uint

	addMultiScript *rueidis.Lua
	addMultiKeys   []string

	removeMultiScript *rueidis.Lua
	removeMultiKeys   []string
}

// NewCountingBloomFilter creates a new Counting Bloom Filter.
// NOTE: 'name:cbf:c' is used as a counter key in the Redis and
// 'name:cbf' is used as a filter key in the Redis
// to keep track of the number of items in the Counting Bloom Filter for Count method.
func NewCountingBloomFilter(
	client rueidis.Client,
	name string,
	expectedNumberOfItems uint,
	falsePositiveRate float64,
) (CountingBloomFilter, error) {
	if len(name) == 0 {
		return nil, ErrEmptyCountingBloomFilterName
	}

	if falsePositiveRate <= 0 {
		return nil, ErrCountingBloomFilterFalsePositiveRateLessThanEqualZero
	}
	if falsePositiveRate >= 1 {
		return nil, ErrCountingBloomFilterFalsePositiveRateGreaterThanOne
	}

	size := numberOfBloomFilterBits(expectedNumberOfItems, falsePositiveRate)
	if size == 0 {
		return nil, ErrCountingBloomFilterBitsSizeZero
	}
	hashIterations := numberOfBloomFilterHashFunctions(size, expectedNumberOfItems)

	// NOTE: https://redis.io/docs/reference/cluster-spec/#hash-tags
	baseName := "{" + name + "}"
	bfName := baseName + ":cbf"
	counterName := bfName + ":c"
	return &countingBloomFilter{
		client:              client,
		name:                bfName,
		counter:             counterName,
		hashIterations:      hashIterations,
		hashIterationString: strconv.FormatUint(uint64(hashIterations), 10),
		size:                size,
		addMultiScript:      rueidis.NewLuaScript(countingBloomFilterAddMultiScript),
		addMultiKeys:        []string{bfName, counterName},
		removeMultiScript:   rueidis.NewLuaScript(countingBloomFilterRemoveMultiScript),
		removeMultiKeys:     []string{bfName, counterName},
	}, nil
}

func (f *countingBloomFilter) Add(ctx context.Context, key string) error {
	return f.AddMulti(ctx, []string{key})
}

func (f *countingBloomFilter) AddMulti(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	indexes := f.indexes(keys)

	args := make([]string, 0, len(indexes)+1)
	args = append(args, strconv.Itoa(len(keys)))
	args = append(args, indexes...)

	resp := f.addMultiScript.Exec(ctx, f.client, f.addMultiKeys, args)
	return resp.Error()
}

func (f *countingBloomFilter) indexes(keys []string) []string {
	allIndexes := make([]string, 0, len(keys)*int(f.hashIterations))
	size := uint64(f.size)
	for _, key := range keys {
		h1, h2 := hash([]byte(key))
		for i := uint(0); i < f.hashIterations; i++ {
			allIndexes = append(allIndexes, strconv.FormatUint(index(h1, h2, i, size), 10))
		}
	}
	return allIndexes
}

func (f *countingBloomFilter) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := f.ExistsMulti(ctx, []string{key})
	if err != nil {
		return false, err
	}

	return exists[0], nil
}

func (f *countingBloomFilter) ExistsMulti(ctx context.Context, keys []string) ([]bool, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	indexes := f.indexes(keys)

	resp := f.client.Do(
		ctx,
		f.client.B().
			Hmget().
			Key(f.name).
			Field(indexes...).
			Build(),
	)
	if resp.Error() != nil {
		return nil, resp.Error()
	}

	messages, err := resp.ToArray()
	if err != nil {
		return nil, err
	}

	result := make([]bool, 0, len(keys))
	isExist := true
	for i, message := range messages {
		cnt, err := message.AsUint64()
		if err != nil {
			if !rueidis.IsRedisNil(err) {
				return nil, err
			}

			isExist = false
		}

		if cnt == 0 {
			isExist = false
		}

		if (i+1)%int(f.hashIterations) == 0 {
			result = append(result, isExist)
			isExist = true
		}
	}

	return result, nil
}

func (f *countingBloomFilter) Remove(ctx context.Context, key string) error {
	return f.RemoveMulti(ctx, []string{key})
}

func (f *countingBloomFilter) RemoveMulti(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	deduplicatedKeys := make([]string, 0, len(keys))
	keySet := make(map[string]struct{}, len(keys))
	for _, key := range keys {
		if _, ok := keySet[key]; !ok {
			keySet[key] = struct{}{}
			deduplicatedKeys = append(deduplicatedKeys, key)
		}
	}

	indexes := f.indexes(deduplicatedKeys)
	args := make([]string, 0, len(indexes)+1)
	args = append(args, f.hashIterationString)
	args = append(args, indexes...)

	resp := f.removeMultiScript.Exec(ctx, f.client, f.removeMultiKeys, args)
	return resp.Error()
}

func (f *countingBloomFilter) Delete(ctx context.Context) error {
	resp := f.client.Do(
		ctx,
		f.client.B().
			Eval().
			Script(countingBloomFilterDeleteScript).
			Numkeys(2).
			Key(f.name, f.counter).
			Build(),
	)
	return resp.Error()
}

func (f *countingBloomFilter) ItemMinCount(ctx context.Context, key string) (uint, error) {
	counts, err := f.ItemMinCountMulti(ctx, []string{key})
	if err != nil {
		return 0, err
	}

	return counts[0], nil
}

func (f *countingBloomFilter) ItemMinCountMulti(ctx context.Context, keys []string) ([]uint, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	indexes := f.indexes(keys)

	resp := f.client.Do(
		ctx,
		f.client.B().
			Hmget().
			Key(f.name).
			Field(indexes...).
			Build(),
	)
	if resp.Error() != nil {
		return nil, resp.Error()
	}

	messages, err := resp.ToArray()
	if err != nil {
		return nil, err
	}

	counts := make([]uint, 0, len(messages))
	minCount := uint64(math.MaxUint64)
	for i, message := range messages {
		cnt, err := message.AsUint64()
		if err != nil {
			if !rueidis.IsRedisNil(err) {
				return nil, err
			}

			minCount = 0
		}

		if cnt < minCount {
			minCount = cnt
		}

		if (i+1)%int(f.hashIterations) == 0 {
			counts = append(counts, uint(minCount))
			minCount = uint64(math.MaxUint64)
		}
	}

	return counts, nil
}

func (f *countingBloomFilter) Count(ctx context.Context) (uint, error) {
	resp := f.client.Do(
		ctx,
		f.client.B().
			Get().
			Key(f.counter).
			Build(),
	)
	count, err := resp.AsUint64()
	if err != nil {
		if rueidis.IsRedisNil(err) {
			return 0, nil
		}

		return 0, err
	}

	return uint(count), nil
}
