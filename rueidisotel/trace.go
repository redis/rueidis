package rueidisotel

import (
	"context"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/trace"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/cmds"
)

var (
	name   = "github.com/redis/rueidis"
	kind   = trace.WithSpanKind(trace.SpanKindClient)
	tracer = otel.Tracer(name)
	meter  = global.Meter(name)
	dbattr = attribute.String("db.system", "redis")

	cscMiss, _ = meter.Int64Counter("rueidis_do_cache_miss")
	cscHits, _ = meter.Int64Counter("rueidis_do_cache_hits")
)

var _ rueidis.Client = (*otelclient)(nil)

// WithClient creates a new rueidis.Client with OpenTelemetry tracing enabled
func WithClient(client rueidis.Client, opts ...Option) rueidis.Client {
	o := &otelclient{client: client}
	for _, fn := range opts {
		fn(o)
	}
	return o
}

// Option is the Functional Options interface
type Option func(o *otelclient)

// MetricAttrs set additional attributes to append to each metric.
func MetricAttrs(attrs ...attribute.KeyValue) Option {
	return func(o *otelclient) {
		o.mAttrs = attrs
	}
}

// TraceAttrs set additional attributes to append to each trace.
func TraceAttrs(attrs ...attribute.KeyValue) Option {
	return func(o *otelclient) {
		o.tAttrs = attrs
	}
}

type otelclient struct {
	client rueidis.Client
	mAttrs []attribute.KeyValue
	tAttrs []attribute.KeyValue
}

func (o *otelclient) B() cmds.Builder {
	return o.client.B()
}

func (o *otelclient) Do(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	ctx, span := start(ctx, first(cmd.Commands()), sum(cmd.Commands()), o.tAttrs)
	resp = o.client.Do(ctx, cmd)
	end(span, resp.Error())
	return
}

func (o *otelclient) DoMulti(ctx context.Context, multi ...rueidis.Completed) (resp []rueidis.RedisResult) {
	ctx, span := start(ctx, multiFirst(multi), multiSum(multi), o.tAttrs)
	resp = o.client.DoMulti(ctx, multi...)
	end(span, firstError(resp))
	return
}

func (o *otelclient) DoCache(ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	ctx, span := start(ctx, first(cmd.Commands()), sum(cmd.Commands()), o.tAttrs)
	resp = o.client.DoCache(ctx, cmd, ttl)
	if resp.NonRedisError() == nil {
		if resp.IsCacheHit() {
			cscHits.Add(ctx, 1, metric.WithAttributes(o.tAttrs...))
		} else {
			cscMiss.Add(ctx, 1, metric.WithAttributes(o.tAttrs...))
		}
	}
	end(span, resp.Error())
	return
}

func (o *otelclient) DoMultiCache(ctx context.Context, multi ...rueidis.CacheableTTL) (resps []rueidis.RedisResult) {
	ctx, span := start(ctx, multiCacheableFirst(multi), multiCacheableSum(multi), o.tAttrs)
	resps = o.client.DoMultiCache(ctx, multi...)
	for _, resp := range resps {
		if resp.NonRedisError() == nil {
			if resp.IsCacheHit() {
				cscHits.Add(ctx, 1, metric.WithAttributes(o.tAttrs...))
			} else {
				cscMiss.Add(ctx, 1, metric.WithAttributes(o.tAttrs...))
			}
		}
	}
	end(span, firstError(resps))
	return
}

func (o *otelclient) Dedicated(fn func(rueidis.DedicatedClient) error) (err error) {
	return o.client.Dedicated(func(client rueidis.DedicatedClient) error {
		return fn(&dedicated{client: client, mAttrs: o.mAttrs, tAttrs: o.tAttrs})
	})
}

func (o *otelclient) Dedicate() (rueidis.DedicatedClient, func()) {
	client, cancel := o.client.Dedicate()
	return &dedicated{client: client, mAttrs: o.mAttrs, tAttrs: o.tAttrs}, cancel
}

func (o *otelclient) Receive(ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	ctx, span := start(ctx, first(subscribe.Commands()), sum(subscribe.Commands()), o.tAttrs)
	err = o.client.Receive(ctx, subscribe, fn)
	end(span, err)
	return
}

func (o *otelclient) Nodes() map[string]rueidis.Client {
	nodes := o.client.Nodes()
	for addr, client := range nodes {
		nodes[addr] = &otelclient{client: client, mAttrs: o.mAttrs, tAttrs: o.tAttrs}
	}
	return nodes
}

func (o *otelclient) Close() {
	o.client.Close()
}

var _ rueidis.DedicatedClient = (*dedicated)(nil)

type dedicated struct {
	client rueidis.DedicatedClient
	mAttrs []attribute.KeyValue
	tAttrs []attribute.KeyValue
}

func (d *dedicated) B() cmds.Builder {
	return d.client.B()
}

func (d *dedicated) Do(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	ctx, span := start(ctx, first(cmd.Commands()), sum(cmd.Commands()), d.tAttrs)
	resp = d.client.Do(ctx, cmd)
	end(span, resp.Error())
	return
}

func (d *dedicated) DoMulti(ctx context.Context, multi ...rueidis.Completed) (resp []rueidis.RedisResult) {
	ctx, span := start(ctx, multiFirst(multi), multiSum(multi), d.tAttrs)
	resp = d.client.DoMulti(ctx, multi...)
	end(span, firstError(resp))
	return
}

func (d *dedicated) Receive(ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	ctx, span := start(ctx, first(subscribe.Commands()), sum(subscribe.Commands()), d.tAttrs)
	err = d.client.Receive(ctx, subscribe, fn)
	end(span, err)
	return
}

func (d *dedicated) SetPubSubHooks(hooks rueidis.PubSubHooks) <-chan error {
	return d.client.SetPubSubHooks(hooks)
}

func (d *dedicated) Close() {
	d.client.Close()
}

func first(s []string) string {
	return s[0]
}

func sum(s []string) (v int) {
	for _, str := range s {
		v += len(str)
	}
	return v
}

func firstError(s []rueidis.RedisResult) error {
	for _, result := range s {
		if err := result.Error(); err != nil && !rueidis.IsRedisNil(err) {
			return err
		}
	}
	return nil
}

func multiSum(multi []rueidis.Completed) (v int) {
	for _, cmd := range multi {
		v += sum(cmd.Commands())
	}
	return v
}

func multiCacheableSum(multi []rueidis.CacheableTTL) (v int) {
	for _, cmd := range multi {
		v += sum(cmd.Cmd.Commands())
	}
	return v
}

func multiFirst(multi []rueidis.Completed) string {
	if len(multi) == 0 {
		return ""
	}

	if len(multi) > 5 {
		multi = multi[:5]
	}
	size := 0
	for _, cmd := range multi {
		size += len(first(cmd.Commands()))
	}
	size += len(multi) - 1

	sb := strings.Builder{}
	sb.Grow(size)
	for i, cmd := range multi {
		sb.WriteString(first(cmd.Commands()))
		if i != len(multi)-1 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func multiCacheableFirst(multi []rueidis.CacheableTTL) string {
	if len(multi) == 0 {
		return ""
	}

	if len(multi) > 5 {
		multi = multi[:5]
	}
	size := 0
	for _, cmd := range multi {
		size += len(first(cmd.Cmd.Commands()))
	}
	size += len(multi) - 1

	sb := strings.Builder{}
	sb.Grow(size)
	for i, cmd := range multi {
		sb.WriteString(first(cmd.Cmd.Commands()))
		if i != len(multi)-1 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func start(ctx context.Context, op string, size int, attrs []attribute.KeyValue) (context.Context, trace.Span) {
	return tracer.Start(ctx, op, kind, attr(op, size), trace.WithAttributes(attrs...))
}

func end(span trace.Span, err error) {
	if err != nil && !rueidis.IsRedisNil(err) {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	} else {
		span.SetStatus(codes.Ok, "")
	}
	span.End()
}

// do not record full db.statement to avoid collecting sensitive data
func attr(op string, size int) trace.SpanStartEventOption {
	return trace.WithAttributes(dbattr, attribute.String("db.operation", op), attribute.Int("db.stmt_size", size))
}
