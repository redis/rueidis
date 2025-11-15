module github.com/redis/rueidis/om

go 1.23.0

toolchain go1.23.4

replace github.com/redis/rueidis => ../

require (
	github.com/oklog/ulid/v2 v2.1.0
	github.com/redis/rueidis v1.0.68
)

require golang.org/x/sys v0.31.0 // indirect
