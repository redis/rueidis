package rueidis

import (
	"context"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func parallel(p int) (chan func(), func()) {
	ch := make(chan func(), p)
	wg := sync.WaitGroup{}
	wg.Add(p)
	for i := 0; i < p; i++ {
		go func() {
			for fn := range ch {
				fn()
			}
			wg.Done()
		}()
	}
	return ch, func() {
		close(ch)
		wg.Wait()
	}
}

func testFlush(t *testing.T, client Client) {
	ctx := context.Background()

	keys := 1000
	para := 8

	kvs := make(map[string]string, keys)
	for i := 0; i < keys; i++ {
		kvs[strconv.Itoa(i)] = strconv.FormatInt(rand.Int63(), 10)
	}

	t.Logf("prepare %d keys for FLUSH\n", keys)
	jobs, wait := parallel(para)
	for i := 0; i < keys; i++ {
		key := strconv.Itoa(i)
		jobs <- func() {
			val, err := client.Do(ctx, client.B().Set().Key(key).Value(kvs[key]).Build()).ToString()
			if err != nil || val != "OK" {
				t.Errorf("unexpected set response %v %v", val, err)
			}
		}
	}
	wait()

	t.Logf("testing client side caching before flush\n")
	jobs, wait = parallel(para)
	for i := 0; i < keys; i++ {
		key := strconv.Itoa(i)
		jobs <- func() {
			resp := client.DoCache(ctx, client.B().Get().Key(key).Cache(), time.Minute)
			val, err := resp.ToString()
			if val != kvs[key] {
				t.Errorf("unexpected csc get response %v %v", val, err)
			}
			if resp.IsCacheHit() {
				t.Errorf("unexpected csc cache hit")
			}
		}
	}
	wait()

	if err := client.Do(ctx, client.B().Flushall().Build()).Error(); err != nil {
		t.Errorf("unexpected flush err %v", err)
	}

	time.Sleep(time.Second)

	t.Logf("testing client side caching after flush\n")
	jobs, wait = parallel(para)
	for i := 0; i < keys; i++ {
		key := strconv.Itoa(i)
		jobs <- func() {
			resp := client.DoCache(ctx, client.B().Get().Key(key).Cache(), time.Minute)
			if !IsRedisNil(resp.Error()) {
				t.Errorf("unexpected csc get response after flush %v", resp)
			}
			if resp.IsCacheHit() {
				t.Errorf("unexpected csc cache hit after flush")
			}
		}
	}
	wait()
}

func testSETGETCSC(t *testing.T, client Client) {
	testSETGET(t, client, true)
}

func testSETGETRESP2(t *testing.T, client Client) {
	testSETGET(t, client, false)
}

//gocyclo:ignore
func testSETGET(t *testing.T, client Client, csc bool) {
	ctx := context.Background()
	keys := 10000
	para := 8

	kvs := make(map[string]string, keys)
	for i := 0; i < keys; i++ {
		kvs[strconv.Itoa(i)] = strconv.FormatInt(rand.Int63(), 10)
	}

	t.Logf("testing SET with %d keys and %d parallelism\n", keys, para)
	jobs, wait := parallel(para)
	for i := 0; i < keys; i++ {
		key := strconv.Itoa(i)
		jobs <- func() {
			val, err := client.Do(ctx, client.B().Set().Key(key).Value(kvs[key]).Build()).ToString()
			if err != nil || val != "OK" {
				t.Errorf("unexpected set response %v %v", val, err)
			}
		}
	}
	wait()

	t.Logf("testing GET with %d keys and %d parallelism\n", keys*2, para)
	jobs, wait = parallel(para)
	for i := 0; i < keys*2; i++ {
		key := strconv.Itoa(rand.Intn(keys * 2))
		jobs <- func() {
			val, err := client.Do(ctx, client.B().Get().Key(key).Build()).ToString()
			if v, ok := kvs[key]; !((ok && val == v) || (!ok && IsRedisNil(err))) {
				t.Errorf("unexpected get response %v %v %v", val, err, ok)
			}
		}
	}
	wait()

	t.Logf("testing client side caching with %d interations and %d parallelism\n", keys*5, para)
	jobs, wait = parallel(para)
	hits, miss := int64(0), int64(0)
	for i := 0; i < keys*10; i++ {
		key := strconv.Itoa(rand.Intn(keys / 100))
		jobs <- func() {
			resp := client.DoCache(ctx, client.B().Get().Key(key).Cache(), time.Minute)
			val, err := resp.ToString()
			if v, ok := kvs[key]; !((ok && val == v) || (!ok && IsRedisNil(err))) {
				t.Errorf("unexpected csc get response %v %v %v", val, err, ok)
			}
			if resp.IsCacheHit() {
				atomic.AddInt64(&hits, 1)
			} else {
				atomic.AddInt64(&miss, 1)
			}
		}
	}
	wait()
	if csc {
		if atomic.LoadInt64(&miss) != 100 || atomic.LoadInt64(&hits) != int64(keys*10-100) {
			t.Fatalf("unexpected client side caching hits and miss %v %v", atomic.LoadInt64(&hits), atomic.LoadInt64(&miss))
		}
	} else {
		if atomic.LoadInt64(&hits) != 0 {
			t.Fatalf("unexpected client side caching hits and miss %v %v", atomic.LoadInt64(&hits), atomic.LoadInt64(&miss))
		}
	}

	t.Logf("testing DEL with %d keys and %d parallelism\n", keys*2, para)
	jobs, wait = parallel(para)
	for i := 0; i < keys*2; i++ {
		key := strconv.Itoa(i)
		jobs <- func() {
			val, err := client.Do(ctx, client.B().Del().Key(key).Build()).AsInt64()
			if _, ok := kvs[key]; !((val == 1 && ok) || (val == 0 && !ok)) {
				t.Errorf("unexpected del response %v %v %v", val, err, ok)
			}
		}
	}
	wait()

	t.Logf("testing client side caching after delete\n")
	jobs, wait = parallel(para)
	for i := 0; i < keys/100; i++ {
		key := strconv.Itoa(i)
		jobs <- func() {
			resp := client.DoCache(ctx, client.B().Get().Key(key).Cache(), time.Minute)
			if !IsRedisNil(resp.Error()) {
				t.Errorf("unexpected csc get response after delete %v", resp)
			}
			if resp.IsCacheHit() {
				t.Errorf("unexpected csc cache hit after delete")
			}
		}
	}
	wait()
}

func testMultiSETGETCSC(t *testing.T, client Client) {
	testMultiSETGET(t, client, true)
}

func testMultiSETGETRESP2(t *testing.T, client Client) {
	testMultiSETGET(t, client, false)
}

//gocyclo:ignore
func testMultiSETGET(t *testing.T, client Client, csc bool) {
	ctx := context.Background()
	keys := 10000
	batch := 100
	para := 8

	kvs := make(map[string]string, keys)
	for i := 0; i < keys; i++ {
		kvs["m"+strconv.Itoa(i)] = strconv.FormatInt(rand.Int63(), 10)
	}

	t.Logf("testing Multi SET with %d keys and %d parallelism\n", keys, para)
	jobs, wait := parallel(para)
	for i := 0; i < keys; i += batch {
		commands := make(Commands, 0, batch)
		for j := 0; j < batch; j++ {
			key := "m" + strconv.Itoa(i+j)
			commands = append(commands, client.B().Set().Key(key).Value(kvs[key]).Build())
		}
		jobs <- func() {
			for _, resp := range client.DoMulti(ctx, commands...) {
				val, err := resp.ToString()
				if err != nil || val != "OK" {
					t.Errorf("unexpected set response %v %v", val, err)
				}
			}
		}
	}
	wait()

	t.Logf("testing GET with %d keys and %d parallelism\n", keys*2, para)
	jobs, wait = parallel(para)
	for i := 0; i < keys*2; i += batch {
		cmdkeys := make([]string, 0, batch)
		commands := make(Commands, 0, batch)
		for j := 0; j < batch; j++ {
			cmdkeys = append(cmdkeys, "m"+strconv.Itoa(rand.Intn(keys*2)))
			commands = append(commands, client.B().Get().Key(cmdkeys[len(cmdkeys)-1]).Build())
		}
		jobs <- func() {
			for j, resp := range client.DoMulti(ctx, commands...) {
				val, err := resp.ToString()
				if v, ok := kvs[cmdkeys[j]]; !((ok && val == v) || (!ok && IsRedisNil(err))) {
					t.Errorf("unexpected get response %v %v %v", val, err, ok)
				}
			}
		}
	}
	wait()

	t.Logf("testing client side caching with %d interations and %d parallelism\n", keys*5, para)
	jobs, wait = parallel(para)
	hits, miss := int64(0), int64(0)
	for i := 0; i < keys*10; i += batch {
		cmdkeys := make([]string, 0, batch)
		commands := make([]CacheableTTL, 0, batch)
		for j := 0; j < batch; j++ {
			cmdkeys = append(cmdkeys, "m"+strconv.Itoa(rand.Intn(keys/100)))
			commands = append(commands, CT(client.B().Get().Key(cmdkeys[len(cmdkeys)-1]).Cache(), time.Minute))
		}
		jobs <- func() {
			for j, resp := range client.DoMultiCache(ctx, commands...) {
				val, err := resp.ToString()
				if v, ok := kvs[cmdkeys[j]]; !((ok && val == v) || (!ok && IsRedisNil(err))) {
					t.Errorf("unexpected csc get response %v %v %v", val, err, ok)
				}
				if resp.IsCacheHit() {
					atomic.AddInt64(&hits, 1)
				} else {
					atomic.AddInt64(&miss, 1)
				}
			}
		}
	}
	wait()
	if csc {
		if atomic.LoadInt64(&miss) != 100 || atomic.LoadInt64(&hits) != int64(keys*10-100) {
			t.Fatalf("unexpected client side caching hits and miss %v %v", atomic.LoadInt64(&hits), atomic.LoadInt64(&miss))
		}
	} else {
		if atomic.LoadInt64(&hits) != 0 {
			t.Fatalf("unexpected client side caching hits and miss %v %v", atomic.LoadInt64(&hits), atomic.LoadInt64(&miss))
		}
	}

	t.Logf("testing DEL with %d keys and %d parallelism\n", keys*2, para)
	jobs, wait = parallel(para)
	for i := 0; i < keys*2; i += batch {
		cmdkeys := make([]string, 0, batch)
		commands := make(Commands, 0, batch)
		for j := 0; j < batch; j++ {
			cmdkeys = append(cmdkeys, "m"+strconv.Itoa(i+j))
			commands = append(commands, client.B().Del().Key(cmdkeys[len(cmdkeys)-1]).Build())
		}
		jobs <- func() {
			for j, resp := range client.DoMulti(ctx, commands...) {
				val, err := resp.AsInt64()
				if _, ok := kvs[cmdkeys[j]]; !((val == 1 && ok) || (val == 0 && !ok)) {
					t.Errorf("unexpected del response %v %v %v", val, err, ok)
				}
			}
		}
	}
	wait()

	t.Logf("testing client side caching after delete\n")
	jobs, wait = parallel(para)
	for i := 0; i < keys/100; i += batch {
		cmdkeys := make([]string, 0, batch)
		commands := make([]CacheableTTL, 0, batch)
		for j := 0; j < batch; j++ {
			cmdkeys = append(cmdkeys, "m"+strconv.Itoa(i+j))
			commands = append(commands, CT(client.B().Get().Key(cmdkeys[len(cmdkeys)-1]).Cache(), time.Minute))
		}
		jobs <- func() {
			for _, resp := range client.DoMultiCache(ctx, commands...) {
				if !IsRedisNil(resp.Error()) {
					t.Errorf("unexpected csc get response after delete %v", resp)
				}
				if resp.IsCacheHit() {
					t.Errorf("unexpected csc cache hit after delete")
				}
			}
		}
	}
	wait()
}

func testBlockingZPOP(t *testing.T, client Client) {
	ctx := context.Background()
	key := "bz_pop_test"
	items := 2000

	client.Do(ctx, client.B().Del().Key(key).Build())

	t.Logf("testing BZPOPMIN blocking concurrently with ZADD with %d items\n", items)
	go func() {
		for i := 0; i < items; i++ {
			v, err := client.Do(ctx, client.B().Zadd().Key(key).ScoreMember().ScoreMember(float64(i), strconv.Itoa(i)).Build()).AsInt64()
			if err != nil || v != 1 {
				t.Errorf("unexpected ZADD response %v %v", v, err)
			}
		}
	}()
	for i := 0; i < items; i++ {
		arr, err := client.Do(ctx, client.B().Bzpopmin().Key(key).Timeout(0).Build()).AsStrSlice()
		if err != nil || (arr[0] != key || arr[1] != arr[2] || arr[1] != strconv.Itoa(i)) {
			t.Errorf("unexpected BZPOPMIN response %v %v", arr, err)
		}
	}
	client.Do(ctx, client.B().Del().Key(key).Build())
}

func testBlockingXREAD(t *testing.T, client Client) {
	ctx := context.Background()
	key := "xread_test"
	items := 2000

	client.Do(ctx, client.B().Del().Key(key).Build())

	t.Logf("testing blocking XREAD concurrently with XADD with %d items\n", items)
	go func() {
		for i := 0; i < items; i++ {
			v := strconv.Itoa(i)
			v, err := client.Do(ctx, client.B().Xadd().Key(key).Id("*").FieldValue().FieldValue(v, v).Build()).ToString()
			if err != nil || v == "" {
				t.Errorf("unexpected XADD response %v %v", v, err)
			}
		}
	}()
	id := "0"
	for i := 0; i < items; i++ {
		m, err := client.Do(ctx, client.B().Xread().Count(1).Block(0).Streams().Key(key).Id(id).Build()).AsXRead()
		if err != nil {
			t.Errorf("unexpected XREAD response %v %v", m, err)
		}
		id = m[key][0].ID
		if len(m[key][0].FieldValues) == 0 {
			t.Errorf("unexpected XREAD response %v %v", m, err)
		}
		for f, v := range m[key][0].FieldValues {
			if f != v || f != strconv.Itoa(i) {
				t.Errorf("unexpected XREAD response %v %v", m, err)
			}
		}
	}
	client.Do(ctx, client.B().Del().Key(key).Build())
}

func testPubSub(t *testing.T, client Client) {
	msgs := 10000
	mmap := make(map[string]struct{})
	for i := 0; i < msgs; i++ {
		mmap[strconv.Itoa(i)] = struct{}{}
	}
	t.Logf("testing pubsub with %v messages\n", msgs)
	jobs, wait := parallel(10)

	ctx := context.Background()

	messages := make(chan string, 10)
	go func() {
		err := client.Receive(ctx, client.B().Subscribe().Channel("ch1").Build(), func(msg PubSubMessage) {
			messages <- msg.Message
		})
		if err != ErrClosing {
			t.Errorf("unexpected subscribe response %v", err)
		}
	}()

	go func() {
		err := client.Receive(ctx, client.B().Psubscribe().Pattern("pat*").Build(), func(msg PubSubMessage) {
			messages <- msg.Message
		})
		if err != ErrClosing {
			t.Errorf("unexpected subscribe response %v", err)
		}
	}()

	go func() {
		time.Sleep(time.Second)
		for i := 0; i < msgs; i++ {
			msg := strconv.Itoa(i)
			ch := "ch1"
			if i%10 == 0 {
				ch = "pat1"
			}
			jobs <- func() {
				if err := client.Do(context.Background(), client.B().Publish().Channel(ch).Message(msg).Build()).Error(); err != nil {
					t.Errorf("unexpected publish response %v", err)
				}
			}
		}
		wait()
	}()

	for message := range messages {
		delete(mmap, message)
		if len(mmap) == 0 {
			close(messages)
		}
	}
}

func run(t *testing.T, client Client, cases ...func(*testing.T, Client)) {
	wg := sync.WaitGroup{}
	wg.Add(len(cases))
	for _, c := range cases {
		c := c
		go func() {
			c(t, client)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestSingleClientIntegration(t *testing.T) {
	client, err := NewClient(ClientOption{
		InitAddress:      []string{"127.0.0.1:6379"},
		ConnWriteTimeout: 180 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	run(t, client, testSETGETCSC, testMultiSETGETCSC, testBlockingZPOP, testBlockingXREAD, testPubSub)
	run(t, client, testFlush)

	client.Close()
	time.Sleep(time.Second * 5) // wait background ping exit
}

func TestSentinelClientIntegration(t *testing.T) {
	client, err := NewClient(ClientOption{
		InitAddress:      []string{"127.0.0.1:26379"},
		ConnWriteTimeout: 180 * time.Second,
		Sentinel: SentinelOption{
			MasterSet: "test",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	run(t, client, testSETGETCSC, testMultiSETGETCSC, testBlockingZPOP, testBlockingXREAD, testPubSub)
	run(t, client, testFlush)

	client.Close()
	time.Sleep(time.Second * 5) // wait background ping exit
}

func TestClusterClientIntegration(t *testing.T) {
	client, err := NewClient(ClientOption{
		InitAddress:      []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"},
		ConnWriteTimeout: 180 * time.Second,
		ShuffleInit:      true,
	})
	if err != nil {
		t.Fatal(err)
	}
	run(t, client, testSETGETCSC, testMultiSETGETCSC, testBlockingZPOP, testBlockingXREAD, testPubSub)

	client.Close()
	time.Sleep(time.Second * 5) // wait background ping exit
}

func TestSingleClient5Integration(t *testing.T) {
	client, err := NewClient(ClientOption{
		InitAddress:      []string{"127.0.0.1:6355"},
		ConnWriteTimeout: 180 * time.Second,
		DisableCache:     true,
	})
	if err != nil {
		t.Fatal(err)
	}

	run(t, client, testSETGETRESP2, testMultiSETGETRESP2, testBlockingZPOP, testBlockingXREAD)

	client.Close()
	time.Sleep(time.Second * 5) // wait background ping exit
}

func TestCluster5ClientIntegration(t *testing.T) {
	client, err := NewClient(ClientOption{
		InitAddress:      []string{"127.0.0.1:7004", "127.0.0.1:7005", "127.0.0.1:7006"},
		ConnWriteTimeout: 180 * time.Second,
		ShuffleInit:      true,
		DisableCache:     true,
	})
	if err != nil {
		t.Fatal(err)
	}
	run(t, client, testSETGETRESP2, testMultiSETGETRESP2, testBlockingZPOP, testBlockingXREAD)

	client.Close()
	time.Sleep(time.Second * 5) // wait background ping exit
}
