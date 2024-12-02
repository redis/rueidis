module github.com/redis/rueidis/rueidishook

go 1.21

replace (
	github.com/redis/rueidis => ../
	github.com/redis/rueidis/mock => ../mock
)

require (
	github.com/redis/rueidis v1.0.51
	github.com/redis/rueidis/mock v1.0.51
	go.uber.org/mock v0.4.0
)

require golang.org/x/sys v0.24.0 // indirect
