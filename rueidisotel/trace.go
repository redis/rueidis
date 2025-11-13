package rueidisotel

import (
	"context"
	"strings"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"github.com/redis/rueidis"
)

var (
	name   = "github.com/redis/rueidis"
	kind   = trace.WithSpanKind(trace.SpanKindClient)
	dbattr = attribute.String("db.system", "redis")
	dbstmt = attribute.Key("db.statement")
)

type contextKey struct{}

var labelerContextKey = contextKey{}

// Labeler is used to allow instrumentation to add additional attributes
// to metrics.
type Labeler struct {
	attrs []attribute.KeyValue
}

// Add appends new attributes to the labeler.
func (l *Labeler) Add(attrs ...attribute.KeyValue) {
	l.attrs = append(l.attrs, attrs...)
}

// Get returns the attributes added to the labeler.
func (l *Labeler) Get() []attribute.KeyValue {
	return l.attrs
}

// LabelerFromContext retrieves a Labeler instance from the provided context.
// It returns the labeler and a boolean indicating whether it was found.
// If no labeler is found, returns a new empty labeler and false.
func LabelerFromContext(ctx context.Context) (*Labeler, bool) {
	if l, ok := ctx.Value(labelerContextKey).(*Labeler); ok {
		return l, true
	}
	return &Labeler{}, false
}

// ContextWithLabeler returns a new context with the provided Labeler.
// Attributes added to the labeler will be included in cache metrics.
//
// Example:
//
//	// Check if labeler exists in context, create new context only if needed
//	labeler, ok := valkeyotel.LabelerFromContext(ctx)
//	if !ok {
//		ctx = valkeyotel.ContextWithLabeler(ctx, labeler)
//	}
//	labeler.Add(attribute.String("key_pattern", "book"))
//	client.DoCache(ctx, client.B().Get().Key("book:123").Cache(), time.Minute)
func ContextWithLabeler(ctx context.Context, labeler *Labeler) context.Context {
	return context.WithValue(ctx, labelerContextKey, labeler)
}

var _ rueidis.Client = (*otelclient)(nil)

// WithClient creates a new rueidis.Client with OpenTelemetry tracing enabled.
//
// Deprecated: use NewClient() instead.
func WithClient(client rueidis.Client, opts ...Option) rueidis.Client {
	cli, err := newClient(opts...)
	if err != nil {
		panic(err)
	}
	cli.client = client
	return cli
}

// Option is the Functional Options interface
type Option func(o *otelclient)

// TraceAttrs set additional attributes to append to each trace.
func TraceAttrs(attrs ...attribute.KeyValue) Option {
	return func(o *otelclient) {
		o.tAttrs = trace.WithAttributes(attrs...)
	}
}

// WithTracerProvider sets the TracerProvider for the otelclient.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return func(o *otelclient) {
		o.tracerProvider = provider
	}
}

// WithDBStatement tells the tracing hook to add raw redis commands to the db.statement attribute.
func WithDBStatement(f StatementFunc) Option {
	return func(o *otelclient) {
		o.dbStmtFunc = f
	}
}

// StatementFunc is the function that maps a command's tokens to a string to put in the db.statement attribute
type StatementFunc func(cmdTokens []string) string

type commandMetrics struct {
	duration   metric.Float64Histogram
	errors     metric.Int64Counter
	addOpts    []metric.AddOption
	recordOpts []metric.RecordOption
	opAttr     bool
}

func (c *commandMetrics) recordDuration(ctx context.Context, op string, now time.Time) {
	opts := c.recordOpts

	labeler, hasLabeler := ctx.Value(labelerContextKey).(*Labeler)
	hasLabelerAttrs := hasLabeler && len(labeler.attrs) > 0

	if c.opAttr && hasLabelerAttrs {
		opts = append(c.recordOpts,
			metric.WithAttributeSet(attribute.NewSet(attribute.String("operation", op))),
			metric.WithAttributeSet(attribute.NewSet(labeler.attrs...)))
	} else if c.opAttr {
		opts = append(c.recordOpts, metric.WithAttributeSet(attribute.NewSet(attribute.String("operation", op))))
	} else if hasLabelerAttrs {
		opts = append(c.recordOpts, metric.WithAttributeSet(attribute.NewSet(labeler.attrs...)))
	}

	c.duration.Record(ctx, time.Since(now).Seconds(), opts...)
}

func (c *commandMetrics) recordError(ctx context.Context, op string, err error) {
	if err != nil && !rueidis.IsRedisNil(err) {
		opts := c.addOpts

		labeler, hasLabeler := ctx.Value(labelerContextKey).(*Labeler)
		hasLabelerAttrs := hasLabeler && len(labeler.attrs) > 0

		if c.opAttr && hasLabelerAttrs {
			opts = append(c.addOpts,
				metric.WithAttributeSet(attribute.NewSet(attribute.String("operation", op))),
				metric.WithAttributeSet(attribute.NewSet(labeler.attrs...)))
		} else if c.opAttr {
			opts = append(c.addOpts, metric.WithAttributeSet(attribute.NewSet(attribute.String("operation", op))))
		} else if hasLabelerAttrs {
			opts = append(c.addOpts, metric.WithAttributeSet(attribute.NewSet(labeler.attrs...)))
		}

		c.errors.Add(ctx, 1, opts...)
	}
}

type otelclient struct {
	client          rueidis.Client
	meterProvider   metric.MeterProvider
	tracerProvider  trace.TracerProvider
	tracer          trace.Tracer
	meter           metric.Meter
	cscMiss         metric.Int64Counter
	cscHits         metric.Int64Counter
	tAttrs          trace.SpanStartEventOption
	dbStmtFunc      StatementFunc
	addOpts         []metric.AddOption
	recordOpts      []metric.RecordOption
	histogramOption HistogramOption
	commandMetrics
}

func (o *otelclient) B() rueidis.Builder {
	return o.client.B()
}

func (o *otelclient) Do(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	op := first(cmd.Commands())
	defer o.recordDuration(ctx, op, time.Now())
	ctx, span := o.start(ctx, op, sum(cmd.Commands()))
	if o.dbStmtFunc != nil {
		span.SetAttributes(dbstmt.String(o.dbStmtFunc(cmd.Commands())))
	}

	resp = o.client.Do(ctx, cmd)
	o.end(span, resp.Error())
	o.recordError(ctx, op, resp.Error())
	return
}

func (o *otelclient) DoMulti(ctx context.Context, multi ...rueidis.Completed) (resp []rueidis.RedisResult) {
	op := multiFirst(multi)
	defer o.recordDuration(ctx, op, time.Now())
	ctx, span := o.start(ctx, op, multiSum(multi))
	resp = o.client.DoMulti(ctx, multi...)
	err := firstError(resp)
	o.end(span, err)
	o.recordError(ctx, op, err)
	return
}

func (o *otelclient) DoStream(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResultStream) {
	op := first(cmd.Commands())
	defer o.recordDuration(ctx, op, time.Now())
	ctx, span := o.start(ctx, op, sum(cmd.Commands()))
	if o.dbStmtFunc != nil {
		span.SetAttributes(dbstmt.String(o.dbStmtFunc(cmd.Commands())))
	}

	resp = o.client.DoStream(ctx, cmd)
	o.end(span, resp.Error())
	o.recordError(ctx, op, resp.Error())
	return
}

func (o *otelclient) DoMultiStream(ctx context.Context, multi ...rueidis.Completed) (resp rueidis.MultiRedisResultStream) {
	op := multiFirst(multi)
	defer o.recordDuration(ctx, op, time.Now())
	ctx, span := o.start(ctx, op, multiSum(multi))
	resp = o.client.DoMultiStream(ctx, multi...)
	o.end(span, resp.Error())
	o.recordError(ctx, op, resp.Error())
	return
}

func (o *otelclient) DoCache(ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	op := first(cmd.Commands())
	defer o.recordDuration(ctx, op, time.Now())
	ctx, span := o.start(ctx, op, sum(cmd.Commands()))
	if o.dbStmtFunc != nil {
		span.SetAttributes(dbstmt.String(o.dbStmtFunc(cmd.Commands())))
	}

	resp = o.client.DoCache(ctx, cmd, ttl)
	if resp.NonRedisError() == nil {
		if resp.IsCacheHit() {
			o.recordCacheHit(ctx)
		} else {
			o.recordCacheMiss(ctx)
		}
	}
	o.end(span, resp.Error())
	o.recordError(ctx, op, resp.Error())
	return
}

func (o *otelclient) DoMultiCache(ctx context.Context, multi ...rueidis.CacheableTTL) (resps []rueidis.RedisResult) {
	op := multiCacheableFirst(multi)
	defer o.recordDuration(ctx, op, time.Now())
	ctx, span := o.start(ctx, op, multiCacheableSum(multi))
	resps = o.client.DoMultiCache(ctx, multi...)
	for _, resp := range resps {
		if resp.NonRedisError() == nil {
			if resp.IsCacheHit() {
				o.recordCacheHit(ctx)
			} else {
				o.recordCacheMiss(ctx)
			}
		}
	}
	err := firstError(resps)
	o.end(span, err)
	o.recordError(ctx, op, err)
	return
}

func (o *otelclient) Dedicated(fn func(rueidis.DedicatedClient) error) (err error) {
	return o.client.Dedicated(func(client rueidis.DedicatedClient) error {
		return fn(&dedicated{
			client:         client,
			tAttrs:         o.tAttrs,
			tracer:         o.tracer,
			dbStmtFunc:     o.dbStmtFunc,
			commandMetrics: o.commandMetrics,
		})
	})
}

func (o *otelclient) Dedicate() (rueidis.DedicatedClient, func()) {
	client, cancel := o.client.Dedicate()
	return &dedicated{
		client:         client,
		tAttrs:         o.tAttrs,
		tracer:         o.tracer,
		dbStmtFunc:     o.dbStmtFunc,
		commandMetrics: o.commandMetrics,
	}, cancel
}

func (o *otelclient) Receive(ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	ctx, span := o.start(ctx, first(subscribe.Commands()), sum(subscribe.Commands()))
	if o.dbStmtFunc != nil {
		span.SetAttributes(dbstmt.String(o.dbStmtFunc(subscribe.Commands())))
	}

	err = o.client.Receive(ctx, subscribe, fn)
	o.end(span, err)
	return
}

func (o *otelclient) Nodes() map[string]rueidis.Client {
	nodes := o.client.Nodes()
	for addr, client := range nodes {
		nodes[addr] = &otelclient{
			client:          client,
			meterProvider:   o.meterProvider,
			tracerProvider:  o.tracerProvider,
			tracer:          o.tracer,
			meter:           o.meter,
			cscMiss:         o.cscMiss,
			cscHits:         o.cscHits,
			addOpts:         o.addOpts,
			recordOpts:      o.recordOpts,
			tAttrs:          o.tAttrs,
			histogramOption: o.histogramOption,
			dbStmtFunc:      o.dbStmtFunc,
			commandMetrics:  o.commandMetrics,
		}
	}
	return nodes
}

func (o *otelclient) Mode() rueidis.ClientMode {
	return o.client.Mode()
}

func (o *otelclient) Close() {
	o.client.Close()
}

func (o *otelclient) recordCacheHit(ctx context.Context) {
	opts := o.addOpts
	if labeler, ok := ctx.Value(labelerContextKey).(*Labeler); ok && len(labeler.attrs) > 0 {
		opts = append(o.addOpts, metric.WithAttributeSet(attribute.NewSet(labeler.attrs...)))
	}
	o.cscHits.Add(ctx, 1, opts...)
}

func (o *otelclient) recordCacheMiss(ctx context.Context) {
	opts := o.addOpts
	if labeler, ok := ctx.Value(labelerContextKey).(*Labeler); ok && len(labeler.attrs) > 0 {
		opts = append(o.addOpts, metric.WithAttributeSet(attribute.NewSet(labeler.attrs...)))
	}
	o.cscMiss.Add(ctx, 1, opts...)
}

var _ rueidis.DedicatedClient = (*dedicated)(nil)

type dedicated struct {
	client     rueidis.DedicatedClient
	tracer     trace.Tracer
	tAttrs     trace.SpanStartEventOption
	dbStmtFunc StatementFunc
	commandMetrics
}

func (d *dedicated) B() rueidis.Builder {
	return d.client.B()
}

func (d *dedicated) Do(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	op := first(cmd.Commands())
	defer d.recordDuration(ctx, op, time.Now())
	ctx, span := d.start(ctx, op, sum(cmd.Commands()))
	if d.dbStmtFunc != nil {
		span.SetAttributes(dbstmt.String(d.dbStmtFunc(cmd.Commands())))
	}

	resp = d.client.Do(ctx, cmd)
	d.end(span, resp.Error())
	d.recordError(ctx, op, resp.Error())
	return
}

func (d *dedicated) DoMulti(ctx context.Context, multi ...rueidis.Completed) (resp []rueidis.RedisResult) {
	op := multiFirst(multi)
	defer d.recordDuration(ctx, op, time.Now())
	ctx, span := d.start(ctx, op, multiSum(multi))
	resp = d.client.DoMulti(ctx, multi...)
	err := firstError(resp)
	d.end(span, err)
	d.recordError(ctx, op, err)
	return
}

func (d *dedicated) Receive(ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	ctx, span := d.start(ctx, first(subscribe.Commands()), sum(subscribe.Commands()))
	if d.dbStmtFunc != nil {
		span.SetAttributes(dbstmt.String(d.dbStmtFunc(subscribe.Commands())))
	}

	err = d.client.Receive(ctx, subscribe, fn)
	d.end(span, err)
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

func (o *otelclient) start(ctx context.Context, op string, size int) (context.Context, trace.Span) {
	return startSpan(o.tracer, ctx, op, size, o.tAttrs)
}

func (o *otelclient) end(span trace.Span, err error) {
	endSpan(span, err)
}

func (d *dedicated) start(ctx context.Context, op string, size int) (context.Context, trace.Span) {
	return startSpan(d.tracer, ctx, op, size, d.tAttrs)
}

func (d *dedicated) end(span trace.Span, err error) {
	endSpan(span, err)
}

func startSpan(tracer trace.Tracer, ctx context.Context, op string, size int, attrs trace.SpanStartEventOption) (context.Context, trace.Span) {
	return tracer.Start(ctx, op, kind, attr(op, size), attrs)
}

func endSpan(span trace.Span, err error) {
	if err != nil && !rueidis.IsRedisNil(err) {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	} else {
		span.SetStatus(codes.Ok, "")
	}
	span.End()
}

// do not record the full db.statement to avoid collecting sensitive data
func attr(op string, size int) trace.SpanStartEventOption {
	return trace.WithAttributes(dbattr, attribute.String("db.operation", op), attribute.Int("db.stmt_size", size))
}
