module github.com/redis/rueidis/rueidisprob

go 1.26.0

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.80
	github.com/twmb/murmur3 v1.2.2
)

require golang.org/x/sys v0.50.0 // indirect
