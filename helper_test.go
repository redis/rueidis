package rueidis

import (
	"context"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

//gocyclo:ignore
func TestMGetCache(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
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
	t.Run("standalone client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newStandaloneClient(
			&ClientOption{
				InitAddress: []string{""},
				Standalone: StandaloneOption{
					ReplicaAddress: []string{""},
				},
				SendToReplicas: func(cmd Completed) bool {
					return cmd.IsReadOnly()
				},
			},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		disabledCacheClient, err := newStandaloneClient(
			&ClientOption{
				InitAddress: []string{""},
				Standalone: StandaloneOption{
					ReplicaAddress: []string{""},
				},
				SendToReplicas: func(cmd Completed) bool {
					return cmd.IsReadOnly()
				},
				DisableCache: true,
			},
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
			keys := []string{"{slot1}a", "{slot1}b", "{slot2}a", "{slot2}b"}
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				result := make([]RedisResult, len(cmd))
				for i, c := range cmd {
					// Each command should be MGET with keys from the same slot
					commands := c.Commands()
					if commands[0] != "MGET" {
						t.Fatalf("expected MGET command, got %v", commands)
						return nil
					}
					// Build response array with values matching the keys
					values := make([]RedisMessage, len(commands)-1)
					for j := 1; j < len(commands); j++ {
						values[j-1] = strmsg('+', commands[j])
					}
					result[i] = newResult(slicemsg('*', values), nil)
				}
				return &redisresults{s: result}
			}
			v, err := MGetCache(disabledCacheClient, context.Background(), 100, keys)
			if err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			for _, key := range keys {
				if vKey, ok := v[key]; !ok || vKey.string() != key {
					t.Fatalf("unexpected response for key %s: %v", key, v)
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
	defer ShouldNotLeak(SetupLeakDetection())
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
	t.Run("standalone client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newStandaloneClient(
			&ClientOption{
				InitAddress: []string{""},
				Standalone: StandaloneOption{
					ReplicaAddress: []string{""},
				},
				SendToReplicas: func(cmd Completed) bool {
					return cmd.IsReadOnly()
				},
			},
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
			keys := []string{"{slot1}a", "{slot1}b", "{slot2}a", "{slot2}b"}
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				result := make([]RedisResult, len(cmd))
				for i, c := range cmd {
					// Each command should be MGET with keys from the same slot
					commands := c.Commands()
					if commands[0] != "MGET" {
						t.Fatalf("expected MGET command, got %v", commands)
						return nil
					}
					// Build response array with values matching the keys
					values := make([]RedisMessage, len(commands)-1)
					for j := 1; j < len(commands); j++ {
						values[j-1] = strmsg('+', commands[j])
					}
					result[i] = newResult(slicemsg('*', values), nil)
				}
				return &redisresults{s: result}
			}
			v, err := MGet(client, context.Background(), keys)
			if err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			for _, key := range keys {
				if vKey, ok := v[key]; !ok || vKey.string() != key {
					t.Fatalf("unexpected response for key %s: %v", key, v)
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
	defer ShouldNotLeak(SetupLeakDetection())
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
	t.Run("standalone client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newStandaloneClient(
			&ClientOption{
				InitAddress: []string{""},
				Standalone: StandaloneOption{
					ReplicaAddress: []string{""},
				},
				SendToReplicas: func(cmd Completed) bool {
					return cmd.IsReadOnly()
				},
			},
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
	defer ShouldNotLeak(SetupLeakDetection())
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
	t.Run("standalone client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newStandaloneClient(
			&ClientOption{
				InitAddress: []string{""},
				Standalone: StandaloneOption{
					ReplicaAddress: []string{""},
				},
				SendToReplicas: func(cmd Completed) bool {
					return cmd.IsReadOnly()
				},
			},
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
			for i := range 100 {
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
	defer ShouldNotLeak(SetupLeakDetection())
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
	t.Run("standalone client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newStandaloneClient(
			&ClientOption{
				InitAddress: []string{""},
				Standalone: StandaloneOption{
					ReplicaAddress: []string{""},
				},
				SendToReplicas: func(cmd Completed) bool {
					return cmd.IsReadOnly()
				},
			},
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
			for i := range 100 {
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
	defer ShouldNotLeak(SetupLeakDetection())
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
	t.Run("standalone client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newStandaloneClient(
			&ClientOption{
				InitAddress: []string{""},
				Standalone: StandaloneOption{
					ReplicaAddress: []string{""},
				},
				SendToReplicas: func(cmd Completed) bool {
					return cmd.IsReadOnly()
				},
			},
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
	defer ShouldNotLeak(SetupLeakDetection())
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
	t.Run("standalone client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newStandaloneClient(
			&ClientOption{
				InitAddress: []string{""},
				Standalone: StandaloneOption{
					ReplicaAddress: []string{""},
				},
				SendToReplicas: func(cmd Completed) bool {
					return cmd.IsReadOnly()
				},
			},
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
	defer ShouldNotLeak(SetupLeakDetection())
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
	t.Run("standalone client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newStandaloneClient(
			&ClientOption{
				InitAddress: []string{""},
				Standalone: StandaloneOption{
					ReplicaAddress: []string{""},
				},
				SendToReplicas: func(cmd Completed) bool {
					return cmd.IsReadOnly()
				},
			},
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
			keys := []string{"{slot1}a", "{slot1}b", "{slot2}a", "{slot2}b"}
			m.DoMultiFn = func(cmd ...Completed) *redisresults {
				result := make([]RedisResult, len(cmd))
				for i, c := range cmd {
					// Each command should be JSON.MGET with keys from the same slot and path at the end
					commands := c.Commands()
					if commands[0] != "JSON.MGET" {
						t.Fatalf("expected JSON.MGET command, got %v", commands)
						return nil
					}
					if commands[len(commands)-1] != "$" {
						t.Fatalf("expected $ as last parameter, got %v", commands)
						return nil
					}
					// Build response array with values matching the keys (exclude the path)
					values := make([]RedisMessage, len(commands)-2)
					for j := 1; j < len(commands)-1; j++ {
						values[j-1] = strmsg('+', commands[j])
					}
					result[i] = newResult(slicemsg('*', values), nil)
				}
				return &redisresults{s: result}
			}
			v, err := JsonMGet(client, context.Background(), keys, "$")
			if err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			for _, key := range keys {
				if vKey, ok := v[key]; !ok || vKey.string() != key {
					t.Fatalf("unexpected response for key %s: %v", key, v)
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
	defer ShouldNotLeak(SetupLeakDetection())
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
	t.Run("standalone client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newStandaloneClient(
			&ClientOption{
				InitAddress: []string{""},
				Standalone: StandaloneOption{
					ReplicaAddress: []string{""},
				},
				SendToReplicas: func(cmd Completed) bool {
					return cmd.IsReadOnly()
				},
			},
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
			for i := range 100 {
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

func TestScannerIter(t *testing.T) {
	tests := []struct {
		name     string
		entries  []ScanEntry
		err      error
		expected []string
		wantErr  bool
	}{
		{
			name: "single page",
			entries: []ScanEntry{
				{Elements: []string{"key1", "key2", "key3"}, Cursor: 0},
			},
			expected: []string{"key1", "key2", "key3"},
		},
		{
			name: "multiple pages",
			entries: []ScanEntry{
				{Elements: []string{"key1", "key2"}, Cursor: 10},
				{Elements: []string{"key3", "key4"}, Cursor: 0},
			},
			expected: []string{"key1", "key2", "key3", "key4"},
		},
		{
			name: "empty result",
			entries: []ScanEntry{
				{Elements: []string{}, Cursor: 0},
			},
			expected: []string{},
		},
		{
			name:    "error case",
			err:     errors.New("scan error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			scanner := NewScanner(func(cursor uint64) (ScanEntry, error) {
				if tt.err != nil {
					return ScanEntry{}, tt.err
				}
				if callCount >= len(tt.entries) {
					return ScanEntry{}, errors.New("unexpected call")
				}
				entry := tt.entries[callCount]
				callCount++
				return entry, nil
			})

			var result []string
			for element := range scanner.Iter() {
				result = append(result, element)
			}

			if tt.wantErr {
				if scanner.Err() == nil {
					t.Error("expected error but got none")
				}
			} else {
				if scanner.Err() != nil {
					t.Errorf("unexpected error: %v", scanner.Err())
				}
				if (len(result) != 0 || len(tt.expected) != 0) && !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("got %v, want %v", result, tt.expected)
				}
			}
		})
	}

	t.Run("early exit", func(t *testing.T) {
		callCount := 0
		entries := []ScanEntry{
			{Elements: []string{"key1"}, Cursor: 10},
		}
		scanner := NewScanner(func(cursor uint64) (ScanEntry, error) {
			if callCount >= len(entries) {
				return ScanEntry{}, errors.New("unexpected call")
			}
			entry := entries[callCount]
			callCount++
			return entry, nil
		})
		for range scanner.Iter() {
			break
		}
		if scanner.Err() != nil {
			t.Errorf("unexpected error: %v", scanner.Err())
		}
	})
}

func TestScannerIter2(t *testing.T) {
	tests := []struct {
		name         string
		entries      []ScanEntry
		err          error
		expectedKeys []string
		expectedVals []string
		wantErr      bool
	}{
		{
			name: "single page pairs",
			entries: []ScanEntry{
				{Elements: []string{"field1", "value1", "field2", "value2"}, Cursor: 0},
			},
			expectedKeys: []string{"field1", "field2"},
			expectedVals: []string{"value1", "value2"},
		},
		{
			name: "multiple pages pairs",
			entries: []ScanEntry{
				{Elements: []string{"field1", "value1"}, Cursor: 10},
				{Elements: []string{"field2", "value2", "field3", "value3"}, Cursor: 0},
			},
			expectedKeys: []string{"field1", "field2", "field3"},
			expectedVals: []string{"value1", "value2", "value3"},
		},
		{
			name: "odd number of elements",
			entries: []ScanEntry{
				{Elements: []string{"field1", "value1", "field2"}, Cursor: 0},
			},
			expectedKeys: []string{"field1"},
			expectedVals: []string{"value1"},
		},
		{
			name: "empty result",
			entries: []ScanEntry{
				{Elements: []string{}, Cursor: 0},
			},
			expectedKeys: []string{},
			expectedVals: []string{},
		},
		{
			name:    "error case",
			err:     errors.New("scan error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			scanner := NewScanner(func(cursor uint64) (ScanEntry, error) {
				if tt.err != nil {
					return ScanEntry{}, tt.err
				}
				if callCount >= len(tt.entries) {
					return ScanEntry{}, errors.New("unexpected call")
				}
				entry := tt.entries[callCount]
				callCount++
				return entry, nil
			})

			var resultKeys, resultVals []string
			for key, val := range scanner.Iter2() {
				resultKeys = append(resultKeys, key)
				resultVals = append(resultVals, val)
			}

			if tt.wantErr {
				if scanner.Err() == nil {
					t.Error("expected error but got none")
				}
			} else {
				if scanner.Err() != nil {
					t.Errorf("unexpected error: %v", scanner.Err())
				}
				if (len(resultKeys) != 0 || len(tt.expectedKeys) != 0) && !reflect.DeepEqual(resultKeys, tt.expectedKeys) {
					t.Errorf("keys: got %v, want %v", resultKeys, tt.expectedKeys)
				}
				if (len(resultVals) != 0 || len(tt.expectedVals) != 0) && !reflect.DeepEqual(resultVals, tt.expectedVals) {
					t.Errorf("values: got %v, want %v", resultVals, tt.expectedVals)
				}
			}
		})
	}

	t.Run("early exit", func(t *testing.T) {
		callCount := 0
		entries := []ScanEntry{
			{Elements: []string{"field1", "value1"}, Cursor: 10},
		}
		scanner := NewScanner(func(cursor uint64) (ScanEntry, error) {
			if callCount >= len(entries) {
				return ScanEntry{}, errors.New("unexpected call")
			}
			entry := entries[callCount]
			callCount++
			return entry, nil
		})
		for range scanner.Iter2() {
			break
		}
		if scanner.Err() != nil {
			t.Errorf("unexpected error: %v", scanner.Err())
		}
	})
}

func TestAZAffinityNodesSelection(t *testing.T) {
	// Setup AZs can be changed per test
	var primAZ, rep1AZ, rep2AZ, rep3AZ string

	t.Run("standalone client node selectors", func(t *testing.T) {
		primary := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "primary"), nil)
			},
			AZFn: func() string { return primAZ },
		}
		replica1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "replica1"), nil)
			},
			AZFn: func() string { return rep1AZ },
		}
		replica2 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "replica2"), nil)
			},
			AZFn: func() string { return rep2AZ },
		}

		buildClient := func(selector ReadNodeSelectorFunc, replicas []string) *standalone {
			c, err := newStandaloneClient(
				&ClientOption{
					InitAddress: []string{"127.0.0.1:6379"},
					Standalone: StandaloneOption{
						ReplicaAddress: replicas,
					},
					EnableReplicaAZInfo: true,
					SendToReplicas:      func(cmd Completed) bool { return true },
					ReadNodeSelector:    selector,
				},
				func(dst string, opt *ClientOption) conn {
					switch dst {
					case "127.0.0.1:6379":
						return primary
					case "127.0.0.1:6380":
						return replica1
					case "127.0.0.1:6381":
						return replica2
					}
					return nil
				},
				newRetryer(defaultRetryDelayFn),
			)
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			return c
		}

		t.Run("PreferReplicaNodeSelector", func(t *testing.T) {
			primAZ, rep1AZ, rep2AZ = "us-west-1", "us-west-1", "us-west-2"
			client := buildClient(PreferReplicaNodeSelector(), []string{"127.0.0.1:6380", "127.0.0.1:6381"})

			// Round-Robin check: must hit replicas, never primary
			for range 10 {
				resp := client.Do(context.Background(), client.B().Get().Key("k").Build())
				val, _ := resp.ToString()
				if val == "primary" {
					t.Fatalf("expected replica, got primary")
				}
				if val != "replica1" && val != "replica2" {
					t.Fatalf("unexpected node: %v", val)
				}
			}
		})

		t.Run("AZAffinityNodeSelector", func(t *testing.T) {
			primAZ, rep1AZ, rep2AZ = "us-west-1", "us-west-1", "us-west-2"
			client := buildClient(AZAffinityNodeSelector("us-west-2"), []string{"127.0.0.1:6380", "127.0.0.1:6381"})

			resp := client.Do(context.Background(), client.B().Get().Key("k").Build())
			if val, _ := resp.ToString(); val != "replica2" {
				t.Fatalf("expected replica2, got %v", val)
			}
		})

		t.Run("AZAffinityNodeSelector Any Replica", func(t *testing.T) {
			primAZ, rep1AZ, rep2AZ = "us-west-1", "us-west-1", "us-west-2"
			client := buildClient(AZAffinityNodeSelector("us-west-3"), []string{"127.0.0.1:6380", "127.0.0.1:6381"})

			// Run multiple times to verify we hit replicas and NOT primary
			for range 10 {
				resp := client.Do(context.Background(), client.B().Get().Key("k").Build())
				val, _ := resp.ToString()

				if val == "primary" {
					t.Fatalf("expected any replica fallback, got primary")
				}
				if val != "replica1" && val != "replica2" {
					t.Fatalf("unexpected node: %v", val)
				}
			}
		})

		t.Run("AZAffinityReplicasAndPrimary Same-AZ Replica", func(t *testing.T) {
			// Client: A, Rep1: A, Primary: B
			primAZ, rep1AZ, rep2AZ = "zone-B", "zone-A", "zone-B"
			client := buildClient(AZAffinityReplicasAndPrimaryNodeSelector("zone-A"), []string{"127.0.0.1:6380", "127.0.0.1:6381"})

			resp := client.Do(context.Background(), client.B().Get().Key("k").Build())
			if val, _ := resp.ToString(); val != "replica1" {
				t.Fatalf("expected replica1, got %v", val)
			}
		})

		t.Run("AZAffinityReplicasAndPrimary Same-AZ Primary", func(t *testing.T) {
			// Client: A, Primary: A, Replicas: B
			primAZ, rep1AZ, rep2AZ = "zone-A", "zone-B", "zone-B"
			client := buildClient(AZAffinityReplicasAndPrimaryNodeSelector("zone-A"), []string{"127.0.0.1:6380", "127.0.0.1:6381"})

			resp := client.Do(context.Background(), client.B().Get().Key("k").Build())
			if val, _ := resp.ToString(); val != "primary" {
				t.Fatalf("expected primary, got %v", val)
			}
		})

		t.Run("AZAffinityReplicasAndPrimary Remote Replica", func(t *testing.T) {
			// Client: A, All Nodes: B
			primAZ, rep1AZ, rep2AZ = "zone-B", "zone-B", "zone-B"
			client := buildClient(AZAffinityReplicasAndPrimaryNodeSelector("zone-A"), []string{"127.0.0.1:6380", "127.0.0.1:6381"})

			for range 10 {
				resp := client.Do(context.Background(), client.B().Get().Key("k").Build())
				val, _ := resp.ToString()
				if val == "primary" {
					t.Fatalf("expected remote replica, got primary")
				}
			}
		})

		t.Run("AZAffinityReplicasAndPrimary Primary Fallback", func(t *testing.T) {
			primAZ = "zone-B"
			client := buildClient(AZAffinityReplicasAndPrimaryNodeSelector("zone-A"), []string{}) // Empty replicas

			resp := client.Do(context.Background(), client.B().Get().Key("k").Build())
			if val, _ := resp.ToString(); val != "primary" {
				t.Fatalf("expected primary, got %v", val)
			}
		})
	})

	t.Run("cluster client node selectors", func(t *testing.T) {
		primary := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Contains(strings.ToUpper(strings.Join(cmd.Commands(), " ")), "CLUSTER SLOTS") {
					return slotsMultiRespWithMultiReplicas
				}
				return newResult(strmsg('+', "primary"), nil)
			},
			AZFn: func() string { return primAZ },
		}
		rep1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return newResult(strmsg('+', "replica1"), nil) },
			AZFn: func() string { return rep1AZ },
		}
		rep2 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return newResult(strmsg('+', "replica2"), nil) },
			AZFn: func() string { return rep2AZ },
		}
		rep3 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return newResult(strmsg('+', "replica3"), nil) },
			AZFn: func() string { return rep3AZ },
		}

		buildCluster := func(selector ReadNodeSelectorFunc) *clusterClient {
			c, err := newClusterClient(
				&ClientOption{
					InitAddress:         []string{"127.0.0.1:0"},
					SendToReplicas:      func(cmd Completed) bool { return true },
					EnableReplicaAZInfo: true,
					ReadNodeSelector:    selector,
				},
				func(dst string, opt *ClientOption) conn {
					switch dst {
					case "127.0.0.2:1", "127.0.1.2:1":
						return rep1
					case "127.0.0.3:2", "127.0.1.3:2":
						return rep2
					case "127.0.0.4:3", "127.0.1.4:3":
						return rep3
					default:
						return primary
					}
				},
				newRetryer(defaultRetryDelayFn),
			)
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			// Force topology refresh
			c.refresh(context.Background())
			return c
		}

		t.Run("AZAffinityNodeSelector", func(t *testing.T) {
			primAZ = "az-1"
			rep1AZ = "az-1" // Same as Primary
			rep2AZ = "az-2" // Target
			rep3AZ = "az-3"

			// Client in az-2, should hit replica2
			client := buildCluster(AZAffinityNodeSelector("az-2"))

			resp := client.Do(context.Background(), client.B().Get().Key("k").Build())
			if val, _ := resp.ToString(); val != "replica2" {
				t.Fatalf("expected replica2 (same AZ), got %v", val)
			}
		})

		t.Run("AZAffinityNodeSelector Fallback", func(t *testing.T) {
			primAZ, rep1AZ, rep2AZ, rep3AZ = "az-1", "az-1", "az-2", "az-3"

			// Client in az-4 (No match), should hit any replica (RR)
			client := buildCluster(AZAffinityNodeSelector("az-4"))

			for range 10 {
				resp := client.Do(context.Background(), client.B().Get().Key("k").Build())
				val, _ := resp.ToString()
				if val == "primary" {
					t.Fatalf("expected any replica, got primary")
				}
			}
		})

		t.Run("AZAffinityReplicasAndPrimary Same-AZ Primary", func(t *testing.T) {
			// Client in az-1. Replica 1 is az-1.
			// Wait, we want to test Priority 2 (Primary).
			// So let's make all Replicas "remote".
			primAZ = "az-1"
			rep1AZ, rep2AZ, rep3AZ = "az-2", "az-3", "az-4"

			client := buildCluster(AZAffinityReplicasAndPrimaryNodeSelector("az-1"))

			// Should prioritize Primary (Same AZ) over Replicas (Different AZ)
			resp := client.Do(context.Background(), client.B().Get().Key("k").Build())
			if val, _ := resp.ToString(); val != "primary" {
				t.Fatalf("expected primary (same AZ), got %v", val)
			}
		})
	})

	t.Run("zero allocation guarantee", func(t *testing.T) {
		count := 200000
		nodes := make([]NodeInfo, count)
		targetAZ := "us-east-1a"

		for i := range count {
			// Every 10th node is in the target AZ
			if i%10 == 0 {
				nodes[i] = NodeInfo{AZ: targetAZ}
			} else {
				nodes[i] = NodeInfo{AZ: "other-az"}
			}
		}
		clientAZ := "us-east-1a"

		t.Run("PreferReplicaNodeSelector", ShouldNotAllocate(func() {
			PreferReplicaNodeSelector()(0, nodes)
		}))
		t.Run("AZAffinityNodeSelector", ShouldNotAllocate(func() {
			AZAffinityNodeSelector(clientAZ)(0, nodes)
		}))
		t.Run("AZAffinityReplicasAndPrimaryNodeSelector", ShouldNotAllocate(func() {
			AZAffinityReplicasAndPrimaryNodeSelector(clientAZ)(0, nodes)
		}))
		t.Run("AZAffinityNodeSelector (Fallback)", ShouldNotAllocate(func() {
			AZAffinityNodeSelector("us-unknown")(0, nodes)
		}))
	})
}
