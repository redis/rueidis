package rueidiscompat

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/cmds"
)

type Cmdable interface {
	Echo(ctx context.Context, message string) *StringCmd
	Ping(ctx context.Context, message string) *StatusCmd
	Quit(ctx context.Context) *StatusCmd
	Del(ctx context.Context, keys ...string) *IntCmd
	Unlink(ctx context.Context, keys ...string) *IntCmd
	Dump(ctx context.Context, key string) *StringCmd
	Exists(ctx context.Context, keys ...string) *IntCmd
	Expire(ctx context.Context, key string, seconds time.Duration) *BoolCmd
	ExpireAt(ctx context.Context, key string, timestamp time.Time) *BoolCmd
	ExpireNX(ctx context.Context, key string, seconds time.Duration) *BoolCmd
	ExpireXX(ctx context.Context, key string, seconds time.Duration) *BoolCmd
	ExpireGT(ctx context.Context, key string, seconds time.Duration) *BoolCmd
	ExpireLT(ctx context.Context, key string, seconds time.Duration) *BoolCmd
	Keys(ctx context.Context, pattern string) *StringSliceCmd
	Migrate(ctx context.Context, host string, port int64, key bool, db int64, timeout time.Duration) *StatusCmd
	Move(ctx context.Context, key string, db int64) *BoolCmd
	ObjectRefCount(ctx context.Context, key string) *IntCmd
	ObjectEncoding(ctx context.Context, key string) *StringCmd
	ObjectIdleTime(ctx context.Context, key string) *IntCmd
	Persist(ctx context.Context, key string) *BoolCmd
	PExpire(ctx context.Context, key string, milliseconds time.Duration) *BoolCmd
	PExpireAt(ctx context.Context, key string, millisecondsTimestamp time.Time) *BoolCmd
	PTTL(ctx context.Context, key string) *IntCmd
	RandomKey(ctx context.Context) *StringCmd
	Rename(ctx context.Context, key, newkey string) *StatusCmd
	RenameNX(ctx context.Context, key, newkey string) *BoolCmd
	Restore(ctx context.Context, key string, ttl time.Duration, serializedValue string) *StatusCmd
	RestoreReplace(ctx context.Context, key string, ttl time.Duration, serializedValue string) *StatusCmd
	Sort(ctx context.Context, key string, sort Sort) *StringSliceCmd
	SortStore(ctx context.Context, key, store string, sort Sort) *IntCmd
	Touch(ctx context.Context, keys ...string) *IntCmd
	TTL(ctx context.Context, key string) *IntCmd
	Type(ctx context.Context, key string) *StatusCmd
	Append(ctx context.Context, key, value string) *IntCmd
	Decr(ctx context.Context, key string) *IntCmd
	DecrBy(ctx context.Context, key string, decrement int64) *IntCmd
	Get(ctx context.Context, key string) *StringCmd
	GetRange(ctx context.Context, key string, start, end int64) *StringCmd
	GetSet(ctx context.Context, key, value string) *StringCmd
	GetEx(ctx context.Context, key string, seconds time.Duration) *StringCmd
	GetDel(ctx context.Context, key string) *StringCmd
	Incr(ctx context.Context, key string) *IntCmd
	IncrBy(ctx context.Context, key string, increment int64) *IntCmd
	IncrByFloat(ctx context.Context, key string, increment float64) *FloatCmd
	MGet(ctx context.Context, keys ...string) *SliceCmd
	MSet(ctx context.Context, keys []string, values []string) *StatusCmd
	MSetNX(ctx context.Context, keys []string, values []string) *BoolCmd
	Set(ctx context.Context, key string, value string, expiration time.Duration) *StatusCmd
	SetArgs(ctx context.Context, key string, value string, a SetArgs) *StatusCmd
	SetEx(ctx context.Context, key string, value string, expiration time.Duration) *StatusCmd
	SetNX(ctx context.Context, key string, value string) *BoolCmd
	SetXX(ctx context.Context, key string, value string, expiration time.Duration) *BoolCmd
	SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd
	StrLen(ctx context.Context, key string) *IntCmd
	Copy(ctx context.Context, source string, destination string, db int64, replace bool) *IntCmd

	GetBit(ctx context.Context, key string, offset int64) *IntCmd
	SetBit(ctx context.Context, key string, offset int64, value int64) *IntCmd
	BitCount(ctx context.Context, key string, bitCount BitCount) *IntCmd
	BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpNot(ctx context.Context, destKey string, key string) *IntCmd
	BitPos(ctx context.Context, key string, bit int64, bitPos BitPos) *IntCmd
	BitField(ctx context.Context, key string, bitField []BitField) *IntSliceCmd

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
	HSet(ctx context.Context, key string, keys []string, values []string) *IntCmd
	HMSet(ctx context.Context, key string, keys []string, values []string) *BoolCmd
	HSetNX(ctx context.Context, key, field string, value string) *BoolCmd
	HVals(ctx context.Context, key string) *StringSliceCmd
	HRandField(ctx context.Context, key string, count int64, withValues bool) *StringSliceCmd

	BLPop(ctx context.Context, timeout time.Duration, key string, keys ...string) *StringSliceCmd
	BRPop(ctx context.Context, timeout time.Duration, key string, keys ...string) *StringSliceCmd
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd
	LIndex(ctx context.Context, key string, index int64) *StringCmd
	LInsert(ctx context.Context, key, op, pivot, element string) *IntCmd
	LInsertBefore(ctx context.Context, key, pivot, element string) *IntCmd
	LInsertAfter(ctx context.Context, key, pivot, element string) *IntCmd
	LLen(ctx context.Context, key string) *IntCmd
	LPop(ctx context.Context, key string) *StringCmd
	LPopCount(ctx context.Context, key string, count int64) *StringSliceCmd
	LPos(ctx context.Context, key string, value string, a LPosArgs) *IntCmd
	LPosCount(ctx context.Context, key string, value string, count int64, args LPosArgs) *IntSliceCmd
	LPush(ctx context.Context, key string, elements ...string) *IntCmd
	LPushX(ctx context.Context, key string, elements ...string) *IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	LRem(ctx context.Context, key string, count int64, element string) *IntCmd
	LSet(ctx context.Context, key string, index int64, element string) *StatusCmd
	LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd
	RPop(ctx context.Context, key string) *StringCmd
	RPopCount(ctx context.Context, key string, count int64) *StringSliceCmd
	RPopLPush(ctx context.Context, source, destination string) *StringCmd
	RPush(ctx context.Context, key string, elements ...string) *IntCmd
	RPushX(ctx context.Context, key string, elements ...string) *IntCmd
	LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd
	BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *StringCmd
	// Implemented until here.
	// TODO:
	//
	// SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd
	// SCard(ctx context.Context, key string) *IntCmd
	// SDiff(ctx context.Context, keys ...string) *StringSliceCmd
	// SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd
	// SInter(ctx context.Context, keys ...string) *StringSliceCmd
	// SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd
	// SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd
	// SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd
	// SMembers(ctx context.Context, key string) *StringSliceCmd
	// SMembersMap(ctx context.Context, key string) *StringStructMapCmd
	// SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd
	// SPop(ctx context.Context, key string) *StringCmd
	// SPopN(ctx context.Context, key string, count int64) *StringSliceCmd
	// SRandMember(ctx context.Context, key string) *StringCmd
	// SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd
	// SRem(ctx context.Context, key string, members ...interface{}) *IntCmd
	// SUnion(ctx context.Context, keys ...string) *StringSliceCmd
	// SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd

	// XAdd(ctx context.Context, a *XAddArgs) *StringCmd
	// XDel(ctx context.Context, stream string, ids ...string) *IntCmd
	// XLen(ctx context.Context, stream string) *IntCmd
	// XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd
	// XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd
	// XRevRange(ctx context.Context, stream string, start, stop string) *XMessageSliceCmd
	// XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *XMessageSliceCmd
	// XRead(ctx context.Context, a *XReadArgs) *XStreamSliceCmd
	// XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd
	// XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd
	// XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd
	// XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd
	// XGroupDestroy(ctx context.Context, stream, group string) *IntCmd
	// XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
	// XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
	// XReadGroup(ctx context.Context, a *XReadGroupArgs) *XStreamSliceCmd
	// XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd
	// XPending(ctx context.Context, stream, group string) *XPendingCmd
	// XPendingExt(ctx context.Context, a *XPendingExtArgs) *XPendingExtCmd
	// XClaim(ctx context.Context, a *XClaimArgs) *XMessageSliceCmd
	// XClaimJustID(ctx context.Context, a *XClaimArgs) *StringSliceCmd
	// XAutoClaim(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimCmd
	// XAutoClaimJustID(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimJustIDCmd

	// // TODO: XTrim and XTrimApprox remove in v9.
	// XTrim(ctx context.Context, key string, maxLen int64) *IntCmd
	// XTrimApprox(ctx context.Context, key string, maxLen int64) *IntCmd
	// XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd
	// XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd
	// XTrimMinID(ctx context.Context, key string, minID string) *IntCmd
	// XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd
	// XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd
	// XInfoStream(ctx context.Context, key string) *XInfoStreamCmd
	// XInfoStreamFull(ctx context.Context, key string, count int) *XInfoStreamFullCmd
	// XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd

	// BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
	// BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd

	// // TODO: remove
	// //		ZAddCh
	// //		ZIncr
	// //		ZAddNXCh
	// //		ZAddXXCh
	// //		ZIncrNX
	// //		ZIncrXX
	// // 	in v9.
	// // 	use ZAddArgs and ZAddArgsIncr.

	// ZAdd(ctx context.Context, key string, members ...*Z) *IntCmd
	// ZAddNX(ctx context.Context, key string, members ...*Z) *IntCmd
	// ZAddXX(ctx context.Context, key string, members ...*Z) *IntCmd
	// ZAddCh(ctx context.Context, key string, members ...*Z) *IntCmd
	// ZAddNXCh(ctx context.Context, key string, members ...*Z) *IntCmd
	// ZAddXXCh(ctx context.Context, key string, members ...*Z) *IntCmd
	// ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd
	// ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd
	// ZIncr(ctx context.Context, key string, member *Z) *FloatCmd
	// ZIncrNX(ctx context.Context, key string, member *Z) *FloatCmd
	// ZIncrXX(ctx context.Context, key string, member *Z) *FloatCmd
	// ZCard(ctx context.Context, key string) *IntCmd
	// ZCount(ctx context.Context, key, min, max string) *IntCmd
	// ZLexCount(ctx context.Context, key, min, max string) *IntCmd
	// ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd
	// ZInter(ctx context.Context, store *ZStore) *StringSliceCmd
	// ZInterWithScores(ctx context.Context, store *ZStore) *ZSliceCmd
	// ZInterStore(ctx context.Context, destination string, store *ZStore) *IntCmd
	// ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd
	// ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd
	// ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd
	// ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	// ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	// ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
	// ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
	// ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
	// ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd
	// ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd
	// ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd
	// ZRank(ctx context.Context, key, member string) *IntCmd
	// ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd
	// ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd
	// ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd
	// ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd
	// ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	// ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	// ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
	// ZRevRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
	// ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
	// ZRevRank(ctx context.Context, key, member string) *IntCmd
	// ZScore(ctx context.Context, key, member string) *FloatCmd
	// ZUnionStore(ctx context.Context, dest string, store *ZStore) *IntCmd
	// ZUnion(ctx context.Context, store ZStore) *StringSliceCmd
	// ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd
	// ZRandMember(ctx context.Context, key string, count int, withScores bool) *StringSliceCmd
	// ZDiff(ctx context.Context, keys ...string) *StringSliceCmd
	// ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd
	// ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd

	// PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd
	// PFCount(ctx context.Context, keys ...string) *IntCmd
	// PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd

	// BgRewriteAOF(ctx context.Context) *StatusCmd
	// BgSave(ctx context.Context) *StatusCmd
	// ClientKill(ctx context.Context, ipPort string) *StatusCmd
	// ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd
	// ClientList(ctx context.Context) *StringCmd
	// ClientPause(ctx context.Context, dur time.Duration) *BoolCmd
	// ClientID(ctx context.Context) *IntCmd
	// ConfigGet(ctx context.Context, parameter string) *SliceCmd
	// ConfigResetStat(ctx context.Context) *StatusCmd
	// ConfigSet(ctx context.Context, parameter, value string) *StatusCmd
	// ConfigRewrite(ctx context.Context) *StatusCmd
	// DBSize(ctx context.Context) *IntCmd
	// FlushAll(ctx context.Context) *StatusCmd
	// FlushAllAsync(ctx context.Context) *StatusCmd
	// FlushDB(ctx context.Context) *StatusCmd
	// FlushDBAsync(ctx context.Context) *StatusCmd
	// Info(ctx context.Context, section ...string) *StringCmd
	// LastSave(ctx context.Context) *IntCmd
	// Save(ctx context.Context) *StatusCmd
	// Shutdown(ctx context.Context) *StatusCmd
	// ShutdownSave(ctx context.Context) *StatusCmd
	// ShutdownNoSave(ctx context.Context) *StatusCmd
	// SlaveOf(ctx context.Context, host, port string) *StatusCmd
	// Time(ctx context.Context) *TimeCmd
	// DebugObject(ctx context.Context, key string) *StringCmd
	// ReadOnly(ctx context.Context) *StatusCmd
	// ReadWrite(ctx context.Context) *StatusCmd
	// MemoryUsage(ctx context.Context, key string, samples ...int) *IntCmd

	// Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
	// EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd
	// ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd
	// ScriptFlush(ctx context.Context) *StatusCmd
	// ScriptKill(ctx context.Context) *StatusCmd
	// ScriptLoad(ctx context.Context, script string) *StringCmd

	// Publish(ctx context.Context, channel string, message interface{}) *IntCmd
	// PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd
	// PubSubNumSub(ctx context.Context, channels ...string) *StringIntMapCmd
	// PubSubNumPat(ctx context.Context) *IntCmd

	// ClusterSlots(ctx context.Context) *ClusterSlotsCmd
	// ClusterNodes(ctx context.Context) *StringCmd
	// ClusterMeet(ctx context.Context, host, port string) *StatusCmd
	// ClusterForget(ctx context.Context, nodeID string) *StatusCmd
	// ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd
	// ClusterResetSoft(ctx context.Context) *StatusCmd
	// ClusterResetHard(ctx context.Context) *StatusCmd
	// ClusterInfo(ctx context.Context) *StringCmd
	// ClusterKeySlot(ctx context.Context, key string) *IntCmd
	// ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *StringSliceCmd
	// ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd
	// ClusterCountKeysInSlot(ctx context.Context, slot int) *IntCmd
	// ClusterDelSlots(ctx context.Context, slots ...int) *StatusCmd
	// ClusterDelSlotsRange(ctx context.Context, min, max int) *StatusCmd
	// ClusterSaveConfig(ctx context.Context) *StatusCmd
	// ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd
	// ClusterFailover(ctx context.Context) *StatusCmd
	// ClusterAddSlots(ctx context.Context, slots ...int) *StatusCmd
	// ClusterAddSlotsRange(ctx context.Context, min, max int) *StatusCmd

	// GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) *IntCmd
	// GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd
	// GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *GeoLocationCmd
	// GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *IntCmd
	// GeoRadiusByMember(ctx context.Context, key, member string, query *GeoRadiusQuery) *GeoLocationCmd
	// GeoRadiusByMemberStore(ctx context.Context, key, member string, query *GeoRadiusQuery) *IntCmd
	// GeoSearch(ctx context.Context, key string, q *GeoSearchQuery) *StringSliceCmd
	// GeoSearchLocation(ctx context.Context, key string, q *GeoSearchLocationQuery) *GeoSearchLocationCmd
	// GeoSearchStore(ctx context.Context, key, store string, q *GeoSearchStoreQuery) *IntCmd
	// GeoDist(ctx context.Context, key string, member1, member2, unit string) *FloatCmd
	// GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd
}

type Compat struct {
	client rueidis.Client
}

func NewAdapter(client rueidis.Client) Cmdable {
	return &Compat{client: client}
}

func (c *Compat) Echo(ctx context.Context, message string) *StringCmd {
	cmd := c.client.B().Echo().Message(message).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringCmd(resp)
}

func (c *Compat) Ping(ctx context.Context, message string) *StatusCmd {
	cmd := c.client.B().Ping().Message(message).Build()
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

func (c *Compat) Migrate(ctx context.Context, host string, port int64, key bool, db int64, timeout time.Duration) *StatusCmd {
	var cmd cmds.Completed
	if key {
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

func (c *Compat) ObjectIdleTime(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().ObjectIdletime().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
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

func (c *Compat) PTTL(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Pttl().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
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

func (c *Compat) Sort(ctx context.Context, key string, sort Sort) *StringSliceCmd {
	cmd := c.client.B().Arbitrary("SORT").Keys(key)
	if sort.By != "" {
		cmd = cmd.Args("BY", sort.By)
	}
	if sort.Offset != 0 || sort.Count != 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(sort.Offset, 10), strconv.FormatInt(sort.Count, 10))
	}
	if len(sort.Get) > 0 {
		cmd = cmd.Args("GET")
		cmd = cmd.Args(sort.Get...)
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
	resp := c.client.Do(ctx, cmd.Build())
	return newStringSliceCmd(resp)
}

func (c *Compat) SortStore(ctx context.Context, key, store string, sort Sort) *IntCmd {
	cmd := c.client.B().Arbitrary("SORT").Keys(key)
	if sort.By != "" {
		cmd = cmd.Args("BY", sort.By)
	}
	if sort.Offset != 0 || sort.Count != 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(sort.Offset, 10), strconv.FormatInt(sort.Count, 10))
	}
	if len(sort.Get) > 0 {
		cmd = cmd.Args("GET")
		cmd = cmd.Args(sort.Get...)
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
	resp := c.client.Do(ctx, cmd.Args("STORE", store).Build())
	return newIntCmd(resp)
}

func (c *Compat) Touch(ctx context.Context, keys ...string) *IntCmd {
	cmd := c.client.B().Touch().Key(keys...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) TTL(ctx context.Context, key string) *IntCmd {
	cmd := c.client.B().Ttl().Key(key).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
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

func (c *Compat) GetSet(ctx context.Context, key, value string) *StringCmd {
	cmd := c.client.B().Getset().Key(key).Value(value).Build()
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

func (c *Compat) MSet(ctx context.Context, keys []string, values []string) *StatusCmd {
	if len(keys) != len(values) {
		panic(fmt.Sprintf("keys and values must be same length %d != %d", len(keys), len(values)))
	}
	partial := c.client.B().Mset().KeyValue()
	for i := 0; i < len(keys); i++ {
		partial = partial.KeyValue(keys[i], values[i])
	}
	cmd := partial.Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) MSetNX(ctx context.Context, keys []string, values []string) *BoolCmd {
	if len(keys) != len(values) {
		panic(fmt.Sprintf("keys and values must be same length %d != %d", len(keys), len(values)))
	}
	partial := c.client.B().Msetnx().KeyValue()
	for i := 0; i < len(keys); i++ {
		partial = partial.KeyValue(keys[i], values[i])
	}
	cmd := partial.Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

// SET key value [expiration]
//
// For no expiration use 0.
//
// For KEEPTTL use -1.
//
// For more options, use SetArgs.
func (c *Compat) Set(ctx context.Context, key string, value string, expiration time.Duration) *StatusCmd {
	var resp rueidis.RedisResult
	if expiration > 0 {
		if usePrecise(expiration) {
			resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(value).PxMilliseconds(formatMs(expiration)).Build())
		} else {
			resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(value).ExSeconds(formatSec(expiration)).Build())
		}
	} else if expiration == -1 {
		resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(value).Keepttl().Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(value).Build())
	}
	return newStatusCmd(resp)
}

func (c *Compat) SetArgs(ctx context.Context, key string, value string, a SetArgs) *StatusCmd {
	cmd := c.client.B().Arbitrary("SET").Keys(key).Args(value)
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
	default:
		panic(fmt.Sprintf("invalid mode for SET: %s", a.Mode))
	}
	if a.Get {
		cmd = cmd.Args("GET")
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newStatusCmd(resp)
}

func (c *Compat) SetEx(ctx context.Context, key string, value string, expiration time.Duration) *StatusCmd {
	cmd := c.client.B().Setex().Key(key).Seconds(formatSec(expiration)).Value(value).Build()
	resp := c.client.Do(ctx, cmd)
	return newStatusCmd(resp)
}

func (c *Compat) SetNX(ctx context.Context, key string, value string) *BoolCmd {
	cmd := c.client.B().Setnx().Key(key).Value(value).Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) SetXX(ctx context.Context, key string, value string, expiration time.Duration) *BoolCmd {
	var resp rueidis.RedisResult
	if expiration > 0 {
		if usePrecise(expiration) {
			resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(value).PxMilliseconds(formatMs(expiration)).Build())
		} else {
			resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(value).ExSeconds(formatSec(expiration)).Build())
		}
	} else if expiration == -1 {
		resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(value).Keepttl().Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Set().Key(key).Value(value).Build())
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

func (c *Compat) BitCount(ctx context.Context, key string, bitCount BitCount) *IntCmd {
	cmd := c.client.B().Bitcount().Key(key).Start(bitCount.Start).End(bitCount.End).Build()
	resp := c.client.Do(ctx, cmd)
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

func (c *Compat) BitPos(ctx context.Context, key string, bit int64, bitPos BitPos) *IntCmd {
	var resp rueidis.RedisResult
	if bitPos.Byte {
		resp = c.client.Do(ctx, c.client.B().Bitpos().Key(key).Bit(bit).Start(bitPos.Start).End(bitPos.End).Byte().Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Bitpos().Key(key).Bit(bit).Start(bitPos.Start).End(bitPos.End).Bit().Build())
	}
	return newIntCmd(resp)
}

func (c *Compat) BitField(ctx context.Context, key string, bitField []BitField) *IntSliceCmd {
	cmd := c.client.B().Arbitrary("BITFIELD").Keys(key)
	for _, a := range bitField {
		if a.Get != nil {
			cmd = cmd.Args("GET", a.Get.Encoding, strconv.FormatInt(a.Get.Offset, 10))
		}
		if a.Set != nil {
			cmd = cmd.Args("SET", a.Set.Encoding, strconv.FormatInt(a.Set.Offset, 10))
		}
		if a.IncrBy != nil {
			cmd = cmd.Args("INCRBY", a.IncrBy.Encoding, strconv.FormatInt(a.IncrBy.Offset, 10), strconv.FormatInt(a.Increment, 10))
		}
		switch overflow := strings.ToUpper(a.Overflow); overflow {
		case "WRAP", "SAT", "FAIL":
			cmd = cmd.Args("OVERFLOW", overflow)
		default:
			panic(fmt.Sprintf("Invalid OVERFLOW argument value: %s", a.Overflow))
		}
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntSliceCmd(resp)
}

func (c *Compat) Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd {
	cmd := c.client.B().Scan().Cursor(int64(cursor)).Match(match).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newScanCmd(resp)
}

func (c *Compat) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd {
	cmd := c.client.B().Scan().Cursor(int64(cursor)).Match(match).Count(count).Type(keyType).Build()
	resp := c.client.Do(ctx, cmd)
	return newScanCmd(resp)
}

func (c *Compat) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	cmd := c.client.B().Sscan().Key(key).Cursor(int64(cursor)).Match(match).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newScanCmd(resp)
}

func (c *Compat) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	cmd := c.client.B().Hscan().Key(key).Cursor(int64(cursor)).Match(match).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
	return newScanCmd(resp)
}

func (c *Compat) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	cmd := c.client.B().Zscan().Key(key).Cursor(int64(cursor)).Match(match).Count(count).Build()
	resp := c.client.Do(ctx, cmd)
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
func (c *Compat) HSet(ctx context.Context, key string, keys []string, values []string) *IntCmd {
	if len(keys) != len(values) {
		panic(fmt.Sprintf("keys and values must be same length %d != %d", len(keys), len(values)))
	}
	partial := c.client.B().Hset().Key(key).FieldValue()
	for i := 0; i < len(keys); i++ {
		partial = partial.FieldValue(keys[i], values[i])
	}
	cmd := partial.Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

// HMSet is a deprecated version of HSet left for compatibility with Redis 3.
func (c *Compat) HMSet(ctx context.Context, key string, keys []string, values []string) *BoolCmd {
	if len(keys) != len(values) {
		panic(fmt.Sprintf("keys and values must be same length %d != %d", len(keys), len(values)))
	}
	partial := c.client.B().Hset().Key(key).FieldValue()
	for i := 0; i < len(keys); i++ {
		partial = partial.FieldValue(keys[i], values[i])
	}
	cmd := partial.Build()
	resp := c.client.Do(ctx, cmd)
	return newBoolCmd(resp)
}

func (c *Compat) HSetNX(ctx context.Context, key, field string, value string) *BoolCmd {
	cmd := c.client.B().Hsetnx().Key(key).Field(field).Value(value).Build()
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
		resp = c.client.Do(ctx, c.client.B().Hrandfield().Key(key).Count(count).Build())
	} else {
		resp = c.client.Do(ctx, c.client.B().Hrandfield().Key(key).Count(count).Withvalues().Build())
	}
	return newStringSliceCmd(resp)
}

func (c *Compat) BLPop(ctx context.Context, timeout time.Duration, key string, keys ...string) *StringSliceCmd {
	cmd := c.client.B().Blpop().Key(key).Key(keys...).Timeout(float64(formatSec(timeout))).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) BRPop(ctx context.Context, timeout time.Duration, key string, keys ...string) *StringSliceCmd {
	cmd := c.client.B().Brpop().Key(key).Key(keys...).Timeout(float64(formatSec(timeout))).Build()
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

func (c *Compat) LInsert(ctx context.Context, key, op, pivot, element string) *IntCmd {
	var resp rueidis.RedisResult
	switch strings.ToUpper(op) {
	case "BEFORE":
		resp = c.client.Do(ctx, c.client.B().Linsert().Key(key).Before().Pivot(pivot).Element(element).Build())
	case "AFTER":
		resp = c.client.Do(ctx, c.client.B().Linsert().Key(key).After().Pivot(pivot).Element(element).Build())
	default:
		panic(fmt.Sprintf("Invalid op argument value: %s", op))
	}
	return newIntCmd(resp)
}

func (c *Compat) LInsertBefore(ctx context.Context, key, pivot, element string) *IntCmd {
	cmd := c.client.B().Linsert().Key(key).Before().Pivot(pivot).Element(element).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LInsertAfter(ctx context.Context, key, pivot, element string) *IntCmd {
	cmd := c.client.B().Linsert().Key(key).After().Pivot(pivot).Element(element).Build()
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
	cmd := c.client.B().Arbitrary("LPOS").Keys(key).Args("ELEMENT", element)
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
	cmd := c.client.B().Arbitrary("LPOS").Keys(key).Args("ELEMENT", element).Args("COUNT", strconv.FormatInt(count, 10))
	if a.Rank != 0 {
		cmd = cmd.Args("RANK", strconv.FormatInt(a.Rank, 10))
	}
	if a.MaxLen != 0 {
		cmd = cmd.Args("MAXLEN", strconv.FormatInt(a.MaxLen, 10))
	}
	resp := c.client.Do(ctx, cmd.Build())
	return newIntSliceCmd(resp)
}

func (c *Compat) LPush(ctx context.Context, key string, elements ...string) *IntCmd {
	cmd := c.client.B().Lpush().Key(key).Element(elements...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LPushX(ctx context.Context, key string, elements ...string) *IntCmd {
	cmd := c.client.B().Lpushx().Key(key).Element(elements...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	cmd := c.client.B().Lrange().Key(key).Start(start).Stop(stop).Build()
	resp := c.client.Do(ctx, cmd)
	return newStringSliceCmd(resp)
}

func (c *Compat) LRem(ctx context.Context, key string, count int64, element string) *IntCmd {
	cmd := c.client.B().Lrem().Key(key).Count(count).Element(element).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) LSet(ctx context.Context, key string, index int64, element string) *StatusCmd {
	cmd := c.client.B().Lset().Key(key).Index(index).Element(element).Build()
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

func (c *Compat) RPush(ctx context.Context, key string, elements ...string) *IntCmd {
	cmd := c.client.B().Rpush().Key(key).Element(elements...).Build()
	resp := c.client.Do(ctx, cmd)
	return newIntCmd(resp)
}

func (c *Compat) RPushX(ctx context.Context, key string, elements ...string) *IntCmd {
	cmd := c.client.B().Rpushx().Key(key).Element(elements...).Build()
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
