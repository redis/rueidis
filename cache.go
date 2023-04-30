package rueidis

import (
	"context"
	"time"
)

// NewCacheStoreFn can be provided in ClientOption for using a custom CacheStore implementation
type NewCacheStoreFn func(CacheStoreOption) CacheStore

// CacheStoreOption will be passed to NewCacheStoreFn
type CacheStoreOption struct {
	// CacheSizeEachConn is redis client side cache size that bind to each TCP connection to a single redis instance.
	// The default is DefaultCacheBytes.
	CacheSizeEachConn int
}

// CacheStore is the store interface for the client side caching
// More detailed interface requirement can be found in cache_test.go
type CacheStore interface {
	// Flight is called when DoCache and DoMultiCache, with the requested client side ttl and the current time.
	// It should look up the store in single-flight manner and return one of the following three combinations:
	// Case 1: (empty RedisMessage, nil CacheEntry)         <- when cache missed, and rueidis will send the request to redis.
	// Case 2: (empty RedisMessage, non-nil CacheEntry)     <- when cache missed, and rueidis will use CacheEntry.Wait to wait for response.
	// Case 3: (non-empty RedisMessage, non-nil CacheEntry) <- when cache hit
	Flight(key, cmd string, ttl time.Duration, now time.Time) (v RedisMessage, e CacheEntry)
	// Update is called when receiving the response of the request sent by the above Flight Case 1 from redis.
	// It should not only update the store but also deliver the response to all CacheEntry.Wait and return a desired client side PTTL of the response.
	// Note that the server side expire time can be retrieved from RedisMessage.CachePXAT.
	Update(key, cmd string, val RedisMessage) (pttl int64)
	// Cancel is called when the request sent by the above Flight Case 1 failed.
	// It should not only deliver the error to all CacheEntry.Wait but also remove the CacheEntry from the store.
	Cancel(key, cmd string, err error)
	// Delete is called when receiving invalidation notifications from redis.
	// If the keys is nil then it should delete all non-pending cached entries under all keys.
	// If the keys is not nil then it should delete all non-pending cached entries under those keys.
	Delete(keys []RedisMessage)
	// Close is called when connection between redis is broken.
	// It should flush all cached entries and deliver the error to all pending CacheEntry.Wait.
	Close(err error)
}

// CacheEntry should be used to wait for single-flight response when cache missed.
type CacheEntry interface {
	Wait(ctx context.Context) (RedisMessage, error)
}
