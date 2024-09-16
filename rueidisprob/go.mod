module github.com/redis/rueidis/rueidisprob

go 1.21

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.46
	github.com/twmb/murmur3 v1.1.8
)

require golang.org/x/sys v0.24.0 // indirect
