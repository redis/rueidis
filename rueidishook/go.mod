module github.com/redis/rueidis/rueidishook

go 1.22.0

replace (
	github.com/redis/rueidis => ../
	github.com/redis/rueidis/mock => ../mock
)

require (
	github.com/redis/rueidis v1.0.54
	github.com/redis/rueidis/mock v1.0.54
	go.uber.org/mock v0.5.0
)

require golang.org/x/sys v0.29.0 // indirect
