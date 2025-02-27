module github.com/dannotripp/rueidis/om

go 1.22.0

replace github.com/dannotripp/rueidis => ../

require (
	github.com/oklog/ulid/v2 v2.1.0
	github.com/dannotripp/rueidis v1.0.55
)

require golang.org/x/sys v0.30.0 // indirect
