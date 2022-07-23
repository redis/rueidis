package rueidis

import (
	"context"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

//gocyclo:ignore
func TestMGetCache(t *testing.T) {
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(&ClientOption{InitAddress: []string{""}}, m, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate DoCache", func(t *testing.T) {
			m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				if !reflect.DeepEqual(cmd.Commands(), []string{"MGET", "1", "2"}) || ttl != 100 {
					t.Fatalf("unexpected command %v, %v", cmd, ttl)
				}
				return newResult(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "1"}, {typ: '+', string: "2"}}}, nil)
			}
			if v, err := MGetCache(client, context.Background(), 100, []string{"1", "2"}); err != nil || v["1"].string != "1" || v["2"].string != "2" {
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
			m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{}, context.Canceled)
			}
			if v, err := MGetCache(client, ctx, 100, []string{"1", "2"}); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
	t.Run("cluster client", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return slotsResp
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate DoCache", func(t *testing.T) {
			keys := make([]string, 100)
			for i := range keys {
				keys[i] = strconv.Itoa(i)
			}
			m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				for _, key := range keys {
					if reflect.DeepEqual(cmd.Commands(), []string{"MGET", key}) && ttl == 100 {
						return newResult(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: key}}}, nil)
					}
				}
				t.Fatalf("unexpected command %v, %v", cmd, ttl)
				return RedisResult{}
			}
			v, err := MGetCache(client, context.Background(), 100, keys)
			if err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			for _, key := range keys {
				if v[key].string != key {
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
			m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{}, context.Canceled)
			}
			if v, err := MGetCache(client, ctx, 100, []string{"1", "2"}); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
}

//gocyclo:ignore
func TestJsonMGetCache(t *testing.T) {
	t.Run("single client", func(t *testing.T) {
		m := &mockConn{}
		client, err := newSingleClient(&ClientOption{InitAddress: []string{""}}, m, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate DoCache", func(t *testing.T) {
			m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				if !reflect.DeepEqual(cmd.Commands(), []string{"JSON.MGET", "1", "2", "$"}) || ttl != 100 {
					t.Fatalf("unexpected command %v, %v", cmd, ttl)
				}
				return newResult(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "1"}, {typ: '+', string: "2"}}}, nil)
			}
			if v, err := JsonMGetCache(client, context.Background(), 100, []string{"1", "2"}, "$"); err != nil || v["1"].string != "1" || v["2"].string != "2" {
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
			m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{}, context.Canceled)
			}
			if v, err := JsonMGetCache(client, ctx, 100, []string{"1", "2"}, "$"); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
	t.Run("cluster client", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return slotsResp
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		t.Run("Delegate DoCache", func(t *testing.T) {
			keys := make([]string, 100)
			for i := range keys {
				keys[i] = strconv.Itoa(i)
			}
			m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				for _, key := range keys {
					if reflect.DeepEqual(cmd.Commands(), []string{"JSON.MGET", key, "$"}) && ttl == 100 {
						return newResult(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: key}}}, nil)
					}
				}
				t.Fatalf("unexpected command %v, %v", cmd, ttl)
				return RedisResult{}
			}
			v, err := JsonMGetCache(client, context.Background(), 100, keys, "$")
			if err != nil {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			for _, key := range keys {
				if v[key].string != key {
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
			m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{}, context.Canceled)
			}
			if v, err := JsonMGetCache(client, ctx, 100, []string{"1", "2"}, "$"); err != context.Canceled {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		})
	})
}
