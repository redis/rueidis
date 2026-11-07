module github.com/redis/rueidis/rueidisgcp

go 1.25.0

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.79
	golang.org/x/oauth2 v0.36.0
)

require (
	cloud.google.com/go/compute/metadata v0.9.0 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
)
