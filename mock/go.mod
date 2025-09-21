module github.com/redis/rueidis/mock

go 1.24.0

toolchain go1.24.2

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.64
	go.uber.org/mock v0.6.0
)

require golang.org/x/sys v0.36.0 // indirect
