package rueidishook

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"go.uber.org/mock/gomock"
)

type hook struct{}

func (h *hook) Do(client rueidis.Client, ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	return client.Do(ctx, cmd)
}

func (h *hook) DoMulti(client rueidis.Client, ctx context.Context, multi ...rueidis.Completed) (resps []rueidis.RedisResult) {
	return client.DoMulti(ctx, multi...)
}

func (h *hook) DoCache(client rueidis.Client, ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	return client.DoCache(ctx, cmd, ttl)
}

func (h *hook) DoMultiCache(client rueidis.Client, ctx context.Context, multi ...rueidis.CacheableTTL) (resps []rueidis.RedisResult) {
	return client.DoMultiCache(ctx, multi...)
}

func (h *hook) Receive(client rueidis.Client, ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	return client.Receive(ctx, subscribe, fn)
}

func (h *hook) DoStream(client rueidis.Client, ctx context.Context, cmd rueidis.Completed) rueidis.RedisResultStream {
	return client.DoStream(ctx, cmd)
}

func (h *hook) DoMultiStream(client rueidis.Client, ctx context.Context, multi ...rueidis.Completed) rueidis.MultiRedisResultStream {
	return client.DoMultiStream(ctx, multi...)
}

type wronghook struct {
	DoFn func(client rueidis.Client)
}

func (w *wronghook) Do(client rueidis.Client, ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	w.DoFn(client)
	return rueidis.RedisResult{}
}

func (w *wronghook) DoMulti(client rueidis.Client, ctx context.Context, multi ...rueidis.Completed) (resps []rueidis.RedisResult) {
	panic("implement me")
}

func (w *wronghook) DoCache(client rueidis.Client, ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	panic("implement me")
}

func (w *wronghook) DoMultiCache(client rueidis.Client, ctx context.Context, multi ...rueidis.CacheableTTL) (resps []rueidis.RedisResult) {
	panic("implement me")
}

func (w *wronghook) Receive(client rueidis.Client, ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	panic("implement me")
}

func (w *wronghook) DoStream(client rueidis.Client, ctx context.Context, cmd rueidis.Completed) rueidis.RedisResultStream {
	panic("implement me")
}

func (w *wronghook) DoMultiStream(client rueidis.Client, ctx context.Context, multi ...rueidis.Completed) rueidis.MultiRedisResultStream {
	panic("implement me")
}

func testHooked(t *testing.T, hooked rueidis.Client, mocked *mock.Client) {
	ctx := context.Background()
	{
		mocked.EXPECT().Mode().Return(rueidis.ClientModeStandalone)
		if mode := hooked.Mode(); mode != rueidis.ClientModeStandalone {
			t.Fatalf("unexpected mode %v", mode)
		}
	}
	{
		mocked.EXPECT().Do(ctx, mock.Match("GET", "a")).Return(mock.Result(mock.RedisNil()))
		if err := hooked.Do(ctx, hooked.B().Get().Key("a").Build()).Error(); !rueidis.IsRedisNil(err) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		mocked.EXPECT().DoCache(ctx, mock.Match("GET", "b"), time.Second).Return(mock.Result(mock.RedisNil()))
		if err := hooked.DoCache(ctx, hooked.B().Get().Key("b").Cache(), time.Second).Error(); !rueidis.IsRedisNil(err) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		mocked.EXPECT().DoMulti(ctx, mock.Match("GET", "c")).Return([]rueidis.RedisResult{mock.Result(mock.RedisNil())})
		for _, resp := range hooked.DoMulti(ctx, hooked.B().Get().Key("c").Build()) {
			if err := resp.Error(); !rueidis.IsRedisNil(err) {
				t.Fatalf("unexpected err %v", err)
			}
		}
	}
	{
		mocked.EXPECT().DoMultiCache(ctx, mock.Match("GET", "e")).Return([]rueidis.RedisResult{mock.Result(mock.RedisNil())})
		for _, resp := range hooked.DoMultiCache(ctx, rueidis.CT(hooked.B().Get().Key("e").Cache(), time.Second)) {
			if err := resp.Error(); !rueidis.IsRedisNil(err) {
				t.Fatalf("unexpected err %v", err)
			}
		}
	}
	{
		mocked.EXPECT().DoStream(ctx, mock.Match("GET", "e")).Return(mock.RedisResultStream(mock.RedisNil()))
		s := hooked.DoStream(ctx, hooked.B().Get().Key("e").Build())
		if _, err := s.WriteTo(io.Discard); !rueidis.IsRedisNil(err) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		mocked.EXPECT().DoMultiStream(ctx, mock.Match("GET", "e"), mock.Match("GET", "f")).Return(mock.MultiRedisResultStream(mock.RedisNil()))
		s := hooked.DoMultiStream(ctx, hooked.B().Get().Key("e").Build(), hooked.B().Get().Key("f").Build())
		if _, err := s.WriteTo(io.Discard); !rueidis.IsRedisNil(err) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		mocked.EXPECT().Receive(ctx, mock.Match("SUBSCRIBE", "a"), gomock.Any()).DoAndReturn(func(ctx context.Context, cmd any, fn func(msg rueidis.PubSubMessage)) error {
			fn(rueidis.PubSubMessage{
				Channel: "s",
				Message: "s",
			})
			return errors.New("any")
		})
		if err := hooked.Receive(ctx, hooked.B().Subscribe().Channel("a").Build(), func(msg rueidis.PubSubMessage) {
			if msg.Message != "s" && msg.Channel != "s" {
				t.Fatalf("unexpected val %v", msg)
			}
		}); err.Error() != "any" {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		mocked.EXPECT().Nodes().Return(map[string]rueidis.Client{"addr": mocked})
		if nodes := hooked.Nodes(); nodes["addr"].(*hookclient).client != mocked {
			t.Fatalf("unexpected val %v", nodes)
		}
	}
	{
		ch := make(chan struct{})
		mocked.EXPECT().Close().Do(func() { close(ch) })
		hooked.Close()
		<-ch
	}
}

func testHookedDedicated(t *testing.T, hooked rueidis.DedicatedClient, mocked *mock.DedicatedClient) {
	ctx := context.Background()
	{
		mocked.EXPECT().Do(ctx, mock.Match("GET", "a")).Return(mock.Result(mock.RedisNil()))
		if err := hooked.Do(ctx, hooked.B().Get().Key("a").Build()).Error(); !rueidis.IsRedisNil(err) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		mocked.EXPECT().DoMulti(ctx, mock.Match("GET", "c")).Return([]rueidis.RedisResult{mock.Result(mock.RedisNil())})
		for _, resp := range hooked.DoMulti(ctx, hooked.B().Get().Key("c").Build()) {
			if err := resp.Error(); !rueidis.IsRedisNil(err) {
				t.Fatalf("unexpected err %v", err)
			}
		}
	}
	{
		mocked.EXPECT().Receive(ctx, mock.Match("SUBSCRIBE", "a"), gomock.Any()).DoAndReturn(func(ctx context.Context, cmd any, fn func(msg rueidis.PubSubMessage)) error {
			fn(rueidis.PubSubMessage{
				Channel: "s",
				Message: "s",
			})
			return errors.New("any")
		})
		if err := hooked.Receive(ctx, hooked.B().Subscribe().Channel("a").Build(), func(msg rueidis.PubSubMessage) {
			if msg.Message != "s" && msg.Channel != "s" {
				t.Fatalf("unexpected val %v", msg)
			}
		}); err.Error() != "any" {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		mocked.EXPECT().SetPubSubHooks(rueidis.PubSubHooks{})
		hooked.SetPubSubHooks(rueidis.PubSubHooks{})
	}
	{
		ch := make(chan struct{})
		mocked.EXPECT().Close().Do(func() { close(ch) })
		hooked.Close()
		<-ch
	}
}

func TestWithHook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocked := mock.NewClient(ctrl)
	hooked := WithHook(mocked, &hook{})

	testHooked(t, hooked, mocked)
	{
		dc := mock.NewDedicatedClient(ctrl)
		mocked.EXPECT().Dedicate().Return(dc, func() {})
		c, _ := hooked.Dedicate()
		testHookedDedicated(t, c, dc)
	}
	{
		dc := mock.NewDedicatedClient(ctrl)
		cb := func(c rueidis.DedicatedClient) error {
			testHookedDedicated(t, c, dc)
			return errors.New("any")
		}
		mocked.EXPECT().Dedicated(gomock.Any()).DoAndReturn(func(fn func(c rueidis.DedicatedClient) error) error {
			return fn(dc)
		})
		if err := hooked.Dedicated(cb); err.Error() != "any" {
			t.Fatalf("unexpected err %v", err)
		}
	}
}

func TestForbiddenMethodForDedicatedClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocked := mock.NewClient(ctrl)

	shouldpanic := func(fn func(client rueidis.Client), msg string) {
		defer func() {
			if err := recover().(string); err != msg {
				t.Fatalf("unexpected err %v", err)
			}
		}()

		hooked := WithHook(mocked, &wronghook{DoFn: fn})
		mocked.EXPECT().Dedicated(gomock.Any()).DoAndReturn(func(fn func(c rueidis.DedicatedClient) error) error {
			return fn(mock.NewDedicatedClient(ctrl))
		})
		hooked.Dedicated(func(client rueidis.DedicatedClient) error {
			return client.Do(context.Background(), client.B().Get().Key("").Build()).Error()
		})
	}
	for _, c := range []struct {
		fn  func(client rueidis.Client)
		msg string
	}{
		{
			fn: func(client rueidis.Client) {
				client.Mode()
			},
			msg: "Mode() is not allowed with rueidis.DedicatedClient",
		},
		{
			fn: func(client rueidis.Client) {
				client.DoCache(context.Background(), client.B().Get().Key("").Cache(), time.Second)
			},
			msg: "DoCache() is not allowed with rueidis.DedicatedClient",
		}, {
			fn: func(client rueidis.Client) {
				client.DoMultiCache(context.Background())
			},
			msg: "DoMultiCache() is not allowed with rueidis.DedicatedClient",
		}, {
			fn: func(client rueidis.Client) {
				client.Dedicated(func(client rueidis.DedicatedClient) error { return nil })
			},
			msg: "Dedicated() is not allowed with rueidis.DedicatedClient",
		}, {
			fn: func(client rueidis.Client) {
				client.Dedicate()
			},
			msg: "Dedicate() is not allowed with rueidis.DedicatedClient",
		}, {
			fn: func(client rueidis.Client) {
				client.Nodes()
			},
			msg: "Nodes() is not allowed with rueidis.DedicatedClient",
		}, {
			fn: func(client rueidis.Client) {
				client.DoStream(context.Background(), client.B().Get().Key("").Build())
			},
			msg: "DoStream() is not allowed with rueidis.DedicatedClient",
		}, {
			fn: func(client rueidis.Client) {
				client.DoMultiStream(context.Background(), client.B().Get().Key("").Build())
			},
			msg: "DoMultiStream() is not allowed with rueidis.DedicatedClient",
		},
	} {
		shouldpanic(c.fn, c.msg)
	}
}

func TestNewErrorResult(t *testing.T) {
	e := errors.New("err")
	r := NewErrorResult(e)
	if r.Error() != e {
		t.Fatal("unexpected err")
	}
}

func TestNewErrorResultStream(t *testing.T) {
	e := errors.New("err")
	r := NewErrorResultStream(e)
	if r.Error() != e {
		t.Fatal("unexpected err")
	}
	if r.HasNext() {
		t.Fatal("unexpected err")
	}
	if n, err := r.WriteTo(io.Discard); err != e || n != 0 {
		t.Fatal("unexpected err or n")
	}
}
