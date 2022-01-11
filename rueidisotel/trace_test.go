package rueidisotel

import (
	"context"
	"testing"
	"time"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/cmds"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestWithClient(t *testing.T) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		t.Fatal(err)
	}
	client = WithClient(client)
	defer client.Close()

	exp := tracetest.NewInMemoryExporter()
	otel.SetTracerProvider(trace.NewTracerProvider(trace.WithSyncer(exp)))

	ctx := context.Background()

	client.Do(ctx, client.B().Set().Key("key").Value("val").Build())
	validate(t, exp, "SET", codes.Ok)

	client.DoCache(ctx, client.B().Get().Key("key").Cache(), time.Minute)
	validate(t, exp, "GET", codes.Ok)

	client.Dedicated(func(client rueidis.DedicatedClient) error {
		client.Do(ctx, client.B().Set().Key("key").Value("val").Build())
		validate(t, exp, "SET", codes.Ok)

		client.DoMulti(
			ctx,
			client.B().Set().Key("key").Value("val").Build(),
			client.B().Set().Key("key").Value("val").Build(),
			client.B().Set().Key("key").Value("val").Build(),
			client.B().Set().Key("key").Value("val").Build(),
			client.B().Set().Key("key").Value("val").Build(),
			client.B().Set().Key("ignored").Value("ignored").Build(),
		)
		validate(t, exp, "SET SET SET SET SET", codes.Ok)

		client.Do(ctx, cmds.NewCompleted([]string{"unknown", "command"}))
		validate(t, exp, "unknown", codes.Error)

		client.DoMulti(ctx, cmds.NewCompleted([]string{"unknown", "command"}))
		validate(t, exp, "unknown", codes.Error)

		return nil
	})

	client.Do(ctx, cmds.NewCompleted([]string{"unknown", "command"}))
	validate(t, exp, "unknown", codes.Error)

	client.Do(ctx, cmds.NewCompleted([]string{"unknown", "command"}))
	validate(t, exp, "unknown", codes.Error)
}

func validate(t *testing.T, exp *tracetest.InMemoryExporter, op string, code codes.Code) {
	if name := exp.GetSpans().Snapshots()[0].Name(); name != op {
		t.Fatalf("unexpected span name %v", name)
	}
	if operation := exp.GetSpans().Snapshots()[0].Attributes()[1].Value.AsString(); operation != op {
		t.Fatalf("unexpected span name %v", operation)
	}
	if c := exp.GetSpans().Snapshots()[0].Status().Code; c != code {
		t.Fatalf("unexpected span statuc code %v", c)
	}
	exp.Reset()
}

func ExampleWithClient_openTelemetry() {
	client, _ := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	client = WithClient(client)
	defer client.Close()
}
