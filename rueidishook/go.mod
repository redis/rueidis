module github.com/redis/rueidis/rueidishook

go 1.23.0

toolchain go1.23.4

replace (
	github.com/redis/rueidis => ../
	github.com/redis/rueidis/mock => ../mock
)

require (
	github.com/redis/rueidis v1.0.68
	github.com/redis/rueidis/mock v1.0.68
	go.uber.org/mock v0.5.0
)

require golang.org/x/sys v0.31.0 // indirect
