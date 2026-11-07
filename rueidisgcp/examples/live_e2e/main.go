// Command live_e2e keeps SET and GET traffic running against a cluster-mode Memorystore for
// Valkey instance with IAM authentication, to check that requests keep succeeding while the
// IAM token is refreshed. rueidisgcp schedules a re-AUTH on each live connection 1m52s to
// 3m45s before its token expires, and rueidis then sends AUTH with a new token. Access
// tokens last an hour, so leave the program running for more than an hour.
//
// It needs a host that can reach the instance's discovery endpoint, such as a VM in the
// instance's VPC network. The tokens come from Application Default Credentials; that
// principal needs roles/memorystore.dbConnectionUser on the instance. Run it from the
// rueidisgcp/examples directory and stop it with Ctrl+C:
//
//	MEMORYSTORE_TEST_DISCOVERY_ENDPOINT=10.0.0.15:6379 go run ./live_e2e
//
// It runs one SET/GET pair every 200ms, logs each failure, and logs the counts after every
// 300 successful pairs, about once a minute. When stopped, it prints the totals and exits
// with status 1 if any pair failed.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisgcp"
)

func main() {
	endpoint := os.Getenv("MEMORYSTORE_TEST_DISCOVERY_ENDPOINT")
	if endpoint == "" {
		log.Fatal("MEMORYSTORE_TEST_DISCOVERY_ENDPOINT environment variable must be set (e.g. 10.0.0.15:6379)")
	}

	ctx := context.Background()
	log.Printf("Connecting to Memorystore at %s with Native IAM Auth...", endpoint)

	client, err := rueidisgcp.NewClient(ctx, rueidis.ClientOption{
		InitAddress: []string{endpoint}, // discovery endpoint
	}, rueidisgcp.WithDefaultIAM()) // IAM tokens from Application Default Credentials
	if err != nil {
		log.Fatalf("Failed to create Memorystore client: %v", err)
	}
	defer client.Close()

	// Initial PING to verify authentication handshake
	if err := client.Do(ctx, client.B().Ping().Build()).Error(); err != nil {
		log.Fatalf("Initial PING failed! Authentication rejected: %v", err)
	}
	log.Println("✅ Initial RESP3 HELLO 3 AUTH handshake succeeded!")

	// Run continuous traffic loop
	var successCount uint64
	var errorCount uint64
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	log.Println("Running continuous traffic test. Monitoring across token rotation (press Ctrl+C to stop)...")

	startTime := time.Now()
	for {
		select {
		case <-ticker.C:
			key := fmt.Sprintf("e2e-iam-key-%d", time.Now().UnixNano()%100)
			val := fmt.Sprintf("val-%d", time.Now().UnixNano())

			// Execute SET
			err := client.Do(ctx, client.B().Set().Key(key).Value(val).Ex(60*time.Second).Build()).Error()
			if err != nil {
				atomic.AddUint64(&errorCount, 1)
				log.Printf("❌ SET failed: %v", err)
				continue
			}

			// Execute GET
			getVal, err := client.Do(ctx, client.B().Get().Key(key).Build()).ToString()
			if err != nil {
				atomic.AddUint64(&errorCount, 1)
				log.Printf("❌ GET failed: %v", err)
				continue
			}
			if getVal != val {
				atomic.AddUint64(&errorCount, 1)
				log.Printf("❌ GET %s returned %q, want %q", key, getVal, val)
				continue
			}

			count := atomic.AddUint64(&successCount, 1)
			if count%300 == 0 {
				log.Printf("[Elapsed: %v] Operations: %d succeeded, %d failed",
					time.Since(startTime).Round(time.Second),
					atomic.LoadUint64(&successCount),
					atomic.LoadUint64(&errorCount),
				)
			}

		case <-stopChan:
			failed := atomic.LoadUint64(&errorCount)
			log.Printf("Stopping test after %v: %d succeeded, %d failed",
				time.Since(startTime).Round(time.Second), atomic.LoadUint64(&successCount), failed)
			if failed > 0 {
				client.Close()
				os.Exit(1) // any failed pair fails the run
			}
			return
		}
	}
}
