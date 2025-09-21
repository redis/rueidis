module github.com/redis/rueidis/rueidisotel

go 1.24.0

toolchain go1.24.2

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.64
	go.opentelemetry.io/otel v1.38.0
	go.opentelemetry.io/otel/metric v1.38.0
	go.opentelemetry.io/otel/sdk v1.38.0
	go.opentelemetry.io/otel/sdk/metric v1.38.0
	go.opentelemetry.io/otel/trace v1.38.0
)

require (
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
)
