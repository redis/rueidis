module github.com/redis/rueidis/om

go 1.24.9

toolchain go1.24.11

replace github.com/redis/rueidis => ../

require (
	github.com/oklog/ulid/v2 v2.1.1
	github.com/redis/rueidis v1.0.72
)

require golang.org/x/sys v0.39.0 // indirect
