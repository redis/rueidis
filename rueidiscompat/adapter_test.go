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

			It("should SortRO with GET patterns", func() {
				Expect(adapter.LPush(ctx, "list_sortro_get", "item1", "item2", "item3").Err()).NotTo(HaveOccurred())
				Expect(adapter.Set(ctx, "weight_item1", "10", 0).Err()).NotTo(HaveOccurred())
				Expect(adapter.Set(ctx, "weight_item2", "5", 0).Err()).NotTo(HaveOccurred())
				Expect(adapter.Set(ctx, "weight_item3", "15", 0).Err()).NotTo(HaveOccurred())
				Expect(adapter.Set(ctx, "data_item1", "alpha", 0).Err()).NotTo(HaveOccurred())
				Expect(adapter.Set(ctx, "data_item2", "beta", 0).Err()).NotTo(HaveOccurred())
				Expect(adapter.Set(ctx, "data_item3", "gamma", 0).Err()).NotTo(HaveOccurred())

				// Sort by external weights and get external data
				sort := Sort{By: "weight_*", Get: []string{"data_*", "#"}}
				sorted, err := adapter.SortRO(ctx, "list_sortro_get", sort).Result()
				Expect(err).NotTo(HaveOccurred())
				// Expected order: item2 (5), item1 (10), item3 (15)
				// Expected result: data_item2, item2, data_item1, item1, data_item3, item3
				Expect(sorted).To(Equal([]string{"beta", "item2", "alpha", "item1", "gamma", "item3"}))

				// Compare with regular Sort
				regularSorted, err := adapter.Sort(ctx, "list_sortro_get", sort).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(sorted).To(Equal(regularSorted))
			})

			It("should SortRO with BY clause not matching", func() {
				Expect(adapter.LPush(ctx, "list_sortro_by_nomatch", "a", "b", "c").Err()).NotTo(HaveOccurred())
				// "nosort" pattern doesn't match any keys, so it should sort lexicographically
				sort := Sort{By: "nosort"}
				sortedRO, err := adapter.SortRO(ctx, "list_sortro_by_nomatch", sort).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(sortedRO).To(Equal([]string{"a", "b", "c"})) // Default lexicographical sort

				sortedNonRO, err := adapter.Sort(ctx, "list_sortro_by_nomatch", sort).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(sortedRO).To(Equal(sortedNonRO))
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

		It("should SortStore with GET and BY clauses", func() {
			Expect(adapter.LPush(ctx, "list_sortstore", "obj1", "obj2", "obj3").Err()).NotTo(HaveOccurred())
			Expect(adapter.Set(ctx, "weight_obj1", "3", 0).Err()).NotTo(HaveOccurred())
			Expect(adapter.Set(ctx, "weight_obj2", "1", 0).Err()).NotTo(HaveOccurred())
			Expect(adapter.Set(ctx, "weight_obj3", "2", 0).Err()).NotTo(HaveOccurred())
			Expect(adapter.Set(ctx, "data_obj1", "gamma", 0).Err()).NotTo(HaveOccurred())
			Expect(adapter.Set(ctx, "data_obj2", "alpha", 0).Err()).NotTo(HaveOccurred())
			Expect(adapter.Set(ctx, "data_obj3", "beta", 0).Err()).NotTo(HaveOccurred())

			storeKey := "list_sortstore_dest"
			sort := Sort{
				By:    "weight_*",               // Sort by external weights
				Get:   []string{"data_*", "#"}, // Get external data and the element itself
				Store: storeKey,
			}
			count, err := adapter.SortStore(ctx, "list_sortstore", storeKey, sort).Result() // SortStore returns count of elements in stored list
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(int64(6))) // 3 elements * 2 GET patterns (data_* and #) = 6 items in destination

			// Expected order in source list (before sort): obj3, obj2, obj1 (due to LPush)
			// Weights: obj1=3, obj2=1, obj3=2
			// Sorted order by weight (ascending): obj2 (1), obj3 (2), obj1 (3)
			// Stored result with GET: data_obj2, obj2, data_obj3, obj3, data_obj1, obj1
			storedValues, err := adapter.LRange(ctx, storeKey, 0, -1).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(storedValues).To(Equal([]string{"alpha", "obj2", "beta", "obj3", "gamma", "obj1"}))
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

		It("should SetWithArgs NX with GET and KEEPTTL", func() {
			// Key does not exist, NX ensures it's set. GET returns nil. KEEPTTL has no effect.
			argsNXGetKeepTTL := SetArgs{Mode: "NX", Get: true, KeepTTL: true}
			res, err := adapter.SetArgs(ctx, "nxgetkeepttlkey1", "value1", argsNXGetKeepTTL).Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue(), "Error should be nil as key is set")
			Expect(res).To(BeEmpty(), "GET should return nil as key didn't exist")
			Expect(adapter.Get(ctx, "nxgetkeepttlkey1").Val()).To(Equal("value1"))
			Expect(adapter.TTL(ctx, "nxgetkeepttlkey1").Val()).To(Equal(time.Duration(-1)), "TTL should be -1 as KEEPTTL has no effect on new key")

			// Key exists, NX prevents set. GET would return old value if supported by this combination (it's not standard for SET NX GET to return old value).
			// For go-redis/rueidiscompat, if SET fails due to NX, GET is not performed.
			Expect(adapter.Set(ctx, "nxgetkeepttlkey2", "oldvalue", 10*time.Second).Err()).NotTo(HaveOccurred()) // Set with TTL
			res, err = adapter.SetArgs(ctx, "nxgetkeepttlkey2", "newvalue", argsNXGetKeepTTL).Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue(), "Error should be redis.Nil as key exists and NX prevents set")
			Expect(res).To(BeEmpty(), "Value should be empty as set failed")
			Expect(adapter.Get(ctx, "nxgetkeepttlkey2").Val()).To(Equal("oldvalue"))                  // Value should be unchanged
			Expect(adapter.TTL(ctx, "nxgetkeepttlkey2").Val()).To(BeNumerically(">", 0*time.Second)) // TTL should be unchanged
		})

		It("should SetWithArgs XX with GET and KEEPTTL", func() {
			// Key does not exist, XX prevents set.
			argsXXGetKeepTTL := SetArgs{Mode: "XX", Get: true, KeepTTL: true}
			res, err := adapter.SetArgs(ctx, "xxgetkeepttlkey1", "value1", argsXXGetKeepTTL).Result()
			Expect(rueidis.IsRedisNil(err)).To(BeTrue(), "Error should be redis.Nil as key doesn't exist and XX prevents set")
			Expect(res).To(BeEmpty())

			// Key exists, XX allows set. GET returns old value. KEEPTTL preserves original TTL.
			Expect(adapter.Set(ctx, "xxgetkeepttlkey2", "oldvalue", 20*time.Second).Err()).NotTo(HaveOccurred())
			res, err = adapter.SetArgs(ctx, "xxgetkeepttlkey2", "newvalue", argsXXGetKeepTTL).Result()
			Expect(err).NotTo(HaveOccurred(), "Set should succeed")
			Expect(res).To(Equal("oldvalue"), "GET should return old value")
			Expect(adapter.Get(ctx, "xxgetkeepttlkey2").Val()).To(Equal("newvalue"))                   // Value should be updated
			Expect(adapter.TTL(ctx, "xxgetkeepttlkey2").Val()).To(BeNumerically(">", 0*time.Second))  // TTL should be preserved
			Expect(adapter.TTL(ctx, "xxgetkeepttlkey2").Val()).To(BeNumerically("<=", 20*time.Second)) // TTL should be preserved
		})

		It("should SetWithArgs with EXAT/PXAT in the past", func() {
			pastTime := time.Now().Add(-5 * time.Hour)
			// pastTimestamp := pastTime.Unix() // Not directly used by SetArgs, but for context
			// pastMillisecondTimestamp := pastTime.UnixNano() / int64(time.Millisecond) // Not directly used

			// EXAT
			argsExatPast := SetArgs{ExpireAt: pastTime}
			_, err := adapter.SetArgs(ctx, "exatpastkey", "value", argsExatPast).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(adapter.Exists(ctx, "exatpastkey").Val()).To(Equal(int64(0)), "Key should not exist as EXAT was in the past")

			// PXAT (simulated via ExpireAt as SetArgs uses time.Time for both)
			argsPxatPast := SetArgs{ExpireAt: pastTime}
			_, err = adapter.SetArgs(ctx, "pxatpastkey", "value", argsPxatPast).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(adapter.Exists(ctx, "pxatpastkey").Val()).To(Equal(int64(0)), "Key should not exist as PXAT (simulated via EXAT) was in the past")

			// Verify TTL of a key set with EXAT/PXAT in the past (it should be -2, meaning not exists)
			Expect(adapter.TTL(ctx, "exatpastkey").Val()).To(Equal(time.Duration(-2)))
			Expect(adapter.TTL(ctx, "pxatpastkey").Val()).To(Equal(time.Duration(-2)))
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
				Expect(res).To(Equal([]int64{1, 1, -2})) // key1, key2 set; key200 noexist
			})

			It("should HPExpire", Label("hash-expiration", "NonRedisEnterprise"), func() {
				res, err := adapter.HPExpire(ctx, "no_such_key", 10*time.Second, "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(res).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err = adapter.HPExpire(ctx, "myhash", 10000*time.Millisecond, "key1", "key2", "key200").Result() // 10 seconds in ms
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, 1, -2}))
			})

			It("should HExpireAt", Label("hash-expiration", "NonRedisEnterprise"), func() {
				expireTime := time.Now().Add(10 * time.Second)
				resEmpty, err := adapter.HExpireAt(ctx, "no_such_key", expireTime, "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err := adapter.HExpireAt(ctx, "myhash", expireTime, "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, 1, -2}))
			})

			It("should HPExpireAt", Label("hash-expiration", "NonRedisEnterprise"), func() {
				expireTime := time.Now().Add(10 * time.Second)
				resEmpty, err := adapter.HPExpireAt(ctx, "no_such_key", expireTime, "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				res, err := adapter.HPExpireAt(ctx, "myhash", expireTime, "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, 1, -2}))
			})

			It("should HPersist", Label("hash-expiration", "NonRedisEnterprise"), func() {
				resEmpty, err := adapter.HPersist(ctx, "no_such_key", "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2})) // Field does not exist in a non-existent key

				for i := 0; i < 3; i++ { // Using fewer keys for clarity
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}

				// Set TTL on key0, key1
				_, err = adapter.HExpire(ctx, "myhash", 10*time.Second, "key0", "key1").Result()
				Expect(err).NotTo(HaveOccurred())

				// key0 had TTL, persist removes it (returns 1)
				// key2 had no TTL, persist does nothing (returns 0)
				// key100 does not exist in hash (returns -2)
				res, err := adapter.HPersist(ctx, "myhash", "key0", "key2", "key100").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, 0, -2}))

				Expect(adapter.HTTL(ctx, "myhash", "key0").Val()[0]).To(Equal(int64(-1))) // TTL removed
				Expect(adapter.HTTL(ctx, "myhash", "key1").Val()[0]).To(BeNumerically(">", 0))  // Still has TTL
				Expect(adapter.HTTL(ctx, "myhash", "key2").Val()[0]).To(Equal(int64(-1))) // No TTL originally
			})

			It("should HExpireWithArgs conditional flags", Label("hash-expiration", "NonRedisEnterprise"), func() {
				hashKey := "hexpireargskey"
				Expect(adapter.HSet(ctx, hashKey, "f1", "v1", "f2", "v2", "f3", "v3", "f4", "v4", "f5", "v5").Err()).NotTo(HaveOccurred())

				// Setup initial TTLs for some fields
				Expect(adapter.HExpire(ctx, hashKey, 20*time.Second, "f1", "f2").Err()).NotTo(HaveOccurred()) // f1, f2 have TTL
				// f3, f4, f5 have no TTL

				// NX: Set TTL only if field has no expiry
				// f1 (has TTL) -> 0 (not set)
				// f3 (no TTL) -> 1 (set)
				// f6 (no exist) -> -2 (not set, field no exist)
				argsNX := HExpireArgs{Condition: "NX"}
				resNX, err := adapter.HExpireArgs(ctx, hashKey, 30*time.Second, argsNX, "f1", "f3", "f6").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resNX).To(Equal([]int64{0, 1, -2}))
				Expect(adapter.HTTL(ctx, hashKey, "f1").Val()[0]).To(BeNumerically("~", 20, 2)) // Unchanged by NX
				Expect(adapter.HTTL(ctx, hashKey, "f3").Val()[0]).To(BeNumerically("~", 30, 2)) // Set by NX

				// XX: Set TTL only if field already has an expiry
				// f1 (has TTL) -> 1 (set)
				// f4 (no TTL) -> 0 (not set)
				argsXX := HExpireArgs{Condition: "XX"}
				resXX, err := adapter.HExpireArgs(ctx, hashKey, 40*time.Second, argsXX, "f1", "f4").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resXX).To(Equal([]int64{1, 0}))
				Expect(adapter.HTTL(ctx, hashKey, "f1").Val()[0]).To(BeNumerically("~", 40, 2)) // Updated by XX
				Expect(adapter.HTTL(ctx, hashKey, "f4").Val()[0]).To(Equal(int64(-1)))          // Unchanged by XX

				// GT: Set TTL only if new TTL is greater than current TTL (only for fields with existing TTL)
				// f1 (TTL ~40s), new TTL 50s -> 1 (set)
				// f2 (TTL ~20s), new TTL 10s -> 0 (not set, 10 < 20)
				// f5 (no TTL) -> 0 (not set, no existing TTL)
				argsGT := HExpireArgs{Condition: "GT"}
				resGT, err := adapter.HExpireArgs(ctx, hashKey, 50*time.Second, argsGT, "f1").Result() // Test one by one due to timing sensitivity
				Expect(err).NotTo(HaveOccurred())
				Expect(resGT).To(Equal([]int64{1}))
				Expect(adapter.HTTL(ctx, hashKey, "f1").Val()[0]).To(BeNumerically("~", 50, 2))

				resGT2, err := adapter.HExpireArgs(ctx, hashKey, 10*time.Second, argsGT, "f2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGT2).To(Equal([]int64{0}))
				Expect(adapter.HTTL(ctx, hashKey, "f2").Val()[0]).To(BeNumerically("~", 20, 2)) // Unchanged

				resGT3, err := adapter.HExpireArgs(ctx, hashKey, 50*time.Second, argsGT, "f5").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resGT3).To(Equal([]int64{0}))
				Expect(adapter.HTTL(ctx, hashKey, "f5").Val()[0]).To(Equal(int64(-1))) // Unchanged

				// LT: Set TTL only if new TTL is less than current TTL (only for fields with existing TTL)
				// f1 (TTL ~50s), new TTL 5s -> 1 (set)
				// f2 (TTL ~20s), new TTL 30s -> 0 (not set, 30 > 20)
				// f5 (no TTL) -> 0 (not set, no existing TTL)
				argsLT := HExpireArgs{Condition: "LT"}
				resLT, err := adapter.HExpireArgs(ctx, hashKey, 5*time.Second, argsLT, "f1").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resLT).To(Equal([]int64{1}))
				Expect(adapter.HTTL(ctx, hashKey, "f1").Val()[0]).To(BeNumerically("~", 5, 2))

				resLT2, err := adapter.HExpireArgs(ctx, hashKey, 30*time.Second, argsLT, "f2").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resLT2).To(Equal([]int64{0}))
				Expect(adapter.HTTL(ctx, hashKey, "f2").Val()[0]).To(BeNumerically("~", 20, 2)) // Unchanged

				resLT3, err := adapter.HExpireArgs(ctx, hashKey, 5*time.Second, argsLT, "f5").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resLT3).To(Equal([]int64{0}))
				Expect(adapter.HTTL(ctx, hashKey, "f5").Val()[0]).To(Equal(int64(-1))) // Unchanged

				// Test HPExpireAtArgs with NX (similar logic for other variants)
				pastTime := time.Now().Add(-1 * time.Hour)
				argsAtNX := HExpireArgs{Condition: "NX"}
				resAtNX, err := adapter.HPExpireAtArgs(ctx, hashKey, pastTime, argsAtNX, "f1", "f4").Result() // f1 has TTL, f4 doesn't
				Expect(err).NotTo(HaveOccurred())
				// f1 not set (0), f4 set and immediately expired (1), but HTTL will show -2 (non-existent)
				// The return value for HPEXPIREAT with past time for a field that *gets* set is 1.
				// Then the field is immediately deleted.
				Expect(resAtNX).To(Equal([]int64{0, 1}))
				Expect(adapter.HExists(ctx, hashKey, "f4").Val()).To(BeFalse()) // f4 should be gone
			})

			It("should HExpireTime", Label("hash-expiration", "NonRedisEnterprise"), func() {
				resEmpty, err := adapter.HExpireTime(ctx, "no_such_key", "field1", "field2", "field3").Result()
				Expect(err).To(BeNil())
				Expect(resEmpty).To(BeEquivalentTo([]int64{-2, -2, -2}))

				for i := 0; i < 100; i++ {
					sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
					Expect(sadd.Err()).NotTo(HaveOccurred())
				}
				expireTime := time.Now().Add(10 * time.Second)
				res, err := adapter.HExpireAt(ctx, "myhash", expireTime, "key1", "key200").Result() // key200 doesn't exist
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal([]int64{1, -2}))

				resExpire, err := adapter.HExpireTime(ctx, "myhash", "key1", "key2", "key200").Result() // key2 no TTL, key200 no exist
				Expect(err).NotTo(HaveOccurred())
				Expect(resExpire[0]).To(BeNumerically("~", expireTime.Unix(), 1))
				Expect(resExpire[1]).To(Equal(int64(-1))) // key2 has no expiry
				Expect(resExpire[2]).To(Equal(int64(-2))) // key200 does not exist
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

				resExpire, err := adapter.HPExpireTime(ctx, "myhash", "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resExpire[0]).To(BeNumerically("~", expireAt.UnixMilli(), 1000)) // allow 1s variance for milliseconds
				Expect(resExpire[1]).To(Equal(int64(-1)))
				Expect(resExpire[2]).To(Equal(int64(-2)))
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

				resTTL, err := adapter.HTTL(ctx, "myhash", "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resTTL[0]).To(BeNumerically("~", 10, 2)) // key1 has TTL of ~10s
				Expect(resTTL[1]).To(Equal(int64(-1)))          // key2 has no TTL
				Expect(resTTL[2]).To(Equal(int64(-2)))          // key200 does not exist
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

				resPTTL, err := adapter.HPTTL(ctx, "myhash", "key1", "key2", "key200").Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(resPTTL[0]).To(BeNumerically("~", (10 * time.Second).Milliseconds(), 2000)) // key1 has PTTL of ~10000ms
				Expect(resPTTL[1]).To(Equal(int64(-1)))                                           // key2 has no PTTL
				Expect(resPTTL[2]).To(Equal(int64(-2)))                                           // key200 does not exist
			})
		}
	})
>>>>>>> REPLACE
rueidiscompat/adapter_test.go
<<<<<<< SEARCH
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
=======
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
				Expect(lcs.Matches).To(ConsistOf([]LCSMatchedPosition{ // Order may vary
					{Key1: LCSPosition{Start: 2, End: 3}, Key2: LCSPosition{Start: 0, End: 1}, MatchLen: 0}, // "my" vs "my"
					{Key1: LCSPosition{Start: 4, End: 7}, Key2: LCSPosition{Start: 5, End: 8}, MatchLen: 0}, // "text" vs "text"
				}))

				lcs, err = adapter.LCS(ctx, &LCSQuery{
					Key1:         "LCSkey1{1}", // "ohmytext"
					Key2:         "LCSkey2{1}", // "mynewtext"
					Idx:          true,
					MinMatchLen:  3,
					WithMatchLen: true,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				Expect(lcs.MatchString).To(Equal("")) // MatchString is not populated when Idx is true
				Expect(lcs.Len).To(Equal(int64(6)))   // Total length of LCS
				Expect(lcs.Matches).To(ConsistOf([]LCSMatchedPosition{
					// "mytext" is the LCS.
					// "my" is part of "ohmytext" (idx 2-3) and "mynewtext" (idx 0-1) -> len 2, not >= MinMatchLen 3
					// "text" is part of "ohmytext" (idx 4-7) and "mynewtext" (idx 5-8) -> len 4, >= MinMatchLen 3
					{Key1: LCSPosition{Start: 4, End: 7}, Key2: LCSPosition{Start: 5, End: 8}, MatchLen: 4},
				}))

				// Test with multiple matches of different lengths
				Expect(adapter.Set(ctx, "lcs_str1", "abchellofg", 0).Err()).NotTo(HaveOccurred())
				Expect(adapter.Set(ctx, "lcs_str2", "xyzhelloabc", 0).Err()).NotTo(HaveOccurred())
				lcsMulti, err := adapter.LCS(ctx, &LCSQuery{
					Key1:         "lcs_str1",    // "abchellofg"
					Key2:         "lcs_str2",    // "xyzhelloabc"
					Idx:          true,
					MinMatchLen:  2,
					WithMatchLen: true,
				}).Result()
				Expect(err).NotTo(HaveOccurred())
				// LCS could be "hello" (len 5) or "abc" (len 3)
				// The command returns all maximal matches.
				Expect(lcsMulti.Matches).To(ConsistOf([]LCSMatchedPosition{
					{Key1: LCSPosition{Start: 3, End: 7}, Key2: LCSPosition{Start: 3, End: 7}, MatchLen: 5}, // "hello"
					{Key1: LCSPosition{Start: 0, End: 2}, Key2: LCSPosition{Start: 8, End: 10}, MatchLen: 3}, // "abc"
				}))
				Expect(lcsMulti.Len).To(Equal(int64(5))) // Length of the longest common subsequence, which is "hello"

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
>>>>>>> REPLACE
