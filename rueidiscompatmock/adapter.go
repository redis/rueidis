package rueidiscompatmock

import (
	"context"
	"encoding"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/cmds"
	"github.com/redis/rueidis/mock"
	"github.com/redis/rueidis/rueidiscompat"
	"go.uber.org/mock/gomock"
)

type Client struct {
	rueidiscompat.Cmdable
	raw  rueidis.Client
	mock *clientMock
}

type ClientMock interface {
	Regexp() *clientMock
	CustomMatch(fn CustomMatch) *clientMock
	MatchExpectationsInOrder(b bool)
	ClearExpect()
	ExpectationsWereMet() error
	ExpectGet(key string) *ExpectedString
	ExpectSet(key string, value any, expiration time.Duration) *ExpectedStatus
	ExpectSetNX(key string, value any, expiration time.Duration) *ExpectedBool
	ExpectGetSet(key string, value any) *ExpectedString
	ExpectAppend(key, value string) *ExpectedInt
	ExpectStrLen(key string) *ExpectedInt
	ExpectDel(keys ...string) *ExpectedInt
	ExpectExists(keys ...string) *ExpectedInt
	ExpectType(key string) *ExpectedStatus
	ExpectTTL(key string) *ExpectedDuration
	ExpectExpire(key string, expiration time.Duration) *ExpectedBool
	ExpectPing() *ExpectedStatus
	ExpectEcho(message any) *ExpectedString
	ExpectIncr(key string) *ExpectedInt
	ExpectIncrBy(key string, value int64) *ExpectedInt
	ExpectDecr(key string) *ExpectedInt
	ExpectDecrBy(key string, value int64) *ExpectedInt
	ExpectMGet(keys ...string) *ExpectedSlice
	ExpectMSet(values ...any) *ExpectedStatus
	ExpectHGet(key, field string) *ExpectedString
	ExpectHSet(key string, values ...any) *ExpectedInt
	ExpectHDel(key string, fields ...string) *ExpectedInt
	ExpectHGetAll(key string) *ExpectedMapStringString
	ExpectLPush(key string, elements ...any) *ExpectedInt
	ExpectRPush(key string, elements ...any) *ExpectedInt
	ExpectLPop(key string) *ExpectedString
	ExpectRPop(key string) *ExpectedString
	ExpectLLen(key string) *ExpectedInt
	ExpectSAdd(key string, members ...any) *ExpectedInt
	ExpectSRem(key string, members ...any) *ExpectedInt
	ExpectSMembers(key string) *ExpectedStringSlice
	ExpectEval(script string, keys []string, args ...any) *ExpectedCmd
	ExpectDo(args ...any) *ExpectedCmd
	ExpectGetEx(key string, expiration time.Duration) *ExpectedString
	ExpectGetDel(key string) *ExpectedString
	ExpectGetRange(key string, start, end int64) *ExpectedString
	ExpectSetEx(key string, value any, expiration time.Duration) *ExpectedStatus
	ExpectSetXX(key string, value any, expiration time.Duration) *ExpectedBool
	ExpectSetRange(key string, offset int64, value string) *ExpectedInt
	ExpectMSetNX(values ...any) *ExpectedBool
	ExpectIncrByFloat(key string, increment float64) *ExpectedFloat
	ExpectGetBit(key string, offset int64) *ExpectedInt
	ExpectSetBit(key string, offset int64, value int) *ExpectedInt
	ExpectBitCount(key string, bitCount *rueidiscompat.BitCount) *ExpectedInt
	ExpectBitOpAnd(destKey string, keys ...string) *ExpectedInt
	ExpectBitOpOr(destKey string, keys ...string) *ExpectedInt
	ExpectBitOpXor(destKey string, keys ...string) *ExpectedInt
	ExpectBitOpNot(destKey, key string) *ExpectedInt
	ExpectUnlink(keys ...string) *ExpectedInt
	ExpectExpireAt(key string, tm time.Time) *ExpectedBool
	ExpectExpireNX(key string, expiration time.Duration) *ExpectedBool
	ExpectExpireXX(key string, expiration time.Duration) *ExpectedBool
	ExpectExpireGT(key string, expiration time.Duration) *ExpectedBool
	ExpectExpireLT(key string, expiration time.Duration) *ExpectedBool
	ExpectExpireTime(key string) *ExpectedDuration
	ExpectPExpire(key string, expiration time.Duration) *ExpectedBool
	ExpectPExpireAt(key string, tm time.Time) *ExpectedBool
	ExpectPExpireTime(key string) *ExpectedDuration
	ExpectPTTL(key string) *ExpectedDuration
	ExpectPersist(key string) *ExpectedBool
	ExpectRandomKey() *ExpectedString
	ExpectRename(key, newkey string) *ExpectedStatus
	ExpectRenameNX(key, newkey string) *ExpectedBool
	ExpectMove(key string, db int) *ExpectedBool
	ExpectObjectEncoding(key string) *ExpectedString
	ExpectObjectIdleTime(key string) *ExpectedDuration
	ExpectObjectRefCount(key string) *ExpectedInt
	ExpectTouch(keys ...string) *ExpectedInt
	ExpectCopy(source, destination string, db int, replace bool) *ExpectedInt
	ExpectHExists(key, field string) *ExpectedBool
	ExpectHKeys(key string) *ExpectedStringSlice
	ExpectHLen(key string) *ExpectedInt
	ExpectHMGet(key string, fields ...string) *ExpectedSlice
	ExpectHMSet(key string, values ...any) *ExpectedBool
	ExpectHSetNX(key, field string, value any) *ExpectedBool
	ExpectHVals(key string) *ExpectedStringSlice
	ExpectHIncrBy(key, field string, incr int64) *ExpectedInt
	ExpectHIncrByFloat(key, field string, incr float64) *ExpectedFloat
	ExpectHRandField(key string, count int) *ExpectedStringSlice
	ExpectLIndex(key string, index int64) *ExpectedString
	ExpectLRange(key string, start, stop int64) *ExpectedStringSlice
	ExpectLRem(key string, count int64, value any) *ExpectedInt
	ExpectLSet(key string, index int64, value any) *ExpectedStatus
	ExpectLTrim(key string, start, stop int64) *ExpectedStatus
	ExpectLInsertBefore(key string, pivot, value any) *ExpectedInt
	ExpectLInsertAfter(key string, pivot, value any) *ExpectedInt
	ExpectLPushX(key string, values ...any) *ExpectedInt
	ExpectRPushX(key string, values ...any) *ExpectedInt
	ExpectRPopLPush(source, destination string) *ExpectedString
	ExpectBLPop(timeout time.Duration, keys ...string) *ExpectedStringSlice
	ExpectBRPop(timeout time.Duration, keys ...string) *ExpectedStringSlice
	ExpectSCard(key string) *ExpectedInt
	ExpectSDiff(keys ...string) *ExpectedStringSlice
	ExpectSDiffStore(destination string, keys ...string) *ExpectedInt
	ExpectSInter(keys ...string) *ExpectedStringSlice
	ExpectSInterStore(destination string, keys ...string) *ExpectedInt
	ExpectSIsMember(key string, member any) *ExpectedBool
	ExpectSMIsMember(key string, members ...any) *ExpectedBoolSlice
	ExpectSMove(source, destination string, member any) *ExpectedBool
	ExpectSPop(key string) *ExpectedString
	ExpectSPopN(key string, count int64) *ExpectedStringSlice
	ExpectSRandMember(key string) *ExpectedString
	ExpectSRandMemberN(key string, count int64) *ExpectedStringSlice
	ExpectSUnion(keys ...string) *ExpectedStringSlice
	ExpectSUnionStore(destination string, keys ...string) *ExpectedInt
	ExpectZCard(key string) *ExpectedInt
	ExpectZCount(key, min, max string) *ExpectedInt
	ExpectZIncrBy(key string, increment float64, member string) *ExpectedFloat
	ExpectZLexCount(key, min, max string) *ExpectedInt
	ExpectZRange(key string, start, stop int64) *ExpectedStringSlice
	ExpectZRevRange(key string, start, stop int64) *ExpectedStringSlice
	ExpectZRank(key, member string) *ExpectedInt
	ExpectZRevRank(key, member string) *ExpectedInt
	ExpectZRem(key string, members ...any) *ExpectedInt
	ExpectZRemRangeByLex(key, min, max string) *ExpectedInt
	ExpectZRemRangeByRank(key string, start, stop int64) *ExpectedInt
	ExpectZRemRangeByScore(key, min, max string) *ExpectedInt
	ExpectZScore(key, member string) *ExpectedFloat
	ExpectZMScore(key string, members ...string) *ExpectedFloatSlice
	ExpectZPopMax(key string, count ...int64) *ExpectedZSlice
	ExpectZPopMin(key string, count ...int64) *ExpectedZSlice
	ExpectPublish(channel string, message any) *ExpectedInt
	ExpectSPublish(channel string, message any) *ExpectedInt
	ExpectPubSubChannels(pattern string) *ExpectedStringSlice
	ExpectPubSubNumPat() *ExpectedInt
	ExpectPFAdd(key string, els ...any) *ExpectedInt
	ExpectPFCount(keys ...string) *ExpectedInt
	ExpectPFMerge(dest string, keys ...string) *ExpectedStatus
	ExpectLastSave() *ExpectedInt
	ExpectTime() *ExpectedTime
	ExpectInfo(sections ...string) *ExpectedString
	ExpectClientID() *ExpectedInt
	ExpectClientGetName() *ExpectedString
	ExpectReadOnly() *ExpectedStatus
	ExpectReadWrite() *ExpectedStatus
	ExpectEvalSha(sha string, keys []string, args ...any) *ExpectedCmd
	ExpectEvalRO(script string, keys []string, args ...any) *ExpectedCmd
	ExpectEvalShaRO(sha string, keys []string, args ...any) *ExpectedCmd
	ExpectClusterInfo() *ExpectedString
	ExpectClusterNodes() *ExpectedString
	ExpectClusterKeySlot(key string) *ExpectedInt
	ExpectBgRewriteAOF() *ExpectedStatus
	ExpectBgSave() *ExpectedStatus
	ExpectConfigGet(parameter string) *ExpectedMapStringString
	ExpectConfigResetStat() *ExpectedStatus
	ExpectConfigRewrite() *ExpectedStatus
	ExpectConfigSet(parameter, value string) *ExpectedStatus
	ExpectDBSize() *ExpectedInt
	ExpectDebugObject(key string) *ExpectedString
	ExpectFlushAll() *ExpectedStatus
	ExpectFlushAllAsync() *ExpectedStatus
	ExpectFlushDB() *ExpectedStatus
	ExpectFlushDBAsync() *ExpectedStatus
	ExpectMemoryUsage(key string, samples ...int) *ExpectedInt
	ExpectSave() *ExpectedStatus
	ExpectShutdown() *ExpectedStatus
	ExpectShutdownNoSave() *ExpectedStatus
	ExpectShutdownSave() *ExpectedStatus
	ExpectSlaveOf(host, port string) *ExpectedStatus
	ExpectSlowLogGet(num int64) *ExpectedSlowLog
	ExpectQuit() *ExpectedStatus
	ExpectCommand() *ExpectedCommandsInfo
	ExpectCommandGetKeys(commands ...any) *ExpectedStringSlice
	ExpectCommandGetKeysAndFlags(commands ...any) *ExpectedKeyFlags
	ExpectCommandList(filter any) *ExpectedStringSlice
	ExpectClientKill(ipPort string) *ExpectedStatus
	ExpectClientKillByFilter(keys ...string) *ExpectedInt
	ExpectClientList() *ExpectedString
	ExpectClientPause(dur time.Duration) *ExpectedBool
	ExpectClientUnpause() *ExpectedBool
	ExpectClientUnblock(id int64) *ExpectedInt
	ExpectClientUnblockWithError(id int64) *ExpectedInt
	ExpectClusterAddSlots(slots ...int) *ExpectedStatus
	ExpectClusterAddSlotsRange(min, max int) *ExpectedStatus
	ExpectClusterCountFailureReports(nodeID string) *ExpectedInt
	ExpectClusterCountKeysInSlot(slot int) *ExpectedInt
	ExpectClusterDelSlots(slots ...int) *ExpectedStatus
	ExpectClusterDelSlotsRange(min, max int) *ExpectedStatus
	ExpectClusterFailover() *ExpectedStatus
	ExpectClusterForget(nodeID string) *ExpectedStatus
	ExpectClusterGetKeysInSlot(slot int, count int) *ExpectedStringSlice
	ExpectClusterLinks() *ExpectedClusterLinks
	ExpectClusterMeet(host, port string) *ExpectedStatus
	ExpectClusterReplicate(nodeID string) *ExpectedStatus
	ExpectClusterResetHard() *ExpectedStatus
	ExpectClusterResetSoft() *ExpectedStatus
	ExpectClusterSaveConfig() *ExpectedStatus
	ExpectClusterShards() *ExpectedClusterShards
	ExpectClusterSlaves(nodeID string) *ExpectedStringSlice
	ExpectClusterSlots() *ExpectedClusterSlots
	ExpectACLDryRun(username string, command ...any) *ExpectedString
	ExpectBitField(key string, args ...any) *ExpectedIntSlice
	ExpectBitPos(key string, bit int64, pos ...int64) *ExpectedInt
	ExpectBitPosSpan(key string, bit int8, start, end int64, span string) *ExpectedInt
	ExpectDump(key string) *ExpectedString
	ExpectKeys(pattern string) *ExpectedStringSlice
	ExpectMigrate(host, port, key string, db int, timeout time.Duration) *ExpectedStatus
	ExpectRestore(key string, ttl time.Duration, serializedValue string) *ExpectedStatus
	ExpectRestoreReplace(key string, ttl time.Duration, serializedValue string) *ExpectedStatus
	ExpectScan(cursor uint64, match string, count int64) *ExpectedScan
	ExpectScanType(cursor uint64, match string, count int64, keyType string) *ExpectedScan
	ExpectHScan(key string, cursor uint64, match string, count int64) *ExpectedScan
	ExpectSScan(key string, cursor uint64, match string, count int64) *ExpectedScan
	ExpectSort(key string, sort *rueidiscompat.Sort) *ExpectedStringSlice
	ExpectSortInterfaces(key string, sort *rueidiscompat.Sort) *ExpectedSlice
	ExpectSortRO(key string, sort *rueidiscompat.Sort) *ExpectedStringSlice
	ExpectSortStore(key, store string, sort *rueidiscompat.Sort) *ExpectedInt
	ExpectBLMPop(timeout time.Duration, direction string, count int64, keys ...string) *ExpectedKeyValues
	ExpectBLMove(source, destination, srcpos, destpos string, timeout time.Duration) *ExpectedString
	ExpectBRPopLPush(source, destination string, timeout time.Duration) *ExpectedString
	ExpectHRandFieldWithValues(key string, count int) *ExpectedKeyValueSlice
	ExpectLCS(q *rueidiscompat.LCSQuery) *ExpectedLCS
	ExpectLInsert(key, op string, pivot, element any) *ExpectedInt
	ExpectLMPop(direction string, count int64, keys ...string) *ExpectedKeyValues
	ExpectLMove(source, destination, srcpos, destpos string) *ExpectedString
	ExpectLPopCount(key string, count int) *ExpectedStringSlice
	ExpectLPos(key string, element string, a rueidiscompat.LPosArgs) *ExpectedInt
	ExpectLPosCount(key string, element string, count int64, a rueidiscompat.LPosArgs) *ExpectedIntSlice
	ExpectRPopCount(key string, count int) *ExpectedStringSlice
	ExpectSInterCard(limit int64, keys ...string) *ExpectedInt
	ExpectSMembersMap(key string) *ExpectedStringStructMap
	ExpectSetArgs(key string, value any, a rueidiscompat.SetArgs) *ExpectedStatus
	ExpectGeoAdd(key string, geoLocation ...any) *ExpectedInt
	ExpectGeoDist(key, member1, member2, unit string) *ExpectedFloat
	ExpectGeoHash(key string, members ...string) *ExpectedStringSlice
	ExpectGeoPos(key string, members ...string) *ExpectedGeoPos
	ExpectGeoRadius(key string, longitude, latitude float64, query *rueidiscompat.GeoRadiusQuery) *ExpectedGeoLocation
	ExpectGeoRadiusByMember(key, member string, query *rueidiscompat.GeoRadiusQuery) *ExpectedGeoLocation
	ExpectGeoRadiusByMemberStore(key, member string, query *rueidiscompat.GeoRadiusQuery) *ExpectedInt
	ExpectGeoRadiusStore(key string, longitude, latitude float64, query *rueidiscompat.GeoRadiusQuery) *ExpectedInt
	ExpectGeoSearch(key string, q *rueidiscompat.GeoSearchQuery) *ExpectedStringSlice
	ExpectGeoSearchLocation(key string, q *rueidiscompat.GeoSearchLocationQuery) *ExpectedGeoSearchLocation
	ExpectGeoSearchStore(src, dest string, q *rueidiscompat.GeoSearchStoreQuery) *ExpectedInt
	ExpectBZMPop(timeout time.Duration, order string, count int64, keys ...string) *ExpectedZSliceWithKey
	ExpectBZPopMax(timeout time.Duration, keys ...string) *ExpectedZWithKey
	ExpectBZPopMin(timeout time.Duration, keys ...string) *ExpectedZWithKey
	ExpectZDiff(keys ...string) *ExpectedStringSlice
	ExpectZDiffStore(destination string, keys ...string) *ExpectedInt
	ExpectZDiffWithScores(keys ...string) *ExpectedZSlice
	ExpectZInterCard(limit int64, keys ...string) *ExpectedInt
	ExpectZMPop(order string, count int64, keys ...string) *ExpectedZSliceWithKey
	ExpectZRandMember(key string, count int) *ExpectedStringSlice
	ExpectZRandMemberWithScores(key string, count int) *ExpectedZSlice
	ExpectZRangeByLex(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedStringSlice
	ExpectZRangeByScore(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedStringSlice
	ExpectZRangeByScoreWithScores(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedZSlice
	ExpectZRangeStore(dst string, z rueidiscompat.ZRangeArgs) *ExpectedInt
	ExpectZRevRangeByLex(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedStringSlice
	ExpectZRevRangeByScore(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedStringSlice
	ExpectZRevRangeByScoreWithScores(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedZSlice
	ExpectZRevRangeWithScores(key string, start, stop int64) *ExpectedZSlice
	ExpectZScan(key string, cursor uint64, match string, count int64) *ExpectedScan
	ExpectZAdd(key string, members ...rueidiscompat.Z) *ExpectedInt
	ExpectZAddNX(key string, members ...rueidiscompat.Z) *ExpectedInt
	ExpectZAddXX(key string, members ...rueidiscompat.Z) *ExpectedInt
	ExpectZAddLT(key string, members ...rueidiscompat.Z) *ExpectedInt
	ExpectZAddGT(key string, members ...rueidiscompat.Z) *ExpectedInt
	ExpectZAddArgs(key string, args rueidiscompat.ZAddArgs) *ExpectedInt
	ExpectZAddArgsIncr(key string, args rueidiscompat.ZAddArgs) *ExpectedFloat
	ExpectZInter(store *rueidiscompat.ZStore) *ExpectedStringSlice
	ExpectZInterWithScores(store *rueidiscompat.ZStore) *ExpectedZSlice
	ExpectZInterStore(destination string, store *rueidiscompat.ZStore) *ExpectedInt
	ExpectZUnion(store rueidiscompat.ZStore) *ExpectedStringSlice
	ExpectZUnionWithScores(store rueidiscompat.ZStore) *ExpectedZSlice
	ExpectZUnionStore(dest string, store *rueidiscompat.ZStore) *ExpectedInt
	ExpectZRangeArgs(z rueidiscompat.ZRangeArgs) *ExpectedStringSlice
	ExpectZRangeArgsWithScores(z rueidiscompat.ZRangeArgs) *ExpectedZSlice
	ExpectZRangeWithScores(key string, start, stop int64) *ExpectedZSlice
	ExpectPubSubNumSub(channels ...string) *ExpectedMapStringInt
	ExpectPubSubShardChannels(pattern string) *ExpectedStringSlice
	ExpectPubSubShardNumSub(channels ...string) *ExpectedMapStringInt
	ExpectScriptExists(hashes ...string) *ExpectedBoolSlice
	ExpectScriptFlush() *ExpectedStatus
	ExpectScriptKill() *ExpectedStatus
	ExpectScriptLoad(script string) *ExpectedString
	ExpectFCall(function string, keys []string, args ...any) *ExpectedCmd
	ExpectFCallRo(function string, keys []string, args ...any) *ExpectedCmd
	ExpectFunctionDelete(libName string) *ExpectedString
	ExpectFunctionDump() *ExpectedString
	ExpectFunctionFlush() *ExpectedString
	ExpectFunctionFlushAsync() *ExpectedString
	ExpectFunctionKill() *ExpectedString
	ExpectFunctionList(q rueidiscompat.FunctionListQuery) *ExpectedFunctionList
	ExpectFunctionLoad(code string) *ExpectedString
	ExpectFunctionLoadReplace(code string) *ExpectedString
	ExpectFunctionRestore(libDump string) *ExpectedString
	ExpectXAck(stream, group string, ids ...string) *ExpectedInt
	ExpectXAdd(args any) *ExpectedString
	ExpectXAutoClaim(args any) *ExpectedXAutoClaim
	ExpectXAutoClaimJustID(args any) *ExpectedXAutoClaimJustID
	ExpectXClaim(args any) *ExpectedXMessageSlice
	ExpectXClaimJustID(args any) *ExpectedStringSlice
	ExpectXDel(stream string, ids ...string) *ExpectedInt
	ExpectXGroupCreate(stream, group, start string) *ExpectedStatus
	ExpectXGroupCreateConsumer(stream, group, consumer string) *ExpectedInt
	ExpectXGroupCreateMkStream(stream, group, start string) *ExpectedStatus
	ExpectXGroupDelConsumer(stream, group, consumer string) *ExpectedInt
	ExpectXGroupDestroy(stream, group string) *ExpectedInt
	ExpectXGroupSetID(stream, group, start string) *ExpectedStatus
	ExpectXInfoConsumers(key, group string) *ExpectedXInfoConsumers
	ExpectXInfoGroups(key string) *ExpectedXInfoGroups
	ExpectXInfoStream(key string) *ExpectedXInfoStream
	ExpectXInfoStreamFull(key string, count int) *ExpectedXInfoStreamFull
	ExpectXLen(stream string) *ExpectedInt
	ExpectXPending(stream, group string) *ExpectedXPending
	ExpectXPendingExt(args any) *ExpectedXPendingExt
	ExpectXRange(stream, start, stop string) *ExpectedXMessageSlice
	ExpectXRangeN(stream, start, stop string, count int64) *ExpectedXMessageSlice
	ExpectXRead(args any) *ExpectedXStreamSlice
	ExpectXReadGroup(args any) *ExpectedXStreamSlice
	ExpectXReadStreams(streams ...string) *ExpectedXStreamSlice
	ExpectXRevRange(stream, start, stop string) *ExpectedXMessageSlice
	ExpectXRevRangeN(stream, start, stop string, count int64) *ExpectedXMessageSlice
	ExpectXTrimMaxLen(key string, maxLen int64) *ExpectedInt
	ExpectXTrimMaxLenApprox(key string, maxLen, limit int64) *ExpectedInt
	ExpectXTrimMinID(key string, minID string) *ExpectedInt
	ExpectXTrimMinIDApprox(key string, minID string, limit int64) *ExpectedInt
	ExpectTxPipeline()
	ExpectTxPipelineExec() *ExpectedSlice
	ExpectWatch(keys ...string) *ExpectedError
}

type clientMock struct {
	raw *mock.Client

	parent *clientMock

	mu        sync.Mutex
	queue     []*expectation
	unmatched []error
	ordered   bool

	expectRegexp bool
	expectCustom CustomMatch
}

type expectation struct {
	matcher     gomock.Matcher
	expected    []string
	result      rueidis.RedisResult
	rawResult   any
	err         error
	resultSet   bool
	rawSet      bool
	redisNil    bool
	regexpMatch bool
	customMatch CustomMatch
}

type resolvedExpectation struct {
	result rueidis.RedisResult
	raw    any
	rawSet bool
}

type CustomMatch func(expected, actual []interface{}) error

type commandMatcher struct {
	gomock.Matcher
	expected []string
}

func (m commandMatcher) expectedCommands() []string {
	return append([]string(nil), m.expected...)
}

func match(cmd ...string) gomock.Matcher {
	return commandMatcher{Matcher: mock.Match(cmd...), expected: append([]string(nil), cmd...)}
}

func matchFn(expected []string, fn func(cmd []string) bool, description ...string) gomock.Matcher {
	return commandMatcher{Matcher: mock.MatchFn(fn, description...), expected: append([]string(nil), expected...)}
}

func matcherCommands(matcher gomock.Matcher) []string {
	withExpected, ok := matcher.(interface{ expectedCommands() []string })
	if !ok {
		return nil
	}
	return withExpected.expectedCommands()
}

func commandInterfaces(cmd []string) []interface{} {
	out := make([]interface{}, len(cmd))
	for i := range cmd {
		out[i] = cmd[i]
	}
	return out
}

func NewAdapter(m *mock.Client) ClientMock {
	cm := newAdapter(m)
	return cm
}

func newAdapter(m *mock.Client) *clientMock {
	cm := &clientMock{raw: m, ordered: true}
	cm.wire()
	return cm
}

func NewClientMock() (*Client, ClientMock) {
	ctrl := gomock.NewController(panicReporter{})
	raw := mock.NewClient(ctrl)
	raw.EXPECT().Close().AnyTimes()
	raw.EXPECT().Nodes().Return(map[string]rueidis.Client{"mock": raw}).AnyTimes()
	cm := newAdapter(raw)
	return &Client{Cmdable: rueidiscompat.NewAdapter(raw), raw: raw, mock: cm}, cm
}

func (c *Client) Close() error {
	c.raw.Close()
	return nil
}

func (c *Client) Do(ctx context.Context, args ...any) *rueidiscompat.Cmd {
	ret := &rueidiscompat.Cmd{}
	if len(args) == 0 {
		ret.SetErr(errors.New("redis: please enter the command to be executed"))
		return ret
	}
	cmd := c.raw.B().Arbitrary(str(args[0]))
	if len(args) > 1 {
		cmd = cmd.Keys(str(args[1]))
		for _, a := range args[2:] {
			cmd = cmd.Args(str(a))
		}
	}
	built := cmd.Build()
	resolved := resolvedExpectation{}
	if c.mock != nil {
		resolved = c.mock.consumeWithRaw(built)
	} else {
		resolved.result = c.raw.Do(ctx, built)
	}
	res := resolved.result
	val, err := res.ToAny()
	if err != nil {
		ret.SetErr(err)
		return ret
	}
	ret.SetVal(val)
	if resolved.rawSet {
		ret.SetRawVal(resolved.raw)
	} else {
		ret.SetRawVal(val)
	}
	return ret
}

type panicReporter struct{}

func (panicReporter) Errorf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

func (panicReporter) Fatalf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

func (m *clientMock) root() *clientMock {
	if m.parent != nil {
		return m.parent.root()
	}
	return m
}

func (m *clientMock) Regexp() *clientMock {
	root := m.root()
	return &clientMock{raw: root.raw, parent: root, expectRegexp: true}
}

func (m *clientMock) CustomMatch(fn CustomMatch) *clientMock {
	root := m.root()
	return &clientMock{raw: root.raw, parent: root, expectCustom: fn}
}

func (m *clientMock) MatchExpectationsInOrder(b bool) {
	root := m.root()
	root.mu.Lock()
	defer root.mu.Unlock()
	root.ordered = b
}

func (m *clientMock) ClearExpect() {
	root := m.root()
	root.mu.Lock()
	defer root.mu.Unlock()
	root.queue = nil
	root.unmatched = nil
}

func (m *clientMock) ExpectationsWereMet() error {
	root := m.root()
	root.mu.Lock()
	defer root.mu.Unlock()
	if len(root.unmatched) > 0 {
		return root.unmatched[0]
	}
	if len(root.queue) > 0 {
		cmds := make([]string, 0, len(root.queue))
		for _, e := range root.queue {
			cmds = append(cmds, e.matcher.String())
		}
		return fmt.Errorf("rueidiscompatmock: there are remaining expectations: %v", cmds)
	}
	return nil
}

func (e *expectation) setErr(err error) {
	e.err = err
}

func (e *expectation) setRedisNil() {
	e.redisNil = true
}

func (m *clientMock) wire() {
	m.raw.EXPECT().
		Do(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, cmd rueidis.Completed) rueidis.RedisResult {
			return m.consume(1, []rueidis.Completed{cmd})[0]
		}).
		AnyTimes()
	m.raw.EXPECT().
		DoMulti(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, cmds ...rueidis.Completed) []rueidis.RedisResult {
			return m.consume(len(cmds), cmds)
		}).
		AnyTimes()
}

func (m *clientMock) consume(n int, cmds []rueidis.Completed) []rueidis.RedisResult {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]rueidis.RedisResult, n)
	for i, cmd := range cmds {
		idx, ok := m.matchLocked(cmd)
		if !ok {
			err := fmt.Errorf("rueidiscompatmock: no expectation for command %v", cmd.Commands())
			m.unmatched = append(m.unmatched, err)
			out[i] = mock.ErrorResult(err)
			for j := i + 1; j < n; j++ {
				out[j] = mock.ErrorResult(err)
			}
			break
		}
		out[i] = m.queue[idx].resolve(cmd)
		m.queue = append(m.queue[:idx], m.queue[idx+1:]...)
	}
	return out
}

func (m *clientMock) consumeWithRaw(cmd rueidis.Completed) resolvedExpectation {
	m.mu.Lock()
	defer m.mu.Unlock()
	idx, ok := m.matchLocked(cmd)
	if !ok {
		err := fmt.Errorf("rueidiscompatmock: no expectation for command %v", cmd.Commands())
		m.unmatched = append(m.unmatched, err)
		return resolvedExpectation{result: mock.ErrorResult(err)}
	}
	e := m.queue[idx]
	resolved := resolvedExpectation{
		result: e.resolve(cmd),
		raw:    e.rawResult,
		rawSet: e.rawSet && e.err == nil && !e.redisNil && e.resultSet,
	}
	m.queue = append(m.queue[:idx], m.queue[idx+1:]...)
	return resolved
}

func (m *clientMock) matchLocked(cmd rueidis.Completed) (int, bool) {
	if m.ordered {
		if len(m.queue) == 0 {
			return 0, false
		}
		if !m.queue[0].matches(cmd) {
			return 0, false
		}
		return 0, true
	}
	for i, e := range m.queue {
		if e.matches(cmd) {
			return i, true
		}
	}
	return 0, false
}

func (e *expectation) matches(cmd rueidis.Completed) bool {
	if e.customMatch == nil && !e.regexpMatch {
		return e.matcher.Matches(cmd)
	}
	if len(e.expected) == 0 {
		return e.matcher.Matches(cmd)
	}
	expected := stringsToAny(e.expected)
	actual := stringsToAny(cmd.Commands())
	if len(expected) != len(actual) || expected[0] != actual[0] {
		return false
	}
	if e.customMatch != nil {
		return e.customMatch(expected, actual) == nil
	}
	if mapArgs(expected[0], &expected) {
		mapArgs(actual[0], &actual)
	}
	for i := range expected {
		if !compareArg(e.regexpMatch, expected[i], actual[i]) {
			return false
		}
	}
	return true
}

func (e *expectation) resolve(cmd rueidis.Completed) rueidis.RedisResult {
	if e.err != nil {
		return mock.ErrorResult(e.err)
	}
	if e.redisNil {
		return mock.Result(mock.RedisNil())
	}
	if !e.resultSet {
		name := ""
		if parts := cmd.Commands(); len(parts) > 0 {
			name = strings.ToLower(parts[0])
		}
		return mock.ErrorResult(fmt.Errorf("cmd(%s), return value is required", name))
	}
	return e.result
}

func (m *clientMock) push(matcher gomock.Matcher, defaultResult rueidis.RedisResult) *expectation {
	return m.pushWithOptions(matcher, defaultResult, false)
}

func (m *clientMock) pushSet(matcher gomock.Matcher, result rueidis.RedisResult) *expectation {
	return m.pushWithOptions(matcher, result, true)
}

func (m *clientMock) pushWithOptions(matcher gomock.Matcher, defaultResult rueidis.RedisResult, resultSet bool) *expectation {
	root := m.root()
	root.mu.Lock()
	defer root.mu.Unlock()
	e := &expectation{
		matcher:     matcher,
		expected:    matcherCommands(matcher),
		result:      defaultResult,
		resultSet:   resultSet,
		regexpMatch: m.expectRegexp,
		customMatch: m.expectCustom,
	}
	root.queue = append(root.queue, e)
	return e
}

func stringsToAny(values []string) []any {
	out := make([]any, len(values))
	for i, value := range values {
		out[i] = value
	}
	return out
}

func compareArg(isRegexp bool, expected, actual any) bool {
	if expectedMap, ok := expected.(map[string]any); ok {
		actualMap, ok := actual.(map[string]any)
		if !ok || len(expectedMap) != len(actualMap) {
			return false
		}
		for k, expectedValue := range expectedMap {
			actualValue, ok := actualMap[k]
			if !ok || !compareArg(isRegexp, expectedValue, actualValue) {
				return false
			}
		}
		return true
	}
	if isRegexp {
		if expr, ok := expected.(string); ok {
			re, err := regexp.Compile(expr)
			return err == nil && re.MatchString(fmt.Sprint(actual))
		}
	}
	return reflect.DeepEqual(expected, actual)
}

func mapArgs(cmd any, args *[]any) bool {
	if len(*args) == 0 {
		return false
	}
	cut := 0
	switch strings.ToLower(fmt.Sprint(cmd)) {
	case "mset", "msetnx":
		cut = 1
	case "hset", "hmset":
		cut = 2
	default:
		return false
	}
	if n := len(*args); n <= cut || (n > cut+1 && (n-cut)%2 != 0) {
		return false
	}
	mapArgs := make(map[string]any)
	rest := (*args)[cut:]
	switch v := rest[0].(type) {
	case []string:
		if len(v)%2 != 0 {
			return false
		}
		for i := 0; i < len(v); i += 2 {
			mapArgs[v[i]] = v[i+1]
		}
	case map[string]any:
		if len(v) > 0 {
			mapArgs = v
		}
	default:
		for i := 0; i < len(rest); i += 2 {
			mapArgs[fmt.Sprint(rest[i])] = rest[i+1]
		}
	}
	if len(mapArgs) == 0 {
		return false
	}
	next := make([]any, cut, cut+1)
	copy(next, (*args)[:cut])
	next = append(next, mapArgs)
	*args = next
	return true
}

const keepTTL = -1

func usePrecise(d time.Duration) bool {
	return d < time.Second || d%time.Second != 0
}

func formatMs(d time.Duration) int64 { return int64(d / time.Millisecond) }

func formatSec(d time.Duration) int64 { return int64(d / time.Second) }

func str(arg any) string {
	if arg == nil {
		return ""
	}
	switch v := arg.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case bool:
		if v {
			return "1"
		}
		return "0"
	case time.Time:
		return v.Format(time.RFC3339Nano)
	case time.Duration:
		return strconv.FormatInt(v.Nanoseconds(), 10)
	case encoding.BinaryMarshaler:
		if data, err := v.MarshalBinary(); err == nil {
			return rueidis.BinaryString(data)
		}
	}
	return fmt.Sprint(arg)
}

func (m *clientMock) ExpectGet(key string) *ExpectedString {
	e := m.push(match("GET", key), mock.Result(mock.RedisNil()))
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectSet(key string, value any, expiration time.Duration) *ExpectedStatus {
	e := m.push(setMatcher(key, value, expiration), mock.Result(mock.RedisString("OK")))
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectSetNX(key string, value any, expiration time.Duration) *ExpectedBool {
	e := m.push(setNXMatcher(key, value, expiration), mock.Result(mock.RedisBool(true)))
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectGetSet(key string, value any) *ExpectedString {
	e := m.push(match("GETSET", key, str(value)), mock.Result(mock.RedisNil()))
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectAppend(key, value string) *ExpectedInt {
	e := m.push(match("APPEND", key, value), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectStrLen(key string) *ExpectedInt {
	e := m.push(match("STRLEN", key), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectDel(keys ...string) *ExpectedInt {
	e := m.push(delMatcher(keys...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectExists(keys ...string) *ExpectedInt {
	args := append([]string{"EXISTS"}, keys...)
	e := m.push(match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectType(key string) *ExpectedStatus {
	e := m.push(match("TYPE", key), mock.Result(mock.RedisString("none")))
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectTTL(key string) *ExpectedDuration {
	e := m.push(match("TTL", key), mock.Result(mock.RedisInt64(-2)))
	return &ExpectedDuration{exp: e, precision: time.Second}
}

func (m *clientMock) ExpectExpire(key string, expiration time.Duration) *ExpectedBool {
	e := m.push(match("EXPIRE", key, strconv.FormatInt(formatSec(expiration), 10)), mock.Result(mock.RedisBool(false)))
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectPing() *ExpectedStatus {
	e := m.push(match("PING"), mock.Result(mock.RedisString("PONG")))
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectEcho(message any) *ExpectedString {
	e := m.push(match("ECHO", str(message)), mock.Result(mock.RedisString("")))
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectIncr(key string) *ExpectedInt {
	e := m.push(match("INCR", key), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectIncrBy(key string, value int64) *ExpectedInt {
	e := m.push(match("INCRBY", key, strconv.FormatInt(value, 10)), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectDecr(key string) *ExpectedInt {
	e := m.push(match("DECR", key), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectDecrBy(key string, value int64) *ExpectedInt {
	e := m.push(match("DECRBY", key, strconv.FormatInt(value, 10)), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectMGet(keys ...string) *ExpectedSlice {
	args := append([]string{"MGET"}, keys...)
	e := m.push(match(args...), mock.Result(mock.RedisArray()))
	return &ExpectedSlice{exp: e}
}

func (m *clientMock) ExpectMSet(values ...any) *ExpectedStatus {
	pairs := argsToSlice(values)
	e := m.push(pairsMatcher("MSET", "", pairs), mock.Result(mock.RedisString("OK")))
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectHGet(key, field string) *ExpectedString {
	e := m.push(match("HGET", key, field), mock.Result(mock.RedisNil()))
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectHSet(key string, values ...any) *ExpectedInt {
	pairs := argsToSlice(values)
	e := m.push(pairsMatcher("HSET", key, pairs), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectHDel(key string, fields ...string) *ExpectedInt {
	args := append([]string{"HDEL", key}, fields...)
	e := m.push(match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectHGetAll(key string) *ExpectedMapStringString {
	e := m.push(match("HGETALL", key), mock.Result(mock.RedisMap(map[string]rueidis.RedisMessage{})))
	return &ExpectedStringStringMap{exp: e}
}

func (m *clientMock) ExpectLPush(key string, elements ...any) *ExpectedInt {
	args := []string{"LPUSH", key}
	for _, el := range elements {
		args = append(args, str(el))
	}
	e := m.push(match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectRPush(key string, elements ...any) *ExpectedInt {
	args := []string{"RPUSH", key}
	for _, el := range elements {
		args = append(args, str(el))
	}
	e := m.push(match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectLPop(key string) *ExpectedString {
	e := m.push(match("LPOP", key), mock.Result(mock.RedisNil()))
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectRPop(key string) *ExpectedString {
	e := m.push(match("RPOP", key), mock.Result(mock.RedisNil()))
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectLLen(key string) *ExpectedInt {
	e := m.push(match("LLEN", key), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectSAdd(key string, members ...any) *ExpectedInt {
	args := []string{"SADD", key}
	for _, mm := range members {
		args = append(args, str(mm))
	}
	e := m.push(match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectSRem(key string, members ...any) *ExpectedInt {
	args := []string{"SREM", key}
	for _, mm := range members {
		args = append(args, str(mm))
	}
	e := m.push(match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectSMembers(key string) *ExpectedStringSlice {
	e := m.push(match("SMEMBERS", key), mock.Result(mock.RedisArray()))
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectEval(script string, keys []string, args ...any) *ExpectedCmd {
	e := m.push(evalMatcher(script, keys, args...), mock.Result(mock.RedisNil()))
	return &ExpectedCmd{exp: e}
}

func setMatcher(key string, value any, expiration time.Duration) gomock.Matcher {
	if expiration > 0 {
		if usePrecise(expiration) {
			return match("SET", key, str(value), "PX", strconv.FormatInt(formatMs(expiration), 10))
		}
		return match("SET", key, str(value), "EX", strconv.FormatInt(formatSec(expiration), 10))
	}
	if expiration == keepTTL {
		return match("SET", key, str(value), "KEEPTTL")
	}
	return match("SET", key, str(value))
}

func setNXMatcher(key string, value any, expiration time.Duration) gomock.Matcher {
	switch expiration {
	case 0:
		return match("SETNX", key, str(value))
	case keepTTL:
		return match("SET", key, str(value), "NX", "KEEPTTL")
	}
	if usePrecise(expiration) {
		return match("SET", key, str(value), "NX", "PX", strconv.FormatInt(formatMs(expiration), 10))
	}
	return match("SET", key, str(value), "NX", "EX", strconv.FormatInt(formatSec(expiration), 10))
}

func delMatcher(keys ...string) gomock.Matcher {
	args := append([]string{"DEL"}, keys...)
	return match(args...)
}

func argsToSlice(src []any) []string {
	if len(src) == 1 {
		return argToSlice(src[0])
	}
	dst := make([]string, 0, len(src))
	for _, v := range src {
		dst = append(dst, str(v))
	}
	return dst
}

func argToSlice(arg any) []string {
	switch arg := arg.(type) {
	case []string:
		return arg
	case []any:
		dst := make([]string, 0, len(arg))
		for _, v := range arg {
			dst = append(dst, str(v))
		}
		return dst
	case map[string]any:
		dst := make([]string, 0, len(arg)*2)
		for k, v := range arg {
			dst = append(dst, k, str(v))
		}
		return dst
	case map[string]string:
		dst := make([]string, 0, len(arg)*2)
		for k, v := range arg {
			dst = append(dst, k, v)
		}
		return dst
	default:
		v := reflect.ValueOf(arg)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return nil
			}
			v = v.Elem()
		}
		if v.Kind() == reflect.Struct {
			return appendStructField(v)
		}
		return []string{str(arg)}
	}
}

func appendStructField(v reflect.Value) []string {
	typ := v.Type()
	dst := make([]string, 0, typ.NumField()*2)
	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("redis")
		if tag == "" || tag == "-" {
			continue
		}
		name, opt, _ := strings.Cut(tag, ",")
		if name == "" {
			continue
		}
		field := v.Field(i)
		if omitEmpty(opt) && isEmptyValue(field) {
			continue
		}
		if field.Kind() == reflect.Pointer && field.IsNil() {
			continue
		}
		if field.Kind() == reflect.Pointer && field.Elem().CanInterface() {
			dst = append(dst, name, str(field.Elem().Interface()))
			continue
		}
		if field.CanInterface() {
			dst = append(dst, name, str(field.Interface()))
		}
	}
	return dst
}

func omitEmpty(opt string) bool {
	for opt != "" {
		var name string
		name, opt, _ = strings.Cut(opt, ",")
		if name == "omitempty" {
			return true
		}
	}
	return false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return v.IsNil()
	}
	return false
}

func pairsMatcher(cmd, key string, pairs []string) gomock.Matcher {
	prefixLen := 1
	if key != "" {
		prefixLen = 2
	}
	wantSorted := sortedPairs(pairs)
	desc := []string{cmd}
	if key != "" {
		desc = append(desc, key)
	}
	desc = append(desc, pairs...)
	return matchFn(desc, func(got []string) bool {
		if len(got) != prefixLen+len(pairs) {
			return false
		}
		if got[0] != cmd {
			return false
		}
		if key != "" && got[1] != key {
			return false
		}
		return equalSorted(wantSorted, sortedPairs(got[prefixLen:]))
	}, desc...)
}

func sortedPairs(flat []string) []string {
	if len(flat)%2 != 0 {
		out := append([]string(nil), flat...)
		return out
	}
	pairs := make([][2]string, 0, len(flat)/2)
	for i := 0; i < len(flat); i += 2 {
		pairs = append(pairs, [2]string{flat[i], flat[i+1]})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
	out := make([]string, 0, len(flat))
	for _, p := range pairs {
		out = append(out, p[0], p[1])
	}
	return out
}

func equalSorted(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func evalMatcher(script string, keys []string, args ...any) gomock.Matcher {
	parts := []string{"EVAL", script, strconv.Itoa(len(keys))}
	parts = append(parts, keys...)
	for _, a := range args {
		parts = append(parts, str(a))
	}
	return match(parts...)
}

func defaultStringResult() rueidis.RedisResult { return mock.Result(mock.RedisNil()) }
func defaultStatusResult() rueidis.RedisResult { return mock.Result(mock.RedisString("OK")) }
func defaultIntResult() rueidis.RedisResult    { return mock.Result(mock.RedisInt64(0)) }
func defaultBoolResult() rueidis.RedisResult   { return mock.Result(mock.RedisBool(false)) }
func defaultFloatResult() rueidis.RedisResult  { return mock.Result(mock.RedisFloat64(0)) }
func defaultArrayResult() rueidis.RedisResult  { return mock.Result(mock.RedisArray()) }
func defaultMapResult() rueidis.RedisResult {
	return mock.Result(mock.RedisMap(map[string]rueidis.RedisMessage{}))
}

func (m *clientMock) ExpectDo(args ...any) *ExpectedCmd {
	e := m.push(match(argsToSlice(args)...), defaultStringResult())
	return &ExpectedCmd{exp: e}
}

func (m *clientMock) ExpectGetEx(key string, expiration time.Duration) *ExpectedString {
	var cmd rueidis.Completed
	if expiration > 0 {
		if usePrecise(expiration) {
			cmd = m.raw.B().Getex().Key(key).PxMilliseconds(formatMs(expiration)).Build()
		} else {
			cmd = m.raw.B().Getex().Key(key).ExSeconds(formatSec(expiration)).Build()
		}
	} else {
		cmd = m.raw.B().Getex().Key(key).Build()
	}
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectGetDel(key string) *ExpectedString {
	cmd := m.raw.B().Getdel().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectGetRange(key string, start, end int64) *ExpectedString {
	cmd := m.raw.B().Getrange().Key(key).Start(start).End(end).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectSetEx(key string, value any, expiration time.Duration) *ExpectedStatus {
	cmd := m.raw.B().Setex().Key(key).Seconds(formatSec(expiration)).Value(str(value)).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectSetXX(key string, value any, expiration time.Duration) *ExpectedBool {
	var cmd rueidis.Completed
	if expiration > 0 {
		if usePrecise(expiration) {
			cmd = m.raw.B().Set().Key(key).Value(str(value)).Xx().PxMilliseconds(formatMs(expiration)).Build()
		} else {
			cmd = m.raw.B().Set().Key(key).Value(str(value)).Xx().ExSeconds(formatSec(expiration)).Build()
		}
	} else if expiration == keepTTL {
		cmd = m.raw.B().Set().Key(key).Value(str(value)).Xx().Keepttl().Build()
	} else {
		cmd = m.raw.B().Set().Key(key).Value(str(value)).Xx().Build()
	}
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectSetRange(key string, offset int64, value string) *ExpectedInt {
	cmd := m.raw.B().Setrange().Key(key).Offset(offset).Value(value).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectMSetNX(values ...any) *ExpectedBool {
	pairs := argsToSlice(values)
	args := append([]string{"MSETNX"}, pairs...)
	e := m.push(pairsMatcher("MSETNX", "", pairs), defaultBoolResult())
	_ = args
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectIncrByFloat(key string, increment float64) *ExpectedFloat {
	cmd := m.raw.B().Incrbyfloat().Key(key).Increment(increment).Build()
	e := m.push(match(cmd.Commands()...), defaultFloatResult())
	return &ExpectedFloat{exp: e}
}

func (m *clientMock) ExpectGetBit(key string, offset int64) *ExpectedInt {
	cmd := m.raw.B().Getbit().Key(key).Offset(offset).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectSetBit(key string, offset int64, value int) *ExpectedInt {
	cmd := m.raw.B().Setbit().Key(key).Offset(offset).Value(int64(value)).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectBitCount(key string, bitCount *rueidiscompat.BitCount) *ExpectedInt {
	var cmd rueidis.Completed
	if bitCount != nil {
		cmd = m.raw.B().Bitcount().Key(key).Start(bitCount.Start).End(bitCount.End).Build()
	} else {
		cmd = m.raw.B().Bitcount().Key(key).Build()
	}
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectBitOpAnd(destKey string, keys ...string) *ExpectedInt {
	cmd := m.raw.B().Bitop().And().Destkey(destKey).Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectBitOpOr(destKey string, keys ...string) *ExpectedInt {
	cmd := m.raw.B().Bitop().Or().Destkey(destKey).Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectBitOpXor(destKey string, keys ...string) *ExpectedInt {
	cmd := m.raw.B().Bitop().Xor().Destkey(destKey).Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectBitOpNot(destKey, key string) *ExpectedInt {
	cmd := m.raw.B().Bitop().Not().Destkey(destKey).Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectUnlink(keys ...string) *ExpectedInt {
	cmd := m.raw.B().Unlink().Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectExpireAt(key string, tm time.Time) *ExpectedBool {
	cmd := m.raw.B().Expireat().Key(key).Timestamp(tm.Unix()).Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectExpireNX(key string, expiration time.Duration) *ExpectedBool {
	cmd := m.raw.B().Expire().Key(key).Seconds(formatSec(expiration)).Nx().Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectExpireXX(key string, expiration time.Duration) *ExpectedBool {
	cmd := m.raw.B().Expire().Key(key).Seconds(formatSec(expiration)).Xx().Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectExpireGT(key string, expiration time.Duration) *ExpectedBool {
	cmd := m.raw.B().Expire().Key(key).Seconds(formatSec(expiration)).Gt().Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectExpireLT(key string, expiration time.Duration) *ExpectedBool {
	cmd := m.raw.B().Expire().Key(key).Seconds(formatSec(expiration)).Lt().Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectExpireTime(key string) *ExpectedDuration {
	cmd := m.raw.B().Expiretime().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedDuration{exp: e, precision: time.Second}
}

func (m *clientMock) ExpectPExpire(key string, expiration time.Duration) *ExpectedBool {
	cmd := m.raw.B().Pexpire().Key(key).Milliseconds(formatMs(expiration)).Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectPExpireAt(key string, tm time.Time) *ExpectedBool {
	cmd := m.raw.B().Pexpireat().Key(key).MillisecondsTimestamp(tm.UnixMilli()).Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectPExpireTime(key string) *ExpectedDuration {
	cmd := m.raw.B().Pexpiretime().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedDuration{exp: e, precision: time.Millisecond}
}

func (m *clientMock) ExpectPTTL(key string) *ExpectedDuration {
	cmd := m.raw.B().Pttl().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedDuration{exp: e, precision: time.Millisecond}
}

func (m *clientMock) ExpectPersist(key string) *ExpectedBool {
	cmd := m.raw.B().Persist().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectRandomKey() *ExpectedString {
	cmd := m.raw.B().Randomkey().Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectRename(key, newkey string) *ExpectedStatus {
	cmd := m.raw.B().Rename().Key(key).Newkey(newkey).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectRenameNX(key, newkey string) *ExpectedBool {
	cmd := m.raw.B().Renamenx().Key(key).Newkey(newkey).Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectMove(key string, db int) *ExpectedBool {
	cmd := m.raw.B().Move().Key(key).Db(int64(db)).Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectObjectEncoding(key string) *ExpectedString {
	cmd := m.raw.B().ObjectEncoding().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectObjectIdleTime(key string) *ExpectedDuration {
	cmd := m.raw.B().ObjectIdletime().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedDuration{exp: e, precision: time.Second}
}

func (m *clientMock) ExpectObjectRefCount(key string) *ExpectedInt {
	cmd := m.raw.B().ObjectRefcount().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectTouch(keys ...string) *ExpectedInt {
	cmd := m.raw.B().Touch().Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectCopy(source, destination string, db int, replace bool) *ExpectedInt {
	var cmd rueidis.Completed
	if replace {
		cmd = m.raw.B().Copy().Source(source).Destination(destination).Db(int64(db)).Replace().Build()
	} else {
		cmd = m.raw.B().Copy().Source(source).Destination(destination).Db(int64(db)).Build()
	}
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectHExists(key, field string) *ExpectedBool {
	cmd := m.raw.B().Hexists().Key(key).Field(field).Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectHKeys(key string) *ExpectedStringSlice {
	cmd := m.raw.B().Hkeys().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectHLen(key string) *ExpectedInt {
	cmd := m.raw.B().Hlen().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectHMGet(key string, fields ...string) *ExpectedSlice {
	cmd := m.raw.B().Hmget().Key(key).Field(fields...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedSlice{exp: e}
}

func (m *clientMock) ExpectHMSet(key string, values ...any) *ExpectedBool {
	pairs := argsToSlice(values)
	e := m.push(pairsMatcher("HMSET", key, pairs), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectHSetNX(key, field string, value any) *ExpectedBool {
	cmd := m.raw.B().Hsetnx().Key(key).Field(field).Value(str(value)).Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectHVals(key string) *ExpectedStringSlice {
	cmd := m.raw.B().Hvals().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectHIncrBy(key, field string, incr int64) *ExpectedInt {
	cmd := m.raw.B().Hincrby().Key(key).Field(field).Increment(incr).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectHIncrByFloat(key, field string, incr float64) *ExpectedFloat {
	cmd := m.raw.B().Hincrbyfloat().Key(key).Field(field).Increment(incr).Build()
	e := m.push(match(cmd.Commands()...), defaultFloatResult())
	return &ExpectedFloat{exp: e}
}

func (m *clientMock) ExpectHRandField(key string, count int) *ExpectedStringSlice {
	cmd := m.raw.B().Hrandfield().Key(key).Count(int64(count)).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectLIndex(key string, index int64) *ExpectedString {
	cmd := m.raw.B().Lindex().Key(key).Index(index).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectLRange(key string, start, stop int64) *ExpectedStringSlice {
	cmd := m.raw.B().Lrange().Key(key).Start(start).Stop(stop).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectLRem(key string, count int64, value any) *ExpectedInt {
	cmd := m.raw.B().Lrem().Key(key).Count(count).Element(str(value)).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectLSet(key string, index int64, value any) *ExpectedStatus {
	cmd := m.raw.B().Lset().Key(key).Index(index).Element(str(value)).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectLTrim(key string, start, stop int64) *ExpectedStatus {
	cmd := m.raw.B().Ltrim().Key(key).Start(start).Stop(stop).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectLInsertBefore(key string, pivot, value any) *ExpectedInt {
	cmd := m.raw.B().Linsert().Key(key).Before().Pivot(str(pivot)).Element(str(value)).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectLInsertAfter(key string, pivot, value any) *ExpectedInt {
	cmd := m.raw.B().Linsert().Key(key).After().Pivot(str(pivot)).Element(str(value)).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectLPushX(key string, values ...any) *ExpectedInt {
	args := make([]string, 0, len(values))
	for _, v := range values {
		args = append(args, str(v))
	}
	cmd := m.raw.B().Lpushx().Key(key).Element(args...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectRPushX(key string, values ...any) *ExpectedInt {
	args := make([]string, 0, len(values))
	for _, v := range values {
		args = append(args, str(v))
	}
	cmd := m.raw.B().Rpushx().Key(key).Element(args...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectRPopLPush(source, destination string) *ExpectedString {
	cmd := m.raw.B().Rpoplpush().Source(source).Destination(destination).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectBLPop(timeout time.Duration, keys ...string) *ExpectedStringSlice {
	cmd := m.raw.B().Blpop().Key(keys...).Timeout(float64(formatSec(timeout))).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectBRPop(timeout time.Duration, keys ...string) *ExpectedStringSlice {
	cmd := m.raw.B().Brpop().Key(keys...).Timeout(float64(formatSec(timeout))).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectSCard(key string) *ExpectedInt {
	cmd := m.raw.B().Scard().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectSDiff(keys ...string) *ExpectedStringSlice {
	cmd := m.raw.B().Sdiff().Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectSDiffStore(destination string, keys ...string) *ExpectedInt {
	cmd := m.raw.B().Sdiffstore().Destination(destination).Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectSInter(keys ...string) *ExpectedStringSlice {
	cmd := m.raw.B().Sinter().Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectSInterStore(destination string, keys ...string) *ExpectedInt {
	cmd := m.raw.B().Sinterstore().Destination(destination).Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectSIsMember(key string, member any) *ExpectedBool {
	cmd := m.raw.B().Sismember().Key(key).Member(str(member)).Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectSMIsMember(key string, members ...any) *ExpectedBoolSlice {
	args := make([]string, 0, len(members))
	for _, mm := range members {
		args = append(args, str(mm))
	}
	cmd := m.raw.B().Smismember().Key(key).Member(args...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedBoolSlice{exp: e}
}

func (m *clientMock) ExpectSMove(source, destination string, member any) *ExpectedBool {
	cmd := m.raw.B().Smove().Source(source).Destination(destination).Member(str(member)).Build()
	e := m.push(match(cmd.Commands()...), defaultBoolResult())
	return &ExpectedBool{exp: e}
}

func (m *clientMock) ExpectSPop(key string) *ExpectedString {
	cmd := m.raw.B().Spop().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectSPopN(key string, count int64) *ExpectedStringSlice {
	cmd := m.raw.B().Spop().Key(key).Count(count).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectSRandMember(key string) *ExpectedString {
	cmd := m.raw.B().Srandmember().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectSRandMemberN(key string, count int64) *ExpectedStringSlice {
	cmd := m.raw.B().Srandmember().Key(key).Count(count).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectSUnion(keys ...string) *ExpectedStringSlice {
	cmd := m.raw.B().Sunion().Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectSUnionStore(destination string, keys ...string) *ExpectedInt {
	cmd := m.raw.B().Sunionstore().Destination(destination).Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZCard(key string) *ExpectedInt {
	cmd := m.raw.B().Zcard().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZCount(key, min, max string) *ExpectedInt {
	cmd := m.raw.B().Zcount().Key(key).Min(min).Max(max).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZIncrBy(key string, increment float64, member string) *ExpectedFloat {
	cmd := m.raw.B().Zincrby().Key(key).Increment(increment).Member(member).Build()
	e := m.push(match(cmd.Commands()...), defaultFloatResult())
	return &ExpectedFloat{exp: e}
}

func (m *clientMock) ExpectZLexCount(key, min, max string) *ExpectedInt {
	cmd := m.raw.B().Zlexcount().Key(key).Min(min).Max(max).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZRange(key string, start, stop int64) *ExpectedStringSlice {
	cmd := m.raw.B().Zrange().Key(key).Min(itoa(start)).Max(itoa(stop)).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZRevRange(key string, start, stop int64) *ExpectedStringSlice {
	cmd := m.raw.B().Zrevrange().Key(key).Start(start).Stop(stop).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZRank(key, member string) *ExpectedInt {
	cmd := m.raw.B().Zrank().Key(key).Member(member).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZRevRank(key, member string) *ExpectedInt {
	cmd := m.raw.B().Zrevrank().Key(key).Member(member).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZRem(key string, members ...any) *ExpectedInt {
	args := make([]string, 0, len(members))
	for _, mm := range members {
		args = append(args, str(mm))
	}
	cmd := m.raw.B().Zrem().Key(key).Member(args...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZRemRangeByLex(key, min, max string) *ExpectedInt {
	cmd := m.raw.B().Zremrangebylex().Key(key).Min(min).Max(max).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZRemRangeByRank(key string, start, stop int64) *ExpectedInt {
	cmd := m.raw.B().Zremrangebyrank().Key(key).Start(start).Stop(stop).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZRemRangeByScore(key, min, max string) *ExpectedInt {
	cmd := m.raw.B().Zremrangebyscore().Key(key).Min(min).Max(max).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZScore(key, member string) *ExpectedFloat {
	cmd := m.raw.B().Zscore().Key(key).Member(member).Build()
	e := m.push(match(cmd.Commands()...), defaultFloatResult())
	return &ExpectedFloat{exp: e}
}

func (m *clientMock) ExpectZMScore(key string, members ...string) *ExpectedFloatSlice {
	cmd := m.raw.B().Zmscore().Key(key).Member(members...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedFloatSlice{exp: e}
}

func (m *clientMock) ExpectZPopMax(key string, count ...int64) *ExpectedZSlice {
	var cmd rueidis.Completed
	if len(count) > 0 {
		cmd = m.raw.B().Zpopmax().Key(key).Count(count[0]).Build()
	} else {
		cmd = m.raw.B().Zpopmax().Key(key).Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

func (m *clientMock) ExpectZPopMin(key string, count ...int64) *ExpectedZSlice {
	var cmd rueidis.Completed
	if len(count) > 0 {
		cmd = m.raw.B().Zpopmin().Key(key).Count(count[0]).Build()
	} else {
		cmd = m.raw.B().Zpopmin().Key(key).Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

func (m *clientMock) ExpectPublish(channel string, message any) *ExpectedInt {
	cmd := m.raw.B().Publish().Channel(channel).Message(str(message)).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectSPublish(channel string, message any) *ExpectedInt {
	cmd := m.raw.B().Spublish().Channel(channel).Message(str(message)).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectPubSubChannels(pattern string) *ExpectedStringSlice {
	cmd := m.raw.B().PubsubChannels().Pattern(pattern).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectPubSubNumPat() *ExpectedInt {
	cmd := m.raw.B().PubsubNumpat().Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectPFAdd(key string, els ...any) *ExpectedInt {
	args := make([]string, 0, len(els))
	for _, v := range els {
		args = append(args, str(v))
	}
	cmd := m.raw.B().Pfadd().Key(key).Element(args...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectPFCount(keys ...string) *ExpectedInt {
	cmd := m.raw.B().Pfcount().Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectPFMerge(dest string, keys ...string) *ExpectedStatus {
	cmd := m.raw.B().Pfmerge().Destkey(dest).Sourcekey(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectLastSave() *ExpectedInt {
	cmd := m.raw.B().Lastsave().Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectTime() *ExpectedTime {
	cmd := m.raw.B().Time().Build()
	e := m.push(match(cmd.Commands()...), mock.Result(mock.RedisArray(mock.RedisString("0"), mock.RedisString("0"))))
	return &ExpectedTime{exp: e}
}

func (m *clientMock) ExpectInfo(sections ...string) *ExpectedString {
	var cmd rueidis.Completed
	if len(sections) > 0 {
		cmd = m.raw.B().Info().Section(sections...).Build()
	} else {
		cmd = m.raw.B().Info().Build()
	}
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectClientID() *ExpectedInt {
	cmd := m.raw.B().ClientId().Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectClientGetName() *ExpectedString {
	cmd := m.raw.B().ClientGetname().Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectReadOnly() *ExpectedStatus {
	cmd := m.raw.B().Readonly().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectReadWrite() *ExpectedStatus {
	cmd := m.raw.B().Readwrite().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectEvalSha(sha string, keys []string, args ...any) *ExpectedCmd {
	strArgs := make([]string, 0, len(args))
	for _, a := range args {
		strArgs = append(strArgs, str(a))
	}
	cmd := m.raw.B().Evalsha().Sha1(sha).Numkeys(int64(len(keys))).Key(keys...).Arg(strArgs...).Build()
	e := m.push(match(cmd.Commands()...), mock.Result(mock.RedisNil()))
	return &ExpectedCmd{exp: e}
}

func (m *clientMock) ExpectEvalRO(script string, keys []string, args ...any) *ExpectedCmd {
	strArgs := make([]string, 0, len(args))
	for _, a := range args {
		strArgs = append(strArgs, str(a))
	}
	cmd := m.raw.B().EvalRo().Script(script).Numkeys(int64(len(keys))).Key(keys...).Arg(strArgs...).Build()
	e := m.push(match(cmd.Commands()...), mock.Result(mock.RedisNil()))
	return &ExpectedCmd{exp: e}
}

func (m *clientMock) ExpectEvalShaRO(sha string, keys []string, args ...any) *ExpectedCmd {
	strArgs := make([]string, 0, len(args))
	for _, a := range args {
		strArgs = append(strArgs, str(a))
	}
	cmd := m.raw.B().EvalshaRo().Sha1(sha).Numkeys(int64(len(keys))).Key(keys...).Arg(strArgs...).Build()
	e := m.push(match(cmd.Commands()...), mock.Result(mock.RedisNil()))
	return &ExpectedCmd{exp: e}
}

func (m *clientMock) ExpectClusterInfo() *ExpectedString {
	cmd := m.raw.B().ClusterInfo().Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectClusterNodes() *ExpectedString {
	cmd := m.raw.B().ClusterNodes().Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectClusterKeySlot(key string) *ExpectedInt {
	cmd := m.raw.B().ClusterKeyslot().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectBgRewriteAOF() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().Bgrewriteaof().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectBgSave() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().Bgsave().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectConfigGet(parameter string) *ExpectedMapStringString {
	cmd := m.raw.B().ConfigGet().Parameter(parameter).Build()
	e := m.push(match(cmd.Commands()...), defaultMapResult())
	return &ExpectedStringStringMap{exp: e}
}

func (m *clientMock) ExpectConfigResetStat() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().ConfigResetstat().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectConfigRewrite() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().ConfigRewrite().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectConfigSet(parameter, value string) *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().ConfigSet().ParameterValue().ParameterValue(parameter, value).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectDBSize() *ExpectedInt {
	m.expectRoleMaster()
	cmd := m.raw.B().Dbsize().Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectDebugObject(key string) *ExpectedString {
	cmd := m.raw.B().DebugObject().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectFlushAll() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().Flushall().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectFlushAllAsync() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().Flushall().Async().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectFlushDB() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().Flushdb().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectFlushDBAsync() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().Flushdb().Async().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectMemoryUsage(key string, samples ...int) *ExpectedInt {
	var cmd rueidis.Completed
	switch len(samples) {
	case 0:
		cmd = m.raw.B().MemoryUsage().Key(key).Build()
	case 1:
		cmd = m.raw.B().MemoryUsage().Key(key).Samples(int64(samples[0])).Build()
	default:
		panic("too many arguments")
	}
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectSave() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().Save().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectShutdown() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().Shutdown().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectShutdownNoSave() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().Shutdown().Nosave().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectShutdownSave() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().Shutdown().Save().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectSlaveOf(host, port string) *ExpectedStatus {
	cmd := m.raw.B().Arbitrary("SLAVEOF").Args(host, port).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectSlowLogGet(num int64) *ExpectedSlowLog {
	cmd := m.raw.B().SlowlogGet().Count(num).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedSlowLog{exp: e}
}

func (m *clientMock) ExpectQuit() *ExpectedStatus {
	cmd := m.raw.B().Quit().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectCommand() *ExpectedCommandsInfo {
	cmd := m.raw.B().Command().Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedCommandsInfo{exp: e}
}

func (m *clientMock) ExpectCommandGetKeys(commands ...any) *ExpectedStringSlice {
	cmd := m.raw.B().CommandGetkeys().Command(commands[0].(string)).Arg(argsToSlice(commands[1:])...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectCommandGetKeysAndFlags(commands ...any) *ExpectedKeyFlags {
	cmd := m.raw.B().CommandGetkeysandflags().Command(commands[0].(string)).Arg(argsToSlice(commands[1:])...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedKeyFlags{exp: e}
}

func (m *clientMock) ExpectCommandList(filter any) *ExpectedStringSlice {
	f := normalizeFilter(filter)
	var cmd rueidis.Completed
	if f.Module != "" {
		cmd = m.raw.B().CommandList().FilterbyModuleName(f.Module).Build()
	} else if f.Pattern != "" {
		cmd = m.raw.B().CommandList().FilterbyPatternPattern(f.Pattern).Build()
	} else if f.ACLCat != "" {
		cmd = m.raw.B().CommandList().FilterbyAclcatCategory(f.ACLCat).Build()
	} else {
		cmd = m.raw.B().CommandList().Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func normalizeFilter(filter any) rueidiscompat.FilterBy {
	switch f := filter.(type) {
	case rueidiscompat.FilterBy:
		return f
	case *rueidiscompat.FilterBy:
		if f != nil {
			return *f
		}
	}
	return rueidiscompat.FilterBy{}
}

func (m *clientMock) ExpectClientKill(ipPort string) *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().ClientKill().IpPort(ipPort).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClientKillByFilter(keys ...string) *ExpectedInt {
	m.expectRoleMaster()
	cmd := m.raw.B().Arbitrary("CLIENT", "KILL").Args(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectClientList() *ExpectedString {
	cmd := m.raw.B().ClientList().Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectClientPause(dur time.Duration) *ExpectedBool {
	m.expectRoleMaster()
	cmd := m.raw.B().ClientPause().Timeout(formatSec(dur)).Build()
	e := m.push(match(cmd.Commands()...), mock.Result(mock.RedisString("OK")))
	return &ExpectedBool{exp: e, encode: okBoolMessage}
}

func (m *clientMock) ExpectClientUnpause() *ExpectedBool {
	m.expectRoleMaster()
	cmd := m.raw.B().ClientUnpause().Build()
	e := m.push(match(cmd.Commands()...), mock.Result(mock.RedisString("OK")))
	return &ExpectedBool{exp: e, encode: okBoolMessage}
}

func (m *clientMock) ExpectClientUnblock(id int64) *ExpectedInt {
	cmd := m.raw.B().ClientUnblock().ClientId(id).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectClientUnblockWithError(id int64) *ExpectedInt {
	cmd := m.raw.B().ClientUnblock().ClientId(id).Error().Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func intsToInt64s(values []int) []int64 {
	out := make([]int64, len(values))
	for i, v := range values {
		out[i] = int64(v)
	}
	return out
}

func (m *clientMock) ExpectClusterAddSlots(slots ...int) *ExpectedStatus {
	cmd := m.raw.B().ClusterAddslots().Slot(intsToInt64s(slots)...).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterAddSlotsRange(min, max int) *ExpectedStatus {
	cmd := m.raw.B().ClusterAddslotsrange().StartSlotEndSlot().StartSlotEndSlot(int64(min), int64(max)).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterCountFailureReports(nodeID string) *ExpectedInt {
	cmd := m.raw.B().ClusterCountFailureReports().NodeId(nodeID).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectClusterCountKeysInSlot(slot int) *ExpectedInt {
	m.expectRoleMaster()
	cmd := m.raw.B().ClusterCountkeysinslot().Slot(int64(slot)).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectClusterDelSlots(slots ...int) *ExpectedStatus {
	cmd := m.raw.B().ClusterDelslots().Slot(intsToInt64s(slots)...).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterDelSlotsRange(min, max int) *ExpectedStatus {
	cmd := m.raw.B().ClusterDelslotsrange().StartSlotEndSlot().StartSlotEndSlot(int64(min), int64(max)).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterFailover() *ExpectedStatus {
	cmd := m.raw.B().ClusterFailover().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterForget(nodeID string) *ExpectedStatus {
	cmd := m.raw.B().ClusterForget().NodeId(nodeID).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterGetKeysInSlot(slot int, count int) *ExpectedStringSlice {
	m.expectRoleMaster()
	cmd := m.raw.B().ClusterGetkeysinslot().Slot(int64(slot)).Count(int64(count)).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectClusterLinks() *ExpectedClusterLinks {
	cmd := m.raw.B().ClusterLinks().Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedClusterLinks{exp: e}
}

func (m *clientMock) ExpectClusterMeet(host, port string) *ExpectedStatus {
	portNum, _ := strconv.ParseInt(port, 10, 64)
	cmd := m.raw.B().ClusterMeet().Ip(host).Port(portNum).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterReplicate(nodeID string) *ExpectedStatus {
	cmd := m.raw.B().ClusterReplicate().NodeId(nodeID).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterResetHard() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().ClusterReset().Hard().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterResetSoft() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().ClusterReset().Soft().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterSaveConfig() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().ClusterSaveconfig().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectClusterShards() *ExpectedClusterShards {
	cmd := m.raw.B().ClusterShards().Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedClusterShards{exp: e}
}

func (m *clientMock) ExpectClusterSlaves(nodeID string) *ExpectedStringSlice {
	cmd := m.raw.B().ClusterSlaves().NodeId(nodeID).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectClusterSlots() *ExpectedClusterSlots {
	cmd := m.raw.B().ClusterSlots().Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedClusterSlots{exp: e}
}

func (m *clientMock) ExpectACLDryRun(username string, command ...any) *ExpectedString {
	cmd := m.raw.B().AclDryrun().Username(username).Command(command[0].(string)).Arg(argsToSlice(command[1:])...).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectBitField(key string, args ...any) *ExpectedIntSlice {
	cmd := m.raw.B().Arbitrary("BITFIELD").Keys(key)
	for _, v := range args {
		cmd = cmd.Args(str(v))
	}
	completed := cmd.Build()
	e := m.push(match(completed.Commands()...), defaultArrayResult())
	return &ExpectedIntSlice{exp: e}
}

func (m *clientMock) ExpectBitPos(key string, bit int64, pos ...int64) *ExpectedInt {
	var cmd rueidis.Completed
	switch len(pos) {
	case 0:
		cmd = m.raw.B().Bitpos().Key(key).Bit(bit).Build()
	case 1:
		cmd = m.raw.B().Bitpos().Key(key).Bit(bit).Start(pos[0]).Build()
	case 2:
		cmd = m.raw.B().Bitpos().Key(key).Bit(bit).Start(pos[0]).End(pos[1]).Build()
	default:
		panic("too many arguments")
	}
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectBitPosSpan(key string, bit int8, start, end int64, span string) *ExpectedInt {
	var cmd rueidis.Completed
	if strings.ToLower(span) == "bit" {
		cmd = m.raw.B().Bitpos().Key(key).Bit(int64(bit)).Start(start).End(end).Bit().Build()
	} else {
		cmd = m.raw.B().Bitpos().Key(key).Bit(int64(bit)).Start(start).End(end).Byte().Build()
	}
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectDump(key string) *ExpectedString {
	cmd := m.raw.B().Dump().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectKeys(pattern string) *ExpectedStringSlice {
	m.expectRoleMaster()
	cmd := m.raw.B().Keys().Pattern(pattern).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectMigrate(host, port, key string, db int, timeout time.Duration) *ExpectedStatus {
	portNum, _ := strconv.ParseInt(port, 10, 64)
	cmd := m.raw.B().Migrate().Host(host).Port(portNum).Key(key).DestinationDb(int64(db)).Timeout(formatSec(timeout)).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectRestore(key string, ttl time.Duration, serializedValue string) *ExpectedStatus {
	cmd := m.raw.B().Restore().Key(key).Ttl(formatMs(ttl)).SerializedValue(serializedValue).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectRestoreReplace(key string, ttl time.Duration, serializedValue string) *ExpectedStatus {
	cmd := m.raw.B().Restore().Key(key).Ttl(formatMs(ttl)).SerializedValue(serializedValue).Replace().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectScan(cursor uint64, pattern string, count int64) *ExpectedScan {
	cmd := m.raw.B().Arbitrary("SCAN", strconv.FormatInt(int64(cursor), 10))
	if pattern != "" {
		cmd = cmd.Args("MATCH", pattern)
	}
	if count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(count, 10))
	}
	completed := cmd.ReadOnly()
	e := m.push(match(completed.Commands()...), defaultArrayResult())
	return &ExpectedScan{exp: e}
}

func (m *clientMock) ExpectScanType(cursor uint64, pattern string, count int64, keyType string) *ExpectedScan {
	cmd := m.raw.B().Arbitrary("SCAN", strconv.FormatInt(int64(cursor), 10))
	if pattern != "" {
		cmd = cmd.Args("MATCH", pattern)
	}
	if count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(count, 10))
	}
	completed := cmd.Args("TYPE", keyType).ReadOnly()
	e := m.push(match(completed.Commands()...), defaultArrayResult())
	return &ExpectedScan{exp: e}
}

func (m *clientMock) ExpectHScan(key string, cursor uint64, pattern string, count int64) *ExpectedScan {
	cmd := m.raw.B().Arbitrary("HSCAN").Keys(key).Args(strconv.FormatInt(int64(cursor), 10))
	if pattern != "" {
		cmd = cmd.Args("MATCH", pattern)
	}
	if count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(count, 10))
	}
	completed := cmd.ReadOnly()
	e := m.push(match(completed.Commands()...), defaultArrayResult())
	return &ExpectedScan{exp: e}
}

func (m *clientMock) ExpectSScan(key string, cursor uint64, pattern string, count int64) *ExpectedScan {
	cmd := m.raw.B().Arbitrary("SSCAN").Keys(key).Args(strconv.FormatInt(int64(cursor), 10))
	if pattern != "" {
		cmd = cmd.Args("MATCH", pattern)
	}
	if count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(count, 10))
	}
	completed := cmd.ReadOnly()
	e := m.push(match(completed.Commands()...), defaultArrayResult())
	return &ExpectedScan{exp: e}
}

func (m *clientMock) ExpectSort(key string, sort *rueidiscompat.Sort) *ExpectedStringSlice {
	cmd := sortCmd(m.raw, "SORT", key, normalizeSort(sort), "")
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectSortInterfaces(key string, sort *rueidiscompat.Sort) *ExpectedSlice {
	cmd := sortCmd(m.raw, "SORT", key, normalizeSort(sort), "")
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedSlice{exp: e}
}

func (m *clientMock) ExpectSortRO(key string, sort *rueidiscompat.Sort) *ExpectedStringSlice {
	cmd := sortCmd(m.raw, "SORT_RO", key, normalizeSort(sort), "")
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectSortStore(key, store string, sort *rueidiscompat.Sort) *ExpectedInt {
	cmd := sortCmd(m.raw, "SORT", key, normalizeSort(sort), store)
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func normalizeSort(sort any) rueidiscompat.Sort {
	switch s := sort.(type) {
	case rueidiscompat.Sort:
		return s
	case *rueidiscompat.Sort:
		if s != nil {
			return *s
		}
	}
	return rueidiscompat.Sort{}
}

func sortCmd(raw rueidis.Client, command, key string, sort rueidiscompat.Sort, store string) rueidis.Completed {
	cmd := raw.B().Arbitrary(command).Keys(key)
	if sort.By != "" {
		cmd = cmd.Args("BY", sort.By)
	}
	if sort.Offset != 0 || sort.Count != 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(sort.Offset, 10), strconv.FormatInt(sort.Count, 10))
	}
	for _, get := range sort.Get {
		cmd = cmd.Args("GET").Args(get)
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
	if store != "" {
		cmd = cmd.Args("STORE", store)
	}
	return cmd.Build()
}

func (m *clientMock) ExpectBLMPop(timeout time.Duration, direction string, count int64, keys ...string) *ExpectedKeyValues {
	cmdBuilder := m.raw.B().Arbitrary("BLMPOP", strconv.FormatInt(formatSec(timeout), 10), strconv.Itoa(len(keys))).Keys(keys...)
	cmdBuilder = cmdBuilder.Args(direction)
	if count > 0 {
		cmdBuilder = cmdBuilder.Args("COUNT", strconv.FormatInt(count, 10))
	}
	cmd := cmdBuilder.Blocking()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedKeyValues{exp: e}
}

func (m *clientMock) ExpectBLMove(source, destination, srcpos, destpos string, timeout time.Duration) *ExpectedString {
	cmd := m.raw.B().Arbitrary("BLMOVE").Keys(source, destination).Args(srcpos, destpos, strconv.FormatFloat(float64(formatSec(timeout)), 'f', -1, 64)).Blocking()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectBRPopLPush(source, destination string, timeout time.Duration) *ExpectedString {
	cmd := m.raw.B().Brpoplpush().Source(source).Destination(destination).Timeout(float64(formatSec(timeout))).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectHRandFieldWithValues(key string, count int) *ExpectedKeyValueSlice {
	cmd := m.raw.B().Hrandfield().Key(key).Count(int64(count)).Withvalues().Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedKeyValueSlice{exp: e}
}

func (m *clientMock) ExpectLCS(q *rueidiscompat.LCSQuery) *ExpectedLCS {
	var cmd rueidis.Completed
	var readType uint8
	var withMatchLen bool

	_cmd := cmds.Incomplete(m.raw.B().Lcs().Key1(q.Key1).Key2(q.Key2))

	if q.Len {
		readType = uint8(2)
		cmd = cmds.LcsKey2(_cmd).Len().Build()
	} else if q.Idx {
		readType = uint8(3)
		withMatchLen = q.WithMatchLen
		if q.MinMatchLen > 0 && q.WithMatchLen {
			cmd = cmds.LcsKey2(_cmd).Idx().Minmatchlen(int64(q.MinMatchLen)).Withmatchlen().Build()
		} else if q.MinMatchLen > 0 {
			cmd = cmds.LcsKey2(_cmd).Idx().Minmatchlen(int64(q.MinMatchLen)).Build()
		} else if q.WithMatchLen {
			cmd = cmds.LcsKey2(_cmd).Idx().Withmatchlen().Build()
		} else {
			cmd = cmds.LcsKey2(_cmd).Idx().Build()
		}
	} else {
		readType = uint8(1)
		cmd = cmds.LcsKey2(_cmd).Build()
	}

	var defaultResult rueidis.RedisResult
	switch readType {
	case 2:
		defaultResult = defaultIntResult()
	case 3:
		defaultResult = defaultMapResult()
	default:
		defaultResult = defaultStringResult()
	}

	e := m.push(match(cmd.Commands()...), defaultResult)
	return &ExpectedLCS{exp: e, readType: readType, withMatchLen: withMatchLen}
}

func (m *clientMock) ExpectLInsert(key, op string, pivot, element any) *ExpectedInt {
	var cmd rueidis.Completed
	switch strings.ToUpper(op) {
	case "BEFORE":
		cmd = m.raw.B().Linsert().Key(key).Before().Pivot(str(pivot)).Element(str(element)).Build()
	case "AFTER":
		cmd = m.raw.B().Linsert().Key(key).After().Pivot(str(pivot)).Element(str(element)).Build()
	default:
		panic(fmt.Sprintf("Invalid op argument value: %s", op))
	}
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectLMPop(direction string, count int64, keys ...string) *ExpectedKeyValues {
	cmdBuilder := m.raw.B().Arbitrary("LMPOP", strconv.Itoa(len(keys))).Keys(keys...)
	cmdBuilder = cmdBuilder.Args(direction)
	if count > 0 {
		cmdBuilder = cmdBuilder.Args("COUNT", strconv.FormatInt(count, 10))
	}
	cmd := cmdBuilder.Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedKeyValues{exp: e}
}

func (m *clientMock) ExpectLMove(source, destination, srcpos, destpos string) *ExpectedString {
	cmd := m.raw.B().Arbitrary("LMOVE").Keys(source, destination).Args(srcpos, destpos).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectLPopCount(key string, count int) *ExpectedStringSlice {
	cmd := m.raw.B().Lpop().Key(key).Count(int64(count)).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectLPos(key string, element string, a rueidiscompat.LPosArgs) *ExpectedInt {
	cmdBuilder := m.raw.B().Arbitrary("LPOS").Keys(key).Args(element)
	if a.Rank != 0 {
		cmdBuilder = cmdBuilder.Args("RANK", strconv.FormatInt(a.Rank, 10))
	}
	if a.MaxLen != 0 {
		cmdBuilder = cmdBuilder.Args("MAXLEN", strconv.FormatInt(a.MaxLen, 10))
	}
	cmd := cmdBuilder.Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectLPosCount(key string, element string, count int64, a rueidiscompat.LPosArgs) *ExpectedIntSlice {
	cmdBuilder := m.raw.B().Arbitrary("LPOS").Keys(key).Args(element).Args("COUNT", strconv.FormatInt(count, 10))
	if a.Rank != 0 {
		cmdBuilder = cmdBuilder.Args("RANK", strconv.FormatInt(a.Rank, 10))
	}
	if a.MaxLen != 0 {
		cmdBuilder = cmdBuilder.Args("MAXLEN", strconv.FormatInt(a.MaxLen, 10))
	}
	cmd := cmdBuilder.Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedIntSlice{exp: e}
}

func (m *clientMock) ExpectRPopCount(key string, count int) *ExpectedStringSlice {
	cmd := m.raw.B().Rpop().Key(key).Count(int64(count)).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectSInterCard(limit int64, keys ...string) *ExpectedInt {
	cmd := m.raw.B().Sintercard().Numkeys(int64(len(keys))).Key(keys...).Limit(limit).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectSMembersMap(key string) *ExpectedStringStructMap {
	cmd := m.raw.B().Smembers().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringStructMap{exp: e}
}

func (m *clientMock) ExpectSetArgs(key string, value any, a rueidiscompat.SetArgs) *ExpectedStatus {
	cmdBuilder := m.raw.B().Arbitrary("SET").Keys(key).Args(str(value))
	if a.KeepTTL {
		cmdBuilder = cmdBuilder.Args("KEEPTTL")
	}
	if !a.ExpireAt.IsZero() {
		cmdBuilder = cmdBuilder.Args("EXAT", strconv.FormatInt(a.ExpireAt.Unix(), 10))
	}
	if a.TTL > 0 {
		if usePrecise(a.TTL) {
			cmdBuilder = cmdBuilder.Args("PX", strconv.FormatInt(formatMs(a.TTL), 10))
		} else {
			cmdBuilder = cmdBuilder.Args("EX", strconv.FormatInt(formatSec(a.TTL), 10))
		}
	}
	switch mode := strings.ToUpper(a.Mode); mode {
	case "XX", "NX":
		cmdBuilder = cmdBuilder.Args(mode)
	case "":
	default:
		panic(fmt.Sprintf("invalid mode for SET: %s", a.Mode))
	}
	if a.Get {
		cmdBuilder = cmdBuilder.Args("GET")
	}
	cmd := cmdBuilder.Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectGeoAdd(key string, geoLocation ...any) *ExpectedInt {
	cmd := m.raw.B().Geoadd().Key(key).LongitudeLatitudeMember()
	for _, loc := range geoLocation {
		normalized := normalizeGeoLocation(loc)
		cmd = cmd.LongitudeLatitudeMember(normalized.Longitude, normalized.Latitude, normalized.Name)
	}
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectGeoDist(key, member1, member2, unit string) *ExpectedFloat {
	var cmd rueidis.Completed
	switch strings.ToUpper(unit) {
	case "M":
		cmd = m.raw.B().Geodist().Key(key).Member1(member1).Member2(member2).M().Build()
	case "MI":
		cmd = m.raw.B().Geodist().Key(key).Member1(member1).Member2(member2).Mi().Build()
	case "FT":
		cmd = m.raw.B().Geodist().Key(key).Member1(member1).Member2(member2).Ft().Build()
	case "KM", "":
		cmd = m.raw.B().Geodist().Key(key).Member1(member1).Member2(member2).Km().Build()
	default:
		panic("invalid unit " + unit)
	}
	e := m.push(match(cmd.Commands()...), defaultFloatResult())
	return &ExpectedFloat{exp: e}
}

func (m *clientMock) ExpectGeoHash(key string, members ...string) *ExpectedStringSlice {
	cmd := m.raw.B().Geohash().Key(key).Member(members...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectGeoPos(key string, members ...string) *ExpectedGeoPos {
	cmd := m.raw.B().Geopos().Key(key).Member(members...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedGeoPos{exp: e}
}

func (m *clientMock) ExpectGeoRadius(key string, longitude, latitude float64, query *rueidiscompat.GeoRadiusQuery) *ExpectedGeoLocation {
	cmd := m.raw.B().Arbitrary("GEORADIUS_RO").Keys(key).Args(strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64))
	cmd = cmd.Args(geoRadiusQueryArgs(normalizeGeoRadiusQuery(query))...)
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultArrayResult())
	return &ExpectedGeoLocation{exp: e}
}

func (m *clientMock) ExpectGeoRadiusByMember(key, member string, query *rueidiscompat.GeoRadiusQuery) *ExpectedGeoLocation {
	cmd := m.raw.B().Arbitrary("GEORADIUSBYMEMBER_RO").Keys(key).Args(member)
	cmd = cmd.Args(geoRadiusQueryArgs(normalizeGeoRadiusQuery(query))...)
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultArrayResult())
	return &ExpectedGeoLocation{exp: e}
}

func (m *clientMock) ExpectGeoRadiusByMemberStore(key, member string, query *rueidiscompat.GeoRadiusQuery) *ExpectedInt {
	cmd := m.raw.B().Arbitrary("GEORADIUSBYMEMBER").Keys(key).Args(member)
	cmd = cmd.Args(geoRadiusQueryArgs(normalizeGeoRadiusQuery(query))...)
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectGeoRadiusStore(key string, longitude, latitude float64, query *rueidiscompat.GeoRadiusQuery) *ExpectedInt {
	cmd := m.raw.B().Arbitrary("GEORADIUS").Keys(key).Args(strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64))
	cmd = cmd.Args(geoRadiusQueryArgs(normalizeGeoRadiusQuery(query))...)
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectGeoSearch(key string, q *rueidiscompat.GeoSearchQuery) *ExpectedStringSlice {
	cmd := m.raw.B().Arbitrary("GEOSEARCH").Keys(key)
	cmd = cmd.Args(geoSearchQueryArgs(normalizeGeoSearchQuery(q))...)
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectGeoSearchLocation(key string, q *rueidiscompat.GeoSearchLocationQuery) *ExpectedGeoSearchLocation {
	cmd := m.raw.B().Arbitrary("GEOSEARCH").Keys(key)
	cmd = cmd.Args(geoSearchLocationQueryArgs(normalizeGeoSearchLocationQuery(q))...)
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultArrayResult())
	return &ExpectedGeoLocation{exp: e}
}

func (m *clientMock) ExpectGeoSearchStore(src, dest string, q *rueidiscompat.GeoSearchStoreQuery) *ExpectedInt {
	query := normalizeGeoSearchStoreQuery(q)
	cmd := m.raw.B().Arbitrary("GEOSEARCHSTORE").Keys(dest, src)
	cmd = cmd.Args(geoSearchQueryArgs(query.GeoSearchQuery)...)
	if query.StoreDist {
		cmd = cmd.Args("STOREDIST")
	}
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func normalizeGeoLocation(loc any) rueidiscompat.GeoLocation {
	switch l := loc.(type) {
	case rueidiscompat.GeoLocation:
		return l
	case *rueidiscompat.GeoLocation:
		if l != nil {
			return *l
		}
	}
	return rueidiscompat.GeoLocation{}
}

func normalizeGeoRadiusQuery(q any) rueidiscompat.GeoRadiusQuery {
	switch query := q.(type) {
	case rueidiscompat.GeoRadiusQuery:
		return query
	case *rueidiscompat.GeoRadiusQuery:
		if query != nil {
			return *query
		}
	}
	return rueidiscompat.GeoRadiusQuery{}
}

func normalizeGeoSearchQuery(q any) rueidiscompat.GeoSearchQuery {
	switch query := q.(type) {
	case rueidiscompat.GeoSearchQuery:
		return query
	case *rueidiscompat.GeoSearchQuery:
		if query != nil {
			return *query
		}
	}
	return rueidiscompat.GeoSearchQuery{}
}

func normalizeGeoSearchLocationQuery(q any) rueidiscompat.GeoSearchLocationQuery {
	switch query := q.(type) {
	case rueidiscompat.GeoSearchLocationQuery:
		return query
	case *rueidiscompat.GeoSearchLocationQuery:
		if query != nil {
			return *query
		}
	}
	return rueidiscompat.GeoSearchLocationQuery{}
}

func normalizeGeoSearchStoreQuery(q any) rueidiscompat.GeoSearchStoreQuery {
	switch query := q.(type) {
	case rueidiscompat.GeoSearchStoreQuery:
		return query
	case *rueidiscompat.GeoSearchStoreQuery:
		if query != nil {
			return *query
		}
	}
	return rueidiscompat.GeoSearchStoreQuery{}
}

// mirrors rueidiscompat.GeoRadiusQuery.args(); keep in sync if upstream changes.
func geoRadiusQueryArgs(q rueidiscompat.GeoRadiusQuery) []string {
	args := make([]string, 0, 2)
	args = append(args, strconv.FormatFloat(q.Radius, 'f', -1, 64))
	if q.Unit != "" {
		args = append(args, q.Unit)
	} else {
		args = append(args, "km")
	}
	if q.WithCoord {
		args = append(args, "WITHCOORD")
	}
	if q.WithDist {
		args = append(args, "WITHDIST")
	}
	if q.WithGeoHash {
		args = append(args, "WITHHASH")
	}
	if q.Count > 0 {
		args = append(args, "COUNT", strconv.FormatInt(q.Count, 10))
	}
	if q.Sort != "" {
		args = append(args, q.Sort)
	}
	if q.Store != "" {
		args = append(args, "STORE")
		args = append(args, q.Store)
	}
	if q.StoreDist != "" {
		args = append(args, "STOREDIST")
		args = append(args, q.StoreDist)
	}
	return args
}

// mirrors rueidiscompat.GeoSearchQuery.args(); keep in sync if upstream changes.
func geoSearchQueryArgs(q rueidiscompat.GeoSearchQuery) []string {
	args := make([]string, 0, 2)
	if q.Member != "" {
		args = append(args, "FROMMEMBER", q.Member)
	} else {
		args = append(args, "FROMLONLAT", strconv.FormatFloat(q.Longitude, 'f', -1, 64), strconv.FormatFloat(q.Latitude, 'f', -1, 64))
	}
	radiusUnit := q.RadiusUnit
	boxUnit := q.BoxUnit
	if q.Radius > 0 {
		if radiusUnit == "" {
			radiusUnit = "KM"
		}
		args = append(args, "BYRADIUS", strconv.FormatFloat(q.Radius, 'f', -1, 64), radiusUnit)
	} else {
		if boxUnit == "" {
			boxUnit = "KM"
		}
		args = append(args, "BYBOX", strconv.FormatFloat(q.BoxWidth, 'f', -1, 64), strconv.FormatFloat(q.BoxHeight, 'f', -1, 64), boxUnit)
	}
	if q.Sort != "" {
		args = append(args, q.Sort)
	}
	if q.Count > 0 {
		args = append(args, "COUNT", strconv.FormatInt(q.Count, 10))
		if q.CountAny {
			args = append(args, "ANY")
		}
	}
	return args
}

// mirrors rueidiscompat.GeoSearchLocationQuery.args(); keep in sync if upstream changes.
func geoSearchLocationQueryArgs(q rueidiscompat.GeoSearchLocationQuery) []string {
	args := geoSearchQueryArgs(q.GeoSearchQuery)
	if q.WithCoord {
		args = append(args, "WITHCOORD")
	}
	if q.WithDist {
		args = append(args, "WITHDIST")
	}
	if q.WithHash {
		args = append(args, "WITHHASH")
	}
	return args
}

func (m *clientMock) ExpectBZMPop(timeout time.Duration, order string, count int64, keys ...string) *ExpectedZSliceWithKey {
	cmdBuilder := m.raw.B().Arbitrary("BZMPOP", strconv.FormatInt(formatSec(timeout), 10), strconv.Itoa(len(keys))).Keys(keys...)
	cmdBuilder = cmdBuilder.Args(order)
	if count > 0 {
		cmdBuilder = cmdBuilder.Args("COUNT", strconv.FormatInt(count, 10))
	}
	cmd := cmdBuilder.Blocking()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSliceWithKey{exp: e}
}

func (m *clientMock) ExpectBZPopMax(timeout time.Duration, keys ...string) *ExpectedZWithKey {
	cmd := m.raw.B().Bzpopmax().Key(keys...).Timeout(float64(formatSec(timeout))).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZWithKey{exp: e}
}

func (m *clientMock) ExpectBZPopMin(timeout time.Duration, keys ...string) *ExpectedZWithKey {
	cmd := m.raw.B().Bzpopmin().Key(keys...).Timeout(float64(formatSec(timeout))).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZWithKey{exp: e}
}

func (m *clientMock) ExpectZDiff(keys ...string) *ExpectedStringSlice {
	cmd := m.raw.B().Zdiff().Numkeys(int64(len(keys))).Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZDiffStore(destination string, keys ...string) *ExpectedInt {
	cmd := m.raw.B().Zdiffstore().Destination(destination).Numkeys(int64(len(keys))).Key(keys...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZDiffWithScores(keys ...string) *ExpectedZSlice {
	cmd := m.raw.B().Zdiff().Numkeys(int64(len(keys))).Key(keys...).Withscores().Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

func (m *clientMock) ExpectZInterCard(limit int64, keys ...string) *ExpectedInt {
	cmd := m.raw.B().Zintercard().Numkeys(int64(len(keys))).Key(keys...).Limit(limit).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZMPop(order string, count int64, keys ...string) *ExpectedZSliceWithKey {
	cmdBuilder := m.raw.B().Arbitrary("ZMPOP", strconv.Itoa(len(keys))).Keys(keys...)
	cmdBuilder = cmdBuilder.Args(order)
	if count > 0 {
		cmdBuilder = cmdBuilder.Args("COUNT", strconv.FormatInt(count, 10))
	}
	cmd := cmdBuilder.Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSliceWithKey{exp: e}
}

func (m *clientMock) ExpectZRandMember(key string, count int) *ExpectedStringSlice {
	cmd := m.raw.B().Zrandmember().Key(key).Count(int64(count)).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZRandMemberWithScores(key string, count int) *ExpectedZSlice {
	cmd := m.raw.B().Zrandmember().Key(key).Count(int64(count)).Withscores().Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

func (m *clientMock) ExpectZRangeByLex(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedStringSlice {
	opt := normalizeZRangeBy(zRangeBy)
	var cmd rueidis.Completed
	if opt.Offset != 0 || opt.Count != 0 {
		cmd = m.raw.B().Zrangebylex().Key(key).Min(opt.Min).Max(opt.Max).Limit(opt.Offset, opt.Count).Build()
	} else {
		cmd = m.raw.B().Zrangebylex().Key(key).Min(opt.Min).Max(opt.Max).Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZRangeByScore(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedStringSlice {
	opt := normalizeZRangeBy(zRangeBy)
	var cmd rueidis.Completed
	if opt.Offset != 0 || opt.Count != 0 {
		cmd = m.raw.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Limit(opt.Offset, opt.Count).Build()
	} else {
		cmd = m.raw.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZRangeByScoreWithScores(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedZSlice {
	opt := normalizeZRangeBy(zRangeBy)
	var cmd rueidis.Completed
	if opt.Offset != 0 || opt.Count != 0 {
		cmd = m.raw.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Withscores().Limit(opt.Offset, opt.Count).Build()
	} else {
		cmd = m.raw.B().Zrangebyscore().Key(key).Min(opt.Min).Max(opt.Max).Withscores().Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

func (m *clientMock) ExpectZRangeStore(dst string, z rueidiscompat.ZRangeArgs) *ExpectedInt {
	cmdBuilder := m.raw.B().Arbitrary("ZRANGESTORE").Keys(dst, z.Key)
	cmdBuilder = cmdBuilder.Args(str(z.Start), str(z.Stop))
	if z.ByScore {
		cmdBuilder = cmdBuilder.Args("BYSCORE")
	} else if z.ByLex {
		cmdBuilder = cmdBuilder.Args("BYLEX")
	}
	if z.Rev {
		cmdBuilder = cmdBuilder.Args("REV")
	}
	if z.Offset != 0 || z.Count != 0 {
		cmdBuilder = cmdBuilder.Args("LIMIT", strconv.FormatInt(z.Offset, 10), strconv.FormatInt(z.Count, 10))
	}
	cmd := cmdBuilder.Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZRevRangeByLex(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedStringSlice {
	opt := normalizeZRangeBy(zRangeBy)
	var cmd rueidis.Completed
	if opt.Offset != 0 || opt.Count != 0 {
		cmd = m.raw.B().Zrevrangebylex().Key(key).Max(opt.Max).Min(opt.Min).Limit(opt.Offset, opt.Count).Build()
	} else {
		cmd = m.raw.B().Zrevrangebylex().Key(key).Max(opt.Max).Min(opt.Min).Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZRevRangeByScore(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedStringSlice {
	opt := normalizeZRangeBy(zRangeBy)
	var cmd rueidis.Completed
	if opt.Offset != 0 || opt.Count != 0 {
		cmd = m.raw.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Limit(opt.Offset, opt.Count).Build()
	} else {
		cmd = m.raw.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZRevRangeByScoreWithScores(key string, zRangeBy *rueidiscompat.ZRangeBy) *ExpectedZSlice {
	opt := normalizeZRangeBy(zRangeBy)
	var cmd rueidis.Completed
	if opt.Offset != 0 || opt.Count != 0 {
		cmd = m.raw.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Withscores().Limit(opt.Offset, opt.Count).Build()
	} else {
		cmd = m.raw.B().Zrevrangebyscore().Key(key).Max(opt.Max).Min(opt.Min).Withscores().Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

func (m *clientMock) ExpectZRevRangeWithScores(key string, start, stop int64) *ExpectedZSlice {
	cmd := m.raw.B().Zrevrange().Key(key).Start(start).Stop(stop).Withscores().Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

func (m *clientMock) ExpectZScan(key string, cursor uint64, pattern string, count int64) *ExpectedScan {
	cmdBuilder := m.raw.B().Arbitrary("ZSCAN").Keys(key).Args(strconv.FormatInt(int64(cursor), 10))
	if pattern != "" {
		cmdBuilder = cmdBuilder.Args("MATCH", pattern)
	}
	if count > 0 {
		cmdBuilder = cmdBuilder.Args("COUNT", strconv.FormatInt(count, 10))
	}
	cmd := cmdBuilder.ReadOnly()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedScan{exp: e}
}

func (m *clientMock) ExpectZAdd(key string, members ...rueidiscompat.Z) *ExpectedInt {
	cmd := compatZAddArgs(m.raw.B(), key, false, rueidiscompat.ZAddArgs{Members: members})
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZAddNX(key string, members ...rueidiscompat.Z) *ExpectedInt {
	cmd := compatZAddArgs(m.raw.B(), key, false, rueidiscompat.ZAddArgs{Members: members, NX: true})
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZAddXX(key string, members ...rueidiscompat.Z) *ExpectedInt {
	cmd := compatZAddArgs(m.raw.B(), key, false, rueidiscompat.ZAddArgs{Members: members, XX: true})
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZAddLT(key string, members ...rueidiscompat.Z) *ExpectedInt {
	cmd := compatZAddArgs(m.raw.B(), key, false, rueidiscompat.ZAddArgs{Members: members, LT: true})
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZAddGT(key string, members ...rueidiscompat.Z) *ExpectedInt {
	cmd := compatZAddArgs(m.raw.B(), key, false, rueidiscompat.ZAddArgs{Members: members, GT: true})
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZAddArgs(key string, args rueidiscompat.ZAddArgs) *ExpectedInt {
	cmd := compatZAddArgs(m.raw.B(), key, false, args)
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZAddArgsIncr(key string, args rueidiscompat.ZAddArgs) *ExpectedFloat {
	cmd := compatZAddArgs(m.raw.B(), key, true, args)
	e := m.push(match(cmd.Commands()...), defaultFloatResult())
	return &ExpectedFloat{exp: e}
}

func (m *clientMock) ExpectZInter(store *rueidiscompat.ZStore) *ExpectedStringSlice {
	cmd := compatZStore(m.raw.B().Arbitrary("ZINTER"), normalizeZStore(store)).ReadOnly()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZInterWithScores(store *rueidiscompat.ZStore) *ExpectedZSlice {
	cmd := compatZStore(m.raw.B().Arbitrary("ZINTER"), normalizeZStore(store)).Args("WITHSCORES").ReadOnly()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

func (m *clientMock) ExpectZInterStore(destination string, store *rueidiscompat.ZStore) *ExpectedInt {
	cmd := compatZStore(m.raw.B().Arbitrary("ZINTERSTORE").Keys(destination), normalizeZStore(store)).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectZUnion(store rueidiscompat.ZStore) *ExpectedStringSlice {
	cmd := compatZStore(m.raw.B().Arbitrary("ZUNION"), store).ReadOnly()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZUnionWithScores(store rueidiscompat.ZStore) *ExpectedZSlice {
	cmd := compatZStore(m.raw.B().Arbitrary("ZUNION"), store).Args("WITHSCORES").ReadOnly()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

func (m *clientMock) ExpectZUnionStore(dest string, store *rueidiscompat.ZStore) *ExpectedInt {
	cmd := compatZStore(m.raw.B().Arbitrary("ZUNIONSTORE").Keys(dest), normalizeZStore(store)).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func normalizeZStore(store any) rueidiscompat.ZStore {
	switch s := store.(type) {
	case rueidiscompat.ZStore:
		return s
	case *rueidiscompat.ZStore:
		if s != nil {
			return *s
		}
	}
	return rueidiscompat.ZStore{}
}

func normalizeZRangeBy(zRangeBy any) rueidiscompat.ZRangeBy {
	switch z := zRangeBy.(type) {
	case rueidiscompat.ZRangeBy:
		return z
	case *rueidiscompat.ZRangeBy:
		if z != nil {
			return *z
		}
	}
	return rueidiscompat.ZRangeBy{}
}

func (m *clientMock) ExpectZRangeArgs(z rueidiscompat.ZRangeArgs) *ExpectedStringSlice {
	cmd := compatZRangeArgs(m.raw.B(), false, z)
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectZRangeArgsWithScores(z rueidiscompat.ZRangeArgs) *ExpectedZSlice {
	cmd := compatZRangeArgs(m.raw.B(), true, z)
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

func (m *clientMock) ExpectZRangeWithScores(key string, start, stop int64) *ExpectedZSlice {
	cmd := compatZRangeArgs(m.raw.B(), true, rueidiscompat.ZRangeArgs{
		Key:   key,
		Start: start,
		Stop:  stop,
	})
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedZSlice{exp: e}
}

// mirrors the unexported zAddArgs helper in rueidiscompat; keep in sync if upstream changes.
func compatZAddArgs(b rueidis.Builder, key string, incr bool, args rueidiscompat.ZAddArgs) rueidis.Completed {
	cmd := b.Arbitrary("ZADD").Keys(key)
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
		cmd = cmd.Args(strconv.FormatFloat(v.Score, 'f', -1, 64), v.Member)
	}
	return cmd.Build()
}

// mirrors the unexported zstore helper in rueidiscompat; keep in sync if upstream changes.
func compatZStore(cmd cmds.Arbitrary, store rueidiscompat.ZStore) cmds.Arbitrary {
	cmd = cmd.Args(strconv.Itoa(len(store.Keys))).Keys(store.Keys...)
	if len(store.Weights) > 0 {
		cmd = cmd.Args("WEIGHTS")
		for _, w := range store.Weights {
			cmd = cmd.Args(strconv.FormatInt(w, 10))
		}
	}
	if store.Aggregate != "" {
		cmd = cmd.Args("AGGREGATE", store.Aggregate)
	}
	return cmd
}

// mirrors the unexported zRangeArgs helper in rueidiscompat; keep in sync if upstream changes.
func compatZRangeArgs(b rueidis.Builder, withScores bool, z rueidiscompat.ZRangeArgs) rueidis.Completed {
	cmd := b.Arbitrary("ZRANGE").Keys(z.Key)
	cmd = cmd.Args(str(z.Start), str(z.Stop))
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

func (m *clientMock) ExpectPubSubNumSub(channels ...string) *ExpectedMapStringInt {
	cmd := m.raw.B().PubsubNumsub().Channel(channels...).Build()
	e := m.push(match(cmd.Commands()...), defaultMapResult())
	return &ExpectedStringIntMap{exp: e}
}

func (m *clientMock) ExpectPubSubShardChannels(pattern string) *ExpectedStringSlice {
	cmd := m.raw.B().PubsubShardchannels().Pattern(pattern).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectPubSubShardNumSub(channels ...string) *ExpectedMapStringInt {
	cmd := m.raw.B().PubsubShardnumsub().Channel(channels...).Build()
	e := m.push(match(cmd.Commands()...), defaultMapResult())
	return &ExpectedStringIntMap{exp: e}
}

func (m *clientMock) ExpectScriptExists(hashes ...string) *ExpectedBoolSlice {
	m.expectRoleMaster()
	cmd := m.raw.B().ScriptExists().Sha1(hashes...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedBoolSlice{exp: e}
}

func (m *clientMock) ExpectScriptFlush() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().ScriptFlush().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectScriptKill() *ExpectedStatus {
	m.expectRoleMaster()
	cmd := m.raw.B().ScriptKill().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectScriptLoad(script string) *ExpectedString {
	m.expectRoleMaster()
	cmd := m.raw.B().ScriptLoad().Script(script).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectFCall(function string, keys []string, args ...any) *ExpectedCmd {
	cmd := m.raw.B().Fcall().Function(function).Numkeys(int64(len(keys))).Key(keys...).Arg(argsToSlice(args)...).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedCmd{exp: e}
}

func (m *clientMock) ExpectFCallRo(function string, keys []string, args ...any) *ExpectedCmd {
	cmd := m.raw.B().FcallRo().Function(function).Numkeys(int64(len(keys))).Key(keys...).Arg(argsToSlice(args)...).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedCmd{exp: e}
}

func (m *clientMock) ExpectFunctionDelete(libName string) *ExpectedString {
	m.expectRoleMaster()
	cmd := m.raw.B().FunctionDelete().LibraryName(libName).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectFunctionDump() *ExpectedString {
	cmd := m.raw.B().FunctionDump().Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectFunctionFlush() *ExpectedString {
	m.expectRoleMaster()
	cmd := m.raw.B().FunctionFlush().Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectFunctionFlushAsync() *ExpectedString {
	m.expectRoleMaster()
	cmd := m.raw.B().FunctionFlush().Async().Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectFunctionKill() *ExpectedString {
	cmd := m.raw.B().FunctionKill().Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectFunctionList(q rueidiscompat.FunctionListQuery) *ExpectedFunctionList {
	builder := m.raw.B().Arbitrary("FUNCTION", "LIST")
	if q.LibraryNamePattern != "" {
		builder = builder.Args("LIBRARYNAME", q.LibraryNamePattern)
	}
	if q.WithCode {
		builder = builder.Args("WITHCODE")
	}
	cmd := builder.Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedFunctionList{exp: e}
}

func (m *clientMock) ExpectFunctionLoad(code string) *ExpectedString {
	m.expectRoleMaster()
	cmd := m.raw.B().FunctionLoad().FunctionCode(code).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectFunctionLoadReplace(code string) *ExpectedString {
	m.expectRoleMaster()
	cmd := m.raw.B().FunctionLoad().Replace().FunctionCode(code).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectFunctionRestore(libDump string) *ExpectedString {
	m.expectRoleMaster()
	cmd := m.raw.B().FunctionRestore().SerializedValue(libDump).Build()
	e := m.push(match(cmd.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func (m *clientMock) ExpectXAck(stream, group string, ids ...string) *ExpectedInt {
	cmd := m.raw.B().Xack().Key(stream).Group(group).Id(ids...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectXAdd(args any) *ExpectedString {
	a := normalizeXAddArgs(args)
	cmd := m.raw.B().Arbitrary("XADD").Keys(a.Stream)
	if a.NoMkStream {
		cmd = cmd.Args("NOMKSTREAM")
	}
	if a.ProducerID != "" {
		if a.IdempotentAuto {
			cmd = cmd.Args("IDMPAUTO", a.ProducerID)
		} else if a.IdempotentID != "" {
			cmd = cmd.Args("IDMP", a.ProducerID, a.IdempotentID)
		}
	}
	switch {
	case a.MaxLen > 0:
		if a.Approx {
			cmd = cmd.Args("MAXLEN", "~", strconv.FormatInt(a.MaxLen, 10))
		} else {
			cmd = cmd.Args("MAXLEN", "=", strconv.FormatInt(a.MaxLen, 10))
		}
	case a.MinID != "":
		if a.Approx {
			cmd = cmd.Args("MINID", "~", a.MinID)
		} else {
			cmd = cmd.Args("MINID", "=", a.MinID)
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
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultStringResult())
	return &ExpectedString{exp: e}
}

func normalizeXAddArgs(args any) rueidiscompat.XAddArgs {
	switch a := args.(type) {
	case rueidiscompat.XAddArgs:
		return a
	case *rueidiscompat.XAddArgs:
		if a != nil {
			return *a
		}
	}
	return rueidiscompat.XAddArgs{}
}

func normalizeXAutoClaimArgs(args any) rueidiscompat.XAutoClaimArgs {
	switch a := args.(type) {
	case rueidiscompat.XAutoClaimArgs:
		return a
	case *rueidiscompat.XAutoClaimArgs:
		if a != nil {
			return *a
		}
	}
	return rueidiscompat.XAutoClaimArgs{}
}

func normalizeXClaimArgs(args any) rueidiscompat.XClaimArgs {
	switch a := args.(type) {
	case rueidiscompat.XClaimArgs:
		return a
	case *rueidiscompat.XClaimArgs:
		if a != nil {
			return *a
		}
	}
	return rueidiscompat.XClaimArgs{}
}

func normalizeXPendingExtArgs(args any) rueidiscompat.XPendingExtArgs {
	switch a := args.(type) {
	case rueidiscompat.XPendingExtArgs:
		return a
	case *rueidiscompat.XPendingExtArgs:
		if a != nil {
			return *a
		}
	}
	return rueidiscompat.XPendingExtArgs{}
}

func normalizeXReadArgs(args any) rueidiscompat.XReadArgs {
	switch a := args.(type) {
	case rueidiscompat.XReadArgs:
		return a
	case *rueidiscompat.XReadArgs:
		if a != nil {
			return *a
		}
	}
	return rueidiscompat.XReadArgs{}
}

func normalizeXReadGroupArgs(args any) rueidiscompat.XReadGroupArgs {
	switch a := args.(type) {
	case rueidiscompat.XReadGroupArgs:
		return a
	case *rueidiscompat.XReadGroupArgs:
		if a != nil {
			return *a
		}
	}
	return rueidiscompat.XReadGroupArgs{}
}

func (m *clientMock) ExpectXAutoClaim(args any) *ExpectedXAutoClaim {
	a := normalizeXAutoClaimArgs(args)
	var cmd rueidis.Completed
	if a.Count > 0 {
		cmd = m.raw.B().Xautoclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Start(a.Start).Count(a.Count).Build()
	} else {
		cmd = m.raw.B().Xautoclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Start(a.Start).Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedXAutoClaim{exp: e}
}

func (m *clientMock) ExpectXAutoClaimJustID(args any) *ExpectedXAutoClaimJustID {
	a := normalizeXAutoClaimArgs(args)
	var cmd rueidis.Completed
	if a.Count > 0 {
		cmd = m.raw.B().Xautoclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Start(a.Start).Count(a.Count).Justid().Build()
	} else {
		cmd = m.raw.B().Xautoclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Start(a.Start).Justid().Build()
	}
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedXAutoClaimJustID{exp: e}
}

func (m *clientMock) ExpectXClaim(args any) *ExpectedXMessageSlice {
	a := normalizeXClaimArgs(args)
	cmd := m.raw.B().Xclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Id(a.Messages...).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedXMessageSlice{exp: e}
}

func (m *clientMock) ExpectXClaimJustID(args any) *ExpectedStringSlice {
	a := normalizeXClaimArgs(args)
	cmd := m.raw.B().Xclaim().Key(a.Stream).Group(a.Group).Consumer(a.Consumer).MinIdleTime(strconv.FormatInt(formatMs(a.MinIdle), 10)).Id(a.Messages...).Justid().Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedStringSlice{exp: e}
}

func (m *clientMock) ExpectXDel(stream string, ids ...string) *ExpectedInt {
	cmd := m.raw.B().Xdel().Key(stream).Id(ids...).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectXGroupCreate(stream, group, start string) *ExpectedStatus {
	cmd := m.raw.B().XgroupCreate().Key(stream).Group(group).Id(start).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectXGroupCreateConsumer(stream, group, consumer string) *ExpectedInt {
	cmd := m.raw.B().XgroupCreateconsumer().Key(stream).Group(group).Consumer(consumer).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectXGroupCreateMkStream(stream, group, start string) *ExpectedStatus {
	cmd := m.raw.B().XgroupCreate().Key(stream).Group(group).Id(start).Mkstream().Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectXGroupDelConsumer(stream, group, consumer string) *ExpectedInt {
	cmd := m.raw.B().XgroupDelconsumer().Key(stream).Group(group).Consumername(consumer).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectXGroupDestroy(stream, group string) *ExpectedInt {
	cmd := m.raw.B().XgroupDestroy().Key(stream).Group(group).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectXGroupSetID(stream, group, start string) *ExpectedStatus {
	cmd := m.raw.B().XgroupSetid().Key(stream).Group(group).Id(start).Build()
	e := m.push(match(cmd.Commands()...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

func (m *clientMock) ExpectXInfoConsumers(key, group string) *ExpectedXInfoConsumers {
	cmd := m.raw.B().XinfoConsumers().Key(key).Group(group).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedXInfoConsumers{exp: e}
}

func (m *clientMock) ExpectXInfoGroups(key string) *ExpectedXInfoGroups {
	cmd := m.raw.B().XinfoGroups().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedXInfoGroups{exp: e}
}

func (m *clientMock) ExpectXInfoStream(key string) *ExpectedXInfoStream {
	cmd := m.raw.B().XinfoStream().Key(key).Build()
	e := m.push(match(cmd.Commands()...), defaultMapResult())
	return &ExpectedXInfoStream{exp: e}
}

func (m *clientMock) ExpectXInfoStreamFull(key string, count int) *ExpectedXInfoStreamFull {
	cmd := m.raw.B().XinfoStream().Key(key).Full().Count(int64(count)).Build()
	e := m.push(match(cmd.Commands()...), defaultMapResult())
	return &ExpectedXInfoStreamFull{exp: e}
}

func (m *clientMock) ExpectXLen(stream string) *ExpectedInt {
	cmd := m.raw.B().Xlen().Key(stream).Build()
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectXPending(stream, group string) *ExpectedXPending {
	cmd := m.raw.B().Xpending().Key(stream).Group(group).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedXPending{exp: e}
}

func (m *clientMock) ExpectXPendingExt(args any) *ExpectedXPendingExt {
	a := normalizeXPendingExtArgs(args)
	cmd := m.raw.B().Arbitrary("XPENDING").Keys(a.Stream).Args(a.Group)
	if a.Idle != 0 {
		cmd = cmd.Args("IDLE", strconv.FormatInt(formatMs(a.Idle), 10))
	}
	cmd = cmd.Args(a.Start, a.End, strconv.FormatInt(a.Count, 10))
	if a.Consumer != "" {
		cmd = cmd.Args(a.Consumer)
	}
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultArrayResult())
	return &ExpectedXPendingExt{exp: e}
}

func (m *clientMock) ExpectXRange(stream, start, stop string) *ExpectedXMessageSlice {
	cmd := m.raw.B().Xrange().Key(stream).Start(start).End(stop).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedXMessageSlice{exp: e}
}

func (m *clientMock) ExpectXRangeN(stream, start, stop string, count int64) *ExpectedXMessageSlice {
	cmd := m.raw.B().Xrange().Key(stream).Start(start).End(stop).Count(count).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedXMessageSlice{exp: e}
}

func (m *clientMock) ExpectXRead(args any) *ExpectedXStreamSlice {
	a := normalizeXReadArgs(args)
	cmd := m.raw.B().Arbitrary("XREAD")
	if a.Count > 0 {
		cmd = cmd.Args("COUNT", strconv.FormatInt(a.Count, 10))
	}
	if a.Block >= 0 {
		cmd = cmd.Args("BLOCK", strconv.FormatInt(formatMs(a.Block), 10))
	}
	cmd = cmd.Args("STREAMS")
	cmd = cmd.Keys(a.Streams[:len(a.Streams)/2]...).Args(a.Streams[len(a.Streams)/2:]...)
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultArrayResult())
	return &ExpectedXStreamSlice{exp: e}
}

func (m *clientMock) ExpectXReadGroup(args any) *ExpectedXStreamSlice {
	a := normalizeXReadGroupArgs(args)
	cmd := m.raw.B().Arbitrary("XREADGROUP")
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
	cmd = cmd.Keys(a.Streams[:len(a.Streams)/2]...).Args(a.Streams[len(a.Streams)/2:]...)
	built := cmd.Build()
	e := m.push(match(built.Commands()...), defaultArrayResult())
	return &ExpectedXStreamSlice{exp: e}
}

func (m *clientMock) ExpectXReadStreams(streams ...string) *ExpectedXStreamSlice {
	return m.ExpectXRead(&rueidiscompat.XReadArgs{Streams: streams, Block: -1})
}

func (m *clientMock) ExpectXRevRange(stream, start, stop string) *ExpectedXMessageSlice {
	cmd := m.raw.B().Xrevrange().Key(stream).End(start).Start(stop).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedXMessageSlice{exp: e}
}

func (m *clientMock) ExpectXRevRangeN(stream, start, stop string, count int64) *ExpectedXMessageSlice {
	cmd := m.raw.B().Xrevrange().Key(stream).End(start).Start(stop).Count(count).Build()
	e := m.push(match(cmd.Commands()...), defaultArrayResult())
	return &ExpectedXMessageSlice{exp: e}
}

func (m *clientMock) ExpectXTrimMaxLen(key string, maxLen int64) *ExpectedInt {
	cmd := compatXTrim(m.raw.B(), key, "MAXLEN", false, strconv.FormatInt(maxLen, 10), 0)
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectXTrimMaxLenApprox(key string, maxLen, limit int64) *ExpectedInt {
	cmd := compatXTrim(m.raw.B(), key, "MAXLEN", true, strconv.FormatInt(maxLen, 10), limit)
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectXTrimMinID(key string, minID string) *ExpectedInt {
	cmd := compatXTrim(m.raw.B(), key, "MINID", false, minID, 0)
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

func (m *clientMock) ExpectXTrimMinIDApprox(key string, minID string, limit int64) *ExpectedInt {
	cmd := compatXTrim(m.raw.B(), key, "MINID", true, minID, limit)
	e := m.push(match(cmd.Commands()...), defaultIntResult())
	return &ExpectedInt{exp: e}
}

// mirrors the unexported xTrim helper in rueidiscompat; keep in sync if upstream changes.
func compatXTrim(b rueidis.Builder, key, strategy string, approx bool, threshold string, limit int64) rueidis.Completed {
	cmd := b.Arbitrary("XTRIM").Keys(key).Args(strategy)
	if approx {
		cmd = cmd.Args("~")
	} else {
		cmd = cmd.Args("=")
	}
	cmd = cmd.Args(threshold)
	if limit > 0 {
		cmd = cmd.Args("LIMIT", strconv.FormatInt(limit, 10))
	}
	return cmd.Build()
}

func (m *clientMock) ExpectTxPipeline() {
	m.pushSet(match("MULTI"), mock.Result(mock.RedisString("OK")))
}

func (m *clientMock) ExpectTxPipelineExec() *ExpectedSlice {
	e := m.pushSet(match("EXEC"), mock.Result(mock.RedisArray()))
	return &ExpectedSlice{exp: e}
}

func (m *clientMock) ExpectWatch(keys ...string) *ExpectedError {
	args := append([]string{"WATCH"}, keys...)
	e := m.pushSet(match(args...), defaultStatusResult())
	return &ExpectedStatus{exp: e}
}

// primes ROLE=master for rueidiscompat fanout helpers (doPrimaries et al).
func (m *clientMock) expectRoleMaster() {
	cmd := m.raw.B().Role().Build()
	m.pushSet(match(cmd.Commands()...), mock.Result(mock.RedisArray(mock.RedisString("master"))))
}
