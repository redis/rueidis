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
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redis/rueidis"
)

var _ = Describe("Scripting", func() {
	var client rueidis.Client
	var rdb Cmdable
	var ctx context.Context

	BeforeEach(func() {
		var err error
		ctx = context.Background()
		// TODO: make this configurable for CI
		client, err = rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
		Expect(err).NotTo(HaveOccurred())
		rdb = NewAdapter(client)
		// It's good practice to flush the DB before each test to ensure independence
		err = rdb.FlushDB(ctx).Err()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		if client != nil {
			client.Close()
		}
	})

	// Keep the ExampleScript logic, slightly adapted for Ginkgo's structure if needed
	// Or, ensure it can run alongside. For now, let's integrate its core logic into a test.
	Context("ExampleScript logic", func() {
		It("should behave as in the example", func() {
			IncrByXX := NewScript(`
				if redis.call("GET", KEYS[1]) ~= false then
					return redis.call("INCRBY", KEYS[1], ARGV[1])
				end
				return false
			`)

			n, err := IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
			Expect(err).To(MatchError("redis: nil")) // Or use a specific error check if rueidis provides one for script returning false like this
			Expect(n).To(BeEquivalentTo(0)) // Or check for nil depending on how rueidis handles 'false' from Lua

			err = rdb.Set(ctx, "xx_counter", "40", 0).Err()
			Expect(err).NotTo(HaveOccurred())

			n, err = IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(BeEquivalentTo(42))
		})
	})

	Describe("Return Types", func() {
		It("should return an integer", func() {
			script := NewScript("return 123")
			val, err := script.Run(ctx, rdb, nil).Int()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal(123))
		})

		It("should return a string", func() {
			script := NewScript("return 'hello world'")
			val, err := script.Run(ctx, rdb, nil).Text()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("hello world"))
		})

		It("should return an array/slice", func() {
			script := NewScript("return {1, 'two', 3}")
			val, err := script.Run(ctx, rdb, nil).Slice()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(HaveLen(3))
			Expect(val[0]).To(BeEquivalentTo(1))
			Expect(val[1]).To(BeEquivalentTo("two"))
			Expect(val[2]).To(BeEquivalentTo(3))
		})

		It("should return a boolean (true)", func() {
			script := NewScript("return true")
			val, err := script.Run(ctx, rdb, nil).Bool()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(BeTrue())
		})

		It("should return a boolean (false)", func() {
			script := NewScript("return false")
			// Note: Redis Lua 'false' is returned as a nil reply.
			// The behavior of .Bool() on a nil reply might vary.
			// For go-redis, it would be an error or false. Let's check rueidis behavior.
			// If it's an error, then the test should expect an error.
			// If it converts nil to false, then the test should expect false.
			// Based on typical Redis clients, a Lua 'false' often results in a 'nil' bulk string reply.
			cmd := script.Run(ctx, rdb, nil)
			Expect(cmd.Err()).To(MatchError("redis: nil"))
		})
	})

	Describe("Argument Types", func() {
		It("should handle multiple arguments", func() {
			// Script sums two numbers and concatenates with a string
			script := NewScript("return tonumber(ARGV[1]) + tonumber(ARGV[2]) .. ARGV[3]")
			keys := []string{}
			args := []interface{}{5, 10, "hello"}
			val, err := script.Run(ctx, rdb, keys, args...).Text()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("15hello"))
		})
	})

	Describe("Error Handling", func() {
		It("should propagate Lua runtime error", func() {
			script := NewScript("return redis.call('unknown_command')")
			err := script.Run(ctx, rdb, nil).Err()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("ERR Error running script")) // Check for part of the typical Redis Lua error message
		})

		It("should propagate Lua syntax error (if Eval catches it)", func() {
			// Note: Syntax errors are usually caught at SCRIPT LOAD time, not EVAL time.
			// If NewScript doesn't implicitly load, this might not be testable here.
			// Let's try loading it explicitly.
			script := NewScript("return 'missing quote")
			err := script.Load(ctx, rdb).Err() // Try to load it
			Expect(err).To(HaveOccurred())     // Expect error during load
			Expect(err.Error()).To(SatisfyAny(
				ContainSubstring("ERR Error compiling script"), // Common error
				ContainSubstring("Unmatched quote"),             // More specific
			))
		})
	})

	Describe("Script Loading and Management", func() {
		var testScript *Script
		var scriptBody string

		BeforeEach(func() {
			scriptBody = "return KEYS[1] .. ARGV[1]"
			testScript = NewScript(scriptBody)
		})

		It("should load a script and allow execution via EVALSHA", func() {
			// Ensure script is not already loaded by flushing scripts
			err := rdb.ScriptFlush(ctx).Err()
			Expect(err).NotTo(HaveOccurred())
			
			sha1, err := testScript.Load(ctx, rdb).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(sha1).To(Equal(testScript.Hash()))

			// Execute using EvalSha
			val, err := rdb.EvalSha(ctx, sha1, []string{"mykey"}, "myvalue").Text()
			Expect(err).NotTo(HaveOccurred())
			Expect(val).To(Equal("mykeymyvalue"))
		})

		It("should check for script existence", func() {
			// Ensure script is not already loaded
			err := rdb.ScriptFlush(ctx).Err()
			Expect(err).NotTo(HaveOccurred())

			exists, err := testScript.Exists(ctx, rdb).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(BeFalse()) // Initially should not exist

			_, err = testScript.Load(ctx, rdb).Result()
			Expect(err).NotTo(HaveOccurred())

			exists, err = testScript.Exists(ctx, rdb).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(BeTrue()) // Should exist after loading

			// Test with a non-existent script (random hash)
			nonExistentScript := NewScript("return 'this was never loaded'")
			// Ensure we use a hash that's unlikely to exist
			// We can't just use nonExistentScript.Hash() if it's not loaded.
			// We need to call Exists with the hash directly.
			// The Script.Exists method handles this internally.
			// Let's create another script that is definitely not loaded.
			otherScript := NewScript("return 'another script'")
			// Ensure scripts are flushed so only testScript is loaded
			err = rdb.ScriptFlush(ctx).Err()
			Expect(err).NotTo(HaveOccurred())
			_, err = testScript.Load(ctx, rdb).Result() // Load the one we want to exist
			Expect(err).NotTo(HaveOccurred())


			exists, err = otherScript.Exists(ctx, rdb).Result() // Check the one not loaded
			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(BeFalse())
		})

		It("should flush scripts", func() {
			_, err := testScript.Load(ctx, rdb).Result()
			Expect(err).NotTo(HaveOccurred())

			exists, err := testScript.Exists(ctx, rdb).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(BeTrue())

			err = rdb.ScriptFlush(ctx).Err()
			Expect(err).NotTo(HaveOccurred())

			exists, err = testScript.Exists(ctx, rdb).Result()
			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(BeFalse()) // Should not exist after flush
		})

		It("should attempt to kill scripts (basic call)", func() {
			// This is hard to test meaningfully without a long-running script.
			// We'll just call it and expect no error.
			// Note: SCRIPT KILL is meant to kill a script currently running.
			// If no script is running, it usually returns OK.
			err := rdb.ScriptKill(ctx).Err()
			// Redis returns "UNKILLABLE" if no script is running, which is not an error for the client library.
			// Or it might return "NOSCRIPT" if no script was ever run or client is in pubsub mode.
			// For a basic call, we expect it not to fail in a way that the client lib reports a connection error.
			// The command itself might return a specific Redis error string, which is fine.
			if err != nil {
				Expect(err.Error()).To(SatisfyAny(
					ContainSubstring("UNKILLABLE"), // No script busy executing.
					ContainSubstring("NOTBUSY"),    // Older Redis versions for the same condition.
					ContainSubstring("NOSCRIPT"),   // If no script was specified to be killed by a SCRIPT KILL <sha> command (not applicable here).
				), "Expected specific Redis messages or no error, but got: "+err.Error())
			} else {
				Expect(err).NotTo(HaveOccurred())
			}
		})
	})
})

// Example function - can be removed if its logic is fully covered by Ginkgo tests
// Or kept for documentation purposes.
// For now, its core logic is in "ExampleScript logic" context.
func ExampleScript() {
	// This is the original example function.
	// To make it runnable as a standalone example, you'd need to set up `ctx` and `rdb`
	// similar to how it's done in BeforeEach.
	// For Ginkgo tests, this setup is handled by BeforeEach.

	// Local setup for example to run independently (if desired)
	localCtx := context.Background()
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		fmt.Printf("Error creating client for example: %v\n", err)
		return
	}
	defer client.Close()
	localRdb := NewAdapter(client)

	// Ensure DB is clean for the example
	_ = localRdb.FlushDB(localCtx).Err()


	IncrByXX := NewScript(`
		if redis.call("GET", KEYS[1]) ~= false then
			return redis.call("INCRBY", KEYS[1], ARGV[1])
		end
		return false
	`)

	n, err := IncrByXX.Run(localCtx, localRdb, []string{"xx_counter"}, 2).Result()
	fmt.Printf("Run 1: %v, %v\n", n, err) // Expect n=0 (or nil), err=redis: nil

	err = localRdb.Set(localCtx, "xx_counter", "40", 0).Err()
	if err != nil {
		fmt.Printf("Error setting counter: %v\n", err)
		return
	}

	n, err = IncrByXX.Run(localCtx, localRdb, []string{"xx_counter"}, 2).Result()
	fmt.Printf("Run 2: %v, %v\n", n, err) // Expect n=42, err=nil

	// Output:
	// Run 1: 0, redis: nil
	// Run 2: 42, <nil>
}
