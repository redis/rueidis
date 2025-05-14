package rueidishook

import (
	"context"
	"time"
	"unsafe"

	"github.com/redis/rueidis"
)

var _ rueidis.Client = (*hookclient)(nil)

// Hook allows user to intercept rueidis.Client by using WithHook
type Hook interface {
	Do(client rueidis.Client, ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult)
	DoMulti(client rueidis.Client, ctx context.Context, multi ...rueidis.Completed) (resps []rueidis.RedisResult)
	DoCache(client rueidis.Client, ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult)
	DoMultiCache(client rueidis.Client, ctx context.Context, multi ...rueidis.CacheableTTL) (resps []rueidis.RedisResult)
	Receive(client rueidis.Client, ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error)
	DoStream(client rueidis.Client, ctx context.Context, cmd rueidis.Completed) rueidis.RedisResultStream
	DoMultiStream(client rueidis.Client, ctx context.Context, multi ...rueidis.Completed) rueidis.MultiRedisResultStream
}

// WithHook wraps rueidis.Client with Hook and allows the user to intercept rueidis.Client
func WithHook(client rueidis.Client, hook Hook) rueidis.Client {
	return &hookclient{client: client, hook: hook}
}

type hookclient struct {
	client rueidis.Client
	hook   Hook
}

func (c *hookclient) B() rueidis.Builder {
	return c.client.B()
}

func (c *hookclient) Do(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	return c.hook.Do(c.client, ctx, cmd)
}

func (c *hookclient) DoMulti(ctx context.Context, multi ...rueidis.Completed) (resp []rueidis.RedisResult) {
	return c.hook.DoMulti(c.client, ctx, multi...)
}

func (c *hookclient) DoCache(ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	return c.hook.DoCache(c.client, ctx, cmd, ttl)
}

func (c *hookclient) DoMultiCache(ctx context.Context, multi ...rueidis.CacheableTTL) (resps []rueidis.RedisResult) {
	return c.hook.DoMultiCache(c.client, ctx, multi...)
}

func (c *hookclient) DoStream(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResultStream {
	return c.hook.DoStream(c.client, ctx, cmd)
}

func (c *hookclient) DoMultiStream(ctx context.Context, multi ...rueidis.Completed) rueidis.MultiRedisResultStream {
	return c.hook.DoMultiStream(c.client, ctx, multi...)
}

func (c *hookclient) Dedicated(fn func(rueidis.DedicatedClient) error) (err error) {
	return c.client.Dedicated(func(client rueidis.DedicatedClient) error {
		return fn(&dedicated{client: &extended{DedicatedClient: client}, hook: c.hook})
	})
}

func (c *hookclient) Dedicate() (rueidis.DedicatedClient, func()) {
	client, cancel := c.client.Dedicate()
	return &dedicated{client: &extended{DedicatedClient: client}, hook: c.hook}, cancel
}

func (c *hookclient) Receive(ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	return c.hook.Receive(c.client, ctx, subscribe, fn)
}

func (c *hookclient) Nodes() map[string]rueidis.Client {
	nodes := c.client.Nodes()
	for addr, client := range nodes {
		nodes[addr] = &hookclient{client: client, hook: c.hook}
	}
	return nodes
}

func (c *hookclient) Mode() rueidis.ClientMode {
	return c.client.Mode()
}

func (c *hookclient) Close() {
	c.client.Close()
}

var _ rueidis.DedicatedClient = (*dedicated)(nil)

type dedicated struct {
	client *extended
	hook   Hook
}

func (d *dedicated) B() rueidis.Builder {
	return d.client.B()
}

func (d *dedicated) Do(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	return d.hook.Do(d.client, ctx, cmd)
}

func (d *dedicated) DoMulti(ctx context.Context, multi ...rueidis.Completed) (resp []rueidis.RedisResult) {
	return d.hook.DoMulti(d.client, ctx, multi...)
}

func (d *dedicated) Receive(ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	return d.hook.Receive(d.client, ctx, subscribe, fn)
}

func (d *dedicated) SetPubSubHooks(hooks rueidis.PubSubHooks) <-chan error {
	return d.client.SetPubSubHooks(hooks)
}

func (d *dedicated) Close() {
	d.client.Close()
}

var _ rueidis.Client = (*extended)(nil)

type extended struct {
	rueidis.DedicatedClient
}

func (e *extended) DoCache(ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	panic("DoCache() is not allowed with rueidis.DedicatedClient")
}

func (e *extended) DoMultiCache(ctx context.Context, multi ...rueidis.CacheableTTL) (resp []rueidis.RedisResult) {
	panic("DoMultiCache() is not allowed with rueidis.DedicatedClient")
}
func (c *extended) DoStream(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResultStream {
	panic("DoStream() is not allowed with rueidis.DedicatedClient")
}

func (c *extended) DoMultiStream(ctx context.Context, multi ...rueidis.Completed) rueidis.MultiRedisResultStream {
	panic("DoMultiStream() is not allowed with rueidis.DedicatedClient")
}

func (e *extended) Dedicated(fn func(rueidis.DedicatedClient) error) (err error) {
	panic("Dedicated() is not allowed with rueidis.DedicatedClient")
}

func (e *extended) Dedicate() (client rueidis.DedicatedClient, cancel func()) {
	panic("Dedicate() is not allowed with rueidis.DedicatedClient")
}

func (e *extended) Nodes() map[string]rueidis.Client {
	panic("Nodes() is not allowed with rueidis.DedicatedClient")
}

func (e *extended) Mode() rueidis.ClientMode {
	panic("Mode() is not allowed with rueidis.DedicatedClient")
}

type result struct {
	err error
	val rueidis.RedisMessage
}

func NewErrorResult(err error) rueidis.RedisResult {
	r := result{err: err}
	return *(*rueidis.RedisResult)(unsafe.Pointer(&r))
}

type stream struct {
	p *int
	w *int
	e error
	n int
}

func NewErrorResultStream(err error) rueidis.RedisResultStream {
	r := stream{e: err}
	return *(*rueidis.RedisResultStream)(unsafe.Pointer(&r))
}
