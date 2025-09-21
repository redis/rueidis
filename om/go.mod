module github.com/redis/rueidis/om

go 1.24.0

toolchain go1.24.2

replace github.com/redis/rueidis => ../

require (
	github.com/oklog/ulid/v2 v2.1.1
	github.com/redis/rueidis v1.0.64
)

require golang.org/x/sys v0.36.0 // indirect
