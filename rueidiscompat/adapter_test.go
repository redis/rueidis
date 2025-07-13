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
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/redis/rueidis"
)

type TimeValue struct {
	time.Time
}

func (t *TimeValue) ScanRedis(s string) (err error) {
	t.Time, err = time.Parse(time.RFC3339Nano, s)
	return
}

func TestAdapter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Adapter Suite")
}

var (
	err                error
	ctx                context.Context
	clientresp2        rueidis.Client
	clientsearchresp2  rueidis.Client
	clusterresp2       rueidis.Client
	clientresp3        rueidis.Client
	clusterresp3       rueidis.Client
	adapterresp2       Cmdable
	adaptersearchresp2 Cmdable
	adaptercluster2    Cmdable
	adapterresp3       Cmdable
	adaptercluster3    Cmdable
)

var _ = BeforeSuite(func() {
	ctx = context.Background()
	clientresp3, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:6378"},
		ClientName:  "rueidis",
	})
	Expect(err).NotTo(HaveOccurred())
	clusterresp3, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:7010"},
		ClientName:  "rueidis",
	})
	Expect(err).NotTo(HaveOccurred())
	adapterresp3 = NewAdapter(clientresp3)
	adaptercluster3 = NewAdapter(clusterresp3)
	clientresp2, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress:  []string{"127.0.0.1:6356"},
		ClientName:   "rueidis",
		DisableCache: true,
	})
	Expect(err).NotTo(HaveOccurred())
	clientsearchresp2, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress:  []string{"127.0.0.1:6381"},
		ClientName:   "rueidis",
		DisableCache: true,
		AlwaysRESP2:  true,
	})
	Expect(err).NotTo(HaveOccurred())
	clusterresp2, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress:  []string{"127.0.0.1:7007"},
		ClientName:   "rueidis",
		DisableCache: true,
	})
	Expect(err).NotTo(HaveOccurred())
	adapterresp2 = NewAdapter(clientresp2)
	adaptersearchresp2 = NewAdapter(clientsearchresp2)
	adaptercluster2 = NewAdapter(clusterresp2)
})

var _ = AfterSuite(func() {
	Expect(adapterresp3.FlushDB(ctx).Err()).NotTo(HaveOccurred())
	Expect(adapterresp3.Quit(ctx).Err()).NotTo(HaveOccurred())
	clientresp3.Close()
	Expect(adapterresp2.FlushDB(ctx).Err()).NotTo(HaveOccurred())
	Expect(adapterresp2.Quit(ctx).Err()).NotTo(HaveOccurred())
	clientresp2.Close()
})

var _ = Describe("RESP3 Commands", func() {
	testAdapter(true)
	testAdapterCache(true)
	testCluster(true)
	testAdapterSearchRESP3()
})

var _ = Describe("RESP2 Commands", func() {
	testAdapter(false)
	testCluster(false)
	testAdapterSearchRESP2()
})

func testCluster(resp3 bool) {
	var adapter Cmdable

	BeforeEach(func() {
		if resp3 {
			adapter = adaptercluster3
		} else {
			adapter = adaptercluster2
		}
	})

	Describe("Cluster", func() {
		if resp3 {
			It("ClusterShards", func() {
				shards, err := adapter.ClusterShards(ctx).Result()
				Expect(err).NotTo(HaveOccurred())
				m := make(map[int64]struct{})
				for _, shard := range shards {
					for _, slot := range shard.Slots {
						for i := slot.Start; i <= slot.End; i++ {
							m[i] = struct{}{}
						}
					}
				}
				Expect(m).To(HaveLen(16384))
			})
			It("ClusterLinks", func() {
				links, err := adapter.ClusterLinks(ctx).Result()
				Expect(err).NotTo(HaveOccurred())

				Expect(links).NotTo(BeEmpty())

				for _, link := range links {
					Expect(link.Direction).NotTo(BeEmpty())
					Expect(link.Node).NotTo(BeEmpty())
					Expect(link.CreateTime).To(BeNumerically(">", 0))
					Expect(link.Events).NotTo(BeEmpty())
					Expect(link.SendBufferAllocated).To(BeNumerically(">=", 0))
					Expect(link.SendBufferUsed).To(BeNumerically(">=", 0))
				}
			})

		}
		It("ClusterSlots", func() {
			slots, err := adapter.ClusterSlots(ctx).Result()
			Expect(err).NotTo(HaveOccurred())
			m := make(map[int64]struct{})
			for _, slot := range slots {
				for i := slot.Start; i <= slot.End; i++ {
					m[i] = struct{}{}
				}
			}
			Expect(m).To(HaveLen(16384))
		})
		It("ClusterNodes", func() {
			nodes, err := adapter.ClusterNodes(ctx).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(strings.Split(strings.TrimSpace(nodes), "\n")).To(HaveLen(3))
		})
		It("ClusterMeet", func() {
			Expect(adapter.ClusterMeet(ctx, "localhost", 8080).Err()).To(MatchError("Invalid node address specified: localhost:8080"))
		})
		It("ClusterForget", func() {
			Expect(adapter.ClusterForget(ctx, "1").Err()).To(MatchError("Unknown node 1"))
		})
		It("ClusterReplicate", func() {
			Expect(adapter.ClusterReplicate(ctx, "1").Err()).To(MatchError("Unknown node 1"))
		})
		It("ClusterInfo", func() {
			info, err := adapter.ClusterInfo(ctx).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(info).NotTo(BeEmpty())
		})
		It("ClusterKeySlot", func() {
			slot, err := adapter.ClusterKeySlot(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(slot).To(Equal(int64(9842)))
		})
		It("ClusterGetKeysInSlot", func() {
			Expect(adapter.Set(ctx, "1", "1", 0).Err()).NotTo(HaveOccurred())
			keys, err := adapter.ClusterGetKeysInSlot(ctx, 9842, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(keys).To(Equal([]string{"1"}))
			kc, err := adapter.ClusterCountKeysInSlot(ctx, 9842).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(kc).To(Equal(int64(1)))
		})
		It("ClusterCountFailureReports", func() {
			Expect(adapter.ClusterCountFailureReports(ctx, "1").Err()).To(MatchError("Unknown node 1"))
		})
		It("ClusterSlaves", func() {
			Expect(adapter.ClusterSlaves(ctx, "1").Err()).To(MatchError("Unknown node 1"))
		})
	})
}

func testAdapter(resp3 bool) {
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

	Describe("server", func() {
		It("should Echo", func() {
			echo := adapter.Echo(ctx, "hello")
			Expect(err).NotTo(HaveOccurred())

			Expect(echo.Err()).NotTo(HaveOccurred())
			Expect(echo.Val()).To(Equal("hello"))
		})

		It("should Ping", func() {
			ping := adapter.Ping(ctx)
			Expect(ping.Err()).NotTo(HaveOccurred())
			Expect(ping.Val()).To(Equal("PONG"))
		})

		It("should Migrate", func() {
			var r *StatusCmd
			if resp3 {
				r = adapter.Migrate(ctx, "127.0.0.1", 6378, "nonkey", 0, 1)
			} else {
				r = adapter.Migrate(ctx, "127.0.0.1", 6356, "nonkey", 0, 1)
			}
			Expect(r.Err()).To(BeNil())
			Expect(r.Val()).To(Equal("NOKEY"))
		})

		It("should Move", func() {
			Expect(adapter.Set(ctx, "movekey", "1", 0).Err()).To(BeNil())
			r := adapter.Move(ctx, "movekey", 1)
			Expect(r.Err()).To(BeNil())
			Expect(r.Val()).To(BeTrue())
		})

		It("should ClientKill", func() {
			r := adapter.ClientKill(ctx, "1.1.1.1:1111")
			Expect(r.Err()).To(MatchError("No such client"))
			Expect(r.Val()).To(Equal(""))
		})

		It("should ClientKillByFilter", func() {
			r := adapter.ClientKillByFilter(ctx, "ID", "12039487")
			Expect(r.Err()).To(BeNil())
			Expect(r.Val()).To(Equal(int64(0)))
		})

		It("should ClientList", func() {
			r := adapter.ClientList(ctx)
			Expect(r.Err()).To(BeNil())
			Expect(r.Val()).NotTo(Equal(""))
		})

		It("should ClientID", func() {
			err := adapter.ClientID(ctx).Err()
			Expect(err).NotTo(HaveOccurred())
			Expect(adapter.ClientID(ctx).Val()).To(BeNumerically(">=", 0))
		})

		It("should ClientGetName", func() {
			r := adapter.ClientGetName(ctx)
			Expect(r.Err()).NotTo(HaveOccurred())
			Expect(r.Val()).To(Equal("rueidis"))
		})

		It("should ConfigGet", func() {
			val, err := adapter.ConfigGet(ctx, "*").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).NotTo(BeEmpty())
		})

		It("should ConfigResetStat", func() {
			r := adapter.ConfigResetStat(ctx)
			Expect(r.Err()).NotTo(HaveOccurred())
			Expect(r.Val()).To(Equal("OK"))
		})

		It("should ConfigSet", func() {
			configGet := adapter.ConfigGet(ctx, "maxmemory")
			Expect(configGet.Err()).NotTo(HaveOccurred())
			Expect(configGet.Val()).To(HaveLen(1))
			Expect(configGet.Val()["maxmemory"]).NotTo(BeEmpty())

			configSet := adapter.ConfigSet(ctx, "maxmemory", configGet.Val()["maxmemory"])
			Expect(configSet.Err()).NotTo(HaveOccurred())
			Expect(configSet.Val()).To(Equal("OK"))
		})

		It("should DBSize", func() {
			size, err := adapter.DBSize(ctx).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(0)))
		})

		It("should Info", func() {
			info := adapter.Info(ctx)
			Expect(info.Err()).NotTo(HaveOccurred())
			Expect(info.Val()).NotTo(Equal(""))
		})

		It("should Info cpu", func() {
			info := adapter.Info(ctx, "cpu")
			Expect(info.Err()).NotTo(HaveOccurred())
			Expect(info.Val()).NotTo(Equal(""))
			Expect(info.Val()).To(ContainSubstring(`used_cpu_sys`))
		})

		It("should LastSave", func() {
			lastSave := adapter.LastSave(ctx)
			Expect(lastSave.Err()).NotTo(HaveOccurred())
			Expect(lastSave.Val()).NotTo(Equal(0))
		})

		It("should Save", func() {
			// workaround for "ERR Background save already in progress"
			Eventually(func() string {
				return adapter.Save(ctx).Val()
			}, "10s").Should(Equal("OK"))
		})

		It("should Time", func() {
			tm, err := adapter.Time(ctx).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(tm).To(BeTemporally("~", time.Now(), 3*time.Second))
		})

		It("should Command", func() {
			cmds, err := adapter.Command(ctx).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(cmds)).To(BeNumerically(">=", 200))

			cmd := cmds["mget"]
			Expect(cmd.Name).To(Equal("mget"))
			Expect(cmd.Arity).To(Equal(int64(-2)))
			Expect(cmd.Flags).To(ContainElement("readonly"))
			Expect(cmd.FirstKeyPos).To(Equal(int64(1)))
			Expect(cmd.LastKeyPos).To(Equal(int64(-1)))
			Expect(cmd.StepCount).To(Equal(int64(1)))

			cmd = cmds["ping"]
			Expect(cmd.Name).To(Equal("ping"))
			Expect(cmd.Arity).To(Equal(int64(-1)))
			Expect(cmd.Flags).To(ContainElement("fast"))
			Expect(cmd.FirstKeyPos).To(Equal(int64(0)))
			Expect(cmd.LastKeyPos).To(Equal(int64(0)))
			Expect(cmd.StepCount).To(Equal(int64(0)))
		})

		if resp3 {
			It("should return all command names", func() {
				cmdList := adapter.CommandList(ctx, FilterBy{})
				Expect(cmdList.Err()).NotTo(HaveOccurred())
				cmdNames := cmdList.Val()

				Expect(cmdNames).NotTo(BeEmpty())

				// Assert that some expected commands are present in the list
				Expect(cmdNames).To(ContainElement("get"))
				Expect(cmdNames).To(ContainElement("set"))
				Expect(cmdNames).To(ContainElement("hset"))
			})

			It("should filter commands by module", func() {
				filter := FilterBy{
					Module: "JSON",
				}
				cmdList := adapter.CommandList(ctx, filter)
				Expect(cmdList.Err()).NotTo(HaveOccurred())
				Expect(cmdList.Val()).To(HaveLen(0))
			})

			It("should filter commands by ACL category", func() {

				filter := FilterBy{
					ACLCat: "admin",
				}

				cmdList := adapter.CommandList(ctx, filter)
				Expect(cmdList.Err()).NotTo(HaveOccurred())
				cmdNames := cmdList.Val()

				// Assert that the returned list only contains commands from the admin ACL category
				Expect(len(cmdNames)).To(BeNumerically(">", 10))
			})

			It("should filter commands by pattern", func() {
				filter := FilterBy{
					Pattern: "*GET*",
				}
				cmdList := adapter.CommandList(ctx, filter)
				Expect(cmdList.Err()).NotTo(HaveOccurred())
				cmdNames := cmdList.Val()

				// Assert that the returned list only contains commands that match the given pattern
				Expect(cmdNames).To(ContainElement("get"))
				Expect(cmdNames).To(ContainElement("getbit"))
				Expect(cmdNames).To(ContainElement("getrange"))
				Expect(cmdNames).NotTo(ContainElement("set"))
			})

			It("Should CommandGetKeys", func() {
				keys, err := adapter.CommandGetKeys(ctx, "MSET", "a", "b", "c", "d", "e", "f").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(keys).To(Equal([]string{"a", "c", "e"}))

				keys, err = adapter.CommandGetKeys(ctx, "EVAL", "not consulted", "3", "key1", "key2", "key3", "arg1", "arg2", "arg3", "argN").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(keys).To(Equal([]string{"key1", "key2", "key3"}))

				keys, err = adapter.CommandGetKeys(ctx, "SORT", "mylist", "ALPHA", "STORE", "outlist").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(keys).To(Equal([]string{"mylist", "outlist"}))

				_, err = adapter.CommandGetKeys(ctx, "FAKECOMMAND", "arg1", "arg2").Result()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Invalid command specified"))
			})

			It("should CommandGetKeysAndFlags", func() {
				keysAndFlags, err := adapter.CommandGetKeysAndFlags(ctx, "LMOVE", "mylist1", "mylist2", "left", "left").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(keysAndFlags).To(Equal([]KeyFlags{
					{
						Key:   "mylist1",
						Flags: []string{"RW", "access", "delete"},
					},
					{
						Key:   "mylist2",
						Flags: []string{"RW", "insert"},
					},
				}))

				_, err = adapter.CommandGetKeysAndFlags(ctx, "FAKECOMMAND", "arg1", "arg2").Result()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Invalid command specified"))
			})
		}
	})

	Describe("debugging", func() {
		It("should MemoryUsage", func() {
			err := adapter.MemoryUsage(ctx, "foo").Err()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue())

			err = adapter.Set(ctx, "foo", "bar", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			n, err := adapter.MemoryUsage(ctx, "foo").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).NotTo(BeZero())

			n, err = adapter.MemoryUsage(ctx, "foo", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).NotTo(BeZero())
		})
	})

	Describe("keys", func() {
		It("should Del", func() {
			err := adapter.Set(ctx, "key1", "Hello", 0).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.Set(ctx, "key2", "World", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			n, err := adapter.Del(ctx, "key1", "key2", "key3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(2)))
		})

		It("should Unlink", func() {
			err := adapter.Set(ctx, "key1", "Hello", 0).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.Set(ctx, "key2", "World", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			n, err := adapter.Unlink(ctx, "key1", "key2", "key3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(2)))
		})

		It("should Dump", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			dump := adapter.Dump(ctx, "key")
			Expect(dump.Err()).NotTo(HaveOccurred())
			Expect(dump.Val()).NotTo(BeEmpty())
		})

		It("should Exists", func() {
			set := adapter.Set(ctx, "key1", "Hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			n, err := adapter.Exists(ctx, "key1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(1)))

			n, err = adapter.Exists(ctx, "key2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(0)))

			n, err = adapter.Exists(ctx, "key1", "key2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(1)))

			n, err = adapter.Exists(ctx, "key1", "key1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(2)))
		})

		It("should Expire", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expire := adapter.Expire(ctx, "key", 10*time.Second)
			Expect(expire.Err()).NotTo(HaveOccurred())
			Expect(expire.Val()).To(Equal(true))

			ttl := adapter.TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(10 * time.Second))

			set = adapter.Set(ctx, "key", "Hello World", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			ttl = adapter.TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(time.Duration(-1)))

			ttl = adapter.TTL(ctx, "nonexistent_key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(time.Duration(-2)))
		})

		if resp3 {
			It("should ExpireNX", func() {
				set := adapter.Set(ctx, "key", "Hello", 0)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				expire := adapter.ExpireNX(ctx, "key", 10*time.Second)
				Expect(expire.Err()).NotTo(HaveOccurred())
				Expect(expire.Val()).To(Equal(true))

				ttl := adapter.TTL(ctx, "key")
				Expect(ttl.Err()).NotTo(HaveOccurred())
				Expect(ttl.Val()).To(Equal(10 * time.Second))

				expire = adapter.ExpireNX(ctx, "key", 20*time.Second)
				Expect(expire.Err()).NotTo(HaveOccurred())
				Expect(expire.Val()).To(Equal(false))
			})

			It("should ExpireXX", func() {
				set := adapter.Set(ctx, "key", "Hello", 0)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				expire := adapter.ExpireXX(ctx, "key", 10*time.Second)
				Expect(expire.Err()).NotTo(HaveOccurred())
				Expect(expire.Val()).To(Equal(false))

				expire = adapter.ExpireNX(ctx, "key", 10*time.Second)
				Expect(expire.Err()).NotTo(HaveOccurred())
				Expect(expire.Val()).To(Equal(true))

				expire = adapter.ExpireXX(ctx, "key", 20*time.Second)
				Expect(expire.Err()).NotTo(HaveOccurred())
				Expect(expire.Val()).To(Equal(true))
			})

			It("should ExpireGT", func() {
				set := adapter.Set(ctx, "key", "Hello", 5*time.Second)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				expire := adapter.ExpireGT(ctx, "key", 10*time.Second)
				Expect(expire.Err()).NotTo(HaveOccurred())
				Expect(expire.Val()).To(Equal(true))

				expire = adapter.ExpireGT(ctx, "key", 5*time.Second)
				Expect(expire.Err()).NotTo(HaveOccurred())
				Expect(expire.Val()).To(Equal(false))
			})

			It("should ExpireLT", func() {
				set := adapter.Set(ctx, "key", "Hello", 10*time.Second)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				expire := adapter.ExpireLT(ctx, "key", 5*time.Second)
				Expect(expire.Err()).NotTo(HaveOccurred())
				Expect(expire.Val()).To(Equal(true))

				expire = adapter.ExpireLT(ctx, "key", 10*time.Second)
				Expect(expire.Err()).NotTo(HaveOccurred())
				Expect(expire.Val()).To(Equal(false))
			})
		}

		if resp3 {
			It("should ExpireAt", func() {
				set := adapter.Set(ctx, "key", "Hello", 0)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				n, err := adapter.Exists(ctx, "key").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(1)))

				// Check correct expiration time is set in the future
				expireAt := time.Now().Add(time.Minute)
				expireAtCmd := adapter.ExpireAt(ctx, "key", expireAt)
				Expect(expireAtCmd.Err()).NotTo(HaveOccurred())
				Expect(expireAtCmd.Val()).To(Equal(true))

				timeCmd := adapter.ExpireTime(ctx, "key")
				Expect(timeCmd.Err()).NotTo(HaveOccurred())
				Expect(timeCmd.Val().Seconds()).To(BeNumerically("==", expireAt.Unix()))

				ptimeCmd := adapter.PExpireTime(ctx, "key")
				Expect(ptimeCmd.Err()).NotTo(HaveOccurred())
				Expect(ptimeCmd.Val().Seconds()).To(BeNumerically("==", expireAt.Unix()))

				// Check correct expiration in the past
				expireAtCmd = adapter.ExpireAt(ctx, "key", time.Now().Add(-time.Hour))
				Expect(expireAtCmd.Err()).NotTo(HaveOccurred())
				Expect(expireAtCmd.Val()).To(Equal(true))

				n, err = adapter.Exists(ctx, "key").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(0)))
			})
		}

		It("should Keys", func() {
			mset := adapter.MSet(ctx, "one", "1", "two", "2", "three", "3", "four", "4")
			Expect(mset.Err()).NotTo(HaveOccurred())
			Expect(mset.Val()).To(Equal("OK"))

			keys := adapter.Keys(ctx, "*o*")
			Expect(keys.Err()).NotTo(HaveOccurred())
			Expect(keys.Val()).To(ConsistOf([]string{"four", "one", "two"}))

			keys = adapter.Keys(ctx, "t??")
			Expect(keys.Err()).NotTo(HaveOccurred())
			Expect(keys.Val()).To(Equal([]string{"two"}))

			keys = adapter.Keys(ctx, "*")
			Expect(keys.Err()).NotTo(HaveOccurred())
			Expect(keys.Val()).To(ConsistOf([]string{"four", "one", "three", "two"}))
		})

		It("should Object", func() {
			start := time.Now()
			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			refCount := adapter.ObjectRefCount(ctx, "key")
			Expect(refCount.Err()).NotTo(HaveOccurred())
			Expect(refCount.Val()).To(Equal(int64(1)))

			err := adapter.ObjectEncoding(ctx, "key").Err()
			Expect(err).NotTo(HaveOccurred())

			idleTime := adapter.ObjectIdleTime(ctx, "key")
			Expect(idleTime.Err()).NotTo(HaveOccurred())

			// Redis returned milliseconds/1000, which may cause ObjectIdleTime to be at a critical value,
			// should be +1s to deal with the critical value problem.
			// if too much time (>1s) is used during command execution, it may also cause the test to fail.
			// so the ObjectIdleTime result should be <=now-start+1s
			// link: https://github.com/redis/redis/blob/5b48d900498c85bbf4772c1d466c214439888115/src/object.c#L1265-L1272
			Expect(idleTime.Val()).To(BeNumerically("<=", time.Now().Sub(start)+time.Second))
		})

		It("should Persist", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expire := adapter.Expire(ctx, "key", 10*time.Second)
			Expect(expire.Err()).NotTo(HaveOccurred())
			Expect(expire.Val()).To(Equal(true))

			ttl := adapter.TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(10 * time.Second))

			persist := adapter.Persist(ctx, "key")
			Expect(persist.Err()).NotTo(HaveOccurred())
			Expect(persist.Val()).To(Equal(true))

			ttl = adapter.TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val() < 0).To(Equal(true))
		})

		It("should PExpire", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expiration := 900 * time.Millisecond
			pexpire := adapter.PExpire(ctx, "key", expiration)
			Expect(pexpire.Err()).NotTo(HaveOccurred())
			Expect(pexpire.Val()).To(Equal(true))

			ttl := adapter.TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(time.Second))

			pttl := adapter.PTTL(ctx, "key")
			Expect(pttl.Err()).NotTo(HaveOccurred())
			Expect(pttl.Val()).To(BeNumerically("~", expiration, 100*time.Millisecond))
		})

		It("should PExpireAt", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expiration := 900 * time.Millisecond
			pexpireat := adapter.PExpireAt(ctx, "key", time.Now().Add(expiration))
			Expect(pexpireat.Err()).NotTo(HaveOccurred())
			Expect(pexpireat.Val()).To(Equal(true))

			ttl := adapter.TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(time.Second))

			pttl := adapter.PTTL(ctx, "key")
			Expect(pttl.Err()).NotTo(HaveOccurred())
			Expect(pttl.Val()).To(BeNumerically("~", expiration, 100*time.Millisecond))
		})

		It("should PTTL", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expiration := time.Second
			expire := adapter.Expire(ctx, "key", expiration)
			Expect(expire.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			pttl := adapter.PTTL(ctx, "key")
			Expect(pttl.Err()).NotTo(HaveOccurred())
			Expect(pttl.Val()).To(BeNumerically("~", expiration, 100*time.Millisecond))
		})

		It("should RandomKey", func() {
			randomKey := adapter.RandomKey(ctx)
			Expect(rueidis.IsRedisNil(randomKey.Err())).To(BeTrue())
			Expect(randomKey.Val()).To(Equal(""))

			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			randomKey = adapter.RandomKey(ctx)
			Expect(randomKey.Err()).NotTo(HaveOccurred())
			Expect(randomKey.Val()).To(Equal("key"))
		})

		It("should Rename", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			status := adapter.Rename(ctx, "key", "key1")
			Expect(status.Err()).NotTo(HaveOccurred())
			Expect(status.Val()).To(Equal("OK"))

			get := adapter.Get(ctx, "key1")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("hello"))
		})

		It("should RenameNX", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			renameNX := adapter.RenameNX(ctx, "key", "key1")
			Expect(renameNX.Err()).NotTo(HaveOccurred())
			Expect(renameNX.Val()).To(Equal(true))

			get := adapter.Get(ctx, "key1")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("hello"))
		})

		It("should Restore", func() {
			err := adapter.Set(ctx, "key", "hello", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			dump := adapter.Dump(ctx, "key")
			Expect(dump.Err()).NotTo(HaveOccurred())

			err = adapter.Del(ctx, "key").Err()
			Expect(err).NotTo(HaveOccurred())

			restore, err := adapter.Restore(ctx, "key", 0, dump.Val()).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(restore).To(Equal("OK"))

			type_, err := adapter.Type(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(type_).To(Equal("string"))

			val, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello"))
		})

		It("should RestoreReplace", func() {
			err := adapter.Set(ctx, "key", "hello", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			dump := adapter.Dump(ctx, "key")
			Expect(dump.Err()).NotTo(HaveOccurred())

			restore, err := adapter.RestoreReplace(ctx, "key", 0, dump.Val()).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(restore).To(Equal("OK"))

			type_, err := adapter.Type(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(type_).To(Equal("string"))

			val, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello"))
		})

		It("should Sort", func() {
			size, err := adapter.LPush(ctx, "list", "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list", "3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(3)))

			els, err := adapter.Sort(ctx, "list", Sort{
				Offset: 0,
				Count:  2,
				Order:  "ASC",
				Alpha:  true,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(els).To(Equal([]string{"1", "2"}))
		})

		It("should Sort By", func() {
			size, err := adapter.LPush(ctx, "list_by", "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list_by", "3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list_by", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(3)))

			els, err := adapter.Sort(ctx, "list_by", Sort{
				Offset: 0,
				Count:  2,
				By:     "nosort",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(els).To(Equal([]string{"2", "3"}))
		})

		if resp3 {
			It("should Sort", func() {
				size, err := adapter.LPush(ctx, "list", "1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(size).To(Equal(int64(1)))

				size, err = adapter.LPush(ctx, "list", "3").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(size).To(Equal(int64(2)))

				size, err = adapter.LPush(ctx, "list", "2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(size).To(Equal(int64(3)))

				els, err := adapter.SortRO(ctx, "list", Sort{
					Offset: 0,
					Count:  2,
					Order:  "ASC",
					Alpha:  true,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(els).To(Equal([]string{"1", "2"}))
			})

			It("should Sort By", func() {
				size, err := adapter.LPush(ctx, "list_by", "1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(size).To(Equal(int64(1)))

				size, err = adapter.LPush(ctx, "list_by", "3").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(size).To(Equal(int64(2)))

				size, err = adapter.LPush(ctx, "list_by", "2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(size).To(Equal(int64(3)))

				els, err := adapter.SortRO(ctx, "list_by", Sort{
					Offset: 0,
					Count:  2,
					By:     "nosort",
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(els).To(Equal([]string{"2", "3"}))
			})
		}

		It("should Sort Panic", func() {
			Expect(func() {
				adapter.Sort(ctx, "list", Sort{Order: "PANIC"})
			}).To(Panic())
		})

		It("should Sort and Get", func() {
			size, err := adapter.LPush(ctx, "list", "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list", "3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(3)))

			err = adapter.Set(ctx, "object_2", "value2", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			{
				els, err := adapter.Sort(ctx, "list", Sort{
					Get: []string{"object_*"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(els).To(Equal([]string{"", "value2", ""}))
			}

			{
				els, err := adapter.SortInterfaces(ctx, "list", Sort{
					Get: []string{"object_*"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(els).To(Equal([]any{nil, "value2", nil}))
			}
		})

		It("should Sort and Store", func() {
			size, err := adapter.LPush(ctx, "list", "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list", "3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(3)))

			n, err := adapter.SortStore(ctx, "list", "list2", Sort{
				Offset: 0,
				Count:  2,
				Order:  "ASC",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(2)))

			els, err := adapter.LRange(ctx, "list2", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(els).To(Equal([]string{"1", "2"}))
		})

		It("should Touch", func() {
			set1 := adapter.Set(ctx, "touch1", "hello", 0)
			Expect(set1.Err()).NotTo(HaveOccurred())
			Expect(set1.Val()).To(Equal("OK"))

			set2 := adapter.Set(ctx, "touch2", "hello", 0)
			Expect(set2.Err()).NotTo(HaveOccurred())
			Expect(set2.Val()).To(Equal("OK"))

			touch := adapter.Touch(ctx, "touch1", "touch2", "touch3")
			Expect(touch.Err()).NotTo(HaveOccurred())
			Expect(touch.Val()).To(Equal(int64(2)))
		})

		It("should TTL", func() {
			ttl := adapter.TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val() < 0).To(Equal(true))

			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expire := adapter.Expire(ctx, "key", 60*time.Second)
			Expect(expire.Err()).NotTo(HaveOccurred())
			Expect(expire.Val()).To(Equal(true))

			ttl = adapter.TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(60 * time.Second))
		})

		It("should Type", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			type_ := adapter.Type(ctx, "key")
			Expect(type_.Err()).NotTo(HaveOccurred())
			Expect(type_.Val()).To(Equal("string"))
		})
	})

	Describe("scanning", func() {
		It("should Scan", func() {
			for i := 0; i < 1000; i++ {
				set := adapter.Set(ctx, fmt.Sprintf("key%d", i), "hello", 0)
				Expect(set.Err()).NotTo(HaveOccurred())
			}

			keys, cursor, err := adapter.Scan(ctx, 0, "key*", 100).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(keys).NotTo(BeEmpty())
			Expect(cursor).NotTo(BeZero())
		})

		if resp3 {
			It("should ScanType", func() {
				for i := 0; i < 1000; i++ {
					set := adapter.Set(ctx, fmt.Sprintf("key%d", i), "hello", 0)
					Expect(set.Err()).NotTo(HaveOccurred())
				}

				keys, cursor, err := adapter.ScanType(ctx, 0, "key*", 100, "string").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(keys).NotTo(BeEmpty())
				Expect(cursor).NotTo(BeZero())
			})
		}

		It("should SScan", func() {
			for i := 0; i < 1000; i++ {
				sadd := adapter.SAdd(ctx, "myset", fmt.Sprintf("member%d", i))
				Expect(sadd.Err()).NotTo(HaveOccurred())
			}

			keys, cursor, err := adapter.SScan(ctx, "myset", 0, "member*", 100).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(keys).NotTo(BeEmpty())
			Expect(cursor).NotTo(BeZero())
		})

		It("should HScan", func() {
			for i := 0; i < 1000; i++ {
				sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
				Expect(sadd.Err()).NotTo(HaveOccurred())
			}

			keys, cursor, err := adapter.HScan(ctx, "myhash", 0, "key*", 100).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(keys).NotTo(BeEmpty())
			Expect(cursor).NotTo(BeZero())
		})

		if resp3 {
			It("should HScan without values", func() {
				for i := 0; i < 1000; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				keys, cursor, err := adapter.HScanNoValues(ctx, "myhash", 0, "", 0).Result()
				Expect(err).NotTo(HaveOccurred())
				// If we don't get at least two items back, it's really strange.
				Expect(cursor).To(BeNumerically(">=", 2))
				Expect(len(keys)).To(BeNumerically(">=", 2))
				Expect(keys[0]).To(HavePrefix("key"))
				Expect(keys[1]).To(HavePrefix("key"))
				Expect(keys).NotTo(BeEmpty())
				Expect(cursor).NotTo(BeZero())
			})
		}

		It("should ZScan", func() {
			for i := 0; i < 1000; i++ {
				err := adapter.ZAdd(ctx, "myset", Z{
					Score:  float64(i),
					Member: fmt.Sprintf("member%d", i),
				}).Err()
				Expect(err).NotTo(HaveOccurred())
			}

			keys, cursor, err := adapter.ZScan(ctx, "myset", 0, "member*", 100).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(keys).NotTo(BeEmpty())
			Expect(cursor).NotTo(BeZero())
		})
	})

	Describe("strings", func() {
		It("should Append", func() {
			n, err := adapter.Exists(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(0)))

			append := adapter.Append(ctx, "key", "Hello")
			Expect(append.Err()).NotTo(HaveOccurred())
			Expect(append.Val()).To(Equal(int64(5)))

			append = adapter.Append(ctx, "key", " World")
			Expect(append.Err()).NotTo(HaveOccurred())
			Expect(append.Val()).To(Equal(int64(11)))

			get := adapter.Get(ctx, "key")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("Hello World"))
		})

		It("should BitCount", func() {
			set := adapter.Set(ctx, "key", "foobar", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			bitCount := adapter.BitCount(ctx, "key", nil)
			Expect(bitCount.Err()).NotTo(HaveOccurred())
			Expect(bitCount.Val()).To(Equal(int64(26)))

			bitCount = adapter.BitCount(ctx, "key", &BitCount{
				Start: 0,
				End:   0,
			})
			Expect(bitCount.Err()).NotTo(HaveOccurred())
			Expect(bitCount.Val()).To(Equal(int64(4)))

			bitCount = adapter.BitCount(ctx, "key", &BitCount{
				Start: 1,
				End:   1,
			})
			Expect(bitCount.Err()).NotTo(HaveOccurred())
			Expect(bitCount.Val()).To(Equal(int64(6)))

			if resp3 {
				bitCount = adapter.BitCount(ctx, "key", &BitCount{
					Start: 1,
					End:   1,
					Unit:  "BYTE",
				})
				Expect(bitCount.Err()).NotTo(HaveOccurred())
				Expect(bitCount.Val()).To(Equal(int64(6)))

				bitCount = adapter.BitCount(ctx, "key", &BitCount{
					Start: 1,
					End:   1,
					Unit:  "BIT",
				})
				Expect(bitCount.Err()).NotTo(HaveOccurred())
				Expect(bitCount.Val()).To(Equal(int64(1)))
			}
		})

		It("should BitOpAnd", func() {
			set := adapter.Set(ctx, "key1", "1", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			set = adapter.Set(ctx, "key2", "0", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			bitOpAnd := adapter.BitOpAnd(ctx, "dest", "key1", "key2")
			Expect(bitOpAnd.Err()).NotTo(HaveOccurred())
			Expect(bitOpAnd.Val()).To(Equal(int64(1)))

			get := adapter.Get(ctx, "dest")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("0"))
		})

		It("should BitOpOr", func() {
			set := adapter.Set(ctx, "key1", "1", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			set = adapter.Set(ctx, "key2", "0", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			bitOpOr := adapter.BitOpOr(ctx, "dest", "key1", "key2")
			Expect(bitOpOr.Err()).NotTo(HaveOccurred())
			Expect(bitOpOr.Val()).To(Equal(int64(1)))

			get := adapter.Get(ctx, "dest")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("1"))
		})

		It("should BitOpXor", func() {
			set := adapter.Set(ctx, "key1", "\xff", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			set = adapter.Set(ctx, "key2", "\x0f", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			bitOpXor := adapter.BitOpXor(ctx, "dest", "key1", "key2")
			Expect(bitOpXor.Err()).NotTo(HaveOccurred())
			Expect(bitOpXor.Val()).To(Equal(int64(1)))

			get := adapter.Get(ctx, "dest")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("\xf0"))
		})

		It("should BitOpNot", func() {
			set := adapter.Set(ctx, "key1", "\x00", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			bitOpNot := adapter.BitOpNot(ctx, "dest", "key1")
			Expect(bitOpNot.Err()).NotTo(HaveOccurred())
			Expect(bitOpNot.Val()).To(Equal(int64(1)))

			get := adapter.Get(ctx, "dest")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("\xff"))
		})

		It("BitPos should panic", func() {
			Expect(func() {
				adapter.BitPos(ctx, "mykey", 0, 0, 0, 0)
			}).To(Panic())
		})

		It("should BitPos", func() {
			err := adapter.Set(ctx, "mykey", "\xff\xf0\x00", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			pos, err := adapter.BitPos(ctx, "mykey", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(12)))

			pos, err = adapter.BitPos(ctx, "mykey", 1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(0)))

			pos, err = adapter.BitPos(ctx, "mykey", 0, 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(16)))

			pos, err = adapter.BitPos(ctx, "mykey", 1, 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(-1)))

			pos, err = adapter.BitPos(ctx, "mykey", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(16)))

			pos, err = adapter.BitPos(ctx, "mykey", 1, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(-1)))

			pos, err = adapter.BitPos(ctx, "mykey", 0, 2, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(-1)))

			pos, err = adapter.BitPos(ctx, "mykey", 0, 0, -3).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(-1)))

			pos, err = adapter.BitPos(ctx, "mykey", 0, 0, 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(-1)))
		})

		if resp3 {
			It("should BitPosSpan", func() {
				err := adapter.Set(ctx, "mykey", "\x00\xff\x00", 0).Err()
				Expect(err).NotTo(HaveOccurred())

				pos, err := adapter.BitPosSpan(ctx, "mykey", 0, 1, 3, "byte").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(pos).To(Equal(int64(16)))

				pos, err = adapter.BitPosSpan(ctx, "mykey", 0, 1, 3, "bit").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(pos).To(Equal(int64(1)))
			})
		}

		It("should BitField", func() {
			nn, err := adapter.BitField(ctx, "mykey", "INCRBY", "i5", 100, 1, "GET", "u4", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(nn).To(Equal([]int64{1, 0}))
		})

		if resp3 {
			It("should BitFieldRO", func() {
				nn, err := adapter.BitField(ctx, "mykey", "SET", "u8", 8, 255).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(nn).To(Equal([]int64{0}))

				nn, err = adapter.BitFieldRO(ctx, "mykey", "u8", 0).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(nn).To(Equal([]int64{0}))

				nn, err = adapter.BitFieldRO(ctx, "mykey", "u8", 0, "u4", 8, "u4", 12, "u4", 13).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(nn).To(Equal([]int64{0, 15, 15, 14}))
			})
		}

		It("should Decr", func() {
			set := adapter.Set(ctx, "key", "10", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			decr := adapter.Decr(ctx, "key")
			Expect(decr.Err()).NotTo(HaveOccurred())
			Expect(decr.Val()).To(Equal(int64(9)))

			set = adapter.Set(ctx, "key", "234293482390480948029348230948", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			decr = adapter.Decr(ctx, "key")
			Expect(decr.Err()).To(MatchError("value is not an integer or out of range"))
			Expect(decr.Val()).To(Equal(int64(0)))
		})

		It("should DecrBy", func() {
			set := adapter.Set(ctx, "key", "10", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			decrBy := adapter.DecrBy(ctx, "key", 5)
			Expect(decrBy.Err()).NotTo(HaveOccurred())
			Expect(decrBy.Val()).To(Equal(int64(5)))
		})

		It("should Get", func() {
			get := adapter.Get(ctx, "_")
			Expect(rueidis.IsRedisNil(get.Err())).To(BeTrue())
			Expect(get.Val()).To(Equal(""))

			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			get = adapter.Get(ctx, "key")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("hello"))
		})

		It("should GetBit", func() {
			setBit := adapter.SetBit(ctx, "key", 7, 1)
			Expect(setBit.Err()).NotTo(HaveOccurred())
			Expect(setBit.Val()).To(Equal(int64(0)))

			getBit := adapter.GetBit(ctx, "key", 0)
			Expect(getBit.Err()).NotTo(HaveOccurred())
			Expect(getBit.Val()).To(Equal(int64(0)))

			getBit = adapter.GetBit(ctx, "key", 7)
			Expect(getBit.Err()).NotTo(HaveOccurred())
			Expect(getBit.Val()).To(Equal(int64(1)))

			getBit = adapter.GetBit(ctx, "key", 100)
			Expect(getBit.Err()).NotTo(HaveOccurred())
			Expect(getBit.Val()).To(Equal(int64(0)))
		})

		It("should GetRange", func() {
			set := adapter.Set(ctx, "key", "This is a string", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			getRange := adapter.GetRange(ctx, "key", 0, 3)
			Expect(getRange.Err()).NotTo(HaveOccurred())
			Expect(getRange.Val()).To(Equal("This"))

			getRange = adapter.GetRange(ctx, "key", -3, -1)
			Expect(getRange.Err()).NotTo(HaveOccurred())
			Expect(getRange.Val()).To(Equal("ing"))

			getRange = adapter.GetRange(ctx, "key", 0, -1)
			Expect(getRange.Err()).NotTo(HaveOccurred())
			Expect(getRange.Val()).To(Equal("This is a string"))

			getRange = adapter.GetRange(ctx, "key", 10, 100)
			Expect(getRange.Err()).NotTo(HaveOccurred())
			Expect(getRange.Val()).To(Equal("string"))
		})

		It("should GetSet", func() {
			incr := adapter.Incr(ctx, "key")
			Expect(incr.Err()).NotTo(HaveOccurred())
			Expect(incr.Val()).To(Equal(int64(1)))

			getSet := adapter.GetSet(ctx, "key", "0")
			Expect(getSet.Err()).NotTo(HaveOccurred())
			Expect(getSet.Val()).To(Equal("1"))

			get := adapter.Get(ctx, "key")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("0"))
		})

		if resp3 {
			It("should GetEX", func() {
				set := adapter.Set(ctx, "key", "value", 100*time.Second)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				ttl := adapter.TTL(ctx, "key")
				Expect(ttl.Err()).NotTo(HaveOccurred())
				Expect(ttl.Val()).To(BeNumerically("~", 100*time.Second, 3*time.Second))

				getEX := adapter.GetEx(ctx, "key", 200*time.Second)
				Expect(getEX.Err()).NotTo(HaveOccurred())
				Expect(getEX.Val()).To(Equal("value"))

				ttl = adapter.TTL(ctx, "key")
				Expect(ttl.Err()).NotTo(HaveOccurred())
				Expect(ttl.Val()).To(BeNumerically("~", 200*time.Second, 3*time.Second))
			})

			It("should GetEX 2", func() {
				set := adapter.Set(ctx, "key", "value", 0)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				getEX := adapter.GetEx(ctx, "key", 0)
				Expect(getEX.Err()).NotTo(HaveOccurred())
				Expect(getEX.Val()).To(Equal("value"))

				ttl := adapter.TTL(ctx, "key")
				Expect(ttl.Err()).NotTo(HaveOccurred())
				Expect(ttl.Val()).To(Equal(time.Duration(-1)))

				getEX = adapter.GetEx(ctx, "key", 100*time.Millisecond)
				Expect(getEX.Err()).NotTo(HaveOccurred())
				Expect(getEX.Val()).To(Equal("value"))

				ttl = adapter.PTTL(ctx, "key")
				Expect(ttl.Err()).NotTo(HaveOccurred())
				Expect(ttl.Val()).To(BeNumerically("~", 100*time.Millisecond, 10*time.Millisecond))
			})

			It("should GetDel", func() {
				set := adapter.Set(ctx, "key", "value", 0)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				getDel := adapter.GetDel(ctx, "key")
				Expect(getDel.Err()).NotTo(HaveOccurred())
				Expect(getDel.Val()).To(Equal("value"))

				get := adapter.Get(ctx, "key")
				Expect(rueidis.IsRedisNil(get.Err())).To(BeTrue())
			})
		}

		It("should Incr", func() {
			set := adapter.Set(ctx, "key", "10", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			incr := adapter.Incr(ctx, "key")
			Expect(incr.Err()).NotTo(HaveOccurred())
			Expect(incr.Val()).To(Equal(int64(11)))

			get := adapter.Get(ctx, "key")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("11"))
		})

		It("should IncrBy", func() {
			set := adapter.Set(ctx, "key", "10", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			incrBy := adapter.IncrBy(ctx, "key", 5)
			Expect(incrBy.Err()).NotTo(HaveOccurred())
			Expect(incrBy.Val()).To(Equal(int64(15)))
		})

		It("should IncrByFloat", func() {
			set := adapter.Set(ctx, "key", "10.50", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			incrByFloat := adapter.IncrByFloat(ctx, "key", 0.1)
			Expect(incrByFloat.Err()).NotTo(HaveOccurred())
			Expect(incrByFloat.Val()).To(Equal(10.6))

			set = adapter.Set(ctx, "key", "5.0e3", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			incrByFloat = adapter.IncrByFloat(ctx, "key", 2.0e2)
			Expect(incrByFloat.Err()).NotTo(HaveOccurred())
			Expect(incrByFloat.Val()).To(Equal(float64(5200)))
		})

		It("should IncrByFloatOverflow", func() {
			incrByFloat := adapter.IncrByFloat(ctx, "key", 996945661)
			Expect(incrByFloat.Err()).NotTo(HaveOccurred())
			Expect(incrByFloat.Val()).To(Equal(float64(996945661)))
		})

		It("should MSetMGet", func() {
			mSet := adapter.MSet(ctx, "key1", "hello1", "key2", "hello2")
			Expect(mSet.Err()).NotTo(HaveOccurred())
			Expect(mSet.Val()).To(Equal("OK"))

			mGet := adapter.MGet(ctx, "key1", "key2", "_")
			Expect(mGet.Err()).NotTo(HaveOccurred())
			Expect(mGet.Val()).To(Equal([]interface{}{"hello1", "hello2", nil}))

			// MSet struct
			type set struct {
				Set1 string                 `redis:"set1"`
				Set2 int16                  `redis:"set2"`
				Set3 time.Duration          `redis:"set3"`
				Set4 interface{}            `redis:"set4"`
				Set5 map[string]interface{} `redis:"-"`
			}
			mSet = adapter.MSet(ctx, &set{
				Set1: "val1",
				Set2: 1024,
				Set3: 2 * time.Millisecond,
				Set4: nil,
				Set5: map[string]interface{}{"k1": 1},
			})
			Expect(mSet.Err()).NotTo(HaveOccurred())
			Expect(mSet.Val()).To(Equal("OK"))

			mGet = adapter.MGet(ctx, "set1", "set2", "set3", "set4")
			Expect(mGet.Err()).NotTo(HaveOccurred())
			Expect(mGet.Val()).To(Equal([]interface{}{
				"val1",
				"1024",
				strconv.Itoa(int(2 * time.Millisecond.Nanoseconds())),
				"",
			}))
		})

		It("should scan Mget", func() {
			now := time.Now()

			err := adapter.MSet(ctx, "key1", "hello1", "key2", 123, "time", now.Format(time.RFC3339Nano)).Err()
			Expect(err).NotTo(HaveOccurred())

			res := adapter.MGet(ctx, "key1", "key2", "_", "time")
			Expect(res.Err()).NotTo(HaveOccurred())

			type data struct {
				Key1 string    `redis:"key1"`
				Key2 int       `redis:"key2"`
				Time TimeValue `redis:"time"`
			}
			var d data
			Expect(res.Scan(&d)).NotTo(HaveOccurred())
			Expect(d.Time.UnixNano()).To(Equal(now.UnixNano()))
			d.Time.Time = time.Time{}
			Expect(d).To(Equal(data{
				Key1: "hello1",
				Key2: 123,
				Time: TimeValue{Time: time.Time{}},
			}))
		})

		It("should MSetNX", func() {
			mSetNX := adapter.MSetNX(ctx, "key1", "hello1", "key2", "hello2")
			Expect(mSetNX.Err()).NotTo(HaveOccurred())
			Expect(mSetNX.Val()).To(Equal(true))

			mSetNX = adapter.MSetNX(ctx, "key2", "hello1", "key3", "hello2")
			Expect(mSetNX.Err()).NotTo(HaveOccurred())
			Expect(mSetNX.Val()).To(Equal(false))

			// set struct
			// MSet struct
			type set struct {
				Set1 string                 `redis:"set1"`
				Set2 int16                  `redis:"set2"`
				Set3 time.Duration          `redis:"set3"`
				Set4 interface{}            `redis:"set4"`
				Set5 map[string]interface{} `redis:"-"`
			}
			mSetNX = adapter.MSetNX(ctx, &set{
				Set1: "val1",
				Set2: 1024,
				Set3: 2 * time.Millisecond,
				Set4: nil,
				Set5: map[string]interface{}{"k1": 1},
			})
			Expect(mSetNX.Err()).NotTo(HaveOccurred())
			Expect(mSetNX.Val()).To(Equal(true))
		})
		It("SetWithArgs should panic wrong mode", func() {
			Expect(func() {
				adapter.SetArgs(ctx, "key", "hello", SetArgs{Mode: "ANY"})
			}).To(Panic())
		})

		It("should SetWithArgs with TTL", func() {
			args := SetArgs{
				TTL: 500 * time.Millisecond,
			}
			err := adapter.SetArgs(ctx, "key", "hello", args).Err()
			Expect(err).NotTo(HaveOccurred())

			val, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello"))

			Eventually(func() bool {
				return rueidis.IsRedisNil(adapter.Get(ctx, "key").Err())
			}, "2s", "100ms").Should(BeTrue())
		})

		if resp3 {
			It("should SetWithArgs with expiration date", func() {
				expireAt := time.Now().AddDate(1, 1, 1)
				args := SetArgs{
					ExpireAt: expireAt,
				}
				err := adapter.SetArgs(ctx, "key", "hello", args).Err()
				Expect(err).NotTo(HaveOccurred())

				val, err := adapter.Get(ctx, "key").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal("hello"))

				// check the key has an expiration date
				// (so a TTL value different of -1)
				ttl := adapter.TTL(ctx, "key")
				Expect(ttl.Err()).NotTo(HaveOccurred())
				Expect(ttl.Val()).ToNot(Equal(-1))
			})

			It("should SetWithArgs with negative expiration date", func() {
				args := SetArgs{
					ExpireAt: time.Now().AddDate(-3, 1, 1),
				}
				// redis accepts a timestamp less than the current date
				// but returns nil when trying to get the key
				err := adapter.SetArgs(ctx, "key", "hello", args).Err()
				Expect(err).NotTo(HaveOccurred())

				val, err := adapter.Get(ctx, "key").Result()
				Expect(rueidis.IsRedisNil(err)).To(BeTrue())
				Expect(val).To(Equal(""))
			})

			It("should SetWithArgs with keepttl", func() {
				// Set with ttl
				argsWithTTL := SetArgs{
					TTL: 5 * time.Second,
				}
				set := adapter.SetArgs(ctx, "key", "hello", argsWithTTL)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Result()).To(Equal("OK"))

				// Set with keepttl
				argsWithKeepTTL := SetArgs{
					KeepTTL: true,
				}
				set = adapter.SetArgs(ctx, "key", "hello", argsWithKeepTTL)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Result()).To(Equal("OK"))

				ttl := adapter.TTL(ctx, "key")
				Expect(ttl.Err()).NotTo(HaveOccurred())
				// set keepttl will Retain the ttl associated with the key
				Expect(ttl.Val().Nanoseconds()).NotTo(Equal(-1))
			})
		}

		It("should SetWithArgs with NX mode and key exists", func() {
			err := adapter.Set(ctx, "key", "hello", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			args := SetArgs{
				Mode: "nx",
			}
			val, err := adapter.SetArgs(ctx, "key", "hello", args).Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue())
			Expect(val).To(Equal(""))
		})

		It("should SetWithArgs with NX mode and key does not exist", func() {
			args := SetArgs{
				Mode: "nx",
			}
			val, err := adapter.SetArgs(ctx, "key", "hello", args).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("OK"))
		})

		It("should SetWithArgs with expiration, NX mode, and key does not exist", func() {
			args := SetArgs{
				TTL:  500 * time.Millisecond,
				Mode: "nx",
			}
			val, err := adapter.SetArgs(ctx, "key", "hello", args).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("OK"))

			Eventually(func() bool {
				return rueidis.IsRedisNil(adapter.Get(ctx, "key").Err())
			}, "1s", "100ms").Should(BeTrue())
		})

		It("should SetWithArgs with expiration, NX mode, and key exists", func() {
			e := adapter.Set(ctx, "key", "hello", 0)
			Expect(e.Err()).NotTo(HaveOccurred())

			args := SetArgs{
				TTL:  500 * time.Millisecond,
				Mode: "nx",
			}
			val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue())
			Expect(val).To(Equal(""))
		})

		It("should SetWithArgs with XX mode and key does not exist", func() {
			args := SetArgs{
				Mode: "xx",
			}
			val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue())
			Expect(val).To(Equal(""))
		})

		It("should SetWithArgs with XX mode and key exists", func() {
			e := adapter.Set(ctx, "key", "hello", 0).Err()
			Expect(e).NotTo(HaveOccurred())

			args := SetArgs{
				Mode: "xx",
			}
			val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("OK"))
		})

		if resp3 {
			It("should SetWithArgs with XX mode and GET option, and key exists", func() {
				e := adapter.Set(ctx, "key", "hello", 0).Err()
				Expect(e).NotTo(HaveOccurred())

				args := SetArgs{
					Mode: "xx",
					Get:  true,
				}
				val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal("hello"))
			})

			It("should SetWithArgs with XX mode and GET option, and key does not exist", func() {
				args := SetArgs{
					Mode: "xx",
					Get:  true,
				}

				val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
				Expect(rueidis.IsRedisNil(err)).To(BeTrue())
				Expect(val).To(Equal(""))
			})

			It("should SetWithArgs with expiration, XX mode, GET option, and key does not exist", func() {
				args := SetArgs{
					TTL:  500 * time.Millisecond,
					Mode: "xx",
					Get:  true,
				}

				val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
				Expect(rueidis.IsRedisNil(err)).To(BeTrue())
				Expect(val).To(Equal(""))
			})

			It("should SetWithArgs with expiration, XX mode, GET option, and key exists", func() {
				e := adapter.Set(ctx, "key", "hello", 0)
				Expect(e.Err()).NotTo(HaveOccurred())

				args := SetArgs{
					TTL:  500 * time.Millisecond,
					Mode: "xx",
					Get:  true,
				}

				val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal("hello"))

				Eventually(func() bool {
					return rueidis.IsRedisNil(adapter.Get(ctx, "key").Err())
				}, "1s", "100ms").Should(BeTrue())
			})

			It("should SetWithArgs with Get and key does not exist yet", func() {
				args := SetArgs{
					Get: true,
				}

				val, err := adapter.SetArgs(ctx, "key", "hello", args).Result()
				Expect(rueidis.IsRedisNil(err)).To(BeTrue())
				Expect(val).To(Equal(""))
			})

			It("should SetWithArgs with Get and key exists", func() {
				e := adapter.Set(ctx, "key", "hello", 0)
				Expect(e.Err()).NotTo(HaveOccurred())

				args := SetArgs{
					Get: true,
				}

				val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal("hello"))
			})

			It("should Set with keepttl", func() {
				// set with ttl
				set := adapter.Set(ctx, "key", "hello", 5*time.Second)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				// set with keepttl
				set = adapter.Set(ctx, "key", "hello1", KeepTTL)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				ttl := adapter.TTL(ctx, "key")
				Expect(ttl.Err()).NotTo(HaveOccurred())
				// set keepttl will Retain the ttl associated with the key
				Expect(ttl.Val().Nanoseconds()).NotTo(Equal(-1))
			})
		}

		It("should Set with expiration", func() {
			err := adapter.Set(ctx, "key", "hello", 100*time.Millisecond).Err()
			Expect(err).NotTo(HaveOccurred())

			val, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello"))

			Eventually(func() bool {
				return rueidis.IsRedisNil(adapter.Get(ctx, "key").Err())
			}, "1s", "100ms").Should(BeTrue())
		})

		It("should SetGet", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			get := adapter.Get(ctx, "key")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("hello"))
		})

		It("should SetEX", func() {
			err := adapter.SetEX(ctx, "key", "hello", 1*time.Second).Err()
			Expect(err).NotTo(HaveOccurred())

			val, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello"))

			Eventually(func() bool {
				return rueidis.IsRedisNil(adapter.Get(ctx, "foo").Err())
			}, "2s", "100ms").Should(BeTrue())
		})

		It("should SetNX", func() {
			setNX := adapter.SetNX(ctx, "key", "hello", 0)
			Expect(setNX.Err()).NotTo(HaveOccurred())
			Expect(setNX.Val()).To(Equal(true))

			setNX = adapter.SetNX(ctx, "key", "hello2", 0)
			Expect(setNX.Err()).NotTo(HaveOccurred())
			Expect(setNX.Val()).To(Equal(false))

			get := adapter.Get(ctx, "key")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("hello"))
		})

		It("should SetNX with expiration", func() {
			isSet, err := adapter.SetNX(ctx, "key", "hello", time.Second).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSet).To(Equal(true))

			isSet, err = adapter.SetNX(ctx, "key", "hello2", time.Second).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSet).To(Equal(false))

			val, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello"))
		})

		It("should SetNX with expiration 2", func() {
			isSet, err := adapter.SetNX(ctx, "key", "hello", 100*time.Millisecond).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSet).To(Equal(true))

			isSet, err = adapter.SetNX(ctx, "key", "hello2", 100*time.Millisecond).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSet).To(Equal(false))

			val, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello"))
		})

		if resp3 {
			It("should SetNX with keepttl", func() {
				isSet, err := adapter.SetNX(ctx, "key", "hello1", KeepTTL).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(isSet).To(Equal(true))

				ttl := adapter.TTL(ctx, "key")
				Expect(ttl.Err()).NotTo(HaveOccurred())
				Expect(ttl.Val().Nanoseconds()).To(Equal(int64(-1)))
			})
		}

		It("should SetXX", func() {
			isSet, err := adapter.SetXX(ctx, "key", "hello2", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSet).To(Equal(false))

			err = adapter.Set(ctx, "key", "hello", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			isSet, err = adapter.SetXX(ctx, "key", "hello2", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSet).To(Equal(true))

			val, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello2"))
		})

		It("should SetXX with expiration", func() {
			isSet, err := adapter.SetXX(ctx, "key", "hello2", time.Second).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSet).To(Equal(false))

			err = adapter.Set(ctx, "key", "hello", time.Second).Err()
			Expect(err).NotTo(HaveOccurred())

			isSet, err = adapter.SetXX(ctx, "key", "hello2", time.Second).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSet).To(Equal(true))

			val, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello2"))
		})

		It("should SetXX with expiration", func() {
			isSet, err := adapter.SetXX(ctx, "key", "hello2", 100*time.Millisecond).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSet).To(Equal(false))

			err = adapter.Set(ctx, "key", "hello", time.Second).Err()
			Expect(err).NotTo(HaveOccurred())

			isSet, err = adapter.SetXX(ctx, "key", "hello2", 100*time.Millisecond).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(isSet).To(Equal(true))

			val, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello2"))
		})

		if resp3 {
			It("should SetXX with keepttl", func() {
				isSet, err := adapter.SetXX(ctx, "key", "hello2", time.Second).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(isSet).To(Equal(false))

				err = adapter.Set(ctx, "key", "hello", time.Second).Err()
				Expect(err).NotTo(HaveOccurred())

				isSet, err = adapter.SetXX(ctx, "key", "hello2", 5*time.Second).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(isSet).To(Equal(true))

				isSet, err = adapter.SetXX(ctx, "key", "hello3", KeepTTL).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(isSet).To(Equal(true))

				val, err := adapter.Get(ctx, "key").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal("hello3"))

				// set keepttl will Retain the ttl associated with the key
				ttl, err := adapter.TTL(ctx, "key").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(ttl).NotTo(Equal(-1))
			})
		}

		It("should SetRange", func() {
			set := adapter.Set(ctx, "key", "Hello World", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			range_ := adapter.SetRange(ctx, "key", 6, "Redis")
			Expect(range_.Err()).NotTo(HaveOccurred())
			Expect(range_.Val()).To(Equal(int64(11)))

			get := adapter.Get(ctx, "key")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("Hello Redis"))
		})

		It("should StrLen", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			strLen := adapter.StrLen(ctx, "key")
			Expect(strLen.Err()).NotTo(HaveOccurred())
			Expect(strLen.Val()).To(Equal(int64(5)))

			strLen = adapter.StrLen(ctx, "_")
			Expect(strLen.Err()).NotTo(HaveOccurred())
			Expect(strLen.Val()).To(Equal(int64(0)))
		})

		if resp3 {
			It("should Copy", func() {
				set := adapter.Set(ctx, "key", "hello", 0)
				Expect(set.Err()).NotTo(HaveOccurred())
				Expect(set.Val()).To(Equal("OK"))

				copy := adapter.Copy(ctx, "key", "newKey", 0, false)
				Expect(copy.Err()).NotTo(HaveOccurred())
				Expect(copy.Val()).To(Equal(int64(1)))

				// Value is available by both keys now
				getOld := adapter.Get(ctx, "key")
				Expect(getOld.Err()).NotTo(HaveOccurred())
				Expect(getOld.Val()).To(Equal("hello"))
				getNew := adapter.Get(ctx, "newKey")
				Expect(getNew.Err()).NotTo(HaveOccurred())
				Expect(getNew.Val()).To(Equal("hello"))

				// Overwriting an existing key should not succeed
				overwrite := adapter.Copy(ctx, "newKey", "key", 0, false)
				Expect(overwrite.Val()).To(Equal(int64(0)))

				// Overwrite is allowed when replace=rue
				replace := adapter.Copy(ctx, "newKey", "key", 0, true)
				Expect(replace.Val()).To(Equal(int64(1)))
			})
		}
	})

	Describe("ACL", func() {
		if resp3 {
			TestUserName := "test"
			It("should ACL LOG", func() {
				Expect(adapter.ACLLogReset(ctx).Err()).NotTo(HaveOccurred())
				err := adapter.ACLSetUser(ctx, "test", ">test", "on", "allkeys", "+get").Err()
				Expect(err).NotTo(HaveOccurred())

				for addr := range clientresp3.Nodes() {
					clientAcl, err := rueidis.NewClient(rueidis.ClientOption{
						InitAddress:  []string{addr},
						Username:     "test",
						Password:     "test",
						DisableCache: true,
					})
					Expect(err).NotTo(HaveOccurred())
					adapterACL := NewAdapter(clientAcl)
					_ = adapterACL.Set(ctx, "mystring", "foo", 0).Err()
					_ = adapterACL.HSet(ctx, "myhash", "foo", "bar").Err()
					_ = adapterACL.SAdd(ctx, "myset", "foo", "bar").Err()
					clientAcl.Close()
					break
				}

				logEntries, err := adapter.ACLLog(ctx, 10).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(logEntries)).To(Equal(6))

				for _, entry := range logEntries {
					Expect(entry.Reason).To(Equal("command"))
					Expect(entry.Context).To(Equal("toplevel"))
					Expect(entry.Object).NotTo(BeEmpty())
					Expect(entry.Username).To(Equal("test"))
					Expect(entry.AgeSeconds).To(BeNumerically(">=", 0))
					Expect(entry.ClientInfo).NotTo(BeNil())
					Expect(entry.EntryID).To(BeNumerically(">=", 0))
					Expect(entry.TimestampCreated).To(BeNumerically(">=", 0))
					Expect(entry.TimestampLastUpdated).To(BeNumerically(">=", 0))
				}

				limitedLogEntries, err := adapter.ACLLog(ctx, 2).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(limitedLogEntries)).To(Equal(2))

				// cleanup after creating the user
				err = adapter.ACLDelUser(ctx, "test").Err()
				Expect(err).NotTo(HaveOccurred())
			})

			It("should ACL LOG RESET", func() {
				// Call ACL LOG RESET
				resetCmd := adapter.ACLLogReset(ctx)
				Expect(resetCmd.Err()).NotTo(HaveOccurred())
				Expect(resetCmd.Val()).To(Equal("OK"))

				// Verify that the log is empty after the reset
				logEntries, err := adapter.ACLLog(ctx, 10).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(logEntries)).To(Equal(0))
			})

			It("list only default user", func() {
				res, err := adapter.ACLList(ctx).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(HaveLen(1))
				Expect(res[0]).To(ContainSubstring("default"))
			})

			It("setuser and deluser", func() {
				res, err := adapter.ACLList(ctx).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(HaveLen(1))
				Expect(res[0]).To(ContainSubstring("default"))

				add, err := adapter.ACLSetUser(ctx, TestUserName, "nopass", "on", "allkeys", "+set", "+get").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(add).To(Equal("OK"))

				resAfter, err := adapter.ACLList(ctx).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resAfter).To(HaveLen(2))
				Expect(resAfter[1]).To(ContainSubstring(TestUserName))

				deletedN, err := adapter.ACLDelUser(ctx, TestUserName).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(deletedN).To(BeNumerically("==", 1))

				resAfterDeletion, err := adapter.ACLList(ctx).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resAfterDeletion).To(HaveLen(1))
				Expect(resAfterDeletion[0]).To(BeEquivalentTo(res[0]))
			})

			It("should acl dryrun", func() {
				dryRun := adapter.ACLDryRun(ctx, "default", "get", "randomKey")
				Expect(dryRun.Err()).NotTo(HaveOccurred())
				Expect(dryRun.Val()).To(Equal("OK"))
			})

			It("lists acl categories and subcategories", func() {
				res, err := adapter.ACLCat(ctx).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(res)).To(BeNumerically(">", 20))
				Expect(res).To(ContainElements(
					"read",
					"write",
					"keyspace",
					"dangerous",
					"slow",
					"set",
					"sortedset",
					"list",
					"hash",
				))

				res, err = adapter.ACLCatArgs(ctx, &ACLCatArgs{Category: "read"}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(ContainElement("get"))
			})
		}
	})

	Describe("hashes", func() {
		It("should HDel", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			Expect(hSet.Err()).NotTo(HaveOccurred())

			hDel := adapter.HDel(ctx, "hash", "key")
			Expect(hDel.Err()).NotTo(HaveOccurred())
			Expect(hDel.Val()).To(Equal(int64(1)))

			hDel = adapter.HDel(ctx, "hash", "key")
			Expect(hDel.Err()).NotTo(HaveOccurred())
			Expect(hDel.Val()).To(Equal(int64(0)))
		})

		It("should HExists", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			Expect(hSet.Err()).NotTo(HaveOccurred())

			hExists := adapter.HExists(ctx, "hash", "key")
			Expect(hExists.Err()).NotTo(HaveOccurred())
			Expect(hExists.Val()).To(Equal(true))

			hExists = adapter.HExists(ctx, "hash", "key1")
			Expect(hExists.Err()).NotTo(HaveOccurred())
			Expect(hExists.Val()).To(Equal(false))
		})

		It("should HGet", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			Expect(hSet.Err()).NotTo(HaveOccurred())

			hGet := adapter.HGet(ctx, "hash", "key")
			Expect(hGet.Err()).NotTo(HaveOccurred())
			Expect(hGet.Val()).To(Equal("hello"))

			hGet = adapter.HGet(ctx, "hash", "key1")
			Expect(rueidis.IsRedisNil(hGet.Err())).To(BeTrue())
			Expect(hGet.Val()).To(Equal(""))
		})

		It("should HGetAll", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1").Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.HSet(ctx, "hash", "key2", "hello2").Err()
			Expect(err).NotTo(HaveOccurred())

			m, err := adapter.HGetAll(ctx, "hash").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(m).To(Equal(map[string]string{"key1": "hello1", "key2": "hello2"}))
		})

		It("should scan", func() {
			now := time.Now()

			err := adapter.HMSet(ctx, "hash", "key1", "hello1", "key2", 123, "time", now.Format(time.RFC3339Nano)).Err()
			Expect(err).NotTo(HaveOccurred())

			res := adapter.HGetAll(ctx, "hash")
			Expect(res.Err()).NotTo(HaveOccurred())

			type data struct {
				Key1 string    `redis:"key1"`
				Key2 int       `redis:"key2"`
				Time TimeValue `redis:"time"`
			}
			var d data
			Expect(res.Scan(&d)).NotTo(HaveOccurred())
			Expect(d.Time.UnixNano()).To(Equal(now.UnixNano()))
			d.Time.Time = time.Time{}
			Expect(d).To(Equal(data{
				Key1: "hello1",
				Key2: 123,
				Time: TimeValue{Time: time.Time{}},
			}))

			type data2 struct {
				Key1 string    `redis:"key1"`
				Key2 int       `redis:"key2"`
				Time time.Time `redis:"time"`
			}
			err = adapter.HSet(ctx, "hash", &data2{
				Key1: "hello2",
				Key2: 200,
				Time: now,
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			var d2 data2
			err = adapter.HMGet(ctx, "hash", "key1", "key2", "time").Scan(&d2)
			Expect(err).NotTo(HaveOccurred())
			Expect(d2.Key1).To(Equal("hello2"))
			Expect(d2.Key2).To(Equal(200))
			Expect(d2.Time.Unix()).To(Equal(now.Unix()))
		})

		It("should HIncrBy", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "5")
			Expect(hSet.Err()).NotTo(HaveOccurred())

			hIncrBy := adapter.HIncrBy(ctx, "hash", "key", 1)
			Expect(hIncrBy.Err()).NotTo(HaveOccurred())
			Expect(hIncrBy.Val()).To(Equal(int64(6)))

			hIncrBy = adapter.HIncrBy(ctx, "hash", "key", -1)
			Expect(hIncrBy.Err()).NotTo(HaveOccurred())
			Expect(hIncrBy.Val()).To(Equal(int64(5)))

			hIncrBy = adapter.HIncrBy(ctx, "hash", "key", -10)
			Expect(hIncrBy.Err()).NotTo(HaveOccurred())
			Expect(hIncrBy.Val()).To(Equal(int64(-5)))
		})

		It("should HIncrByFloat", func() {
			hSet := adapter.HSet(ctx, "hash", "field", "10.50")
			Expect(hSet.Err()).NotTo(HaveOccurred())
			Expect(hSet.Val()).To(Equal(int64(1)))

			hIncrByFloat := adapter.HIncrByFloat(ctx, "hash", "field", 0.1)
			Expect(hIncrByFloat.Err()).NotTo(HaveOccurred())
			Expect(hIncrByFloat.Val()).To(Equal(10.6))

			hSet = adapter.HSet(ctx, "hash", "field", "5.0e3")
			Expect(hSet.Err()).NotTo(HaveOccurred())
			Expect(hSet.Val()).To(Equal(int64(0)))

			hIncrByFloat = adapter.HIncrByFloat(ctx, "hash", "field", 2.0e2)
			Expect(hIncrByFloat.Err()).NotTo(HaveOccurred())
			Expect(hIncrByFloat.Val()).To(Equal(float64(5200)))
		})

		It("should HKeys", func() {
			hkeys := adapter.HKeys(ctx, "hash")
			Expect(hkeys.Err()).NotTo(HaveOccurred())
			Expect(hkeys.Val()).To(Equal([]string{}))

			hset := adapter.HSet(ctx, "hash", "key1", "hello1")
			Expect(hset.Err()).NotTo(HaveOccurred())
			hset = adapter.HSet(ctx, "hash", "key2", "hello2")
			Expect(hset.Err()).NotTo(HaveOccurred())

			hkeys = adapter.HKeys(ctx, "hash")
			Expect(hkeys.Err()).NotTo(HaveOccurred())
			Expect(hkeys.Val()).To(Equal([]string{"key1", "key2"}))
		})

		It("should HLen", func() {
			hSet := adapter.HSet(ctx, "hash", "key1", "hello1")
			Expect(hSet.Err()).NotTo(HaveOccurred())
			hSet = adapter.HSet(ctx, "hash", "key2", "hello2")
			Expect(hSet.Err()).NotTo(HaveOccurred())

			hLen := adapter.HLen(ctx, "hash")
			Expect(hLen.Err()).NotTo(HaveOccurred())
			Expect(hLen.Val()).To(Equal(int64(2)))
		})

		It("should HMGet", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1", "key2", "hello2").Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.HMGet(ctx, "hash", "key1", "key2", "_").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]any{"hello1", "hello2", nil}))
		})

		It("should HSet", func() {
			ok, err := adapter.HSet(ctx, "hash", map[string]interface{}{
				"key1": "hello1",
				"key2": "hello2",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(ok).To(Equal(int64(2)))

			v, err := adapter.HGet(ctx, "hash", "key1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal("hello1"))

			v, err = adapter.HGet(ctx, "hash", "key2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal("hello2"))

			keys, err := adapter.HKeys(ctx, "hash").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(keys).To(ConsistOf([]string{"key1", "key2"}))
		})

		It("should HSet", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			Expect(hSet.Err()).NotTo(HaveOccurred())
			Expect(hSet.Val()).To(Equal(int64(1)))

			hGet := adapter.HGet(ctx, "hash", "key")
			Expect(hGet.Err()).NotTo(HaveOccurred())
			Expect(hGet.Val()).To(Equal("hello"))

			// set struct
			// MSet struct
			type set struct {
				Set1 string                 `redis:"set1"`
				Set2 int16                  `redis:"set2"`
				Set3 time.Duration          `redis:"set3"`
				Set4 interface{}            `redis:"set4"`
				Set5 map[string]interface{} `redis:"-"`
				Set6 string                 `redis:"set6,omitempty"`
				Set7 *string                `redis:"set7"`
				Set8 *string                `redis:"set8"`
			}
			str := "str"
			hSet = adapter.HSet(ctx, "hash", &set{
				Set1: "val1",
				Set2: 1024,
				Set3: 2 * time.Millisecond,
				Set4: nil,
				Set5: map[string]interface{}{"k1": 1},
				Set7: &str,
				Set8: nil,
			})
			Expect(hSet.Err()).NotTo(HaveOccurred())
			Expect(hSet.Val()).To(Equal(int64(5)))

			hMGet := adapter.HMGet(ctx, "hash", "set1", "set2", "set3", "set4", "set5", "set6", "set7", "set8")
			Expect(hMGet.Err()).NotTo(HaveOccurred())
			Expect(hMGet.Val()).To(Equal([]interface{}{
				"val1",
				"1024",
				strconv.Itoa(int(2 * time.Millisecond.Nanoseconds())),
				"",
				nil,
				nil,
				str,
				nil,
			}))

			hSet = adapter.HSet(ctx, "hash2", &set{
				Set1: "val2",
				Set6: "val",
			})
			Expect(hSet.Err()).NotTo(HaveOccurred())
			Expect(hSet.Val()).To(Equal(int64(5)))

			hMGet = adapter.HMGet(ctx, "hash2", "set1", "set6")
			Expect(hMGet.Err()).NotTo(HaveOccurred())
			Expect(hMGet.Val()).To(Equal([]interface{}{
				"val2",
				"val",
			}))
		})

		It("should HSetNX", func() {
			hSetNX := adapter.HSetNX(ctx, "hash", "key", "hello")
			Expect(hSetNX.Err()).NotTo(HaveOccurred())
			Expect(hSetNX.Val()).To(Equal(true))

			hSetNX = adapter.HSetNX(ctx, "hash", "key", "hello")
			Expect(hSetNX.Err()).NotTo(HaveOccurred())
			Expect(hSetNX.Val()).To(Equal(false))

			hGet := adapter.HGet(ctx, "hash", "key")
			Expect(hGet.Err()).NotTo(HaveOccurred())
			Expect(hGet.Val()).To(Equal("hello"))
		})

		It("should HVals", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1").Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.HSet(ctx, "hash", "key2", "hello2").Err()
			Expect(err).NotTo(HaveOccurred())

			v, err := adapter.HVals(ctx, "hash").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal([]string{"hello1", "hello2"}))

			// TODO
			// var slice []string
			// err = adapter.HVals(ctx, "hash").ScanSlice(&slice)
			// Expect(err).NotTo(HaveOccurred())
			// Expect(slice).To(Equal([]string{"hello1", "hello2"}))
		})

		if resp3 {
			It("should HRandField", func() {
				err := adapter.HSet(ctx, "hash", "key1", "hello1").Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.HSet(ctx, "hash", "key2", "hello2").Err()
				Expect(err).NotTo(HaveOccurred())

				v := adapter.HRandField(ctx, "hash", 1)
				Expect(v.Err()).NotTo(HaveOccurred())
				Expect(v.Val()).To(Or(Equal([]string{"key1"}), Equal([]string{"key2"})))

				v = adapter.HRandField(ctx, "hash", 0)
				Expect(v.Err()).NotTo(HaveOccurred())
				Expect(v.Val()).To(HaveLen(0))

				kv, err := adapter.HRandFieldWithValues(ctx, "hash", -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(kv).To(Or(
					Equal([]KeyValue{{Key: "key1", Value: "hello1"}}),
					Equal([]KeyValue{{Key: "key2", Value: "hello2"}}),
				))
			})

			It("should HStrLen", func() {
				hSet := adapter.HSet(ctx, "hash", "key", "hello")
				Expect(hSet.Err()).NotTo(HaveOccurred())

				hStrLen := adapter.HStrLen(ctx, "hash", "key")
				Expect(hStrLen.Err()).NotTo(HaveOccurred())
				Expect(hStrLen.Val()).To(Equal(int64(len("hello"))))

				nonHStrLen := adapter.HStrLen(ctx, "hash", "keyNon")
				Expect(hStrLen.Err()).NotTo(HaveOccurred())
				Expect(nonHStrLen.Val()).To(Equal(int64(0)))

				hDel := adapter.HDel(ctx, "hash", "key")
				Expect(hDel.Err()).NotTo(HaveOccurred())
				Expect(hDel.Val()).To(Equal(int64(1)))
			})

			It("should HExpire", Label("hash-expiration", "NonRedisEnterprise"), func() {
				res, err := adapter.HExpire(ctx, "no_such_key", 10*time.Second, "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(res).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err = adapter.HExpire(ctx, "myhash", 10*time.Second, "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, 1, -2}))
			})

			It("should HPExpire", Label("hash-expiration", "NonRedisEnterprise"), func() {
				res, err := adapter.HPExpire(ctx, "no_such_key", 10*time.Second, "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(res).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err = adapter.HPExpire(ctx, "myhash", 10*time.Second, "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, 1, -2}))
			})

			It("should HExpireAt", Label("hash-expiration", "NonRedisEnterprise"), func() {
				resEmpty, err := adapter.HExpireAt(ctx, "no_such_key", time.Now().Add(10*time.Second), "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err := adapter.HExpireAt(ctx, "myhash", time.Now().Add(10*time.Second), "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, 1, -2}))
			})

			It("should HPExpireAt", Label("hash-expiration", "NonRedisEnterprise"), func() {
				resEmpty, err := adapter.HPExpireAt(ctx, "no_such_key", time.Now().Add(10*time.Second), "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err := adapter.HPExpireAt(ctx, "myhash", time.Now().Add(10*time.Second), "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, 1, -2}))
			})

			It("should HPersist", Label("hash-expiration", "NonRedisEnterprise"), func() {
				resEmpty, err := adapter.HPersist(ctx, "no_such_key", "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err := adapter.HPersist(ctx, "myhash", "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{-1, -1, -2}))

				res, err = adapter.HExpire(ctx, "myhash", 10*time.Second, "key1", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, -2}))

				res, err = adapter.HPersist(ctx, "myhash", "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, -1, -2}))
			})

			It("should HExpireTime", Label("hash-expiration", "NonRedisEnterprise"), func() {
				resEmpty, err := adapter.HExpireTime(ctx, "no_such_key", "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err := adapter.HExpire(ctx, "myhash", 10*time.Second, "key1", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, -2}))

				res, err = adapter.HExpireTime(ctx, "myhash", "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res[0]).To(BeNumerically("~", time.Now().Add(10*time.Second).Unix(), 1))
			})

			It("should HPExpireTime", Label("hash-expiration", "NonRedisEnterprise"), func() {
				resEmpty, err := adapter.HPExpireTime(ctx, "no_such_key", "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				expireAt := time.Now().Add(10 * time.Second)
				res, err := adapter.HPExpireAt(ctx, "myhash", expireAt, "key1", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, -2}))

				res, err = adapter.HPExpireTime(ctx, "myhash", "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(BeEquivalentTo([]int64{expireAt.UnixMilli(), -1, -2}))
			})

			It("should HTTL", Label("hash-expiration", "NonRedisEnterprise"), func() {
				resEmpty, err := adapter.HTTL(ctx, "no_such_key", "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err := adapter.HExpire(ctx, "myhash", 10*time.Second, "key1", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, -2}))

				res, err = adapter.HTTL(ctx, "myhash", "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{10, -1, -2}))
			})

			It("should HPTTL", Label("hash-expiration", "NonRedisEnterprise"), func() {
				resEmpty, err := adapter.HPTTL(ctx, "no_such_key", "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err := adapter.HExpire(ctx, "myhash", 10*time.Second, "key1", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, -2}))

				res, err = adapter.HPTTL(ctx, "myhash", "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res[0]).To(BeNumerically("~", 10*time.Second.Milliseconds(), 1))
			})
		}
	})

	Describe("hyperloglog", func() {
		It("should PFMerge", func() {
			pfAdd := adapter.PFAdd(ctx, "hll1", "1", "2", "3", "4", "5")
			Expect(pfAdd.Err()).NotTo(HaveOccurred())

			pfCount := adapter.PFCount(ctx, "hll1")
			Expect(pfCount.Err()).NotTo(HaveOccurred())
			Expect(pfCount.Val()).To(Equal(int64(5)))

			pfAdd = adapter.PFAdd(ctx, "hll2", "a", "b", "c", "d", "e")
			Expect(pfAdd.Err()).NotTo(HaveOccurred())

			pfMerge := adapter.PFMerge(ctx, "hllMerged", "hll1", "hll2")
			Expect(pfMerge.Err()).NotTo(HaveOccurred())

			pfCount = adapter.PFCount(ctx, "hllMerged")
			Expect(pfCount.Err()).NotTo(HaveOccurred())
			Expect(pfCount.Val()).To(Equal(int64(10)))

			pfCount = adapter.PFCount(ctx, "hll1", "hll2")
			Expect(pfCount.Err()).NotTo(HaveOccurred())
			Expect(pfCount.Val()).To(Equal(int64(10)))
		})
	})

	Describe("lists", func() {
		It("should BLPop", func() {
			rPush := adapter.RPush(ctx, "list1", "a", "b", "c")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			bLPop := adapter.BLPop(ctx, 0, "list1", "list2")
			Expect(bLPop.Err()).NotTo(HaveOccurred())
			Expect(bLPop.Val()).To(Equal([]string{"list1", "a"}))
		})

		It("should BLPopBlocks", func() {
			started := make(chan bool)
			done := make(chan bool)
			go func() {
				defer GinkgoRecover()

				started <- true
				bLPop := adapter.BLPop(ctx, 0, "list")
				Expect(bLPop.Err()).NotTo(HaveOccurred())
				Expect(bLPop.Val()).To(Equal([]string{"list", "a"}))
				done <- true
			}()
			<-started

			select {
			case <-done:
				Fail("BLPop is not blocked")
			case <-time.After(time.Second):
				// ok
			}

			rPush := adapter.RPush(ctx, "list", "a")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			select {
			case <-done:
				// ok
			case <-time.After(time.Second):
				Fail("BLPop is still blocked")
			}
		})

		It("should BLPop timeout", func() {
			val, err := adapter.BLPop(ctx, time.Second, "list1").Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue())
			Expect(val).To(BeNil())

			Expect(adapter.Ping(ctx).Err()).NotTo(HaveOccurred())
		})

		It("should BRPop", func() {
			rPush := adapter.RPush(ctx, "list1", "a", "b", "c")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			bRPop := adapter.BRPop(ctx, 0, "list1", "list2")
			Expect(bRPop.Err()).NotTo(HaveOccurred())
			Expect(bRPop.Val()).To(Equal([]string{"list1", "c"}))
		})

		It("should BRPop blocks", func() {
			started := make(chan bool)
			done := make(chan bool)
			go func() {
				defer GinkgoRecover()

				started <- true
				brpop := adapter.BRPop(ctx, 0, "list")
				Expect(brpop.Err()).NotTo(HaveOccurred())
				Expect(brpop.Val()).To(Equal([]string{"list", "a"}))
				done <- true
			}()
			<-started

			select {
			case <-done:
				Fail("BRPop is not blocked")
			case <-time.After(time.Second):
				// ok
			}

			rPush := adapter.RPush(ctx, "list", "a")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			select {
			case <-done:
				// ok
			case <-time.After(time.Second):
				Fail("BRPop is still blocked")
				// ok
			}
		})

		It("should BRPopLPush", func() {
			_, err := adapter.BRPopLPush(ctx, "list1", "list2", time.Second).Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue())

			err = adapter.RPush(ctx, "list1", "a", "b", "c").Err()
			Expect(err).NotTo(HaveOccurred())

			v, err := adapter.BRPopLPush(ctx, "list1", "list2", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal("c"))
		})

		if resp3 {
			It("should LCS", func() {
				err := adapter.MSet(ctx, "LCSkey1{1}", "ohmytext", "LCSkey2{1}", "mynewtext").Err()
				Expect(err).NotTo(HaveOccurred())

				lcs, err := adapter.LCS(ctx, &LCSQuery{
					Key1: "LCSkey1{1}",
					Key2: "LCSkey2{1}",
				}).Result()

				Expect(err).NotTo(HaveOccurred())
				Expect(lcs.MatchString).To(Equal("mytext"))

				lcs, err = adapter.LCS(ctx, &LCSQuery{
					Key1: "LCSnonexistent_key1{1}",
					Key2: "LCSkey2{1}",
				}).Result()

				Expect(err).NotTo(HaveOccurred())
				Expect(lcs.MatchString).To(Equal(""))

				lcs, err = adapter.LCS(ctx, &LCSQuery{
					Key1: "LCSkey1{1}",
					Key2: "LCSkey2{1}",
					Len:  true,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(lcs.MatchString).To(Equal(""))
				Expect(lcs.Len).To(Equal(int64(6)))

				lcs, err = adapter.LCS(ctx, &LCSQuery{
					Key1: "LCSkey1{1}",
					Key2: "LCSkey2{1}",
					Idx:  true,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(lcs.MatchString).To(Equal(""))
				Expect(lcs.Len).To(Equal(int64(6)))
				Expect(lcs.Matches).To(Equal([]LCSMatchedPosition{
					{
						Key1:     LCSPosition{Start: 4, End: 7},
						Key2:     LCSPosition{Start: 5, End: 8},
						MatchLen: 0,
					},
					{
						Key1:     LCSPosition{Start: 2, End: 3},
						Key2:     LCSPosition{Start: 0, End: 1},
						MatchLen: 0,
					},
				}))

				lcs, err = adapter.LCS(ctx, &LCSQuery{
					Key1:         "LCSkey1{1}",
					Key2:         "LCSkey2{1}",
					Idx:          true,
					MinMatchLen:  3,
					WithMatchLen: true,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(lcs.MatchString).To(Equal(""))
				Expect(lcs.Len).To(Equal(int64(6)))
				Expect(lcs.Matches).To(Equal([]LCSMatchedPosition{
					{
						Key1:     LCSPosition{Start: 4, End: 7},
						Key2:     LCSPosition{Start: 5, End: 8},
						MatchLen: 4,
					},
				}))

				_, err = adapter.Set(ctx, "keywithstringvalue{1}", "golang", 0).Result()
				Expect(err).NotTo(HaveOccurred())
				_, err = adapter.LPush(ctx, "keywithnonstringvalue{1}", "somevalue").Result()
				Expect(err).NotTo(HaveOccurred())
				_, err = adapter.LCS(ctx, &LCSQuery{
					Key1: "keywithstringvalue{1}",
					Key2: "keywithnonstringvalue{1}",
				}).Result()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("The specified keys must contain string values"))
			})
		}

		It("should LIndex", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			Expect(lPush.Err()).NotTo(HaveOccurred())
			lPush = adapter.LPush(ctx, "list", "Hello")
			Expect(lPush.Err()).NotTo(HaveOccurred())

			lIndex := adapter.LIndex(ctx, "list", 0)
			Expect(lIndex.Err()).NotTo(HaveOccurred())
			Expect(lIndex.Val()).To(Equal("Hello"))

			lIndex = adapter.LIndex(ctx, "list", -1)
			Expect(lIndex.Err()).NotTo(HaveOccurred())
			Expect(lIndex.Val()).To(Equal("World"))

			lIndex = adapter.LIndex(ctx, "list", 3)
			Expect(rueidis.IsRedisNil(lIndex.Err())).To(BeTrue())
			Expect(lIndex.Val()).To(Equal(""))
		})

		It("LInsert should panic", func() {
			Expect(func() {
				adapter.LInsert(ctx, "list", "ANY", "World", "There")
			}).To(Panic())
		})

		It("should LInsert", func() {
			rPush := adapter.RPush(ctx, "list", "Hello")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "World")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			lInsert := adapter.LInsert(ctx, "list", "BEFORE", "World", "There")
			Expect(lInsert.Err()).NotTo(HaveOccurred())
			Expect(lInsert.Val()).To(Equal(int64(3)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"Hello", "There", "World"}))

			lInsert = adapter.LInsert(ctx, "list", "AFTER", "World", "There")
			Expect(lInsert.Err()).NotTo(HaveOccurred())
			Expect(lInsert.Val()).To(Equal(int64(4)))

			lRange = adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"Hello", "There", "World", "There"}))
		})

		It("should LInsert", func() {
			rPush := adapter.RPush(ctx, "list", "Hello")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "World")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			lInsert := adapter.LInsertBefore(ctx, "list", "World", "There")
			Expect(lInsert.Err()).NotTo(HaveOccurred())
			Expect(lInsert.Val()).To(Equal(int64(3)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"Hello", "There", "World"}))

			lInsert = adapter.LInsertAfter(ctx, "list", "World", "There")
			Expect(lInsert.Err()).NotTo(HaveOccurred())
			Expect(lInsert.Val()).To(Equal(int64(4)))

			lRange = adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"Hello", "There", "World", "There"}))
		})

		if resp3 {
			It("should LMPop", func() {
				err := adapter.LPush(ctx, "list1", "one", "two", "three", "four", "five").Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.LPush(ctx, "list2", "a", "b", "c", "d", "e").Err()
				Expect(err).NotTo(HaveOccurred())

				key, val, err := adapter.LMPop(ctx, "left", 3, "list1", "list2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(key).To(Equal("list1"))
				Expect(val).To(Equal([]string{"five", "four", "three"}))

				key, val, err = adapter.LMPop(ctx, "right", 3, "list1", "list2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(key).To(Equal("list1"))
				Expect(val).To(Equal([]string{"one", "two"}))

				key, val, err = adapter.LMPop(ctx, "left", 1, "list1", "list2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(key).To(Equal("list2"))
				Expect(val).To(Equal([]string{"e"}))

				key, val, err = adapter.LMPop(ctx, "right", 10, "list1", "list2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(key).To(Equal("list2"))
				Expect(val).To(Equal([]string{"a", "b", "c", "d"}))

				err = adapter.LMPop(ctx, "left", 10, "list1", "list2").Err()
				Expect(err).To(Equal(rueidis.Nil))

				err = adapter.Set(ctx, "list3", 1024, 0).Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.LMPop(ctx, "left", 10, "list1", "list2", "list3").Err()
				Expect(err.Error()).To(Equal("WRONGTYPE Operation against a key holding the wrong kind of value"))

				err = adapter.LMPop(ctx, "right", 0, "list1", "list2").Err()
				Expect(err).To(HaveOccurred())
			})

			It("should BLMPop", func() {
				err := adapter.LPush(ctx, "list1", "one", "two", "three", "four", "five").Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.LPush(ctx, "list2", "a", "b", "c", "d", "e").Err()
				Expect(err).NotTo(HaveOccurred())

				key, val, err := adapter.BLMPop(ctx, 0, "left", 3, "list1", "list2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(key).To(Equal("list1"))
				Expect(val).To(Equal([]string{"five", "four", "three"}))

				key, val, err = adapter.BLMPop(ctx, 0, "right", 3, "list1", "list2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(key).To(Equal("list1"))
				Expect(val).To(Equal([]string{"one", "two"}))

				key, val, err = adapter.BLMPop(ctx, 0, "left", 1, "list1", "list2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(key).To(Equal("list2"))
				Expect(val).To(Equal([]string{"e"}))

				key, val, err = adapter.BLMPop(ctx, 0, "right", 10, "list1", "list2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(key).To(Equal("list2"))
				Expect(val).To(Equal([]string{"a", "b", "c", "d"}))

			})

			It("should BLMPopBlocks", func() {
				started := make(chan bool)
				done := make(chan bool)
				go func() {
					defer GinkgoRecover()

					started <- true
					key, val, err := adapter.BLMPop(ctx, 0, "left", 1, "list_list").Result()
					Expect(err).NotTo(HaveOccurred())
					Expect(key).To(Equal("list_list"))
					Expect(val).To(Equal([]string{"a"}))
					done <- true
				}()
				<-started

				select {
				case <-done:
					Fail("BLMPop is not blocked")
				case <-time.After(time.Second):
					// ok
				}

				_, err := adapter.LPush(ctx, "list_list", "a").Result()
				Expect(err).NotTo(HaveOccurred())

				select {
				case <-done:
					// ok
				case <-time.After(time.Second):
					Fail("BLMPop is still blocked")
				}
			})

			It("should BLMPop timeout", func() {
				_, val, err := adapter.BLMPop(ctx, time.Second, "left", 1, "list1").Result()
				Expect(err).To(Equal(rueidis.Nil))
				Expect(val).To(BeNil())

				Expect(adapter.Ping(ctx).Err()).NotTo(HaveOccurred())
			})
		}

		It("should LLen", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			Expect(lPush.Err()).NotTo(HaveOccurred())
			lPush = adapter.LPush(ctx, "list", "Hello")
			Expect(lPush.Err()).NotTo(HaveOccurred())

			lLen := adapter.LLen(ctx, "list")
			Expect(lLen.Err()).NotTo(HaveOccurred())
			Expect(lLen.Val()).To(Equal(int64(2)))
		})

		It("should LPop", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			lPop := adapter.LPop(ctx, "list")
			Expect(lPop.Err()).NotTo(HaveOccurred())
			Expect(lPop.Val()).To(Equal("one"))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"two", "three"}))
		})

		if resp3 {
			It("should LPopCount", func() {
				rPush := adapter.RPush(ctx, "list", "one")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "two")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "three")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "four")
				Expect(rPush.Err()).NotTo(HaveOccurred())

				lPopCount := adapter.LPopCount(ctx, "list", 2)
				Expect(lPopCount.Err()).NotTo(HaveOccurred())
				Expect(lPopCount.Val()).To(Equal([]string{"one", "two"}))

				lRange := adapter.LRange(ctx, "list", 0, -1)
				Expect(lRange.Err()).NotTo(HaveOccurred())
				Expect(lRange.Val()).To(Equal([]string{"three", "four"}))
			})

			It("should LPos", func() {
				rPush := adapter.RPush(ctx, "list", "a")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "b")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "c")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "b")
				Expect(rPush.Err()).NotTo(HaveOccurred())

				lPos := adapter.LPos(ctx, "list", "b", LPosArgs{})
				Expect(lPos.Err()).NotTo(HaveOccurred())
				Expect(lPos.Val()).To(Equal(int64(1)))

				lPos = adapter.LPos(ctx, "list", "b", LPosArgs{Rank: 2})
				Expect(lPos.Err()).NotTo(HaveOccurred())
				Expect(lPos.Val()).To(Equal(int64(3)))

				lPos = adapter.LPos(ctx, "list", "b", LPosArgs{Rank: -2})
				Expect(lPos.Err()).NotTo(HaveOccurred())
				Expect(lPos.Val()).To(Equal(int64(1)))

				lPos = adapter.LPos(ctx, "list", "b", LPosArgs{Rank: 2, MaxLen: 1})
				Expect(rueidis.IsRedisNil(lPos.Err())).To(BeTrue())

				lPos = adapter.LPos(ctx, "list", "z", LPosArgs{})
				Expect(rueidis.IsRedisNil(lPos.Err())).To(BeTrue())
			})

			It("should LPosCount", func() {
				rPush := adapter.RPush(ctx, "list", "a")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "b")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "c")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "b")
				Expect(rPush.Err()).NotTo(HaveOccurred())

				lPos := adapter.LPosCount(ctx, "list", "b", 2, LPosArgs{})
				Expect(lPos.Err()).NotTo(HaveOccurred())
				Expect(lPos.Val()).To(Equal([]int64{1, 3}))

				lPos = adapter.LPosCount(ctx, "list", "b", 2, LPosArgs{Rank: 2})
				Expect(lPos.Err()).NotTo(HaveOccurred())
				Expect(lPos.Val()).To(Equal([]int64{3}))

				lPos = adapter.LPosCount(ctx, "list", "b", 1, LPosArgs{Rank: 1, MaxLen: 1})
				Expect(lPos.Err()).NotTo(HaveOccurred())
				Expect(lPos.Val()).To(Equal([]int64{}))

				lPos = adapter.LPosCount(ctx, "list", "b", 1, LPosArgs{Rank: 1, MaxLen: 0})
				Expect(lPos.Err()).NotTo(HaveOccurred())
				Expect(lPos.Val()).To(Equal([]int64{1}))
			})
		}

		It("should LPush", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			Expect(lPush.Err()).NotTo(HaveOccurred())
			lPush = adapter.LPush(ctx, "list", "Hello")
			Expect(lPush.Err()).NotTo(HaveOccurred())

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"Hello", "World"}))
		})

		It("should LPushX", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			Expect(lPush.Err()).NotTo(HaveOccurred())

			lPushX := adapter.LPushX(ctx, "list", "Hello")
			Expect(lPushX.Err()).NotTo(HaveOccurred())
			Expect(lPushX.Val()).To(Equal(int64(2)))

			lPush = adapter.LPush(ctx, "list1", "three")
			Expect(lPush.Err()).NotTo(HaveOccurred())
			Expect(lPush.Val()).To(Equal(int64(1)))

			lPushX = adapter.LPushX(ctx, "list1", "two", "one")
			Expect(lPushX.Err()).NotTo(HaveOccurred())
			Expect(lPushX.Val()).To(Equal(int64(3)))

			lPushX = adapter.LPushX(ctx, "list2", "Hello")
			Expect(lPushX.Err()).NotTo(HaveOccurred())
			Expect(lPushX.Val()).To(Equal(int64(0)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"Hello", "World"}))

			lRange = adapter.LRange(ctx, "list1", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"one", "two", "three"}))

			lRange = adapter.LRange(ctx, "list2", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{}))
		})

		It("should LRange", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			lRange := adapter.LRange(ctx, "list", 0, 0)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"one"}))

			lRange = adapter.LRange(ctx, "list", -3, 2)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"one", "two", "three"}))

			lRange = adapter.LRange(ctx, "list", -100, 100)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"one", "two", "three"}))

			lRange = adapter.LRange(ctx, "list", 5, 10)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{}))
		})

		It("should LRem", func() {
			rPush := adapter.RPush(ctx, "list", "hello")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "hello")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "key")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "hello")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			lRem := adapter.LRem(ctx, "list", -2, "hello")
			Expect(lRem.Err()).NotTo(HaveOccurred())
			Expect(lRem.Val()).To(Equal(int64(2)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"hello", "key"}))
		})

		It("should LSet", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			lSet := adapter.LSet(ctx, "list", 0, "four")
			Expect(lSet.Err()).NotTo(HaveOccurred())
			Expect(lSet.Val()).To(Equal("OK"))

			lSet = adapter.LSet(ctx, "list", -2, "five")
			Expect(lSet.Err()).NotTo(HaveOccurred())
			Expect(lSet.Val()).To(Equal("OK"))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"four", "five", "three"}))
		})

		It("should LTrim", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			lTrim := adapter.LTrim(ctx, "list", 1, -1)
			Expect(lTrim.Err()).NotTo(HaveOccurred())
			Expect(lTrim.Val()).To(Equal("OK"))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"two", "three"}))
		})

		It("should RPop", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			rPop := adapter.RPop(ctx, "list")
			Expect(rPop.Err()).NotTo(HaveOccurred())
			Expect(rPop.Val()).To(Equal("three"))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"one", "two"}))
		})

		if resp3 {
			It("should RPopCount", func() {
				rPush := adapter.RPush(ctx, "list", "one", "two", "three", "four")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				Expect(rPush.Val()).To(Equal(int64(4)))

				rPopCount := adapter.RPopCount(ctx, "list", 2)
				Expect(rPopCount.Err()).NotTo(HaveOccurred())
				Expect(rPopCount.Val()).To(Equal([]string{"four", "three"}))

				lRange := adapter.LRange(ctx, "list", 0, -1)
				Expect(lRange.Err()).NotTo(HaveOccurred())
				Expect(lRange.Val()).To(Equal([]string{"one", "two"}))
			})
		}

		It("should RPopLPush", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			rPopLPush := adapter.RPopLPush(ctx, "list", "list2")
			Expect(rPopLPush.Err()).NotTo(HaveOccurred())
			Expect(rPopLPush.Val()).To(Equal("three"))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"one", "two"}))

			lRange = adapter.LRange(ctx, "list2", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"three"}))
		})

		It("should RPush", func() {
			rPush := adapter.RPush(ctx, "list", "Hello")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			Expect(rPush.Val()).To(Equal(int64(1)))

			rPush = adapter.RPush(ctx, "list", "World")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			Expect(rPush.Val()).To(Equal(int64(2)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"Hello", "World"}))
		})

		It("should RPushX", func() {
			rPush := adapter.RPush(ctx, "list", "Hello")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			Expect(rPush.Val()).To(Equal(int64(1)))

			rPushX := adapter.RPushX(ctx, "list", "World")
			Expect(rPushX.Err()).NotTo(HaveOccurred())
			Expect(rPushX.Val()).To(Equal(int64(2)))

			rPush = adapter.RPush(ctx, "list1", "one")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			Expect(rPush.Val()).To(Equal(int64(1)))

			rPushX = adapter.RPushX(ctx, "list1", "two", "three")
			Expect(rPushX.Err()).NotTo(HaveOccurred())
			Expect(rPushX.Val()).To(Equal(int64(3)))

			rPushX = adapter.RPushX(ctx, "list2", "World")
			Expect(rPushX.Err()).NotTo(HaveOccurred())
			Expect(rPushX.Val()).To(Equal(int64(0)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"Hello", "World"}))

			lRange = adapter.LRange(ctx, "list1", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"one", "two", "three"}))

			lRange = adapter.LRange(ctx, "list2", 0, -1)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{}))
		})

		if resp3 {
			It("should LMove", func() {
				rPush := adapter.RPush(ctx, "lmove1", "ichi")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				Expect(rPush.Val()).To(Equal(int64(1)))

				rPush = adapter.RPush(ctx, "lmove1", "ni")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				Expect(rPush.Val()).To(Equal(int64(2)))

				rPush = adapter.RPush(ctx, "lmove1", "san")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				Expect(rPush.Val()).To(Equal(int64(3)))

				lMove := adapter.LMove(ctx, "lmove1", "lmove2", "RIGHT", "LEFT")
				Expect(lMove.Err()).NotTo(HaveOccurred())
				Expect(lMove.Val()).To(Equal("san"))

				lRange := adapter.LRange(ctx, "lmove2", 0, -1)
				Expect(lRange.Err()).NotTo(HaveOccurred())
				Expect(lRange.Val()).To(Equal([]string{"san"}))
			})

			It("should BLMove", func() {
				rPush := adapter.RPush(ctx, "blmove1", "ichi")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				Expect(rPush.Val()).To(Equal(int64(1)))

				rPush = adapter.RPush(ctx, "blmove1", "ni")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				Expect(rPush.Val()).To(Equal(int64(2)))

				rPush = adapter.RPush(ctx, "blmove1", "san")
				Expect(rPush.Err()).NotTo(HaveOccurred())
				Expect(rPush.Val()).To(Equal(int64(3)))

				blMove := adapter.BLMove(ctx, "blmove1", "blmove2", "RIGHT", "LEFT", time.Second)
				Expect(blMove.Err()).NotTo(HaveOccurred())
				Expect(blMove.Val()).To(Equal("san"))

				lRange := adapter.LRange(ctx, "blmove2", 0, -1)
				Expect(lRange.Err()).NotTo(HaveOccurred())
				Expect(lRange.Val()).To(Equal([]string{"san"}))
			})
		}
	})

	Describe("sets", func() {
		It("should SAdd", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			Expect(sAdd.Val()).To(Equal(int64(1)))

			sAdd = adapter.SAdd(ctx, "set", "World")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			Expect(sAdd.Val()).To(Equal(int64(1)))

			sAdd = adapter.SAdd(ctx, "set", "World")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			Expect(sAdd.Val()).To(Equal(int64(0)))

			sMembers := adapter.SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(ConsistOf([]string{"Hello", "World"}))
		})

		It("should SAdd strings", func() {
			set := []string{"Hello", "World", "World"}
			sAdd := adapter.SAdd(ctx, "set", set)
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			Expect(sAdd.Val()).To(Equal(int64(2)))

			sMembers := adapter.SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(ConsistOf([]string{"Hello", "World"}))
		})

		It("should SCard", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			Expect(sAdd.Val()).To(Equal(int64(1)))

			sAdd = adapter.SAdd(ctx, "set", "World")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			Expect(sAdd.Val()).To(Equal(int64(1)))

			sCard := adapter.SCard(ctx, "set")
			Expect(sCard.Err()).NotTo(HaveOccurred())
			Expect(sCard.Val()).To(Equal(int64(2)))
		})

		It("should SDiff", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sDiff := adapter.SDiff(ctx, "set1", "set2")
			Expect(sDiff.Err()).NotTo(HaveOccurred())
			Expect(sDiff.Val()).To(ConsistOf([]string{"a", "b"}))
		})

		It("should SDiffStore", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sDiffStore := adapter.SDiffStore(ctx, "set", "set1", "set2")
			Expect(sDiffStore.Err()).NotTo(HaveOccurred())
			Expect(sDiffStore.Val()).To(Equal(int64(2)))

			sMembers := adapter.SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(ConsistOf([]string{"a", "b"}))
		})

		It("should SInter", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sInter := adapter.SInter(ctx, "set1", "set2")
			Expect(sInter.Err()).NotTo(HaveOccurred())
			Expect(sInter.Val()).To(Equal([]string{"c"}))
		})

		if resp3 {
			It("should SInterCard", func() {
				sAdd := adapter.SAdd(ctx, "set1", "a")
				Expect(sAdd.Err()).NotTo(HaveOccurred())
				sAdd = adapter.SAdd(ctx, "set1", "b")
				Expect(sAdd.Err()).NotTo(HaveOccurred())
				sAdd = adapter.SAdd(ctx, "set1", "c")
				Expect(sAdd.Err()).NotTo(HaveOccurred())

				sAdd = adapter.SAdd(ctx, "set2", "b")
				Expect(sAdd.Err()).NotTo(HaveOccurred())
				sAdd = adapter.SAdd(ctx, "set2", "c")
				Expect(sAdd.Err()).NotTo(HaveOccurred())
				sAdd = adapter.SAdd(ctx, "set2", "d")
				Expect(sAdd.Err()).NotTo(HaveOccurred())
				sAdd = adapter.SAdd(ctx, "set2", "e")
				Expect(sAdd.Err()).NotTo(HaveOccurred())
				// limit 0 means no limit,see https://redis.io/commands/sintercard/ for more details
				sInterCard := adapter.SInterCard(ctx, 0, "set1", "set2")
				Expect(sInterCard.Err()).NotTo(HaveOccurred())
				Expect(sInterCard.Val()).To(Equal(int64(2)))

				sInterCard = adapter.SInterCard(ctx, 1, "set1", "set2")
				Expect(sInterCard.Err()).NotTo(HaveOccurred())
				Expect(sInterCard.Val()).To(Equal(int64(1)))

				sInterCard = adapter.SInterCard(ctx, 3, "set1", "set2")
				Expect(sInterCard.Err()).NotTo(HaveOccurred())
				Expect(sInterCard.Val()).To(Equal(int64(2)))
			})
		}

		It("should SInterStore", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sInterStore := adapter.SInterStore(ctx, "set", "set1", "set2")
			Expect(sInterStore.Err()).NotTo(HaveOccurred())
			Expect(sInterStore.Val()).To(Equal(int64(1)))

			sMembers := adapter.SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(Equal([]string{"c"}))
		})

		It("should IsMember", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sIsMember := adapter.SIsMember(ctx, "set", "one")
			Expect(sIsMember.Err()).NotTo(HaveOccurred())
			Expect(sIsMember.Val()).To(Equal(true))

			sIsMember = adapter.SIsMember(ctx, "set", "two")
			Expect(sIsMember.Err()).NotTo(HaveOccurred())
			Expect(sIsMember.Val()).To(Equal(false))
		})

		if resp3 {
			It("should SMIsMember", func() {
				sAdd := adapter.SAdd(ctx, "set", "one")
				Expect(sAdd.Err()).NotTo(HaveOccurred())

				sMIsMember := adapter.SMIsMember(ctx, "set", "one", "two")
				Expect(sMIsMember.Err()).NotTo(HaveOccurred())
				Expect(sMIsMember.Val()).To(Equal([]bool{true, false}))
			})
		}

		It("should SMembers", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "World")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sMembers := adapter.SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(ConsistOf([]string{"Hello", "World"}))
		})

		It("should SMembersMap", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "World")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sMembersMap := adapter.SMembersMap(ctx, "set")
			Expect(sMembersMap.Err()).NotTo(HaveOccurred())
			Expect(sMembersMap.Val()).To(Equal(map[string]struct{}{"Hello": {}, "World": {}}))
		})

		It("should SMove", func() {
			sAdd := adapter.SAdd(ctx, "set1", "one")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "two")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "three")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sMove := adapter.SMove(ctx, "set1", "set2", "two")
			Expect(sMove.Err()).NotTo(HaveOccurred())
			Expect(sMove.Val()).To(Equal(true))

			sMembers := adapter.SMembers(ctx, "set1")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(Equal([]string{"one"}))

			sMembers = adapter.SMembers(ctx, "set2")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(ConsistOf([]string{"three", "two"}))
		})

		It("should SPop", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "two")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "three")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sPop := adapter.SPop(ctx, "set")
			Expect(sPop.Err()).NotTo(HaveOccurred())
			Expect(sPop.Val()).NotTo(Equal(""))

			sMembers := adapter.SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(HaveLen(2))
		})

		It("should SPopN", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "two")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "three")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "four")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sPopN := adapter.SPopN(ctx, "set", 1)
			Expect(sPopN.Err()).NotTo(HaveOccurred())
			Expect(sPopN.Val()).NotTo(Equal([]string{""}))

			sMembers := adapter.SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(HaveLen(3))

			sPopN = adapter.SPopN(ctx, "set", 4)
			Expect(sPopN.Err()).NotTo(HaveOccurred())
			Expect(sPopN.Val()).To(HaveLen(3))

			sMembers = adapter.SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(HaveLen(0))
		})

		It("should SRandMember and SRandMemberN", func() {
			err := adapter.SAdd(ctx, "set", "one").Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.SAdd(ctx, "set", "two").Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.SAdd(ctx, "set", "three").Err()
			Expect(err).NotTo(HaveOccurred())

			members, err := adapter.SMembers(ctx, "set").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(HaveLen(3))

			member, err := adapter.SRandMember(ctx, "set").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(member).NotTo(Equal(""))

			members, err = adapter.SRandMemberN(ctx, "set", 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(HaveLen(2))
		})

		It("should SRem", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "two")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "three")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sRem := adapter.SRem(ctx, "set", "one")
			Expect(sRem.Err()).NotTo(HaveOccurred())
			Expect(sRem.Val()).To(Equal(int64(1)))

			sRem = adapter.SRem(ctx, "set", "four")
			Expect(sRem.Err()).NotTo(HaveOccurred())
			Expect(sRem.Val()).To(Equal(int64(0)))

			sMembers := adapter.SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(ConsistOf([]string{"three", "two"}))
		})

		It("should SUnion", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sUnion := adapter.SUnion(ctx, "set1", "set2")
			Expect(sUnion.Err()).NotTo(HaveOccurred())
			Expect(sUnion.Val()).To(HaveLen(5))
		})

		It("should SUnionStore", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sUnionStore := adapter.SUnionStore(ctx, "set", "set1", "set2")
			Expect(sUnionStore.Err()).NotTo(HaveOccurred())
			Expect(sUnionStore.Val()).To(Equal(int64(5)))

			sMembers := adapter.SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(HaveLen(5))
		})
	})

	Describe("sorted sets", func() {
		It("should BZPopMax", func() {
			err := adapter.ZAdd(ctx, "zset1", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{
				Score:  3,
				Member: "three",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			member, err := adapter.BZPopMax(ctx, 0, "zset1", "zset2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(member).To(Equal(ZWithKey{
				Z: Z{
					Score:  3,
					Member: "three",
				},
				Key: "zset1",
			}))
		})

		It("should BZPopMax blocks", func() {
			started := make(chan bool)
			done := make(chan bool)
			go func() {
				defer GinkgoRecover()

				started <- true
				bZPopMax := adapter.BZPopMax(ctx, 0, "zset")
				Expect(bZPopMax.Err()).NotTo(HaveOccurred())
				Expect(bZPopMax.Val()).To(Equal(ZWithKey{
					Z: Z{
						Member: "a",
						Score:  1,
					},
					Key: "zset",
				}))
				done <- true
			}()
			<-started

			select {
			case <-done:
				Fail("BZPopMax is not blocked")
			case <-time.After(time.Second):
				// ok
			}

			zAdd := adapter.ZAdd(ctx, "zset", Z{
				Member: "a",
				Score:  1,
			})
			Expect(zAdd.Err()).NotTo(HaveOccurred())

			select {
			case <-done:
				// ok
			case <-time.After(time.Second):
				Fail("BZPopMax is still blocked")
			}
		})

		It("should BZPopMax timeout", func() {
			_, err := adapter.BZPopMax(ctx, time.Second, "zset1").Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue())

			Expect(adapter.Ping(ctx).Err()).NotTo(HaveOccurred())
		})

		It("should BZPopMin", func() {
			err := adapter.ZAdd(ctx, "zset1", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{
				Score:  3,
				Member: "three",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			member, err := adapter.BZPopMin(ctx, 0, "zset1", "zset2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(member).To(Equal(ZWithKey{
				Z: Z{
					Score:  1,
					Member: "one",
				},
				Key: "zset1",
			}))
		})

		It("should BZPopMin blocks", func() {
			started := make(chan bool)
			done := make(chan bool)
			go func() {
				defer GinkgoRecover()

				started <- true
				bZPopMin := adapter.BZPopMin(ctx, 0, "zset")
				Expect(bZPopMin.Err()).NotTo(HaveOccurred())
				Expect(bZPopMin.Val()).To(Equal(ZWithKey{
					Z: Z{
						Member: "a",
						Score:  1,
					},
					Key: "zset",
				}))
				done <- true
			}()
			<-started

			select {
			case <-done:
				Fail("BZPopMin is not blocked")
			case <-time.After(time.Second):
				// ok
			}

			zAdd := adapter.ZAdd(ctx, "zset", Z{
				Member: "a",
				Score:  1,
			})
			Expect(zAdd.Err()).NotTo(HaveOccurred())

			select {
			case <-done:
				// ok
			case <-time.After(time.Second):
				Fail("BZPopMin is still blocked")
			}
		})

		It("should BZPopMin timeout", func() {
			_, err := adapter.BZPopMin(ctx, time.Second, "zset1").Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue())

			Expect(adapter.Ping(ctx).Err()).NotTo(HaveOccurred())
		})

		It("should ZAdd", func() {
			added, err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "uno",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "two",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(0)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  1,
				Member: "uno",
			}, {
				Score:  3,
				Member: "two",
			}}))
		})

		It("should ZAdd bytes", func() {
			added, err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "uno",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "two",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(0)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  1,
				Member: "uno",
			}, {
				Score:  3,
				Member: "two",
			}}))
		})

		if resp3 {
			It("should ZAddArgs", func() {
				// Test only the GT+LT options.
				added, err := adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					GT:      true,
					Members: []Z{{Score: 1, Member: "one"}},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(1)))

				vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]Z{{Score: 1, Member: "one"}}))

				added, err = adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					GT:      true,
					Members: []Z{{Score: 2, Member: "one"}},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]Z{{Score: 2, Member: "one"}}))

				added, err = adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					LT:      true,
					Members: []Z{{Score: 1, Member: "one"}},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]Z{{Score: 1, Member: "one"}}))
			})
		}

		if resp3 {
			It("should ZAddArgsLT", func() {
				added, err := adapter.ZAddLT(ctx, "zset", Z{
					Score:  2,
					Member: "one",
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(1)))

				vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]Z{{Score: 2, Member: "one"}}))

				added, err = adapter.ZAddLT(ctx, "zset", Z{
					Score:  3,
					Member: "one",
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]Z{{Score: 2, Member: "one"}}))

				added, err = adapter.ZAddLT(ctx, "zset", Z{
					Score:  1,
					Member: "one",
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]Z{{Score: 1, Member: "one"}}))
			})

			It("should ZAddArgsGT", func() {
				added, err := adapter.ZAddGT(ctx, "zset", Z{
					Score:  2,
					Member: "one",
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(1)))

				vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]Z{{Score: 2, Member: "one"}}))

				added, err = adapter.ZAddGT(ctx, "zset", Z{
					Score:  3,
					Member: "one",
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]Z{{Score: 3, Member: "one"}}))

				added, err = adapter.ZAddGT(ctx, "zset", Z{
					Score:  1,
					Member: "one",
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]Z{{Score: 3, Member: "one"}}))
			})
		}

		It("should ZAddNX", func() {
			added, err := adapter.ZAddNX(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(1)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{Score: 1, Member: "one"}}))

			added, err = adapter.ZAddNX(ctx, "zset", Z{
				Score:  2,
				Member: "one",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(0)))

			vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{Score: 1, Member: "one"}}))
		})

		It("should ZAddXX", func() {
			added, err := adapter.ZAddXX(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(0)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(BeEmpty())

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(1)))

			added, err = adapter.ZAddXX(ctx, "zset", Z{
				Score:  2,
				Member: "one",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(0)))

			vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{Score: 2, Member: "one"}}))
		})

		It("should ZCard", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			card, err := adapter.ZCard(ctx, "zset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(card).To(Equal(int64(2)))
		})

		It("should ZCount", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			count, err := adapter.ZCount(ctx, "zset", "-inf", "+inf").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(int64(3)))

			count, err = adapter.ZCount(ctx, "zset", "(1", "3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(int64(2)))

			count, err = adapter.ZLexCount(ctx, "zset", "-", "+").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(int64(3)))
		})

		It("should ZIncrBy", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			n, err := adapter.ZIncrBy(ctx, "zset", 2, "one").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(float64(3)))

			val, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "one",
			}}))
		})

		It("should ZInterStore", func() {
			err := adapter.ZAdd(ctx, "zset1", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset3", Z{Score: 3, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())

			n, err := adapter.ZInterStore(ctx, "out", ZStore{
				Keys:    []string{"zset1", "zset2"},
				Weights: []int64{2, 3},
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(2)))

			vals, err := adapter.ZRangeWithScores(ctx, "out", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  5,
				Member: "one",
			}, {
				Score:  10,
				Member: "two",
			}}))
		})

		if resp3 {
			It("should ZMScore", func() {
				zmScore := adapter.ZMScore(ctx, "zset", "one", "three")
				Expect(zmScore.Err()).NotTo(HaveOccurred())
				Expect(zmScore.Val()).To(HaveLen(2))
				Expect(zmScore.Val()[0]).To(Equal(float64(0)))

				err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
				Expect(err).NotTo(HaveOccurred())

				zmScore = adapter.ZMScore(ctx, "zset", "one", "three")
				Expect(zmScore.Err()).NotTo(HaveOccurred())
				Expect(zmScore.Val()).To(HaveLen(2))
				Expect(zmScore.Val()[0]).To(Equal(float64(1)))

				zmScore = adapter.ZMScore(ctx, "zset", "four")
				Expect(zmScore.Err()).NotTo(HaveOccurred())
				Expect(zmScore.Val()).To(HaveLen(1))

				zmScore = adapter.ZMScore(ctx, "zset", "four", "one")
				Expect(zmScore.Err()).NotTo(HaveOccurred())
				Expect(zmScore.Val()).To(HaveLen(2))
			})
		}

		It("should ZPopMax", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			members, err := adapter.ZPopMax(ctx, "zset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}}))

			// adding back 3
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			members, err = adapter.ZPopMax(ctx, "zset", 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}}))

			// adding back 2 & 3
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			members, err = adapter.ZPopMax(ctx, "zset", 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))
		})

		It("should ZPopMin", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			members, err := adapter.ZPopMin(ctx, "zset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}}))

			// adding back 1
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			members, err = adapter.ZPopMin(ctx, "zset", 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  2,
				Member: "two",
			}}))

			// adding back 1 & 2
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			members, err = adapter.ZPopMin(ctx, "zset", 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(members).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))
		})

		It("should ZRange", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRange := adapter.ZRange(ctx, "zset", 0, -1)
			Expect(zRange.Err()).NotTo(HaveOccurred())
			Expect(zRange.Val()).To(Equal([]string{"one", "two", "three"}))

			zRange = adapter.ZRange(ctx, "zset", 2, 3)
			Expect(zRange.Err()).NotTo(HaveOccurred())
			Expect(zRange.Val()).To(Equal([]string{"three"}))

			zRange = adapter.ZRange(ctx, "zset", -2, -1)
			Expect(zRange.Err()).NotTo(HaveOccurred())
			Expect(zRange.Val()).To(Equal([]string{"two", "three"}))
		})

		It("should ZRangeWithScores", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))

			vals, err = adapter.ZRangeWithScores(ctx, "zset", 2, 3).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{Score: 3, Member: "three"}}))

			vals, err = adapter.ZRangeWithScores(ctx, "zset", -2, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))
		})

		if resp3 {
			It("should ZRangeArgs", func() {
				added, err := adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					Members: []Z{
						{Score: 1, Member: "one"},
						{Score: 2, Member: "two"},
						{Score: 3, Member: "three"},
						{Score: 4, Member: "four"},
					},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(4)))

				zRange, err := adapter.ZRangeArgs(ctx, ZRangeArgs{
					Key:     "zset",
					Start:   1,
					Stop:    4,
					ByScore: true,
					Rev:     true,
					Offset:  1,
					Count:   2,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(zRange).To(Equal([]string{"three", "two"}))

				zRange, err = adapter.ZRangeArgs(ctx, ZRangeArgs{
					Key:    "zset",
					Start:  "-",
					Stop:   "+",
					ByLex:  true,
					Rev:    true,
					Offset: 2,
					Count:  2,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(zRange).To(Equal([]string{"two", "one"}))

				zRange, err = adapter.ZRangeArgs(ctx, ZRangeArgs{
					Key:     "zset",
					Start:   "(1",
					Stop:    "(4",
					ByScore: true,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(zRange).To(Equal([]string{"two", "three"}))

				// withScores.
				zSlice, err := adapter.ZRangeArgsWithScores(ctx, ZRangeArgs{
					Key:     "zset",
					Start:   1,
					Stop:    4,
					ByScore: true,
					Rev:     true,
					Offset:  1,
					Count:   2,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(zSlice).To(Equal([]Z{
					{Score: 3, Member: "three"},
					{Score: 2, Member: "two"},
				}))
			})
		}

		It("should ZRangeByScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRangeByScore := adapter.ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "-inf",
				Max: "+inf",
			})
			Expect(zRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRangeByScore.Val()).To(Equal([]string{"one", "two", "three"}))

			zRangeByScore = adapter.ZRangeByScore(ctx, "zset", ZRangeBy{
				Min:    "-inf",
				Max:    "+inf",
				Offset: 1,
				Count:  2,
			})
			Expect(zRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRangeByScore.Val()).To(Equal([]string{"two", "three"}))

			zRangeByScore = adapter.ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "1",
				Max: "2",
			})
			Expect(zRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRangeByScore.Val()).To(Equal([]string{"one", "two"}))

			zRangeByScore = adapter.ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "2",
			})
			Expect(zRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRangeByScore.Val()).To(Equal([]string{"two"}))

			zRangeByScore = adapter.ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "(2",
			})
			Expect(zRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRangeByScore.Val()).To(Equal([]string{}))
		})

		It("should ZRangeByLex", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "a",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "b",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "c",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRangeByLex := adapter.ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "-",
				Max: "+",
			})
			Expect(zRangeByLex.Err()).NotTo(HaveOccurred())
			Expect(zRangeByLex.Val()).To(Equal([]string{"a", "b", "c"}))

			zRangeByLex = adapter.ZRangeByLex(ctx, "zset", ZRangeBy{
				Min:    "-",
				Max:    "+",
				Offset: 1,
				Count:  2,
			})
			Expect(zRangeByLex.Err()).NotTo(HaveOccurred())
			Expect(zRangeByLex.Val()).To(Equal([]string{"b", "c"}))

			zRangeByLex = adapter.ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "[a",
				Max: "[b",
			})
			Expect(zRangeByLex.Err()).NotTo(HaveOccurred())
			Expect(zRangeByLex.Val()).To(Equal([]string{"a", "b"}))

			zRangeByLex = adapter.ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "(a",
				Max: "[b",
			})
			Expect(zRangeByLex.Err()).NotTo(HaveOccurred())
			Expect(zRangeByLex.Val()).To(Equal([]string{"b"}))

			zRangeByLex = adapter.ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "(a",
				Max: "(b",
			})
			Expect(zRangeByLex.Err()).NotTo(HaveOccurred())
			Expect(zRangeByLex.Val()).To(Equal([]string{}))
		})

		It("should ZRangeByScoreWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "-inf",
				Max: "+inf",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))

			vals, err = adapter.ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min:    "-inf",
				Max:    "+inf",
				Offset: 1,
				Count:  2,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))

			vals, err = adapter.ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "1",
				Max: "2",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  2,
				Member: "two",
			}}))

			vals, err = adapter.ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "2",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{Score: 2, Member: "two"}}))

			vals, err = adapter.ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "(2",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{}))
		})

		if resp3 {
			It("should ZRangeStore", func() {
				added, err := adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					Members: []Z{
						{Score: 1, Member: "one"},
						{Score: 2, Member: "two"},
						{Score: 3, Member: "three"},
						{Score: 4, Member: "four"},
					},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(4)))

				rangeStore, err := adapter.ZRangeStore(ctx, "new-zset", ZRangeArgs{
					Key:    "zset",
					Start:  "-",
					Stop:   "+",
					ByLex:  true,
					Rev:    false,
					Offset: 1,
					Count:  2,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(rangeStore).To(Equal(int64(2)))

				zRange, err := adapter.ZRange(ctx, "new-zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(zRange).To(Equal([]string{"two", "three"}))
			})
			It("should ZRangeStore Rev", func() {
				added, err := adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					Members: []Z{
						{Score: 1, Member: "one"},
						{Score: 2, Member: "two"},
						{Score: 3, Member: "three"},
						{Score: 4, Member: "four"},
					},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(added).To(Equal(int64(4)))

				rangeStore, err := adapter.ZRangeStore(ctx, "new-zset", ZRangeArgs{
					Key:     "zset",
					Start:   1,
					Stop:    4,
					ByScore: true,
					Rev:     true,
					Offset:  1,
					Count:   2,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(rangeStore).To(Equal(int64(2)))

				zRange, err := adapter.ZRange(ctx, "new-zset", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(zRange).To(Equal([]string{"two", "three"}))
			})
		}

		It("should ZRank", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRank := adapter.ZRank(ctx, "zset", "three")
			Expect(zRank.Err()).NotTo(HaveOccurred())
			Expect(zRank.Val()).To(Equal(int64(2)))

			zRank = adapter.ZRank(ctx, "zset", "four")
			Expect(rueidis.IsRedisNil(zRank.Err())).To(BeTrue())
			Expect(zRank.Val()).To(Equal(int64(0)))
		})

		if resp3 {
			It("should ZRankWithScore", func() {
				err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
				Expect(err).NotTo(HaveOccurred())

				zRankWithScore := adapter.ZRankWithScore(ctx, "zset", "one")
				Expect(zRankWithScore.Err()).NotTo(HaveOccurred())
				Expect(zRankWithScore.Result()).To(Equal(RankScore{Rank: 0, Score: 1}))

				zRankWithScore = adapter.ZRankWithScore(ctx, "zset", "two")
				Expect(zRankWithScore.Err()).NotTo(HaveOccurred())
				Expect(zRankWithScore.Result()).To(Equal(RankScore{Rank: 1, Score: 2}))

				zRankWithScore = adapter.ZRankWithScore(ctx, "zset", "three")
				Expect(zRankWithScore.Err()).NotTo(HaveOccurred())
				Expect(zRankWithScore.Result()).To(Equal(RankScore{Rank: 2, Score: 3}))

				zRankWithScore = adapter.ZRankWithScore(ctx, "zset", "four")
				Expect(zRankWithScore.Err()).To(HaveOccurred())
				Expect(zRankWithScore.Err()).To(Equal(rueidis.Nil))
			})
		}

		It("should ZRem", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRem := adapter.ZRem(ctx, "zset", "two")
			Expect(zRem.Err()).NotTo(HaveOccurred())
			Expect(zRem.Val()).To(Equal(int64(1)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  3,
				Member: "three",
			}}))
		})

		It("should ZRemRangeByRank", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRemRangeByRank := adapter.ZRemRangeByRank(ctx, "zset", 0, 1)
			Expect(zRemRangeByRank.Err()).NotTo(HaveOccurred())
			Expect(zRemRangeByRank.Val()).To(Equal(int64(2)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}}))
		})

		It("should ZRemRangeByScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRemRangeByScore := adapter.ZRemRangeByScore(ctx, "zset", "-inf", "(2")
			Expect(zRemRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRemRangeByScore.Val()).To(Equal(int64(1)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))
		})

		It("should ZRemRangeByLex", func() {
			zz := []Z{
				{Score: 0, Member: "aaaa"},
				{Score: 0, Member: "b"},
				{Score: 0, Member: "c"},
				{Score: 0, Member: "d"},
				{Score: 0, Member: "e"},
				{Score: 0, Member: "foo"},
				{Score: 0, Member: "zap"},
				{Score: 0, Member: "zip"},
				{Score: 0, Member: "ALPHA"},
				{Score: 0, Member: "alpha"},
			}
			for _, z := range zz {
				err := adapter.ZAdd(ctx, "zset", z).Err()
				Expect(err).NotTo(HaveOccurred())
			}

			n, err := adapter.ZRemRangeByLex(ctx, "zset", "[alpha", "[omega").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(6)))

			vals, err := adapter.ZRange(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"ALPHA", "aaaa", "zap", "zip"}))
		})

		It("should ZRevRange", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRevRange := adapter.ZRevRange(ctx, "zset", 0, -1)
			Expect(zRevRange.Err()).NotTo(HaveOccurred())
			Expect(zRevRange.Val()).To(Equal([]string{"three", "two", "one"}))

			zRevRange = adapter.ZRevRange(ctx, "zset", 2, 3)
			Expect(zRevRange.Err()).NotTo(HaveOccurred())
			Expect(zRevRange.Val()).To(Equal([]string{"one"}))

			zRevRange = adapter.ZRevRange(ctx, "zset", -2, -1)
			Expect(zRevRange.Err()).NotTo(HaveOccurred())
			Expect(zRevRange.Val()).To(Equal([]string{"two", "one"}))
		})

		It("should ZRevRangeWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			val, err := adapter.ZRevRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))

			val, err = adapter.ZRevRangeWithScores(ctx, "zset", 2, 3).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]Z{{Score: 1, Member: "one"}}))

			val, err = adapter.ZRevRangeWithScores(ctx, "zset", -2, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))
		})

		It("should ZRevRangeByScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"three", "two", "one"}))

			vals, err = adapter.ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf", Offset: 1, Count: 2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"two", "one"}))

			vals, err = adapter.ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "2", Min: "(1"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"two"}))

			vals, err = adapter.ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "(2", Min: "(1"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{}))
		})

		It("should ZRevRangeByLex", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "a"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "b"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "c"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "+", Min: "-"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"c", "b", "a"}))

			vals, err = adapter.ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "+", Min: "-", Offset: 1, Count: 2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"b", "a"}))

			vals, err = adapter.ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "[b", Min: "(a"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"b"}))

			vals, err = adapter.ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "(b", Min: "(a"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{}))
		})

		It("should ZRevRangeByScoreWithScores", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))

			vals, err = adapter.ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf", Offset: 1, Count: 2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))
		})

		It("should ZRevRangeByScoreWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))

			vals, err = adapter.ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "2", Min: "(1"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{Score: 2, Member: "two"}}))

			vals, err = adapter.ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "(2", Min: "(1"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{}))
		})

		It("should ZRevRank", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRevRank := adapter.ZRevRank(ctx, "zset", "one")
			Expect(zRevRank.Err()).NotTo(HaveOccurred())
			Expect(zRevRank.Val()).To(Equal(int64(2)))

			zRevRank = adapter.ZRevRank(ctx, "zset", "four")
			Expect(rueidis.IsRedisNil(zRevRank.Err())).To(BeTrue())
			Expect(zRevRank.Val()).To(Equal(int64(0)))
		})

		if resp3 {
			It("should ZRevRankWithScore", func() {
				err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
				Expect(err).NotTo(HaveOccurred())

				zRevRankWithScore := adapter.ZRevRankWithScore(ctx, "zset", "one")
				Expect(zRevRankWithScore.Err()).NotTo(HaveOccurred())
				Expect(zRevRankWithScore.Result()).To(Equal(RankScore{Rank: 2, Score: 1}))

				zRevRankWithScore = adapter.ZRevRankWithScore(ctx, "zset", "two")
				Expect(zRevRankWithScore.Err()).NotTo(HaveOccurred())
				Expect(zRevRankWithScore.Result()).To(Equal(RankScore{Rank: 1, Score: 2}))

				zRevRankWithScore = adapter.ZRevRankWithScore(ctx, "zset", "three")
				Expect(zRevRankWithScore.Err()).NotTo(HaveOccurred())
				Expect(zRevRankWithScore.Result()).To(Equal(RankScore{Rank: 0, Score: 3}))

				zRevRankWithScore = adapter.ZRevRankWithScore(ctx, "zset", "four")
				Expect(zRevRankWithScore.Err()).To(HaveOccurred())
				Expect(zRevRankWithScore.Err()).To(Equal(rueidis.Nil))
			})
		}

		It("should ZScore", func() {
			zAdd := adapter.ZAdd(ctx, "zset", Z{Score: 1.001, Member: "one"})
			Expect(zAdd.Err()).NotTo(HaveOccurred())

			zScore := adapter.ZScore(ctx, "zset", "one")
			Expect(zScore.Err()).NotTo(HaveOccurred())
			Expect(zScore.Val()).To(Equal(float64(1.001)))
		})

		if resp3 {
			It("should ZUnion", func() {
				err := adapter.ZAddArgs(ctx, "zset1", ZAddArgs{
					Members: []Z{
						{Score: 1, Member: "one"},
						{Score: 2, Member: "two"},
					},
				}).Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.ZAddArgs(ctx, "zset2", ZAddArgs{
					Members: []Z{
						{Score: 1, Member: "one"},
						{Score: 2, Member: "two"},
						{Score: 3, Member: "three"},
					},
				}).Err()
				Expect(err).NotTo(HaveOccurred())

				union, err := adapter.ZUnion(ctx, ZStore{
					Keys:      []string{"zset1", "zset2"},
					Weights:   []int64{2, 3},
					Aggregate: "sum",
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(union).To(Equal([]string{"one", "three", "two"}))

				unionScores, err := adapter.ZUnionWithScores(ctx, ZStore{
					Keys:      []string{"zset1", "zset2"},
					Weights:   []int64{2, 3},
					Aggregate: "sum",
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(unionScores).To(Equal([]Z{
					{Score: 5, Member: "one"},
					{Score: 9, Member: "three"},
					{Score: 10, Member: "two"},
				}))
			})
		}

		It("should ZUnionStore", func() {
			err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())

			err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			n, err := adapter.ZUnionStore(ctx, "out", ZStore{
				Keys:    []string{"zset1", "zset2"},
				Weights: []int64{2, 3},
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(3)))

			val, err := adapter.ZRangeWithScores(ctx, "out", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]Z{{
				Score:  5,
				Member: "one",
			}, {
				Score:  9,
				Member: "three",
			}, {
				Score:  10,
				Member: "two",
			}}))
		})

		if resp3 {
			It("should ZRandMember", func() {
				err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())

				v := adapter.ZRandMember(ctx, "zset", 1)
				Expect(v.Err()).NotTo(HaveOccurred())
				Expect(v.Val()).To(Or(Equal([]string{"one"}), Equal([]string{"two"})))

				v = adapter.ZRandMember(ctx, "zset", 0)
				Expect(v.Err()).NotTo(HaveOccurred())
				Expect(v.Val()).To(HaveLen(0))

				kv, err := adapter.ZRandMemberWithScores(ctx, "zset", 1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(kv).To(Or(
					Equal([]Z{{Member: "one", Score: 1}}),
					Equal([]Z{{Member: "two", Score: 2}}),
				))
			})

			It("should ZDiff", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 3, Member: "three"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())

				v, err := adapter.ZDiff(ctx, "zset1", "zset2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(v).To(Equal([]string{"two", "three"}))
			})

			It("should ZDiffWithScores", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 3, Member: "three"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())

				v, err := adapter.ZDiffWithScores(ctx, "zset1", "zset2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(v).To(Equal([]Z{
					{
						Member: "two",
						Score:  2,
					},
					{
						Member: "three",
						Score:  3,
					},
				}))
			})

			It("should ZInter", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
				Expect(err).NotTo(HaveOccurred())

				v, err := adapter.ZInter(ctx, ZStore{
					Keys: []string{"zset1", "zset2"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(v).To(Equal([]string{"one", "two"}))
			})

			It("should ZInterCard", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
				Expect(err).NotTo(HaveOccurred())

				// limit 0 means no limit
				sInterCard := adapter.ZInterCard(ctx, 0, "zset1", "zset2")
				Expect(sInterCard.Err()).NotTo(HaveOccurred())
				Expect(sInterCard.Val()).To(Equal(int64(2)))

				sInterCard = adapter.ZInterCard(ctx, 1, "zset1", "zset2")
				Expect(sInterCard.Err()).NotTo(HaveOccurred())
				Expect(sInterCard.Val()).To(Equal(int64(1)))

				sInterCard = adapter.ZInterCard(ctx, 3, "zset1", "zset2")
				Expect(sInterCard.Err()).NotTo(HaveOccurred())
				Expect(sInterCard.Val()).To(Equal(int64(2)))
			})

			It("should ZInterWithScores", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
				Expect(err).NotTo(HaveOccurred())

				v, err := adapter.ZInterWithScores(ctx, ZStore{
					Keys:      []string{"zset1", "zset2"},
					Weights:   []int64{2, 3},
					Aggregate: "Max",
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(v).To(Equal([]Z{
					{
						Member: "one",
						Score:  3,
					},
					{
						Member: "two",
						Score:  6,
					},
				}))
			})

			It("should ZDiffStore", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
				Expect(err).NotTo(HaveOccurred())
				v, err := adapter.ZDiffStore(ctx, "out1", "zset1", "zset2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(v).To(Equal(int64(0)))
				v, err = adapter.ZDiffStore(ctx, "out1", "zset2", "zset1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(v).To(Equal(int64(1)))
				vals, err := adapter.ZRangeWithScores(ctx, "out1", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]Z{{
					Score:  3,
					Member: "three",
				}}))
			})
		}
	})

	Describe("streams", func() {
		BeforeEach(func() {
			if resp3 {
				_, err := adapter.XAdd(ctx, XAddArgs{
					Stream:     "stream",
					ID:         "1-0",
					Values:     map[string]any{"uno": "un"},
					NoMkStream: true,
				}).Result()
				Expect(rueidis.IsRedisNil(err)).To(BeTrue())
			}

			id, err := adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				ID:     "1-0",
				Values: map[string]any{"uno": "un"},
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(id).To(Equal("1-0"))

			// Values supports []any.
			id, err = adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				ID:     "2-0",
				Values: []any{"dos", "deux"},
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(id).To(Equal("2-0"))

			// Value supports []string.
			id, err = adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				ID:     "3-0",
				Values: []string{"tres", "troix"},
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(id).To(Equal("3-0"))
		})

		It("should XTrimMaxLen", func() {
			n, err := adapter.XTrimMaxLen(ctx, "stream", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(3)))
		})

		It("should XTrimMaxLenApprox", func() {
			n, err := adapter.XTrimMaxLenApprox(ctx, "stream", 0, 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(3)))
		})

		if resp3 {
			It("should XTrimMaxLenApprox Limit", func() {
				n, err := adapter.XTrimMaxLenApprox(ctx, "stream", 0, 1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(0)))
			})

			It("should XTrimMinID", func() {
				n, err := adapter.XTrimMinID(ctx, "stream", "4-0").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(3)))
			})

			It("should XTrimMinIDApprox", func() {
				n, err := adapter.XTrimMinIDApprox(ctx, "stream", "4-0", 0).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(3)))
			})
		}

		It("should XAdd", func() {
			id, err := adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				Values: map[string]any{"quatro": "quatre"},
			}).Result()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
				{ID: id, Values: map[string]any{"quatro": "quatre"}},
			}))
		})

		// TODO XAdd There is a bug in the limit parameter.
		// TODO Don't test it for now.
		// TODO link: https://github.com/redis/redis/issues/9046
		It("should XAdd with MaxLen", func() {
			id, err := adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				MaxLen: 1,
				Values: map[string]any{"quatro": "quatre"},
			}).Result()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]XMessage{
				{ID: id, Values: map[string]any{"quatro": "quatre"}},
			}))
		})

		It("should XAdd with MaxLen Approx", func() {
			id, err := adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				MaxLen: 1,
				Approx: true,
				Values: map[string]any{"quatro": "quatre"},
			}).Result()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
				{ID: id, Values: map[string]any{"quatro": "quatre"}},
			}))
		})

		if resp3 {
			It("should XAdd with MinID", func() {
				id, err := adapter.XAdd(ctx, XAddArgs{
					Stream: "stream",
					MinID:  "5-0",
					ID:     "4-0",
					Values: map[string]any{"quatro": "quatre"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(id).To(Equal("4-0"))

				vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(HaveLen(0))
			})

			It("should XAdd with MinID Approx", func() {
				id, err := adapter.XAdd(ctx, XAddArgs{
					Stream: "stream",
					MinID:  "5-0",
					ID:     "4-0",
					Approx: true,
					Values: map[string]any{"quatro": "quatre"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(id).To(Equal("4-0"))

				vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(HaveLen(0))
			})

			It("should XAdd with MinID Limit", func() {
				id, err := adapter.XAdd(ctx, XAddArgs{
					Stream: "stream",
					MinID:  "5-0",
					ID:     "4-0",
					Approx: true,
					Values: map[string]any{"quatro": "quatre"},
					Limit:  1,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(id).To(Equal("4-0"))

				vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]XMessage{
					{ID: "1-0", Values: map[string]any{"uno": "un"}},
					{ID: "2-0", Values: map[string]any{"dos": "deux"}},
					{ID: "3-0", Values: map[string]any{"tres": "troix"}},
					{ID: id, Values: map[string]any{"quatro": "quatre"}},
				}))
			})
		}

		It("should XDel", func() {
			n, err := adapter.XDel(ctx, "stream", "1-0", "2-0", "3-0").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(3)))
		})

		It("should XLen", func() {
			n, err := adapter.XLen(ctx, "stream").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(3)))
		})

		It("should XRange", func() {
			msgs, err := adapter.XRange(ctx, "stream", "-", "+").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
			}))

			msgs, err = adapter.XRange(ctx, "stream", "2", "+").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]XMessage{
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
			}))

			msgs, err = adapter.XRange(ctx, "stream", "-", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
			}))
		})

		It("should XRangeN", func() {
			msgs, err := adapter.XRangeN(ctx, "stream", "-", "+", 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
			}))

			msgs, err = adapter.XRangeN(ctx, "stream", "2", "+", 1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]XMessage{
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
			}))

			msgs, err = adapter.XRangeN(ctx, "stream", "-", "2", 1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
			}))
		})

		It("should XRevRange", func() {
			msgs, err := adapter.XRevRange(ctx, "stream", "+", "-").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]XMessage{
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
			}))

			msgs, err = adapter.XRevRange(ctx, "stream", "+", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]XMessage{
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
			}))
		})

		It("should XRevRangeN", func() {
			msgs, err := adapter.XRevRangeN(ctx, "stream", "+", "-", 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]XMessage{
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
			}))

			msgs, err = adapter.XRevRangeN(ctx, "stream", "+", "2", 1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(msgs).To(Equal([]XMessage{
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
			}))
		})

		It("should XRead", func() {
			res, err := adapter.XReadStreams(ctx, "stream", "0").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal([]XStream{
				{
					Stream: "stream",
					Messages: []XMessage{
						{ID: "1-0", Values: map[string]any{"uno": "un"}},
						{ID: "2-0", Values: map[string]any{"dos": "deux"}},
						{ID: "3-0", Values: map[string]any{"tres": "troix"}},
					},
				},
			}))

			_, err = adapter.XReadStreams(ctx, "stream", "3").Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue())
		})

		It("should XRead", func() {
			res, err := adapter.XRead(ctx, XReadArgs{
				Streams: []string{"stream", "0"},
				Count:   2,
				Block:   100 * time.Millisecond,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal([]XStream{
				{
					Stream: "stream",
					Messages: []XMessage{
						{ID: "1-0", Values: map[string]any{"uno": "un"}},
						{ID: "2-0", Values: map[string]any{"dos": "deux"}},
					},
				},
			}))

			_, err = adapter.XRead(ctx, XReadArgs{
				Streams: []string{"stream", "3"},
				Count:   1,
				Block:   100 * time.Millisecond,
			}).Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue())
		})

		Describe("group", func() {
			BeforeEach(func() {
				err := adapter.XGroupCreate(ctx, "stream", "group", "0").Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.XGroupSetID(ctx, "stream", "group", "0").Err()
				Expect(err).NotTo(HaveOccurred())

				res, err := adapter.XReadGroup(ctx, XReadGroupArgs{
					Group:    "group",
					Consumer: "consumer",
					Streams:  []string{"stream", ">"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]XStream{
					{
						Stream: "stream",
						Messages: []XMessage{
							{ID: "1-0", Values: map[string]any{"uno": "un"}},
							{ID: "2-0", Values: map[string]any{"dos": "deux"}},
							{ID: "3-0", Values: map[string]any{"tres": "troix"}},
						},
					},
				}))
			})

			AfterEach(func() {
				n, err := adapter.XGroupDestroy(ctx, "stream", "group").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(1)))
			})

			It("should XReadGroup skip empty", func() {
				n, err := adapter.XDel(ctx, "stream", "2-0").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(1)))

				res, err := adapter.XReadGroup(ctx, XReadGroupArgs{
					Group:    "group",
					Consumer: "consumer",
					Streams:  []string{"stream", "0"},
					NoAck:    true,
					Block:    -1,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]XStream{
					{
						Stream: "stream",
						Messages: []XMessage{
							{ID: "1-0", Values: map[string]any{"uno": "un"}},
							{ID: "2-0", Values: nil},
							{ID: "3-0", Values: map[string]any{"tres": "troix"}},
						},
					},
				}))
			})

			It("should XGroupCreateMkStream", func() {
				err := adapter.XGroupCreateMkStream(ctx, "stream2", "group", "0").Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.XGroupCreateMkStream(ctx, "stream2", "group", "0").Err()
				Expect(err.Error()).To(Equal("BUSYGROUP Consumer Group name already exists"))

				n, err := adapter.XGroupDestroy(ctx, "stream2", "group").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(1)))

				n, err = adapter.Del(ctx, "stream2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(1)))
			})

			if resp3 {
				It("should XPending", func() {
					info, err := adapter.XPending(ctx, "stream", "group").Result()
					Expect(err).NotTo(HaveOccurred())
					Expect(info).To(Equal(XPending{
						Count:     3,
						Lower:     "1-0",
						Higher:    "3-0",
						Consumers: map[string]int64{"consumer": 3},
					}))
					args := XPendingExtArgs{
						Stream:   "stream",
						Group:    "group",
						Start:    "-",
						End:      "+",
						Count:    10,
						Consumer: "consumer",
					}
					infoExt, err := adapter.XPendingExt(ctx, args).Result()
					Expect(err).NotTo(HaveOccurred())
					for i := range infoExt {
						infoExt[i].Idle = 0
					}
					Expect(infoExt).To(Equal([]XPendingExt{
						{ID: "1-0", Consumer: "consumer", Idle: 0, RetryCount: 1},
						{ID: "2-0", Consumer: "consumer", Idle: 0, RetryCount: 1},
						{ID: "3-0", Consumer: "consumer", Idle: 0, RetryCount: 1},
					}))

					args.Idle = 72 * time.Hour
					infoExt, err = adapter.XPendingExt(ctx, args).Result()
					Expect(err).NotTo(HaveOccurred())
					Expect(infoExt).To(HaveLen(0))
				})

				It("should XGroup Create Delete Consumer", func() {
					n, err := adapter.XGroupCreateConsumer(ctx, "stream", "group", "c1").Result()
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(int64(1)))

					n, err = adapter.XGroupDelConsumer(ctx, "stream", "group", "consumer").Result()
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(int64(3)))
				})

				It("should XAutoClaim", func() {
					xca := XAutoClaimArgs{
						Stream:   "stream",
						Group:    "group",
						Consumer: "consumer",
						Start:    "-",
						Count:    2,
					}
					msgs, start, err := adapter.XAutoClaim(ctx, xca).Result()
					Expect(err).NotTo(HaveOccurred())
					Expect(start).To(Equal("3-0"))
					Expect(msgs).To(Equal([]XMessage{{
						ID:     "1-0",
						Values: map[string]any{"uno": "un"},
					}, {
						ID:     "2-0",
						Values: map[string]any{"dos": "deux"},
					}}))

					xca.Start = start
					msgs, start, err = adapter.XAutoClaim(ctx, xca).Result()
					Expect(err).NotTo(HaveOccurred())
					Expect(start).To(Equal("0-0"))
					Expect(msgs).To(Equal([]XMessage{{
						ID:     "3-0",
						Values: map[string]any{"tres": "troix"},
					}}))

					ids, start, err := adapter.XAutoClaimJustID(ctx, xca).Result()
					Expect(err).NotTo(HaveOccurred())
					Expect(start).To(Equal("0-0"))
					Expect(ids).To(Equal([]string{"3-0"}))
				})

				It("should XAutoClaim NoCount", func() {
					xca := XAutoClaimArgs{
						Stream:   "stream",
						Group:    "group",
						Consumer: "consumer",
						Start:    "-",
					}
					msgs, start, err := adapter.XAutoClaim(ctx, xca).Result()
					Expect(err).NotTo(HaveOccurred())
					Expect(start).To(Equal("0-0"))
					Expect(msgs).To(Equal([]XMessage{{
						ID:     "1-0",
						Values: map[string]any{"uno": "un"},
					}, {
						ID:     "2-0",
						Values: map[string]any{"dos": "deux"},
					}, {
						ID:     "3-0",
						Values: map[string]any{"tres": "troix"},
					}}))

					ids, start, err := adapter.XAutoClaimJustID(ctx, xca).Result()
					Expect(err).NotTo(HaveOccurred())
					Expect(start).To(Equal("0-0"))
					Expect(ids).To(Equal([]string{"1-0", "2-0", "3-0"}))
				})
			}

			It("should XClaim", func() {
				msgs, err := adapter.XClaim(ctx, XClaimArgs{
					Stream:   "stream",
					Group:    "group",
					Consumer: "consumer",
					Messages: []string{"1-0", "2-0", "3-0"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(msgs).To(Equal([]XMessage{{
					ID:     "1-0",
					Values: map[string]any{"uno": "un"},
				}, {
					ID:     "2-0",
					Values: map[string]any{"dos": "deux"},
				}, {
					ID:     "3-0",
					Values: map[string]any{"tres": "troix"},
				}}))

				ids, err := adapter.XClaimJustID(ctx, XClaimArgs{
					Stream:   "stream",
					Group:    "group",
					Consumer: "consumer",
					Messages: []string{"1-0", "2-0", "3-0"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(ids).To(Equal([]string{"1-0", "2-0", "3-0"}))
			})

			It("should XAck", func() {
				n, err := adapter.XAck(ctx, "stream", "group", "1-0", "2-0", "4-0").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(2)))
			})
		})

		Describe("xinfo", func() {
			BeforeEach(func() {
				err := adapter.XGroupCreate(ctx, "stream", "group1", "0").Err()
				Expect(err).NotTo(HaveOccurred())

				res, err := adapter.XReadGroup(ctx, XReadGroupArgs{
					Group:    "group1",
					Consumer: "consumer1",
					Streams:  []string{"stream", ">"},
					Count:    2,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]XStream{
					{
						Stream: "stream",
						Messages: []XMessage{
							{ID: "1-0", Values: map[string]any{"uno": "un"}},
							{ID: "2-0", Values: map[string]any{"dos": "deux"}},
						},
					},
				}))

				res, err = adapter.XReadGroup(ctx, XReadGroupArgs{
					Group:    "group1",
					Consumer: "consumer2",
					Streams:  []string{"stream", ">"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]XStream{
					{
						Stream: "stream",
						Messages: []XMessage{
							{ID: "3-0", Values: map[string]any{"tres": "troix"}},
						},
					},
				}))

				err = adapter.XGroupCreate(ctx, "stream", "group2", "1-0").Err()
				Expect(err).NotTo(HaveOccurred())

				res, err = adapter.XReadGroup(ctx, XReadGroupArgs{
					Group:    "group2",
					Consumer: "consumer1",
					Streams:  []string{"stream", ">"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]XStream{
					{
						Stream: "stream",
						Messages: []XMessage{
							{ID: "2-0", Values: map[string]any{"dos": "deux"}},
							{ID: "3-0", Values: map[string]any{"tres": "troix"}},
						},
					},
				}))
			})

			AfterEach(func() {
				n, err := adapter.XGroupDestroy(ctx, "stream", "group1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(1)))
				n, err = adapter.XGroupDestroy(ctx, "stream", "group2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(1)))
			})

			It("should XINFO STREAM", func() {
				res, err := adapter.XInfoStream(ctx, "stream").Result()
				Expect(err).NotTo(HaveOccurred())
				res.RadixTreeKeys = 0
				res.RadixTreeNodes = 0

				if resp3 {
					Expect(res).To(Equal(XInfoStream{
						Length:            3,
						RadixTreeKeys:     0,
						RadixTreeNodes:    0,
						Groups:            2,
						LastGeneratedID:   "3-0",
						MaxDeletedEntryID: "0-0",
						EntriesAdded:      3,
						FirstEntry: XMessage{
							ID:     "1-0",
							Values: map[string]any{"uno": "un"},
						},
						LastEntry: XMessage{
							ID:     "3-0",
							Values: map[string]any{"tres": "troix"},
						},
						RecordedFirstEntryID: "1-0",
					}))
				} else {
					Expect(res).To(Equal(XInfoStream{
						Length:          3,
						RadixTreeKeys:   0,
						RadixTreeNodes:  0,
						Groups:          2,
						LastGeneratedID: "3-0",
						FirstEntry: XMessage{
							ID:     "1-0",
							Values: map[string]any{"uno": "un"},
						},
						LastEntry: XMessage{
							ID:     "3-0",
							Values: map[string]any{"tres": "troix"},
						},
					}))
				}

				// stream is empty
				n, err := adapter.XDel(ctx, "stream", "1-0", "2-0", "3-0").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(Equal(int64(3)))

				res, err = adapter.XInfoStream(ctx, "stream").Result()
				Expect(err).NotTo(HaveOccurred())
				res.RadixTreeKeys = 0
				res.RadixTreeNodes = 0

				if resp3 {
					Expect(res).To(Equal(XInfoStream{
						Length:               0,
						RadixTreeKeys:        0,
						RadixTreeNodes:       0,
						Groups:               2,
						LastGeneratedID:      "3-0",
						MaxDeletedEntryID:    "3-0",
						EntriesAdded:         3,
						FirstEntry:           XMessage{},
						LastEntry:            XMessage{},
						RecordedFirstEntryID: "0-0",
					}))
				} else {
					Expect(res).To(Equal(XInfoStream{
						Length:          0,
						RadixTreeKeys:   0,
						RadixTreeNodes:  0,
						Groups:          2,
						LastGeneratedID: "3-0",
						FirstEntry:      XMessage{},
						LastEntry:       XMessage{},
					}))
				}
			})

			if resp3 {
				It("should XINFO STREAM FULL", func() {
					res, err := adapter.XInfoStreamFull(ctx, "stream", 2).Result()
					Expect(err).NotTo(HaveOccurred())
					res.RadixTreeKeys = 0
					res.RadixTreeNodes = 0

					// Verify DeliveryTime
					now := time.Now()
					maxElapsed := 10 * time.Minute
					for k, g := range res.Groups {
						for k2, p := range g.Pending {
							Expect(now.Sub(p.DeliveryTime)).To(BeNumerically("<=", maxElapsed))
							res.Groups[k].Pending[k2].DeliveryTime = time.Time{}
						}
						for k3, c := range g.Consumers {
							Expect(now.Sub(c.SeenTime)).To(BeNumerically("<=", maxElapsed))
							res.Groups[k].Consumers[k3].SeenTime = time.Time{}

							for k4, p := range c.Pending {
								Expect(now.Sub(p.DeliveryTime)).To(BeNumerically("<=", maxElapsed))
								res.Groups[k].Consumers[k3].Pending[k4].DeliveryTime = time.Time{}
							}
						}
					}

					Expect(res).To(Equal(XInfoStreamFull{
						Length:            3,
						RadixTreeKeys:     0,
						RadixTreeNodes:    0,
						LastGeneratedID:   "3-0",
						MaxDeletedEntryID: "0-0",
						EntriesAdded:      3,
						Entries: []XMessage{
							{ID: "1-0", Values: map[string]any{"uno": "un"}},
							{ID: "2-0", Values: map[string]any{"dos": "deux"}},
						},
						Groups: []XInfoStreamGroup{
							{
								Name:            "group1",
								LastDeliveredID: "3-0",
								PelCount:        3,
								EntriesRead:     3,
								Pending: []XInfoStreamGroupPending{
									{
										ID:            "1-0",
										Consumer:      "consumer1",
										DeliveryTime:  time.Time{},
										DeliveryCount: 1,
									},
									{
										ID:            "2-0",
										Consumer:      "consumer1",
										DeliveryTime:  time.Time{},
										DeliveryCount: 1,
									},
								},
								Consumers: []XInfoStreamConsumer{
									{
										Name:     "consumer1",
										SeenTime: time.Time{},
										PelCount: 2,
										Pending: []XInfoStreamConsumerPending{
											{
												ID:            "1-0",
												DeliveryTime:  time.Time{},
												DeliveryCount: 1,
											},
											{
												ID:            "2-0",
												DeliveryTime:  time.Time{},
												DeliveryCount: 1,
											},
										},
									},
									{
										Name:     "consumer2",
										SeenTime: time.Time{},
										PelCount: 1,
										Pending: []XInfoStreamConsumerPending{
											{
												ID:            "3-0",
												DeliveryTime:  time.Time{},
												DeliveryCount: 1,
											},
										},
									},
								},
							},
							{
								Name:            "group2",
								LastDeliveredID: "3-0",
								PelCount:        2,
								EntriesRead:     3,
								Pending: []XInfoStreamGroupPending{
									{
										ID:            "2-0",
										Consumer:      "consumer1",
										DeliveryTime:  time.Time{},
										DeliveryCount: 1,
									},
									{
										ID:            "3-0",
										Consumer:      "consumer1",
										DeliveryTime:  time.Time{},
										DeliveryCount: 1,
									},
								},
								Consumers: []XInfoStreamConsumer{
									{
										Name:     "consumer1",
										SeenTime: time.Time{},
										PelCount: 2,
										Pending: []XInfoStreamConsumerPending{
											{
												ID:            "2-0",
												DeliveryTime:  time.Time{},
												DeliveryCount: 1,
											},
											{
												ID:            "3-0",
												DeliveryTime:  time.Time{},
												DeliveryCount: 1,
											},
										},
									},
								},
							},
						},
						RecordedFirstEntryID: "1-0",
					}))
				})
			}

			It("should XINFO GROUPS", func() {
				res, err := adapter.XInfoGroups(ctx, "stream").Result()
				Expect(err).NotTo(HaveOccurred())
				if resp3 {
					Expect(res).To(Equal([]XInfoGroup{
						{Name: "group1", Consumers: 2, Pending: 3, LastDeliveredID: "3-0", EntriesRead: 3},
						{Name: "group2", Consumers: 1, Pending: 2, LastDeliveredID: "3-0", EntriesRead: 3},
					}))
				} else {
					Expect(res).To(Equal([]XInfoGroup{
						{Name: "group1", Consumers: 2, Pending: 3, LastDeliveredID: "3-0"},
						{Name: "group2", Consumers: 1, Pending: 2, LastDeliveredID: "3-0"},
					}))
				}
			})

			It("should XINFO CONSUMERS", func() {
				time.Sleep(time.Millisecond * 2) // make consumer idle > 0
				res, err := adapter.XInfoConsumers(ctx, "stream", "group1").Result()
				Expect(err).NotTo(HaveOccurred())
				for i := range res {
					Expect(res[i].Idle > 0).To(BeTrue())
					res[i].Idle = 0
				}
				Expect(res).To(Equal([]XInfoConsumer{
					{Name: "consumer1", Pending: 2, Idle: 0},
					{Name: "consumer2", Pending: 1, Idle: 0},
				}))
			})
		})
	})

	Describe("Geo add and radius search", func() {
		BeforeEach(func() {
			n, err := adapter.GeoAdd(
				ctx,
				"Sicily",
				GeoLocation{Longitude: 13.361389, Latitude: 38.115556, Name: "Palermo"},
				GeoLocation{Longitude: 15.087269, Latitude: 37.502669, Name: "Catania"},
			).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(2)))
		})

		It("should not add same geo location", func() {
			geoAdd := adapter.GeoAdd(
				ctx,
				"Sicily",
				GeoLocation{Longitude: 13.361389, Latitude: 38.115556, Name: "Palermo"},
			)
			Expect(geoAdd.Err()).NotTo(HaveOccurred())
			Expect(geoAdd.Val()).To(Equal(int64(0)))
		})

		It("should search geo radius", func() {
			res, err := adapter.GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius: 200,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveLen(2))
			Expect(res[0].Name).To(Equal("Palermo"))
			Expect(res[1].Name).To(Equal("Catania"))
		})

		It("should geo radius and store the result", func() {
			n, err := adapter.GeoRadiusStore(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius: 200,
				Store:  "result",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(2)))

			res, err := adapter.ZRangeWithScores(ctx, "result", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(ContainElement(Z{
				Score:  3.479099956230698e+15,
				Member: "Palermo",
			}))
			Expect(res).To(ContainElement(Z{
				Score:  3.479447370796909e+15,
				Member: "Catania",
			}))
		})

		It("should geo radius and store dist", func() {
			n, err := adapter.GeoRadiusStore(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius:    200,
				StoreDist: "result",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(2)))

			res, err := adapter.ZRangeWithScores(ctx, "result", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			for i, z := range res {
				res[i].Score = float64(int(z.Score))
			}
			Expect(res).To(ContainElement(Z{
				Score:  190.,
				Member: "Palermo",
			}))
			Expect(res).To(ContainElement(Z{
				Score:  56.,
				Member: "Catania",
			}))
		})

		It("should search geo radius with options", func() {
			res, err := adapter.GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveLen(2))
			Expect(res[1].Name).To(Equal("Palermo"))
			Expect(res[1].Dist).To(Equal(190.4424))
			Expect(res[1].GeoHash).To(Equal(int64(3479099956230698)))
			Expect(res[1].Longitude).To(Equal(13.361389338970184))
			Expect(res[1].Latitude).To(Equal(38.115556395496299))
			Expect(res[0].Name).To(Equal("Catania"))
			Expect(res[0].Dist).To(Equal(56.4413))
			Expect(res[0].GeoHash).To(Equal(int64(3479447370796909)))
			Expect(res[0].Longitude).To(Equal(15.087267458438873))
			Expect(res[0].Latitude).To(Equal(37.50266842333162))
		})

		It("should search geo radius with WithDist=false", func() {
			res, err := adapter.GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveLen(2))
			Expect(res[1].Name).To(Equal("Palermo"))
			Expect(res[1].Dist).To(Equal(float64(0)))
			Expect(res[1].GeoHash).To(Equal(int64(3479099956230698)))
			Expect(res[1].Longitude).To(Equal(13.361389338970184))
			Expect(res[1].Latitude).To(Equal(38.115556395496299))
			Expect(res[0].Name).To(Equal("Catania"))
			Expect(res[0].Dist).To(Equal(float64(0)))
			Expect(res[0].GeoHash).To(Equal(int64(3479447370796909)))
			Expect(res[0].Longitude).To(Equal(15.087267458438873))
			Expect(res[0].Latitude).To(Equal(37.50266842333162))
		})

		It("should search geo radius by member with options", func() {
			res, err := adapter.GeoRadiusByMember(ctx, "Sicily", "Catania", GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveLen(2))
			Expect(res[0].Name).To(Equal("Catania"))
			Expect(res[0].Dist).To(Equal(0.0))
			Expect(res[0].GeoHash).To(Equal(int64(3479447370796909)))
			Expect(res[0].Longitude).To(Equal(15.087267458438873))
			Expect(res[0].Latitude).To(Equal(37.50266842333162))
			Expect(res[1].Name).To(Equal("Palermo"))
			Expect(res[1].Dist).To(Equal(166.2742))
			Expect(res[1].GeoHash).To(Equal(int64(3479099956230698)))
			Expect(res[1].Longitude).To(Equal(13.361389338970184))
			Expect(res[1].Latitude).To(Equal(38.115556395496299))

			ress, err := adapter.GeoRadiusByMemberStore(ctx, "Sicily", "Catania", GeoRadiusQuery{
				Radius: 200,
				Unit:   "km",
				Count:  2,
				Store:  "Sicily2",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(ress).To(Equal(int64(2)))
		})

		It("should search geo radius with no results", func() {
			res, err := adapter.GeoRadius(ctx, "Sicily", 99, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveLen(0))
		})

		It("should get geo distance with unit options", func() {
			// From Redis CLI, note the difference in rounding in m vs
			// km on Redis itself.
			//
			// GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
			// GEODIST Sicily Palermo Catania m
			// "166274.15156960033"
			// GEODIST Sicily Palermo Catania km
			// "166.27415156960032"
			dist, err := adapter.GeoDist(ctx, "Sicily", "Palermo", "Catania", "km").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(dist).To(BeNumerically("~", 166.27, 0.01))

			dist, err = adapter.GeoDist(ctx, "Sicily", "Palermo", "Catania", "m").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(dist).To(BeNumerically("~", 166274.15, 0.01))

			_, err = adapter.GeoDist(ctx, "Sicily", "Palermo", "Catania", "mi").Result()
			Expect(err).NotTo(HaveOccurred())

			_, err = adapter.GeoDist(ctx, "Sicily", "Palermo", "Catania", "ft").Result()
			Expect(err).NotTo(HaveOccurred())
		})

		It("should get geo hash in string representation", func() {
			hashes, err := adapter.GeoHash(ctx, "Sicily", "Palermo", "Catania").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(hashes).To(ConsistOf([]string{"sqc8b49rny0", "sqdtr74hyu0"}))
		})

		It("should return geo position", func() {
			pos, err := adapter.GeoPos(ctx, "Sicily", "Palermo", "Catania", "NonExisting").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(ConsistOf([]*GeoPos{
				{
					Longitude: 13.361389338970184,
					Latitude:  38.1155563954963,
				},
				{
					Longitude: 15.087267458438873,
					Latitude:  37.50266842333162,
				},
				nil,
			}))
		})

		if resp3 {
			It("should geo search", func() {
				q := GeoSearchQuery{
					Member:    "Catania",
					BoxWidth:  400,
					BoxHeight: 100,
					BoxUnit:   "km",
					Sort:      "asc",
				}
				val, err := adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania"}))

				q.BoxHeight = 400
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania", "Palermo"}))

				q.Count = 1
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania"}))

				q.CountAny = true
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Palermo"}))

				q = GeoSearchQuery{
					Member:     "Catania",
					Radius:     100,
					RadiusUnit: "km",
					Sort:       "asc",
				}
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania"}))

				q.Radius = 400
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania", "Palermo"}))

				q.Count = 1
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania"}))

				q.CountAny = true
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Palermo"}))

				q = GeoSearchQuery{
					Longitude: 15,
					Latitude:  37,
					BoxWidth:  200,
					BoxHeight: 200,
					BoxUnit:   "km",
					Sort:      "asc",
				}
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania"}))

				q.BoxWidth, q.BoxHeight = 400, 400
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania", "Palermo"}))

				q.Count = 1
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania"}))

				q.CountAny = true
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Palermo"}))

				q = GeoSearchQuery{
					Longitude:  15,
					Latitude:   37,
					Radius:     100,
					RadiusUnit: "km",
					Sort:       "asc",
				}
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania"}))

				q.Radius = 200
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania", "Palermo"}))

				q.Count = 1
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Catania"}))

				q.CountAny = true
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]string{"Palermo"}))
			})

			It("should geo search with options", func() {
				q := GeoSearchLocationQuery{
					GeoSearchQuery: GeoSearchQuery{
						Longitude:  15,
						Latitude:   37,
						Radius:     200,
						RadiusUnit: "km",
						Sort:       "asc",
					},
					WithHash:  true,
					WithDist:  true,
					WithCoord: true,
				}
				val, err := adapter.GeoSearchLocation(ctx, "Sicily", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal([]GeoLocation{
					{
						Name:      "Catania",
						Longitude: 15.08726745843887329,
						Latitude:  37.50266842333162032,
						Dist:      56.4413,
						GeoHash:   3479447370796909,
					},
					{
						Name:      "Palermo",
						Longitude: 13.36138933897018433,
						Latitude:  38.11555639549629859,
						Dist:      190.4424,
						GeoHash:   3479099956230698,
					},
				}))
			})

			It("should geo search store", func() {
				q := GeoSearchStoreQuery{
					GeoSearchQuery: GeoSearchQuery{
						Longitude:  15,
						Latitude:   37,
						Radius:     200,
						RadiusUnit: "km",
						Sort:       "asc",
					},
					StoreDist: false,
				}

				val, err := adapter.GeoSearchStore(ctx, "Sicily", "key1", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal(int64(2)))

				q.StoreDist = true
				val, err = adapter.GeoSearchStore(ctx, "Sicily", "key2", q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(Equal(int64(2)))

				loc, err := adapter.GeoSearchLocation(ctx, "key1", GeoSearchLocationQuery{
					GeoSearchQuery: q.GeoSearchQuery,
					WithCoord:      true,
					WithDist:       true,
					WithHash:       true,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(loc).To(Equal([]GeoLocation{
					{
						Name:      "Catania",
						Longitude: 15.08726745843887329,
						Latitude:  37.50266842333162032,
						Dist:      56.4413,
						GeoHash:   3479447370796909,
					},
					{
						Name:      "Palermo",
						Longitude: 13.36138933897018433,
						Latitude:  38.11555639549629859,
						Dist:      190.4424,
						GeoHash:   3479099956230698,
					},
				}))

				v, err := adapter.ZRangeWithScores(ctx, "key2", 0, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				for i, z := range v {
					v[i].Score = float64(int(z.Score))
				}
				Expect(v).To(Equal([]Z{
					{
						Score:  56.,
						Member: "Catania",
					},
					{
						Score:  190.,
						Member: "Palermo",
					},
				}))
			})
		}
	})

	Describe("marshaling/unmarshaling", func() {
		type convTest struct {
			value  any
			dest   any
			wanted string
		}

		convTests := []convTest{
			// TODO
			// {nil, "", nil},
			{"hello", new(string), "hello"},
			{[]byte("hello"), new([]byte), "hello"},
			{int(1), new(int), "1"},
			{int8(1), new(int8), "1"},
			{int16(1), new(int16), "1"},
			{int32(1), new(int32), "1"},
			{int64(1), new(int64), "1"},
			{uint(1), new(uint), "1"},
			{uint8(1), new(uint8), "1"},
			{uint16(1), new(uint16), "1"},
			{uint32(1), new(uint32), "1"},
			{uint64(1), new(uint64), "1"},
			{float32(1.0), new(float32), "1"},
			{float64(1.0), new(float64), "1"},
			{true, new(bool), "1"},
			{false, new(bool), "0"},
		}

		It("should convert to string", func() {
			for _, test := range convTests {
				err := adapter.Set(ctx, "key", test.value, 0).Err()
				Expect(err).NotTo(HaveOccurred())

				s, err := adapter.Get(ctx, "key").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(s).To(Equal(test.wanted))

				if test.dest == nil {
					continue
				}
				// TODO
				// err = adapter.Get(ctx, "key").Scan(test.dest)
				// Expect(err).NotTo(HaveOccurred())
				// Expect(deref(test.dest)).To(Equal(test.value))
			}
		})
	})

	Describe("json marshaling/unmarshaling", func() {
		BeforeEach(func() {
			value := &numberStruct{Number: 42}
			err := adapter.Set(ctx, "key", value, 0).Err()
			Expect(err).NotTo(HaveOccurred())
		})

		It("should marshal custom values using json", func() {
			s, err := adapter.Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(s).To(Equal(`{"Number":42}`))
		})

		// TODO
		// It("should scan custom values using json", func() {
		//	value := &numberStruct{}
		//	err := adapter.Get(ctx, "key").Scan(value)
		//	Expect(err).NotTo(HaveOccurred())
		//	Expect(value.Number).To(Equal(42))
		// })
	})

	Describe("Pub/Sub", func() {
		It("Publish", func() {
			v, err := adapter.Publish(ctx, "ch", "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal(int64(0)))
		})

		It("PubSubChannels", func() {
			v, err := adapter.PubSubChannels(ctx, "*").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal([]string{}))
		})

		It("PubSubNumPat", func() {
			v, err := adapter.PubSubNumPat(ctx).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal(int64(0)))
		})

		It("PubSubNumSub", func() {
			v, err := adapter.PubSubNumSub(ctx, "ch").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal(map[string]int64{"ch": 0}))
		})

		if resp3 {
			It("SPublish", func() {
				v, err := adapter.SPublish(ctx, "ch", "1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(v).To(Equal(int64(0)))
			})

			It("PubSubShardChannels", func() {
				v, err := adapter.PubSubShardChannels(ctx, "*").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(v).To(Equal([]string{}))
			})

			It("PubSubShardNumSub", func() {
				v, err := adapter.PubSubShardNumSub(ctx, "ch").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(v).To(Equal(map[string]int64{"ch": 0}))
			})
		}
	})

	Describe("Script", func() {
		It("returns keys and values", func() {
			vals, err := adapter.Eval(
				ctx,
				"return {KEYS[1],ARGV[1]}",
				[]string{"key"},
				"hello",
			).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]any{"key", "hello"}))
		})

		It("returns all values after an error", func() {
			vals, err := adapter.Eval(
				ctx,
				`return {12, {err="error"}, "abc"}`,
				nil,
			).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals.([]any)[0]).To(Equal(int64(12)))
			Expect(vals.([]any)[1].(error).Error()).To(Equal("error"))
			Expect(vals.([]any)[2]).To(Equal("abc"))
		})

		It("script load", func() {
			val, err := adapter.ScriptLoad(
				ctx,
				"return {KEYS[1],ARGV[1]}",
			).Result()
			Expect(err).NotTo(HaveOccurred())
			ret, err := adapter.ScriptExists(ctx, val).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(ret).To(Equal([]bool{true}))

			vals, err := adapter.EvalSha(
				ctx,
				val,
				[]string{"key"},
				"hello",
			).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]any{"key", "hello"}))
		})

		if resp3 {
			It("script load", func() {
				val, err := adapter.ScriptLoad(
					ctx,
					"return {KEYS[1],ARGV[1]}",
				).Result()
				Expect(err).NotTo(HaveOccurred())
				ret, err := adapter.ScriptExists(ctx, val).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(ret).To(Equal([]bool{true}))

				valsRo, err := adapter.EvalShaRO(
					ctx,
					val,
					[]string{"key"},
					"hello",
				).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(valsRo).To(Equal([]any{"key", "hello"}))
			})
		}

		It("script kill & flush", func() {
			Expect(adapter.ScriptKill(ctx).Err()).To(MatchError("NOTBUSY No scripts in execution right now."))
			Expect(adapter.ScriptFlush(ctx).Err()).NotTo(HaveOccurred())
		})
	})
}

func testAdapterCache(resp3 bool) {

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

	Describe("Config", func() {
		It("Flush", func() {
			Expect(adapter.FlushDBAsync(ctx).Err()).NotTo(HaveOccurred())
			time.Sleep(2 * time.Second)
			Expect(adapter.FlushAllAsync(ctx).Err()).NotTo(HaveOccurred())
			time.Sleep(2 * time.Second)
		})
		It("BgSave", func() {
			Expect(adapter.BgRewriteAOF(ctx).Err()).NotTo(HaveOccurred())
			time.Sleep(2 * time.Second)
			Expect(adapter.BgSave(ctx).Err()).NotTo(HaveOccurred())
			time.Sleep(2 * time.Second)
		})
		It("Config Rewrite", func() {
			Expect(adapter.ConfigRewrite(ctx).Err()).To(MatchError("The server is running without a config file"))
		})
		It("DebugObject", func() {
			Expect(adapter.DebugObject(ctx, "non").Err().Error()).To(HavePrefix("DEBUG command not allowed."))
		})
		It("ReadOnly & ReadWrite", func() {
			Expect(adapter.ReadOnly(ctx).Err()).To(MatchError("This instance has cluster support disabled"))
			Expect(adapter.ReadWrite(ctx).Err()).To(MatchError("This instance has cluster support disabled"))
		})
	})

	Describe("Client", func() {
		It("should ClientUnblock", func() {
			id := adapter.ClientID(ctx).Val()
			r, err := adapter.ClientUnblock(ctx, id).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(r).To(Equal(int64(0)))
		})

		It("should ClientUnblockWithError", func() {
			id := adapter.ClientID(ctx).Val()
			r, err := adapter.ClientUnblockWithError(ctx, id).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(r).To(Equal(int64(0)))
		})

		It("should ClientInfo", func() {
			info, err := adapter.ClientInfo(ctx).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(info).NotTo(BeNil())
		})

		It("ClientPause", func() {
			Expect(adapter.ClientPause(ctx, time.Second).Err()).NotTo(HaveOccurred())
		})

		It("ClientUnpause", func() {
			Expect(adapter.ClientUnpause(ctx).Err()).NotTo(HaveOccurred())
		})
	})

	Describe("EvalRO", func() {
		It("returns keys and values", func() {
			vals, err := adapter.EvalRO(
				ctx,
				"return {KEYS[1],ARGV[1]}",
				[]string{"key"},
				"hello",
			).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]any{"key", "hello"}))
		})

		It("returns all values after an error", func() {
			vals, err := adapter.EvalRO(
				ctx,
				`return {12, {err="error"}, "abc"}`,
				nil,
			).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals.([]any)[0]).To(Equal(int64(12)))
			Expect(vals.([]any)[1].(error).Error()).To(Equal("error"))
			Expect(vals.([]any)[2]).To(Equal("abc"))
		})
	})

	if resp3 {
		Describe("Functions", func() {
			var (
				q        FunctionListQuery
				lib1Code string
				lib2Code string
				lib1     Library
				lib2     Library
			)

			BeforeEach(func() {
				flush := adapter.FunctionFlush(ctx)
				Expect(flush.Err()).NotTo(HaveOccurred())

				lib1 = Library{
					Name:   "mylib1",
					Engine: "LUA",
					Functions: []Function{
						{
							Name:        "lib1_func1",
							Description: "This is the func-1 of lib 1",
							Flags:       []string{"allow-oom", "allow-stale"},
						},
					},
					Code: `#!lua name=%s

                     local function f1(keys, args)
                        local hash = keys[1]  -- Get the key name
                        local time = redis.call('TIME')[1]  -- Get the current time from the Redis server

                        -- Add the current timestamp to the arguments that the user passed to the function, stored in args
                        table.insert(args, '_updated_at')
                        table.insert(args, time)

                        -- Run HSET with the updated argument list
                        return redis.call('HSET', hash, unpack(args))
                     end

					redis.register_function{
						function_name='%s',
						description ='%s',
						callback=f1,
						flags={'%s', '%s'}
					}`,
				}

				lib2 = Library{
					Name:   "mylib2",
					Engine: "LUA",
					Functions: []Function{
						{
							Name:  "lib2_func1",
							Flags: []string{},
						},
						{
							Name:        "lib2_func2",
							Description: "This is the func-2 of lib 2",
							Flags:       []string{"no-writes"},
						},
					},
					Code: `#!lua name=%s

					local function f1(keys, args)
						 return 'Function 1'
					end

					local function f2(keys, args)
						 return 'Function 2'
					end

					redis.register_function('%s', f1)
					redis.register_function{
						function_name='%s',
						description ='%s',
						callback=f2,
						flags={'%s'}
					}`,
				}

				lib1Code = fmt.Sprintf(lib1.Code, lib1.Name, lib1.Functions[0].Name,
					lib1.Functions[0].Description, lib1.Functions[0].Flags[0], lib1.Functions[0].Flags[1])
				lib2Code = fmt.Sprintf(lib2.Code, lib2.Name, lib2.Functions[0].Name,
					lib2.Functions[1].Name, lib2.Functions[1].Description, lib2.Functions[1].Flags[0])

				q = FunctionListQuery{}
			})

			It("Loads a new library", func() {
				functionLoad := adapter.FunctionLoad(ctx, lib1Code)
				Expect(functionLoad.Err()).NotTo(HaveOccurred())
				Expect(functionLoad.Val()).To(Equal(lib1.Name))

				functionList := adapter.FunctionList(ctx, q)
				Expect(functionList.Err()).NotTo(HaveOccurred())
				Expect(functionList.Val()).To(HaveLen(1))
			})

			It("Loads and replaces a new library", func() {
				// Load a library for the first time
				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				Expect(err).NotTo(HaveOccurred())

				newFuncName := "replaces_func_name"
				newFuncDesc := "replaces_func_desc"
				flag1, flag2 := "allow-stale", "no-cluster"
				newCode := fmt.Sprintf(lib1.Code, lib1.Name, newFuncName, newFuncDesc, flag1, flag2)

				// And then replace it
				functionLoadReplace := adapter.FunctionLoadReplace(ctx, newCode)
				Expect(functionLoadReplace.Err()).NotTo(HaveOccurred())
				Expect(functionLoadReplace.Val()).To(Equal(lib1.Name))

				lib, err := adapter.FunctionList(ctx, q).First()
				Expect(err).NotTo(HaveOccurred())
				Expect(lib.Functions).To(Equal([]Function{
					{
						Name:        newFuncName,
						Description: newFuncDesc,
						Flags:       []string{flag1, flag2},
					},
				}))
			})

			It("Deletes a library", func() {
				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.FunctionDelete(ctx, lib1.Name).Err()
				Expect(err).NotTo(HaveOccurred())

				val, err := adapter.FunctionList(ctx, FunctionListQuery{
					LibraryNamePattern: lib1.Name,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(HaveLen(0))
			})

			It("Flushes all libraries", func() {
				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.FunctionLoad(ctx, lib2Code).Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.FunctionFlush(ctx).Err()
				Expect(err).NotTo(HaveOccurred())

				val, err := adapter.FunctionList(ctx, q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(HaveLen(0))
			})

			It("Flushes all libraries asynchronously", func() {
				functionLoad := adapter.FunctionLoad(ctx, lib1Code)
				Expect(functionLoad.Err()).NotTo(HaveOccurred())

				// we only verify the command result.
				functionFlush := adapter.FunctionFlushAsync(ctx)
				Expect(functionFlush.Err()).NotTo(HaveOccurred())
			})

			It("Kills a running function", func() {
				functionKill := adapter.FunctionKill(ctx)
				Expect(functionKill.Err()).To(MatchError("NOTBUSY No scripts in execution right now."))

				// Add test for a long-running function, once we make the test for `function stats` pass
			})

			It("Lists registered functions", func() {
				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				Expect(err).NotTo(HaveOccurred())

				val, err := adapter.FunctionList(ctx, FunctionListQuery{
					LibraryNamePattern: "*",
					WithCode:           true,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(HaveLen(1))
				Expect(val[0].Name).To(Equal(lib1.Name))
				Expect(val[0].Engine).To(Equal(lib1.Engine))
				Expect(val[0].Code).To(Equal(lib1Code))
				Expect(val[0].Functions).Should(ConsistOf(lib1.Functions))

				err = adapter.FunctionLoad(ctx, lib2Code).Err()
				Expect(err).NotTo(HaveOccurred())

				val, err = adapter.FunctionList(ctx, FunctionListQuery{
					WithCode: true,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(val).To(HaveLen(2))

				lib, err := adapter.FunctionList(ctx, FunctionListQuery{
					LibraryNamePattern: lib2.Name,
					WithCode:           false,
				}).First()
				Expect(err).NotTo(HaveOccurred())
				Expect(lib.Name).To(Equal(lib2.Name))
				Expect(lib.Code).To(Equal(""))

				_, err = adapter.FunctionList(ctx, FunctionListQuery{
					LibraryNamePattern: "non_lib",
					WithCode:           true,
				}).First()
				Expect(err).To(Equal(rueidis.Nil))
			})

			It("Dump and restores all libraries", func() {
				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.FunctionLoad(ctx, lib2Code).Err()
				Expect(err).NotTo(HaveOccurred())

				dump, err := adapter.FunctionDump(ctx).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(dump).NotTo(BeEmpty())

				err = adapter.FunctionRestore(ctx, dump).Err()
				Expect(err).To(HaveOccurred())

				err = adapter.FunctionFlush(ctx).Err()
				Expect(err).NotTo(HaveOccurred())

				list, err := adapter.FunctionList(ctx, q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(list).To(HaveLen(0))

				err = adapter.FunctionRestore(ctx, dump).Err()
				Expect(err).NotTo(HaveOccurred())

				list, err = adapter.FunctionList(ctx, q).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(list).To(HaveLen(2))
			})

			It("Calls a function", func() {
				lib1Code = fmt.Sprintf(lib1.Code, lib1.Name, lib1.Functions[0].Name,
					lib1.Functions[0].Description, lib1.Functions[0].Flags[0], lib1.Functions[0].Flags[1])

				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				Expect(err).NotTo(HaveOccurred())

				x := adapter.FCall(ctx, lib1.Functions[0].Name, []string{"my_hash"}, "a", 1, "b", 2)
				Expect(x.Err()).NotTo(HaveOccurred())
				Expect(x.Int()).To(Equal(3))
			})

			It("Calls a function as read-only", func() {
				lib1Code = fmt.Sprintf(lib1.Code, lib1.Name, lib1.Functions[0].Name,
					lib1.Functions[0].Description, lib1.Functions[0].Flags[0], lib1.Functions[0].Flags[1])

				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				Expect(err).NotTo(HaveOccurred())

				// This function doesn't have a "no-writes" flag
				x := adapter.FCallRO(ctx, lib1.Functions[0].Name, []string{"my_hash"}, "a", 1, "b", 2)

				Expect(x.Err()).To(HaveOccurred())

				lib2Code = fmt.Sprintf(lib2.Code, lib2.Name, lib2.Functions[0].Name, lib2.Functions[1].Name,
					lib2.Functions[1].Description, lib2.Functions[1].Flags[0])

				// This function has a "no-writes" flag
				err = adapter.FunctionLoad(ctx, lib2Code).Err()
				Expect(err).NotTo(HaveOccurred())

				x = adapter.FCallRO(ctx, lib2.Functions[1].Name, []string{})

				Expect(x.Err()).NotTo(HaveOccurred())
				Expect(x.Text()).To(Equal("Function 2"))
			})

			It("Shows function stats", func() {
				defer func() {
					for i := 0; i < 30; i++ {
						adapter.FunctionKill(ctx)
					}
				}()

				// We cannot run blocking commands in Redis functions, so we're using an infinite loop,
				// but we're killing the function after calling FUNCTION STATS
				lib := Library{
					Name:   "mylib1",
					Engine: "LUA",
					Functions: []Function{
						{
							Name:        "lib1_func1",
							Description: "This is the func-1 of lib 1",
							Flags:       []string{"no-writes"},
						},
					},
					Code: `#!lua name=%s
					local function f1(keys, args)
						local i = 0
					   	while true do
							i = i + 1
					   	end
					end

					redis.register_function{
						function_name='%s',
						description ='%s',
						callback=f1,
						flags={'%s'}
					}`,
				}
				libCode := fmt.Sprintf(lib.Code, lib.Name, lib.Functions[0].Name,
					lib.Functions[0].Description, lib.Functions[0].Flags[0])
				err := adapter.FunctionLoad(ctx, libCode).Err()

				Expect(err).NotTo(HaveOccurred())

				r, err := adapter.FunctionStats(ctx).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(r.Engines)).To(Equal(1))
				Expect(r.Running()).To(BeFalse())

				started := make(chan bool, 1)
				go func() {
					defer GinkgoRecover()

					for addr := range clientresp3.Nodes() {
						client2, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{addr}})
						Expect(err).NotTo(HaveOccurred())
						adapter2 := NewAdapter(client2)
						started <- true
						adapter2.FCall(ctx, lib.Functions[0].Name, nil).Result()
						break
					}
				}()

				<-started
				time.Sleep(1 * time.Second)
				r, err = adapter.FunctionStats(ctx).Result()

				Expect(err).NotTo(HaveOccurred())
				Expect(len(r.Engines)).To(Equal(1))
				rs, isRunning := r.RunningScript()
				Expect(isRunning).To(BeTrue())
				Expect(rs.Name).To(Equal(lib.Functions[0].Name))
				Expect(rs.Duration > 0).To(BeTrue())

				close(started)
			})
		})

		Describe("SlowLogGet", func() {
			It("returns slow query result", func() {
				const key = "slowlog-log-slower-than"

				old := adapter.ConfigGet(ctx, key).Val()
				adapter.ConfigSet(ctx, key, "0")
				defer adapter.ConfigSet(ctx, key, old[key])

				err := adapter.SlowLogReset(ctx).Err()
				Expect(err).NotTo(HaveOccurred())

				adapter.Set(ctx, "test", "true", 0)

				result, err := adapter.SlowLogGet(ctx, -1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(result)).NotTo(BeZero())
			})
		})
	}

	Describe("keys", func() {

		It("should Expire", func() {
			ttl := adapter.Cache(time.Hour).TTL(ctx, "nonexistent_key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(time.Duration(-2)))

			set := adapter.Set(ctx, "key", "Hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expire := adapter.Expire(ctx, "key", 10*time.Second)
			Expect(expire.Err()).NotTo(HaveOccurred())
			Expect(expire.Val()).To(Equal(true))

			ttl = adapter.Cache(time.Hour).TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(10 * time.Second))

			set = adapter.Set(ctx, "key", "Hello World", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			time.Sleep(time.Millisecond * 100)

			ttl = adapter.Cache(time.Hour).TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(time.Duration(-1)))
		})

		It("should PExpire", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expiration := 900 * time.Millisecond
			pexpire := adapter.PExpire(ctx, "key", expiration)
			Expect(pexpire.Err()).NotTo(HaveOccurred())
			Expect(pexpire.Val()).To(Equal(true))

			ttl := adapter.Cache(time.Hour).TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(time.Second))

			pttl := adapter.Cache(time.Hour).PTTL(ctx, "key")
			Expect(pttl.Err()).NotTo(HaveOccurred())
			Expect(pttl.Val()).To(BeNumerically("~", expiration, 100*time.Millisecond))
		})

		It("should PExpireAt", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expiration := 900 * time.Millisecond
			pexpireat := adapter.PExpireAt(ctx, "key", time.Now().Add(expiration))
			Expect(pexpireat.Err()).NotTo(HaveOccurred())
			Expect(pexpireat.Val()).To(Equal(true))

			ttl := adapter.Cache(time.Hour).TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(time.Second))

			pttl := adapter.Cache(time.Hour).PTTL(ctx, "key")
			Expect(pttl.Err()).NotTo(HaveOccurred())
			Expect(pttl.Val()).To(BeNumerically("~", expiration, 100*time.Millisecond))
		})

		It("should PTTL", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expiration := time.Second
			expire := adapter.Expire(ctx, "key", expiration)
			Expect(expire.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			pttl := adapter.Cache(time.Hour).PTTL(ctx, "key")
			Expect(pttl.Err()).NotTo(HaveOccurred())
			Expect(pttl.Val()).To(BeNumerically("~", expiration, 100*time.Millisecond))
		})

		It("should Sort", func() {
			Expect(func() {
				adapter.Cache(time.Hour).SortRO(ctx, "list", Sort{
					Order: "PANIC",
				})
			}).To(Panic())
		})

		It("should Sort", func() {
			size, err := adapter.LPush(ctx, "list", "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list", "3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(3)))

			els, err := adapter.Cache(time.Hour).SortRO(ctx, "list", Sort{
				Offset: 0,
				Count:  2,
				Order:  "ASC",
				Alpha:  true,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(els).To(Equal([]string{"1", "2"}))
		})

		It("should Sort By", func() {
			size, err := adapter.LPush(ctx, "list_by", "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list_by", "3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list_by", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(3)))

			els, err := adapter.Cache(time.Hour).SortRO(ctx, "list_by", Sort{
				Offset: 0,
				Count:  2,
				By:     "nosort",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(els).To(Equal([]string{"2", "3"}))
		})

		It("should Sort and Get", func() {
			size, err := adapter.LPush(ctx, "list", "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list", "3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(size).To(Equal(int64(3)))

			err = adapter.Set(ctx, "object_2", "value2", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			{
				els, err := adapter.Cache(time.Hour).SortRO(ctx, "list", Sort{
					Get: []string{"object_*"},
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(els).To(Equal([]string{"", "value2", ""}))
			}

		})

		It("should TTL", func() {
			ttl := adapter.Cache(time.Hour).TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val() < 0).To(Equal(true))

			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			expire := adapter.Expire(ctx, "key", 60*time.Second)
			Expect(expire.Err()).NotTo(HaveOccurred())
			Expect(expire.Val()).To(Equal(true))

			ttl = adapter.Cache(time.Hour).TTL(ctx, "key")
			Expect(ttl.Err()).NotTo(HaveOccurred())
			Expect(ttl.Val()).To(Equal(60 * time.Second))
		})

		It("should Type", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			type_ := adapter.Cache(time.Hour).Type(ctx, "key")
			Expect(type_.Err()).NotTo(HaveOccurred())
			Expect(type_.Val()).To(Equal("string"))
		})
	})

	Describe("strings", func() {

		It("should BitCount", func() {
			set := adapter.Set(ctx, "key", "foobar", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			bitCount := adapter.Cache(time.Hour).BitCount(ctx, "key", nil)
			Expect(bitCount.Err()).NotTo(HaveOccurred())
			Expect(bitCount.Val()).To(Equal(int64(26)))

			bitCount = adapter.Cache(time.Hour).BitCount(ctx, "key", &BitCount{
				Start: 0,
				End:   0,
			})
			Expect(bitCount.Err()).NotTo(HaveOccurred())
			Expect(bitCount.Val()).To(Equal(int64(4)))

			bitCount = adapter.Cache(time.Hour).BitCount(ctx, "key", &BitCount{
				Start: 1,
				End:   1,
			})
			Expect(bitCount.Err()).NotTo(HaveOccurred())
			Expect(bitCount.Val()).To(Equal(int64(6)))

			if resp3 {
				bitCount = adapter.Cache(time.Hour).BitCount(ctx, "key", &BitCount{
					Start: 1,
					End:   1,
					Unit:  "BIT",
				})
				Expect(bitCount.Err()).NotTo(HaveOccurred())
				Expect(bitCount.Val()).To(Equal(int64(1)))

				bitCount = adapter.Cache(time.Hour).BitCount(ctx, "key", &BitCount{
					Start: 1,
					End:   1,
					Unit:  "BYTE",
				})
				Expect(bitCount.Err()).NotTo(HaveOccurred())
				Expect(bitCount.Val()).To(Equal(int64(6)))
			}
		})

		It("BitPos should panic", func() {
			Expect(func() {
				adapter.BitPos(ctx, "mykey", 0, 0, 0, 0)
			}).To(Panic())
		})

		It("should panic on too many arguments in BitPos", func() {
			defer func() {
				if r := recover(); r == nil {
					Fail("The code did not panic")
				}
			}()

			// This should cause the function to panic due to too many arguments
			adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, 0, 1, 2)
		})

		It("should BitPos", func() {
			err := adapter.Set(ctx, "mykey", "\xff\xf0\x00", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			pos, err := adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(12)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(0)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(16)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 1, 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(-1)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(16)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 1, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(-1)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, 2, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(-1)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, 0, -3).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(-1)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, 0, 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(-1)))
		})

		It("should BitPosSpan", func() {
			err := adapter.Set(ctx, "mykey", "\x00\xff\x00", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			pos, err := adapter.Cache(time.Hour).BitPosSpan(ctx, "mykey", 0, 1, 3, "byte").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(16)))

			pos, err = adapter.Cache(time.Hour).BitPosSpan(ctx, "mykey", 0, 1, 3, "bit").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(Equal(int64(1)))
		})

		It("should BitFieldRO", func() {
			nn, err := adapter.BitField(ctx, "mykey", "SET", "u8", 8, 255).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(nn).To(Equal([]int64{0}))

			nn, err = adapter.Cache(time.Hour).BitFieldRO(ctx, "mykey", "u8", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(nn).To(Equal([]int64{0}))

			nn, err = adapter.Cache(time.Hour).BitFieldRO(ctx, "mykey", "u8", 0, "u4", 8, "u4", 12, "u4", 13).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(nn).To(Equal([]int64{0, 15, 15, 14}))
		})

		Describe("EvalRO", func() {
			It("returns keys and values", func() {
				vals, err := adapter.Cache(time.Hour).EvalRO(
					ctx,
					"return {KEYS[1],ARGV[1]}",
					[]string{"key"},
					"hello",
				).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals).To(Equal([]any{"key", "hello"}))
			})

			It("returns all values after an error", func() {
				vals, err := adapter.Cache(time.Hour).EvalRO(
					ctx,
					`return {12, {err="error"}, "abc"}`,
					[]string{"key"},
				).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(vals.([]any)[0]).To(Equal(int64(12)))
				Expect(vals.([]any)[1].(error).Error()).To(Equal("error"))
				Expect(vals.([]any)[2]).To(Equal("abc"))
			})
		})

		It("script load", func() {
			val, err := adapter.ScriptLoad(
				ctx,
				"return {KEYS[1],ARGV[1]}",
			).Result()
			Expect(err).NotTo(HaveOccurred())
			ret, err := adapter.ScriptExists(ctx, val).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(ret).To(Equal([]bool{true}))

			valsRo, err := adapter.Cache(time.Hour).EvalShaRO(
				ctx,
				val,
				[]string{"key"},
				"hello",
			).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(valsRo).To(Equal([]any{"key", "hello"}))
		})

		Describe("Functions", func() {
			var (
				lib1Code string
				lib2Code string
				lib1     Library
				lib2     Library
			)

			BeforeEach(func() {
				flush := adapter.FunctionFlush(ctx)
				Expect(flush.Err()).NotTo(HaveOccurred())

				lib1 = Library{
					Name:   "mylib1",
					Engine: "LUA",
					Functions: []Function{
						{
							Name:        "lib1_func1",
							Description: "This is the func-1 of lib 1",
							Flags:       []string{"allow-oom", "allow-stale"},
						},
					},
					Code: `#!lua name=%s

                     local function f1(keys, args)
                        local hash = keys[1]  -- Get the key name
                        local time = redis.call('TIME')[1]  -- Get the current time from the Redis server

                        -- Add the current timestamp to the arguments that the user passed to the function, stored in args
                        table.insert(args, '_updated_at')
                        table.insert(args, time)

                        -- Run HSET with the updated argument list
                        return redis.call('HSET', hash, unpack(args))
                     end

					redis.register_function{
						function_name='%s',
						description ='%s',
						callback=f1,
						flags={'%s', '%s'}
					}`,
				}

				lib2 = Library{
					Name:   "mylib2",
					Engine: "LUA",
					Functions: []Function{
						{
							Name:  "lib2_func1",
							Flags: []string{},
						},
						{
							Name:        "lib2_func2",
							Description: "This is the func-2 of lib 2",
							Flags:       []string{"no-writes"},
						},
					},
					Code: `#!lua name=%s

					local function f1(keys, args)
						 return 'Function 1'
					end

					local function f2(keys, args)
						 return 'Function 2'
					end

					redis.register_function('%s', f1)
					redis.register_function{
						function_name='%s',
						description ='%s',
						callback=f2,
						flags={'%s'}
					}`,
				}

				lib1Code = fmt.Sprintf(lib1.Code, lib1.Name, lib1.Functions[0].Name,
					lib1.Functions[0].Description, lib1.Functions[0].Flags[0], lib1.Functions[0].Flags[1])
				lib2Code = fmt.Sprintf(lib2.Code, lib2.Name, lib2.Functions[0].Name,
					lib2.Functions[1].Name, lib2.Functions[1].Description, lib2.Functions[1].Flags[0])
			})

			It("Calls a function as read-only", func() {
				lib1Code = fmt.Sprintf(lib1.Code, lib1.Name, lib1.Functions[0].Name,
					lib1.Functions[0].Description, lib1.Functions[0].Flags[0], lib1.Functions[0].Flags[1])

				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				Expect(err).NotTo(HaveOccurred())

				// This function doesn't have a "no-writes" flag
				x := adapter.Cache(time.Hour).FCallRO(ctx, lib1.Functions[0].Name, []string{"my_hash"}, "a", 1, "b", 2)

				Expect(x.Err()).To(HaveOccurred())

				lib2Code = fmt.Sprintf(lib2.Code, lib2.Name, lib2.Functions[0].Name, lib2.Functions[1].Name,
					lib2.Functions[1].Description, lib2.Functions[1].Flags[0])

				// This function has a "no-writes" flag
				err = adapter.FunctionLoad(ctx, lib2Code).Err()
				Expect(err).NotTo(HaveOccurred())

				x = adapter.Cache(time.Hour).FCallRO(ctx, lib2.Functions[1].Name, []string{"my_hash"})

				Expect(x.Err()).NotTo(HaveOccurred())
				Expect(x.Text()).To(Equal("Function 2"))
			})
		})

		It("should Get", func() {
			get := adapter.Cache(time.Hour).Get(ctx, "_")
			Expect(rueidis.IsRedisNil(get.Err())).To(BeTrue())
			Expect(get.Val()).To(Equal(""))

			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			get = adapter.Cache(time.Hour).Get(ctx, "key")
			Expect(get.Err()).NotTo(HaveOccurred())
			Expect(get.Val()).To(Equal("hello"))
		})

		It("should GetBit", func() {
			setBit := adapter.SetBit(ctx, "key", 7, 1)
			Expect(setBit.Err()).NotTo(HaveOccurred())
			Expect(setBit.Val()).To(Equal(int64(0)))

			getBit := adapter.Cache(time.Hour).GetBit(ctx, "key", 0)
			Expect(getBit.Err()).NotTo(HaveOccurred())
			Expect(getBit.Val()).To(Equal(int64(0)))

			getBit = adapter.Cache(time.Hour).GetBit(ctx, "key", 7)
			Expect(getBit.Err()).NotTo(HaveOccurred())
			Expect(getBit.Val()).To(Equal(int64(1)))

			getBit = adapter.Cache(time.Hour).GetBit(ctx, "key", 100)
			Expect(getBit.Err()).NotTo(HaveOccurred())
			Expect(getBit.Val()).To(Equal(int64(0)))
		})

		It("should GetRange", func() {
			set := adapter.Set(ctx, "key", "This is a string", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			getRange := adapter.Cache(time.Hour).GetRange(ctx, "key", 0, 3)
			Expect(getRange.Err()).NotTo(HaveOccurred())
			Expect(getRange.Val()).To(Equal("This"))

			getRange = adapter.Cache(time.Hour).GetRange(ctx, "key", -3, -1)
			Expect(getRange.Err()).NotTo(HaveOccurred())
			Expect(getRange.Val()).To(Equal("ing"))

			getRange = adapter.Cache(time.Hour).GetRange(ctx, "key", 0, -1)
			Expect(getRange.Err()).NotTo(HaveOccurred())
			Expect(getRange.Val()).To(Equal("This is a string"))

			getRange = adapter.Cache(time.Hour).GetRange(ctx, "key", 10, 100)
			Expect(getRange.Err()).NotTo(HaveOccurred())
			Expect(getRange.Val()).To(Equal("string"))
		})

		It("should StrLen", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			Expect(set.Err()).NotTo(HaveOccurred())
			Expect(set.Val()).To(Equal("OK"))

			strLen := adapter.Cache(time.Hour).StrLen(ctx, "key")
			Expect(strLen.Err()).NotTo(HaveOccurred())
			Expect(strLen.Val()).To(Equal(int64(5)))

			strLen = adapter.Cache(time.Hour).StrLen(ctx, "_")
			Expect(strLen.Err()).NotTo(HaveOccurred())
			Expect(strLen.Val()).To(Equal(int64(0)))
		})
	})

	Describe("hashes", func() {

		It("should HExists", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			Expect(hSet.Err()).NotTo(HaveOccurred())

			hExists := adapter.Cache(time.Hour).HExists(ctx, "hash", "key")
			Expect(hExists.Err()).NotTo(HaveOccurred())
			Expect(hExists.Val()).To(Equal(true))

			hExists = adapter.Cache(time.Hour).HExists(ctx, "hash", "key1")
			Expect(hExists.Err()).NotTo(HaveOccurred())
			Expect(hExists.Val()).To(Equal(false))
		})

		It("should HGet", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			Expect(hSet.Err()).NotTo(HaveOccurred())

			hGet := adapter.Cache(time.Hour).HGet(ctx, "hash", "key")
			Expect(hGet.Err()).NotTo(HaveOccurred())
			Expect(hGet.Val()).To(Equal("hello"))

			hGet = adapter.Cache(time.Hour).HGet(ctx, "hash", "key1")
			Expect(rueidis.IsRedisNil(hGet.Err())).To(BeTrue())
			Expect(hGet.Val()).To(Equal(""))
		})

		It("should HGetAll", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1").Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.HSet(ctx, "hash", "key2", "hello2").Err()
			Expect(err).NotTo(HaveOccurred())

			m, err := adapter.Cache(time.Hour).HGetAll(ctx, "hash").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(m).To(Equal(map[string]string{"key1": "hello1", "key2": "hello2"}))
		})

		It("should HKeys", func() {
			hkeys := adapter.HKeys(ctx, "hash")
			Expect(hkeys.Err()).NotTo(HaveOccurred())
			Expect(hkeys.Val()).To(Equal([]string{}))

			hset := adapter.HSet(ctx, "hash", "key1", "hello1")
			Expect(hset.Err()).NotTo(HaveOccurred())
			hset = adapter.HSet(ctx, "hash", "key2", "hello2")
			Expect(hset.Err()).NotTo(HaveOccurred())

			hkeys = adapter.Cache(time.Hour).HKeys(ctx, "hash")
			Expect(hkeys.Err()).NotTo(HaveOccurred())
			Expect(hkeys.Val()).To(Equal([]string{"key1", "key2"}))
		})

		It("should HLen", func() {
			hSet := adapter.HSet(ctx, "hash", "key1", "hello1")
			Expect(hSet.Err()).NotTo(HaveOccurred())
			hSet = adapter.HSet(ctx, "hash", "key2", "hello2")
			Expect(hSet.Err()).NotTo(HaveOccurred())

			hLen := adapter.Cache(time.Hour).HLen(ctx, "hash")
			Expect(hLen.Err()).NotTo(HaveOccurred())
			Expect(hLen.Val()).To(Equal(int64(2)))
		})

		It("should HMGet", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1", "key2", "hello2").Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.Cache(time.Hour).HMGet(ctx, "hash", "key1", "key2", "_").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]any{"hello1", "hello2", nil}))
		})

		It("should HVals", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1").Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.HSet(ctx, "hash", "key2", "hello2").Err()
			Expect(err).NotTo(HaveOccurred())

			v, err := adapter.Cache(time.Hour).HVals(ctx, "hash").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal([]string{"hello1", "hello2"}))

			// TODO
			// var slice []string
			// err = adapter.Cache(time.Hour).HVals(ctx, "hash").ScanSlice(&slice)
			// Expect(err).NotTo(HaveOccurred())
			// Expect(slice).To(Equal([]string{"hello1", "hello2"}))
		})

		It("should HStrLen", func() {
			hSet := adapter.HSet(ctx, "hash", "field", "hello")
			Expect(hSet.Err()).NotTo(HaveOccurred())

			hStrLen := adapter.Cache(time.Hour).HStrLen(ctx, "hash", "field")
			Expect(hStrLen.Err()).NotTo(HaveOccurred())
			Expect(hStrLen.Val()).To(Equal(int64(5)))
		})
	})

	Describe("lists", func() {

		It("should LIndex", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			Expect(lPush.Err()).NotTo(HaveOccurred())
			lPush = adapter.LPush(ctx, "list", "Hello")
			Expect(lPush.Err()).NotTo(HaveOccurred())

			lIndex := adapter.Cache(time.Hour).LIndex(ctx, "list", 0)
			Expect(lIndex.Err()).NotTo(HaveOccurred())
			Expect(lIndex.Val()).To(Equal("Hello"))

			lIndex = adapter.Cache(time.Hour).LIndex(ctx, "list", -1)
			Expect(lIndex.Err()).NotTo(HaveOccurred())
			Expect(lIndex.Val()).To(Equal("World"))

			lIndex = adapter.Cache(time.Hour).LIndex(ctx, "list", 3)
			Expect(rueidis.IsRedisNil(lIndex.Err())).To(BeTrue())
			Expect(lIndex.Val()).To(Equal(""))
		})

		It("should LLen", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			Expect(lPush.Err()).NotTo(HaveOccurred())
			lPush = adapter.LPush(ctx, "list", "Hello")
			Expect(lPush.Err()).NotTo(HaveOccurred())

			lLen := adapter.Cache(time.Hour).LLen(ctx, "list")
			Expect(lLen.Err()).NotTo(HaveOccurred())
			Expect(lLen.Val()).To(Equal(int64(2)))
		})

		It("should LPos", func() {
			rPush := adapter.RPush(ctx, "list", "a")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "b")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "c")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "b")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			lPos := adapter.Cache(time.Hour).LPos(ctx, "list", "b", LPosArgs{})
			Expect(lPos.Err()).NotTo(HaveOccurred())
			Expect(lPos.Val()).To(Equal(int64(1)))

			lPos = adapter.Cache(time.Hour).LPos(ctx, "list", "b", LPosArgs{Rank: 2})
			Expect(lPos.Err()).NotTo(HaveOccurred())
			Expect(lPos.Val()).To(Equal(int64(3)))

			lPos = adapter.Cache(time.Hour).LPos(ctx, "list", "b", LPosArgs{Rank: -2})
			Expect(lPos.Err()).NotTo(HaveOccurred())
			Expect(lPos.Val()).To(Equal(int64(1)))

			lPos = adapter.Cache(time.Hour).LPos(ctx, "list", "b", LPosArgs{Rank: 2, MaxLen: 1})
			Expect(rueidis.IsRedisNil(lPos.Err())).To(BeTrue())

			lPos = adapter.Cache(time.Hour).LPos(ctx, "list", "z", LPosArgs{})
			Expect(rueidis.IsRedisNil(lPos.Err())).To(BeTrue())
		})

		It("should LRange", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			Expect(rPush.Err()).NotTo(HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			Expect(rPush.Err()).NotTo(HaveOccurred())

			lRange := adapter.Cache(time.Hour).LRange(ctx, "list", 0, 0)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"one"}))

			lRange = adapter.Cache(time.Hour).LRange(ctx, "list", -3, 2)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"one", "two", "three"}))

			lRange = adapter.Cache(time.Hour).LRange(ctx, "list", -100, 100)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{"one", "two", "three"}))

			lRange = adapter.Cache(time.Hour).LRange(ctx, "list", 5, 10)
			Expect(lRange.Err()).NotTo(HaveOccurred())
			Expect(lRange.Val()).To(Equal([]string{}))
		})
	})

	Describe("sets", func() {

		It("should SCard", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			Expect(sAdd.Val()).To(Equal(int64(1)))

			sAdd = adapter.SAdd(ctx, "set", "World")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			Expect(sAdd.Val()).To(Equal(int64(1)))

			sCard := adapter.Cache(time.Hour).SCard(ctx, "set")
			Expect(sCard.Err()).NotTo(HaveOccurred())
			Expect(sCard.Val()).To(Equal(int64(2)))
		})

		It("should IsMember", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sIsMember := adapter.Cache(time.Hour).SIsMember(ctx, "set", "one")
			Expect(sIsMember.Err()).NotTo(HaveOccurred())
			Expect(sIsMember.Val()).To(Equal(true))

			sIsMember = adapter.Cache(time.Hour).SIsMember(ctx, "set", "two")
			Expect(sIsMember.Err()).NotTo(HaveOccurred())
			Expect(sIsMember.Val()).To(Equal(false))
		})

		It("should SMIsMember", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sMIsMember := adapter.Cache(time.Hour).SMIsMember(ctx, "set", "one", "two")
			Expect(sMIsMember.Err()).NotTo(HaveOccurred())
			Expect(sMIsMember.Val()).To(Equal([]bool{true, false}))
		})

		It("should SMembers", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			Expect(sAdd.Err()).NotTo(HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "World")
			Expect(sAdd.Err()).NotTo(HaveOccurred())

			sMembers := adapter.Cache(time.Hour).SMembers(ctx, "set")
			Expect(sMembers.Err()).NotTo(HaveOccurred())
			Expect(sMembers.Val()).To(ConsistOf([]string{"Hello", "World"}))
		})
	})

	Describe("sorted sets", func() {

		It("should ZCard", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			card, err := adapter.Cache(time.Hour).ZCard(ctx, "zset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(card).To(Equal(int64(2)))
		})

		It("should ZCount", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			count, err := adapter.Cache(time.Hour).ZCount(ctx, "zset", "-inf", "+inf").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(int64(3)))

			count, err = adapter.Cache(time.Hour).ZCount(ctx, "zset", "(1", "3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(int64(2)))

			count, err = adapter.Cache(time.Hour).ZLexCount(ctx, "zset", "-", "+").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(int64(3)))
		})

		It("should ZRangeWithScores", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))

			vals, err = adapter.Cache(time.Hour).ZRangeWithScores(ctx, "zset", 2, 3).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{Score: 3, Member: "three"}}))

			vals, err = adapter.Cache(time.Hour).ZRangeWithScores(ctx, "zset", -2, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))
		})

		It("should ZRangeArgs", func() {
			added, err := adapter.ZAddArgs(ctx, "zset", ZAddArgs{
				Members: []Z{
					{Score: 1, Member: "one"},
					{Score: 2, Member: "two"},
					{Score: 3, Member: "three"},
				},
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(3)))

			added, err = adapter.ZAddArgs(ctx, "zset", ZAddArgs{
				NX: true,
				Members: []Z{
					{Score: 4, Member: "four"},
				},
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(1)))

			added, err = adapter.ZAddArgs(ctx, "zsetxx", ZAddArgs{
				XX: true,
				Members: []Z{
					{Score: 1, Member: "one"},
				},
				Ch: true,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(added).To(Equal(int64(0)))

			score, err := adapter.ZAddArgsIncr(ctx, "zsetxx", ZAddArgs{
				Members: []Z{
					{Score: 1, Member: "one"},
				},
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(score).To(Equal(float64(1)))

			zRange, err := adapter.Cache(time.Hour).ZRangeArgs(ctx, ZRangeArgs{
				Key:     "zset",
				Start:   1,
				Stop:    4,
				ByScore: true,
				Rev:     true,
				Offset:  1,
				Count:   2,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(zRange).To(Equal([]string{"three", "two"}))

			zRange, err = adapter.Cache(time.Hour).ZRangeArgs(ctx, ZRangeArgs{
				Key:    "zset",
				Start:  "-",
				Stop:   "+",
				ByLex:  true,
				Rev:    true,
				Offset: 2,
				Count:  2,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(zRange).To(Equal([]string{"two", "one"}))

			zRange, err = adapter.Cache(time.Hour).ZRangeArgs(ctx, ZRangeArgs{
				Key:     "zset",
				Start:   "(1",
				Stop:    "(4",
				ByScore: true,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(zRange).To(Equal([]string{"two", "three"}))

			// withScores.
			zSlice, err := adapter.Cache(time.Hour).ZRangeArgsWithScores(ctx, ZRangeArgs{
				Key:     "zset",
				Start:   1,
				Stop:    4,
				ByScore: true,
				Rev:     true,
				Offset:  1,
				Count:   2,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(zSlice).To(Equal([]Z{
				{Score: 3, Member: "three"},
				{Score: 2, Member: "two"},
			}))
		})

		It("should ZRangeByScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRangeByScore := adapter.Cache(time.Hour).ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "-inf",
				Max: "+inf",
			})
			Expect(zRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRangeByScore.Val()).To(Equal([]string{"one", "two", "three"}))

			zRangeByScore = adapter.Cache(time.Hour).ZRangeByScore(ctx, "zset", ZRangeBy{
				Min:    "-inf",
				Max:    "+inf",
				Offset: 1,
				Count:  2,
			})
			Expect(zRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRangeByScore.Val()).To(Equal([]string{"two", "three"}))

			zRangeByScore = adapter.Cache(time.Hour).ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "1",
				Max: "2",
			})
			Expect(zRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRangeByScore.Val()).To(Equal([]string{"one", "two"}))

			zRangeByScore = adapter.Cache(time.Hour).ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "2",
			})
			Expect(zRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRangeByScore.Val()).To(Equal([]string{"two"}))

			zRangeByScore = adapter.Cache(time.Hour).ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "(2",
			})
			Expect(zRangeByScore.Err()).NotTo(HaveOccurred())
			Expect(zRangeByScore.Val()).To(Equal([]string{}))
		})

		It("should ZRangeByLex", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "a",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "b",
			}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "c",
			}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRangeByLex := adapter.Cache(time.Hour).ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "-",
				Max: "+",
			})
			Expect(zRangeByLex.Err()).NotTo(HaveOccurred())
			Expect(zRangeByLex.Val()).To(Equal([]string{"a", "b", "c"}))

			zRangeByLex = adapter.Cache(time.Hour).ZRangeByLex(ctx, "zset", ZRangeBy{
				Min:    "-",
				Max:    "+",
				Offset: 1,
				Count:  2,
			})
			Expect(zRangeByLex.Err()).NotTo(HaveOccurred())
			Expect(zRangeByLex.Val()).To(Equal([]string{"b", "c"}))

			zRangeByLex = adapter.Cache(time.Hour).ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "[a",
				Max: "[b",
			})
			Expect(zRangeByLex.Err()).NotTo(HaveOccurred())
			Expect(zRangeByLex.Val()).To(Equal([]string{"a", "b"}))

			zRangeByLex = adapter.Cache(time.Hour).ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "(a",
				Max: "[b",
			})
			Expect(zRangeByLex.Err()).NotTo(HaveOccurred())
			Expect(zRangeByLex.Val()).To(Equal([]string{"b"}))

			zRangeByLex = adapter.Cache(time.Hour).ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "(a",
				Max: "(b",
			})
			Expect(zRangeByLex.Err()).NotTo(HaveOccurred())
			Expect(zRangeByLex.Val()).To(Equal([]string{}))
		})

		It("should ZRangeByScoreWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "-inf",
				Max: "+inf",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))

			vals, err = adapter.Cache(time.Hour).ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min:    "-inf",
				Max:    "+inf",
				Offset: 1,
				Count:  2,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))

			vals, err = adapter.Cache(time.Hour).ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "1",
				Max: "2",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  2,
				Member: "two",
			}}))

			vals, err = adapter.Cache(time.Hour).ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "2",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{Score: 2, Member: "two"}}))

			vals, err = adapter.Cache(time.Hour).ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "(2",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{}))
		})

		It("should ZRank", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRank := adapter.Cache(time.Hour).ZRank(ctx, "zset", "three")
			Expect(zRank.Err()).NotTo(HaveOccurred())
			Expect(zRank.Val()).To(Equal(int64(2)))

			zRank = adapter.Cache(time.Hour).ZRank(ctx, "zset", "four")
			Expect(rueidis.IsRedisNil(zRank.Err())).To(BeTrue())
			Expect(zRank.Val()).To(Equal(int64(0)))
		})

		It("should ZRankWithScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRankWithScore := adapter.Cache(time.Hour).ZRankWithScore(ctx, "zset", "one")
			Expect(zRankWithScore.Err()).NotTo(HaveOccurred())
			Expect(zRankWithScore.Result()).To(Equal(RankScore{Rank: 0, Score: 1}))

			zRankWithScore = adapter.Cache(time.Hour).ZRankWithScore(ctx, "zset", "two")
			Expect(zRankWithScore.Err()).NotTo(HaveOccurred())
			Expect(zRankWithScore.Result()).To(Equal(RankScore{Rank: 1, Score: 2}))

			zRankWithScore = adapter.Cache(time.Hour).ZRankWithScore(ctx, "zset", "three")
			Expect(zRankWithScore.Err()).NotTo(HaveOccurred())
			Expect(zRankWithScore.Result()).To(Equal(RankScore{Rank: 2, Score: 3}))

			zRankWithScore = adapter.Cache(time.Hour).ZRankWithScore(ctx, "zset", "four")
			Expect(zRankWithScore.Err()).To(HaveOccurred())
			Expect(zRankWithScore.Err()).To(Equal(rueidis.Nil))
		})

		It("should ZRevRange", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRevRange := adapter.Cache(time.Hour).ZRevRange(ctx, "zset", 0, -1)
			Expect(zRevRange.Err()).NotTo(HaveOccurred())
			Expect(zRevRange.Val()).To(Equal([]string{"three", "two", "one"}))

			zRevRange = adapter.Cache(time.Hour).ZRevRange(ctx, "zset", 2, 3)
			Expect(zRevRange.Err()).NotTo(HaveOccurred())
			Expect(zRevRange.Val()).To(Equal([]string{"one"}))

			zRevRange = adapter.Cache(time.Hour).ZRevRange(ctx, "zset", -2, -1)
			Expect(zRevRange.Err()).NotTo(HaveOccurred())
			Expect(zRevRange.Val()).To(Equal([]string{"two", "one"}))
		})

		It("should ZRevRangeWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			val, err := adapter.Cache(time.Hour).ZRevRangeWithScores(ctx, "zset", 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))

			val, err = adapter.Cache(time.Hour).ZRevRangeWithScores(ctx, "zset", 2, 3).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]Z{{Score: 1, Member: "one"}}))

			val, err = adapter.Cache(time.Hour).ZRevRangeWithScores(ctx, "zset", -2, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))
		})

		It("should ZRevRangeByScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"three", "two", "one"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf", Offset: 1, Count: 2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"two", "one"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "2", Min: "(1"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"two"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "(2", Min: "(1"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{}))
		})

		It("should ZRevRangeByLex", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "a"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "b"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "c"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "+", Min: "-"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"c", "b", "a"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "+", Min: "-", Offset: 1, Count: 2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"b", "a"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "[b", Min: "(a"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{"b"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "(b", Min: "(a"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]string{}))
		})

		It("should ZRevRangeByScoreWithScores", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf", Offset: 1, Count: 2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))
		})

		It("should ZRevRangeByScoreWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "2", Min: "(1"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{{Score: 2, Member: "two"}}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "(2", Min: "(1"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(vals).To(Equal([]Z{}))
		})

		It("should ZRevRank", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRevRank := adapter.Cache(time.Hour).ZRevRank(ctx, "zset", "one")
			Expect(zRevRank.Err()).NotTo(HaveOccurred())
			Expect(zRevRank.Val()).To(Equal(int64(2)))

			zRevRank = adapter.Cache(time.Hour).ZRevRank(ctx, "zset", "four")
			Expect(rueidis.IsRedisNil(zRevRank.Err())).To(BeTrue())
			Expect(zRevRank.Val()).To(Equal(int64(0)))
		})

		It("should ZRevRankWithScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zRevRankWithScore := adapter.Cache(time.Hour).ZRevRankWithScore(ctx, "zset", "one")
			Expect(zRevRankWithScore.Err()).NotTo(HaveOccurred())
			Expect(zRevRankWithScore.Result()).To(Equal(RankScore{Rank: 2, Score: 1}))

			zRevRankWithScore = adapter.Cache(time.Hour).ZRevRankWithScore(ctx, "zset", "two")
			Expect(zRevRankWithScore.Err()).NotTo(HaveOccurred())
			Expect(zRevRankWithScore.Result()).To(Equal(RankScore{Rank: 1, Score: 2}))

			zRevRankWithScore = adapter.Cache(time.Hour).ZRevRankWithScore(ctx, "zset", "three")
			Expect(zRevRankWithScore.Err()).NotTo(HaveOccurred())
			Expect(zRevRankWithScore.Result()).To(Equal(RankScore{Rank: 0, Score: 3}))

			zRevRankWithScore = adapter.Cache(time.Hour).ZRevRankWithScore(ctx, "zset", "four")
			Expect(zRevRankWithScore.Err()).To(HaveOccurred())
			Expect(zRevRankWithScore.Err()).To(Equal(rueidis.Nil))
		})

		It("should ZScore", func() {
			zAdd := adapter.ZAdd(ctx, "zset", Z{Score: 1.001, Member: "one"})
			Expect(zAdd.Err()).NotTo(HaveOccurred())

			zScore := adapter.Cache(time.Hour).ZScore(ctx, "zset", "one")
			Expect(zScore.Err()).NotTo(HaveOccurred())
			Expect(zScore.Val()).To(Equal(float64(1.001)))
		})

		It("should ZMPop", func() {

			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			key, elems, err := adapter.ZMPop(ctx, "min", 1, "zset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(key).To(Equal("zset"))
			Expect(elems).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}}))

			_, _, err = adapter.ZMPop(ctx, "min", 1, "nosuchkey").Result()
			Expect(err).To(Equal(rueidis.Nil))

			err = adapter.ZAdd(ctx, "myzset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "myzset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "myzset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			key, elems, err = adapter.ZMPop(ctx, "min", 1, "myzset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(key).To(Equal("myzset"))
			Expect(elems).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}}))

			key, elems, err = adapter.ZMPop(ctx, "max", 10, "myzset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(key).To(Equal("myzset"))
			Expect(elems).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}}))

			err = adapter.ZAdd(ctx, "myzset2", Z{Score: 4, Member: "four"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "myzset2", Z{Score: 5, Member: "five"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "myzset2", Z{Score: 6, Member: "six"}).Err()
			Expect(err).NotTo(HaveOccurred())

			key, elems, err = adapter.ZMPop(ctx, "min", 10, "myzset", "myzset2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(key).To(Equal("myzset2"))
			Expect(elems).To(Equal([]Z{{
				Score:  4,
				Member: "four",
			}, {
				Score:  5,
				Member: "five",
			}, {
				Score:  6,
				Member: "six",
			}}))
		})

		It("should BZMPop", func() {

			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			key, elems, err := adapter.BZMPop(ctx, 0, "min", 1, "zset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(key).To(Equal("zset"))
			Expect(elems).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}}))
			key, elems, err = adapter.BZMPop(ctx, 0, "max", 1, "zset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(key).To(Equal("zset"))
			Expect(elems).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}}))
			key, elems, err = adapter.BZMPop(ctx, 0, "min", 10, "zset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(key).To(Equal("zset"))
			Expect(elems).To(Equal([]Z{{
				Score:  2,
				Member: "two",
			}}))

			key, elems, err = adapter.BZMPop(ctx, 0, "max", 10, "zset2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(key).To(Equal("zset2"))
			Expect(elems).To(Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))

			err = adapter.ZAdd(ctx, "myzset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			key, elems, err = adapter.BZMPop(ctx, 0, "min", 10, "myzset").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(key).To(Equal("myzset"))
			Expect(elems).To(Equal([]Z{{
				Score:  1,
				Member: "one",
			}}))

			err = adapter.ZAdd(ctx, "myzset2", Z{Score: 4, Member: "four"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "myzset2", Z{Score: 5, Member: "five"}).Err()
			Expect(err).NotTo(HaveOccurred())

			key, elems, err = adapter.BZMPop(ctx, 0, "min", 10, "myzset", "myzset2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(key).To(Equal("myzset2"))
			Expect(elems).To(Equal([]Z{{
				Score:  4,
				Member: "four",
			}, {
				Score:  5,
				Member: "five",
			}}))
		})

		It("should BFExists", func() {
			bfExists := adapter.Cache(time.Hour).BFExists(ctx, "key", "element")
			Expect(bfExists.Err()).NotTo(HaveOccurred())
		})

		It("should BFInfo", func() {
			bfAdd, err := adapter.BFAdd(ctx, "key", "element").Result()

			Expect(err).NotTo(HaveOccurred())
			Expect(bfAdd).To(BeTrue())

			// Check if the key exists
			bfExists := adapter.Cache(time.Hour).BFExists(ctx, "key", "element")
			Expect(bfExists.Val()).To(BeTrue())

			// Call BFInfo
			bfInfo := adapter.Cache(time.Hour).BFInfo(ctx, "key")
			Expect(bfInfo.Err()).NotTo(HaveOccurred())
		})

		It("should BFInfoArg with CAPACITY", func() {
			// Add the key
			bfAdd, err := adapter.BFAdd(ctx, "key", "element").Result()

			Expect(err).NotTo(HaveOccurred())
			Expect(bfAdd).To(BeTrue())

			// Call BFInfoArg with CAPACITY
			bfInfo := adapter.Cache(time.Hour).BFInfoArg(ctx, "key", "CAPACITY")
			Expect(bfInfo.Err()).NotTo(HaveOccurred())
		})

		It("should BFInfoArg with SIZE", func() {
			// Add the key
			bfAdd, err := adapter.BFAdd(ctx, "key", "element").Result()

			Expect(err).NotTo(HaveOccurred())
			Expect(bfAdd).To(BeTrue())

			// Call BFInfoArg with SIZE
			bfInfo := adapter.Cache(time.Hour).BFInfoArg(ctx, "key", "SIZE")
			Expect(bfInfo.Err()).NotTo(HaveOccurred())
		})

		It("should BFInfoArg with FILTERS", func() {
			// Add the key
			bfAdd, err := adapter.BFAdd(ctx, "key", "element").Result()

			Expect(err).NotTo(HaveOccurred())
			Expect(bfAdd).To(BeTrue())

			// Call BFInfoArg with FILTERS
			bfInfo := adapter.Cache(time.Hour).BFInfoArg(ctx, "key", "FILTERS")
			Expect(bfInfo.Err()).NotTo(HaveOccurred())
		})

		It("should BFInfoArg with ITEMS", func() {
			// Add the key
			bfAdd, err := adapter.BFAdd(ctx, "key", "element").Result()

			Expect(err).NotTo(HaveOccurred())
			Expect(bfAdd).To(BeTrue())

			// Call BFInfoArg with ITEMS
			bfInfo := adapter.Cache(time.Hour).BFInfoArg(ctx, "key", "ITEMS")
			Expect(bfInfo.Err()).NotTo(HaveOccurred())
		})

		It("should BFInfoArg with EXPANSION", func() {
			// Add the key
			bfAdd, err := adapter.BFAdd(ctx, "key", "element").Result()

			Expect(err).NotTo(HaveOccurred())
			Expect(bfAdd).To(BeTrue())

			// Call BFInfoArg with EXPANSION
			bfInfo := adapter.Cache(time.Hour).BFInfoArg(ctx, "key", "EXPANSION")
			Expect(bfInfo.Err()).NotTo(HaveOccurred())
		})

		It("should CFCount", func() {
			cfAdd := adapter.CFAdd(ctx, "cf_key", "element")
			Expect(cfAdd.Err()).NotTo(HaveOccurred())

			// Call CFCount
			cache := adapter.Cache(time.Hour)
			cfCount := cache.CFCount(ctx, "cf_key", "element")
			Expect(cfCount.Err()).NotTo(HaveOccurred())
		})

		It("should CFExists", func() {
			cache := adapter.Cache(time.Hour)
			// Add the key
			cfAdd := adapter.CFAdd(ctx, "cf_key", "element")
			Expect(cfAdd.Err()).NotTo(HaveOccurred())

			// Call CFExists
			cfExists := cache.CFExists(ctx, "cf_key", "element")
			Expect(cfExists.Err()).NotTo(HaveOccurred())
		})

		It("should CFInfo", func() {
			cache := adapter.Cache(time.Hour)

			// Add the key
			cfAdd := adapter.CFAdd(ctx, "cf_key", "element")
			Expect(cfAdd.Err()).NotTo(HaveOccurred())

			// Call CFInfo
			cfInfo := cache.CFInfo(ctx, "cf_key")
			Expect(cfInfo.Err()).NotTo(HaveOccurred())
		})

		It("should BZMPopBlocks", func() {
			started := make(chan bool)
			done := make(chan bool)
			go func() {
				defer GinkgoRecover()

				started <- true
				key, elems, err := adapter.BZMPop(ctx, 0, "min", 1, "list_list").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(key).To(Equal("list_list"))
				Expect(elems).To(Equal([]Z{{
					Score:  1,
					Member: "one",
				}}))
				done <- true
			}()
			<-started

			select {
			case <-done:
				Fail("BZMPop is not blocked")
			case <-time.After(time.Second):
				// ok
			}

			err := adapter.ZAdd(ctx, "list_list", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())

			select {
			case <-done:
				// ok
			case <-time.After(time.Second):
				Fail("BZMPop is still blocked")
			}
		})

		It("should BZMPop timeout", func() {
			_, val, err := adapter.BZMPop(ctx, time.Second, "min", 1, "list1").Result()
			Expect(err).To(Equal(rueidis.Nil))
			Expect(val).To(BeNil())

			Expect(adapter.Ping(ctx).Err()).NotTo(HaveOccurred())
		})

		It("should ZMScore", func() {
			zmScore := adapter.Cache(time.Hour).ZMScore(ctx, "zset", "one", "three")
			Expect(zmScore.Err()).NotTo(HaveOccurred())
			Expect(zmScore.Val()).To(HaveLen(2))
			Expect(zmScore.Val()[0]).To(Equal(float64(0)))

			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			Expect(err).NotTo(HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			Expect(err).NotTo(HaveOccurred())

			zmScore = adapter.Cache(time.Hour).ZMScore(ctx, "zset", "one", "three")
			Expect(zmScore.Err()).NotTo(HaveOccurred())
			Expect(zmScore.Val()).To(HaveLen(2))
			Expect(zmScore.Val()[0]).To(Equal(float64(1)))

			zmScore = adapter.Cache(time.Hour).ZMScore(ctx, "zset", "four")
			Expect(zmScore.Err()).NotTo(HaveOccurred())
			Expect(zmScore.Val()).To(HaveLen(1))

			zmScore = adapter.Cache(time.Hour).ZMScore(ctx, "zset", "four", "one")
			Expect(zmScore.Err()).NotTo(HaveOccurred())
			Expect(zmScore.Val()).To(HaveLen(2))
		})
	})

	Describe("Geo add and radius search", func() {
		BeforeEach(func() {
			n, err := adapter.GeoAdd(
				ctx,
				"Sicily",
				GeoLocation{Longitude: 13.361389, Latitude: 38.115556, Name: "Palermo"},
				GeoLocation{Longitude: 15.087269, Latitude: 37.502669, Name: "Catania"},
			).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(int64(2)))
		})

		It("should search geo radius", func() {
			res, err := adapter.Cache(time.Hour).GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius: 200,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveLen(2))
			Expect(res[0].Name).To(Equal("Palermo"))
			Expect(res[1].Name).To(Equal("Catania"))
		})

		It("should search geo radius with options", func() {
			res, err := adapter.Cache(time.Hour).GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveLen(2))
			Expect(res[1].Name).To(Equal("Palermo"))
			Expect(res[1].Dist).To(Equal(190.4424))
			Expect(res[1].GeoHash).To(Equal(int64(3479099956230698)))
			Expect(res[1].Longitude).To(Equal(13.361389338970184))
			Expect(res[1].Latitude).To(Equal(38.115556395496299))
			Expect(res[0].Name).To(Equal("Catania"))
			Expect(res[0].Dist).To(Equal(56.4413))
			Expect(res[0].GeoHash).To(Equal(int64(3479447370796909)))
			Expect(res[0].Longitude).To(Equal(15.087267458438873))
			Expect(res[0].Latitude).To(Equal(37.50266842333162))
		})

		It("should search geo radius with WithDist=false", func() {
			res, err := adapter.Cache(time.Hour).GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveLen(2))
			Expect(res[1].Name).To(Equal("Palermo"))
			Expect(res[1].Dist).To(Equal(float64(0)))
			Expect(res[1].GeoHash).To(Equal(int64(3479099956230698)))
			Expect(res[1].Longitude).To(Equal(13.361389338970184))
			Expect(res[1].Latitude).To(Equal(38.115556395496299))
			Expect(res[0].Name).To(Equal("Catania"))
			Expect(res[0].Dist).To(Equal(float64(0)))
			Expect(res[0].GeoHash).To(Equal(int64(3479447370796909)))
			Expect(res[0].Longitude).To(Equal(15.087267458438873))
			Expect(res[0].Latitude).To(Equal(37.50266842333162))
		})

		It("should search geo radius by member with options", func() {
			res, err := adapter.Cache(time.Hour).GeoRadiusByMember(ctx, "Sicily", "Catania", GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveLen(2))
			Expect(res[0].Name).To(Equal("Catania"))
			Expect(res[0].Dist).To(Equal(0.0))
			Expect(res[0].GeoHash).To(Equal(int64(3479447370796909)))
			Expect(res[0].Longitude).To(Equal(15.087267458438873))
			Expect(res[0].Latitude).To(Equal(37.50266842333162))
			Expect(res[1].Name).To(Equal("Palermo"))
			Expect(res[1].Dist).To(Equal(166.2742))
			Expect(res[1].GeoHash).To(Equal(int64(3479099956230698)))
			Expect(res[1].Longitude).To(Equal(13.361389338970184))
			Expect(res[1].Latitude).To(Equal(38.115556395496299))
		})

		It("should search geo radius with no results", func() {
			res, err := adapter.Cache(time.Hour).GeoRadius(ctx, "Sicily", 99, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveLen(0))
		})

		It("should panic on invalid unit in GeoDist", func() {
			defer func() {
				if r := recover(); r == nil {
					Fail("The code did not panic")
				}
			}()

			// This should cause the function to panic due to an invalid unit
			adapter.Cache(time.Hour).GeoDist(ctx, "Sicily", "Palermo", "Catania", "invalid_unit")
		})

		It("should get geo distance with unit options", func() {
			// From Redis CLI, note the difference in rounding in m vs
			// km on Redis itself.
			//
			// GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
			// GEODIST Sicily Palermo Catania m
			// "166274.15156960033"
			// GEODIST Sicily Palermo Catania km
			// "166.27415156960032"
			dist, err := adapter.Cache(time.Hour).GeoDist(ctx, "Sicily", "Palermo", "Catania", "km").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(dist).To(BeNumerically("~", 166.27, 0.01))

			dist, err = adapter.Cache(time.Hour).GeoDist(ctx, "Sicily", "Palermo", "Catania", "m").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(dist).To(BeNumerically("~", 166274.15, 0.01))

			_, err = adapter.Cache(time.Hour).GeoDist(ctx, "Sicily", "Palermo", "Catania", "mi").Result()
			Expect(err).NotTo(HaveOccurred())

			_, err = adapter.Cache(time.Hour).GeoDist(ctx, "Sicily", "Palermo", "Catania", "ft").Result()
			Expect(err).NotTo(HaveOccurred())
		})

		It("should get geo hash in string representation", func() {
			hashes, err := adapter.Cache(time.Hour).GeoHash(ctx, "Sicily", "Palermo", "Catania").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(hashes).To(ConsistOf([]string{"sqc8b49rny0", "sqdtr74hyu0"}))
		})

		It("should return geo position", func() {
			pos, err := adapter.Cache(time.Hour).GeoPos(ctx, "Sicily", "Palermo", "Catania", "NonExisting").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(pos).To(ConsistOf([]*GeoPos{
				{
					Longitude: 13.361389338970184,
					Latitude:  38.1155563954963,
				},
				{
					Longitude: 15.087267458438873,
					Latitude:  37.50266842333162,
				},
				nil,
			}))
		})

		It("should geo search", func() {
			q := GeoSearchQuery{
				Member:    "Catania",
				BoxWidth:  400,
				BoxHeight: 100,
				BoxUnit:   "km",
				Sort:      "asc",
			}
			val, err := adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania"}))

			q.BoxHeight = 400
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania", "Palermo"}))

			q.Count = 1
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania"}))

			q.CountAny = true
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Palermo"}))

			q = GeoSearchQuery{
				Member:     "Catania",
				Radius:     100,
				RadiusUnit: "km",
				Sort:       "asc",
			}
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania"}))

			q.Radius = 400
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania", "Palermo"}))

			q.Count = 1
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania"}))

			q.CountAny = true
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Palermo"}))

			q = GeoSearchQuery{
				Longitude: 15,
				Latitude:  37,
				BoxWidth:  200,
				BoxHeight: 200,
				BoxUnit:   "km",
				Sort:      "asc",
			}
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania"}))

			q.BoxWidth, q.BoxHeight = 400, 400
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania", "Palermo"}))

			q.Count = 1
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania"}))

			q.CountAny = true
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Palermo"}))

			q = GeoSearchQuery{
				Longitude:  15,
				Latitude:   37,
				Radius:     100,
				RadiusUnit: "km",
				Sort:       "asc",
			}
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania"}))

			q.Radius = 200
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania", "Palermo"}))

			q.Count = 1
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Catania"}))

			q.CountAny = true
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal([]string{"Palermo"}))
		})
	})

	Describe("marshaling/unmarshaling", func() {
		type convTest struct {
			value  any
			dest   any
			wanted string
		}

		convTests := []convTest{
			// TODO
			// {nil, "", nil},
			{"hello", new(string), "hello"},
			{[]byte("hello"), new([]byte), "hello"},
			{int(1), new(int), "1"},
			{int8(1), new(int8), "1"},
			{int16(1), new(int16), "1"},
			{int32(1), new(int32), "1"},
			{int64(1), new(int64), "1"},
			{uint(1), new(uint), "1"},
			{uint8(1), new(uint8), "1"},
			{uint16(1), new(uint16), "1"},
			{uint32(1), new(uint32), "1"},
			{uint64(1), new(uint64), "1"},
			{float32(1.0), new(float32), "1"},
			{float64(1.0), new(float64), "1"},
			{true, new(bool), "1"},
			{false, new(bool), "0"},
		}

		It("should convert to string", func() {
			for _, test := range convTests {
				err := adapter.Set(ctx, "key", test.value, 0).Err()
				Expect(err).NotTo(HaveOccurred())
				time.Sleep(time.Millisecond * 10)
				s, err := adapter.Cache(time.Hour).Get(ctx, "key").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(s).To(Equal(test.wanted))

				if test.dest == nil {
					continue
				}
				// TODO
				// err = adapter.Cache(time.Hour).Get(ctx, "key").Scan(test.dest)
				// Expect(err).NotTo(HaveOccurred())
				// Expect(deref(test.dest)).To(Equal(test.value))
			}
		})
	})

	Describe("json marshaling/unmarshaling", func() {
		BeforeEach(func() {
			value := &numberStruct{Number: 42}
			err := adapter.Set(ctx, "key", value, 0).Err()
			Expect(err).NotTo(HaveOccurred())
		})

		It("should marshal custom values using json", func() {
			s, err := adapter.Cache(time.Hour).Get(ctx, "key").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(s).To(Equal(`{"Number":42}`))
		})

		// TODO
		// It("should scan custom values using json", func() {
		//	value := &numberStruct{}
		//	err := adapter.Cache(time.Hour).Get(ctx, "key").Scan(value)
		//	Expect(err).NotTo(HaveOccurred())
		//	Expect(value.Number).To(Equal(42))
		// })
	})

	Describe("GearsCmdable", func() {
		BeforeEach(func() {
			Expect(adapter.FlushDB(ctx).Err()).NotTo(HaveOccurred())
			adapter.TFunctionDelete(ctx, "lib1")
			adapter.TFunctionDelete(ctx, "lib2")
		})
		// Copied from go-redis
		// https://github.com/redis/go-redis/blob/f994ff1cd96299a5c8029ae3403af7b17ef06e8a/gears_commands_test.go
		It("should TFunctionLoad, TFunctionLoadArgs and TFunctionDelete ", Label("gears", "tfunctionload"), func() {
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("OK"))
			opt := &TFunctionLoadOptions{Replace: true, Config: `{"last_update_field_name":"last_update"}`}
			resultAdd, err = adapter.TFunctionLoadArgs(ctx, libCodeWithConfig("lib1"), opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("OK"))
			opt.Replace = false
			resultAdd, err = adapter.TFunctionLoadArgs(ctx, libCodeWithConfig("lib2"), opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("OK"))
		})
		It("should TFunctionList", Label("gears", "tfunctionlist"), func() {
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("OK"))
			resultList, err := adapter.TFunctionList(ctx).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultList[0]["engine"]).To(BeEquivalentTo("js"))
			opt := &TFunctionListOptions{Withcode: true, Verbose: 2}
			resultListArgs, err := adapter.TFunctionListArgs(ctx, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultListArgs[0]["code"]).NotTo(BeEquivalentTo(""))
			opt.Library = "VERBOSE"
			resultListArgs, err = adapter.TFunctionListArgs(ctx, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultListArgs[0]["code"]).NotTo(BeEquivalentTo(""))
		})

		It("should TFCall", Label("gears", "tfcall"), func() {
			var resultAdd interface{}
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("OK"))
			resultAdd, err = adapter.TFCall(ctx, "lib1", "foo", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("bar"))
		})

		It("should TFCallArgs", Label("gears", "tfcallargs"), func() {
			var resultAdd interface{}
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("OK"))
			opt := &TFCallOptions{Arguments: []string{"foo", "bar"}}
			resultAdd, err = adapter.TFCallArgs(ctx, "lib1", "foo", 0, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("bar"))
		})

		It("should TFCallASYNC", Label("gears", "TFCallASYNC"), func() {
			var resultAdd interface{}
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("OK"))
			resultAdd, err = adapter.TFCallASYNC(ctx, "lib1", "foo", 0).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("bar"))
		})

		It("should TFCallASYNCArgs", Label("gears", "TFCallASYNCargs"), func() {
			var resultAdd interface{}
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("OK"))
			opt := &TFCallOptions{Arguments: []string{"foo", "bar"}}
			resultAdd, err = adapter.TFCallASYNCArgs(ctx, "lib1", "foo", 0, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo("bar"))
		})
	})
	// https://github.com/redis/go-redis/blob/master/probabilistic_test.go#L14
	Describe("ProbabilisticCmdable", func() {
		ctx := context.TODO()
		BeforeEach(func() {
			Expect(adapter.FlushDB(ctx).Err()).NotTo(HaveOccurred())
		})
		Describe("bloom", Label("bloom"), func() {
			It("should BFAdd", Label("bloom", "bfadd"), func() {
				resultAdd, err := adapter.BFAdd(ctx, "testbf1", 1).Result()

				Expect(err).NotTo(HaveOccurred())
				Expect(resultAdd).To(BeTrue())

				resultInfo, err := adapter.BFInfo(ctx, "testbf1").Result()

				Expect(err).NotTo(HaveOccurred())
				Expect(resultInfo).To(BeAssignableToTypeOf(BFInfo{}))
				Expect(resultInfo.ItemsInserted).To(BeEquivalentTo(int64(1)))
			})

			It("should get Bloom filter information with specific options", Label("bloom", "bfinfoarg"), func() {
				// Set up the test data
				err := adapter.BFAdd(ctx, "testbf1", "element").Err()
				Expect(err).NotTo(HaveOccurred())

				// Test CAPACITY option
				info, err := adapter.BFInfoArg(ctx, "testbf1", "CAPACITY").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info).NotTo(BeNil())

				// Test SIZE option
				info, err = adapter.BFInfoArg(ctx, "testbf1", "SIZE").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info).NotTo(BeNil())

				// Test FILTERS option
				info, err = adapter.BFInfoArg(ctx, "testbf1", "FILTERS").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info).NotTo(BeNil())

				// Test ITEMS option
				info, err = adapter.BFInfoArg(ctx, "testbf1", "ITEMS").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info).NotTo(BeNil())

				// Test EXPANSION option
				info, err = adapter.BFInfoArg(ctx, "testbf1", "EXPANSION").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info).NotTo(BeNil())
			})

			It("should panic on unknown option in BFInfoArg", Label("bloom", "bfinfoarg"), func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("The code did not panic")
					}
				}()

				// This should cause the function to panic due to an unknown option
				adapter.BFInfoArg(ctx, "testbf1", "UNKNOWN_OPTION")
			})

			It("should BFCard", Label("bloom", "bfcard"), func() {
				// This is a probabilistic data structure, and it's not always guaranteed that we will get back
				// the exact number of inserted items, during hash collisions
				// But with such a low number of items (only 3),
				// the probability of a collision is very low, so we can expect to get back the exact number of items
				_, err := adapter.BFAdd(ctx, "testbf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				_, err = adapter.BFAdd(ctx, "testbf1", "item2").Result()
				Expect(err).NotTo(HaveOccurred())
				_, err = adapter.BFAdd(ctx, "testbf1", 3).Result()
				Expect(err).NotTo(HaveOccurred())

				result, err := adapter.BFCard(ctx, "testbf1").Result()

				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeEquivalentTo(int64(3)))
			})

			It("should BFExists", Label("bloom", "bfexists"), func() {
				exists, err := adapter.BFExists(ctx, "testbf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(exists).To(BeFalse())

				_, err = adapter.BFAdd(ctx, "testbf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())

				exists, err = adapter.BFExists(ctx, "testbf1", "item1").Result()

				Expect(err).NotTo(HaveOccurred())
				Expect(exists).To(BeTrue())
			})

			It("should BFInfo and BFReserve", Label("bloom", "bfinfo", "bfreserve"), func() {
				err := adapter.BFReserve(ctx, "testbf1", 0.001, 2000).Err()
				Expect(err).NotTo(HaveOccurred())

				result, err := adapter.BFInfo(ctx, "testbf1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeAssignableToTypeOf(BFInfo{}))
				Expect(result.Capacity).To(BeEquivalentTo(int64(2000)))
			})

			It("should BFInfoCapacity, BFInfoSize, BFInfoFilters, BFInfoItems, BFInfoExpansion, ", Label("bloom", "bfinfocapacity", "bfinfosize", "bfinfofilters", "bfinfoitems", "bfinfoexpansion"), func() {
				err := adapter.BFReserve(ctx, "testbf1", 0.001, 2000).Err()
				Expect(err).NotTo(HaveOccurred())

				result, err := adapter.BFInfoCapacity(ctx, "testbf1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Capacity).To(BeEquivalentTo(int64(2000)))

				result, err = adapter.BFInfoItems(ctx, "testbf1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result.ItemsInserted).To(BeEquivalentTo(int64(0)))

				result, err = adapter.BFInfoSize(ctx, "testbf1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Size).To(BeEquivalentTo(int64(4056)))

				err = adapter.BFReserveExpansion(ctx, "testbf2", 0.001, 2000, 3).Err()
				Expect(err).NotTo(HaveOccurred())

				result, err = adapter.BFInfoFilters(ctx, "testbf2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result.Filters).To(BeEquivalentTo(int64(1)))

				result, err = adapter.BFInfoExpansion(ctx, "testbf2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result.ExpansionRate).To(BeEquivalentTo(int64(3)))
			})

			It("should BFInsert", Label("bloom", "bfinsert"), func() {
				options := &BFInsertOptions{
					Capacity:   2000,
					Error:      0.001,
					Expansion:  3,
					NonScaling: false,
					NoCreate:   true,
				}

				resultInsert, err := adapter.BFInsert(ctx, "testbf1", options, "item1").Result()
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("not found"))

				options = &BFInsertOptions{
					Capacity:   2000,
					Error:      0.001,
					Expansion:  3,
					NonScaling: false,
					NoCreate:   false,
				}

				resultInsert, err = adapter.BFInsert(ctx, "testbf1", options, "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(resultInsert)).To(BeEquivalentTo(1))

				exists, err := adapter.BFExists(ctx, "testbf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(exists).To(BeTrue())

				result, err := adapter.BFInfo(ctx, "testbf1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeAssignableToTypeOf(BFInfo{}))
				Expect(result.Capacity).To(BeEquivalentTo(int64(2000)))
				Expect(result.ExpansionRate).To(BeEquivalentTo(int64(3)))

				options.NonScaling = true
				resultInsert, err = adapter.BFInsert(ctx, "testbf2", options, "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(resultInsert)).To(BeEquivalentTo(1))

				exists, err = adapter.BFExists(ctx, "testbf2", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(exists).To(BeTrue())

				result, err = adapter.BFInfo(ctx, "testbf2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeAssignableToTypeOf(BFInfo{}))
				Expect(result.Capacity).To(BeEquivalentTo(int64(2000)))
				Expect(result.ExpansionRate).To(BeEquivalentTo(int64(0)))
			})

			It("should BFMAdd", Label("bloom", "bfmadd"), func() {
				resultAdd, err := adapter.BFMAdd(ctx, "testbf1", "item1", "item2", "item3").Result()

				Expect(err).NotTo(HaveOccurred())
				Expect(len(resultAdd)).To(Equal(3))

				resultInfo, err := adapter.BFInfo(ctx, "testbf1").Result()

				Expect(err).NotTo(HaveOccurred())
				Expect(resultInfo).To(BeAssignableToTypeOf(BFInfo{}))
				Expect(resultInfo.ItemsInserted).To(BeEquivalentTo(int64(3)))
				resultAdd2, err := adapter.BFMAdd(ctx, "testbf1", "item1", "item2", "item4").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resultAdd2[0]).To(BeFalse())
				Expect(resultAdd2[1]).To(BeFalse())
				Expect(resultAdd2[2]).To(BeTrue())
			})

			It("should BFMExists", Label("bloom", "bfmexists"), func() {
				exist, err := adapter.BFMExists(ctx, "testbf1", "item1", "item2", "item3").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(exist)).To(Equal(3))
				Expect(exist[0]).To(BeFalse())
				Expect(exist[1]).To(BeFalse())
				Expect(exist[2]).To(BeFalse())

				_, err = adapter.BFMAdd(ctx, "testbf1", "item1", "item2", "item3").Result()
				Expect(err).NotTo(HaveOccurred())

				exist, err = adapter.BFMExists(ctx, "testbf1", "item1", "item2", "item3", "item4").Result()

				Expect(err).NotTo(HaveOccurred())
				Expect(len(exist)).To(Equal(4))
				Expect(exist[0]).To(BeTrue())
				Expect(exist[1]).To(BeTrue())
				Expect(exist[2]).To(BeTrue())
				Expect(exist[3]).To(BeFalse())
			})

			It("should BFReserveExpansion", Label("bloom", "bfreserveexpansion"), func() {
				err := adapter.BFReserveExpansion(ctx, "testbf1", 0.001, 2000, 3).Err()
				Expect(err).NotTo(HaveOccurred())

				result, err := adapter.BFInfo(ctx, "testbf1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeAssignableToTypeOf(BFInfo{}))
				Expect(result.Capacity).To(BeEquivalentTo(int64(2000)))
				Expect(result.ExpansionRate).To(BeEquivalentTo(int64(3)))
			})

			It("should BFReserveNonScaling", Label("bloom", "bfreservenonscaling"), func() {
				err := adapter.BFReserveNonScaling(ctx, "testbfns1", 0.001, 1000).Err()
				Expect(err).NotTo(HaveOccurred())

				_, err = adapter.BFInfo(ctx, "testbfns1").Result()
				Expect(err).NotTo(HaveOccurred())
			})

			It("should BFScanDump and BFLoadChunk", Label("bloom", "bfscandump", "bfloadchunk"), func() {
				err := adapter.BFReserve(ctx, "testbfsd1", 0.001, 3000).Err()
				Expect(err).NotTo(HaveOccurred())
				for i := 0; i < 1000; i++ {
					adapter.BFAdd(ctx, "testbfsd1", i)
				}
				infBefore := adapter.BFInfoSize(ctx, "testbfsd1")
				fd := []ScanDump{}
				sd, err := adapter.BFScanDump(ctx, "testbfsd1", 0).Result()
				for {
					if sd.Iter == 0 {
						break
					}
					Expect(err).NotTo(HaveOccurred())
					fd = append(fd, sd)
					sd, err = adapter.BFScanDump(ctx, "testbfsd1", sd.Iter).Result()
				}
				adapter.Del(ctx, "testbfsd1")
				for _, e := range fd {
					adapter.BFLoadChunk(ctx, "testbfsd1", e.Iter, e.Data)
				}
				infAfter := adapter.BFInfoSize(ctx, "testbfsd1")
				Expect(infBefore).To(BeEquivalentTo(infAfter))
			})

			It("should BFReserveWithArgs", Label("bloom", "bfreserveargs"), func() {
				options := &BFReserveOptions{
					Capacity:   2000,
					Error:      0.001,
					Expansion:  3,
					NonScaling: false,
				}
				err := adapter.BFReserveWithArgs(ctx, "testbf", options).Err()
				Expect(err).NotTo(HaveOccurred())

				result, err := adapter.BFInfo(ctx, "testbf").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeAssignableToTypeOf(BFInfo{}))
				Expect(result.Capacity).To(BeEquivalentTo(int64(2000)))
				Expect(result.ExpansionRate).To(BeEquivalentTo(int64(3)))

				options.NonScaling = true
				err = adapter.BFReserveWithArgs(ctx, "testbf2", options).Err()
				Expect(err).To(HaveOccurred())
			})
		})

		Describe("cuckoo", Label("cuckoo"), func() {
			It("should CFAdd", Label("cuckoo", "cfadd"), func() {
				add, err := adapter.CFAdd(ctx, "testcf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(add).To(BeTrue())

				exists, err := adapter.CFExists(ctx, "testcf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(exists).To(BeTrue())

				info, err := adapter.CFInfo(ctx, "testcf1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info).To(BeAssignableToTypeOf(CFInfo{}))
				Expect(info.NumItemsInserted).To(BeEquivalentTo(int64(1)))
			})

			It("should CFAddNX", Label("cuckoo", "cfaddnx"), func() {
				add, err := adapter.CFAddNX(ctx, "testcf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(add).To(BeTrue())

				exists, err := adapter.CFExists(ctx, "testcf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(exists).To(BeTrue())

				result, err := adapter.CFAddNX(ctx, "testcf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeFalse())

				info, err := adapter.CFInfo(ctx, "testcf1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info).To(BeAssignableToTypeOf(CFInfo{}))
				Expect(info.NumItemsInserted).To(BeEquivalentTo(int64(1)))
			})

			It("should CFCount", Label("cuckoo", "cfcount"), func() {
				err := adapter.CFAdd(ctx, "testcf1", "item1").Err()
				cnt, err := adapter.CFCount(ctx, "testcf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cnt).To(BeEquivalentTo(int64(1)))

				err = adapter.CFAdd(ctx, "testcf1", "item1").Err()
				Expect(err).NotTo(HaveOccurred())

				cnt, err = adapter.CFCount(ctx, "testcf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cnt).To(BeEquivalentTo(int64(2)))
			})

			It("should CFDel and CFExists", Label("cuckoo", "cfdel", "cfexists"), func() {
				err := adapter.CFAdd(ctx, "testcf1", "item1").Err()
				Expect(err).NotTo(HaveOccurred())

				exists, err := adapter.CFExists(ctx, "testcf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(exists).To(BeTrue())

				del, err := adapter.CFDel(ctx, "testcf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(del).To(BeTrue())

				exists, err = adapter.CFExists(ctx, "testcf1", "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(exists).To(BeFalse())
			})

			It("should CFInfo and CFReserve", Label("cuckoo", "cfinfo", "cfreserve"), func() {
				err := adapter.CFReserve(ctx, "testcf1", 1000).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.CFReserveExpansion(ctx, "testcfe1", 1000, 1).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.CFReserveBucketSize(ctx, "testcfbs1", 1000, 4).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.CFReserveMaxIterations(ctx, "testcfmi1", 1000, 10).Err()
				Expect(err).NotTo(HaveOccurred())

				result, err := adapter.CFInfo(ctx, "testcf1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeAssignableToTypeOf(CFInfo{}))
			})

			It("should CFScanDump and CFLoadChunk", Label("bloom", "cfscandump", "cfloadchunk"), func() {
				err := adapter.CFReserve(ctx, "testcfsd1", 1000).Err()
				Expect(err).NotTo(HaveOccurred())
				for i := 0; i < 1000; i++ {
					Item := fmt.Sprintf("item%d", i)
					adapter.CFAdd(ctx, "testcfsd1", Item)
				}
				infBefore := adapter.CFInfo(ctx, "testcfsd1")
				fd := []ScanDump{}
				sd, err := adapter.CFScanDump(ctx, "testcfsd1", 0).Result()
				for {
					if sd.Iter == 0 {
						break
					}
					Expect(err).NotTo(HaveOccurred())
					fd = append(fd, sd)
					sd, err = adapter.CFScanDump(ctx, "testcfsd1", sd.Iter).Result()
				}
				adapter.Del(ctx, "testcfsd1")
				for _, e := range fd {
					adapter.CFLoadChunk(ctx, "testcfsd1", e.Iter, e.Data)
				}
				infAfter := adapter.CFInfo(ctx, "testcfsd1")
				Expect(infBefore).To(BeEquivalentTo(infAfter))
			})

			It("should CFInfo and CFReserveWithArgs", Label("cuckoo", "cfinfo", "cfreserveargs"), func() {
				args := &CFReserveOptions{
					Capacity:      2048,
					BucketSize:    3,
					MaxIterations: 15,
					Expansion:     2,
				}

				err := adapter.CFReserveWithArgs(ctx, "testcf1", args).Err()
				Expect(err).NotTo(HaveOccurred())

				result, err := adapter.CFInfo(ctx, "testcf1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeAssignableToTypeOf(CFInfo{}))
				Expect(result.BucketSize).To(BeEquivalentTo(int64(3)))
				Expect(result.MaxIteration).To(BeEquivalentTo(int64(15)))
				Expect(result.ExpansionRate).To(BeEquivalentTo(int64(2)))
			})

			It("should CFInsert", Label("cuckoo", "cfinsert"), func() {
				args := &CFInsertOptions{
					Capacity: 3000,
					NoCreate: true,
				}

				result, err := adapter.CFInsert(ctx, "testcf1", args, "item1", "item2", "item3").Result()
				Expect(err).To(HaveOccurred())

				args = &CFInsertOptions{
					Capacity: 3000,
					NoCreate: false,
				}

				result, err = adapter.CFInsert(ctx, "testcf1", args, "item1", "item2", "item3").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(result)).To(BeEquivalentTo(3))
			})

			It("should CFInsertNX", Label("cuckoo", "cfinsertnx"), func() {
				args := &CFInsertOptions{
					Capacity: 3000,
					NoCreate: true,
				}

				result, err := adapter.CFInsertNX(ctx, "testcf1", args, "item1", "item2", "item2").Result()
				Expect(err).To(HaveOccurred())

				args = &CFInsertOptions{
					Capacity: 3000,
					NoCreate: false,
				}

				result, err = adapter.CFInsertNX(ctx, "testcf2", args, "item1", "item2", "item2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(result)).To(BeEquivalentTo(3))
				Expect(result[0]).To(BeEquivalentTo(int64(1)))
				Expect(result[1]).To(BeEquivalentTo(int64(1)))
				Expect(result[2]).To(BeEquivalentTo(int64(0)))
			})

			It("should CFMexists", Label("cuckoo", "cfmexists"), func() {
				err := adapter.CFInsert(ctx, "testcf1", nil, "item1", "item2", "item3").Err()
				Expect(err).NotTo(HaveOccurred())

				result, err := adapter.CFMExists(ctx, "testcf1", "item1", "item2", "item3", "item4").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(result)).To(BeEquivalentTo(4))
				Expect(result[0]).To(BeTrue())
				Expect(result[1]).To(BeTrue())
				Expect(result[2]).To(BeTrue())
				Expect(result[3]).To(BeFalse())
			})
		})

		Describe("CMS", Label("cms"), func() {
			It("should CMSIncrBy", Label("cms", "cmsincrby"), func() {
				err := adapter.CMSInitByDim(ctx, "testcms1", 5, 10).Err()
				Expect(err).NotTo(HaveOccurred())

				result, err := adapter.CMSIncrBy(ctx, "testcms1", "item1", 1, "item2", 2, "item3", 3).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(result)).To(BeEquivalentTo(3))
				Expect(result[0]).To(BeEquivalentTo(int64(1)))
				Expect(result[1]).To(BeEquivalentTo(int64(2)))
				Expect(result[2]).To(BeEquivalentTo(int64(3)))
			})

			It("should CMSInitByDim and CMSInfo", Label("cms", "cmsinitbydim", "cmsinfo"), func() {
				err := adapter.CMSInitByDim(ctx, "testcms1", 5, 10).Err()
				Expect(err).NotTo(HaveOccurred())

				info, err := adapter.CMSInfo(ctx, "testcms1").Result()
				Expect(err).NotTo(HaveOccurred())

				Expect(info).To(BeAssignableToTypeOf(CMSInfo{}))
				Expect(info.Width).To(BeEquivalentTo(int64(5)))
				Expect(info.Depth).To(BeEquivalentTo(int64(10)))
			})

			It("should CMSInitByProb", Label("cms", "cmsinitbyprob"), func() {
				err := adapter.CMSInitByProb(ctx, "testcms1", 0.002, 0.01).Err()
				Expect(err).NotTo(HaveOccurred())

				info, err := adapter.CMSInfo(ctx, "testcms1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info).To(BeAssignableToTypeOf(CMSInfo{}))
			})

			It("should CMSMerge, CMSMergeWithWeight and CMSQuery", Label("cms", "cmsmerge", "cmsquery"), func() {
				err := adapter.CMSMerge(ctx, "destCms1", "testcms2", "testcms3").Err()
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("CMS: key does not exist"))

				err = adapter.CMSInitByDim(ctx, "destCms1", 5, 10).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.CMSInitByDim(ctx, "destCms2", 5, 10).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.CMSInitByDim(ctx, "cms1", 2, 20).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.CMSInitByDim(ctx, "cms2", 3, 20).Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.CMSMerge(ctx, "destCms1", "cms1", "cms2").Err()
				Expect(err).To(MatchError("CMS: width/depth is not equal"))

				adapter.Del(ctx, "cms1", "cms2")

				err = adapter.CMSInitByDim(ctx, "cms1", 5, 10).Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.CMSInitByDim(ctx, "cms2", 5, 10).Err()
				Expect(err).NotTo(HaveOccurred())

				adapter.CMSIncrBy(ctx, "cms1", "item1", 1, "item2", 2)
				adapter.CMSIncrBy(ctx, "cms2", "item2", 2, "item3", 3)

				err = adapter.CMSMerge(ctx, "destCms1", "cms1", "cms2").Err()
				Expect(err).NotTo(HaveOccurred())

				result, err := adapter.CMSQuery(ctx, "destCms1", "item1", "item2", "item3").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(result)).To(BeEquivalentTo(3))
				Expect(result[0]).To(BeEquivalentTo(int64(1)))
				Expect(result[1]).To(BeEquivalentTo(int64(4)))
				Expect(result[2]).To(BeEquivalentTo(int64(3)))

				sourceSketches := map[string]int64{
					"cms1": 1,
					"cms2": 2,
				}
				err = adapter.CMSMergeWithWeight(ctx, "destCms2", sourceSketches).Err()
				Expect(err).NotTo(HaveOccurred())

				result, err = adapter.CMSQuery(ctx, "destCms2", "item1", "item2", "item3").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(result)).To(BeEquivalentTo(3))
				Expect(result[0]).To(BeEquivalentTo(int64(1)))
				Expect(result[1]).To(BeEquivalentTo(int64(6)))
				Expect(result[2]).To(BeEquivalentTo(int64(6)))
			})
		})

		Describe("TopK", Label("topk"), func() {
			It("should TopKReserve, TopKInfo, TopKAdd, TopKQuery, TopKCount, TopKIncrBy, TopKList, TopKListWithCount", Label("topk", "topkreserve", "topkinfo", "topkadd", "topkquery", "topkcount", "topkincrby", "topklist", "topklistwithcount"), func() {
				err := adapter.TopKReserve(ctx, "topk1", 3).Err()
				Expect(err).NotTo(HaveOccurred())

				resultInfo, err := adapter.TopKInfo(ctx, "topk1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resultInfo.K).To(BeEquivalentTo(int64(3)))

				resultAdd, err := adapter.TopKAdd(ctx, "topk1", "item1", "item2", 3, "item1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(resultAdd)).To(BeEquivalentTo(int64(4)))

				resultQuery, err := adapter.TopKQuery(ctx, "topk1", "item1", "item2", 4, 3).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(resultQuery)).To(BeEquivalentTo(4))
				Expect(resultQuery[0]).To(BeTrue())
				Expect(resultQuery[1]).To(BeTrue())
				Expect(resultQuery[2]).To(BeFalse())
				Expect(resultQuery[3]).To(BeTrue())

				resultCount, err := adapter.TopKCount(ctx, "topk1", "item1", "item2", "item3").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(resultCount)).To(BeEquivalentTo(3))
				Expect(resultCount[0]).To(BeEquivalentTo(int64(2)))
				Expect(resultCount[1]).To(BeEquivalentTo(int64(1)))
				Expect(resultCount[2]).To(BeEquivalentTo(int64(0)))

				resultIncr, err := adapter.TopKIncrBy(ctx, "topk1", "item1", 5, "item2", 10).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(resultIncr)).To(BeEquivalentTo(2))

				resultCount, err = adapter.TopKCount(ctx, "topk1", "item1", "item2", "item3").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(resultCount)).To(BeEquivalentTo(3))
				Expect(resultCount[0]).To(BeEquivalentTo(int64(7)))
				Expect(resultCount[1]).To(BeEquivalentTo(int64(11)))
				Expect(resultCount[2]).To(BeEquivalentTo(int64(0)))

				resultList, err := adapter.TopKList(ctx, "topk1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(resultList)).To(BeEquivalentTo(3))
				Expect(resultList).To(ContainElements("item2", "item1", "3"))

				resultListWithCount, err := adapter.TopKListWithCount(ctx, "topk1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(resultListWithCount)).To(BeEquivalentTo(3))
				Expect(resultListWithCount["3"]).To(BeEquivalentTo(int64(1)))
				Expect(resultListWithCount["item1"]).To(BeEquivalentTo(int64(7)))
				Expect(resultListWithCount["item2"]).To(BeEquivalentTo(int64(11)))
			})

			It("should TopKReserveWithOptions", Label("topk", "topkreservewithoptions"), func() {
				err := adapter.TopKReserveWithOptions(ctx, "topk1", 3, 1500, 8, 0.5).Err()
				Expect(err).NotTo(HaveOccurred())

				resultInfo, err := adapter.TopKInfo(ctx, "topk1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resultInfo.K).To(BeEquivalentTo(int64(3)))
				Expect(resultInfo.Width).To(BeEquivalentTo(int64(1500)))
				Expect(resultInfo.Depth).To(BeEquivalentTo(int64(8)))
				Expect(resultInfo.Decay).To(BeEquivalentTo(0.5))
			})
		})

		Describe("t-digest", Label("tdigest"), func() {
			It("should TDigestAdd, TDigestCreate, TDigestInfo, TDigestByRank, TDigestByRevRank, TDigestCDF, TDigestMax, TDigestMin, TDigestQuantile, TDigestRank, TDigestRevRank, TDigestTrimmedMean, TDigestReset, ", Label("tdigest", "tdigestadd", "tdigestcreate", "tdigestinfo", "tdigestbyrank", "tdigestbyrevrank", "tdigestcdf", "tdigestmax", "tdigestmin", "tdigestquantile", "tdigestrank", "tdigestrevrank", "tdigesttrimmedmean", "tdigestreset"), func() {
				err := adapter.TDigestCreate(ctx, "tdigest1").Err()
				Expect(err).NotTo(HaveOccurred())

				info, err := adapter.TDigestInfo(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info.Observations).To(BeEquivalentTo(int64(0)))

				// Test with empty sketch
				byRank, err := adapter.TDigestByRank(ctx, "tdigest1", 0, 1, 2, 3).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(byRank)).To(BeEquivalentTo(4))

				byRevRank, err := adapter.TDigestByRevRank(ctx, "tdigest1", 0, 1, 2).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(byRevRank)).To(BeEquivalentTo(3))

				cdf, err := adapter.TDigestCDF(ctx, "tdigest1", 15, 35, 70).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(cdf)).To(BeEquivalentTo(3))

				max, err := adapter.TDigestMax(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(math.IsNaN(max)).To(BeTrue())

				min, err := adapter.TDigestMin(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(math.IsNaN(min)).To(BeTrue())

				quantile, err := adapter.TDigestQuantile(ctx, "tdigest1", 0.1, 0.2).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(quantile)).To(BeEquivalentTo(2))

				rank, err := adapter.TDigestRank(ctx, "tdigest1", 10, 20).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(rank)).To(BeEquivalentTo(2))

				revRank, err := adapter.TDigestRevRank(ctx, "tdigest1", 10, 20).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(revRank)).To(BeEquivalentTo(2))

				trimmedMean, err := adapter.TDigestTrimmedMean(ctx, "tdigest1", 0.1, 0.6).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(math.IsNaN(trimmedMean)).To(BeTrue())

				// Add elements
				err = adapter.TDigestAdd(ctx, "tdigest1", 10, 20, 30, 40, 50, 60, 70, 80, 90, 100).Err()
				Expect(err).NotTo(HaveOccurred())

				info, err = adapter.TDigestInfo(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info.Observations).To(BeEquivalentTo(int64(10)))

				byRank, err = adapter.TDigestByRank(ctx, "tdigest1", 0, 1, 2).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(byRank)).To(BeEquivalentTo(3))
				Expect(byRank[0]).To(BeEquivalentTo(float64(10)))
				Expect(byRank[1]).To(BeEquivalentTo(float64(20)))
				Expect(byRank[2]).To(BeEquivalentTo(float64(30)))

				byRevRank, err = adapter.TDigestByRevRank(ctx, "tdigest1", 0, 1, 2).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(byRevRank)).To(BeEquivalentTo(3))
				Expect(byRevRank[0]).To(BeEquivalentTo(float64(100)))
				Expect(byRevRank[1]).To(BeEquivalentTo(float64(90)))
				Expect(byRevRank[2]).To(BeEquivalentTo(float64(80)))

				cdf, err = adapter.TDigestCDF(ctx, "tdigest1", 15, 35, 70).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(cdf)).To(BeEquivalentTo(3))
				Expect(cdf[0]).To(BeEquivalentTo(0.1))
				Expect(cdf[1]).To(BeEquivalentTo(0.3))
				Expect(cdf[2]).To(BeEquivalentTo(0.65))

				max, err = adapter.TDigestMax(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(max).To(BeEquivalentTo(float64(100)))

				min, err = adapter.TDigestMin(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(min).To(BeEquivalentTo(float64(10)))

				quantile, err = adapter.TDigestQuantile(ctx, "tdigest1", 0.1, 0.2).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(quantile)).To(BeEquivalentTo(2))
				Expect(quantile[0]).To(BeEquivalentTo(float64(20)))
				Expect(quantile[1]).To(BeEquivalentTo(float64(30)))

				rank, err = adapter.TDigestRank(ctx, "tdigest1", 10, 20).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(rank)).To(BeEquivalentTo(2))
				Expect(rank[0]).To(BeEquivalentTo(int64(0)))
				Expect(rank[1]).To(BeEquivalentTo(int64(1)))

				revRank, err = adapter.TDigestRevRank(ctx, "tdigest1", 10, 20).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(len(revRank)).To(BeEquivalentTo(2))
				Expect(revRank[0]).To(BeEquivalentTo(int64(9)))
				Expect(revRank[1]).To(BeEquivalentTo(int64(8)))

				trimmedMean, err = adapter.TDigestTrimmedMean(ctx, "tdigest1", 0.1, 0.6).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(trimmedMean).To(BeEquivalentTo(float64(40)))

				reset, err := adapter.TDigestReset(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(reset).To(BeEquivalentTo("OK"))
			})

			It("should TDigestCreateWithCompression", Label("tdigest", "tcreatewithcompression"), func() {
				err := adapter.TDigestCreateWithCompression(ctx, "tdigest1", 2000).Err()
				Expect(err).NotTo(HaveOccurred())

				info, err := adapter.TDigestInfo(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info.Compression).To(BeEquivalentTo(int64(2000)))
			})

			It("should TDigestMerge", Label("tdigest", "tmerge"), func() {
				err := adapter.TDigestCreate(ctx, "tdigest1").Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.TDigestAdd(ctx, "tdigest1", 10, 20, 30, 40, 50, 60, 70, 80, 90, 100).Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.TDigestCreate(ctx, "tdigest2").Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.TDigestAdd(ctx, "tdigest2", 15, 25, 35, 45, 55, 65, 75, 85, 95, 105).Err()
				Expect(err).NotTo(HaveOccurred())

				err = adapter.TDigestCreate(ctx, "tdigest3").Err()
				Expect(err).NotTo(HaveOccurred())
				err = adapter.TDigestAdd(ctx, "tdigest3", 50, 60, 70, 80, 90, 100, 110, 120, 130, 140).Err()
				Expect(err).NotTo(HaveOccurred())

				options := &TDigestMergeOptions{
					Compression: 1000,
					Override:    false,
				}
				err = adapter.TDigestMerge(ctx, "tdigest1", options, "tdigest2", "tdigest3").Err()
				Expect(err).NotTo(HaveOccurred())

				info, err := adapter.TDigestInfo(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info.Observations).To(BeEquivalentTo(int64(30)))
				Expect(info.Compression).To(BeEquivalentTo(int64(1000)))

				max, err := adapter.TDigestMax(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(max).To(BeEquivalentTo(float64(140)))

				options.Override = true
				err = adapter.TDigestMerge(ctx, "tdigest1", options, "tdigest2", "tdigest3").Err()
				Expect(err).NotTo(HaveOccurred())

				info, err = adapter.TDigestInfo(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(info.Observations).To(BeEquivalentTo(int64(20)))
				Expect(info.Compression).To(BeEquivalentTo(int64(1000)))

				max, err = adapter.TDigestMax(ctx, "tdigest1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(max).To(BeEquivalentTo(float64(140)))
			})
		})
	})
	Describe("RedisTimeseries commands", Label("timeseries"), func() {
		ctx := context.TODO()

		BeforeEach(func() {
			Expect(adapter.FlushDB(ctx).Err()).NotTo(HaveOccurred())
		})

		It("should TSCreate and TSCreateWithArgs", Label("timeseries", "tscreate", "tscreateWithArgs"), func() {
			result, err := adapter.TSCreate(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo("OK"))
			// Test TSCreateWithArgs
			opt := &TSOptions{Retention: 5}
			result, err = adapter.TSCreateWithArgs(ctx, "2", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"Redis": "Labs"}}
			result, err = adapter.TSCreateWithArgs(ctx, "3", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"Time": "Series"}, Retention: 20}
			result, err = adapter.TSCreateWithArgs(ctx, "4", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo("OK"))
			resultInfo, err := adapter.TSInfo(ctx, "4").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["labels"].(map[string]interface{})["Time"]).To(BeEquivalentTo("Series"))
			// Test chunk size
			opt = &TSOptions{ChunkSize: 128}
			result, err = adapter.TSCreateWithArgs(ctx, "ts-cs-1", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo("OK"))
			resultInfo, err = adapter.TSInfo(ctx, "ts-cs-1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["chunkSize"]).To(BeEquivalentTo(128))
			// Test duplicate policy
			duplicate_policies := []string{"BLOCK", "LAST", "FIRST", "MIN", "MAX"}
			for _, dup := range duplicate_policies {
				keyName := "ts-dup-" + dup
				opt = &TSOptions{DuplicatePolicy: dup}
				result, err = adapter.TSCreateWithArgs(ctx, keyName, opt).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeEquivalentTo("OK"))
				resultInfo, err = adapter.TSInfo(ctx, keyName).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(strings.ToUpper(resultInfo["duplicatePolicy"].(string))).To(BeEquivalentTo(dup))

			}
		})
		It("should TSAdd and TSAddWithArgs", Label("timeseries", "tsadd", "tsaddWithArgs"), func() {
			result, err := adapter.TSAdd(ctx, "1", 1, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(1))
			// Test TSAddWithArgs
			opt := &TSOptions{Retention: 10}
			result, err = adapter.TSAddWithArgs(ctx, "2", 2, 3, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(2))
			opt = &TSOptions{Labels: map[string]string{"Redis": "Labs"}}
			result, err = adapter.TSAddWithArgs(ctx, "3", 3, 2, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(3))
			opt = &TSOptions{Labels: map[string]string{"Redis": "Labs", "Time": "Series"}, Retention: 10}
			result, err = adapter.TSAddWithArgs(ctx, "4", 4, 2, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(4))
			resultInfo, err := adapter.TSInfo(ctx, "4").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["labels"].(map[string]interface{})["Time"]).To(BeEquivalentTo("Series"))
			// Test chunk size
			opt = &TSOptions{ChunkSize: 128}
			result, err = adapter.TSAddWithArgs(ctx, "ts-cs-1", 1, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(1))
			resultInfo, err = adapter.TSInfo(ctx, "ts-cs-1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["chunkSize"]).To(BeEquivalentTo(128))
			// Test duplicate policy
			// LAST
			opt = &TSOptions{DuplicatePolicy: "LAST"}
			result, err = adapter.TSAddWithArgs(ctx, "tsal-1", 1, 5, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(1))
			result, err = adapter.TSAddWithArgs(ctx, "tsal-1", 1, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(1))
			resultGet, err := adapter.TSGet(ctx, "tsal-1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet.Value).To(BeEquivalentTo(10))
			// FIRST
			opt = &TSOptions{DuplicatePolicy: "FIRST"}
			result, err = adapter.TSAddWithArgs(ctx, "tsaf-1", 1, 5, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(1))
			result, err = adapter.TSAddWithArgs(ctx, "tsaf-1", 1, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(1))
			resultGet, err = adapter.TSGet(ctx, "tsaf-1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet.Value).To(BeEquivalentTo(5))
			// MAX
			opt = &TSOptions{DuplicatePolicy: "MAX"}
			result, err = adapter.TSAddWithArgs(ctx, "tsam-1", 1, 5, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(1))
			result, err = adapter.TSAddWithArgs(ctx, "tsam-1", 1, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(1))
			resultGet, err = adapter.TSGet(ctx, "tsam-1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet.Value).To(BeEquivalentTo(10))
			// MIN
			opt = &TSOptions{DuplicatePolicy: "MIN"}
			result, err = adapter.TSAddWithArgs(ctx, "tsami-1", 1, 5, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(1))
			result, err = adapter.TSAddWithArgs(ctx, "tsami-1", 1, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo(1))
			resultGet, err = adapter.TSGet(ctx, "tsami-1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet.Value).To(BeEquivalentTo(5))
		})

		It("should TSAlter", Label("timeseries", "tsalter"), func() {
			result, err := adapter.TSCreate(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo("OK"))
			resultInfo, err := adapter.TSInfo(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["retentionTime"]).To(BeEquivalentTo(0))

			opt := &TSAlterOptions{Retention: 10}
			resultAlter, err := adapter.TSAlter(ctx, "1", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAlter).To(BeEquivalentTo("OK"))

			resultInfo, err = adapter.TSInfo(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["retentionTime"]).To(BeEquivalentTo(10))

			resultInfo, err = adapter.TSInfo(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["labels"]).To(BeEquivalentTo(map[string]interface{}{}))

			opt = &TSAlterOptions{Labels: map[string]string{"Time": "Series"}}
			resultAlter, err = adapter.TSAlter(ctx, "1", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAlter).To(BeEquivalentTo("OK"))

			resultInfo, err = adapter.TSInfo(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["labels"].(map[string]interface{})["Time"]).To(BeEquivalentTo("Series"))
			Expect(resultInfo["retentionTime"]).To(BeEquivalentTo(10))
			Expect(resultInfo["duplicatePolicy"]).To(BeNil())
			opt = &TSAlterOptions{DuplicatePolicy: "min"}
			resultAlter, err = adapter.TSAlter(ctx, "1", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAlter).To(BeEquivalentTo("OK"))

			resultInfo, err = adapter.TSInfo(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["duplicatePolicy"]).To(BeEquivalentTo("min"))
		})

		It("should TSCreateRule and TSDeleteRule", Label("timeseries", "tscreaterule", "tsdeleterule"), func() {
			result, err := adapter.TSCreate(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo("OK"))
			result, err = adapter.TSCreate(ctx, "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo("OK"))
			result, err = adapter.TSCreateRule(ctx, "1", "2", Avg, 100).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo("OK"))
			for i := 0; i < 50; i++ {
				resultAdd, err := adapter.TSAdd(ctx, "1", 100+i*2, 1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resultAdd).To(BeEquivalentTo(100 + i*2))
				resultAdd, err = adapter.TSAdd(ctx, "1", 100+i*2+1, 2).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resultAdd).To(BeEquivalentTo(100 + i*2 + 1))

			}
			resultAdd, err := adapter.TSAdd(ctx, "1", 100*2, 1.5).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultAdd).To(BeEquivalentTo(100 * 2))
			resultGet, err := adapter.TSGet(ctx, "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet.Value).To(BeEquivalentTo(1.5))
			Expect(resultGet.Timestamp).To(BeEquivalentTo(100))

			resultDeleteRule, err := adapter.TSDeleteRule(ctx, "1", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultDeleteRule).To(BeEquivalentTo("OK"))
			resultInfo, err := adapter.TSInfo(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["rules"]).To(BeEquivalentTo(map[string]interface{}{}))
		})

		It("should TSIncrBy, TSIncrByWithArgs, TSDecrBy and TSDecrByWithArgs", Label("timeseries", "tsincrby", "tsdecrby", "tsincrbyWithArgs", "tsdecrbyWithArgs"), func() {
			for i := 0; i < 100; i++ {
				_, err := adapter.TSIncrBy(ctx, "1", 1).Result()
				Expect(err).NotTo(HaveOccurred())
			}
			result, err := adapter.TSGet(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Value).To(BeEquivalentTo(100))

			for i := 0; i < 100; i++ {
				_, err := adapter.TSDecrBy(ctx, "1", 1).Result()
				Expect(err).NotTo(HaveOccurred())
			}
			result, err = adapter.TSGet(ctx, "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Value).To(BeEquivalentTo(0))

			opt := &TSIncrDecrOptions{Timestamp: 5}
			_, err = adapter.TSIncrByWithArgs(ctx, "2", 1.5, opt).Result()
			Expect(err).NotTo(HaveOccurred())

			result, err = adapter.TSGet(ctx, "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Timestamp).To(BeEquivalentTo(5))
			Expect(result.Value).To(BeEquivalentTo(1.5))

			opt = &TSIncrDecrOptions{Timestamp: 7}
			_, err = adapter.TSIncrByWithArgs(ctx, "2", 2.25, opt).Result()
			Expect(err).NotTo(HaveOccurred())

			result, err = adapter.TSGet(ctx, "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Timestamp).To(BeEquivalentTo(7))
			Expect(result.Value).To(BeEquivalentTo(3.75))

			opt = &TSIncrDecrOptions{Timestamp: 15}
			_, err = adapter.TSDecrByWithArgs(ctx, "2", 1.5, opt).Result()
			Expect(err).NotTo(HaveOccurred())

			result, err = adapter.TSGet(ctx, "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Timestamp).To(BeEquivalentTo(15))
			Expect(result.Value).To(BeEquivalentTo(2.25))

			// Test chunk size INCRBY
			opt = &TSIncrDecrOptions{ChunkSize: 128}
			_, err = adapter.TSIncrByWithArgs(ctx, "3", 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())

			resultInfo, err := adapter.TSInfo(ctx, "3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["chunkSize"]).To(BeEquivalentTo(128))

			// Test chunk size DECRBY
			opt = &TSIncrDecrOptions{ChunkSize: 128}
			_, err = adapter.TSDecrByWithArgs(ctx, "4", 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())

			resultInfo, err = adapter.TSInfo(ctx, "4").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultInfo["chunkSize"]).To(BeEquivalentTo(128))
		})

		It("should TSGet", Label("timeseries", "tsget"), func() {
			opt := &TSOptions{DuplicatePolicy: "max"}
			resultGet, err := adapter.TSAddWithArgs(ctx, "foo", 2265985, 151, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet).To(BeEquivalentTo(2265985))
			result, err := adapter.TSGet(ctx, "foo").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Timestamp).To(BeEquivalentTo(2265985))
			Expect(result.Value).To(BeEquivalentTo(151))
		})

		It("should TSGet Latest", Label("timeseries", "tsgetlatest"), func() {
			resultGet, err := adapter.TSCreate(ctx, "tsgl-1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet).To(BeEquivalentTo("OK"))
			resultGet, err = adapter.TSCreate(ctx, "tsgl-2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet).To(BeEquivalentTo("OK"))
			resultGet, err = adapter.TSCreateRule(ctx, "tsgl-1", "tsgl-2", Sum, 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet).To(BeEquivalentTo("OK"))
			_, err = adapter.TSAdd(ctx, "tsgl-1", 1, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "tsgl-1", 2, 3).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "tsgl-1", 11, 7).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "tsgl-1", 13, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			result, errGet := adapter.TSGet(ctx, "tsgl-2").Result()
			Expect(errGet).NotTo(HaveOccurred())
			Expect(result.Timestamp).To(BeEquivalentTo(0))
			Expect(result.Value).To(BeEquivalentTo(4))
			result, errGet = adapter.TSGetWithArgs(ctx, "tsgl-2", &TSGetOptions{Latest: true}).Result()
			Expect(errGet).NotTo(HaveOccurred())
			Expect(result.Timestamp).To(BeEquivalentTo(10))
			Expect(result.Value).To(BeEquivalentTo(8))
		})

		It("should TSInfo", Label("timeseries", "tsinfo"), func() {
			resultGet, err := adapter.TSAdd(ctx, "foo", 2265985, 151).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet).To(BeEquivalentTo(2265985))
			result, err := adapter.TSInfo(ctx, "foo").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["firstTimestamp"]).To(BeEquivalentTo(2265985))
		})

		It("should TSMAdd", Label("timeseries", "tsmadd"), func() {
			resultGet, err := adapter.TSCreate(ctx, "a").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultGet).To(BeEquivalentTo("OK"))
			ktvSlices := make([][]interface{}, 3)
			for i := 0; i < 3; i++ {
				ktvSlices[i] = make([]interface{}, 3)
				ktvSlices[i][0] = "a"
				for j := 1; j < 3; j++ {
					ktvSlices[i][j] = (i + j) * j
				}
			}
			result, err := adapter.TSMAdd(ctx, ktvSlices).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo([]int64{1, 2, 3}))
		})

		It("should TSMGet and TSMGetWithArgs", Label("timeseries", "tsmget", "tsmgetWithArgs"), func() {
			opt := &TSOptions{Labels: map[string]string{"Test": "This"}}
			resultCreate, err := adapter.TSCreateWithArgs(ctx, "a", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"Test": "This", "Taste": "That"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			_, err = adapter.TSAdd(ctx, "a", "*", 15).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "b", "*", 25).Result()
			Expect(err).NotTo(HaveOccurred())

			result, err := adapter.TSMGet(ctx, []string{"Test=This"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][1].([]interface{})[1]).To(BeEquivalentTo(15))
			Expect(result["b"][1].([]interface{})[1]).To(BeEquivalentTo(25))
			mgetOpt := &TSMGetOptions{WithLabels: true}
			result, err = adapter.TSMGetWithArgs(ctx, []string{"Test=This"}, mgetOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["b"][0]).To(BeEquivalentTo(map[string]interface{}{"Test": "This", "Taste": "That"}))

			resultCreate, err = adapter.TSCreate(ctx, "c").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"is_compaction": "true"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "d", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			resultCreateRule, err := adapter.TSCreateRule(ctx, "c", "d", Sum, 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreateRule).To(BeEquivalentTo("OK"))
			_, err = adapter.TSAdd(ctx, "c", 1, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 2, 3).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 11, 7).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 13, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			result, err = adapter.TSMGet(ctx, []string{"is_compaction=true"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["d"][1]).To(BeEquivalentTo([]interface{}{int64(0), 4.0}))
			mgetOpt = &TSMGetOptions{Latest: true}
			result, err = adapter.TSMGetWithArgs(ctx, []string{"is_compaction=true"}, mgetOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["d"][1]).To(BeEquivalentTo([]interface{}{int64(10), 8.0}))
		})

		It("should TSQueryIndex", Label("timeseries", "tsqueryindex"), func() {
			opt := &TSOptions{Labels: map[string]string{"Test": "This"}}
			resultCreate, err := adapter.TSCreateWithArgs(ctx, "a", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"Test": "This", "Taste": "That"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			result, err := adapter.TSQueryIndex(ctx, []string{"Test=This"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(2))
			result, err = adapter.TSQueryIndex(ctx, []string{"Taste=That"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(1))
		})

		It("should TSDel and TSRange", Label("timeseries", "tsdel", "tsrange"), func() {
			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i, float64(i%7)).Result()
				Expect(err).NotTo(HaveOccurred())
			}
			resultDelete, err := adapter.TSDel(ctx, "a", 0, 21).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultDelete).To(BeEquivalentTo(22))

			resultRange, err := adapter.TSRange(ctx, "a", 0, 21).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange).To(BeEquivalentTo([]TSTimestampValue{}))

			resultRange, err = adapter.TSRange(ctx, "a", 22, 22).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 22, Value: 1}))
		})

		It("should TSRange, TSRangeWithArgs", Label("timeseries", "tsrange", "tsrangeWithArgs"), func() {
			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i, float64(i%7)).Result()
				Expect(err).NotTo(HaveOccurred())

			}
			result, err := adapter.TSRange(ctx, "a", 0, 200).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(100))
			for i := 0; i < 100; i++ {
				adapter.TSAdd(ctx, "a", i+200, float64(i%7))
			}
			result, err = adapter.TSRange(ctx, "a", 0, 500).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(200))
			fts := make([]int, 0)
			for i := 10; i < 20; i++ {
				fts = append(fts, i)
			}
			opt := &TSRangeOptions{FilterByTS: fts, FilterByValue: []int{1, 2}}
			result, err = adapter.TSRangeWithArgs(ctx, "a", 0, 500, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(2))
			opt = &TSRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "+"}
			result, err = adapter.TSRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo([]TSTimestampValue{{Timestamp: 0, Value: 10}, {Timestamp: 10, Value: 1}}))
			opt = &TSRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "5"}
			result, err = adapter.TSRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo([]TSTimestampValue{{Timestamp: 0, Value: 5}, {Timestamp: 5, Value: 6}}))
			opt = &TSRangeOptions{Aggregator: Twa, BucketDuration: 10}
			result, err = adapter.TSRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo([]TSTimestampValue{{Timestamp: 0, Value: 2.55}, {Timestamp: 10, Value: 3}}))
			// Test Range Latest
			resultCreate, err := adapter.TSCreate(ctx, "t1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			resultCreate, err = adapter.TSCreate(ctx, "t2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			resultRule, err := adapter.TSCreateRule(ctx, "t1", "t2", Sum, 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRule).To(BeEquivalentTo("OK"))
			_, errAdd := adapter.TSAdd(ctx, "t1", 1, 1).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 2, 3).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 11, 7).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 13, 1).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			resultRange, err := adapter.TSRange(ctx, "t1", 0, 20).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 1, Value: 1}))

			opt = &TSRangeOptions{Latest: true}
			resultRange, err = adapter.TSRangeWithArgs(ctx, "t2", 0, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 0, Value: 4}))
			// Test Bucket Timestamp
			resultCreate, err = adapter.TSCreate(ctx, "t3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			_, errAdd = adapter.TSAdd(ctx, "t3", 15, 1).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 17, 4).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 51, 3).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 73, 5).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 75, 3).Result()
			Expect(errAdd).NotTo(HaveOccurred())

			opt = &TSRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10}
			resultRange, err = adapter.TSRangeWithArgs(ctx, "t3", 0, 100, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 10, Value: 4}))
			Expect(len(resultRange)).To(BeEquivalentTo(3))

			opt = &TSRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10, BucketTimestamp: "+"}
			resultRange, err = adapter.TSRangeWithArgs(ctx, "t3", 0, 100, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 20, Value: 4}))
			Expect(len(resultRange)).To(BeEquivalentTo(3))
			// Test Empty
			_, errAdd = adapter.TSAdd(ctx, "t4", 15, 1).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 17, 4).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 51, 3).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 73, 5).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 75, 3).Result()
			Expect(errAdd).NotTo(HaveOccurred())

			opt = &TSRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10}
			resultRange, err = adapter.TSRangeWithArgs(ctx, "t4", 0, 100, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 10, Value: 4}))
			Expect(len(resultRange)).To(BeEquivalentTo(3))

			opt = &TSRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10, Empty: true}
			resultRange, err = adapter.TSRangeWithArgs(ctx, "t4", 0, 100, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 10, Value: 4}))
			Expect(len(resultRange)).To(BeEquivalentTo(7))
		})

		It("should TSRevRange, TSRevRangeWithArgs", Label("timeseries", "tsrevrange", "tsrevrangeWithArgs"), func() {
			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i, float64(i%7)).Result()
				Expect(err).NotTo(HaveOccurred())

			}
			result, err := adapter.TSRange(ctx, "a", 0, 200).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(100))
			for i := 0; i < 100; i++ {
				adapter.TSAdd(ctx, "a", i+200, float64(i%7))
			}
			result, err = adapter.TSRange(ctx, "a", 0, 500).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(200))

			opt := &TSRevRangeOptions{Aggregator: Avg, BucketDuration: 10}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 500, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(20))

			opt = &TSRevRangeOptions{Count: 10}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 500, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(10))

			fts := make([]int, 0)
			for i := 10; i < 20; i++ {
				fts = append(fts, i)
			}
			opt = &TSRevRangeOptions{FilterByTS: fts, FilterByValue: []int{1, 2}}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 500, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(2))

			opt = &TSRevRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "+"}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo([]TSTimestampValue{{Timestamp: 10, Value: 1}, {Timestamp: 0, Value: 10}}))

			opt = &TSRevRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "1"}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo([]TSTimestampValue{{Timestamp: 1, Value: 10}, {Timestamp: 0, Value: 1}}))

			opt = &TSRevRangeOptions{Aggregator: Twa, BucketDuration: 10}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeEquivalentTo([]TSTimestampValue{{Timestamp: 10, Value: 3}, {Timestamp: 0, Value: 2.55}}))
			// Test Range Latest
			resultCreate, err := adapter.TSCreate(ctx, "t1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			resultCreate, err = adapter.TSCreate(ctx, "t2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			resultRule, err := adapter.TSCreateRule(ctx, "t1", "t2", Sum, 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRule).To(BeEquivalentTo("OK"))
			_, errAdd := adapter.TSAdd(ctx, "t1", 1, 1).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 2, 3).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 11, 7).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 13, 1).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			resultRange, err := adapter.TSRange(ctx, "t2", 0, 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 0, Value: 4}))
			opt = &TSRevRangeOptions{Latest: true}
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t2", 0, 10, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 10, Value: 8}))
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t2", 0, 9, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 0, Value: 4}))
			// Test Bucket Timestamp
			resultCreate, err = adapter.TSCreate(ctx, "t3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			_, errAdd = adapter.TSAdd(ctx, "t3", 15, 1).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 17, 4).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 51, 3).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 73, 5).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 75, 3).Result()
			Expect(errAdd).NotTo(HaveOccurred())

			opt = &TSRevRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10}
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t3", 0, 100, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 70, Value: 5}))
			Expect(len(resultRange)).To(BeEquivalentTo(3))

			opt = &TSRevRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10, BucketTimestamp: "+"}
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t3", 0, 100, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 80, Value: 5}))
			Expect(len(resultRange)).To(BeEquivalentTo(3))
			// Test Empty
			_, errAdd = adapter.TSAdd(ctx, "t4", 15, 1).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 17, 4).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 51, 3).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 73, 5).Result()
			Expect(errAdd).NotTo(HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 75, 3).Result()
			Expect(errAdd).NotTo(HaveOccurred())

			opt = &TSRevRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10}
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t4", 0, 100, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 70, Value: 5}))
			Expect(len(resultRange)).To(BeEquivalentTo(3))

			opt = &TSRevRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10, Empty: true}
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t4", 0, 100, opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultRange[0]).To(BeEquivalentTo(TSTimestampValue{Timestamp: 70, Value: 5}))
			Expect(len(resultRange)).To(BeEquivalentTo(7))
		})

		It("should TSMRange and TSMRangeWithArgs", Label("timeseries", "tsmrange", "tsmrangeWithArgs"), func() {
			createOpt := &TSOptions{Labels: map[string]string{"Test": "This", "team": "ny"}}
			resultCreate, err := adapter.TSCreateWithArgs(ctx, "a", createOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			createOpt = &TSOptions{Labels: map[string]string{"Test": "This", "Taste": "That", "team": "sf"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", createOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))

			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i, float64(i%7)).Result()
				Expect(err).NotTo(HaveOccurred())
				_, err = adapter.TSAdd(ctx, "b", i, float64(i%11)).Result()
				Expect(err).NotTo(HaveOccurred())
			}

			result, err := adapter.TSMRange(ctx, 0, 200, []string{"Test=This"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(2))
			Expect(len(result["a"][2].([]interface{}))).To(BeEquivalentTo(100))
			// Test Count
			mrangeOpt := &TSMRangeOptions{Count: 10}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result["a"][2].([]interface{}))).To(BeEquivalentTo(10))
			// Test Aggregation and BucketDuration
			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i+200, float64(i%7)).Result()
				Expect(err).NotTo(HaveOccurred())
			}
			mrangeOpt = &TSMRangeOptions{Aggregator: Avg, BucketDuration: 10}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 500, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(2))
			Expect(len(result["a"][2].([]interface{}))).To(BeEquivalentTo(20))
			// Test WithLabels
			Expect(result["a"][0]).To(BeEquivalentTo(map[string]interface{}{}))
			mrangeOpt = &TSMRangeOptions{WithLabels: true}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][0]).To(BeEquivalentTo(map[string]interface{}{"Test": "This", "team": "ny"}))
			// Test SelectedLabels
			mrangeOpt = &TSMRangeOptions{SelectedLabels: []interface{}{"team"}}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][0]).To(BeEquivalentTo(map[string]interface{}{"team": "ny"}))
			Expect(result["b"][0]).To(BeEquivalentTo(map[string]interface{}{"team": "sf"}))
			// Test FilterBy
			fts := make([]int, 0)
			for i := 10; i < 20; i++ {
				fts = append(fts, i)
			}
			mrangeOpt = &TSMRangeOptions{FilterByTS: fts, FilterByValue: []int{1, 2}}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][2]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(15), 1.0}, []interface{}{int64(16), 2.0}}))
			// Test GroupBy
			mrangeOpt = &TSMRangeOptions{GroupByLabel: "Test", Reducer: "sum"}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["Test=This"][3]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(0), 0.0}, []interface{}{int64(1), 2.0}, []interface{}{int64(2), 4.0}, []interface{}{int64(3), 6.0}}))

			mrangeOpt = &TSMRangeOptions{GroupByLabel: "Test", Reducer: "max"}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["Test=This"][3]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(0), 0.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(3), 3.0}}))

			mrangeOpt = &TSMRangeOptions{GroupByLabel: "team", Reducer: "min"}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(2))
			Expect(result["team=ny"][3]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(0), 0.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(3), 3.0}}))
			Expect(result["team=sf"][3]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(0), 0.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(3), 3.0}}))
			// Test Align
			mrangeOpt = &TSMRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "-"}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 10, []string{"team=ny"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][2]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(0), 10.0}, []interface{}{int64(10), 1.0}}))

			mrangeOpt = &TSMRangeOptions{Aggregator: Count, BucketDuration: 10, Align: 5}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 10, []string{"team=ny"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][2]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(0), 5.0}, []interface{}{int64(5), 6.0}}))
		})

		It("should TSMRangeWithArgs Latest", Label("timeseries", "tsmrangeWithArgs", "tsmrangelatest"), func() {
			resultCreate, err := adapter.TSCreate(ctx, "a").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			opt := &TSOptions{Labels: map[string]string{"is_compaction": "true"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))

			resultCreate, err = adapter.TSCreate(ctx, "c").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"is_compaction": "true"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "d", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))

			resultCreateRule, err := adapter.TSCreateRule(ctx, "a", "b", Sum, 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreateRule).To(BeEquivalentTo("OK"))
			resultCreateRule, err = adapter.TSCreateRule(ctx, "c", "d", Sum, 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreateRule).To(BeEquivalentTo("OK"))

			_, err = adapter.TSAdd(ctx, "a", 1, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 2, 3).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 11, 7).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 13, 1).Result()
			Expect(err).NotTo(HaveOccurred())

			_, err = adapter.TSAdd(ctx, "c", 1, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 2, 3).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 11, 7).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 13, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			mrangeOpt := &TSMRangeOptions{Latest: true}
			result, err := adapter.TSMRangeWithArgs(ctx, 0, 10, []string{"is_compaction=true"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["b"][2]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(0), 4.0}, []interface{}{int64(10), 8.0}}))
			Expect(result["d"][2]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(0), 4.0}, []interface{}{int64(10), 8.0}}))
		})
		It("should TSMRevRange and TSMRevRangeWithArgs", Label("timeseries", "tsmrevrange", "tsmrevrangeWithArgs"), func() {
			createOpt := &TSOptions{Labels: map[string]string{"Test": "This", "team": "ny"}}
			resultCreate, err := adapter.TSCreateWithArgs(ctx, "a", createOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			createOpt = &TSOptions{Labels: map[string]string{"Test": "This", "Taste": "That", "team": "sf"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", createOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))

			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i, float64(i%7)).Result()
				Expect(err).NotTo(HaveOccurred())
				_, err = adapter.TSAdd(ctx, "b", i, float64(i%11)).Result()
				Expect(err).NotTo(HaveOccurred())
			}
			result, err := adapter.TSMRevRange(ctx, 0, 200, []string{"Test=This"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(2))
			Expect(len(result["a"][2].([]interface{}))).To(BeEquivalentTo(100))
			// Test Count
			mrangeOpt := &TSMRevRangeOptions{Count: 10}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result["a"][2].([]interface{}))).To(BeEquivalentTo(10))
			// Test Aggregation and BucketDuration
			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i+200, float64(i%7)).Result()
				Expect(err).NotTo(HaveOccurred())
			}
			mrangeOpt = &TSMRevRangeOptions{Aggregator: Avg, BucketDuration: 10}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 500, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(2))
			Expect(len(result["a"][2].([]interface{}))).To(BeEquivalentTo(20))
			Expect(result["a"][0]).To(BeEquivalentTo(map[string]interface{}{}))
			// Test WithLabels
			Expect(result["a"][0]).To(BeEquivalentTo(map[string]interface{}{}))
			mrangeOpt = &TSMRevRangeOptions{WithLabels: true}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][0]).To(BeEquivalentTo(map[string]interface{}{"Test": "This", "team": "ny"}))
			// Test SelectedLabels
			mrangeOpt = &TSMRevRangeOptions{SelectedLabels: []interface{}{"team"}}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][0]).To(BeEquivalentTo(map[string]interface{}{"team": "ny"}))
			Expect(result["b"][0]).To(BeEquivalentTo(map[string]interface{}{"team": "sf"}))
			// Test FilterBy
			fts := make([]int, 0)
			for i := 10; i < 20; i++ {
				fts = append(fts, i)
			}
			mrangeOpt = &TSMRevRangeOptions{FilterByTS: fts, FilterByValue: []int{1, 2}}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][2]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(16), 2.0}, []interface{}{int64(15), 1.0}}))
			// Test GroupBy
			mrangeOpt = &TSMRevRangeOptions{GroupByLabel: "Test", Reducer: "sum"}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["Test=This"][3]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(3), 6.0}, []interface{}{int64(2), 4.0}, []interface{}{int64(1), 2.0}, []interface{}{int64(0), 0.0}}))

			mrangeOpt = &TSMRevRangeOptions{GroupByLabel: "Test", Reducer: "max"}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["Test=This"][3]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(3), 3.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(0), 0.0}}))

			mrangeOpt = &TSMRevRangeOptions{GroupByLabel: "team", Reducer: "min"}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(len(result)).To(BeEquivalentTo(2))
			Expect(result["team=ny"][3]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(3), 3.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(0), 0.0}}))
			Expect(result["team=sf"][3]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(3), 3.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(0), 0.0}}))
			// Test Align
			mrangeOpt = &TSMRevRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "-"}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 10, []string{"team=ny"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][2]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(10), 1.0}, []interface{}{int64(0), 10.0}}))

			mrangeOpt = &TSMRevRangeOptions{Aggregator: Count, BucketDuration: 10, Align: 1}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 10, []string{"team=ny"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["a"][2]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(1), 10.0}, []interface{}{int64(0), 1.0}}))
		})

		It("should TSMRevRangeWithArgs Latest", Label("timeseries", "tsmrevrangeWithArgs", "tsmrevrangelatest"), func() {
			resultCreate, err := adapter.TSCreate(ctx, "a").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			opt := &TSOptions{Labels: map[string]string{"is_compaction": "true"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))

			resultCreate, err = adapter.TSCreate(ctx, "c").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"is_compaction": "true"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "d", opt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreate).To(BeEquivalentTo("OK"))

			resultCreateRule, err := adapter.TSCreateRule(ctx, "a", "b", Sum, 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreateRule).To(BeEquivalentTo("OK"))
			resultCreateRule, err = adapter.TSCreateRule(ctx, "c", "d", Sum, 10).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultCreateRule).To(BeEquivalentTo("OK"))

			_, err = adapter.TSAdd(ctx, "a", 1, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 2, 3).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 11, 7).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 13, 1).Result()
			Expect(err).NotTo(HaveOccurred())

			_, err = adapter.TSAdd(ctx, "c", 1, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 2, 3).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 11, 7).Result()
			Expect(err).NotTo(HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 13, 1).Result()
			Expect(err).NotTo(HaveOccurred())
			mrangeOpt := &TSMRevRangeOptions{Latest: true}
			result, err := adapter.TSMRevRangeWithArgs(ctx, 0, 10, []string{"is_compaction=true"}, mrangeOpt).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(result["b"][2]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(10), 8.0}, []interface{}{int64(0), 4.0}}))
			Expect(result["d"][2]).To(BeEquivalentTo([]interface{}{[]interface{}{int64(10), 8.0}, []interface{}{int64(0), 4.0}}))
		})
	})
	Describe("JSON Commands", Label("json"), func() {
		BeforeEach(func() {
			Expect(adapter.FlushDB(ctx).Err()).NotTo(HaveOccurred())
		})

		Describe("arrays", Label("arrays"), func() {
			It("should JSONArrAppend", Label("json.arrappend", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "append2", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONArrAppend(ctx, "append2", "$..a", 10)
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(Equal([]int64{2, 3}))
			})

			It("should JSONArrIndex and JSONArrIndexWithArgs", Label("json.arrindex", "json"), func() {
				cmd1, err := adapter.JSONSet(ctx, "index1", "$", `{"a": [10], "b": {"a": [12, 10]}}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd1).To(Equal("OK"))

				cmd2, err := adapter.JSONArrIndex(ctx, "index1", "$.b.a", 10).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd2).To(Equal([]int64{1}))

				cmd3, err := adapter.JSONSet(ctx, "index2", "$", `[0,1,2,3,4]`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd3).To(Equal("OK"))

				res, err := adapter.JSONArrIndex(ctx, "index2", "$", 1).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res[0]).To(Equal(int64(1)))

				res, err = adapter.JSONArrIndex(ctx, "index2", "$", 1, 2).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res[0]).To(Equal(int64(-1)))

				res, err = adapter.JSONArrIndex(ctx, "index2", "$", 4).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res[0]).To(Equal(int64(4)))

				res, err = adapter.JSONArrIndexWithArgs(ctx, "index2", "$", &JSONArrIndexArgs{}, 4).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res[0]).To(Equal(int64(4)))

				stop := 5000
				res, err = adapter.JSONArrIndexWithArgs(ctx, "index2", "$", &JSONArrIndexArgs{Stop: &stop}, 4).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res[0]).To(Equal(int64(4)))

				stop = -1
				res, err = adapter.JSONArrIndexWithArgs(ctx, "index2", "$", &JSONArrIndexArgs{Stop: &stop}, 4).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res[0]).To(Equal(int64(-1)))

			})

			// FIXME: how to deal with expanded ?
			It("should JSONArrIndex and JSONArrIndexWithArgs with $", Label("json.arrindex", "json"), func() {
				doc := `{
					"store": {
						"book": [
							{
								"category": "reference",
								"author": "Nigel Rees",
								"title": "Sayings of the Century",
								"price": 8.95,
								"size": [10, 20, 30, 40]
							},
							{
								"category": "fiction",
								"author": "Evelyn Waugh",
								"title": "Sword of Honour",
								"price": 12.99,
								"size": [50, 60, 70, 80]
							},
							{
								"category": "fiction",
								"author": "Herman Melville",
								"title": "Moby Dick",
								"isbn": "0-553-21311-3",
								"price": 8.99,
								"size": [5, 10, 20, 30]
							},
							{
								"category": "fiction",
								"author": "J. R. R. Tolkien",
								"title": "The Lord of the Rings",
								"isbn": "0-395-19395-8",
								"price": 22.99,
								"size": [5, 6, 7, 8]
							}
						],
						"bicycle": {"color": "red", "price": 19.95}
					}
				}`
				res, err := adapter.JSONSet(ctx, "doc1", "$", doc).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				resGet, err := adapter.JSONGet(ctx, "doc1", "$.store.book[?(@.price<10)].size").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGet).To(Equal("[[10,20,30,40],[5,10,20,30]]"))

				resArr, err := adapter.JSONArrIndex(ctx, "doc1", "$.store.book[?(@.price<10)].size", 20).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resArr).To(Equal([]int64{1, 2}))
			})

			It("should JSONArrInsert", Label("json.arrinsert", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "insert2", "$", `[100, 200, 300, 200]`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONArrInsert(ctx, "insert2", "$", -1, 1, 2)
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(Equal([]int64{6}))

				cmd3 := adapter.JSONGet(ctx, "insert2")
				Expect(cmd3.Err()).NotTo(HaveOccurred())
				Expect(cmd3.Val()).To(Equal(`[100,200,300,1,2,200]`))
			})

			It("should JSONArrLen", Label("json.arrlen", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "length2", "$", `{"a": [10], "b": {"a": [12, 10, 20, 12, 90, 10]}}`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONArrLen(ctx, "length2", "$..a")
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(Equal([]int64{1, 6}))
			})

			It("should JSONArrPop", Label("json.arrpop"), func() {
				cmd1 := adapter.JSONSet(ctx, "pop4", "$", `[100, 200, 300, 200]`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONArrPop(ctx, "pop4", "$", 2)
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(Equal([]string{"300"}))

				cmd3 := adapter.JSONGet(ctx, "pop4", "$")
				Expect(cmd3.Err()).NotTo(HaveOccurred())
				Expect(cmd3.Val()).To(Equal("[[100,200,200]]"))
			})

			It("should JSONArrTrim", Label("json.arrtrim", "json"), func() {
				cmd1, err := adapter.JSONSet(ctx, "trim1", "$", `[0,1,2,3,4]`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd1).To(Equal("OK"))

				stop := 3
				cmd2, err := adapter.JSONArrTrimWithArgs(ctx, "trim1", "$", &JSONArrTrimArgs{Start: 1, Stop: &stop}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd2).To(Equal([]int64{3}))

				res, err := adapter.JSONGet(ctx, "trim1", "$").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(`[[1,2,3]]`))

				cmd3, err := adapter.JSONSet(ctx, "trim2", "$", `[0,1,2,3,4]`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd3).To(Equal("OK"))

				stop = 3
				cmd4, err := adapter.JSONArrTrimWithArgs(ctx, "trim2", "$", &JSONArrTrimArgs{Start: -1, Stop: &stop}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd4).To(Equal([]int64{0}))

				cmd5, err := adapter.JSONSet(ctx, "trim3", "$", `[0,1,2,3,4]`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd5).To(Equal("OK"))

				stop = 99
				cmd6, err := adapter.JSONArrTrimWithArgs(ctx, "trim3", "$", &JSONArrTrimArgs{Start: 3, Stop: &stop}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd6).To(Equal([]int64{2}))

				cmd7, err := adapter.JSONSet(ctx, "trim4", "$", `[0,1,2,3,4]`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd7).To(Equal("OK"))

				stop = 1
				cmd8, err := adapter.JSONArrTrimWithArgs(ctx, "trim4", "$", &JSONArrTrimArgs{Start: 9, Stop: &stop}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd8).To(Equal([]int64{0}))

				cmd9, err := adapter.JSONSet(ctx, "trim5", "$", `[0,1,2,3,4]`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd9).To(Equal("OK"))

				stop = 11
				cmd10, err := adapter.JSONArrTrimWithArgs(ctx, "trim5", "$", &JSONArrTrimArgs{Start: 9, Stop: &stop}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd10).To(Equal([]int64{0}))
			})

			It("should JSONArrPop", Label("json.arrpop", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "pop4", "$", `[100, 200, 300, 200]`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONArrPop(ctx, "pop4", "$", 2)
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(Equal([]string{"300"}))

				cmd3 := adapter.JSONGet(ctx, "pop4", "$")
				Expect(cmd3.Err()).NotTo(HaveOccurred())
				Expect(cmd3.Val()).To(Equal("[[100,200,200]]"))
			})

		})

		Describe("get/set", Label("getset"), func() {
			It("should JSONSet", Label("json.set", "json"), func() {
				cmd := adapter.JSONSet(ctx, "set1", "$", `{"a": 1, "b": 2, "hello": "world"}`)
				Expect(cmd.Err()).NotTo(HaveOccurred())
				Expect(cmd.Val()).To(Equal("OK"))
			})

			It("should JSONGet", Label("json.get", "json"), func() {
				res, err := adapter.JSONSet(ctx, "get3", "$", `{"a": 1, "b": 2}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				res, err = adapter.JSONGetWithArgs(ctx, "get3", &JSONGetArgs{Indent: "-"}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(`{-"a":1,-"b":2}`))

				res, err = adapter.JSONGetWithArgs(ctx, "get3", &JSONGetArgs{Indent: "-", Newline: `~`, Space: `!`}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(`{~-"a":!1,~-"b":!2~}`))
			})

			It("should JSONMerge", Label("json.merge", "json"), func() {
				res, err := adapter.JSONSet(ctx, "merge1", "$", `{"a": 1, "b": 2}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				res, err = adapter.JSONMerge(ctx, "merge1", "$", `{"b": 3, "c": 4}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				res, err = adapter.JSONGet(ctx, "merge1", "$").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(`[{"a":1,"b":3,"c":4}]`))
			})

			It("should JSONSetMode with NX", Label("json.setmode", "json"), func() {
				cmd := adapter.JSONSetMode(ctx, "setmode1", "$", `{"a": 1, "b": 2, "hello": "world"}`, "NX")
				Expect(cmd.Err()).NotTo(HaveOccurred())
				Expect(cmd.Val()).To(Equal("OK"))
			})

			It("should panic with invalid mode", Label("json.setmode", "json"), func() {
				Expect(func() {
					adapter.JSONSetMode(ctx, "setmode3", "$", `{"a": 1, "b": 2, "hello": "world"}`, "INVALID")
				}).To(Panic())
			})

			It("should JSONMSet", Label("json.mset", "json"), func() {
				doc1 := JSONSetArgs{Key: "mset1", Path: "$", Value: `{"a": 1}`}
				doc2 := JSONSetArgs{Key: "mset2", Path: "$", Value: 2}
				docs := []JSONSetArgs{doc1, doc2}

				mSetResult, err := adapter.JSONMSetArgs(ctx, docs).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(mSetResult).To(Equal("OK"))

				res, err := adapter.JSONMGet(ctx, "$", "mset1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]interface{}{`[{"a":1}]`}))

				res, err = adapter.JSONMGet(ctx, "$", "mset1", "mset2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]interface{}{`[{"a":1}]`, "[2]"}))

				mSetResult, err = adapter.JSONMSet(ctx, "mset1", "$.a", 2, "mset3", "$", `[1]`).Result()
				Expect(err).NotTo(HaveOccurred())
			})

			It("should JSONMGet", Label("json.mget", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "mget2a", "$", `{"a": ["aa", "ab", "ac", "ad"], "b": {"a": ["ba", "bb", "bc", "bd"]}}`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))
				cmd2 := adapter.JSONSet(ctx, "mget2b", "$", `{"a": [100, 200, 300, 200], "b": {"a": [100, 200, 300, 200]}}`)
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(Equal("OK"))

				cmd3 := adapter.JSONMGet(ctx, "$..a", "mget2a", "mget2b")
				Expect(cmd3.Err()).NotTo(HaveOccurred())
				Expect(cmd3.Val()).To(HaveLen(2))
				Expect(cmd3.Val()[0]).To(Equal(`[["aa","ab","ac","ad"],["ba","bb","bc","bd"]]`))
				Expect(cmd3.Val()[1]).To(Equal(`[[100,200,300,200],[100,200,300,200]]`))
			})

			It("should JSONMget with $", Label("json.mget", "json"), func() {
				res, err := adapter.JSONSet(ctx, "doc1", "$", `{"a": 1, "b": 2, "nested": {"a": 3}, "c": "", "nested2": {"a": ""}}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				res, err = adapter.JSONSet(ctx, "doc2", "$", `{"a": 4, "b": 5, "nested": {"a": 6}, "c": "", "nested2": {"a": [""]}}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				iRes, err := adapter.JSONMGet(ctx, "$..a", "doc1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal([]interface{}{`[1,3,""]`}))

				iRes, err = adapter.JSONMGet(ctx, "$..a", "doc1", "doc2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal([]interface{}{`[1,3,""]`, `[4,6,[""]]`}))

				iRes, err = adapter.JSONMGet(ctx, "$..a", "non_existing_doc", "non_existing_doc1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal([]interface{}{nil, nil}))
			})
		})

		Describe("Misc", Label("misc"), func() {

			It("should JSONClear", Label("json.clear", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "clear1", "$", `[1]`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONClear(ctx, "clear1", "$")
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(Equal(int64(1)))

				cmd3 := adapter.JSONGet(ctx, "clear1", "$")
				Expect(cmd3.Err()).NotTo(HaveOccurred())
				Expect(cmd3.Val()).To(Equal(`[[]]`))
			})

			It("should JSONClear with $", Label("json.clear", "json"), func() {
				doc := `{
					"nested1": {"a": {"foo": 10, "bar": 20}},
					"a": ["foo"],
					"nested2": {"a": "claro"},
					"nested3": {"a": {"baz": 50}}
				}`
				res, err := adapter.JSONSet(ctx, "doc1", "$", doc).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				iRes, err := adapter.JSONClear(ctx, "doc1", "$..a").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal(int64(3)))

				resGet, err := adapter.JSONGet(ctx, "doc1", `$`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGet).To(Equal(`[{"nested1":{"a":{}},"a":[],"nested2":{"a":"claro"},"nested3":{"a":{}}}]`))

				res, err = adapter.JSONSet(ctx, "doc1", "$", doc).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				iRes, err = adapter.JSONClear(ctx, "doc1", "$.nested1.a").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal(int64(1)))

				resGet, err = adapter.JSONGet(ctx, "doc1", `$`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGet).To(Equal(`[{"nested1":{"a":{}},"a":["foo"],"nested2":{"a":"claro"},"nested3":{"a":{"baz":50}}}]`))
			})

			It("should JSONDel", Label("json.del", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "del1", "$", `[1]`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONDel(ctx, "del1", "$")
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(Equal(int64(1)))

				cmd3 := adapter.JSONGet(ctx, "del1", "$")
				// go-redis's test assertion is wrong.
				// based on the result from redis/redis-stack:7.2.0-v3,
				// cmd3.Err() should be rueidis.Nil, not nil
				Expect(cmd3.Err()).To(Equal(rueidis.Nil))
				Expect(cmd3.Val()).To(HaveLen(0))
			})

			It("should JSONDel with $", Label("json.del", "json"), func() {
				res, err := adapter.JSONSet(ctx, "del1", "$", `{"a": 1, "nested": {"a": 2, "b": 3}}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				iRes, err := adapter.JSONDel(ctx, "del1", "$..a").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal(int64(2)))

				resGet, err := adapter.JSONGet(ctx, "del1", "$").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGet).To(Equal(`[{"nested":{"b":3}}]`))

				res, err = adapter.JSONSet(ctx, "del2", "$", `{"a": {"a": 2, "b": 3}, "b": ["a", "b"], "nested": {"b": [true, "a", "b"]}}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				iRes, err = adapter.JSONDel(ctx, "del2", "$..a").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal(int64(1)))

				resGet, err = adapter.JSONGet(ctx, "del2", "$").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGet).To(Equal(`[{"nested":{"b":[true,"a","b"]},"b":["a","b"]}]`))

				doc := `[
					{
						"ciao": ["non ancora"],
						"nested": [
							{"ciao": [1, "a"]},
							{"ciao": [2, "a"]},
							{"ciaoc": [3, "non", "ciao"]},
							{"ciao": [4, "a"]},
							{"e": [5, "non", "ciao"]}
						]
					}
				]`
				res, err = adapter.JSONSet(ctx, "del3", "$", doc).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				iRes, err = adapter.JSONDel(ctx, "del3", `$.[0]["nested"]..ciao`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal(int64(3)))

				resVal := `[[{"ciao":["non ancora"],"nested":[{},{},{"ciaoc":[3,"non","ciao"]},{},{"e":[5,"non","ciao"]}]}]]`
				resGet, err = adapter.JSONGet(ctx, "del3", "$").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGet).To(Equal(resVal))
			})

			It("should JSONForget", Label("json.forget", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "forget3", "$", `{"a": [1,2,3], "b": {"a": [1,2,3], "b": "annie"}}`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONForget(ctx, "forget3", "$..a")
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(Equal(int64(2)))

				cmd3 := adapter.JSONGet(ctx, "forget3", "$")
				Expect(cmd3.Err()).NotTo(HaveOccurred())
				Expect(cmd3.Val()).To(Equal(`[{"b":{"b":"annie"}}]`))

			})

			It("should JSONForget with $", Label("json.forget", "json"), func() {
				res, err := adapter.JSONSet(ctx, "doc1", "$", `{"a": 1, "nested": {"a": 2, "b": 3}}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				iRes, err := adapter.JSONForget(ctx, "doc1", "$..a").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal(int64(2)))

				resGet, err := adapter.JSONGet(ctx, "doc1", "$").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGet).To(Equal(`[{"nested":{"b":3}}]`))

				res, err = adapter.JSONSet(ctx, "doc2", "$", `{"a": {"a": 2, "b": 3}, "b": ["a", "b"], "nested": {"b": [true, "a", "b"]}}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				iRes, err = adapter.JSONForget(ctx, "doc2", "$..a").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal(int64(1)))

				resGet, err = adapter.JSONGet(ctx, "doc2", "$").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGet).To(Equal(`[{"nested":{"b":[true,"a","b"]},"b":["a","b"]}]`))

				doc := `[
					{
						"ciao": ["non ancora"],
						"nested": [
							{"ciao": [1, "a"]},
							{"ciao": [2, "a"]},
							{"ciaoc": [3, "non", "ciao"]},
							{"ciao": [4, "a"]},
							{"e": [5, "non", "ciao"]}
						]
					}
				]`
				res, err = adapter.JSONSet(ctx, "doc3", "$", doc).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				iRes, err = adapter.JSONForget(ctx, "doc3", `$.[0]["nested"]..ciao`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(iRes).To(Equal(int64(3)))

				resVal := `[[{"ciao":["non ancora"],"nested":[{},{},{"ciaoc":[3,"non","ciao"]},{},{"e":[5,"non","ciao"]}]}]]`
				resGet, err = adapter.JSONGet(ctx, "doc3", "$").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGet).To(Equal(resVal))
			})

			It("should JSONNumIncrBy", Label("json.numincrby", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "incr3", "$", `{"a": [1, 2], "b": {"a": [0, -1]}}`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONNumIncrBy(ctx, "incr3", "$..a[1]", float64(1))
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(Equal(`[3,0]`))

				cmd3 := adapter.JSONSet(ctx, "incr4", "$", `{"a": [1, 2], "b": {"a": [0, -1], "c": "z"}, "c": 2}`)
				Expect(cmd3.Err()).NotTo(HaveOccurred())
				Expect(cmd3.Val()).To(Equal("OK"))

				cmd4 := adapter.JSONNumIncrBy(ctx, "incr4", "$..c", float64(1))
				Expect(cmd4.Err()).NotTo(HaveOccurred())
				// for NaN field, it should be null
				Expect(cmd4.Val()).To(Equal(`[3,null]`))
			})

			It("should JSONNumIncrBy with $", Label("json.numincrby", "json"), func() {
				res, err := adapter.JSONSet(ctx, "doc1", "$", `{"a": "b", "b": [{"a": 2}, {"a": 5.0}, {"a": "c"}]}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				res, err = adapter.JSONNumIncrBy(ctx, "doc1", "$.b[1].a", 2).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(`[7]`))

				res, err = adapter.JSONNumIncrBy(ctx, "doc1", "$.b[1].a", 3.5).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(`[10.5]`))

				res, err = adapter.JSONSet(ctx, "doc2", "$", `{"a": "b", "b": [{"a": 2}, {"a": 5.0}, {"a": "c"}]}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				res, err = adapter.JSONNumIncrBy(ctx, "doc2", "$.b[0].a", 3).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(`[5]`))
			})

			It("should JSONObjKeys", Label("json.objkeys", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "objkeys1", "$", `{"a": [1, 2], "b": {"a": [0, -1]}}`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONObjKeys(ctx, "objkeys1", "$..*")
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(HaveLen(7))
				Expect(cmd2.Val()).To(Equal([]interface{}{nil, []interface{}{"a"}, nil, nil, nil, nil, nil}))
			})

			It("should JSONObjKeys with $", Label("json.objkeys", "json"), func() {
				doc := `{
					"nested1": {"a": {"foo": 10, "bar": 20}},
					"a": ["foo"],
					"nested2": {"a": {"baz": 50}}
				}`
				cmd1, err := adapter.JSONSet(ctx, "objkeys1", "$", doc).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd1).To(Equal("OK"))

				cmd2, err := adapter.JSONObjKeys(ctx, "objkeys1", "$.nested1.a").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd2).To(Equal([]interface{}{[]interface{}{"foo", "bar"}}))

				cmd2, err = adapter.JSONObjKeys(ctx, "objkeys1", ".*.a").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd2).To(Equal([]interface{}{"foo", "bar"}))

				cmd2, err = adapter.JSONObjKeys(ctx, "objkeys1", ".nested2.a").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd2).To(Equal([]interface{}{"baz"}))

				_, err = adapter.JSONObjKeys(ctx, "non_existing_doc", "..a").Result()
				Expect(err).To(HaveOccurred())
			})

			It("should JSONObjLen", Label("json.objlen", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "objlen2", "$", `{"a": [1, 2], "b": {"a": [0, -1]}}`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONObjLen(ctx, "objlen2", "$..*")
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(HaveLen(7))
				Expect(cmd2.Val()[0]).To(BeNil())
				Expect(*cmd2.Val()[1]).To(Equal(int64(1)))
			})

			It("should JSONStrLen", Label("json.strlen", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "strlen2", "$", `{"a": "alice", "b": "bob", "c": {"a": "alice", "b": "bob"}}`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONStrLen(ctx, "strlen2", "$..*")
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(HaveLen(5))
				var tmp int64 = 20
				Expect(cmd2.Val()[0]).To(BeAssignableToTypeOf(&tmp))
				Expect(*cmd2.Val()[0]).To(Equal(int64(5)))
				Expect(*cmd2.Val()[1]).To(Equal(int64(3)))
				Expect(cmd2.Val()[2]).To(BeNil())
				Expect(*cmd2.Val()[3]).To(Equal(int64(5)))
				Expect(*cmd2.Val()[4]).To(Equal(int64(3)))
			})

			It("should JSONStrAppend", Label("json.strappend", "json"), func() {
				cmd1, err := adapter.JSONSet(ctx, "strapp1", "$", `"foo"`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd1).To(Equal("OK"))
				cmd2, err := adapter.JSONStrAppend(ctx, "strapp1", "$", `"bar"`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(*cmd2[0]).To(Equal(int64(6)))
				cmd3, err := adapter.JSONGet(ctx, "strapp1", "$").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(cmd3).To(Equal(`["foobar"]`))

			})

			It("should JSONStrAppend and JSONStrLen with $", Label("json.strappend", "json.strlen", "json"), func() {
				res, err := adapter.JSONSet(ctx, "doc1", "$", `{"a": "foo", "nested1": {"a": "hello"}, "nested2": {"a": 31}}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				intArrayResult, err := adapter.JSONStrAppend(ctx, "doc1", "$.nested1.a", `"baz"`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(*intArrayResult[0]).To(Equal(int64(8)))

				res, err = adapter.JSONSet(ctx, "doc2", "$", `{"a": "foo", "nested1": {"a": "hello"}, "nested2": {"a": 31}}`).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal("OK"))

				intResult, err := adapter.JSONStrLen(ctx, "doc2", "$.nested1.a").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(*intResult[0]).To(Equal(int64(5)))
			})

			It("should JSONToggle", Label("json.toggle", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "toggle1", "$", `[true]`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONToggle(ctx, "toggle1", "$[0]")
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(HaveLen(1))
				Expect(*cmd2.Val()[0]).To(Equal(int64(0)))
			})

			It("should JSONType", Label("json.type", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "type1", "$", `[true]`)
				Expect(cmd1.Err()).NotTo(HaveOccurred())
				Expect(cmd1.Val()).To(Equal("OK"))

				cmd2 := adapter.JSONType(ctx, "type1", "$[0]")
				Expect(cmd2.Err()).NotTo(HaveOccurred())
				Expect(cmd2.Val()).To(HaveLen(1))
				// RESP2 v RESP3
				Expect(cmd2.Val()[0]).To(Or(Equal([]interface{}{"boolean"}), Equal("boolean")))
			})
		})
	})
}

func testAdapterSearchRESP3() {
	var adapter Cmdable
	var client rueidis.Client

	BeforeEach(func() {
		adapter = adapterresp3
		client = clientresp3
		Expect(adapter.FlushDB(ctx).Err()).NotTo(HaveOccurred())
		Expect(adapter.FlushAll(ctx).Err()).NotTo(HaveOccurred())
	})

	Describe("RediSearch commands Resp 3", Label("search"), func() {
		ctx := context.TODO()

		It("should handle FTAggregate with Unstable RESP3 Search Module and without stability", Label("search", "ftcreate", "ftaggregate"), func() {
			text1 := &FieldSchema{FieldName: "PrimaryKey", FieldType: SearchFieldTypeText, Sortable: true}
			num1 := &FieldSchema{FieldName: "CreatedDateTimeUTC", FieldType: SearchFieldTypeNumeric, Sortable: true}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1, num1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 3)

			adapter.HSet(ctx, "doc1", "PrimaryKey", "9::362330", "CreatedDateTimeUTC", "637387878524969984")
			adapter.HSet(ctx, "doc2", "PrimaryKey", "9::362329", "CreatedDateTimeUTC", "637387875859270016")

			options := &FTAggregateOptions{Apply: []FTAggregateApply{{Field: "@CreatedDateTimeUTC * 10", As: "CreatedDateTimeUTC"}}}
			res, err := adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).RawResult()
			// results := res.(map[interface{}]interface{})["results"].([]interface{})
			results := res.(map[string]interface{})["results"].([]interface{})
			// Expect(results[0].(map[interface{}]interface{})["extra_attributes"].(map[interface{}]interface{})["CreatedDateTimeUTC"]).
			Expect(results[0].(map[string]interface{})["extra_attributes"].(map[string]interface{})["CreatedDateTimeUTC"]).
				To(Or(BeEquivalentTo("6373878785249699840"), BeEquivalentTo("6373878758592700416")))
			// Expect(results[1].(map[interface{}]interface{})["extra_attributes"].(map[interface{}]interface{})["CreatedDateTimeUTC"]).
			Expect(results[1].(map[string]interface{})["extra_attributes"].(map[string]interface{})["CreatedDateTimeUTC"]).
				To(Or(BeEquivalentTo("6373878785249699840"), BeEquivalentTo("6373878758592700416")))

			rawVal := adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).RawVal()
			// rawValResults := rawVal.(map[interface{}]interface{})["results"].([]interface{})
			rawValResults := rawVal.(map[string]interface{})["results"].([]interface{})
			Expect(err).NotTo(HaveOccurred())
			Expect(rawValResults[0]).To(Or(BeEquivalentTo(results[0]), BeEquivalentTo(results[1])))
			Expect(rawValResults[1]).To(Or(BeEquivalentTo(results[0]), BeEquivalentTo(results[1])))

			// NOTE: rueidis can't support this behavior because we cannot know whether UnstableResp3 is enabled or not
			// Test with UnstableResp3 false
			// Expect(func() {
			// 	options = &FTAggregateOptions{Apply: []FTAggregateApply{{Field: "@CreatedDateTimeUTC * 10", As: "CreatedDateTimeUTC"}}}
			// 	rawRes, _ := adapter2resp3.FTAggregateWithArgs(ctx, "idx1", "*", options).RawResult()
			// 	rawVal = adapter2resp3.FTAggregateWithArgs(ctx, "idx1", "*", options).RawVal()
			// 	Expect(rawRes).To(BeNil())
			// 	Expect(rawVal).To(BeNil())
			// }).Should(Panic())
		})

		It("should handle FTInfo with Unstable RESP3 Search Module and without stability", Label("search", "ftcreate", "ftinfo"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{FieldName: "txt", FieldType: SearchFieldTypeText, Sortable: true, NoStem: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 3)

			resInfo, err := adapter.FTInfo(ctx, "idx1").RawResult()
			Expect(err).NotTo(HaveOccurred())
			// attributes := resInfo.(map[interface{}]interface{})["attributes"].([]interface{})
			attributes := resInfo.(map[string]interface{})["attributes"].([]interface{})
			// flags := attributes[0].(map[interface{}]interface{})["flags"].([]interface{})
			flags := attributes[0].(map[string]interface{})["flags"].([]interface{})
			Expect(flags).To(ConsistOf("SORTABLE", "NOSTEM"))

			valInfo, err := adapter.FTInfo(ctx, "idx1").RawResult()
			Expect(err).NotTo(HaveOccurred())
			// attributes = valInfo.(map[interface{}]interface{})["attributes"].([]interface{})
			attributes = valInfo.(map[string]interface{})["attributes"].([]interface{})
			// flags = attributes[0].(map[interface{}]interface{})["flags"].([]interface{})
			flags = attributes[0].(map[string]interface{})["flags"].([]interface{})
			Expect(flags).To(ConsistOf("SORTABLE", "NOSTEM"))

			// NOTE: rueidis can't support this behavior because we cannot know whether UnstableResp3 is enabled or not
			// Test with UnstableResp3 false
			// Expect(func() {
			// 	rawResInfo, _ := adapter2resp3.FTInfo(ctx, "idx1").RawResult()
			// 	rawValInfo := adapter2resp3.FTInfo(ctx, "idx1").RawVal()
			// 	Expect(rawResInfo).To(BeNil())
			// 	Expect(rawValInfo).To(BeNil())
			// }).Should(Panic())
		})

		It("should handle FTSpellCheck with Unstable RESP3 Search Module and without stability", Label("search", "ftcreate", "ftspellcheck"), func() {
			text1 := &FieldSchema{FieldName: "f1", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "f2", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 3)

			adapter.HSet(ctx, "doc1", "f1", "some valid content", "f2", "this is sample text")
			adapter.HSet(ctx, "doc2", "f1", "very important", "f2", "lorem ipsum")

			resSpellCheck, err := adapter.FTSpellCheck(ctx, "idx1", "impornant").RawResult()
			valSpellCheck := adapter.FTSpellCheck(ctx, "idx1", "impornant").RawVal()
			Expect(err).NotTo(HaveOccurred())
			Expect(valSpellCheck).To(BeEquivalentTo(resSpellCheck))
			// results := resSpellCheck.(map[interface{}]interface{})["results"].(map[interface{}]interface{})
			results := resSpellCheck.(map[string]interface{})["results"].(map[string]interface{})
			// Expect(results["impornant"].([]interface{})[0].(map[interface{}]interface{})["important"]).To(BeEquivalentTo(0.5))
			Expect(results["impornant"].([]interface{})[0].(map[string]interface{})["important"]).To(BeEquivalentTo(0.5))

			// NOTE: rueidis can't support this behavior because we cannot know whether UnstableResp3 is enabled or not
			// Test with UnstableResp3 false
			// Expect(func() {
			// 	rawResSpellCheck, _ := adapter2resp3.FTSpellCheck(ctx, "idx1", "impornant").RawResult()
			// 	rawValSpellCheck := adapter2resp3.FTSpellCheck(ctx, "idx1", "impornant").RawVal()
			// 	Expect(rawResSpellCheck).To(BeNil())
			// 	Expect(rawValSpellCheck).To(BeNil())
			// }).Should(Panic())
		})

		It("should handle FTSearch with Unstable RESP3 Search Module and without stability", Label("search", "ftcreate", "ftsearch"), func() {
			val, err := adapter.FTCreate(ctx, "txt", &FTCreateOptions{StopWords: []interface{}{"foo", "bar", "baz"}}, &FieldSchema{FieldName: "txt", FieldType: SearchFieldTypeText}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "txt", 3)
			adapter.HSet(ctx, "doc1", "txt", "foo baz")
			adapter.HSet(ctx, "doc2", "txt", "hello world")
			res1, err := adapter.FTSearchWithArgs(ctx, "txt", "foo bar", &FTSearchOptions{NoContent: true}).RawResult()
			val1 := adapter.FTSearchWithArgs(ctx, "txt", "foo bar", &FTSearchOptions{NoContent: true}).RawVal()
			Expect(err).NotTo(HaveOccurred())
			Expect(val1).To(BeEquivalentTo(res1))
			totalResults := res1.(map[string]interface{})["total_results"]
			// totalResults := res1.(map[interface{}]interface{})["total_results"]
			Expect(totalResults).To(BeEquivalentTo(int64(0)))
			res2, err := adapter.FTSearchWithArgs(ctx, "txt", "foo bar hello world", &FTSearchOptions{NoContent: true}).RawResult()
			Expect(err).NotTo(HaveOccurred())
			totalResults2 := res2.(map[string]interface{})["total_results"]
			// totalResults2 := res2.(map[interface{}]interface{})["total_results"]
			Expect(totalResults2).To(BeEquivalentTo(int64(1)))

			// NOTE: rueidis can't support this behavior because we cannot know whether UnstableResp3 is enabled or not
			// Test with UnstableResp3 false
			// Expect(func() {
			// 	rawRes2, _ := adapter2resp3.FTSearchWithArgs(ctx, "txt", "foo bar hello world", &FTSearchOptions{NoContent: true}).RawResult()
			// 	rawVal2 := adapter2resp3.FTSearchWithArgs(ctx, "txt", "foo bar hello world", &FTSearchOptions{NoContent: true}).RawVal()
			// 	Expect(rawRes2).To(BeNil())
			// 	Expect(rawVal2).To(BeNil())
			// }).Should(Panic())
		})
		It("should handle FTSynDump with Unstable RESP3 Search Module and without stability", Label("search", "ftsyndump"), func() {
			text1 := &FieldSchema{FieldName: "title", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "body", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{OnHash: true}, text1, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 3)

			resSynUpdate, err := adapter.FTSynUpdate(ctx, "idx1", "id1", []interface{}{"boy", "child", "offspring"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSynUpdate).To(BeEquivalentTo("OK"))

			resSynUpdate, err = adapter.FTSynUpdate(ctx, "idx1", "id1", []interface{}{"baby", "child"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSynUpdate).To(BeEquivalentTo("OK"))

			resSynUpdate, err = adapter.FTSynUpdate(ctx, "idx1", "id1", []interface{}{"tree", "wood"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSynUpdate).To(BeEquivalentTo("OK"))

			resSynDump, err := adapter.FTSynDump(ctx, "idx1").RawResult()
			valSynDump := adapter.FTSynDump(ctx, "idx1").RawVal()
			Expect(err).NotTo(HaveOccurred())
			Expect(valSynDump).To(BeEquivalentTo(resSynDump))
			// Expect(resSynDump.(map[interface{}]interface{})["baby"]).To(BeEquivalentTo([]interface{}{"id1"}))
			Expect(resSynDump.(map[string]interface{})["baby"]).To(BeEquivalentTo([]interface{}{"id1"}))

			// NOTE: rueidis can't support this behavior because we cannot know whether UnstableResp3 is enabled or not
			// Test with UnstableResp3 false
			// Expect(func() {
			// 	rawResSynDump, _ := adapter2resp3.FTSynDump(ctx, "idx1").RawResult()
			// 	rawValSynDump := adapter2resp3.FTSynDump(ctx, "idx1").RawVal()
			// 	Expect(rawResSynDump).To(BeNil())
			// 	Expect(rawValSynDump).To(BeNil())
			// }).Should(Panic())
		})

		It("should test not affected Resp 3 Search method - FTExplain", Label("search", "ftexplain"), func() {
			text1 := &FieldSchema{FieldName: "f1", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "f2", FieldType: SearchFieldTypeText}
			text3 := &FieldSchema{FieldName: "f3", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "txt", &FTCreateOptions{}, text1, text2, text3).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "txt", 3)
			res1, err := adapter.FTExplain(ctx, "txt", "@f3:f3_val @f2:f2_val @f1:f1_val").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1).ToNot(BeEmpty())

			// NOTE: rueidis can't support this behavior because we cannot know whether UnstableResp3 is enabled or not
			// Test with UnstableResp3 false
			// Expect(func() {
			// 	res2, err := adapter2resp3.FTExplain(ctx, "txt", "@f3:f3_val @f2:f2_val @f1:f1_val").Result()
			// 	Expect(err).NotTo(HaveOccurred())
			// 	Expect(res2).ToNot(BeEmpty())
			// }).ShouldNot(Panic())
		})
	})
}

func testAdapterSearchRESP2() {
	var adapter Cmdable
	var client rueidis.Client

	BeforeEach(func() {
		adapter = adaptersearchresp2
		client = clientsearchresp2

		Expect(adapter.FlushDB(ctx).Err()).NotTo(HaveOccurred())
		Expect(adapter.FlushAll(ctx).Err()).NotTo(HaveOccurred())
	})

	Describe("RediSearch commands Resp 2", Label("search"), func() {
		ctx := context.TODO()

		BeforeEach(func() {
			Expect(adapter.FlushDB(ctx).Err()).NotTo(HaveOccurred())
		})

		It("should FTCreate and FTSearch WithScores", Label("search", "ftcreate", "ftsearch"), func() {
			// FIXME: FieldType
			val, err := adapter.FTCreate(ctx, "txt", &FTCreateOptions{}, &FieldSchema{FieldName: "txt", FieldType: SearchFieldTypeText}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "txt", 2)
			adapter.HSet(ctx, "doc1", "txt", "foo baz")
			adapter.HSet(ctx, "doc2", "txt", "foo bar")
			res, err := adapter.FTSearchWithArgs(ctx, "txt", "foo ~bar", &FTSearchOptions{WithScores: true}).Result()

			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(int64(2)))
			for _, doc := range res.Docs {
				Expect(*doc.Score).To(BeNumerically(">", 0))
				Expect(doc.ID).To(Or(Equal("doc1"), Equal("doc2")))
			}
		})

		It("should FTCreate and FTSearch stopwords", Label("search", "ftcreate", "ftsearch"), func() {
			val, err := adapter.FTCreate(ctx, "txt", &FTCreateOptions{StopWords: []interface{}{"foo", "bar", "baz"}}, &FieldSchema{FieldName: "txt", FieldType: SearchFieldTypeText}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "txt", 2)
			adapter.HSet(ctx, "doc1", "txt", "foo baz")
			adapter.HSet(ctx, "doc2", "txt", "hello world")
			res1, err := adapter.FTSearchWithArgs(ctx, "txt", "foo bar", &FTSearchOptions{NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(0)))
			res2, err := adapter.FTSearchWithArgs(ctx, "txt", "foo bar hello world", &FTSearchOptions{NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res2.Total).To(BeEquivalentTo(int64(1)))
		})

		It("should FTCreate and FTSearch filters", Label("search", "ftcreate", "ftsearch"), func() {
			val, err := adapter.FTCreate(ctx, "txt", &FTCreateOptions{}, &FieldSchema{FieldName: "txt", FieldType: SearchFieldTypeText}, &FieldSchema{FieldName: "num", FieldType: SearchFieldTypeNumeric}, &FieldSchema{FieldName: "loc", FieldType: SearchFieldTypeGeo}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "txt", 2)
			adapter.HSet(ctx, "doc1", "txt", "foo bar", "num", 3.141, "loc", "-0.441,51.458")
			adapter.HSet(ctx, "doc2", "txt", "foo baz", "num", 2, "loc", "-0.1,51.2")
			res1, err := adapter.FTSearchWithArgs(ctx, "txt", "foo", &FTSearchOptions{Filters: []FTSearchFilter{{FieldName: "num", Min: 0, Max: 2}}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(1)))
			Expect(res1.Docs[0].ID).To(BeEquivalentTo("doc2"))
			res2, err := adapter.FTSearchWithArgs(ctx, "txt", "foo", &FTSearchOptions{Filters: []FTSearchFilter{{FieldName: "num", Min: 0, Max: "+inf"}}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res2.Total).To(BeEquivalentTo(int64(2)))
			Expect(res2.Docs[0].ID).To(BeEquivalentTo("doc1"))
			// Test Geo filter
			geoFilter1 := FTSearchGeoFilter{FieldName: "loc", Longitude: -0.44, Latitude: 51.45, Radius: 10, Unit: "km"}
			geoFilter2 := FTSearchGeoFilter{FieldName: "loc", Longitude: -0.44, Latitude: 51.45, Radius: 100, Unit: "km"}
			res3, err := adapter.FTSearchWithArgs(ctx, "txt", "foo", &FTSearchOptions{GeoFilter: []FTSearchGeoFilter{geoFilter1}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res3.Total).To(BeEquivalentTo(int64(1)))
			Expect(res3.Docs[0].ID).To(BeEquivalentTo("doc1"))
			res4, err := adapter.FTSearchWithArgs(ctx, "txt", "foo", &FTSearchOptions{GeoFilter: []FTSearchGeoFilter{geoFilter2}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res4.Total).To(BeEquivalentTo(int64(2)))
			docs := []interface{}{res4.Docs[0].ID, res4.Docs[1].ID}
			Expect(docs).To(ContainElement("doc1"))
			Expect(docs).To(ContainElement("doc2"))

		})

		It("should FTCreate and FTSearch sortby", Label("search", "ftcreate", "ftsearch"), func() {
			val, err := adapter.FTCreate(ctx, "num", &FTCreateOptions{}, &FieldSchema{FieldName: "txt", FieldType: SearchFieldTypeText}, &FieldSchema{FieldName: "num", FieldType: SearchFieldTypeNumeric, Sortable: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "num", 2)
			adapter.HSet(ctx, "doc1", "txt", "foo bar", "num", 1)
			adapter.HSet(ctx, "doc2", "txt", "foo baz", "num", 2)
			adapter.HSet(ctx, "doc3", "txt", "foo qux", "num", 3)

			sortBy1 := FTSearchSortBy{FieldName: "num", Asc: true}
			sortBy2 := FTSearchSortBy{FieldName: "num", Desc: true}
			res1, err := adapter.FTSearchWithArgs(ctx, "num", "foo", &FTSearchOptions{NoContent: true, SortBy: []FTSearchSortBy{sortBy1}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(3)))
			Expect(res1.Docs[0].ID).To(BeEquivalentTo("doc1"))
			Expect(res1.Docs[1].ID).To(BeEquivalentTo("doc2"))
			Expect(res1.Docs[2].ID).To(BeEquivalentTo("doc3"))

			res2, err := adapter.FTSearchWithArgs(ctx, "num", "foo", &FTSearchOptions{NoContent: true, SortBy: []FTSearchSortBy{sortBy2}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res2.Total).To(BeEquivalentTo(int64(3)))
			Expect(res2.Docs[2].ID).To(BeEquivalentTo("doc1"))
			Expect(res2.Docs[1].ID).To(BeEquivalentTo("doc2"))
			Expect(res2.Docs[0].ID).To(BeEquivalentTo("doc3"))

		})

		It("should FTCreate and FTSearch example", Label("search", "ftcreate", "ftsearch"), func() {
			val, err := adapter.FTCreate(ctx, "txt", &FTCreateOptions{}, &FieldSchema{FieldName: "title", FieldType: SearchFieldTypeText, Weight: 5}, &FieldSchema{FieldName: "body", FieldType: SearchFieldTypeText}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "txt", 2)
			adapter.HSet(ctx, "doc1", "title", "RediSearch", "body", "Redisearch implements a search engine on top of redis")
			res1, err := adapter.FTSearchWithArgs(ctx, "txt", "search engine", &FTSearchOptions{NoContent: true, Verbatim: true, LimitOffset: 0, Limit: 5}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(1)))

		})

		It("should FTCreate NoIndex", Label("search", "ftcreate", "ftsearch"), func() {
			text1 := &FieldSchema{FieldName: "field", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "text", FieldType: SearchFieldTypeText, NoIndex: true, Sortable: true}
			num := &FieldSchema{FieldName: "numeric", FieldType: SearchFieldTypeNumeric, NoIndex: true, Sortable: true}
			geo := &FieldSchema{FieldName: "geo", FieldType: SearchFieldTypeGeo, NoIndex: true, Sortable: true}
			tag := &FieldSchema{FieldName: "tag", FieldType: SearchFieldTypeTag, NoIndex: true, Sortable: true}
			val, err := adapter.FTCreate(ctx, "idx", &FTCreateOptions{}, text1, text2, num, geo, tag).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx", 2)
			adapter.HSet(ctx, "doc1", "field", "aaa", "text", "1", "numeric", 1, "geo", "1,1", "tag", "1")
			adapter.HSet(ctx, "doc2", "field", "aab", "text", "2", "numeric", 2, "geo", "2,2", "tag", "2")
			res1, err := adapter.FTSearch(ctx, "idx", "@text:aa*").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(0)))
			res2, err := adapter.FTSearch(ctx, "idx", "@field:aa*").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res2.Total).To(BeEquivalentTo(int64(2)))
			res3, err := adapter.FTSearchWithArgs(ctx, "idx", "*", &FTSearchOptions{SortBy: []FTSearchSortBy{{FieldName: "text", Desc: true}}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res3.Total).To(BeEquivalentTo(int64(2)))
			Expect(res3.Docs[0].ID).To(BeEquivalentTo("doc2"))
			res4, err := adapter.FTSearchWithArgs(ctx, "idx", "*", &FTSearchOptions{SortBy: []FTSearchSortBy{{FieldName: "text", Asc: true}}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res4.Total).To(BeEquivalentTo(int64(2)))
			Expect(res4.Docs[0].ID).To(BeEquivalentTo("doc1"))
			res5, err := adapter.FTSearchWithArgs(ctx, "idx", "*", &FTSearchOptions{SortBy: []FTSearchSortBy{{FieldName: "numeric", Asc: true}}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res5.Docs[0].ID).To(BeEquivalentTo("doc1"))
			res6, err := adapter.FTSearchWithArgs(ctx, "idx", "*", &FTSearchOptions{SortBy: []FTSearchSortBy{{FieldName: "geo", Asc: true}}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res6.Docs[0].ID).To(BeEquivalentTo("doc1"))
			res7, err := adapter.FTSearchWithArgs(ctx, "idx", "*", &FTSearchOptions{SortBy: []FTSearchSortBy{{FieldName: "tag", Asc: true}}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res7.Docs[0].ID).To(BeEquivalentTo("doc1"))

		})

		It("should FTExplain", Label("search", "ftexplain"), func() {
			text1 := &FieldSchema{FieldName: "f1", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "f2", FieldType: SearchFieldTypeText}
			text3 := &FieldSchema{FieldName: "f3", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "txt", &FTCreateOptions{}, text1, text2, text3).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "txt", 2)
			res1, err := adapter.FTExplain(ctx, "txt", "@f3:f3_val @f2:f2_val @f1:f1_val").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1).ToNot(BeEmpty())

		})

		It("should FTAlias", Label("search", "ftexplain"), func() {
			text1 := &FieldSchema{FieldName: "name", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "name", FieldType: SearchFieldTypeText}
			val1, err := adapter.FTCreate(ctx, "testAlias", &FTCreateOptions{Prefix: []interface{}{"index1:"}}, text1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val1).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "testAlias", 2)
			val2, err := adapter.FTCreate(ctx, "testAlias2", &FTCreateOptions{Prefix: []interface{}{"index2:"}}, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val2).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "testAlias2", 2)

			adapter.HSet(ctx, "index1:lonestar", "name", "lonestar")
			adapter.HSet(ctx, "index2:yogurt", "name", "yogurt")

			res1, err := adapter.FTSearch(ctx, "testAlias", "*").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Docs[0].ID).To(BeEquivalentTo("index1:lonestar"))

			aliasAddRes, err := adapter.FTAliasAdd(ctx, "testAlias", "mj23").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(aliasAddRes).To(BeEquivalentTo("OK"))

			res1, err = adapter.FTSearch(ctx, "mj23", "*").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Docs[0].ID).To(BeEquivalentTo("index1:lonestar"))

			aliasUpdateRes, err := adapter.FTAliasUpdate(ctx, "testAlias2", "kb24").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(aliasUpdateRes).To(BeEquivalentTo("OK"))

			res3, err := adapter.FTSearch(ctx, "kb24", "*").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res3.Docs[0].ID).To(BeEquivalentTo("index2:yogurt"))

			aliasDelRes, err := adapter.FTAliasDel(ctx, "mj23").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(aliasDelRes).To(BeEquivalentTo("OK"))

		})

		It("should FTCreate and FTSearch textfield, sortable and nostem ", Label("search", "ftcreate", "ftsearch"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{FieldName: "txt", FieldType: SearchFieldTypeText, Sortable: true, NoStem: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			resInfo, err := adapter.FTInfo(ctx, "idx1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resInfo.Attributes[0].Sortable).To(BeTrue())
			Expect(resInfo.Attributes[0].NoStem).To(BeTrue())

		})

		It("should FTAlter", Label("search", "ftcreate", "ftsearch", "ftalter"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{FieldName: "txt", FieldType: SearchFieldTypeText}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			resAlter, err := adapter.FTAlter(ctx, "idx1", false, []interface{}{"body", SearchFieldTypeText.String()}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resAlter).To(BeEquivalentTo("OK"))

			adapter.HSet(ctx, "doc1", "title", "MyTitle", "body", "Some content only in the body")
			res1, err := adapter.FTSearch(ctx, "idx1", "only in the body").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(1)))

		})

		It("should FTSpellCheck", Label("search", "ftcreate", "ftsearch", "ftspellcheck"), func() {
			text1 := &FieldSchema{FieldName: "f1", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "f2", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "f1", "some valid content", "f2", "this is sample text")
			adapter.HSet(ctx, "doc2", "f1", "very important", "f2", "lorem ipsum")

			resSpellCheck, err := adapter.FTSpellCheck(ctx, "idx1", "impornant").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSpellCheck[0].Suggestions[0].Suggestion).To(BeEquivalentTo("important"))

			resSpellCheck2, err := adapter.FTSpellCheck(ctx, "idx1", "contnt").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSpellCheck2[0].Suggestions[0].Suggestion).To(BeEquivalentTo("content"))

			// test spellcheck with Levenshtein distance
			resSpellCheck3, err := adapter.FTSpellCheck(ctx, "idx1", "vlis").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSpellCheck3[0].Term).To(BeEquivalentTo("vlis"))

			resSpellCheck4, err := adapter.FTSpellCheckWithArgs(ctx, "idx1", "vlis", &FTSpellCheckOptions{Distance: 2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSpellCheck4[0].Suggestions[0].Suggestion).To(BeEquivalentTo("valid"))

			// test spellcheck include
			resDictAdd, err := adapter.FTDictAdd(ctx, "dict", "lore", "lorem", "lorm").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resDictAdd).To(BeEquivalentTo(3))
			terms := &FTSpellCheckTerms{Inclusion: "INCLUDE", Dictionary: "dict"}
			resSpellCheck5, err := adapter.FTSpellCheckWithArgs(ctx, "idx1", "lorm", &FTSpellCheckOptions{Terms: terms}).Result()
			Expect(err).NotTo(HaveOccurred())
			lorm := resSpellCheck5[0].Suggestions
			Expect(len(lorm)).To(BeEquivalentTo(3))
			Expect(lorm[0].Score).To(BeEquivalentTo(0.5))
			Expect(lorm[1].Score).To(BeEquivalentTo(0))
			Expect(lorm[2].Score).To(BeEquivalentTo(0))

			terms2 := &FTSpellCheckTerms{Inclusion: "EXCLUDE", Dictionary: "dict"}
			resSpellCheck6, err := adapter.FTSpellCheckWithArgs(ctx, "idx1", "lorm", &FTSpellCheckOptions{Terms: terms2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSpellCheck6).To(BeEmpty())
		})

		It("should FTDict operations", Label("search", "ftdictdump", "ftdictdel", "ftdictadd"), func() {
			text1 := &FieldSchema{FieldName: "f1", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "f2", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			resDictAdd, err := adapter.FTDictAdd(ctx, "custom_dict", "item1", "item2", "item3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resDictAdd).To(BeEquivalentTo(3))

			resDictDel, err := adapter.FTDictDel(ctx, "custom_dict", "item2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resDictDel).To(BeEquivalentTo(1))

			resDictDump, err := adapter.FTDictDump(ctx, "custom_dict").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resDictDump).To(BeEquivalentTo([]string{"item1", "item3"}))

			resDictDel2, err := adapter.FTDictDel(ctx, "custom_dict", "item1", "item3").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resDictDel2).To(BeEquivalentTo(2))
		})

		It("should FTSearch phonetic matcher", Label("search", "ftsearch"), func() {
			text1 := &FieldSchema{FieldName: "name", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "name", "Jon")
			adapter.HSet(ctx, "doc2", "name", "John")

			res1, err := adapter.FTSearch(ctx, "idx1", "Jon").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(1)))
			Expect(res1.Docs[0].Fields["name"]).To(BeEquivalentTo("Jon"))

			adapter.FlushDB(ctx)
			text2 := &FieldSchema{FieldName: "name", FieldType: SearchFieldTypeText, PhoneticMatcher: "dm:en"}
			val2, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val2).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "name", "Jon")
			adapter.HSet(ctx, "doc2", "name", "John")

			res2, err := adapter.FTSearch(ctx, "idx1", "Jon").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res2.Total).To(BeEquivalentTo(int64(2)))
			names := []interface{}{res2.Docs[0].Fields["name"], res2.Docs[1].Fields["name"]}
			Expect(names).To(ContainElement("Jon"))
			Expect(names).To(ContainElement("John"))
		})

		It("should FTSearch WithScores", Label("search", "ftsearch"), func() {
			text1 := &FieldSchema{FieldName: "description", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "description", "The quick brown fox jumps over the lazy dog")
			adapter.HSet(ctx, "doc2", "description", "Quick alice was beginning to get very tired of sitting by her quick sister on the bank, and of having nothing to do.")

			res, err := adapter.FTSearchWithArgs(ctx, "idx1", "quick", &FTSearchOptions{WithScores: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(*res.Docs[0].Score).To(BeEquivalentTo(float64(1)))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "quick", &FTSearchOptions{WithScores: true, Scorer: "TFIDF"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(*res.Docs[0].Score).To(BeEquivalentTo(float64(1)))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "quick", &FTSearchOptions{WithScores: true, Scorer: "TFIDF.DOCNORM"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(*res.Docs[0].Score).To(BeEquivalentTo(0.14285714285714285))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "quick", &FTSearchOptions{WithScores: true, Scorer: "BM25"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(*res.Docs[0].Score).To(BeNumerically("<=", 0.22471909420069797))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "quick", &FTSearchOptions{WithScores: true, Scorer: "DISMAX"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(*res.Docs[0].Score).To(BeEquivalentTo(float64(2)))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "quick", &FTSearchOptions{WithScores: true, Scorer: "DOCSCORE"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(*res.Docs[0].Score).To(BeEquivalentTo(float64(1)))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "quick", &FTSearchOptions{WithScores: true, Scorer: "HAMMING"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(*res.Docs[0].Score).To(BeEquivalentTo(float64(0)))
		})

		It("should FTConfigSet and FTConfigGet ", Label("search", "ftconfigget", "ftconfigset", "NonRedisEnterprise"), func() {
			val, err := adapter.FTConfigSet(ctx, "TIMEOUT", "100").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))

			res, err := adapter.FTConfigGet(ctx, "*").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res["TIMEOUT"]).To(BeEquivalentTo("100"))

			res, err = adapter.FTConfigGet(ctx, "TIMEOUT").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(BeEquivalentTo(map[string]interface{}{"TIMEOUT": "100"}))

		})

		It("should FTAggregate GroupBy ", Label("search", "ftaggregate"), func() {
			text1 := &FieldSchema{FieldName: "title", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "body", FieldType: SearchFieldTypeText}
			text3 := &FieldSchema{FieldName: "parent", FieldType: SearchFieldTypeText}
			num := &FieldSchema{FieldName: "random_num", FieldType: SearchFieldTypeNumeric}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1, text2, text3, num).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "search", "title", "RediSearch",
				"body", "Redisearch implements a search engine on top of redis",
				"parent", "redis",
				"random_num", 10)
			adapter.HSet(ctx, "ai", "title", "RedisAI",
				"body", "RedisAI executes Deep Learning/Machine Learning models and managing their data.",
				"parent", "redis",
				"random_num", 3)
			adapter.HSet(ctx, "json", "title", "RedisJson",
				"body", "RedisJSON implements ECMA-404 The JSON Data Interchange Standard as a native data type.",
				"parent", "redis",
				"random_num", 8)

			reducer := FTAggregateReducer{Reducer: SearchCount}
			options := &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err := adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["__generated_aliascount"]).To(BeEquivalentTo("3"))

			reducer = FTAggregateReducer{Reducer: SearchCountDistinct, Args: []interface{}{"@title"}}
			options = &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["__generated_aliascount_distincttitle"]).To(BeEquivalentTo("3"))

			reducer = FTAggregateReducer{Reducer: SearchSum, Args: []interface{}{"@random_num"}}
			options = &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["__generated_aliassumrandom_num"]).To(BeEquivalentTo("21"))

			reducer = FTAggregateReducer{Reducer: SearchMin, Args: []interface{}{"@random_num"}}
			options = &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["__generated_aliasminrandom_num"]).To(BeEquivalentTo("3"))

			reducer = FTAggregateReducer{Reducer: SearchMax, Args: []interface{}{"@random_num"}}
			options = &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["__generated_aliasmaxrandom_num"]).To(BeEquivalentTo("10"))

			reducer = FTAggregateReducer{Reducer: SearchAvg, Args: []interface{}{"@random_num"}}
			options = &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["__generated_aliasavgrandom_num"]).To(BeEquivalentTo("7"))

			reducer = FTAggregateReducer{Reducer: SearchStdDev, Args: []interface{}{"@random_num"}}
			options = &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["__generated_aliasstddevrandom_num"]).To(BeEquivalentTo("3.60555127546"))

			reducer = FTAggregateReducer{Reducer: SearchQuantile, Args: []interface{}{"@random_num", 0.5}}
			options = &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["__generated_aliasquantilerandom_num,0.5"]).To(BeEquivalentTo("8"))

			reducer = FTAggregateReducer{Reducer: SearchToList, Args: []interface{}{"@title"}}
			options = &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["__generated_aliastolisttitle"]).To(ContainElements("RediSearch", "RedisAI", "RedisJson"))

			reducer = FTAggregateReducer{Reducer: SearchFirstValue, Args: []interface{}{"@title"}, As: "first"}
			options = &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["first"]).To(Or(BeEquivalentTo("RediSearch"), BeEquivalentTo("RedisAI"), BeEquivalentTo("RedisJson")))

			reducer = FTAggregateReducer{Reducer: SearchRandomSample, Args: []interface{}{"@title", 2}, As: "random"}
			options = &FTAggregateOptions{GroupBy: []FTAggregateGroupBy{{Fields: []interface{}{"@parent"}, Reduce: []FTAggregateReducer{reducer}}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "redis", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["parent"]).To(BeEquivalentTo("redis"))
			Expect(res.Rows[0].Fields["random"]).To(Or(
				ContainElement("RediSearch"),
				ContainElement("RedisAI"),
				ContainElement("RedisJson"),
			))

		})

		It("should FTAggregate sort and limit", Label("search", "ftaggregate"), func() {
			text1 := &FieldSchema{FieldName: "t1", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "t2", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "t1", "a", "t2", "b")
			adapter.HSet(ctx, "doc2", "t1", "b", "t2", "a")

			options := &FTAggregateOptions{SortBy: []FTAggregateSortBy{{FieldName: "@t2", Asc: true}, {FieldName: "@t1", Desc: true}}}
			res, err := adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["t1"]).To(BeEquivalentTo("b"))
			Expect(res.Rows[1].Fields["t1"]).To(BeEquivalentTo("a"))
			Expect(res.Rows[0].Fields["t2"]).To(BeEquivalentTo("a"))
			Expect(res.Rows[1].Fields["t2"]).To(BeEquivalentTo("b"))

			options = &FTAggregateOptions{SortBy: []FTAggregateSortBy{{FieldName: "@t1"}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["t1"]).To(BeEquivalentTo("a"))
			Expect(res.Rows[1].Fields["t1"]).To(BeEquivalentTo("b"))

			options = &FTAggregateOptions{SortBy: []FTAggregateSortBy{{FieldName: "@t1"}}, SortByMax: 1}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["t1"]).To(BeEquivalentTo("a"))

			options = &FTAggregateOptions{SortBy: []FTAggregateSortBy{{FieldName: "@t1"}}, Limit: 1, LimitOffset: 1}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["t1"]).To(BeEquivalentTo("b"))

			options = &FTAggregateOptions{SortBy: []FTAggregateSortBy{{FieldName: "@t1"}}, Limit: 1, LimitOffset: 0}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(res)
			Expect(res.Rows[0].Fields["t1"]).To(BeEquivalentTo("a"))
		})

		It("should FTAggregate load ", Label("search", "ftaggregate"), func() {
			text1 := &FieldSchema{FieldName: "t1", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "t2", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "t1", "hello", "t2", "world")

			options := &FTAggregateOptions{Load: []FTAggregateLoad{{Field: "t1"}}}
			res, err := adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["t1"]).To(BeEquivalentTo("hello"))

			options = &FTAggregateOptions{Load: []FTAggregateLoad{{Field: "t2"}}}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["t2"]).To(BeEquivalentTo("world"))

			options = &FTAggregateOptions{LoadAll: true}
			res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["t1"]).To(BeEquivalentTo("hello"))
			Expect(res.Rows[0].Fields["t2"]).To(BeEquivalentTo("world"))
		})

		It("should FTAggregate apply", Label("search", "ftaggregate"), func() {
			text1 := &FieldSchema{FieldName: "PrimaryKey", FieldType: SearchFieldTypeText, Sortable: true}
			num1 := &FieldSchema{FieldName: "CreatedDateTimeUTC", FieldType: SearchFieldTypeNumeric, Sortable: true}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1, num1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "PrimaryKey", "9::362330", "CreatedDateTimeUTC", "637387878524969984")
			adapter.HSet(ctx, "doc2", "PrimaryKey", "9::362329", "CreatedDateTimeUTC", "637387875859270016")

			options := &FTAggregateOptions{Apply: []FTAggregateApply{{Field: "@CreatedDateTimeUTC * 10", As: "CreatedDateTimeUTC"}}}
			res, err := adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Rows[0].Fields["CreatedDateTimeUTC"]).To(Or(BeEquivalentTo("6373878785249699840"), BeEquivalentTo("6373878758592700416")))
			Expect(res.Rows[1].Fields["CreatedDateTimeUTC"]).To(Or(BeEquivalentTo("6373878785249699840"), BeEquivalentTo("6373878758592700416")))

		})

		It("should FTAggregate filter", Label("search", "ftaggregate"), func() {
			text1 := &FieldSchema{FieldName: "name", FieldType: SearchFieldTypeText, Sortable: true}
			num1 := &FieldSchema{FieldName: "age", FieldType: SearchFieldTypeNumeric, Sortable: true}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, text1, num1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "name", "bar", "age", "25")
			adapter.HSet(ctx, "doc2", "name", "foo", "age", "19")

			for _, dlc := range []int{1, 2} {
				options := &FTAggregateOptions{Filter: "@name=='foo' && @age < 20", DialectVersion: dlc}
				res, err := adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res.Total).To(Or(BeEquivalentTo(2), BeEquivalentTo(1)))
				Expect(res.Rows[0].Fields["name"]).To(BeEquivalentTo("foo"))

				options = &FTAggregateOptions{Filter: "@age > 15", DialectVersion: dlc, SortBy: []FTAggregateSortBy{{FieldName: "@age"}}}
				res, err = adapter.FTAggregateWithArgs(ctx, "idx1", "*", options).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res.Total).To(BeEquivalentTo(2))
				Expect(res.Rows[0].Fields["age"]).To(BeEquivalentTo("19"))
				Expect(res.Rows[1].Fields["age"]).To(BeEquivalentTo("25"))
			}

		})

		It("should FTSearch SkipInitialScan", Label("search", "ftsearch"), func() {
			adapter.HSet(ctx, "doc1", "foo", "bar")

			text1 := &FieldSchema{FieldName: "foo", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{SkipInitialScan: true}, text1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			res, err := adapter.FTSearch(ctx, "idx1", "@foo:bar").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(int64(0)))
		})

		It("should FTCreate json", Label("search", "ftcreate"), func() {

			text1 := &FieldSchema{FieldName: "$.name", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{OnJSON: true, Prefix: []interface{}{"king:"}}, text1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.JSONSet(ctx, "king:1", "$", `{"name": "henry"}`)
			adapter.JSONSet(ctx, "king:2", "$", `{"name": "james"}`)

			res, err := adapter.FTSearch(ctx, "idx1", "henry").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("king:1"))
			Expect(res.Docs[0].Fields["$"]).To(BeEquivalentTo(`{"name":"henry"}`))
		})

		It("should FTCreate json fields as names", Label("search", "ftcreate"), func() {

			text1 := &FieldSchema{FieldName: "$.name", FieldType: SearchFieldTypeText, As: "name"}
			num1 := &FieldSchema{FieldName: "$.age", FieldType: SearchFieldTypeNumeric, As: "just_a_number"}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{OnJSON: true}, text1, num1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.JSONSet(ctx, "doc:1", "$", `{"name": "Jon", "age": 25}`)

			res, err := adapter.FTSearchWithArgs(ctx, "idx1", "Jon", &FTSearchOptions{Return: []FTSearchReturn{{FieldName: "name"}, {FieldName: "just_a_number"}}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("doc:1"))
			Expect(res.Docs[0].Fields["name"]).To(BeEquivalentTo("Jon"))
			Expect(res.Docs[0].Fields["just_a_number"]).To(BeEquivalentTo("25"))
		})

		It("should FTCreate CaseSensitive", Label("search", "ftcreate"), func() {

			tag1 := &FieldSchema{FieldName: "t", FieldType: SearchFieldTypeTag, CaseSensitive: false}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, tag1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "1", "t", "HELLO")
			adapter.HSet(ctx, "2", "t", "hello")

			res, err := adapter.FTSearch(ctx, "idx1", "@t:{HELLO}").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(2))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("1"))
			Expect(res.Docs[1].ID).To(BeEquivalentTo("2"))

			resDrop, err := adapter.FTDropIndex(ctx, "idx1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resDrop).To(BeEquivalentTo("OK"))

			tag2 := &FieldSchema{FieldName: "t", FieldType: SearchFieldTypeTag, CaseSensitive: true}
			val, err = adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, tag2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			res, err = adapter.FTSearch(ctx, "idx1", "@t:{HELLO}").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("1"))

		})

		It("should FTSearch ReturnFields", Label("search", "ftsearch"), func() {
			resJson, err := adapter.JSONSet(ctx, "doc:1", "$", `{"t": "riceratops","t2": "telmatosaurus", "n": 9072, "flt": 97.2}`).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resJson).To(BeEquivalentTo("OK"))

			text1 := &FieldSchema{FieldName: "$.t", FieldType: SearchFieldTypeText}
			num1 := &FieldSchema{FieldName: "$.flt", FieldType: SearchFieldTypeNumeric}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{OnJSON: true}, text1, num1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			res, err := adapter.FTSearchWithArgs(ctx, "idx1", "*", &FTSearchOptions{Return: []FTSearchReturn{{FieldName: "$.t", As: "txt"}}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("doc:1"))
			Expect(res.Docs[0].Fields["txt"]).To(BeEquivalentTo("riceratops"))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "*", &FTSearchOptions{Return: []FTSearchReturn{{FieldName: "$.t2", As: "txt"}}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("doc:1"))
			Expect(res.Docs[0].Fields["txt"]).To(BeEquivalentTo("telmatosaurus"))
		})

		It("should FTSynUpdate", Label("search", "ftsynupdate"), func() {

			text1 := &FieldSchema{FieldName: "title", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "body", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{OnHash: true}, text1, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			resSynUpdate, err := adapter.FTSynUpdateWithArgs(ctx, "idx1", "id1", &FTSynUpdateOptions{SkipInitialScan: true}, []interface{}{"boy", "child", "offspring"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSynUpdate).To(BeEquivalentTo("OK"))
			adapter.HSet(ctx, "doc1", "title", "he is a baby", "body", "this is a test")

			resSynUpdate, err = adapter.FTSynUpdateWithArgs(ctx, "idx1", "id1", &FTSynUpdateOptions{SkipInitialScan: true}, []interface{}{"baby"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSynUpdate).To(BeEquivalentTo("OK"))
			adapter.HSet(ctx, "doc2", "title", "he is another baby", "body", "another test")

			res, err := adapter.FTSearchWithArgs(ctx, "idx1", "child", &FTSearchOptions{Expander: "SYNONYM"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("doc2"))
			Expect(res.Docs[0].Fields["title"]).To(BeEquivalentTo("he is another baby"))
			Expect(res.Docs[0].Fields["body"]).To(BeEquivalentTo("another test"))
		})

		It("should FTSynDump", Label("search", "ftsyndump"), func() {
			text1 := &FieldSchema{FieldName: "title", FieldType: SearchFieldTypeText}
			text2 := &FieldSchema{FieldName: "body", FieldType: SearchFieldTypeText}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{OnHash: true}, text1, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			resSynUpdate, err := adapter.FTSynUpdate(ctx, "idx1", "id1", []interface{}{"boy", "child", "offspring"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSynUpdate).To(BeEquivalentTo("OK"))

			resSynUpdate, err = adapter.FTSynUpdate(ctx, "idx1", "id1", []interface{}{"baby", "child"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSynUpdate).To(BeEquivalentTo("OK"))

			resSynUpdate, err = adapter.FTSynUpdate(ctx, "idx1", "id1", []interface{}{"tree", "wood"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSynUpdate).To(BeEquivalentTo("OK"))

			resSynDump, err := adapter.FTSynDump(ctx, "idx1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resSynDump[0].Term).To(BeEquivalentTo("baby"))
			Expect(resSynDump[0].Synonyms).To(BeEquivalentTo([]string{"id1"}))
			Expect(resSynDump[1].Term).To(BeEquivalentTo("wood"))
			Expect(resSynDump[1].Synonyms).To(BeEquivalentTo([]string{"id1"}))
			Expect(resSynDump[2].Term).To(BeEquivalentTo("boy"))
			Expect(resSynDump[2].Synonyms).To(BeEquivalentTo([]string{"id1"}))
			Expect(resSynDump[3].Term).To(BeEquivalentTo("tree"))
			Expect(resSynDump[3].Synonyms).To(BeEquivalentTo([]string{"id1"}))
			Expect(resSynDump[4].Term).To(BeEquivalentTo("child"))
			Expect(resSynDump[4].Synonyms).To(Or(BeEquivalentTo([]string{"id1"}), BeEquivalentTo([]string{"id1", "id1"})))
			Expect(resSynDump[5].Term).To(BeEquivalentTo("offspring"))
			Expect(resSynDump[5].Synonyms).To(BeEquivalentTo([]string{"id1"}))

		})

		It("should FTCreate json with alias", Label("search", "ftcreate"), func() {

			text1 := &FieldSchema{FieldName: "$.name", FieldType: SearchFieldTypeText, As: "name"}
			num1 := &FieldSchema{FieldName: "$.num", FieldType: SearchFieldTypeNumeric, As: "num"}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{OnJSON: true, Prefix: []interface{}{"king:"}}, text1, num1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.JSONSet(ctx, "king:1", "$", `{"name": "henry", "num": 42}`)
			adapter.JSONSet(ctx, "king:2", "$", `{"name": "james", "num": 3.14}`)

			res, err := adapter.FTSearch(ctx, "idx1", "@name:henry").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("king:1"))
			Expect(res.Docs[0].Fields["$"]).To(BeEquivalentTo(`{"name":"henry","num":42}`))

			res, err = adapter.FTSearch(ctx, "idx1", "@num:[0 10]").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("king:2"))
			Expect(res.Docs[0].Fields["$"]).To(BeEquivalentTo(`{"name":"james","num":3.14}`))
		})

		It("should FTCreate json with multipath", Label("search", "ftcreate"), func() {

			tag1 := &FieldSchema{FieldName: "$..name", FieldType: SearchFieldTypeTag, As: "name"}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{OnJSON: true, Prefix: []interface{}{"king:"}}, tag1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.JSONSet(ctx, "king:1", "$", `{"name": "henry", "country": {"name": "england"}}`)

			res, err := adapter.FTSearch(ctx, "idx1", "@name:{england}").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("king:1"))
			Expect(res.Docs[0].Fields["$"]).To(BeEquivalentTo(`{"name":"henry","country":{"name":"england"}}`))
		})

		It("should FTCreate json with jsonpath", Label("search", "ftcreate"), func() {

			text1 := &FieldSchema{FieldName: `$["prod:name"]`, FieldType: SearchFieldTypeText, As: "name"}
			text2 := &FieldSchema{FieldName: `$.prod:name`, FieldType: SearchFieldTypeText, As: "name_unsupported"}
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{OnJSON: true}, text1, text2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.JSONSet(ctx, "doc:1", "$", `{"prod:name": "RediSearch"}`)

			res, err := adapter.FTSearch(ctx, "idx1", "@name:RediSearch").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("doc:1"))
			Expect(res.Docs[0].Fields["$"]).To(BeEquivalentTo(`{"prod:name":"RediSearch"}`))

			res, err = adapter.FTSearch(ctx, "idx1", "@name_unsupported:RediSearch").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "@name:RediSearch", &FTSearchOptions{Return: []FTSearchReturn{{FieldName: "name"}}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(1))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("doc:1"))
			Expect(res.Docs[0].Fields["name"]).To(BeEquivalentTo("RediSearch"))

		})

		It("should FTCreate VECTOR", Label("search", "ftcreate"), func() {
			hnswOptions := &FTHNSWOptions{Type: "FLOAT32", Dim: 2, DistanceMetric: "L2"}
			val, err := adapter.FTCreate(ctx, "idx1",
				&FTCreateOptions{},
				&FieldSchema{FieldName: "v", FieldType: SearchFieldTypeVector, VectorArgs: &FTVectorArgs{HNSWOptions: hnswOptions}}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "a", "v", "aaaaaaaa")
			adapter.HSet(ctx, "b", "v", "aaaabaaa")
			adapter.HSet(ctx, "c", "v", "aaaaabaa")

			searchOptions := &FTSearchOptions{
				Return:         []FTSearchReturn{{FieldName: "__v_score"}},
				SortBy:         []FTSearchSortBy{{FieldName: "__v_score", Asc: true}},
				DialectVersion: 2,
				Params:         map[string]interface{}{"vec": "aaaaaaaa"},
			}
			res, err := adapter.FTSearchWithArgs(ctx, "idx1", "*=>[KNN 2 @v $vec]", searchOptions).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("a"))
			Expect(res.Docs[0].Fields["__v_score"]).To(BeEquivalentTo("0"))
		})

		It("should FTCreate and FTSearch text params", Label("search", "ftcreate", "ftsearch"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{FieldName: "name", FieldType: SearchFieldTypeText}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "name", "Alice")
			adapter.HSet(ctx, "doc2", "name", "Bob")
			adapter.HSet(ctx, "doc3", "name", "Carol")

			res1, err := adapter.FTSearchWithArgs(ctx, "idx1", "@name:($name1 | $name2 )", &FTSearchOptions{Params: map[string]interface{}{"name1": "Alice", "name2": "Bob"}, DialectVersion: 2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(2)))
			Expect(res1.Docs[0].ID).To(BeEquivalentTo("doc1"))
			Expect(res1.Docs[1].ID).To(BeEquivalentTo("doc2"))

		})

		It("should FTCreate and FTSearch numeric params", Label("search", "ftcreate", "ftsearch"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{FieldName: "numval", FieldType: SearchFieldTypeNumeric}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "numval", 101)
			adapter.HSet(ctx, "doc2", "numval", 102)
			adapter.HSet(ctx, "doc3", "numval", 103)

			res1, err := adapter.FTSearchWithArgs(ctx, "idx1", "@numval:[$min $max]", &FTSearchOptions{Params: map[string]interface{}{"min": 101, "max": 102}, DialectVersion: 2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(2)))
			Expect(res1.Docs[0].ID).To(BeEquivalentTo("doc1"))
			Expect(res1.Docs[1].ID).To(BeEquivalentTo("doc2"))

		})

		It("should FTCreate and FTSearch geo params", Label("search", "ftcreate", "ftsearch"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{FieldName: "g", FieldType: SearchFieldTypeGeo}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "doc1", "g", "29.69465, 34.95126")
			adapter.HSet(ctx, "doc2", "g", "29.69350, 34.94737")
			adapter.HSet(ctx, "doc3", "g", "29.68746, 34.94882")

			res1, err := adapter.FTSearchWithArgs(ctx, "idx1", "@g:[$lon $lat $radius $units]", &FTSearchOptions{Params: map[string]interface{}{"lat": "34.95126", "lon": "29.69465", "radius": 1000, "units": "km"}, DialectVersion: 2}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(3)))
			Expect(res1.Docs[0].ID).To(BeEquivalentTo("doc1"))
			Expect(res1.Docs[1].ID).To(BeEquivalentTo("doc2"))
			Expect(res1.Docs[2].ID).To(BeEquivalentTo("doc3"))

		})

		It("should FTConfigSet and FTConfigGet dialect", Label("search", "ftconfigget", "ftconfigset", "NonRedisEnterprise"), func() {
			res, err := adapter.FTConfigSet(ctx, "DEFAULT_DIALECT", "1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(BeEquivalentTo("OK"))

			defDialect, err := adapter.FTConfigGet(ctx, "DEFAULT_DIALECT").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(defDialect).To(BeEquivalentTo(map[string]interface{}{"DEFAULT_DIALECT": "1"}))

			res, err = adapter.FTConfigSet(ctx, "DEFAULT_DIALECT", "2").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(BeEquivalentTo("OK"))

			defDialect, err = adapter.FTConfigGet(ctx, "DEFAULT_DIALECT").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(defDialect).To(BeEquivalentTo(map[string]interface{}{"DEFAULT_DIALECT": "2"}))
		})

		It("should FTCreate WithSuffixtrie", Label("search", "ftcreate", "ftinfo"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{FieldName: "txt", FieldType: SearchFieldTypeText}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			res, err := adapter.FTInfo(ctx, "idx1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Attributes[0].Attribute).To(BeEquivalentTo("txt"))

			resDrop, err := adapter.FTDropIndex(ctx, "idx1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resDrop).To(BeEquivalentTo("OK"))

			// create withsuffixtrie index - text field
			val, err = adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{FieldName: "txt", FieldType: SearchFieldTypeText, WithSuffixtrie: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			res, err = adapter.FTInfo(ctx, "idx1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Attributes[0].WithSuffixtrie).To(BeTrue())

			resDrop, err = adapter.FTDropIndex(ctx, "idx1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(resDrop).To(BeEquivalentTo("OK"))

			// create withsuffixtrie index - tag field
			val, err = adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{FieldName: "t", FieldType: SearchFieldTypeTag, WithSuffixtrie: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			res, err = adapter.FTInfo(ctx, "idx1").Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Attributes[0].WithSuffixtrie).To(BeTrue())
		})

		It("should test dialect 4", Label("search", "ftcreate", "ftsearch", "NonRedisEnterprise"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{
				Prefix: []interface{}{"resource:"},
			}, &FieldSchema{
				FieldName: "uuid",
				FieldType: SearchFieldTypeTag,
			}, &FieldSchema{
				FieldName: "tags",
				FieldType: SearchFieldTypeTag,
			}, &FieldSchema{
				FieldName: "description",
				FieldType: SearchFieldTypeText,
			}, &FieldSchema{
				FieldName: "rating",
				FieldType: SearchFieldTypeNumeric,
			}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))

			adapter.HSet(ctx, "resource:1", map[string]interface{}{
				"uuid":        "123e4567-e89b-12d3-a456-426614174000",
				"tags":        "finance|crypto|$btc|blockchain",
				"description": "Analysis of blockchain technologies & Bitcoin's potential.",
				"rating":      5,
			})
			adapter.HSet(ctx, "resource:2", map[string]interface{}{
				"uuid":        "987e6543-e21c-12d3-a456-426614174999",
				"tags":        "health|well-being|fitness|new-year's-resolutions",
				"description": "Health trends for the new year, including fitness regimes.",
				"rating":      4,
			})

			res, err := adapter.FTSearchWithArgs(ctx, "idx1", "@uuid:{$uuid}",
				&FTSearchOptions{
					DialectVersion: 2,
					Params:         map[string]interface{}{"uuid": "123e4567-e89b-12d3-a456-426614174000"},
				}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(int64(1)))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("resource:1"))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "@uuid:{$uuid}",
				&FTSearchOptions{
					DialectVersion: 4,
					Params:         map[string]interface{}{"uuid": "123e4567-e89b-12d3-a456-426614174000"},
				}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Total).To(BeEquivalentTo(int64(1)))
			Expect(res.Docs[0].ID).To(BeEquivalentTo("resource:1"))

			adapter.HSet(ctx, "test:1", map[string]interface{}{
				"uuid":  "3d3586fe-0416-4572-8ce",
				"email": "adriano@acme.com.ie",
				"num":   5,
			})

			// Create the index
			ftCreateOptions := &FTCreateOptions{
				Prefix: []interface{}{"test:"},
			}
			schema := []*FieldSchema{
				{
					FieldName: "uuid",
					FieldType: SearchFieldTypeTag,
				},
				{
					FieldName: "email",
					FieldType: SearchFieldTypeTag,
				},
				{
					FieldName: "num",
					FieldType: SearchFieldTypeNumeric,
				},
			}

			val, err = adapter.FTCreate(ctx, "idx_hash", ftCreateOptions, schema...).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("OK"))

			ftSearchOptions := &FTSearchOptions{
				DialectVersion: 4,
				Params: map[string]interface{}{
					"uuid":  "3d3586fe-0416-4572-8ce",
					"email": "adriano@acme.com.ie",
				},
			}

			res, err = adapter.FTSearchWithArgs(ctx, "idx_hash", "@uuid:{$uuid}", ftSearchOptions).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("test:1"))
			Expect(res.Docs[0].Fields["uuid"]).To(BeEquivalentTo("3d3586fe-0416-4572-8ce"))

			res, err = adapter.FTSearchWithArgs(ctx, "idx_hash", "@email:{$email}", ftSearchOptions).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("test:1"))
			Expect(res.Docs[0].Fields["email"]).To(BeEquivalentTo("adriano@acme.com.ie"))

			ftSearchOptions.Params = map[string]interface{}{"num": 5}
			res, err = adapter.FTSearchWithArgs(ctx, "idx_hash", "@num:[5]", ftSearchOptions).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("test:1"))
			Expect(res.Docs[0].Fields["num"]).To(BeEquivalentTo("5"))
		})

		It("should FTCreate GeoShape", Label("search", "ftcreate", "ftsearch"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{FieldName: "geom", FieldType: SearchFieldTypeGeoShape, GeoShapeFieldType: "FLAT"}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "small", "geom", "POLYGON((1 1, 1 100, 100 100, 100 1, 1 1))")
			adapter.HSet(ctx, "large", "geom", "POLYGON((1 1, 1 200, 200 200, 200 1, 1 1))")

			res1, err := adapter.FTSearchWithArgs(ctx, "idx1", "@geom:[WITHIN $poly]",
				&FTSearchOptions{
					DialectVersion: 3,
					Params:         map[string]interface{}{"poly": "POLYGON((0 0, 0 150, 150 150, 150 0, 0 0))"},
				}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Total).To(BeEquivalentTo(int64(1)))
			Expect(res1.Docs[0].ID).To(BeEquivalentTo("small"))

			res2, err := adapter.FTSearchWithArgs(ctx, "idx1", "@geom:[CONTAINS $poly]",
				&FTSearchOptions{
					DialectVersion: 3,
					Params:         map[string]interface{}{"poly": "POLYGON((2 2, 2 50, 50 50, 50 2, 2 2))"},
				}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res2.Total).To(BeEquivalentTo(int64(2)))
		})

		It("should create search index with FLOAT16 and BFLOAT16 vectors", Label("search", "ftcreate", "NonRedisEnterprise"), func() {
			val, err := adapter.FTCreate(ctx, "index", &FTCreateOptions{},
				&FieldSchema{FieldName: "float16", FieldType: SearchFieldTypeVector, VectorArgs: &FTVectorArgs{FlatOptions: &FTFlatOptions{Type: "FLOAT16", Dim: 768, DistanceMetric: "COSINE"}}},
				&FieldSchema{FieldName: "bfloat16", FieldType: SearchFieldTypeVector, VectorArgs: &FTVectorArgs{FlatOptions: &FTFlatOptions{Type: "BFLOAT16", Dim: 768, DistanceMetric: "COSINE"}}},
			).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "index", 2)
		})

		It("should test geoshapes query intersects and disjoint", Label("NonRedisEnterprise"), func() {
			_, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{}, &FieldSchema{
				FieldName:         "g",
				FieldType:         SearchFieldTypeGeoShape,
				GeoShapeFieldType: "FLAT",
			}).Result()
			Expect(err).NotTo(HaveOccurred())

			adapter.HSet(ctx, "doc_point1", "g", "POINT (10 10)")
			adapter.HSet(ctx, "doc_point2", "g", "POINT (50 50)")
			adapter.HSet(ctx, "doc_polygon1", "g", "POLYGON ((20 20, 25 35, 35 25, 20 20))")
			adapter.HSet(ctx, "doc_polygon2", "g", "POLYGON ((60 60, 65 75, 70 70, 65 55, 60 60))")

			intersection, err := adapter.FTSearchWithArgs(ctx, "idx1", "@g:[intersects $shape]",
				&FTSearchOptions{
					DialectVersion: 3,
					Params:         map[string]interface{}{"shape": "POLYGON((15 15, 75 15, 50 70, 20 40, 15 15))"},
				}).Result()
			Expect(err).NotTo(HaveOccurred())
			_assert_geosearch_result(&intersection, []string{"doc_point2", "doc_polygon1"})

			disjunction, err := adapter.FTSearchWithArgs(ctx, "idx1", "@g:[disjoint $shape]",
				&FTSearchOptions{
					DialectVersion: 3,
					Params:         map[string]interface{}{"shape": "POLYGON((15 15, 75 15, 50 70, 20 40, 15 15))"},
				}).Result()
			Expect(err).NotTo(HaveOccurred())
			_assert_geosearch_result(&disjunction, []string{"doc_point1", "doc_polygon2"})
		})

		It("should test geoshapes query contains and within", func() {
			_, err := adapter.FTCreate(ctx, "idx2", &FTCreateOptions{}, &FieldSchema{
				FieldName:         "g",
				FieldType:         SearchFieldTypeGeoShape,
				GeoShapeFieldType: "FLAT",
			}).Result()
			Expect(err).NotTo(HaveOccurred())

			adapter.HSet(ctx, "doc_point1", "g", "POINT (10 10)")
			adapter.HSet(ctx, "doc_point2", "g", "POINT (50 50)")
			adapter.HSet(ctx, "doc_polygon1", "g", "POLYGON ((20 20, 25 35, 35 25, 20 20))")
			adapter.HSet(ctx, "doc_polygon2", "g", "POLYGON ((60 60, 65 75, 70 70, 65 55, 60 60))")

			containsA, err := adapter.FTSearchWithArgs(ctx, "idx2", "@g:[contains $shape]",
				&FTSearchOptions{
					DialectVersion: 3,
					Params:         map[string]interface{}{"shape": "POINT(25 25)"},
				}).Result()
			Expect(err).NotTo(HaveOccurred())
			_assert_geosearch_result(&containsA, []string{"doc_polygon1"})

			containsB, err := adapter.FTSearchWithArgs(ctx, "idx2", "@g:[contains $shape]",
				&FTSearchOptions{
					DialectVersion: 3,
					Params:         map[string]interface{}{"shape": "POLYGON((24 24, 24 26, 25 25, 24 24))"},
				}).Result()
			Expect(err).NotTo(HaveOccurred())
			_assert_geosearch_result(&containsB, []string{"doc_polygon1"})

			within, err := adapter.FTSearchWithArgs(ctx, "idx2", "@g:[within $shape]",
				&FTSearchOptions{
					DialectVersion: 3,
					Params:         map[string]interface{}{"shape": "POLYGON((15 15, 75 15, 50 70, 20 40, 15 15))"},
				}).Result()
			Expect(err).NotTo(HaveOccurred())
			_assert_geosearch_result(&within, []string{"doc_point2", "doc_polygon1"})
		})

		It("should search missing fields", Label("search", "ftcreate", "ftsearch", "NonRedisEnterprise"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{Prefix: []interface{}{"property:"}},
				&FieldSchema{FieldName: "title", FieldType: SearchFieldTypeText, Sortable: true},
				&FieldSchema{FieldName: "features", FieldType: SearchFieldTypeTag, IndexMissing: true},
				&FieldSchema{FieldName: "description", FieldType: SearchFieldTypeText, IndexMissing: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "property:1", map[string]interface{}{
				"title":       "Luxury Villa in Malibu",
				"features":    "pool,sea view,modern",
				"description": "A stunning modern villa overlooking the Pacific Ocean.",
			})

			adapter.HSet(ctx, "property:2", map[string]interface{}{
				"title":       "Downtown Flat",
				"description": "Modern flat in central Paris with easy access to metro.",
			})

			adapter.HSet(ctx, "property:3", map[string]interface{}{
				"title":    "Beachfront Bungalow",
				"features": "beachfront,sun deck",
			})

			res, err := adapter.FTSearchWithArgs(ctx, "idx1", "ismissing(@features)", &FTSearchOptions{DialectVersion: 4, Return: []FTSearchReturn{{FieldName: "id"}}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("property:2"))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "-ismissing(@features)", &FTSearchOptions{DialectVersion: 4, Return: []FTSearchReturn{{FieldName: "id"}}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("property:1"))
			Expect(res.Docs[1].ID).To(BeEquivalentTo("property:3"))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "ismissing(@description)", &FTSearchOptions{DialectVersion: 4, Return: []FTSearchReturn{{FieldName: "id"}}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("property:3"))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "-ismissing(@description)", &FTSearchOptions{DialectVersion: 4, Return: []FTSearchReturn{{FieldName: "id"}}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("property:1"))
			Expect(res.Docs[1].ID).To(BeEquivalentTo("property:2"))
		})

		It("should search empty fields", Label("search", "ftcreate", "ftsearch", "NonRedisEnterprise"), func() {
			val, err := adapter.FTCreate(ctx, "idx1", &FTCreateOptions{Prefix: []interface{}{"property:"}},
				&FieldSchema{FieldName: "title", FieldType: SearchFieldTypeText, Sortable: true},
				&FieldSchema{FieldName: "features", FieldType: SearchFieldTypeTag, IndexEmpty: true},
				&FieldSchema{FieldName: "description", FieldType: SearchFieldTypeText, IndexEmpty: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeEquivalentTo("OK"))
			WaitForIndexing(client, "idx1", 2)

			adapter.HSet(ctx, "property:1", map[string]interface{}{
				"title":       "Luxury Villa in Malibu",
				"features":    "pool,sea view,modern",
				"description": "A stunning modern villa overlooking the Pacific Ocean.",
			})

			adapter.HSet(ctx, "property:2", map[string]interface{}{
				"title":       "Downtown Flat",
				"features":    "",
				"description": "Modern flat in central Paris with easy access to metro.",
			})

			adapter.HSet(ctx, "property:3", map[string]interface{}{
				"title":       "Beachfront Bungalow",
				"features":    "beachfront,sun deck",
				"description": "",
			})

			res, err := adapter.FTSearchWithArgs(ctx, "idx1", "@features:{\"\"}", &FTSearchOptions{DialectVersion: 4, Return: []FTSearchReturn{{FieldName: "id"}}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("property:2"))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "-@features:{\"\"}", &FTSearchOptions{DialectVersion: 4, Return: []FTSearchReturn{{FieldName: "id"}}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("property:1"))
			Expect(res.Docs[1].ID).To(BeEquivalentTo("property:3"))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "@description:''", &FTSearchOptions{DialectVersion: 4, Return: []FTSearchReturn{{FieldName: "id"}}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("property:3"))

			res, err = adapter.FTSearchWithArgs(ctx, "idx1", "-@description:''", &FTSearchOptions{DialectVersion: 4, Return: []FTSearchReturn{{FieldName: "id"}}, NoContent: true}).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Docs[0].ID).To(BeEquivalentTo("property:1"))
			Expect(res.Docs[1].ID).To(BeEquivalentTo("property:2"))
		})
	})
}

func libCode(libName string) string {
	return fmt.Sprintf("#!js api_version=1.0 name=%s\n redis.registerFunction('foo', ()=>{{return 'bar'}})", libName)
}

func libCodeWithConfig(libName string) string {
	lib := `#!js api_version=1.0 name=%s

	var last_update_field_name = "__last_update__"

	if (redis.config.last_update_field_name !== undefined) {
		if (typeof redis.config.last_update_field_name != 'string') {
			throw "last_update_field_name must be a string";
		}
		last_update_field_name = redis.config.last_update_field_name
	}

	redis.registerFunction("hset", function(client, key, field, val){
		// get the current time in ms
		var curr_time = client.call("time")[0];
		return client.call('hset', key, field, val, last_update_field_name, curr_time);
	});`
	return fmt.Sprintf(lib, libName)
}

type numberStruct struct {
	Number int
}

func (s *numberStruct) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *numberStruct) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, s)
}

func WaitForIndexing(c rueidis.Client, index string, protocolVersion int) {
	adapter := NewAdapter(c)
	time.Sleep(1000 * time.Millisecond)
	for {
		res, err := adapter.FTInfo(context.Background(), index).Result()
		Expect(err).NotTo(HaveOccurred())
		if protocolVersion == 2 {
			if res.Indexing == 0 {
				return
			}
			time.Sleep(100 * time.Millisecond)
		} else {
			return
		}
	}
}

func _assert_geosearch_result(result *FTSearchResult, expectedDocIDs []string) {
	ids := make([]string, len(result.Docs))
	for i, doc := range result.Docs {
		ids[i] = doc.ID
	}
	Expect(ids).To(ConsistOf(expectedDocIDs))
	Expect(result.Total).To(BeEquivalentTo(len(expectedDocIDs)))
}
