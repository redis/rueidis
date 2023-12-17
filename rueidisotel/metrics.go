package rueidisotel

import (
	"context"
	"crypto/tls"
	"net"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"

	"github.com/redis/rueidis"
)

var (
	defaultHistogramBuckets = []float64{
		.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10,
	}
)

type HistogramOption struct {
	Buckets []float64
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

	if clientOption.DialFn == nil {
		clientOption.DialFn = defaultDialFn
	}

	attempt, err := oclient.meter.Int64Counter("rueidis_dial_attempt")
	if err != nil {
		return nil, err
	}
	oclient.attempt = attempt

	success, err := oclient.meter.Int64Counter("rueidis_dial_success")
	if err != nil {
		return nil, err
	}
	oclient.success = success

	conns, err := oclient.meter.Int64UpDownCounter("rueidis_dial_conns")
	if err != nil {
		return nil, err
	}
	oclient.conns = conns

	dialLatency, err := oclient.meter.Float64Histogram(
		"rueidis_dial_latency",
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(oclient.histogramOption.Buckets...),
	)
	if err != nil {
		return nil, err
	}
	oclient.dialLatency = dialLatency

	clientOption.DialFn = trackDialing(
		attempt, success, conns, dialLatency, clientOption.DialFn,
	)
	cli, err := rueidis.NewClient(clientOption)
	if err != nil {
		return nil, err
	}
	oclient.client = cli

	return oclient, nil
}

func newClient(opts ...Option) (*otelclient, error) {
	cli := &otelclient{}
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
	return cli, nil
}

func trackDialing(
	attempt metric.Int64Counter,
	success metric.Int64Counter,
	conns metric.Int64UpDownCounter,
	dialLatency metric.Float64Histogram,
	dialFn func(string, *net.Dialer, *tls.Config) (conn net.Conn, err error),
) func(string, *net.Dialer, *tls.Config) (conn net.Conn, err error) {
	return func(network string, dialer *net.Dialer, tlsConfig *tls.Config) (conn net.Conn, err error) {
		ctx := context.Background()
		attempt.Add(ctx, 1)

		start := time.Now()

		conn, err = dialFn(network, dialer, tlsConfig)
		if err != nil {
			return nil, err
		}

		// Use floating point division for higher precision (instead of Seconds method).
		dialLatency.Record(ctx, float64(time.Since(start))/float64(time.Second))
		success.Add(ctx, 1)
		conns.Add(ctx, 1)

		return &connTracker{
			Conn:  conn,
			conns: conns,
			once:  0,
		}, nil
	}
}

type connTracker struct {
	net.Conn
	conns metric.Int64UpDownCounter
	once  int32
}

func (t *connTracker) Close() error {
	if atomic.CompareAndSwapInt32(&t.once, 0, 1) {
		t.conns.Add(context.Background(), -1)
	}

	return t.Conn.Close()
}

func defaultDialFn(dst string, dialer *net.Dialer, cfg *tls.Config) (conn net.Conn, err error) {
	if cfg != nil {
		return tls.DialWithDialer(dialer, "tcp", dst, cfg)
	}
	return dialer.Dial("tcp", dst)
}
