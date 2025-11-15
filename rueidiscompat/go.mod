module github.com/redis/rueidis/rueidiscompat

go 1.23.0

toolchain go1.23.4

replace github.com/redis/rueidis => ../

replace github.com/redis/rueidis/mock => ../mock

require (
	github.com/onsi/ginkgo/v2 v2.22.2
	github.com/onsi/gomega v1.36.2
	github.com/redis/rueidis v1.0.68
	github.com/redis/rueidis/mock v1.0.68
	go.uber.org/mock v0.5.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/pprof v0.0.0-20250208200701-d0013a598941 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/tools v0.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
