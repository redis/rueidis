package rueidisotel

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"

	metricapi "go.opentelemetry.io/otel/metric"

	"github.com/redis/rueidis"
)

var (
	errMocked = errors.New("ERROR_MOCKED")
)

// MockMeterProvider for testing purposes
type MockMeterProvider struct {
	metric.MeterProvider
	testName string
}

func (m *MockMeterProvider) Meter(name string, opts ...metricapi.MeterOption) metricapi.Meter {
	return &mockMeter{testName: m.testName}
}

type mockMeter struct {
	metricapi.Meter
	testName string
}

func (m *mockMeter) Int64Counter(name string, options ...metricapi.Int64CounterOption) (metricapi.Int64Counter, error) {
	if m.testName == name {
		return nil, fmt.Errorf("%w: %s", errMocked, m.testName)
	}
	return nil, nil
}

func (m *mockMeter) Int64UpDownCounter(name string, options ...metricapi.Int64UpDownCounterOption) (metricapi.Int64UpDownCounter, error) {
	if m.testName == name {
		return nil, fmt.Errorf("%w: %s", errMocked, m.testName)
	}
	return nil, nil
}
func (m *mockMeter) Float64Histogram(name string, options ...metricapi.Float64HistogramOption) (metricapi.Float64Histogram, error) {
	if m.testName == name {
		return nil, fmt.Errorf("%w: %s", errMocked, m.testName)
	}
	return nil, nil
}

func TestWithClientGlobalProvider(t *testing.T) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		t.Fatal(err)
	}

	exp := tracetest.NewInMemoryExporter()
	tracerProvider := trace.NewTracerProvider(trace.WithSyncer(exp))
	otel.SetTracerProvider(tracerProvider)

	mxp := metric.NewManualReader()
	meterProvider := metric.NewMeterProvider(metric.WithReader(mxp))
	otel.SetMeterProvider(meterProvider)

	client = WithClient(
		client,
		TraceAttrs(attribute.String("any", "label")),
		MetricAttrs(attribute.String("any", "label")),
	)
	defer client.Close()

	testWithClient(t, client, exp, mxp)
}

func TestWithClient(t *testing.T) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		t.Fatal(err)
	}

	exp := tracetest.NewInMemoryExporter()
	tracerProvider := trace.NewTracerProvider(trace.WithSyncer(exp))

	mxp := metric.NewManualReader()
	meterProvider := metric.NewMeterProvider(metric.WithReader(mxp))

	client = WithClient(
		client,
		TraceAttrs(attribute.String("any", "label")),
		MetricAttrs(attribute.String("any", "label")),
		WithTracerProvider(tracerProvider),
		WithMeterProvider(meterProvider),
	)
	defer client.Close()
	testWithClient(t, client, exp, mxp)
}

func testWithClient(t *testing.T, client rueidis.Client, exp *tracetest.InMemoryExporter, mxp metric.Reader) {
	ctx := context.Background()

	// test empty trace
	var emptyCompletedArr []rueidis.Completed
	resps := client.DoMulti(ctx, emptyCompletedArr...)
	if resps != nil {
		t.Error("unexpected response : ", resps)
	}
	validateTrace(t, exp, "", codes.Ok)

	var emtpyCacheableArr []rueidis.CacheableTTL
	resps = client.DoMultiCache(ctx, emtpyCacheableArr...)
	if resps != nil {
		t.Error("unexpected response : ", resps)
	}
	validateTrace(t, exp, "", codes.Ok)

	client.Do(ctx, client.B().Set().Key("key").Value("val").Build())
	validateTrace(t, exp, "SET", codes.Ok)

	client.DoMulti(ctx, client.B().Set().Key("key").Value("val").Build(), client.B().Set().Key("key").Value("val").Build())
	validateTrace(t, exp, "SET SET", codes.Ok)

	client.DoStream(ctx, client.B().Set().Key("key").Value("val").Build())
	validateTrace(t, exp, "SET", codes.Ok)

	client.DoMultiStream(ctx, client.B().Set().Key("key").Value("val").Build(), client.B().Set().Key("key").Value("val").Build())
	validateTrace(t, exp, "SET SET", codes.Ok)

	// first DoCache
	client.DoCache(ctx, client.B().Get().Key("key").Cache(), time.Minute)
	validateTrace(t, exp, "GET", codes.Ok)

	// second DoCache
	client.DoCache(ctx, client.B().Get().Key("key").Cache(), time.Minute)
	validateTrace(t, exp, "GET", codes.Ok)

	// first DoMultiCache
	client.DoMultiCache(ctx,
		rueidis.CT(client.B().Get().Key("key1").Cache(), time.Minute),
		rueidis.CT(client.B().Get().Key("key2").Cache(), time.Minute),
		rueidis.CT(client.B().Get().Key("key3").Cache(), time.Minute),
		rueidis.CT(client.B().Get().Key("key4").Cache(), time.Minute),
		rueidis.CT(client.B().Get().Key("key5").Cache(), time.Minute),
		rueidis.CT(client.B().Get().Key("key6").Cache(), time.Minute))
	validateTrace(t, exp, "GET GET GET GET GET", codes.Ok)

	// second DoMultiCache
	client.DoMultiCache(ctx,
		rueidis.CT(client.B().Get().Key("key1").Cache(), time.Minute),
		rueidis.CT(client.B().Get().Key("key2").Cache(), time.Minute),
		rueidis.CT(client.B().Get().Key("key3").Cache(), time.Minute),
		rueidis.CT(client.B().Get().Key("key4").Cache(), time.Minute),
		rueidis.CT(client.B().Get().Key("key5").Cache(), time.Minute),
		rueidis.CT(client.B().Get().Key("key6").Cache(), time.Minute))
	validateTrace(t, exp, "GET GET GET GET GET", codes.Ok)

	metrics := metricdata.ResourceMetrics{}
	if err := mxp.Collect(ctx, &metrics); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	validateMetrics(t, metrics, "rueidis_do_cache_miss", 7) // 1 (DoCache) + 6 (DoMultiCache)
	validateMetrics(t, metrics, "rueidis_do_cache_hits", 7) // 1 (DoCache) + 6 (DoMultiCache)

	ctx2, cancel := context.WithTimeout(ctx, time.Second/2)
	client.Receive(ctx2, client.B().Subscribe().Channel("ch").Build(), func(msg rueidis.PubSubMessage) {})
	cancel()
	validateTrace(t, exp, "SUBSCRIBE", codes.Error)

	var hookCh <-chan error

	client.Dedicated(func(client rueidis.DedicatedClient) error {
		client.Do(ctx, client.B().Set().Key("key").Value("val").Build())
		validateTrace(t, exp, "SET", codes.Ok)

		client.DoMulti(
			ctx,
			client.B().Set().Key("key").Value("val").Build(),
			client.B().Set().Key("key").Value("val").Build(),
			client.B().Set().Key("key").Value("val").Build(),
			client.B().Set().Key("key").Value("val").Build(),
			client.B().Set().Key("key").Value("val").Build(),
			client.B().Set().Key("ignored").Value("ignored").Build(),
		)
		validateTrace(t, exp, "SET SET SET SET SET", codes.Ok)

		client.Do(ctx, client.B().Arbitrary("unknown", "command").Build())
		validateTrace(t, exp, "unknown", codes.Error)

		client.DoMulti(ctx, client.B().Arbitrary("unknown", "command").Build())
		validateTrace(t, exp, "unknown", codes.Error)

		ctx2, cancel := context.WithTimeout(ctx, time.Second/2)
		client.Receive(ctx2, client.B().Subscribe().Channel("ch").Build(), func(msg rueidis.PubSubMessage) {})
		cancel()
		validateTrace(t, exp, "SUBSCRIBE", codes.Error)

		hookCh = client.SetPubSubHooks(rueidis.PubSubHooks{OnMessage: func(m rueidis.PubSubMessage) {}})

		client.Close()

		return nil
	})
	<-hookCh

	c, cancel := client.Dedicate()
	{
		c.Do(ctx, c.B().Set().Key("key").Value("val").Build())
		validateTrace(t, exp, "SET", codes.Ok)

		c.DoMulti(
			ctx,
			c.B().Set().Key("key").Value("val").Build(),
			c.B().Set().Key("key").Value("val").Build(),
			c.B().Set().Key("key").Value("val").Build(),
			c.B().Set().Key("key").Value("val").Build(),
			c.B().Set().Key("key").Value("val").Build(),
			c.B().Set().Key("ignored").Value("ignored").Build(),
		)
		validateTrace(t, exp, "SET SET SET SET SET", codes.Ok)

		c.Do(ctx, client.B().Arbitrary("unknown", "command").Build())
		validateTrace(t, exp, "unknown", codes.Error)

		c.DoMulti(ctx, client.B().Arbitrary("unknown", "command").Build())
		validateTrace(t, exp, "unknown", codes.Error)

		ctx2, cancel := context.WithTimeout(ctx, time.Second/2)
		c.Receive(ctx2, c.B().Subscribe().Channel("ch").Build(), func(msg rueidis.PubSubMessage) {})
		cancel()
		validateTrace(t, exp, "SUBSCRIBE", codes.Error)

		hookCh = c.SetPubSubHooks(rueidis.PubSubHooks{OnMessage: func(m rueidis.PubSubMessage) {}})

		c.Close()
	}
	cancel()
	<-hookCh

	client.Do(ctx, client.B().Arbitrary("unknown", "command").Build())
	validateTrace(t, exp, "unknown", codes.Error)

	client.Do(ctx, client.B().Arbitrary("unknown", "command").Build())
	validateTrace(t, exp, "unknown", codes.Error)

	nodes := client.Nodes()
	if len(nodes) == 0 {
		t.Fatalf("unexpected nodes count %v", len(nodes))
	}
	for _, client := range nodes {
		client.Do(ctx, client.B().Set().Key("key").Value("val").Build())
		validateTrace(t, exp, "SET", codes.Ok)

		client.DoMulti(ctx, client.B().Set().Key("key").Value("val").Build(), client.B().Set().Key("key").Value("val").Build())
		validateTrace(t, exp, "SET SET", codes.Ok)
	}
}

func validateTrace(t *testing.T, exp *tracetest.InMemoryExporter, op string, code codes.Code) {
	if name := exp.GetSpans().Snapshots()[0].Name(); name != op {
		t.Fatalf("unexpected span name %v", name)
	}
	if operation := exp.GetSpans().Snapshots()[0].Attributes()[1].Value.AsString(); operation != op {
		t.Fatalf("unexpected span name %v", operation)
	}
	customAttr := exp.GetSpans().Snapshots()[0].Attributes()[3]
	if string(customAttr.Key) != "any" || customAttr.Value.AsString() != "label" {
		t.Fatalf("unexpected custom attr %v", customAttr)
	}
	if c := exp.GetSpans().Snapshots()[0].Status().Code; c != code {
		t.Fatalf("unexpected span statuc code %v", c)
	}
	exp.Reset()
}

func TestWithMeterProvider(t *testing.T) {
	mockMeterProvider := &MockMeterProvider{}

	client := &otelclient{}
	option := WithMeterProvider(mockMeterProvider)
	option(client)

	if client.meterProvider != mockMeterProvider {
		t.Fatalf("unexpected MeterProvider: got %v, expected %v", client.meterProvider, mockMeterProvider)
	}
}

func TestWithClientSimple(t *testing.T) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		t.Fatal(err)
	}

	exp := tracetest.NewInMemoryExporter()
	tracerProvider := trace.NewTracerProvider(trace.WithSyncer(exp))

	mxp := metric.NewManualReader()
	meterProvider := metric.NewMeterProvider(metric.WithReader(mxp))

	client = WithClient(
		client,
		TraceAttrs(attribute.String("any", "label")),
		MetricAttrs(attribute.String("any", "label")),
		WithTracerProvider(tracerProvider),
		WithMeterProvider(meterProvider),
	)
	defer client.Close()

	cmd := client.B().Set().Key("key").Value("val").Build()
	client.Do(context.Background(), cmd)

	// Validate trace
	spans := exp.GetSpans().Snapshots()
	if len(spans) != 1 {
		t.Fatalf("expected 1 span, got %d", len(spans))
	}
	span := spans[0]
	if span.Name() != "SET" {
		t.Fatalf("unexpected span name: got %s, expected %s", span.Name(), "Set")
	}
	var found bool
	for _, attr := range span.Attributes() {
		if string(attr.Key) == "any" && attr.Value.AsString() == "label" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected attribute 'any: label' not found in span attributes")
	}
}

func validateMetrics(t *testing.T, metrics metricdata.ResourceMetrics, name string, value int64) {
	for _, sm := range metrics.ScopeMetrics {
		for _, m := range sm.Metrics {
			if m.Name == name {
				data := m.Data.(metricdata.Sum[int64])
				metricdatatest.AssertHasAttributes(t, data, attribute.String("any", "label"))
				if data.DataPoints[0].Value != value {
					t.Fatalf("unexpected metric value %v", data.DataPoints[0].Value)
				}
				return
			}
		}
	}
	t.Fatalf("metrics not found %v", name)
}

func ExampleWithClient_openTelemetry() {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	client = WithClient(client)
	defer client.Close()
}
