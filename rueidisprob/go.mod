module github.com/dannotripp/rueidis/rueidisprob

go 1.22.0

replace github.com/dannotripp/rueidis => ../

require (
	github.com/dannotripp/rueidis v1.0.55
	github.com/twmb/murmur3 v1.1.8
)

require golang.org/x/sys v0.30.0 // indirect
