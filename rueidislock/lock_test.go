package rueidislock

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/redis/rueidis"
)

const testDB = 10

var address = []string{"127.0.0.1:6379"}

func newLocker(t *testing.T, noLoop, setpx, nocsc bool) *locker {
	impl, err := NewLocker(LockerOption{
		ClientOption:   rueidis.ClientOption{InitAddress: address, DisableCache: nocsc, SelectDB: testDB},
		NoLoopTracking: noLoop,
		FallbackSETPX:  setpx,
	})
	if err != nil {
		t.Fatal(err)
	}
	return impl.(*locker)
}

func newClient(t *testing.T) rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: address, SelectDB: testDB})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestNewLocker(t *testing.T) {
	_, err := NewLocker(LockerOption{ClientOption: rueidis.ClientOption{InitAddress: nil, SelectDB: testDB}})
	if err == nil {
		t.Fatal(err)
	}
	l, err := NewLocker(LockerOption{ClientOption: rueidis.ClientOption{InitAddress: address, SelectDB: testDB}})
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	impl := l.(*locker)
	if impl.validity != 5*time.Second {
		t.Fatalf("unexpected default validity %v", impl.validity)
	}
	if impl.majority != 2 {
		t.Fatalf("unexpected default majority %v", impl.majority)
	}
	if impl.totalcnt != impl.majority*2-1 {
		t.Fatalf("unexpected default totalcnt %v", impl.totalcnt)
	}
}

func TestNewLocker_WithClientBuilder(t *testing.T) {
	var client rueidis.Client
	l, err := NewLocker(LockerOption{
		ClientOption: rueidis.ClientOption{InitAddress: address, SelectDB: testDB},
		ClientBuilder: func(option rueidis.ClientOption) (_ rueidis.Client, err error) {
			client, err = rueidis.NewClient(option)
			return client, err
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	if l.Client() != client {
		t.Fatal("client mismatched")
	}
}

func TestLocker_WithContext_MultipleLocker(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		lockers := make([]*locker, 10)
		sum := make([]int, len(lockers))
		for i := 0; i < len(lockers); i++ {
			lockers[i] = newLocker(t, noLoop, setpx, nocsc)
			lockers[i].timeout = time.Second
		}
		defer func() {
			for _, locker := range lockers {
				locker.Close()
			}
		}()
		cnt := 100
		lck := strconv.Itoa(rand.Int())
		ctx := context.Background()
		var wg sync.WaitGroup
		wg.Add(len(lockers))
		for i, l := range lockers {
			go func(i int, l *locker) {
				defer wg.Done()
				for j := 0; j < cnt; j++ {
					_, cancel, err := l.WithContext(ctx, lck)
					if err != nil {
						t.Error(err)
						return
					}
					sum[i]++
					cancel()
				}
			}(i, l)
		}
		wg.Wait()
		for i, s := range sum {
			if s != cnt {
				t.Fatalf("unexpected sum %v %v %v", i, s, cnt)
			}
		}
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_WithContext_UnlockByClientSideCaching(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx bool) {
		locker := newLocker(t, noLoop, setpx, false)
		locker.timeout = time.Second
		defer locker.Close()
		lck := strconv.Itoa(rand.Int())
		ctx, cancel, err := locker.WithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			client := newClient(t)
			defer client.Close()
			for i := int32(0); i < locker.majority; i++ {
				if err := client.Do(context.Background(), client.B().Del().Key(keyname(locker.prefix, lck, i)).Build()).Error(); err != nil {
					t.Error(err)
				}
			}
		}()
		<-ctx.Done()
		cancel()
		if !errors.Is(ctx.Err(), context.Canceled) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	t.Run("Tracking Loop", func(t *testing.T) {
		test(t, false, false)
	})
	t.Run("Tracking NoLoop", func(t *testing.T) {
		test(t, true, false)
	})
	t.Run("SET PX", func(t *testing.T) {
		test(t, true, true)
	})
}

func TestLocker_WithContext_UnlockBySelfForceWithContext(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx bool) {
		locker := newLocker(t, noLoop, setpx, false)
		locker.timeout = time.Second
		defer locker.Close()
		lck := strconv.Itoa(rand.Int())
		ctx, cancel, err := locker.WithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			_, cancel2, err2 := locker.ForceWithContext(context.Background(), lck)
			if err2 != nil {
				t.Errorf("unexpected err %v", err2)
				return
			}
			cancel2()
		}()
		<-ctx.Done()
		cancel()
		if !errors.Is(ctx.Err(), context.Canceled) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	t.Run("Tracking Loop", func(t *testing.T) {
		test(t, false, false)
	})
	t.Run("Tracking NoLoop", func(t *testing.T) {
		test(t, true, false)
	})
	t.Run("SET PX", func(t *testing.T) {
		test(t, true, true)
	})
}

func TestLocker_WithContext_UnlockByOtherForceWithContext(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx bool) {
		locker := newLocker(t, noLoop, setpx, false)
		locker.timeout = time.Second
		defer locker.Close()
		lck := strconv.Itoa(rand.Int())
		ctx, cancel, err := locker.WithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			locker2 := newLocker(t, noLoop, setpx, false)
			locker2.timeout = time.Second
			defer locker2.Close()
			_, cancel2, err2 := locker2.ForceWithContext(context.Background(), lck)
			if err2 != nil {
				t.Errorf("unexpected err %v", err2)
				return
			}
			cancel2()
		}()
		<-ctx.Done()
		cancel()
		if !errors.Is(ctx.Err(), context.Canceled) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	t.Run("Tracking Loop", func(t *testing.T) {
		test(t, false, false)
	})
	t.Run("Tracking NoLoop", func(t *testing.T) {
		test(t, true, false)
	})
	t.Run("SET PX", func(t *testing.T) {
		test(t, true, true)
	})
}

func TestLocker_ForceWithContext_UnlockBySelfForceWithContext(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx bool) {
		locker := newLocker(t, noLoop, setpx, false)
		locker.timeout = time.Second
		defer locker.Close()
		lck := strconv.Itoa(rand.Int())
		ctx, cancel, err := locker.ForceWithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			_, cancel2, err2 := locker.ForceWithContext(context.Background(), lck)
			if err2 != nil {
				t.Errorf("unexpected err %v", err2)
				return
			}
			cancel2()
		}()
		<-ctx.Done()
		cancel()
		if !errors.Is(ctx.Err(), context.Canceled) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	t.Run("Tracking Loop", func(t *testing.T) {
		test(t, false, false)
	})
	t.Run("Tracking NoLoop", func(t *testing.T) {
		test(t, true, false)
	})
	t.Run("SET PX", func(t *testing.T) {
		test(t, true, true)
	})
}

func TestLocker_ForceWithContext_UnlockByOtherForceWithContext(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx bool) {
		locker := newLocker(t, noLoop, setpx, false)
		locker.timeout = time.Second
		defer locker.Close()
		lck := strconv.Itoa(rand.Int())
		ctx, cancel, err := locker.ForceWithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			locker2 := newLocker(t, noLoop, setpx, false)
			locker2.timeout = time.Second
			defer locker2.Close()
			_, cancel2, err2 := locker2.ForceWithContext(context.Background(), lck)
			if err2 != nil {
				t.Errorf("unexpected err %v", err2)
				return
			}
			cancel2()
		}()
		<-ctx.Done()
		cancel()
		if !errors.Is(ctx.Err(), context.Canceled) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	t.Run("Tracking Loop", func(t *testing.T) {
		test(t, false, false)
	})
	t.Run("Tracking NoLoop", func(t *testing.T) {
		test(t, true, false)
	})
	t.Run("SET PX", func(t *testing.T) {
		test(t, true, true)
	})
}

func TestLocker_WithContext_ExtendByClientSideCaching(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx bool) {
		locker := newLocker(t, noLoop, setpx, false)
		locker.timeout = time.Second
		defer locker.Close()
		lck := strconv.Itoa(rand.Int())
		ctx, cancel, err := locker.WithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			client := newClient(t)
			defer client.Close()
			for i := int32(0); i < locker.majority; i++ {
				if err := client.Do(context.Background(), client.B().Pexpire().Key(keyname(locker.prefix, lck, i)).Milliseconds(100).Build()).Error(); err != nil {
					t.Error(err)
				}
			}
		}()
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			t.Fatalf("unexpected err %v", ctx.Err())
		default:
		}
		cancel()
	}
	t.Run("Tracking Loop", func(t *testing.T) {
		test(t, false, false)
	})
	t.Run("Tracking NoLoop", func(t *testing.T) {
		test(t, true, false)
	})
	t.Run("SET PX", func(t *testing.T) {
		test(t, true, true)
	})
}

func TestLocker_WithContext_AutoExtendConcurrent(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)
		locker.validity = time.Second
		locker.interval = time.Second / 2
		defer locker.Close()

		key := strconv.Itoa(rand.Int())

		ctx1, cancel1, err1 := locker.WithContext(context.Background(), key)
		if err1 != nil {
			t.Fatal(err1)
		}
		go func() {
			for i := 0; i < 4; i++ {
				select {
				case <-ctx1.Done():
					t.Errorf("unexpected context canceled %v", ctx1.Err())
				default:
					time.Sleep(locker.validity)
				}
			}
			cancel1()
		}()
		ctx2, cancel2, err2 := locker.WithContext(context.Background(), key)
		if err2 != nil {
			t.Fatal(err2)
		}
		if !errors.Is(ctx1.Err(), context.Canceled) {
			t.Fatalf("unexpected context canceled %v", ctx1.Err())
		}
		if ctx2.Err() != nil {
			t.Fatalf("unexpected context canceled %v", ctx2.Err())
		}
		cancel2()
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_WithContext_AutoExtend(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)
		locker.validity = time.Second * 2
		locker.interval = time.Second
		defer locker.Close()

		ctx, cancel, err := locker.WithContext(context.Background(), strconv.Itoa(rand.Int()))
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 2; i++ {
			select {
			case <-ctx.Done():
				t.Fatalf("unexpected context canceled %v", ctx.Err())
			default:
				time.Sleep(locker.validity)
			}
		}
		if ctx.Err() != nil {
			t.Fatalf("unexpected context canceled %v", ctx.Err())
		}
		cancel()
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_WithContext_DeadContext(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)
		defer locker.Close()

		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, _, err := locker.WithContext(ctx, strconv.Itoa(rand.Int()))
		if !errors.Is(err, context.Canceled) {
			t.Fatal(err)
		}
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_WithContext_CancelContext(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)
		defer locker.Close()

		lck := strconv.Itoa(rand.Int())
		ctx, cancel, err := locker.WithContext(context.Background(), lck)

		wg := sync.WaitGroup{}
		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func() {
				if _, _, err := locker.WithContext(ctx, lck); !errors.Is(err, context.Canceled) {
					t.Error(err)
				}
				wg.Done()
			}()
		}
		time.Sleep(time.Second * 3)
		cancel()
		wg.Wait()
		if !errors.Is(ctx.Err(), context.Canceled) {
			t.Fatal(err)
		}
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_WithContext_ShorterTimeoutContext(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)
		locker.validity = time.Second * 5
		locker.interval = time.Second * 3
		defer locker.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		ctx, cancel, err := locker.WithContext(ctx, strconv.Itoa(rand.Int()))
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second * 2)
		if !errors.Is(ctx.Err(), context.DeadlineExceeded) {
			t.Fatalf("unexpected context canceled %v", ctx.Err())
		}
		cancel()
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_TryWithContext(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)
		locker.timeout = time.Second
		defer locker.Close()

		lck := strconv.Itoa(rand.Int())
		ctx, cancel, err := locker.TryWithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}
		if _, _, err := locker.TryWithContext(ctx, lck); !errors.Is(err, ErrNotLocked) {
			t.Fatal(err)
		}
		cancel()
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_ForceWithContextThenTryWithContext(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)
		locker.timeout = time.Second
		defer locker.Close()

		lck := strconv.Itoa(rand.Int())
		ctx, cancel, err := locker.ForceWithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}
		if _, _, err := locker.TryWithContext(ctx, lck); !errors.Is(err, ErrNotLocked) {
			t.Fatal(err)
		}
		cancel()
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_TryWithContext_MultipleLocker(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		lockers := make([]*locker, 10)
		sum := make([]int, len(lockers))
		for i := 0; i < len(lockers); i++ {
			lockers[i] = newLocker(t, noLoop, setpx, nocsc)
			lockers[i].timeout = time.Second
		}
		defer func() {
			for _, locker := range lockers {
				locker.Close()
			}
		}()
		cnt := 1000
		lck := strconv.Itoa(rand.Int())
		ctx := context.Background()
		var wg sync.WaitGroup
		wg.Add(len(lockers))
		for i, l := range lockers {
			go func(i int, l *locker) {
				defer wg.Done()
				for j := 0; j < cnt; j++ {
					for {
						_, cancel, err := l.TryWithContext(ctx, lck)
						if err != nil && !errors.Is(err, ErrNotLocked) {
							t.Error(err)
							return
						}
						if cancel != nil {
							cancel()
							sum[i]++
							break
						}
						time.Sleep(time.Millisecond)
					}
				}
			}(i, l)
		}
		wg.Wait()
		for i, s := range sum {
			if s != cnt {
				t.Fatalf("unexpected sum %v %v %v", i, s, cnt)
			}
		}
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_WithContext_MissingPeers(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)
		defer locker.Close()

		lck := strconv.Itoa(rand.Int())
		for i := 0; i < 100; i++ {
			ctx, cancel, err := locker.WithContext(context.Background(), lck)
			if err != nil {
				t.Fatal(err)
			}
			locker.mu.RLock()
			g := locker.gates[lck]
			locker.mu.RUnlock()
			var wg sync.WaitGroup
			for i := 0; i < 100; i++ {
				wg.Add(1)
				go func() {
					if _, _, err := locker.WithContext(ctx, lck); !errors.Is(err, context.Canceled) {
						t.Error(err)
					}
					wg.Done()
				}()
			}
			cancel()
			select {
			case g.ch <- struct{}{}:
			default:
			}
			wg.Wait()
		}
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_WithContext_Cleanup(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)
		defer locker.Close()

		lck := strconv.Itoa(rand.Int())
		_, cancel1, err := locker.WithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}
		locker.mu.Lock()
		locker.gates[lck].w--
		locker.mu.Unlock()
		ctx2, cancel2 := context.WithCancel(context.Background())
		go func() {
			time.Sleep(time.Second)
			cancel2()
			cancel1()
		}()
		if _, _, err := locker.WithContext(ctx2, lck); !errors.Is(err, context.Canceled) {
			t.Fatal(err)
		}
		locker.mu.Lock()
		keys := len(locker.gates)
		locker.mu.Unlock()
		if keys != 0 {
			t.Fatalf("unexpected length %v", keys)
		}
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_Close(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)

		lck := strconv.Itoa(rand.Int())
		ctx, _, err := locker.WithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}

		wg := sync.WaitGroup{}
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() {
				if _, _, err := locker.WithContext(context.Background(), lck); !errors.Is(err, ErrLockerClosed) {
					t.Error(err)
				}
				wg.Done()
			}()
		}
		time.Sleep(time.Second)
		locker.Close()
		wg.Wait()
		<-ctx.Done()
		if err := ctx.Err(); !errors.Is(err, context.Canceled) {
			t.Fatal(err)
		}
		if _, _, err := locker.WithContext(context.Background(), lck); !errors.Is(err, ErrLockerClosed) {
			t.Fatal(err)
		}
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_RetryErrLockerClosed(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		locker := newLocker(t, noLoop, setpx, nocsc)

		lck := strconv.Itoa(rand.Int())
		ctx, cancel, err := locker.WithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			_, cancel, err := locker.WithContext(context.Background(), lck)
			if err != nil {
				t.Error(err)
			}
			wg.Done()
			cancel()
		}()
		go func() {
			for {
				time.Sleep(time.Second)
				locker.onInvalidations(nil) // create ErrLockerClosed
				cancel()
			}
		}()
		wg.Wait()
		if err := ctx.Err(); !errors.Is(err, context.Canceled) {
			t.Fatal(err)
		}
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}

func TestLocker_Flush(t *testing.T) {
	test := func(t *testing.T, noLoop, setpx, nocsc bool) {
		client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: address, SelectDB: testDB})
		if err != nil {
			t.Fatal(err)
		}
		defer client.Close()

		locker := newLocker(t, noLoop, setpx, nocsc)

		lck := strconv.Itoa(rand.Int())
		ctx, _, err := locker.WithContext(context.Background(), lck)
		if err != nil {
			t.Fatal(err)
		}

		if err := client.Do(context.Background(), client.B().Flushdb().Build()).Error(); err != nil {
			t.Fatal(err)
		}

		<-ctx.Done()

		if err := ctx.Err(); !errors.Is(err, context.Canceled) {
			t.Fatal(err)
		}

		_, cancel, err := locker.WithContext(context.Background(), strconv.Itoa(rand.Int()))
		if err != nil {
			t.Fatal(err)
		}
		cancel()
	}
	for _, nocsc := range []bool{false, true} {
		t.Run("Tracking Loop", func(t *testing.T) {
			test(t, false, false, nocsc)
		})
		t.Run("Tracking NoLoop", func(t *testing.T) {
			test(t, true, false, nocsc)
		})
		t.Run("SET PX", func(t *testing.T) {
			test(t, true, true, nocsc)
		})
	}
}
