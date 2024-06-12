module github.com/redis/rueidis/rueidishook

go 1.20

replace (
	github.com/redis/rueidis => ../
	github.com/redis/rueidis/mock => ../mock
)

require (
	github.com/redis/rueidis v1.0.38
	github.com/redis/rueidis/mock v1.0.38
	go.uber.org/mock v0.4.0
)

require golang.org/x/sys v0.19.0 // indirect
