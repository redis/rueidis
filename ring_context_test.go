package rueidis

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

func TestRingCanceledBeforeEnqueue(t *testing.T) {
	r := newRing(1)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if ch, err := r.PutOne(ctx, cmds.PingCmd); ch != nil || !errors.Is(err, context.Canceled) {
		t.Errorf("PutOne = (%v, %v), want (nil, context.Canceled)", ch, err)
	}
	if ch, err := r.PutMulti(ctx, []Completed{cmds.PingCmd}, nil); ch != nil || !errors.Is(err, context.Canceled) {
		t.Errorf("PutMulti = (%v, %v), want (nil, context.Canceled)", ch, err)
	}
	if _, _, ch := r.NextWriteCmd(); ch != nil {
		t.Error("already canceled context enqueued a command")
	}
}

func TestRingFullContext(t *testing.T) {
	for _, useMulti := range []bool{false, true} {
		for _, useDeadline := range []bool{false, true} {
			t.Run(ringContextTestName(useMulti, useDeadline), func(t *testing.T) {
				synctest.Test(t, func(t *testing.T) {
					r := newRing(1)
					for range len(r.store) {
						if _, err := r.PutOne(context.Background(), cmds.PingCmd); err != nil {
							t.Fatal(err)
						}
					}

					ctx, cancel := context.WithCancel(context.Background())
					want := context.Canceled
					if useDeadline {
						cancel()
						ctx, cancel = context.WithTimeout(context.Background(), time.Second)
						want = context.DeadlineExceeded
					}
					defer cancel()
					done := make(chan error, 1)
					go func() {
						if useMulti {
							_, err := r.PutMulti(ctx, []Completed{cmds.PingCmd}, make([]RedisResult, 1))
							done <- err
							return
						}
						_, err := r.PutOne(ctx, cmds.PingCmd)
						done <- err
					}()

					synctest.Wait()
					select {
					case err := <-done:
						t.Fatalf("enqueue returned while the ring was full: %v", err)
					default:
					}
					if useDeadline {
						time.Sleep(time.Second)
					} else {
						cancel()
					}
					synctest.Wait()
					if err := <-done; !errors.Is(err, want) {
						t.Fatalf("enqueue error = %v, want %v", err, want)
					}

					// 回收已填充的槽位，确认取消的请求没有留下空槽。
					for range len(r.store) {
						one, multi, ch := r.NextWriteCmd()
						if (one.IsEmpty() && multi == nil) || ch == nil {
							t.Fatal("filled slot was lost after cancellation")
						}
						_, _, resultCh, _ := r.NextResultCh()
						if resultCh != ch {
							t.Fatal("ring returned a different response channel")
						}
						r.FinishResult()
					}

					ch, err := r.PutOne(context.Background(), cmds.PingCmd)
					if err != nil {
						t.Fatal(err)
					}
					one, _, written := r.NextWriteCmd()
					_, _, read, _ := r.NextResultCh()
					r.FinishResult()
					if one.IsEmpty() || written != ch || read != ch {
						t.Fatal("ring did not advance after a canceled enqueue")
					}
				})
			})
		}
	}
}

func ringContextTestName(useMulti, useDeadline bool) string {
	name := "PutOne"
	if useMulti {
		name = "PutMulti"
	}
	if useDeadline {
		return name + "/DeadlineExceeded"
	}
	return name + "/Canceled"
}

func TestRingWaiterCancellationDoesNotBlockLaterRequests(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		r := newRing(1)
		for range len(r.store) {
			r.PutOne(context.Background(), cmds.PingCmd)
		}
		ctx, cancel := context.WithCancel(context.Background())
		canceled := make(chan error, 1)
		go func() {
			_, err := r.PutOne(ctx, cmds.PingCmd)
			canceled <- err
		}()
		synctest.Wait()

		live := make(chan chan RedisResult, 1)
		go func() {
			ch, err := r.PutOne(context.Background(), cmds.PingCmd)
			if err != nil {
				live <- nil
				return
			}
			live <- ch
		}()
		cancel()
		synctest.Wait()
		if err := <-canceled; !errors.Is(err, context.Canceled) {
			t.Fatal(err)
		}

		for range len(r.store) {
			r.NextWriteCmd()
			r.NextResultCh()
			r.FinishResult()
		}
		synctest.Wait()
		ch := <-live
		if ch == nil {
			t.Fatal("live request did not receive the released slot")
		}
		_, _, written := r.NextWriteCmd()
		_, _, read, _ := r.NextResultCh()
		r.FinishResult()
		if written != ch || read != ch {
			t.Fatal("live request was not paired after canceled waiter")
		}
	})
}

func TestRingCancellationDuringRelease(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		r := newRing(1)
		for range len(r.store) {
			r.PutOne(context.Background(), cmds.PingCmd)
		}
		const requests = 32
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		liveCtx, stop := context.WithCancel(context.Background())
		defer stop()
		done := make(chan error, requests)
		for i := range requests {
			go func() {
				requestCtx := liveCtx
				if i%2 == 0 {
					requestCtx = ctx
				}
				var err error
				if i%4 < 2 {
					_, err = r.PutOne(requestCtx, cmds.PingCmd)
				} else {
					_, err = r.PutMulti(requestCtx, []Completed{cmds.PingCmd}, nil)
				}
				done <- err
			}()
		}
		synctest.Wait()
		// 让取消与槽位释放竞争，确认剩余等待者仍能全部完成入队。
		go cancel()
		completed := 0
		for {
			synctest.Wait()
			for len(done) != 0 {
				if err := <-done; err != nil && !errors.Is(err, context.Canceled) {
					t.Fatal(err)
				}
				completed++
			}
			if completed == requests {
				break
			}
			_, _, written := r.NextWriteCmd()
			if written == nil {
				stop()
				synctest.Wait()
				t.Fatalf("%d requests completed, but the empty ring still has blocked waiters", completed)
			}
			_, _, response, _ := r.NextResultCh()
			r.FinishResult()
			if response != written {
				t.Fatal("response channel does not match the written command")
			}
		}
	})
}

func TestRingSequenceWraparound(t *testing.T) {
	r := newRing(1)
	// 从 uint32 边界前开始，验证容量计算和槽位索引一起回绕。
	r.write = ^uint32(0) - 1
	r.recycled, r.read1, r.read2 = r.write, r.write, r.write
	for range 3 {
		for range len(r.store) {
			if _, err := r.PutOne(context.Background(), cmds.PingCmd); err != nil {
				t.Fatal(err)
			}
		}
		for range len(r.store) {
			_, _, written := r.NextWriteCmd()
			_, _, response, _ := r.NextResultCh()
			r.FinishResult()
			if written == nil || response != written {
				t.Fatal("ring lost a command across sequence wraparound")
			}
		}
	}
}

func BenchmarkRingQueue(b *testing.B) {
	for _, cancellable := range []bool{false, true} {
		name := "Background"
		ctx := context.Background()
		if cancellable {
			name = "Cancellable"
			var cancel context.CancelFunc
			ctx, cancel = context.WithCancel(ctx)
			defer cancel()
		}
		b.Run(name+"/Sequential", func(b *testing.B) {
			r := newRing(DefaultRingScale)
			b.ReportAllocs()
			b.ResetTimer()
			for range b.N {
				if _, err := r.PutOne(ctx, cmds.PingCmd); err != nil {
					b.Fatal(err)
				}
				r.NextWriteCmd()
				r.NextResultCh()
				r.FinishResult()
			}
		})
		b.Run(name+"/Parallel", func(b *testing.B) {
			r := newRing(DefaultRingScale)
			done := make(chan struct{})
			go func() {
				defer close(done)
				for range b.N {
					r.WaitForWrite()
					_, _, ch, _ := r.NextResultCh()
					ch <- RedisResult{}
					r.FinishResult()
				}
			}()
			b.ReportAllocs()
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					ch, err := r.PutOne(ctx, cmds.PingCmd)
					if err != nil {
						b.Error(err)
						return
					}
					<-ch
				}
			})
			<-done
		})
	}
}
