package rueidisotel

import (
	"context"
	"testing"
	"time"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/cmds"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/metrictest"
	"go.opentelemetry.io/otel/metric/number"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestWithClient(t *testing.T) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		t.Fatal(err)
	}
	client = WithClient(client, TraceAttrs(attribute.String("any", "label")), MetricAttrs(attribute.String("any", "label")))
	defer client.Close()

	exp := tracetest.NewInMemoryExporter()
	otel.SetTracerProvider(trace.NewTracerProvider(trace.WithSyncer(exp)))

	mxp := metrictest.NewMeterProvider()
	global.SetMeterProvider(mxp)

	ctx := context.Background()

	client.Do(ctx, client.B().Set().Key("key").Value("val").Build())
	validateTrace(t, exp, "SET", codes.Ok)

	// first DoCache
	client.DoCache(ctx, client.B().Get().Key("key").Cache(), time.Minute)
	validateTrace(t, exp, "GET", codes.Ok)
	validateMetrics(t, mxp, 0, cscMiss.SyncImpl().Descriptor().Name(), 1)

	// second DoCache
	client.DoCache(ctx, client.B().Get().Key("key").Cache(), time.Minute)
	validateTrace(t, exp, "GET", codes.Ok)
	validateMetrics(t, mxp, 0, cscMiss.SyncImpl().Descriptor().Name(), 1)
	validateMetrics(t, mxp, 1, cscHits.SyncImpl().Descriptor().Name(), 1)

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

		client.Do(ctx, cmds.NewCompleted([]string{"unknown", "command"}))
		validateTrace(t, exp, "unknown", codes.Error)

		client.DoMulti(ctx, cmds.NewCompleted([]string{"unknown", "command"}))
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

		c.Do(ctx, cmds.NewCompleted([]string{"unknown", "command"}))
		validateTrace(t, exp, "unknown", codes.Error)

		c.DoMulti(ctx, cmds.NewCompleted([]string{"unknown", "command"}))
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

	client.Do(ctx, cmds.NewCompleted([]string{"unknown", "command"}))
	validateTrace(t, exp, "unknown", codes.Error)

	client.Do(ctx, cmds.NewCompleted([]string{"unknown", "command"}))
	validateTrace(t, exp, "unknown", codes.Error)
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

func validateMetrics(t *testing.T, mxp *metrictest.MeterProvider, idx int, name string, value uint64) {
	if customAttr := mxp.MeasurementBatches[idx].Labels[0]; string(customAttr.Key) != "any" || customAttr.Value.AsString() != "label" {
		t.Fatalf("unexpected custom attr %v", customAttr)
	}
	if v := mxp.MeasurementBatches[idx].Measurements[0].Instrument.Descriptor().Name(); v != name {
		t.Fatalf("unexpected metric name %v", v)
	}
	if v := mxp.MeasurementBatches[idx].Measurements[0].Number; v != number.Number(value) {
		t.Fatalf("unexpected metric value %v", v)
	}
}

func ExampleWithClient_openTelemetry() {
	client, _ := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	client = WithClient(client)
	defer client.Close()
}
