module github.com/redis/rueidis/rueidisgcp

go 1.26.0

replace github.com/redis/rueidis => ../

require (
	github.com/redis/rueidis v1.0.79
	golang.org/x/oauth2 v0.38.0
)

require (
	cloud.google.com/go/compute/metadata v0.10.0 // indirect
	golang.org/x/sys v0.50.0 // indirect
)
