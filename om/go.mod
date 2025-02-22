module github.com/redis/rueidis/om

go 1.22.0

replace github.com/redis/rueidis => ../

require (
	github.com/oklog/ulid/v2 v2.1.0
	github.com/redis/rueidis v1.0.54
)

require golang.org/x/sys v0.30.0 // indirect
