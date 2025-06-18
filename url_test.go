package rueidis

import (
	"context"
	"strings"
	"testing"
)

func TestParseURL(t *testing.T) {
	if opt, err := ParseURL("re dis://"); err == nil {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL(""); !strings.HasPrefix(err.Error(), "redis: invalid URL scheme") {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("rediss://"); err != nil || opt.TLSConfig == nil {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("valkeys://"); err != nil || opt.TLSConfig == nil {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("unix://"); err != nil || opt.DialCtxFn == nil {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("valkey://"); err != nil {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://"); err != nil || opt.InitAddress[0] != "localhost:6379" {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://localhost"); err != nil || opt.InitAddress[0] != "localhost:6379" {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://?addr=:6380&addr=:6381"); err != nil || opt.InitAddress[0] != "localhost:6379" || opt.InitAddress[1] != "localhost:6380" || opt.InitAddress[2] != "localhost:6381" {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://myhost:1234"); err != nil || opt.InitAddress[0] != "myhost:1234" {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://ooo:xxx@"); err != nil || opt.Username != "ooo" || opt.Password != "xxx" {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis:///1"); err != nil || opt.SelectDB != 1 {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis:///a"); !strings.HasPrefix(err.Error(), "redis: invalid database number") {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis:///1?db=a"); !strings.HasPrefix(err.Error(), "redis: invalid database number") {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis:////1"); !strings.HasPrefix(err.Error(), "redis: invalid URL path") {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://?dial_timeout=a"); !strings.HasPrefix(err.Error(), "redis: invalid dial timeout") {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://?write_timeout=a"); !strings.HasPrefix(err.Error(), "redis: invalid write timeout") {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("rediss://?skip_verify"); err != nil || opt.TLSConfig == nil || !opt.TLSConfig.InsecureSkipVerify {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("rediss://?skip_verify=true"); err != nil || opt.TLSConfig == nil || !opt.TLSConfig.InsecureSkipVerify {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("rediss://?skip_verify=a"); !strings.HasPrefix(err.Error(), "redis: invalid skip verify") {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://?protocol=2"); !opt.AlwaysRESP2 {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://?client_cache=0"); !opt.DisableCache {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://?max_retries=0"); !opt.DisableRetry {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://?client_name=0"); opt.ClientName != "0" {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("redis://?master_set=0"); opt.Sentinel.MasterSet != "0" {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("rediss://myhost:6379"); err != nil || opt.TLSConfig.ServerName != "myhost" {
		t.Fatalf("unexpected %v %v", opt, err)
	}
	if opt, err := ParseURL("unix:///path/to/redis.sock?db=1"); opt.DialCtxFn == nil || opt.InitAddress[0] != "/path/to/redis.sock" || opt.SelectDB != 1 {
		t.Fatalf("unexpected %v %v", opt, err)
	}
}

func TestMustParseURL(t *testing.T) {
	defer func() {
		if err := recover(); !strings.HasPrefix(err.(error).Error(), "redis: invalid URL path") {
			t.Failed()
		}
	}()
	MustParseURL("redis:////1")
}

func TestMustParseURLUnix(t *testing.T) {
	opt := MustParseURL("unix://")
	if conn, err := opt.DialCtxFn(context.Background(), "", &opt.Dialer, nil); !strings.Contains(err.Error(), "unix") {
		t.Fatalf("unexpected %v %v", conn, err) // the error should be "dial unix: missing address"
	}
}
