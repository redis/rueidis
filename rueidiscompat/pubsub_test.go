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
	"bytes"
	"context"
	"sync"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PubSub", func() {
	var client Cmdable

	BeforeEach(func() {
		client = adapterresp3
	})

	It("implements Stringer", func() {
		pubsub := client.PSubscribe(ctx, "mychannel*")
		defer pubsub.Close()

		Expect(pubsub.String()).To(Equal("PubSub(mychannel*)"))
	})

	It("should support pattern matching", func() {
		pubsub := client.PSubscribe(ctx, "mychannel*")
		defer pubsub.Close()

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Subscription)
			Expect(subscr.Kind).To(Equal("psubscribe"))
			Expect(subscr.Channel).To(Equal("mychannel*"))
			Expect(subscr.Count).To(Equal(1))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).To(MatchError(context.DeadlineExceeded))
			Expect(msgi).To(BeNil())
		}

		n, err := client.Publish(ctx, "mychannel1", "hello").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(Equal(int64(1)))

		Expect(pubsub.PUnsubscribe(ctx, "mychannel*")).NotTo(HaveOccurred())

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Message)
			Expect(subscr.Channel).To(Equal("mychannel1"))
			Expect(subscr.Pattern).To(Equal("mychannel*"))
			Expect(subscr.Payload).To(Equal("hello"))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Subscription)
			Expect(subscr.Kind).To(Equal("punsubscribe"))
			Expect(subscr.Channel).To(Equal("mychannel*"))
			Expect(subscr.Count).To(Equal(0))
		}
	})

	It("should pub/sub channels", func() {
		channels, err := client.PubSubChannels(ctx, "mychannel*").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(channels).To(BeEmpty())

		pubsub := client.Subscribe(ctx, "mychannel", "mychannel2")
		defer pubsub.Close()

		channels, err = client.PubSubChannels(ctx, "mychannel*").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(channels).To(ConsistOf([]string{"mychannel", "mychannel2"}))

		channels, err = client.PubSubChannels(ctx, "").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(channels).To(BeEmpty())

		channels, err = client.PubSubChannels(ctx, "*").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(len(channels)).To(BeNumerically(">=", 2))
	})

	It("should sharded pub/sub channels", func() {
		channels, err := client.PubSubShardChannels(ctx, "mychannel*").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(channels).To(BeEmpty())

		pubsub := client.SSubscribe(ctx, "mychannel", "mychannel2")
		defer pubsub.Close()

		channels, err = client.PubSubShardChannels(ctx, "mychannel*").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(channels).To(ConsistOf([]string{"mychannel", "mychannel2"}))

		channels, err = client.PubSubShardChannels(ctx, "").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(channels).To(BeEmpty())

		channels, err = client.PubSubShardChannels(ctx, "*").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(len(channels)).To(BeNumerically(">=", 2))

		nums, err := client.PubSubShardNumSub(ctx, "mychannel", "mychannel2", "mychannel3").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(nums).To(Equal(map[string]int64{
			"mychannel":  1,
			"mychannel2": 1,
			"mychannel3": 0,
		}))
	})

	It("should return the numbers of subscribers", func() {
		pubsub := client.Subscribe(ctx, "mychannel", "mychannel2")
		defer pubsub.Close()

		channels, err := client.PubSubNumSub(ctx, "mychannel", "mychannel2", "mychannel3").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(channels).To(Equal(map[string]int64{
			"mychannel":  1,
			"mychannel2": 1,
			"mychannel3": 0,
		}))
	})

	It("should return the numbers of subscribers by pattern", func() {
		num, err := client.PubSubNumPat(ctx).Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(num).To(Equal(int64(0)))

		pubsub := client.PSubscribe(ctx, "*")
		defer pubsub.Close()

		num, err = client.PubSubNumPat(ctx).Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(num).To(Equal(int64(1)))
	})

	It("should pub/sub", func() {
		pubsub := client.Subscribe(ctx, "mychannel", "mychannel2")
		defer pubsub.Close()

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Subscription)
			Expect(subscr.Kind).To(Equal("subscribe"))
			Expect(subscr.Channel).To(Equal("mychannel"))
			Expect(subscr.Count).To(Equal(1))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Subscription)
			Expect(subscr.Kind).To(Equal("subscribe"))
			Expect(subscr.Channel).To(Equal("mychannel2"))
			Expect(subscr.Count).To(Equal(2))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).To(MatchError(context.DeadlineExceeded))
			Expect(msgi).NotTo(HaveOccurred())
		}

		n, err := client.Publish(ctx, "mychannel", "hello").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(Equal(int64(1)))

		n, err = client.Publish(ctx, "mychannel2", "hello2").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(Equal(int64(1)))

		Expect(pubsub.Unsubscribe(ctx, "mychannel", "mychannel2")).NotTo(HaveOccurred())

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			msg := msgi.(*Message)
			Expect(msg.Channel).To(Equal("mychannel"))
			Expect(msg.Payload).To(Equal("hello"))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			msg := msgi.(*Message)
			Expect(msg.Channel).To(Equal("mychannel2"))
			Expect(msg.Payload).To(Equal("hello2"))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Subscription)
			Expect(subscr.Kind).To(Equal("unsubscribe"))
			Expect(subscr.Channel).To(Equal("mychannel"))
			Expect(subscr.Count).To(Equal(1))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Subscription)
			Expect(subscr.Kind).To(Equal("unsubscribe"))
			Expect(subscr.Channel).To(Equal("mychannel2"))
			Expect(subscr.Count).To(Equal(0))
		}
	})

	It("should sharded pub/sub", func() {
		pubsub := client.SSubscribe(ctx, "mychannel", "mychannel2")
		defer pubsub.Close()

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Subscription)
			Expect(subscr.Kind).To(Equal("ssubscribe"))
			Expect(subscr.Channel).To(Equal("mychannel"))
			Expect(subscr.Count).To(Equal(1))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Subscription)
			Expect(subscr.Kind).To(Equal("ssubscribe"))
			Expect(subscr.Channel).To(Equal("mychannel2"))
			Expect(subscr.Count).To(Equal(2))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).To(MatchError(context.DeadlineExceeded))
			Expect(msgi).NotTo(HaveOccurred())
		}

		n, err := client.SPublish(ctx, "mychannel", "hello").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(Equal(int64(1)))

		n, err = client.SPublish(ctx, "mychannel2", "hello2").Result()
		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(Equal(int64(1)))

		Expect(pubsub.SUnsubscribe(ctx, "mychannel", "mychannel2")).NotTo(HaveOccurred())

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			msg := msgi.(*Message)
			Expect(msg.Channel).To(Equal("mychannel"))
			Expect(msg.Payload).To(Equal("hello"))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			msg := msgi.(*Message)
			Expect(msg.Channel).To(Equal("mychannel2"))
			Expect(msg.Payload).To(Equal("hello2"))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Subscription)
			Expect(subscr.Kind).To(Equal("sunsubscribe"))
			Expect(subscr.Channel).To(Equal("mychannel"))
			Expect(subscr.Count).To(Equal(1))
		}

		{
			msgi, err := pubsub.ReceiveTimeout(ctx, time.Second)
			Expect(err).NotTo(HaveOccurred())
			subscr := msgi.(*Subscription)
			Expect(subscr.Kind).To(Equal("sunsubscribe"))
			Expect(subscr.Channel).To(Equal("mychannel2"))
			Expect(subscr.Count).To(Equal(0))
		}
	})

	It("should multi-ReceiveMessage", func() {
		pubsub := client.Subscribe(ctx, "mychannel")
		defer pubsub.Close()

		subscr, err := pubsub.ReceiveTimeout(ctx, time.Second)
		Expect(err).NotTo(HaveOccurred())
		Expect(subscr).To(Equal(&Subscription{
			Kind:    "subscribe",
			Channel: "mychannel",
			Count:   1,
		}))

		err = client.Publish(ctx, "mychannel", "hello").Err()
		Expect(err).NotTo(HaveOccurred())

		err = client.Publish(ctx, "mychannel", "world").Err()
		Expect(err).NotTo(HaveOccurred())

		msg, err := pubsub.ReceiveMessage(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(msg.Channel).To(Equal("mychannel"))
		Expect(msg.Payload).To(Equal("hello"))

		msg, err = pubsub.ReceiveMessage(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(msg.Channel).To(Equal("mychannel"))
		Expect(msg.Payload).To(Equal("world"))
	})

	It("should return on Close", func() {
		pubsub := client.Subscribe(ctx, "mychannel")
		defer pubsub.Close()

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer GinkgoRecover()

			wg.Done()
			defer wg.Done()

			_, err := pubsub.ReceiveMessage(ctx)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(SatisfyAny(
				Equal("redis: client is closed"),
				ContainSubstring("use of closed network connection"),
			))
		}()

		wg.Wait()
		wg.Add(1)

		Expect(pubsub.Close()).NotTo(HaveOccurred())

		wg.Wait()
	})

	It("should ReceiveMessage without a subscription", func() {
		timeout := 100 * time.Millisecond

		pubsub := client.Subscribe(ctx)
		defer pubsub.Close()

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer GinkgoRecover()
			defer wg.Done()

			time.Sleep(timeout)

			err := pubsub.Subscribe(ctx, "mychannel")
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(timeout)

			err = client.Publish(ctx, "mychannel", "hello").Err()
			Expect(err).NotTo(HaveOccurred())
		}()

		msg, err := pubsub.ReceiveMessage(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(msg.Channel).To(Equal("mychannel"))
		Expect(msg.Payload).To(Equal("hello"))

		wg.Wait()
	})

	It("handles big message payload", func() {
		pubsub := client.Subscribe(ctx, "mychannel")
		defer pubsub.Close()

		ch := pubsub.Channel()

		bigVal := bigVal()
		err := client.Publish(ctx, "mychannel", bigVal).Err()
		Expect(err).NotTo(HaveOccurred())

		var msg *Message
		Eventually(ch).WithTimeout(5 * time.Second).Should(Receive(&msg))
		Expect(msg.Channel).To(Equal("mychannel"))
		Expect(msg.Payload).To(Equal(string(bigVal)))
	})

	It("supports concurrent Publish and Receive", func() {
		const N = 100

		pubsub := client.Subscribe(ctx, "mychannel")
		defer pubsub.Close()

		done := make(chan struct{})
		go func() {
			defer GinkgoRecover()

			for i := 0; i < N; i++ {
				_, err := pubsub.ReceiveTimeout(ctx, 5*time.Second)
				Expect(err).NotTo(HaveOccurred())
			}
			close(done)
		}()

		for i := 0; i < N; i++ {
			err := client.Publish(ctx, "mychannel", "hello").Err()
			Expect(err).NotTo(HaveOccurred())
		}

		select {
		case <-done:
		case <-time.After(30 * time.Second):
			Fail("timeout")
		}
	})

	It("should ChannelMessage", func() {
		pubsub := client.Subscribe(ctx, "mychannel")
		defer pubsub.Close()

		ch := pubsub.Channel(
			WithChannelSize(10),
			WithChannelHealthCheckInterval(time.Second),
		)

		text := "test channel message"
		err := client.Publish(ctx, "mychannel", text).Err()
		Expect(err).NotTo(HaveOccurred())

		var msg *Message
		Eventually(ch).Should(Receive(&msg))
		Expect(msg.Channel).To(Equal("mychannel"))
		Expect(msg.Payload).To(Equal(text))
	})
})

func bigVal() []byte {
	return bytes.Repeat([]byte{'*'}, 1<<17) // 128kb
}

var _ = Describe("WithChannelSize", func() {
	It("should set the channel size correctly", func() {
		customSize := 500
		cfg := &chopt{}
		WithChannelSize(customSize)(cfg)

		Expect(cfg.chanSize).To(Equal(customSize))
	})
})

var _ = Describe("Subscription", func() {
	It("should return the correct string representation", func() {
		sub := &Subscription{Kind: "subscribe", Channel: "channel"}
		Expect(sub.String()).To(Equal("subscribe: channel"))
	})
})

var _ = Describe("Message", func() {
	It("should return the correct string representation", func() {
		m := &Message{Channel: "channel", Payload: "new"}
		Expect(m.String()).To(Equal("Message<channel: new>"))
	})
})
