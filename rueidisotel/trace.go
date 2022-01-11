package rueidisotel

import (
	"context"
	"strings"
	"time"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/cmds"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var (
	kind   = trace.WithSpanKind(trace.SpanKindClient)
	tracer = otel.Tracer("github.com/rueian/rueidis")
	dbattr = attribute.String("db.system", "redis")
)

var _ rueidis.Client = (*otelclient)(nil)

// WithClient creates a new rueidis.Client with OpenTelemetry tracing enabled
func WithClient(client rueidis.Client) *otelclient {
	return &otelclient{client: client}
}

type otelclient struct {
	client rueidis.Client
}

func (o *otelclient) B() *cmds.Builder {
	return o.client.B()
}

func (o *otelclient) Do(ctx context.Context, cmd cmds.Completed) (resp rueidis.RedisResult) {
	ctx, span := start(ctx, first(cmd.Commands()), sum(cmd.Commands()))
	resp = o.client.Do(ctx, cmd)
	end(span, resp.Error())
	return
}

func (o *otelclient) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	ctx, span := start(ctx, first(cmd.Commands()), sum(cmd.Commands()))
	resp = o.client.DoCache(ctx, cmd, ttl)
	end(span, resp.Error())
	return
}

func (o *otelclient) Dedicated(fn func(rueidis.DedicatedClient) error) (err error) {
	return o.client.Dedicated(func(client rueidis.DedicatedClient) error { return fn(&dedicated{client: client}) })
}

func (o *otelclient) Close() {
	o.client.Close()
}

var _ rueidis.DedicatedClient = (*dedicated)(nil)

type dedicated struct {
	client rueidis.DedicatedClient
}

func (d *dedicated) B() *cmds.Builder {
	return d.client.B()
}

func (d *dedicated) Do(ctx context.Context, cmd cmds.Completed) (resp rueidis.RedisResult) {
	ctx, span := start(ctx, first(cmd.Commands()), sum(cmd.Commands()))
	resp = d.client.Do(ctx, cmd)
	end(span, resp.Error())
	return
}

func (d *dedicated) DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []rueidis.RedisResult) {
	ctx, span := start(ctx, multiFirst(multi), multiSum(multi))
	resp = d.client.DoMulti(ctx, multi...)
	end(span, firstError(resp))
	return
}

func first(s []string) string {
	if len(s) != 0 {
		return s[0]
	}
	return ""
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

func multiSum(multi []cmds.Completed) (v int) {
	for _, cmd := range multi {
		v += sum(cmd.Commands())
	}
	return v
}

func multiFirst(multi []cmds.Completed) string {
	sb := strings.Builder{}
	for i, cmd := range multi {
		if i == 5 {
			break
		}
		sb.WriteString(first(cmd.Commands()))
		sb.WriteString(" ")
	}
	return sb.String()
}

func start(ctx context.Context, op string, size int) (context.Context, trace.Span) {
	return tracer.Start(ctx, op, kind, attr(op, size))
}

func end(span trace.Span, err error) {
	if err != nil && !rueidis.IsRedisNil(err) {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	span.End()
}

func attr(op string, size int) trace.SpanStartEventOption {
	return trace.WithAttributes(dbattr, attribute.String("db.operation", op), attribute.Int("db.statement_size", size))
}
