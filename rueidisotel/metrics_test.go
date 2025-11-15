package rueidisotel

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"strings"
	"testing"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"

	"github.com/redis/rueidis"
)

func TestNewClient(t *testing.T) {
	t.Run("client option only (no ctx)", func(t *testing.T) {
		c, err := NewClient(rueidis.ClientOption{
			InitAddress: []string{"127.0.0.1:6379"},
			DialFn: func(dst string, dialer *net.Dialer, _ *tls.Config) (conn net.Conn, err error) {
				return dialer.Dial("tcp", dst)
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		defer c.Close()
	})

	t.Run("client option only", func(t *testing.T) {
		c, err := NewClient(rueidis.ClientOption{
			InitAddress: []string{"127.0.0.1:6379"},
			DialCtxFn: func(ctx context.Context, dst string, dialer *net.Dialer, _ *tls.Config) (conn net.Conn, err error) {
				return dialer.DialContext(ctx, "tcp", dst)
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		defer c.Close()
	})

	t.Run("meter provider", func(t *testing.T) {
		mr := metric.NewManualReader()
		meterProvider := metric.NewMeterProvider(metric.WithReader(mr))
		c, err := NewClient(
			rueidis.ClientOption{
				InitAddress: []string{"127.0.0.1:6379"},
				DialCtxFn: func(ctx context.Context, dst string, dialer *net.Dialer, _ *tls.Config) (conn net.Conn, err error) {
					return dialer.DialContext(ctx, "tcp", dst)
				},
			},
			WithMeterProvider(meterProvider),
		)
		if err != nil {
			t.Fatal(err)
		}
		defer c.Close()
	})

	t.Run("dial latency histogram option", func(t *testing.T) {
		c, err := NewClient(
			rueidis.ClientOption{
				InitAddress: []string{"127.0.0.1:6379"},
				DialCtxFn: func(ctx context.Context, dst string, dialer *net.Dialer, _ *tls.Config) (conn net.Conn, err error) {
					return dialer.DialContext(ctx, "tcp", dst)
				},
			},
			WithHistogramOption(HistogramOption{
				Buckets: []float64{1, 2, 3},
			}),
		)
		if err != nil {
			t.Fatal(err)
		}
		defer c.Close()
	})

	t.Run("DialFn by default", func(t *testing.T) {
		_, err := NewClient(rueidis.ClientOption{
			InitAddress: []string{"127.0.0.1:6379"},
		},
		)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestNewClientError(t *testing.T) {
	t.Run("invalid client option", func(t *testing.T) {
		_, err := NewClient(rueidis.ClientOption{
			InitAddress: []string{""},
			DialCtxFn: func(ctx context.Context, dst string, dialer *net.Dialer, _ *tls.Config) (conn net.Conn, err error) {
				return dialer.DialContext(ctx, "tcp", dst)
			},
		})
		if err == nil {
			t.Error(err)
		}
	})
}

func TestNewClientMeterError(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"rueidis_dial_attempt"},
		{"rueidis_dial_success"},
		{"rueidis_do_cache_miss"},
		{"rueidis_do_cache_hits"},
		{"rueidis_dial_conns"},
		{"rueidis_dial_latency"},
		{"rueidis_command_duration_seconds"},
		{"rueidis_command_errors"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			meterProvider := &MockMeterProvider{testName: tt.name}
			_, err := NewClient(
				rueidis.ClientOption{
					InitAddress: []string{"127.0.0.1:6379"},
				},
				WithMeterProvider(meterProvider),
			)
			if !errors.Is(err, errMocked) || !strings.Contains(err.Error(), tt.name) {
				t.Errorf("mocked error: got %s, want %s", err, errMocked)
			}
		})
	}
}

func TestTrackDialing(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mr := metric.NewManualReader()
		meterProvider := metric.NewMeterProvider(metric.WithReader(mr))
		c, err := NewClient(
			rueidis.ClientOption{
				InitAddress: []string{"127.0.0.1:6379"},
				DialCtxFn: func(ctx context.Context, dst string, dialer *net.Dialer, _ *tls.Config) (conn net.Conn, err error) {
					return dialer.DialContext(ctx, "tcp", dst)
				},
			},
			WithMeterProvider(meterProvider),
		)
		if err != nil {
			t.Fatal(err)
		}

		metrics := metricdata.ResourceMetrics{}
		if err := mr.Collect(context.Background(), &metrics); err != nil {
			t.Fatal(err)
		}
		attempt := int64CountMetric(metrics, "rueidis_dial_attempt")
		if attempt != 1 {
			t.Errorf("attempt: got %d, want 1", attempt)
		}
		success := int64CountMetric(metrics, "rueidis_dial_success")
		if success != 1 {
			t.Errorf("success: got %d, want 1", success)
		}
		conns := int64CountMetric(metrics, "rueidis_dial_conns")
		if conns != 1 {
			t.Errorf("conns: got %d, want 1", conns)
		}
		dialLatency := float64HistogramMetric(metrics, "rueidis_dial_latency")
		if dialLatency == 0 {
			t.Error("dial latency: got 0, want > 0")
		}

		c.Close()

		metrics = metricdata.ResourceMetrics{}
		if err := mr.Collect(context.Background(), &metrics); err != nil {
			t.Fatal(err)
		}
		conns = int64CountMetric(metrics, "rueidis_dial_conns")
		if conns != 0 {
			t.Errorf("conns: got %d, want 0", conns)
		}
	})

	t.Run("deduplicated closed connection conns metric", func(t *testing.T) {
		mr := metric.NewManualReader()
		meterProvider := metric.NewMeterProvider(metric.WithReader(mr))
		c, err := NewClient(
			rueidis.ClientOption{
				InitAddress: []string{"127.0.0.1:6379"},
				DialCtxFn: func(ctx context.Context, dst string, dialer *net.Dialer, _ *tls.Config) (conn net.Conn, err error) {
					return dialer.DialContext(ctx, "tcp", dst)
				},
			},
			WithMeterProvider(meterProvider),
		)
		if err != nil {
			t.Fatal(err)
		}

		c.Close()
		c.Close()

		metrics := metricdata.ResourceMetrics{}
		if err := mr.Collect(context.Background(), &metrics); err != nil {
			t.Fatal(err)
		}
		conns := int64CountMetric(metrics, "rueidis_dial_conns")
		if conns != 0 {
			t.Errorf("conns: got %d, want 0", conns)
		}
	})

	t.Run("failed to dial", func(t *testing.T) {
		mr := metric.NewManualReader()
		meterProvider := metric.NewMeterProvider(metric.WithReader(mr))
		_, err := NewClient(
			rueidis.ClientOption{
				InitAddress: []string{""},
				DialCtxFn: func(ctx context.Context, dst string, dialer *net.Dialer, _ *tls.Config) (conn net.Conn, err error) {
					return dialer.DialContext(ctx, "tcp", dst)
				},
			},
			WithMeterProvider(meterProvider),
		)
		if err == nil {
			t.Fatal(err)
		}

		metrics := metricdata.ResourceMetrics{}
		if err := mr.Collect(context.Background(), &metrics); err != nil {
			t.Fatal(err)
		}
		attempt := int64CountMetric(metrics, "rueidis_dial_attempt")
		if attempt != 1 {
			t.Errorf("attempt: got %d, want 1", attempt)
		}
		success := int64CountMetric(metrics, "rueidis_dial_success")
		if success != 0 {
			t.Errorf("success: got %d, want 0", success)
		}
		conns := int64CountMetric(metrics, "rueidis_dial_conns")
		if conns != 0 {
			t.Errorf("conns: got %d, want 0", conns)
		}
		dialLatency := float64HistogramMetric(metrics, "rueidis_dial_latency")
		if dialLatency != 0 {
			t.Error("dial latency: got 0, want 0")
		}
	})
}

func findMetric(metrics metricdata.ResourceMetrics, name string) metricdata.Aggregation {
	for _, sm := range metrics.ScopeMetrics {
		for _, m := range sm.Metrics {
			if m.Name == name {
				return m.Data
			}
		}
	}
	return nil
}

func int64CountMetric(metrics metricdata.ResourceMetrics, name string) int64 {
	m := findMetric(metrics, name)
	if data, ok := m.(metricdata.Sum[int64]); ok {
		return data.DataPoints[0].Value
	}
	return 0
}

func float64HistogramMetric(metrics metricdata.ResourceMetrics, name string) float64 {
	m := findMetric(metrics, name)
	if data, ok := m.(metricdata.Histogram[float64]); ok {
		return data.DataPoints[0].Sum
	}
	return 0
}

func int64CountMetricWithAttrs(metrics metricdata.ResourceMetrics, name string, attrKey string, attrValue string) int64 {
	m := findMetric(metrics, name)
	if data, ok := m.(metricdata.Sum[int64]); ok {
		for _, dp := range data.DataPoints {
			for _, attr := range dp.Attributes.ToSlice() {
				if string(attr.Key) == attrKey && attr.Value.AsString() == attrValue {
					return dp.Value
				}
			}
		}
	}
	return 0
}

func TestLabeler(t *testing.T) {
	t.Run("labeler basic operations", func(t *testing.T) {
		ctx := context.Background()

		// Test ContextWithLabeler
		labeler := &Labeler{}
		labeler.Add(attribute.String("key_pattern", "book"))
		ctxWithLabeler := ContextWithLabeler(ctx, labeler)

		retrieved, ok := LabelerFromContext(ctxWithLabeler)
		if !ok {
			t.Error("LabelerFromContext: labeler not found in context, want found")
		}
		attrs := retrieved.Get()
		if len(attrs) != 1 {
			t.Errorf("labeler attributes length: got %d, want 1", len(attrs))
		}
		if attrs[0].Key != "key_pattern" || attrs[0].Value.AsString() != "book" {
			t.Errorf("labeler attribute: got %v, want key_pattern=book", attrs[0])
		}

		// Test LabelerFromContext without labeler
		emptyLabeler, ok := LabelerFromContext(ctx)
		if ok {
			t.Error("LabelerFromContext without labeler: got found, want not found")
		}
		if len(emptyLabeler.Get()) != 0 {
			t.Errorf("empty labeler should have 0 attributes, got %d", len(emptyLabeler.Get()))
		}

		// Test Add with multiple attributes
		labeler2 := &Labeler{}
		labeler2.Add(
			attribute.String("key_pattern", "author"),
			attribute.String("tenant", "acme"),
		)
		if len(labeler2.Get()) != 2 {
			t.Errorf("labeler with 2 attributes: got %d, want 2", len(labeler2.Get()))
		}
	})

	t.Run("cache metrics with labeler attributes", func(t *testing.T) {
		client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
		if err != nil {
			t.Fatal(err)
		}

		mxp := metric.NewManualReader()
		meterProvider := metric.NewMeterProvider(metric.WithReader(mxp))

		oclient := WithClient(client, WithMeterProvider(meterProvider))
		defer oclient.Close()

		ctx := context.Background()

		// Set up some test data
		oclient.Do(ctx, oclient.B().Set().Key("book:1").Value("Book One").Build())
		oclient.Do(ctx, oclient.B().Set().Key("author:1").Value("Author One").Build())

		// DoCache with "book" labeler - cache miss
		bookLabeler := &Labeler{}
		bookLabeler.Add(attribute.String("key_pattern", "book"))
		ctxBook := ContextWithLabeler(ctx, bookLabeler)
		oclient.DoCache(ctxBook, oclient.B().Get().Key("book:1").Cache(), time.Minute)

		// DoCache with "book" labeler again - cache hit
		oclient.DoCache(ctxBook, oclient.B().Get().Key("book:1").Cache(), time.Minute)

		// DoCache with "author" labeler and multiple attributes - cache miss
		authorLabeler := &Labeler{}
		authorLabeler.Add(
			attribute.String("key_pattern", "author"),
			attribute.String("tenant", "test"),
		)
		ctxAuthor := ContextWithLabeler(ctx, authorLabeler)
		oclient.DoCache(ctxAuthor, oclient.B().Get().Key("author:1").Cache(), time.Minute)

		// DoCache with "author" labeler again - cache hit
		oclient.DoCache(ctxAuthor, oclient.B().Get().Key("author:1").Cache(), time.Minute)

		// DoCache without labeler - cache miss
		oclient.DoCache(ctx, oclient.B().Get().Key("other:1").Cache(), time.Minute)

		// Collect metrics
		metrics := metricdata.ResourceMetrics{}
		if err := mxp.Collect(ctx, &metrics); err != nil {
			t.Fatal(err)
		}

		// Validate "book" metrics
		bookMiss := int64CountMetricWithAttrs(metrics, "rueidis_do_cache_miss", "key_pattern", "book")
		if bookMiss != 1 {
			t.Errorf("book cache miss: got %d, want 1", bookMiss)
		}

		bookHits := int64CountMetricWithAttrs(metrics, "rueidis_do_cache_hits", "key_pattern", "book")
		if bookHits != 1 {
			t.Errorf("book cache hits: got %d, want 1", bookHits)
		}

		// Validate "author" metrics (should have both key_pattern and tenant attributes)
		authorMiss := int64CountMetricWithAttrs(metrics, "rueidis_do_cache_miss", "key_pattern", "author")
		if authorMiss != 1 {
			t.Errorf("author cache miss: got %d, want 1", authorMiss)
		}

		authorHits := int64CountMetricWithAttrs(metrics, "rueidis_do_cache_hits", "key_pattern", "author")
		if authorHits != 1 {
			t.Errorf("author cache hits: got %d, want 1", authorHits)
		}
	})

	t.Run("command metrics with labeler attributes", func(t *testing.T) {
		client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
		if err != nil {
			t.Fatal(err)
		}

		mxp := metric.NewManualReader()
		meterProvider := metric.NewMeterProvider(metric.WithReader(mxp))

		oclient := WithClient(client, WithMeterProvider(meterProvider))
		defer oclient.Close()

		ctx := context.Background()

		// Execute command with labeler
		labeler := &Labeler{}
		labeler.Add(attribute.String("service", "api"))
		ctxWithLabeler := ContextWithLabeler(ctx, labeler)
		oclient.Do(ctxWithLabeler, oclient.B().Set().Key("test:1").Value("value").Build())

		// Collect metrics
		metrics := metricdata.ResourceMetrics{}
		if err := mxp.Collect(ctx, &metrics); err != nil {
			t.Fatal(err)
		}

		// Verify command duration has the label
		m := findMetric(metrics, "rueidis_command_duration_seconds")
		if m == nil {
			t.Fatal("rueidis_command_duration_seconds metric not found")
		}

		data, ok := m.(metricdata.Histogram[float64])
		if !ok {
			t.Fatalf("unexpected metric type: %T", m)
		}

		found := false
		for _, dp := range data.DataPoints {
			for _, attr := range dp.Attributes.ToSlice() {
				if string(attr.Key) == "service" && attr.Value.AsString() == "api" {
					found = true
					break
				}
			}
		}

		if !found {
			t.Error("command duration metric should have service=api attribute")
		}
	})

	t.Run("command metrics with operation attr and labeler", func(t *testing.T) {
		client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
		if err != nil {
			t.Fatal(err)
		}

		mxp := metric.NewManualReader()
		meterProvider := metric.NewMeterProvider(metric.WithReader(mxp))

		// Enable operation metric attribute
		oclient := WithClient(client, WithMeterProvider(meterProvider), WithOperationMetricAttr())
		defer oclient.Close()

		ctx := context.Background()

		// Execute command with labeler and operation attr enabled
		labeler := &Labeler{}
		labeler.Add(attribute.String("tenant", "test-tenant"))
		ctxWithLabeler := ContextWithLabeler(ctx, labeler)
		oclient.Do(ctxWithLabeler, oclient.B().Set().Key("test:op:1").Value("value").Build())

		// Collect metrics
		metrics := metricdata.ResourceMetrics{}
		if err := mxp.Collect(ctx, &metrics); err != nil {
			t.Fatal(err)
		}

		// Verify command duration has both operation and labeler attributes
		m := findMetric(metrics, "rueidis_command_duration_seconds")
		if m == nil {
			t.Fatal("rueidis_command_duration_seconds metric not found")
		}

		data, ok := m.(metricdata.Histogram[float64])
		if !ok {
			t.Fatalf("unexpected metric type: %T", m)
		}

		foundTenant := false
		foundOperation := false
		for _, dp := range data.DataPoints {
			attrs := dp.Attributes.ToSlice()
			for _, attr := range attrs {
				if string(attr.Key) == "tenant" && attr.Value.AsString() == "test-tenant" {
					foundTenant = true
				}
				if string(attr.Key) == "operation" && attr.Value.AsString() == "SET" {
					foundOperation = true
				}
			}
		}

		if !foundTenant {
			t.Error("command duration metric should have tenant=test-tenant attribute")
		}
		if !foundOperation {
			t.Error("command duration metric should have operation=SET attribute")
		}
	})

	t.Run("command error metrics with operation attr and labeler", func(t *testing.T) {
		client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
		if err != nil {
			t.Fatal(err)
		}

		mxp := metric.NewManualReader()
		meterProvider := metric.NewMeterProvider(metric.WithReader(mxp))

		// Enable operation metric attribute
		oclient := WithClient(client, WithMeterProvider(meterProvider), WithOperationMetricAttr())
		defer oclient.Close()

		ctx := context.Background()

		labeler := &Labeler{}
		labeler.Add(attribute.String("service", "error-test"))
		ctxWithLabeler := ContextWithLabeler(ctx, labeler)

		oclient.Do(ctx, oclient.B().Set().Key("test:error:1").Value("not-a-list").Build())
		oclient.Do(ctxWithLabeler, oclient.B().Lpop().Key("test:error:1").Build())

		metrics := metricdata.ResourceMetrics{}
		if err := mxp.Collect(ctx, &metrics); err != nil {
			t.Fatal(err)
		}

		m := findMetric(metrics, "rueidis_command_errors")
		if m == nil {
			t.Fatal("rueidis_command_errors metric not found")
		}

		data, ok := m.(metricdata.Sum[int64])
		if !ok {
			t.Fatalf("unexpected metric type: %T", m)
		}

		foundService := false
		foundOperation := false
		for _, dp := range data.DataPoints {
			attrs := dp.Attributes.ToSlice()
			for _, attr := range attrs {
				if string(attr.Key) == "service" && attr.Value.AsString() == "error-test" {
					foundService = true
				}
				if string(attr.Key) == "operation" && attr.Value.AsString() == "LPOP" {
					foundOperation = true
				}
			}
		}

		if !foundService {
			t.Error("command error metric should have service=error-test attribute")
		}
		if !foundOperation {
			t.Error("command error metric should have operation=LPOP attribute")
		}
	})

	t.Run("command error metrics with labeler only", func(t *testing.T) {
		client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
		if err != nil {
			t.Fatal(err)
		}

		mxp := metric.NewManualReader()
		meterProvider := metric.NewMeterProvider(metric.WithReader(mxp))

		oclient := WithClient(client, WithMeterProvider(meterProvider))
		defer oclient.Close()

		ctx := context.Background()

		labeler := &Labeler{}
		labeler.Add(attribute.String("app", "test-app"))
		ctxWithLabeler := ContextWithLabeler(ctx, labeler)

		oclient.Do(ctx, oclient.B().Set().Key("test:error:2").Value("not-a-list").Build())
		oclient.Do(ctxWithLabeler, oclient.B().Lpop().Key("test:error:2").Build())

		metrics := metricdata.ResourceMetrics{}
		if err := mxp.Collect(ctx, &metrics); err != nil {
			t.Fatal(err)
		}

		m := findMetric(metrics, "rueidis_command_errors")
		if m == nil {
			t.Fatal("rueidis_command_errors metric not found")
		}

		data, ok := m.(metricdata.Sum[int64])
		if !ok {
			t.Fatalf("unexpected metric type: %T", m)
		}

		foundApp := false
		for _, dp := range data.DataPoints {
			attrs := dp.Attributes.ToSlice()
			for _, attr := range attrs {
				if string(attr.Key) == "app" && attr.Value.AsString() == "test-app" {
					foundApp = true
					break
				}
			}
		}

		if !foundApp {
			t.Error("command error metric should have app=test-app attribute")
		}
	})

	t.Run("command error metrics with operation attr only", func(t *testing.T) {
		client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
		if err != nil {
			t.Fatal(err)
		}

		mxp := metric.NewManualReader()
		meterProvider := metric.NewMeterProvider(metric.WithReader(mxp))

		oclient := WithClient(client, WithMeterProvider(meterProvider), WithOperationMetricAttr())
		defer oclient.Close()

		ctx := context.Background()

		oclient.Do(ctx, oclient.B().Set().Key("test:error:3").Value("not-a-list").Build())
		oclient.Do(ctx, oclient.B().Lpop().Key("test:error:3").Build())

		metrics := metricdata.ResourceMetrics{}
		if err := mxp.Collect(ctx, &metrics); err != nil {
			t.Fatal(err)
		}

		m := findMetric(metrics, "rueidis_command_errors")
		if m == nil {
			t.Fatal("rueidis_command_errors metric not found")
		}

		data, ok := m.(metricdata.Sum[int64])
		if !ok {
			t.Fatalf("unexpected metric type: %T", m)
		}

		foundOperation := false
		for _, dp := range data.DataPoints {
			attrs := dp.Attributes.ToSlice()
			for _, attr := range attrs {
				if string(attr.Key) == "operation" && attr.Value.AsString() == "LPOP" {
					foundOperation = true
					break
				}
			}
		}

		if !foundOperation {
			t.Error("command error metric should have operation=LPOP attribute")
		}
	})
}
