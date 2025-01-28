module github.com/redis/rueidis/rueidiscompat

go 1.22

replace github.com/redis/rueidis => ../

replace github.com/redis/rueidis/mock => ../mock

require (
	github.com/onsi/ginkgo/v2 v2.20.1
	github.com/onsi/gomega v1.34.1
	github.com/redis/rueidis v1.0.53
	github.com/redis/rueidis/mock v1.0.53
	go.uber.org/mock v0.4.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20240727154555-813a5fbdbec8 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
