module github.com/redis/rueidis/rueidisotel

go 1.20

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.28
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/metric v1.21.0
	go.opentelemetry.io/otel/sdk v1.21.0
	go.opentelemetry.io/otel/sdk/metric v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
)

require (
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	golang.org/x/sys v0.14.0 // indirect
)
