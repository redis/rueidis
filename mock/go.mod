module github.com/redis/rueidis/mock

go 1.22.0

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.55
	go.uber.org/mock v0.5.0
)

require golang.org/x/sys v0.30.0 // indirect
