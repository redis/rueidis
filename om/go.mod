module github.com/redis/rueidis/om

go 1.21

replace github.com/redis/rueidis => ../

require (
	github.com/oklog/ulid/v2 v2.1.0
	github.com/redis/rueidis v1.0.52
)

require golang.org/x/sys v0.24.0 // indirect
