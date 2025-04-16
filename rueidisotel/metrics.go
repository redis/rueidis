package rueidisotel

import (
	"context"
	"crypto/tls"
	"net"
	"sync/atomic"
	"time"

	"github.com/redis/rueidis"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

var (
	defaultHistogramBuckets = []float64{
		.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10,
	}
)

// MetricAttrs set additional attributes to append to each metric.
func MetricAttrs(attrs ...attribute.KeyValue) Option {
	return func(o *otelclient) {
		mAttrs := metric.WithAttributeSet(attribute.NewSet(attrs...))
		// Allocate slices once and use many times
		o.addOpts = []metric.AddOption{mAttrs}
		o.recordOpts = []metric.RecordOption{mAttrs}
	}
}

// WithMeterProvider sets the MeterProvider for the otelclient.
func WithMeterProvider(provider metric.MeterProvider) Option {
	return func(o *otelclient) {
		o.meterProvider = provider
	}
}

// WithOperationMetricAttr sets the operation name as an attribute for duration and error metrics.
// This may cause memory usage to increase with the number of commands used.
func WithOperationMetricAttr() Option {
	return func(cli *otelclient) {
		cli.commandMetrics.opAttr = true
	}
}

type HistogramOption struct {
	Buckets []float64
}

type dialMetrics struct {
	attempt    metric.Int64Counter
	success    metric.Int64Counter
	counts     metric.Int64UpDownCounter
	latency    metric.Float64Histogram
	addOpts    []metric.AddOption
	recordOpts []metric.RecordOption
}

type dialTracer struct {
	trace.Tracer
	tAttrs trace.SpanStartEventOption
}

// WithHistogramOption sets the HistogramOption.
// If not set, DefaultHistogramBuckets will be used.
func WithHistogramOption(histogramOption HistogramOption) Option {
	return func(cli *otelclient) {
		cli.histogramOption = histogramOption
	}
}

// NewClient creates a new Client.
// The following metrics are recorded:
// - rueidis_dial_attempt: number of dial attempts
// - rueidis_dial_success: number of successful dials
// - rueidis_dial_conns: number of active connections
// - rueidis_dial_latency: dial latency in seconds
func NewClient(clientOption rueidis.ClientOption, opts ...Option) (rueidis.Client, error) {
	oclient, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	if clientOption.DialCtxFn == nil {
		clientOption.DialCtxFn = defaultDialFn
		if clientOption.DialFn != nil {
			clientOption.DialCtxFn = func(_ context.Context, s string, dialer *net.Dialer, config *tls.Config) (conn net.Conn, err error) {
				return clientOption.DialFn(s, dialer, config)
			}
		}
	}

	metrics := dialMetrics{
		addOpts:    oclient.addOpts,
		recordOpts: oclient.recordOpts,
	}

	metrics.attempt, err = oclient.meter.Int64Counter("rueidis_dial_attempt")
	if err != nil {
		return nil, err
	}

	metrics.success, err = oclient.meter.Int64Counter("rueidis_dial_success")
	if err != nil {
		return nil, err
	}

	metrics.counts, err = oclient.meter.Int64UpDownCounter("rueidis_dial_conns")
	if err != nil {
		return nil, err
	}

	metrics.latency, err = oclient.meter.Float64Histogram(
		"rueidis_dial_latency",
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(oclient.histogramOption.Buckets...),
	)
	if err != nil {
		return nil, err
	}

	clientOption.DialCtxFn = trackDialing(metrics, dialTracer{Tracer: oclient.tracer, tAttrs: oclient.tAttrs}, clientOption.DialCtxFn)

	cli, err := rueidis.NewClient(clientOption)
	if err != nil {
		return nil, err
	}
	oclient.client = cli

	return oclient, nil
}

func newClient(opts ...Option) (*otelclient, error) {
	cli := &otelclient{
		tAttrs: trace.WithAttributes(),
	}
	for _, opt := range opts {
		opt(cli)
	}
	if cli.histogramOption.Buckets == nil {
		cli.histogramOption.Buckets = defaultHistogramBuckets
	}
	if cli.meterProvider == nil {
		cli.meterProvider = otel.GetMeterProvider() // Default to global MeterProvider
	}
	if cli.tracerProvider == nil {
		cli.tracerProvider = otel.GetTracerProvider() // Default to global TracerProvider
	}

	// Now that we have the meterProvider and tracerProvider, get the Meter and Tracer
	cli.meter = cli.meterProvider.Meter(name)
	cli.tracer = cli.tracerProvider.Tracer(name)
	// Now create the counters using the meter
	var err error
	cli.cscMiss, err = cli.meter.Int64Counter("rueidis_do_cache_miss")
	if err != nil {
		return nil, err
	}
	cli.cscHits, err = cli.meter.Int64Counter("rueidis_do_cache_hits")
	if err != nil {
		return nil, err
	}
	cli.commandMetrics.addOpts = cli.addOpts
	cli.commandMetrics.recordOpts = cli.recordOpts
	cli.commandMetrics.duration, err = cli.meter.Float64Histogram(
		"rueidis_command_duration_seconds",
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(defaultHistogramBuckets...),
	)
	if err != nil {
		return nil, err
	}
	cli.commandMetrics.errors, err = cli.meter.Int64Counter("rueidis_command_errors")
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func trackDialing(m dialMetrics, t dialTracer, dialFn func(context.Context, string, *net.Dialer, *tls.Config) (conn net.Conn, err error)) func(context.Context, string, *net.Dialer, *tls.Config) (conn net.Conn, err error) {
	return func(ctx context.Context, dst string, dialer *net.Dialer, tlsConfig *tls.Config) (conn net.Conn, err error) {
		ctx, span := t.Start(ctx, "redis.dial", kind, trace.WithAttributes(dbattr, attribute.String("server.address", dst)), t.tAttrs)
		defer span.End()

		m.attempt.Add(ctx, 1, m.addOpts...)

		start := time.Now()

		conn, err = dialFn(ctx, dst, dialer, tlsConfig)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		span.SetStatus(codes.Ok, "")

		// Use floating point division for higher precision (instead of Seconds method).
		m.latency.Record(ctx, float64(time.Since(start))/float64(time.Second), m.recordOpts...)
		m.success.Add(ctx, 1, m.addOpts...)
		m.counts.Add(ctx, 1, m.addOpts...)

		return &connTracker{
			Conn:    conn,
			counts:  m.counts,
			addOpts: m.addOpts,
			once:    0,
		}, nil
	}
}

type connTracker struct {
	net.Conn
	counts  metric.Int64UpDownCounter
	addOpts []metric.AddOption
	once    int32
}

func (t *connTracker) Close() error {
	if atomic.CompareAndSwapInt32(&t.once, 0, 1) {
		t.counts.Add(context.Background(), -1, t.addOpts...)
	}

	return t.Conn.Close()
}

func defaultDialFn(ctx context.Context, dst string, dialer *net.Dialer, cfg *tls.Config) (conn net.Conn, err error) {
	if cfg != nil {
		td := tls.Dialer{NetDialer: dialer, Config: cfg}
		return td.DialContext(ctx, "tcp", dst)
	}
	return dialer.DialContext(ctx, "tcp", dst)
}
