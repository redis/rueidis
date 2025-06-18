package rueidisprob

import (
	"context"
	"errors"
	"math"
	"strconv"

	"github.com/redis/rueidis"
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

local numElements = tonumber(#ARGV) - 1
local hashIterations = tonumber(ARGV[#ARGV])
local filterKey = KEYS[1]
local counterKey = KEYS[2]

local indexCounter = {}
for i=1, numElements do
	local index = ARGV[i]
	local count = redis.call('HGET', filterKey, index)

	if (not indexCounter[index]) then
		if (not count) then
			indexCounter[index] = 0
		else
			indexCounter[index] = tonumber(count)
		end
	end
end

local decreaseIndexes = {}
local deleteItemCount = 0
for i=1, numElements, hashIterations do
	local isAbleToRemove = true
	local temp = {}
	local rollbackIndex = i

	for j=i, i+hashIterations-1 do
		local index = ARGV[j]

		table.insert(temp, index)
		indexCounter[index] = indexCounter[index] - 1
		
		if indexCounter[index] < 0 then
			isAbleToRemove = false
			rollbackIndex = j
			break
		end
	end

	if isAbleToRemove then
		decreaseIndexes = MergeTables(decreaseIndexes, temp)
		deleteItemCount = deleteItemCount + 1
	else
		for j=i, rollbackIndex do
			local index = ARGV[j]
			
			indexCounter[index] = indexCounter[index] + 1
		end
	end
end

for i=1, #decreaseIndexes do
    redis.call('HINCRBY', filterKey, decreaseIndexes[i], -1)
end

return redis.call('DECRBY', counterKey, deleteItemCount)
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
// CountingBloomFilter uses a 128-bit murmur3 hash function.
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
	// NOTE: If keys are too many, it can block the Redis server for a long time.
	RemoveMulti(ctx context.Context, keys []string) error

	// Delete deletes the Counting Bloom Filter.
	Delete(ctx context.Context) error

	// ItemMinCount returns the minimum count of item in the Counting Bloom Filter.
	// If the item is not in the Counting Bloom Filter, it returns a zero value.
	// A minimum count is not always accurate because of the hash collisions.
	ItemMinCount(ctx context.Context, key string) (uint64, error)

	// ItemMinCountMulti returns the minimum count of items in the Counting Bloom Filter.
	// If the item is not in the Counting Bloom Filter, it returns a zero value.
	// A minimum count is not always accurate because of the hash collisions.
	ItemMinCountMulti(ctx context.Context, keys []string) ([]uint64, error)

	// Count returns count of items in Counting Bloom Filter.
	Count(ctx context.Context) (uint64, error)
}

type countingBloomFilter struct {
	client rueidis.Client

	addMultiScript *rueidis.Lua

	removeMultiScript *rueidis.Lua

	// name is the name of the Counting Bloom Filter.
	// It is used as a key in the Redis.
	name string

	// counter is the name of the counter.
	counter string

	hashIterationString string

	addMultiKeys []string

	removeMultiKeys []string

	// hashIterations is the number of hash functions to use.
	hashIterations uint

	// size is the number of bits to use.
	size uint
}

// NewCountingBloomFilter creates a new Counting Bloom Filter.
// NOTE: 'name:cbf:c' is used as a counter-key in the Redis and
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

	buf := bytesPool.Get(0, len(keys)*int(f.hashIterations)*8)
	defer bytesPool.Put(buf)

	indexes := f.indexes(keys, &buf.s)

	args := make([]string, 0, len(indexes)+1)
	args = append(args, strconv.Itoa(len(keys)))
	args = append(args, indexes...)

	resp := f.addMultiScript.Exec(ctx, f.client, f.addMultiKeys, args)
	return resp.Error()
}

func (f *countingBloomFilter) indexes(keys []string, buf *[]byte) []string {
	allIndexes := make([]string, 0, len(keys)*int(f.hashIterations))
	size := uint64(f.size)
	for _, key := range keys {
		h1, h2 := hash([]byte(key))
		for i := uint(0); i < f.hashIterations; i++ {
			offset := len(*buf)
			*buf = strconv.AppendUint(*buf, index(h1, h2, i, size), 10)
			allIndexes = append(allIndexes, rueidis.BinaryString((*buf)[offset:]))
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

	buf := bytesPool.Get(0, len(keys)*int(f.hashIterations)*8)
	defer bytesPool.Put(buf)

	indexes := f.indexes(keys, &buf.s)

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

	buf := bytesPool.Get(0, len(keys)*int(f.hashIterations)*8)
	defer bytesPool.Put(buf)

	indexes := f.indexes(keys, &buf.s)

	args := make([]string, 0, len(indexes)+1)
	args = append(args, indexes...)
	args = append(args, f.hashIterationString)

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

func (f *countingBloomFilter) ItemMinCount(ctx context.Context, key string) (uint64, error) {
	counts, err := f.ItemMinCountMulti(ctx, []string{key})
	if err != nil {
		return 0, err
	}

	return counts[0], nil
}

func (f *countingBloomFilter) ItemMinCountMulti(ctx context.Context, keys []string) ([]uint64, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	buf := bytesPool.Get(0, len(keys)*int(f.hashIterations)*8)
	defer bytesPool.Put(buf)

	indexes := f.indexes(keys, &buf.s)

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

	counts := make([]uint64, 0, len(messages))
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
			counts = append(counts, minCount)
			minCount = uint64(math.MaxUint64)
		}
	}

	return counts, nil
}

func (f *countingBloomFilter) Count(ctx context.Context) (uint64, error) {
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

	return count, nil
}
