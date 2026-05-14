package rueidiscompatmock

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/redis/rueidis/mock"
	"github.com/redis/rueidis/rueidiscompat"
	"go.uber.org/mock/gomock"
)

func TestNewAdapter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	{
		cm.ExpectGet("k").SetVal("v")
		if v, err := rdb.Get(ctx, "k").Result(); err != nil || v != "v" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
	}
	{
		cm.ExpectGet("missing").RedisNil()
		if err := rdb.Get(ctx, "missing").Err(); !errors.Is(err, rueidiscompat.Nil) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		cm.ExpectGet("k").SetErr(errors.New("any"))
		if err := rdb.Get(ctx, "k").Err(); err == nil || err.Error() != "any" {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		cm.ExpectSet("k", "v", 0).SetVal("OK")
		if err := rdb.Set(ctx, "k", "v", 0).Err(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		cm.ExpectSet("k", "v", 5*time.Second).SetVal("OK")
		if err := rdb.Set(ctx, "k", "v", 5*time.Second).Err(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		cm.ExpectSet("k", "v", 500*time.Millisecond).SetVal("OK")
		if err := rdb.Set(ctx, "k", "v", 500*time.Millisecond).Err(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		cm.ExpectSet("k", "v", rueidiscompat.KeepTTL).SetVal("OK")
		if err := rdb.Set(ctx, "k", "v", rueidiscompat.KeepTTL).Err(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		cm.ExpectSetNX("k", "v", 0).SetVal(true)
		if ok, err := rdb.SetNX(ctx, "k", "v", 0).Result(); err != nil || !ok {
			t.Fatalf("unexpected val %v err %v", ok, err)
		}
	}
	{
		cm.ExpectSetNX("k", "v", rueidiscompat.KeepTTL).SetVal(true)
		if ok, err := rdb.SetNX(ctx, "k", "v", rueidiscompat.KeepTTL).Result(); err != nil || !ok {
			t.Fatalf("unexpected val %v err %v", ok, err)
		}
	}
	{
		cm.ExpectSetNX("k", "v", 5*time.Second).SetVal(true)
		if ok, err := rdb.SetNX(ctx, "k", "v", 5*time.Second).Result(); err != nil || !ok {
			t.Fatalf("unexpected val %v err %v", ok, err)
		}
	}
	{
		cm.ExpectSetNX("k", "v", 500*time.Millisecond).SetVal(true)
		if ok, err := rdb.SetNX(ctx, "k", "v", 500*time.Millisecond).Result(); err != nil || !ok {
			t.Fatalf("unexpected val %v err %v", ok, err)
		}
	}
	{
		cm.ExpectGetSet("k", "v").SetVal("old")
		if v, err := rdb.GetSet(ctx, "k", "v").Result(); err != nil || v != "old" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
	}
	{
		cm.ExpectAppend("k", "v").SetVal(1)
		if n, err := rdb.Append(ctx, "k", "v").Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectStrLen("k").SetVal(1)
		if n, err := rdb.StrLen(ctx, "k").Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectIncr("k").SetVal(1)
		if n, err := rdb.Incr(ctx, "k").Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectIncrBy("k", 5).SetVal(6)
		if n, err := rdb.IncrBy(ctx, "k", 5).Result(); err != nil || n != 6 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectDecr("k").SetVal(5)
		if n, err := rdb.Decr(ctx, "k").Result(); err != nil || n != 5 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectDecrBy("k", 2).SetVal(3)
		if n, err := rdb.DecrBy(ctx, "k", 2).Result(); err != nil || n != 3 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectMSet("k1", "v1", "k2", "v2").SetVal("OK")
		if err := rdb.MSet(ctx, "k1", "v1", "k2", "v2").Err(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		cm.ExpectMGet("k1", "k2").SetVal([]any{"v1", "v2"})
		v, err := rdb.MGet(ctx, "k1", "k2").Result()
		if err != nil || len(v) != 2 || v[0] != "v1" || v[1] != "v2" {
			t.Fatalf("unexpected val %v err %v", v, err)
		}
	}
	{
		cm.ExpectDel("k1", "k2").SetVal(2)
		if n, err := rdb.Del(ctx, "k1", "k2").Result(); err != nil || n != 2 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectExists("a", "b").SetVal(2)
		if n, err := rdb.Exists(ctx, "a", "b").Result(); err != nil || n != 2 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectType("k").SetVal("string")
		if v, err := rdb.Type(ctx, "k").Result(); err != nil || v != "string" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
	}
	{
		cm.ExpectExpire("k", 10*time.Second).SetVal(true)
		if ok, err := rdb.Expire(ctx, "k", 10*time.Second).Result(); err != nil || !ok {
			t.Fatalf("unexpected val %v err %v", ok, err)
		}
	}
	{
		cm.ExpectTTL("k").SetVal(10 * time.Second)
		if d, err := rdb.TTL(ctx, "k").Result(); err != nil || d != 10*time.Second {
			t.Fatalf("unexpected val %v err %v", d, err)
		}
	}
	{
		cm.ExpectPing().SetVal("PONG")
		if v, err := rdb.Ping(ctx).Result(); err != nil || v != "PONG" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
	}
	{
		cm.ExpectEcho("hi").SetVal("hi")
		if v, err := rdb.Echo(ctx, "hi").Result(); err != nil || v != "hi" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
	}
	{
		cm.ExpectHSet("h", "f", "v").SetVal(1)
		if n, err := rdb.HSet(ctx, "h", "f", "v").Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectHSet("h", "f1", "v1", "f2", "v2").SetVal(2)
		if n, err := rdb.HSet(ctx, "h", "f1", "v1", "f2", "v2").Result(); err != nil || n != 2 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectHSet("h", []string{"f1", "v1"}).SetVal(1)
		if n, err := rdb.HSet(ctx, "h", []string{"f1", "v1"}).Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectHSet("h", []any{"f1", "v1"}).SetVal(1)
		if n, err := rdb.HSet(ctx, "h", []any{"f1", "v1"}).Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectHSet("h", map[string]any{"f": "v"}).SetVal(1)
		if n, err := rdb.HSet(ctx, "h", map[string]any{"f": "v"}).Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectHSet("h", map[string]string{"f": "v"}).SetVal(1)
		if n, err := rdb.HSet(ctx, "h", map[string]string{"f": "v"}).Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectHGet("h", "f").SetVal("v")
		if v, err := rdb.HGet(ctx, "h", "f").Result(); err != nil || v != "v" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
	}
	{
		cm.ExpectHDel("h", "f").SetVal(1)
		if n, err := rdb.HDel(ctx, "h", "f").Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectHGetAll("h").SetVal(map[string]string{"f": "v"})
		v, err := rdb.HGetAll(ctx, "h").Result()
		if err != nil || v["f"] != "v" {
			t.Fatalf("unexpected val %v err %v", v, err)
		}
	}
	{
		cm.ExpectLPush("l", "a", "b", "c").SetVal(3)
		if n, err := rdb.LPush(ctx, "l", "a", "b", "c").Result(); err != nil || n != 3 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectRPush("l", "a").SetVal(1)
		if n, err := rdb.RPush(ctx, "l", "a").Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectLPop("l").SetVal("a")
		if v, err := rdb.LPop(ctx, "l").Result(); err != nil || v != "a" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
	}
	{
		cm.ExpectRPop("l").SetVal("a")
		if v, err := rdb.RPop(ctx, "l").Result(); err != nil || v != "a" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
	}
	{
		cm.ExpectLLen("l").SetVal(1)
		if n, err := rdb.LLen(ctx, "l").Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectSAdd("s", "x", "y").SetVal(2)
		if n, err := rdb.SAdd(ctx, "s", "x", "y").Result(); err != nil || n != 2 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectSRem("s", "x").SetVal(1)
		if n, err := rdb.SRem(ctx, "s", "x").Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		cm.ExpectSMembers("s").SetVal([]string{"x"})
		v, err := rdb.SMembers(ctx, "s").Result()
		if err != nil || len(v) != 1 || v[0] != "x" {
			t.Fatalf("unexpected val %v err %v", v, err)
		}
	}
	{
		cm.ExpectEval("return 1", []string{}).SetValInt(1)
		v, err := rdb.Eval(ctx, "return 1", []string{}).Result()
		if err != nil || v.(int64) != 1 {
			t.Fatalf("unexpected val %v err %v", v, err)
		}
	}
	{
		cm.ExpectEval("return KEYS[1]", []string{"k"}, "a").SetVal("ok")
		v, err := rdb.Eval(ctx, "return KEYS[1]", []string{"k"}, "a").Result()
		if err != nil || v.(string) != "ok" {
			t.Fatalf("unexpected val %v err %v", v, err)
		}
	}
	{
		cm.ExpectEval("return nil", []string{}).RedisNil()
		if err := rdb.Eval(ctx, "return nil", []string{}).Err(); !errors.Is(err, rueidiscompat.Nil) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		cm.ExpectMGet("a", "b", "c", "d", "e").SetVal([]any{int64(1), true, "s", nil, 3.14})
		v, err := rdb.MGet(ctx, "a", "b", "c", "d", "e").Result()
		if err != nil || len(v) != 5 {
			t.Fatalf("unexpected val %v err %v", v, err)
		}
	}
	{
		cm.ExpectHGetAll("h").SetVal(map[string]string{})
		v, err := rdb.HGetAll(ctx, "h").Result()
		if err != nil || len(v) != 0 {
			t.Fatalf("unexpected val %v err %v", v, err)
		}
	}
	{
		cm.ExpectSMembers("s").SetVal([]string{})
		v, err := rdb.SMembers(ctx, "s").Result()
		if err != nil || len(v) != 0 {
			t.Fatalf("unexpected val %v err %v", v, err)
		}
	}
}

func TestPipeline(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	{
		cm.ExpectGet("k1").SetVal("v1")
		cm.ExpectSet("k2", "v2", 0).SetVal("OK")
		cm.ExpectDel("k3").SetVal(1)

		p := rdb.Pipeline()
		g := p.Get(ctx, "k1")
		s := p.Set(ctx, "k2", "v2", 0)
		d := p.Del(ctx, "k3")
		if _, err := p.Exec(ctx); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := g.Result(); err != nil || v != "v1" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
		if err := s.Err(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if n, err := d.Result(); err != nil || n != 1 {
			t.Fatalf("unexpected val %d err %v", n, err)
		}
	}
	{
		any := errors.New("any")
		cm.ExpectGet("k1").SetVal("v1")
		cm.ExpectGet("k2").SetErr(any)
		cm.ExpectGet("k3").RedisNil()

		p := rdb.Pipeline()
		c1 := p.Get(ctx, "k1")
		c2 := p.Get(ctx, "k2")
		c3 := p.Get(ctx, "k3")
		_, _ = p.Exec(ctx)

		if v, err := c1.Result(); err != nil || v != "v1" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
		if err := c2.Err(); err == nil || err.Error() != "any" {
			t.Fatalf("unexpected err %v", err)
		}
		if err := c3.Err(); !errors.Is(err, rueidiscompat.Nil) {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		cm.ExpectSetNX("k", "v", 0).SetVal(true)
		cm.ExpectEval("return 1", []string{}).SetValInt(1)

		p := rdb.Pipeline()
		nx := p.SetNX(ctx, "k", "v", 0)
		ev := p.Eval(ctx, "return 1", []string{})
		if _, err := p.Exec(ctx); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if ok, err := nx.Result(); err != nil || !ok {
			t.Fatalf("unexpected val %v err %v", ok, err)
		}
		if v, err := ev.Result(); err != nil || v.(int64) != 1 {
			t.Fatalf("unexpected val %v err %v", v, err)
		}
	}
	{
		cm.ExpectGet("solo").SetVal("v0")
		cm.ExpectGet("k1").SetVal("v1")
		cm.ExpectGet("k2").SetVal("v2")

		if v, err := rdb.Get(ctx, "solo").Result(); err != nil || v != "v0" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
		p := rdb.Pipeline()
		c1 := p.Get(ctx, "k1")
		c2 := p.Get(ctx, "k2")
		if _, err := p.Exec(ctx); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, _ := c1.Result(); v != "v1" {
			t.Fatalf("unexpected val %q", v)
		}
		if v, _ := c2.Result(); v != "v2" {
			t.Fatalf("unexpected val %q", v)
		}
	}
	{
		cm.ExpectGet("k1").SetVal("v1")
		cm.ExpectGet("k2").SetVal("v2")
		cm.ExpectGet("k3").SetVal("v3")

		p := rdb.Pipeline()
		c1 := p.Get(ctx, "k1")
		c2 := p.Get(ctx, "WRONG")
		c3 := p.Get(ctx, "k3")
		_, _ = p.Exec(ctx)

		if v, _ := c1.Result(); v != "v1" {
			t.Fatalf("unexpected val %q", v)
		}
		if err := c2.Err(); err == nil {
			t.Fatalf("unexpected err %v", err)
		}
		_ = c3
	}
}

func TestExpectations(t *testing.T) {
	{
		ctrl := gomock.NewController(t)
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)

		if err := rdb.Get(context.Background(), "k").Err(); err == nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := cm.ExpectationsWereMet(); err == nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		ctrl := gomock.NewController(t)
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)

		cm.ExpectGet("k").SetVal("v")
		rdb.Get(context.Background(), "k")
		if err := cm.ExpectationsWereMet(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		ctrl := gomock.NewController(t)
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)

		cm.ExpectGet("k").SetVal("v")
		if err := cm.ExpectationsWereMet(); err == nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		ctrl := gomock.NewController(t)
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)

		cm.ExpectGet("k").SetVal("v")
		cm.ClearExpect()
		if err := cm.ExpectationsWereMet(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		ctrl := gomock.NewController(t)
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)

		cm.ExpectGet("k1").SetVal("v1")
		cm.ExpectGet("k2").SetVal("v2")
		if _, err := rdb.Get(context.Background(), "k2").Result(); err == nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		ctrl := gomock.NewController(t)
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.MatchExpectationsInOrder(false)

		cm.ExpectGet("k1").SetVal("v1")
		cm.ExpectGet("k2").SetVal("v2")
		ctx := context.Background()
		if v, err := rdb.Get(ctx, "k2").Result(); err != nil || v != "v2" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
		if v, err := rdb.Get(ctx, "k1").Result(); err != nil || v != "v1" {
			t.Fatalf("unexpected val %q err %v", v, err)
		}
		if err := cm.ExpectationsWereMet(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		ctrl := gomock.NewController(t)
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.MatchExpectationsInOrder(false)

		cm.ExpectGet("a").SetVal("v")
		if err := rdb.Get(context.Background(), "b").Err(); err == nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		ctrl := gomock.NewController(t)
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.MatchExpectationsInOrder(false)

		cm.ExpectSet("a", "1", 0).SetVal("OK")
		cm.ExpectGet("a").SetVal("1")
		ctx := context.Background()
		p := rdb.Pipeline()
		g := p.Get(ctx, "a")
		s := p.Set(ctx, "a", "1", 0)
		if _, err := p.Exec(ctx); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, _ := g.Result(); v != "1" {
			t.Fatalf("unexpected val %q", v)
		}
		if err := s.Err(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
}

func TestDefaults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cm.ExpectGet("k")
	cm.ExpectDel("d")
	cm.ExpectSet("s", "v", 0)
	cm.ExpectSetNX("nx", "v", 0)
	cm.ExpectStrLen("sl")
	cm.ExpectIncr("i")

	if err := rdb.Get(ctx, "k").Err(); err == nil {
		t.Fatalf("unexpected err %v", err)
	}
	if n, _ := rdb.Del(ctx, "d").Result(); n != 0 {
		t.Fatalf("unexpected val %d", n)
	}
	if err := rdb.Set(ctx, "s", "v", 0).Err(); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	if ok, _ := rdb.SetNX(ctx, "nx", "v", 0).Result(); !ok {
		t.Fatalf("unexpected val %v", ok)
	}
	if n, _ := rdb.StrLen(ctx, "sl").Result(); n != 0 {
		t.Fatalf("unexpected val %d", n)
	}
	if n, _ := rdb.Incr(ctx, "i").Result(); n != 0 {
		t.Fatalf("unexpected val %d", n)
	}
}

func TestSetErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)
	any := errors.New("any")

	cases := []struct {
		setup func()
		call  func() error
	}{
		{func() { cm.ExpectGet("k").SetErr(any) }, func() error { return rdb.Get(ctx, "k").Err() }},
		{func() { cm.ExpectSet("k", "v", 0).SetErr(any) }, func() error { return rdb.Set(ctx, "k", "v", 0).Err() }},
		{func() { cm.ExpectSetNX("k", "v", 0).SetErr(any) }, func() error { return rdb.SetNX(ctx, "k", "v", 0).Err() }},
		{func() { cm.ExpectGetSet("k", "v").SetErr(any) }, func() error { return rdb.GetSet(ctx, "k", "v").Err() }},
		{func() { cm.ExpectAppend("k", "v").SetErr(any) }, func() error { return rdb.Append(ctx, "k", "v").Err() }},
		{func() { cm.ExpectStrLen("k").SetErr(any) }, func() error { return rdb.StrLen(ctx, "k").Err() }},
		{func() { cm.ExpectIncr("k").SetErr(any) }, func() error { return rdb.Incr(ctx, "k").Err() }},
		{func() { cm.ExpectIncrBy("k", 1).SetErr(any) }, func() error { return rdb.IncrBy(ctx, "k", 1).Err() }},
		{func() { cm.ExpectDecr("k").SetErr(any) }, func() error { return rdb.Decr(ctx, "k").Err() }},
		{func() { cm.ExpectDecrBy("k", 1).SetErr(any) }, func() error { return rdb.DecrBy(ctx, "k", 1).Err() }},
		{func() { cm.ExpectMGet("k").SetErr(any) }, func() error { return rdb.MGet(ctx, "k").Err() }},
		{func() { cm.ExpectMSet("k", "v").SetErr(any) }, func() error { return rdb.MSet(ctx, "k", "v").Err() }},
		{func() { cm.ExpectDel("k").SetErr(any) }, func() error { return rdb.Del(ctx, "k").Err() }},
		{func() { cm.ExpectExists("k").SetErr(any) }, func() error { return rdb.Exists(ctx, "k").Err() }},
		{func() { cm.ExpectType("k").SetErr(any) }, func() error { return rdb.Type(ctx, "k").Err() }},
		{func() { cm.ExpectExpire("k", time.Second).SetErr(any) }, func() error { return rdb.Expire(ctx, "k", time.Second).Err() }},
		{func() { cm.ExpectTTL("k").SetErr(any) }, func() error { return rdb.TTL(ctx, "k").Err() }},
		{func() { cm.ExpectPing().SetErr(any) }, func() error { return rdb.Ping(ctx).Err() }},
		{func() { cm.ExpectEcho("hi").SetErr(any) }, func() error { return rdb.Echo(ctx, "hi").Err() }},
		{func() { cm.ExpectHGet("h", "f").SetErr(any) }, func() error { return rdb.HGet(ctx, "h", "f").Err() }},
		{func() { cm.ExpectHSet("h", "f", "v").SetErr(any) }, func() error { return rdb.HSet(ctx, "h", "f", "v").Err() }},
		{func() { cm.ExpectHDel("h", "f").SetErr(any) }, func() error { return rdb.HDel(ctx, "h", "f").Err() }},
		{func() { cm.ExpectHGetAll("h").SetErr(any) }, func() error { return rdb.HGetAll(ctx, "h").Err() }},
		{func() { cm.ExpectLPush("l", "x").SetErr(any) }, func() error { return rdb.LPush(ctx, "l", "x").Err() }},
		{func() { cm.ExpectRPush("l", "x").SetErr(any) }, func() error { return rdb.RPush(ctx, "l", "x").Err() }},
		{func() { cm.ExpectLPop("l").SetErr(any) }, func() error { return rdb.LPop(ctx, "l").Err() }},
		{func() { cm.ExpectRPop("l").SetErr(any) }, func() error { return rdb.RPop(ctx, "l").Err() }},
		{func() { cm.ExpectLLen("l").SetErr(any) }, func() error { return rdb.LLen(ctx, "l").Err() }},
		{func() { cm.ExpectSAdd("s", "x").SetErr(any) }, func() error { return rdb.SAdd(ctx, "s", "x").Err() }},
		{func() { cm.ExpectSRem("s", "x").SetErr(any) }, func() error { return rdb.SRem(ctx, "s", "x").Err() }},
		{func() { cm.ExpectSMembers("s").SetErr(any) }, func() error { return rdb.SMembers(ctx, "s").Err() }},
		{func() { cm.ExpectEval("x", []string{}).SetErr(any) }, func() error { return rdb.Eval(ctx, "x", []string{}).Err() }},
	}
	for _, c := range cases {
		c.setup()
		if err := c.call(); err == nil || err.Error() != "any" {
			t.Fatalf("unexpected err %v", err)
		}
	}
	if err := cm.ExpectationsWereMet(); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestStr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	for _, in := range []any{"s", []byte("b"), int(1), int64(2), 3.14, true, false, nil} {
		cm.ExpectEcho(in).SetVal("ok")
		if _, err := rdb.Echo(ctx, in).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		now := time.Now()
		cm.ExpectEcho(now).SetVal("ok")
		if _, err := rdb.Echo(ctx, now).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		dur := 5 * time.Second
		cm.ExpectEcho(dur).SetVal("ok")
		if _, err := rdb.Echo(ctx, dur).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		bm := binMarshaler{data: "hello"}
		cm.ExpectEcho(bm).SetVal("ok")
		if _, err := rdb.Echo(ctx, bm).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
}

type binMarshaler struct{ data string }

func (b binMarshaler) MarshalBinary() ([]byte, error) { return []byte(b.data), nil }
