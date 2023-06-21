package util

import (
	"sync"
	"sync/atomic"
)

func ParallelKeys[K comparable, V any](p map[K]V, fn func(k K)) {
	ch := make(chan K, len(p))
	for k := range p {
		ch <- k
	}
	closeThenParallel(ch, fn)
}

func ParallelVals[K comparable, V any](p map[K]V, fn func(k V)) {
	ch := make(chan V, len(p))
	for _, v := range p {
		ch <- v
	}
	closeThenParallel(ch, fn)
}

func ParallelArrI[V any](p []V, fn func(index int)) {
	index := int32(-1)
	wg := sync.WaitGroup{}
	wg.Add(len(p))
	for i := 1; i < len(p); i++ {
		go func() {
			for i := int(atomic.AddInt32(&index, 1)); i < len(p); i = int(atomic.AddInt32(&index, 1)) {
				fn(i)
			}
			wg.Done()
		}()
	}
	for i := int(atomic.AddInt32(&index, 1)); i < len(p); i = int(atomic.AddInt32(&index, 1)) {
		fn(i)
	}
	wg.Done()
	wg.Wait()
}

func closeThenParallel[V any](ch chan V, fn func(k V)) {
	close(ch)
	concurrency := len(ch)
	// TODO runtime.GOMAXPROCS(0) is heavy, we should avoid doing it every time
	//if cpus := runtime.GOMAXPROCS(0); concurrency > cpus {
	//	concurrency = cpus
	//}
	wg := sync.WaitGroup{}
	wg.Add(concurrency)
	for i := 1; i < concurrency; i++ {
		go func() {
			for v := range ch {
				fn(v)
			}
			wg.Done()
		}()
	}
	for v := range ch {
		fn(v)
	}
	wg.Done()
	wg.Wait()
}
