module github.com/redis/rueidis/rueidiscompatrate

go 1.26.0

replace github.com/redis/rueidis => ../

replace github.com/redis/rueidis/rueidiscompat => ../rueidiscompat

replace github.com/redis/rueidis/rueidislimiter => ../rueidislimiter

require (
	github.com/onsi/ginkgo/v2 v2.34.0
	github.com/onsi/gomega v1.44.0
	github.com/redis/rueidis v1.0.79
	github.com/redis/rueidis/rueidiscompat v1.0.79
	github.com/redis/rueidis/rueidislimiter v1.0.79
)

require (
	github.com/Masterminds/semver/v3 v3.5.0 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/pprof v0.0.0-20261111200114-fac431ef4d26 // indirect
	github.com/sergi/go-diff v1.4.0 // indirect
	go.yaml.in/yaml/v3 v3.0.6 // indirect
	golang.org/x/mod v0.43.0 // indirect
	golang.org/x/net v0.62.0 // indirect
	golang.org/x/sync v0.25.0 // indirect
	golang.org/x/sys v0.50.0 // indirect
	golang.org/x/text v0.44.0 // indirect
	golang.org/x/tools v0.52.0 // indirect
)
