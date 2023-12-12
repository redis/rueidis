package rueidisotel

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"

	"github.com/redis/rueidis"
)

const metricNamespace = "github.com/redis/rueidis"

var DefaultDialLatencyHistogramDefaultBucketBoundaries = []float64{
	.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10,
}

var ErrDialFnRequired = errors.New("rueidisotel: clientOption.DialFn is required")

type DialLatencyHistogramOption struct {
	ExplicitBucketBoundaries []float64
}

type config struct {
	clientOption               rueidis.ClientOption
	meter                      metric.MeterProvider
	dialLatencyHistogramOption DialLatencyHistogramOption
}

// OptionFunc is a function that configures config.
type OptionFunc func(*config)

// WithClientOption sets the rueidis.ClientOption.
func WithClientOption(clientOption rueidis.ClientOption) OptionFunc {
	return func(o *config) {
		o.clientOption = clientOption
	}
}

// WithMetricMeterProvider sets the metric.MeterProvider.
func WithMetricMeterProvider(meter metric.MeterProvider) OptionFunc {
	return func(o *config) {
		o.meter = meter
	}
}

// WithDialLatencyHistogramOption sets the DialLatencyHistogramOption.
// If not set, DefaultDialLatencyHistogramDefaultBucketBoundaries will be used.
func WithDialLatencyHistogramOption(histogramOption DialLatencyHistogramOption) OptionFunc {
	return func(o *config) {
		o.dialLatencyHistogramOption = histogramOption
	}
}

// Client is a rueidis.Client that tracks connection metrics.
// It tracks the following metrics:
// - rueidis_dial_attempt: number of dial attempts
// - rueidis_dial_success: number of successful dials
// - rueidis_dial_conns: number of active connections
// - rueidis_dial_latency: dial latency in seconds
type Client struct {
	rueidis.Client
	attempt     metric.Int64Counter
	success     metric.Int64Counter
	conns       metric.Int64UpDownCounter
	dialLatency metric.Float64Histogram
}

// NewClient creates a new Client.
// It requires a rueidis.ClientOption with a DialFn.
func NewClient(opts ...OptionFunc) (rueidis.Client, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	if cfg.clientOption.DialFn == nil {
		return nil, ErrDialFnRequired
	}
	if cfg.meter == nil {
		cfg.meter = otel.GetMeterProvider()
	}
	if cfg.dialLatencyHistogramOption.ExplicitBucketBoundaries == nil {
		cfg.dialLatencyHistogramOption.ExplicitBucketBoundaries = DefaultDialLatencyHistogramDefaultBucketBoundaries
	}

	meter := cfg.meter.Meter(metricNamespace)
	attempt, err := meter.Int64Counter("rueidis_dial_attempt")
	if err != nil {
		return nil, err
	}
	success, err := meter.Int64Counter("rueidis_dial_success")
	if err != nil {
		return nil, err
	}
	conns, err := meter.Int64UpDownCounter("rueidis_dial_conns")
	if err != nil {
		return nil, err
	}
	dialLatency, err := meter.Float64Histogram(
		"rueidis_dial_latency",
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(cfg.dialLatencyHistogramOption.ExplicitBucketBoundaries...),
	)
	if err != nil {
		return nil, err
	}

	cfg.clientOption.DialFn = trackDialing(
		attempt, success, conns, dialLatency, cfg.clientOption.DialFn,
	)
	cli, err := rueidis.NewClient(cfg.clientOption)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:      cli,
		attempt:     attempt,
		success:     success,
		conns:       conns,
		dialLatency: dialLatency,
	}, nil
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

		dialLatency.Record(ctx, time.Since(start).Seconds())
		success.Add(ctx, 1)
		conns.Add(ctx, 1)

		return &connTracker{
			Conn:  conn,
			conns: conns,
		}, nil
	}
}

type connTracker struct {
	net.Conn
	conns metric.Int64UpDownCounter
	once  sync.Once
}

func (t *connTracker) Close() error {
	err := t.Conn.Close()
	t.once.Do(func() {
		t.conns.Add(context.Background(), -1)
	})

	if err != nil {
		return err
	}
	return nil
}
