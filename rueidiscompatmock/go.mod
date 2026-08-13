module github.com/redis/rueidis/rueidiscompatmock

go 1.25.0

replace github.com/redis/rueidis => ../

replace github.com/redis/rueidis/mock => ../mock

replace github.com/redis/rueidis/rueidiscompat => ../rueidiscompat

require (
	github.com/redis/rueidis v1.0.77
	github.com/redis/rueidis/mock v1.0.77
	github.com/redis/rueidis/rueidiscompat v1.0.77
	go.uber.org/mock v0.6.0
)

require golang.org/x/sys v0.47.0 // indirect
