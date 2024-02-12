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

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

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
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Adapter Suite")
}

var (
	err             error
	ctx             context.Context
	clientresp2     rueidis.Client
	clusterresp2    rueidis.Client
	clientresp3     rueidis.Client
	clusterresp3    rueidis.Client
	adapterresp2    Cmdable
	adaptercluster2 Cmdable
	adapterresp3    Cmdable
	adaptercluster3 Cmdable
)

var _ = ginkgo.BeforeSuite(func() {
	ctx = context.Background()
	clientresp3, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:6378"},
		ClientName:  "rueidis",
	})
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	clusterresp3, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:7010"},
		ClientName:  "rueidis",
	})
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	adapterresp3 = NewAdapter(clientresp3)
	adaptercluster3 = NewAdapter(clusterresp3)
	clientresp2, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress:  []string{"127.0.0.1:6356"},
		ClientName:   "rueidis",
		DisableCache: true,
	})
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	clusterresp2, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress:  []string{"127.0.0.1:7007"},
		ClientName:   "rueidis",
		DisableCache: true,
	})
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	adapterresp2 = NewAdapter(clientresp2)
	adaptercluster2 = NewAdapter(clusterresp2)
})

var _ = ginkgo.AfterSuite(func() {
	gomega.Expect(adapterresp3.FlushDB(ctx).Err()).NotTo(gomega.HaveOccurred())
	gomega.Expect(adapterresp3.Quit(ctx).Err()).NotTo(gomega.HaveOccurred())
	clientresp3.Close()
	gomega.Expect(adapterresp2.FlushDB(ctx).Err()).NotTo(gomega.HaveOccurred())
	gomega.Expect(adapterresp2.Quit(ctx).Err()).NotTo(gomega.HaveOccurred())
	clientresp2.Close()
})

var _ = ginkgo.Describe("RESP3 Commands", func() {
	testAdapter(true)
	testAdapterCache(true)
	testCluster(true)
})

var _ = ginkgo.Describe("RESP2 Commands", func() {
	testAdapter(false)
	testCluster(false)
})

func testCluster(resp3 bool) {
	var adapter Cmdable

	ginkgo.BeforeEach(func() {
		if resp3 {
			adapter = adaptercluster3
		} else {
			adapter = adaptercluster2
		}
	})

	ginkgo.Describe("Cluster", func() {
		if resp3 {
			ginkgo.It("ClusterShards", func() {
				shards, err := adapter.ClusterShards(ctx).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				m := make(map[int64]struct{})
				for _, shard := range shards {
					for _, slot := range shard.Slots {
						for i := slot.Start; i <= slot.End; i++ {
							m[i] = struct{}{}
						}
					}
				}
				gomega.Expect(m).To(gomega.HaveLen(16384))
			})
		}
		ginkgo.It("ClusterSlots", func() {
			slots, err := adapter.ClusterSlots(ctx).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			m := make(map[int64]struct{})
			for _, slot := range slots {
				for i := slot.Start; i <= slot.End; i++ {
					m[i] = struct{}{}
				}
			}
			gomega.Expect(m).To(gomega.HaveLen(16384))
		})
		ginkgo.It("ClusterNodes", func() {
			nodes, err := adapter.ClusterNodes(ctx).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(strings.Split(strings.TrimSpace(nodes), "\n")).To(gomega.HaveLen(3))
		})
		ginkgo.It("ClusterInfo", func() {
			info, err := adapter.ClusterInfo(ctx).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(info).NotTo(gomega.BeEmpty())
		})
		ginkgo.It("ClusterKeySlot", func() {
			slot, err := adapter.ClusterKeySlot(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(slot).To(gomega.Equal(int64(9842)))
		})
		ginkgo.It("ClusterGetKeysInSlot", func() {
			gomega.Expect(adapter.Set(ctx, "1", "1", 0).Err()).NotTo(gomega.HaveOccurred())
			keys, err := adapter.ClusterGetKeysInSlot(ctx, 9842, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(keys).To(gomega.Equal([]string{"1"}))
			kc, err := adapter.ClusterCountKeysInSlot(ctx, 9842).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(kc).To(gomega.Equal(int64(1)))
		})
		ginkgo.It("ClusterCountFailureReports", func() {
			gomega.Expect(adapter.ClusterCountFailureReports(ctx, "1").Err()).To(gomega.MatchError("Unknown node 1"))
		})
		ginkgo.It("ClusterSlaves", func() {
			gomega.Expect(adapter.ClusterSlaves(ctx, "1").Err()).To(gomega.MatchError("Unknown node 1"))
		})
	})
}

func testAdapter(resp3 bool) {
	var adapter Cmdable

	ginkgo.BeforeEach(func() {
		if resp3 {
			adapter = adapterresp3
		} else {
			adapter = adapterresp2
		}
		gomega.Expect(adapter.FlushDB(ctx).Err()).NotTo(gomega.HaveOccurred())
		gomega.Expect(adapter.FlushAll(ctx).Err()).NotTo(gomega.HaveOccurred())
	})

	ginkgo.Describe("server", func() {
		ginkgo.It("should Echo", func() {
			echo := adapter.Echo(ctx, "hello")
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			gomega.Expect(echo.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(echo.Val()).To(gomega.Equal("hello"))
		})

		ginkgo.It("should Ping", func() {
			ping := adapter.Ping(ctx)
			gomega.Expect(ping.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ping.Val()).To(gomega.Equal("PONG"))
		})

		ginkgo.It("should Migrate", func() {
			var r *StatusCmd
			if resp3 {
				r = adapter.Migrate(ctx, "127.0.0.1", 6378, "nonkey", 0, 1)
			} else {
				r = adapter.Migrate(ctx, "127.0.0.1", 6356, "nonkey", 0, 1)
			}
			gomega.Expect(r.Err()).To(gomega.BeNil())
			gomega.Expect(r.Val()).To(gomega.Equal("NOKEY"))
		})

		ginkgo.It("should Move", func() {
			gomega.Expect(adapter.Set(ctx, "movekey", "1", 0).Err()).To(gomega.BeNil())
			r := adapter.Move(ctx, "movekey", 1)
			gomega.Expect(r.Err()).To(gomega.BeNil())
			gomega.Expect(r.Val()).To(gomega.BeTrue())
		})

		ginkgo.It("should ClientKill", func() {
			r := adapter.ClientKill(ctx, "1.1.1.1:1111")
			gomega.Expect(r.Err()).To(gomega.MatchError("No such client"))
			gomega.Expect(r.Val()).To(gomega.Equal(""))
		})

		ginkgo.It("should ClientKillByFilter", func() {
			r := adapter.ClientKillByFilter(ctx, "ID", "12039487")
			gomega.Expect(r.Err()).To(gomega.BeNil())
			gomega.Expect(r.Val()).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("should ClientList", func() {
			r := adapter.ClientList(ctx)
			gomega.Expect(r.Err()).To(gomega.BeNil())
			gomega.Expect(r.Val()).NotTo(gomega.Equal(""))
		})

		ginkgo.It("should ClientID", func() {
			err := adapter.ClientID(ctx).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(adapter.ClientID(ctx).Val()).To(gomega.BeNumerically(">=", 0))
		})

		ginkgo.It("should ClientGetName", func() {
			r := adapter.ClientGetName(ctx)
			gomega.Expect(r.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(r.Val()).To(gomega.Equal("rueidis"))
		})

		ginkgo.It("should ConfigGet", func() {
			val, err := adapter.ConfigGet(ctx, "*").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).NotTo(gomega.BeEmpty())
		})

		ginkgo.It("should ConfigResetStat", func() {
			r := adapter.ConfigResetStat(ctx)
			gomega.Expect(r.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(r.Val()).To(gomega.Equal("OK"))
		})

		ginkgo.It("should ConfigSet", func() {
			configGet := adapter.ConfigGet(ctx, "maxmemory")
			gomega.Expect(configGet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(configGet.Val()).To(gomega.HaveLen(1))
			gomega.Expect(configGet.Val()["maxmemory"]).NotTo(gomega.BeEmpty())

			configSet := adapter.ConfigSet(ctx, "maxmemory", configGet.Val()["maxmemory"])
			gomega.Expect(configSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(configSet.Val()).To(gomega.Equal("OK"))
		})

		ginkgo.It("should DBSize", func() {
			size, err := adapter.DBSize(ctx).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("should Info", func() {
			info := adapter.Info(ctx)
			gomega.Expect(info.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(info.Val()).NotTo(gomega.Equal(""))
		})

		ginkgo.It("should Info cpu", func() {
			info := adapter.Info(ctx, "cpu")
			gomega.Expect(info.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(info.Val()).NotTo(gomega.Equal(""))
			gomega.Expect(info.Val()).To(gomega.ContainSubstring(`used_cpu_sys`))
		})

		ginkgo.It("should LastSave", func() {
			lastSave := adapter.LastSave(ctx)
			gomega.Expect(lastSave.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lastSave.Val()).NotTo(gomega.Equal(0))
		})

		ginkgo.It("should Save", func() {
			// workaround for "ERR Background save already in progress"
			gomega.Eventually(func() string {
				return adapter.Save(ctx).Val()
			}, "10s").Should(gomega.Equal("OK"))
		})

		ginkgo.It("should Time", func() {
			tm, err := adapter.Time(ctx).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(tm).To(gomega.BeTemporally("~", time.Now(), 3*time.Second))
		})

		ginkgo.It("should Command", func() {
			cmds, err := adapter.Command(ctx).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(cmds)).To(gomega.BeNumerically(">=", 200))

			cmd := cmds["mget"]
			gomega.Expect(cmd.Name).To(gomega.Equal("mget"))
			gomega.Expect(cmd.Arity).To(gomega.Equal(int64(-2)))
			gomega.Expect(cmd.Flags).To(gomega.ContainElement("readonly"))
			gomega.Expect(cmd.FirstKeyPos).To(gomega.Equal(int64(1)))
			gomega.Expect(cmd.LastKeyPos).To(gomega.Equal(int64(-1)))
			gomega.Expect(cmd.StepCount).To(gomega.Equal(int64(1)))

			cmd = cmds["ping"]
			gomega.Expect(cmd.Name).To(gomega.Equal("ping"))
			gomega.Expect(cmd.Arity).To(gomega.Equal(int64(-1)))
			gomega.Expect(cmd.Flags).To(gomega.ContainElement("fast"))
			gomega.Expect(cmd.FirstKeyPos).To(gomega.Equal(int64(0)))
			gomega.Expect(cmd.LastKeyPos).To(gomega.Equal(int64(0)))
			gomega.Expect(cmd.StepCount).To(gomega.Equal(int64(0)))
		})

		if resp3 {
			ginkgo.It("should return all command names", func() {
				cmdList := adapter.CommandList(ctx, FilterBy{})
				gomega.Expect(cmdList.Err()).NotTo(gomega.HaveOccurred())
				cmdNames := cmdList.Val()

				gomega.Expect(cmdNames).NotTo(gomega.BeEmpty())

				// Assert that some expected commands are present in the list
				gomega.Expect(cmdNames).To(gomega.ContainElement("get"))
				gomega.Expect(cmdNames).To(gomega.ContainElement("set"))
				gomega.Expect(cmdNames).To(gomega.ContainElement("hset"))
			})

			ginkgo.It("should filter commands by module", func() {
				filter := FilterBy{
					Module: "JSON",
				}
				cmdList := adapter.CommandList(ctx, filter)
				gomega.Expect(cmdList.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmdList.Val()).To(gomega.HaveLen(0))
			})

			ginkgo.It("should filter commands by ACL category", func() {

				filter := FilterBy{
					ACLCat: "admin",
				}

				cmdList := adapter.CommandList(ctx, filter)
				gomega.Expect(cmdList.Err()).NotTo(gomega.HaveOccurred())
				cmdNames := cmdList.Val()

				// Assert that the returned list only contains commands from the admin ACL category
				gomega.Expect(len(cmdNames)).To(gomega.BeNumerically(">", 10))
			})

			ginkgo.It("should filter commands by pattern", func() {
				filter := FilterBy{
					Pattern: "*GET*",
				}
				cmdList := adapter.CommandList(ctx, filter)
				gomega.Expect(cmdList.Err()).NotTo(gomega.HaveOccurred())
				cmdNames := cmdList.Val()

				// Assert that the returned list only contains commands that match the given pattern
				gomega.Expect(cmdNames).To(gomega.ContainElement("get"))
				gomega.Expect(cmdNames).To(gomega.ContainElement("getbit"))
				gomega.Expect(cmdNames).To(gomega.ContainElement("getrange"))
				gomega.Expect(cmdNames).NotTo(gomega.ContainElement("set"))
			})

			ginkgo.It("Should CommandGetKeys", func() {
				keys, err := adapter.CommandGetKeys(ctx, "MSET", "a", "b", "c", "d", "e", "f").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(keys).To(gomega.Equal([]string{"a", "c", "e"}))

				keys, err = adapter.CommandGetKeys(ctx, "EVAL", "not consulted", "3", "key1", "key2", "key3", "arg1", "arg2", "arg3", "argN").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(keys).To(gomega.Equal([]string{"key1", "key2", "key3"}))

				keys, err = adapter.CommandGetKeys(ctx, "SORT", "mylist", "ALPHA", "STORE", "outlist").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(keys).To(gomega.Equal([]string{"mylist", "outlist"}))

				_, err = adapter.CommandGetKeys(ctx, "FAKECOMMAND", "arg1", "arg2").Result()
				gomega.Expect(err).To(gomega.HaveOccurred())
				gomega.Expect(err.Error()).To(gomega.Equal("Invalid command specified"))
			})

			ginkgo.It("should CommandGetKeysAndFlags", func() {
				keysAndFlags, err := adapter.CommandGetKeysAndFlags(ctx, "LMOVE", "mylist1", "mylist2", "left", "left").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(keysAndFlags).To(gomega.Equal([]KeyFlags{
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
				gomega.Expect(err).To(gomega.HaveOccurred())
				gomega.Expect(err.Error()).To(gomega.Equal("Invalid command specified"))
			})
		}
	})

	ginkgo.Describe("debugging", func() {
		ginkgo.It("should MemoryUsage", func() {
			err := adapter.MemoryUsage(ctx, "foo").Err()
			gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())

			err = adapter.Set(ctx, "foo", "bar", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			n, err := adapter.MemoryUsage(ctx, "foo").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).NotTo(gomega.BeZero())

			n, err = adapter.MemoryUsage(ctx, "foo", 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).NotTo(gomega.BeZero())
		})
	})

	ginkgo.Describe("keys", func() {
		ginkgo.It("should Del", func() {
			err := adapter.Set(ctx, "key1", "Hello", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.Set(ctx, "key2", "World", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			n, err := adapter.Del(ctx, "key1", "key2", "key3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should Unlink", func() {
			err := adapter.Set(ctx, "key1", "Hello", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.Set(ctx, "key2", "World", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			n, err := adapter.Unlink(ctx, "key1", "key2", "key3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should Dump", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			dump := adapter.Dump(ctx, "key")
			gomega.Expect(dump.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(dump.Val()).NotTo(gomega.BeEmpty())
		})

		ginkgo.It("should Exists", func() {
			set := adapter.Set(ctx, "key1", "Hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			n, err := adapter.Exists(ctx, "key1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(1)))

			n, err = adapter.Exists(ctx, "key2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(0)))

			n, err = adapter.Exists(ctx, "key1", "key2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(1)))

			n, err = adapter.Exists(ctx, "key1", "key1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should Expire", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expire := adapter.Expire(ctx, "key", 10*time.Second)
			gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(expire.Val()).To(gomega.Equal(true))

			ttl := adapter.TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(10 * time.Second))

			set = adapter.Set(ctx, "key", "Hello World", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			ttl = adapter.TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(time.Duration(-1)))

			ttl = adapter.TTL(ctx, "nonexistent_key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(time.Duration(-2)))
		})

		if resp3 {
			ginkgo.It("should ExpireNX", func() {
				set := adapter.Set(ctx, "key", "Hello", 0)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				expire := adapter.ExpireNX(ctx, "key", 10*time.Second)
				gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expire.Val()).To(gomega.Equal(true))

				ttl := adapter.TTL(ctx, "key")
				gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(ttl.Val()).To(gomega.Equal(10 * time.Second))

				expire = adapter.ExpireNX(ctx, "key", 20*time.Second)
				gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expire.Val()).To(gomega.Equal(false))
			})

			ginkgo.It("should ExpireXX", func() {
				set := adapter.Set(ctx, "key", "Hello", 0)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				expire := adapter.ExpireXX(ctx, "key", 10*time.Second)
				gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expire.Val()).To(gomega.Equal(false))

				expire = adapter.ExpireNX(ctx, "key", 10*time.Second)
				gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expire.Val()).To(gomega.Equal(true))

				expire = adapter.ExpireXX(ctx, "key", 20*time.Second)
				gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expire.Val()).To(gomega.Equal(true))
			})

			ginkgo.It("should ExpireGT", func() {
				set := adapter.Set(ctx, "key", "Hello", 5*time.Second)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				expire := adapter.ExpireGT(ctx, "key", 10*time.Second)
				gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expire.Val()).To(gomega.Equal(true))

				expire = adapter.ExpireGT(ctx, "key", 5*time.Second)
				gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expire.Val()).To(gomega.Equal(false))
			})

			ginkgo.It("should ExpireLT", func() {
				set := adapter.Set(ctx, "key", "Hello", 10*time.Second)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				expire := adapter.ExpireLT(ctx, "key", 5*time.Second)
				gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expire.Val()).To(gomega.Equal(true))

				expire = adapter.ExpireLT(ctx, "key", 10*time.Second)
				gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expire.Val()).To(gomega.Equal(false))
			})
		}

		if resp3 {
			ginkgo.It("should ExpireAt", func() {
				set := adapter.Set(ctx, "key", "Hello", 0)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				n, err := adapter.Exists(ctx, "key").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(1)))

				// Check correct expiration time is set in the future
				expireAt := time.Now().Add(time.Minute)
				expireAtCmd := adapter.ExpireAt(ctx, "key", expireAt)
				gomega.Expect(expireAtCmd.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expireAtCmd.Val()).To(gomega.Equal(true))

				timeCmd := adapter.ExpireTime(ctx, "key")
				gomega.Expect(timeCmd.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(timeCmd.Val().Seconds()).To(gomega.BeNumerically("==", expireAt.Unix()))

				ptimeCmd := adapter.PExpireTime(ctx, "key")
				gomega.Expect(ptimeCmd.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(ptimeCmd.Val().Seconds()).To(gomega.BeNumerically("==", expireAt.Unix()))

				// Check correct expiration in the past
				expireAtCmd = adapter.ExpireAt(ctx, "key", time.Now().Add(-time.Hour))
				gomega.Expect(expireAtCmd.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(expireAtCmd.Val()).To(gomega.Equal(true))

				n, err = adapter.Exists(ctx, "key").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(0)))
			})
		}

		ginkgo.It("should Keys", func() {
			mset := adapter.MSet(ctx, "one", "1", "two", "2", "three", "3", "four", "4")
			gomega.Expect(mset.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(mset.Val()).To(gomega.Equal("OK"))

			keys := adapter.Keys(ctx, "*o*")
			gomega.Expect(keys.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(keys.Val()).To(gomega.ConsistOf([]string{"four", "one", "two"}))

			keys = adapter.Keys(ctx, "t??")
			gomega.Expect(keys.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(keys.Val()).To(gomega.Equal([]string{"two"}))

			keys = adapter.Keys(ctx, "*")
			gomega.Expect(keys.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(keys.Val()).To(gomega.ConsistOf([]string{"four", "one", "three", "two"}))
		})

		ginkgo.It("should Object", func() {
			start := time.Now()
			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			refCount := adapter.ObjectRefCount(ctx, "key")
			gomega.Expect(refCount.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(refCount.Val()).To(gomega.Equal(int64(1)))

			err := adapter.ObjectEncoding(ctx, "key").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			idleTime := adapter.ObjectIdleTime(ctx, "key")
			gomega.Expect(idleTime.Err()).NotTo(gomega.HaveOccurred())

			// Redis returned milliseconds/1000, which may cause ObjectIdleTime to be at a critical value,
			// should be +1s to deal with the critical value problem.
			// if too much time (>1s) is used during command execution, it may also cause the test to fail.
			// so the ObjectIdleTime result should be <=now-start+1s
			// link: https://github.com/redis/redis/blob/5b48d900498c85bbf4772c1d466c214439888115/src/object.c#L1265-L1272
			gomega.Expect(idleTime.Val()).To(gomega.BeNumerically("<=", time.Now().Sub(start)+time.Second))
		})

		ginkgo.It("should Persist", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expire := adapter.Expire(ctx, "key", 10*time.Second)
			gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(expire.Val()).To(gomega.Equal(true))

			ttl := adapter.TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(10 * time.Second))

			persist := adapter.Persist(ctx, "key")
			gomega.Expect(persist.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(persist.Val()).To(gomega.Equal(true))

			ttl = adapter.TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val() < 0).To(gomega.Equal(true))
		})

		ginkgo.It("should PExpire", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expiration := 900 * time.Millisecond
			pexpire := adapter.PExpire(ctx, "key", expiration)
			gomega.Expect(pexpire.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pexpire.Val()).To(gomega.Equal(true))

			ttl := adapter.TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(time.Second))

			pttl := adapter.PTTL(ctx, "key")
			gomega.Expect(pttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pttl.Val()).To(gomega.BeNumerically("~", expiration, 100*time.Millisecond))
		})

		ginkgo.It("should PExpireAt", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expiration := 900 * time.Millisecond
			pexpireat := adapter.PExpireAt(ctx, "key", time.Now().Add(expiration))
			gomega.Expect(pexpireat.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pexpireat.Val()).To(gomega.Equal(true))

			ttl := adapter.TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(time.Second))

			pttl := adapter.PTTL(ctx, "key")
			gomega.Expect(pttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pttl.Val()).To(gomega.BeNumerically("~", expiration, 100*time.Millisecond))
		})

		ginkgo.It("should PTTL", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expiration := time.Second
			expire := adapter.Expire(ctx, "key", expiration)
			gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			pttl := adapter.PTTL(ctx, "key")
			gomega.Expect(pttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pttl.Val()).To(gomega.BeNumerically("~", expiration, 100*time.Millisecond))
		})

		ginkgo.It("should RandomKey", func() {
			randomKey := adapter.RandomKey(ctx)
			gomega.Expect(rueidis.IsRedisNil(randomKey.Err())).To(gomega.BeTrue())
			gomega.Expect(randomKey.Val()).To(gomega.Equal(""))

			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			randomKey = adapter.RandomKey(ctx)
			gomega.Expect(randomKey.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(randomKey.Val()).To(gomega.Equal("key"))
		})

		ginkgo.It("should Rename", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			status := adapter.Rename(ctx, "key", "key1")
			gomega.Expect(status.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(status.Val()).To(gomega.Equal("OK"))

			get := adapter.Get(ctx, "key1")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("hello"))
		})

		ginkgo.It("should RenameNX", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			renameNX := adapter.RenameNX(ctx, "key", "key1")
			gomega.Expect(renameNX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(renameNX.Val()).To(gomega.Equal(true))

			get := adapter.Get(ctx, "key1")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("hello"))
		})

		ginkgo.It("should Restore", func() {
			err := adapter.Set(ctx, "key", "hello", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			dump := adapter.Dump(ctx, "key")
			gomega.Expect(dump.Err()).NotTo(gomega.HaveOccurred())

			err = adapter.Del(ctx, "key").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			restore, err := adapter.Restore(ctx, "key", 0, dump.Val()).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(restore).To(gomega.Equal("OK"))

			type_, err := adapter.Type(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(type_).To(gomega.Equal("string"))

			val, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("hello"))
		})

		ginkgo.It("should RestoreReplace", func() {
			err := adapter.Set(ctx, "key", "hello", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			dump := adapter.Dump(ctx, "key")
			gomega.Expect(dump.Err()).NotTo(gomega.HaveOccurred())

			restore, err := adapter.RestoreReplace(ctx, "key", 0, dump.Val()).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(restore).To(gomega.Equal("OK"))

			type_, err := adapter.Type(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(type_).To(gomega.Equal("string"))

			val, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("hello"))
		})

		ginkgo.It("should Sort", func() {
			size, err := adapter.LPush(ctx, "list", "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list", "3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list", "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(3)))

			els, err := adapter.Sort(ctx, "list", Sort{
				Offset: 0,
				Count:  2,
				Order:  "ASC",
				Alpha:  true,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(els).To(gomega.Equal([]string{"1", "2"}))
		})

		ginkgo.It("should Sort By", func() {
			size, err := adapter.LPush(ctx, "list_by", "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list_by", "3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list_by", "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(3)))

			els, err := adapter.Sort(ctx, "list_by", Sort{
				Offset: 0,
				Count:  2,
				By:     "nosort",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(els).To(gomega.Equal([]string{"2", "3"}))
		})

		if resp3 {
			ginkgo.It("should Sort", func() {
				size, err := adapter.LPush(ctx, "list", "1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(size).To(gomega.Equal(int64(1)))

				size, err = adapter.LPush(ctx, "list", "3").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(size).To(gomega.Equal(int64(2)))

				size, err = adapter.LPush(ctx, "list", "2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(size).To(gomega.Equal(int64(3)))

				els, err := adapter.SortRO(ctx, "list", Sort{
					Offset: 0,
					Count:  2,
					Order:  "ASC",
					Alpha:  true,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(els).To(gomega.Equal([]string{"1", "2"}))
			})

			ginkgo.It("should Sort By", func() {
				size, err := adapter.LPush(ctx, "list_by", "1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(size).To(gomega.Equal(int64(1)))

				size, err = adapter.LPush(ctx, "list_by", "3").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(size).To(gomega.Equal(int64(2)))

				size, err = adapter.LPush(ctx, "list_by", "2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(size).To(gomega.Equal(int64(3)))

				els, err := adapter.SortRO(ctx, "list_by", Sort{
					Offset: 0,
					Count:  2,
					By:     "nosort",
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(els).To(gomega.Equal([]string{"2", "3"}))
			})
		}

		ginkgo.It("should Sort Panic", func() {
			gomega.Expect(func() {
				adapter.Sort(ctx, "list", Sort{Order: "PANIC"})
			}).To(gomega.Panic())
		})

		ginkgo.It("should Sort and Get", func() {
			size, err := adapter.LPush(ctx, "list", "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list", "3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list", "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(3)))

			err = adapter.Set(ctx, "object_2", "value2", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			{
				els, err := adapter.Sort(ctx, "list", Sort{
					Get: []string{"object_*"},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(els).To(gomega.Equal([]string{"", "value2", ""}))
			}

			{
				els, err := adapter.SortInterfaces(ctx, "list", Sort{
					Get: []string{"object_*"},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(els).To(gomega.Equal([]any{nil, "value2", nil}))
			}
		})

		ginkgo.It("should Sort and Store", func() {
			size, err := adapter.LPush(ctx, "list", "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list", "3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list", "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(3)))

			n, err := adapter.SortStore(ctx, "list", "list2", Sort{
				Offset: 0,
				Count:  2,
				Order:  "ASC",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(2)))

			els, err := adapter.LRange(ctx, "list2", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(els).To(gomega.Equal([]string{"1", "2"}))
		})

		ginkgo.It("should Touch", func() {
			set1 := adapter.Set(ctx, "touch1", "hello", 0)
			gomega.Expect(set1.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set1.Val()).To(gomega.Equal("OK"))

			set2 := adapter.Set(ctx, "touch2", "hello", 0)
			gomega.Expect(set2.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set2.Val()).To(gomega.Equal("OK"))

			touch := adapter.Touch(ctx, "touch1", "touch2", "touch3")
			gomega.Expect(touch.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(touch.Val()).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should TTL", func() {
			ttl := adapter.TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val() < 0).To(gomega.Equal(true))

			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expire := adapter.Expire(ctx, "key", 60*time.Second)
			gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(expire.Val()).To(gomega.Equal(true))

			ttl = adapter.TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(60 * time.Second))
		})

		ginkgo.It("should Type", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			type_ := adapter.Type(ctx, "key")
			gomega.Expect(type_.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(type_.Val()).To(gomega.Equal("string"))
		})
	})

	ginkgo.Describe("scanning", func() {
		ginkgo.It("should Scan", func() {
			for i := 0; i < 1000; i++ {
				set := adapter.Set(ctx, fmt.Sprintf("key%d", i), "hello", 0)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			}

			keys, cursor, err := adapter.Scan(ctx, 0, "key*", 100).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(keys).NotTo(gomega.BeEmpty())
			gomega.Expect(cursor).NotTo(gomega.BeZero())
		})

		if resp3 {
			ginkgo.It("should ScanType", func() {
				for i := 0; i < 1000; i++ {
					set := adapter.Set(ctx, fmt.Sprintf("key%d", i), "hello", 0)
					gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				}

				keys, cursor, err := adapter.ScanType(ctx, 0, "key*", 100, "string").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(keys).NotTo(gomega.BeEmpty())
				gomega.Expect(cursor).NotTo(gomega.BeZero())
			})
		}

		ginkgo.It("should SScan", func() {
			for i := 0; i < 1000; i++ {
				sadd := adapter.SAdd(ctx, "myset", fmt.Sprintf("member%d", i))
				gomega.Expect(sadd.Err()).NotTo(gomega.HaveOccurred())
			}

			keys, cursor, err := adapter.SScan(ctx, "myset", 0, "member*", 100).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(keys).NotTo(gomega.BeEmpty())
			gomega.Expect(cursor).NotTo(gomega.BeZero())
		})

		ginkgo.It("should HScan", func() {
			for i := 0; i < 1000; i++ {
				sadd := adapter.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
				gomega.Expect(sadd.Err()).NotTo(gomega.HaveOccurred())
			}

			keys, cursor, err := adapter.HScan(ctx, "myhash", 0, "key*", 100).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(keys).NotTo(gomega.BeEmpty())
			gomega.Expect(cursor).NotTo(gomega.BeZero())
		})

		ginkgo.It("should ZScan", func() {
			for i := 0; i < 1000; i++ {
				err := adapter.ZAdd(ctx, "myset", Z{
					Score:  float64(i),
					Member: fmt.Sprintf("member%d", i),
				}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			}

			keys, cursor, err := adapter.ZScan(ctx, "myset", 0, "member*", 100).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(keys).NotTo(gomega.BeEmpty())
			gomega.Expect(cursor).NotTo(gomega.BeZero())
		})
	})

	ginkgo.Describe("strings", func() {
		ginkgo.It("should Append", func() {
			n, err := adapter.Exists(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(0)))

			append := adapter.Append(ctx, "key", "Hello")
			gomega.Expect(append.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(append.Val()).To(gomega.Equal(int64(5)))

			append = adapter.Append(ctx, "key", " World")
			gomega.Expect(append.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(append.Val()).To(gomega.Equal(int64(11)))

			get := adapter.Get(ctx, "key")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("Hello World"))
		})

		ginkgo.It("should BitCount", func() {
			set := adapter.Set(ctx, "key", "foobar", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			bitCount := adapter.BitCount(ctx, "key", nil)
			gomega.Expect(bitCount.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bitCount.Val()).To(gomega.Equal(int64(26)))

			bitCount = adapter.BitCount(ctx, "key", &BitCount{
				Start: 0,
				End:   0,
			})
			gomega.Expect(bitCount.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bitCount.Val()).To(gomega.Equal(int64(4)))

			bitCount = adapter.BitCount(ctx, "key", &BitCount{
				Start: 1,
				End:   1,
			})
			gomega.Expect(bitCount.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bitCount.Val()).To(gomega.Equal(int64(6)))
		})

		ginkgo.It("should BitOpAnd", func() {
			set := adapter.Set(ctx, "key1", "1", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			set = adapter.Set(ctx, "key2", "0", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			bitOpAnd := adapter.BitOpAnd(ctx, "dest", "key1", "key2")
			gomega.Expect(bitOpAnd.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bitOpAnd.Val()).To(gomega.Equal(int64(1)))

			get := adapter.Get(ctx, "dest")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("0"))
		})

		ginkgo.It("should BitOpOr", func() {
			set := adapter.Set(ctx, "key1", "1", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			set = adapter.Set(ctx, "key2", "0", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			bitOpOr := adapter.BitOpOr(ctx, "dest", "key1", "key2")
			gomega.Expect(bitOpOr.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bitOpOr.Val()).To(gomega.Equal(int64(1)))

			get := adapter.Get(ctx, "dest")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("1"))
		})

		ginkgo.It("should BitOpXor", func() {
			set := adapter.Set(ctx, "key1", "\xff", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			set = adapter.Set(ctx, "key2", "\x0f", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			bitOpXor := adapter.BitOpXor(ctx, "dest", "key1", "key2")
			gomega.Expect(bitOpXor.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bitOpXor.Val()).To(gomega.Equal(int64(1)))

			get := adapter.Get(ctx, "dest")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("\xf0"))
		})

		ginkgo.It("should BitOpNot", func() {
			set := adapter.Set(ctx, "key1", "\x00", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			bitOpNot := adapter.BitOpNot(ctx, "dest", "key1")
			gomega.Expect(bitOpNot.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bitOpNot.Val()).To(gomega.Equal(int64(1)))

			get := adapter.Get(ctx, "dest")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("\xff"))
		})

		ginkgo.It("BitPos should panic", func() {
			gomega.Expect(func() {
				adapter.BitPos(ctx, "mykey", 0, 0, 0, 0)
			}).To(gomega.Panic())
		})

		ginkgo.It("should BitPos", func() {
			err := adapter.Set(ctx, "mykey", "\xff\xf0\x00", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			pos, err := adapter.BitPos(ctx, "mykey", 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(12)))

			pos, err = adapter.BitPos(ctx, "mykey", 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(0)))

			pos, err = adapter.BitPos(ctx, "mykey", 0, 2).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(16)))

			pos, err = adapter.BitPos(ctx, "mykey", 1, 2).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(-1)))

			pos, err = adapter.BitPos(ctx, "mykey", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(16)))

			pos, err = adapter.BitPos(ctx, "mykey", 1, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(-1)))

			pos, err = adapter.BitPos(ctx, "mykey", 0, 2, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(-1)))

			pos, err = adapter.BitPos(ctx, "mykey", 0, 0, -3).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(-1)))

			pos, err = adapter.BitPos(ctx, "mykey", 0, 0, 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(-1)))
		})

		if resp3 {
			ginkgo.It("should BitPosSpan", func() {
				err := adapter.Set(ctx, "mykey", "\x00\xff\x00", 0).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				pos, err := adapter.BitPosSpan(ctx, "mykey", 0, 1, 3, "byte").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(pos).To(gomega.Equal(int64(16)))

				pos, err = adapter.BitPosSpan(ctx, "mykey", 0, 1, 3, "bit").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(pos).To(gomega.Equal(int64(1)))
			})
		}

		ginkgo.It("should BitField", func() {
			nn, err := adapter.BitField(ctx, "mykey", "INCRBY", "i5", 100, 1, "GET", "u4", 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(nn).To(gomega.Equal([]int64{1, 0}))
		})

		ginkgo.It("should Decr", func() {
			set := adapter.Set(ctx, "key", "10", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			decr := adapter.Decr(ctx, "key")
			gomega.Expect(decr.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(decr.Val()).To(gomega.Equal(int64(9)))

			set = adapter.Set(ctx, "key", "234293482390480948029348230948", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			decr = adapter.Decr(ctx, "key")
			gomega.Expect(decr.Err()).To(gomega.MatchError("value is not an integer or out of range"))
			gomega.Expect(decr.Val()).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("should DecrBy", func() {
			set := adapter.Set(ctx, "key", "10", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			decrBy := adapter.DecrBy(ctx, "key", 5)
			gomega.Expect(decrBy.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(decrBy.Val()).To(gomega.Equal(int64(5)))
		})

		ginkgo.It("should Get", func() {
			get := adapter.Get(ctx, "_")
			gomega.Expect(rueidis.IsRedisNil(get.Err())).To(gomega.BeTrue())
			gomega.Expect(get.Val()).To(gomega.Equal(""))

			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			get = adapter.Get(ctx, "key")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("hello"))
		})

		ginkgo.It("should GetBit", func() {
			setBit := adapter.SetBit(ctx, "key", 7, 1)
			gomega.Expect(setBit.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(setBit.Val()).To(gomega.Equal(int64(0)))

			getBit := adapter.GetBit(ctx, "key", 0)
			gomega.Expect(getBit.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getBit.Val()).To(gomega.Equal(int64(0)))

			getBit = adapter.GetBit(ctx, "key", 7)
			gomega.Expect(getBit.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getBit.Val()).To(gomega.Equal(int64(1)))

			getBit = adapter.GetBit(ctx, "key", 100)
			gomega.Expect(getBit.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getBit.Val()).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("should GetRange", func() {
			set := adapter.Set(ctx, "key", "This is a string", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			getRange := adapter.GetRange(ctx, "key", 0, 3)
			gomega.Expect(getRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getRange.Val()).To(gomega.Equal("This"))

			getRange = adapter.GetRange(ctx, "key", -3, -1)
			gomega.Expect(getRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getRange.Val()).To(gomega.Equal("ing"))

			getRange = adapter.GetRange(ctx, "key", 0, -1)
			gomega.Expect(getRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getRange.Val()).To(gomega.Equal("This is a string"))

			getRange = adapter.GetRange(ctx, "key", 10, 100)
			gomega.Expect(getRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getRange.Val()).To(gomega.Equal("string"))
		})

		ginkgo.It("should GetSet", func() {
			incr := adapter.Incr(ctx, "key")
			gomega.Expect(incr.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(incr.Val()).To(gomega.Equal(int64(1)))

			getSet := adapter.GetSet(ctx, "key", "0")
			gomega.Expect(getSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getSet.Val()).To(gomega.Equal("1"))

			get := adapter.Get(ctx, "key")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("0"))
		})

		if resp3 {
			ginkgo.It("should GetEX", func() {
				set := adapter.Set(ctx, "key", "value", 100*time.Second)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				ttl := adapter.TTL(ctx, "key")
				gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(ttl.Val()).To(gomega.BeNumerically("~", 100*time.Second, 3*time.Second))

				getEX := adapter.GetEx(ctx, "key", 200*time.Second)
				gomega.Expect(getEX.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(getEX.Val()).To(gomega.Equal("value"))

				ttl = adapter.TTL(ctx, "key")
				gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(ttl.Val()).To(gomega.BeNumerically("~", 200*time.Second, 3*time.Second))
			})

			ginkgo.It("should GetEX 2", func() {
				set := adapter.Set(ctx, "key", "value", 0)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				getEX := adapter.GetEx(ctx, "key", 0)
				gomega.Expect(getEX.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(getEX.Val()).To(gomega.Equal("value"))

				ttl := adapter.TTL(ctx, "key")
				gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(ttl.Val()).To(gomega.Equal(time.Duration(-1)))

				getEX = adapter.GetEx(ctx, "key", 100*time.Millisecond)
				gomega.Expect(getEX.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(getEX.Val()).To(gomega.Equal("value"))

				ttl = adapter.PTTL(ctx, "key")
				gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(ttl.Val()).To(gomega.BeNumerically("~", 100*time.Millisecond, 10*time.Millisecond))
			})

			ginkgo.It("should GetDel", func() {
				set := adapter.Set(ctx, "key", "value", 0)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				getDel := adapter.GetDel(ctx, "key")
				gomega.Expect(getDel.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(getDel.Val()).To(gomega.Equal("value"))

				get := adapter.Get(ctx, "key")
				gomega.Expect(rueidis.IsRedisNil(get.Err())).To(gomega.BeTrue())
			})
		}

		ginkgo.It("should Incr", func() {
			set := adapter.Set(ctx, "key", "10", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			incr := adapter.Incr(ctx, "key")
			gomega.Expect(incr.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(incr.Val()).To(gomega.Equal(int64(11)))

			get := adapter.Get(ctx, "key")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("11"))
		})

		ginkgo.It("should IncrBy", func() {
			set := adapter.Set(ctx, "key", "10", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			incrBy := adapter.IncrBy(ctx, "key", 5)
			gomega.Expect(incrBy.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(incrBy.Val()).To(gomega.Equal(int64(15)))
		})

		ginkgo.It("should IncrByFloat", func() {
			set := adapter.Set(ctx, "key", "10.50", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			incrByFloat := adapter.IncrByFloat(ctx, "key", 0.1)
			gomega.Expect(incrByFloat.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(incrByFloat.Val()).To(gomega.Equal(10.6))

			set = adapter.Set(ctx, "key", "5.0e3", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			incrByFloat = adapter.IncrByFloat(ctx, "key", 2.0e2)
			gomega.Expect(incrByFloat.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(incrByFloat.Val()).To(gomega.Equal(float64(5200)))
		})

		ginkgo.It("should IncrByFloatOverflow", func() {
			incrByFloat := adapter.IncrByFloat(ctx, "key", 996945661)
			gomega.Expect(incrByFloat.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(incrByFloat.Val()).To(gomega.Equal(float64(996945661)))
		})

		ginkgo.It("should MSetMGet", func() {
			mSet := adapter.MSet(ctx, "key1", "hello1", "key2", "hello2")
			gomega.Expect(mSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(mSet.Val()).To(gomega.Equal("OK"))

			mGet := adapter.MGet(ctx, "key1", "key2", "_")
			gomega.Expect(mGet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(mGet.Val()).To(gomega.Equal([]interface{}{"hello1", "hello2", nil}))

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
			gomega.Expect(mSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(mSet.Val()).To(gomega.Equal("OK"))

			mGet = adapter.MGet(ctx, "set1", "set2", "set3", "set4")
			gomega.Expect(mGet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(mGet.Val()).To(gomega.Equal([]interface{}{
				"val1",
				"1024",
				strconv.Itoa(int(2 * time.Millisecond.Nanoseconds())),
				"",
			}))
		})

		ginkgo.It("should scan Mget", func() {
			now := time.Now()

			err := adapter.MSet(ctx, "key1", "hello1", "key2", 123, "time", now.Format(time.RFC3339Nano)).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res := adapter.MGet(ctx, "key1", "key2", "_", "time")
			gomega.Expect(res.Err()).NotTo(gomega.HaveOccurred())

			type data struct {
				Key1 string    `redis:"key1"`
				Key2 int       `redis:"key2"`
				Time TimeValue `redis:"time"`
			}
			var d data
			gomega.Expect(res.Scan(&d)).NotTo(gomega.HaveOccurred())
			gomega.Expect(d.Time.UnixNano()).To(gomega.Equal(now.UnixNano()))
			d.Time.Time = time.Time{}
			gomega.Expect(d).To(gomega.Equal(data{
				Key1: "hello1",
				Key2: 123,
				Time: TimeValue{Time: time.Time{}},
			}))
		})

		ginkgo.It("should MSetNX", func() {
			mSetNX := adapter.MSetNX(ctx, "key1", "hello1", "key2", "hello2")
			gomega.Expect(mSetNX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(mSetNX.Val()).To(gomega.Equal(true))

			mSetNX = adapter.MSetNX(ctx, "key2", "hello1", "key3", "hello2")
			gomega.Expect(mSetNX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(mSetNX.Val()).To(gomega.Equal(false))

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
			gomega.Expect(mSetNX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(mSetNX.Val()).To(gomega.Equal(true))
		})
		ginkgo.It("SetWithArgs should panic wrong mode", func() {
			gomega.Expect(func() {
				adapter.SetArgs(ctx, "key", "hello", SetArgs{Mode: "ANY"})
			}).To(gomega.Panic())
		})

		ginkgo.It("should SetWithArgs with TTL", func() {
			args := SetArgs{
				TTL: 500 * time.Millisecond,
			}
			err := adapter.SetArgs(ctx, "key", "hello", args).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			val, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("hello"))

			gomega.Eventually(func() bool {
				return rueidis.IsRedisNil(adapter.Get(ctx, "key").Err())
			}, "2s", "100ms").Should(gomega.BeTrue())
		})

		if resp3 {
			ginkgo.It("should SetWithArgs with expiration date", func() {
				expireAt := time.Now().AddDate(1, 1, 1)
				args := SetArgs{
					ExpireAt: expireAt,
				}
				err := adapter.SetArgs(ctx, "key", "hello", args).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				val, err := adapter.Get(ctx, "key").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal("hello"))

				// check the key has an expiration date
				// (so a TTL value different of -1)
				ttl := adapter.TTL(ctx, "key")
				gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(ttl.Val()).ToNot(gomega.Equal(-1))
			})

			ginkgo.It("should SetWithArgs with negative expiration date", func() {
				args := SetArgs{
					ExpireAt: time.Now().AddDate(-3, 1, 1),
				}
				// redis accepts a timestamp less than the current date
				// but returns nil when trying to get the key
				err := adapter.SetArgs(ctx, "key", "hello", args).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				val, err := adapter.Get(ctx, "key").Result()
				gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
				gomega.Expect(val).To(gomega.Equal(""))
			})

			ginkgo.It("should SetWithArgs with keepttl", func() {
				// Set with ttl
				argsWithTTL := SetArgs{
					TTL: 5 * time.Second,
				}
				set := adapter.SetArgs(ctx, "key", "hello", argsWithTTL)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Result()).To(gomega.Equal("OK"))

				// Set with keepttl
				argsWithKeepTTL := SetArgs{
					KeepTTL: true,
				}
				set = adapter.SetArgs(ctx, "key", "hello", argsWithKeepTTL)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Result()).To(gomega.Equal("OK"))

				ttl := adapter.TTL(ctx, "key")
				gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
				// set keepttl will Retain the ttl associated with the key
				gomega.Expect(ttl.Val().Nanoseconds()).NotTo(gomega.Equal(-1))
			})
		}

		ginkgo.It("should SetWithArgs with NX mode and key exists", func() {
			err := adapter.Set(ctx, "key", "hello", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			args := SetArgs{
				Mode: "nx",
			}
			val, err := adapter.SetArgs(ctx, "key", "hello", args).Result()
			gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
			gomega.Expect(val).To(gomega.Equal(""))
		})

		ginkgo.It("should SetWithArgs with NX mode and key does not exist", func() {
			args := SetArgs{
				Mode: "nx",
			}
			val, err := adapter.SetArgs(ctx, "key", "hello", args).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("OK"))
		})

		ginkgo.It("should SetWithArgs with expiration, NX mode, and key does not exist", func() {
			args := SetArgs{
				TTL:  500 * time.Millisecond,
				Mode: "nx",
			}
			val, err := adapter.SetArgs(ctx, "key", "hello", args).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("OK"))

			gomega.Eventually(func() bool {
				return rueidis.IsRedisNil(adapter.Get(ctx, "key").Err())
			}, "1s", "100ms").Should(gomega.BeTrue())
		})

		ginkgo.It("should SetWithArgs with expiration, NX mode, and key exists", func() {
			e := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(e.Err()).NotTo(gomega.HaveOccurred())

			args := SetArgs{
				TTL:  500 * time.Millisecond,
				Mode: "nx",
			}
			val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
			gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
			gomega.Expect(val).To(gomega.Equal(""))
		})

		ginkgo.It("should SetWithArgs with XX mode and key does not exist", func() {
			args := SetArgs{
				Mode: "xx",
			}
			val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
			gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
			gomega.Expect(val).To(gomega.Equal(""))
		})

		ginkgo.It("should SetWithArgs with XX mode and key exists", func() {
			e := adapter.Set(ctx, "key", "hello", 0).Err()
			gomega.Expect(e).NotTo(gomega.HaveOccurred())

			args := SetArgs{
				Mode: "xx",
			}
			val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("OK"))
		})

		if resp3 {
			ginkgo.It("should SetWithArgs with XX mode and GET option, and key exists", func() {
				e := adapter.Set(ctx, "key", "hello", 0).Err()
				gomega.Expect(e).NotTo(gomega.HaveOccurred())

				args := SetArgs{
					Mode: "xx",
					Get:  true,
				}
				val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal("hello"))
			})

			ginkgo.It("should SetWithArgs with XX mode and GET option, and key does not exist", func() {
				args := SetArgs{
					Mode: "xx",
					Get:  true,
				}

				val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
				gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
				gomega.Expect(val).To(gomega.Equal(""))
			})

			ginkgo.It("should SetWithArgs with expiration, XX mode, GET option, and key does not exist", func() {
				args := SetArgs{
					TTL:  500 * time.Millisecond,
					Mode: "xx",
					Get:  true,
				}

				val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
				gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
				gomega.Expect(val).To(gomega.Equal(""))
			})

			ginkgo.It("should SetWithArgs with expiration, XX mode, GET option, and key exists", func() {
				e := adapter.Set(ctx, "key", "hello", 0)
				gomega.Expect(e.Err()).NotTo(gomega.HaveOccurred())

				args := SetArgs{
					TTL:  500 * time.Millisecond,
					Mode: "xx",
					Get:  true,
				}

				val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal("hello"))

				gomega.Eventually(func() bool {
					return rueidis.IsRedisNil(adapter.Get(ctx, "key").Err())
				}, "1s", "100ms").Should(gomega.BeTrue())
			})

			ginkgo.It("should SetWithArgs with Get and key does not exist yet", func() {
				args := SetArgs{
					Get: true,
				}

				val, err := adapter.SetArgs(ctx, "key", "hello", args).Result()
				gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
				gomega.Expect(val).To(gomega.Equal(""))
			})

			ginkgo.It("should SetWithArgs with Get and key exists", func() {
				e := adapter.Set(ctx, "key", "hello", 0)
				gomega.Expect(e.Err()).NotTo(gomega.HaveOccurred())

				args := SetArgs{
					Get: true,
				}

				val, err := adapter.SetArgs(ctx, "key", "world", args).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal("hello"))
			})

			ginkgo.It("should Set with keepttl", func() {
				// set with ttl
				set := adapter.Set(ctx, "key", "hello", 5*time.Second)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				// set with keepttl
				set = adapter.Set(ctx, "key", "hello1", KeepTTL)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				ttl := adapter.TTL(ctx, "key")
				gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
				// set keepttl will Retain the ttl associated with the key
				gomega.Expect(ttl.Val().Nanoseconds()).NotTo(gomega.Equal(-1))
			})
		}

		ginkgo.It("should Set with expiration", func() {
			err := adapter.Set(ctx, "key", "hello", 100*time.Millisecond).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			val, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("hello"))

			gomega.Eventually(func() bool {
				return rueidis.IsRedisNil(adapter.Get(ctx, "key").Err())
			}, "1s", "100ms").Should(gomega.BeTrue())
		})

		ginkgo.It("should SetGet", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			get := adapter.Get(ctx, "key")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("hello"))
		})

		ginkgo.It("should SetEX", func() {
			err := adapter.SetEX(ctx, "key", "hello", 1*time.Second).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			val, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("hello"))

			gomega.Eventually(func() bool {
				return rueidis.IsRedisNil(adapter.Get(ctx, "foo").Err())
			}, "2s", "100ms").Should(gomega.BeTrue())
		})

		ginkgo.It("should SetNX", func() {
			setNX := adapter.SetNX(ctx, "key", "hello", 0)
			gomega.Expect(setNX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(setNX.Val()).To(gomega.Equal(true))

			setNX = adapter.SetNX(ctx, "key", "hello2", 0)
			gomega.Expect(setNX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(setNX.Val()).To(gomega.Equal(false))

			get := adapter.Get(ctx, "key")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("hello"))
		})

		ginkgo.It("should SetNX with expiration", func() {
			isSet, err := adapter.SetNX(ctx, "key", "hello", time.Second).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(isSet).To(gomega.Equal(true))

			isSet, err = adapter.SetNX(ctx, "key", "hello2", time.Second).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(isSet).To(gomega.Equal(false))

			val, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("hello"))
		})

		ginkgo.It("should SetNX with expiration 2", func() {
			isSet, err := adapter.SetNX(ctx, "key", "hello", 100*time.Millisecond).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(isSet).To(gomega.Equal(true))

			isSet, err = adapter.SetNX(ctx, "key", "hello2", 100*time.Millisecond).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(isSet).To(gomega.Equal(false))

			val, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("hello"))
		})

		if resp3 {
			ginkgo.It("should SetNX with keepttl", func() {
				isSet, err := adapter.SetNX(ctx, "key", "hello1", KeepTTL).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(isSet).To(gomega.Equal(true))

				ttl := adapter.TTL(ctx, "key")
				gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(ttl.Val().Nanoseconds()).To(gomega.Equal(int64(-1)))
			})
		}

		ginkgo.It("should SetXX", func() {
			isSet, err := adapter.SetXX(ctx, "key", "hello2", 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(isSet).To(gomega.Equal(false))

			err = adapter.Set(ctx, "key", "hello", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			isSet, err = adapter.SetXX(ctx, "key", "hello2", 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(isSet).To(gomega.Equal(true))

			val, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("hello2"))
		})

		ginkgo.It("should SetXX with expiration", func() {
			isSet, err := adapter.SetXX(ctx, "key", "hello2", time.Second).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(isSet).To(gomega.Equal(false))

			err = adapter.Set(ctx, "key", "hello", time.Second).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			isSet, err = adapter.SetXX(ctx, "key", "hello2", time.Second).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(isSet).To(gomega.Equal(true))

			val, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("hello2"))
		})

		ginkgo.It("should SetXX with expiration", func() {
			isSet, err := adapter.SetXX(ctx, "key", "hello2", 100*time.Millisecond).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(isSet).To(gomega.Equal(false))

			err = adapter.Set(ctx, "key", "hello", time.Second).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			isSet, err = adapter.SetXX(ctx, "key", "hello2", 100*time.Millisecond).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(isSet).To(gomega.Equal(true))

			val, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal("hello2"))
		})

		if resp3 {
			ginkgo.It("should SetXX with keepttl", func() {
				isSet, err := adapter.SetXX(ctx, "key", "hello2", time.Second).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(isSet).To(gomega.Equal(false))

				err = adapter.Set(ctx, "key", "hello", time.Second).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				isSet, err = adapter.SetXX(ctx, "key", "hello2", 5*time.Second).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(isSet).To(gomega.Equal(true))

				isSet, err = adapter.SetXX(ctx, "key", "hello3", KeepTTL).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(isSet).To(gomega.Equal(true))

				val, err := adapter.Get(ctx, "key").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal("hello3"))

				// set keepttl will Retain the ttl associated with the key
				ttl, err := adapter.TTL(ctx, "key").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(ttl).NotTo(gomega.Equal(-1))
			})
		}

		ginkgo.It("should SetRange", func() {
			set := adapter.Set(ctx, "key", "Hello World", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			range_ := adapter.SetRange(ctx, "key", 6, "Redis")
			gomega.Expect(range_.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(range_.Val()).To(gomega.Equal(int64(11)))

			get := adapter.Get(ctx, "key")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("Hello Redis"))
		})

		ginkgo.It("should StrLen", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			strLen := adapter.StrLen(ctx, "key")
			gomega.Expect(strLen.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(strLen.Val()).To(gomega.Equal(int64(5)))

			strLen = adapter.StrLen(ctx, "_")
			gomega.Expect(strLen.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(strLen.Val()).To(gomega.Equal(int64(0)))
		})

		if resp3 {
			ginkgo.It("should Copy", func() {
				set := adapter.Set(ctx, "key", "hello", 0)
				gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(set.Val()).To(gomega.Equal("OK"))

				copy := adapter.Copy(ctx, "key", "newKey", 0, false)
				gomega.Expect(copy.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(copy.Val()).To(gomega.Equal(int64(1)))

				// Value is available by both keys now
				getOld := adapter.Get(ctx, "key")
				gomega.Expect(getOld.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(getOld.Val()).To(gomega.Equal("hello"))
				getNew := adapter.Get(ctx, "newKey")
				gomega.Expect(getNew.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(getNew.Val()).To(gomega.Equal("hello"))

				// Overwriting an existing key should not succeed
				overwrite := adapter.Copy(ctx, "newKey", "key", 0, false)
				gomega.Expect(overwrite.Val()).To(gomega.Equal(int64(0)))

				// Overwrite is allowed when replace=rue
				replace := adapter.Copy(ctx, "newKey", "key", 0, true)
				gomega.Expect(replace.Val()).To(gomega.Equal(int64(1)))
			})

			ginkgo.It("should acl dryrun", func() {
				dryRun := adapter.ACLDryRun(ctx, "default", "get", "randomKey")
				gomega.Expect(dryRun.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(dryRun.Val()).To(gomega.Equal("OK"))
			})
		}
	})

	ginkgo.Describe("hashes", func() {
		ginkgo.It("should HDel", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())

			hDel := adapter.HDel(ctx, "hash", "key")
			gomega.Expect(hDel.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hDel.Val()).To(gomega.Equal(int64(1)))

			hDel = adapter.HDel(ctx, "hash", "key")
			gomega.Expect(hDel.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hDel.Val()).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("should HExists", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())

			hExists := adapter.HExists(ctx, "hash", "key")
			gomega.Expect(hExists.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hExists.Val()).To(gomega.Equal(true))

			hExists = adapter.HExists(ctx, "hash", "key1")
			gomega.Expect(hExists.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hExists.Val()).To(gomega.Equal(false))
		})

		ginkgo.It("should HGet", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())

			hGet := adapter.HGet(ctx, "hash", "key")
			gomega.Expect(hGet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hGet.Val()).To(gomega.Equal("hello"))

			hGet = adapter.HGet(ctx, "hash", "key1")
			gomega.Expect(rueidis.IsRedisNil(hGet.Err())).To(gomega.BeTrue())
			gomega.Expect(hGet.Val()).To(gomega.Equal(""))
		})

		ginkgo.It("should HGetAll", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.HSet(ctx, "hash", "key2", "hello2").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			m, err := adapter.HGetAll(ctx, "hash").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(m).To(gomega.Equal(map[string]string{"key1": "hello1", "key2": "hello2"}))
		})

		ginkgo.It("should scan", func() {
			now := time.Now()

			err := adapter.HMSet(ctx, "hash", "key1", "hello1", "key2", 123, "time", now.Format(time.RFC3339Nano)).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res := adapter.HGetAll(ctx, "hash")
			gomega.Expect(res.Err()).NotTo(gomega.HaveOccurred())

			type data struct {
				Key1 string    `redis:"key1"`
				Key2 int       `redis:"key2"`
				Time TimeValue `redis:"time"`
			}
			var d data
			gomega.Expect(res.Scan(&d)).NotTo(gomega.HaveOccurred())
			gomega.Expect(d.Time.UnixNano()).To(gomega.Equal(now.UnixNano()))
			d.Time.Time = time.Time{}
			gomega.Expect(d).To(gomega.Equal(data{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			var d2 data2
			err = adapter.HMGet(ctx, "hash", "key1", "key2", "time").Scan(&d2)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(d2.Key1).To(gomega.Equal("hello2"))
			gomega.Expect(d2.Key2).To(gomega.Equal(200))
			gomega.Expect(d2.Time.Unix()).To(gomega.Equal(now.Unix()))
		})

		ginkgo.It("should HIncrBy", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "5")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())

			hIncrBy := adapter.HIncrBy(ctx, "hash", "key", 1)
			gomega.Expect(hIncrBy.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hIncrBy.Val()).To(gomega.Equal(int64(6)))

			hIncrBy = adapter.HIncrBy(ctx, "hash", "key", -1)
			gomega.Expect(hIncrBy.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hIncrBy.Val()).To(gomega.Equal(int64(5)))

			hIncrBy = adapter.HIncrBy(ctx, "hash", "key", -10)
			gomega.Expect(hIncrBy.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hIncrBy.Val()).To(gomega.Equal(int64(-5)))
		})

		ginkgo.It("should HIncrByFloat", func() {
			hSet := adapter.HSet(ctx, "hash", "field", "10.50")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hSet.Val()).To(gomega.Equal(int64(1)))

			hIncrByFloat := adapter.HIncrByFloat(ctx, "hash", "field", 0.1)
			gomega.Expect(hIncrByFloat.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hIncrByFloat.Val()).To(gomega.Equal(10.6))

			hSet = adapter.HSet(ctx, "hash", "field", "5.0e3")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hSet.Val()).To(gomega.Equal(int64(0)))

			hIncrByFloat = adapter.HIncrByFloat(ctx, "hash", "field", 2.0e2)
			gomega.Expect(hIncrByFloat.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hIncrByFloat.Val()).To(gomega.Equal(float64(5200)))
		})

		ginkgo.It("should HKeys", func() {
			hkeys := adapter.HKeys(ctx, "hash")
			gomega.Expect(hkeys.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hkeys.Val()).To(gomega.Equal([]string{}))

			hset := adapter.HSet(ctx, "hash", "key1", "hello1")
			gomega.Expect(hset.Err()).NotTo(gomega.HaveOccurred())
			hset = adapter.HSet(ctx, "hash", "key2", "hello2")
			gomega.Expect(hset.Err()).NotTo(gomega.HaveOccurred())

			hkeys = adapter.HKeys(ctx, "hash")
			gomega.Expect(hkeys.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hkeys.Val()).To(gomega.Equal([]string{"key1", "key2"}))
		})

		ginkgo.It("should HLen", func() {
			hSet := adapter.HSet(ctx, "hash", "key1", "hello1")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())
			hSet = adapter.HSet(ctx, "hash", "key2", "hello2")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())

			hLen := adapter.HLen(ctx, "hash")
			gomega.Expect(hLen.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hLen.Val()).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should HMGet", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1", "key2", "hello2").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.HMGet(ctx, "hash", "key1", "key2", "_").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]any{"hello1", "hello2", nil}))
		})

		ginkgo.It("should HSet", func() {
			ok, err := adapter.HSet(ctx, "hash", map[string]interface{}{
				"key1": "hello1",
				"key2": "hello2",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(ok).To(gomega.Equal(int64(2)))

			v, err := adapter.HGet(ctx, "hash", "key1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(v).To(gomega.Equal("hello1"))

			v, err = adapter.HGet(ctx, "hash", "key2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(v).To(gomega.Equal("hello2"))

			keys, err := adapter.HKeys(ctx, "hash").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(keys).To(gomega.ConsistOf([]string{"key1", "key2"}))
		})

		ginkgo.It("should HSet", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hSet.Val()).To(gomega.Equal(int64(1)))

			hGet := adapter.HGet(ctx, "hash", "key")
			gomega.Expect(hGet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hGet.Val()).To(gomega.Equal("hello"))

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
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hSet.Val()).To(gomega.Equal(int64(5)))

			hMGet := adapter.HMGet(ctx, "hash", "set1", "set2", "set3", "set4", "set5", "set6", "set7", "set8")
			gomega.Expect(hMGet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hMGet.Val()).To(gomega.Equal([]interface{}{
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
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hSet.Val()).To(gomega.Equal(int64(5)))

			hMGet = adapter.HMGet(ctx, "hash2", "set1", "set6")
			gomega.Expect(hMGet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hMGet.Val()).To(gomega.Equal([]interface{}{
				"val2",
				"val",
			}))
		})

		ginkgo.It("should HSetNX", func() {
			hSetNX := adapter.HSetNX(ctx, "hash", "key", "hello")
			gomega.Expect(hSetNX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hSetNX.Val()).To(gomega.Equal(true))

			hSetNX = adapter.HSetNX(ctx, "hash", "key", "hello")
			gomega.Expect(hSetNX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hSetNX.Val()).To(gomega.Equal(false))

			hGet := adapter.HGet(ctx, "hash", "key")
			gomega.Expect(hGet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hGet.Val()).To(gomega.Equal("hello"))
		})

		ginkgo.It("should HVals", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.HSet(ctx, "hash", "key2", "hello2").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			v, err := adapter.HVals(ctx, "hash").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(v).To(gomega.Equal([]string{"hello1", "hello2"}))

			// TODO
			// var slice []string
			// err = adapter.HVals(ctx, "hash").ScanSlice(&slice)
			// gomega.Expect(err).NotTo(gomega.HaveOccurred())
			// gomega.Expect(slice).To(gomega.Equal([]string{"hello1", "hello2"}))
		})

		if resp3 {
			ginkgo.It("should HRandField", func() {
				err := adapter.HSet(ctx, "hash", "key1", "hello1").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.HSet(ctx, "hash", "key2", "hello2").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				v := adapter.HRandField(ctx, "hash", 1)
				gomega.Expect(v.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(v.Val()).To(gomega.Or(gomega.Equal([]string{"key1"}), gomega.Equal([]string{"key2"})))

				v = adapter.HRandField(ctx, "hash", 0)
				gomega.Expect(v.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(v.Val()).To(gomega.HaveLen(0))

				kv, err := adapter.HRandFieldWithValues(ctx, "hash", -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(kv).To(gomega.Or(
					gomega.Equal([]KeyValue{{Key: "key1", Value: "hello1"}}),
					gomega.Equal([]KeyValue{{Key: "key2", Value: "hello2"}}),
				))
			})
		}
	})

	ginkgo.Describe("hyperloglog", func() {
		ginkgo.It("should PFMerge", func() {
			pfAdd := adapter.PFAdd(ctx, "hll1", "1", "2", "3", "4", "5")
			gomega.Expect(pfAdd.Err()).NotTo(gomega.HaveOccurred())

			pfCount := adapter.PFCount(ctx, "hll1")
			gomega.Expect(pfCount.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pfCount.Val()).To(gomega.Equal(int64(5)))

			pfAdd = adapter.PFAdd(ctx, "hll2", "a", "b", "c", "d", "e")
			gomega.Expect(pfAdd.Err()).NotTo(gomega.HaveOccurred())

			pfMerge := adapter.PFMerge(ctx, "hllMerged", "hll1", "hll2")
			gomega.Expect(pfMerge.Err()).NotTo(gomega.HaveOccurred())

			pfCount = adapter.PFCount(ctx, "hllMerged")
			gomega.Expect(pfCount.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pfCount.Val()).To(gomega.Equal(int64(10)))

			pfCount = adapter.PFCount(ctx, "hll1", "hll2")
			gomega.Expect(pfCount.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pfCount.Val()).To(gomega.Equal(int64(10)))
		})
	})

	ginkgo.Describe("lists", func() {
		ginkgo.It("should BLPop", func() {
			rPush := adapter.RPush(ctx, "list1", "a", "b", "c")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			bLPop := adapter.BLPop(ctx, 0, "list1", "list2")
			gomega.Expect(bLPop.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bLPop.Val()).To(gomega.Equal([]string{"list1", "a"}))
		})

		ginkgo.It("should BLPopBlocks", func() {
			started := make(chan bool)
			done := make(chan bool)
			go func() {
				defer ginkgo.GinkgoRecover()

				started <- true
				bLPop := adapter.BLPop(ctx, 0, "list")
				gomega.Expect(bLPop.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(bLPop.Val()).To(gomega.Equal([]string{"list", "a"}))
				done <- true
			}()
			<-started

			select {
			case <-done:
				ginkgo.Fail("BLPop is not blocked")
			case <-time.After(time.Second):
				// ok
			}

			rPush := adapter.RPush(ctx, "list", "a")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			select {
			case <-done:
				// ok
			case <-time.After(time.Second):
				ginkgo.Fail("BLPop is still blocked")
			}
		})

		ginkgo.It("should BLPop timeout", func() {
			val, err := adapter.BLPop(ctx, time.Second, "list1").Result()
			gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
			gomega.Expect(val).To(gomega.BeNil())

			gomega.Expect(adapter.Ping(ctx).Err()).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("should BRPop", func() {
			rPush := adapter.RPush(ctx, "list1", "a", "b", "c")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			bRPop := adapter.BRPop(ctx, 0, "list1", "list2")
			gomega.Expect(bRPop.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bRPop.Val()).To(gomega.Equal([]string{"list1", "c"}))
		})

		ginkgo.It("should BRPop blocks", func() {
			started := make(chan bool)
			done := make(chan bool)
			go func() {
				defer ginkgo.GinkgoRecover()

				started <- true
				brpop := adapter.BRPop(ctx, 0, "list")
				gomega.Expect(brpop.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(brpop.Val()).To(gomega.Equal([]string{"list", "a"}))
				done <- true
			}()
			<-started

			select {
			case <-done:
				ginkgo.Fail("BRPop is not blocked")
			case <-time.After(time.Second):
				// ok
			}

			rPush := adapter.RPush(ctx, "list", "a")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			select {
			case <-done:
				// ok
			case <-time.After(time.Second):
				ginkgo.Fail("BRPop is still blocked")
				// ok
			}
		})

		ginkgo.It("should BRPopLPush", func() {
			_, err := adapter.BRPopLPush(ctx, "list1", "list2", time.Second).Result()
			gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())

			err = adapter.RPush(ctx, "list1", "a", "b", "c").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			v, err := adapter.BRPopLPush(ctx, "list1", "list2", 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(v).To(gomega.Equal("c"))
		})

		ginkgo.It("should LIndex", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())
			lPush = adapter.LPush(ctx, "list", "Hello")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())

			lIndex := adapter.LIndex(ctx, "list", 0)
			gomega.Expect(lIndex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lIndex.Val()).To(gomega.Equal("Hello"))

			lIndex = adapter.LIndex(ctx, "list", -1)
			gomega.Expect(lIndex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lIndex.Val()).To(gomega.Equal("World"))

			lIndex = adapter.LIndex(ctx, "list", 3)
			gomega.Expect(rueidis.IsRedisNil(lIndex.Err())).To(gomega.BeTrue())
			gomega.Expect(lIndex.Val()).To(gomega.Equal(""))
		})

		ginkgo.It("LInsert should panic", func() {
			gomega.Expect(func() {
				adapter.LInsert(ctx, "list", "ANY", "World", "There")
			}).To(gomega.Panic())
		})

		ginkgo.It("should LInsert", func() {
			rPush := adapter.RPush(ctx, "list", "Hello")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "World")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			lInsert := adapter.LInsert(ctx, "list", "BEFORE", "World", "There")
			gomega.Expect(lInsert.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lInsert.Val()).To(gomega.Equal(int64(3)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"Hello", "There", "World"}))

			lInsert = adapter.LInsert(ctx, "list", "AFTER", "World", "There")
			gomega.Expect(lInsert.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lInsert.Val()).To(gomega.Equal(int64(4)))

			lRange = adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"Hello", "There", "World", "There"}))
		})

		ginkgo.It("should LInsert", func() {
			rPush := adapter.RPush(ctx, "list", "Hello")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "World")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			lInsert := adapter.LInsertBefore(ctx, "list", "World", "There")
			gomega.Expect(lInsert.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lInsert.Val()).To(gomega.Equal(int64(3)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"Hello", "There", "World"}))

			lInsert = adapter.LInsertAfter(ctx, "list", "World", "There")
			gomega.Expect(lInsert.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lInsert.Val()).To(gomega.Equal(int64(4)))

			lRange = adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"Hello", "There", "World", "There"}))
		})

		if resp3 {
			ginkgo.It("should LMPop", func() {
				err := adapter.LPush(ctx, "list1", "one", "two", "three", "four", "five").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.LPush(ctx, "list2", "a", "b", "c", "d", "e").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				key, val, err := adapter.LMPop(ctx, "left", 3, "list1", "list2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(key).To(gomega.Equal("list1"))
				gomega.Expect(val).To(gomega.Equal([]string{"five", "four", "three"}))

				key, val, err = adapter.LMPop(ctx, "right", 3, "list1", "list2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(key).To(gomega.Equal("list1"))
				gomega.Expect(val).To(gomega.Equal([]string{"one", "two"}))

				key, val, err = adapter.LMPop(ctx, "left", 1, "list1", "list2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(key).To(gomega.Equal("list2"))
				gomega.Expect(val).To(gomega.Equal([]string{"e"}))

				key, val, err = adapter.LMPop(ctx, "right", 10, "list1", "list2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(key).To(gomega.Equal("list2"))
				gomega.Expect(val).To(gomega.Equal([]string{"a", "b", "c", "d"}))

				err = adapter.LMPop(ctx, "left", 10, "list1", "list2").Err()
				gomega.Expect(err).To(gomega.Equal(rueidis.Nil))

				err = adapter.Set(ctx, "list3", 1024, 0).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.LMPop(ctx, "left", 10, "list1", "list2", "list3").Err()
				gomega.Expect(err.Error()).To(gomega.Equal("WRONGTYPE Operation against a key holding the wrong kind of value"))

				err = adapter.LMPop(ctx, "right", 0, "list1", "list2").Err()
				gomega.Expect(err).To(gomega.HaveOccurred())
			})

			ginkgo.It("should BLMPop", func() {
				err := adapter.LPush(ctx, "list1", "one", "two", "three", "four", "five").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.LPush(ctx, "list2", "a", "b", "c", "d", "e").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				key, val, err := adapter.BLMPop(ctx, 0, "left", 3, "list1", "list2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(key).To(gomega.Equal("list1"))
				gomega.Expect(val).To(gomega.Equal([]string{"five", "four", "three"}))

				key, val, err = adapter.BLMPop(ctx, 0, "right", 3, "list1", "list2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(key).To(gomega.Equal("list1"))
				gomega.Expect(val).To(gomega.Equal([]string{"one", "two"}))

				key, val, err = adapter.BLMPop(ctx, 0, "left", 1, "list1", "list2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(key).To(gomega.Equal("list2"))
				gomega.Expect(val).To(gomega.Equal([]string{"e"}))

				key, val, err = adapter.BLMPop(ctx, 0, "right", 10, "list1", "list2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(key).To(gomega.Equal("list2"))
				gomega.Expect(val).To(gomega.Equal([]string{"a", "b", "c", "d"}))

			})

			ginkgo.It("should BLMPopBlocks", func() {
				started := make(chan bool)
				done := make(chan bool)
				go func() {
					defer ginkgo.GinkgoRecover()

					started <- true
					key, val, err := adapter.BLMPop(ctx, 0, "left", 1, "list_list").Result()
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					gomega.Expect(key).To(gomega.Equal("list_list"))
					gomega.Expect(val).To(gomega.Equal([]string{"a"}))
					done <- true
				}()
				<-started

				select {
				case <-done:
					ginkgo.Fail("BLMPop is not blocked")
				case <-time.After(time.Second):
					// ok
				}

				_, err := adapter.LPush(ctx, "list_list", "a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				select {
				case <-done:
					// ok
				case <-time.After(time.Second):
					ginkgo.Fail("BLMPop is still blocked")
				}
			})

			ginkgo.It("should BLMPop timeout", func() {
				_, val, err := adapter.BLMPop(ctx, time.Second, "left", 1, "list1").Result()
				gomega.Expect(err).To(gomega.Equal(rueidis.Nil))
				gomega.Expect(val).To(gomega.BeNil())

				gomega.Expect(adapter.Ping(ctx).Err()).NotTo(gomega.HaveOccurred())
			})
		}

		ginkgo.It("should LLen", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())
			lPush = adapter.LPush(ctx, "list", "Hello")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())

			lLen := adapter.LLen(ctx, "list")
			gomega.Expect(lLen.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lLen.Val()).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should LPop", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			lPop := adapter.LPop(ctx, "list")
			gomega.Expect(lPop.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lPop.Val()).To(gomega.Equal("one"))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"two", "three"}))
		})

		if resp3 {
			ginkgo.It("should LPopCount", func() {
				rPush := adapter.RPush(ctx, "list", "one")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "two")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "three")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "four")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

				lPopCount := adapter.LPopCount(ctx, "list", 2)
				gomega.Expect(lPopCount.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lPopCount.Val()).To(gomega.Equal([]string{"one", "two"}))

				lRange := adapter.LRange(ctx, "list", 0, -1)
				gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"three", "four"}))
			})

			ginkgo.It("should LPos", func() {
				rPush := adapter.RPush(ctx, "list", "a")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "b")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "c")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "b")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

				lPos := adapter.LPos(ctx, "list", "b", LPosArgs{})
				gomega.Expect(lPos.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lPos.Val()).To(gomega.Equal(int64(1)))

				lPos = adapter.LPos(ctx, "list", "b", LPosArgs{Rank: 2})
				gomega.Expect(lPos.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lPos.Val()).To(gomega.Equal(int64(3)))

				lPos = adapter.LPos(ctx, "list", "b", LPosArgs{Rank: -2})
				gomega.Expect(lPos.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lPos.Val()).To(gomega.Equal(int64(1)))

				lPos = adapter.LPos(ctx, "list", "b", LPosArgs{Rank: 2, MaxLen: 1})
				gomega.Expect(rueidis.IsRedisNil(lPos.Err())).To(gomega.BeTrue())

				lPos = adapter.LPos(ctx, "list", "z", LPosArgs{})
				gomega.Expect(rueidis.IsRedisNil(lPos.Err())).To(gomega.BeTrue())
			})

			ginkgo.It("should LPosCount", func() {
				rPush := adapter.RPush(ctx, "list", "a")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "b")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "c")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				rPush = adapter.RPush(ctx, "list", "b")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

				lPos := adapter.LPosCount(ctx, "list", "b", 2, LPosArgs{})
				gomega.Expect(lPos.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lPos.Val()).To(gomega.Equal([]int64{1, 3}))

				lPos = adapter.LPosCount(ctx, "list", "b", 2, LPosArgs{Rank: 2})
				gomega.Expect(lPos.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lPos.Val()).To(gomega.Equal([]int64{3}))

				lPos = adapter.LPosCount(ctx, "list", "b", 1, LPosArgs{Rank: 1, MaxLen: 1})
				gomega.Expect(lPos.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lPos.Val()).To(gomega.Equal([]int64{}))

				lPos = adapter.LPosCount(ctx, "list", "b", 1, LPosArgs{Rank: 1, MaxLen: 0})
				gomega.Expect(lPos.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lPos.Val()).To(gomega.Equal([]int64{1}))
			})
		}

		ginkgo.It("should LPush", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())
			lPush = adapter.LPush(ctx, "list", "Hello")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"Hello", "World"}))
		})

		ginkgo.It("should LPushX", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())

			lPushX := adapter.LPushX(ctx, "list", "Hello")
			gomega.Expect(lPushX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lPushX.Val()).To(gomega.Equal(int64(2)))

			lPush = adapter.LPush(ctx, "list1", "three")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lPush.Val()).To(gomega.Equal(int64(1)))

			lPushX = adapter.LPushX(ctx, "list1", "two", "one")
			gomega.Expect(lPushX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lPushX.Val()).To(gomega.Equal(int64(3)))

			lPushX = adapter.LPushX(ctx, "list2", "Hello")
			gomega.Expect(lPushX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lPushX.Val()).To(gomega.Equal(int64(0)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"Hello", "World"}))

			lRange = adapter.LRange(ctx, "list1", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one", "two", "three"}))

			lRange = adapter.LRange(ctx, "list2", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{}))
		})

		ginkgo.It("should LRange", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			lRange := adapter.LRange(ctx, "list", 0, 0)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one"}))

			lRange = adapter.LRange(ctx, "list", -3, 2)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one", "two", "three"}))

			lRange = adapter.LRange(ctx, "list", -100, 100)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one", "two", "three"}))

			lRange = adapter.LRange(ctx, "list", 5, 10)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{}))
		})

		ginkgo.It("should LRem", func() {
			rPush := adapter.RPush(ctx, "list", "hello")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "hello")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "key")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "hello")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			lRem := adapter.LRem(ctx, "list", -2, "hello")
			gomega.Expect(lRem.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRem.Val()).To(gomega.Equal(int64(2)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"hello", "key"}))
		})

		ginkgo.It("should LSet", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			lSet := adapter.LSet(ctx, "list", 0, "four")
			gomega.Expect(lSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lSet.Val()).To(gomega.Equal("OK"))

			lSet = adapter.LSet(ctx, "list", -2, "five")
			gomega.Expect(lSet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lSet.Val()).To(gomega.Equal("OK"))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"four", "five", "three"}))
		})

		ginkgo.It("should LTrim", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			lTrim := adapter.LTrim(ctx, "list", 1, -1)
			gomega.Expect(lTrim.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lTrim.Val()).To(gomega.Equal("OK"))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"two", "three"}))
		})

		ginkgo.It("should RPop", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			rPop := adapter.RPop(ctx, "list")
			gomega.Expect(rPop.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(rPop.Val()).To(gomega.Equal("three"))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one", "two"}))
		})

		if resp3 {
			ginkgo.It("should RPopCount", func() {
				rPush := adapter.RPush(ctx, "list", "one", "two", "three", "four")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(rPush.Val()).To(gomega.Equal(int64(4)))

				rPopCount := adapter.RPopCount(ctx, "list", 2)
				gomega.Expect(rPopCount.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(rPopCount.Val()).To(gomega.Equal([]string{"four", "three"}))

				lRange := adapter.LRange(ctx, "list", 0, -1)
				gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one", "two"}))
			})
		}

		ginkgo.It("should RPopLPush", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			rPopLPush := adapter.RPopLPush(ctx, "list", "list2")
			gomega.Expect(rPopLPush.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(rPopLPush.Val()).To(gomega.Equal("three"))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one", "two"}))

			lRange = adapter.LRange(ctx, "list2", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"three"}))
		})

		ginkgo.It("should RPush", func() {
			rPush := adapter.RPush(ctx, "list", "Hello")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(rPush.Val()).To(gomega.Equal(int64(1)))

			rPush = adapter.RPush(ctx, "list", "World")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(rPush.Val()).To(gomega.Equal(int64(2)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"Hello", "World"}))
		})

		ginkgo.It("should RPushX", func() {
			rPush := adapter.RPush(ctx, "list", "Hello")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(rPush.Val()).To(gomega.Equal(int64(1)))

			rPushX := adapter.RPushX(ctx, "list", "World")
			gomega.Expect(rPushX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(rPushX.Val()).To(gomega.Equal(int64(2)))

			rPush = adapter.RPush(ctx, "list1", "one")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(rPush.Val()).To(gomega.Equal(int64(1)))

			rPushX = adapter.RPushX(ctx, "list1", "two", "three")
			gomega.Expect(rPushX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(rPushX.Val()).To(gomega.Equal(int64(3)))

			rPushX = adapter.RPushX(ctx, "list2", "World")
			gomega.Expect(rPushX.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(rPushX.Val()).To(gomega.Equal(int64(0)))

			lRange := adapter.LRange(ctx, "list", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"Hello", "World"}))

			lRange = adapter.LRange(ctx, "list1", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one", "two", "three"}))

			lRange = adapter.LRange(ctx, "list2", 0, -1)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{}))
		})

		if resp3 {
			ginkgo.It("should LMove", func() {
				rPush := adapter.RPush(ctx, "lmove1", "ichi")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(rPush.Val()).To(gomega.Equal(int64(1)))

				rPush = adapter.RPush(ctx, "lmove1", "ni")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(rPush.Val()).To(gomega.Equal(int64(2)))

				rPush = adapter.RPush(ctx, "lmove1", "san")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(rPush.Val()).To(gomega.Equal(int64(3)))

				lMove := adapter.LMove(ctx, "lmove1", "lmove2", "RIGHT", "LEFT")
				gomega.Expect(lMove.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lMove.Val()).To(gomega.Equal("san"))

				lRange := adapter.LRange(ctx, "lmove2", 0, -1)
				gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"san"}))
			})

			ginkgo.It("should BLMove", func() {
				rPush := adapter.RPush(ctx, "blmove1", "ichi")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(rPush.Val()).To(gomega.Equal(int64(1)))

				rPush = adapter.RPush(ctx, "blmove1", "ni")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(rPush.Val()).To(gomega.Equal(int64(2)))

				rPush = adapter.RPush(ctx, "blmove1", "san")
				gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(rPush.Val()).To(gomega.Equal(int64(3)))

				blMove := adapter.BLMove(ctx, "blmove1", "blmove2", "RIGHT", "LEFT", time.Second)
				gomega.Expect(blMove.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(blMove.Val()).To(gomega.Equal("san"))

				lRange := adapter.LRange(ctx, "blmove2", 0, -1)
				gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"san"}))
			})
		}
	})

	ginkgo.Describe("sets", func() {
		ginkgo.It("should SAdd", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sAdd.Val()).To(gomega.Equal(int64(1)))

			sAdd = adapter.SAdd(ctx, "set", "World")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sAdd.Val()).To(gomega.Equal(int64(1)))

			sAdd = adapter.SAdd(ctx, "set", "World")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sAdd.Val()).To(gomega.Equal(int64(0)))

			sMembers := adapter.SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.ConsistOf([]string{"Hello", "World"}))
		})

		ginkgo.It("should SAdd strings", func() {
			set := []string{"Hello", "World", "World"}
			sAdd := adapter.SAdd(ctx, "set", set)
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sAdd.Val()).To(gomega.Equal(int64(2)))

			sMembers := adapter.SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.ConsistOf([]string{"Hello", "World"}))
		})

		ginkgo.It("should SCard", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sAdd.Val()).To(gomega.Equal(int64(1)))

			sAdd = adapter.SAdd(ctx, "set", "World")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sAdd.Val()).To(gomega.Equal(int64(1)))

			sCard := adapter.SCard(ctx, "set")
			gomega.Expect(sCard.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sCard.Val()).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should SDiff", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sDiff := adapter.SDiff(ctx, "set1", "set2")
			gomega.Expect(sDiff.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sDiff.Val()).To(gomega.ConsistOf([]string{"a", "b"}))
		})

		ginkgo.It("should SDiffStore", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sDiffStore := adapter.SDiffStore(ctx, "set", "set1", "set2")
			gomega.Expect(sDiffStore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sDiffStore.Val()).To(gomega.Equal(int64(2)))

			sMembers := adapter.SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.ConsistOf([]string{"a", "b"}))
		})

		ginkgo.It("should SInter", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sInter := adapter.SInter(ctx, "set1", "set2")
			gomega.Expect(sInter.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sInter.Val()).To(gomega.Equal([]string{"c"}))
		})

		if resp3 {
			ginkgo.It("should SInterCard", func() {
				sAdd := adapter.SAdd(ctx, "set1", "a")
				gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
				sAdd = adapter.SAdd(ctx, "set1", "b")
				gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
				sAdd = adapter.SAdd(ctx, "set1", "c")
				gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

				sAdd = adapter.SAdd(ctx, "set2", "b")
				gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
				sAdd = adapter.SAdd(ctx, "set2", "c")
				gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
				sAdd = adapter.SAdd(ctx, "set2", "d")
				gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
				sAdd = adapter.SAdd(ctx, "set2", "e")
				gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
				// limit 0 means no limit,see https://redis.io/commands/sintercard/ for more details
				sInterCard := adapter.SInterCard(ctx, 0, "set1", "set2")
				gomega.Expect(sInterCard.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(sInterCard.Val()).To(gomega.Equal(int64(2)))

				sInterCard = adapter.SInterCard(ctx, 1, "set1", "set2")
				gomega.Expect(sInterCard.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(sInterCard.Val()).To(gomega.Equal(int64(1)))

				sInterCard = adapter.SInterCard(ctx, 3, "set1", "set2")
				gomega.Expect(sInterCard.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(sInterCard.Val()).To(gomega.Equal(int64(2)))
			})
		}

		ginkgo.It("should SInterStore", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sInterStore := adapter.SInterStore(ctx, "set", "set1", "set2")
			gomega.Expect(sInterStore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sInterStore.Val()).To(gomega.Equal(int64(1)))

			sMembers := adapter.SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.Equal([]string{"c"}))
		})

		ginkgo.It("should IsMember", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sIsMember := adapter.SIsMember(ctx, "set", "one")
			gomega.Expect(sIsMember.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sIsMember.Val()).To(gomega.Equal(true))

			sIsMember = adapter.SIsMember(ctx, "set", "two")
			gomega.Expect(sIsMember.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sIsMember.Val()).To(gomega.Equal(false))
		})

		if resp3 {
			ginkgo.It("should SMIsMember", func() {
				sAdd := adapter.SAdd(ctx, "set", "one")
				gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

				sMIsMember := adapter.SMIsMember(ctx, "set", "one", "two")
				gomega.Expect(sMIsMember.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(sMIsMember.Val()).To(gomega.Equal([]bool{true, false}))
			})
		}

		ginkgo.It("should SMembers", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "World")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sMembers := adapter.SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.ConsistOf([]string{"Hello", "World"}))
		})

		ginkgo.It("should SMembersMap", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "World")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sMembersMap := adapter.SMembersMap(ctx, "set")
			gomega.Expect(sMembersMap.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembersMap.Val()).To(gomega.Equal(map[string]struct{}{"Hello": {}, "World": {}}))
		})

		ginkgo.It("should SMove", func() {
			sAdd := adapter.SAdd(ctx, "set1", "one")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "two")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "three")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sMove := adapter.SMove(ctx, "set1", "set2", "two")
			gomega.Expect(sMove.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMove.Val()).To(gomega.Equal(true))

			sMembers := adapter.SMembers(ctx, "set1")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.Equal([]string{"one"}))

			sMembers = adapter.SMembers(ctx, "set2")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.ConsistOf([]string{"three", "two"}))
		})

		ginkgo.It("should SPop", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "two")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "three")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sPop := adapter.SPop(ctx, "set")
			gomega.Expect(sPop.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sPop.Val()).NotTo(gomega.Equal(""))

			sMembers := adapter.SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.HaveLen(2))
		})

		ginkgo.It("should SPopN", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "two")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "three")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "four")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sPopN := adapter.SPopN(ctx, "set", 1)
			gomega.Expect(sPopN.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sPopN.Val()).NotTo(gomega.Equal([]string{""}))

			sMembers := adapter.SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.HaveLen(3))

			sPopN = adapter.SPopN(ctx, "set", 4)
			gomega.Expect(sPopN.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sPopN.Val()).To(gomega.HaveLen(3))

			sMembers = adapter.SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.HaveLen(0))
		})

		ginkgo.It("should SRandMember and SRandMemberN", func() {
			err := adapter.SAdd(ctx, "set", "one").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.SAdd(ctx, "set", "two").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.SAdd(ctx, "set", "three").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			members, err := adapter.SMembers(ctx, "set").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(members).To(gomega.HaveLen(3))

			member, err := adapter.SRandMember(ctx, "set").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(member).NotTo(gomega.Equal(""))

			members, err = adapter.SRandMemberN(ctx, "set", 2).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(members).To(gomega.HaveLen(2))
		})

		ginkgo.It("should SRem", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "two")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "three")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sRem := adapter.SRem(ctx, "set", "one")
			gomega.Expect(sRem.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sRem.Val()).To(gomega.Equal(int64(1)))

			sRem = adapter.SRem(ctx, "set", "four")
			gomega.Expect(sRem.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sRem.Val()).To(gomega.Equal(int64(0)))

			sMembers := adapter.SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.ConsistOf([]string{"three", "two"}))
		})

		ginkgo.It("should SUnion", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sUnion := adapter.SUnion(ctx, "set1", "set2")
			gomega.Expect(sUnion.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sUnion.Val()).To(gomega.HaveLen(5))
		})

		ginkgo.It("should SUnionStore", func() {
			sAdd := adapter.SAdd(ctx, "set1", "a")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "b")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set1", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sAdd = adapter.SAdd(ctx, "set2", "c")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "d")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set2", "e")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sUnionStore := adapter.SUnionStore(ctx, "set", "set1", "set2")
			gomega.Expect(sUnionStore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sUnionStore.Val()).To(gomega.Equal(int64(5)))

			sMembers := adapter.SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.HaveLen(5))
		})
	})

	ginkgo.Describe("sorted sets", func() {
		ginkgo.It("should BZPopMax", func() {
			err := adapter.ZAdd(ctx, "zset1", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{
				Score:  3,
				Member: "three",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			member, err := adapter.BZPopMax(ctx, 0, "zset1", "zset2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(member).To(gomega.Equal(ZWithKey{
				Z: Z{
					Score:  3,
					Member: "three",
				},
				Key: "zset1",
			}))
		})

		ginkgo.It("should BZPopMax blocks", func() {
			started := make(chan bool)
			done := make(chan bool)
			go func() {
				defer ginkgo.GinkgoRecover()

				started <- true
				bZPopMax := adapter.BZPopMax(ctx, 0, "zset")
				gomega.Expect(bZPopMax.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(bZPopMax.Val()).To(gomega.Equal(ZWithKey{
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
				ginkgo.Fail("BZPopMax is not blocked")
			case <-time.After(time.Second):
				// ok
			}

			zAdd := adapter.ZAdd(ctx, "zset", Z{
				Member: "a",
				Score:  1,
			})
			gomega.Expect(zAdd.Err()).NotTo(gomega.HaveOccurred())

			select {
			case <-done:
				// ok
			case <-time.After(time.Second):
				ginkgo.Fail("BZPopMax is still blocked")
			}
		})

		ginkgo.It("should BZPopMax timeout", func() {
			_, err := adapter.BZPopMax(ctx, time.Second, "zset1").Result()
			gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())

			gomega.Expect(adapter.Ping(ctx).Err()).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("should BZPopMin", func() {
			err := adapter.ZAdd(ctx, "zset1", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{
				Score:  3,
				Member: "three",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			member, err := adapter.BZPopMin(ctx, 0, "zset1", "zset2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(member).To(gomega.Equal(ZWithKey{
				Z: Z{
					Score:  1,
					Member: "one",
				},
				Key: "zset1",
			}))
		})

		ginkgo.It("should BZPopMin blocks", func() {
			started := make(chan bool)
			done := make(chan bool)
			go func() {
				defer ginkgo.GinkgoRecover()

				started <- true
				bZPopMin := adapter.BZPopMin(ctx, 0, "zset")
				gomega.Expect(bZPopMin.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(bZPopMin.Val()).To(gomega.Equal(ZWithKey{
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
				ginkgo.Fail("BZPopMin is not blocked")
			case <-time.After(time.Second):
				// ok
			}

			zAdd := adapter.ZAdd(ctx, "zset", Z{
				Member: "a",
				Score:  1,
			})
			gomega.Expect(zAdd.Err()).NotTo(gomega.HaveOccurred())

			select {
			case <-done:
				// ok
			case <-time.After(time.Second):
				ginkgo.Fail("BZPopMin is still blocked")
			}
		})

		ginkgo.It("should BZPopMin timeout", func() {
			_, err := adapter.BZPopMin(ctx, time.Second, "zset1").Result()
			gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())

			gomega.Expect(adapter.Ping(ctx).Err()).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("should ZAdd", func() {
			added, err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "uno",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "two",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(0)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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

		ginkgo.It("should ZAdd bytes", func() {
			added, err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "uno",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(1)))

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "two",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(0)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			ginkgo.It("should ZAddArgs", func() {
				// Test only the GT+LT options.
				added, err := adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					GT:      true,
					Members: []Z{{Score: 1, Member: "one"}},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(1)))

				vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 1, Member: "one"}}))

				added, err = adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					GT:      true,
					Members: []Z{{Score: 2, Member: "one"}},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 2, Member: "one"}}))

				added, err = adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					LT:      true,
					Members: []Z{{Score: 1, Member: "one"}},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 1, Member: "one"}}))
			})
		}

		if resp3 {
			ginkgo.It("should ZAddArgsLT", func() {
				added, err := adapter.ZAddLT(ctx, "zset", Z{
					Score:  2,
					Member: "one",
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(1)))

				vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 2, Member: "one"}}))

				added, err = adapter.ZAddLT(ctx, "zset", Z{
					Score:  3,
					Member: "one",
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 2, Member: "one"}}))

				added, err = adapter.ZAddLT(ctx, "zset", Z{
					Score:  1,
					Member: "one",
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 1, Member: "one"}}))
			})

			ginkgo.It("should ZAddArgsGT", func() {
				added, err := adapter.ZAddGT(ctx, "zset", Z{
					Score:  2,
					Member: "one",
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(1)))

				vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 2, Member: "one"}}))

				added, err = adapter.ZAddGT(ctx, "zset", Z{
					Score:  3,
					Member: "one",
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 3, Member: "one"}}))

				added, err = adapter.ZAddGT(ctx, "zset", Z{
					Score:  1,
					Member: "one",
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(0)))

				vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 3, Member: "one"}}))
			})
		}

		ginkgo.It("should ZAddNX", func() {
			added, err := adapter.ZAddNX(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(1)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 1, Member: "one"}}))

			added, err = adapter.ZAddNX(ctx, "zset", Z{
				Score:  2,
				Member: "one",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(0)))

			vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 1, Member: "one"}}))
		})

		ginkgo.It("should ZAddXX", func() {
			added, err := adapter.ZAddXX(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(0)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.BeEmpty())

			added, err = adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(1)))

			added, err = adapter.ZAddXX(ctx, "zset", Z{
				Score:  2,
				Member: "one",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(0)))

			vals, err = adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 2, Member: "one"}}))
		})

		ginkgo.It("should ZCard", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			card, err := adapter.ZCard(ctx, "zset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(card).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should ZCount", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			count, err := adapter.ZCount(ctx, "zset", "-inf", "+inf").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(count).To(gomega.Equal(int64(3)))

			count, err = adapter.ZCount(ctx, "zset", "(1", "3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(count).To(gomega.Equal(int64(2)))

			count, err = adapter.ZLexCount(ctx, "zset", "-", "+").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(count).To(gomega.Equal(int64(3)))
		})

		ginkgo.It("should ZIncrBy", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			n, err := adapter.ZIncrBy(ctx, "zset", 2, "one").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(float64(3)))

			val, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "one",
			}}))
		})

		ginkgo.It("should ZInterStore", func() {
			err := adapter.ZAdd(ctx, "zset1", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset3", Z{Score: 3, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			n, err := adapter.ZInterStore(ctx, "out", ZStore{
				Keys:    []string{"zset1", "zset2"},
				Weights: []int64{2, 3},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(2)))

			vals, err := adapter.ZRangeWithScores(ctx, "out", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
				Score:  5,
				Member: "one",
			}, {
				Score:  10,
				Member: "two",
			}}))
		})

		if resp3 {
			ginkgo.It("should ZMScore", func() {
				zmScore := adapter.ZMScore(ctx, "zset", "one", "three")
				gomega.Expect(zmScore.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(zmScore.Val()).To(gomega.HaveLen(2))
				gomega.Expect(zmScore.Val()[0]).To(gomega.Equal(float64(0)))

				err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				zmScore = adapter.ZMScore(ctx, "zset", "one", "three")
				gomega.Expect(zmScore.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(zmScore.Val()).To(gomega.HaveLen(2))
				gomega.Expect(zmScore.Val()[0]).To(gomega.Equal(float64(1)))

				zmScore = adapter.ZMScore(ctx, "zset", "four")
				gomega.Expect(zmScore.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(zmScore.Val()).To(gomega.HaveLen(1))

				zmScore = adapter.ZMScore(ctx, "zset", "four", "one")
				gomega.Expect(zmScore.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(zmScore.Val()).To(gomega.HaveLen(2))
			})
		}

		ginkgo.It("should ZPopMax", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			members, err := adapter.ZPopMax(ctx, "zset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(members).To(gomega.Equal([]Z{{
				Score:  3,
				Member: "three",
			}}))

			// adding back 3
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			members, err = adapter.ZPopMax(ctx, "zset", 2).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(members).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			members, err = adapter.ZPopMax(ctx, "zset", 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(members).To(gomega.Equal([]Z{{
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

		ginkgo.It("should ZPopMin", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			members, err := adapter.ZPopMin(ctx, "zset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(members).To(gomega.Equal([]Z{{
				Score:  1,
				Member: "one",
			}}))

			// adding back 1
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			members, err = adapter.ZPopMin(ctx, "zset", 2).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(members).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			members, err = adapter.ZPopMin(ctx, "zset", 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(members).To(gomega.Equal([]Z{{
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

		ginkgo.It("should ZRange", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRange := adapter.ZRange(ctx, "zset", 0, -1)
			gomega.Expect(zRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRange.Val()).To(gomega.Equal([]string{"one", "two", "three"}))

			zRange = adapter.ZRange(ctx, "zset", 2, 3)
			gomega.Expect(zRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRange.Val()).To(gomega.Equal([]string{"three"}))

			zRange = adapter.ZRange(ctx, "zset", -2, -1)
			gomega.Expect(zRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRange.Val()).To(gomega.Equal([]string{"two", "three"}))
		})

		ginkgo.It("should ZRangeWithScores", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 3, Member: "three"}}))

			vals, err = adapter.ZRangeWithScores(ctx, "zset", -2, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))
		})

		if resp3 {
			ginkgo.It("should ZRangeArgs", func() {
				added, err := adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					Members: []Z{
						{Score: 1, Member: "one"},
						{Score: 2, Member: "two"},
						{Score: 3, Member: "three"},
						{Score: 4, Member: "four"},
					},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(4)))

				zRange, err := adapter.ZRangeArgs(ctx, ZRangeArgs{
					Key:     "zset",
					Start:   1,
					Stop:    4,
					ByScore: true,
					Rev:     true,
					Offset:  1,
					Count:   2,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRange).To(gomega.Equal([]string{"three", "two"}))

				zRange, err = adapter.ZRangeArgs(ctx, ZRangeArgs{
					Key:    "zset",
					Start:  "-",
					Stop:   "+",
					ByLex:  true,
					Rev:    true,
					Offset: 2,
					Count:  2,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRange).To(gomega.Equal([]string{"two", "one"}))

				zRange, err = adapter.ZRangeArgs(ctx, ZRangeArgs{
					Key:     "zset",
					Start:   "(1",
					Stop:    "(4",
					ByScore: true,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRange).To(gomega.Equal([]string{"two", "three"}))

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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(zSlice).To(gomega.Equal([]Z{
					{Score: 3, Member: "three"},
					{Score: 2, Member: "two"},
				}))
			})
		}

		ginkgo.It("should ZRangeByScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRangeByScore := adapter.ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "-inf",
				Max: "+inf",
			})
			gomega.Expect(zRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByScore.Val()).To(gomega.Equal([]string{"one", "two", "three"}))

			zRangeByScore = adapter.ZRangeByScore(ctx, "zset", ZRangeBy{
				Min:    "-inf",
				Max:    "+inf",
				Offset: 1,
				Count:  2,
			})
			gomega.Expect(zRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByScore.Val()).To(gomega.Equal([]string{"two", "three"}))

			zRangeByScore = adapter.ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "1",
				Max: "2",
			})
			gomega.Expect(zRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByScore.Val()).To(gomega.Equal([]string{"one", "two"}))

			zRangeByScore = adapter.ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "2",
			})
			gomega.Expect(zRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByScore.Val()).To(gomega.Equal([]string{"two"}))

			zRangeByScore = adapter.ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "(2",
			})
			gomega.Expect(zRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByScore.Val()).To(gomega.Equal([]string{}))
		})

		ginkgo.It("should ZRangeByLex", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "a",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "b",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "c",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRangeByLex := adapter.ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "-",
				Max: "+",
			})
			gomega.Expect(zRangeByLex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByLex.Val()).To(gomega.Equal([]string{"a", "b", "c"}))

			zRangeByLex = adapter.ZRangeByLex(ctx, "zset", ZRangeBy{
				Min:    "-",
				Max:    "+",
				Offset: 1,
				Count:  2,
			})
			gomega.Expect(zRangeByLex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByLex.Val()).To(gomega.Equal([]string{"b", "c"}))

			zRangeByLex = adapter.ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "[a",
				Max: "[b",
			})
			gomega.Expect(zRangeByLex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByLex.Val()).To(gomega.Equal([]string{"a", "b"}))

			zRangeByLex = adapter.ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "(a",
				Max: "[b",
			})
			gomega.Expect(zRangeByLex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByLex.Val()).To(gomega.Equal([]string{"b"}))

			zRangeByLex = adapter.ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "(a",
				Max: "(b",
			})
			gomega.Expect(zRangeByLex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByLex.Val()).To(gomega.Equal([]string{}))
		})

		ginkgo.It("should ZRangeByScoreWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "-inf",
				Max: "+inf",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 2, Member: "two"}}))

			vals, err = adapter.ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "(2",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{}))
		})

		if resp3 {
			ginkgo.It("should ZRangeStore", func() {
				added, err := adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					Members: []Z{
						{Score: 1, Member: "one"},
						{Score: 2, Member: "two"},
						{Score: 3, Member: "three"},
						{Score: 4, Member: "four"},
					},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(4)))

				rangeStore, err := adapter.ZRangeStore(ctx, "new-zset", ZRangeArgs{
					Key:    "zset",
					Start:  "-",
					Stop:   "+",
					ByLex:  true,
					Rev:    false,
					Offset: 1,
					Count:  2,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(rangeStore).To(gomega.Equal(int64(2)))

				zRange, err := adapter.ZRange(ctx, "new-zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRange).To(gomega.Equal([]string{"two", "three"}))
			})
			ginkgo.It("should ZRangeStore Rev", func() {
				added, err := adapter.ZAddArgs(ctx, "zset", ZAddArgs{
					Members: []Z{
						{Score: 1, Member: "one"},
						{Score: 2, Member: "two"},
						{Score: 3, Member: "three"},
						{Score: 4, Member: "four"},
					},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(added).To(gomega.Equal(int64(4)))

				rangeStore, err := adapter.ZRangeStore(ctx, "new-zset", ZRangeArgs{
					Key:     "zset",
					Start:   1,
					Stop:    4,
					ByScore: true,
					Rev:     true,
					Offset:  1,
					Count:   2,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(rangeStore).To(gomega.Equal(int64(2)))

				zRange, err := adapter.ZRange(ctx, "new-zset", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRange).To(gomega.Equal([]string{"two", "three"}))
			})
		}

		ginkgo.It("should ZRank", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRank := adapter.ZRank(ctx, "zset", "three")
			gomega.Expect(zRank.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRank.Val()).To(gomega.Equal(int64(2)))

			zRank = adapter.ZRank(ctx, "zset", "four")
			gomega.Expect(rueidis.IsRedisNil(zRank.Err())).To(gomega.BeTrue())
			gomega.Expect(zRank.Val()).To(gomega.Equal(int64(0)))
		})

		if resp3 {
			ginkgo.It("should ZRankWithScore", func() {
				err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				zRankWithScore := adapter.ZRankWithScore(ctx, "zset", "one")
				gomega.Expect(zRankWithScore.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 0, Score: 1}))

				zRankWithScore = adapter.ZRankWithScore(ctx, "zset", "two")
				gomega.Expect(zRankWithScore.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 1, Score: 2}))

				zRankWithScore = adapter.ZRankWithScore(ctx, "zset", "three")
				gomega.Expect(zRankWithScore.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 2, Score: 3}))

				zRankWithScore = adapter.ZRankWithScore(ctx, "zset", "four")
				gomega.Expect(zRankWithScore.Err()).To(gomega.HaveOccurred())
				gomega.Expect(zRankWithScore.Err()).To(gomega.Equal(rueidis.Nil))
			})
		}

		ginkgo.It("should ZRem", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRem := adapter.ZRem(ctx, "zset", "two")
			gomega.Expect(zRem.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRem.Val()).To(gomega.Equal(int64(1)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
				Score:  1,
				Member: "one",
			}, {
				Score:  3,
				Member: "three",
			}}))
		})

		ginkgo.It("should ZRemRangeByRank", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRemRangeByRank := adapter.ZRemRangeByRank(ctx, "zset", 0, 1)
			gomega.Expect(zRemRangeByRank.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRemRangeByRank.Val()).To(gomega.Equal(int64(2)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
				Score:  3,
				Member: "three",
			}}))
		})

		ginkgo.It("should ZRemRangeByScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRemRangeByScore := adapter.ZRemRangeByScore(ctx, "zset", "-inf", "(2")
			gomega.Expect(zRemRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRemRangeByScore.Val()).To(gomega.Equal(int64(1)))

			vals, err := adapter.ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))
		})

		ginkgo.It("should ZRemRangeByLex", func() {
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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			}

			n, err := adapter.ZRemRangeByLex(ctx, "zset", "[alpha", "[omega").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(6)))

			vals, err := adapter.ZRange(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"ALPHA", "aaaa", "zap", "zip"}))
		})

		ginkgo.It("should ZRevRange", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRevRange := adapter.ZRevRange(ctx, "zset", 0, -1)
			gomega.Expect(zRevRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRange.Val()).To(gomega.Equal([]string{"three", "two", "one"}))

			zRevRange = adapter.ZRevRange(ctx, "zset", 2, 3)
			gomega.Expect(zRevRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRange.Val()).To(gomega.Equal([]string{"one"}))

			zRevRange = adapter.ZRevRange(ctx, "zset", -2, -1)
			gomega.Expect(zRevRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRange.Val()).To(gomega.Equal([]string{"two", "one"}))
		})

		ginkgo.It("should ZRevRangeWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			val, err := adapter.ZRevRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]Z{{Score: 1, Member: "one"}}))

			val, err = adapter.ZRevRangeWithScores(ctx, "zset", -2, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))
		})

		ginkgo.It("should ZRevRangeByScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"three", "two", "one"}))

			vals, err = adapter.ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf", Offset: 1, Count: 2}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"two", "one"}))

			vals, err = adapter.ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "2", Min: "(1"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"two"}))

			vals, err = adapter.ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "(2", Min: "(1"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{}))
		})

		ginkgo.It("should ZRevRangeByLex", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "a"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "b"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "c"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "+", Min: "-"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"c", "b", "a"}))

			vals, err = adapter.ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "+", Min: "-", Offset: 1, Count: 2}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"b", "a"}))

			vals, err = adapter.ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "[b", Min: "(a"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"b"}))

			vals, err = adapter.ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "(b", Min: "(a"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{}))
		})

		ginkgo.It("should ZRevRangeByScoreWithScores", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))
		})

		ginkgo.It("should ZRevRangeByScoreWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 2, Member: "two"}}))

			vals, err = adapter.ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "(2", Min: "(1"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{}))
		})

		ginkgo.It("should ZRevRank", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRevRank := adapter.ZRevRank(ctx, "zset", "one")
			gomega.Expect(zRevRank.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRank.Val()).To(gomega.Equal(int64(2)))

			zRevRank = adapter.ZRevRank(ctx, "zset", "four")
			gomega.Expect(rueidis.IsRedisNil(zRevRank.Err())).To(gomega.BeTrue())
			gomega.Expect(zRevRank.Val()).To(gomega.Equal(int64(0)))
		})

		if resp3 {
			ginkgo.It("should ZRevRankWithScore", func() {
				err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				zRevRankWithScore := adapter.ZRevRankWithScore(ctx, "zset", "one")
				gomega.Expect(zRevRankWithScore.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRevRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 2, Score: 1}))

				zRevRankWithScore = adapter.ZRevRankWithScore(ctx, "zset", "two")
				gomega.Expect(zRevRankWithScore.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRevRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 1, Score: 2}))

				zRevRankWithScore = adapter.ZRevRankWithScore(ctx, "zset", "three")
				gomega.Expect(zRevRankWithScore.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(zRevRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 0, Score: 3}))

				zRevRankWithScore = adapter.ZRevRankWithScore(ctx, "zset", "four")
				gomega.Expect(zRevRankWithScore.Err()).To(gomega.HaveOccurred())
				gomega.Expect(zRevRankWithScore.Err()).To(gomega.Equal(rueidis.Nil))
			})
		}

		ginkgo.It("should ZScore", func() {
			zAdd := adapter.ZAdd(ctx, "zset", Z{Score: 1.001, Member: "one"})
			gomega.Expect(zAdd.Err()).NotTo(gomega.HaveOccurred())

			zScore := adapter.ZScore(ctx, "zset", "one")
			gomega.Expect(zScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zScore.Val()).To(gomega.Equal(float64(1.001)))
		})

		if resp3 {
			ginkgo.It("should ZUnion", func() {
				err := adapter.ZAddArgs(ctx, "zset1", ZAddArgs{
					Members: []Z{
						{Score: 1, Member: "one"},
						{Score: 2, Member: "two"},
					},
				}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.ZAddArgs(ctx, "zset2", ZAddArgs{
					Members: []Z{
						{Score: 1, Member: "one"},
						{Score: 2, Member: "two"},
						{Score: 3, Member: "three"},
					},
				}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				union, err := adapter.ZUnion(ctx, ZStore{
					Keys:      []string{"zset1", "zset2"},
					Weights:   []int64{2, 3},
					Aggregate: "sum",
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(union).To(gomega.Equal([]string{"one", "three", "two"}))

				unionScores, err := adapter.ZUnionWithScores(ctx, ZStore{
					Keys:      []string{"zset1", "zset2"},
					Weights:   []int64{2, 3},
					Aggregate: "sum",
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(unionScores).To(gomega.Equal([]Z{
					{Score: 5, Member: "one"},
					{Score: 9, Member: "three"},
					{Score: 10, Member: "two"},
				}))
			})
		}

		ginkgo.It("should ZUnionStore", func() {
			err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			n, err := adapter.ZUnionStore(ctx, "out", ZStore{
				Keys:    []string{"zset1", "zset2"},
				Weights: []int64{2, 3},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(3)))

			val, err := adapter.ZRangeWithScores(ctx, "out", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]Z{{
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
			ginkgo.It("should ZRandMember", func() {
				err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				v := adapter.ZRandMember(ctx, "zset", 1)
				gomega.Expect(v.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(v.Val()).To(gomega.Or(gomega.Equal([]string{"one"}), gomega.Equal([]string{"two"})))

				v = adapter.ZRandMember(ctx, "zset", 0)
				gomega.Expect(v.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(v.Val()).To(gomega.HaveLen(0))

				kv, err := adapter.ZRandMemberWithScores(ctx, "zset", 1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(kv).To(gomega.Or(
					gomega.Equal([]Z{{Member: "one", Score: 1}}),
					gomega.Equal([]Z{{Member: "two", Score: 2}}),
				))
			})

			ginkgo.It("should ZDiff", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 3, Member: "three"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				v, err := adapter.ZDiff(ctx, "zset1", "zset2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(v).To(gomega.Equal([]string{"two", "three"}))
			})

			ginkgo.It("should ZDiffWithScores", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 3, Member: "three"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				v, err := adapter.ZDiffWithScores(ctx, "zset1", "zset2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(v).To(gomega.Equal([]Z{
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

			ginkgo.It("should ZInter", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				v, err := adapter.ZInter(ctx, ZStore{
					Keys: []string{"zset1", "zset2"},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(v).To(gomega.Equal([]string{"one", "two"}))
			})

			ginkgo.It("should ZInterCard", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				// limit 0 means no limit
				sInterCard := adapter.ZInterCard(ctx, 0, "zset1", "zset2")
				gomega.Expect(sInterCard.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(sInterCard.Val()).To(gomega.Equal(int64(2)))

				sInterCard = adapter.ZInterCard(ctx, 1, "zset1", "zset2")
				gomega.Expect(sInterCard.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(sInterCard.Val()).To(gomega.Equal(int64(1)))

				sInterCard = adapter.ZInterCard(ctx, 3, "zset1", "zset2")
				gomega.Expect(sInterCard.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(sInterCard.Val()).To(gomega.Equal(int64(2)))
			})

			ginkgo.It("should ZInterWithScores", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				v, err := adapter.ZInterWithScores(ctx, ZStore{
					Keys:      []string{"zset1", "zset2"},
					Weights:   []int64{2, 3},
					Aggregate: "Max",
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(v).To(gomega.Equal([]Z{
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

			ginkgo.It("should ZDiffStore", func() {
				err := adapter.ZAdd(ctx, "zset1", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset1", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				v, err := adapter.ZDiffStore(ctx, "out1", "zset1", "zset2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(v).To(gomega.Equal(int64(0)))
				v, err = adapter.ZDiffStore(ctx, "out1", "zset2", "zset1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(v).To(gomega.Equal(int64(1)))
				vals, err := adapter.ZRangeWithScores(ctx, "out1", 0, -1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]Z{{
					Score:  3,
					Member: "three",
				}}))
			})
		}
	})

	ginkgo.Describe("streams", func() {
		ginkgo.BeforeEach(func() {
			if resp3 {
				_, err := adapter.XAdd(ctx, XAddArgs{
					Stream:     "stream",
					ID:         "1-0",
					Values:     map[string]any{"uno": "un"},
					NoMkStream: true,
				}).Result()
				gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
			}

			id, err := adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				ID:     "1-0",
				Values: map[string]any{"uno": "un"},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(id).To(gomega.Equal("1-0"))

			// Values supports []any.
			id, err = adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				ID:     "2-0",
				Values: []any{"dos", "deux"},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(id).To(gomega.Equal("2-0"))

			// Value supports []string.
			id, err = adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				ID:     "3-0",
				Values: []string{"tres", "troix"},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(id).To(gomega.Equal("3-0"))
		})

		ginkgo.It("should XTrimMaxLen", func() {
			n, err := adapter.XTrimMaxLen(ctx, "stream", 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(3)))
		})

		ginkgo.It("should XTrimMaxLenApprox", func() {
			n, err := adapter.XTrimMaxLenApprox(ctx, "stream", 0, 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(3)))
		})

		if resp3 {
			ginkgo.It("should XTrimMaxLenApprox Limit", func() {
				n, err := adapter.XTrimMaxLenApprox(ctx, "stream", 0, 1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(0)))
			})

			ginkgo.It("should XTrimMinID", func() {
				n, err := adapter.XTrimMinID(ctx, "stream", "4-0").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(3)))
			})

			ginkgo.It("should XTrimMinIDApprox", func() {
				n, err := adapter.XTrimMinIDApprox(ctx, "stream", "4-0", 0).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(3)))
			})
		}

		ginkgo.It("should XAdd", func() {
			id, err := adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				Values: map[string]any{"quatro": "quatre"},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
				{ID: id, Values: map[string]any{"quatro": "quatre"}},
			}))
		})

		// TODO XAdd There is a bug in the limit parameter.
		// TODO Don't test it for now.
		// TODO link: https://github.com/redis/redis/issues/9046
		ginkgo.It("should XAdd with MaxLen", func() {
			id, err := adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				MaxLen: 1,
				Values: map[string]any{"quatro": "quatre"},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]XMessage{
				{ID: id, Values: map[string]any{"quatro": "quatre"}},
			}))
		})

		ginkgo.It("should XAdd with MaxLen Approx", func() {
			id, err := adapter.XAdd(ctx, XAddArgs{
				Stream: "stream",
				MaxLen: 1,
				Approx: true,
				Values: map[string]any{"quatro": "quatre"},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
				{ID: id, Values: map[string]any{"quatro": "quatre"}},
			}))
		})

		if resp3 {
			ginkgo.It("should XAdd with MinID", func() {
				id, err := adapter.XAdd(ctx, XAddArgs{
					Stream: "stream",
					MinID:  "5-0",
					ID:     "4-0",
					Values: map[string]any{"quatro": "quatre"},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(id).To(gomega.Equal("4-0"))

				vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.HaveLen(0))
			})

			ginkgo.It("should XAdd with MinID Approx", func() {
				id, err := adapter.XAdd(ctx, XAddArgs{
					Stream: "stream",
					MinID:  "5-0",
					ID:     "4-0",
					Approx: true,
					Values: map[string]any{"quatro": "quatre"},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(id).To(gomega.Equal("4-0"))

				vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.HaveLen(0))
			})

			ginkgo.It("should XAdd with MinID Limit", func() {
				id, err := adapter.XAdd(ctx, XAddArgs{
					Stream: "stream",
					MinID:  "5-0",
					ID:     "4-0",
					Approx: true,
					Values: map[string]any{"quatro": "quatre"},
					Limit:  1,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(id).To(gomega.Equal("4-0"))

				vals, err := adapter.XRange(ctx, "stream", "-", "+").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]XMessage{
					{ID: "1-0", Values: map[string]any{"uno": "un"}},
					{ID: "2-0", Values: map[string]any{"dos": "deux"}},
					{ID: "3-0", Values: map[string]any{"tres": "troix"}},
					{ID: id, Values: map[string]any{"quatro": "quatre"}},
				}))
			})
		}

		ginkgo.It("should XDel", func() {
			n, err := adapter.XDel(ctx, "stream", "1-0", "2-0", "3-0").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(3)))
		})

		ginkgo.It("should XLen", func() {
			n, err := adapter.XLen(ctx, "stream").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(3)))
		})

		ginkgo.It("should XRange", func() {
			msgs, err := adapter.XRange(ctx, "stream", "-", "+").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(msgs).To(gomega.Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
			}))

			msgs, err = adapter.XRange(ctx, "stream", "2", "+").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(msgs).To(gomega.Equal([]XMessage{
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
			}))

			msgs, err = adapter.XRange(ctx, "stream", "-", "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(msgs).To(gomega.Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
			}))
		})

		ginkgo.It("should XRangeN", func() {
			msgs, err := adapter.XRangeN(ctx, "stream", "-", "+", 2).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(msgs).To(gomega.Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
			}))

			msgs, err = adapter.XRangeN(ctx, "stream", "2", "+", 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(msgs).To(gomega.Equal([]XMessage{
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
			}))

			msgs, err = adapter.XRangeN(ctx, "stream", "-", "2", 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(msgs).To(gomega.Equal([]XMessage{
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
			}))
		})

		ginkgo.It("should XRevRange", func() {
			msgs, err := adapter.XRevRange(ctx, "stream", "+", "-").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(msgs).To(gomega.Equal([]XMessage{
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
				{ID: "1-0", Values: map[string]any{"uno": "un"}},
			}))

			msgs, err = adapter.XRevRange(ctx, "stream", "+", "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(msgs).To(gomega.Equal([]XMessage{
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
			}))
		})

		ginkgo.It("should XRevRangeN", func() {
			msgs, err := adapter.XRevRangeN(ctx, "stream", "+", "-", 2).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(msgs).To(gomega.Equal([]XMessage{
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
				{ID: "2-0", Values: map[string]any{"dos": "deux"}},
			}))

			msgs, err = adapter.XRevRangeN(ctx, "stream", "+", "2", 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(msgs).To(gomega.Equal([]XMessage{
				{ID: "3-0", Values: map[string]any{"tres": "troix"}},
			}))
		})

		ginkgo.It("should XRead", func() {
			res, err := adapter.XReadStreams(ctx, "stream", "0").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.Equal([]XStream{
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
			gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
		})

		ginkgo.It("should XRead", func() {
			res, err := adapter.XRead(ctx, XReadArgs{
				Streams: []string{"stream", "0"},
				Count:   2,
				Block:   100 * time.Millisecond,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.Equal([]XStream{
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
			gomega.Expect(rueidis.IsRedisNil(err)).To(gomega.BeTrue())
		})

		ginkgo.Describe("group", func() {
			ginkgo.BeforeEach(func() {
				err := adapter.XGroupCreate(ctx, "stream", "group", "0").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.XGroupSetID(ctx, "stream", "group", "0").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				res, err := adapter.XReadGroup(ctx, XReadGroupArgs{
					Group:    "group",
					Consumer: "consumer",
					Streams:  []string{"stream", ">"},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal([]XStream{
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

			ginkgo.AfterEach(func() {
				n, err := adapter.XGroupDestroy(ctx, "stream", "group").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(1)))
			})

			ginkgo.It("should XReadGroup skip empty", func() {
				n, err := adapter.XDel(ctx, "stream", "2-0").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(1)))

				res, err := adapter.XReadGroup(ctx, XReadGroupArgs{
					Group:    "group",
					Consumer: "consumer",
					Streams:  []string{"stream", "0"},
					NoAck:    true,
					Block:    -1,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal([]XStream{
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

			ginkgo.It("should XGroupCreateMkStream", func() {
				err := adapter.XGroupCreateMkStream(ctx, "stream2", "group", "0").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.XGroupCreateMkStream(ctx, "stream2", "group", "0").Err()
				gomega.Expect(err.Error()).To(gomega.Equal("BUSYGROUP Consumer Group name already exists"))

				n, err := adapter.XGroupDestroy(ctx, "stream2", "group").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(1)))

				n, err = adapter.Del(ctx, "stream2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(1)))
			})

			if resp3 {
				ginkgo.It("should XPending", func() {
					info, err := adapter.XPending(ctx, "stream", "group").Result()
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					gomega.Expect(info).To(gomega.Equal(XPending{
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
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					for i := range infoExt {
						infoExt[i].Idle = 0
					}
					gomega.Expect(infoExt).To(gomega.Equal([]XPendingExt{
						{ID: "1-0", Consumer: "consumer", Idle: 0, RetryCount: 1},
						{ID: "2-0", Consumer: "consumer", Idle: 0, RetryCount: 1},
						{ID: "3-0", Consumer: "consumer", Idle: 0, RetryCount: 1},
					}))

					args.Idle = 72 * time.Hour
					infoExt, err = adapter.XPendingExt(ctx, args).Result()
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					gomega.Expect(infoExt).To(gomega.HaveLen(0))
				})

				ginkgo.It("should XGroup Create Delete Consumer", func() {
					n, err := adapter.XGroupCreateConsumer(ctx, "stream", "group", "c1").Result()
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					gomega.Expect(n).To(gomega.Equal(int64(1)))

					n, err = adapter.XGroupDelConsumer(ctx, "stream", "group", "consumer").Result()
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					gomega.Expect(n).To(gomega.Equal(int64(3)))
				})

				ginkgo.It("should XAutoClaim", func() {
					xca := XAutoClaimArgs{
						Stream:   "stream",
						Group:    "group",
						Consumer: "consumer",
						Start:    "-",
						Count:    2,
					}
					msgs, start, err := adapter.XAutoClaim(ctx, xca).Result()
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					gomega.Expect(start).To(gomega.Equal("3-0"))
					gomega.Expect(msgs).To(gomega.Equal([]XMessage{{
						ID:     "1-0",
						Values: map[string]any{"uno": "un"},
					}, {
						ID:     "2-0",
						Values: map[string]any{"dos": "deux"},
					}}))

					xca.Start = start
					msgs, start, err = adapter.XAutoClaim(ctx, xca).Result()
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					gomega.Expect(start).To(gomega.Equal("0-0"))
					gomega.Expect(msgs).To(gomega.Equal([]XMessage{{
						ID:     "3-0",
						Values: map[string]any{"tres": "troix"},
					}}))

					ids, start, err := adapter.XAutoClaimJustID(ctx, xca).Result()
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					gomega.Expect(start).To(gomega.Equal("0-0"))
					gomega.Expect(ids).To(gomega.Equal([]string{"3-0"}))
				})

				ginkgo.It("should XAutoClaim NoCount", func() {
					xca := XAutoClaimArgs{
						Stream:   "stream",
						Group:    "group",
						Consumer: "consumer",
						Start:    "-",
					}
					msgs, start, err := adapter.XAutoClaim(ctx, xca).Result()
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					gomega.Expect(start).To(gomega.Equal("0-0"))
					gomega.Expect(msgs).To(gomega.Equal([]XMessage{{
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
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					gomega.Expect(start).To(gomega.Equal("0-0"))
					gomega.Expect(ids).To(gomega.Equal([]string{"1-0", "2-0", "3-0"}))
				})
			}

			ginkgo.It("should XClaim", func() {
				msgs, err := adapter.XClaim(ctx, XClaimArgs{
					Stream:   "stream",
					Group:    "group",
					Consumer: "consumer",
					Messages: []string{"1-0", "2-0", "3-0"},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(msgs).To(gomega.Equal([]XMessage{{
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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(ids).To(gomega.Equal([]string{"1-0", "2-0", "3-0"}))
			})

			ginkgo.It("should XAck", func() {
				n, err := adapter.XAck(ctx, "stream", "group", "1-0", "2-0", "4-0").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(2)))
			})
		})

		ginkgo.Describe("xinfo", func() {
			ginkgo.BeforeEach(func() {
				err := adapter.XGroupCreate(ctx, "stream", "group1", "0").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				res, err := adapter.XReadGroup(ctx, XReadGroupArgs{
					Group:    "group1",
					Consumer: "consumer1",
					Streams:  []string{"stream", ">"},
					Count:    2,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal([]XStream{
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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal([]XStream{
					{
						Stream: "stream",
						Messages: []XMessage{
							{ID: "3-0", Values: map[string]any{"tres": "troix"}},
						},
					},
				}))

				err = adapter.XGroupCreate(ctx, "stream", "group2", "1-0").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				res, err = adapter.XReadGroup(ctx, XReadGroupArgs{
					Group:    "group2",
					Consumer: "consumer1",
					Streams:  []string{"stream", ">"},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal([]XStream{
					{
						Stream: "stream",
						Messages: []XMessage{
							{ID: "2-0", Values: map[string]any{"dos": "deux"}},
							{ID: "3-0", Values: map[string]any{"tres": "troix"}},
						},
					},
				}))
			})

			ginkgo.AfterEach(func() {
				n, err := adapter.XGroupDestroy(ctx, "stream", "group1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(1)))
				n, err = adapter.XGroupDestroy(ctx, "stream", "group2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(1)))
			})

			ginkgo.It("should XINFO STREAM", func() {
				res, err := adapter.XInfoStream(ctx, "stream").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				res.RadixTreeKeys = 0
				res.RadixTreeNodes = 0

				if resp3 {
					gomega.Expect(res).To(gomega.Equal(XInfoStream{
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
					gomega.Expect(res).To(gomega.Equal(XInfoStream{
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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(n).To(gomega.Equal(int64(3)))

				res, err = adapter.XInfoStream(ctx, "stream").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				res.RadixTreeKeys = 0
				res.RadixTreeNodes = 0

				if resp3 {
					gomega.Expect(res).To(gomega.Equal(XInfoStream{
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
					gomega.Expect(res).To(gomega.Equal(XInfoStream{
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
				ginkgo.It("should XINFO STREAM FULL", func() {
					res, err := adapter.XInfoStreamFull(ctx, "stream", 2).Result()
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					res.RadixTreeKeys = 0
					res.RadixTreeNodes = 0

					// Verify DeliveryTime
					now := time.Now()
					maxElapsed := 10 * time.Minute
					for k, g := range res.Groups {
						for k2, p := range g.Pending {
							gomega.Expect(now.Sub(p.DeliveryTime)).To(gomega.BeNumerically("<=", maxElapsed))
							res.Groups[k].Pending[k2].DeliveryTime = time.Time{}
						}
						for k3, c := range g.Consumers {
							gomega.Expect(now.Sub(c.SeenTime)).To(gomega.BeNumerically("<=", maxElapsed))
							res.Groups[k].Consumers[k3].SeenTime = time.Time{}

							for k4, p := range c.Pending {
								gomega.Expect(now.Sub(p.DeliveryTime)).To(gomega.BeNumerically("<=", maxElapsed))
								res.Groups[k].Consumers[k3].Pending[k4].DeliveryTime = time.Time{}
							}
						}
					}

					gomega.Expect(res).To(gomega.Equal(XInfoStreamFull{
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

			ginkgo.It("should XINFO GROUPS", func() {
				res, err := adapter.XInfoGroups(ctx, "stream").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				if resp3 {
					gomega.Expect(res).To(gomega.Equal([]XInfoGroup{
						{Name: "group1", Consumers: 2, Pending: 3, LastDeliveredID: "3-0", EntriesRead: 3},
						{Name: "group2", Consumers: 1, Pending: 2, LastDeliveredID: "3-0", EntriesRead: 3},
					}))
				} else {
					gomega.Expect(res).To(gomega.Equal([]XInfoGroup{
						{Name: "group1", Consumers: 2, Pending: 3, LastDeliveredID: "3-0"},
						{Name: "group2", Consumers: 1, Pending: 2, LastDeliveredID: "3-0"},
					}))
				}
			})

			ginkgo.It("should XINFO CONSUMERS", func() {
				time.Sleep(time.Millisecond * 2) // make consumer idle > 0
				res, err := adapter.XInfoConsumers(ctx, "stream", "group1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				for i := range res {
					gomega.Expect(res[i].Idle > 0).To(gomega.BeTrue())
					res[i].Idle = 0
				}
				gomega.Expect(res).To(gomega.Equal([]XInfoConsumer{
					{Name: "consumer1", Pending: 2, Idle: 0},
					{Name: "consumer2", Pending: 1, Idle: 0},
				}))
			})
		})
	})

	ginkgo.Describe("Geo add and radius search", func() {
		ginkgo.BeforeEach(func() {
			n, err := adapter.GeoAdd(
				ctx,
				"Sicily",
				GeoLocation{Longitude: 13.361389, Latitude: 38.115556, Name: "Palermo"},
				GeoLocation{Longitude: 15.087269, Latitude: 37.502669, Name: "Catania"},
			).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should not add same geo location", func() {
			geoAdd := adapter.GeoAdd(
				ctx,
				"Sicily",
				GeoLocation{Longitude: 13.361389, Latitude: 38.115556, Name: "Palermo"},
			)
			gomega.Expect(geoAdd.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(geoAdd.Val()).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("should search geo radius", func() {
			res, err := adapter.GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius: 200,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.HaveLen(2))
			gomega.Expect(res[0].Name).To(gomega.Equal("Palermo"))
			gomega.Expect(res[1].Name).To(gomega.Equal("Catania"))
		})

		ginkgo.It("should geo radius and store the result", func() {
			n, err := adapter.GeoRadiusStore(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius: 200,
				Store:  "result",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(2)))

			res, err := adapter.ZRangeWithScores(ctx, "result", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.ContainElement(Z{
				Score:  3.479099956230698e+15,
				Member: "Palermo",
			}))
			gomega.Expect(res).To(gomega.ContainElement(Z{
				Score:  3.479447370796909e+15,
				Member: "Catania",
			}))
		})

		ginkgo.It("should geo radius and store dist", func() {
			n, err := adapter.GeoRadiusStore(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius:    200,
				StoreDist: "result",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(2)))

			res, err := adapter.ZRangeWithScores(ctx, "result", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.ContainElement(Z{
				Score:  190.44242984775784,
				Member: "Palermo",
			}))
			gomega.Expect(res).To(gomega.ContainElement(Z{
				Score:  56.4412578701582,
				Member: "Catania",
			}))
		})

		ginkgo.It("should search geo radius with options", func() {
			res, err := adapter.GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.HaveLen(2))
			gomega.Expect(res[1].Name).To(gomega.Equal("Palermo"))
			gomega.Expect(res[1].Dist).To(gomega.Equal(190.4424))
			gomega.Expect(res[1].GeoHash).To(gomega.Equal(int64(3479099956230698)))
			gomega.Expect(res[1].Longitude).To(gomega.Equal(13.361389338970184))
			gomega.Expect(res[1].Latitude).To(gomega.Equal(38.115556395496299))
			gomega.Expect(res[0].Name).To(gomega.Equal("Catania"))
			gomega.Expect(res[0].Dist).To(gomega.Equal(56.4413))
			gomega.Expect(res[0].GeoHash).To(gomega.Equal(int64(3479447370796909)))
			gomega.Expect(res[0].Longitude).To(gomega.Equal(15.087267458438873))
			gomega.Expect(res[0].Latitude).To(gomega.Equal(37.50266842333162))
		})

		ginkgo.It("should search geo radius with WithDist=false", func() {
			res, err := adapter.GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.HaveLen(2))
			gomega.Expect(res[1].Name).To(gomega.Equal("Palermo"))
			gomega.Expect(res[1].Dist).To(gomega.Equal(float64(0)))
			gomega.Expect(res[1].GeoHash).To(gomega.Equal(int64(3479099956230698)))
			gomega.Expect(res[1].Longitude).To(gomega.Equal(13.361389338970184))
			gomega.Expect(res[1].Latitude).To(gomega.Equal(38.115556395496299))
			gomega.Expect(res[0].Name).To(gomega.Equal("Catania"))
			gomega.Expect(res[0].Dist).To(gomega.Equal(float64(0)))
			gomega.Expect(res[0].GeoHash).To(gomega.Equal(int64(3479447370796909)))
			gomega.Expect(res[0].Longitude).To(gomega.Equal(15.087267458438873))
			gomega.Expect(res[0].Latitude).To(gomega.Equal(37.50266842333162))
		})

		ginkgo.It("should search geo radius by member with options", func() {
			res, err := adapter.GeoRadiusByMember(ctx, "Sicily", "Catania", GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.HaveLen(2))
			gomega.Expect(res[0].Name).To(gomega.Equal("Catania"))
			gomega.Expect(res[0].Dist).To(gomega.Equal(0.0))
			gomega.Expect(res[0].GeoHash).To(gomega.Equal(int64(3479447370796909)))
			gomega.Expect(res[0].Longitude).To(gomega.Equal(15.087267458438873))
			gomega.Expect(res[0].Latitude).To(gomega.Equal(37.50266842333162))
			gomega.Expect(res[1].Name).To(gomega.Equal("Palermo"))
			gomega.Expect(res[1].Dist).To(gomega.Equal(166.2742))
			gomega.Expect(res[1].GeoHash).To(gomega.Equal(int64(3479099956230698)))
			gomega.Expect(res[1].Longitude).To(gomega.Equal(13.361389338970184))
			gomega.Expect(res[1].Latitude).To(gomega.Equal(38.115556395496299))

			ress, err := adapter.GeoRadiusByMemberStore(ctx, "Sicily", "Catania", GeoRadiusQuery{
				Radius: 200,
				Unit:   "km",
				Count:  2,
				Store:  "Sicily2",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(ress).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should search geo radius with no results", func() {
			res, err := adapter.GeoRadius(ctx, "Sicily", 99, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.HaveLen(0))
		})

		ginkgo.It("should get geo distance with unit options", func() {
			// From Redis CLI, note the difference in rounding in m vs
			// km on Redis itself.
			//
			// GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
			// GEODIST Sicily Palermo Catania m
			// "166274.15156960033"
			// GEODIST Sicily Palermo Catania km
			// "166.27415156960032"
			dist, err := adapter.GeoDist(ctx, "Sicily", "Palermo", "Catania", "km").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(dist).To(gomega.BeNumerically("~", 166.27, 0.01))

			dist, err = adapter.GeoDist(ctx, "Sicily", "Palermo", "Catania", "m").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(dist).To(gomega.BeNumerically("~", 166274.15, 0.01))

			_, err = adapter.GeoDist(ctx, "Sicily", "Palermo", "Catania", "mi").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			_, err = adapter.GeoDist(ctx, "Sicily", "Palermo", "Catania", "ft").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("should get geo hash in string representation", func() {
			hashes, err := adapter.GeoHash(ctx, "Sicily", "Palermo", "Catania").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(hashes).To(gomega.ConsistOf([]string{"sqc8b49rny0", "sqdtr74hyu0"}))
		})

		ginkgo.It("should return geo position", func() {
			pos, err := adapter.GeoPos(ctx, "Sicily", "Palermo", "Catania", "NonExisting").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.ConsistOf([]*GeoPos{
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
			ginkgo.It("should geo search", func() {
				q := GeoSearchQuery{
					Member:    "Catania",
					BoxWidth:  400,
					BoxHeight: 100,
					BoxUnit:   "km",
					Sort:      "asc",
				}
				val, err := adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

				q.BoxHeight = 400
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania", "Palermo"}))

				q.Count = 1
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

				q.CountAny = true
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Palermo"}))

				q = GeoSearchQuery{
					Member:     "Catania",
					Radius:     100,
					RadiusUnit: "km",
					Sort:       "asc",
				}
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

				q.Radius = 400
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania", "Palermo"}))

				q.Count = 1
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

				q.CountAny = true
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Palermo"}))

				q = GeoSearchQuery{
					Longitude: 15,
					Latitude:  37,
					BoxWidth:  200,
					BoxHeight: 200,
					BoxUnit:   "km",
					Sort:      "asc",
				}
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

				q.BoxWidth, q.BoxHeight = 400, 400
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania", "Palermo"}))

				q.Count = 1
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

				q.CountAny = true
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Palermo"}))

				q = GeoSearchQuery{
					Longitude:  15,
					Latitude:   37,
					Radius:     100,
					RadiusUnit: "km",
					Sort:       "asc",
				}
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

				q.Radius = 200
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania", "Palermo"}))

				q.Count = 1
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

				q.CountAny = true
				val, err = adapter.GeoSearch(ctx, "Sicily", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]string{"Palermo"}))
			})

			ginkgo.It("should geo search with options", func() {
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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal([]GeoLocation{
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

			ginkgo.It("should geo search store", func() {
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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal(int64(2)))

				q.StoreDist = true
				val, err = adapter.GeoSearchStore(ctx, "Sicily", "key2", q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.Equal(int64(2)))

				loc, err := adapter.GeoSearchLocation(ctx, "key1", GeoSearchLocationQuery{
					GeoSearchQuery: q.GeoSearchQuery,
					WithCoord:      true,
					WithDist:       true,
					WithHash:       true,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(loc).To(gomega.Equal([]GeoLocation{
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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(v).To(gomega.Equal([]Z{
					{
						Score:  56.441257870158204,
						Member: "Catania",
					},
					{
						Score:  190.44242984775784,
						Member: "Palermo",
					},
				}))
			})
		}
	})

	ginkgo.Describe("marshaling/unmarshaling", func() {
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

		ginkgo.It("should convert to string", func() {
			for _, test := range convTests {
				err := adapter.Set(ctx, "key", test.value, 0).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				s, err := adapter.Get(ctx, "key").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(s).To(gomega.Equal(test.wanted))

				if test.dest == nil {
					continue
				}
				// TODO
				// err = adapter.Get(ctx, "key").Scan(test.dest)
				// gomega.Expect(err).NotTo(gomega.HaveOccurred())
				// gomega.Expect(deref(test.dest)).To(gomega.Equal(test.value))
			}
		})
	})

	ginkgo.Describe("json marshaling/unmarshaling", func() {
		ginkgo.BeforeEach(func() {
			value := &numberStruct{Number: 42}
			err := adapter.Set(ctx, "key", value, 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("should marshal custom values using json", func() {
			s, err := adapter.Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(s).To(gomega.Equal(`{"Number":42}`))
		})

		// TODO
		// ginkgo.It("should scan custom values using json", func() {
		//	value := &numberStruct{}
		//	err := adapter.Get(ctx, "key").Scan(value)
		//	gomega.Expect(err).NotTo(gomega.HaveOccurred())
		//	gomega.Expect(value.Number).To(gomega.Equal(42))
		// })
	})

	ginkgo.Describe("Pub/Sub", func() {
		ginkgo.It("Publish", func() {
			v, err := adapter.Publish(ctx, "ch", "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(v).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("PubSubChannels", func() {
			v, err := adapter.PubSubChannels(ctx, "*").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(v).To(gomega.Equal([]string{}))
		})

		ginkgo.It("PubSubNumPat", func() {
			v, err := adapter.PubSubNumPat(ctx).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(v).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("PubSubNumSub", func() {
			v, err := adapter.PubSubNumSub(ctx, "ch").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(v).To(gomega.Equal(map[string]int64{"ch": 0}))
		})

		if resp3 {
			ginkgo.It("SPublish", func() {
				v, err := adapter.SPublish(ctx, "ch", "1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(v).To(gomega.Equal(int64(0)))
			})

			ginkgo.It("PubSubShardChannels", func() {
				v, err := adapter.PubSubShardChannels(ctx, "*").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(v).To(gomega.Equal([]string{}))
			})

			ginkgo.It("PubSubShardNumSub", func() {
				v, err := adapter.PubSubShardNumSub(ctx, "ch").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(v).To(gomega.Equal(map[string]int64{"ch": 0}))
			})
		}
	})

	ginkgo.Describe("Script", func() {
		ginkgo.It("returns keys and values", func() {
			vals, err := adapter.Eval(
				ctx,
				"return {KEYS[1],ARGV[1]}",
				[]string{"key"},
				"hello",
			).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]any{"key", "hello"}))
		})

		ginkgo.It("returns all values after an error", func() {
			vals, err := adapter.Eval(
				ctx,
				`return {12, {err="error"}, "abc"}`,
				nil,
			).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals.([]any)[0]).To(gomega.Equal(int64(12)))
			gomega.Expect(vals.([]any)[1].(error).Error()).To(gomega.Equal("error"))
			gomega.Expect(vals.([]any)[2]).To(gomega.Equal("abc"))
		})

		ginkgo.It("script load", func() {
			val, err := adapter.ScriptLoad(
				ctx,
				"return {KEYS[1],ARGV[1]}",
			).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			ret, err := adapter.ScriptExists(ctx, val).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(ret).To(gomega.Equal([]bool{true}))

			vals, err := adapter.EvalSha(
				ctx,
				val,
				[]string{"key"},
				"hello",
			).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]any{"key", "hello"}))
		})

		if resp3 {
			ginkgo.It("script load", func() {
				val, err := adapter.ScriptLoad(
					ctx,
					"return {KEYS[1],ARGV[1]}",
				).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				ret, err := adapter.ScriptExists(ctx, val).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(ret).To(gomega.Equal([]bool{true}))

				valsRo, err := adapter.EvalShaRO(
					ctx,
					val,
					[]string{"key"},
					"hello",
				).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(valsRo).To(gomega.Equal([]any{"key", "hello"}))
			})
		}

		ginkgo.It("script kill & flush", func() {
			gomega.Expect(adapter.ScriptKill(ctx).Err()).To(gomega.MatchError("NOTBUSY No scripts in execution right now."))
			gomega.Expect(adapter.ScriptFlush(ctx).Err()).NotTo(gomega.HaveOccurred())
		})
	})
}

func testAdapterCache(resp3 bool) {

	var adapter Cmdable

	ginkgo.BeforeEach(func() {
		if resp3 {
			adapter = adapterresp3
		} else {
			adapter = adapterresp2
		}
		gomega.Expect(adapter.FlushDB(ctx).Err()).NotTo(gomega.HaveOccurred())
		gomega.Expect(adapter.FlushAll(ctx).Err()).NotTo(gomega.HaveOccurred())
	})

	ginkgo.Describe("Config", func() {
		ginkgo.It("Flush", func() {
			gomega.Expect(adapter.FlushDBAsync(ctx).Err()).NotTo(gomega.HaveOccurred())
			time.Sleep(2 * time.Second)
			gomega.Expect(adapter.FlushAllAsync(ctx).Err()).NotTo(gomega.HaveOccurred())
			time.Sleep(2 * time.Second)
		})
		ginkgo.It("BgSave", func() {
			gomega.Expect(adapter.BgRewriteAOF(ctx).Err()).NotTo(gomega.HaveOccurred())
			time.Sleep(2 * time.Second)
			gomega.Expect(adapter.BgSave(ctx).Err()).NotTo(gomega.HaveOccurred())
			time.Sleep(2 * time.Second)
		})
		ginkgo.It("Config Rewrite", func() {
			gomega.Expect(adapter.ConfigRewrite(ctx).Err()).To(gomega.MatchError("The server is running without a config file"))
		})
		ginkgo.It("DebugObject", func() {
			gomega.Expect(adapter.DebugObject(ctx, "non").Err().Error()).To(gomega.HavePrefix("DEBUG command not allowed."))
		})
		ginkgo.It("ReadOnly & ReadWrite", func() {
			gomega.Expect(adapter.ReadOnly(ctx).Err()).To(gomega.MatchError("This instance has cluster support disabled"))
			gomega.Expect(adapter.ReadWrite(ctx).Err()).To(gomega.MatchError("This instance has cluster support disabled"))
		})
	})

	ginkgo.Describe("Client", func() {
		ginkgo.It("should ClientUnblock", func() {
			id := adapter.ClientID(ctx).Val()
			r, err := adapter.ClientUnblock(ctx, id).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(r).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("should ClientUnblockWithError", func() {
			id := adapter.ClientID(ctx).Val()
			r, err := adapter.ClientUnblockWithError(ctx, id).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(r).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("ClientPause", func() {
			gomega.Expect(adapter.ClientPause(ctx, time.Second).Err()).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("ClientUnpause", func() {
			gomega.Expect(adapter.ClientUnpause(ctx).Err()).NotTo(gomega.HaveOccurred())
		})
	})

	ginkgo.Describe("EvalRO", func() {
		ginkgo.It("returns keys and values", func() {
			vals, err := adapter.EvalRO(
				ctx,
				"return {KEYS[1],ARGV[1]}",
				[]string{"key"},
				"hello",
			).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]any{"key", "hello"}))
		})

		ginkgo.It("returns all values after an error", func() {
			vals, err := adapter.EvalRO(
				ctx,
				`return {12, {err="error"}, "abc"}`,
				nil,
			).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals.([]any)[0]).To(gomega.Equal(int64(12)))
			gomega.Expect(vals.([]any)[1].(error).Error()).To(gomega.Equal("error"))
			gomega.Expect(vals.([]any)[2]).To(gomega.Equal("abc"))
		})
	})

	if resp3 {
		ginkgo.Describe("Functions", func() {
			var (
				q        FunctionListQuery
				lib1Code string
				lib2Code string
				lib1     Library
				lib2     Library
			)

			ginkgo.BeforeEach(func() {
				flush := adapter.FunctionFlush(ctx)
				gomega.Expect(flush.Err()).NotTo(gomega.HaveOccurred())

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

			ginkgo.It("Loads a new library", func() {
				functionLoad := adapter.FunctionLoad(ctx, lib1Code)
				gomega.Expect(functionLoad.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(functionLoad.Val()).To(gomega.Equal(lib1.Name))

				functionList := adapter.FunctionList(ctx, q)
				gomega.Expect(functionList.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(functionList.Val()).To(gomega.HaveLen(1))
			})

			ginkgo.It("Loads and replaces a new library", func() {
				// Load a library for the first time
				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				newFuncName := "replaces_func_name"
				newFuncDesc := "replaces_func_desc"
				flag1, flag2 := "allow-stale", "no-cluster"
				newCode := fmt.Sprintf(lib1.Code, lib1.Name, newFuncName, newFuncDesc, flag1, flag2)

				// And then replace it
				functionLoadReplace := adapter.FunctionLoadReplace(ctx, newCode)
				gomega.Expect(functionLoadReplace.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(functionLoadReplace.Val()).To(gomega.Equal(lib1.Name))

				lib, err := adapter.FunctionList(ctx, q).First()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(lib.Functions).To(gomega.Equal([]Function{
					{
						Name:        newFuncName,
						Description: newFuncDesc,
						Flags:       []string{flag1, flag2},
					},
				}))
			})

			ginkgo.It("Deletes a library", func() {
				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.FunctionDelete(ctx, lib1.Name).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				val, err := adapter.FunctionList(ctx, FunctionListQuery{
					LibraryNamePattern: lib1.Name,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.HaveLen(0))
			})

			ginkgo.It("Flushes all libraries", func() {
				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.FunctionLoad(ctx, lib2Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.FunctionFlush(ctx).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				val, err := adapter.FunctionList(ctx, q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.HaveLen(0))
			})

			ginkgo.It("Flushes all libraries asynchronously", func() {
				functionLoad := adapter.FunctionLoad(ctx, lib1Code)
				gomega.Expect(functionLoad.Err()).NotTo(gomega.HaveOccurred())

				// we only verify the command result.
				functionFlush := adapter.FunctionFlushAsync(ctx)
				gomega.Expect(functionFlush.Err()).NotTo(gomega.HaveOccurred())
			})

			ginkgo.It("Kills a running function", func() {
				functionKill := adapter.FunctionKill(ctx)
				gomega.Expect(functionKill.Err()).To(gomega.MatchError("NOTBUSY No scripts in execution right now."))

				// Add test for a long-running function, once we make the test for `function stats` pass
			})

			ginkgo.It("Lists registered functions", func() {
				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				val, err := adapter.FunctionList(ctx, FunctionListQuery{
					LibraryNamePattern: "*",
					WithCode:           true,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.HaveLen(1))
				gomega.Expect(val[0].Name).To(gomega.Equal(lib1.Name))
				gomega.Expect(val[0].Engine).To(gomega.Equal(lib1.Engine))
				gomega.Expect(val[0].Code).To(gomega.Equal(lib1Code))
				gomega.Expect(val[0].Functions).Should(gomega.ConsistOf(lib1.Functions))

				err = adapter.FunctionLoad(ctx, lib2Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				val, err = adapter.FunctionList(ctx, FunctionListQuery{
					WithCode: true,
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(val).To(gomega.HaveLen(2))

				lib, err := adapter.FunctionList(ctx, FunctionListQuery{
					LibraryNamePattern: lib2.Name,
					WithCode:           false,
				}).First()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(lib.Name).To(gomega.Equal(lib2.Name))
				gomega.Expect(lib.Code).To(gomega.Equal(""))

				_, err = adapter.FunctionList(ctx, FunctionListQuery{
					LibraryNamePattern: "non_lib",
					WithCode:           true,
				}).First()
				gomega.Expect(err).To(gomega.Equal(rueidis.Nil))
			})

			ginkgo.It("Dump and restores all libraries", func() {
				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.FunctionLoad(ctx, lib2Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				dump, err := adapter.FunctionDump(ctx).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(dump).NotTo(gomega.BeEmpty())

				err = adapter.FunctionRestore(ctx, dump).Err()
				gomega.Expect(err).To(gomega.HaveOccurred())

				err = adapter.FunctionFlush(ctx).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				list, err := adapter.FunctionList(ctx, q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(list).To(gomega.HaveLen(0))

				err = adapter.FunctionRestore(ctx, dump).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				list, err = adapter.FunctionList(ctx, q).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(list).To(gomega.HaveLen(2))
			})

			ginkgo.It("Calls a function", func() {
				lib1Code = fmt.Sprintf(lib1.Code, lib1.Name, lib1.Functions[0].Name,
					lib1.Functions[0].Description, lib1.Functions[0].Flags[0], lib1.Functions[0].Flags[1])

				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				x := adapter.FCall(ctx, lib1.Functions[0].Name, []string{"my_hash"}, "a", 1, "b", 2)
				gomega.Expect(x.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(x.Int()).To(gomega.Equal(3))
			})

			ginkgo.It("Calls a function as read-only", func() {
				lib1Code = fmt.Sprintf(lib1.Code, lib1.Name, lib1.Functions[0].Name,
					lib1.Functions[0].Description, lib1.Functions[0].Flags[0], lib1.Functions[0].Flags[1])

				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				// This function doesn't have a "no-writes" flag
				x := adapter.FCallRO(ctx, lib1.Functions[0].Name, []string{"my_hash"}, "a", 1, "b", 2)

				gomega.Expect(x.Err()).To(gomega.HaveOccurred())

				lib2Code = fmt.Sprintf(lib2.Code, lib2.Name, lib2.Functions[0].Name, lib2.Functions[1].Name,
					lib2.Functions[1].Description, lib2.Functions[1].Flags[0])

				// This function has a "no-writes" flag
				err = adapter.FunctionLoad(ctx, lib2Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				x = adapter.FCallRO(ctx, lib2.Functions[1].Name, []string{})

				gomega.Expect(x.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(x.Text()).To(gomega.Equal("Function 2"))
			})
		})
	}

	ginkgo.Describe("keys", func() {

		ginkgo.It("should Expire", func() {
			ttl := adapter.Cache(time.Hour).TTL(ctx, "nonexistent_key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(time.Duration(-2)))

			set := adapter.Set(ctx, "key", "Hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expire := adapter.Expire(ctx, "key", 10*time.Second)
			gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(expire.Val()).To(gomega.Equal(true))

			ttl = adapter.Cache(time.Hour).TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(10 * time.Second))

			set = adapter.Set(ctx, "key", "Hello World", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			time.Sleep(time.Millisecond * 100)

			ttl = adapter.Cache(time.Hour).TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(time.Duration(-1)))
		})

		ginkgo.It("should PExpire", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expiration := 900 * time.Millisecond
			pexpire := adapter.PExpire(ctx, "key", expiration)
			gomega.Expect(pexpire.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pexpire.Val()).To(gomega.Equal(true))

			ttl := adapter.Cache(time.Hour).TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(time.Second))

			pttl := adapter.Cache(time.Hour).PTTL(ctx, "key")
			gomega.Expect(pttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pttl.Val()).To(gomega.BeNumerically("~", expiration, 100*time.Millisecond))
		})

		ginkgo.It("should PExpireAt", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expiration := 900 * time.Millisecond
			pexpireat := adapter.PExpireAt(ctx, "key", time.Now().Add(expiration))
			gomega.Expect(pexpireat.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pexpireat.Val()).To(gomega.Equal(true))

			ttl := adapter.Cache(time.Hour).TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(time.Second))

			pttl := adapter.Cache(time.Hour).PTTL(ctx, "key")
			gomega.Expect(pttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pttl.Val()).To(gomega.BeNumerically("~", expiration, 100*time.Millisecond))
		})

		ginkgo.It("should PTTL", func() {
			set := adapter.Set(ctx, "key", "Hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expiration := time.Second
			expire := adapter.Expire(ctx, "key", expiration)
			gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			pttl := adapter.Cache(time.Hour).PTTL(ctx, "key")
			gomega.Expect(pttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(pttl.Val()).To(gomega.BeNumerically("~", expiration, 100*time.Millisecond))
		})

		ginkgo.It("should Sort", func() {
			gomega.Expect(func() {
				adapter.Cache(time.Hour).SortRO(ctx, "list", Sort{
					Order: "PANIC",
				})
			}).To(gomega.Panic())
		})

		ginkgo.It("should Sort", func() {
			size, err := adapter.LPush(ctx, "list", "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list", "3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list", "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(3)))

			els, err := adapter.Cache(time.Hour).SortRO(ctx, "list", Sort{
				Offset: 0,
				Count:  2,
				Order:  "ASC",
				Alpha:  true,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(els).To(gomega.Equal([]string{"1", "2"}))
		})

		ginkgo.It("should Sort By", func() {
			size, err := adapter.LPush(ctx, "list_by", "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list_by", "3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list_by", "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(3)))

			els, err := adapter.Cache(time.Hour).SortRO(ctx, "list_by", Sort{
				Offset: 0,
				Count:  2,
				By:     "nosort",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(els).To(gomega.Equal([]string{"2", "3"}))
		})

		ginkgo.It("should Sort and Get", func() {
			size, err := adapter.LPush(ctx, "list", "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(1)))

			size, err = adapter.LPush(ctx, "list", "3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(2)))

			size, err = adapter.LPush(ctx, "list", "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(size).To(gomega.Equal(int64(3)))

			err = adapter.Set(ctx, "object_2", "value2", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			{
				els, err := adapter.Cache(time.Hour).SortRO(ctx, "list", Sort{
					Get: []string{"object_*"},
				}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(els).To(gomega.Equal([]string{"", "value2", ""}))
			}

		})

		ginkgo.It("should TTL", func() {
			ttl := adapter.Cache(time.Hour).TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val() < 0).To(gomega.Equal(true))

			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			expire := adapter.Expire(ctx, "key", 60*time.Second)
			gomega.Expect(expire.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(expire.Val()).To(gomega.Equal(true))

			ttl = adapter.Cache(time.Hour).TTL(ctx, "key")
			gomega.Expect(ttl.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(ttl.Val()).To(gomega.Equal(60 * time.Second))
		})

		ginkgo.It("should Type", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			type_ := adapter.Cache(time.Hour).Type(ctx, "key")
			gomega.Expect(type_.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(type_.Val()).To(gomega.Equal("string"))
		})
	})

	ginkgo.Describe("strings", func() {

		ginkgo.It("should BitCount", func() {
			set := adapter.Set(ctx, "key", "foobar", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			bitCount := adapter.Cache(time.Hour).BitCount(ctx, "key", nil)
			gomega.Expect(bitCount.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bitCount.Val()).To(gomega.Equal(int64(26)))

			bitCount = adapter.Cache(time.Hour).BitCount(ctx, "key", &BitCount{
				Start: 0,
				End:   0,
			})
			gomega.Expect(bitCount.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bitCount.Val()).To(gomega.Equal(int64(4)))

			bitCount = adapter.Cache(time.Hour).BitCount(ctx, "key", &BitCount{
				Start: 1,
				End:   1,
			})
			gomega.Expect(bitCount.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(bitCount.Val()).To(gomega.Equal(int64(6)))
		})

		ginkgo.It("should BitPos", func() {
			err := adapter.Set(ctx, "mykey", "\xff\xf0\x00", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			pos, err := adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(12)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(0)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, 2).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(16)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 1, 2).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(-1)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(16)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 1, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(-1)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, 2, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(-1)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, 0, -3).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(-1)))

			pos, err = adapter.Cache(time.Hour).BitPos(ctx, "mykey", 0, 0, 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(-1)))
		})

		ginkgo.It("should BitPosSpan", func() {
			err := adapter.Set(ctx, "mykey", "\x00\xff\x00", 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			pos, err := adapter.Cache(time.Hour).BitPosSpan(ctx, "mykey", 0, 1, 3, "byte").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(16)))

			pos, err = adapter.Cache(time.Hour).BitPosSpan(ctx, "mykey", 0, 1, 3, "bit").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.Equal(int64(1)))
		})

		ginkgo.Describe("EvalRO", func() {
			ginkgo.It("returns keys and values", func() {
				vals, err := adapter.Cache(time.Hour).EvalRO(
					ctx,
					"return {KEYS[1],ARGV[1]}",
					[]string{"key"},
					"hello",
				).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals).To(gomega.Equal([]any{"key", "hello"}))
			})

			ginkgo.It("returns all values after an error", func() {
				vals, err := adapter.Cache(time.Hour).EvalRO(
					ctx,
					`return {12, {err="error"}, "abc"}`,
					[]string{"key"},
				).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(vals.([]any)[0]).To(gomega.Equal(int64(12)))
				gomega.Expect(vals.([]any)[1].(error).Error()).To(gomega.Equal("error"))
				gomega.Expect(vals.([]any)[2]).To(gomega.Equal("abc"))
			})
		})

		ginkgo.It("script load", func() {
			val, err := adapter.ScriptLoad(
				ctx,
				"return {KEYS[1],ARGV[1]}",
			).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			ret, err := adapter.ScriptExists(ctx, val).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(ret).To(gomega.Equal([]bool{true}))

			valsRo, err := adapter.Cache(time.Hour).EvalShaRO(
				ctx,
				val,
				[]string{"key"},
				"hello",
			).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(valsRo).To(gomega.Equal([]any{"key", "hello"}))
		})

		ginkgo.Describe("Functions", func() {
			var (
				lib1Code string
				lib2Code string
				lib1     Library
				lib2     Library
			)

			ginkgo.BeforeEach(func() {
				flush := adapter.FunctionFlush(ctx)
				gomega.Expect(flush.Err()).NotTo(gomega.HaveOccurred())

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

			ginkgo.It("Calls a function as read-only", func() {
				lib1Code = fmt.Sprintf(lib1.Code, lib1.Name, lib1.Functions[0].Name,
					lib1.Functions[0].Description, lib1.Functions[0].Flags[0], lib1.Functions[0].Flags[1])

				err := adapter.FunctionLoad(ctx, lib1Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				// This function doesn't have a "no-writes" flag
				x := adapter.Cache(time.Hour).FCallRO(ctx, lib1.Functions[0].Name, []string{"my_hash"}, "a", 1, "b", 2)

				gomega.Expect(x.Err()).To(gomega.HaveOccurred())

				lib2Code = fmt.Sprintf(lib2.Code, lib2.Name, lib2.Functions[0].Name, lib2.Functions[1].Name,
					lib2.Functions[1].Description, lib2.Functions[1].Flags[0])

				// This function has a "no-writes" flag
				err = adapter.FunctionLoad(ctx, lib2Code).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				x = adapter.Cache(time.Hour).FCallRO(ctx, lib2.Functions[1].Name, []string{"my_hash"})

				gomega.Expect(x.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(x.Text()).To(gomega.Equal("Function 2"))
			})
		})

		ginkgo.It("should Get", func() {
			get := adapter.Cache(time.Hour).Get(ctx, "_")
			gomega.Expect(rueidis.IsRedisNil(get.Err())).To(gomega.BeTrue())
			gomega.Expect(get.Val()).To(gomega.Equal(""))

			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			get = adapter.Cache(time.Hour).Get(ctx, "key")
			gomega.Expect(get.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(get.Val()).To(gomega.Equal("hello"))
		})

		ginkgo.It("should GetBit", func() {
			setBit := adapter.SetBit(ctx, "key", 7, 1)
			gomega.Expect(setBit.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(setBit.Val()).To(gomega.Equal(int64(0)))

			getBit := adapter.Cache(time.Hour).GetBit(ctx, "key", 0)
			gomega.Expect(getBit.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getBit.Val()).To(gomega.Equal(int64(0)))

			getBit = adapter.Cache(time.Hour).GetBit(ctx, "key", 7)
			gomega.Expect(getBit.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getBit.Val()).To(gomega.Equal(int64(1)))

			getBit = adapter.Cache(time.Hour).GetBit(ctx, "key", 100)
			gomega.Expect(getBit.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getBit.Val()).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("should GetRange", func() {
			set := adapter.Set(ctx, "key", "This is a string", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			getRange := adapter.Cache(time.Hour).GetRange(ctx, "key", 0, 3)
			gomega.Expect(getRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getRange.Val()).To(gomega.Equal("This"))

			getRange = adapter.Cache(time.Hour).GetRange(ctx, "key", -3, -1)
			gomega.Expect(getRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getRange.Val()).To(gomega.Equal("ing"))

			getRange = adapter.Cache(time.Hour).GetRange(ctx, "key", 0, -1)
			gomega.Expect(getRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getRange.Val()).To(gomega.Equal("This is a string"))

			getRange = adapter.Cache(time.Hour).GetRange(ctx, "key", 10, 100)
			gomega.Expect(getRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(getRange.Val()).To(gomega.Equal("string"))
		})

		ginkgo.It("should StrLen", func() {
			set := adapter.Set(ctx, "key", "hello", 0)
			gomega.Expect(set.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(set.Val()).To(gomega.Equal("OK"))

			strLen := adapter.Cache(time.Hour).StrLen(ctx, "key")
			gomega.Expect(strLen.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(strLen.Val()).To(gomega.Equal(int64(5)))

			strLen = adapter.Cache(time.Hour).StrLen(ctx, "_")
			gomega.Expect(strLen.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(strLen.Val()).To(gomega.Equal(int64(0)))
		})
	})

	ginkgo.Describe("hashes", func() {

		ginkgo.It("should HExists", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())

			hExists := adapter.Cache(time.Hour).HExists(ctx, "hash", "key")
			gomega.Expect(hExists.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hExists.Val()).To(gomega.Equal(true))

			hExists = adapter.Cache(time.Hour).HExists(ctx, "hash", "key1")
			gomega.Expect(hExists.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hExists.Val()).To(gomega.Equal(false))
		})

		ginkgo.It("should HGet", func() {
			hSet := adapter.HSet(ctx, "hash", "key", "hello")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())

			hGet := adapter.Cache(time.Hour).HGet(ctx, "hash", "key")
			gomega.Expect(hGet.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hGet.Val()).To(gomega.Equal("hello"))

			hGet = adapter.Cache(time.Hour).HGet(ctx, "hash", "key1")
			gomega.Expect(rueidis.IsRedisNil(hGet.Err())).To(gomega.BeTrue())
			gomega.Expect(hGet.Val()).To(gomega.Equal(""))
		})

		ginkgo.It("should HGetAll", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.HSet(ctx, "hash", "key2", "hello2").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			m, err := adapter.Cache(time.Hour).HGetAll(ctx, "hash").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(m).To(gomega.Equal(map[string]string{"key1": "hello1", "key2": "hello2"}))
		})

		ginkgo.It("should HKeys", func() {
			hkeys := adapter.HKeys(ctx, "hash")
			gomega.Expect(hkeys.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hkeys.Val()).To(gomega.Equal([]string{}))

			hset := adapter.HSet(ctx, "hash", "key1", "hello1")
			gomega.Expect(hset.Err()).NotTo(gomega.HaveOccurred())
			hset = adapter.HSet(ctx, "hash", "key2", "hello2")
			gomega.Expect(hset.Err()).NotTo(gomega.HaveOccurred())

			hkeys = adapter.Cache(time.Hour).HKeys(ctx, "hash")
			gomega.Expect(hkeys.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hkeys.Val()).To(gomega.Equal([]string{"key1", "key2"}))
		})

		ginkgo.It("should HLen", func() {
			hSet := adapter.HSet(ctx, "hash", "key1", "hello1")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())
			hSet = adapter.HSet(ctx, "hash", "key2", "hello2")
			gomega.Expect(hSet.Err()).NotTo(gomega.HaveOccurred())

			hLen := adapter.Cache(time.Hour).HLen(ctx, "hash")
			gomega.Expect(hLen.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(hLen.Val()).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should HMGet", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1", "key2", "hello2").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.Cache(time.Hour).HMGet(ctx, "hash", "key1", "key2", "_").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]any{"hello1", "hello2", nil}))
		})

		ginkgo.It("should HVals", func() {
			err := adapter.HSet(ctx, "hash", "key1", "hello1").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.HSet(ctx, "hash", "key2", "hello2").Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			v, err := adapter.Cache(time.Hour).HVals(ctx, "hash").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(v).To(gomega.Equal([]string{"hello1", "hello2"}))

			// TODO
			// var slice []string
			// err = adapter.Cache(time.Hour).HVals(ctx, "hash").ScanSlice(&slice)
			// gomega.Expect(err).NotTo(gomega.HaveOccurred())
			// gomega.Expect(slice).To(gomega.Equal([]string{"hello1", "hello2"}))
		})
	})

	ginkgo.Describe("lists", func() {

		ginkgo.It("should LIndex", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())
			lPush = adapter.LPush(ctx, "list", "Hello")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())

			lIndex := adapter.Cache(time.Hour).LIndex(ctx, "list", 0)
			gomega.Expect(lIndex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lIndex.Val()).To(gomega.Equal("Hello"))

			lIndex = adapter.Cache(time.Hour).LIndex(ctx, "list", -1)
			gomega.Expect(lIndex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lIndex.Val()).To(gomega.Equal("World"))

			lIndex = adapter.Cache(time.Hour).LIndex(ctx, "list", 3)
			gomega.Expect(rueidis.IsRedisNil(lIndex.Err())).To(gomega.BeTrue())
			gomega.Expect(lIndex.Val()).To(gomega.Equal(""))
		})

		ginkgo.It("should LLen", func() {
			lPush := adapter.LPush(ctx, "list", "World")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())
			lPush = adapter.LPush(ctx, "list", "Hello")
			gomega.Expect(lPush.Err()).NotTo(gomega.HaveOccurred())

			lLen := adapter.Cache(time.Hour).LLen(ctx, "list")
			gomega.Expect(lLen.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lLen.Val()).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should LPos", func() {
			rPush := adapter.RPush(ctx, "list", "a")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "b")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "c")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "b")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			lPos := adapter.Cache(time.Hour).LPos(ctx, "list", "b", LPosArgs{})
			gomega.Expect(lPos.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lPos.Val()).To(gomega.Equal(int64(1)))

			lPos = adapter.Cache(time.Hour).LPos(ctx, "list", "b", LPosArgs{Rank: 2})
			gomega.Expect(lPos.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lPos.Val()).To(gomega.Equal(int64(3)))

			lPos = adapter.Cache(time.Hour).LPos(ctx, "list", "b", LPosArgs{Rank: -2})
			gomega.Expect(lPos.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lPos.Val()).To(gomega.Equal(int64(1)))

			lPos = adapter.Cache(time.Hour).LPos(ctx, "list", "b", LPosArgs{Rank: 2, MaxLen: 1})
			gomega.Expect(rueidis.IsRedisNil(lPos.Err())).To(gomega.BeTrue())

			lPos = adapter.Cache(time.Hour).LPos(ctx, "list", "z", LPosArgs{})
			gomega.Expect(rueidis.IsRedisNil(lPos.Err())).To(gomega.BeTrue())
		})

		ginkgo.It("should LRange", func() {
			rPush := adapter.RPush(ctx, "list", "one")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "two")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())
			rPush = adapter.RPush(ctx, "list", "three")
			gomega.Expect(rPush.Err()).NotTo(gomega.HaveOccurred())

			lRange := adapter.Cache(time.Hour).LRange(ctx, "list", 0, 0)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one"}))

			lRange = adapter.Cache(time.Hour).LRange(ctx, "list", -3, 2)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one", "two", "three"}))

			lRange = adapter.Cache(time.Hour).LRange(ctx, "list", -100, 100)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{"one", "two", "three"}))

			lRange = adapter.Cache(time.Hour).LRange(ctx, "list", 5, 10)
			gomega.Expect(lRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(lRange.Val()).To(gomega.Equal([]string{}))
		})
	})

	ginkgo.Describe("sets", func() {

		ginkgo.It("should SCard", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sAdd.Val()).To(gomega.Equal(int64(1)))

			sAdd = adapter.SAdd(ctx, "set", "World")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sAdd.Val()).To(gomega.Equal(int64(1)))

			sCard := adapter.Cache(time.Hour).SCard(ctx, "set")
			gomega.Expect(sCard.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sCard.Val()).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should IsMember", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sIsMember := adapter.Cache(time.Hour).SIsMember(ctx, "set", "one")
			gomega.Expect(sIsMember.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sIsMember.Val()).To(gomega.Equal(true))

			sIsMember = adapter.Cache(time.Hour).SIsMember(ctx, "set", "two")
			gomega.Expect(sIsMember.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sIsMember.Val()).To(gomega.Equal(false))
		})

		ginkgo.It("should SMIsMember", func() {
			sAdd := adapter.SAdd(ctx, "set", "one")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sMIsMember := adapter.Cache(time.Hour).SMIsMember(ctx, "set", "one", "two")
			gomega.Expect(sMIsMember.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMIsMember.Val()).To(gomega.Equal([]bool{true, false}))
		})

		ginkgo.It("should SMembers", func() {
			sAdd := adapter.SAdd(ctx, "set", "Hello")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())
			sAdd = adapter.SAdd(ctx, "set", "World")
			gomega.Expect(sAdd.Err()).NotTo(gomega.HaveOccurred())

			sMembers := adapter.Cache(time.Hour).SMembers(ctx, "set")
			gomega.Expect(sMembers.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(sMembers.Val()).To(gomega.ConsistOf([]string{"Hello", "World"}))
		})
	})

	ginkgo.Describe("sorted sets", func() {

		ginkgo.It("should ZCard", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			card, err := adapter.Cache(time.Hour).ZCard(ctx, "zset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(card).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should ZCount", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  1,
				Member: "one",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  2,
				Member: "two",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  3,
				Member: "three",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			count, err := adapter.Cache(time.Hour).ZCount(ctx, "zset", "-inf", "+inf").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(count).To(gomega.Equal(int64(3)))

			count, err = adapter.Cache(time.Hour).ZCount(ctx, "zset", "(1", "3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(count).To(gomega.Equal(int64(2)))

			count, err = adapter.Cache(time.Hour).ZLexCount(ctx, "zset", "-", "+").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(count).To(gomega.Equal(int64(3)))
		})

		ginkgo.It("should ZRangeWithScores", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 3, Member: "three"}}))

			vals, err = adapter.Cache(time.Hour).ZRangeWithScores(ctx, "zset", -2, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  3,
				Member: "three",
			}}))
		})

		ginkgo.It("should ZRangeArgs", func() {
			added, err := adapter.ZAddArgs(ctx, "zset", ZAddArgs{
				Members: []Z{
					{Score: 1, Member: "one"},
					{Score: 2, Member: "two"},
					{Score: 3, Member: "three"},
				},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(3)))

			added, err = adapter.ZAddArgs(ctx, "zset", ZAddArgs{
				NX: true,
				Members: []Z{
					{Score: 4, Member: "four"},
				},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(1)))

			added, err = adapter.ZAddArgs(ctx, "zsetxx", ZAddArgs{
				XX: true,
				Members: []Z{
					{Score: 1, Member: "one"},
				},
				Ch: true,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(added).To(gomega.Equal(int64(0)))

			score, err := adapter.ZAddArgsIncr(ctx, "zsetxx", ZAddArgs{
				Members: []Z{
					{Score: 1, Member: "one"},
				},
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(score).To(gomega.Equal(float64(1)))

			zRange, err := adapter.Cache(time.Hour).ZRangeArgs(ctx, ZRangeArgs{
				Key:     "zset",
				Start:   1,
				Stop:    4,
				ByScore: true,
				Rev:     true,
				Offset:  1,
				Count:   2,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRange).To(gomega.Equal([]string{"three", "two"}))

			zRange, err = adapter.Cache(time.Hour).ZRangeArgs(ctx, ZRangeArgs{
				Key:    "zset",
				Start:  "-",
				Stop:   "+",
				ByLex:  true,
				Rev:    true,
				Offset: 2,
				Count:  2,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRange).To(gomega.Equal([]string{"two", "one"}))

			zRange, err = adapter.Cache(time.Hour).ZRangeArgs(ctx, ZRangeArgs{
				Key:     "zset",
				Start:   "(1",
				Stop:    "(4",
				ByScore: true,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRange).To(gomega.Equal([]string{"two", "three"}))

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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(zSlice).To(gomega.Equal([]Z{
				{Score: 3, Member: "three"},
				{Score: 2, Member: "two"},
			}))
		})

		ginkgo.It("should ZRangeByScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRangeByScore := adapter.Cache(time.Hour).ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "-inf",
				Max: "+inf",
			})
			gomega.Expect(zRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByScore.Val()).To(gomega.Equal([]string{"one", "two", "three"}))

			zRangeByScore = adapter.Cache(time.Hour).ZRangeByScore(ctx, "zset", ZRangeBy{
				Min:    "-inf",
				Max:    "+inf",
				Offset: 1,
				Count:  2,
			})
			gomega.Expect(zRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByScore.Val()).To(gomega.Equal([]string{"two", "three"}))

			zRangeByScore = adapter.Cache(time.Hour).ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "1",
				Max: "2",
			})
			gomega.Expect(zRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByScore.Val()).To(gomega.Equal([]string{"one", "two"}))

			zRangeByScore = adapter.Cache(time.Hour).ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "2",
			})
			gomega.Expect(zRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByScore.Val()).To(gomega.Equal([]string{"two"}))

			zRangeByScore = adapter.Cache(time.Hour).ZRangeByScore(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "(2",
			})
			gomega.Expect(zRangeByScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByScore.Val()).To(gomega.Equal([]string{}))
		})

		ginkgo.It("should ZRangeByLex", func() {
			err := adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "a",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "b",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{
				Score:  0,
				Member: "c",
			}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRangeByLex := adapter.Cache(time.Hour).ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "-",
				Max: "+",
			})
			gomega.Expect(zRangeByLex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByLex.Val()).To(gomega.Equal([]string{"a", "b", "c"}))

			zRangeByLex = adapter.Cache(time.Hour).ZRangeByLex(ctx, "zset", ZRangeBy{
				Min:    "-",
				Max:    "+",
				Offset: 1,
				Count:  2,
			})
			gomega.Expect(zRangeByLex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByLex.Val()).To(gomega.Equal([]string{"b", "c"}))

			zRangeByLex = adapter.Cache(time.Hour).ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "[a",
				Max: "[b",
			})
			gomega.Expect(zRangeByLex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByLex.Val()).To(gomega.Equal([]string{"a", "b"}))

			zRangeByLex = adapter.Cache(time.Hour).ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "(a",
				Max: "[b",
			})
			gomega.Expect(zRangeByLex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByLex.Val()).To(gomega.Equal([]string{"b"}))

			zRangeByLex = adapter.Cache(time.Hour).ZRangeByLex(ctx, "zset", ZRangeBy{
				Min: "(a",
				Max: "(b",
			})
			gomega.Expect(zRangeByLex.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRangeByLex.Val()).To(gomega.Equal([]string{}))
		})

		ginkgo.It("should ZRangeByScoreWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "-inf",
				Max: "+inf",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 2, Member: "two"}}))

			vals, err = adapter.Cache(time.Hour).ZRangeByScoreWithScores(ctx, "zset", ZRangeBy{
				Min: "(1",
				Max: "(2",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{}))
		})

		ginkgo.It("should ZRank", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRank := adapter.Cache(time.Hour).ZRank(ctx, "zset", "three")
			gomega.Expect(zRank.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRank.Val()).To(gomega.Equal(int64(2)))

			zRank = adapter.Cache(time.Hour).ZRank(ctx, "zset", "four")
			gomega.Expect(rueidis.IsRedisNil(zRank.Err())).To(gomega.BeTrue())
			gomega.Expect(zRank.Val()).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("should ZRankWithScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRankWithScore := adapter.Cache(time.Hour).ZRankWithScore(ctx, "zset", "one")
			gomega.Expect(zRankWithScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 0, Score: 1}))

			zRankWithScore = adapter.Cache(time.Hour).ZRankWithScore(ctx, "zset", "two")
			gomega.Expect(zRankWithScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 1, Score: 2}))

			zRankWithScore = adapter.Cache(time.Hour).ZRankWithScore(ctx, "zset", "three")
			gomega.Expect(zRankWithScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 2, Score: 3}))

			zRankWithScore = adapter.Cache(time.Hour).ZRankWithScore(ctx, "zset", "four")
			gomega.Expect(zRankWithScore.Err()).To(gomega.HaveOccurred())
			gomega.Expect(zRankWithScore.Err()).To(gomega.Equal(rueidis.Nil))
		})

		ginkgo.It("should ZRevRange", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRevRange := adapter.Cache(time.Hour).ZRevRange(ctx, "zset", 0, -1)
			gomega.Expect(zRevRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRange.Val()).To(gomega.Equal([]string{"three", "two", "one"}))

			zRevRange = adapter.Cache(time.Hour).ZRevRange(ctx, "zset", 2, 3)
			gomega.Expect(zRevRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRange.Val()).To(gomega.Equal([]string{"one"}))

			zRevRange = adapter.Cache(time.Hour).ZRevRange(ctx, "zset", -2, -1)
			gomega.Expect(zRevRange.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRange.Val()).To(gomega.Equal([]string{"two", "one"}))
		})

		ginkgo.It("should ZRevRangeWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			val, err := adapter.Cache(time.Hour).ZRevRangeWithScores(ctx, "zset", 0, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]Z{{Score: 1, Member: "one"}}))

			val, err = adapter.Cache(time.Hour).ZRevRangeWithScores(ctx, "zset", -2, -1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))
		})

		ginkgo.It("should ZRevRangeByScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"three", "two", "one"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf", Offset: 1, Count: 2}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"two", "one"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "2", Min: "(1"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"two"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByScore(
				ctx, "zset", ZRangeBy{Max: "(2", Min: "(1"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{}))
		})

		ginkgo.It("should ZRevRangeByLex", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "a"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "b"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 0, Member: "c"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "+", Min: "-"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"c", "b", "a"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "+", Min: "-", Offset: 1, Count: 2}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"b", "a"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "[b", Min: "(a"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{"b"}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByLex(
				ctx, "zset", ZRangeBy{Max: "(b", Min: "(a"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]string{}))
		})

		ginkgo.It("should ZRevRangeByScoreWithScores", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
				Score:  2,
				Member: "two",
			}, {
				Score:  1,
				Member: "one",
			}}))
		})

		ginkgo.It("should ZRevRangeByScoreWithScoresMap", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			vals, err := adapter.Cache(time.Hour).ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "+inf", Min: "-inf"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{{Score: 2, Member: "two"}}))

			vals, err = adapter.Cache(time.Hour).ZRevRangeByScoreWithScores(
				ctx, "zset", ZRangeBy{Max: "(2", Min: "(1"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(vals).To(gomega.Equal([]Z{}))
		})

		ginkgo.It("should ZRevRank", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRevRank := adapter.Cache(time.Hour).ZRevRank(ctx, "zset", "one")
			gomega.Expect(zRevRank.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRank.Val()).To(gomega.Equal(int64(2)))

			zRevRank = adapter.Cache(time.Hour).ZRevRank(ctx, "zset", "four")
			gomega.Expect(rueidis.IsRedisNil(zRevRank.Err())).To(gomega.BeTrue())
			gomega.Expect(zRevRank.Val()).To(gomega.Equal(int64(0)))
		})

		ginkgo.It("should ZRevRankWithScore", func() {
			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zRevRankWithScore := adapter.Cache(time.Hour).ZRevRankWithScore(ctx, "zset", "one")
			gomega.Expect(zRevRankWithScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 2, Score: 1}))

			zRevRankWithScore = adapter.Cache(time.Hour).ZRevRankWithScore(ctx, "zset", "two")
			gomega.Expect(zRevRankWithScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 1, Score: 2}))

			zRevRankWithScore = adapter.Cache(time.Hour).ZRevRankWithScore(ctx, "zset", "three")
			gomega.Expect(zRevRankWithScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zRevRankWithScore.Result()).To(gomega.Equal(RankScore{Rank: 0, Score: 3}))

			zRevRankWithScore = adapter.Cache(time.Hour).ZRevRankWithScore(ctx, "zset", "four")
			gomega.Expect(zRevRankWithScore.Err()).To(gomega.HaveOccurred())
			gomega.Expect(zRevRankWithScore.Err()).To(gomega.Equal(rueidis.Nil))
		})

		ginkgo.It("should ZScore", func() {
			zAdd := adapter.ZAdd(ctx, "zset", Z{Score: 1.001, Member: "one"})
			gomega.Expect(zAdd.Err()).NotTo(gomega.HaveOccurred())

			zScore := adapter.Cache(time.Hour).ZScore(ctx, "zset", "one")
			gomega.Expect(zScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zScore.Val()).To(gomega.Equal(float64(1.001)))
		})

		ginkgo.It("should ZMPop", func() {

			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			key, elems, err := adapter.ZMPop(ctx, "min", 1, "zset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(key).To(gomega.Equal("zset"))
			gomega.Expect(elems).To(gomega.Equal([]Z{{
				Score:  1,
				Member: "one",
			}}))

			_, _, err = adapter.ZMPop(ctx, "min", 1, "nosuchkey").Result()
			gomega.Expect(err).To(gomega.Equal(rueidis.Nil))

			err = adapter.ZAdd(ctx, "myzset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "myzset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "myzset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			key, elems, err = adapter.ZMPop(ctx, "min", 1, "myzset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(key).To(gomega.Equal("myzset"))
			gomega.Expect(elems).To(gomega.Equal([]Z{{
				Score:  1,
				Member: "one",
			}}))

			key, elems, err = adapter.ZMPop(ctx, "max", 10, "myzset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(key).To(gomega.Equal("myzset"))
			gomega.Expect(elems).To(gomega.Equal([]Z{{
				Score:  3,
				Member: "three",
			}, {
				Score:  2,
				Member: "two",
			}}))

			err = adapter.ZAdd(ctx, "myzset2", Z{Score: 4, Member: "four"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "myzset2", Z{Score: 5, Member: "five"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "myzset2", Z{Score: 6, Member: "six"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			key, elems, err = adapter.ZMPop(ctx, "min", 10, "myzset", "myzset2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(key).To(gomega.Equal("myzset2"))
			gomega.Expect(elems).To(gomega.Equal([]Z{{
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

		ginkgo.It("should BZMPop", func() {

			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			err = adapter.ZAdd(ctx, "zset2", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset2", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			key, elems, err := adapter.BZMPop(ctx, 0, "min", 1, "zset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(key).To(gomega.Equal("zset"))
			gomega.Expect(elems).To(gomega.Equal([]Z{{
				Score:  1,
				Member: "one",
			}}))
			key, elems, err = adapter.BZMPop(ctx, 0, "max", 1, "zset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(key).To(gomega.Equal("zset"))
			gomega.Expect(elems).To(gomega.Equal([]Z{{
				Score:  3,
				Member: "three",
			}}))
			key, elems, err = adapter.BZMPop(ctx, 0, "min", 10, "zset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(key).To(gomega.Equal("zset"))
			gomega.Expect(elems).To(gomega.Equal([]Z{{
				Score:  2,
				Member: "two",
			}}))

			key, elems, err = adapter.BZMPop(ctx, 0, "max", 10, "zset2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(key).To(gomega.Equal("zset2"))
			gomega.Expect(elems).To(gomega.Equal([]Z{{
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
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			key, elems, err = adapter.BZMPop(ctx, 0, "min", 10, "myzset").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(key).To(gomega.Equal("myzset"))
			gomega.Expect(elems).To(gomega.Equal([]Z{{
				Score:  1,
				Member: "one",
			}}))

			err = adapter.ZAdd(ctx, "myzset2", Z{Score: 4, Member: "four"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "myzset2", Z{Score: 5, Member: "five"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			key, elems, err = adapter.BZMPop(ctx, 0, "min", 10, "myzset", "myzset2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(key).To(gomega.Equal("myzset2"))
			gomega.Expect(elems).To(gomega.Equal([]Z{{
				Score:  4,
				Member: "four",
			}, {
				Score:  5,
				Member: "five",
			}}))
		})

		ginkgo.It("should BZMPopBlocks", func() {
			started := make(chan bool)
			done := make(chan bool)
			go func() {
				defer ginkgo.GinkgoRecover()

				started <- true
				key, elems, err := adapter.BZMPop(ctx, 0, "min", 1, "list_list").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(key).To(gomega.Equal("list_list"))
				gomega.Expect(elems).To(gomega.Equal([]Z{{
					Score:  1,
					Member: "one",
				}}))
				done <- true
			}()
			<-started

			select {
			case <-done:
				ginkgo.Fail("BZMPop is not blocked")
			case <-time.After(time.Second):
				// ok
			}

			err := adapter.ZAdd(ctx, "list_list", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			select {
			case <-done:
				// ok
			case <-time.After(time.Second):
				ginkgo.Fail("BZMPop is still blocked")
			}
		})

		ginkgo.It("should BZMPop timeout", func() {
			_, val, err := adapter.BZMPop(ctx, time.Second, "min", 1, "list1").Result()
			gomega.Expect(err).To(gomega.Equal(rueidis.Nil))
			gomega.Expect(val).To(gomega.BeNil())

			gomega.Expect(adapter.Ping(ctx).Err()).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("should ZMScore", func() {
			zmScore := adapter.Cache(time.Hour).ZMScore(ctx, "zset", "one", "three")
			gomega.Expect(zmScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zmScore.Val()).To(gomega.HaveLen(2))
			gomega.Expect(zmScore.Val()[0]).To(gomega.Equal(float64(0)))

			err := adapter.ZAdd(ctx, "zset", Z{Score: 1, Member: "one"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 2, Member: "two"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			err = adapter.ZAdd(ctx, "zset", Z{Score: 3, Member: "three"}).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			zmScore = adapter.Cache(time.Hour).ZMScore(ctx, "zset", "one", "three")
			gomega.Expect(zmScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zmScore.Val()).To(gomega.HaveLen(2))
			gomega.Expect(zmScore.Val()[0]).To(gomega.Equal(float64(1)))

			zmScore = adapter.Cache(time.Hour).ZMScore(ctx, "zset", "four")
			gomega.Expect(zmScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zmScore.Val()).To(gomega.HaveLen(1))

			zmScore = adapter.Cache(time.Hour).ZMScore(ctx, "zset", "four", "one")
			gomega.Expect(zmScore.Err()).NotTo(gomega.HaveOccurred())
			gomega.Expect(zmScore.Val()).To(gomega.HaveLen(2))
		})
	})

	ginkgo.Describe("Geo add and radius search", func() {
		ginkgo.BeforeEach(func() {
			n, err := adapter.GeoAdd(
				ctx,
				"Sicily",
				GeoLocation{Longitude: 13.361389, Latitude: 38.115556, Name: "Palermo"},
				GeoLocation{Longitude: 15.087269, Latitude: 37.502669, Name: "Catania"},
			).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(n).To(gomega.Equal(int64(2)))
		})

		ginkgo.It("should search geo radius", func() {
			res, err := adapter.Cache(time.Hour).GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius: 200,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.HaveLen(2))
			gomega.Expect(res[0].Name).To(gomega.Equal("Palermo"))
			gomega.Expect(res[1].Name).To(gomega.Equal("Catania"))
		})

		ginkgo.It("should search geo radius with options", func() {
			res, err := adapter.Cache(time.Hour).GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.HaveLen(2))
			gomega.Expect(res[1].Name).To(gomega.Equal("Palermo"))
			gomega.Expect(res[1].Dist).To(gomega.Equal(190.4424))
			gomega.Expect(res[1].GeoHash).To(gomega.Equal(int64(3479099956230698)))
			gomega.Expect(res[1].Longitude).To(gomega.Equal(13.361389338970184))
			gomega.Expect(res[1].Latitude).To(gomega.Equal(38.115556395496299))
			gomega.Expect(res[0].Name).To(gomega.Equal("Catania"))
			gomega.Expect(res[0].Dist).To(gomega.Equal(56.4413))
			gomega.Expect(res[0].GeoHash).To(gomega.Equal(int64(3479447370796909)))
			gomega.Expect(res[0].Longitude).To(gomega.Equal(15.087267458438873))
			gomega.Expect(res[0].Latitude).To(gomega.Equal(37.50266842333162))
		})

		ginkgo.It("should search geo radius with WithDist=false", func() {
			res, err := adapter.Cache(time.Hour).GeoRadius(ctx, "Sicily", 15, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.HaveLen(2))
			gomega.Expect(res[1].Name).To(gomega.Equal("Palermo"))
			gomega.Expect(res[1].Dist).To(gomega.Equal(float64(0)))
			gomega.Expect(res[1].GeoHash).To(gomega.Equal(int64(3479099956230698)))
			gomega.Expect(res[1].Longitude).To(gomega.Equal(13.361389338970184))
			gomega.Expect(res[1].Latitude).To(gomega.Equal(38.115556395496299))
			gomega.Expect(res[0].Name).To(gomega.Equal("Catania"))
			gomega.Expect(res[0].Dist).To(gomega.Equal(float64(0)))
			gomega.Expect(res[0].GeoHash).To(gomega.Equal(int64(3479447370796909)))
			gomega.Expect(res[0].Longitude).To(gomega.Equal(15.087267458438873))
			gomega.Expect(res[0].Latitude).To(gomega.Equal(37.50266842333162))
		})

		ginkgo.It("should search geo radius by member with options", func() {
			res, err := adapter.Cache(time.Hour).GeoRadiusByMember(ctx, "Sicily", "Catania", GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
				Count:       2,
				Sort:        "ASC",
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.HaveLen(2))
			gomega.Expect(res[0].Name).To(gomega.Equal("Catania"))
			gomega.Expect(res[0].Dist).To(gomega.Equal(0.0))
			gomega.Expect(res[0].GeoHash).To(gomega.Equal(int64(3479447370796909)))
			gomega.Expect(res[0].Longitude).To(gomega.Equal(15.087267458438873))
			gomega.Expect(res[0].Latitude).To(gomega.Equal(37.50266842333162))
			gomega.Expect(res[1].Name).To(gomega.Equal("Palermo"))
			gomega.Expect(res[1].Dist).To(gomega.Equal(166.2742))
			gomega.Expect(res[1].GeoHash).To(gomega.Equal(int64(3479099956230698)))
			gomega.Expect(res[1].Longitude).To(gomega.Equal(13.361389338970184))
			gomega.Expect(res[1].Latitude).To(gomega.Equal(38.115556395496299))
		})

		ginkgo.It("should search geo radius with no results", func() {
			res, err := adapter.Cache(time.Hour).GeoRadius(ctx, "Sicily", 99, 37, GeoRadiusQuery{
				Radius:      200,
				Unit:        "km",
				WithGeoHash: true,
				WithCoord:   true,
				WithDist:    true,
			}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(res).To(gomega.HaveLen(0))
		})

		ginkgo.It("should get geo distance with unit options", func() {
			// From Redis CLI, note the difference in rounding in m vs
			// km on Redis itself.
			//
			// GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
			// GEODIST Sicily Palermo Catania m
			// "166274.15156960033"
			// GEODIST Sicily Palermo Catania km
			// "166.27415156960032"
			dist, err := adapter.Cache(time.Hour).GeoDist(ctx, "Sicily", "Palermo", "Catania", "km").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(dist).To(gomega.BeNumerically("~", 166.27, 0.01))

			dist, err = adapter.Cache(time.Hour).GeoDist(ctx, "Sicily", "Palermo", "Catania", "m").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(dist).To(gomega.BeNumerically("~", 166274.15, 0.01))

			_, err = adapter.Cache(time.Hour).GeoDist(ctx, "Sicily", "Palermo", "Catania", "mi").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			_, err = adapter.Cache(time.Hour).GeoDist(ctx, "Sicily", "Palermo", "Catania", "ft").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("should get geo hash in string representation", func() {
			hashes, err := adapter.Cache(time.Hour).GeoHash(ctx, "Sicily", "Palermo", "Catania").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(hashes).To(gomega.ConsistOf([]string{"sqc8b49rny0", "sqdtr74hyu0"}))
		})

		ginkgo.It("should return geo position", func() {
			pos, err := adapter.Cache(time.Hour).GeoPos(ctx, "Sicily", "Palermo", "Catania", "NonExisting").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(pos).To(gomega.ConsistOf([]*GeoPos{
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

		ginkgo.It("should geo search", func() {
			q := GeoSearchQuery{
				Member:    "Catania",
				BoxWidth:  400,
				BoxHeight: 100,
				BoxUnit:   "km",
				Sort:      "asc",
			}
			val, err := adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

			q.BoxHeight = 400
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania", "Palermo"}))

			q.Count = 1
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

			q.CountAny = true
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Palermo"}))

			q = GeoSearchQuery{
				Member:     "Catania",
				Radius:     100,
				RadiusUnit: "km",
				Sort:       "asc",
			}
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

			q.Radius = 400
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania", "Palermo"}))

			q.Count = 1
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

			q.CountAny = true
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Palermo"}))

			q = GeoSearchQuery{
				Longitude: 15,
				Latitude:  37,
				BoxWidth:  200,
				BoxHeight: 200,
				BoxUnit:   "km",
				Sort:      "asc",
			}
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

			q.BoxWidth, q.BoxHeight = 400, 400
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania", "Palermo"}))

			q.Count = 1
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

			q.CountAny = true
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Palermo"}))

			q = GeoSearchQuery{
				Longitude:  15,
				Latitude:   37,
				Radius:     100,
				RadiusUnit: "km",
				Sort:       "asc",
			}
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

			q.Radius = 200
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania", "Palermo"}))

			q.Count = 1
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Catania"}))

			q.CountAny = true
			val, err = adapter.Cache(time.Hour).GeoSearch(ctx, "Sicily", q).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(val).To(gomega.Equal([]string{"Palermo"}))
		})
	})

	ginkgo.Describe("marshaling/unmarshaling", func() {
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

		ginkgo.It("should convert to string", func() {
			for _, test := range convTests {
				err := adapter.Set(ctx, "key", test.value, 0).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				time.Sleep(time.Millisecond * 10)
				s, err := adapter.Cache(time.Hour).Get(ctx, "key").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(s).To(gomega.Equal(test.wanted))

				if test.dest == nil {
					continue
				}
				// TODO
				// err = adapter.Cache(time.Hour).Get(ctx, "key").Scan(test.dest)
				// gomega.Expect(err).NotTo(gomega.HaveOccurred())
				// gomega.Expect(deref(test.dest)).To(gomega.Equal(test.value))
			}
		})
	})

	ginkgo.Describe("json marshaling/unmarshaling", func() {
		ginkgo.BeforeEach(func() {
			value := &numberStruct{Number: 42}
			err := adapter.Set(ctx, "key", value, 0).Err()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("should marshal custom values using json", func() {
			s, err := adapter.Cache(time.Hour).Get(ctx, "key").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(s).To(gomega.Equal(`{"Number":42}`))
		})

		// TODO
		// ginkgo.It("should scan custom values using json", func() {
		//	value := &numberStruct{}
		//	err := adapter.Cache(time.Hour).Get(ctx, "key").Scan(value)
		//	gomega.Expect(err).NotTo(gomega.HaveOccurred())
		//	gomega.Expect(value.Number).To(gomega.Equal(42))
		// })
	})

	ginkgo.Describe("GearsCmdable", func() {
		ginkgo.BeforeEach(func() {
			gomega.Expect(adapter.FlushDB(ctx).Err()).NotTo(gomega.HaveOccurred())
			adapter.TFunctionDelete(ctx, "lib1")
		})
		// Copied from go-redis
		// https://github.com/redis/go-redis/blob/f994ff1cd96299a5c8029ae3403af7b17ef06e8a/gears_commands_test.go
		ginkgo.It("should TFunctionLoad, TFunctionLoadArgs and TFunctionDelete ", ginkgo.Label("gears", "tfunctionload"), func() {
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("OK"))
			opt := &TFunctionLoadOptions{Replace: true, Config: `{"last_update_field_name":"last_update"}`}
			resultAdd, err = adapter.TFunctionLoadArgs(ctx, libCodeWithConfig("lib1"), opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("OK"))
		})
		ginkgo.It("should TFunctionList", ginkgo.Label("gears", "tfunctionlist"), func() {
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("OK"))
			resultList, err := adapter.TFunctionList(ctx).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultList[0]["engine"]).To(gomega.BeEquivalentTo("js"))
			opt := &TFunctionListOptions{Withcode: true, Verbose: 2}
			resultListArgs, err := adapter.TFunctionListArgs(ctx, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultListArgs[0]["code"]).NotTo(gomega.BeEquivalentTo(""))
		})

		ginkgo.It("should TFCall", ginkgo.Label("gears", "tfcall"), func() {
			var resultAdd interface{}
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("OK"))
			resultAdd, err = adapter.TFCall(ctx, "lib1", "foo", 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("bar"))
		})

		ginkgo.It("should TFCallArgs", ginkgo.Label("gears", "tfcallargs"), func() {
			var resultAdd interface{}
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("OK"))
			opt := &TFCallOptions{Arguments: []string{"foo", "bar"}}
			resultAdd, err = adapter.TFCallArgs(ctx, "lib1", "foo", 0, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("bar"))
		})

		ginkgo.It("should TFCallASYNC", ginkgo.Label("gears", "TFCallASYNC"), func() {
			var resultAdd interface{}
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("OK"))
			resultAdd, err = adapter.TFCallASYNC(ctx, "lib1", "foo", 0).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("bar"))
		})

		ginkgo.It("should TFCallASYNCArgs", ginkgo.Label("gears", "TFCallASYNCargs"), func() {
			var resultAdd interface{}
			resultAdd, err := adapter.TFunctionLoad(ctx, libCode("lib1")).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("OK"))
			opt := &TFCallOptions{Arguments: []string{"foo", "bar"}}
			resultAdd, err = adapter.TFCallASYNCArgs(ctx, "lib1", "foo", 0, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo("bar"))
		})
	})
	// https://github.com/redis/go-redis/blob/master/probabilistic_test.go#L14
	ginkgo.Describe("ProbabilisticCmdable", func() {
		ctx := context.TODO()
		ginkgo.BeforeEach(func() {
			gomega.Expect(adapter.FlushDB(ctx).Err()).NotTo(gomega.HaveOccurred())
		})
		ginkgo.Describe("bloom", ginkgo.Label("bloom"), func() {
			ginkgo.It("should BFAdd", ginkgo.Label("bloom", "bfadd"), func() {
				resultAdd, err := adapter.BFAdd(ctx, "testbf1", 1).Result()

				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resultAdd).To(gomega.BeTrue())

				resultInfo, err := adapter.BFInfo(ctx, "testbf1").Result()

				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resultInfo).To(gomega.BeAssignableToTypeOf(BFInfo{}))
				gomega.Expect(resultInfo.ItemsInserted).To(gomega.BeEquivalentTo(int64(1)))
			})

			ginkgo.It("should BFCard", ginkgo.Label("bloom", "bfcard"), func() {
				// This is a probabilistic data structure, and it's not always guaranteed that we will get back
				// the exact number of inserted items, during hash collisions
				// But with such a low number of items (only 3),
				// the probability of a collision is very low, so we can expect to get back the exact number of items
				_, err := adapter.BFAdd(ctx, "testbf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				_, err = adapter.BFAdd(ctx, "testbf1", "item2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				_, err = adapter.BFAdd(ctx, "testbf1", 3).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err := adapter.BFCard(ctx, "testbf1").Result()

				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result).To(gomega.BeEquivalentTo(int64(3)))
			})

			ginkgo.It("should BFExists", ginkgo.Label("bloom", "bfexists"), func() {
				exists, err := adapter.BFExists(ctx, "testbf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(exists).To(gomega.BeFalse())

				_, err = adapter.BFAdd(ctx, "testbf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				exists, err = adapter.BFExists(ctx, "testbf1", "item1").Result()

				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(exists).To(gomega.BeTrue())
			})

			ginkgo.It("should BFInfo and BFReserve", ginkgo.Label("bloom", "bfinfo", "bfreserve"), func() {
				err := adapter.BFReserve(ctx, "testbf1", 0.001, 2000).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err := adapter.BFInfo(ctx, "testbf1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result).To(gomega.BeAssignableToTypeOf(BFInfo{}))
				gomega.Expect(result.Capacity).To(gomega.BeEquivalentTo(int64(2000)))
			})

			ginkgo.It("should BFInfoCapacity, BFInfoSize, BFInfoFilters, BFInfoItems, BFInfoExpansion, ", ginkgo.Label("bloom", "bfinfocapacity", "bfinfosize", "bfinfofilters", "bfinfoitems", "bfinfoexpansion"), func() {
				err := adapter.BFReserve(ctx, "testbf1", 0.001, 2000).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err := adapter.BFInfoCapacity(ctx, "testbf1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result.Capacity).To(gomega.BeEquivalentTo(int64(2000)))

				result, err = adapter.BFInfoItems(ctx, "testbf1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result.ItemsInserted).To(gomega.BeEquivalentTo(int64(0)))

				result, err = adapter.BFInfoSize(ctx, "testbf1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result.Size).To(gomega.BeEquivalentTo(int64(4056)))

				err = adapter.BFReserveExpansion(ctx, "testbf2", 0.001, 2000, 3).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err = adapter.BFInfoFilters(ctx, "testbf2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result.Filters).To(gomega.BeEquivalentTo(int64(1)))

				result, err = adapter.BFInfoExpansion(ctx, "testbf2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result.ExpansionRate).To(gomega.BeEquivalentTo(int64(3)))
			})

			ginkgo.It("should BFInsert", ginkgo.Label("bloom", "bfinsert"), func() {
				options := &BFInsertOptions{
					Capacity:   2000,
					Error:      0.001,
					Expansion:  3,
					NonScaling: false,
					NoCreate:   true,
				}

				resultInsert, err := adapter.BFInsert(ctx, "testbf1", options, "item1").Result()
				gomega.Expect(err).To(gomega.HaveOccurred())
				gomega.Expect(err).To(gomega.MatchError("not found"))

				options = &BFInsertOptions{
					Capacity:   2000,
					Error:      0.001,
					Expansion:  3,
					NonScaling: false,
					NoCreate:   false,
				}

				resultInsert, err = adapter.BFInsert(ctx, "testbf1", options, "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(resultInsert)).To(gomega.BeEquivalentTo(1))

				exists, err := adapter.BFExists(ctx, "testbf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(exists).To(gomega.BeTrue())

				result, err := adapter.BFInfo(ctx, "testbf1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result).To(gomega.BeAssignableToTypeOf(BFInfo{}))
				gomega.Expect(result.Capacity).To(gomega.BeEquivalentTo(int64(2000)))
				gomega.Expect(result.ExpansionRate).To(gomega.BeEquivalentTo(int64(3)))
			})

			ginkgo.It("should BFMAdd", ginkgo.Label("bloom", "bfmadd"), func() {
				resultAdd, err := adapter.BFMAdd(ctx, "testbf1", "item1", "item2", "item3").Result()

				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(resultAdd)).To(gomega.Equal(3))

				resultInfo, err := adapter.BFInfo(ctx, "testbf1").Result()

				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resultInfo).To(gomega.BeAssignableToTypeOf(BFInfo{}))
				gomega.Expect(resultInfo.ItemsInserted).To(gomega.BeEquivalentTo(int64(3)))
				resultAdd2, err := adapter.BFMAdd(ctx, "testbf1", "item1", "item2", "item4").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resultAdd2[0]).To(gomega.BeFalse())
				gomega.Expect(resultAdd2[1]).To(gomega.BeFalse())
				gomega.Expect(resultAdd2[2]).To(gomega.BeTrue())
			})

			ginkgo.It("should BFMExists", ginkgo.Label("bloom", "bfmexists"), func() {
				exist, err := adapter.BFMExists(ctx, "testbf1", "item1", "item2", "item3").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(exist)).To(gomega.Equal(3))
				gomega.Expect(exist[0]).To(gomega.BeFalse())
				gomega.Expect(exist[1]).To(gomega.BeFalse())
				gomega.Expect(exist[2]).To(gomega.BeFalse())

				_, err = adapter.BFMAdd(ctx, "testbf1", "item1", "item2", "item3").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				exist, err = adapter.BFMExists(ctx, "testbf1", "item1", "item2", "item3", "item4").Result()

				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(exist)).To(gomega.Equal(4))
				gomega.Expect(exist[0]).To(gomega.BeTrue())
				gomega.Expect(exist[1]).To(gomega.BeTrue())
				gomega.Expect(exist[2]).To(gomega.BeTrue())
				gomega.Expect(exist[3]).To(gomega.BeFalse())
			})

			ginkgo.It("should BFReserveExpansion", ginkgo.Label("bloom", "bfreserveexpansion"), func() {
				err := adapter.BFReserveExpansion(ctx, "testbf1", 0.001, 2000, 3).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err := adapter.BFInfo(ctx, "testbf1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result).To(gomega.BeAssignableToTypeOf(BFInfo{}))
				gomega.Expect(result.Capacity).To(gomega.BeEquivalentTo(int64(2000)))
				gomega.Expect(result.ExpansionRate).To(gomega.BeEquivalentTo(int64(3)))
			})

			ginkgo.It("should BFReserveNonScaling", ginkgo.Label("bloom", "bfreservenonscaling"), func() {
				err := adapter.BFReserveNonScaling(ctx, "testbfns1", 0.001, 1000).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				_, err = adapter.BFInfo(ctx, "testbfns1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			})

			ginkgo.It("should BFScanDump and BFLoadChunk", ginkgo.Label("bloom", "bfscandump", "bfloadchunk"), func() {
				err := adapter.BFReserve(ctx, "testbfsd1", 0.001, 3000).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
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
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					fd = append(fd, sd)
					sd, err = adapter.BFScanDump(ctx, "testbfsd1", sd.Iter).Result()
				}
				adapter.Del(ctx, "testbfsd1")
				for _, e := range fd {
					adapter.BFLoadChunk(ctx, "testbfsd1", e.Iter, e.Data)
				}
				infAfter := adapter.BFInfoSize(ctx, "testbfsd1")
				gomega.Expect(infBefore).To(gomega.BeEquivalentTo(infAfter))
			})

			ginkgo.It("should BFReserveWithArgs", ginkgo.Label("bloom", "bfreserveargs"), func() {
				options := &BFReserveOptions{
					Capacity:   2000,
					Error:      0.001,
					Expansion:  3,
					NonScaling: false,
				}
				err := adapter.BFReserveWithArgs(ctx, "testbf", options).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err := adapter.BFInfo(ctx, "testbf").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result).To(gomega.BeAssignableToTypeOf(BFInfo{}))
				gomega.Expect(result.Capacity).To(gomega.BeEquivalentTo(int64(2000)))
				gomega.Expect(result.ExpansionRate).To(gomega.BeEquivalentTo(int64(3)))
			})
		})

		ginkgo.Describe("cuckoo", ginkgo.Label("cuckoo"), func() {
			ginkgo.It("should CFAdd", ginkgo.Label("cuckoo", "cfadd"), func() {
				add, err := adapter.CFAdd(ctx, "testcf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(add).To(gomega.BeTrue())

				exists, err := adapter.CFExists(ctx, "testcf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(exists).To(gomega.BeTrue())

				info, err := adapter.CFInfo(ctx, "testcf1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(info).To(gomega.BeAssignableToTypeOf(CFInfo{}))
				gomega.Expect(info.NumItemsInserted).To(gomega.BeEquivalentTo(int64(1)))
			})

			ginkgo.It("should CFAddNX", ginkgo.Label("cuckoo", "cfaddnx"), func() {
				add, err := adapter.CFAddNX(ctx, "testcf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(add).To(gomega.BeTrue())

				exists, err := adapter.CFExists(ctx, "testcf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(exists).To(gomega.BeTrue())

				result, err := adapter.CFAddNX(ctx, "testcf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result).To(gomega.BeFalse())

				info, err := adapter.CFInfo(ctx, "testcf1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(info).To(gomega.BeAssignableToTypeOf(CFInfo{}))
				gomega.Expect(info.NumItemsInserted).To(gomega.BeEquivalentTo(int64(1)))
			})

			ginkgo.It("should CFCount", ginkgo.Label("cuckoo", "cfcount"), func() {
				err := adapter.CFAdd(ctx, "testcf1", "item1").Err()
				cnt, err := adapter.CFCount(ctx, "testcf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cnt).To(gomega.BeEquivalentTo(int64(1)))

				err = adapter.CFAdd(ctx, "testcf1", "item1").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				cnt, err = adapter.CFCount(ctx, "testcf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cnt).To(gomega.BeEquivalentTo(int64(2)))
			})

			ginkgo.It("should CFDel and CFExists", ginkgo.Label("cuckoo", "cfdel", "cfexists"), func() {
				err := adapter.CFAdd(ctx, "testcf1", "item1").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				exists, err := adapter.CFExists(ctx, "testcf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(exists).To(gomega.BeTrue())

				del, err := adapter.CFDel(ctx, "testcf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(del).To(gomega.BeTrue())

				exists, err = adapter.CFExists(ctx, "testcf1", "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(exists).To(gomega.BeFalse())
			})

			ginkgo.It("should CFInfo and CFReserve", ginkgo.Label("cuckoo", "cfinfo", "cfreserve"), func() {
				err := adapter.CFReserve(ctx, "testcf1", 1000).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.CFReserveExpansion(ctx, "testcfe1", 1000, 1).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.CFReserveBucketSize(ctx, "testcfbs1", 1000, 4).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.CFReserveMaxIterations(ctx, "testcfmi1", 1000, 10).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err := adapter.CFInfo(ctx, "testcf1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result).To(gomega.BeAssignableToTypeOf(CFInfo{}))
			})

			ginkgo.It("should CFScanDump and CFLoadChunk", ginkgo.Label("bloom", "cfscandump", "cfloadchunk"), func() {
				err := adapter.CFReserve(ctx, "testcfsd1", 1000).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
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
					gomega.Expect(err).NotTo(gomega.HaveOccurred())
					fd = append(fd, sd)
					sd, err = adapter.CFScanDump(ctx, "testcfsd1", sd.Iter).Result()
				}
				adapter.Del(ctx, "testcfsd1")
				for _, e := range fd {
					adapter.CFLoadChunk(ctx, "testcfsd1", e.Iter, e.Data)
				}
				infAfter := adapter.CFInfo(ctx, "testcfsd1")
				gomega.Expect(infBefore).To(gomega.BeEquivalentTo(infAfter))
			})

			ginkgo.It("should CFInfo and CFReserveWithArgs", ginkgo.Label("cuckoo", "cfinfo", "cfreserveargs"), func() {
				args := &CFReserveOptions{
					Capacity:      2048,
					BucketSize:    3,
					MaxIterations: 15,
					Expansion:     2,
				}

				err := adapter.CFReserveWithArgs(ctx, "testcf1", args).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err := adapter.CFInfo(ctx, "testcf1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result).To(gomega.BeAssignableToTypeOf(CFInfo{}))
				gomega.Expect(result.BucketSize).To(gomega.BeEquivalentTo(int64(3)))
				gomega.Expect(result.MaxIteration).To(gomega.BeEquivalentTo(int64(15)))
				gomega.Expect(result.ExpansionRate).To(gomega.BeEquivalentTo(int64(2)))
			})

			ginkgo.It("should CFInsert", ginkgo.Label("cuckoo", "cfinsert"), func() {
				args := &CFInsertOptions{
					Capacity: 3000,
					NoCreate: true,
				}

				result, err := adapter.CFInsert(ctx, "testcf1", args, "item1", "item2", "item3").Result()
				gomega.Expect(err).To(gomega.HaveOccurred())

				args = &CFInsertOptions{
					Capacity: 3000,
					NoCreate: false,
				}

				result, err = adapter.CFInsert(ctx, "testcf1", args, "item1", "item2", "item3").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(result)).To(gomega.BeEquivalentTo(3))
			})

			ginkgo.It("should CFInsertNX", ginkgo.Label("cuckoo", "cfinsertnx"), func() {
				args := &CFInsertOptions{
					Capacity: 3000,
					NoCreate: true,
				}

				result, err := adapter.CFInsertNX(ctx, "testcf1", args, "item1", "item2", "item2").Result()
				gomega.Expect(err).To(gomega.HaveOccurred())

				args = &CFInsertOptions{
					Capacity: 3000,
					NoCreate: false,
				}

				result, err = adapter.CFInsertNX(ctx, "testcf2", args, "item1", "item2", "item2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(result)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(result[0]).To(gomega.BeEquivalentTo(int64(1)))
				gomega.Expect(result[1]).To(gomega.BeEquivalentTo(int64(1)))
				gomega.Expect(result[2]).To(gomega.BeEquivalentTo(int64(0)))
			})

			ginkgo.It("should CFMexists", ginkgo.Label("cuckoo", "cfmexists"), func() {
				err := adapter.CFInsert(ctx, "testcf1", nil, "item1", "item2", "item3").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err := adapter.CFMExists(ctx, "testcf1", "item1", "item2", "item3", "item4").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(result)).To(gomega.BeEquivalentTo(4))
				gomega.Expect(result[0]).To(gomega.BeTrue())
				gomega.Expect(result[1]).To(gomega.BeTrue())
				gomega.Expect(result[2]).To(gomega.BeTrue())
				gomega.Expect(result[3]).To(gomega.BeFalse())
			})
		})

		ginkgo.Describe("CMS", ginkgo.Label("cms"), func() {
			ginkgo.It("should CMSIncrBy", ginkgo.Label("cms", "cmsincrby"), func() {
				err := adapter.CMSInitByDim(ctx, "testcms1", 5, 10).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err := adapter.CMSIncrBy(ctx, "testcms1", "item1", 1, "item2", 2, "item3", 3).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(result)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(result[0]).To(gomega.BeEquivalentTo(int64(1)))
				gomega.Expect(result[1]).To(gomega.BeEquivalentTo(int64(2)))
				gomega.Expect(result[2]).To(gomega.BeEquivalentTo(int64(3)))
			})

			ginkgo.It("should CMSInitByDim and CMSInfo", ginkgo.Label("cms", "cmsinitbydim", "cmsinfo"), func() {
				err := adapter.CMSInitByDim(ctx, "testcms1", 5, 10).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				info, err := adapter.CMSInfo(ctx, "testcms1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				gomega.Expect(info).To(gomega.BeAssignableToTypeOf(CMSInfo{}))
				gomega.Expect(info.Width).To(gomega.BeEquivalentTo(int64(5)))
				gomega.Expect(info.Depth).To(gomega.BeEquivalentTo(int64(10)))
			})

			ginkgo.It("should CMSInitByProb", ginkgo.Label("cms", "cmsinitbyprob"), func() {
				err := adapter.CMSInitByProb(ctx, "testcms1", 0.002, 0.01).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				info, err := adapter.CMSInfo(ctx, "testcms1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(info).To(gomega.BeAssignableToTypeOf(CMSInfo{}))
			})

			ginkgo.It("should CMSMerge, CMSMergeWithWeight and CMSQuery", ginkgo.Label("cms", "cmsmerge", "cmsquery"), func() {
				err := adapter.CMSMerge(ctx, "destCms1", "testcms2", "testcms3").Err()
				gomega.Expect(err).To(gomega.HaveOccurred())
				gomega.Expect(err).To(gomega.MatchError("CMS: key does not exist"))

				err = adapter.CMSInitByDim(ctx, "destCms1", 5, 10).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.CMSInitByDim(ctx, "destCms2", 5, 10).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.CMSInitByDim(ctx, "cms1", 2, 20).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.CMSInitByDim(ctx, "cms2", 3, 20).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.CMSMerge(ctx, "destCms1", "cms1", "cms2").Err()
				gomega.Expect(err).To(gomega.MatchError("CMS: width/depth is not equal"))

				adapter.Del(ctx, "cms1", "cms2")

				err = adapter.CMSInitByDim(ctx, "cms1", 5, 10).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.CMSInitByDim(ctx, "cms2", 5, 10).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				adapter.CMSIncrBy(ctx, "cms1", "item1", 1, "item2", 2)
				adapter.CMSIncrBy(ctx, "cms2", "item2", 2, "item3", 3)

				err = adapter.CMSMerge(ctx, "destCms1", "cms1", "cms2").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err := adapter.CMSQuery(ctx, "destCms1", "item1", "item2", "item3").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(result)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(result[0]).To(gomega.BeEquivalentTo(int64(1)))
				gomega.Expect(result[1]).To(gomega.BeEquivalentTo(int64(4)))
				gomega.Expect(result[2]).To(gomega.BeEquivalentTo(int64(3)))

				sourceSketches := map[string]int64{
					"cms1": 1,
					"cms2": 2,
				}
				err = adapter.CMSMergeWithWeight(ctx, "destCms2", sourceSketches).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				result, err = adapter.CMSQuery(ctx, "destCms2", "item1", "item2", "item3").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(result)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(result[0]).To(gomega.BeEquivalentTo(int64(1)))
				gomega.Expect(result[1]).To(gomega.BeEquivalentTo(int64(6)))
				gomega.Expect(result[2]).To(gomega.BeEquivalentTo(int64(6)))
			})
		})

		ginkgo.Describe("TopK", ginkgo.Label("topk"), func() {
			ginkgo.It("should TopKReserve, TopKInfo, TopKAdd, TopKQuery, TopKCount, TopKIncrBy, TopKList, TopKListWithCount", ginkgo.Label("topk", "topkreserve", "topkinfo", "topkadd", "topkquery", "topkcount", "topkincrby", "topklist", "topklistwithcount"), func() {
				err := adapter.TopKReserve(ctx, "topk1", 3).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				resultInfo, err := adapter.TopKInfo(ctx, "topk1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resultInfo.K).To(gomega.BeEquivalentTo(int64(3)))

				resultAdd, err := adapter.TopKAdd(ctx, "topk1", "item1", "item2", 3, "item1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(resultAdd)).To(gomega.BeEquivalentTo(int64(4)))

				resultQuery, err := adapter.TopKQuery(ctx, "topk1", "item1", "item2", 4, 3).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(resultQuery)).To(gomega.BeEquivalentTo(4))
				gomega.Expect(resultQuery[0]).To(gomega.BeTrue())
				gomega.Expect(resultQuery[1]).To(gomega.BeTrue())
				gomega.Expect(resultQuery[2]).To(gomega.BeFalse())
				gomega.Expect(resultQuery[3]).To(gomega.BeTrue())

				resultCount, err := adapter.TopKCount(ctx, "topk1", "item1", "item2", "item3").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(resultCount)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(resultCount[0]).To(gomega.BeEquivalentTo(int64(2)))
				gomega.Expect(resultCount[1]).To(gomega.BeEquivalentTo(int64(1)))
				gomega.Expect(resultCount[2]).To(gomega.BeEquivalentTo(int64(0)))

				resultIncr, err := adapter.TopKIncrBy(ctx, "topk1", "item1", 5, "item2", 10).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(resultIncr)).To(gomega.BeEquivalentTo(2))

				resultCount, err = adapter.TopKCount(ctx, "topk1", "item1", "item2", "item3").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(resultCount)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(resultCount[0]).To(gomega.BeEquivalentTo(int64(7)))
				gomega.Expect(resultCount[1]).To(gomega.BeEquivalentTo(int64(11)))
				gomega.Expect(resultCount[2]).To(gomega.BeEquivalentTo(int64(0)))

				resultList, err := adapter.TopKList(ctx, "topk1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(resultList)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(resultList).To(gomega.ContainElements("item2", "item1", "3"))

				resultListWithCount, err := adapter.TopKListWithCount(ctx, "topk1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(resultListWithCount)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(resultListWithCount["3"]).To(gomega.BeEquivalentTo(int64(1)))
				gomega.Expect(resultListWithCount["item1"]).To(gomega.BeEquivalentTo(int64(7)))
				gomega.Expect(resultListWithCount["item2"]).To(gomega.BeEquivalentTo(int64(11)))
			})

			ginkgo.It("should TopKReserveWithOptions", ginkgo.Label("topk", "topkreservewithoptions"), func() {
				err := adapter.TopKReserveWithOptions(ctx, "topk1", 3, 1500, 8, 0.5).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				resultInfo, err := adapter.TopKInfo(ctx, "topk1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resultInfo.K).To(gomega.BeEquivalentTo(int64(3)))
				gomega.Expect(resultInfo.Width).To(gomega.BeEquivalentTo(int64(1500)))
				gomega.Expect(resultInfo.Depth).To(gomega.BeEquivalentTo(int64(8)))
				gomega.Expect(resultInfo.Decay).To(gomega.BeEquivalentTo(0.5))
			})
		})

		ginkgo.Describe("t-digest", ginkgo.Label("tdigest"), func() {
			ginkgo.It("should TDigestAdd, TDigestCreate, TDigestInfo, TDigestByRank, TDigestByRevRank, TDigestCDF, TDigestMax, TDigestMin, TDigestQuantile, TDigestRank, TDigestRevRank, TDigestTrimmedMean, TDigestReset, ", ginkgo.Label("tdigest", "tdigestadd", "tdigestcreate", "tdigestinfo", "tdigestbyrank", "tdigestbyrevrank", "tdigestcdf", "tdigestmax", "tdigestmin", "tdigestquantile", "tdigestrank", "tdigestrevrank", "tdigesttrimmedmean", "tdigestreset"), func() {
				err := adapter.TDigestCreate(ctx, "tdigest1").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				info, err := adapter.TDigestInfo(ctx, "tdigest1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(info.Observations).To(gomega.BeEquivalentTo(int64(0)))

				// Test with empty sketch
				byRank, err := adapter.TDigestByRank(ctx, "tdigest1", 0, 1, 2, 3).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(byRank)).To(gomega.BeEquivalentTo(4))

				byRevRank, err := adapter.TDigestByRevRank(ctx, "tdigest1", 0, 1, 2).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(byRevRank)).To(gomega.BeEquivalentTo(3))

				cdf, err := adapter.TDigestCDF(ctx, "tdigest1", 15, 35, 70).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(cdf)).To(gomega.BeEquivalentTo(3))

				max, err := adapter.TDigestMax(ctx, "tdigest1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(math.IsNaN(max)).To(gomega.BeTrue())

				min, err := adapter.TDigestMin(ctx, "tdigest1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(math.IsNaN(min)).To(gomega.BeTrue())

				quantile, err := adapter.TDigestQuantile(ctx, "tdigest1", 0.1, 0.2).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(quantile)).To(gomega.BeEquivalentTo(2))

				rank, err := adapter.TDigestRank(ctx, "tdigest1", 10, 20).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(rank)).To(gomega.BeEquivalentTo(2))

				revRank, err := adapter.TDigestRevRank(ctx, "tdigest1", 10, 20).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(revRank)).To(gomega.BeEquivalentTo(2))

				trimmedMean, err := adapter.TDigestTrimmedMean(ctx, "tdigest1", 0.1, 0.6).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(math.IsNaN(trimmedMean)).To(gomega.BeTrue())

				// Add elements
				err = adapter.TDigestAdd(ctx, "tdigest1", 10, 20, 30, 40, 50, 60, 70, 80, 90, 100).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				info, err = adapter.TDigestInfo(ctx, "tdigest1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(info.Observations).To(gomega.BeEquivalentTo(int64(10)))

				byRank, err = adapter.TDigestByRank(ctx, "tdigest1", 0, 1, 2).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(byRank)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(byRank[0]).To(gomega.BeEquivalentTo(float64(10)))
				gomega.Expect(byRank[1]).To(gomega.BeEquivalentTo(float64(20)))
				gomega.Expect(byRank[2]).To(gomega.BeEquivalentTo(float64(30)))

				byRevRank, err = adapter.TDigestByRevRank(ctx, "tdigest1", 0, 1, 2).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(byRevRank)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(byRevRank[0]).To(gomega.BeEquivalentTo(float64(100)))
				gomega.Expect(byRevRank[1]).To(gomega.BeEquivalentTo(float64(90)))
				gomega.Expect(byRevRank[2]).To(gomega.BeEquivalentTo(float64(80)))

				cdf, err = adapter.TDigestCDF(ctx, "tdigest1", 15, 35, 70).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(cdf)).To(gomega.BeEquivalentTo(3))
				gomega.Expect(cdf[0]).To(gomega.BeEquivalentTo(0.1))
				gomega.Expect(cdf[1]).To(gomega.BeEquivalentTo(0.3))
				gomega.Expect(cdf[2]).To(gomega.BeEquivalentTo(0.65))

				max, err = adapter.TDigestMax(ctx, "tdigest1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(max).To(gomega.BeEquivalentTo(float64(100)))

				min, err = adapter.TDigestMin(ctx, "tdigest1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(min).To(gomega.BeEquivalentTo(float64(10)))

				quantile, err = adapter.TDigestQuantile(ctx, "tdigest1", 0.1, 0.2).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(quantile)).To(gomega.BeEquivalentTo(2))
				gomega.Expect(quantile[0]).To(gomega.BeEquivalentTo(float64(20)))
				gomega.Expect(quantile[1]).To(gomega.BeEquivalentTo(float64(30)))

				rank, err = adapter.TDigestRank(ctx, "tdigest1", 10, 20).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(rank)).To(gomega.BeEquivalentTo(2))
				gomega.Expect(rank[0]).To(gomega.BeEquivalentTo(int64(0)))
				gomega.Expect(rank[1]).To(gomega.BeEquivalentTo(int64(1)))

				revRank, err = adapter.TDigestRevRank(ctx, "tdigest1", 10, 20).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(len(revRank)).To(gomega.BeEquivalentTo(2))
				gomega.Expect(revRank[0]).To(gomega.BeEquivalentTo(int64(9)))
				gomega.Expect(revRank[1]).To(gomega.BeEquivalentTo(int64(8)))

				trimmedMean, err = adapter.TDigestTrimmedMean(ctx, "tdigest1", 0.1, 0.6).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(trimmedMean).To(gomega.BeEquivalentTo(float64(40)))

				reset, err := adapter.TDigestReset(ctx, "tdigest1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(reset).To(gomega.BeEquivalentTo("OK"))
			})

			ginkgo.It("should TDigestCreateWithCompression", ginkgo.Label("tdigest", "tcreatewithcompression"), func() {
				err := adapter.TDigestCreateWithCompression(ctx, "tdigest1", 2000).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				info, err := adapter.TDigestInfo(ctx, "tdigest1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(info.Compression).To(gomega.BeEquivalentTo(int64(2000)))
			})

			ginkgo.It("should TDigestMerge", ginkgo.Label("tdigest", "tmerge"), func() {
				err := adapter.TDigestCreate(ctx, "tdigest1").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.TDigestAdd(ctx, "tdigest1", 10, 20, 30, 40, 50, 60, 70, 80, 90, 100).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.TDigestCreate(ctx, "tdigest2").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.TDigestAdd(ctx, "tdigest2", 15, 25, 35, 45, 55, 65, 75, 85, 95, 105).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				err = adapter.TDigestCreate(ctx, "tdigest3").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				err = adapter.TDigestAdd(ctx, "tdigest3", 50, 60, 70, 80, 90, 100, 110, 120, 130, 140).Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				options := &TDigestMergeOptions{
					Compression: 1000,
					Override:    false,
				}
				err = adapter.TDigestMerge(ctx, "tdigest1", options, "tdigest2", "tdigest3").Err()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

				info, err := adapter.TDigestInfo(ctx, "tdigest1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(info.Observations).To(gomega.BeEquivalentTo(int64(30)))
				gomega.Expect(info.Compression).To(gomega.BeEquivalentTo(int64(1000)))

				max, err := adapter.TDigestMax(ctx, "tdigest1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(max).To(gomega.BeEquivalentTo(float64(140)))
			})
		})
	})
	ginkgo.Describe("RedisTimeseries commands", ginkgo.Label("timeseries"), func() {
		ctx := context.TODO()

		ginkgo.BeforeEach(func() {
			gomega.Expect(adapter.FlushDB(ctx).Err()).NotTo(gomega.HaveOccurred())
		})

		ginkgo.It("should TSCreate and TSCreateWithArgs", ginkgo.Label("timeseries", "tscreate", "tscreateWithArgs"), func() {
			result, err := adapter.TSCreate(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo("OK"))
			// Test TSCreateWithArgs
			opt := &TSOptions{Retention: 5}
			result, err = adapter.TSCreateWithArgs(ctx, "2", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"Redis": "Labs"}}
			result, err = adapter.TSCreateWithArgs(ctx, "3", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"Time": "Series"}, Retention: 20}
			result, err = adapter.TSCreateWithArgs(ctx, "4", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo("OK"))
			resultInfo, err := adapter.TSInfo(ctx, "4").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["labels"].(map[string]interface{})["Time"]).To(gomega.BeEquivalentTo("Series"))
			// Test chunk size
			opt = &TSOptions{ChunkSize: 128}
			result, err = adapter.TSCreateWithArgs(ctx, "ts-cs-1", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo("OK"))
			resultInfo, err = adapter.TSInfo(ctx, "ts-cs-1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["chunkSize"]).To(gomega.BeEquivalentTo(128))
			// Test duplicate policy
			duplicate_policies := []string{"BLOCK", "LAST", "FIRST", "MIN", "MAX"}
			for _, dup := range duplicate_policies {
				keyName := "ts-dup-" + dup
				opt = &TSOptions{DuplicatePolicy: dup}
				result, err = adapter.TSCreateWithArgs(ctx, keyName, opt).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(result).To(gomega.BeEquivalentTo("OK"))
				resultInfo, err = adapter.TSInfo(ctx, keyName).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(strings.ToUpper(resultInfo["duplicatePolicy"].(string))).To(gomega.BeEquivalentTo(dup))

			}
		})
		ginkgo.It("should TSAdd and TSAddWithArgs", ginkgo.Label("timeseries", "tsadd", "tsaddWithArgs"), func() {
			result, err := adapter.TSAdd(ctx, "1", 1, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(1))
			// Test TSAddWithArgs
			opt := &TSOptions{Retention: 10}
			result, err = adapter.TSAddWithArgs(ctx, "2", 2, 3, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(2))
			opt = &TSOptions{Labels: map[string]string{"Redis": "Labs"}}
			result, err = adapter.TSAddWithArgs(ctx, "3", 3, 2, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(3))
			opt = &TSOptions{Labels: map[string]string{"Redis": "Labs", "Time": "Series"}, Retention: 10}
			result, err = adapter.TSAddWithArgs(ctx, "4", 4, 2, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(4))
			resultInfo, err := adapter.TSInfo(ctx, "4").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["labels"].(map[string]interface{})["Time"]).To(gomega.BeEquivalentTo("Series"))
			// Test chunk size
			opt = &TSOptions{ChunkSize: 128}
			result, err = adapter.TSAddWithArgs(ctx, "ts-cs-1", 1, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(1))
			resultInfo, err = adapter.TSInfo(ctx, "ts-cs-1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["chunkSize"]).To(gomega.BeEquivalentTo(128))
			// Test duplicate policy
			// LAST
			opt = &TSOptions{DuplicatePolicy: "LAST"}
			result, err = adapter.TSAddWithArgs(ctx, "tsal-1", 1, 5, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(1))
			result, err = adapter.TSAddWithArgs(ctx, "tsal-1", 1, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(1))
			resultGet, err := adapter.TSGet(ctx, "tsal-1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet.Value).To(gomega.BeEquivalentTo(10))
			// FIRST
			opt = &TSOptions{DuplicatePolicy: "FIRST"}
			result, err = adapter.TSAddWithArgs(ctx, "tsaf-1", 1, 5, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(1))
			result, err = adapter.TSAddWithArgs(ctx, "tsaf-1", 1, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(1))
			resultGet, err = adapter.TSGet(ctx, "tsaf-1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet.Value).To(gomega.BeEquivalentTo(5))
			// MAX
			opt = &TSOptions{DuplicatePolicy: "MAX"}
			result, err = adapter.TSAddWithArgs(ctx, "tsam-1", 1, 5, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(1))
			result, err = adapter.TSAddWithArgs(ctx, "tsam-1", 1, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(1))
			resultGet, err = adapter.TSGet(ctx, "tsam-1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet.Value).To(gomega.BeEquivalentTo(10))
			// MIN
			opt = &TSOptions{DuplicatePolicy: "MIN"}
			result, err = adapter.TSAddWithArgs(ctx, "tsami-1", 1, 5, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(1))
			result, err = adapter.TSAddWithArgs(ctx, "tsami-1", 1, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo(1))
			resultGet, err = adapter.TSGet(ctx, "tsami-1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet.Value).To(gomega.BeEquivalentTo(5))
		})

		ginkgo.It("should TSAlter", ginkgo.Label("timeseries", "tsalter"), func() {
			result, err := adapter.TSCreate(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo("OK"))
			resultInfo, err := adapter.TSInfo(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["retentionTime"]).To(gomega.BeEquivalentTo(0))

			opt := &TSAlterOptions{Retention: 10}
			resultAlter, err := adapter.TSAlter(ctx, "1", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAlter).To(gomega.BeEquivalentTo("OK"))

			resultInfo, err = adapter.TSInfo(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["retentionTime"]).To(gomega.BeEquivalentTo(10))

			resultInfo, err = adapter.TSInfo(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["labels"]).To(gomega.BeEquivalentTo(map[string]interface{}{}))

			opt = &TSAlterOptions{Labels: map[string]string{"Time": "Series"}}
			resultAlter, err = adapter.TSAlter(ctx, "1", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAlter).To(gomega.BeEquivalentTo("OK"))

			resultInfo, err = adapter.TSInfo(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["labels"].(map[string]interface{})["Time"]).To(gomega.BeEquivalentTo("Series"))
			gomega.Expect(resultInfo["retentionTime"]).To(gomega.BeEquivalentTo(10))
			gomega.Expect(resultInfo["duplicatePolicy"]).To(gomega.BeNil())
			opt = &TSAlterOptions{DuplicatePolicy: "min"}
			resultAlter, err = adapter.TSAlter(ctx, "1", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAlter).To(gomega.BeEquivalentTo("OK"))

			resultInfo, err = adapter.TSInfo(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["duplicatePolicy"]).To(gomega.BeEquivalentTo("min"))
		})

		ginkgo.It("should TSCreateRule and TSDeleteRule", ginkgo.Label("timeseries", "tscreaterule", "tsdeleterule"), func() {
			result, err := adapter.TSCreate(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo("OK"))
			result, err = adapter.TSCreate(ctx, "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo("OK"))
			result, err = adapter.TSCreateRule(ctx, "1", "2", Avg, 100).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo("OK"))
			for i := 0; i < 50; i++ {
				resultAdd, err := adapter.TSAdd(ctx, "1", 100+i*2, 1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resultAdd).To(gomega.BeEquivalentTo(100 + i*2))
				resultAdd, err = adapter.TSAdd(ctx, "1", 100+i*2+1, 2).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resultAdd).To(gomega.BeEquivalentTo(100 + i*2 + 1))

			}
			resultAdd, err := adapter.TSAdd(ctx, "1", 100*2, 1.5).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultAdd).To(gomega.BeEquivalentTo(100 * 2))
			resultGet, err := adapter.TSGet(ctx, "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet.Value).To(gomega.BeEquivalentTo(1.5))
			gomega.Expect(resultGet.Timestamp).To(gomega.BeEquivalentTo(100))

			resultDeleteRule, err := adapter.TSDeleteRule(ctx, "1", "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultDeleteRule).To(gomega.BeEquivalentTo("OK"))
			resultInfo, err := adapter.TSInfo(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["rules"]).To(gomega.BeEquivalentTo(map[string]interface{}{}))
		})

		ginkgo.It("should TSIncrBy, TSIncrByWithArgs, TSDecrBy and TSDecrByWithArgs", ginkgo.Label("timeseries", "tsincrby", "tsdecrby", "tsincrbyWithArgs", "tsdecrbyWithArgs"), func() {
			for i := 0; i < 100; i++ {
				_, err := adapter.TSIncrBy(ctx, "1", 1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			}
			result, err := adapter.TSGet(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result.Value).To(gomega.BeEquivalentTo(100))

			for i := 0; i < 100; i++ {
				_, err := adapter.TSDecrBy(ctx, "1", 1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			}
			result, err = adapter.TSGet(ctx, "1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result.Value).To(gomega.BeEquivalentTo(0))

			opt := &TSIncrDecrOptions{Timestamp: 5}
			_, err = adapter.TSIncrByWithArgs(ctx, "2", 1.5, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			result, err = adapter.TSGet(ctx, "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result.Timestamp).To(gomega.BeEquivalentTo(5))
			gomega.Expect(result.Value).To(gomega.BeEquivalentTo(1.5))

			opt = &TSIncrDecrOptions{Timestamp: 7}
			_, err = adapter.TSIncrByWithArgs(ctx, "2", 2.25, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			result, err = adapter.TSGet(ctx, "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result.Timestamp).To(gomega.BeEquivalentTo(7))
			gomega.Expect(result.Value).To(gomega.BeEquivalentTo(3.75))

			opt = &TSIncrDecrOptions{Timestamp: 15}
			_, err = adapter.TSDecrByWithArgs(ctx, "2", 1.5, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			result, err = adapter.TSGet(ctx, "2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result.Timestamp).To(gomega.BeEquivalentTo(15))
			gomega.Expect(result.Value).To(gomega.BeEquivalentTo(2.25))

			// Test chunk size INCRBY
			opt = &TSIncrDecrOptions{ChunkSize: 128}
			_, err = adapter.TSIncrByWithArgs(ctx, "3", 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			resultInfo, err := adapter.TSInfo(ctx, "3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["chunkSize"]).To(gomega.BeEquivalentTo(128))

			// Test chunk size DECRBY
			opt = &TSIncrDecrOptions{ChunkSize: 128}
			_, err = adapter.TSDecrByWithArgs(ctx, "4", 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			resultInfo, err = adapter.TSInfo(ctx, "4").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultInfo["chunkSize"]).To(gomega.BeEquivalentTo(128))
		})

		ginkgo.It("should TSGet", ginkgo.Label("timeseries", "tsget"), func() {
			opt := &TSOptions{DuplicatePolicy: "max"}
			resultGet, err := adapter.TSAddWithArgs(ctx, "foo", 2265985, 151, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet).To(gomega.BeEquivalentTo(2265985))
			result, err := adapter.TSGet(ctx, "foo").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result.Timestamp).To(gomega.BeEquivalentTo(2265985))
			gomega.Expect(result.Value).To(gomega.BeEquivalentTo(151))
		})

		ginkgo.It("should TSGet Latest", ginkgo.Label("timeseries", "tsgetlatest"), func() {
			resultGet, err := adapter.TSCreate(ctx, "tsgl-1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet).To(gomega.BeEquivalentTo("OK"))
			resultGet, err = adapter.TSCreate(ctx, "tsgl-2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet).To(gomega.BeEquivalentTo("OK"))
			resultGet, err = adapter.TSCreateRule(ctx, "tsgl-1", "tsgl-2", Sum, 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet).To(gomega.BeEquivalentTo("OK"))
			_, err = adapter.TSAdd(ctx, "tsgl-1", 1, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "tsgl-1", 2, 3).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "tsgl-1", 11, 7).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "tsgl-1", 13, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			result, errGet := adapter.TSGet(ctx, "tsgl-2").Result()
			gomega.Expect(errGet).NotTo(gomega.HaveOccurred())
			gomega.Expect(result.Timestamp).To(gomega.BeEquivalentTo(0))
			gomega.Expect(result.Value).To(gomega.BeEquivalentTo(4))
			result, errGet = adapter.TSGetWithArgs(ctx, "tsgl-2", &TSGetOptions{Latest: true}).Result()
			gomega.Expect(errGet).NotTo(gomega.HaveOccurred())
			gomega.Expect(result.Timestamp).To(gomega.BeEquivalentTo(10))
			gomega.Expect(result.Value).To(gomega.BeEquivalentTo(8))
		})

		ginkgo.It("should TSInfo", ginkgo.Label("timeseries", "tsinfo"), func() {
			resultGet, err := adapter.TSAdd(ctx, "foo", 2265985, 151).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet).To(gomega.BeEquivalentTo(2265985))
			result, err := adapter.TSInfo(ctx, "foo").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["firstTimestamp"]).To(gomega.BeEquivalentTo(2265985))
		})

		ginkgo.It("should TSMAdd", ginkgo.Label("timeseries", "tsmadd"), func() {
			resultGet, err := adapter.TSCreate(ctx, "a").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultGet).To(gomega.BeEquivalentTo("OK"))
			ktvSlices := make([][]interface{}, 3)
			for i := 0; i < 3; i++ {
				ktvSlices[i] = make([]interface{}, 3)
				ktvSlices[i][0] = "a"
				for j := 1; j < 3; j++ {
					ktvSlices[i][j] = (i + j) * j
				}
			}
			result, err := adapter.TSMAdd(ctx, ktvSlices).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo([]int64{1, 2, 3}))
		})

		ginkgo.It("should TSMGet and TSMGetWithArgs", ginkgo.Label("timeseries", "tsmget", "tsmgetWithArgs"), func() {
			opt := &TSOptions{Labels: map[string]string{"Test": "This"}}
			resultCreate, err := adapter.TSCreateWithArgs(ctx, "a", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"Test": "This", "Taste": "That"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			_, err = adapter.TSAdd(ctx, "a", "*", 15).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "b", "*", 25).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			result, err := adapter.TSMGet(ctx, []string{"Test=This"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][1].([]interface{})[1]).To(gomega.BeEquivalentTo(15))
			gomega.Expect(result["b"][1].([]interface{})[1]).To(gomega.BeEquivalentTo(25))
			mgetOpt := &TSMGetOptions{WithLabels: true}
			result, err = adapter.TSMGetWithArgs(ctx, []string{"Test=This"}, mgetOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["b"][0]).To(gomega.BeEquivalentTo(map[string]interface{}{"Test": "This", "Taste": "That"}))

			resultCreate, err = adapter.TSCreate(ctx, "c").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"is_compaction": "true"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "d", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			resultCreateRule, err := adapter.TSCreateRule(ctx, "c", "d", Sum, 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreateRule).To(gomega.BeEquivalentTo("OK"))
			_, err = adapter.TSAdd(ctx, "c", 1, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 2, 3).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 11, 7).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 13, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			result, err = adapter.TSMGet(ctx, []string{"is_compaction=true"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["d"][1]).To(gomega.BeEquivalentTo([]interface{}{int64(0), 4.0}))
			mgetOpt = &TSMGetOptions{Latest: true}
			result, err = adapter.TSMGetWithArgs(ctx, []string{"is_compaction=true"}, mgetOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["d"][1]).To(gomega.BeEquivalentTo([]interface{}{int64(10), 8.0}))
		})

		ginkgo.It("should TSQueryIndex", ginkgo.Label("timeseries", "tsqueryindex"), func() {
			opt := &TSOptions{Labels: map[string]string{"Test": "This"}}
			resultCreate, err := adapter.TSCreateWithArgs(ctx, "a", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"Test": "This", "Taste": "That"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			result, err := adapter.TSQueryIndex(ctx, []string{"Test=This"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(2))
			result, err = adapter.TSQueryIndex(ctx, []string{"Taste=That"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(1))
		})

		ginkgo.It("should TSDel and TSRange", ginkgo.Label("timeseries", "tsdel", "tsrange"), func() {
			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i, float64(i%7)).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			}
			resultDelete, err := adapter.TSDel(ctx, "a", 0, 21).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultDelete).To(gomega.BeEquivalentTo(22))

			resultRange, err := adapter.TSRange(ctx, "a", 0, 21).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange).To(gomega.BeEquivalentTo([]TSTimestampValue{}))

			resultRange, err = adapter.TSRange(ctx, "a", 22, 22).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 22, Value: 1}))
		})

		ginkgo.It("should TSRange, TSRangeWithArgs", ginkgo.Label("timeseries", "tsrange", "tsrangeWithArgs"), func() {
			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i, float64(i%7)).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

			}
			result, err := adapter.TSRange(ctx, "a", 0, 200).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(100))
			for i := 0; i < 100; i++ {
				adapter.TSAdd(ctx, "a", i+200, float64(i%7))
			}
			result, err = adapter.TSRange(ctx, "a", 0, 500).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(200))
			fts := make([]int, 0)
			for i := 10; i < 20; i++ {
				fts = append(fts, i)
			}
			opt := &TSRangeOptions{FilterByTS: fts, FilterByValue: []int{1, 2}}
			result, err = adapter.TSRangeWithArgs(ctx, "a", 0, 500, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(2))
			opt = &TSRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "+"}
			result, err = adapter.TSRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo([]TSTimestampValue{{Timestamp: 0, Value: 10}, {Timestamp: 10, Value: 1}}))
			opt = &TSRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "5"}
			result, err = adapter.TSRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo([]TSTimestampValue{{Timestamp: 0, Value: 5}, {Timestamp: 5, Value: 6}}))
			opt = &TSRangeOptions{Aggregator: Twa, BucketDuration: 10}
			result, err = adapter.TSRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo([]TSTimestampValue{{Timestamp: 0, Value: 2.55}, {Timestamp: 10, Value: 3}}))
			// Test Range Latest
			resultCreate, err := adapter.TSCreate(ctx, "t1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			resultCreate, err = adapter.TSCreate(ctx, "t2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			resultRule, err := adapter.TSCreateRule(ctx, "t1", "t2", Sum, 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRule).To(gomega.BeEquivalentTo("OK"))
			_, errAdd := adapter.TSAdd(ctx, "t1", 1, 1).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 2, 3).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 11, 7).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 13, 1).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			resultRange, err := adapter.TSRange(ctx, "t1", 0, 20).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 1, Value: 1}))

			opt = &TSRangeOptions{Latest: true}
			resultRange, err = adapter.TSRangeWithArgs(ctx, "t2", 0, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 0, Value: 4}))
			// Test Bucket Timestamp
			resultCreate, err = adapter.TSCreate(ctx, "t3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			_, errAdd = adapter.TSAdd(ctx, "t3", 15, 1).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 17, 4).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 51, 3).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 73, 5).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 75, 3).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())

			opt = &TSRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10}
			resultRange, err = adapter.TSRangeWithArgs(ctx, "t3", 0, 100, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 10, Value: 4}))
			gomega.Expect(len(resultRange)).To(gomega.BeEquivalentTo(3))

			opt = &TSRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10, BucketTimestamp: "+"}
			resultRange, err = adapter.TSRangeWithArgs(ctx, "t3", 0, 100, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 20, Value: 4}))
			gomega.Expect(len(resultRange)).To(gomega.BeEquivalentTo(3))
			// Test Empty
			_, errAdd = adapter.TSAdd(ctx, "t4", 15, 1).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 17, 4).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 51, 3).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 73, 5).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 75, 3).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())

			opt = &TSRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10}
			resultRange, err = adapter.TSRangeWithArgs(ctx, "t4", 0, 100, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 10, Value: 4}))
			gomega.Expect(len(resultRange)).To(gomega.BeEquivalentTo(3))

			opt = &TSRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10, Empty: true}
			resultRange, err = adapter.TSRangeWithArgs(ctx, "t4", 0, 100, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 10, Value: 4}))
			gomega.Expect(len(resultRange)).To(gomega.BeEquivalentTo(7))
		})

		ginkgo.It("should TSRevRange, TSRevRangeWithArgs", ginkgo.Label("timeseries", "tsrevrange", "tsrevrangeWithArgs"), func() {
			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i, float64(i%7)).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())

			}
			result, err := adapter.TSRange(ctx, "a", 0, 200).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(100))
			for i := 0; i < 100; i++ {
				adapter.TSAdd(ctx, "a", i+200, float64(i%7))
			}
			result, err = adapter.TSRange(ctx, "a", 0, 500).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(200))

			opt := &TSRevRangeOptions{Aggregator: Avg, BucketDuration: 10}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 500, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(20))

			opt = &TSRevRangeOptions{Count: 10}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 500, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(10))

			fts := make([]int, 0)
			for i := 10; i < 20; i++ {
				fts = append(fts, i)
			}
			opt = &TSRevRangeOptions{FilterByTS: fts, FilterByValue: []int{1, 2}}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 500, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(2))

			opt = &TSRevRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "+"}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo([]TSTimestampValue{{Timestamp: 10, Value: 1}, {Timestamp: 0, Value: 10}}))

			opt = &TSRevRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "1"}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo([]TSTimestampValue{{Timestamp: 1, Value: 10}, {Timestamp: 0, Value: 1}}))

			opt = &TSRevRangeOptions{Aggregator: Twa, BucketDuration: 10}
			result, err = adapter.TSRevRangeWithArgs(ctx, "a", 0, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result).To(gomega.BeEquivalentTo([]TSTimestampValue{{Timestamp: 10, Value: 3}, {Timestamp: 0, Value: 2.55}}))
			// Test Range Latest
			resultCreate, err := adapter.TSCreate(ctx, "t1").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			resultCreate, err = adapter.TSCreate(ctx, "t2").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			resultRule, err := adapter.TSCreateRule(ctx, "t1", "t2", Sum, 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRule).To(gomega.BeEquivalentTo("OK"))
			_, errAdd := adapter.TSAdd(ctx, "t1", 1, 1).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 2, 3).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 11, 7).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t1", 13, 1).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			resultRange, err := adapter.TSRange(ctx, "t2", 0, 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 0, Value: 4}))
			opt = &TSRevRangeOptions{Latest: true}
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t2", 0, 10, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 10, Value: 8}))
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t2", 0, 9, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 0, Value: 4}))
			// Test Bucket Timestamp
			resultCreate, err = adapter.TSCreate(ctx, "t3").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			_, errAdd = adapter.TSAdd(ctx, "t3", 15, 1).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 17, 4).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 51, 3).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 73, 5).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t3", 75, 3).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())

			opt = &TSRevRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10}
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t3", 0, 100, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 70, Value: 5}))
			gomega.Expect(len(resultRange)).To(gomega.BeEquivalentTo(3))

			opt = &TSRevRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10, BucketTimestamp: "+"}
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t3", 0, 100, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 80, Value: 5}))
			gomega.Expect(len(resultRange)).To(gomega.BeEquivalentTo(3))
			// Test Empty
			_, errAdd = adapter.TSAdd(ctx, "t4", 15, 1).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 17, 4).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 51, 3).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 73, 5).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())
			_, errAdd = adapter.TSAdd(ctx, "t4", 75, 3).Result()
			gomega.Expect(errAdd).NotTo(gomega.HaveOccurred())

			opt = &TSRevRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10}
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t4", 0, 100, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 70, Value: 5}))
			gomega.Expect(len(resultRange)).To(gomega.BeEquivalentTo(3))

			opt = &TSRevRangeOptions{Aggregator: Max, Align: 0, BucketDuration: 10, Empty: true}
			resultRange, err = adapter.TSRevRangeWithArgs(ctx, "t4", 0, 100, opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultRange[0]).To(gomega.BeEquivalentTo(TSTimestampValue{Timestamp: 70, Value: 5}))
			gomega.Expect(len(resultRange)).To(gomega.BeEquivalentTo(7))
		})

		ginkgo.It("should TSMRange and TSMRangeWithArgs", ginkgo.Label("timeseries", "tsmrange", "tsmrangeWithArgs"), func() {
			createOpt := &TSOptions{Labels: map[string]string{"Test": "This", "team": "ny"}}
			resultCreate, err := adapter.TSCreateWithArgs(ctx, "a", createOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			createOpt = &TSOptions{Labels: map[string]string{"Test": "This", "Taste": "That", "team": "sf"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", createOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))

			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i, float64(i%7)).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				_, err = adapter.TSAdd(ctx, "b", i, float64(i%11)).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			}

			result, err := adapter.TSMRange(ctx, 0, 200, []string{"Test=This"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(2))
			gomega.Expect(len(result["a"][2].([]interface{}))).To(gomega.BeEquivalentTo(100))
			// Test Count
			mrangeOpt := &TSMRangeOptions{Count: 10}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result["a"][2].([]interface{}))).To(gomega.BeEquivalentTo(10))
			// Test Aggregation and BucketDuration
			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i+200, float64(i%7)).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			}
			mrangeOpt = &TSMRangeOptions{Aggregator: Avg, BucketDuration: 10}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 500, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(2))
			gomega.Expect(len(result["a"][2].([]interface{}))).To(gomega.BeEquivalentTo(20))
			// Test WithLabels
			gomega.Expect(result["a"][0]).To(gomega.BeEquivalentTo(map[string]interface{}{}))
			mrangeOpt = &TSMRangeOptions{WithLabels: true}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][0]).To(gomega.BeEquivalentTo(map[string]interface{}{"Test": "This", "team": "ny"}))
			// Test SelectedLabels
			mrangeOpt = &TSMRangeOptions{SelectedLabels: []interface{}{"team"}}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][0]).To(gomega.BeEquivalentTo(map[string]interface{}{"team": "ny"}))
			gomega.Expect(result["b"][0]).To(gomega.BeEquivalentTo(map[string]interface{}{"team": "sf"}))
			// Test FilterBy
			fts := make([]int, 0)
			for i := 10; i < 20; i++ {
				fts = append(fts, i)
			}
			mrangeOpt = &TSMRangeOptions{FilterByTS: fts, FilterByValue: []int{1, 2}}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][2]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(15), 1.0}, []interface{}{int64(16), 2.0}}))
			// Test GroupBy
			mrangeOpt = &TSMRangeOptions{GroupByLabel: "Test", Reducer: "sum"}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["Test=This"][3]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(0), 0.0}, []interface{}{int64(1), 2.0}, []interface{}{int64(2), 4.0}, []interface{}{int64(3), 6.0}}))

			mrangeOpt = &TSMRangeOptions{GroupByLabel: "Test", Reducer: "max"}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["Test=This"][3]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(0), 0.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(3), 3.0}}))

			mrangeOpt = &TSMRangeOptions{GroupByLabel: "team", Reducer: "min"}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(2))
			gomega.Expect(result["team=ny"][3]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(0), 0.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(3), 3.0}}))
			gomega.Expect(result["team=sf"][3]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(0), 0.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(3), 3.0}}))
			// Test Align
			mrangeOpt = &TSMRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "-"}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 10, []string{"team=ny"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][2]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(0), 10.0}, []interface{}{int64(10), 1.0}}))

			mrangeOpt = &TSMRangeOptions{Aggregator: Count, BucketDuration: 10, Align: 5}
			result, err = adapter.TSMRangeWithArgs(ctx, 0, 10, []string{"team=ny"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][2]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(0), 5.0}, []interface{}{int64(5), 6.0}}))
		})

		ginkgo.It("should TSMRangeWithArgs Latest", ginkgo.Label("timeseries", "tsmrangeWithArgs", "tsmrangelatest"), func() {
			resultCreate, err := adapter.TSCreate(ctx, "a").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			opt := &TSOptions{Labels: map[string]string{"is_compaction": "true"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))

			resultCreate, err = adapter.TSCreate(ctx, "c").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"is_compaction": "true"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "d", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))

			resultCreateRule, err := adapter.TSCreateRule(ctx, "a", "b", Sum, 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreateRule).To(gomega.BeEquivalentTo("OK"))
			resultCreateRule, err = adapter.TSCreateRule(ctx, "c", "d", Sum, 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreateRule).To(gomega.BeEquivalentTo("OK"))

			_, err = adapter.TSAdd(ctx, "a", 1, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 2, 3).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 11, 7).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 13, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			_, err = adapter.TSAdd(ctx, "c", 1, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 2, 3).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 11, 7).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 13, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			mrangeOpt := &TSMRangeOptions{Latest: true}
			result, err := adapter.TSMRangeWithArgs(ctx, 0, 10, []string{"is_compaction=true"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["b"][2]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(0), 4.0}, []interface{}{int64(10), 8.0}}))
			gomega.Expect(result["d"][2]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(0), 4.0}, []interface{}{int64(10), 8.0}}))
		})
		ginkgo.It("should TSMRevRange and TSMRevRangeWithArgs", ginkgo.Label("timeseries", "tsmrevrange", "tsmrevrangeWithArgs"), func() {
			createOpt := &TSOptions{Labels: map[string]string{"Test": "This", "team": "ny"}}
			resultCreate, err := adapter.TSCreateWithArgs(ctx, "a", createOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			createOpt = &TSOptions{Labels: map[string]string{"Test": "This", "Taste": "That", "team": "sf"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", createOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))

			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i, float64(i%7)).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				_, err = adapter.TSAdd(ctx, "b", i, float64(i%11)).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			}
			result, err := adapter.TSMRevRange(ctx, 0, 200, []string{"Test=This"}).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(2))
			gomega.Expect(len(result["a"][2].([]interface{}))).To(gomega.BeEquivalentTo(100))
			// Test Count
			mrangeOpt := &TSMRevRangeOptions{Count: 10}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result["a"][2].([]interface{}))).To(gomega.BeEquivalentTo(10))
			// Test Aggregation and BucketDuration
			for i := 0; i < 100; i++ {
				_, err := adapter.TSAdd(ctx, "a", i+200, float64(i%7)).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			}
			mrangeOpt = &TSMRevRangeOptions{Aggregator: Avg, BucketDuration: 10}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 500, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(2))
			gomega.Expect(len(result["a"][2].([]interface{}))).To(gomega.BeEquivalentTo(20))
			gomega.Expect(result["a"][0]).To(gomega.BeEquivalentTo(map[string]interface{}{}))
			// Test WithLabels
			gomega.Expect(result["a"][0]).To(gomega.BeEquivalentTo(map[string]interface{}{}))
			mrangeOpt = &TSMRevRangeOptions{WithLabels: true}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][0]).To(gomega.BeEquivalentTo(map[string]interface{}{"Test": "This", "team": "ny"}))
			// Test SelectedLabels
			mrangeOpt = &TSMRevRangeOptions{SelectedLabels: []interface{}{"team"}}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][0]).To(gomega.BeEquivalentTo(map[string]interface{}{"team": "ny"}))
			gomega.Expect(result["b"][0]).To(gomega.BeEquivalentTo(map[string]interface{}{"team": "sf"}))
			// Test FilterBy
			fts := make([]int, 0)
			for i := 10; i < 20; i++ {
				fts = append(fts, i)
			}
			mrangeOpt = &TSMRevRangeOptions{FilterByTS: fts, FilterByValue: []int{1, 2}}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 200, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][2]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(16), 2.0}, []interface{}{int64(15), 1.0}}))
			// Test GroupBy
			mrangeOpt = &TSMRevRangeOptions{GroupByLabel: "Test", Reducer: "sum"}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["Test=This"][3]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(3), 6.0}, []interface{}{int64(2), 4.0}, []interface{}{int64(1), 2.0}, []interface{}{int64(0), 0.0}}))

			mrangeOpt = &TSMRevRangeOptions{GroupByLabel: "Test", Reducer: "max"}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["Test=This"][3]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(3), 3.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(0), 0.0}}))

			mrangeOpt = &TSMRevRangeOptions{GroupByLabel: "team", Reducer: "min"}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 3, []string{"Test=This"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(len(result)).To(gomega.BeEquivalentTo(2))
			gomega.Expect(result["team=ny"][3]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(3), 3.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(0), 0.0}}))
			gomega.Expect(result["team=sf"][3]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(3), 3.0}, []interface{}{int64(2), 2.0}, []interface{}{int64(1), 1.0}, []interface{}{int64(0), 0.0}}))
			// Test Align
			mrangeOpt = &TSMRevRangeOptions{Aggregator: Count, BucketDuration: 10, Align: "-"}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 10, []string{"team=ny"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][2]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(10), 1.0}, []interface{}{int64(0), 10.0}}))

			mrangeOpt = &TSMRevRangeOptions{Aggregator: Count, BucketDuration: 10, Align: 1}
			result, err = adapter.TSMRevRangeWithArgs(ctx, 0, 10, []string{"team=ny"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["a"][2]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(1), 10.0}, []interface{}{int64(0), 1.0}}))
		})

		ginkgo.It("should TSMRevRangeWithArgs Latest", ginkgo.Label("timeseries", "tsmrevrangeWithArgs", "tsmrevrangelatest"), func() {
			resultCreate, err := adapter.TSCreate(ctx, "a").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			opt := &TSOptions{Labels: map[string]string{"is_compaction": "true"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "b", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))

			resultCreate, err = adapter.TSCreate(ctx, "c").Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))
			opt = &TSOptions{Labels: map[string]string{"is_compaction": "true"}}
			resultCreate, err = adapter.TSCreateWithArgs(ctx, "d", opt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreate).To(gomega.BeEquivalentTo("OK"))

			resultCreateRule, err := adapter.TSCreateRule(ctx, "a", "b", Sum, 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreateRule).To(gomega.BeEquivalentTo("OK"))
			resultCreateRule, err = adapter.TSCreateRule(ctx, "c", "d", Sum, 10).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(resultCreateRule).To(gomega.BeEquivalentTo("OK"))

			_, err = adapter.TSAdd(ctx, "a", 1, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 2, 3).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 11, 7).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "a", 13, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			_, err = adapter.TSAdd(ctx, "c", 1, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 2, 3).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 11, 7).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			_, err = adapter.TSAdd(ctx, "c", 13, 1).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			mrangeOpt := &TSMRevRangeOptions{Latest: true}
			result, err := adapter.TSMRevRangeWithArgs(ctx, 0, 10, []string{"is_compaction=true"}, mrangeOpt).Result()
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			gomega.Expect(result["b"][2]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(10), 8.0}, []interface{}{int64(0), 4.0}}))
			gomega.Expect(result["d"][2]).To(gomega.BeEquivalentTo([]interface{}{[]interface{}{int64(10), 8.0}, []interface{}{int64(0), 4.0}}))
		})
	})
	ginkgo.Describe("JSON Commands", ginkgo.Label("json"), func() {
		ginkgo.BeforeEach(func() {
			gomega.Expect(adapter.FlushDB(ctx).Err()).NotTo(gomega.HaveOccurred())
		})

		ginkgo.Describe("arrays", ginkgo.Label("arrays"), func() {
			ginkgo.It("should JSONArrAppend", ginkgo.Label("json.arrappend", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "append2", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONArrAppend(ctx, "append2", "$..a", 10)
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.Equal([]int64{2, 3}))
			})

			ginkgo.It("should JSONArrIndex and JSONArrIndexWithArgs", ginkgo.Label("json.arrindex", "json"), func() {
				cmd1, err := adapter.JSONSet(ctx, "index1", "$", `{"a": [10], "b": {"a": [12, 10]}}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1).To(gomega.Equal("OK"))

				cmd2, err := adapter.JSONArrIndex(ctx, "index1", "$.b.a", 10).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2).To(gomega.Equal([]int64{1}))

				cmd3, err := adapter.JSONSet(ctx, "index2", "$", `[0,1,2,3,4]`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd3).To(gomega.Equal("OK"))

				res, err := adapter.JSONArrIndex(ctx, "index2", "$", 1).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res[0]).To(gomega.Equal(int64(1)))

				res, err = adapter.JSONArrIndex(ctx, "index2", "$", 1, 2).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res[0]).To(gomega.Equal(int64(-1)))

				res, err = adapter.JSONArrIndex(ctx, "index2", "$", 4).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res[0]).To(gomega.Equal(int64(4)))

				res, err = adapter.JSONArrIndexWithArgs(ctx, "index2", "$", &JSONArrIndexArgs{}, 4).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res[0]).To(gomega.Equal(int64(4)))

				stop := 5000
				res, err = adapter.JSONArrIndexWithArgs(ctx, "index2", "$", &JSONArrIndexArgs{Stop: &stop}, 4).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res[0]).To(gomega.Equal(int64(4)))

				stop = -1
				res, err = adapter.JSONArrIndexWithArgs(ctx, "index2", "$", &JSONArrIndexArgs{Stop: &stop}, 4).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res[0]).To(gomega.Equal(int64(-1)))

			})

			// FIXME: how to deal with expanded ?
			ginkgo.It("should JSONArrIndex and JSONArrIndexWithArgs with $", ginkgo.Label("json.arrindex", "json"), func() {
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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				resGet, err := adapter.JSONGet(ctx, "doc1", "$.store.book[?(@.price<10)].size").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resGet).To(gomega.Equal("[[10,20,30,40],[5,10,20,30]]"))

				resArr, err := adapter.JSONArrIndex(ctx, "doc1", "$.store.book[?(@.price<10)].size", 20).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resArr).To(gomega.Equal([]int64{1, 2}))
			})

			ginkgo.It("should JSONArrInsert", ginkgo.Label("json.arrinsert", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "insert2", "$", `[100, 200, 300, 200]`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONArrInsert(ctx, "insert2", "$", -1, 1, 2)
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.Equal([]int64{6}))

				cmd3 := adapter.JSONGet(ctx, "insert2")
				gomega.Expect(cmd3.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd3.Val()).To(gomega.Or(
					gomega.Equal(`[[100,200,300,1,2,200]]`)))
			})

			ginkgo.It("should JSONArrLen", ginkgo.Label("json.arrlen", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "length2", "$", `{"a": [10], "b": {"a": [12, 10, 20, 12, 90, 10]}}`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONArrLen(ctx, "length2", "$..a")
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.Equal([]int64{1, 6}))
			})

			ginkgo.It("should JSONArrPop", ginkgo.Label("json.arrpop"), func() {
				cmd1 := adapter.JSONSet(ctx, "pop4", "$", `[100, 200, 300, 200]`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONArrPop(ctx, "pop4", "$", 2)
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.Equal([]string{"300"}))

				cmd3 := adapter.JSONGet(ctx, "pop4", "$")
				gomega.Expect(cmd3.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd3.Val()).To(gomega.Equal("[[100,200,200]]"))
			})

			ginkgo.It("should JSONArrTrim", ginkgo.Label("json.arrtrim", "json"), func() {
				cmd1, err := adapter.JSONSet(ctx, "trim1", "$", `[0,1,2,3,4]`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1).To(gomega.Equal("OK"))

				stop := 3
				cmd2, err := adapter.JSONArrTrimWithArgs(ctx, "trim1", "$", &JSONArrTrimArgs{Start: 1, Stop: &stop}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2).To(gomega.Equal([]int64{3}))

				res, err := adapter.JSONGet(ctx, "trim1", "$").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal(`[[1,2,3]]`))

				cmd3, err := adapter.JSONSet(ctx, "trim2", "$", `[0,1,2,3,4]`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd3).To(gomega.Equal("OK"))

				stop = 3
				cmd4, err := adapter.JSONArrTrimWithArgs(ctx, "trim2", "$", &JSONArrTrimArgs{Start: -1, Stop: &stop}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd4).To(gomega.Equal([]int64{0}))

				cmd5, err := adapter.JSONSet(ctx, "trim3", "$", `[0,1,2,3,4]`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd5).To(gomega.Equal("OK"))

				stop = 99
				cmd6, err := adapter.JSONArrTrimWithArgs(ctx, "trim3", "$", &JSONArrTrimArgs{Start: 3, Stop: &stop}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd6).To(gomega.Equal([]int64{2}))

				cmd7, err := adapter.JSONSet(ctx, "trim4", "$", `[0,1,2,3,4]`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd7).To(gomega.Equal("OK"))

				stop = 1
				cmd8, err := adapter.JSONArrTrimWithArgs(ctx, "trim4", "$", &JSONArrTrimArgs{Start: 9, Stop: &stop}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd8).To(gomega.Equal([]int64{0}))

				cmd9, err := adapter.JSONSet(ctx, "trim5", "$", `[0,1,2,3,4]`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd9).To(gomega.Equal("OK"))

				stop = 11
				cmd10, err := adapter.JSONArrTrimWithArgs(ctx, "trim5", "$", &JSONArrTrimArgs{Start: 9, Stop: &stop}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd10).To(gomega.Equal([]int64{0}))
			})

			ginkgo.It("should JSONArrPop", ginkgo.Label("json.arrpop", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "pop4", "$", `[100, 200, 300, 200]`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONArrPop(ctx, "pop4", "$", 2)
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.Equal([]string{"300"}))

				cmd3 := adapter.JSONGet(ctx, "pop4", "$")
				gomega.Expect(cmd3.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd3.Val()).To(gomega.Equal("[[100,200,200]]"))
			})

		})

		ginkgo.Describe("get/set", ginkgo.Label("getset"), func() {
			ginkgo.It("should JSONSet", ginkgo.Label("json.set", "json"), func() {
				cmd := adapter.JSONSet(ctx, "set1", "$", `{"a": 1, "b": 2, "hello": "world"}`)
				gomega.Expect(cmd.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd.Val()).To(gomega.Equal("OK"))
			})

			ginkgo.It("should JSONGet", ginkgo.Label("json.get", "json"), func() {
				res, err := adapter.JSONSet(ctx, "get3", "$", `{"a": 1, "b": 2}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				res, err = adapter.JSONGetWithArgs(ctx, "get3", &JSONGetArgs{Indent: "-"}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal(`[-{--"a":1,--"b":2-}]`))

				res, err = adapter.JSONGetWithArgs(ctx, "get3", &JSONGetArgs{Indent: "-", Newline: `~`, Space: `!`}).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal(`[~-{~--"a":!1,~--"b":!2~-}~]`))
			})

			ginkgo.It("should JSONMerge", ginkgo.Label("json.merge", "json"), func() {
				res, err := adapter.JSONSet(ctx, "merge1", "$", `{"a": 1, "b": 2}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				res, err = adapter.JSONMerge(ctx, "merge1", "$", `{"b": 3, "c": 4}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				res, err = adapter.JSONGet(ctx, "merge1", "$").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal(`[{"a":1,"b":3,"c":4}]`))
			})

			ginkgo.It("should JSONMSet", ginkgo.Label("json.mset", "json"), func() {
				doc1 := JSONSetArgs{Key: "mset1", Path: "$", Value: `{"a": 1}`}
				doc2 := JSONSetArgs{Key: "mset2", Path: "$", Value: 2}
				docs := []JSONSetArgs{doc1, doc2}

				mSetResult, err := adapter.JSONMSetArgs(ctx, docs).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(mSetResult).To(gomega.Equal("OK"))

				res, err := adapter.JSONMGet(ctx, "$", "mset1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal([]interface{}{`[{"a":1}]`}))

				res, err = adapter.JSONMGet(ctx, "$", "mset1", "mset2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal([]interface{}{`[{"a":1}]`, "[2]"}))

				mSetResult, err = adapter.JSONMSet(ctx, "mset1", "$.a", 2, "mset3", "$", `[1]`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
			})

			ginkgo.It("should JSONMGet", ginkgo.Label("json.mget", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "mget2a", "$", `{"a": ["aa", "ab", "ac", "ad"], "b": {"a": ["ba", "bb", "bc", "bd"]}}`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))
				cmd2 := adapter.JSONSet(ctx, "mget2b", "$", `{"a": [100, 200, 300, 200], "b": {"a": [100, 200, 300, 200]}}`)
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.Equal("OK"))

				cmd3 := adapter.JSONMGet(ctx, "$..a", "mget2a", "mget2b")
				gomega.Expect(cmd3.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd3.Val()).To(gomega.HaveLen(2))
				gomega.Expect(cmd3.Val()[0]).To(gomega.Equal(`[["aa","ab","ac","ad"],["ba","bb","bc","bd"]]`))
				gomega.Expect(cmd3.Val()[1]).To(gomega.Equal(`[[100,200,300,200],[100,200,300,200]]`))
			})

			ginkgo.It("should JSONMget with $", ginkgo.Label("json.mget", "json"), func() {
				res, err := adapter.JSONSet(ctx, "doc1", "$", `{"a": 1, "b": 2, "nested": {"a": 3}, "c": "", "nested2": {"a": ""}}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				res, err = adapter.JSONSet(ctx, "doc2", "$", `{"a": 4, "b": 5, "nested": {"a": 6}, "c": "", "nested2": {"a": [""]}}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				iRes, err := adapter.JSONMGet(ctx, "$..a", "doc1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal([]interface{}{`[1,3,""]`}))

				iRes, err = adapter.JSONMGet(ctx, "$..a", "doc1", "doc2").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal([]interface{}{`[1,3,""]`, `[4,6,[""]]`}))

				iRes, err = adapter.JSONMGet(ctx, "$..a", "non_existing_doc", "non_existing_doc1").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal([]interface{}{nil, nil}))
			})
		})

		ginkgo.Describe("Misc", ginkgo.Label("misc"), func() {

			ginkgo.It("should JSONClear", ginkgo.Label("json.clear", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "clear1", "$", `[1]`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONClear(ctx, "clear1", "$")
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.Equal(int64(1)))

				cmd3 := adapter.JSONGet(ctx, "clear1", "$")
				gomega.Expect(cmd3.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd3.Val()).To(gomega.Equal(`[[]]`))
			})

			ginkgo.It("should JSONClear with $", ginkgo.Label("json.clear", "json"), func() {
				doc := `{
					"nested1": {"a": {"foo": 10, "bar": 20}},
					"a": ["foo"],
					"nested2": {"a": "claro"},
					"nested3": {"a": {"baz": 50}}
				}`
				res, err := adapter.JSONSet(ctx, "doc1", "$", doc).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				iRes, err := adapter.JSONClear(ctx, "doc1", "$..a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal(int64(3)))

				resGet, err := adapter.JSONGet(ctx, "doc1", `$`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resGet).To(gomega.Equal(`[{"nested1":{"a":{}},"a":[],"nested2":{"a":"claro"},"nested3":{"a":{}}}]`))

				res, err = adapter.JSONSet(ctx, "doc1", "$", doc).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				iRes, err = adapter.JSONClear(ctx, "doc1", "$.nested1.a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal(int64(1)))

				resGet, err = adapter.JSONGet(ctx, "doc1", `$`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resGet).To(gomega.Equal(`[{"nested1":{"a":{}},"a":["foo"],"nested2":{"a":"claro"},"nested3":{"a":{"baz":50}}}]`))
			})

			ginkgo.It("should JSONDel", ginkgo.Label("json.del", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "del1", "$", `[1]`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONDel(ctx, "del1", "$")
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.Equal(int64(1)))

				cmd3 := adapter.JSONGet(ctx, "del1", "$")
				// go-redis's test assertion is wrong.
				// based on the result from redis/redis-stack:7.2.0-v3,
				// cmd3.Err() should be rueidis.Nil, not nil
				gomega.Expect(cmd3.Err()).To(gomega.Equal(rueidis.Nil))
				gomega.Expect(cmd3.Val()).To(gomega.HaveLen(0))
			})

			ginkgo.It("should JSONDel with $", ginkgo.Label("json.del", "json"), func() {
				res, err := adapter.JSONSet(ctx, "del1", "$", `{"a": 1, "nested": {"a": 2, "b": 3}}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				iRes, err := adapter.JSONDel(ctx, "del1", "$..a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal(int64(2)))

				resGet, err := adapter.JSONGet(ctx, "del1", "$").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resGet).To(gomega.Equal(`[{"nested":{"b":3}}]`))

				res, err = adapter.JSONSet(ctx, "del2", "$", `{"a": {"a": 2, "b": 3}, "b": ["a", "b"], "nested": {"b": [true, "a", "b"]}}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				iRes, err = adapter.JSONDel(ctx, "del2", "$..a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal(int64(1)))

				resGet, err = adapter.JSONGet(ctx, "del2", "$").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resGet).To(gomega.Equal(`[{"nested":{"b":[true,"a","b"]},"b":["a","b"]}]`))

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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				iRes, err = adapter.JSONDel(ctx, "del3", `$.[0]["nested"]..ciao`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal(int64(3)))

				resVal := `[[{"ciao":["non ancora"],"nested":[{},{},{"ciaoc":[3,"non","ciao"]},{},{"e":[5,"non","ciao"]}]}]]`
				resGet, err = adapter.JSONGet(ctx, "del3", "$").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resGet).To(gomega.Equal(resVal))
			})

			ginkgo.It("should JSONForget", ginkgo.Label("json.forget", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "forget3", "$", `{"a": [1,2,3], "b": {"a": [1,2,3], "b": "annie"}}`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONForget(ctx, "forget3", "$..a")
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.Equal(int64(2)))

				cmd3 := adapter.JSONGet(ctx, "forget3", "$")
				gomega.Expect(cmd3.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd3.Val()).To(gomega.Equal(`[{"b":{"b":"annie"}}]`))

			})

			ginkgo.It("should JSONForget with $", ginkgo.Label("json.forget", "json"), func() {
				res, err := adapter.JSONSet(ctx, "doc1", "$", `{"a": 1, "nested": {"a": 2, "b": 3}}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				iRes, err := adapter.JSONForget(ctx, "doc1", "$..a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal(int64(2)))

				resGet, err := adapter.JSONGet(ctx, "doc1", "$").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resGet).To(gomega.Equal(`[{"nested":{"b":3}}]`))

				res, err = adapter.JSONSet(ctx, "doc2", "$", `{"a": {"a": 2, "b": 3}, "b": ["a", "b"], "nested": {"b": [true, "a", "b"]}}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				iRes, err = adapter.JSONForget(ctx, "doc2", "$..a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal(int64(1)))

				resGet, err = adapter.JSONGet(ctx, "doc2", "$").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resGet).To(gomega.Equal(`[{"nested":{"b":[true,"a","b"]},"b":["a","b"]}]`))

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
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				iRes, err = adapter.JSONForget(ctx, "doc3", `$.[0]["nested"]..ciao`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(iRes).To(gomega.Equal(int64(3)))

				resVal := `[[{"ciao":["non ancora"],"nested":[{},{},{"ciaoc":[3,"non","ciao"]},{},{"e":[5,"non","ciao"]}]}]]`
				resGet, err = adapter.JSONGet(ctx, "doc3", "$").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(resGet).To(gomega.Equal(resVal))
			})

			ginkgo.It("should JSONNumIncrBy", ginkgo.Label("json.numincrby", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "incr3", "$", `{"a": [1, 2], "b": {"a": [0, -1]}}`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONNumIncrBy(ctx, "incr3", "$..a[1]", float64(1))
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.Equal(`[3,0]`))

				cmd3 := adapter.JSONSet(ctx, "incr4", "$", `{"a": [1, 2], "b": {"a": [0, -1], "c": "z"}, "c": 2}`)
				gomega.Expect(cmd3.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd3.Val()).To(gomega.Equal("OK"))

				cmd4 := adapter.JSONNumIncrBy(ctx, "incr4", "$..c", float64(1))
				gomega.Expect(cmd4.Err()).NotTo(gomega.HaveOccurred())
				// for NaN field, it should be null
				gomega.Expect(cmd4.Val()).To(gomega.Equal(`[3,null]`))
			})

			ginkgo.It("should JSONNumIncrBy with $", ginkgo.Label("json.numincrby", "json"), func() {
				res, err := adapter.JSONSet(ctx, "doc1", "$", `{"a": "b", "b": [{"a": 2}, {"a": 5.0}, {"a": "c"}]}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				res, err = adapter.JSONNumIncrBy(ctx, "doc1", "$.b[1].a", 2).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal(`[7]`))

				res, err = adapter.JSONNumIncrBy(ctx, "doc1", "$.b[1].a", 3.5).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal(`[10.5]`))

				res, err = adapter.JSONSet(ctx, "doc2", "$", `{"a": "b", "b": [{"a": 2}, {"a": 5.0}, {"a": "c"}]}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				res, err = adapter.JSONNumIncrBy(ctx, "doc2", "$.b[0].a", 3).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal(`[5]`))
			})

			ginkgo.It("should JSONObjKeys", ginkgo.Label("json.objkeys", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "objkeys1", "$", `{"a": [1, 2], "b": {"a": [0, -1]}}`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONObjKeys(ctx, "objkeys1", "$..*")
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.HaveLen(7))
				gomega.Expect(cmd2.Val()).To(gomega.Equal([]interface{}{nil, []interface{}{"a"}, nil, nil, nil, nil, nil}))
			})

			ginkgo.It("should JSONObjKeys with $", ginkgo.Label("json.objkeys", "json"), func() {
				doc := `{
					"nested1": {"a": {"foo": 10, "bar": 20}},
					"a": ["foo"],
					"nested2": {"a": {"baz": 50}}
				}`
				cmd1, err := adapter.JSONSet(ctx, "objkeys1", "$", doc).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1).To(gomega.Equal("OK"))

				cmd2, err := adapter.JSONObjKeys(ctx, "objkeys1", "$.nested1.a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2).To(gomega.Equal([]interface{}{[]interface{}{"foo", "bar"}}))

				cmd2, err = adapter.JSONObjKeys(ctx, "objkeys1", ".*.a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2).To(gomega.Equal([]interface{}{"foo", "bar"}))

				cmd2, err = adapter.JSONObjKeys(ctx, "objkeys1", ".nested2.a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2).To(gomega.Equal([]interface{}{"baz"}))

				_, err = adapter.JSONObjKeys(ctx, "non_existing_doc", "..a").Result()
				gomega.Expect(err).To(gomega.HaveOccurred())
			})

			ginkgo.It("should JSONObjLen", ginkgo.Label("json.objlen", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "objlen2", "$", `{"a": [1, 2], "b": {"a": [0, -1]}}`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONObjLen(ctx, "objlen2", "$..*")
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.HaveLen(7))
				gomega.Expect(cmd2.Val()[0]).To(gomega.BeNil())
				gomega.Expect(*cmd2.Val()[1]).To(gomega.Equal(int64(1)))
			})

			ginkgo.It("should JSONStrLen", ginkgo.Label("json.strlen", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "strlen2", "$", `{"a": "alice", "b": "bob", "c": {"a": "alice", "b": "bob"}}`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONStrLen(ctx, "strlen2", "$..*")
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.HaveLen(5))
				var tmp int64 = 20
				gomega.Expect(cmd2.Val()[0]).To(gomega.BeAssignableToTypeOf(&tmp))
				gomega.Expect(*cmd2.Val()[0]).To(gomega.Equal(int64(5)))
				gomega.Expect(*cmd2.Val()[1]).To(gomega.Equal(int64(3)))
				gomega.Expect(cmd2.Val()[2]).To(gomega.BeNil())
				gomega.Expect(*cmd2.Val()[3]).To(gomega.Equal(int64(5)))
				gomega.Expect(*cmd2.Val()[4]).To(gomega.Equal(int64(3)))
			})

			ginkgo.It("should JSONStrAppend", ginkgo.Label("json.strappend", "json"), func() {
				cmd1, err := adapter.JSONSet(ctx, "strapp1", "$", `"foo"`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1).To(gomega.Equal("OK"))
				cmd2, err := adapter.JSONStrAppend(ctx, "strapp1", "$", `"bar"`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(*cmd2[0]).To(gomega.Equal(int64(6)))
				cmd3, err := adapter.JSONGet(ctx, "strapp1", "$").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd3).To(gomega.Equal(`["foobar"]`))

			})

			ginkgo.It("should JSONStrAppend and JSONStrLen with $", ginkgo.Label("json.strappend", "json.strlen", "json"), func() {
				res, err := adapter.JSONSet(ctx, "doc1", "$", `{"a": "foo", "nested1": {"a": "hello"}, "nested2": {"a": 31}}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				intArrayResult, err := adapter.JSONStrAppend(ctx, "doc1", "$.nested1.a", `"baz"`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(*intArrayResult[0]).To(gomega.Equal(int64(8)))

				res, err = adapter.JSONSet(ctx, "doc2", "$", `{"a": "foo", "nested1": {"a": "hello"}, "nested2": {"a": 31}}`).Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(res).To(gomega.Equal("OK"))

				intResult, err := adapter.JSONStrLen(ctx, "doc2", "$.nested1.a").Result()
				gomega.Expect(err).NotTo(gomega.HaveOccurred())
				gomega.Expect(*intResult[0]).To(gomega.Equal(int64(5)))
			})

			ginkgo.It("should JSONToggle", ginkgo.Label("json.toggle", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "toggle1", "$", `[true]`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONToggle(ctx, "toggle1", "$[0]")
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.HaveLen(1))
				gomega.Expect(*cmd2.Val()[0]).To(gomega.Equal(int64(0)))
			})

			ginkgo.It("should JSONType", ginkgo.Label("json.type", "json"), func() {
				cmd1 := adapter.JSONSet(ctx, "type1", "$", `[true]`)
				gomega.Expect(cmd1.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd1.Val()).To(gomega.Equal("OK"))

				cmd2 := adapter.JSONType(ctx, "type1", "$[0]")
				gomega.Expect(cmd2.Err()).NotTo(gomega.HaveOccurred())
				gomega.Expect(cmd2.Val()).To(gomega.HaveLen(1))
				// RESP2 v RESP3
				gomega.Expect(cmd2.Val()[0]).To(gomega.Or(gomega.Equal([]interface{}{"boolean"}), gomega.Equal("boolean")))
			})
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
