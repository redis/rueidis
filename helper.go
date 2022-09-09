package rueidis

import (
	"context"
	"runtime"
	"sync"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

// MGetCache is a helper that consults the client-side caches with multiple keys by grouping keys within same slot into MGETs
func MGetCache(client Client, ctx context.Context, ttl time.Duration, keys []string) (ret map[string]RedisMessage, err error) {
	if len(keys) == 0 {
		return make(map[string]RedisMessage), nil
	}
	ret = make(map[string]RedisMessage, len(keys))
	if cc, ok := client.(*clusterClient); ok {
		return clusterMGetCache(cc, ctx, ttl, cmds.MGets(keys), keys)
	}
	return clientMGetCache(client, ctx, ttl, client.B().Mget().Key(keys...).Cache(), keys)
}

// JsonMGetCache is a helper that consults the client-side caches with multiple keys by grouping keys within same slot into JSON.MGETs
func JsonMGetCache(client Client, ctx context.Context, ttl time.Duration, keys []string, path string) (ret map[string]RedisMessage, err error) {
	if len(keys) == 0 {
		return make(map[string]RedisMessage), nil
	}
	if cc, ok := client.(*clusterClient); ok {
		return clusterMGetCache(cc, ctx, ttl, cmds.JsonMGets(keys, path), keys)
	}
	return clientMGetCache(client, ctx, ttl, client.B().JsonMget().Key(keys...).Path(path).Cache(), keys)
}

func clientMGetCache(client Client, ctx context.Context, ttl time.Duration, cmd cmds.Cacheable, keys []string) (ret map[string]RedisMessage, err error) {
	arr, err := client.DoCache(ctx, cmd, ttl).ToArray()
	if err != nil {
		return nil, err
	}
	ret = make(map[string]RedisMessage, len(keys))
	for i, resp := range arr {
		ret[keys[i]] = resp
	}
	return ret, nil
}

func clusterMGetCache(cc *clusterClient, ctx context.Context, ttl time.Duration, mgets map[uint16]cmds.Completed, keys []string) (ret map[string]RedisMessage, err error) {
	var mu sync.Mutex
	ret = make(map[string]RedisMessage, len(keys))
	parallelVals(mgets, func(cmd cmds.Completed) {
		c := cmds.Cacheable(cmd)
		arr, err2 := cc.doCache(ctx, c, ttl).ToArray()
		mu.Lock()
		if err2 != nil {
			err = err2
		} else {
			for i, resp := range arr {
				ret[c.MGetCacheKey(i)] = resp
			}
		}
		mu.Unlock()
	})
	if err != nil {
		return nil, err
	}
	for _, mget := range mgets { // not recycle cmds if error, since cmds may be used later in pipe. consider recycle them by pipe
		cmds.Put(mget.CommandSlice())
	}
	return ret, nil
}

func parallelKeys[K comparable, V any](p map[K]V, fn func(k K)) {
	ch := make(chan K, len(p))
	for k := range p {
		ch <- k
	}
	closeThenParallel(ch, fn)
}

func parallelVals[K comparable, V any](p map[K]V, fn func(k V)) {
	ch := make(chan V, len(p))
	for _, v := range p {
		ch <- v
	}
	closeThenParallel(ch, fn)
}

func closeThenParallel[V any](ch chan V, fn func(k V)) {
	close(ch)
	concurrency := len(ch)
	if cpus := runtime.NumCPU(); concurrency > cpus {
		concurrency = cpus
	}
	wg := sync.WaitGroup{}
	wg.Add(concurrency)
	for i := 1; i < concurrency; i++ {
		go func() {
			for v := range ch {
				fn(v)
			}
			wg.Done()
		}()
	}
	for v := range ch {
		fn(v)
	}
	wg.Done()
	wg.Wait()
}
