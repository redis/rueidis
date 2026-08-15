package rueidiscompat_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/redis/rueidis/rueidiscompat"
)

// fakeCache is an external-package test double that implements the exported
// CacheCompat surface literally, exactly as mockery/gomock test doubles do.
// Because it lives outside the package and cannot implement unexported
// methods, it fails to compile if any unexported method is ever added to the
// interface — protecting the exported-only constraint that makes client-side
// caching mockable (regression for issue #972).
type fakeCache struct{}

func (c fakeCache) BFExists(ctx context.Context, key string, element interface{}) *rueidiscompat.BoolCmd {
	return nil
}
func (c fakeCache) BFInfo(ctx context.Context, key string) *rueidiscompat.BFInfoCmd { return nil }
func (c fakeCache) BFInfoArg(ctx context.Context, key string, option string) *rueidiscompat.BFInfoCmd {
	return nil
}
func (c fakeCache) BFInfoCapacity(ctx context.Context, key string) *rueidiscompat.BFInfoCmd {
	return nil
}
func (c fakeCache) BFInfoExpansion(ctx context.Context, key string) *rueidiscompat.BFInfoCmd {
	return nil
}
func (c fakeCache) BFInfoFilters(ctx context.Context, key string) *rueidiscompat.BFInfoCmd {
	return nil
}
func (c fakeCache) BFInfoItems(ctx context.Context, key string) *rueidiscompat.BFInfoCmd { return nil }
func (c fakeCache) BFInfoSize(ctx context.Context, key string) *rueidiscompat.BFInfoCmd  { return nil }
func (c fakeCache) BitCount(ctx context.Context, key string, bitCount *rueidiscompat.BitCount) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) BitFieldRO(ctx context.Context, key string, args ...any) *rueidiscompat.IntSliceCmd {
	return nil
}
func (c fakeCache) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) BitPosSpan(ctx context.Context, key string, bit int64, start int64, end int64, span string) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) CFCount(ctx context.Context, key string, element interface{}) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) CFExists(ctx context.Context, key string, element interface{}) *rueidiscompat.BoolCmd {
	return nil
}
func (c fakeCache) CFInfo(ctx context.Context, key string) *rueidiscompat.CFInfoCmd   { return nil }
func (c fakeCache) CMSInfo(ctx context.Context, key string) *rueidiscompat.CMSInfoCmd { return nil }
func (c fakeCache) CMSQuery(ctx context.Context, key string, elements ...interface{}) *rueidiscompat.IntSliceCmd {
	return nil
}
func (c fakeCache) EvalRO(ctx context.Context, script string, keys []string, args ...any) *rueidiscompat.Cmd {
	return nil
}
func (c fakeCache) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...any) *rueidiscompat.Cmd {
	return nil
}
func (c fakeCache) FCallRO(ctx context.Context, function string, keys []string, args ...any) *rueidiscompat.Cmd {
	return nil
}
func (c fakeCache) GeoDist(ctx context.Context, key string, member1 string, member2 string, unit string) *rueidiscompat.FloatCmd {
	return nil
}
func (c fakeCache) GeoHash(ctx context.Context, key string, members ...string) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) GeoPos(ctx context.Context, key string, members ...string) *rueidiscompat.GeoPosCmd {
	return nil
}
func (c fakeCache) GeoRadius(ctx context.Context, key string, longitude float64, latitude float64, query rueidiscompat.GeoRadiusQuery) *rueidiscompat.GeoLocationCmd {
	return nil
}
func (c fakeCache) GeoRadiusByMember(ctx context.Context, key string, member string, query rueidiscompat.GeoRadiusQuery) *rueidiscompat.GeoLocationCmd {
	return nil
}
func (c fakeCache) GeoSearch(ctx context.Context, key string, q rueidiscompat.GeoSearchQuery) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) Get(ctx context.Context, key string) *rueidiscompat.StringCmd {
	return &rueidiscompat.StringCmd{}
}
func (c fakeCache) GetBit(ctx context.Context, key string, offset int64) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) GetRange(ctx context.Context, key string, start int64, end int64) *rueidiscompat.StringCmd {
	return nil
}
func (c fakeCache) HExists(ctx context.Context, key string, field string) *rueidiscompat.BoolCmd {
	return nil
}
func (c fakeCache) HGet(ctx context.Context, key string, field string) *rueidiscompat.StringCmd {
	return nil
}
func (c fakeCache) HGetAll(ctx context.Context, key string) *rueidiscompat.StringStringMapCmd {
	return nil
}
func (c fakeCache) HKeys(ctx context.Context, key string) *rueidiscompat.StringSliceCmd { return nil }
func (c fakeCache) HLen(ctx context.Context, key string) *rueidiscompat.IntCmd          { return nil }
func (c fakeCache) HMGet(ctx context.Context, key string, fields ...string) *rueidiscompat.SliceCmd {
	return nil
}
func (c fakeCache) HStrLen(ctx context.Context, key string, field string) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) HVals(ctx context.Context, key string) *rueidiscompat.StringSliceCmd { return nil }
func (c fakeCache) JSONArrIndex(ctx context.Context, key string, path string, value ...interface{}) *rueidiscompat.IntSliceCmd {
	return nil
}
func (c fakeCache) JSONArrLen(ctx context.Context, key string, path string) *rueidiscompat.IntSliceCmd {
	return nil
}
func (c fakeCache) JSONGet(ctx context.Context, key string, paths ...string) *rueidiscompat.JSONCmd {
	return nil
}
func (c fakeCache) JSONMGet(ctx context.Context, path string, keys ...string) *rueidiscompat.JSONSliceCmd {
	return nil
}
func (c fakeCache) JSONObjKeys(ctx context.Context, key string, path string) *rueidiscompat.SliceCmd {
	return nil
}
func (c fakeCache) JSONObjLen(ctx context.Context, key string, path string) *rueidiscompat.IntPointerSliceCmd {
	return nil
}
func (c fakeCache) JSONStrLen(ctx context.Context, key string, path string) *rueidiscompat.IntPointerSliceCmd {
	return nil
}
func (c fakeCache) JSONType(ctx context.Context, key string, path string) *rueidiscompat.JSONSliceCmd {
	return nil
}
func (c fakeCache) LIndex(ctx context.Context, key string, index int64) *rueidiscompat.StringCmd {
	return nil
}
func (c fakeCache) LLen(ctx context.Context, key string) *rueidiscompat.IntCmd { return nil }
func (c fakeCache) LPos(ctx context.Context, key string, element string, a rueidiscompat.LPosArgs) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) LRange(ctx context.Context, key string, start int64, stop int64) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) MGet(ctx context.Context, keys ...string) *rueidiscompat.SliceCmd { return nil }
func (c fakeCache) PTTL(ctx context.Context, key string) *rueidiscompat.DurationCmd  { return nil }
func (c fakeCache) SCard(ctx context.Context, key string) *rueidiscompat.IntCmd      { return nil }
func (c fakeCache) SIsMember(ctx context.Context, key string, member any) *rueidiscompat.BoolCmd {
	return nil
}
func (c fakeCache) SMIsMember(ctx context.Context, key string, members ...any) *rueidiscompat.BoolSliceCmd {
	return nil
}
func (c fakeCache) SMembers(ctx context.Context, key string) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) SortRO(ctx context.Context, key string, sort rueidiscompat.Sort) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) StrLen(ctx context.Context, key string) *rueidiscompat.IntCmd        { return nil }
func (c fakeCache) TTL(ctx context.Context, key string) *rueidiscompat.DurationCmd      { return nil }
func (c fakeCache) TopKInfo(ctx context.Context, key string) *rueidiscompat.TopKInfoCmd { return nil }
func (c fakeCache) TopKList(ctx context.Context, key string) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) TopKQuery(ctx context.Context, key string, elements ...interface{}) *rueidiscompat.BoolSliceCmd {
	return nil
}
func (c fakeCache) Type(ctx context.Context, key string) *rueidiscompat.StatusCmd { return nil }
func (c fakeCache) ZCard(ctx context.Context, key string) *rueidiscompat.IntCmd   { return nil }
func (c fakeCache) ZCount(ctx context.Context, key string, min string, max string) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) ZLexCount(ctx context.Context, key string, min string, max string) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) ZMScore(ctx context.Context, key string, members ...string) *rueidiscompat.FloatSliceCmd {
	return nil
}
func (c fakeCache) ZRangeArgs(ctx context.Context, z rueidiscompat.ZRangeArgs) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) ZRangeArgsWithScores(ctx context.Context, z rueidiscompat.ZRangeArgs) *rueidiscompat.ZSliceCmd {
	return nil
}
func (c fakeCache) ZRangeByLex(ctx context.Context, key string, opt rueidiscompat.ZRangeBy) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) ZRangeByScore(ctx context.Context, key string, opt rueidiscompat.ZRangeBy) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) ZRangeByScoreWithScores(ctx context.Context, key string, opt rueidiscompat.ZRangeBy) *rueidiscompat.ZSliceCmd {
	return nil
}
func (c fakeCache) ZRangeWithScores(ctx context.Context, key string, start int64, stop int64) *rueidiscompat.ZSliceCmd {
	return nil
}
func (c fakeCache) ZRank(ctx context.Context, key string, member string) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) ZRankWithScore(ctx context.Context, key string, member string) *rueidiscompat.RankWithScoreCmd {
	return nil
}
func (c fakeCache) ZRevRange(ctx context.Context, key string, start int64, stop int64) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) ZRevRangeByLex(ctx context.Context, key string, opt rueidiscompat.ZRangeBy) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) ZRevRangeByScore(ctx context.Context, key string, opt rueidiscompat.ZRangeBy) *rueidiscompat.StringSliceCmd {
	return nil
}
func (c fakeCache) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt rueidiscompat.ZRangeBy) *rueidiscompat.ZSliceCmd {
	return nil
}
func (c fakeCache) ZRevRangeWithScores(ctx context.Context, key string, start int64, stop int64) *rueidiscompat.ZSliceCmd {
	return nil
}
func (c fakeCache) ZRevRank(ctx context.Context, key string, member string) *rueidiscompat.IntCmd {
	return nil
}
func (c fakeCache) ZRevRankWithScore(ctx context.Context, key string, member string) *rueidiscompat.RankWithScoreCmd {
	return nil
}
func (c fakeCache) ZScore(ctx context.Context, key string, member string) *rueidiscompat.FloatCmd {
	return nil
}

var _ rueidiscompat.CacheCompat = fakeCache{}

func TestCacheCompatMockability(t *testing.T) {
	// A user test double satisfies the interface (the #972 failure mode).
	var cc rueidiscompat.CacheCompat = fakeCache{}
	// Non-vacuous dispatch check: the fake returns a distinguishable payload, so
	// this fails if the method does not actually round-trip through the interface.
	if got := cc.Get(context.Background(), "key"); got == nil {
		t.Fatal("CacheCompat test double did not dispatch Get through the interface")
	}
	// The concrete handle from Cache() also flows through the interface (chained-call pattern).
	_ = (&rueidiscompat.Compat{}).Cache(time.Second)
	// Pin the interface surface (the #972 guarantee) so that a method REMOVED from
	// the interface (surface shrink) or dropped from the fake fails the test, not
	// only additions that the external-package compile guard already catches.
	const want = 82
	if n := reflect.TypeOf((*rueidiscompat.CacheCompat)(nil)).Elem().NumMethod(); n != want {
		t.Fatalf("CacheCompat interface has %d methods, want %d", n, want)
	}
	if n := reflect.TypeOf(fakeCache{}).NumMethod(); n != want {
		t.Fatalf("fakeCache implements %d methods, want %d", n, want)
	}
}
