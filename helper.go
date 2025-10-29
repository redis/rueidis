package rueidis

import (
	"context"
	"errors"
	"iter"
	"time"

	intl "github.com/redis/rueidis/internal/cmds"
)

var crc16tab = [256]uint16{
	0x0000, 0x1021, 0x2042, 0x3063, 0x4084, 0x50a5, 0x60c6, 0x70e7,
	0x8108, 0x9129, 0xa14a, 0xb16b, 0xc18c, 0xd1ad, 0xe1ce, 0xf1ef,
	0x1231, 0x0210, 0x3273, 0x2252, 0x52b5, 0x4294, 0x72f7, 0x62d6,
	0x9339, 0x8318, 0xb37b, 0xa35a, 0xd3bd, 0xc39c, 0xf3ff, 0xe3de,
	0x2462, 0x3443, 0x0420, 0x1401, 0x64e6, 0x74c7, 0x44a4, 0x5485,
	0xa56a, 0xb54b, 0x8528, 0x9509, 0xe5ee, 0xf5cf, 0xc5ac, 0xd58d,
	0x3653, 0x2672, 0x1611, 0x0630, 0x76d7, 0x66f6, 0x5695, 0x46b4,
	0xb75b, 0xa77a, 0x9719, 0x8738, 0xf7df, 0xe7fe, 0xd79d, 0xc7bc,
	0x48c4, 0x58e5, 0x6886, 0x78a7, 0x0840, 0x1861, 0x2802, 0x3823,
	0xc9cc, 0xd9ed, 0xe98e, 0xf9af, 0x8948, 0x9969, 0xa90a, 0xb92b,
	0x5af5, 0x4ad4, 0x7ab7, 0x6a96, 0x1a71, 0x0a50, 0x3a33, 0x2a12,
	0xdbfd, 0xcbdc, 0xfbbf, 0xeb9e, 0x9b79, 0x8b58, 0xbb3b, 0xab1a,
	0x6ca6, 0x7c87, 0x4ce4, 0x5cc5, 0x2c22, 0x3c03, 0x0c60, 0x1c41,
	0xedae, 0xfd8f, 0xcdec, 0xddcd, 0xad2a, 0xbd0b, 0x8d68, 0x9d49,
	0x7e97, 0x6eb6, 0x5ed5, 0x4ef4, 0x3e13, 0x2e32, 0x1e51, 0x0e70,
	0xff9f, 0xefbe, 0xdfdd, 0xcffc, 0xbf1b, 0xaf3a, 0x9f59, 0x8f78,
	0x9188, 0x81a9, 0xb1ca, 0xa1eb, 0xd10c, 0xc12d, 0xf14e, 0xe16f,
	0x1080, 0x00a1, 0x30c2, 0x20e3, 0x5004, 0x4025, 0x7046, 0x6067,
	0x83b9, 0x9398, 0xa3fb, 0xb3da, 0xc33d, 0xd31c, 0xe37f, 0xf35e,
	0x02b1, 0x1290, 0x22f3, 0x32d2, 0x4235, 0x5214, 0x6277, 0x7256,
	0xb5ea, 0xa5cb, 0x95a8, 0x8589, 0xf56e, 0xe54f, 0xd52c, 0xc50d,
	0x34e2, 0x24c3, 0x14a0, 0x0481, 0x7466, 0x6447, 0x5424, 0x4405,
	0xa7db, 0xb7fa, 0x8799, 0x97b8, 0xe75f, 0xf77e, 0xc71d, 0xd73c,
	0x26d3, 0x36f2, 0x0691, 0x16b0, 0x6657, 0x7676, 0x4615, 0x5634,
	0xd94c, 0xc96d, 0xf90e, 0xe92f, 0x99c8, 0x89e9, 0xb98a, 0xa9ab,
	0x5844, 0x4865, 0x7806, 0x6827, 0x18c0, 0x08e1, 0x3882, 0x28a3,
	0xcb7d, 0xdb5c, 0xeb3f, 0xfb1e, 0x8bf9, 0x9bd8, 0xabbb, 0xbb9a,
	0x4a75, 0x5a54, 0x6a37, 0x7a16, 0x0af1, 0x1ad0, 0x2ab3, 0x3a92,
	0xfd2e, 0xed0f, 0xdd6c, 0xcd4d, 0xbdaa, 0xad8b, 0x9de8, 0x8dc9,
	0x7c26, 0x6c07, 0x5c64, 0x4c45, 0x3ca2, 0x2c83, 0x1ce0, 0x0cc1,
	0xef1f, 0xff3e, 0xcf5d, 0xdf7c, 0xaf9b, 0xbfba, 0x8fd9, 0x9ff8,
	0x6e17, 0x7e36, 0x4e55, 0x5e74, 0x2e93, 0x3eb2, 0x0ed1, 0x1ef0,
}

func crc16(key string) (crc uint16) {
	for i := 0; i < len(key); i++ {
		crc = (crc << 8) ^ crc16tab[(uint8(crc>>8)^key[i])&0x00FF]
	}
	return crc
}

func slot(key string) uint16 {
	var s, e int
	for ; s < len(key); s++ {
		if key[s] == '{' {
			break
		}
	}
	if s == len(key) {
		return crc16(key) & 16383
	}
	for e = s + 1; e < len(key); e++ {
		if key[e] == '}' {
			break
		}
	}
	if e == len(key) || e == s+1 {
		return crc16(key) & 16383
	}
	return crc16(key[s+1:e]) & 16383
}

// MGetCache is a helper that consults the client-side caches with multiple keys by grouping keys within the same slot into multiple GETs
func MGetCache(client Client, ctx context.Context, ttl time.Duration, keys []string) (ret map[string]RedisMessage, err error) {
	if len(keys) == 0 {
		return make(map[string]RedisMessage), nil
	}
	if isCacheDisabled(client) {
		return MGet(client, ctx, keys)
	}
	cmds := mgetcachecmdsp.Get(len(keys), len(keys))
	defer mgetcachecmdsp.Put(cmds)
	for i := range cmds.s {
		cmds.s[i] = CT(client.B().Get().Key(keys[i]).Cache(), ttl)
	}
	return doMultiCache(client, ctx, cmds.s, keys)
}

func isCacheDisabled(client Client) bool {
	switch c := client.(type) {
	case *singleClient:
		return c.DisableCache
	case *standalone:
		return c.primary.Load().DisableCache
	case *sentinelClient:
		return c.mOpt != nil && c.mOpt.DisableCache
	case *clusterClient:
		return c.opt != nil && c.opt.DisableCache
	}
	return false
}

// MGet is a helper that consults the redis directly with multiple keys by grouping keys within the same slot into MGET or multiple GETs
func MGet(client Client, ctx context.Context, keys []string) (ret map[string]RedisMessage, err error) {
	if len(keys) == 0 {
		return make(map[string]RedisMessage), nil
	}

	switch client.(type) {
	case *singleClient, *standalone, *sentinelClient:
		return clientMGet(client, ctx, client.B().Mget().Key(keys...).Build(), keys)
	}

	return clusterMGet(client, ctx, keys)
}

// MSet is a helper that consults the redis directly with multiple keys by grouping keys within the same slot into MSETs or multiple SETs
func MSet(client Client, ctx context.Context, kvs map[string]string) map[string]error {
	if len(kvs) == 0 {
		return make(map[string]error)
	}

	switch client.(type) {
	case *singleClient, *standalone, *sentinelClient:
		return clientMSet(client, ctx, "MSET", kvs, make(map[string]error, len(kvs)))
	}

	cmds := mgetcmdsp.Get(0, len(kvs))
	defer mgetcmdsp.Put(cmds)
	for k, v := range kvs {
		cmds.s = append(cmds.s, client.B().Set().Key(k).Value(v).Build().Pin())
	}
	return doMultiSet(client, ctx, cmds.s)
}

// MDel is a helper that consults the redis directly with multiple keys by grouping keys within the same slot into DELs
func MDel(client Client, ctx context.Context, keys []string) map[string]error {
	if len(keys) == 0 {
		return make(map[string]error)
	}

	switch client.(type) {
	case *singleClient, *standalone, *sentinelClient:
		return clientMDel(client, ctx, keys)
	}

	cmds := mgetcmdsp.Get(len(keys), len(keys))
	defer mgetcmdsp.Put(cmds)
	for i, k := range keys {
		cmds.s[i] = client.B().Del().Key(k).Build().Pin()
	}
	return doMultiSet(client, ctx, cmds.s)
}

// MSetNX is a helper that consults the redis directly with multiple keys by grouping keys within the same slot into MSETNXs or multiple SETNXs
func MSetNX(client Client, ctx context.Context, kvs map[string]string) map[string]error {
	if len(kvs) == 0 {
		return make(map[string]error)
	}

	switch client.(type) {
	case *singleClient, *standalone, *sentinelClient:
		return clientMSet(client, ctx, "MSETNX", kvs, make(map[string]error, len(kvs)))
	}

	cmds := mgetcmdsp.Get(0, len(kvs))
	defer mgetcmdsp.Put(cmds)
	for k, v := range kvs {
		cmds.s = append(cmds.s, client.B().Set().Key(k).Value(v).Nx().Build().Pin())
	}
	return doMultiSet(client, ctx, cmds.s)
}

// JsonMGetCache is a helper that consults the client-side caches with multiple keys by grouping keys within the same slot into multiple JSON.GETs
func JsonMGetCache(client Client, ctx context.Context, ttl time.Duration, keys []string, path string) (ret map[string]RedisMessage, err error) {
	if len(keys) == 0 {
		return make(map[string]RedisMessage), nil
	}
	cmds := mgetcachecmdsp.Get(len(keys), len(keys))
	defer mgetcachecmdsp.Put(cmds)
	for i := range cmds.s {
		cmds.s[i] = CT(client.B().JsonGet().Key(keys[i]).Path(path).Cache(), ttl)
	}
	return doMultiCache(client, ctx, cmds.s, keys)
}

// JsonMGet is a helper that consults redis directly with multiple keys by grouping keys within the same slot into JSON.MGETs or multiple JSON.GETs
func JsonMGet(client Client, ctx context.Context, keys []string, path string) (ret map[string]RedisMessage, err error) {
	if len(keys) == 0 {
		return make(map[string]RedisMessage), nil
	}

	switch client.(type) {
	case *singleClient, *standalone, *sentinelClient:
		return clientMGet(client, ctx, client.B().JsonMget().Key(keys...).Path(path).Build(), keys)
	}

	return clusterJsonMGet(client, ctx, keys, path)
}

// JsonMSet is a helper that consults redis directly with multiple keys by grouping keys within the same slot into JSON.MSETs or multiple JSON.SETs
func JsonMSet(client Client, ctx context.Context, kvs map[string]string, path string) map[string]error {
	if len(kvs) == 0 {
		return make(map[string]error)
	}

	switch client.(type) {
	case *singleClient, *standalone, *sentinelClient:
		return clientJSONMSet(client, ctx, kvs, path, make(map[string]error, len(kvs)))
	}

	cmds := mgetcmdsp.Get(0, len(kvs))
	defer mgetcmdsp.Put(cmds)
	for k, v := range kvs {
		cmds.s = append(cmds.s, client.B().JsonSet().Key(k).Path(path).Value(v).Build().Pin())
	}
	return doMultiSet(client, ctx, cmds.s)
}

// DecodeSliceOfJSON is a helper that struct-scans each RedisMessage into dest, which must be a slice of the pointer.
func DecodeSliceOfJSON[T any](result RedisResult, dest *[]T) error {
	values, err := result.ToArray()
	if err != nil {
		return err
	}

	ts := make([]T, len(values))
	for i, v := range values {
		var t T
		if err = v.DecodeJSON(&t); err != nil {
			if IsRedisNil(err) {
				continue
			}
			return err
		}
		ts[i] = t
	}
	*dest = ts
	return nil
}

func clientMGet(client Client, ctx context.Context, cmd Completed, keys []string) (ret map[string]RedisMessage, err error) {
	arr, err := client.Do(ctx, cmd).ToArray()
	if err != nil {
		return nil, err
	}
	return arrayToKV(make(map[string]RedisMessage, len(keys)), arr, keys), nil
}

func clientMSet(client Client, ctx context.Context, mset string, kvs map[string]string, ret map[string]error) map[string]error {
	cmd := client.B().Arbitrary(mset)
	for k, v := range kvs {
		cmd = cmd.Args(k, v)
	}
	ok, err := client.Do(ctx, cmd.Build()).AsBool()
	if err == nil && !ok {
		err = ErrMSetNXNotSet
	}
	for k := range kvs {
		ret[k] = err
	}
	return ret
}

func clientJSONMSet(client Client, ctx context.Context, kvs map[string]string, path string, ret map[string]error) map[string]error {
	cmd := intl.JsonMsetTripletValue(client.B().JsonMset())
	for k, v := range kvs {
		cmd = cmd.Key(k).Path(path).Value(v)
	}
	err := client.Do(ctx, cmd.Build()).Error()
	for k := range kvs {
		ret[k] = err
	}
	return ret
}

func clientMDel(client Client, ctx context.Context, keys []string) map[string]error {
	err := client.Do(ctx, client.B().Del().Key(keys...).Build()).Error()
	ret := make(map[string]error, len(keys))
	for _, k := range keys {
		ret[k] = err
	}
	return ret
}

func doMultiCache(cc Client, ctx context.Context, cmds []CacheableTTL, keys []string) (ret map[string]RedisMessage, err error) {
	ret = make(map[string]RedisMessage, len(keys))
	resps := cc.DoMultiCache(ctx, cmds...)
	defer resultsp.Put(&redisresults{s: resps})
	for i, resp := range resps {
		if err := resp.NonRedisError(); err != nil {
			return nil, err
		}
		ret[keys[i]] = resp.val
	}
	return ret, nil
}

func doMultiGet(cc Client, ctx context.Context, cmds []Completed, keys []string) (ret map[string]RedisMessage, err error) {
	ret = make(map[string]RedisMessage, len(keys))
	resps := cc.DoMulti(ctx, cmds...)
	defer resultsp.Put(&redisresults{s: resps})
	for i, resp := range resps {
		if err := resp.NonRedisError(); err != nil {
			return nil, err
		}
		ret[keys[i]] = resp.val
	}
	return ret, nil
}

func doMultiSet(cc Client, ctx context.Context, cmds []Completed) (ret map[string]error) {
	ret = make(map[string]error, len(cmds))
	resps := cc.DoMulti(ctx, cmds...)
	for i, resp := range resps {
		if ret[cmds[i].Commands()[1]] = resp.Error(); resp.NonRedisError() == nil {
			intl.PutCompletedForce(cmds[i])
		}
	}
	resultsp.Put(&redisresults{s: resps})
	return ret
}

func arrayToKV(m map[string]RedisMessage, arr []RedisMessage, keys []string) map[string]RedisMessage {
	for i, resp := range arr {
		m[keys[i]] = resp
	}
	return m
}

func clusterMGet(client Client, ctx context.Context, keys []string) (ret map[string]RedisMessage, err error) {
	ret = make(map[string]RedisMessage, len(keys))
	slotGroups := make(map[uint16][]string)
	for _, key := range keys {
		ks := slot(key)
		slotGroups[ks] = append(slotGroups[ks], key)
	}
	cmds := mgetcmdsp.Get(0, len(slotGroups))
	defer mgetcmdsp.Put(cmds)
	var cmdKeys [][]string
	for _, group := range slotGroups {
		cmd := client.B().Mget().Key(group...).Build().Pin()
		cmds.s = append(cmds.s, cmd)
		cmdKeys = append(cmdKeys, group)
	}
	resps := client.DoMulti(ctx, cmds.s...)
	defer resultsp.Put(&redisresults{s: resps})
	for i, resp := range resps {
		arr, err := resp.ToArray()
		if err != nil {
			return nil, err
		}
		ret = arrayToKV(ret, arr, cmdKeys[i])
	}
	for i := range cmds.s {
		intl.PutCompletedForce(cmds.s[i])
	}
	return ret, nil
}

func clusterJsonMGet(client Client, ctx context.Context, keys []string, path string) (ret map[string]RedisMessage, err error) {
	ret = make(map[string]RedisMessage, len(keys))
	slotGroups := make(map[uint16][]string)
	for _, key := range keys {
		ks := slot(key)
		slotGroups[ks] = append(slotGroups[ks], key)
	}
	if len(slotGroups) == 0 {
		return ret, nil
	}
	cmds := mgetcmdsp.Get(0, len(slotGroups))
	defer mgetcmdsp.Put(cmds)
	var cmdKeys [][]string
	for _, group := range slotGroups {
		cmd := client.B().JsonMget().Key(group...).Path(path).Build().Pin()
		cmds.s = append(cmds.s, cmd)
		cmdKeys = append(cmdKeys, group)
	}
	resps := client.DoMulti(ctx, cmds.s...)
	defer resultsp.Put(&redisresults{s: resps})
	for i, resp := range resps {
		arr, err := resp.ToArray()
		if err != nil {
			return nil, err
		}
		ret = arrayToKV(ret, arr, cmdKeys[i])
	}
	for i := range cmds.s {
		intl.PutCompletedForce(cmds.s[i])
	}
	return ret, nil
}

// ErrMSetNXNotSet is used in the MSetNX helper when the underlying MSETNX response is 0.
// Ref: https://redis.io/commands/msetnx/
var ErrMSetNXNotSet = errors.New("MSETNX: no key was set")

type Scanner struct {
	next func(cursor uint64) (ScanEntry, error)
	err  error
}

func NewScanner(next func(cursor uint64) (ScanEntry, error)) *Scanner {
	return &Scanner{next: next}
}

func (s *Scanner) scan() iter.Seq[[]string] {
	return func(yield func([]string) bool) {
		var e ScanEntry
		for e, s.err = s.next(0); s.err == nil && yield(e.Elements) && e.Cursor != 0; {
			e, s.err = s.next(e.Cursor)
		}
	}
}

func (s *Scanner) Iter() iter.Seq[string] {
	return func(yield func(string) bool) {
		for vs := range s.scan() {
			for _, v := range vs {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func (s *Scanner) Iter2() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for vs := range s.scan() {
			for i := 0; i+1 < len(vs); i += 2 {
				if !yield(vs[i], vs[i+1]) {
					return
				}
			}
		}
	}
}

func (s *Scanner) Err() error {
	return s.err
}
