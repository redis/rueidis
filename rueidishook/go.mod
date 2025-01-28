module github.com/redis/rueidis/rueidishook

go 1.22

replace (
	github.com/redis/rueidis => ../
	github.com/redis/rueidis/mock => ../mock
)

require (
	github.com/redis/rueidis v1.0.53
	github.com/redis/rueidis/mock v1.0.53
	go.uber.org/mock v0.4.0
)

require golang.org/x/sys v0.24.0 // indirect
