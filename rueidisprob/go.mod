module github.com/redis/rueidis/rueidisprob

go 1.24.9

toolchain go1.24.11

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.71
	github.com/twmb/murmur3 v1.1.8
)

require golang.org/x/sys v0.39.0 // indirect
