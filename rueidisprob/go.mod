module github.com/redis/rueidis/rueidisprob

go 1.23.0

toolchain go1.23.4

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.68
	github.com/twmb/murmur3 v1.1.8
)

require golang.org/x/sys v0.31.0 // indirect
