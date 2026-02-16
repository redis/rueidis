module github.com/redis/rueidis/rueidishook

go 1.24.9

toolchain go1.24.11

replace (
	github.com/redis/rueidis => ../
	github.com/redis/rueidis/mock => ../mock
)

require (
	github.com/redis/rueidis v1.0.71
	github.com/redis/rueidis/mock v1.0.71
	go.uber.org/mock v0.6.0
)

require golang.org/x/sys v0.39.0 // indirect
