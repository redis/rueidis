module github.com/redis/rueidis/mock

go 1.24.9

toolchain go1.24.11

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.72
	go.uber.org/mock v0.5.0
)

require golang.org/x/sys v0.39.0 // indirect
