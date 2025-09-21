module github.com/redis/rueidis/rueidiscompat

go 1.24.0

toolchain go1.24.2

replace github.com/redis/rueidis => ../

replace github.com/redis/rueidis/mock => ../mock

require (
	github.com/onsi/ginkgo/v2 v2.25.3
	github.com/onsi/gomega v1.38.2
	github.com/redis/rueidis v1.0.64
	github.com/redis/rueidis/mock v1.0.64
	go.uber.org/mock v0.6.0
)

require (
	github.com/Masterminds/semver/v3 v3.4.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/pprof v0.0.0-20250403155104-27863c87afa6 // indirect
	go.uber.org/automaxprocs v1.6.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	golang.org/x/tools v0.36.0 // indirect
)
