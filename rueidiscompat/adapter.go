package rueidiscompat

import (
	"context"
	"encoding"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/cmds"
)

const KeepTTL = -1

type Cmdable interface {
	Cache(ttl time.Duration) CacheCompat

	Command(ctx context.Context) *CommandsInfoCmd
	ClientGetName(ctx context.Context) *StringCmd
	Echo(ctx context.Context, message interface{}) *StringCmd
	Ping(ctx context.Context) *StatusCmd
	Quit(ctx context.Context) *StatusCmd
	Del(ctx context.Context, keys ...string) *IntCmd
	Unlink(ctx context.Context, keys ...string) *IntCmd
	Dump(ctx context.Context, key string) *StringCmd
	Exists(ctx context.Context, keys ...string) *IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
	ExpireNX(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireXX(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireGT(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireLT(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	Keys(ctx context.Context, pattern string) *StringSliceCmd
	Migrate(ctx context.Context, host string, port int64, key string, db int64, timeout time.Duration) *StatusCmd
	Move(ctx context.Context, key string, db int64) *BoolCmd
	ObjectRefCount(ctx context.Context, key string) *IntCmd
	ObjectEncoding(ctx context.Context, key string) *StringCmd
	ObjectIdleTime(ctx context.Context, key string) *DurationCmd
	Persist(ctx context.Context, key string) *BoolCmd
	PExpire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	PExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
	PTTL(ctx context.Context, key string) *DurationCmd
	RandomKey(ctx context.Context) *StringCmd
	Rename(ctx context.Context, key, newkey string) *StatusCmd
	RenameNX(ctx context.Context, key, newkey string) *BoolCmd
	Restore(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd
	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd
	Sort(ctx context.Context, key string, sort Sort) *StringSliceCmd
	SortStore(ctx context.Context, key, store string, sort Sort) *IntCmd
	SortInterfaces(ctx context.Context, key string, sort Sort) *SliceCmd
	Touch(ctx context.Context, keys ...string) *IntCmd
	TTL(ctx context.Context, key string) *DurationCmd
	Type(ctx context.Context, key string) *StatusCmd
	Append(ctx context.Context, key, value string) *IntCmd
	Decr(ctx context.Context, key string) *IntCmd
	DecrBy(ctx context.Context, key string, decrement int64) *IntCmd
	Get(ctx context.Context, key string) *StringCmd
	GetRange(ctx context.Context, key string, start, end int64) *StringCmd
	GetSet(ctx context.Context, key string, value interface{}) *StringCmd
	GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd
	GetDel(ctx context.Context, key string) *StringCmd
	Incr(ctx context.Context, key string) *IntCmd
	IncrBy(ctx context.Context, key string, value int64) *IntCmd
	IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd
	MGet(ctx context.Context, keys ...string) *SliceCmd
	MSet(ctx context.Context, values ...interface{}) *StatusCmd
	MSetNX(ctx context.Context, values ...interface{}) *BoolCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *StatusCmd
	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
	SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd
	StrLen(ctx context.Context, key string) *IntCmd
	Copy(ctx context.Context, sourceKey string, destKey string, db int64, replace bool) *IntCmd

	GetBit(ctx context.Context, key string, offset int64) *IntCmd
	SetBit(ctx context.Context, key string, offset int64, value int64) *IntCmd
	BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd
	BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpNot(ctx context.Context, destKey string, key string) *IntCmd
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd
	BitField(ctx context.Context, key string, args ...interface{}) *IntSliceCmd

	Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd
	ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd

	HDel(ctx context.Context, key string, fields ...string) *IntCmd
	HExists(ctx context.Context, key, field string) *BoolCmd
	HGet(ctx context.Context, key, field string) *StringCmd
	HGetAll(ctx context.Context, key string) *StringStringMapCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd
	HKeys(ctx context.Context, key string) *StringSliceCmd
	HLen(ctx context.Context, key string) *IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *SliceCmd
	HSet(ctx context.Context, key string, values ...interface{}) *IntCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *BoolCmd
	HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd
	HVals(ctx context.Context, key string) *StringSliceCmd
	HRandField(ctx context.Context, key string, count int64, withValues bool) *StringSliceCmd

	BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd
	LIndex(ctx context.Context, key string, index int64) *StringCmd
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *IntCmd
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *IntCmd
	LLen(ctx context.Context, key string) *IntCmd
	LPop(ctx context.Context, key string) *StringCmd
	LPopCount(ctx context.Context, key string, count int64) *StringSliceCmd
	LPos(ctx context.Context, key string, value string, args LPosArgs) *IntCmd
	LPosCount(ctx context.Context, key string, value string, count int64, args LPosArgs) *IntSliceCmd
	LPush(ctx context.Context, key string, values ...interface{}) *IntCmd
	LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd
	LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd
	LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd
	RPop(ctx context.Context, key string) *StringCmd
	RPopCount(ctx context.Context, key string, count int64) *StringSliceCmd
	RPopLPush(ctx context.Context, source, destination string) *StringCmd
	RPush(ctx context.Context, key string, values ...interface{}) *IntCmd
	RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
	LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd
	BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *StringCmd

	SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd
	SCard(ctx context.Context, key string) *IntCmd
	SDiff(ctx context.Context, keys ...string) *StringSliceCmd
	SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd
	SInter(ctx context.Context, keys ...string) *StringSliceCmd
	SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd
	SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd
	SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd
	SMembers(ctx context.Context, key string) *StringSliceCmd
	SMembersMap(ctx context.Context, key string) *StringStructMapCmd
	SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd
	SPop(ctx context.Context, key string) *StringCmd
	SPopN(ctx context.Context, key string, count int64) *StringSliceCmd
	SRandMember(ctx context.Context, key string) *StringCmd
	SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd
	SRem(ctx context.Context, key string, members ...interface{}) *IntCmd
	SUnion(ctx context.Context, keys ...string) *StringSliceCmd
	SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd

	XAdd(ctx context.Context, a XAddArgs) *StringCmd
	XDel(ctx context.Context, stream string, ids ...string) *IntCmd
	XLen(ctx context.Context, stream string) *IntCmd
	XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd
	XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd
	XRevRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd
	XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *XMessageSliceCmd
	XRead(ctx context.Context, a XReadArgs) *XStreamSliceCmd
	XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd
	XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd
	XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd
	XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd
	XGroupDestroy(ctx context.Context, stream, group string) *IntCmd
	XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
	XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
	XReadGroup(ctx context.Context, a XReadGroupArgs) *XStreamSliceCmd
	XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd
	XPending(ctx context.Context, stream, group string) *XPendingCmd
	XPendingExt(ctx context.Context, a XPendingExtArgs) *XPendingExtCmd
	XClaim(ctx context.Context, a XClaimArgs) *XMessageSliceCmd
	XClaimJustID(ctx context.Context, a XClaimArgs) *StringSliceCmd
	XAutoClaim(ctx context.Context, a XAutoClaimArgs) *XAutoClaimCmd
	XAutoClaimJustID(ctx context.Context, a XAutoClaimArgs) *XAutoClaimJustIDCmd

	XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd
	XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd
	XTrimMinID(ctx context.Context, key string, minID string) *IntCmd
	XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd
	XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd
	XInfoStream(ctx context.Context, key string) *XInfoStreamCmd
	XInfoStreamFull(ctx context.Context, key string, count int64) *XInfoStreamFullCmd
	XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd

	BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
	BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd

	ZAdd(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddNX(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddXX(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd
	ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd
	ZCard(ctx context.Context, key string) *IntCmd
	ZCount(ctx context.Context, key string, min, max string) *IntCmd
	ZLexCount(ctx context.Context, key, min, max string) *IntCmd
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd
	ZInter(ctx context.Context, store ZStore) *StringSliceCmd
	ZInterWithScores(ctx context.Context, store ZStore) *ZSliceCmd
	ZInterStore(ctx context.Context, destination string, store ZStore) *IntCmd
	ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd
	ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd
	ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd
	ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd
	ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd
	ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd
	ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd
	ZRank(ctx context.Context, key, member string) *IntCmd
	ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd
	ZRemRangeByScore(ctx context.Context, key string, min, max string) *IntCmd
	ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd
	ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRevRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd
	ZRevRank(ctx context.Context, key, member string) *IntCmd
	ZScore(ctx context.Context, key, member string) *FloatCmd
	ZUnionStore(ctx context.Context, dest string, store ZStore) *IntCmd
	ZUnion(ctx context.Context, store ZStore) *StringSliceCmd
	ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd
	ZRandMember(ctx context.Context, key string, count int64, withScores bool) *StringSliceCmd
	ZDiff(ctx context.Context, keys ...string) *StringSliceCmd
	ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd
	ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd

	PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd
	PFCount(ctx context.Context, keys ...string) *IntCmd
	PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd

	BgRewriteAOF(ctx context.Context) *StatusCmd
	BgSave(ctx context.Context) *StatusCmd
	ClientKill(ctx context.Context, ipPort string) *StatusCmd
	ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd
	ClientList(ctx context.Context) *StringCmd
	ClientPause(ctx context.Context, dur time.Duration) *BoolCmd
	ClientID(ctx context.Context) *IntCmd
	ConfigGet(ctx context.Context, parameter string) *SliceCmd
	ConfigResetStat(ctx context.Context) *StatusCmd
	ConfigSet(ctx context.Context, parameter, value string) *StatusCmd
	ConfigRewrite(ctx context.Context) *StatusCmd
	DBSize(ctx context.Context) *IntCmd
	FlushAll(ctx context.Context) *StatusCmd
	FlushAllAsync(ctx context.Context) *StatusCmd
	FlushDB(ctx context.Context) *StatusCmd
	FlushDBAsync(ctx context.Context) *StatusCmd
	Info(ctx context.Context, section ...string) *StringCmd
	LastSave(ctx context.Context) *IntCmd
	Save(ctx context.Context) *StatusCmd
	Shutdown(ctx context.Context) *StatusCmd
	ShutdownSave(ctx context.Context) *StatusCmd
	ShutdownNoSave(ctx context.Context) *StatusCmd
	SlaveOf(ctx context.Context, host string, port string) *StatusCmd
	Time(ctx context.Context) *TimeCmd
	DebugObject(ctx context.Context, key string) *StringCmd
	ReadOnly(ctx context.Context) *StatusCmd
	ReadWrite(ctx context.Context) *StatusCmd
	MemoryUsage(ctx context.Context, key string, samples ...int64) *IntCmd

	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd
	ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd
	ScriptFlush(ctx context.Context) *StatusCmd
	ScriptKill(ctx context.Context) *StatusCmd
	ScriptLoad(ctx context.Context, script string) *StringCmd

	Publish(ctx context.Context, channel string, message interface{}) *IntCmd
	PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd
	PubSubNumSub(ctx context.Context, channels ...string) *StringIntMapCmd
	PubSubNumPat(ctx context.Context) *IntCmd

	ClusterSlots(ctx context.Context) *ClusterSlotsCmd
	ClusterNodes(ctx context.Context) *StringCmd
	ClusterMeet(ctx context.Context, host string, port int64) *StatusCmd
	ClusterForget(ctx context.Context, nodeID string) *StatusCmd
	ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd
	ClusterResetSoft(ctx context.Context) *StatusCmd
	ClusterResetHard(ctx context.Context) *StatusCmd
	ClusterInfo(ctx context.Context) *StringCmd
	ClusterKeySlot(ctx context.Context, key string) *IntCmd
	ClusterGetKeysInSlot(ctx context.Context, slot int64, count int64) *StringSliceCmd
	ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd
	ClusterCountKeysInSlot(ctx context.Context, slot int64) *IntCmd
	ClusterDelSlots(ctx context.Context, slots ...int64) *StatusCmd
	ClusterDelSlotsRange(ctx context.Context, min, max int64) *StatusCmd
	ClusterSaveConfig(ctx context.Context) *StatusCmd
	ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd
	ClusterFailover(ctx context.Context) *StatusCmd
	ClusterAddSlots(ctx context.Context, slots ...int64) *StatusCmd
	ClusterAddSlotsRange(ctx context.Context, min, max int64) *StatusCmd

	GeoAdd(ctx context.Context, key string, geoLocation ...GeoLocation) *IntCmd
	GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *IntCmd
	GeoRadiusByMember(ctx context.Context, key, member string, query GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query GeoRadiusQuery) *IntCmd
	GeoSearch(ctx context.Context, key string, q GeoSearchQuery) *StringSliceCmd
	GeoSearchLocation(ctx context.Context, key string, q GeoSearchLocationQuery) *GeoLocationCmd
	GeoSearchStore(ctx context.Context, src, dest string, q GeoSearchStoreQuery) *IntCmd
	GeoDist(ctx context.Context, key, member1, member2, unit string) *FloatCmd
	GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd
}

type Compat struct {
	client rueidis.Client
}

type CacheCompat struct {
	client rueidis.Client
	ttl    time.Duration
}

func NewAdapter(client rueidis.Client) Cmdable {
	return &Compat{client: client}
}

func (c *Compat) Cache(ttl time.Duration) CacheCompat {
	return CacheCompat{client: c.client, ttl: ttl}
}

func (c *Compat) Command(ctx context.Context) *CommandsInfoCmd {
	cmd := c.client.B().Command().Build()
	resp := c.client.Do(ctx, cmd)
	return newCommandsInfoCmd(resp)
}

func (c *Compat) ClientGetName(ctx context.Context) *StringCmd {
	cmd := c.client.B().ClientGetname().Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) Echo(ctx context.Context, message interface{}) *StringCmd {
	cmd := c.client.B().Echo().Message(str(message)).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) Ping(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Ping().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) Quit(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Quit().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) Del(ctx context.Context, keys ...string) *IntCmd {
	cmd := c.client.B().Del().Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) Unlink(ctx context.Context, keys ...string) *IntCmd {
	cmd := c.client.B().Unlink().Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) Dump(ctx context.Context, key string) *StringCmd {
	cmd := c.client.B().Dump().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) Exists(ctx context.Context, keys ...string) *IntCmd {
	cmd := c.client.B().Exists().Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) Expire(ctx context.Context, key string, seconds time.Duration) *BoolCmd {
	cmd := c.client.B().Expire().Key(key).Seconds(formatSec(seconds)).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) ExpireAt(ctx context.Context, key string, timestamp time.Time) *BoolCmd {
	cmd := c.client.B().Expireat().Key(key).Timestamp(timestamp.Unix()).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) ExpireNX(ctx context.Context, key string, seconds time.Duration) *BoolCmd {
	cmd := c.client.B().Expire().Key(key).Seconds(formatSec(seconds)).Nx().Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) ExpireXX(ctx context.Context, key string, seconds time.Duration) *BoolCmd {
	cmd := c.client.B().Expire().Key(key).Seconds(formatSec(seconds)).Xx().Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) ExpireGT(ctx context.Context, key string, seconds time.Duration) *BoolCmd {
	cmd := c.client.B().Expire().Key(key).Seconds(formatSec(seconds)).Gt().Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) ExpireLT(ctx context.Context, key string, seconds time.Duration) *BoolCmd {
	cmd := c.client.B().Expire().Key(key).Seconds(formatSec(seconds)).Lt().Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) Keys(ctx context.Context, pattern string) *StringSliceCmd {
	cmd := c.client.B().Keys().Pattern(pattern).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) Migrate(ctx context.Context, host string, port int64, key string, db int64, timeout time.Duration) *StatusCmd {
	var cmd cmds.Completed
	if key == "key" {
		cmd = c.client.B().Migrate().Host(host).Port(port).Key().DestinationDb(db).Timeout(formatSec(timeout)).Build()
	} else {
		cmd = c.client.B().Migrate().Host(host).Port(port).Empty().DestinationDb(db).Timeout(formatSec(timeout)).Build()
	}
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) Move(ctx context.Context, key string, db int64) *BoolCmd {
	cmd := c.client.B().Move().Key(key).Db(db).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) ObjectRefCount(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().ObjectRefcount().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ObjectEncoding(ctx context.Context, key string) *StringCmd {
	cmd := c.client.B().ObjectEncoding().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) ObjectIdleTime(ctx context.Context, key string) *DurationCmd {
	cmd := c.client.B().ObjectIdletime().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newDurationCmd(resp, time.Second)
}
func (c *Compat) Persist(ctx context.Context, key string) *BoolCmd {
	cmd := c.client.B().Persist().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) PExpire(ctx context.Context, key string, milliseconds time.Duration) *BoolCmd {
	cmd := c.client.B().Pexpire().Key(key).Milliseconds(formatMs(milliseconds)).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) PExpireAt(ctx context.Context, key string, millisecondsTimestamp time.Time) *BoolCmd {
	cmd := c.client.B().Pexpireat().Key(key).MillisecondsTimestamp(millisecondsTimestamp.UnixNano() / int64(time.Millisecond)).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) PTTL(ctx context.Context, key string) *DurationCmd {
	cmd := c.client.B().Pttl().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newDurationCmd(resp, time.Millisecond)
}

func (c *Compat) RandomKey(ctx context.Context) *StringCmd {
	cmd := c.client.B().Randomkey().Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) Rename(ctx context.Context, key, newkey string) *StatusCmd {
	cmd := c.client.B().Rename().Key(key).Newkey(newkey).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) RenameNX(ctx context.Context, key, newkey string) *BoolCmd {
	cmd := c.client.B().Renamenx().Key(key).Newkey(newkey).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) Restore(ctx context.Context, key string, ttl time.Duration, serializedValue string) *StatusCmd {
	cmd := c.client.B().Restore().Key(key).Ttl(formatMs(ttl)).SerializedValue(serializedValue).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) RestoreReplace(ctx context.Context, key string, ttl time.Duration, serializedValue string) *StatusCmd {
	cmd := c.client.B().Restore().Key(key).Ttl(formatMs(ttl)).SerializedValue(serializedValue).Replace().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) sort(key string, sort Sort) cmds.Arbitrary {
	cmd := c.client.B().Arbitrary("SORT").Keys(key)
	if sort.By != "" {
		cmd = cmd.Args("BY", sort.By)
	}
	if sort.Offset != 0 || sort.Count != 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(sort.Offset, 10), strconv.FormatInt(sort.Count, 10))
	}
	if len(sort.Get) > 0 {
		cmd = cmd.Args("GET").Args(sort.Get...)
	}
	switch order := strings.ToUpper(sort.Order); order {
	case "ASC", "DESC":
		cmd = cmd.Args(order)
	case "":
	default:
		panic(fmt.Sprintf("invalid sort order %s", sort.Order))
	}
	if sort.Alpha {
		cmd = cmd.Args("ALPHA")
	}
	return cmd
}

func (c *Compat) Sort(ctx context.Context, key string, sort Sort) *StringSliceCmd {
	resp := c.client.Do(ctx, c.sort(key, sort).Build())
	return newStringSliceCmd(resp)
}

func (c *Compat) SortStore(ctx context.Context, key, store string, sort Sort) *IntCmd {
	resp := c.client.Do(ctx, c.sort(key, sort).Args("STORE", store).Build())
	return newIntCmd(resp)
}

func (c *Compat) SortInterfaces(ctx context.Context, key string, sort Sort) *SliceCmd {
	resp := c.client.Do(ctx, c.sort(key, sort).Build())
	return newSliceCmd(resp)
}

func (c *Compat) Touch(ctx context.Context, keys ...string) *IntCmd {
	cmd := c.client.B().Touch().Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) TTL(ctx context.Context, key string) *DurationCmd {
	cmd := c.client.B().Ttl().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newDurationCmd(resp, time.Second)
}

func (c *Compat) Type(ctx context.Context, key string) *StatusCmd {
	cmd := c.client.B().Type().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) Append(ctx context.Context, key, value string) *IntCmd {
	cmd := c.client.B().Append().Key(key).Value(value).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) Decr(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Decr().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) DecrBy(ctx context.Context, key string, decrement int64) *IntCmd {
	cmd := c.client.B().Decrby().Key(key).Decrement(decrement).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) Get(ctx context.Context, key string) *StringCmd {
	cmd := c.client.B().Get().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) GetRange(ctx context.Context, key string, start, end int64) *StringCmd {
	cmd := c.client.B().Getrange().Key(key).Start(start).End(end).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) GetSet(ctx context.Context, key string, value interface{}) *StringCmd {
	cmd := c.client.B().Getset().Key(key).Value(str(value)).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

// GetEx An expiration of zero removes the TTL associated with the key (i.e. GETEX key persist).
// Requires Redis >= 6.2.0.
func (c *Compat) GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd {
	var resp rueidis.RedisResult
	if expiration > 0 {
		if usePrecise(expiration) {
			resp = c.client.Do(ctx, c.client.B().Getex().Key(key).PxMilliseconds(formatMs(expiration)).Build())
		} else {
			resp = c.client.Do(ctx, c.client.B().Getex().Key(key).ExSeconds(formatSec(expiration)).Build())
		}
	} else {
		resp = c.client.Do(ctx, c.client.B().Getex().Key(key).Build())
	}
	return newStringCmd(resp)
}

func (c *Compat) GetDel(ctx context.Context, key string) *StringCmd {
	cmd := c.client.B().Getdel().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) Incr(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Incr().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) IncrBy(ctx context.Context, key string, increment int64) *IntCmd {
	cmd := c.client.B().Incrby().Key(key).Increment(increment).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) IncrByFloat(ctx context.Context, key string, increment float64) *FloatCmd {
	cmd := c.client.B().Incrbyfloat().Key(key).Increment(increment).Build()
	resp := c.client.Do(ctx, cmd)
	return newFloatCmd(resp)
}

func (c *Compat) MGet(ctx context.Context, keys ...string) *SliceCmd {
	cmd := c.client.B().Mget().Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newSliceCmd(resp)
}

func (c *Compat) MSet(ctx context.Context, values ...interface{}) *StatusCmd {
	partial := c.client.B().Mset().KeyValue()

	args := argsToSlice(values)
	for i := 0; i < len(args); i += 2 {
		partial = partial.KeyValue(args[i], args[i+1])
	}
	cmd := partial.Build()

	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) MSetNX(ctx context.Context, values ...interface{}) *BoolCmd {
	partial := c.client.B().Msetnx().KeyValue()

	args := argsToSlice(values)
	for i := 0; i < len(args); i += 2 {
		partial = partial.KeyValue(args[i], args[i+1])
	}
	cmd := partial.Build()

	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

// Set key value [expiration]
//
// For no expiration use 0.
//
// For KEEPTTL use -1.
//
// For more options, use SetArgs.
func (c *Compat) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd {
	var resp rueidis.RedisResult
	if expiration > 0 {
		if usePrecise(expiration) {
			resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).PxMilliseconds(formatMs(expiration)).Build())
		} else {
			resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).ExSeconds(formatSec(expiration)).Build())
		}
	} else if expiration == KeepTTL {
		resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).Keepttl().Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).Build())
	}
	return newStatusCmd(resp)
}

func (c *Compat) SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *StatusCmd {
	cmd := c.client.B().Arbitrary("SET").Keys(key).Args(str(value))
	if a.KeepTTL {
		cmd = cmd.Args("KEEPTTL")
	}
	if !a.ExpireAt.IsZero() {
		cmd = cmd.Args("EXAT", strconv.FormatInt(a.ExpireAt.Unix(), 10))
	}
	if a.TTL > 0 {
		if usePrecise(a.TTL) {
			cmd = cmd.Args("PX", strconv.FormatInt(formatMs(a.TTL), 10))
		} else {
			cmd = cmd.Args("EX", strconv.FormatInt(formatSec(a.TTL), 10))
		}
	}
	switch mode := strings.ToUpper(a.Mode); mode {
	case "XX", "NX":
		cmd = cmd.Args(mode)
	case "":
	default:
		panic(fmt.Sprintf("invalid mode for SET: %s", a.Mode))
	}
	if a.Get {
		cmd = cmd.Args("GET")
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newStatusCmd(resp)
}

func (c *Compat) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd {
	cmd := c.client.B().Setex().Key(key).Seconds(formatSec(expiration)).Value(str(value)).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd {
	var resp rueidis.RedisResult

	switch expiration {
	case 0:
		resp = c.client.Do(ctx, c.client.B().Setnx().Key(key).Value(str(value)).Build())
	case KeepTTL:
		resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).Keepttl().Nx().Build())
	default:
		if usePrecise(expiration) {
			resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).PxMilliseconds(formatMs(expiration)).Nx().Build())
		} else {
			resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).ExSeconds(formatSec(expiration)).Nx().Build())
		}
	}

	return newBoolCmd(resp)
}

func (c *Compat) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd {
	var resp rueidis.RedisResult
	if expiration > 0 {
		if usePrecise(expiration) {
			resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).PxMilliseconds(formatMs(expiration)).Xx().Build())
		} else {
			resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).ExSeconds(formatSec(expiration)).Xx().Build())
		}
	} else if expiration == KeepTTL {
		resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).Keepttl().Xx().Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(str(value)).Xx().Build())
	}
	return newBoolCmd(resp)
}

func (c *Compat) SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd {
	cmd := c.client.B().Setrange().Key(key).Offset(offset).Value(value).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) StrLen(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Strlen().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) Copy(ctx context.Context, source string, destination string, db int64, replace bool) *IntCmd {
	var resp rueidis.RedisResult
	if replace {
		resp = c.client.Do(ctx, c.client.B().Copy().Source(source).Destination(destination).Db(db).Replace().Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Copy().Source(source).Destination(destination).Db(db).Build())
	}
	return newIntCmd(resp)
}

func (c *Compat) GetBit(ctx context.Context, key string, offset int64) *IntCmd {
	cmd := c.client.B().Getbit().Key(key).Offset(offset).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) SetBit(ctx context.Context, key string, offset int64, value int64) *IntCmd {
	cmd := c.client.B().Setbit().Key(key).Offset(offset).Value(value).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd {
	var resp rueidis.RedisResult
	if bitCount == nil {
		resp = c.client.Do(ctx, c.client.B().Bitcount().Key(key).Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Bitcount().Key(key).Start(bitCount.Start).End(bitCount.End).Build())
	}
	return newIntCmd(resp)
}

func (c *Compat) BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd {
	cmd := c.client.B().Bitop().Operation("AND").Destkey(destKey).Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd {
	cmd := c.client.B().Bitop().Operation("OR").Destkey(destKey).Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd {
	cmd := c.client.B().Bitop().Operation("XOR").Destkey(destKey).Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) BitOpNot(ctx context.Context, destKey string, key string) *IntCmd {
	cmd := c.client.B().Bitop().Operation("NOT").Destkey(destKey).Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd {
	var resp rueidis.RedisResult
	switch len(pos) {
	case 0:
		resp = c.client.Do(ctx, c.client.B().Bitpos().Key(key).Bit(bit).Build())
	case 1:
		resp = c.client.Do(ctx, c.client.B().Bitpos().Key(key).Bit(bit).Start(pos[0]).Build())
	case 2:
		resp = c.client.Do(ctx, c.client.B().Bitpos().Key(key).Bit(bit).Start(pos[0]).End(pos[1]).Build())
	default:
		panic("too many arguments")
	}
	return newIntCmd(resp)
}

func (c *Compat) BitField(ctx context.Context, key string, args ...interface{}) *IntSliceCmd {
	cmd := c.client.B().Arbitrary("BITFIELD").Keys(key)
	for _, v := range args {
		cmd = cmd.Args(str(v))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntSliceCmd(resp)
}

func (c *Compat) Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd {
	cmd := c.client.B().Arbitrary("SCAN", strconv.FormatInt(int64(cursor), 10))
	if match != "" {
		cmd = cmd.Args("MATCH", match)
	}
	if count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(count, 10))
	}
	resp := c.client.Do(ctx, cmd.ReadOnly())
	return newScanCmd(resp)
}

func (c *Compat) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd {
	cmd := c.client.B().Arbitrary("SCAN", strconv.FormatInt(int64(cursor), 10))
	if match != "" {
		cmd = cmd.Args("MATCH", match)
	}
	if count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(count, 10))
	}
	resp := c.client.Do(ctx, cmd.Args("TYPE", keyType).ReadOnly())
	return newScanCmd(resp)
}

func (c *Compat) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	cmd := c.client.B().Arbitrary("SSCAN").Keys(key).Args(strconv.FormatInt(int64(cursor), 10))
	if match != "" {
		cmd = cmd.Args("MATCH", match)
	}
	if count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(count, 10))
	}
	resp := c.client.Do(ctx, cmd.ReadOnly())
	return newScanCmd(resp)
}

func (c *Compat) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	cmd := c.client.B().Arbitrary("HSCAN").Keys(key).Args(strconv.FormatInt(int64(cursor), 10))
	if match != "" {
		cmd = cmd.Args("MATCH", match)
	}
	if count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(count, 10))
	}
	resp := c.client.Do(ctx, cmd.ReadOnly())
	return newScanCmd(resp)
}

func (c *Compat) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	cmd := c.client.B().Arbitrary("ZSCAN").Keys(key).Args(strconv.FormatInt(int64(cursor), 10))
	if match != "" {
		cmd = cmd.Args("MATCH", match)
	}
	if count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(count, 10))
	}
	resp := c.client.Do(ctx, cmd.ReadOnly())
	return newScanCmd(resp)
}

func (c *Compat) HDel(ctx context.Context, key string, fields ...string) *IntCmd {
	cmd := c.client.B().Hdel().Key(key).Field(fields...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) HExists(ctx context.Context, key, field string) *BoolCmd {
	cmd := c.client.B().Hexists().Key(key).Field(field).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) HGet(ctx context.Context, key, field string) *StringCmd {
	cmd := c.client.B().Hget().Key(key).Field(field).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) HGetAll(ctx context.Context, key string) *StringStringMapCmd {
	cmd := c.client.B().Hgetall().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringStringMapCmd(resp)
}

func (c *Compat) HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd {
	cmd := c.client.B().Hincrby().Key(key).Field(field).Increment(incr).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd {
	cmd := c.client.B().Hincrbyfloat().Key(key).Field(field).Increment(incr).Build()
	resp := c.client.Do(ctx, cmd)
	return newFloatCmd(resp)
}

func (c *Compat) HKeys(ctx context.Context, key string) *StringSliceCmd {
	cmd := c.client.B().Hkeys().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) HLen(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Hlen().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) HMGet(ctx context.Context, key string, fields ...string) *SliceCmd {
	cmd := c.client.B().Hmget().Key(key).Field(fields...).Build()
	resp := c.client.Do(ctx, cmd)
	return newSliceCmd(resp)
}

// HSet requires Redis v4 for multiple field/value pairs support.
func (c *Compat) HSet(ctx context.Context, key string, values ...interface{}) *IntCmd {
	partial := c.client.B().Hset().Key(key).FieldValue()

	args := argsToSlice(values)
	for i := 0; i < len(args); i += 2 {
		partial = partial.FieldValue(args[i], args[i+1])
	}
	cmd := partial.Build()

	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

// HMSet is a deprecated version of HSet left for compatibility with Redis 3.
func (c *Compat) HMSet(ctx context.Context, key string, values ...interface{}) *BoolCmd {
	partial := c.client.B().Hset().Key(key).FieldValue()

	args := argsToSlice(values)
	for i := 0; i < len(args); i += 2 {
		partial = partial.FieldValue(args[i], args[i+1])
	}
	cmd := partial.Build()

	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd {
	cmd := c.client.B().Hsetnx().Key(key).Field(field).Value(str(value)).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) HVals(ctx context.Context, key string) *StringSliceCmd {
	cmd := c.client.B().Hvals().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) HRandField(ctx context.Context, key string, count int64, withValues bool) *StringSliceCmd {
	var resp rueidis.RedisResult
	if withValues {
		resp = c.client.Do(ctx, c.client.B().Hrandfield().Key(key).Count(count).Withvalues().Build())
		return flattenStringSliceCmd(resp)
	} else {
		resp = c.client.Do(ctx, c.client.B().Hrandfield().Key(key).Count(count).Build())
		return newStringSliceCmd(resp)
	}
}

func (c *Compat) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {
	cmd := c.client.B().Blpop().Key(keys...).Timeout(float64(formatSec(timeout))).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {
	cmd := c.client.B().Brpop().Key(keys...).Timeout(float64(formatSec(timeout))).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd {
	cmd := c.client.B().Brpoplpush().Source(source).Destination(destination).Timeout(float64(formatSec(timeout))).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) LIndex(ctx context.Context, key string, index int64) *StringCmd {
	cmd := c.client.B().Lindex().Key(key).Index(index).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) LInsert(ctx context.Context, key, op string, pivot, element interface{}) *IntCmd {
	var resp rueidis.RedisResult
	switch strings.ToUpper(op) {
	case "BEFORE":
		resp = c.client.Do(ctx, c.client.B().Linsert().Key(key).Before().Pivot(str(pivot)).Element(str(element)).Build())
	case "AFTER":
		resp = c.client.Do(ctx, c.client.B().Linsert().Key(key).After().Pivot(str(pivot)).Element(str(element)).Build())
	default:
		panic(fmt.Sprintf("Invalid op argument value: %s", op))
	}
	return newIntCmd(resp)
}

func (c *Compat) LInsertBefore(ctx context.Context, key string, pivot, element interface{}) *IntCmd {
	cmd := c.client.B().Linsert().Key(key).Before().Pivot(str(pivot)).Element(str(element)).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LInsertAfter(ctx context.Context, key string, pivot, element interface{}) *IntCmd {
	cmd := c.client.B().Linsert().Key(key).After().Pivot(str(pivot)).Element(str(element)).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LLen(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Llen().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LPop(ctx context.Context, key string) *StringCmd {
	cmd := c.client.B().Lpop().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) LPopCount(ctx context.Context, key string, count int64) *StringSliceCmd {
	cmd := c.client.B().Lpop().Key(key).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) LPos(ctx context.Context, key string, element string, a LPosArgs) *IntCmd {
	cmd := c.client.B().Arbitrary("LPOS").Keys(key).Args(element)
	if a.Rank != 0 {
		cmd = cmd.Args("RANK", strconv.FormatInt(a.Rank, 10))
	}
	if a.MaxLen != 0 {
		cmd = cmd.Args("MAXLEN", strconv.FormatInt(a.MaxLen, 10))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

func (c *Compat) LPosCount(ctx context.Context, key string, element string, count int64, a LPosArgs) *IntSliceCmd {
	cmd := c.client.B().Arbitrary("LPOS").Keys(key).Args(element).Args("COUNT", strconv.FormatInt(count, 10))
	if a.Rank != 0 {
		cmd = cmd.Args("RANK", strconv.FormatInt(a.Rank, 10))
	}
	if a.MaxLen != 0 {
		cmd = cmd.Args("MAXLEN", strconv.FormatInt(a.MaxLen, 10))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntSliceCmd(resp)
}

func (c *Compat) LPush(ctx context.Context, key string, elements ...interface{}) *IntCmd {
	cmd := c.client.B().Lpush().Key(key).Element(argsToSlice(elements)...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LPushX(ctx context.Context, key string, elements ...interface{}) *IntCmd {
	cmd := c.client.B().Lpushx().Key(key).Element(argsToSlice(elements)...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	cmd := c.client.B().Lrange().Key(key).Start(start).Stop(stop).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) LRem(ctx context.Context, key string, count int64, element interface{}) *IntCmd {
	cmd := c.client.B().Lrem().Key(key).Count(count).Element(str(element)).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LSet(ctx context.Context, key string, index int64, element interface{}) *StatusCmd {
	cmd := c.client.B().Lset().Key(key).Index(index).Element(str(element)).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd {
	cmd := c.client.B().Ltrim().Key(key).Start(start).Stop(stop).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) RPop(ctx context.Context, key string) *StringCmd {
	cmd := c.client.B().Rpop().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) RPopCount(ctx context.Context, key string, count int64) *StringSliceCmd {
	cmd := c.client.B().Rpop().Key(key).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) RPopLPush(ctx context.Context, source, destination string) *StringCmd {
	cmd := c.client.B().Rpoplpush().Source(source).Destination(destination).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) RPush(ctx context.Context, key string, elements ...interface{}) *IntCmd {
	cmd := c.client.B().Rpush().Key(key).Element(argsToSlice(elements)...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) RPushX(ctx context.Context, key string, elements ...interface{}) *IntCmd {
	cmd := c.client.B().Rpushx().Key(key).Element(argsToSlice(elements)...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd {
	var resp rueidis.RedisResult
	switch strings.ToUpper(srcpos + destpos) {
	case "LEFTLEFT":
		resp = c.client.Do(ctx, c.client.B().Lmove().Source(source).Destination(destination).Left().Left().Build())
	case "LEFTRIGHT":
		resp = c.client.Do(ctx, c.client.B().Lmove().Source(source).Destination(destination).Left().Right().Build())
	case "RIGHTLEFT":
		resp = c.client.Do(ctx, c.client.B().Lmove().Source(source).Destination(destination).Right().Left().Build())
	case "RIGHTRIGHT":
		resp = c.client.Do(ctx, c.client.B().Lmove().Source(source).Destination(destination).Right().Right().Build())
	default:
		panic(fmt.Sprintf("Invalid srcpost + destpos argument value: %s", srcpos+destpos))
	}
	return newStringCmd(resp)
}

func (c *Compat) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *StringCmd {
	var resp rueidis.RedisResult
	switch strings.ToUpper(srcpos + destpos) {
	case "LEFTLEFT":
		resp = c.client.Do(ctx, c.client.B().Blmove().Source(source).Destination(destination).Left().Left().Timeout(float64(formatSec(timeout))).Build())
	case "LEFTRIGHT":
		resp = c.client.Do(ctx, c.client.B().Blmove().Source(source).Destination(destination).Left().Right().Timeout(float64(formatSec(timeout))).Build())
	case "RIGHTLEFT":
		resp = c.client.Do(ctx, c.client.B().Blmove().Source(source).Destination(destination).Right().Left().Timeout(float64(formatSec(timeout))).Build())
	case "RIGHTRIGHT":
		resp = c.client.Do(ctx, c.client.B().Blmove().Source(source).Destination(destination).Right().Right().Timeout(float64(formatSec(timeout))).Build())
	default:
		panic(fmt.Sprintf("Invalid srcpost + destpos argument value: %s", srcpos+destpos))
	}
	return newStringCmd(resp)
}

func (c *Compat) SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd {
	cmd := c.client.B().Sadd().Key(key).Member()
	for _, m := range argsToSlice(members) {
		cmd = cmd.Member(str(m))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

func (c *Compat) SCard(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Scard().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) SDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	cmd := c.client.B().Sdiff().Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	cmd := c.client.B().Sdiffstore().Destination(destination).Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) SInter(ctx context.Context, keys ...string) *StringSliceCmd {
	cmd := c.client.B().Sinter().Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	cmd := c.client.B().Sinterstore().Destination(destination).Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd {
	cmd := c.client.B().Sismember().Key(key).Member(str(member)).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd {
	cmd := c.client.B().Smismember().Key(key).Member(argsToSlice(members)...).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolSliceCmd(resp)
}

func (c *Compat) SMembers(ctx context.Context, key string) *StringSliceCmd {
	cmd := c.client.B().Smembers().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) SMembersMap(ctx context.Context, key string) *StringStructMapCmd {
	cmd := c.client.B().Smembers().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringStructMapCmd(resp)
}

func (c *Compat) SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd {
	cmd := c.client.B().Smove().Source(source).Destination(destination).Member(str(member)).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) SPop(ctx context.Context, key string) *StringCmd {
	cmd := c.client.B().Spop().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) SPopN(ctx context.Context, key string, count int64) *StringSliceCmd {
	cmd := c.client.B().Spop().Key(key).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) SRandMember(ctx context.Context, key string) *StringCmd {
	cmd := c.client.B().Srandmember().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd {
	cmd := c.client.B().Srandmember().Key(key).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) SRem(ctx context.Context, key string, members ...interface{}) *IntCmd {
	cmd := c.client.B().Srem().Key(key).Member(argsToSlice(members)...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) SUnion(ctx context.Context, keys ...string) *StringSliceCmd {
	cmd := c.client.B().Sunion().Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	cmd := c.client.B().Sunionstore().Destination(destination).Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) XAdd(ctx context.Context, a XAddArgs) *StringCmd {
	cmd := c.client.B().Arbitrary("XADD").Keys(a.Stream)
	if a.NoMkStream {
		cmd = cmd.Args("NOMKSTREAM")
	}
	switch {
	case a.MaxLen > 0:
		if a.Approx {
			cmd = cmd.Args("MAXLEN", "~", strconv.FormatInt(a.MaxLen, 10))
		} else {
			cmd = cmd.Args("MAXLEN", strconv.FormatInt(a.MaxLen, 10))
		}
	case a.MinID != "":
		if a.Approx {
			cmd = cmd.Args("MINID", "~", a.MinID)
		} else {
			cmd = cmd.Args("MINID", a.MinID)
		}
	}
	if a.Limit > 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(a.Limit, 10))
	}
	if a.ID != "" {
		cmd = cmd.Args(a.ID)
	} else {
		cmd = cmd.Args("*")
	}
	cmd = cmd.Args(argToSlice(a.Values)...)
	resp := c.client.Do(ctx, cmd.Build())
	return newStringCmd(resp)
}

func (c *Compat) XDel(ctx context.Context, stream string, ids ...string) *IntCmd {
	cmd := c.client.B().Xdel().Key(stream).Id(ids...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) XLen(ctx context.Context, stream string) *IntCmd {
	cmd := c.client.B().Xlen().Key(stream).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd {
	cmd := c.client.B().Xrange().Key(stream).Start(start).End(stop).Build()
	resp := c.client.Do(ctx, cmd)
	return newXMessageSliceCmd(resp)
}

func (c *Compat) XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd {
	cmd := c.client.B().Xrange().Key(stream).Start(start).End(stop).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newXMessageSliceCmd(resp)
}

func (c *Compat) XRevRange(ctx context.Context, stream, stop, start string) *XMessageSliceCmd {
	cmd := c.client.B().Xrevrange().Key(stream).End(stop).Start(start).Build()
	resp := c.client.Do(ctx, cmd)
	return newXMessageSliceCmd(resp)
}

func (c *Compat) XRevRangeN(ctx context.Context, stream, stop, start string, count int64) *XMessageSliceCmd {
	cmd := c.client.B().Xrevrange().Key(stream).End(stop).Start(start).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newXMessageSliceCmd(resp)
}

func (c *Compat) XRead(ctx context.Context, a XReadArgs) *XStreamSliceCmd {
	cmd := c.client.B().Arbitrary("XREAD")
	if a.Count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(a.Count, 10))
	}
	if a.Block >= 0 {
		cmd = cmd.Args("BLOCK", strconv.FormatInt(formatMs(a.Block), 10))
	}
	cmd = cmd.Args("STREAMS")
	cmd = cmd.Keys(a.Streams[:a.Count]...).Args(a.Streams[a.Count:]...)
	resp := c.client.Do(ctx, cmd.Build())
	return newXStreamSliceCmd(resp)
}

func (c *Compat) XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd {
	return c.XRead(ctx, XReadArgs{Streams: streams, Block: -1})
}

func (c *Compat) XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd {
	cmd := c.client.B().XgroupCreate().Key(stream).Groupname(group).Id(start).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd {
	cmd := c.client.B().XgroupCreate().Key(stream).Groupname(group).Id(start).Mkstream().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd {
	cmd := c.client.B().XgroupSetid().Key(stream).Groupname(group).Id(start).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) XGroupDestroy(ctx context.Context, stream, group string) *IntCmd {
	cmd := c.client.B().XgroupDestroy().Key(stream).Groupname(group).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd {
	cmd := c.client.B().XgroupCreateconsumer().Key(stream).Groupname(group).Consumername(consumer).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd {
	cmd := c.client.B().XgroupDelconsumer().Key(stream).Groupname(group).Consumername(consumer).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) XReadGroup(ctx context.Context, a XReadGroupArgs) *XStreamSliceCmd {
	cmd := c.client.B().Arbitrary("XREADGROUP")
	cmd = cmd.Args("GROUP", a.Group, a.Consumer)
	if a.Count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(a.Count, 10))
	}
	if a.Block >= 0 {
		cmd = cmd.Args("BLOCK", strconv.FormatInt(formatMs(a.Block), 10))
	}
	if a.NoAck {
		cmd = cmd.Args("NOACK")
	}
	cmd = cmd.Args("STREAMS")
	cmd = cmd.Keys(a.Streams[:a.Count]...).Args(a.Streams[a.Count:]...)
	resp := c.client.Do(ctx, cmd.Build())
	return newXStreamSliceCmd(resp)
}

func (c *Compat) XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd {
	cmd := c.client.B().Xack().Key(stream).Group(group).Id(ids...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) XPending(ctx context.Context, stream, group string) *XPendingCmd {
	cmd := c.client.B().Xpending().Key(stream).Group(group).Build()
	resp := c.client.Do(ctx, cmd)
	return newXPendingCmd(resp)
}

func (c *Compat) XPendingExt(ctx context.Context, a XPendingExtArgs) *XPendingExtCmd {
	cmd := c.client.B().Arbitrary("XPENDING").Keys(a.Stream).Args(a.Group)
	if a.Idle != 0 {
		cmd = cmd.Args("IDLE", strconv.FormatInt(formatMs(a.Idle), 10))
	}
	cmd = cmd.Args(a.Start, a.End, strconv.FormatInt(a.Count, 10))
	if a.Consumer != "" {
		cmd = cmd.Args(a.Consumer)
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newXPendingExtCmd(resp)
}

func (c *Compat) XClaim(ctx context.Context, a XClaimArgs) *XMessageSliceCmd {
	cmd := c.client.B().Xclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Id(a.Messages...).Build()
	resp := c.client.Do(ctx, cmd)
	return newXMessageSliceCmd(resp)
}

func (c *Compat) XClaimJustID(ctx context.Context, a XClaimArgs) *StringSliceCmd {
	cmd := c.client.B().Xclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Id(a.Messages...).Justid().Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) XAutoClaim(ctx context.Context, a XAutoClaimArgs) *XAutoClaimCmd {
	var resp rueidis.RedisResult
	if a.Count > 0 {
		resp = c.client.Do(ctx, c.client.B().Xautoclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Start(a.Start).Count(a.Count).Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Xautoclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Start(a.Start).Build())
	}
	return newXAutoClaimCmd(resp)
}

func (c *Compat) XAutoClaimJustID(ctx context.Context, a XAutoClaimArgs) *XAutoClaimJustIDCmd {
	var resp rueidis.RedisResult
	if a.Count > 0 {
		resp = c.client.Do(ctx, c.client.B().Xautoclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Start(a.Start).Count(a.Count).Justid().Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Xautoclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Start(a.Start).Justid().Build())
	}
	return newXAutoClaimJustIDCmd(resp)
}

// xTrim If approx is true, add the "~" parameter, otherwise it is the default "=" (redis default).
// example:
//		XTRIM key MAXLEN/MINID threshold LIMIT limit.
//		XTRIM key MAXLEN/MINID ~ threshold LIMIT limit.
// The redis-server version is lower than 6.2, please set limit to 0.
func (c *Compat) xTrim(ctx context.Context, key, strategy string,
	approx bool, threshold string, limit int64) *IntCmd {
	cmd := c.client.B().Arbitrary("XTRIM").Keys(key).Args(strategy)
	if approx {
		cmd = cmd.Args("~")
	}
	cmd = cmd.Args(threshold)
	if limit > 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(limit, 10))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

// XTrimMaxLen No `~` rules are used, `limit` cannot be used.
// cmd: XTRIM key MAXLEN maxLen
func (c *Compat) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd {
	return c.xTrim(ctx, key, "MAXLEN", false, strconv.FormatInt(maxLen, 10), 0)
}

// XTrimMaxLenApprox LIMIT has a bug, please confirm it and use it.
// issue: https://github.com/redis/redis/issues/9046
// cmd: XTRIM key MAXLEN ~ maxLen LIMIT limit
func (c *Compat) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd {
	return c.xTrim(ctx, key, "MAXLEN", true, strconv.FormatInt(maxLen, 10), limit)
}

// XTrimMinID No `~` rules are used, `limit` cannot be used.
// cmd: XTRIM key MINID minID
func (c *Compat) XTrimMinID(ctx context.Context, key string, minID string) *IntCmd {
	return c.xTrim(ctx, key, "MINID", false, minID, 0)
}

// XTrimMinIDApprox LIMIT has a bug, please confirm it and use it.
// issue: https://github.com/redis/redis/issues/9046
// cmd: XTRIM key MINID ~ minID LIMIT limit
func (c *Compat) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd {
	return c.xTrim(ctx, key, "MINID", true, minID, limit)
}

func (c *Compat) XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd {
	cmd := c.client.B().XinfoGroups().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newXInfoGroupsCmd(resp)
}

func (c *Compat) XInfoStream(ctx context.Context, key string) *XInfoStreamCmd {
	cmd := c.client.B().XinfoStream().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newXInfoStreamCmd(resp)
}

func (c *Compat) XInfoStreamFull(ctx context.Context, key string, count int64) *XInfoStreamFullCmd {
	cmd := c.client.B().XinfoStream().Key(key).Full().Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newXInfoStreamFullCmd(resp)
}

func (c *Compat) XInfoConsumers(ctx context.Context, key, group string) *XInfoConsumersCmd {
	cmd := c.client.B().XinfoConsumers().Key(key).Groupname(group).Build()
	resp := c.client.Do(ctx, cmd)
	return newXInfoConsumersCmd(resp)
}

// BZPopMax Redis `BZPOPMAX key [key ...] timeout` command.
func (c *Compat) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {
	cmd := c.client.B().Bzpopmax().Key(keys...).Timeout(float64(formatSec(timeout))).Build()
	resp := c.client.Do(ctx, cmd)
	return newZWithKeyCmd(resp)
}

// BZPopMin Redis `BZPOPMIN key [key ...] timeout` command.
func (c *Compat) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {
	cmd := c.client.B().Bzpopmin().Key(keys...).Timeout(float64(formatSec(timeout))).Build()
	resp := c.client.Do(ctx, cmd)
	return newZWithKeyCmd(resp)
}

// ZAdd Redis `ZADD key score member [score member ...]` command.
func (c *Compat) ZAdd(ctx context.Context, key string, members ...Z) *IntCmd {
	cmd := c.client.B().Zadd().Key(key).ScoreMember()
	for _, v := range members {
		cmd = cmd.ScoreMember(v.Score, str(v.Member))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

// ZAddNX Redis `ZADD key NX score member [score member ...]` command.
func (c *Compat) ZAddNX(ctx context.Context, key string, members ...Z) *IntCmd {
	cmd := c.client.B().Zadd().Key(key).Nx().ScoreMember()
	for _, v := range members {
		cmd = cmd.ScoreMember(v.Score, str(v.Member))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

// ZAddXX Redis `ZADD key XX score member [score member ...]` command.
func (c *Compat) ZAddXX(ctx context.Context, key string, members ...Z) *IntCmd {
	cmd := c.client.B().Zadd().Key(key).Xx().ScoreMember()
	for _, v := range members {
		cmd = cmd.ScoreMember(v.Score, str(v.Member))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

func (c *Compat) zAddArgs(ctx context.Context, key string, incr bool, args ZAddArgs) rueidis.RedisResult {
	cmd := c.client.B().Arbitrary("ZADD").Keys(key)
	// The GT, LT and NX options are mutually exclusive.
	if args.NX {
		cmd = cmd.Args("NX")
	} else {
		if args.XX {
			cmd = cmd.Args("XX")
		}
		if args.GT {
			cmd = cmd.Args("GT")
		} else if args.LT {
			cmd = cmd.Args("LT")
		}
	}
	if args.Ch {
		cmd = cmd.Args("CH")
	}
	if incr {
		cmd = cmd.Args("INCR")
	}
	for _, v := range args.Members {
		cmd = cmd.Args(strconv.FormatFloat(v.Score, 'f', -1, 64), str(v.Member))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return resp
}

func (c *Compat) ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd {
	return newIntCmd(c.zAddArgs(ctx, key, false, args))
}

func (c *Compat) ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd {
	return newFloatCmd(c.zAddArgs(ctx, key, true, args))
}

func (c *Compat) ZCard(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Zcard().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ZCount(ctx context.Context, key string, min, max string) *IntCmd {
	cmd := c.client.B().Zcount().Key(key).Min(min).Max(max).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ZLexCount(ctx context.Context, key, min, max string) *IntCmd {
	cmd := c.client.B().Zlexcount().Key(key).Min(min).Max(max).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd {
	cmd := c.client.B().Zincrby().Key(key).Increment(increment).Member(member).Build()
	resp := c.client.Do(ctx, cmd)
	return newFloatCmd(resp)
}

func (c *Compat) ZInter(ctx context.Context, store ZStore) *StringSliceCmd {
	var resp rueidis.RedisResult
	if len(store.Weights) > 0 {
		switch strings.ToUpper(store.Aggregate) {
		case "SUM":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateSum().Build())
		case "MIN":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMin().Build())
		case "MAX":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMax().Build())
		case "":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).Build())
		default:
			panic(fmt.Sprintf("invalid aggregate argument value %s", store.Aggregate))
		}
	} else {
		switch strings.ToUpper(store.Aggregate) {
		case "SUM":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).AggregateSum().Build())
		case "MIN":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).AggregateMin().Build())
		case "MAX":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).AggregateMax().Build())
		case "":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Build())
		default:
			panic(fmt.Sprintf("invalid aggregate argument value %s", store.Aggregate))
		}
	}
	return newStringSliceCmd(resp)
}

func (c *Compat) ZInterWithScores(ctx context.Context, store ZStore) *ZSliceCmd {
	var resp rueidis.RedisResult
	if len(store.Weights) > 0 {
		switch strings.ToUpper(store.Aggregate) {
		case "SUM":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateSum().Withscores().Build())
		case "MIN":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMin().Withscores().Build())
		case "MAX":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMax().Withscores().Build())
		case "":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).Withscores().Build())
		default:
			panic(fmt.Sprintf("invalid aggregate argument value %s", store.Aggregate))
		}
	} else {
		switch strings.ToUpper(store.Aggregate) {
		case "SUM":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).AggregateSum().Withscores().Build())
		case "MIN":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).AggregateMin().Withscores().Build())
		case "MAX":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).AggregateMax().Withscores().Build())
		case "":
			resp = c.client.Do(ctx, c.client.B().Zinter().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Withscores().Build())
		default:
			panic(fmt.Sprintf("invalid aggregate argument value %s", store.Aggregate))
		}
	}
	return newZSliceCmd(resp)
}

func (c *Compat) ZInterStore(ctx context.Context, destination string, store ZStore) *IntCmd {
	var resp rueidis.RedisResult
	switch strings.ToUpper(store.Aggregate) {
	case "SUM":
		resp = c.client.Do(ctx, c.client.B().Zinterstore().Destination(destination).Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateSum().Build())
	case "MIN":
		resp = c.client.Do(ctx, c.client.B().Zinterstore().Destination(destination).Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMin().Build())
	case "MAX":
		resp = c.client.Do(ctx, c.client.B().Zinterstore().Destination(destination).Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMax().Build())
	case "":
		resp = c.client.Do(ctx, c.client.B().Zinterstore().Destination(destination).Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).Build())
	default:
		panic(fmt.Sprintf("invalid aggregate argument value %s", store.Aggregate))
	}
	return newIntCmd(resp)
}

func (c *Compat) ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd {
	cmd := c.client.B().Zmscore().Key(key).Member(members...).Build()
	resp := c.client.Do(ctx, cmd)
	return newFloatSliceCmd(resp)
}

func (c *Compat) ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd {
	var resp rueidis.RedisResult
	switch len(count) {
	case 0:
		resp = c.client.Do(ctx, c.client.B().Zpopmax().Key(key).Build())
	case 1:
		resp = c.client.Do(ctx, c.client.B().Zpopmax().Key(key).Count(count[0]).Build())
		if count[0] > 1 {
			return newZSliceCmd(resp)
		}
	default:
		panic("too many arguments")
	}
	return newZSliceSingleCmd(resp)
}

func (c *Compat) ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd {
	var resp rueidis.RedisResult
	switch len(count) {
	case 0:
		resp = c.client.Do(ctx, c.client.B().Zpopmin().Key(key).Build())
	case 1:
		resp = c.client.Do(ctx, c.client.B().Zpopmin().Key(key).Count(count[0]).Build())
		if count[0] > 1 {
			return newZSliceCmd(resp)
		}
	default:
		panic("too many arguments")
	}
	return newZSliceSingleCmd(resp)
}

func (c *Compat) zRangeArgs(withScores bool, z ZRangeArgs) cmds.Completed {
	cmd := c.client.B().Arbitrary("ZRANGE").Keys(z.Key)
	if z.Rev && (z.ByScore || z.ByLex) {
		cmd = cmd.Args(str(z.Stop), str(z.Start))
	} else {
		cmd = cmd.Args(str(z.Start), str(z.Stop))
	}
	if z.ByScore {
		cmd = cmd.Args("BYSCORE")
	} else if z.ByLex {
		cmd = cmd.Args("BYLEX")
	}
	if z.Rev {
		cmd = cmd.Args("REV")
	}
	if z.Offset != 0 || z.Count != 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(z.Offset, 10), strconv.FormatInt(z.Count, 10))
	}
	if withScores {
		cmd = cmd.Args("WITHSCORES")
	}
	return cmd.Build()
}

func (c *Compat) ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	cmd := c.zRangeArgs(false, ZRangeArgs{
		Key:   key,
		Start: start,
		Stop:  stop,
	})
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	cmd := c.zRangeArgs(true, ZRangeArgs{
		Key:   key,
		Start: start,
		Stop:  stop,
	})
	resp := c.client.Do(ctx, cmd)
	return newZSliceCmd(resp)
}

func (c *Compat) ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.Do(ctx, c.client.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Limit(opt.Offset, opt.Count).Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Build())
	}
	return newStringSliceCmd(resp)
}

func (c *Compat) ZRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.Do(ctx, c.client.B().Zrangebylex().Key(key).Min(opt.Min).Max(opt.Max).Limit(opt.Offset, opt.Count).Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Zrangebylex().Key(key).Min(opt.Min).Max(opt.Max).Build())
	}
	return newStringSliceCmd(resp)
}

func (c *Compat) ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.Do(ctx, c.client.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Withscores().Limit(opt.Offset, opt.Count).Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Withscores().Build())
	}
	return newZSliceCmd(resp)
}

func (c *Compat) ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd {
	cmd := c.zRangeArgs(false, z)
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd {
	cmd := c.zRangeArgs(true, z)
	resp := c.client.Do(ctx, cmd)
	return newZSliceCmd(resp)
}

func (c *Compat) ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd {
	cmd := c.client.B().Arbitrary("ZRANGESTORE").Keys(dst, z.Key)
	if z.Rev && (z.ByScore || z.ByLex) {
		cmd = cmd.Args(str(z.Stop), str(z.Start))
	} else {
		cmd = cmd.Args(str(z.Start), str(z.Stop))
	}
	if z.ByScore {
		cmd = cmd.Args("BYSCORE")
	} else if z.ByLex {
		cmd = cmd.Args("BYLEX")
	}
	if z.Rev {
		cmd = cmd.Args("REV")
	}
	if z.Offset != 0 || z.Count != 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(z.Offset, 10), strconv.FormatInt(z.Count, 10))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

func (c *Compat) ZRank(ctx context.Context, key, member string) *IntCmd {
	cmd := c.client.B().Zrank().Key(key).Member(member).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd {
	cmd := c.client.B().Zrem().Key(key).Member(argsToSlice(members)...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd {
	cmd := c.client.B().Zremrangebyrank().Key(key).Start(start).Stop(stop).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}
func (c *Compat) ZRemRangeByScore(ctx context.Context, key string, min, max string) *IntCmd {
	cmd := c.client.B().Zremrangebyscore().Key(key).Min(min).Max(max).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ZRemRangeByLex(ctx context.Context, key string, min, max string) *IntCmd {
	cmd := c.client.B().Zremrangebylex().Key(key).Min(min).Max(max).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	cmd := c.client.B().Zrevrange().Key(key).Start(start).Stop(stop).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	cmd := c.client.B().Zrevrange().Key(key).Start(start).Stop(stop).Withscores().Build()
	resp := c.client.Do(ctx, cmd)
	return newZSliceCmd(resp)
}

func (c *Compat) ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.Do(ctx, c.client.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Limit(opt.Offset, opt.Count).Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Build())
	}
	return newStringSliceCmd(resp)
}

func (c *Compat) ZRevRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.Do(ctx, c.client.B().Zrevrangebylex().Key(key).Max(opt.Max).Min(opt.Min).Limit(opt.Offset, opt.Count).Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Zrevrangebylex().Key(key).Max(opt.Max).Min(opt.Min).Build())
	}
	return newStringSliceCmd(resp)
}

func (c *Compat) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.Do(ctx, c.client.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Withscores().Limit(opt.Offset, opt.Count).Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Withscores().Build())
	}
	return newZSliceCmd(resp)
}

func (c *Compat) ZRevRank(ctx context.Context, key, member string) *IntCmd {
	cmd := c.client.B().Zrevrank().Key(key).Member(member).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ZScore(ctx context.Context, key, member string) *FloatCmd {
	cmd := c.client.B().Zscore().Key(key).Member(member).Build()
	resp := c.client.Do(ctx, cmd)
	return newFloatCmd(resp)
}

func (c *Compat) ZUnionStore(ctx context.Context, dest string, store ZStore) *IntCmd {
	var resp rueidis.RedisResult
	switch strings.ToUpper(store.Aggregate) {
	case "SUM":
		resp = c.client.Do(ctx, c.client.B().Zunionstore().Destination(dest).Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateSum().Build())
	case "MIN":
		resp = c.client.Do(ctx, c.client.B().Zunionstore().Destination(dest).Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMin().Build())
	case "MAX":
		resp = c.client.Do(ctx, c.client.B().Zunionstore().Destination(dest).Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMax().Build())
	case "":
		resp = c.client.Do(ctx, c.client.B().Zunionstore().Destination(dest).Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).Build())
	default:
		panic(fmt.Sprintf("invalid aggregate argument value %s", store.Aggregate))
	}
	return newIntCmd(resp)
}

func (c *Compat) ZUnion(ctx context.Context, store ZStore) *StringSliceCmd {
	var resp rueidis.RedisResult
	switch strings.ToUpper(store.Aggregate) {
	case "SUM":
		resp = c.client.Do(ctx, c.client.B().Zunion().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateSum().Build())
	case "MIN":
		resp = c.client.Do(ctx, c.client.B().Zunion().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMin().Build())
	case "MAX":
		resp = c.client.Do(ctx, c.client.B().Zunion().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMax().Build())
	case "":
		resp = c.client.Do(ctx, c.client.B().Zunion().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).Build())
	default:
		panic(fmt.Sprintf("invalid aggregate argument value %s", store.Aggregate))
	}
	return newStringSliceCmd(resp)
}

func (c *Compat) ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd {
	var resp rueidis.RedisResult
	switch strings.ToUpper(store.Aggregate) {
	case "SUM":
		resp = c.client.Do(ctx, c.client.B().Zunion().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateSum().Withscores().Build())
	case "MIN":
		resp = c.client.Do(ctx, c.client.B().Zunion().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMin().Withscores().Build())
	case "MAX":
		resp = c.client.Do(ctx, c.client.B().Zunion().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).AggregateMax().Withscores().Build())
	case "":
		resp = c.client.Do(ctx, c.client.B().Zunion().Numkeys(int64(len(store.Keys))).Key(store.Keys...).Weights(store.Weights...).Withscores().Build())
	default:
		panic(fmt.Sprintf("invalid aggregate argument value %s", store.Aggregate))
	}
	return newZSliceCmd(resp)
}

func (c *Compat) ZRandMember(ctx context.Context, key string, count int64, withScores bool) *StringSliceCmd {
	var resp rueidis.RedisResult
	if withScores {
		resp = c.client.Do(ctx, c.client.B().Zrandmember().Key(key).Count(count).Withscores().Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Zrandmember().Key(key).Count(count).Build())
	}
	return newStringSliceCmd(resp)
}

func (c *Compat) ZDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	cmd := c.client.B().Zdiff().Numkeys(int64(len(keys))).Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd {
	cmd := c.client.B().Zdiff().Numkeys(int64(len(keys))).Key(keys...).Withscores().Build()
	resp := c.client.Do(ctx, cmd)
	return newZSliceCmd(resp)
}

func (c *Compat) ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	cmd := c.client.B().Zdiffstore().Destination(destination).Numkeys(int64(len(keys))).Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd {
	cmd := c.client.B().Pfadd().Key(key).Element(argsToSlice(els)...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) PFCount(ctx context.Context, keys ...string) *IntCmd {
	cmd := c.client.B().Pfcount().Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd {
	cmd := c.client.B().Pfmerge().Destkey(dest).Sourcekey(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) BgRewriteAOF(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Bgrewriteaof().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) BgSave(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Bgsave().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClientKill(ctx context.Context, ipPort string) *StatusCmd {
	cmd := c.client.B().ClientKill().IpPort(ipPort).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

// ClientKillByFilter is new style syntax, while the ClientKill is old
//
//   CLIENT KILL <option> [value] ... <option> [value]
func (c *Compat) ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd {
	cmd := c.client.B().Arbitrary("CLIENT KILL").Args(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ClientList(ctx context.Context) *StringCmd {
	cmd := c.client.B().ClientList().Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) ClientPause(ctx context.Context, dur time.Duration) *BoolCmd {
	cmd := c.client.B().ClientPause().Timeout(formatSec(dur)).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) ClientID(ctx context.Context) *IntCmd {
	cmd := c.client.B().ClientId().Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ConfigGet(ctx context.Context, parameter string) *SliceCmd {
	cmd := c.client.B().ConfigGet().Parameter(parameter).Build()
	resp := c.client.Do(ctx, cmd)
	return newSliceCmdFromMap(resp)
}

func (c *Compat) ConfigResetStat(ctx context.Context) *StatusCmd {
	cmd := c.client.B().ConfigResetstat().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ConfigSet(ctx context.Context, parameter, value string) *StatusCmd {
	cmd := c.client.B().ConfigSet().ParameterValue().ParameterValue(parameter, value).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ConfigRewrite(ctx context.Context) *StatusCmd {
	cmd := c.client.B().ConfigRewrite().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) DBSize(ctx context.Context) *IntCmd {
	cmd := c.client.B().Dbsize().Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) FlushAll(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Flushall().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) FlushAllAsync(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Flushall().Async().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) FlushDB(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Flushdb().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) FlushDBAsync(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Flushdb().Async().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) Info(ctx context.Context, section ...string) *StringCmd {
	cmd := c.client.B().Info().Section(section...).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) LastSave(ctx context.Context) *IntCmd {
	cmd := c.client.B().Lastsave().Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) Save(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Save().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) Shutdown(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Shutdown().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ShutdownSave(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Shutdown().Save().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ShutdownNoSave(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Shutdown().Nosave().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) SlaveOf(ctx context.Context, host string, port string) *StatusCmd {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return &StatusCmd{err: err}
	}
	cmd := c.client.B().Slaveof().Host(host).Port(p).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) Time(ctx context.Context) *TimeCmd {
	cmd := c.client.B().Time().Build()
	resp := c.client.Do(ctx, cmd)
	return newTimeCmd(resp)
}

func (c *Compat) DebugObject(ctx context.Context, key string) *StringCmd {
	cmd := c.client.B().DebugObject().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) ReadOnly(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Readonly().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ReadWrite(ctx context.Context) *StatusCmd {
	cmd := c.client.B().Readwrite().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) MemoryUsage(ctx context.Context, key string, samples ...int64) *IntCmd {
	var resp rueidis.RedisResult
	switch len(samples) {
	case 0:
		resp = c.client.Do(ctx, c.client.B().MemoryUsage().Key(key).Build())
	case 1:
		resp = c.client.Do(ctx, c.client.B().MemoryUsage().Key(key).Samples(samples[0]).Build())
	default:
		panic("too many arguments")
	}
	return newIntCmd(resp)
}

func (c *Compat) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd {
	cmd := c.client.B().Eval().Script(script).Numkeys(int64(len(keys))).Key(keys...).Arg(argsToSlice(args)...).Build()
	return newCmd(c.client.Do(ctx, cmd))
}

func (c *Compat) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd {
	cmd := c.client.B().Evalsha().Sha1(sha1).Numkeys(int64(len(keys))).Key(keys...).Arg(argsToSlice(args)...).Build()
	return newCmd(c.client.Do(ctx, cmd))
}

func (c *Compat) ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd {
	cmd := c.client.B().ScriptExists().Sha1(hashes...).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolSliceCmd(resp)
}

func (c *Compat) ScriptFlush(ctx context.Context) *StatusCmd {
	cmd := c.client.B().ScriptFlush().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ScriptKill(ctx context.Context) *StatusCmd {
	cmd := c.client.B().ScriptKill().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ScriptLoad(ctx context.Context, script string) *StringCmd {
	cmd := c.client.B().ScriptLoad().Script(script).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) Publish(ctx context.Context, channel string, message interface{}) *IntCmd {
	cmd := c.client.B().Publish().Channel(channel).Message(str(message)).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd {
	cmd := c.client.B().PubsubChannels().Pattern(pattern).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) PubSubNumSub(ctx context.Context, channels ...string) *StringIntMapCmd {
	cmd := c.client.B().PubsubNumsub().Channel(channels...).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringIntMapCmd(resp)
}

func (c *Compat) PubSubNumPat(ctx context.Context) *IntCmd {
	cmd := c.client.B().PubsubNumpat().Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ClusterSlots(ctx context.Context) *ClusterSlotsCmd {
	cmd := c.client.B().ClusterSlots().Build()
	resp := c.client.Do(ctx, cmd)
	return newClusterSlotsCmd(resp)
}

func (c *Compat) ClusterNodes(ctx context.Context) *StringCmd {
	cmd := c.client.B().ClusterNodes().Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) ClusterMeet(ctx context.Context, host string, port int64) *StatusCmd {
	cmd := c.client.B().ClusterMeet().Ip(host).Port(port).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClusterForget(ctx context.Context, nodeID string) *StatusCmd {
	cmd := c.client.B().ClusterForget().NodeId(nodeID).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd {
	cmd := c.client.B().ClusterReplicate().NodeId(nodeID).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClusterResetSoft(ctx context.Context) *StatusCmd {
	cmd := c.client.B().ClusterReset().Soft().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClusterResetHard(ctx context.Context) *StatusCmd {
	cmd := c.client.B().ClusterReset().Hard().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClusterInfo(ctx context.Context) *StringCmd {
	cmd := c.client.B().ClusterInfo().Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) ClusterKeySlot(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().ClusterKeyslot().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ClusterGetKeysInSlot(ctx context.Context, slot int64, count int64) *StringSliceCmd {
	cmd := c.client.B().ClusterGetkeysinslot().Slot(slot).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd {
	cmd := c.client.B().ClusterCountFailureReports().NodeId(nodeID).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ClusterCountKeysInSlot(ctx context.Context, slot int64) *IntCmd {
	cmd := c.client.B().ClusterCountkeysinslot().Slot(slot).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) ClusterDelSlots(ctx context.Context, slots ...int64) *StatusCmd {
	cmd := c.client.B().ClusterDelslots().Slot(slots...).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClusterDelSlotsRange(ctx context.Context, min, max int64) *StatusCmd {
	cmd := c.client.B().ClusterDelslotsrange().StartSlotEndSlot().StartSlotEndSlot(min, max).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClusterSaveConfig(ctx context.Context) *StatusCmd {
	cmd := c.client.B().ClusterSaveconfig().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd {
	cmd := c.client.B().ClusterSlaves().NodeId(nodeID).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) ClusterFailover(ctx context.Context) *StatusCmd {
	cmd := c.client.B().ClusterFailover().Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClusterAddSlots(ctx context.Context, slots ...int64) *StatusCmd {
	cmd := c.client.B().ClusterAddslots().Slot(slots...).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) ClusterAddSlotsRange(ctx context.Context, min, max int64) *StatusCmd {
	cmd := c.client.B().ClusterAddslotsrange().StartSlotEndSlot().StartSlotEndSlot(min, max).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) GeoAdd(ctx context.Context, key string, geoLocation ...GeoLocation) *IntCmd {
	cmd := c.client.B().Geoadd().Key(key).LongitudeLatitudeMember()
	for _, loc := range geoLocation {
		cmd = cmd.LongitudeLatitudeMember(loc.Longitude, loc.Latitude, loc.Name)
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

func (c *Compat) GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd {
	cmd := c.client.B().Geopos().Key(key).Member(members...).Build()
	resp := c.client.Do(ctx, cmd)
	return newGeoPosCmd(resp)
}

// GeoRadius is a read-only GEORADIUS_RO command.
func (c *Compat) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *GeoLocationCmd {
	cmd := c.client.B().Arbitrary("GEORADIUS_RO").Keys(key).Args(strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64))
	if query.Store != "" || query.StoreDist != "" {
		panic("GeoRadius does not support Store or StoreDist")
	}
	cmd = cmd.Args(query.args()...)
	resp := c.client.Do(ctx, cmd.Build())
	return newGeoLocationCmd(resp, query.WithDist, query.WithGeoHash, query.WithCoord)
}

// GeoRadiusStore is a writing GEORADIUS command.
func (c *Compat) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *IntCmd {
	cmd := c.client.B().Arbitrary("GEORADIUS").Keys(key).Args(strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64))
	if query.Store == "" && query.StoreDist == "" {
		panic("GeoRadiusStore requires Store or StoreDist")
	}
	cmd = cmd.Args(query.args()...)
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

// GeoRadiusByMember is a read-only GEORADIUSBYMEMBER_RO command.
func (c *Compat) GeoRadiusByMember(ctx context.Context, key, member string, query GeoRadiusQuery) *GeoLocationCmd {
	cmd := c.client.B().Arbitrary("GEORADIUSBYMEMBER_RO").Keys(key).Args(member)
	if query.Store != "" || query.StoreDist != "" {
		panic("GeoRadiusByMember does not support Store or StoreDist")
	}
	cmd = cmd.Args(query.args()...)
	resp := c.client.Do(ctx, cmd.Build())
	return newGeoLocationCmd(resp, query.WithDist, query.WithGeoHash, query.WithCoord)
}

// GeoRadiusByMemberStore is a writing GEORADIUSBYMEMBER command.
func (c *Compat) GeoRadiusByMemberStore(ctx context.Context, key, member string, query GeoRadiusQuery) *IntCmd {
	cmd := c.client.B().Arbitrary("GEORADIUSBYMEMBER").Keys(key).Args(member)
	if query.Store != "" && query.StoreDist != "" {
		panic("GeoRadiusByMemberStore requires Store or StoreDist")
	}
	cmd = cmd.Args(query.args()...)
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

func (c *Compat) GeoSearch(ctx context.Context, key string, q GeoSearchQuery) *StringSliceCmd {
	cmd := c.client.B().Arbitrary("GEOSEARCH").Keys(key)
	cmd = cmd.Args(q.args()...)
	resp := c.client.Do(ctx, cmd.Build())
	return newStringSliceCmd(resp)
}

func (c *Compat) GeoSearchLocation(ctx context.Context, key string, q GeoSearchLocationQuery) *GeoLocationCmd {
	cmd := c.client.B().Arbitrary("GEOSEARCH").Keys(key)
	cmd = cmd.Args(q.args()...)
	resp := c.client.Do(ctx, cmd.Build())
	return newGeoLocationCmd(resp, q.WithDist, q.WithHash, q.WithCoord)
}

func (c *Compat) GeoSearchStore(ctx context.Context, src, dest string, q GeoSearchStoreQuery) *IntCmd {
	cmd := c.client.B().Arbitrary("GEOSEARCHSTORE").Keys(dest, src)
	cmd = cmd.Args(q.args()...)
	if q.StoreDist {
		cmd = cmd.Args("STOREDIST")
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntCmd(resp)
}

func (c *Compat) GeoDist(ctx context.Context, key, member1, member2, unit string) *FloatCmd {
	var resp rueidis.RedisResult
	switch strings.ToUpper(unit) {
	case "M":
		resp = c.client.Do(ctx, c.client.B().Geodist().Key(key).Member1(member1).Member2(member2).M().Build())
	case "MI":
		resp = c.client.Do(ctx, c.client.B().Geodist().Key(key).Member1(member1).Member2(member2).Mi().Build())
	case "FT":
		resp = c.client.Do(ctx, c.client.B().Geodist().Key(key).Member1(member1).Member2(member2).Ft().Build())
	case "KM", "":
		resp = c.client.Do(ctx, c.client.B().Geodist().Key(key).Member1(member1).Member2(member2).Km().Build())
	default:
		panic(fmt.Sprintf("invalid unit %s", unit))
	}
	return newFloatCmd(resp)
}

func (c *Compat) GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd {
	cmd := c.client.B().Geohash().Key(key).Member(members...).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c CacheCompat) BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd {
	var resp rueidis.RedisResult
	if bitCount == nil {
		resp = c.client.DoCache(ctx, c.client.B().Bitcount().Key(key).Cache(), c.ttl)
	} else {
		resp = c.client.DoCache(ctx, c.client.B().Bitcount().Key(key).Start(bitCount.Start).End(bitCount.End).Cache(), c.ttl)
	}
	return newIntCmd(resp)
}

func (c CacheCompat) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd {
	var resp rueidis.RedisResult
	switch len(pos) {
	case 0:
		resp = c.client.DoCache(ctx, c.client.B().Bitpos().Key(key).Bit(bit).Cache(), c.ttl)
	case 1:
		resp = c.client.DoCache(ctx, c.client.B().Bitpos().Key(key).Bit(bit).Start(pos[0]).Cache(), c.ttl)
	case 2:
		resp = c.client.DoCache(ctx, c.client.B().Bitpos().Key(key).Bit(bit).Start(pos[0]).End(pos[1]).Cache(), c.ttl)
	default:
		panic("too many arguments")
	}
	return newIntCmd(resp)
}

func (c CacheCompat) GeoDist(ctx context.Context, key, member1, member2, unit string) *FloatCmd {
	var resp rueidis.RedisResult
	switch strings.ToUpper(unit) {
	case "M":
		resp = c.client.DoCache(ctx, c.client.B().Geodist().Key(key).Member1(member1).Member2(member2).M().Cache(), c.ttl)
	case "MI":
		resp = c.client.DoCache(ctx, c.client.B().Geodist().Key(key).Member1(member1).Member2(member2).Mi().Cache(), c.ttl)
	case "FT":
		resp = c.client.DoCache(ctx, c.client.B().Geodist().Key(key).Member1(member1).Member2(member2).Ft().Cache(), c.ttl)
	case "KM", "":
		resp = c.client.DoCache(ctx, c.client.B().Geodist().Key(key).Member1(member1).Member2(member2).Km().Cache(), c.ttl)
	default:
		panic(fmt.Sprintf("invalid unit %s", unit))
	}
	return newFloatCmd(resp)
}

func (c CacheCompat) GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd {
	cmd := c.client.B().Geohash().Key(key).Member(members...).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringSliceCmd(resp)
}

func (c CacheCompat) GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd {
	cmd := c.client.B().Geopos().Key(key).Member(members...).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newGeoPosCmd(resp)
}

// GeoRadius is a read-only GEORADIUS_RO command.
func (c CacheCompat) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *GeoLocationCmd {
	cmd := c.client.B().Arbitrary("GEORADIUS_RO").Keys(key).Args(strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64))
	if query.Store != "" || query.StoreDist != "" {
		panic("GeoRadius does not support Store or StoreDist")
	}
	cmd = cmd.Args(query.args()...)
	resp := c.client.DoCache(ctx, cmds.Cacheable(cmd.Build()), c.ttl)
	return newGeoLocationCmd(resp, query.WithDist, query.WithGeoHash, query.WithCoord)
}

// GeoRadiusByMember is a read-only GEORADIUSBYMEMBER_RO command.
func (c CacheCompat) GeoRadiusByMember(ctx context.Context, key, member string, query GeoRadiusQuery) *GeoLocationCmd {
	cmd := c.client.B().Arbitrary("GEORADIUSBYMEMBER_RO").Keys(key).Args(member)
	if query.Store != "" || query.StoreDist != "" {
		panic("GeoRadiusByMember does not support Store or StoreDist")
	}
	cmd = cmd.Args(query.args()...)
	resp := c.client.DoCache(ctx, cmds.Cacheable(cmd.Build()), c.ttl)
	return newGeoLocationCmd(resp, query.WithDist, query.WithGeoHash, query.WithCoord)
}

func (c CacheCompat) GeoSearch(ctx context.Context, key string, q GeoSearchQuery) *StringSliceCmd {
	cmd := c.client.B().Arbitrary("GEOSEARCH").Keys(key)
	cmd = cmd.Args(q.args()...)
	resp := c.client.DoCache(ctx, cmds.Cacheable(cmd.Build()), c.ttl)
	return newStringSliceCmd(resp)
}

func (c CacheCompat) Get(ctx context.Context, key string) *StringCmd {
	cmd := c.client.B().Get().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringCmd(resp)
}

func (c CacheCompat) GetBit(ctx context.Context, key string, offset int64) *IntCmd {
	cmd := c.client.B().Getbit().Key(key).Offset(offset).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) GetRange(ctx context.Context, key string, start, end int64) *StringCmd {
	cmd := c.client.B().Getrange().Key(key).Start(start).End(end).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringCmd(resp)
}

func (c CacheCompat) HExists(ctx context.Context, key, field string) *BoolCmd {
	cmd := c.client.B().Hexists().Key(key).Field(field).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newBoolCmd(resp)
}

func (c CacheCompat) HGet(ctx context.Context, key, field string) *StringCmd {
	cmd := c.client.B().Hget().Key(key).Field(field).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringCmd(resp)
}

func (c CacheCompat) HGetAll(ctx context.Context, key string) *StringStringMapCmd {
	cmd := c.client.B().Hgetall().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringStringMapCmd(resp)
}

func (c CacheCompat) HKeys(ctx context.Context, key string) *StringSliceCmd {
	cmd := c.client.B().Hkeys().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringSliceCmd(resp)
}

func (c CacheCompat) HLen(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Hlen().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) HMGet(ctx context.Context, key string, fields ...string) *SliceCmd {
	cmd := c.client.B().Hmget().Key(key).Field(fields...).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newSliceCmd(resp)
}

func (c CacheCompat) HVals(ctx context.Context, key string) *StringSliceCmd {
	cmd := c.client.B().Hvals().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringSliceCmd(resp)
}

func (c CacheCompat) LIndex(ctx context.Context, key string, index int64) *StringCmd {
	cmd := c.client.B().Lindex().Key(key).Index(index).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringCmd(resp)
}

func (c CacheCompat) LLen(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Llen().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) LPos(ctx context.Context, key string, element string, a LPosArgs) *IntCmd {
	cmd := c.client.B().Arbitrary("LPOS").Keys(key).Args(element)
	if a.Rank != 0 {
		cmd = cmd.Args("RANK", strconv.FormatInt(a.Rank, 10))
	}
	if a.MaxLen != 0 {
		cmd = cmd.Args("MAXLEN", strconv.FormatInt(a.MaxLen, 10))
	}
	resp := c.client.DoCache(ctx, cmds.Cacheable(cmd.Build()), c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	cmd := c.client.B().Lrange().Key(key).Start(start).Stop(stop).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringSliceCmd(resp)
}

func (c CacheCompat) PTTL(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Pttl().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) SCard(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Scard().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd {
	cmd := c.client.B().Sismember().Key(key).Member(str(member)).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newBoolCmd(resp)
}

func (c CacheCompat) SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd {
	cmd := c.client.B().Smismember().Key(key).Member(argsToSlice(members)...).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newBoolSliceCmd(resp)
}

func (c CacheCompat) SMembers(ctx context.Context, key string) *StringSliceCmd {
	cmd := c.client.B().Smembers().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringSliceCmd(resp)
}

func (c CacheCompat) Sort(ctx context.Context, key string, sort Sort) *StringSliceCmd {
	cmd := c.client.B().Arbitrary("SORT").Keys(key)
	if sort.By != "" {
		cmd = cmd.Args("BY", sort.By)
	}
	if sort.Offset != 0 || sort.Count != 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(sort.Offset, 10), strconv.FormatInt(sort.Count, 10))
	}
	if len(sort.Get) > 0 {
		cmd = cmd.Args("GET").Args(sort.Get...)
	}
	switch order := strings.ToUpper(sort.Order); order {
	case "ASC", "DESC":
		cmd = cmd.Args(order)
	case "":
	default:
		panic(fmt.Sprintf("invalid sort order %s", sort.Order))
	}
	if sort.Alpha {
		cmd = cmd.Args("ALPHA")
	}
	resp := c.client.DoCache(ctx, cmds.Cacheable(cmd.Build()), c.ttl)
	return newStringSliceCmd(resp)
}

func (c CacheCompat) StrLen(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Strlen().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) TTL(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Ttl().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) Type(ctx context.Context, key string) *StatusCmd {
	cmd := c.client.B().Type().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStatusCmd(resp)
}

func (c CacheCompat) ZCard(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Zcard().Key(key).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) ZCount(ctx context.Context, key string, min, max string) *IntCmd {
	cmd := c.client.B().Zcount().Key(key).Min(min).Max(max).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) ZLexCount(ctx context.Context, key, min, max string) *IntCmd {
	cmd := c.client.B().Zlexcount().Key(key).Min(min).Max(max).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd {
	cmd := c.client.B().Zmscore().Key(key).Member(members...).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newFloatSliceCmd(resp)
}

func (c CacheCompat) zRangeArgs(withScores bool, z ZRangeArgs) cmds.Cacheable {
	cmd := c.client.B().Arbitrary("ZRANGE").Keys(z.Key)
	if z.Rev && (z.ByScore || z.ByLex) {
		cmd = cmd.Args(str(z.Stop), str(z.Start))
	} else {
		cmd = cmd.Args(str(z.Start), str(z.Stop))
	}
	if z.ByScore {
		cmd = cmd.Args("BYSCORE")
	} else if z.ByLex {
		cmd = cmd.Args("BYLEX")
	}
	if z.Rev {
		cmd = cmd.Args("REV")
	}
	if z.Offset != 0 || z.Count != 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(z.Offset, 10), strconv.FormatInt(z.Count, 10))
	}
	if withScores {
		cmd = cmd.Args("WITHSCORES")
	}
	return cmds.Cacheable(cmd.Build())
}

func (c CacheCompat) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	cmd := c.zRangeArgs(true, ZRangeArgs{
		Key:   key,
		Start: start,
		Stop:  stop,
	})
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newZSliceCmd(resp)
}

func (c CacheCompat) ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.DoCache(ctx, c.client.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Limit(opt.Offset, opt.Count).Cache(), c.ttl)
	} else {
		resp = c.client.DoCache(ctx, c.client.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Cache(), c.ttl)
	}
	return newStringSliceCmd(resp)
}

func (c CacheCompat) ZRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.DoCache(ctx, c.client.B().Zrangebylex().Key(key).Min(opt.Min).Max(opt.Max).Limit(opt.Offset, opt.Count).Cache(), c.ttl)
	} else {
		resp = c.client.DoCache(ctx, c.client.B().Zrangebylex().Key(key).Min(opt.Min).Max(opt.Max).Cache(), c.ttl)
	}
	return newStringSliceCmd(resp)
}

func (c CacheCompat) ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.DoCache(ctx, c.client.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Withscores().Limit(opt.Offset, opt.Count).Cache(), c.ttl)
	} else {
		resp = c.client.DoCache(ctx, c.client.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Withscores().Cache(), c.ttl)
	}
	return newZSliceCmd(resp)
}

func (c CacheCompat) ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd {
	cmd := c.zRangeArgs(false, z)
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringSliceCmd(resp)
}

func (c CacheCompat) ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd {
	cmd := c.zRangeArgs(true, z)
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newZSliceCmd(resp)
}

func (c CacheCompat) ZRank(ctx context.Context, key, member string) *IntCmd {
	cmd := c.client.B().Zrank().Key(key).Member(member).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	cmd := c.client.B().Zrevrange().Key(key).Start(start).Stop(stop).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newStringSliceCmd(resp)
}

func (c CacheCompat) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	cmd := c.client.B().Zrevrange().Key(key).Start(start).Stop(stop).Withscores().Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newZSliceCmd(resp)
}

func (c CacheCompat) ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.DoCache(ctx, c.client.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Limit(opt.Offset, opt.Count).Cache(), c.ttl)
	} else {
		resp = c.client.DoCache(ctx, c.client.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Cache(), c.ttl)
	}
	return newStringSliceCmd(resp)
}

func (c CacheCompat) ZRevRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.DoCache(ctx, c.client.B().Zrevrangebylex().Key(key).Max(opt.Max).Min(opt.Min).Limit(opt.Offset, opt.Count).Cache(), c.ttl)
	} else {
		resp = c.client.DoCache(ctx, c.client.B().Zrevrangebylex().Key(key).Max(opt.Max).Min(opt.Min).Cache(), c.ttl)
	}
	return newStringSliceCmd(resp)
}

func (c CacheCompat) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd {
	var resp rueidis.RedisResult
	if opt.Offset != 0 || opt.Count != 0 {
		resp = c.client.DoCache(ctx, c.client.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Withscores().Limit(opt.Offset, opt.Count).Cache(), c.ttl)
	} else {
		resp = c.client.DoCache(ctx, c.client.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Withscores().Cache(), c.ttl)
	}
	return newZSliceCmd(resp)
}

func (c CacheCompat) ZRevRank(ctx context.Context, key, member string) *IntCmd {
	cmd := c.client.B().Zrevrank().Key(key).Member(member).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newIntCmd(resp)
}

func (c CacheCompat) ZScore(ctx context.Context, key, member string) *FloatCmd {
	cmd := c.client.B().Zscore().Key(key).Member(member).Cache()
	resp := c.client.DoCache(ctx, cmd, c.ttl)
	return newFloatCmd(resp)
}

func str(arg interface{}) string {
	if v, ok := arg.(encoding.BinaryMarshaler); ok {
		if data, err := v.MarshalBinary(); err == nil {
			return rueidis.BinaryString(data)
		}
	}
	if v, ok := arg.(string); ok {
		return v
	}
	if v, ok := arg.([]byte); ok {
		return string(v)
	}
	if v, ok := arg.(bool); ok {
		if v {
			return "1"
		}
		return "0"
	}
	return fmt.Sprint(arg)
}

func argsToSlice(src []interface{}) []string {
	if len(src) == 1 {
		return argToSlice(src[0])
	}
	dst := make([]string, 0, len(src))
	for _, v := range src {
		dst = append(dst, str(v))
	}
	return dst
}

func argToSlice(arg interface{}) []string {
	switch arg := arg.(type) {
	case []string:
		return arg
	case []interface{}:
		dst := make([]string, 0, len(arg))
		for _, v := range arg {
			dst = append(dst, str(v))
		}
		return dst
	case map[string]interface{}:
		dst := make([]string, 0, len(arg))
		for k, v := range arg {
			dst = append(dst, k, str(v))
		}
		return dst
	case map[string]string:
		dst := make([]string, 0, len(arg))
		for k, v := range arg {
			dst = append(dst, k, v)
		}
		return dst
	default:
		return []string{str(arg)}
	}
}
