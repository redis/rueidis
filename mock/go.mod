module github.com/redis/rueidis/mock

go 1.26.0

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.80
	go.uber.org/mock v0.6.0
)

require golang.org/x/sys v0.50.0 // indirect
