package util

import (
	"runtime"
	"sync/atomic"
	"testing"
)

func TestParallelKeys(t *testing.T) {
	var sum int64
	data, sk, _ := gen(int64(runtime.NumCPU() * 1000))
	ParallelKeys(runtime.GOMAXPROCS(0), data, func(i int64) {
		atomic.AddInt64(&sum, i)
	})
	if atomic.LoadInt64(&sum) != sk {
		t.Fatalf("unexpected")
	}
}

func TestParallelVals(t *testing.T) {
	var sum int64
	data, _, sv := gen(int64(runtime.NumCPU() * 1000))
	ParallelVals(runtime.GOMAXPROCS(0), data, func(i int64) {
		atomic.AddInt64(&sum, i)
	})
	if atomic.LoadInt64(&sum) != sv {
		t.Fatalf("unexpected")
	}
}

func gen(count int64) (data map[int64]int64, sk, sv int64) {
	data = make(map[int64]int64, count)
	for i := int64(1); i <= count; i++ {
		sk += i
		sv += i * -1
		data[i] = i * -1
	}
	return
}
