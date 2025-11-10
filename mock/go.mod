module github.com/redis/rueidis/mock

go 1.23.0

toolchain go1.23.4

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.68
	go.uber.org/mock v0.5.0
)

require golang.org/x/sys v0.31.0 // indirect
