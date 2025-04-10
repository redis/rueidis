package rueidis

import (
	"context"
	"errors"
	"io"
	"reflect"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

//gocyclo:ignore
func TestSentinelClientInit(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("Init no nodes", func(t *testing.T) {
		if _, err := newSentinelClient(
			&ClientOption{InitAddress: []string{}},
			func(dst string, opt *ClientOption) conn { return nil },
			newRetryer(defaultRetryDelayFn),
		); err != ErrNoAddr {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Init no dialable", func(t *testing.T) {
		v := errors.New("dial err")
		if _, err := newSentinelClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return &mockConn{DialFn: func() error { return v }} },
			newRetryer(defaultRetryDelayFn),
		); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh err", func(t *testing.T) {
		v := errors.New("refresh err")
		if _, err := newSentinelClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoMultiFn: func(cmd ...Completed) *redisresults { return &redisresults{s: []RedisResult{newErrResult(v)}} },
				}
			},
			newRetryer(defaultRetryDelayFn),
		); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh retry", func(t *testing.T) {
		v := errors.New("refresh err")
		s0 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					newErrResult(v),
					newErrResult(v),
				}}
			},
		}
		s1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "0"),
						}),
					})},
					newErrResult(v),
				}}
			},
		}
		s2 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "3"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "5"),
					})},
				}}
			},
		}
		s3 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "4"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "6"),
					})},
				}}
			},
		}
		s4 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "2"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "7"),
					})},
				}}
			},
		}
		client, err := newSentinelClient(
			&ClientOption{InitAddress: []string{":0", ":1", ":2"}},
			func(dst string, opt *ClientOption) conn {
				if dst == ":0" {
					return s0
				}
				if dst == ":1" {
					return s1
				}
				if dst == ":2" {
					return s2
				}
				if dst == ":3" {
					return s3
				}
				if dst == ":4" {
					return s4
				}
				if dst == ":5" {
					return &mockConn{
						DialFn: func() error { return v },
					}
				}
				if dst == ":6" {
					return &mockConn{
						DoFn: func(cmd Completed) RedisResult {
							return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "slave")})}
						},
					}
				}
				if dst == ":7" {
					return &mockConn{
						DoFn: func(cmd Completed) RedisResult {
							return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
						},
					}
				}
				return nil
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if client.sConn == nil {
			t.Fatalf("unexpected nil sentinel conn")
		}
		if client.mConn.Load() == nil {
			t.Fatalf("unexpected nil master conn")
		}
		client.Close()
	})

	t.Run("Refresh retry replica-only client", func(t *testing.T) {
		v := errors.New("refresh err")
		slaveWithMultiError := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					newErrResult(v),
					newErrResult(v),
				}}
			},
		}
		slaveWithReplicaResponseErr := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "0"),
						}),
					})},
					newErrResult(v),
				}}
			},
		}
		sentinelWithFaultiSlave := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "3"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "6"),
						}),
					})},
				}}
			},
		}
		// this connection will fail because OK slave is in 's-down' status
		// since the next 2 connections won't update sentinel list,
		// we update the list here.
		sentinelWithHealthySlaveInSDown := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "4"),
						}),
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "32"),
						}),
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "31"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "6"),
						}),
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "8"),
							strmsg('+', "s-down-time"), strmsg('+', "1"),
						}),
					})},
				}}
			},
		}
		sentinelWithoutEligibleSlave := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "32"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "8"),
							strmsg('+', "s-down-time"), strmsg('+', "1"),
						}),
					})},
				}}
			},
		}

		sentinelWithInvalidMapResponse := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "4"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						RedisMessage(*Nil),
					})},
				}}
			},
		}
		sentinelWithMasterRoleAsSlave := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "5"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "7"),
						}),
					})},
				}}
			},
		}
		sentinelWithOKResponse := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "2"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "8"),
						}),
					})},
				}}
			},
		}

		client, err := newSentinelClient(
			&ClientOption{InitAddress: []string{":0", ":1", ":2"}, ReplicaOnly: true},
			func(dst string, opt *ClientOption) conn {
				if dst == ":0" {
					return slaveWithMultiError
				}
				if dst == ":1" {
					return slaveWithReplicaResponseErr
				}
				if dst == ":2" {
					return sentinelWithFaultiSlave
				}
				if dst == ":3" {
					return sentinelWithHealthySlaveInSDown
				}
				if dst == ":31" {
					return sentinelWithoutEligibleSlave
				}

				if dst == ":32" {
					return sentinelWithInvalidMapResponse
				}

				if dst == ":4" {
					return sentinelWithMasterRoleAsSlave
				}
				if dst == ":5" {
					return sentinelWithOKResponse
				}
				if dst == ":6" {
					return &mockConn{
						DialFn: func() error { return v },
					}
				}
				if dst == ":7" {
					return &mockConn{
						DoFn: func(cmd Completed) RedisResult {
							return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
						},
					}
				}
				if dst == ":8" {
					return &mockConn{
						DoFn: func(cmd Completed) RedisResult {
							return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "slave")})}
						},
					}
				}
				return nil
			},
			newRetryer(defaultRetryDelayFn),
		)
		if client.sAddr != ":5" && err == nil {
			t.Fatalf("expected error but got nil with sentinel %s", client.sAddr)
		}

		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if client.sConn == nil {
			t.Fatalf("unexpected nil sentinel conn")
		}
		if client.mConn.Load() == nil {
			t.Fatalf("unexpected nil slave conn")
		}
		client.Close()
	})

	t.Run("Refresh retry 2", func(t *testing.T) {
		v := errors.New("refresh err")
		s0 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "1"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "2"),
					})},
				}}
			},
		}
		s1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "0"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "3"),
					})},
				}}
			},
		}
		client, err := newSentinelClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				if dst == ":0" {
					return s0
				}
				if dst == ":1" {
					return s1
				}
				if dst == ":2" {
					return &mockConn{
						DoFn: func(cmd Completed) RedisResult { return newErrResult(v) },
					}
				}
				if dst == ":3" {
					return &mockConn{
						DoFn: func(cmd Completed) RedisResult {
							return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
						},
					}
				}
				return nil
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if client.sConn == nil {
			t.Fatalf("unexpected nil sentinel conn")
		}
		if client.mConn.Load() == nil {
			t.Fatalf("unexpected nil master conn")
		}
		client.Close()
	})

	t.Run("sentinel disconnect", func(t *testing.T) {
		trigger := make(chan error)
		disconnect := int32(0)
		s0closed := int32(0)
		r3closed := int32(0)
		s0 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if atomic.LoadInt32(&disconnect) == 1 {
					return newErrResult(errors.New("die"))
				}
				return RedisResult{}
			},
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "1"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "3"),
					})},
				}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				if err, ok := <-trigger; ok {
					return err
				}
				return ErrClosing
			},
			CloseFn: func() {
				atomic.StoreInt32(&s0closed, 1)
			},
		}
		s1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "2"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "3"),
					})},
				}}
			},
		}
		s2 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "0"),
						}),
					})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "4"),
					})},
				}}
			},
		}
		r3 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if atomic.LoadInt32(&disconnect) == 1 {
					return newErrResult(errors.New("die"))
				}
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
			},
			CloseFn: func() {
				atomic.StoreInt32(&r3closed, 1)
			},
			ErrorFn: func() error {
				if atomic.LoadInt32(&disconnect) == 1 {
					return errClosing
				}
				return nil
			},
		}
		r4 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
			},
		}
		client, err := newSentinelClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				if dst == ":0" {
					return s0
				}
				if dst == ":1" {
					return s1
				}
				if dst == ":2" {
					return s2
				}
				if dst == ":3" {
					return r3
				}
				if dst == ":4" {
					return r4
				}
				return nil
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		atomic.StoreInt32(&disconnect, 1)
		trigger <- errors.New("reconnect")
		close(trigger)
		for {
			t.Log("wait switch master")
			if client.mConn.Load().(*mockConn) == r4 {
				break
			}
		}
		if atomic.LoadInt32(&s0closed) != 1 {
			t.Fatalf("s0 not closed")
		}
		if atomic.LoadInt32(&r3closed) != 1 {
			t.Fatalf("r3 not closed")
		}
		client.Close()
	})
}

func TestSentinelRefreshAfterClose(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	first := true
	s0 := &mockConn{
		DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
		DoMultiFn: func(multi ...Completed) *redisresults {
			if first {
				first = true
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "1"),
					})},
				}}
			}
			return &redisresults{s: []RedisResult{newErrResult(ErrClosing), newErrResult(ErrClosing)}}
		},
	}
	m := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
		},
	}
	client, err := newSentinelClient(
		&ClientOption{InitAddress: []string{":0"}},
		func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			if dst == ":1" {
				return m
			}
			return nil
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	client.Close()
	if err := client.refresh(); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestSentinelSwitchAfterClose(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	first := true
	s0 := &mockConn{
		DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
		DoMultiFn: func(multi ...Completed) *redisresults {
			return &redisresults{s: []RedisResult{
				{val: slicemsg('*', []RedisMessage{})},
				{val: slicemsg('*', []RedisMessage{
					strmsg('+', ""), strmsg('+', "1"),
				})},
			}}
		},
	}
	m := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			if first {
				first = false
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
			}
			return newErrResult(ErrClosing)
		},
	}
	client, err := newSentinelClient(
		&ClientOption{InitAddress: []string{":0"}},
		func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			if dst == ":1" {
				return m
			}
			return nil
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	client.Close()
	if err := client._switchTarget(":1"); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
}

//gocyclo:ignore
func TestSentinelClientDelegate(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	s0 := &mockConn{
		DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
		DoMultiFn: func(multi ...Completed) *redisresults {
			return &redisresults{s: []RedisResult{
				{val: slicemsg('*', []RedisMessage{})},
				{val: slicemsg('*', []RedisMessage{
					strmsg('+', ""), strmsg('+', "1"),
				})},
			}}
		},
	}
	m := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
		},
		AddrFn: func() string { return ":1" },
	}
	client, err := newSentinelClient(
		&ClientOption{InitAddress: []string{":0"}},
		func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			if dst == ":1" {
				return m
			}
			return nil
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	defer client.Close()

	disabledCacheClient, err := newSentinelClient(
		&ClientOption{InitAddress: []string{":0"}, DisableCache: true},
		func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			if dst == ":1" {
				return m
			}
			return nil
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	defer disabledCacheClient.Close()

	t.Run("Delegate MGetCache", func(t *testing.T) {
		keys := []string{"key1", "key2"}
		expectedCommand := []string{"MGET", "key1", "key2"}

		m.DoMultiFn = func(cmd ...Completed) *redisresults {
			if !reflect.DeepEqual(cmd[0].Commands(), expectedCommand) {
				t.Fatalf("unexpected command %v", cmd)
			}
			return &redisresults{s: []RedisResult{
				newResult(strmsg('+', "master"), nil),
			}}
		}

		ret, err := MGetCache(disabledCacheClient, context.Background(), time.Second, keys)
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}

		expected := map[string]RedisMessage{
			"key1": strmsg('+', "master"),
		}
		if !reflect.DeepEqual(ret, expected) {
			t.Fatalf("unexpected result %v, expected %v", ret, expected)
		}
	})

	t.Run("Nodes", func(t *testing.T) {
		if nodes := client.Nodes(); len(nodes) != 1 || nodes[":1"] == nil {
			t.Fatalf("unexpected nodes")
		}
	})

	t.Run("Mode", func(t *testing.T) {
		if mode := client.Mode(); mode != ClientModeSentinel {
			t.Fatalf("unexpected mode %v", mode)
		}
	})

	t.Run("Delegate Do", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		m.DoFn = func(cmd Completed) RedisResult {
			if !reflect.DeepEqual(cmd.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", cmd)
			}
			return newResult(strmsg('+', "Do"), nil)
		}
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoStream", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		m.DoStreamFn = func(cmd Completed) RedisResultStream {
			return RedisResultStream{e: errors.New(cmd.Commands()[1])}
		}
		if s := client.DoStream(context.Background(), c); s.Error().Error() != "Do" {
			t.Fatalf("unexpected response %v", s.Error())
		}
	})

	t.Run("Delegate DoMulti", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		m.DoMultiFn = func(cmd ...Completed) *redisresults {
			if !reflect.DeepEqual(cmd[0].Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", cmd)
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "Do"), nil)}}
		}
		if len(client.DoMulti(context.Background())) != 0 {
			t.Fatalf("unexpected response length")
		}
		if v, err := client.DoMulti(context.Background(), c)[0].ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiStream", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		m.DoMultiStreamFn = func(cmd ...Completed) MultiRedisResultStream {
			return MultiRedisResultStream{e: errors.New(cmd[0].Commands()[1])}
		}
		if s := client.DoMultiStream(context.Background()); s.Error() != io.EOF {
			t.Fatalf("unexpected response %v", err)
		}
		if s := client.DoMultiStream(context.Background(), c); s.Error().Error() != "Do" {
			t.Fatalf("unexpected response %v", s.Error())
		}
	})

	t.Run("Delegate DoCache", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		m.DoCacheFn = func(cmd Cacheable, ttl time.Duration) RedisResult {
			if !reflect.DeepEqual(cmd.Commands(), c.Commands()) || ttl != 100 {
				t.Fatalf("unexpected command %v, %v", cmd, ttl)
			}
			return newResult(strmsg('+', "DoCache"), nil)
		}
		if v, err := client.DoCache(context.Background(), c, 100).ToString(); err != nil || v != "DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiCache", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			if !reflect.DeepEqual(multi[0].Cmd.Commands(), c.Commands()) || multi[0].TTL != 100 {
				t.Fatalf("unexpected command %v, %v", multi[0].Cmd, multi[0].TTL)
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "DoCache"), nil)}}
		}
		if len(client.DoMultiCache(context.Background())) != 0 {
			t.Fatalf("unexpected response length")
		}
		if v, err := client.DoMultiCache(context.Background(), CT(c, 100))[0].ToString(); err != nil || v != "DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Receive", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		hdl := func(message PubSubMessage) {}
		m.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			if !reflect.DeepEqual(subscribe.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", subscribe)
			}
			return nil
		}
		if err := client.Receive(context.Background(), c, hdl); err != nil {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Receive Redis Err", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		e := &RedisError{}
		m.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return e
		}
		if err := client.Receive(context.Background(), c, func(message PubSubMessage) {}); err != e {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Close", func(t *testing.T) {
		called := false
		m.CloseFn = func() { called = true }
		client.Close()
		if !called {
			t.Fatalf("Close is not delegated")
		}
	})

	t.Run("Dedicated Err", func(t *testing.T) {
		v := errors.New("fn err")
		if err := client.Dedicated(func(client DedicatedClient) error {
			return v
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated Delegate", func(t *testing.T) {
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{newResult(strmsg('+', "Delegate"), nil)}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			ErrorFn: func() error {
				return ErrClosing
			},
		}
		m.AcquireFn = func() wire {
			return w
		}
		stored := false
		m.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			if v := c.DoMulti(context.Background()); len(v) != 0 {
				t.Fatalf("received unexpected response %v", v)
			}
			for _, resp := range c.DoMulti(context.Background(), c.B().Get().Key("a").Build()) {
				if v, err := resp.ToString(); err != nil || v != "Delegate" {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			return nil
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
	})

	t.Run("Dedicate Delegate", func(t *testing.T) {
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{newResult(strmsg('+', "Delegate"), nil)}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			ErrorFn: func() error {
				return ErrClosing
			},
		}
		m.AcquireFn = func() wire {
			return w
		}
		stored := false
		m.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		c, cancel := client.Dedicate()
		if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v := c.DoMulti(context.Background()); len(v) != 0 {
			t.Fatalf("received unexpected response %v", v)
		}
		for _, resp := range c.DoMulti(context.Background(), c.B().Get().Key("a").Build()) {
			if v, err := resp.ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		cancel()
		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
	})
}

//gocyclo:ignore
func TestSentinelClientDelegateRetry(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	setup := func(t *testing.T) (client *sentinelClient, cb func()) {
		retry := uint32(0)
		trigger := make(chan error)
		s0 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				if atomic.LoadUint32(&retry) == 0 {
					return &redisresults{s: []RedisResult{
						{val: slicemsg('*', []RedisMessage{})},
						{val: slicemsg('*', []RedisMessage{
							strmsg('+', ""), strmsg('+', "1"),
						})},
					}}
				}
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "2"),
					})},
				}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				if err, ok := <-trigger; ok {
					return err
				}
				return ErrClosing
			},
			ErrorFn: func() error {
				return ErrClosing
			},
		}
		m1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if cmd == cmds.RoleCmd {
					return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
				}
				atomic.AddUint32(&retry, 1)
				return newErrResult(ErrClosing)
			},
			DoMultiFn: func(multi ...Completed) *redisresults {
				atomic.AddUint32(&retry, 1)
				return &redisresults{s: []RedisResult{newErrResult(ErrClosing)}}
			},
			DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
				atomic.AddUint32(&retry, 1)
				return newErrResult(ErrClosing)
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				atomic.AddUint32(&retry, 1)
				return ErrClosing
			},
		}
		m2 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if cmd == cmds.RoleCmd {
					return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
				}
				return RedisResult{val: strmsg('+', "OK")}
			},
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{{val: strmsg('+', "OK")}}}
			},
			DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
				return RedisResult{val: strmsg('+', "OK")}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return nil
			},
		}
		client, err := newSentinelClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				if dst == ":0" {
					return s0
				}
				if dst == ":1" {
					return m1
				}
				if dst == ":2" {
					return m2
				}
				return nil
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return client, func() {
			for atomic.LoadUint32(&retry) < 10 {
				time.Sleep(time.Millisecond * 100)
			}
			trigger <- errors.New("die")
			close(trigger)
		}
	}

	t.Run("Delegate Do", func(t *testing.T) {
		client, cb := setup(t)

		go func() {
			cb()
		}()

		v, err := client.Do(context.Background(), client.B().Get().Key("k").Build()).ToString()
		if err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}

		client.Close()
	})

	t.Run("Delegate DoMulti", func(t *testing.T) {
		client, cb := setup(t)

		go func() {
			cb()
		}()

		v, err := client.DoMulti(context.Background(), client.B().Get().Key("k").Build())[0].ToString()
		if err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}

		client.Close()
	})

	t.Run("Delegate DoCache", func(t *testing.T) {
		client, cb := setup(t)

		go func() {
			cb()
		}()

		v, err := client.DoCache(context.Background(), client.B().Get().Key("k").Cache(), time.Minute).ToString()
		if err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}

		client.Close()
	})

	t.Run("Delegate Receive", func(t *testing.T) {
		client, cb := setup(t)

		go func() {
			cb()
		}()

		err := client.Receive(context.Background(), client.B().Subscribe().Channel("k").Build(), func(msg PubSubMessage) {

		})
		if err != nil {
			t.Fatalf("unexpected resp %v", err)
		}

		client.Close()
	})
}

//gocyclo:ignore
func TestSentinelClientPubSub(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	var s0count, s0close, m1close, m2close, m4close int32

	messages := make(chan PubSubMessage)

	s0 := &mockConn{
		DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
		DoMultiFn: func(multi ...Completed) *redisresults {
			count := atomic.AddInt32(&s0count, 1)
			if (count-1)%2 == 0 {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "1"),
					})},
				}}
			}
			return &redisresults{s: []RedisResult{
				{val: slicemsg('*', []RedisMessage{})},
				{val: slicemsg('*', []RedisMessage{
					strmsg('+', ""), strmsg('+', "2"),
				})},
			}}
		},
		ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			for msg := range messages {
				fn(msg)
			}
			return ErrClosing
		},
		CloseFn: func() { atomic.AddInt32(&s0close, 1) },
	}
	m1 := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			if cmd == cmds.RoleCmd {
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
			}
			return RedisResult{val: strmsg('+', "OK")}
		},
		CloseFn: func() {
			atomic.AddInt32(&m1close, 1)
		},
	}
	m2 := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "slave")})}
		},
		CloseFn: func() { atomic.AddInt32(&m2close, 1) },
	}
	s3 := &mockConn{
		DoMultiFn: func(cmd ...Completed) *redisresults { return &redisresults{s: []RedisResult{newErrResult(errClosing)}} },
	}
	m4 := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			if cmd == cmds.RoleCmd {
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
			}
			return RedisResult{val: strmsg('+', "OK4")}
		},
		CloseFn: func() { atomic.AddInt32(&m4close, 1) },
	}

	client, err := newSentinelClient(
		&ClientOption{
			InitAddress: []string{":0"},
			Sentinel: SentinelOption{
				MasterSet: "test",
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			if dst == ":1" {
				return m1
			}
			if dst == ":2" {
				return m2
			}
			if dst == ":3" {
				return s3
			}
			if dst == ":4" {
				return m4
			}
			return nil
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	messages <- PubSubMessage{Channel: "+sentinel", Message: "sentinel 000000  3"}

	var added bool
	for !added {
		client.mu.Lock()
		added = client.sentinels.Front().Value.(string) == ":3"
		client.mu.Unlock()
		t.Log("wait +sentinel")
		time.Sleep(time.Millisecond * 100)
	}

	// switch to false master
	messages <- PubSubMessage{Channel: "+switch-master", Message: "test  1  2"}

	for atomic.LoadInt32(&m2close) < 2 {
		t.Log("wait false m2 to be close", atomic.LoadInt32(&m2close))
		time.Sleep(time.Millisecond * 100)
	}

	for atomic.LoadInt32(&s0count) < 3 {
		t.Log("wait s0 to be call third time", atomic.LoadInt32(&s0count))
		time.Sleep(time.Millisecond * 100)
	}

	v, err := client.Do(context.Background(), client.B().Get().Key("k").Build()).ToString()
	if err != nil || v != "OK" {
		t.Fatalf("unexpected resp %v %v", v, err)
	}

	// switch to master by reboot
	messages <- PubSubMessage{Channel: "+reboot", Message: "master test  4"}

	for atomic.LoadInt32(&m1close) < 1 {
		t.Log("wait old m1 to be close", atomic.LoadInt32(&m1close))
		time.Sleep(time.Millisecond * 100)
	}

	v, err = client.Do(context.Background(), client.B().Get().Key("k").Build()).ToString()
	if err != nil || v != "OK4" {
		t.Fatalf("unexpected resp %v %v", v, err)
	}

	close(messages)
	client.Close()

	for atomic.LoadInt32(&s0close) < 4 {
		t.Log("wait old s0 to be close", atomic.LoadInt32(&s0close))
		time.Sleep(time.Millisecond * 100)
	}
	for atomic.LoadInt32(&m4close) < 1 {
		t.Log("wait old m1 to be close", atomic.LoadInt32(&m4close))
		time.Sleep(time.Millisecond * 100)
	}
}

//gocyclo:ignore
func TestSentinelReplicaOnlyClientPubSub(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	var s0count, s0close, slave1close, slave2close, slave4close int32

	messages := make(chan PubSubMessage)

	s0 := &mockConn{
		DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
		DoMultiFn: func(multi ...Completed) *redisresults {
			count := atomic.AddInt32(&s0count, 1)
			remainder := (count - 1) % 3
			if remainder == 0 {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{})},
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "1"),
						}),
					})},
				}}
			} else if remainder == 1 {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{})},
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "2"),
						}),
					})},
				}}
			} else {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{})},
					{val: slicemsg('*', []RedisMessage{
						slicemsg('%', []RedisMessage{
							strmsg('+', "ip"), strmsg('+', ""),
							strmsg('+', "port"), strmsg('+', "4"),
						}),
					})},
				}}
			}
		},
		ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			for msg := range messages {
				fn(msg)
			}
			return ErrClosing
		},
		CloseFn: func() { atomic.AddInt32(&s0close, 1) },
	}
	slave1 := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			if cmd == cmds.RoleCmd {
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "slave")})}
			}
			return RedisResult{val: strmsg('+', "OK")}
		},
		CloseFn: func() {
			atomic.AddInt32(&slave1close, 1)
		},
	}
	slave2 := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
		},
		CloseFn: func() { atomic.AddInt32(&slave2close, 1) },
	}
	s3 := &mockConn{
		DoMultiFn: func(cmd ...Completed) *redisresults { return &redisresults{s: []RedisResult{newErrResult(errClosing)}} },
	}
	slave4 := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			if cmd == cmds.RoleCmd {
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "slave")})}
			}
			return RedisResult{val: strmsg('+', "OK4")}
		},
		CloseFn: func() { atomic.AddInt32(&slave4close, 1) },
	}

	client, err := newSentinelClient(
		&ClientOption{
			InitAddress: []string{":0"},
			Sentinel: SentinelOption{
				MasterSet: "replicaonly",
			},
			ReplicaOnly: true,
		},
		func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			if dst == ":1" {
				return slave1
			}
			if dst == ":2" {
				return slave2
			}
			if dst == ":3" {
				return s3
			}
			if dst == ":4" {
				return slave4
			}
			return nil
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	messages <- PubSubMessage{Channel: "+sentinel", Message: "sentinel 000000  3"}

	var added bool
	for !added {
		client.mu.Lock()
		added = client.sentinels.Front().Value.(string) == ":3"
		client.mu.Unlock()
		t.Log("wait +sentinel")
		time.Sleep(time.Millisecond * 100)
	}

	// event will be skipped because of first word
	messages <- PubSubMessage{Channel: "+slave", Message: "sla_ve 0:0 0 2 @ replicaonly 0 0"}

	v, err := client.Do(context.Background(), client.B().Get().Key("k").Build()).ToString()
	if err != nil || v != "OK" {
		t.Fatalf("unexpected resp %v %v", v, err)
	}

	// event will be skipped because of wrong master set name
	messages <- PubSubMessage{Channel: "+slave", Message: "slave 0:0 0 2 @ test 0 0"}

	v, err = client.Do(context.Background(), client.B().Get().Key("k").Build()).ToString()
	if err != nil || v != "OK" {
		t.Fatalf("unexpected resp %v %v", v, err)
	}

	// new slave with wrong role (2)
	// this won't directly switch to :2 like master
	// it will cause s0 to return :2 in DoMulti response
	messages <- PubSubMessage{Channel: "+slave", Message: "slave 0:0 0 2 @ replicaonly 0 0"}

	for atomic.LoadInt32(&slave2close) < 1 {
		t.Log("wait false slave2 to be close", atomic.LoadInt32(&slave2close))
		time.Sleep(time.Millisecond * 100)
	}

	for atomic.LoadInt32(&s0count) < 3 {
		t.Log("wait s0 to be call third time", atomic.LoadInt32(&s0count))
		time.Sleep(time.Millisecond * 100)
	}

	for atomic.LoadInt32(&slave1close) < 1 {
		t.Log("wait for slave1 to close (and for client to use slave4)", atomic.LoadInt32(&slave1close))
		time.Sleep(time.Millisecond * 100)
	}

	v, err = client.Do(context.Background(), client.B().Get().Key("k").Build()).ToString()
	if err != nil || v != "OK4" {
		t.Fatalf("unexpected resp %v %v", v, err)
	}

	// switch to new slave by reboot
	messages <- PubSubMessage{Channel: "+reboot", Message: "slave 0:0 0 1 @ replicaonly 0 0"}

	for atomic.LoadInt32(&slave4close) < 1 {
		t.Log("wait old slave4 to be close", atomic.LoadInt32(&slave4close))
		time.Sleep(time.Millisecond * 100)
	}

	v, err = client.Do(context.Background(), client.B().Get().Key("k").Build()).ToString()
	if err != nil || v != "OK" {
		t.Fatalf("unexpected resp %v %v", v, err)
	}

	close(messages)
	client.Close()

	for atomic.LoadInt32(&s0close) < 4 {
		t.Log("wait old s0 to be close", atomic.LoadInt32(&s0close))
		time.Sleep(time.Millisecond * 100)
	}
	for atomic.LoadInt32(&slave1close) < 2 {
		t.Log("wait old slave1 to be close", atomic.LoadInt32(&slave1close))
		time.Sleep(time.Millisecond * 100)
	}
}

func TestSentinelClientRetry(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	SetupClientRetry(t, func(m *mockConn) Client {
		m.DoOverride = map[string]func(cmd Completed) RedisResult{
			"SENTINEL SENTINELS masters": func(cmd Completed) RedisResult {
				return RedisResult{val: slicemsg('*', []RedisMessage{})}
			},
			"SENTINEL GET-MASTER-ADDR-BY-NAME masters": func(cmd Completed) RedisResult {
				return RedisResult{val: slicemsg('*', []RedisMessage{
					strmsg('+', ""), strmsg('+', "5"),
				})}
			},
			"ROLE": func(cmd Completed) RedisResult {
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
			},
		}
		m.ReceiveOverride = map[string]func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error{
			"SUBSCRIBE +sentinel +slave -sdown +sdown +switch-master +reboot": func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return nil
			},
		}
		c, err := newSentinelClient(
			&ClientOption{
				InitAddress: []string{":0"},
				Sentinel:    SentinelOption{MasterSet: "masters"},
			},
			func(dst string, opt *ClientOption) conn {
				return m
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return c
	})
}

func TestSentinelClientLoadingRetry(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())

	setup := func() (*sentinelClient, *mockConn, *mockConn) {
		s0 := &mockConn{
			DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{
					{val: slicemsg('*', []RedisMessage{})},
					{val: slicemsg('*', []RedisMessage{
						strmsg('+', ""), strmsg('+', "1"),
					})},
				}}
			},
		}
		m1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if cmd == cmds.RoleCmd {
					return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
				}
				return RedisResult{}
			},
		}
		client, err := newSentinelClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				if dst == ":0" {
					return s0
				}
				if dst == ":1" {
					return m1
				}
				return nil
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return client, s0, m1
	}

	t.Run("Do Retry on Loading", func(t *testing.T) {
		client, _, m1 := setup()
		attempts := 0
		m1.DoFn = func(cmd Completed) RedisResult {
			if cmd == cmds.RoleCmd {
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
			}
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}

		if v, err := client.Do(context.Background(), client.B().Get().Key("test").Build()).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if attempts != 2 {
			t.Fatalf("expected 2 attempts, got %v", attempts)
		}
	})

	t.Run("Do not retry on non-loading errors", func(t *testing.T) {
		client, _, m1 := setup()
		attempts := 0
		m1.DoFn = func(cmd Completed) RedisResult {
			if cmd == cmds.RoleCmd {
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
			}
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "ERR some other error"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}

		if err := client.Do(context.Background(), client.B().Get().Key("test").Build()).Error(); err == nil {
			t.Fatal("expected error but got nil")
		}
		if attempts != 1 {
			t.Fatalf("unexpected attempts %v, expected no retry", attempts)
		}
	})

	t.Run("DoMulti Retry on Loading", func(t *testing.T) {
		client, _, m1 := setup()
		attempts := 0
		m1.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}

		cmd := client.B().Get().Key("test").Build()
		resps := client.DoMulti(context.Background(), cmd)
		if len(resps) != 1 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoCache Retry on Loading", func(t *testing.T) {
		client, _, m1 := setup()
		attempts := 0
		m1.DoCacheFn = func(cmd Cacheable, ttl time.Duration) RedisResult {
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}

		cmd := client.B().Get().Key("test").Cache()
		if v, err := client.DoCache(context.Background(), cmd, time.Minute).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Retry on Loading", func(t *testing.T) {
		client, _, m1 := setup()
		attempts := 0
		m1.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}

		cmd := client.B().Get().Key("test").Cache()
		resps := client.DoMultiCache(context.Background(), CT(cmd, time.Minute))
		if len(resps) != 1 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Dedicated Do Retry on Loading", func(t *testing.T) {
		client, _, m1 := setup()
		attempts := 0
		m1.DoFn = func(cmd Completed) RedisResult {
			if cmd == cmds.RoleCmd {
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
			}
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}
		m1.AcquireFn = func() wire { return &mockWire{DoFn: m1.DoFn} }

		err := client.Dedicated(func(c DedicatedClient) error {
			if v, err := c.Do(context.Background(), c.B().Get().Key("test").Build()).ToString(); err != nil || v != "OK" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated DoMulti Retry on Loading", func(t *testing.T) {
		client, _, m1 := setup()
		attempts := 0
		m1.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		m1.AcquireFn = func() wire { return &mockWire{DoMultiFn: m1.DoMultiFn} }

		err := client.Dedicated(func(c DedicatedClient) error {
			resps := c.DoMulti(context.Background(), c.B().Get().Key("test").Build())
			if len(resps) != 1 {
				t.Fatalf("unexpected response length %v", len(resps))
			}
			if v, err := resps[0].ToString(); err != nil || v != "OK" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})
}
