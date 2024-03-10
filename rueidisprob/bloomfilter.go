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
	addMultiScript = `
local hashIterations = tonumber(ARGV[1])
local numElements = tonumber(#ARGV) - 1
local filterKey = KEYS[1]
local counterKey = KEYS[2]

local bitfieldArgs = { filterKey }
for i=2, numElements+1 do
    table.insert(bitfieldArgs, 'SET')
    table.insert(bitfieldArgs, 'u1')
    table.insert(bitfieldArgs, ARGV[i])
    table.insert(bitfieldArgs, '1')
end

local setBits = redis.call('BITFIELD', unpack(bitfieldArgs))

local counter = 0
local oneBits = 0
for i=1, #setBits do
	oneBits = oneBits + setBits[i]

	if i % hashIterations == 0 then
		if oneBits ~= hashIterations then
			counter = counter + 1
		end

		oneBits = 0
	end
end

return redis.call('INCRBY', counterKey, counter)
`

	existsMultiScript = `
local numElements = tonumber(#ARGV)
local filterKey = KEYS[1]
local result = {}

local bitfieldArgs = { filterKey }
for i=1, numElements do
	local index = tonumber(ARGV[i])

	table.insert(bitfieldArgs, 'GET')
	table.insert(bitfieldArgs, 'u1')
	table.insert(bitfieldArgs, index)
end

return redis.call('BITFIELD', unpack(bitfieldArgs))
`

	resetScript = `
local filterKey = KEYS[1]
local counterKey = KEYS[2]

redis.call('SET', filterKey, "")
redis.call('SET', counterKey, 0)

return 1
`

	deleteScript = `
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

// BloomFilter based on Redis Bitmaps.
// BloomFilter uses 128-bit murmur3 hash function.
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
	Count(ctx context.Context) (uint, error)
}

type bloomFilter struct {
	client rueidis.Client

	// name is the name of the Bloom filter.
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

	existsMultiScript *rueidis.Lua
	existsMultiKeys   []string
}

// NewBloomFilter creates a new Bloom filter.
// NOTE: 'name:c' is used as a counter key in the Redis
// to keep track of the number of items in the Bloom filter for Count method.
func NewBloomFilter(
	client rueidis.Client,
	name string,
	expectedNumberOfItems uint,
	falsePositiveRate float64,
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

	size := numberOfBits(expectedNumberOfItems, falsePositiveRate)
	if size == 0 {
		return nil, ErrBitsSizeZero
	}
	if size > maxSize {
		return nil, ErrBitsSizeTooLarge
	}
	hashIterations := numberOfHashFunctions(size, expectedNumberOfItems)

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
		addMultiScript:      rueidis.NewLuaScript(addMultiScript),
		addMultiKeys:        []string{bfName, counterName},
		existsMultiScript:   rueidis.NewLuaScript(existsMultiScript),
		existsMultiKeys:     []string{bfName},
	}, nil
}

func numberOfBits(n uint, r float64) uint {
	return uint(math.Ceil(-float64(n) * math.Log(r) / math.Pow(math.Log(2), 2)))
}

func numberOfHashFunctions(s uint, n uint) uint {
	return uint(math.Round(float64(s) / float64(n) * math.Log(2)))
}

func (c *bloomFilter) Add(ctx context.Context, key string) error {
	return c.AddMulti(ctx, []string{key})
}

func (c *bloomFilter) AddMulti(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	indexes := c.indexes(keys)

	var args []string
	args = append(args, c.hashIterationString)
	args = append(args, indexes...)

	resp := c.addMultiScript.Exec(ctx, c.client, c.addMultiKeys, args)
	if resp.Error() != nil {
		return resp.Error()
	}

	return nil
}

func (c *bloomFilter) indexes(keys []string) []string {
	var allIndexes []string
	for _, key := range keys {
		indices := indexes([]byte(key), c.hashIterations, c.size)
		for _, index := range indices {
			allIndexes = append(allIndexes, strconv.FormatUint(index, 10))
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

	indexes := c.indexes(keys)

	resp := c.existsMultiScript.Exec(ctx, c.client, c.existsMultiKeys, indexes)
	if resp.Error() != nil {
		return nil, resp.Error()
	}

	is, err := resp.AsIntSlice()
	if err != nil {
		return nil, err
	}

	result := make([]bool, len(keys))
	resultIdx := 0
	oneBits := 0
	iterations := int(c.hashIterations)
	for i, v := range is {
		if v == 1 {
			oneBits++
		}

		if (i+1)%iterations == 0 {
			if oneBits == iterations {
				result[resultIdx] = true
			}

			resultIdx++
			oneBits = 0
		}
	}
	return result, nil
}

func (c *bloomFilter) Reset(ctx context.Context) error {
	resp := c.client.Do(
		ctx,
		c.client.B().
			Eval().
			Script(resetScript).
			Numkeys(2).
			Key(c.name, c.counter).
			Build(),
	)
	if resp.Error() != nil {
		return resp.Error()
	}

	return nil
}

func (c *bloomFilter) Delete(ctx context.Context) error {
	resp := c.client.Do(
		ctx,
		c.client.B().
			Eval().
			Script(deleteScript).
			Numkeys(2).
			Key(c.name, c.counter).
			Build(),
	)
	if resp.Error() != nil {
		return resp.Error()
	}

	return nil
}

func (c *bloomFilter) Count(ctx context.Context) (uint, error) {
	resp := c.client.Do(
		ctx,
		c.client.B().
			Get().
			Key(c.counter).
			Build(),
	)
	if resp.Error() != nil {
		if rueidis.IsRedisNil(resp.Error()) {
			return 0, nil
		}

		return 0, resp.Error()
	}

	count, err := resp.AsUint64()
	if err != nil {
		return 0, err
	}

	return uint(count), nil
}
