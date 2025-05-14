package rueidisprob

import (
	"context"
	"errors"
	"math"
	"strconv"

	"github.com/redis/rueidis"
)

const (
	// NOTE: https://redis.io/docs/data-types/bitmaps/
	maxSize = 1 << 32
)

const (
	bloomFilterAddMultiScript = `
local hashIterations = tonumber(ARGV[1])
local numElements = tonumber(#ARGV) - 1
local filterKey = KEYS[1]
local counterKey = KEYS[2]

local counter = 0
local oneBits = 0
for i=1, numElements do
	local bitset = redis.call('BITFIELD', filterKey, 'SET', 'u1', ARGV[i+1], '1')

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		if oneBits ~= hashIterations then
			counter = counter + 1
		end

		oneBits = 0
	end
end

return redis.call('INCRBY', counterKey, counter)
`

	bloomFilterExistsMultiScript = `
local hashIterations = tonumber(ARGV[1])
local numElements = tonumber(#ARGV) - 1
local filterKey = KEYS[1]

local result = {}
local oneBits = 0
for i=1, numElements do
	local index = tonumber(ARGV[i+1])
	local bitset = redis.call('BITFIELD', filterKey, 'GET', 'u1', index)

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		table.insert(result, oneBits == hashIterations)

		oneBits = 0
	end
end

return result
`

	bloomFilterExistsMultiReadOnlyScript = `
local hashIterations = tonumber(ARGV[1])
local numElements = tonumber(#ARGV) - 1
local filterKey = KEYS[1]

local result = {}
local oneBits = 0
for i=1, numElements do
	local index = tonumber(ARGV[i+1])
	local bitset = redis.call('BITFIELD_RO', filterKey, 'GET', 'u1', index)

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		table.insert(result, oneBits == hashIterations)

		oneBits = 0
	end
end

return result
`

	bloomFilterResetScript = `
local filterKey = KEYS[1]
local counterKey = KEYS[2]

redis.call('SET', filterKey, "")
redis.call('SET', counterKey, 0)

return 1
`

	bloomFilterDeleteScript = `
local filterKey = KEYS[1]
local counterKey = KEYS[2]

redis.call('DEL', filterKey)
redis.call('DEL', counterKey)

return 1
`
)

var (
	ErrEmptyName                          = errors.New("name cannot be empty")
	ErrFalsePositiveRateLessThanEqualZero = errors.New("false positive rate cannot be less than or equal to zero")
	ErrFalsePositiveRateGreaterThanOne    = errors.New("false positive rate cannot be greater than 1")
	ErrBitsSizeZero                       = errors.New("bits size cannot be zero")
	ErrBitsSizeTooLarge                   = errors.New("bits size is too large")
)

// BloomFilterOptions is used to configure BloomFilter.
type BloomFilterOptions struct {
	enableReadOperation bool
}

// BloomFilterOptionFunc is used to configure BloomFilter.
type BloomFilterOptionFunc func(*BloomFilterOptions)

// WithEnableReadOperation enables read operation.
// If enabled, Exists and ExistsMulti methods will be available as read-only operations.
// NOTE: If enabled, minimum redis version should be 7.0.0.
func WithEnableReadOperation(enableReadOperations bool) BloomFilterOptionFunc {
	return func(o *BloomFilterOptions) {
		o.enableReadOperation = enableReadOperations
	}
}

// BloomFilter based on Redis Bitmaps.
// BloomFilter uses a 128-bit murmur3 hash function.
type BloomFilter interface {
	// Add adds an item to the Bloom filter.
	Add(ctx context.Context, key string) error

	// AddMulti adds one or more items to the Bloom filter.
	// NOTE: If keys are too many, it can block the Redis server for a long time.
	AddMulti(ctx context.Context, keys []string) error

	// Exists checks if an item is in the Bloom filter.
	Exists(ctx context.Context, key string) (bool, error)

	// ExistsMulti checks if one or more items are in the Bloom filter.
	// Returns a slice of bool values where each bool indicates whether the corresponding key was found.
	ExistsMulti(ctx context.Context, keys []string) ([]bool, error)

	// Reset resets the Bloom filter.
	Reset(ctx context.Context) error

	// Delete deletes the Bloom filter.
	Delete(ctx context.Context) error

	// Count returns count of items in Bloom filter.
	Count(ctx context.Context) (uint64, error)
}

type bloomFilter struct {
	client rueidis.Client

	addMultiScript *rueidis.Lua

	existsMultiScript *rueidis.Lua

	// name is the name of the Bloom filter.
	// It is used as a key in the Redis.
	name string

	// counter is the name of the counter.
	counter string

	hashIterationString string

	addMultiKeys []string

	existsMultiKeys []string

	// hashIterations is the number of hash functions to use.
	hashIterations uint

	// size is the number of bits to use.
	size uint
}

// NewBloomFilter creates a new Bloom filter.
// NOTE: 'name:c' is used as a counter-key in the Redis
// to keep track of the number of items in the Bloom filter for Count method.
func NewBloomFilter(
	client rueidis.Client,
	name string,
	expectedNumberOfItems uint,
	falsePositiveRate float64,
	opts ...BloomFilterOptionFunc,
) (BloomFilter, error) {
	if len(name) == 0 {
		return nil, ErrEmptyName
	}

	if falsePositiveRate <= 0 {
		return nil, ErrFalsePositiveRateLessThanEqualZero
	}
	if falsePositiveRate > 1 {
		return nil, ErrFalsePositiveRateGreaterThanOne
	}

	size := numberOfBloomFilterBits(expectedNumberOfItems, falsePositiveRate)
	if size == 0 {
		return nil, ErrBitsSizeZero
	}
	if size > maxSize {
		return nil, ErrBitsSizeTooLarge
	}
	hashIterations := numberOfBloomFilterHashFunctions(size, expectedNumberOfItems)

	options := &BloomFilterOptions{}
	for _, opt := range opts {
		opt(options)
	}

	var existsMultiScript *rueidis.Lua
	if options.enableReadOperation {
		existsMultiScript = rueidis.NewLuaScriptReadOnly(bloomFilterExistsMultiReadOnlyScript)
	} else {
		existsMultiScript = rueidis.NewLuaScript(bloomFilterExistsMultiScript)
	}

	// NOTE: https://redis.io/docs/reference/cluster-spec/#hash-tags
	bfName := "{" + name + "}"
	counterName := bfName + ":c"
	return &bloomFilter{
		client:              client,
		name:                bfName,
		counter:             counterName,
		hashIterations:      hashIterations,
		hashIterationString: strconv.FormatUint(uint64(hashIterations), 10),
		size:                size,
		addMultiScript:      rueidis.NewLuaScript(bloomFilterAddMultiScript),
		addMultiKeys:        []string{bfName, counterName},
		existsMultiScript:   existsMultiScript,
		existsMultiKeys:     []string{bfName},
	}, nil
}

func numberOfBloomFilterBits(n uint, r float64) uint {
	return uint(math.Ceil(-float64(n) * math.Log(r) / math.Pow(math.Log(2), 2)))
}

func numberOfBloomFilterHashFunctions(s uint, n uint) uint {
	return uint(math.Round(float64(s) / float64(n) * math.Log(2)))
}

func (c *bloomFilter) Add(ctx context.Context, key string) error {
	return c.AddMulti(ctx, []string{key})
}

func (c *bloomFilter) AddMulti(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	buf := bytesPool.Get(0, len(keys)*int(c.hashIterations)*8)
	defer bytesPool.Put(buf)

	indexes := c.indexes(keys, &buf.s)

	args := make([]string, 0, len(indexes)+1)
	args = append(args, c.hashIterationString)
	args = append(args, indexes...)

	resp := c.addMultiScript.Exec(ctx, c.client, c.addMultiKeys, args)
	return resp.Error()
}

func (c *bloomFilter) indexes(keys []string, buf *[]byte) []string {
	allIndexes := make([]string, 0, len(keys)*int(c.hashIterations))
	size := uint64(c.size)
	for _, key := range keys {
		h1, h2 := hash([]byte(key))
		for i := uint(0); i < c.hashIterations; i++ {
			offset := len(*buf)
			*buf = strconv.AppendUint(*buf, index(h1, h2, i, size), 10)
			allIndexes = append(allIndexes, rueidis.BinaryString((*buf)[offset:]))
		}
	}
	return allIndexes
}

func (c *bloomFilter) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := c.ExistsMulti(ctx, []string{key})
	if err != nil {
		return false, err
	}

	return exists[0], nil
}

func (c *bloomFilter) ExistsMulti(ctx context.Context, keys []string) ([]bool, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	buf := bytesPool.Get(0, len(keys)*int(c.hashIterations)*8)
	defer bytesPool.Put(buf)

	indexes := c.indexes(keys, &buf.s)

	args := make([]string, 0, len(indexes)+1)
	args = append(args, c.hashIterationString)
	args = append(args, indexes...)

	resp := c.existsMultiScript.Exec(ctx, c.client, c.existsMultiKeys, args)
	if resp.Error() != nil {
		return nil, resp.Error()
	}

	arr, err := resp.ToArray()
	if err != nil {
		return nil, err
	}

	result := make([]bool, len(keys))
	for i, el := range arr {
		v, err := el.AsBool()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				result[i] = false
				continue
			}

			return nil, err
		}

		result[i] = v
	}
	return result, nil
}

func (c *bloomFilter) Reset(ctx context.Context) error {
	resp := c.client.Do(
		ctx,
		c.client.B().
			Eval().
			Script(bloomFilterResetScript).
			Numkeys(2).
			Key(c.name, c.counter).
			Build(),
	)
	return resp.Error()
}

func (c *bloomFilter) Delete(ctx context.Context) error {
	resp := c.client.Do(
		ctx,
		c.client.B().
			Eval().
			Script(bloomFilterDeleteScript).
			Numkeys(2).
			Key(c.name, c.counter).
			Build(),
	)
	return resp.Error()
}

func (c *bloomFilter) Count(ctx context.Context) (uint64, error) {
	resp := c.client.Do(
		ctx,
		c.client.B().
			Get().
			Key(c.counter).
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
