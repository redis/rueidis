module github.com/redis/rueidis/rueidisotel

go 1.21

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.50
	go.opentelemetry.io/otel v1.28.0
	go.opentelemetry.io/otel/metric v1.28.0
	go.opentelemetry.io/otel/sdk v1.28.0
	go.opentelemetry.io/otel/sdk/metric v1.28.0
	go.opentelemetry.io/otel/trace v1.28.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
)
