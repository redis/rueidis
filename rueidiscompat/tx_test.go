// Copyright (c) 2013 The github.com/go-redis/redis Authors.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package rueidiscompat

import (
	"context"
	"errors"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"go.uber.org/mock/gomock"
)

var _ = Describe("RESP3 TxPipeline Commands", func() {
	testAdapterTxPipeline(true)
})

var _ = Describe("RESP2 TxPipeline Commands", func() {
	testAdapterTxPipeline(false)
})

func testAdapterTxPipeline(resp3 bool) {
	var adapter Cmdable

	BeforeEach(func() {
		if resp3 {
			adapter = adapterresp3
		} else {
			adapter = adapterresp2
		}
		Expect(adapter.FlushDB(ctx).Err()).NotTo(HaveOccurred())
		Expect(adapter.FlushAll(ctx).Err()).NotTo(HaveOccurred())
	})

	It("should TxPipelined", func() {
		var echo, ping *StringCmd
		rets, err := adapter.TxPipelined(ctx, func(pipe Pipeliner) error {
			echo = pipe.Echo(ctx, "hello")
			ping = pipe.Ping(ctx)
			Expect(echo.Err()).To(MatchError(errPipelineNotExecuted))
			Expect(ping.Err()).To(MatchError(errPipelineNotExecuted))
			return nil
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(rets).To(HaveLen(2))
		Expect(rets[0]).To(Equal(echo))
		Expect(rets[1]).To(Equal(ping))
		Expect(echo.Err()).NotTo(HaveOccurred())
		Expect(echo.Val()).To(Equal("hello"))
		Expect(ping.Err()).NotTo(HaveOccurred())
		Expect(ping.Val()).To(Equal("PONG"))
	})

	It("should TxPipeline", func() {
		pipe := adapter.TxPipeline()
		echo := pipe.Echo(ctx, "hello")
		ping := pipe.Ping(ctx)
		Expect(echo.Err()).To(MatchError(errPipelineNotExecuted))
		Expect(ping.Err()).To(MatchError(errPipelineNotExecuted))
		Expect(pipe.Len()).To(Equal(2))

		rets, err := pipe.Exec(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(pipe.Len()).To(Equal(0))
		Expect(rets).To(HaveLen(2))
		Expect(rets[0]).To(Equal(echo))
		Expect(rets[1]).To(Equal(ping))
		Expect(echo.Err()).NotTo(HaveOccurred())
		Expect(echo.Val()).To(Equal("hello"))
		Expect(ping.Err()).NotTo(HaveOccurred())
		Expect(ping.Val()).To(Equal("PONG"))
	})

	It("should Discard", func() {
		pipe := adapter.TxPipeline()
		echo := pipe.Echo(ctx, "hello")
		ping := pipe.Ping(ctx)

		pipe.Discard()
		Expect(pipe.Len()).To(Equal(0))

		rets, err := pipe.Exec(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(rets).To(HaveLen(0))

		Expect(echo.Err()).To(MatchError(errPipelineNotExecuted))
		Expect(ping.Err()).To(MatchError(errPipelineNotExecuted))
	})

	It("should Watch", func() {
		k1 := "_k1_"
		k2 := "_k2_"
		err := adapter.Watch(ctx, func(t Tx) error {
			if t.Get(ctx, k1).Err() != Nil {
				return errors.New("unclean")
			}
			if t.Get(ctx, k2).Err() != Nil {
				return errors.New("unclean")
			}
			_, err := t.TxPipelined(ctx, func(pipe Pipeliner) error {
				pipe.Set(ctx, k1, k1, 0)
				pipe.Set(ctx, k2, k2, 0)
				return nil
			})
			return err
		}, k1, k2)
		Expect(err).NotTo(HaveOccurred())
		Expect(adapter.Get(ctx, k1).Val()).To(Equal(k1))
		Expect(adapter.Get(ctx, k2).Val()).To(Equal(k2))
	})

	It("should Watch Abort", func() {
		k1 := "_k1_"
		ch := make(chan error)
		var del *IntCmd
		go func() {
			ch <- adapter.Watch(ctx, func(t Tx) error {
				ch <- nil
				<-ch
				_, err := t.TxPipelined(ctx, func(pipe Pipeliner) error {
					del = pipe.Del(ctx, k1)
					return nil
				})
				return err
			}, k1)
		}()
		<-ch
		Expect(adapter.Set(ctx, k1, k1, 0).Err()).NotTo(HaveOccurred())
		ch <- nil
		Expect(<-ch).To(MatchError(TxFailedErr))
		Expect(del.Err()).To(MatchError(TxFailedErr))
	})

	It("should set the command errors when EXEC is aborted", func() {
		k1 := "_k1_"
		var incr *IntCmd
		var set *Cmd
		_, err := adapter.TxPipelined(ctx, func(pipe Pipeliner) error {
			incr = pipe.Incr(ctx, k1)
			set = pipe.Do(ctx, "SET", k1)
			return nil
		})
		Expect(err).To(MatchError(ContainSubstring("EXECABORT")))
		Expect(incr.Err()).To(Equal(err))
		Expect(set.Err()).To(MatchError(ContainSubstring("wrong number of arguments")))
		Expect(adapter.Exists(ctx, k1).Val()).To(Equal(int64(0)))
	})

	It("should Unwatch and Close", func() {
		k1 := "_k1_"
		err := adapter.Watch(ctx, func(t Tx) error {
			Expect(t.Unwatch(ctx).Err()).NotTo(HaveOccurred())
			Expect(t.Close(ctx)).NotTo(HaveOccurred())
			_, err := t.TxPipelined(ctx, func(pipe Pipeliner) error {
				pipe.Del(ctx, k1)
				return nil
			})
			return err
		}, k1)
		Expect(err).To(MatchError(rueidis.ErrDedicatedClientRecycled))
	})

	if resp3 {
		It("should catch first cmder error", func() {
			k1 := "_k1_not_exists"
			var cmder *JSONCmd
			_, err := adapter.TxPipelined(ctx, func(pipe Pipeliner) error {
				pipe.Del(ctx, k1)
				cmder = pipe.JSONGet(ctx, k1)
				return nil
			})
			Expect(err).To(Equal(cmder.Err()))
		})
	}

	It("should panic for GetToBuffer in TxPipeline", func() {
		pipe := adapter.TxPipeline()

		Expect(func() {
			pipe.GetToBuffer(ctx, "key", make([]byte, 10))
		}).To(Panic())
	})
}

func TestTxPipelineExecFailure(t *testing.T) {
	queued := mock.Result(mock.RedisString("QUEUED"))
	for _, tc := range []struct {
		name  string
		queue rueidis.RedisResult // the reply to queuing the second command
		exec  rueidis.RedisResult
		err   error
	}{
		{
			name:  "WATCH conflict",
			queue: queued,
			exec:  mock.Result(mock.RedisNil()),
			err:   TxFailedErr,
		},
		{
			name:  "EXECABORT",
			queue: mock.Result(mock.RedisError("ERR wrong number of arguments for 'set' command")),
			exec:  mock.Result(mock.RedisError("EXECABORT Transaction discarded because of previous errors.")),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			client := mock.NewClient(ctrl)
			client.EXPECT().DoMulti(gomock.Any(), gomock.Any()).Return([]rueidis.RedisResult{
				mock.Result(mock.RedisString("OK")), queued, tc.queue, tc.exec,
			})

			pipe := newTxPipeline(client)
			incr := pipe.Incr(context.Background(), "k")
			set := pipe.Do(context.Background(), "SET", "k")
			_, err := pipe.Exec(context.Background())

			want := tc.err
			if want == nil {
				want = tc.exec.Error()
			}
			if err == nil || err.Error() != want.Error() {
				t.Fatalf("Exec error: got %v, want %v", err, want)
			}
			if incr.Err() == nil || incr.Err().Error() != want.Error() {
				t.Fatalf("INCR error: got %v, want %v", incr.Err(), want)
			}
			wantSet := want
			if qerr := tc.queue.Error(); qerr != nil {
				wantSet = qerr
			}
			if set.Err() == nil || set.Err().Error() != wantSet.Error() {
				t.Fatalf("SET error: got %v, want %v", set.Err(), wantSet)
			}
		})
	}
}
