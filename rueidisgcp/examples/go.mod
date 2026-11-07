module github.com/redis/rueidis/rueidisgcp/examples

go 1.25.0

replace (
	github.com/redis/rueidis => ../../
	github.com/redis/rueidis/rueidisgcp => ../
)

require (
	github.com/redis/rueidis v1.0.79
	github.com/redis/rueidis/rueidisgcp v1.0.79
	golang.org/x/oauth2 v0.36.0
)

require (
	cloud.google.com/go/compute/metadata v0.9.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
)
