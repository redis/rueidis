module github.com/redis/rueidis/om

go 1.23.0

toolchain go1.23.4

replace github.com/redis/rueidis => ../

require (
	github.com/kr/pretty v0.1.0
	github.com/oklog/ulid/v2 v2.1.0
	github.com/redis/rueidis v1.0.67
)

require (
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
)
