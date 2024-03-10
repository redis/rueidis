module github.com/redis/rueidis/rueidisprob

go 1.22.0

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v0.0.0-00010101000000-000000000000
	github.com/twmb/murmur3 v1.1.8
)

require golang.org/x/sys v0.17.0 // indirect
