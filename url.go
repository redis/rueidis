package rueidis

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ParseURL parses a redis URL into ClientOption.
// https://github.com/redis/redis-specifications/blob/master/uri/redis.txt
// Example:
//
//	redis://<user>:<password>@<host>:<port>/<db_number>
//	unix://<user>:<password>@</path/to/redis.sock>?db=<db_number>
func ParseURL(str string) (opt ClientOption, err error) {
	u, _ := url.Parse(str)
	switch u.Scheme {
	case "unix":
		opt.DialFn = func(s string, dialer *net.Dialer, config *tls.Config) (conn net.Conn, err error) {
			return dialer.Dial("unix", s)
		}
		opt.InitAddress = []string{strings.TrimSpace(u.Path)}
	case "rediss":
		opt.TLSConfig = &tls.Config{}
	case "redis":
	default:
		return opt, fmt.Errorf("redis: invalid URL scheme: %s", u.Scheme)
	}
	if opt.InitAddress == nil {
		host, port, _ := net.SplitHostPort(u.Host)
		if host == "" {
			host = u.Host
		}
		if host == "" {
			host = "localhost"
		}
		if port == "" {
			port = "6379"
		}
		opt.InitAddress = []string{net.JoinHostPort(host, port)}
	}
	if u.User != nil {
		opt.Username = u.User.Username()
		opt.Password, _ = u.User.Password()
	}
	if ps := strings.Split(u.Path, "/"); len(ps) == 2 {
		if opt.SelectDB, err = strconv.Atoi(ps[1]); err != nil {
			return opt, fmt.Errorf("redis: invalid database number: %q", ps[1])
		}
	} else if len(ps) > 2 {
		return opt, fmt.Errorf("redis: invalid URL path: %s", u.Path)
	}
	q := u.Query()
	if q.Has("db") {
		if opt.SelectDB, err = strconv.Atoi(q.Get("db")); err != nil {
			return opt, fmt.Errorf("redis: invalid database number: %q", q.Get("db"))
		}
	}
	if q.Has("dial_timeout") {
		if opt.Dialer.Timeout, err = time.ParseDuration(q.Get("dial_timeout")); err != nil {
			return opt, fmt.Errorf("redis: invalid dial timeout: %q", q.Get("dial_timeout"))
		}
	}
	if q.Has("write_timeout") {
		if opt.Dialer.Timeout, err = time.ParseDuration(q.Get("write_timeout")); err != nil {
			return opt, fmt.Errorf("redis: invalid write timeout: %q", q.Get("write_timeout"))
		}
	}
	opt.AlwaysRESP2 = q.Get("protocol") == "2"
	opt.DisableRetry = q.Get("max_retries") == "0"
	opt.ClientName = q.Get("client_name")
	opt.Sentinel.MasterSet = q.Get("master_set")
	return
}

func MustParseURL(str string) ClientOption {
	opt, err := ParseURL(str)
	if err != nil {
		panic(err)
	}
	return opt
}
