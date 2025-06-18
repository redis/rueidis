package rueidisprob

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/redis/rueidis"
)

const (
	slidingBloomFilterInitializeScript = `
local filterKey = KEYS[1]
local nextFilterKey = KEYS[2]
local counterKey = KEYS[3]
local nextCounterKey = KEYS[4]
local lastRotationKey = KEYS[5]
local windowHalf = tonumber(ARGV[1])

if redis.call('EXISTS', filterKey, nextFilterKey, counterKey, nextCounterKey, lastRotationKey) == 0 then
	local time = redis.call('TIME')
	local current_time = tonumber(time[1]) * 1000 + math.floor(tonumber(time[2]) / 1000)

	redis.call('MSET', filterKey, "", counterKey, 0, nextFilterKey, "", nextCounterKey, 0)
	redis.call('SET', lastRotationKey, tostring(current_time), 'PX', windowHalf, 'NX')
end

return 1
`

	slidingBloomFilterAddMultiScript = `
local hashIterations = tonumber(ARGV[1])
local windowHalf = tonumber(ARGV[2])
local numElements = tonumber(#ARGV) - 2

local filterKey = KEYS[1]
local nextFilterKey = KEYS[2]
local counterKey = KEYS[3]
local nextCounterKey = KEYS[4]
local lastRotationKey = KEYS[5]

local time = redis.call('TIME')
local current_time = tonumber(time[1]) * 1000 + math.floor(tonumber(time[2])/1000)
local acquiredLock = redis.call('SET', lastRotationKey, tostring(current_time), 'PX', windowHalf, 'NX')

if acquiredLock then
	redis.call('RENAME', nextFilterKey, filterKey)
	redis.call('RENAME', nextCounterKey, counterKey)
	redis.call('SET', nextFilterKey, "")
	redis.call('SET', nextCounterKey, 0)
end

local counter = 0
local oneBits = 0
for i=1, numElements do
	local bitset = redis.call('BITFIELD', filterKey, 'SET', 'u1', ARGV[i+2], '1')
	redis.call('BITFIELD', nextFilterKey, 'SET', 'u1', ARGV[i+2], '1')

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		if oneBits ~= hashIterations then
			counter = counter + 1
		end

		oneBits = 0
	end
end

redis.call('INCRBY', nextCounterKey, counter)
return redis.call('INCRBY', counterKey, counter)
`
	slidingBloomFilterExistsMultiScript = `
local hashIterations = tonumber(ARGV[1])
local windowHalf = tonumber(ARGV[2])
local numElements = tonumber(#ARGV) - 2

local filterKey = KEYS[1]
local nextFilterKey = KEYS[2]
local counterKey = KEYS[3]
local nextCounterKey = KEYS[4]
local lastRotationKey = KEYS[5]

local time = redis.call('TIME')
local current_time = tonumber(time[1]) * 1000 + math.floor(tonumber(time[2])/1000)
local acquiredLock = redis.call('SET', lastRotationKey, tostring(current_time), 'PX', windowHalf, 'NX')

if acquiredLock then
	redis.call('RENAME', nextFilterKey, filterKey)
	redis.call('RENAME', nextCounterKey, counterKey)
	redis.call('SET', nextFilterKey, "")
	redis.call('SET', nextCounterKey, 0)
end

local result = {}
local oneBits = 0
for i=1, numElements do
	local index = tonumber(ARGV[i+2])
	local bitset = redis.call('BITFIELD', filterKey, 'GET', 'u1', index)

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		table.insert(result, oneBits == hashIterations)

		oneBits = 0
	end
end

return result
`

	slidingBloomFilterExistsReadOnlyMultiScript = `
local hashIterations = tonumber(ARGV[1])
local windowHalf = tonumber(ARGV[2])
local numElements = tonumber(#ARGV) - 2

local filterKey = KEYS[1]
local nextFilterKey = KEYS[2]
local counterKey = KEYS[3]
local nextCounterKey = KEYS[4]
local lastRotationKey = KEYS[5]

local time = redis.call('TIME')
local current_time = tonumber(time[1]) * 1000 + math.floor(tonumber(time[2])/1000)
local acquiredLock = redis.call('SET', lastRotationKey, tostring(current_time), 'PX', windowHalf, 'NX')

if acquiredLock then
	redis.call('RENAME', nextFilterKey, filterKey)
	redis.call('RENAME', nextCounterKey, counterKey)
	redis.call('SET', nextFilterKey, "")
	redis.call('SET', nextCounterKey, 0)
end

local result = {}
local oneBits = 0
for i=1, numElements do
	local index = tonumber(ARGV[i+2])
	local bitset = redis.call('BITFIELD_RO', filterKey, 'GET', 'u1', index)

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		table.insert(result, oneBits == hashIterations)

		oneBits = 0
	end
end

return result
`

	slidingBloomFilterResetScript = `
local filterKey = KEYS[1]
local nextFilterKey = KEYS[2]
local counterKey = KEYS[3]
local nextCounterKey = KEYS[4]

redis.call('RENAME', nextFilterKey, filterKey)
redis.call('RENAME', nextCounterKey, counterKey)
redis.call('SET', nextFilterKey, "")
redis.call('SET', nextCounterKey, 0)
`

	// Redis key suffixes
	counterSuffix      = ":c"
	nextFilterSuffix   = ":n"
	nextCounterSuffix  = ":nc"
	lastRotationSuffix = ":lr"
)

var (
	_                              BloomFilter = (*slidingBloomFilter)(nil)
	ErrWindowSizeLessThanOneSecond             = errors.New("window size cannot be less than 1 second")
)

type SlidingBloomFilterOptions struct {
	enableReadOperation bool
}

type SlidingBloomFilterOptionFunc func(o *SlidingBloomFilterOptions)

func WithReadOnlyExists(enableReadOperations bool) SlidingBloomFilterOptionFunc {
	return func(o *SlidingBloomFilterOptions) {
		o.enableReadOperation = enableReadOperations
	}
}

type slidingBloomFilter struct {
	client rueidis.Client

	addMultiScript *rueidis.Lua

	existsMultiScript *rueidis.Lua

	// Pre-calculated window half in milliseconds
	windowHalfMs string

	// name is the name of the sliding Bloom filter.
	// It is used as a key in the Redis.
	name string

	// counter is the name of the counter.
	counter string

	hashIterationString string

	addMultiKeys []string

	existsMultiKeys []string

	// window is the duration of the sliding window.
	window time.Duration

	// hashIterations is the number of hash functions to use.
	hashIterations uint

	// size is the number of bits to use.
	size uint
}

// NewSlidingBloomFilter creates a new sliding window Bloom filter.
// NOTE: 'name:c' is used as a counter-key in the Redis
// 'name:n' is used as a next filter key in the Redis
// 'name:nc' is used as a next counter key in the Redis
// 'name:lr' is used as a last rotation key in the Redis
// to keep track of the items in the window.
func NewSlidingBloomFilter(
	redisClient rueidis.Client,
	name string,
	expectedNumberOfItems uint,
	falsePositiveRate float64,
	windowSize time.Duration,
	opts ...SlidingBloomFilterOptionFunc,
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
	if windowSize < time.Second {
		return nil, ErrWindowSizeLessThanOneSecond
	}

	size := numberOfBloomFilterBits(expectedNumberOfItems, falsePositiveRate)
	if size == 0 {
		return nil, ErrBitsSizeZero
	}
	if size > maxSize {
		return nil, ErrBitsSizeTooLarge
	}
	hashIterations := numberOfBloomFilterHashFunctions(size, expectedNumberOfItems)

	options := &SlidingBloomFilterOptions{}
	for _, opt := range opts {
		opt(options)
	}

	var existsMultiScript string
	if options.enableReadOperation {
		existsMultiScript = slidingBloomFilterExistsReadOnlyMultiScript
	} else {
		existsMultiScript = slidingBloomFilterExistsMultiScript
	}

	// NOTE: https://redis.io/docs/reference/cluster-spec/#hash-tags
	bfName := "{" + name + "}"
	counterName := bfName + counterSuffix
	nextFilterName := bfName + nextFilterSuffix
	nextCounterName := bfName + nextCounterSuffix
	lastRotationName := bfName + lastRotationSuffix

	s := &slidingBloomFilter{
		client:              redisClient,
		name:                bfName,
		counter:             counterName,
		window:              windowSize,
		windowHalfMs:        strconv.FormatInt(windowSize.Milliseconds()/2, 10),
		hashIterations:      hashIterations,
		hashIterationString: strconv.FormatUint(uint64(hashIterations), 10),
		size:                size,
		addMultiScript:      rueidis.NewLuaScript(slidingBloomFilterAddMultiScript),
		addMultiKeys:        []string{bfName, nextFilterName, counterName, nextCounterName, lastRotationName},
		existsMultiScript:   rueidis.NewLuaScript(existsMultiScript),
		existsMultiKeys:     []string{bfName, nextFilterName, counterName, nextCounterName, lastRotationName},
	}

	err := s.initialize()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *slidingBloomFilter) initialize() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	initializeScript := rueidis.NewLuaScript(slidingBloomFilterInitializeScript)
	resp := initializeScript.Exec(ctx, s.client, s.addMultiKeys, []string{s.windowHalfMs})
	if resp.Error() != nil && !rueidis.IsRedisNil(resp.Error()) {
		return resp.Error()
	}

	v, err := resp.AsInt64()
	if err != nil {
		return err
	}

	if v != 1 {
		return errors.New("failed to initialize sliding Bloom filter")
	}

	return nil
}

func (s *slidingBloomFilter) Add(ctx context.Context, key string) error {
	return s.AddMulti(ctx, []string{key})
}

func (s *slidingBloomFilter) AddMulti(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	buf := bytesPool.Get(0, len(keys)*int(s.hashIterations)*8)
	defer bytesPool.Put(buf)

	indexes := s.indexes(keys, &buf.s)

	args := make([]string, 0, len(indexes)+2)
	args = append(args, s.hashIterationString)
	args = append(args, s.windowHalfMs)
	args = append(args, indexes...)

	resp := s.addMultiScript.Exec(ctx, s.client, s.addMultiKeys, args)
	return resp.Error()
}

func (s *slidingBloomFilter) indexes(keys []string, buf *[]byte) []string {
	allIndexes := make([]string, 0, len(keys)*int(s.hashIterations))
	size := uint64(s.size)

	for _, key := range keys {
		h1, h2 := hash([]byte(key))
		for i := uint(0); i < s.hashIterations; i++ {
			offset := len(*buf)
			*buf = strconv.AppendUint(*buf, index(h1, h2, i, size), 10)
			allIndexes = append(allIndexes, rueidis.BinaryString((*buf)[offset:]))
		}
	}
	return allIndexes
}

func (s *slidingBloomFilter) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := s.ExistsMulti(ctx, []string{key})
	if err != nil {
		return false, err
	}

	return exists[0], nil
}

func (s *slidingBloomFilter) ExistsMulti(ctx context.Context, keys []string) ([]bool, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	buf := bytesPool.Get(0, len(keys)*int(s.hashIterations)*8)
	defer bytesPool.Put(buf)

	indexes := s.indexes(keys, &buf.s)

	args := make([]string, 0, len(indexes)+2)
	args = append(args, s.hashIterationString)
	args = append(args, s.windowHalfMs)
	args = append(args, indexes...)

	resp := s.existsMultiScript.Exec(ctx, s.client, s.existsMultiKeys, args)
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

func (s *slidingBloomFilter) Reset(ctx context.Context) error {
	resp := s.client.Do(ctx,
		s.client.B().
			Eval().
			Script(slidingBloomFilterResetScript).
			Numkeys(4).
			Key(s.addMultiKeys...).
			Build(),
	)
	return resp.Error()
}

func (s *slidingBloomFilter) Delete(ctx context.Context) error {
	resp := s.client.Do(ctx, s.client.B().Del().Key(s.addMultiKeys...).Build())
	return resp.Error()
}

func (s *slidingBloomFilter) Count(ctx context.Context) (uint64, error) {
	resp := s.client.Do(
		ctx,
		s.client.B().
			Get().
			Key(s.counter).
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
