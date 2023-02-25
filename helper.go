package rueidis

import (
	"context"
	"sync"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/util"
)

// MGetCache is a helper that consults the client-side caches with multiple keys by grouping keys within same slot into MGETs
func MGetCache(client Client, ctx context.Context, ttl time.Duration, keys []string) (ret map[string]RedisMessage, err error) {
	if len(keys) == 0 {
		return make(map[string]RedisMessage), nil
	}
	if cc, ok := client.(*clusterClient); ok {
		return clusterMGetCache(cc, ctx, ttl, cmds.MGets(keys), keys)
	}
	return clientMGetCache(client, ctx, ttl, client.B().Mget().Key(keys...).Cache(), keys)
}

// MGet is a helper that consults the redis directly with multiple keys by grouping keys within same slot into MGETs
func MGet(client Client, ctx context.Context, keys []string) (ret map[string]RedisMessage, err error) {
	if len(keys) == 0 {
		return make(map[string]RedisMessage), nil
	}
	if cc, ok := client.(*clusterClient); ok {
		return clusterMGet(cc, ctx, cmds.MGets(keys), keys)
	}
	return clientMGet(client, ctx, client.B().Mget().Key(keys...).Build(), keys)
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

// JsonMGet is a helper that consults redis directly with multiple keys by grouping keys within same slot into JSON.MGETs
func JsonMGet(client Client, ctx context.Context, keys []string, path string) (ret map[string]RedisMessage, err error) {
	if len(keys) == 0 {
		return make(map[string]RedisMessage), nil
	}
	if cc, ok := client.(*clusterClient); ok {
		return clusterMGet(cc, ctx, cmds.JsonMGets(keys, path), keys)
	}
	return clientMGet(client, ctx, client.B().JsonMget().Key(keys...).Path(path).Build(), keys)
}

func clientMGetCache(client Client, ctx context.Context, ttl time.Duration, cmd cmds.Cacheable, keys []string) (ret map[string]RedisMessage, err error) {
	arr, err := client.DoCache(ctx, cmd, ttl).ToArray()
	if err != nil {
		return nil, err
	}
	return arrayToKV(make(map[string]RedisMessage, len(keys)), arr, keys), nil
}

func clientMGet(client Client, ctx context.Context, cmd cmds.Completed, keys []string) (ret map[string]RedisMessage, err error) {
	arr, err := client.Do(ctx, cmd).ToArray()
	if err != nil {
		return nil, err
	}
	return arrayToKV(make(map[string]RedisMessage, len(keys)), arr, keys), nil
}

func clusterMGetCache(cc *clusterClient, ctx context.Context, ttl time.Duration, mgets map[uint16]cmds.Completed, keys []string) (ret map[string]RedisMessage, err error) {
	return doMGets(make(map[string]RedisMessage, len(keys)), mgets, func(cmd cmds.Completed) RedisResult {
		return cc.doCache(ctx, cmds.Cacheable(cmd), ttl)
	})
}

func clusterMGet(cc *clusterClient, ctx context.Context, mgets map[uint16]cmds.Completed, keys []string) (ret map[string]RedisMessage, err error) {
	return doMGets(make(map[string]RedisMessage, len(keys)), mgets, func(cmd cmds.Completed) RedisResult {
		return cc.do(ctx, cmd)
	})
}

func doMGets(m map[string]RedisMessage, mgets map[uint16]cmds.Completed, fn func(cmd cmds.Completed) RedisResult) (ret map[string]RedisMessage, err error) {
	var mu sync.Mutex
	util.ParallelVals(mgets, func(cmd cmds.Completed) {
		arr, err2 := fn(cmd).ToArray()
		mu.Lock()
		if err2 != nil {
			err = err2
		} else {
			arrayToKV(m, arr, cmd.Commands()[1:])
		}
		mu.Unlock()
	})
	if err != nil {
		return nil, err
	}
	for _, mget := range mgets { // not recycle cmds if error, since cmds may be used later in pipe. consider recycle them by pipe
		cmds.Put(mget.CommandSlice())
	}
	return m, nil
}

func arrayToKV(m map[string]RedisMessage, arr []RedisMessage, keys []string) map[string]RedisMessage {
	for i, resp := range arr {
		m[keys[i]] = resp
	}
	return m
}
