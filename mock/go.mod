module github.com/redis/rueidis/mock

go 1.21

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.53
	go.uber.org/mock v0.4.0
)

require golang.org/x/sys v0.24.0 // indirect
