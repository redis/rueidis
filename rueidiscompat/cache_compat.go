package rueidiscompat

import (
	"context"
	"time"

	"github.com/redis/rueidis"
)

// CacheCompat is the surface for client-side cached commands, returned by Cache().
// Only exported methods are part of the interface, so it can be implemented by
// external test doubles such as mockery or gomock mocks.
type CacheCompat interface {
	BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd
	BitPosSpan(ctx context.Context, key string, bit, start, end int64, span string) *IntCmd
	BitFieldRO(ctx context.Context, key string, args ...any) *IntSliceCmd
	EvalRO(ctx context.Context, script string, keys []string, args ...any) *Cmd
	EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...any) *Cmd
	FCallRO(ctx context.Context, function string, keys []string, args ...any) *Cmd
	GeoDist(ctx context.Context, key, member1, member2, unit string) *FloatCmd
	GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd
	GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusByMember(ctx context.Context, key, member string, query GeoRadiusQuery) *GeoLocationCmd
	GeoSearch(ctx context.Context, key string, q GeoSearchQuery) *StringSliceCmd
	Get(ctx context.Context, key string) *StringCmd
	MGet(ctx context.Context, keys ...string) *SliceCmd
	GetBit(ctx context.Context, key string, offset int64) *IntCmd
	GetRange(ctx context.Context, key string, start, end int64) *StringCmd
	HExists(ctx context.Context, key, field string) *BoolCmd
	HGet(ctx context.Context, key, field string) *StringCmd
	HGetAll(ctx context.Context, key string) *StringStringMapCmd
	HKeys(ctx context.Context, key string) *StringSliceCmd
	HLen(ctx context.Context, key string) *IntCmd
	HStrLen(ctx context.Context, key, field string) *IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *SliceCmd
	HVals(ctx context.Context, key string) *StringSliceCmd
	LIndex(ctx context.Context, key string, index int64) *StringCmd
	LLen(ctx context.Context, key string) *IntCmd
	LPos(ctx context.Context, key string, element string, a LPosArgs) *IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	PTTL(ctx context.Context, key string) *DurationCmd
	SCard(ctx context.Context, key string) *IntCmd
	SIsMember(ctx context.Context, key string, member any) *BoolCmd
	SMIsMember(ctx context.Context, key string, members ...any) *BoolSliceCmd
	SMembers(ctx context.Context, key string) *StringSliceCmd
	SortRO(ctx context.Context, key string, sort Sort) *StringSliceCmd
	StrLen(ctx context.Context, key string) *IntCmd
	TTL(ctx context.Context, key string) *DurationCmd
	Type(ctx context.Context, key string) *StatusCmd
	ZCard(ctx context.Context, key string) *IntCmd
	ZCount(ctx context.Context, key, min, max string) *IntCmd
	ZLexCount(ctx context.Context, key, min, max string) *IntCmd
	ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd
	ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd
	ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd
	ZRank(ctx context.Context, key, member string) *IntCmd
	ZRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd
	ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRevRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd
	ZRevRank(ctx context.Context, key, member string) *IntCmd
	ZRevRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd
	ZScore(ctx context.Context, key, member string) *FloatCmd
	BFExists(ctx context.Context, key string, element interface{}) *BoolCmd
	BFInfo(ctx context.Context, key string) *BFInfoCmd
	BFInfoArg(ctx context.Context, key, option string) *BFInfoCmd
	BFInfoCapacity(ctx context.Context, key string) *BFInfoCmd
	BFInfoSize(ctx context.Context, key string) *BFInfoCmd
	BFInfoFilters(ctx context.Context, key string) *BFInfoCmd
	BFInfoItems(ctx context.Context, key string) *BFInfoCmd
	BFInfoExpansion(ctx context.Context, key string) *BFInfoCmd
	CFCount(ctx context.Context, key string, element interface{}) *IntCmd
	CFExists(ctx context.Context, key string, element interface{}) *BoolCmd
	CFInfo(ctx context.Context, key string) *CFInfoCmd
	CMSInfo(ctx context.Context, key string) *CMSInfoCmd
	CMSQuery(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd
	TopKInfo(ctx context.Context, key string) *TopKInfoCmd
	TopKList(ctx context.Context, key string) *StringSliceCmd
	TopKQuery(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd
	JSONArrIndex(ctx context.Context, key, path string, value ...interface{}) *IntSliceCmd
	JSONArrLen(ctx context.Context, key, path string) *IntSliceCmd
	JSONGet(ctx context.Context, key string, paths ...string) *JSONCmd
	JSONMGet(ctx context.Context, path string, keys ...string) *JSONSliceCmd
	JSONObjKeys(ctx context.Context, key, path string) *SliceCmd
	JSONObjLen(ctx context.Context, key, path string) *IntPointerSliceCmd
	JSONStrLen(ctx context.Context, key, path string) *IntPointerSliceCmd
	JSONType(ctx context.Context, key, path string) *JSONSliceCmd
}

type cacheCompat struct {
	client rueidis.Client
	ttl    time.Duration
}

var _ CacheCompat = cacheCompat{}
