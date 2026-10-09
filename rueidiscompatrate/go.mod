module github.com/redis/rueidis/rueidiscompatrate

go 1.25.0

replace github.com/redis/rueidis => ../

replace github.com/redis/rueidis/rueidiscompat => ../rueidiscompat

replace github.com/redis/rueidis/rueidislimiter => ../rueidislimiter

require (
	github.com/onsi/ginkgo/v2 v2.32.2
	github.com/onsi/gomega v1.42.1
	github.com/redis/rueidis v1.0.78
	github.com/redis/rueidis/rueidiscompat v1.0.78
	github.com/redis/rueidis/rueidislimiter v1.0.78
)

require (
	github.com/Masterminds/semver/v3 v3.5.0 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/pprof v0.0.0-20260802141513-ef3492d7dac3 // indirect
	github.com/sergi/go-diff v1.4.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/mod v0.39.0 // indirect
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	golang.org/x/tools v0.48.0 // indirect
)
