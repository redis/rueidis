package rueidislock

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/rueian/rueidis"
)

var address = []string{"127.0.0.1:6376"}

func newLocker(t *testing.T) *locker {
	impl, err := NewLocker(LockerOption{
		ClientOption: rueidis.ClientOption{InitAddress: address},
	})
	if err != nil {
		t.Fatal(err)
	}
	return impl.(*locker)
}

func newClient(t *testing.T) rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: address})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestNewLocker(t *testing.T) {
	l, err := NewLocker(LockerOption{ClientOption: rueidis.ClientOption{InitAddress: nil}})
	if err == nil {
		t.Fatal(err)
	}
	l, err = NewLocker(LockerOption{ClientOption: rueidis.ClientOption{InitAddress: address}})
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
	l, err := NewLocker(LockerOption{
		ClientOption: rueidis.ClientOption{InitAddress: address},
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			return rueidis.NewClient(option)
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
}

func TestLocker_WithContext_MultipleLocker(t *testing.T) {
	lockers := make([]*locker, 10)
	sum := make([]int, len(lockers))
	for i := 0; i < len(lockers); i++ {
		lockers[i] = newLocker(t)
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

func TestLocker_WithContext_UnlockByClientSideCaching(t *testing.T) {
	locker := newLocker(t)
	defer locker.Close()
	lck := strconv.Itoa(rand.Int())
	ctx, cancel, err := locker.WithContext(context.Background(), lck)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		client := newClient(t)
		defer client.Close()
		for i := 0; i < locker.majority; i++ {
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

func TestLocker_WithContext_AutoExtend(t *testing.T) {
	locker := newLocker(t)
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

func TestLocker_WithContext_DeadContext(t *testing.T) {
	locker := newLocker(t)
	defer locker.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	ctx, cancel, err := locker.WithContext(ctx, strconv.Itoa(rand.Int()))
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}

func TestLocker_WithContext_CancelContext(t *testing.T) {
	locker := newLocker(t)
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

func TestLocker_TryWithContext(t *testing.T) {
	locker := newLocker(t)
	defer locker.Close()

	lck := strconv.Itoa(rand.Int())
	ctx, cancel, err := locker.TryWithContext(context.Background(), lck)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := locker.TryWithContext(ctx, lck); err == nil {
		t.Fatal(err)
	}
	cancel()
}

func TestLocker_Close(t *testing.T) {
	locker := newLocker(t)

	lck := strconv.Itoa(rand.Int())
	ctx, _, err := locker.WithContext(context.Background(), lck)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			if _, _, err := locker.WithContext(context.Background(), lck); err != ErrLockerClosed {
				t.Error(err)
			}
			wg.Done()
		}()
	}
	time.Sleep(time.Second)
	locker.Close()
	wg.Wait()
	if !errors.Is(ctx.Err(), context.Canceled) {
		t.Fatal(err)
	}
}
