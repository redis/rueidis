package singleflight

import (
	"sync"
)

type Call struct {
	mu sync.Mutex
	wg *sync.WaitGroup
}

func (c *Call) Do(fn func() error) error {
	var wg *sync.WaitGroup
	c.mu.Lock()
	if c.wg != nil {
		wg = c.wg
		c.mu.Unlock()
		wg.Wait()
		return nil
	}
	wg = &sync.WaitGroup{}
	wg.Add(1)
	c.wg = wg
	c.mu.Unlock()

	err := fn()
	c.mu.Lock()
	c.wg = nil
	c.mu.Unlock()
	wg.Done()
	return err
}
