package rueidiscompatmock

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/redis/rueidis"
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
		anyErr := errors.New("any")
		cm.ExpectGet("k1").SetVal("v1")
		cm.ExpectGet("k2").SetErr(anyErr)
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

	for _, call := range []func() error{
		func() error { return rdb.Get(ctx, "k").Err() },
		func() error { return rdb.Del(ctx, "d").Err() },
		func() error { return rdb.Set(ctx, "s", "v", 0).Err() },
		func() error { return rdb.SetNX(ctx, "nx", "v", 0).Err() },
		func() error { return rdb.StrLen(ctx, "sl").Err() },
		func() error { return rdb.Incr(ctx, "i").Err() },
	} {
		if err := call(); err == nil || !strings.Contains(err.Error(), "return value is required") {
			t.Fatalf("unexpected err %v", err)
		}
	}
}

func TestRedismockRuntimeParity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	anyErr := errors.New("any")
	errExp := cm.ExpectGet("err")
	errExp.SetErr(anyErr)
	errExp.SetVal("ignored")
	if err := rdb.Get(ctx, "err").Err(); !errors.Is(err, anyErr) {
		t.Fatalf("unexpected err %v", err)
	}

	nilExp := cm.ExpectGet("nil")
	nilExp.RedisNil()
	nilExp.SetVal("ignored")
	if err := rdb.Get(ctx, "nil").Err(); !errors.Is(err, rueidiscompat.Nil) {
		t.Fatalf("unexpected err %v", err)
	}

	members := []string{"a"}
	cm.ExpectSMembers("s").SetVal(members)
	members[0] = "b"
	if got, err := rdb.SMembers(ctx, "s").Result(); err != nil || len(got) != 1 || got[0] != "a" {
		t.Fatalf("unexpected val %v err %v", got, err)
	}

	cm.Regexp().ExpectGet("user:[0-9]+").SetVal("regexp")
	if got, err := rdb.Get(ctx, "user:42").Result(); err != nil || got != "regexp" {
		t.Fatalf("unexpected val %q err %v", got, err)
	}

	cm.CustomMatch(func(_, actual []any) error {
		if len(actual) == 2 && actual[1] == "dynamic" {
			return nil
		}
		return errors.New("no match")
	}).ExpectGet("ignored").SetVal("custom")
	if got, err := rdb.Get(ctx, "dynamic").Result(); err != nil || got != "custom" {
		t.Fatalf("unexpected val %q err %v", got, err)
	}

	cm.ExpectSlowLogGet(1).SetVal([]rueidiscompat.SlowLog{{
		ID:         1,
		Time:       time.Unix(1, 0),
		Duration:   time.Microsecond,
		Args:       []string{"GET", "k"},
		ClientName: "client",
	}})
	slow, err := rdb.SlowLogGet(ctx, 1).Result()
	if err != nil || len(slow) != 1 || slow[0].ClientAddr != "" || slow[0].ClientName != "client" {
		t.Fatalf("unexpected slowlog %v err %v", slow, err)
	}

	cm.ExpectFunctionList(rueidiscompat.FunctionListQuery{}).SetVal([]rueidiscompat.Library{{
		Name:      "lib",
		Engine:    "LUA",
		Functions: []rueidiscompat.Function{{Name: "fn"}},
	}})
	libs, err := rdb.FunctionList(ctx, rueidiscompat.FunctionListQuery{}).Result()
	if err != nil || len(libs) != 1 || len(libs[0].Functions) != 1 || libs[0].Functions[0].Flags == nil {
		t.Fatalf("unexpected libraries %v err %v", libs, err)
	}

	geoQuery := rueidiscompat.GeoSearchLocationQuery{
		GeoSearchQuery: rueidiscompat.GeoSearchQuery{Member: "origin", Radius: 10, RadiusUnit: "km"},
		WithCoord:      true,
		WithDist:       true,
		WithHash:       true,
	}
	cm.ExpectGeoSearchLocation("geo", &geoQuery).SetVal([]rueidiscompat.GeoLocation{{
		Name:      "place",
		Longitude: 1.25,
		Latitude:  2.5,
		Dist:      3.75,
		GeoHash:   99,
	}})
	geo, err := rdb.GeoSearchLocation(ctx, "geo", geoQuery).Result()
	if err != nil || len(geo) != 1 || geo[0].Longitude != 1.25 || geo[0].Latitude != 2.5 || geo[0].Dist != 3.75 || geo[0].GeoHash != 99 {
		t.Fatalf("unexpected geo %v err %v", geo, err)
	}

	cm.ExpectXReadStreams("s1", "s2", "0", "0").SetVal([]rueidiscompat.XStream{
		{Stream: "s2", Messages: []rueidiscompat.XMessage{{ID: "2-0", Values: map[string]any{"b": "2", "a": "1"}}}},
		{Stream: "s1", Messages: []rueidiscompat.XMessage{{ID: "1-0", Values: map[string]any{"a": "1"}}}},
	})
	streams, err := rdb.XReadStreams(ctx, "s1", "s2", "0", "0").Result()
	if err != nil || len(streams) != 2 || streams[0].Stream != "s2" || streams[1].Stream != "s1" {
		t.Fatalf("unexpected streams %v err %v", streams, err)
	}

	if err := cm.ExpectationsWereMet(); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestDoSetValRawValue(t *testing.T) {
	ctx := context.Background()
	rdb, cm := NewClientMock()
	defer rdb.Close()

	type rawPayload struct {
		Name  string
		Count int
	}

	want := rawPayload{Name: "node", Count: 2}
	cm.ExpectDo("CUSTOM", "k").SetVal(want)
	cmd := rdb.Do(ctx, "CUSTOM", "k")
	if got, err := cmd.RawResult(); err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected raw result %v err %v", got, err)
	}
	if got := cmd.RawVal(); !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected raw val %v", got)
	}

	cm.ExpectDo("PING").SetVal("PONG")
	cmd = rdb.Do(ctx, "PING")
	if got, err := cmd.RawResult(); err != nil || got != "PONG" {
		t.Fatalf("unexpected raw result %v err %v", got, err)
	}
	if got, err := cmd.Result(); err != nil || got != "PONG" {
		t.Fatalf("unexpected result %v err %v", got, err)
	}

	if err := cm.ExpectationsWereMet(); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestSetErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)
	anyErr := errors.New("any")

	cases := []struct {
		setup func()
		call  func() error
	}{
		{func() { cm.ExpectGet("k").SetErr(anyErr) }, func() error { return rdb.Get(ctx, "k").Err() }},
		{func() { cm.ExpectSet("k", "v", 0).SetErr(anyErr) }, func() error { return rdb.Set(ctx, "k", "v", 0).Err() }},
		{func() { cm.ExpectSetNX("k", "v", 0).SetErr(anyErr) }, func() error { return rdb.SetNX(ctx, "k", "v", 0).Err() }},
		{func() { cm.ExpectGetSet("k", "v").SetErr(anyErr) }, func() error { return rdb.GetSet(ctx, "k", "v").Err() }},
		{func() { cm.ExpectAppend("k", "v").SetErr(anyErr) }, func() error { return rdb.Append(ctx, "k", "v").Err() }},
		{func() { cm.ExpectStrLen("k").SetErr(anyErr) }, func() error { return rdb.StrLen(ctx, "k").Err() }},
		{func() { cm.ExpectIncr("k").SetErr(anyErr) }, func() error { return rdb.Incr(ctx, "k").Err() }},
		{func() { cm.ExpectIncrBy("k", 1).SetErr(anyErr) }, func() error { return rdb.IncrBy(ctx, "k", 1).Err() }},
		{func() { cm.ExpectDecr("k").SetErr(anyErr) }, func() error { return rdb.Decr(ctx, "k").Err() }},
		{func() { cm.ExpectDecrBy("k", 1).SetErr(anyErr) }, func() error { return rdb.DecrBy(ctx, "k", 1).Err() }},
		{func() { cm.ExpectMGet("k").SetErr(anyErr) }, func() error { return rdb.MGet(ctx, "k").Err() }},
		{func() { cm.ExpectMSet("k", "v").SetErr(anyErr) }, func() error { return rdb.MSet(ctx, "k", "v").Err() }},
		{func() { cm.ExpectDel("k").SetErr(anyErr) }, func() error { return rdb.Del(ctx, "k").Err() }},
		{func() { cm.ExpectExists("k").SetErr(anyErr) }, func() error { return rdb.Exists(ctx, "k").Err() }},
		{func() { cm.ExpectType("k").SetErr(anyErr) }, func() error { return rdb.Type(ctx, "k").Err() }},
		{func() { cm.ExpectExpire("k", time.Second).SetErr(anyErr) }, func() error { return rdb.Expire(ctx, "k", time.Second).Err() }},
		{func() { cm.ExpectTTL("k").SetErr(anyErr) }, func() error { return rdb.TTL(ctx, "k").Err() }},
		{func() { cm.ExpectPing().SetErr(anyErr) }, func() error { return rdb.Ping(ctx).Err() }},
		{func() { cm.ExpectEcho("hi").SetErr(anyErr) }, func() error { return rdb.Echo(ctx, "hi").Err() }},
		{func() { cm.ExpectHGet("h", "f").SetErr(anyErr) }, func() error { return rdb.HGet(ctx, "h", "f").Err() }},
		{func() { cm.ExpectHSet("h", "f", "v").SetErr(anyErr) }, func() error { return rdb.HSet(ctx, "h", "f", "v").Err() }},
		{func() { cm.ExpectHDel("h", "f").SetErr(anyErr) }, func() error { return rdb.HDel(ctx, "h", "f").Err() }},
		{func() { cm.ExpectHGetAll("h").SetErr(anyErr) }, func() error { return rdb.HGetAll(ctx, "h").Err() }},
		{func() { cm.ExpectLPush("l", "x").SetErr(anyErr) }, func() error { return rdb.LPush(ctx, "l", "x").Err() }},
		{func() { cm.ExpectRPush("l", "x").SetErr(anyErr) }, func() error { return rdb.RPush(ctx, "l", "x").Err() }},
		{func() { cm.ExpectLPop("l").SetErr(anyErr) }, func() error { return rdb.LPop(ctx, "l").Err() }},
		{func() { cm.ExpectRPop("l").SetErr(anyErr) }, func() error { return rdb.RPop(ctx, "l").Err() }},
		{func() { cm.ExpectLLen("l").SetErr(anyErr) }, func() error { return rdb.LLen(ctx, "l").Err() }},
		{func() { cm.ExpectSAdd("s", "x").SetErr(anyErr) }, func() error { return rdb.SAdd(ctx, "s", "x").Err() }},
		{func() { cm.ExpectSRem("s", "x").SetErr(anyErr) }, func() error { return rdb.SRem(ctx, "s", "x").Err() }},
		{func() { cm.ExpectSMembers("s").SetErr(anyErr) }, func() error { return rdb.SMembers(ctx, "s").Err() }},
		{func() { cm.ExpectEval("x", []string{}).SetErr(anyErr) }, func() error { return rdb.Eval(ctx, "x", []string{}).Err() }},
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

type hsetStruct struct {
	Name string `redis:"name"`
	Age  int    `redis:"age"`
}

type hsetStructWithOmit struct {
	Name string  `redis:"name"`
	Note string  `redis:"note,omitempty"`
	Ptr  *string `redis:"ptr"`
}

func TestHSetVariants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.ExpectHSet("h", map[string]string{"a": "1", "b": "2", "c": "3"}).SetVal(3)
		if _, err := rdb.HSet(ctx, "h", map[string]string{"c": "3", "a": "1", "b": "2"}).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.ExpectHSet("h", map[string]any{"a": 1, "b": "two"}).SetVal(2)
		if _, err := rdb.HSet(ctx, "h", map[string]any{"b": "two", "a": 1}).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.ExpectHSet("h", hsetStruct{Name: "alice", Age: 30}).SetVal(2)
		if _, err := rdb.HSet(ctx, "h", hsetStruct{Name: "alice", Age: 30}).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.ExpectHSet("h", &hsetStruct{Name: "bob", Age: 0}).SetVal(2)
		if _, err := rdb.HSet(ctx, "h", &hsetStruct{Name: "bob", Age: 0}).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.ExpectHSet("h", hsetStructWithOmit{Name: "x"}).SetVal(1)
		if _, err := rdb.HSet(ctx, "h", hsetStructWithOmit{Name: "x"}).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		s := "v"
		cm.ExpectHSet("h", hsetStructWithOmit{Name: "x", Ptr: &s}).SetVal(2)
		if _, err := rdb.HSet(ctx, "h", hsetStructWithOmit{Name: "x", Ptr: &s}).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.ExpectHSet("h", []any{"a", 1, "b", 2}).SetVal(2)
		if _, err := rdb.HSet(ctx, "h", []any{"a", 1, "b", 2}).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.ExpectHSet("h", []string{"a", "1"}).SetVal(1)
		if _, err := rdb.HSet(ctx, "h", []string{"a", "1"}).Result(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
}

func TestMSetVariants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.ExpectMSet(map[string]string{"a": "1", "b": "2"}).SetVal("OK")
		if err := rdb.MSet(ctx, map[string]string{"b": "2", "a": "1"}).Err(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.ExpectMSet(map[string]any{"a": 1, "b": "two"}).SetVal("OK")
		if err := rdb.MSet(ctx, map[string]any{"b": "two", "a": 1}).Err(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
	{
		raw := mock.NewClient(ctrl)
		cm := NewAdapter(raw)
		rdb := rueidiscompat.NewAdapter(raw)
		cm.ExpectMSet([]string{"a", "1", "b", "2"}).SetVal("OK")
		if err := rdb.MSet(ctx, []string{"a", "1", "b", "2"}).Err(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	}
}

func TestStringCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectGetEx("k", time.Second).SetVal("v")
			_, err := rdb.GetEx(ctx, "k", time.Second).Result()
			return err
		},
		func() error {
			cm.ExpectGetEx("k", 0).SetVal("v")
			_, err := rdb.GetEx(ctx, "k", 0).Result()
			return err
		},
		func() error {
			cm.ExpectGetEx("k", 50*time.Millisecond).SetVal("v")
			_, err := rdb.GetEx(ctx, "k", 50*time.Millisecond).Result()
			return err
		},
		func() error {
			cm.ExpectGetDel("k").SetVal("v")
			_, err := rdb.GetDel(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectGetRange("k", 0, 5).SetVal("v")
			_, err := rdb.GetRange(ctx, "k", 0, 5).Result()
			return err
		},
		func() error {
			cm.ExpectSetEx("k", "v", time.Second).SetVal("OK")
			return rdb.SetEX(ctx, "k", "v", time.Second).Err()
		},
		func() error {
			cm.ExpectSetXX("k", "v", time.Second).SetVal(true)
			_, err := rdb.SetXX(ctx, "k", "v", time.Second).Result()
			return err
		},
		func() error {
			cm.ExpectSetXX("k", "v", 50*time.Millisecond).SetVal(true)
			_, err := rdb.SetXX(ctx, "k", "v", 50*time.Millisecond).Result()
			return err
		},
		func() error {
			cm.ExpectSetXX("k", "v", -1).SetVal(true)
			_, err := rdb.SetXX(ctx, "k", "v", -1).Result()
			return err
		},
		func() error {
			cm.ExpectSetXX("k", "v", 0).SetVal(true)
			_, err := rdb.SetXX(ctx, "k", "v", 0).Result()
			return err
		},
		func() error {
			cm.ExpectSetRange("k", 0, "v").SetVal(1)
			_, err := rdb.SetRange(ctx, "k", 0, "v").Result()
			return err
		},
		func() error {
			cm.ExpectMSetNX("k1", "v1", "k2", "v2").SetVal(true)
			_, err := rdb.MSetNX(ctx, "k1", "v1", "k2", "v2").Result()
			return err
		},
		func() error {
			cm.ExpectIncrByFloat("k", 1.5).SetVal(2.5)
			_, err := rdb.IncrByFloat(ctx, "k", 1.5).Result()
			return err
		},
		func() error {
			cm.ExpectGetBit("k", 7).SetVal(1)
			_, err := rdb.GetBit(ctx, "k", 7).Result()
			return err
		},
		func() error {
			cm.ExpectSetBit("k", 7, 1).SetVal(0)
			_, err := rdb.SetBit(ctx, "k", 7, 1).Result()
			return err
		},
		func() error {
			cm.ExpectBitCount("k", &rueidiscompat.BitCount{Start: 0, End: 10}).SetVal(5)
			_, err := rdb.BitCount(ctx, "k", &rueidiscompat.BitCount{Start: 0, End: 10}).Result()
			return err
		},
		func() error {
			cm.ExpectBitOpAnd("dst", "k1", "k2").SetVal(3)
			_, err := rdb.BitOpAnd(ctx, "dst", "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectBitOpOr("dst", "k1", "k2").SetVal(3)
			_, err := rdb.BitOpOr(ctx, "dst", "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectBitOpXor("dst", "k1", "k2").SetVal(3)
			_, err := rdb.BitOpXor(ctx, "dst", "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectBitOpNot("dst", "k").SetVal(3)
			_, err := rdb.BitOpNot(ctx, "dst", "k").Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestGenericCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectUnlink("k1", "k2").SetVal(2)
			_, err := rdb.Unlink(ctx, "k1", "k2").Result()
			return err
		},
		func() error {
			now := time.Now()
			cm.ExpectExpireAt("k", now).SetVal(true)
			_, err := rdb.ExpireAt(ctx, "k", now).Result()
			return err
		},
		func() error {
			cm.ExpectExpireNX("k", time.Second).SetVal(true)
			_, err := rdb.ExpireNX(ctx, "k", time.Second).Result()
			return err
		},
		func() error {
			cm.ExpectExpireXX("k", time.Second).SetVal(true)
			_, err := rdb.ExpireXX(ctx, "k", time.Second).Result()
			return err
		},
		func() error {
			cm.ExpectExpireGT("k", time.Second).SetVal(true)
			_, err := rdb.ExpireGT(ctx, "k", time.Second).Result()
			return err
		},
		func() error {
			cm.ExpectExpireLT("k", time.Second).SetVal(true)
			_, err := rdb.ExpireLT(ctx, "k", time.Second).Result()
			return err
		},
		func() error {
			cm.ExpectExpireTime("k").SetVal(time.Hour)
			_, err := rdb.ExpireTime(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectPExpire("k", time.Second).SetVal(true)
			_, err := rdb.PExpire(ctx, "k", time.Second).Result()
			return err
		},
		func() error {
			now := time.Now()
			cm.ExpectPExpireAt("k", now).SetVal(true)
			_, err := rdb.PExpireAt(ctx, "k", now).Result()
			return err
		},
		func() error {
			cm.ExpectPExpireTime("k").SetVal(time.Hour)
			_, err := rdb.PExpireTime(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectPTTL("k").SetVal(time.Second)
			_, err := rdb.PTTL(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectPersist("k").SetVal(true)
			_, err := rdb.Persist(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectRandomKey().SetVal("k")
			_, err := rdb.RandomKey(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectRename("k", "n").SetVal("OK")
			return rdb.Rename(ctx, "k", "n").Err()
		},
		func() error {
			cm.ExpectRenameNX("k", "n").SetVal(true)
			_, err := rdb.RenameNX(ctx, "k", "n").Result()
			return err
		},
		func() error {
			cm.ExpectMove("k", 1).SetVal(true)
			_, err := rdb.Move(ctx, "k", 1).Result()
			return err
		},
		func() error {
			cm.ExpectObjectEncoding("k").SetVal("raw")
			_, err := rdb.ObjectEncoding(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectObjectIdleTime("k").SetVal(time.Second)
			_, err := rdb.ObjectIdleTime(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectObjectRefCount("k").SetVal(1)
			_, err := rdb.ObjectRefCount(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectTouch("k1", "k2").SetVal(2)
			_, err := rdb.Touch(ctx, "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectCopy("s", "d", 0, false).SetVal(1)
			_, err := rdb.Copy(ctx, "s", "d", 0, false).Result()
			return err
		},
		func() error {
			cm.ExpectCopy("s", "d", 0, true).SetVal(1)
			_, err := rdb.Copy(ctx, "s", "d", 0, true).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestHashCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectHExists("h", "f").SetVal(true)
			_, err := rdb.HExists(ctx, "h", "f").Result()
			return err
		},
		func() error {
			cm.ExpectHKeys("h").SetVal([]string{"f1", "f2"})
			_, err := rdb.HKeys(ctx, "h").Result()
			return err
		},
		func() error {
			cm.ExpectHLen("h").SetVal(2)
			_, err := rdb.HLen(ctx, "h").Result()
			return err
		},
		func() error {
			cm.ExpectHMGet("h", "f1", "f2").SetVal([]any{"v1", "v2"})
			_, err := rdb.HMGet(ctx, "h", "f1", "f2").Result()
			return err
		},
		func() error {
			cm.ExpectHMSet("h", "f1", "v1", "f2", "v2").SetVal(true)
			_, err := rdb.HMSet(ctx, "h", "f1", "v1", "f2", "v2").Result()
			return err
		},
		func() error {
			cm.ExpectHSetNX("h", "f", "v").SetVal(true)
			_, err := rdb.HSetNX(ctx, "h", "f", "v").Result()
			return err
		},
		func() error {
			cm.ExpectHVals("h").SetVal([]string{"v1", "v2"})
			_, err := rdb.HVals(ctx, "h").Result()
			return err
		},
		func() error {
			cm.ExpectHIncrBy("h", "f", 1).SetVal(2)
			_, err := rdb.HIncrBy(ctx, "h", "f", 1).Result()
			return err
		},
		func() error {
			cm.ExpectHIncrByFloat("h", "f", 1.5).SetVal(2.5)
			_, err := rdb.HIncrByFloat(ctx, "h", "f", 1.5).Result()
			return err
		},
		func() error {
			cm.ExpectHRandField("h", 2).SetVal([]string{"f1", "f2"})
			_, err := rdb.HRandField(ctx, "h", 2).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestListCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectLIndex("l", 0).SetVal("v")
			_, err := rdb.LIndex(ctx, "l", 0).Result()
			return err
		},
		func() error {
			cm.ExpectLRange("l", 0, -1).SetVal([]string{"v"})
			_, err := rdb.LRange(ctx, "l", 0, -1).Result()
			return err
		},
		func() error {
			cm.ExpectLRem("l", 0, "v").SetVal(1)
			_, err := rdb.LRem(ctx, "l", 0, "v").Result()
			return err
		},
		func() error {
			cm.ExpectLSet("l", 0, "v").SetVal("OK")
			return rdb.LSet(ctx, "l", 0, "v").Err()
		},
		func() error {
			cm.ExpectLTrim("l", 0, -1).SetVal("OK")
			return rdb.LTrim(ctx, "l", 0, -1).Err()
		},
		func() error {
			cm.ExpectLInsertBefore("l", "p", "v").SetVal(1)
			_, err := rdb.LInsertBefore(ctx, "l", "p", "v").Result()
			return err
		},
		func() error {
			cm.ExpectLInsertAfter("l", "p", "v").SetVal(1)
			_, err := rdb.LInsertAfter(ctx, "l", "p", "v").Result()
			return err
		},
		func() error {
			cm.ExpectLPushX("l", "v").SetVal(1)
			_, err := rdb.LPushX(ctx, "l", "v").Result()
			return err
		},
		func() error {
			cm.ExpectRPushX("l", "v").SetVal(1)
			_, err := rdb.RPushX(ctx, "l", "v").Result()
			return err
		},
		func() error {
			cm.ExpectRPopLPush("s", "d").SetVal("v")
			_, err := rdb.RPopLPush(ctx, "s", "d").Result()
			return err
		},
		func() error {
			cm.ExpectBLPop(time.Second, "k1").SetVal([]string{"k1", "v"})
			_, err := rdb.BLPop(ctx, time.Second, "k1").Result()
			return err
		},
		func() error {
			cm.ExpectBRPop(time.Second, "k1").SetVal([]string{"k1", "v"})
			_, err := rdb.BRPop(ctx, time.Second, "k1").Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestSetCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectSCard("s").SetVal(2)
			_, err := rdb.SCard(ctx, "s").Result()
			return err
		},
		func() error {
			cm.ExpectSDiff("s1", "s2").SetVal([]string{"a"})
			_, err := rdb.SDiff(ctx, "s1", "s2").Result()
			return err
		},
		func() error {
			cm.ExpectSDiffStore("d", "s1", "s2").SetVal(1)
			_, err := rdb.SDiffStore(ctx, "d", "s1", "s2").Result()
			return err
		},
		func() error {
			cm.ExpectSInter("s1", "s2").SetVal([]string{"a"})
			_, err := rdb.SInter(ctx, "s1", "s2").Result()
			return err
		},
		func() error {
			cm.ExpectSInterStore("d", "s1", "s2").SetVal(1)
			_, err := rdb.SInterStore(ctx, "d", "s1", "s2").Result()
			return err
		},
		func() error {
			cm.ExpectSIsMember("s", "m").SetVal(true)
			_, err := rdb.SIsMember(ctx, "s", "m").Result()
			return err
		},
		func() error {
			cm.ExpectSMIsMember("s", "m1", "m2").SetVal([]bool{true, false})
			_, err := rdb.SMIsMember(ctx, "s", "m1", "m2").Result()
			return err
		},
		func() error {
			cm.ExpectSMove("s", "d", "m").SetVal(true)
			_, err := rdb.SMove(ctx, "s", "d", "m").Result()
			return err
		},
		func() error {
			cm.ExpectSPop("s").SetVal("m")
			_, err := rdb.SPop(ctx, "s").Result()
			return err
		},
		func() error {
			cm.ExpectSPopN("s", 2).SetVal([]string{"a", "b"})
			_, err := rdb.SPopN(ctx, "s", 2).Result()
			return err
		},
		func() error {
			cm.ExpectSRandMember("s").SetVal("m")
			_, err := rdb.SRandMember(ctx, "s").Result()
			return err
		},
		func() error {
			cm.ExpectSRandMemberN("s", 2).SetVal([]string{"a", "b"})
			_, err := rdb.SRandMemberN(ctx, "s", 2).Result()
			return err
		},
		func() error {
			cm.ExpectSUnion("s1", "s2").SetVal([]string{"a", "b"})
			_, err := rdb.SUnion(ctx, "s1", "s2").Result()
			return err
		},
		func() error {
			cm.ExpectSUnionStore("d", "s1", "s2").SetVal(2)
			_, err := rdb.SUnionStore(ctx, "d", "s1", "s2").Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestSortedSetCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectZCard("z").SetVal(2)
			_, err := rdb.ZCard(ctx, "z").Result()
			return err
		},
		func() error {
			cm.ExpectZCount("z", "0", "10").SetVal(2)
			_, err := rdb.ZCount(ctx, "z", "0", "10").Result()
			return err
		},
		func() error {
			cm.ExpectZIncrBy("z", 1.0, "m").SetVal(2.0)
			_, err := rdb.ZIncrBy(ctx, "z", 1.0, "m").Result()
			return err
		},
		func() error {
			cm.ExpectZLexCount("z", "[a", "[z").SetVal(2)
			_, err := rdb.ZLexCount(ctx, "z", "[a", "[z").Result()
			return err
		},
		func() error {
			cm.ExpectZRange("z", 0, -1).SetVal([]string{"a", "b"})
			_, err := rdb.ZRange(ctx, "z", 0, -1).Result()
			return err
		},
		func() error {
			cm.ExpectZRevRange("z", 0, -1).SetVal([]string{"b", "a"})
			_, err := rdb.ZRevRange(ctx, "z", 0, -1).Result()
			return err
		},
		func() error {
			cm.ExpectZRank("z", "m").SetVal(0)
			_, err := rdb.ZRank(ctx, "z", "m").Result()
			return err
		},
		func() error {
			cm.ExpectZRevRank("z", "m").SetVal(0)
			_, err := rdb.ZRevRank(ctx, "z", "m").Result()
			return err
		},
		func() error {
			cm.ExpectZRem("z", "m1", "m2").SetVal(2)
			_, err := rdb.ZRem(ctx, "z", "m1", "m2").Result()
			return err
		},
		func() error {
			cm.ExpectZRemRangeByLex("z", "[a", "[z").SetVal(2)
			_, err := rdb.ZRemRangeByLex(ctx, "z", "[a", "[z").Result()
			return err
		},
		func() error {
			cm.ExpectZRemRangeByRank("z", 0, -1).SetVal(2)
			_, err := rdb.ZRemRangeByRank(ctx, "z", 0, -1).Result()
			return err
		},
		func() error {
			cm.ExpectZRemRangeByScore("z", "0", "10").SetVal(2)
			_, err := rdb.ZRemRangeByScore(ctx, "z", "0", "10").Result()
			return err
		},
		func() error {
			cm.ExpectZScore("z", "m").SetVal(1.5)
			_, err := rdb.ZScore(ctx, "z", "m").Result()
			return err
		},
		func() error {
			cm.ExpectZMScore("z", "m1", "m2").SetVal([]float64{1.0, 2.0})
			_, err := rdb.ZMScore(ctx, "z", "m1", "m2").Result()
			return err
		},
		func() error {
			cm.ExpectZPopMax("z").SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZPopMax(ctx, "z").Result()
			return err
		},
		func() error {
			cm.ExpectZPopMax("z", 2).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZPopMax(ctx, "z", 2).Result()
			return err
		},
		func() error {
			cm.ExpectZPopMin("z").SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZPopMin(ctx, "z").Result()
			return err
		},
		func() error {
			cm.ExpectZPopMin("z", 2).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZPopMin(ctx, "z", 2).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestServerCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectLastSave().SetVal(1)
			_, err := rdb.LastSave(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectTime().SetVal(time.Unix(1, 0))
			_, err := rdb.Time(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectInfo().SetVal("info")
			_, err := rdb.Info(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectInfo("server", "memory").SetVal("info")
			_, err := rdb.Info(ctx, "server", "memory").Result()
			return err
		},
		func() error {
			cm.ExpectClientID().SetVal(1)
			_, err := rdb.ClientID(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectClientGetName().SetVal("name")
			_, err := rdb.ClientGetName(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectReadOnly().SetVal("OK")
			return rdb.ReadOnly(ctx).Err()
		},
		func() error {
			cm.ExpectReadWrite().SetVal("OK")
			return rdb.ReadWrite(ctx).Err()
		},
	}
	runCases(t, cm, cases)
}

func TestPubSubCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectPublish("ch", "msg").SetVal(1)
			_, err := rdb.Publish(ctx, "ch", "msg").Result()
			return err
		},
		func() error {
			cm.ExpectSPublish("ch", "msg").SetVal(1)
			_, err := rdb.SPublish(ctx, "ch", "msg").Result()
			return err
		},
		func() error {
			cm.ExpectPubSubChannels("p*").SetVal([]string{"ch1"})
			_, err := rdb.PubSubChannels(ctx, "p*").Result()
			return err
		},
		func() error {
			cm.ExpectPubSubNumPat().SetVal(1)
			_, err := rdb.PubSubNumPat(ctx).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestHyperLogLogCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectPFAdd("k", "v").SetVal(1)
			_, err := rdb.PFAdd(ctx, "k", "v").Result()
			return err
		},
		func() error {
			cm.ExpectPFCount("k1", "k2").SetVal(2)
			_, err := rdb.PFCount(ctx, "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectPFMerge("d", "k1", "k2").SetVal("OK")
			return rdb.PFMerge(ctx, "d", "k1", "k2").Err()
		},
	}
	runCases(t, cm, cases)
}

func TestScriptingCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectEvalSha("sha", []string{"k"}, "a").SetVal("ok")
			_, err := rdb.EvalSha(ctx, "sha", []string{"k"}, "a").Result()
			return err
		},
		func() error {
			cm.ExpectEvalRO("script", []string{"k"}, "a").SetVal("ok")
			_, err := rdb.EvalRO(ctx, "script", []string{"k"}, "a").Result()
			return err
		},
		func() error {
			cm.ExpectEvalShaRO("sha", []string{"k"}, "a").SetVal("ok")
			_, err := rdb.EvalShaRO(ctx, "sha", []string{"k"}, "a").Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestClusterCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectClusterInfo().SetVal("ok")
			_, err := rdb.ClusterInfo(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectClusterNodes().SetVal("ok")
			_, err := rdb.ClusterNodes(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectClusterKeySlot("k").SetVal(123)
			_, err := rdb.ClusterKeySlot(ctx, "k").Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestServerAdminCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	raw.EXPECT().Nodes().Return(map[string]rueidis.Client{"mock": raw}).AnyTimes()
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectBgRewriteAOF().SetVal("OK")
			return rdb.BgRewriteAOF(ctx).Err()
		},
		func() error {
			cm.ExpectBgSave().SetVal("OK")
			return rdb.BgSave(ctx).Err()
		},
		func() error {
			cm.ExpectConfigGet("maxmemory").SetVal(map[string]string{"maxmemory": "0"})
			_, err := rdb.ConfigGet(ctx, "maxmemory").Result()
			return err
		},
		func() error {
			cm.ExpectConfigResetStat().SetVal("OK")
			return rdb.ConfigResetStat(ctx).Err()
		},
		func() error {
			cm.ExpectConfigRewrite().SetVal("OK")
			return rdb.ConfigRewrite(ctx).Err()
		},
		func() error {
			cm.ExpectConfigSet("maxmemory", "0").SetVal("OK")
			return rdb.ConfigSet(ctx, "maxmemory", "0").Err()
		},
		func() error {
			cm.ExpectDBSize().SetVal(1)
			_, err := rdb.DBSize(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectDebugObject("k").SetVal("debug")
			_, err := rdb.DebugObject(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectFlushAll().SetVal("OK")
			return rdb.FlushAll(ctx).Err()
		},
		func() error {
			cm.ExpectFlushAllAsync().SetVal("OK")
			return rdb.FlushAllAsync(ctx).Err()
		},
		func() error {
			cm.ExpectFlushDB().SetVal("OK")
			return rdb.FlushDB(ctx).Err()
		},
		func() error {
			cm.ExpectFlushDBAsync().SetVal("OK")
			return rdb.FlushDBAsync(ctx).Err()
		},
		func() error {
			cm.ExpectMemoryUsage("k").SetVal(100)
			_, err := rdb.MemoryUsage(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectSave().SetVal("OK")
			return rdb.Save(ctx).Err()
		},
		func() error {
			cm.ExpectShutdown().SetVal("OK")
			return rdb.Shutdown(ctx).Err()
		},
		func() error {
			cm.ExpectShutdownNoSave().SetVal("OK")
			return rdb.ShutdownNoSave(ctx).Err()
		},
		func() error {
			cm.ExpectShutdownSave().SetVal("OK")
			return rdb.ShutdownSave(ctx).Err()
		},
		func() error {
			cm.ExpectSlaveOf("127.0.0.1", "6379").SetVal("OK")
			return rdb.SlaveOf(ctx, "127.0.0.1", "6379").Err()
		},
		func() error {
			cm.ExpectSlowLogGet(10).SetVal([]rueidiscompat.SlowLog{{
				ID:       1,
				Time:     time.Unix(1, 0),
				Duration: time.Microsecond,
				Args:     []string{"GET", "k"},
			}})
			_, err := rdb.SlowLogGet(ctx, 10).Result()
			return err
		},
		func() error {
			cm.ExpectQuit().SetVal("OK")
			return rdb.Quit(ctx).Err()
		},
		func() error {
			cm.ExpectCommand().SetVal([]*rueidiscompat.CommandInfo{{Name: "get"}})
			_, err := rdb.Command(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectCommandGetKeys("GET", "k").SetVal([]string{"k"})
			_, err := rdb.CommandGetKeys(ctx, "GET", "k").Result()
			return err
		},
		func() error {
			cm.ExpectCommandGetKeysAndFlags("SET", "k", "v").SetVal([]rueidiscompat.KeyFlags{{Key: "k", Flags: []string{"RW"}}})
			_, err := rdb.CommandGetKeysAndFlags(ctx, "SET", "k", "v").Result()
			return err
		},
		func() error {
			cm.ExpectCommandList(&rueidiscompat.FilterBy{Module: "mod"}).SetVal([]string{"cmd"})
			_, err := rdb.CommandList(ctx, rueidiscompat.FilterBy{Module: "mod"}).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestClientClusterCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	raw.EXPECT().Nodes().Return(map[string]rueidis.Client{"mock": raw}).AnyTimes()
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectClientKill("127.0.0.1:6379").SetVal("OK")
			return rdb.ClientKill(ctx, "127.0.0.1:6379").Err()
		},
		func() error {
			cm.ExpectClientKillByFilter("ID", "1").SetVal(1)
			_, err := rdb.ClientKillByFilter(ctx, "ID", "1").Result()
			return err
		},
		func() error {
			cm.ExpectClientList().SetVal("id=1")
			_, err := rdb.ClientList(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectClientPause(time.Second).SetVal(true)
			_, err := rdb.ClientPause(ctx, time.Second).Result()
			return err
		},
		func() error {
			cm.ExpectClientUnpause().SetVal(true)
			_, err := rdb.ClientUnpause(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectClientUnblock(1).SetVal(1)
			_, err := rdb.ClientUnblock(ctx, 1).Result()
			return err
		},
		func() error {
			cm.ExpectClientUnblockWithError(1).SetVal(0)
			_, err := rdb.ClientUnblockWithError(ctx, 1).Result()
			return err
		},
		func() error {
			cm.ExpectClusterAddSlots(1, 2).SetVal("OK")
			return rdb.ClusterAddSlots(ctx, 1, 2).Err()
		},
		func() error {
			cm.ExpectClusterAddSlotsRange(1, 10).SetVal("OK")
			return rdb.ClusterAddSlotsRange(ctx, 1, 10).Err()
		},
		func() error {
			cm.ExpectClusterCountFailureReports("node").SetVal(0)
			_, err := rdb.ClusterCountFailureReports(ctx, "node").Result()
			return err
		},
		func() error {
			cm.ExpectClusterCountKeysInSlot(1).SetVal(3)
			_, err := rdb.ClusterCountKeysInSlot(ctx, 1).Result()
			return err
		},
		func() error {
			cm.ExpectClusterDelSlots(1, 2).SetVal("OK")
			return rdb.ClusterDelSlots(ctx, 1, 2).Err()
		},
		func() error {
			cm.ExpectClusterDelSlotsRange(1, 10).SetVal("OK")
			return rdb.ClusterDelSlotsRange(ctx, 1, 10).Err()
		},
		func() error {
			cm.ExpectClusterFailover().SetVal("OK")
			return rdb.ClusterFailover(ctx).Err()
		},
		func() error {
			cm.ExpectClusterForget("node").SetVal("OK")
			return rdb.ClusterForget(ctx, "node").Err()
		},
		func() error {
			cm.ExpectClusterGetKeysInSlot(1, 10).SetVal([]string{"k"})
			_, err := rdb.ClusterGetKeysInSlot(ctx, 1, 10).Result()
			return err
		},
		func() error {
			cm.ExpectClusterLinks().SetVal([]rueidiscompat.ClusterLink{{
				Direction:           "to",
				Node:                "abc",
				CreateTime:          1,
				Events:              "r",
				SendBufferAllocated: 0,
				SendBufferUsed:      0,
			}})
			_, err := rdb.ClusterLinks(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectClusterMeet("127.0.0.1", "6379").SetVal("OK")
			return rdb.ClusterMeet(ctx, "127.0.0.1", 6379).Err()
		},
		func() error {
			cm.ExpectClusterReplicate("node").SetVal("OK")
			return rdb.ClusterReplicate(ctx, "node").Err()
		},
		func() error {
			cm.ExpectClusterResetHard().SetVal("OK")
			return rdb.ClusterResetHard(ctx).Err()
		},
		func() error {
			cm.ExpectClusterResetSoft().SetVal("OK")
			return rdb.ClusterResetSoft(ctx).Err()
		},
		func() error {
			cm.ExpectClusterSaveConfig().SetVal("OK")
			return rdb.ClusterSaveConfig(ctx).Err()
		},
		func() error {
			cm.ExpectClusterShards().SetVal([]rueidiscompat.ClusterShard{{
				Slots: []rueidiscompat.SlotRange{{Start: 0, End: 16383}},
				Nodes: []rueidiscompat.Node{{ID: "abc", Role: "master", Port: 6379}},
			}})
			_, err := rdb.ClusterShards(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectClusterSlaves("node").SetVal([]string{"slave"})
			_, err := rdb.ClusterSlaves(ctx, "node").Result()
			return err
		},
		func() error {
			cm.ExpectClusterSlots().SetVal([]rueidiscompat.ClusterSlot{{
				Start: 0,
				End:   16383,
				Nodes: []rueidiscompat.ClusterNode{{Addr: "127.0.0.1:6379", ID: "abc"}},
			}})
			_, err := rdb.ClusterSlots(ctx).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestGenericBitmapScanSortCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	raw.EXPECT().Nodes().Return(map[string]rueidis.Client{"mock": raw}).AnyTimes()
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	sort := rueidiscompat.Sort{Order: "ASC"}

	cases := []func() error{
		func() error {
			cm.ExpectACLDryRun("default", "get", "randomKey").SetVal("OK")
			_, err := rdb.ACLDryRun(ctx, "default", "get", "randomKey").Result()
			return err
		},
		func() error {
			cm.ExpectBitField("k", "GET", "u8", 0).SetVal([]int64{255})
			_, err := rdb.BitField(ctx, "k", "GET", "u8", 0).Result()
			return err
		},
		func() error {
			cm.ExpectBitPos("k", 1).SetVal(3)
			_, err := rdb.BitPos(ctx, "k", 1).Result()
			return err
		},
		func() error {
			cm.ExpectBitPos("k", 1, 0).SetVal(3)
			_, err := rdb.BitPos(ctx, "k", 1, 0).Result()
			return err
		},
		func() error {
			cm.ExpectBitPos("k", 1, 0, 10).SetVal(3)
			_, err := rdb.BitPos(ctx, "k", 1, 0, 10).Result()
			return err
		},
		func() error {
			cm.ExpectBitPosSpan("k", 1, 0, 10, "bit").SetVal(3)
			_, err := rdb.BitPosSpan(ctx, "k", 1, 0, 10, "bit").Result()
			return err
		},
		func() error {
			cm.ExpectBitPosSpan("k", 1, 0, 10, "byte").SetVal(3)
			_, err := rdb.BitPosSpan(ctx, "k", 1, 0, 10, "byte").Result()
			return err
		},
		func() error {
			cm.ExpectDump("k").SetVal("dump")
			_, err := rdb.Dump(ctx, "k").Result()
			return err
		},
		func() error {
			cm.ExpectKeys("*").SetVal([]string{"k1"})
			_, err := rdb.Keys(ctx, "*").Result()
			return err
		},
		func() error {
			cm.ExpectMigrate("127.0.0.1", "6379", "k", 0, time.Second).SetVal("OK")
			return rdb.Migrate(ctx, "127.0.0.1", 6379, "k", 0, time.Second).Err()
		},
		func() error {
			cm.ExpectRestore("k", time.Hour, "val").SetVal("OK")
			return rdb.Restore(ctx, "k", time.Hour, "val").Err()
		},
		func() error {
			cm.ExpectRestoreReplace("k", time.Hour, "val").SetVal("OK")
			return rdb.RestoreReplace(ctx, "k", time.Hour, "val").Err()
		},
		func() error {
			cm.ExpectScan(0, "k*", 10).SetVal([]string{"k1"}, 0)
			_, _, err := rdb.Scan(ctx, 0, "k*", 10).Result()
			return err
		},
		func() error {
			cm.ExpectScanType(0, "k*", 10, "string").SetVal([]string{"k1"}, 0)
			_, _, err := rdb.ScanType(ctx, 0, "k*", 10, "string").Result()
			return err
		},
		func() error {
			cm.ExpectHScan("h", 0, "f*", 10).SetVal([]string{"f1", "v1"}, 0)
			_, _, err := rdb.HScan(ctx, "h", 0, "f*", 10).Result()
			return err
		},
		func() error {
			cm.ExpectSScan("s", 0, "m*", 10).SetVal([]string{"m1"}, 0)
			_, _, err := rdb.SScan(ctx, "s", 0, "m*", 10).Result()
			return err
		},
		func() error {
			cm.ExpectSort("k", &sort).SetVal([]string{"a", "b"})
			_, err := rdb.Sort(ctx, "k", sort).Result()
			return err
		},
		func() error {
			cm.ExpectSortInterfaces("k", &sort).SetVal([]any{"a", "b"})
			_, err := rdb.SortInterfaces(ctx, "k", sort).Result()
			return err
		},
		func() error {
			cm.ExpectSortRO("k", &sort).SetVal([]string{"a", "b"})
			_, err := rdb.SortRO(ctx, "k", sort).Result()
			return err
		},
		func() error {
			cm.ExpectSortStore("k", "dest", &sort).SetVal(2)
			_, err := rdb.SortStore(ctx, "k", "dest", sort).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestListSetHashStringCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectBLMPop(time.Second, "LEFT", 2, "k1", "k2").SetVal("k1", []string{"v1", "v2"})
			_, _, err := rdb.BLMPop(ctx, time.Second, "LEFT", 2, "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectBLMove("src", "dst", "LEFT", "RIGHT", time.Second).SetVal("v")
			_, err := rdb.BLMove(ctx, "src", "dst", "LEFT", "RIGHT", time.Second).Result()
			return err
		},
		func() error {
			cm.ExpectBRPopLPush("src", "dst", time.Second).SetVal("v")
			_, err := rdb.BRPopLPush(ctx, "src", "dst", time.Second).Result()
			return err
		},
		func() error {
			cm.ExpectHRandFieldWithValues("h", 2).SetVal([]rueidiscompat.KeyValue{{Key: "f1", Value: "v1"}})
			_, err := rdb.HRandFieldWithValues(ctx, "h", 2).Result()
			return err
		},
		func() error {
			q := &rueidiscompat.LCSQuery{Key1: "k1", Key2: "k2"}
			cm.ExpectLCS(q).SetVal(&rueidiscompat.LCSMatch{MatchString: "match"})
			_, err := rdb.LCS(ctx, q).Result()
			return err
		},
		func() error {
			q := &rueidiscompat.LCSQuery{Key1: "k1", Key2: "k2", Len: true}
			cm.ExpectLCS(q).SetVal(&rueidiscompat.LCSMatch{Len: 5})
			_, err := rdb.LCS(ctx, q).Result()
			return err
		},
		func() error {
			q := &rueidiscompat.LCSQuery{Key1: "k1", Key2: "k2", Idx: true, MinMatchLen: 4, WithMatchLen: true}
			cm.ExpectLCS(q).SetVal(&rueidiscompat.LCSMatch{Len: 6, Matches: []rueidiscompat.LCSMatchedPosition{{Key1: rueidiscompat.LCSPosition{Start: 0, End: 3}, Key2: rueidiscompat.LCSPosition{Start: 1, End: 4}, MatchLen: 4}}})
			_, err := rdb.LCS(ctx, q).Result()
			return err
		},
		func() error {
			cm.ExpectLInsert("l", "BEFORE", "p", "v").SetVal(1)
			_, err := rdb.LInsert(ctx, "l", "BEFORE", "p", "v").Result()
			return err
		},
		func() error {
			cm.ExpectLMPop("LEFT", 2, "k1", "k2").SetVal("k1", []string{"v1", "v2"})
			_, _, err := rdb.LMPop(ctx, "LEFT", 2, "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectLMove("src", "dst", "LEFT", "RIGHT").SetVal("v")
			_, err := rdb.LMove(ctx, "src", "dst", "LEFT", "RIGHT").Result()
			return err
		},
		func() error {
			cm.ExpectLPopCount("l", 2).SetVal([]string{"v1", "v2"})
			_, err := rdb.LPopCount(ctx, "l", 2).Result()
			return err
		},
		func() error {
			cm.ExpectLPos("l", "v", rueidiscompat.LPosArgs{Rank: 1}).SetVal(0)
			_, err := rdb.LPos(ctx, "l", "v", rueidiscompat.LPosArgs{Rank: 1}).Result()
			return err
		},
		func() error {
			cm.ExpectLPosCount("l", "v", 2, rueidiscompat.LPosArgs{Rank: 1}).SetVal([]int64{0, 3})
			_, err := rdb.LPosCount(ctx, "l", "v", 2, rueidiscompat.LPosArgs{Rank: 1}).Result()
			return err
		},
		func() error {
			cm.ExpectRPopCount("l", 2).SetVal([]string{"v1", "v2"})
			_, err := rdb.RPopCount(ctx, "l", 2).Result()
			return err
		},
		func() error {
			cm.ExpectSInterCard(0, "s1", "s2").SetVal(1)
			_, err := rdb.SInterCard(ctx, 0, "s1", "s2").Result()
			return err
		},
		func() error {
			cm.ExpectSMembersMap("s").SetVal([]string{"m1", "m2"})
			_, err := rdb.SMembersMap(ctx, "s").Result()
			return err
		},
		func() error {
			cm.ExpectSetArgs("k", "v", rueidiscompat.SetArgs{TTL: time.Second, Mode: "NX"}).SetVal("OK")
			return rdb.SetArgs(ctx, "k", "v", rueidiscompat.SetArgs{TTL: time.Second, Mode: "NX"}).Err()
		},
	}
	runCases(t, cm, cases)
}

func TestGeoCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	loc := rueidiscompat.GeoLocation{Name: "k1", Longitude: 122.4194, Latitude: 37.7749}
	radiusQuery := rueidiscompat.GeoRadiusQuery{Radius: 200}
	radiusStoreQuery := rueidiscompat.GeoRadiusQuery{Radius: 200, StoreDist: "result"}
	searchQuery := rueidiscompat.GeoSearchQuery{Member: "Catania", BoxWidth: 400, BoxHeight: 100, BoxUnit: "km", Sort: "asc"}
	searchLocationQuery := rueidiscompat.GeoSearchLocationQuery{
		GeoSearchQuery: rueidiscompat.GeoSearchQuery{Longitude: 15, Latitude: 37, Radius: 200, RadiusUnit: "km", Sort: "asc"},
		WithCoord:      true,
		WithDist:       true,
		WithHash:       true,
	}
	searchStoreQuery := rueidiscompat.GeoSearchStoreQuery{
		GeoSearchQuery: rueidiscompat.GeoSearchQuery{Longitude: 15, Latitude: 37, Radius: 200, RadiusUnit: "km", Sort: "asc"},
		StoreDist:      false,
	}

	cases := []func() error{
		func() error {
			cm.ExpectGeoAdd("1", &loc).SetVal(1)
			_, err := rdb.GeoAdd(ctx, "1", loc).Result()
			return err
		},
		func() error {
			cm.ExpectGeoDist("1", "2", "3", "M").SetVal(100.5)
			_, err := rdb.GeoDist(ctx, "1", "2", "3", "M").Result()
			return err
		},
		func() error {
			cm.ExpectGeoHash("1", "2", "3").SetVal([]string{"hash1", "hash2"})
			_, err := rdb.GeoHash(ctx, "1", "2", "3").Result()
			return err
		},
		func() error {
			cm.ExpectGeoPos("1", "1", "2").SetVal([]*rueidiscompat.GeoPos{{Longitude: 1, Latitude: 2}})
			_, err := rdb.GeoPos(ctx, "1", "1", "2").Result()
			return err
		},
		func() error {
			cm.ExpectGeoRadius("1", 1, 2, &radiusQuery).SetVal([]rueidiscompat.GeoLocation{{Name: "m1"}})
			_, err := rdb.GeoRadius(ctx, "1", 1, 2, radiusQuery).Result()
			return err
		},
		func() error {
			cm.ExpectGeoRadiusByMember("1", "2", &radiusQuery).SetVal([]rueidiscompat.GeoLocation{{Name: "m1"}})
			_, err := rdb.GeoRadiusByMember(ctx, "1", "2", radiusQuery).Result()
			return err
		},
		func() error {
			cm.ExpectGeoRadiusByMemberStore("1", "2", &radiusStoreQuery).SetVal(3)
			_, err := rdb.GeoRadiusByMemberStore(ctx, "1", "2", radiusStoreQuery).Result()
			return err
		},
		func() error {
			cm.ExpectGeoRadiusStore("1", 1, 2, &radiusStoreQuery).SetVal(3)
			_, err := rdb.GeoRadiusStore(ctx, "1", 1, 2, radiusStoreQuery).Result()
			return err
		},
		func() error {
			cm.ExpectGeoSearch("1", &searchQuery).SetVal([]string{"member1"})
			_, err := rdb.GeoSearch(ctx, "1", searchQuery).Result()
			return err
		},
		func() error {
			cm.ExpectGeoSearchLocation("1", &searchLocationQuery).SetVal([]rueidiscompat.GeoLocation{{Name: "m1"}})
			_, err := rdb.GeoSearchLocation(ctx, "1", searchLocationQuery).Result()
			return err
		},
		func() error {
			cm.ExpectGeoSearchStore("1", "2", &searchStoreQuery).SetVal(3)
			_, err := rdb.GeoSearchStore(ctx, "1", "2", searchStoreQuery).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestSortedSetCommandsExtra(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	opt := rueidiscompat.ZRangeBy{Min: "0", Max: "10"}
	optLimit := rueidiscompat.ZRangeBy{Min: "0", Max: "10", Offset: 1, Count: 2}
	rangeArgs := rueidiscompat.ZRangeArgs{Key: "z", Start: int64(0), Stop: int64(-1)}

	cases := []func() error{
		func() error {
			cm.ExpectBZMPop(time.Second, "MIN", 1, "z").SetVal("z", []rueidiscompat.Z{{Member: "m", Score: 1}})
			_, _, err := rdb.BZMPop(ctx, time.Second, "MIN", 1, "z").Result()
			return err
		},
		func() error {
			cm.ExpectBZPopMax(time.Second, "z").SetVal(&rueidiscompat.ZWithKey{Key: "z", Z: rueidiscompat.Z{Member: "m", Score: 1}})
			_, err := rdb.BZPopMax(ctx, time.Second, "z").Result()
			return err
		},
		func() error {
			cm.ExpectBZPopMin(time.Second, "z").SetVal(&rueidiscompat.ZWithKey{Key: "z", Z: rueidiscompat.Z{Member: "m", Score: 1}})
			_, err := rdb.BZPopMin(ctx, time.Second, "z").Result()
			return err
		},
		func() error {
			cm.ExpectZDiff("k1", "k2").SetVal([]string{"a"})
			_, err := rdb.ZDiff(ctx, "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectZDiffStore("dst", "k1", "k2").SetVal(1)
			_, err := rdb.ZDiffStore(ctx, "dst", "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectZDiffWithScores("k1", "k2").SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZDiffWithScores(ctx, "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectZInterCard(0, "k1", "k2").SetVal(1)
			_, err := rdb.ZInterCard(ctx, 0, "k1", "k2").Result()
			return err
		},
		func() error {
			cm.ExpectZMPop("MIN", 1, "z").SetVal("z", []rueidiscompat.Z{{Member: "m", Score: 1}})
			_, _, err := rdb.ZMPop(ctx, "MIN", 1, "z").Result()
			return err
		},
		func() error {
			cm.ExpectZRandMember("z", 2).SetVal([]string{"a", "b"})
			_, err := rdb.ZRandMember(ctx, "z", 2).Result()
			return err
		},
		func() error {
			cm.ExpectZRandMemberWithScores("z", 2).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZRandMemberWithScores(ctx, "z", 2).Result()
			return err
		},
		func() error {
			cm.ExpectZRangeByLex("z", &opt).SetVal([]string{"a"})
			_, err := rdb.ZRangeByLex(ctx, "z", opt).Result()
			return err
		},
		func() error {
			cm.ExpectZRangeByLex("z", &optLimit).SetVal([]string{"a"})
			_, err := rdb.ZRangeByLex(ctx, "z", optLimit).Result()
			return err
		},
		func() error {
			cm.ExpectZRangeByScore("z", &opt).SetVal([]string{"a"})
			_, err := rdb.ZRangeByScore(ctx, "z", opt).Result()
			return err
		},
		func() error {
			cm.ExpectZRangeByScore("z", &optLimit).SetVal([]string{"a"})
			_, err := rdb.ZRangeByScore(ctx, "z", optLimit).Result()
			return err
		},
		func() error {
			cm.ExpectZRangeByScoreWithScores("z", &opt).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZRangeByScoreWithScores(ctx, "z", opt).Result()
			return err
		},
		func() error {
			cm.ExpectZRangeByScoreWithScores("z", &optLimit).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZRangeByScoreWithScores(ctx, "z", optLimit).Result()
			return err
		},
		func() error {
			cm.ExpectZRangeStore("dst", rangeArgs).SetVal(1)
			_, err := rdb.ZRangeStore(ctx, "dst", rangeArgs).Result()
			return err
		},
		func() error {
			cm.ExpectZRevRangeByLex("z", &opt).SetVal([]string{"a"})
			_, err := rdb.ZRevRangeByLex(ctx, "z", opt).Result()
			return err
		},
		func() error {
			cm.ExpectZRevRangeByLex("z", &optLimit).SetVal([]string{"a"})
			_, err := rdb.ZRevRangeByLex(ctx, "z", optLimit).Result()
			return err
		},
		func() error {
			cm.ExpectZRevRangeByScore("z", &opt).SetVal([]string{"a"})
			_, err := rdb.ZRevRangeByScore(ctx, "z", opt).Result()
			return err
		},
		func() error {
			cm.ExpectZRevRangeByScore("z", &optLimit).SetVal([]string{"a"})
			_, err := rdb.ZRevRangeByScore(ctx, "z", optLimit).Result()
			return err
		},
		func() error {
			cm.ExpectZRevRangeByScoreWithScores("z", &opt).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZRevRangeByScoreWithScores(ctx, "z", opt).Result()
			return err
		},
		func() error {
			cm.ExpectZRevRangeByScoreWithScores("z", &optLimit).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZRevRangeByScoreWithScores(ctx, "z", optLimit).Result()
			return err
		},
		func() error {
			cm.ExpectZRevRangeWithScores("z", 0, -1).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZRevRangeWithScores(ctx, "z", 0, -1).Result()
			return err
		},
		func() error {
			cm.ExpectZScan("z", 0, "m*", 10).SetVal([]string{"m1"}, 0)
			_, _, err := rdb.ZScan(ctx, "z", 0, "m*", 10).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestSortedSetHelpersCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	member := rueidiscompat.Z{Member: "m", Score: 1}
	store := rueidiscompat.ZStore{Keys: []string{"k1", "k2"}, Weights: []int64{1, 2}, Aggregate: "SUM"}
	rangeArgs := rueidiscompat.ZRangeArgs{Key: "z", Start: int64(0), Stop: int64(-1)}

	cases := []func() error{
		func() error {
			cm.ExpectZAdd("z", member).SetVal(1)
			_, err := rdb.ZAdd(ctx, "z", member).Result()
			return err
		},
		func() error {
			cm.ExpectZAddNX("z", member).SetVal(1)
			_, err := rdb.ZAddNX(ctx, "z", member).Result()
			return err
		},
		func() error {
			cm.ExpectZAddXX("z", member).SetVal(1)
			_, err := rdb.ZAddXX(ctx, "z", member).Result()
			return err
		},
		func() error {
			cm.ExpectZAddLT("z", member).SetVal(1)
			_, err := rdb.ZAddLT(ctx, "z", member).Result()
			return err
		},
		func() error {
			cm.ExpectZAddGT("z", member).SetVal(1)
			_, err := rdb.ZAddGT(ctx, "z", member).Result()
			return err
		},
		func() error {
			cm.ExpectZAddArgs("z", rueidiscompat.ZAddArgs{Members: []rueidiscompat.Z{member}, Ch: true}).SetVal(1)
			_, err := rdb.ZAddArgs(ctx, "z", rueidiscompat.ZAddArgs{Members: []rueidiscompat.Z{member}, Ch: true}).Result()
			return err
		},
		func() error {
			cm.ExpectZAddArgsIncr("z", rueidiscompat.ZAddArgs{Members: []rueidiscompat.Z{member}}).SetVal(2.0)
			_, err := rdb.ZAddArgsIncr(ctx, "z", rueidiscompat.ZAddArgs{Members: []rueidiscompat.Z{member}}).Result()
			return err
		},
		func() error {
			cm.ExpectZInter(&store).SetVal([]string{"a"})
			_, err := rdb.ZInter(ctx, store).Result()
			return err
		},
		func() error {
			cm.ExpectZInterWithScores(&store).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZInterWithScores(ctx, store).Result()
			return err
		},
		func() error {
			cm.ExpectZInterStore("dst", &store).SetVal(1)
			_, err := rdb.ZInterStore(ctx, "dst", store).Result()
			return err
		},
		func() error {
			cm.ExpectZUnion(store).SetVal([]string{"a"})
			_, err := rdb.ZUnion(ctx, store).Result()
			return err
		},
		func() error {
			cm.ExpectZUnionWithScores(store).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZUnionWithScores(ctx, store).Result()
			return err
		},
		func() error {
			cm.ExpectZUnionStore("dst", &store).SetVal(1)
			_, err := rdb.ZUnionStore(ctx, "dst", store).Result()
			return err
		},
		func() error {
			cm.ExpectZRangeArgs(rangeArgs).SetVal([]string{"a"})
			_, err := rdb.ZRangeArgs(ctx, rangeArgs).Result()
			return err
		},
		func() error {
			cm.ExpectZRangeArgsWithScores(rangeArgs).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZRangeArgsWithScores(ctx, rangeArgs).Result()
			return err
		},
		func() error {
			cm.ExpectZRangeWithScores("z", 0, -1).SetVal([]rueidiscompat.Z{{Member: "m", Score: 1}})
			_, err := rdb.ZRangeWithScores(ctx, "z", 0, -1).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestPubSubScriptingFunctionsCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	raw.EXPECT().Nodes().Return(map[string]rueidis.Client{"mock": raw}).AnyTimes()
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectPubSubNumSub("ch1", "ch2").SetVal(map[string]int64{"ch1": 1, "ch2": 2})
			_, err := rdb.PubSubNumSub(ctx, "ch1", "ch2").Result()
			return err
		},
		func() error {
			cm.ExpectPubSubShardChannels("p*").SetVal([]string{"ch1"})
			_, err := rdb.PubSubShardChannels(ctx, "p*").Result()
			return err
		},
		func() error {
			cm.ExpectPubSubShardNumSub("ch1").SetVal(map[string]int64{"ch1": 3})
			_, err := rdb.PubSubShardNumSub(ctx, "ch1").Result()
			return err
		},
		func() error {
			cm.ExpectScriptExists("h1", "h2").SetVal([]bool{true, false})
			_, err := rdb.ScriptExists(ctx, "h1", "h2").Result()
			return err
		},
		func() error {
			cm.ExpectScriptFlush().SetVal("OK")
			return rdb.ScriptFlush(ctx).Err()
		},
		func() error {
			cm.ExpectScriptKill().SetVal("OK")
			return rdb.ScriptKill(ctx).Err()
		},
		func() error {
			cm.ExpectScriptLoad("return 1").SetVal("sha1")
			_, err := rdb.ScriptLoad(ctx, "return 1").Result()
			return err
		},
		func() error {
			cm.ExpectFCall("fn", []string{"k1"}, "a", "b").SetVal("ok")
			_, err := rdb.FCall(ctx, "fn", []string{"k1"}, "a", "b").Result()
			return err
		},
		func() error {
			cm.ExpectFunctionDelete("lib").SetVal("OK")
			_, err := rdb.FunctionDelete(ctx, "lib").Result()
			return err
		},
		func() error {
			cm.ExpectFunctionDump().SetVal("dump")
			_, err := rdb.FunctionDump(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectFunctionFlush().SetVal("OK")
			_, err := rdb.FunctionFlush(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectFunctionFlushAsync().SetVal("OK")
			_, err := rdb.FunctionFlushAsync(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectFunctionKill().SetVal("OK")
			_, err := rdb.FunctionKill(ctx).Result()
			return err
		},
		func() error {
			cm.ExpectFunctionList(rueidiscompat.FunctionListQuery{
				LibraryNamePattern: "*",
				WithCode:           true,
			}).SetVal([]rueidiscompat.Library{{
				Name:   "lib",
				Engine: "LUA",
				Code:   "code",
				Functions: []rueidiscompat.Function{{
					Name:        "fn",
					Description: "desc",
					Flags:       []string{"no-writes"},
				}},
			}})
			_, err := rdb.FunctionList(ctx, rueidiscompat.FunctionListQuery{
				LibraryNamePattern: "*",
				WithCode:           true,
			}).Result()
			return err
		},
		func() error {
			cm.ExpectFunctionLoad("code").SetVal("lib")
			_, err := rdb.FunctionLoad(ctx, "code").Result()
			return err
		},
		func() error {
			cm.ExpectFunctionLoadReplace("code").SetVal("lib")
			_, err := rdb.FunctionLoadReplace(ctx, "code").Result()
			return err
		},
		func() error {
			cm.ExpectFunctionRestore("dump").SetVal("OK")
			_, err := rdb.FunctionRestore(ctx, "dump").Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestStreamsCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	xAddArgs := rueidiscompat.XAddArgs{
		Stream: "s",
		Values: map[string]string{"f": "v"},
	}
	xAutoClaimArgs := rueidiscompat.XAutoClaimArgs{
		Stream:   "s",
		Group:    "g",
		Consumer: "c",
		Start:    "0-0",
		MinIdle:  time.Second,
		Count:    10,
	}
	xClaimArgs := rueidiscompat.XClaimArgs{
		Stream:   "s",
		Group:    "g",
		Consumer: "c",
		Messages: []string{"1-0"},
		MinIdle:  time.Second,
	}
	xPendingExtArgs := rueidiscompat.XPendingExtArgs{
		Stream: "s",
		Group:  "g",
		Start:  "-",
		End:    "+",
		Count:  10,
	}
	xReadArgs := rueidiscompat.XReadArgs{
		Streams: []string{"s", "0"},
		Count:   10,
		Block:   -1,
	}
	xReadGroupArgs := rueidiscompat.XReadGroupArgs{
		Group:    "g",
		Consumer: "c",
		Streams:  []string{"s", ">"},
		Count:    10,
		Block:    -1,
	}
	xMsg := rueidiscompat.XMessage{ID: "1-0", Values: map[string]any{"f": "v"}}
	xStream := rueidiscompat.XStream{Stream: "s", Messages: []rueidiscompat.XMessage{xMsg}}

	cases := []func() error{
		func() error {
			cm.ExpectXAck("s", "g", "1-0").SetVal(1)
			_, err := rdb.XAck(ctx, "s", "g", "1-0").Result()
			return err
		},
		func() error {
			cm.ExpectXAdd(&xAddArgs).SetVal("1-0")
			_, err := rdb.XAdd(ctx, xAddArgs).Result()
			return err
		},
		func() error {
			cm.ExpectXAutoClaim(&xAutoClaimArgs).SetVal([]rueidiscompat.XMessage{xMsg}, "0-0")
			_, _, err := rdb.XAutoClaim(ctx, xAutoClaimArgs).Result()
			return err
		},
		func() error {
			cm.ExpectXAutoClaimJustID(&xAutoClaimArgs).SetVal([]string{"1-0"}, "0-0")
			_, _, err := rdb.XAutoClaimJustID(ctx, xAutoClaimArgs).Result()
			return err
		},
		func() error {
			cm.ExpectXClaim(&xClaimArgs).SetVal([]rueidiscompat.XMessage{xMsg})
			_, err := rdb.XClaim(ctx, xClaimArgs).Result()
			return err
		},
		func() error {
			cm.ExpectXClaimJustID(&xClaimArgs).SetVal([]string{"1-0"})
			_, err := rdb.XClaimJustID(ctx, xClaimArgs).Result()
			return err
		},
		func() error {
			cm.ExpectXDel("s", "1-0").SetVal(1)
			_, err := rdb.XDel(ctx, "s", "1-0").Result()
			return err
		},
		func() error {
			cm.ExpectXGroupCreate("s", "g", "0").SetVal("OK")
			return rdb.XGroupCreate(ctx, "s", "g", "0").Err()
		},
		func() error {
			cm.ExpectXGroupCreateConsumer("s", "g", "c").SetVal(1)
			_, err := rdb.XGroupCreateConsumer(ctx, "s", "g", "c").Result()
			return err
		},
		func() error {
			cm.ExpectXGroupCreateMkStream("s", "g", "0").SetVal("OK")
			return rdb.XGroupCreateMkStream(ctx, "s", "g", "0").Err()
		},
		func() error {
			cm.ExpectXGroupDelConsumer("s", "g", "c").SetVal(1)
			_, err := rdb.XGroupDelConsumer(ctx, "s", "g", "c").Result()
			return err
		},
		func() error {
			cm.ExpectXGroupDestroy("s", "g").SetVal(1)
			_, err := rdb.XGroupDestroy(ctx, "s", "g").Result()
			return err
		},
		func() error {
			cm.ExpectXGroupSetID("s", "g", "0").SetVal("OK")
			return rdb.XGroupSetID(ctx, "s", "g", "0").Err()
		},
		func() error {
			cm.ExpectXInfoConsumers("s", "g").SetVal([]rueidiscompat.XInfoConsumer{{Name: "c", Pending: 1}})
			_, err := rdb.XInfoConsumers(ctx, "s", "g").Result()
			return err
		},
		func() error {
			cm.ExpectXInfoGroups("s").SetVal([]rueidiscompat.XInfoGroup{{Name: "g"}})
			_, err := rdb.XInfoGroups(ctx, "s").Result()
			return err
		},
		func() error {
			cm.ExpectXInfoStream("s").SetVal(&rueidiscompat.XInfoStream{Length: 1})
			_, err := rdb.XInfoStream(ctx, "s").Result()
			return err
		},
		func() error {
			cm.ExpectXInfoStreamFull("s", 10).SetVal(&rueidiscompat.XInfoStreamFull{Length: 1})
			_, err := rdb.XInfoStreamFull(ctx, "s", 10).Result()
			return err
		},
		func() error {
			deliveryTime := time.UnixMilli(1000)
			seenTime := time.UnixMilli(2000)
			cm.ExpectXInfoStreamFull("s", 10).SetVal(&rueidiscompat.XInfoStreamFull{
				Length: 1,
				Groups: []rueidiscompat.XInfoStreamGroup{{
					Name:            "g",
					LastDeliveredID: "1-0",
					EntriesRead:     1,
					Lag:             2,
					PelCount:        1,
					Pending: []rueidiscompat.XInfoStreamGroupPending{{
						ID:            "1-0",
						Consumer:      "c",
						DeliveryTime:  deliveryTime,
						DeliveryCount: 1,
					}},
					Consumers: []rueidiscompat.XInfoStreamConsumer{{
						Name:     "c",
						SeenTime: seenTime,
						PelCount: 1,
						Pending: []rueidiscompat.XInfoStreamConsumerPending{{
							ID:            "1-0",
							DeliveryTime:  deliveryTime,
							DeliveryCount: 1,
						}},
					}},
				}},
			})
			_, err := rdb.XInfoStreamFull(ctx, "s", 10).Result()
			return err
		},
		func() error {
			cm.ExpectXLen("s").SetVal(1)
			_, err := rdb.XLen(ctx, "s").Result()
			return err
		},
		func() error {
			cm.ExpectXPending("s", "g").SetVal(&rueidiscompat.XPending{Count: 1, Lower: "0-0", Higher: "1-0"})
			_, err := rdb.XPending(ctx, "s", "g").Result()
			return err
		},
		func() error {
			cm.ExpectXPendingExt(&xPendingExtArgs).SetVal([]rueidiscompat.XPendingExt{{ID: "1-0", Consumer: "c"}})
			_, err := rdb.XPendingExt(ctx, xPendingExtArgs).Result()
			return err
		},
		func() error {
			cm.ExpectXRange("s", "-", "+").SetVal([]rueidiscompat.XMessage{xMsg})
			_, err := rdb.XRange(ctx, "s", "-", "+").Result()
			return err
		},
		func() error {
			cm.ExpectXRangeN("s", "-", "+", 10).SetVal([]rueidiscompat.XMessage{xMsg})
			_, err := rdb.XRangeN(ctx, "s", "-", "+", 10).Result()
			return err
		},
		func() error {
			cm.ExpectXRead(&xReadArgs).SetVal([]rueidiscompat.XStream{xStream})
			_, err := rdb.XRead(ctx, xReadArgs).Result()
			return err
		},
		func() error {
			cm.ExpectXReadGroup(&xReadGroupArgs).SetVal([]rueidiscompat.XStream{xStream})
			_, err := rdb.XReadGroup(ctx, xReadGroupArgs).Result()
			return err
		},
		func() error {
			cm.ExpectXReadStreams("s", "0").SetVal([]rueidiscompat.XStream{xStream})
			_, err := rdb.XReadStreams(ctx, "s", "0").Result()
			return err
		},
		func() error {
			cm.ExpectXRevRange("s", "+", "-").SetVal([]rueidiscompat.XMessage{xMsg})
			_, err := rdb.XRevRange(ctx, "s", "+", "-").Result()
			return err
		},
		func() error {
			cm.ExpectXRevRangeN("s", "+", "-", 10).SetVal([]rueidiscompat.XMessage{xMsg})
			_, err := rdb.XRevRangeN(ctx, "s", "+", "-", 10).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestStreamsHelpersCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectXTrimMaxLen("s", 100).SetVal(1)
			_, err := rdb.XTrimMaxLen(ctx, "s", 100).Result()
			return err
		},
		func() error {
			cm.ExpectXTrimMaxLenApprox("s", 100, 10).SetVal(1)
			_, err := rdb.XTrimMaxLenApprox(ctx, "s", 100, 10).Result()
			return err
		},
		func() error {
			cm.ExpectXTrimMinID("s", "0-0").SetVal(1)
			_, err := rdb.XTrimMinID(ctx, "s", "0-0").Result()
			return err
		},
		func() error {
			cm.ExpectXTrimMinIDApprox("s", "0-0", 10).SetVal(1)
			_, err := rdb.XTrimMinIDApprox(ctx, "s", "0-0", 10).Result()
			return err
		},
	}
	runCases(t, cm, cases)
}

func TestTxPipelineCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cm.ExpectTxPipeline()
	cm.ExpectSet("k", "v", 0).SetVal("OK")
	cm.ExpectGet("k").SetVal("v")
	cm.ExpectTxPipelineExec().SetVal([]any{"OK", "v"})

	tx := rdb.TxPipeline()
	s := tx.Set(ctx, "k", "v", 0)
	g := tx.Get(ctx, "k")
	if _, err := tx.Exec(ctx); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	if err := s.Err(); err != nil {
		t.Fatalf("unexpected set err %v", err)
	}
	if v, err := g.Result(); err != nil || v != "v" {
		t.Fatalf("unexpected get val %q err %v", v, err)
	}
	if err := cm.ExpectationsWereMet(); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestNewExpectedSetErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	raw.EXPECT().Nodes().Return(map[string]rueidis.Client{"mock": raw}).AnyTimes()
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	errAny := errors.New("any")
	xAutoClaimArgs := rueidiscompat.XAutoClaimArgs{
		Stream:   "s",
		Group:    "g",
		Consumer: "c",
		Start:    "0-0",
		MinIdle:  time.Second,
		Count:    10,
	}
	xPendingExtArgs := rueidiscompat.XPendingExtArgs{
		Stream: "s",
		Group:  "g",
		Start:  "-",
		End:    "+",
		Count:  10,
	}
	lcsQuery := &rueidiscompat.LCSQuery{Key1: "k1", Key2: "k2"}

	cases := []func() error{
		func() error {
			cm.ExpectBitField("k", "GET", "u8", 0).SetErr(errAny)
			return wantAnyErr(rdb.BitField(ctx, "k", "GET", "u8", 0).Err())
		},
		func() error {
			cm.ExpectBLMPop(time.Second, "LEFT", 2, "k1", "k2").SetErr(errAny)
			return wantAnyErr(rdb.BLMPop(ctx, time.Second, "LEFT", 2, "k1", "k2").Err())
		},
		func() error {
			cm.ExpectBZPopMax(time.Second, "z").SetErr(errAny)
			return wantAnyErr(rdb.BZPopMax(ctx, time.Second, "z").Err())
		},
		func() error {
			cm.ExpectClientPause(time.Second).SetErr(errAny)
			return wantAnyErr(rdb.ClientPause(ctx, time.Second).Err())
		},
		func() error {
			cm.ExpectClusterLinks().SetErr(errAny)
			return wantAnyErr(rdb.ClusterLinks(ctx).Err())
		},
		func() error {
			cm.ExpectClusterShards().SetErr(errAny)
			return wantAnyErr(rdb.ClusterShards(ctx).Err())
		},
		func() error {
			cm.ExpectClusterSlots().SetErr(errAny)
			return wantAnyErr(rdb.ClusterSlots(ctx).Err())
		},
		func() error {
			cm.ExpectCommand().SetErr(errAny)
			return wantAnyErr(rdb.Command(ctx).Err())
		},
		func() error {
			cm.ExpectCommandGetKeysAndFlags("SET", "k", "v").SetErr(errAny)
			return wantAnyErr(rdb.CommandGetKeysAndFlags(ctx, "SET", "k", "v").Err())
		},
		func() error {
			cm.ExpectFunctionList(rueidiscompat.FunctionListQuery{LibraryNamePattern: "*"}).SetErr(errAny)
			return wantAnyErr(rdb.FunctionList(ctx, rueidiscompat.FunctionListQuery{LibraryNamePattern: "*"}).Err())
		},
		func() error {
			cm.ExpectGeoPos("1", "1", "2").SetErr(errAny)
			return wantAnyErr(rdb.GeoPos(ctx, "1", "1", "2").Err())
		},
		func() error {
			cm.ExpectGeoRadius("1", 1, 2, &rueidiscompat.GeoRadiusQuery{Radius: 200}).SetErr(errAny)
			return wantAnyErr(rdb.GeoRadius(ctx, "1", 1, 2, rueidiscompat.GeoRadiusQuery{Radius: 200}).Err())
		},
		func() error {
			cm.ExpectHRandFieldWithValues("h", 2).SetErr(errAny)
			return wantAnyErr(rdb.HRandFieldWithValues(ctx, "h", 2).Err())
		},
		func() error {
			cm.ExpectIncrByFloat("k", 1.5).SetErr(errAny)
			return wantAnyErr(rdb.IncrByFloat(ctx, "k", 1.5).Err())
		},
		func() error {
			cm.ExpectLCS(lcsQuery).SetErr(errAny)
			return wantAnyErr(rdb.LCS(ctx, lcsQuery).Err())
		},
		func() error {
			cm.ExpectPubSubNumSub("ch1", "ch2").SetErr(errAny)
			return wantAnyErr(rdb.PubSubNumSub(ctx, "ch1", "ch2").Err())
		},
		func() error {
			cm.ExpectScan(0, "k*", 10).SetErr(errAny)
			return wantAnyErr(rdb.Scan(ctx, 0, "k*", 10).Err())
		},
		func() error {
			cm.ExpectScriptExists("h1", "h2").SetErr(errAny)
			return wantAnyErr(rdb.ScriptExists(ctx, "h1", "h2").Err())
		},
		func() error {
			cm.ExpectSMembersMap("s").SetErr(errAny)
			return wantAnyErr(rdb.SMembersMap(ctx, "s").Err())
		},
		func() error {
			cm.ExpectSMIsMember("s", "m1", "m2").SetErr(errAny)
			return wantAnyErr(rdb.SMIsMember(ctx, "s", "m1", "m2").Err())
		},
		func() error {
			cm.ExpectSlowLogGet(10).SetErr(errAny)
			return wantAnyErr(rdb.SlowLogGet(ctx, 10).Err())
		},
		func() error {
			cm.ExpectTime().SetErr(errAny)
			return wantAnyErr(rdb.Time(ctx).Err())
		},
		func() error {
			cm.ExpectXAutoClaim(&xAutoClaimArgs).SetErr(errAny)
			return wantAnyErr(rdb.XAutoClaim(ctx, xAutoClaimArgs).Err())
		},
		func() error {
			cm.ExpectXAutoClaimJustID(&xAutoClaimArgs).SetErr(errAny)
			return wantAnyErr(rdb.XAutoClaimJustID(ctx, xAutoClaimArgs).Err())
		},
		func() error {
			cm.ExpectXInfoConsumers("s", "g").SetErr(errAny)
			return wantAnyErr(rdb.XInfoConsumers(ctx, "s", "g").Err())
		},
		func() error {
			cm.ExpectXInfoGroups("s").SetErr(errAny)
			return wantAnyErr(rdb.XInfoGroups(ctx, "s").Err())
		},
		func() error {
			cm.ExpectXInfoStream("s").SetErr(errAny)
			return wantAnyErr(rdb.XInfoStream(ctx, "s").Err())
		},
		func() error {
			cm.ExpectXInfoStreamFull("s", 10).SetErr(errAny)
			return wantAnyErr(rdb.XInfoStreamFull(ctx, "s", 10).Err())
		},
		func() error {
			cm.ExpectXPending("s", "g").SetErr(errAny)
			return wantAnyErr(rdb.XPending(ctx, "s", "g").Err())
		},
		func() error {
			cm.ExpectXPendingExt(&xPendingExtArgs).SetErr(errAny)
			return wantAnyErr(rdb.XPendingExt(ctx, xPendingExtArgs).Err())
		},
		func() error {
			cm.ExpectXRange("s", "-", "+").SetErr(errAny)
			return wantAnyErr(rdb.XRange(ctx, "s", "-", "+").Err())
		},
		func() error {
			cm.ExpectXReadStreams("s", "0").SetErr(errAny)
			return wantAnyErr(rdb.XReadStreams(ctx, "s", "0").Err())
		},
		func() error {
			cm.ExpectZMPop("MIN", 1, "z").SetErr(errAny)
			return wantAnyErr(rdb.ZMPop(ctx, "MIN", 1, "z").Err())
		},
		func() error {
			cm.ExpectZPopMax("z").SetErr(errAny)
			return wantAnyErr(rdb.ZPopMax(ctx, "z").Err())
		},
		func() error {
			cm.ExpectZMScore("z", "m1", "m2").SetErr(errAny)
			return wantAnyErr(rdb.ZMScore(ctx, "z", "m1", "m2").Err())
		},
	}
	runCases(t, cm, cases)
}

func TestNewExpectedRedisNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	raw := mock.NewClient(ctrl)
	cm := NewAdapter(raw)
	rdb := rueidiscompat.NewAdapter(raw)

	cases := []func() error{
		func() error {
			cm.ExpectBZPopMax(time.Second, "z").RedisNil()
			if err := rdb.BZPopMax(ctx, time.Second, "z").Err(); !errors.Is(err, rueidiscompat.Nil) {
				return fmt.Errorf("expected Nil, got %v", err)
			}
			return nil
		},
	}
	runCases(t, cm, cases)
}

func wantAnyErr(err error) error {
	if err == nil || err.Error() != "any" {
		return fmt.Errorf("expected err \"any\", got %v", err)
	}
	return nil
}

func runCases(t *testing.T, cm ClientMock, cases []func() error) {
	t.Helper()
	for i, c := range cases {
		if err := c(); err != nil {
			t.Fatalf("case %d unexpected err %v", i, err)
		}
	}
	if err := cm.ExpectationsWereMet(); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
}
