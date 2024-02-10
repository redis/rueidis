module mock

go 1.20

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.28
	go.uber.org/mock v0.3.0
)

require golang.org/x/sys v0.14.0 // indirect
