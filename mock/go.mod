module github.com/redis/rueidis/mock

go 1.20

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.38
	go.uber.org/mock v0.3.0
)

require golang.org/x/sys v0.19.0 // indirect
