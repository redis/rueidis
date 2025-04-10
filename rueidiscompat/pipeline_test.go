// Copyright (c) 2013 The github.com/go-redis/redis Authors.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package rueidiscompat

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/rueidis/mock"
	"go.uber.org/mock/gomock"
)

var _ = Describe("RESP3 Pipeline Commands", func() {
	testAdapterPipeline(true)
})

var _ = Describe("RESP2 Pipeline Commands", func() {
	testAdapterPipeline(false)
})

func testAdapterPipeline(resp3 bool) {
	var adapter Cmdable

	BeforeEach(func() {
		if resp3 {
			adapter = adapterresp3
		} else {
			adapter = adapterresp2
		}
		Expect(adapter.FlushDB(ctx).Err()).NotTo(HaveOccurred())
		Expect(adapter.FlushAll(ctx).Err()).NotTo(HaveOccurred())
	})

	It("should Pipelined", func() {
		var echo, ping *StringCmd
		rets, err := adapter.Pipelined(ctx, func(pipe Pipeliner) error {
			echo = pipe.Echo(ctx, "hello")
			ping = pipe.Ping(ctx)
			Expect(echo.Err()).To(MatchError(placeholder.err))
			Expect(ping.Err()).To(MatchError(placeholder.err))
			return nil
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(rets).To(HaveLen(2))
		Expect(rets[0]).To(Equal(echo))
		Expect(rets[1]).To(Equal(ping))
		Expect(echo.Err()).NotTo(HaveOccurred())
		Expect(echo.Val()).To(Equal("hello"))
		Expect(ping.Err()).NotTo(HaveOccurred())
		Expect(ping.Val()).To(Equal("PONG"))
	})

	It("should Pipeline", func() {
		pipe := adapter.Pipeline()
		echo := pipe.Echo(ctx, "hello")
		ping := pipe.Ping(ctx)
		Expect(echo.Err()).To(MatchError(placeholder.err))
		Expect(ping.Err()).To(MatchError(placeholder.err))
		Expect(pipe.Len()).To(Equal(2))

		rets, err := pipe.Exec(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(pipe.Len()).To(Equal(0))
		Expect(rets).To(HaveLen(2))
		Expect(rets[0]).To(Equal(echo))
		Expect(rets[1]).To(Equal(ping))
		Expect(echo.Err()).NotTo(HaveOccurred())
		Expect(echo.Val()).To(Equal("hello"))
		Expect(ping.Err()).NotTo(HaveOccurred())
		Expect(ping.Val()).To(Equal("PONG"))
	})

	It("should Discard", func() {
		pipe := adapter.Pipeline()
		echo := pipe.Echo(ctx, "hello")
		ping := pipe.Ping(ctx)

		pipe.Discard()
		Expect(pipe.Len()).To(Equal(0))

		rets, err := pipe.Exec(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(rets).To(HaveLen(0))

		Expect(echo.Err()).To(MatchError(placeholder.err))
		Expect(ping.Err()).To(MatchError(placeholder.err))
	})
}

func TestPipeliner(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewClient(ctrl)

	var testPipeline = func(t *testing.T, p *Pipeline) {
		p.Command(ctx)
		p.CommandList(ctx, FilterBy{})
		p.CommandGetKeys(ctx, "1", "2")
		p.CommandGetKeysAndFlags(ctx, "1", "2")
		p.ClientGetName(ctx)
		p.Echo(ctx, "1")
		p.Ping(ctx)
		p.Quit(ctx)
		p.Del(ctx, "1", "2")
		p.Unlink(ctx, "1", "2")
		p.Dump(ctx, "1")
		p.Exists(ctx, "1", "2")
		p.Expire(ctx, "1", time.Second)
		p.ExpireAt(ctx, "1", time.UnixMilli(1724164643))
		p.ExpireTime(ctx, "1")
		p.ExpireNX(ctx, "1", time.Second)
		p.ExpireXX(ctx, "1", time.Second)
		p.ExpireGT(ctx, "1", time.Second)
		p.ExpireLT(ctx, "1", time.Second)
		p.Keys(ctx, "*")
		p.Migrate(ctx, "host", 0, "1", 0, time.Second)
		p.Move(ctx, "1", 0)
		p.ObjectRefCount(ctx, "1")
		p.ObjectEncoding(ctx, "1")
		p.ObjectIdleTime(ctx, "1")
		p.Persist(ctx, "1")
		p.PExpire(ctx, "1", time.Second)
		p.PExpireAt(ctx, "1", time.UnixMilli(1724164643))
		p.PExpireTime(ctx, "1")
		p.PTTL(ctx, "1")
		p.RandomKey(ctx)
		p.Rename(ctx, "1", "2")
		p.RenameNX(ctx, "1", "2")
		p.Restore(ctx, "1", time.Second, "v")
		p.RestoreReplace(ctx, "1", time.Second, "v")
		p.Sort(ctx, "1", Sort{})
		p.SortRO(ctx, "1", Sort{})
		p.SortStore(ctx, "1", "2", Sort{})
		p.SortInterfaces(ctx, "1", Sort{})
		p.Touch(ctx, "1", "2")
		p.TTL(ctx, "1")
		p.Type(ctx, "1")
		p.Append(ctx, "1", "v")
		p.Decr(ctx, "1")
		p.DecrBy(ctx, "1", 1)
		p.Get(ctx, "1")
		p.GetRange(ctx, "1", 0, 1)
		p.GetSet(ctx, "1", 2)
		p.GetEx(ctx, "1", time.Second)
		p.GetDel(ctx, "1")
		p.Incr(ctx, "1")
		p.IncrBy(ctx, "1", 1)
		p.IncrByFloat(ctx, "1", 1)
		p.MGet(ctx, "1", "2")
		p.MSet(ctx, 1, 2)
		p.MSetNX(ctx, 1, 2)
		p.Set(ctx, "1", 2, time.Second)
		p.SetArgs(ctx, "1", 2, SetArgs{})
		p.SetEX(ctx, "1", 2, time.Second)
		p.SetNX(ctx, "1", 2, time.Second)
		p.SetXX(ctx, "1", 2, time.Second)
		p.SetRange(ctx, "1", 2, "v")
		p.StrLen(ctx, "1")
		p.Copy(ctx, "1", "2", 0, true)
		p.GetBit(ctx, "1", 2)
		p.SetBit(ctx, "1", 2, 3)
		p.BitCount(ctx, "1", &BitCount{})
		p.BitOpAnd(ctx, "1", "1", "2")
		p.BitOpOr(ctx, "1", "1", "2")
		p.BitOpXor(ctx, "1", "1", "2")
		p.BitOpNot(ctx, "1", "2")
		p.BitPos(ctx, "1", 0, 0, 1)
		p.BitPosSpan(ctx, "1", 0, 1, 2, "1")
		p.BitField(ctx, "1", "2", "3")
		p.BitFieldRO(ctx, "1", "2", "3")
		p.Scan(ctx, 0, "*", 1)
		p.ScanType(ctx, 0, "*", 1, "1")
		p.SScan(ctx, "1", 0, "*", 1)
		p.HScan(ctx, "1", 0, "*", 1)
		p.HScanNoValues(ctx, "1", 0, "*", 1)
		p.ZScan(ctx, "1", 0, "*", 1)
		p.HDel(ctx, "1", "2", "3")
		p.HExists(ctx, "1", "2")
		p.HGet(ctx, "1", "2")
		p.HGetAll(ctx, "1")
		p.HIncrBy(ctx, "1", "2", 3)
		p.HIncrByFloat(ctx, "1", "2", 3)
		p.HKeys(ctx, "1")
		p.HLen(ctx, "1")
		p.HMGet(ctx, "1", "1", "2")
		p.HSet(ctx, "1", "1", "2")
		p.HMSet(ctx, "1", "1", "2")
		p.HSetNX(ctx, "1", "1", "2")
		p.HVals(ctx, "1")
		p.HRandField(ctx, "1", 1)
		p.HRandFieldWithValues(ctx, "1", 1)
		p.HExpire(ctx, "1", time.Second, "2", "3")
		p.HExpireWithArgs(ctx, "1", time.Second, HExpireArgs{NX: true}, "2", "3")
		p.HExpireWithArgs(ctx, "1", time.Second, HExpireArgs{XX: true}, "2", "3")
		p.HExpireWithArgs(ctx, "1", time.Second, HExpireArgs{GT: true}, "2", "3")
		p.HExpireWithArgs(ctx, "1", time.Second, HExpireArgs{LT: true}, "2", "3")
		p.HExpireWithArgs(ctx, "1", time.Second, HExpireArgs{}, "2", "3")
		p.HPExpire(ctx, "1", time.Second, "2", "3")
		p.HPExpireWithArgs(ctx, "1", time.Second, HExpireArgs{NX: true}, "2", "3")
		p.HPExpireWithArgs(ctx, "1", time.Second, HExpireArgs{XX: true}, "2", "3")
		p.HPExpireWithArgs(ctx, "1", time.Second, HExpireArgs{GT: true}, "2", "3")
		p.HPExpireWithArgs(ctx, "1", time.Second, HExpireArgs{LT: true}, "2", "3")
		p.HPExpireWithArgs(ctx, "1", time.Second, HExpireArgs{}, "2", "3")
		p.HExpireAt(ctx, "1", time.UnixMilli(1724164643), "2", "3")
		p.HExpireAtWithArgs(ctx, "1", time.UnixMilli(1724164643), HExpireArgs{NX: true}, "2", "3")
		p.HExpireAtWithArgs(ctx, "1", time.UnixMilli(1724164643), HExpireArgs{XX: true}, "2", "3")
		p.HExpireAtWithArgs(ctx, "1", time.UnixMilli(1724164643), HExpireArgs{GT: true}, "2", "3")
		p.HExpireAtWithArgs(ctx, "1", time.UnixMilli(1724164643), HExpireArgs{LT: true}, "2", "3")
		p.HExpireAtWithArgs(ctx, "1", time.UnixMilli(1724164643), HExpireArgs{}, "2", "3")
		p.HPExpireAt(ctx, "1", time.UnixMilli(1724164643), "2", "3")
		p.HPExpireAtWithArgs(ctx, "1", time.UnixMilli(1724164643), HExpireArgs{NX: true}, "2", "3")
		p.HPExpireAtWithArgs(ctx, "1", time.UnixMilli(1724164643), HExpireArgs{XX: true}, "2", "3")
		p.HPExpireAtWithArgs(ctx, "1", time.UnixMilli(1724164643), HExpireArgs{GT: true}, "2", "3")
		p.HPExpireAtWithArgs(ctx, "1", time.UnixMilli(1724164643), HExpireArgs{LT: true}, "2", "3")
		p.HPExpireAtWithArgs(ctx, "1", time.UnixMilli(1724164643), HExpireArgs{}, "2", "3")
		p.HPersist(ctx, "1", "2", "3")
		p.HExpireTime(ctx, "1", "2", "3")
		p.HPExpireTime(ctx, "1", "2", "3")
		p.HTTL(ctx, "1", "2", "3")
		p.HPTTL(ctx, "1", "2", "3")
		p.BLPop(ctx, time.Second, "1", "2")
		p.BLMPop(ctx, time.Second, "1", 1, "1", "2")
		p.BRPop(ctx, time.Second, "1", "2")
		p.BRPopLPush(ctx, "1", "2", time.Second)
		p.LIndex(ctx, "1", 1)
		p.LInsert(ctx, "1", "BEFORE", "3", "4")
		p.LInsertBefore(ctx, "1", "2", "3")
		p.LInsertAfter(ctx, "1", "2", "3")
		p.LLen(ctx, "1")
		p.LMPop(ctx, "1", 1, "1", "2")
		p.LPop(ctx, "1")
		p.LPopCount(ctx, "1", 1)
		p.LPos(ctx, "1", "v", LPosArgs{})
		p.LPosCount(ctx, "1", "v", 1, LPosArgs{})
		p.LPush(ctx, "1", "1", "2")
		p.LPushX(ctx, "1", "1", "2")
		p.LRange(ctx, "1", 1, 2)
		p.LRem(ctx, "1", 1, 2)
		p.LSet(ctx, "1", 1, 2)
		p.LTrim(ctx, "1", 1, 2)
		p.LCS(ctx, &LCSQuery{Key1: "a", Key2: "b"})
		p.RPop(ctx, "1")
		p.RPopCount(ctx, "1", 1)
		p.RPopLPush(ctx, "1", "2")
		p.RPush(ctx, "1", "1", "2")
		p.RPushX(ctx, "1", "1", "2")
		p.LMove(ctx, "1", "2", "3", "4")
		p.BLMove(ctx, "1", "2", "3", "4", time.Second)
		p.SAdd(ctx, "1", "2", "3")
		p.SCard(ctx, "1")
		p.SDiff(ctx, "1", "2")
		p.SDiffStore(ctx, "1", "1", "2")
		p.SInter(ctx, "1", "2")
		p.SInterCard(ctx, 1, "1", "2")
		p.SInterStore(ctx, "1", "1", "2")
		p.SIsMember(ctx, "1", "2")
		p.SMIsMember(ctx, "1", "2", "3")
		p.SMembers(ctx, "1")
		p.SMembersMap(ctx, "1")
		p.SMove(ctx, "1", "2", "3")
		p.SPop(ctx, "1")
		p.SPopN(ctx, "1", 1)
		p.SRandMember(ctx, "1")
		p.SRandMemberN(ctx, "1", 1)
		p.SRem(ctx, "1", "2", "3")
		p.SUnion(ctx, "1", "2")
		p.SUnionStore(ctx, "1", "2", "3")
		p.XAdd(ctx, XAddArgs{Stream: "stream", ID: "1-0", Values: map[string]any{"uno": "un"}})
		p.XDel(ctx, "1", "2", "3")
		p.XLen(ctx, "1")
		p.XRange(ctx, "1", "2", "3")
		p.XRangeN(ctx, "1", "2", "3", 1)
		p.XRevRange(ctx, "1", "2", "3")
		p.XRevRangeN(ctx, "1", "2", "3", 1)
		p.XRead(ctx, XReadArgs{Streams: []string{"stream", "0"}, Count: 2, Block: 100 * time.Millisecond})
		p.XReadStreams(ctx, "1", "2", "3")
		p.XGroupCreate(ctx, "1", "2", "3")
		p.XGroupCreateMkStream(ctx, "1", "2", "3")
		p.XGroupSetID(ctx, "1", "2", "3")
		p.XGroupDestroy(ctx, "1", "2")
		p.XGroupCreateConsumer(ctx, "1", "2", "3")
		p.XGroupDelConsumer(ctx, "1", "2", "3")
		p.XReadGroup(ctx, XReadGroupArgs{Group: "group", Consumer: "consumer", Streams: []string{"stream", ">"}})
		p.XAck(ctx, "1", "2", "3", "4")
		p.XPending(ctx, "1", "2")
		p.XPendingExt(ctx, XPendingExtArgs{Stream: "stream", Group: "group", Start: "-", End: "+", Count: 10, Consumer: "consumer"})
		p.XClaim(ctx, XClaimArgs{Stream: "stream", Group: "group", Consumer: "consumer", Messages: []string{"1-0", "2-0", "3-0"}})
		p.XClaimJustID(ctx, XClaimArgs{Stream: "stream", Group: "group", Consumer: "consumer", Messages: []string{"1-0", "2-0", "3-0"}})
		p.XAutoClaim(ctx, XAutoClaimArgs{Stream: "stream", Group: "group", Consumer: "consumer", Start: "-", Count: 2})
		p.XAutoClaimJustID(ctx, XAutoClaimArgs{Stream: "stream", Group: "group", Consumer: "consumer", Start: "-", Count: 2})
		p.XTrimMaxLen(ctx, "1", 1)
		p.XTrimMaxLenApprox(ctx, "1", 1, 2)
		p.XTrimMinID(ctx, "1", "2")
		p.XTrimMinIDApprox(ctx, "1", "2", 1)
		p.XInfoGroups(ctx, "1")
		p.XInfoStream(ctx, "1")
		p.XInfoStreamFull(ctx, "1", 1)
		p.XInfoConsumers(ctx, "1", "2")
		p.BZPopMax(ctx, time.Second, "1", "2")
		p.BZPopMin(ctx, time.Second, "1", "2")
		p.BZMPop(ctx, time.Second, "1", 0, "1", "2")
		p.ZAdd(ctx, "1", Z{Score: 2, Member: "two"}, Z{Score: 2, Member: "two"})
		p.ZAddLT(ctx, "1", Z{Score: 2, Member: "two"}, Z{Score: 2, Member: "two"})
		p.ZAddGT(ctx, "1", Z{Score: 2, Member: "two"}, Z{Score: 2, Member: "two"})
		p.ZAddNX(ctx, "1", Z{Score: 2, Member: "two"}, Z{Score: 2, Member: "two"})
		p.ZAddXX(ctx, "1", Z{Score: 2, Member: "two"}, Z{Score: 2, Member: "two"})
		p.ZAddArgs(ctx, "1", ZAddArgs{GT: true, Members: []Z{{Score: 1, Member: "one"}}})
		p.ZAddArgsIncr(ctx, "1", ZAddArgs{GT: true, Members: []Z{{Score: 1, Member: "one"}}})
		p.ZCard(ctx, "1")
		p.ZCount(ctx, "1", "2", "3")
		p.ZLexCount(ctx, "1", "2", "3")
		p.ZIncrBy(ctx, "1", 1, "2")
		p.ZInter(ctx, ZStore{Keys: []string{"zset1", "zset2"}, Weights: []int64{2, 3}})
		p.ZInterWithScores(ctx, ZStore{Keys: []string{"zset1", "zset2"}, Weights: []int64{2, 3}})
		p.ZInterCard(ctx, 1, "1", "2")
		p.ZInterStore(ctx, "1", ZStore{Keys: []string{"zset1", "zset2"}, Weights: []int64{2, 3}})
		p.ZMPop(ctx, "1", 1, "1", "2")
		p.ZMScore(ctx, "1", "2", "3")
		p.ZPopMax(ctx, "1", 1)
		p.ZPopMin(ctx, "1", 1)
		p.ZRange(ctx, "1", 1, 2)
		p.ZRangeWithScores(ctx, "1", 1, 2)
		p.ZRangeByScore(ctx, "1", ZRangeBy{Min: "-inf", Max: "+inf", Offset: 1, Count: 2})
		p.ZRangeByLex(ctx, "1", ZRangeBy{Min: "-inf", Max: "+inf", Offset: 1, Count: 2})
		p.ZRangeByScoreWithScores(ctx, "1", ZRangeBy{Min: "-inf", Max: "+inf", Offset: 1, Count: 2})
		p.ZRangeArgs(ctx, ZRangeArgs{Key: "zset", Start: 1, Stop: 4, ByScore: true, Rev: true, Offset: 1, Count: 2})
		p.ZRangeArgsWithScores(ctx, ZRangeArgs{Key: "zset", Start: 1, Stop: 4, ByScore: true, Rev: true, Offset: 1, Count: 2})
		p.ZRangeStore(ctx, "1", ZRangeArgs{Key: "zset", Start: 1, Stop: 4, ByScore: true, Rev: true, Offset: 1, Count: 2})
		p.ZRank(ctx, "1", "2")
		p.ZRankWithScore(ctx, "1", "2")
		p.ZRem(ctx, "1", "1", "2")
		p.ZRemRangeByRank(ctx, "1", 1, 2)
		p.ZRemRangeByScore(ctx, "1", "2", "3")
		p.ZRemRangeByLex(ctx, "1", "2", "3")
		p.ZRevRange(ctx, "1", 1, 2)
		p.ZRevRangeWithScores(ctx, "1", 1, 2)
		p.ZRevRangeByScore(ctx, "1", ZRangeBy{Min: "-inf", Max: "+inf", Offset: 1, Count: 2})
		p.ZRevRangeByLex(ctx, "1", ZRangeBy{Min: "-inf", Max: "+inf", Offset: 1, Count: 2})
		p.ZRevRangeByScoreWithScores(ctx, "1", ZRangeBy{Min: "-inf", Max: "+inf", Offset: 1, Count: 2})
		p.ZRevRank(ctx, "1", "2")
		p.ZRevRankWithScore(ctx, "1", "2")
		p.ZScore(ctx, "1", "2")
		p.ZUnionStore(ctx, "1", ZStore{Keys: []string{"zset1", "zset2"}, Weights: []int64{2, 3}})
		p.ZRandMember(ctx, "1", 1)
		p.ZRandMemberWithScores(ctx, "1", 1)
		p.ZUnion(ctx, ZStore{Keys: []string{"zset1", "zset2"}, Weights: []int64{2, 3}})
		p.ZUnionWithScores(ctx, ZStore{Keys: []string{"zset1", "zset2"}, Weights: []int64{2, 3}})
		p.ZDiff(ctx, "1", "2")
		p.ZDiffWithScores(ctx, "1", "2")
		p.ZDiffStore(ctx, "1", "1", "2")
		p.PFAdd(ctx, "1", "1", "2")
		p.PFCount(ctx, "1", "2")
		p.PFMerge(ctx, "1", "1", "2")
		p.BgRewriteAOF(ctx)
		p.BgSave(ctx)
		p.ClientKill(ctx, "1")
		p.ClientKillByFilter(ctx, "1", "2")
		p.ClientList(ctx)
		p.ClientPause(ctx, time.Second)
		p.ClientUnpause(ctx)
		p.ClientID(ctx)
		p.ClientUnblock(ctx, 1)
		p.ClientUnblockWithError(ctx, 1)
		p.ClientInfo(ctx)
		p.ConfigGet(ctx, "1")
		p.ConfigResetStat(ctx)
		p.ConfigSet(ctx, "1", "v")
		p.ConfigRewrite(ctx)
		p.DBSize(ctx)
		p.FlushAll(ctx)
		p.FlushAllAsync(ctx)
		p.FlushDB(ctx)
		p.FlushDBAsync(ctx)
		p.Info(ctx, "1", "2")
		p.LastSave(ctx)
		p.Save(ctx)
		p.Shutdown(ctx)
		p.ShutdownSave(ctx)
		p.ShutdownNoSave(ctx)
		p.Time(ctx)
		p.DebugObject(ctx, "1")
		p.ReadOnly(ctx)
		p.ReadWrite(ctx)
		p.MemoryUsage(ctx, "1", 1)
		p.Eval(ctx, "1", []string{"1"}, "2", "3")
		p.EvalSha(ctx, "1", []string{"1"}, "2", "3")
		p.EvalRO(ctx, "1", []string{"1"}, "2", "3")
		p.EvalShaRO(ctx, "1", []string{"1"}, "2", "3")
		p.ScriptExists(ctx, "1", "2")
		p.ScriptFlush(ctx)
		p.ScriptKill(ctx)
		p.ScriptLoad(ctx, "1")
		p.FunctionLoad(ctx, "1")
		p.FunctionLoadReplace(ctx, "1")
		p.FunctionDelete(ctx, "1")
		p.FunctionFlush(ctx)
		p.FunctionKill(ctx)
		p.FunctionFlushAsync(ctx)
		p.FunctionList(ctx, FunctionListQuery{LibraryNamePattern: "*", WithCode: true})
		p.FunctionDump(ctx)
		p.FunctionRestore(ctx, "1")
		p.FunctionStats(ctx)
		p.FCall(ctx, "1", []string{"1"}, "2", "3")
		p.FCallRO(ctx, "1", []string{"1"}, "2", "3")
		p.Publish(ctx, "1", "2")
		p.SPublish(ctx, "1", "2")
		p.PubSubChannels(ctx, "*")
		p.PubSubNumSub(ctx, "1", "2")
		p.PubSubNumPat(ctx)
		p.PubSubShardChannels(ctx, "*")
		p.PubSubShardNumSub(ctx, "1", "2")
		p.ClusterSlots(ctx)
		p.ClusterShards(ctx)
		p.ClusterNodes(ctx)
		p.ClusterMeet(ctx, "1", 1)
		p.ClusterForget(ctx, "1")
		p.ClusterReplicate(ctx, "1")
		p.ClusterResetSoft(ctx)
		p.ClusterResetHard(ctx)
		p.ClusterInfo(ctx)
		p.ClusterKeySlot(ctx, "1")
		p.ClusterGetKeysInSlot(ctx, 1, 2)
		p.ClusterCountFailureReports(ctx, "1")
		p.ClusterCountKeysInSlot(ctx, 1)
		p.ClusterDelSlots(ctx, 1, 2)
		p.ClusterDelSlotsRange(ctx, 1, 2)
		p.ClusterSaveConfig(ctx)
		p.ClusterSlaves(ctx, "1")
		p.ClusterFailover(ctx)
		p.ClusterAddSlots(ctx, 1, 2)
		p.ClusterAddSlotsRange(ctx, 1, 2)
		p.GeoAdd(ctx, "1", GeoLocation{Name: "k1", Longitude: 122.4194, Latitude: 37.7749}, GeoLocation{Name: "k1", Longitude: 122.4194, Latitude: 37.7749})
		p.GeoPos(ctx, "1", "1", "2")
		p.GeoRadius(ctx, "1", 1, 2, GeoRadiusQuery{Radius: 200})
		p.GeoRadiusStore(ctx, "1", 1, 2, GeoRadiusQuery{Radius: 200, StoreDist: "result"})
		p.GeoRadiusByMember(ctx, "1", "2", GeoRadiusQuery{Radius: 200})
		p.GeoRadiusByMemberStore(ctx, "1", "2", GeoRadiusQuery{Radius: 200, StoreDist: "result"})
		p.GeoSearch(ctx, "1", GeoSearchQuery{Member: "Catania", BoxWidth: 400, BoxHeight: 100, BoxUnit: "km", Sort: "asc"})
		p.GeoSearchLocation(ctx, "1", GeoSearchLocationQuery{GeoSearchQuery: GeoSearchQuery{Longitude: 15, Latitude: 37, Radius: 200, RadiusUnit: "km", Sort: "asc"}, WithCoord: true, WithDist: true, WithHash: true})
		p.GeoSearchStore(ctx, "1", "2", GeoSearchStoreQuery{GeoSearchQuery: GeoSearchQuery{Longitude: 15, Latitude: 37, Radius: 200, RadiusUnit: "km", Sort: "asc"}, StoreDist: false})
		p.GeoDist(ctx, "1", "2", "3", "M")
		p.GeoHash(ctx, "1", "2", "3", "4")
		p.ACLDryRun(ctx, "1", "2", "3")
		p.ACLSetUser(ctx, "1", "2", "3")
		p.ACLDelUser(ctx, "1")
		p.ACLLog(ctx, 1)
		p.ACLCat(ctx)
		p.ACLList(ctx)
		p.ACLLogReset(ctx)
		p.ACLCatArgs(ctx, &ACLCatArgs{Category: "read"})
		p.TFunctionLoad(ctx, "1")
		p.TFunctionLoadArgs(ctx, "1", &TFunctionLoadOptions{Replace: true, Config: `{"last_update_field_name":"last_update"}`})
		p.TFunctionDelete(ctx, "1")
		p.TFunctionList(ctx)
		p.TFunctionListArgs(ctx, &TFunctionListOptions{Withcode: true, Verbose: 2})
		p.TFCall(ctx, "1", "2", 3)
		p.TFCallArgs(ctx, "1", "2", 3, &TFCallOptions{Arguments: []string{"foo", "bar"}})
		p.TFCallASYNC(ctx, "1", "2", 3)
		p.TFCallASYNCArgs(ctx, "1", "2", 3, &TFCallOptions{Arguments: []string{"foo", "bar"}})
		p.BFAdd(ctx, "1", "2")
		p.BFCard(ctx, "1")
		p.BFExists(ctx, "1", "2")
		p.BFInfo(ctx, "1")
		p.BFInfoArg(ctx, "1", "SIZE")
		p.BFInfoCapacity(ctx, "1")
		p.BFInfoSize(ctx, "1")
		p.BFInfoFilters(ctx, "1")
		p.BFInfoItems(ctx, "1")
		p.BFInfoExpansion(ctx, "1")
		p.BFInsert(ctx, "1", &BFInsertOptions{Capacity: 2000, Error: 0.001, Expansion: 3, NonScaling: false, NoCreate: true}, "1", "2")
		p.BFMAdd(ctx, "1", "1", "2")
		p.BFMExists(ctx, "1", "1", "2")
		p.BFReserve(ctx, "1", 1, 2)
		p.BFReserveExpansion(ctx, "1", 1, 2, 3)
		p.BFReserveNonScaling(ctx, "1", 1, 2)
		p.BFReserveWithArgs(ctx, "1", &BFReserveOptions{Capacity: 2000, Error: 0.001, Expansion: 3, NonScaling: false})
		p.BFScanDump(ctx, "1", 1)
		p.BFLoadChunk(ctx, "1", 2, 3)
		p.CFAdd(ctx, "1", 1)
		p.CFAddNX(ctx, "1", 1)
		p.CFCount(ctx, "1", 1)
		p.CFDel(ctx, "1", 1)
		p.CFExists(ctx, "1", 1)
		p.CFInfo(ctx, "1")
		p.CFInsert(ctx, "1", &CFInsertOptions{Capacity: 3000, NoCreate: false}, 1, 2)
		p.CFInsertNX(ctx, "1", &CFInsertOptions{Capacity: 3000, NoCreate: false}, 1, 2)
		p.CFMExists(ctx, "1", 1, 2)
		p.CFReserve(ctx, "1", 1)
		p.CFReserveWithArgs(ctx, "1", &CFReserveOptions{Capacity: 2048, BucketSize: 3, MaxIterations: 15, Expansion: 2})
		p.CFReserveExpansion(ctx, "1", 1, 2)
		p.CFReserveBucketSize(ctx, "1", 1, 2)
		p.CFReserveMaxIterations(ctx, "1", 1, 2)
		p.CFScanDump(ctx, "1", 1)
		p.CFLoadChunk(ctx, "1", 1, 2)
		p.CMSIncrBy(ctx, "1", 1, 2)
		p.CMSInfo(ctx, "1")
		p.CMSInitByDim(ctx, "1", 1, 2)
		p.CMSInitByProb(ctx, "1", 1, 2)
		p.CMSMerge(ctx, "1", "2", "3")
		p.CMSMergeWithWeight(ctx, "1", map[string]int64{"1": 1})
		p.CMSQuery(ctx, "1", 1, 2)
		p.TopKAdd(ctx, "1", 1, 2)
		p.TopKCount(ctx, "1", 1, 2)
		p.TopKIncrBy(ctx, "1", 1, 2)
		p.TopKInfo(ctx, "1")
		p.TopKList(ctx, "1")
		p.TopKListWithCount(ctx, "1")
		p.TopKQuery(ctx, "1", 1, 2)
		p.TopKReserve(ctx, "1", 1)
		p.TopKReserveWithOptions(ctx, "1", 1, 2, 3, 4)
		p.TDigestAdd(ctx, "1", 1, 2)
		p.TDigestByRank(ctx, "1", 1, 2)
		p.TDigestByRevRank(ctx, "1", 1, 2)
		p.TDigestCDF(ctx, "1", 1, 2)
		p.TDigestCreate(ctx, "1")
		p.TDigestCreateWithCompression(ctx, "1", 1)
		p.TDigestInfo(ctx, "1")
		p.TDigestMax(ctx, "1")
		p.TDigestMin(ctx, "1")
		p.TDigestMerge(ctx, "1", &TDigestMergeOptions{Compression: 1000, Override: false}, "1", "2")
		p.TDigestQuantile(ctx, "1", 1, 2)
		p.TDigestRank(ctx, "1", 1, 2)
		p.TDigestReset(ctx, "1")
		p.TDigestRevRank(ctx, "1", 1, 2)
		p.TDigestTrimmedMean(ctx, "1", 1, 2)
		p.TSAdd(ctx, "1", 1, 1)
		p.TSAddWithArgs(ctx, "1", 1, 1, &TSOptions{Labels: map[string]string{"Time": "Series"}, Retention: 20})
		p.TSCreate(ctx, "1")
		p.TSCreateWithArgs(ctx, "1", &TSOptions{Labels: map[string]string{"Time": "Series"}, Retention: 20})
		p.TSAlter(ctx, "1", &TSAlterOptions{DuplicatePolicy: "min"})
		p.TSCreateRule(ctx, "1", "1", 1, 1)
		p.TSCreateRuleWithArgs(ctx, "1", "1", 1, 1, &TSCreateRuleOptions{AlignTimestamp: 1})
		p.TSIncrBy(ctx, "1", 1)
		p.TSIncrByWithArgs(ctx, "1", 1, &TSIncrDecrOptions{Timestamp: 5})
		p.TSDecrBy(ctx, "1", 1)
		p.TSDecrByWithArgs(ctx, "1", 1, &TSIncrDecrOptions{Timestamp: 5})
		p.TSDel(ctx, "1", 1, 1)
		p.TSDeleteRule(ctx, "1", "1")
		p.TSGet(ctx, "1")
		p.TSGetWithArgs(ctx, "1", &TSGetOptions{Latest: true})
		p.TSInfo(ctx, "1")
		p.TSInfoWithArgs(ctx, "1", &TSInfoOptions{Debug: true})
		p.TSMAdd(ctx, [][]interface{}{{1, 2, 3}, {1, 2, 3}})
		p.TSQueryIndex(ctx, []string{"1"})
		p.TSRevRange(ctx, "1", 1, 1)
		p.TSRevRangeWithArgs(ctx, "1", 1, 1, &TSRevRangeOptions{Count: 10})
		p.TSRange(ctx, "1", 1, 1)
		p.TSRangeWithArgs(ctx, "1", 1, 1, &TSRangeOptions{Count: 10})
		p.TSMRange(ctx, 1, 1, []string{"1"})
		p.TSMRangeWithArgs(ctx, 1, 1, []string{"1"}, &TSMRangeOptions{Count: 10})
		p.TSMRevRange(ctx, 1, 1, []string{"1"})
		p.TSMRevRangeWithArgs(ctx, 1, 1, []string{"1"}, &TSMRevRangeOptions{Count: 10})
		p.TSMGet(ctx, []string{"1"})
		p.TSMGetWithArgs(ctx, []string{"1"}, &TSMGetOptions{Latest: true})
		p.JSONArrAppend(ctx, "1", "1", "1", "2")
		p.JSONArrIndex(ctx, "1", "1", 1, 2)
		p.JSONArrIndexWithArgs(ctx, "1", "1", &JSONArrIndexArgs{Start: 1}, "1", "2")
		p.JSONArrInsert(ctx, "1", "1", 1, "1", "2")
		p.JSONArrLen(ctx, "1", "1")
		p.JSONArrPop(ctx, "1", "1", 1)
		p.JSONArrTrim(ctx, "1", "1")
		p.JSONArrTrimWithArgs(ctx, "1", "1", &JSONArrTrimArgs{Start: 1})
		p.JSONClear(ctx, "1", "1")
		p.JSONDebugMemory(ctx, "1", "1")
		p.JSONDel(ctx, "1", "1")
		p.JSONForget(ctx, "1", "1")
		p.JSONGet(ctx, "1", "1", "2")
		p.JSONGetWithArgs(ctx, "1", &JSONGetArgs{}, "1", "2")
		p.JSONMerge(ctx, "1", "1", "1")
		p.JSONMSetArgs(ctx, []JSONSetArgs{{Key: "1", Value: "1"}})
		p.JSONMSet(ctx, "mset1", "$.a", 2, "mset3", "$", `[1]`)
		p.JSONMGet(ctx, "1", "1", "1")
		p.JSONNumIncrBy(ctx, "1", "1", 1)
		p.JSONObjKeys(ctx, "1", "1")
		p.JSONObjLen(ctx, "1", "1")
		p.JSONSet(ctx, "1", "1", "2")
		p.JSONSetMode(ctx, "1", "1", "2", "NX")
		p.JSONStrAppend(ctx, "1", "1", "1")
		p.JSONStrLen(ctx, "1", "1")
		p.JSONToggle(ctx, "1", "1")
		p.JSONType(ctx, "1", "1")
		p.SlaveOf(ctx, "NO", "ONE")
		p.SlowLogGet(ctx, 1)
		p.SlowLogReset(ctx)
		p.ClusterMyShardID(ctx)
		p.ModuleLoadex(ctx, &ModuleLoadexConfig{
			Path: "/",
			Conf: map[string]any{"k": "v"},
			Args: []any{"1", "2"},
		})

		if n := len(p.rets); n != 491 {
			t.Fatalf("unexpected pipeline calls: %v", n)
		}
		for i, cmd := range p.rets {
			if err := cmd.Err(); !errors.Is(err, placeholder.err) {
				t.Fatalf("unexpected pipeline placeholder err(%d): %v", i, err)
			}
		}
		if n := len(p.comp.client.(*proxy).cmds); n != 491 {
			t.Fatalf("unexpected pipeline commands: %v", n)
		}
		var pipeline [][]string
		for _, cmd := range p.comp.client.(*proxy).cmds {
			pipeline = append(pipeline, cmd.Commands())
		}
		var expected [][]string
		if err := json.Unmarshal([]byte(golden), &expected); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(expected, pipeline) {
			actual, _ := json.Marshal(pipeline)
			expect, _ := json.Marshal(expected)
			t.Fatalf("unexpected pipeline results:\n%v\n%v", string(actual), string(expect))
		}
	}

	t.Run("Pipeline", func(t *testing.T) {
		testPipeline(t, newPipeline(m))
	})

	t.Run("Pipeline.Pipeline()", func(t *testing.T) {
		testPipeline(t, newPipeline(m).Pipeline().(*Pipeline))
	})

	t.Run("Pipeline.TxPipeline()", func(t *testing.T) {
		testPipeline(t, newPipeline(m).TxPipeline().(*Pipeline))
	})
}

var golden = `[
    ["COMMAND"],
    ["COMMAND","LIST"],
    ["COMMAND","GETKEYS","1","2"],
    ["COMMAND","GETKEYSANDFLAGS","1","2"],
    ["CLIENT","GETNAME"],
    ["ECHO","1"],
    ["PING"],
    ["QUIT"],
    ["DEL","1","2"],
    ["UNLINK","1","2"],
    ["DUMP","1"],
    ["EXISTS","1","2"],
    ["EXPIRE","1","1"],
    ["EXPIREAT","1","1724164"],
    ["EXPIRETIME","1"],
    ["EXPIRE","1","1","NX"],
    ["EXPIRE","1","1","XX"],
    ["EXPIRE","1","1","GT"],
    ["EXPIRE","1","1","LT"],
    ["KEYS","*"],
    ["MIGRATE","host","0","1","0","1"],
    ["MOVE","1","0"],
    ["OBJECT","REFCOUNT","1"],
    ["OBJECT","ENCODING","1"],
    ["OBJECT","IDLETIME","1"],
    ["PERSIST","1"],
    ["PEXPIRE","1","1000"],
    ["PEXPIREAT","1","1724164643"],
    ["PEXPIRETIME","1"],
    ["PTTL","1"],
    ["RANDOMKEY"],
    ["RENAME","1","2"],
    ["RENAMENX","1","2"],
    ["RESTORE","1","1000","v"],
    ["RESTORE","1","1000","v","REPLACE"],
    ["SORT","1"],
    ["SORT_RO","1"],
    ["SORT","1","STORE","2"],
    ["SORT","1"],
    ["TOUCH","1","2"],
    ["TTL","1"],
    ["TYPE","1"],
    ["APPEND","1","v"],
    ["DECR","1"],
    ["DECRBY","1","1"],
    ["GET","1"],
    ["GETRANGE","1","0","1"],
    ["GETSET","1","2"],
    ["GETEX","1","EX","1"],
    ["GETDEL","1"],
    ["INCR","1"],
    ["INCRBY","1","1"],
    ["INCRBYFLOAT","1","1"],
    ["MGET","1","2"],
    ["MSET","1","2"],
    ["MSETNX","1","2"],
    ["SET","1","2","EX","1"],
    ["SET","1","2"],
    ["SETEX","1","1","2"],
    ["SET","1","2","NX","EX","1"],
    ["SET","1","2","XX","EX","1"],
    ["SETRANGE","1","2","v"],
    ["STRLEN","1"],
    ["COPY","1","2","DB","0","REPLACE"],
    ["GETBIT","1","2"],
    ["SETBIT","1","2","3"],
    ["BITCOUNT","1","0","0"],
    ["BITOP","AND","1","1","2"],
    ["BITOP","OR","1","1","2"],
    ["BITOP","XOR","1","1","2"],
    ["BITOP","NOT","1","2"],
    ["BITPOS","1","0","0","1"],
    ["BITPOS","1","0","1","2","BYTE"],
    ["BITFIELD","1","2","3"],
    ["BITFIELD_RO","1","GET","2","3"],
    ["SCAN","0","MATCH","*","COUNT","1"],
    ["SCAN","0","MATCH","*","COUNT","1","TYPE","1"],
    ["SSCAN","1","0","MATCH","*","COUNT","1"],
    ["HSCAN","1","0","MATCH","*","COUNT","1"],
    ["HSCAN","1","0","MATCH","*","COUNT","1","NOVALUES"],
    ["ZSCAN","1","0","MATCH","*","COUNT","1"],
    ["HDEL","1","2","3"],
    ["HEXISTS","1","2"],
    ["HGET","1","2"],
    ["HGETALL","1"],
    ["HINCRBY","1","2","3"],
    ["HINCRBYFLOAT","1","2","3"],
    ["HKEYS","1"],
    ["HLEN","1"],
    ["HMGET","1","1","2"],
    ["HSET","1","1","2"],
    ["HSET","1","1","2"],
    ["HSETNX","1","1","2"],
    ["HVALS","1"],
    ["HRANDFIELD","1","1"],
    ["HRANDFIELD","1","1","WITHVALUES"],
    ["HEXPIRE","1","1","FIELDS","2","2","3"],
    ["HEXPIRE","1","1","NX","FIELDS","2","2","3"],
    ["HEXPIRE","1","1","XX","FIELDS","2","2","3"],
    ["HEXPIRE","1","1","GT","FIELDS","2","2","3"],
    ["HEXPIRE","1","1","LT","FIELDS","2","2","3"],
    ["HEXPIRE","1","1","FIELDS","2","2","3"],
    ["HPEXPIRE","1","1000","FIELDS","2","2","3"],
    ["HPEXPIRE","1","1000","NX","FIELDS","2","2","3"],
    ["HPEXPIRE","1","1000","XX","FIELDS","2","2","3"],
    ["HPEXPIRE","1","1000","GT","FIELDS","2","2","3"],
    ["HPEXPIRE","1","1000","LT","FIELDS","2","2","3"],
    ["HPEXPIRE","1","1000","FIELDS","2","2","3"],
    ["HEXPIREAT","1","1724164","FIELDS","2","2","3"],
    ["HEXPIREAT","1","1724164","NX","FIELDS","2","2","3"],
    ["HEXPIREAT","1","1724164","XX","FIELDS","2","2","3"],
    ["HEXPIREAT","1","1724164","GT","FIELDS","2","2","3"],
    ["HEXPIREAT","1","1724164","LT","FIELDS","2","2","3"],
    ["HEXPIREAT","1","1724164","FIELDS","2","2","3"],
    ["HPEXPIREAT","1","1724164643","FIELDS","2","2","3"],
    ["HPEXPIREAT","1","1724164643","NX","FIELDS","2","2","3"],
    ["HPEXPIREAT","1","1724164643","XX","FIELDS","2","2","3"],
    ["HPEXPIREAT","1","1724164643","GT","FIELDS","2","2","3"],
    ["HPEXPIREAT","1","1724164643","LT","FIELDS","2","2","3"],
    ["HPEXPIREAT","1","1724164643","FIELDS","2","2","3"],
    ["HPERSIST","1","FIELDS","2","2","3"],
    ["HEXPIRETIME","1","FIELDS","2","2","3"],
    ["HPEXPIRETIME","1","FIELDS","2","2","3"],
    ["HTTL","1","FIELDS","2","2","3"],
    ["HPTTL","1","FIELDS","2","2","3"],
    ["BLPOP","1","2","1"],
    ["BLMPOP","1","2","1","2","1","COUNT","1"],
    ["BRPOP","1","2","1"],
    ["BRPOPLPUSH","1","2","1"],
    ["LINDEX","1","1"],
    ["LINSERT","1","BEFORE","3","4"],
    ["LINSERT","1","BEFORE","2","3"],
    ["LINSERT","1","AFTER","2","3"],
    ["LLEN","1"],
    ["LMPOP","2","1","2","1","COUNT","1"],
    ["LPOP","1"],
    ["LPOP","1","1"],
    ["LPOS","1","v"],
    ["LPOS","1","v","COUNT","1"],
    ["LPUSH","1","1","2"],
    ["LPUSHX","1","1","2"],
    ["LRANGE","1","1","2"],
    ["LREM","1","1","2"],
    ["LSET","1","1","2"],
    ["LTRIM","1","1","2"],
    ["LCS","a","b"],
    ["RPOP","1"],
    ["RPOP","1","1"],
    ["RPOPLPUSH","1","2"],
    ["RPUSH","1","1","2"],
    ["RPUSHX","1","1","2"],
    ["LMOVE","1","2","3","4"],
    ["BLMOVE","1","2","3","4","1"],
    ["SADD","1","2","3"],
    ["SCARD","1"],
    ["SDIFF","1","2"],
    ["SDIFFSTORE","1","1","2"],
    ["SINTER","1","2"],
    ["SINTERCARD","2","1","2","LIMIT","1"],
    ["SINTERSTORE","1","1","2"],
    ["SISMEMBER","1","2"],
    ["SMISMEMBER","1","2","3"],
    ["SMEMBERS","1"],
    ["SMEMBERS","1"],
    ["SMOVE","1","2","3"],
    ["SPOP","1"],
    ["SPOP","1","1"],
    ["SRANDMEMBER","1"],
    ["SRANDMEMBER","1","1"],
    ["SREM","1","2","3"],
    ["SUNION","1","2"],
    ["SUNIONSTORE","1","2","3"],
    ["XADD","stream","1-0","uno","un"],
    ["XDEL","1","2","3"],
    ["XLEN","1"],
    ["XRANGE","1","2","3"],
    ["XRANGE","1","2","3","COUNT","1"],
    ["XREVRANGE","1","2","3"],
    ["XREVRANGE","1","2","3","COUNT","1"],
    ["XREAD","COUNT","2","BLOCK","100","STREAMS","stream","0"],
    ["XREAD","STREAMS","1","2","3"],
    ["XGROUP","CREATE","1","2","3"],
    ["XGROUP","CREATE","1","2","3","MKSTREAM"],
    ["XGROUP","SETID","1","2","3"],
    ["XGROUP","DESTROY","1","2"],
    ["XGROUP","CREATECONSUMER","1","2","3"],
    ["XGROUP","DELCONSUMER","1","2","3"],
    ["XREADGROUP","GROUP","group","consumer","BLOCK","0","STREAMS","stream","\u003e"],
    ["XACK","1","2","3","4"],
    ["XPENDING","1","2"],
    ["XPENDING","stream","group","-","+","10","consumer"],
    ["XCLAIM","stream","group","consumer","0","1-0","2-0","3-0"],
    ["XCLAIM","stream","group","consumer","0","1-0","2-0","3-0","JUSTID"],
    ["XAUTOCLAIM","stream","group","consumer","0","-","COUNT","2"],
    ["XAUTOCLAIM","stream","group","consumer","0","-","COUNT","2","JUSTID"],
    ["XTRIM","1","MAXLEN","1"],
    ["XTRIM","1","MAXLEN","~","1","LIMIT","2"],
    ["XTRIM","1","MINID","2"],
    ["XTRIM","1","MINID","~","2","LIMIT","1"],
    ["XINFO","GROUPS","1"],
    ["XINFO","STREAM","1"],
    ["XINFO","STREAM","1","FULL","COUNT","1"],
    ["XINFO","CONSUMERS","1","2"],
    ["BZPOPMAX","1","2","1"],
    ["BZPOPMIN","1","2","1"],
    ["BZMPOP","1","2","1","2","1"],
    ["ZADD","1","2","two","2","two"],
    ["ZADD","1","LT","2","two","2","two"],
    ["ZADD","1","GT","2","two","2","two"],
    ["ZADD","1","NX","2","two","2","two"],
    ["ZADD","1","XX","2","two","2","two"],
    ["ZADD","1","GT","1","one"],
    ["ZADD","1","GT","INCR","1","one"],
    ["ZCARD","1"],
    ["ZCOUNT","1","2","3"],
    ["ZLEXCOUNT","1","2","3"],
    ["ZINCRBY","1","1","2"],
    ["ZINTER","2","zset1","zset2","WEIGHTS","2","3"],
    ["ZINTER","2","zset1","zset2","WEIGHTS","2","3","WITHSCORES"],
    ["ZINTERCARD","2","1","2","LIMIT","1"],
    ["ZINTERSTORE","1","2","zset1","zset2","WEIGHTS","2","3"],
    ["ZMPOP","2","1","2","1","COUNT","1"],
    ["ZMSCORE","1","2","3"],
    ["ZPOPMAX","1","1"],
    ["ZPOPMIN","1","1"],
    ["ZRANGE","1","1","2"],
    ["ZRANGE","1","1","2","WITHSCORES"],
    ["ZRANGEBYSCORE","1","-inf","+inf","LIMIT","1","2"],
    ["ZRANGEBYLEX","1","-inf","+inf","LIMIT","1","2"],
    ["ZRANGEBYSCORE","1","-inf","+inf","WITHSCORES","LIMIT","1","2"],
    ["ZRANGE","zset","4","1","BYSCORE","REV","LIMIT","1","2"],
    ["ZRANGE","zset","4","1","BYSCORE","REV","LIMIT","1","2","WITHSCORES"],
    ["ZRANGESTORE","1","zset","4","1","BYSCORE","REV","LIMIT","1","2"],
    ["ZRANK","1","2"],
    ["ZRANK","1","2","WITHSCORE"],
    ["ZREM","1","1","2"],
    ["ZREMRANGEBYRANK","1","1","2"],
    ["ZREMRANGEBYSCORE","1","2","3"],
    ["ZREMRANGEBYLEX","1","2","3"],
    ["ZREVRANGE","1","1","2"],
    ["ZREVRANGE","1","1","2","WITHSCORES"],
    ["ZREVRANGEBYSCORE","1","+inf","-inf","LIMIT","1","2"],
    ["ZREVRANGEBYLEX","1","+inf","-inf","LIMIT","1","2"],
    ["ZREVRANGEBYSCORE","1","+inf","-inf","WITHSCORES","LIMIT","1","2"],
    ["ZREVRANK","1","2"],
    ["ZREVRANK","1","2","WITHSCORE"],
    ["ZSCORE","1","2"],
    ["ZUNIONSTORE","1","2","zset1","zset2","WEIGHTS","2","3"],
    ["ZRANDMEMBER","1","1"],
    ["ZRANDMEMBER","1","1","WITHSCORES"],
    ["ZUNION","2","zset1","zset2","WEIGHTS","2","3"],
    ["ZUNION","2","zset1","zset2","WEIGHTS","2","3","WITHSCORES"],
    ["ZDIFF","2","1","2"],
    ["ZDIFF","2","1","2","WITHSCORES"],
    ["ZDIFFSTORE","1","2","1","2"],
    ["PFADD","1","1","2"],
    ["PFCOUNT","1","2"],
    ["PFMERGE","1","1","2"],
    ["BGREWRITEAOF"],
    ["BGSAVE"],
    ["CLIENT","KILL","1"],
    ["CLIENT","KILL","1","2"],
    ["CLIENT","LIST"],
    ["CLIENT","PAUSE","1"],
    ["CLIENT","UNPAUSE"],
    ["CLIENT","ID"],
    ["CLIENT","UNBLOCK","1"],
    ["CLIENT","UNBLOCK","1","ERROR"],
    ["CLIENT","INFO"],
    ["CONFIG","GET","1"],
    ["CONFIG","RESETSTAT"],
    ["CONFIG","SET","1","v"],
    ["CONFIG","REWRITE"],
    ["DBSIZE"],
    ["FLUSHALL"],
    ["FLUSHALL","ASYNC"],
    ["FLUSHDB"],
    ["FLUSHDB","ASYNC"],
    ["INFO","1","2"],
    ["LASTSAVE"],
    ["SAVE"],
    ["SHUTDOWN"],
    ["SHUTDOWN","SAVE"],
    ["SHUTDOWN","NOSAVE"],
    ["TIME"],
    ["DEBUG","OBJECT","1"],
    ["READONLY"],
    ["READWRITE"],
    ["MEMORY","USAGE","1","SAMPLES","1"],
    ["EVAL","1","1","1","2","3"],
    ["EVALSHA","1","1","1","2","3"],
    ["EVAL_RO","1","1","1","2","3"],
    ["EVALSHA_RO","1","1","1","2","3"],
    ["SCRIPT","EXISTS","1","2"],
    ["SCRIPT","FLUSH"],
    ["SCRIPT","KILL"],
    ["SCRIPT","LOAD","1"],
    ["FUNCTION","LOAD","1"],
    ["FUNCTION","LOAD","REPLACE","1"],
    ["FUNCTION","DELETE","1"],
    ["FUNCTION","FLUSH"],
    ["FUNCTION","KILL"],
    ["FUNCTION","FLUSH","ASYNC"],
    ["FUNCTION","LIST","LIBRARYNAME","*","WITHCODE"],
    ["FUNCTION","DUMP"],
    ["FUNCTION","RESTORE","1"],
    ["FUNCTION","STATS"],
    ["FCALL","1","1","1","2","3"],
    ["FCALL_RO","1","1","1","2","3"],
    ["PUBLISH","1","2"],
    ["SPUBLISH","1","2"],
    ["PUBSUB","CHANNELS","*"],
    ["PUBSUB","NUMSUB","1","2"],
    ["PUBSUB","NUMPAT"],
    ["PUBSUB","SHARDCHANNELS","*"],
    ["PUBSUB","SHARDNUMSUB","1","2"],
    ["CLUSTER","SLOTS"],
    ["CLUSTER","SHARDS"],
    ["CLUSTER","NODES"],
    ["CLUSTER","MEET","1","1"],
    ["CLUSTER","FORGET","1"],
    ["CLUSTER","REPLICATE","1"],
    ["CLUSTER","RESET","SOFT"],
    ["CLUSTER","RESET","HARD"],
    ["CLUSTER","INFO"],
    ["CLUSTER","KEYSLOT","1"],
    ["CLUSTER","GETKEYSINSLOT","1","2"],
    ["CLUSTER","COUNT-FAILURE-REPORTS","1"],
    ["CLUSTER","COUNTKEYSINSLOT","1"],
    ["CLUSTER","DELSLOTS","1","2"],
    ["CLUSTER","DELSLOTSRANGE","1","2"],
    ["CLUSTER","SAVECONFIG"],
    ["CLUSTER","SLAVES","1"],
    ["CLUSTER","FAILOVER"],
    ["CLUSTER","ADDSLOTS","1","2"],
    ["CLUSTER","ADDSLOTSRANGE","1","2"],
    ["GEOADD","1","122.4194","37.7749","k1","122.4194","37.7749","k1"],
    ["GEOPOS","1","1","2"],
    ["GEORADIUS_RO","1","1","2","200","km"],
    ["GEORADIUS","1","1","2","200","km","STOREDIST","result"],
    ["GEORADIUSBYMEMBER_RO","1","2","200","km"],
    ["GEORADIUSBYMEMBER","1","2","200","km","STOREDIST","result"],
    ["GEOSEARCH","1","FROMMEMBER","Catania","BYBOX","400","100","km","asc"],
    ["GEOSEARCH","1","FROMLONLAT","15","37","BYRADIUS","200","km","asc","WITHCOORD","WITHDIST","WITHHASH"],
    ["GEOSEARCHSTORE","2","1","FROMLONLAT","15","37","BYRADIUS","200","km","asc"],
    ["GEODIST","1","2","3","m"],
    ["GEOHASH","1","2","3","4"],
    ["ACL","DRYRUN","1","2","3"],
    ["ACL","SETUSER","1","2","3"],
    ["ACL","DELUSER","1"],
    ["ACL","LOG","1"],
    ["ACL","CAT"],
    ["ACL","LIST"],
    ["ACL","LOG","RESET"],
    ["ACL","CAT","read"],
    ["TFUNCTION","LOAD","1"],
    ["TFUNCTION","LOAD","REPLACE","CONFIG","{\"last_update_field_name\":\"last_update\"}","1"],
    ["TFUNCTION","DELETE","1"],
    ["TFUNCTION","LIST"],
    ["TFUNCTION","LIST","WITHCODE","VERBOSE","V","V"],
    ["TFCALL","1.2","3"],
    ["TFCALL","1.2","3","foo","bar"],
    ["TFCALLASYNC","1.2","3"],
    ["TFCALLASYNC","1.2","3","foo","bar"],
    ["BF.ADD","1","2"],
    ["BF.CARD","1"],
    ["BF.EXISTS","1","2"],
    ["BF.INFO","1"],
    ["BF.INFO","1","SIZE"],
    ["BF.INFO","1","CAPACITY"],
    ["BF.INFO","1","SIZE"],
    ["BF.INFO","1","FILTERS"],
    ["BF.INFO","1","ITEMS"],
    ["BF.INFO","1","EXPANSION"],
    ["BF.INSERT","1","CAPACITY","2000","ERROR","0.001","EXPANSION","3","NOCREATE","ITEMS","1","2"],
    ["BF.MADD","1","1","2"],
    ["BF.MEXISTS","1","1","2"],
    ["BF.RESERVE","1","1","2"],
    ["BF.RESERVE","1","1","2","EXPANSION","3"],
    ["BF.RESERVE","1","1","2","NONSCALING"],
    ["BF.RESERVE","1","0.001","2000","EXPANSION","3"],
    ["BF.SCANDUMP","1","1"],
    ["BF.LOADCHUNK","1","2","3"],
    ["CF.ADD","1","1"],
    ["CF.ADDNX","1","1"],
    ["CF.COUNT","1","1"],
    ["CF.DEL","1","1"],
    ["CF.EXISTS","1","1"],
    ["CF.INFO","1"],
    ["CF.INSERT","1","CAPACITY","3000","ITEMS","1","2"],
    ["CF.INSERTNX","1","CAPACITY","3000","ITEMS","1","2"],
    ["CF.MEXISTS","1","1","2"],
    ["CF.RESERVE","1","1"],
    ["CF.RESERVE","1","2048","BUCKETSIZE","3","MAXITERATIONS","15","EXPANSION","2"],
    ["CF.RESERVE","1","1","EXPANSION","2"],
    ["CF.RESERVE","1","1","BUCKETSIZE","2"],
    ["CF.RESERVE","1","1","MAXITERATIONS","2"],
    ["CF.SCANDUMP","1","1"],
    ["CF.LOADCHUNK","1","1","2"],
    ["CMS.INCRBY","1","1","2"],
    ["CMS.INFO","1"],
    ["CMS.INITBYDIM","1","1","2"],
    ["CMS.INITBYPROB","1","1","2"],
    ["CMS.MERGE","1","2","2","3"],
    ["CMS.MERGE","1","1","1","WEIGHTS","1"],
    ["CMS.QUERY","1","1","2"],
    ["TOPK.ADD","1","1","2"],
    ["TOPK.COUNT","1","1","2"],
    ["TOPK.INCRBY","1","1","2"],
    ["TOPK.INFO","1"],
    ["TOPK.LIST","1"],
    ["TOPK.LIST","1","WITHCOUNT"],
    ["TOPK.QUERY","1","1","2"],
    ["TOPK.RESERVE","1","1"],
    ["TOPK.RESERVE","1","1","2","3","4"],
    ["TDIGEST.ADD","1","1","2"],
    ["TDIGEST.BYRANK","1","1","2"],
    ["TDIGEST.BYREVRANK","1","1","2"],
    ["TDIGEST.CDF","1","1","2"],
    ["TDIGEST.CREATE","1"],
    ["TDIGEST.CREATE","1","COMPRESSION","1"],
    ["TDIGEST.INFO","1"],
    ["TDIGEST.MAX","1"],
    ["TDIGEST.MIN","1"],
    ["TDIGEST.MERGE","1","2","1","2","COMPRESSION","1000"],
    ["TDIGEST.QUANTILE","1","1","2"],
    ["TDIGEST.RANK","1","1","2"],
    ["TDIGEST.RESET","1"],
    ["TDIGEST.REVRANK","1","1","2"],
    ["TDIGEST.TRIMMED_MEAN","1","1","2"],
    ["TS.ADD","1","1","1"],
    ["TS.ADD","1","1","1","RETENTION","20","ENCODING","COMPRESSED","LABELS","Time","Series"],
    ["TS.CREATE","1"],
    ["TS.CREATE","1","RETENTION","20","ENCODING","UNCOMPRESSED","LABELS","Time","Series"],
    ["TS.ALTER","1","DUPLICATE_POLICY","MIN"],
    ["TS.CREATERULE","1","1","AGGREGATION","AVG","1"],
    ["TS.CREATERULE","1","1","AGGREGATION","AVG","1","1"],
    ["TS.INCRBY","1","1"],
    ["TS.INCRBY","1","1","TIMESTAMP","5"],
    ["TS.DECRBY","1","1"],
    ["TS.DECRBY","1","1","TIMESTAMP","5"],
    ["TS.DEL","1","1","1"],
    ["TS.DELETERULE","1","1"],
    ["TS.GET","1"],
    ["TS.GET","1","LATEST"],
    ["TS.INFO","1"],
    ["TS.INFO","1","DEBUG"],
    ["TS.MADD","1","2","3","1","2","3"],
    ["TS.QUERYINDEX","1"],
    ["TS.REVRANGE","1","1","1"],
    ["TS.REVRANGE","1","1","1","COUNT","10"],
    ["TS.RANGE","1","1","1"],
    ["TS.RANGE","1","1","1","COUNT","10"],
    ["TS.MRANGE","1","1","FILTER","1"],
    ["TS.MRANGE","1","1","COUNT","10","FILTER","1"],
    ["TS.MRANGE","1","1","FILTER","1"],
    ["TS.MREVRANGE","1","1","COUNT","10","FILTER","1"],
    ["TS.MGET","FILTER","1"],
    ["TS.MGET","LATEST","FILTER","1"],
    ["JSON.ARRAPPEND","1","1","1","2"],
    ["JSON.ARRINDEX","1","1","1","2"],
    ["JSON.ARRINDEX","1","1","1","1"],
    ["JSON.ARRINSERT","1","1","1","1","2"],
    ["JSON.ARRLEN","1","1"],
    ["JSON.ARRPOP","1","1","1"],
    ["JSON.ARRTRIM","1","1","0","0"],
    ["JSON.ARRTRIM","1","1","1","0"],
    ["JSON.CLEAR","1","1"],
    ["JSON.DEBUG","MEMORY","1","1"],
    ["JSON.DEL","1","1"],
    ["JSON.FORGET","1","1"],
    ["JSON.GET","1","1","2"],
    ["JSON.GET","1","1","2"],
    ["JSON.MERGE","1","1","1"],
    ["JSON.MSET","1","","1"],
    ["JSON.MSET","mset1","$.a","2","mset3","$","[1]"],
    ["JSON.MGET","1","1","1"],
    ["JSON.NUMINCRBY","1","1","1"],
    ["JSON.OBJKEYS","1","1"],
    ["JSON.OBJLEN","1","1"],
    ["JSON.SET","1","1","2"],
    ["JSON.SET","1","1","2","NX"],
    ["JSON.STRAPPEND","1","1","1"],
    ["JSON.STRLEN","1","1"],
    ["JSON.TOGGLE","1","1"],
    ["JSON.TYPE","1","1"],
    ["SLAVEOF","NO","ONE"],
    ["SLOWLOG","GET","1"],
    ["SLOWLOG","RESET"],
    ["CLUSTER","MYSHARDID"],
    ["MODULE","LOADEX","/","CONFIG","k","v","ARGS","1","2"]
]`
