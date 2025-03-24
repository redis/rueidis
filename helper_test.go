package rueidis

import (
	"context"
	"reflect"
	"strconv"
	"testing"
)

//gocyclo:ignore
func TestMGetCache(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		disabledCacheClient, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}, DisableCache: true},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate DisabledCache MGetCache", func(t *testing.T) {
			m.DoFn = func(cmd Completed) RedisResult {
				if !reflect.DeepEqual(cmd.Commands(), []string{"MGET", "1", "2"}) {
					t.Fatalf("unexpected command %v", cmd)
				}
				return newResult(slicemsg('*', []RedisMessage{strmsg('+', "1"), strmsg('+', "2")}), nil)
			}
			if v, err := MGetCache(disabledCacheClient, context.Background(), 100, []string{"1", "2"}); err != nil || v == nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate DoCache", func(t *testing.T) {
			m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
				if reflect.DeepEqual(multi[0].Cmd.Commands(), []string{"GET", "1"}) && multi[0].TTL == 100 &&
					reflect.DeepEqual(multi[1].Cmd.Commands(), []string{"GET", "2"}) && multi[1].TTL == 100 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "1"), nil),
						newResult(strmsg('+', "2"), nil),
					}}
				}
				t.Fatalf("unexpected command %v", multi)
				return nil
			}
			if v, err := MGetCache(client, context.Background(), 100, []string{"1", "2"}); err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			} else if v1, v2 := v["1"], v["2"]; v1.string() != "1" || v2.string() != "2" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate DoCache Empty", func(t *testing.T) {
			if v, err := MGetCache(client, context.Background(), 100, []string{}); err != nil || v == nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate DoCache Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
				return &redisresults{s: []RedisResult{newResult(RedisMessage{}, context.Canceled), newResult(RedisMessage{}, context.Canceled)}}
			}
			if v, err := MGetCache(client, ctx, 100, []string{"1", "2"}); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
	t.Run("cluster client", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return slotsResp
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		disabledCacheClient, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}, DisableCache: true},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate DisabledCache DoCache", func(t *testing.T) {
			keys := make([]string, 100)
			for i := range keys {
				keys[i] = strconv.Itoa(i)
			}
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				result := make([]RedisResult, len(cmd))
				for i, key := range keys {
					if !reflect.DeepEqual(cmd[i].Commands(), []string{"GET", key}) {
						t.Fatalf("unexpected command %v", cmd)
						return nil
					}
					result[i] = newResult(strmsg('+', key), nil)
				}
				return &redisresults{s: result}
			}
			v, err := MGetCache(disabledCacheClient, context.Background(), 100, keys)
			if err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			for _, key := range keys {
				if vKey, ok := v[key]; !ok || vKey.string() != key {
					t.Fatalf("unexpected response %v", v)
				}
			}
		})

		t.Run("Delegate DoCache", func(t *testing.T) {
			keys := make([]string, 100)
			for i := range keys {
				keys[i] = strconv.Itoa(i)
			}
			m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
				result := make([]RedisResult, len(multi))
				for i, key := range keys {
					if !reflect.DeepEqual(multi[i].Cmd.Commands(), []string{"GET", key}) || multi[i].TTL != 100 {
						t.Fatalf("unexpected command %v", multi)
						return nil
					}
					result[i] = newResult(strmsg('+', key), nil)
				}
				return &redisresults{s: result}
			}
			v, err := MGetCache(client, context.Background(), 100, keys)
			if err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			for _, key := range keys {
				if vKey, ok := v[key]; !ok || vKey.string() != key {
					t.Fatalf("unexpected response %v", v)
				}
			}
		})
		t.Run("Delegate DoCache Empty", func(t *testing.T) {
			if v, err := MGetCache(client, context.Background(), 100, []string{}); err != nil || v == nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate DoCache Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
				result := make([]RedisResult, len(multi))
				for i := range result {
					result[i] = newErrResult(context.Canceled)
				}
				return &redisresults{s: result}
			}
			if v, err := MGetCache(client, ctx, 100, []string{"1", "2"}); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
}

//gocyclo:ignore
func TestMGet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			m.DoFn = func(cmd Completed) RedisResult {
				if !reflect.DeepEqual(cmd.Commands(), []string{"MGET", "1", "2"}) {
					t.Fatalf("unexpected command %v", cmd)
				}
				return newResult(slicemsg('*', []RedisMessage{strmsg('+', "1"), strmsg('+', "2")}), nil)
			}
			if v, err := MGet(client, context.Background(), []string{"1", "2"}); err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			} else if v1, v2 := v["1"], v["2"]; v1.string() != "1" || v2.string() != "2" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if v, err := MGet(client, context.Background(), []string{}); err != nil || v == nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoFn = func(cmd Completed) RedisResult {
				return newResult(RedisMessage{}, context.Canceled)
			}
			if v, err := MGet(client, ctx, []string{"1", "2"}); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
	t.Run("cluster client", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return slotsResp
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			keys := make([]string, 100)
			for i := range keys {
				keys[i] = strconv.Itoa(i)
			}
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				result := make([]RedisResult, len(cmd))
				for i, key := range keys {
					if !reflect.DeepEqual(cmd[i].Commands(), []string{"GET", key}) {
						t.Fatalf("unexpected command %v", cmd)
						return nil
					}
					result[i] = newResult(strmsg('+', key), nil)
				}
				return &redisresults{s: result}
			}
			v, err := MGet(client, context.Background(), keys)
			if err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			for _, key := range keys {
				if vKey, ok := v[key]; !ok || vKey.string() != key {
					t.Fatalf("unexpected response %v", v)
				}
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if v, err := MGet(client, context.Background(), []string{}); err != nil || v == nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{newErrResult(context.Canceled), newErrResult(context.Canceled)}}
			}
			if v, err := MGet(client, ctx, []string{"1", "2"}); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
}

//gocyclo:ignore
func TestMDel(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			m.DoFn = func(cmd Completed) RedisResult {
				if !reflect.DeepEqual(cmd.Commands(), []string{"DEL", "1", "2"}) {
					t.Fatalf("unexpected command %v", cmd)
				}
				return newResult(RedisMessage{typ: ':', intlen: 2}, nil)
			}
			if v := MDel(client, context.Background(), []string{"1", "2"}); v["1"] != nil || v["2"] != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if v := MDel(client, context.Background(), []string{}); len(v) != 0 {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoFn = func(cmd Completed) RedisResult {
				return newResult(RedisMessage{}, context.Canceled)
			}
			if v := MDel(client, ctx, []string{"1", "2"}); v["1"] != context.Canceled || v["2"] != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
	t.Run("cluster client", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return slotsResp
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			keys := make([]string, 100)
			for i := range keys {
				keys[i] = strconv.Itoa(i)
			}
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				result := make([]RedisResult, len(cmd))
				for i, key := range keys {
					if !reflect.DeepEqual(cmd[i].Commands(), []string{"DEL", key}) {
						t.Fatalf("unexpected command %v", cmd)
						return nil
					}
					result[i] = newResult(RedisMessage{typ: ':', intlen: 1}, nil)
				}
				return &redisresults{s: result}
			}
			v := MDel(client, context.Background(), keys)
			for _, key := range keys {
				if v[key] != nil {
					t.Fatalf("unexpected response %v", v)
				}
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if v := MDel(client, context.Background(), []string{}); len(v) != 0 {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{newErrResult(context.Canceled), newErrResult(context.Canceled)}}
			}
			if v := MDel(client, ctx, []string{"1", "2"}); v["1"] != context.Canceled || v["2"] != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
}

func TestMSet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			m.DoFn = func(cmd Completed) RedisResult {
				if !reflect.DeepEqual(cmd.Commands(), []string{"MSET", "1", "1", "2", "2"}) &&
					!reflect.DeepEqual(cmd.Commands(), []string{"MSET", "2", "2", "1", "1"}) {
					t.Fatalf("unexpected command %v", cmd)
				}
				return newResult(strmsg('+', "OK"), nil)
			}
			if err := MSet(client, context.Background(), map[string]string{"1": "1", "2": "2"}); err["1"] != nil || err["2"] != nil {
				t.Fatalf("unexpected response %v", err)
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if err := MSet(client, context.Background(), map[string]string{}); len(err) != 0 {
				t.Fatalf("unexpected response %v", err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoFn = func(cmd Completed) RedisResult {
				return newResult(RedisMessage{}, context.Canceled)
			}
			if err := MSet(client, ctx, map[string]string{"1": "1", "2": "2"}); err["1"] != context.Canceled || err["2"] != context.Canceled {
				t.Fatalf("unexpected response %v", err)
			}
		})
	})
	t.Run("cluster client", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return slotsResp
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			keys := make(map[string]string, 100)
			for i := 0; i < 100; i++ {
				keys[strconv.Itoa(i)] = strconv.Itoa(i)
			}
			cpy := make(map[string]struct{}, len(keys))
			for k := range keys {
				cpy[k] = struct{}{}
			}
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				result := make([]RedisResult, len(cmd))
				for i, c := range cmd {
					delete(cpy, c.Commands()[1])
					if c.Commands()[0] != "SET" || keys[c.Commands()[1]] != c.Commands()[2] {
						t.Fatalf("unexpected command %v", cmd)
						return nil
					}
					result[i] = newResult(strmsg('+', "OK"), nil)
				}
				if len(cpy) != 0 {
					t.Fatalf("unexpected command %v", cmd)
					return nil
				}
				return &redisresults{s: result}
			}
			err := MSet(client, context.Background(), keys)
			for key := range keys {
				if err[key] != nil {
					t.Fatalf("unexpected response %v", err)
				}
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if err := MSet(client, context.Background(), map[string]string{}); len(err) != 0 {
				t.Fatalf("unexpected response %v", err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{newErrResult(context.Canceled), newErrResult(context.Canceled)}}
			}
			if err := MSet(client, ctx, map[string]string{"1": "1", "2": "2"}); err["1"] != context.Canceled || err["2"] != context.Canceled {
				t.Fatalf("unexpected response %v", err)
			}
		})
	})
}

func TestMSetNX(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			m.DoFn = func(cmd Completed) RedisResult {
				if !reflect.DeepEqual(cmd.Commands(), []string{"MSETNX", "1", "1", "2", "2"}) &&
					!reflect.DeepEqual(cmd.Commands(), []string{"MSETNX", "2", "2", "1", "1"}) {
					t.Fatalf("unexpected command %v", cmd)
				}
				return newResult(strmsg('+', "OK"), nil)
			}
			if err := MSetNX(client, context.Background(), map[string]string{"1": "1", "2": "2"}); err["1"] != nil || err["2"] != nil {
				t.Fatalf("unexpected response %v", err)
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if err := MSetNX(client, context.Background(), map[string]string{}); len(err) != 0 {
				t.Fatalf("unexpected response %v", err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoFn = func(cmd Completed) RedisResult {
				return newResult(RedisMessage{}, context.Canceled)
			}
			if err := MSetNX(client, ctx, map[string]string{"1": "1", "2": "2"}); err["1"] != context.Canceled || err["2"] != context.Canceled {
				t.Fatalf("unexpected response %v", err)
			}
		})
	})
	t.Run("cluster client", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return slotsResp
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			keys := make(map[string]string, 100)
			for i := 0; i < 100; i++ {
				keys[strconv.Itoa(i)] = strconv.Itoa(i)
			}
			cpy := make(map[string]struct{}, len(keys))
			for k := range keys {
				cpy[k] = struct{}{}
			}
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				result := make([]RedisResult, len(cmd))
				for i, c := range cmd {
					delete(cpy, c.Commands()[1])
					if c.Commands()[0] != "SET" || c.Commands()[3] != "NX" || keys[c.Commands()[1]] != c.Commands()[2] {
						t.Fatalf("unexpected command %v", cmd)
						return nil
					}
					result[i] = newResult(strmsg('+', "OK"), nil)
				}
				if len(cpy) != 0 {
					t.Fatalf("unexpected command %v", cmd)
					return nil
				}
				return &redisresults{s: result}
			}
			err := MSetNX(client, context.Background(), keys)
			for key := range keys {
				if err[key] != nil {
					t.Fatalf("unexpected response %v", err)
				}
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if err := MSetNX(client, context.Background(), map[string]string{}); len(err) != 0 {
				t.Fatalf("unexpected response %v", err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{newErrResult(context.Canceled), newErrResult(context.Canceled)}}
			}
			if err := MSetNX(client, ctx, map[string]string{"1": "1", "2": "2"}); err["1"] != context.Canceled || err["2"] != context.Canceled {
				t.Fatalf("unexpected response %v", err)
			}
		})
	})
}

func TestMSetNXNotSet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do Not Set", func(t *testing.T) {
			m.DoFn = func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: ':', intlen: 0}, nil)
			}
			if err := MSetNX(client, context.Background(), map[string]string{"1": "1", "2": "2"}); err["1"] != ErrMSetNXNotSet || err["2"] != ErrMSetNXNotSet {
				t.Fatalf("unexpected response %v", err)
			}
		})
	})
}

//gocyclo:ignore
func TestJsonMGetCache(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate DoCache", func(t *testing.T) {
			m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
				if reflect.DeepEqual(multi[0].Cmd.Commands(), []string{"JSON.GET", "1", "$"}) && multi[0].TTL == 100 &&
					reflect.DeepEqual(multi[1].Cmd.Commands(), []string{"JSON.GET", "2", "$"}) && multi[1].TTL == 100 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "1"), nil),
						newResult(strmsg('+', "2"), nil),
					}}
				}
				t.Fatalf("unexpected command %v", multi)
				return nil
			}
			if v, err := JsonMGetCache(client, context.Background(), 100, []string{"1", "2"}, "$"); err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			} else if v1, v2 := v["1"], v["2"]; v1.string() != "1" || v2.string() != "2" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate DoCache Empty", func(t *testing.T) {
			if v, err := JsonMGetCache(client, context.Background(), 100, []string{}, "$"); err != nil || v == nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate DoCache Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
				return &redisresults{s: []RedisResult{newResult(RedisMessage{}, context.Canceled), newResult(RedisMessage{}, context.Canceled)}}
			}
			if v, err := JsonMGetCache(client, ctx, 100, []string{"1", "2"}, "$"); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
	t.Run("cluster client", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return slotsResp
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate DoCache", func(t *testing.T) {
			keys := make([]string, 100)
			for i := range keys {
				keys[i] = strconv.Itoa(i)
			}
			m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
				result := make([]RedisResult, len(multi))
				for i, key := range keys {
					if !reflect.DeepEqual(multi[i].Cmd.Commands(), []string{"JSON.GET", key, "$"}) || multi[i].TTL != 100 {
						t.Fatalf("unexpected command %v", multi)
						return nil
					}
					result[i] = newResult(strmsg('+', key), nil)
				}
				return &redisresults{s: result}
			}
			v, err := JsonMGetCache(client, context.Background(), 100, keys, "$")
			if err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			for _, key := range keys {
				if vKey, ok := v[key]; !ok || vKey.string() != key {
					t.Fatalf("unexpected response %v", v)
				}
			}
		})
		t.Run("Delegate DoCache Empty", func(t *testing.T) {
			if v, err := JsonMGetCache(client, context.Background(), 100, []string{}, "$"); err != nil || v == nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate DoCache Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
				result := make([]RedisResult, len(multi))
				for i := range result {
					result[i] = newErrResult(context.Canceled)
				}
				return &redisresults{s: result}
			}
			if v, err := JsonMGetCache(client, ctx, 100, []string{"1", "2"}, "$"); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
}

//gocyclo:ignore
func TestJsonMGet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			m.DoFn = func(cmd Completed) RedisResult {
				if !reflect.DeepEqual(cmd.Commands(), []string{"JSON.MGET", "1", "2", "$"}) {
					t.Fatalf("unexpected command %v", cmd)
				}
				return newResult(slicemsg('*', []RedisMessage{strmsg('+', "1"), strmsg('+', "2")}), nil)
			}
			if v, err := JsonMGet(client, context.Background(), []string{"1", "2"}, "$"); err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			} else if v1, v2 := v["1"], v["2"]; v1.string() != "1" || v2.string() != "2" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if v, err := JsonMGet(client, context.Background(), []string{}, "$"); err != nil || v == nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoFn = func(cmd Completed) RedisResult {
				return newResult(RedisMessage{}, context.Canceled)
			}
			if v, err := JsonMGet(client, ctx, []string{"1", "2"}, "$"); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
	t.Run("cluster client", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return slotsResp
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			keys := make([]string, 100)
			for i := range keys {
				keys[i] = strconv.Itoa(i)
			}
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				result := make([]RedisResult, len(cmd))
				for i, key := range keys {
					if !reflect.DeepEqual(cmd[i].Commands(), []string{"JSON.GET", key, "$"}) {
						t.Fatalf("unexpected command %v", cmd)
						return nil
					}
					result[i] = newResult(strmsg('+', key), nil)
				}
				return &redisresults{s: result}
			}
			v, err := JsonMGet(client, context.Background(), keys, "$")
			if err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			for _, key := range keys {
				if vKey, ok := v[key]; !ok || vKey.string() != key {
					t.Fatalf("unexpected response %v", v)
				}
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if v, err := JsonMGet(client, context.Background(), []string{}, "$"); err != nil || v == nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{newErrResult(context.Canceled), newErrResult(context.Canceled)}}
			}
			if v, err := JsonMGet(client, ctx, []string{"1", "2"}, "$"); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
}

func TestJsonMSet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			m.DoFn = func(cmd Completed) RedisResult {
				if !reflect.DeepEqual(cmd.Commands(), []string{"JSON.MSET", "1", "$", "1", "2", "$", "2"}) &&
					!reflect.DeepEqual(cmd.Commands(), []string{"JSON.MSET", "2", "$", "2", "1", "$", "1"}) {
					t.Fatalf("unexpected command %v", cmd)
				}
				return newResult(strmsg('+', "OK"), nil)
			}
			if err := JsonMSet(client, context.Background(), map[string]string{"1": "1", "2": "2"}, "$"); err["1"] != nil || err["2"] != nil {
				t.Fatalf("unexpected response %v", err)
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if err := JsonMSet(client, context.Background(), map[string]string{}, "$"); len(err) != 0 {
				t.Fatalf("unexpected response %v", err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoFn = func(cmd Completed) RedisResult {
				return newResult(RedisMessage{}, context.Canceled)
			}
			if err := JsonMSet(client, ctx, map[string]string{"1": "1", "2": "2"}, "$"); err["1"] != context.Canceled || err["2"] != context.Canceled {
				t.Fatalf("unexpected response %v", err)
			}
		})
	})
	t.Run("cluster client", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return slotsResp
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate Do", func(t *testing.T) {
			keys := make(map[string]string, 100)
			for i := 0; i < 100; i++ {
				keys[strconv.Itoa(i)] = strconv.Itoa(i)
			}
			cpy := make(map[string]struct{}, len(keys))
			for k := range keys {
				cpy[k] = struct{}{}
			}
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				result := make([]RedisResult, len(cmd))
				for i, c := range cmd {
					delete(cpy, c.Commands()[1])
					if c.Commands()[0] != "JSON.SET" || keys[c.Commands()[1]] != c.Commands()[3] || c.Commands()[2] != "$" {
						t.Fatalf("unexpected command %v", cmd)
						return nil
					}
					result[i] = newResult(strmsg('+', "OK"), nil)
				}
				if len(cpy) != 0 {
					t.Fatalf("unexpected command %v", cmd)
					return nil
				}
				return &redisresults{s: result}
			}
			err := JsonMSet(client, context.Background(), keys, "$")
			for key := range keys {
				if err[key] != nil {
					t.Fatalf("unexpected response %v", err)
				}
			}
		})
		t.Run("Delegate Do Empty", func(t *testing.T) {
			if err := JsonMSet(client, context.Background(), map[string]string{}, "$"); len(err) != 0 {
				t.Fatalf("unexpected response %v", err)
			}
		})
		t.Run("Delegate Do Err", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{newErrResult(context.Canceled), newErrResult(context.Canceled)}}
			}
			if err := JsonMSet(client, ctx, map[string]string{"1": "1", "2": "2"}, "$"); err["1"] != context.Canceled || err["2"] != context.Canceled {
				t.Fatalf("unexpected response %v", err)
			}
		})
	})
}

func TestDecodeSliceOfJSON(t *testing.T) {
	type Inner struct {
		Field string
	}
	type T struct {
		Name   string
		Inners []*Inner
		ID     int
	}
	values := []RedisMessage{
		strmsg('+', `{"ID":1, "Name": "n1", "Inners": [{"Field": "f1"}]}`),
		strmsg('+', `{"ID":2, "Name": "n2", "Inners": [{"Field": "f2"}]}`),
	}
	result := RedisResult{val: slicemsg('*', values)}

	t.Run("Scan []*T", func(t *testing.T) {
		got := make([]*T, 0)
		want := []*T{
			{ID: 1, Name: "n1", Inners: []*Inner{{Field: "f1"}}},
			{ID: 2, Name: "n2", Inners: []*Inner{{Field: "f2"}}},
		}
		if err := DecodeSliceOfJSON(result, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("DecodeSliceOfJSON not get value as expected %v", got)
		}
	})

	t.Run("Scan []T", func(t *testing.T) {
		got := make([]T, 0)
		want := []T{
			{ID: 1, Name: "n1", Inners: []*Inner{{Field: "f1"}}},
			{ID: 2, Name: "n2", Inners: []*Inner{{Field: "f2"}}},
		}
		if err := DecodeSliceOfJSON(result, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("DecodeSliceOfJSON not get value as expected %v", got)
		}
	})

	t.Run("Scan []*T: has nil error message", func(t *testing.T) {
		hasNilValues := []RedisMessage{
			strmsg('+', `{"ID":1, "Name": "n1", "Inners": [{"Field": "f1"}]}`),
			{typ: '_'},
			strmsg('+', `{"ID":2, "Name": "n2", "Inners": [{"Field": "f2"}]}`),
		}
		hasNilResult := RedisResult{val: slicemsg('*', hasNilValues)}

		got := make([]*T, 0)
		want := []*T{
			{ID: 1, Name: "n1", Inners: []*Inner{{Field: "f1"}}},
			nil,
			{ID: 2, Name: "n2", Inners: []*Inner{{Field: "f2"}}},
		}
		if err := DecodeSliceOfJSON(hasNilResult, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("DecodeSliceOfJSON not get value as expected %v", got)
		}
	})

	t.Run("Scan []T: has nil error message", func(t *testing.T) {
		hasNilValues := []RedisMessage{
			strmsg('+', `{"ID":1, "Name": "n1", "Inners": [{"Field": "f1"}]}`),
			{typ: '_'},
			strmsg('+', `{"ID":2, "Name": "n2", "Inners": [{"Field": "f2"}]}`),
		}
		hasNilResult := RedisResult{val: slicemsg('*', hasNilValues)}

		got := make([]T, 0)
		want := []T{
			{ID: 1, Name: "n1", Inners: []*Inner{{Field: "f1"}}},
			{ID: 0, Name: "", Inners: nil},
			{ID: 2, Name: "n2", Inners: []*Inner{{Field: "f2"}}},
		}
		if err := DecodeSliceOfJSON(hasNilResult, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("DecodeSliceOfJSON not get value as expected %v", got)
		}
	})

	t.Run("error result", func(t *testing.T) {
		if err := DecodeSliceOfJSON(RedisResult{val: RedisMessage{typ: '-'}}, &[]*T{}); err == nil {
			t.Fatal("DecodeSliceOfJSON not failed as expected")
		}
	})

	t.Run("has non-nil error message in result", func(t *testing.T) {
		hasErrValues := []RedisMessage{
			strmsg('+', `{"ID":1, "Name": "n1", "Inners": [{"Field": "f1"}]}`),
			strmsg('-', `invalid`),
		}
		hasErrResult := RedisResult{val: slicemsg('*', hasErrValues)}

		got := make([]*T, 0)
		if err := DecodeSliceOfJSON(hasErrResult, &got); err == nil {
			t.Fatal("DecodeSliceOfJSON not failed as expected")
		}
	})
}
