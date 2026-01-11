# Implementation Plan: DoWithReader API

## Overview

Add a new `DoWithReader()` method to rueidis that provides zero-allocation parsing of Redis responses by giving applications direct access to the underlying `bufio.Reader`. This method will automatically handle all corner cases (cluster redirects, retries, connection expiration) just like `Do()`, but delegate successful response parsing to the application.

## Motivation

Currently, `RedisResult` methods like `ToString()`, `AsStrSlice()`, `ToArray()` etc. allocate intermediate Go structures when parsing Redis responses. For high-performance applications, these allocations can be significant overhead.

The existing `DoStream()` method provides zero-allocation access but:
1. Only works with string, integer, or float responses
2. Does NOT handle cluster redirects automatically
3. Cannot parse arrays, maps, or nested structures

`DoWithReader()` will:
1. Work with ALL Redis response types (arrays, maps, sets, nested structures)
2. Handle cluster redirects (MOVED/ASK) automatically
3. Handle retries (TRYAGAIN, LOADING, connection errors) automatically
4. Only call user callback for successful (non-error) responses

## Architecture

### Two-Layer Design

```
┌─────────────────────────────────────────────────────────────────┐
│                     Client Layer (High-Level)                    │
│  Handles: Redirects, Retries, Connection Selection, Cmd Recycle │
├─────────────────────────────────────────────────────────────────┤
│  clusterClient.DoWithReader()  - MOVED/ASK handling             │
│  sentinelClient.DoWithReader() - Retry handling                 │
│  standalone.DoWithReader()     - REDIRECT handling              │
│  singleClient.DoWithReader()   - Retry handling                 │
│  mux.DoWithReader()            - Pool management                │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                     Wire Layer (Low-Level)                       │
│  Handles: Protocol parsing, Error detection, Push messages       │
├─────────────────────────────────────────────────────────────────┤
│  pipe.DoWithReader()                                             │
│  - Sends command                                                 │
│  - Reads response type                                           │
│  - If ERROR: Parse fully, return RedisError (for redirect check) │
│  - If SUCCESS: Call user callback with raw reader                │
└─────────────────────────────────────────────────────────────────┘
```

## Files to Modify

### 1. `rueidis.go` - Interface Definitions

**Add new type:**
```go
// ReaderFunc is called with direct access to the bufio.Reader containing
// the Redis response. The respType parameter is the RESP type byte.
// The reader is only valid during callback execution.
type ReaderFunc func(reader *bufio.Reader, respType byte) error
```

**Modify `Client` interface (add after DoMultiStream):**
```go
// DoWithReader sends a command and provides direct access to the raw RESP response
// through a callback function for zero-allocation parsing.
//
// Unlike DoStream, DoWithReader:
// - Works with ALL Redis response types (arrays, maps, sets, nested structures)
// - Automatically handles cluster redirects (MOVED/ASK)
// - Automatically handles retries (TRYAGAIN, LOADING, connection errors)
//
// The callback is ONLY invoked for successful (non-error) responses.
// All error handling (redirects, retries) is done automatically by rueidis.
//
// The reader is only valid during callback execution and must not be stored.
// The cmd parameter is recycled after DoWithReader returns.
DoWithReader(ctx context.Context, cmd Completed, fn ReaderFunc) error
```

### 2. `pipe.go` - Wire Layer Implementation

**Add to `wire` interface:**
```go
DoWithReader(ctx context.Context, pool *pool, cmd Completed, fn ReaderFunc) error
```

**Add implementation (after DoMultiStream):**

Key behaviors:
- Set connection deadline (same as DoStream)
- Write and flush command
- Skip RESP3 push messages (type `>`)
- Read response type byte
- **If error type (`-` or `!`)**: Parse fully as RedisError, return it (allows cluster to detect MOVED/ASK)
- **If null type (`_`)**: Return `Nil` error
- **If success**: Call user callback with reader and type byte
- Properly handle connection lifecycle (blcksig, waits, pool.Store)

### 3. `mux.go` - Connection Multiplexer

**Add to `conn` interface:**
```go
DoWithReader(ctx context.Context, cmd Completed, fn ReaderFunc) error
```

**Add implementation:**
```go
func (m *mux) DoWithReader(ctx context.Context, cmd Completed, fn ReaderFunc) error {
    wire := m.spool.Acquire(ctx)
    return wire.DoWithReader(ctx, m.spool, cmd, fn)
}
```

### 4. `client.go` - Single Client

**Add to singleClient:**
```go
func (c *singleClient) DoWithReader(ctx context.Context, cmd Completed, fn ReaderFunc) error {
    attempts := 1
retry:
    err := c.conn.DoWithReader(ctx, cmd, fn)
    if err != nil {
        if err == errConnExpired {
            goto retry
        }
        if c.retry && cmd.IsRetryable() && c.isRetryable(err, ctx) {
            if c.retryHandler.WaitOrSkipRetry(ctx, attempts, cmd, err) {
                attempts++
                goto retry
            }
        }
    }
    if err == nil {
        cmds.PutCompleted(cmd)
    }
    return err
}
```

**Add to dedicatedSingleClient:**
```go
func (c *dedicatedSingleClient) DoWithReader(ctx context.Context, cmd Completed, fn ReaderFunc) error {
    // Similar pattern with wire access
}
```

### 5. `cluster.go` - Cluster Client

**Add to clusterClient:**

Corner cases to handle:
- `errConnExpired` - Retry on same node
- `MOVED` error - Get new node address, retry (with redirect counter)
- `ASK` error - Send ASKING command first, then retry
- `TRYAGAIN` / `CLUSTERDOWN` / `LOADING` - Wait and retry (if retryable)
- Max redirect limit check
- Connection picking based on slot and replica preference

```go
func (c *clusterClient) DoWithReader(ctx context.Context, cmd Completed, fn ReaderFunc) error {
    slot := cmd.Slot()
    cc, err := c.pick(ctx, slot, c.toReplica(cmd))
    if err != nil {
        return err
    }

    redirects := 0
    attempts := 0

retry:
    err = cc.DoWithReader(ctx, cmd, fn)

    if err == errConnExpired {
        goto retry
    }

process:
    switch addr, mode := c.shouldRefreshRetry(err, ctx); mode {
    case RedirectMove:
        redirects++
        if c.opt.ClusterOption.MaxMovedRedirections > 0 &&
           redirects > c.opt.ClusterOption.MaxMovedRedirections {
            cmds.PutCompleted(cmd)
            return err
        }
        ncc := c.redirectOrNew(addr, cc, slot, mode)
    recover1:
        err = ncc.DoWithReader(ctx, cmd, fn)
        if err == errConnExpired {
            goto recover1
        }
        cc = ncc
        goto process

    case RedirectAsk:
        redirects++
        if c.opt.ClusterOption.MaxMovedRedirections > 0 &&
           redirects > c.opt.ClusterOption.MaxMovedRedirections {
            cmds.PutCompleted(cmd)
            return err
        }
        ncc := c.redirectOrNew(addr, cc, slot, mode)
    recover2:
        // Must send ASKING command first for ASK redirect
        if askResp := ncc.Do(ctx, cmds.AskingCmd); askResp.NonRedisError() == nil {
            err = ncc.DoWithReader(ctx, cmd, fn)
            if err == errConnExpired {
                goto recover2
            }
        } else {
            err = askResp.NonRedisError()
        }
        cc = ncc
        goto process

    case RedirectRetry:
        if !c.retry || !cmd.IsRetryable() {
            break
        }
        shouldRetry := c.retryHandler.WaitOrSkipRetry(ctx, attempts, cmd, err)
        if shouldRetry {
            attempts++
            if ncc, nerr := c.pick(ctx, slot, c.toReplica(cmd)); nerr == nil {
                cc = ncc
                goto retry
            }
        }
    }

    cmds.PutCompleted(cmd)
    return err
}
```

**Add to dedicatedClusterClient:**
```go
func (c *dedicatedClusterClient) DoWithReader(ctx context.Context, cmd Completed, fn ReaderFunc) error {
    // Similar to dedicatedClusterClient.Do() but with DoWithReader
    // Handle slot verification, retry logic
}
```

### 6. `sentinel.go` - Sentinel Client

**Add implementation:**
```go
func (c *sentinelClient) DoWithReader(ctx context.Context, cmd Completed, fn ReaderFunc) error {
    attempts := 1
retry:
    cc := c.pick(cmd)
    err := cc.DoWithReader(ctx, cmd, fn)
    if err != nil {
        if err == errConnExpired {
            goto retry
        }
        if c.retry && cmd.IsRetryable() && c.isRetryable(err, ctx) {
            if c.retryHandler.WaitOrSkipRetry(ctx, attempts, cmd, err) {
                attempts++
                goto retry
            }
        }
    }
    if err == nil {
        cmds.PutCompleted(cmd)
    }
    return err
}
```

### 7. `standalone.go` - Standalone with Replicas

**Add implementation:**
```go
func (s *standalone) DoWithReader(ctx context.Context, cmd Completed, fn ReaderFunc) error {
    attempts := 1

    if s.enableRedirect {
        cmd = cmd.Pin()
    }

retry:
    var err error
    if s.toReplicas != nil && s.toReplicas(cmd) {
        err = s.pick(cmd.Slot()).DoWithReader(ctx, cmd, fn)
    } else {
        err = s.primary.Load().DoWithReader(ctx, cmd, fn)
    }

    if s.enableRedirect {
        // handleRedirect returns (redirectError, wasRedirect)
        // redirectError is error from redirect operation itself
        // wasRedirect indicates if this was a REDIRECT error
        if redirectErr, ok := s.handleRedirect(ctx, err); ok {
            // If redirect succeeded (redirectErr == nil), retry
            // OR if retry handler says to retry
            if redirectErr == nil || s.retryer.WaitOrSkipRetry(ctx, attempts, cmd, err) {
                attempts++
                goto retry
            }
        }
        if err == nil {
            cmds.PutCompletedForce(cmd)
        }
    }

    return err
}
```

### 8. `resp.go` - Export Helper Functions

Export parsing helpers for applications to use in callbacks:

```go
// ReadInt reads a RESP integer value from the reader.
// Use after receiving type byte ':' or for reading lengths after '$', '*', etc.
func ReadInt(r *bufio.Reader) (int64, error) {
    return readI(r)
}

// ReadBlobBytes reads a RESP blob string and returns the raw bytes.
// Call after receiving type byte '$'. Allocates a new []byte.
func ReadBlobBytes(r *bufio.Reader) ([]byte, error) {
    ptr, length, err := readB(r)
    if err != nil {
        return nil, err
    }
    bs := make([]byte, length)
    copy(bs, unsafe.Slice(ptr, length))
    return bs, nil
}

// ReadSimpleString reads a RESP simple string.
// Call after receiving type byte '+'.
// Returns a string that references the reader's internal buffer (zero-copy).
// The string is only valid until the next read operation.
func ReadSimpleString(r *bufio.Reader) (string, error) {
    ptr, length, err := readS(r)
    if err != nil {
        return "", err
    }
    return unsafe.String(ptr, length), nil
}

// DiscardCRLF discards the trailing \r\n after a simple value.
func DiscardCRLF(r *bufio.Reader) error {
    _, err := r.Discard(2)
    return err
}
```

### 9. `rueidishook/hook.go` - Hook Wrapper (Optional)

**Add to Hook interface:**
```go
DoWithReader(client rueidis.Client, ctx context.Context, cmd rueidis.Completed, fn rueidis.ReaderFunc) error
```

**Add implementation in hookclient and dedicated types**

### 10. `rueidisotel/trace.go` - OpenTelemetry Wrapper (Optional)

**Add tracing implementation:**
```go
func (o *otelclient) DoWithReader(ctx context.Context, cmd rueidis.Completed, fn rueidis.ReaderFunc) error {
    op := first(cmd.Commands())
    defer o.recordDuration(ctx, op, time.Now())
    ctx, span := o.start(ctx, op, sum(cmd.Commands()))
    if o.dbStmtFunc != nil {
        span.SetAttributes(dbstmt.String(o.dbStmtFunc(cmd.Commands())))
    }

    err := o.client.DoWithReader(ctx, cmd, fn)
    o.end(span, err)
    o.recordError(ctx, op, err)
    return err
}
```

## Corner Cases Checklist

### Cluster Mode
- [x] MOVED redirect - Retry on new node
- [x] ASK redirect - Send ASKING first, then retry on new node
- [x] TRYAGAIN error - Wait and retry (if retryable)
- [x] CLUSTERDOWN error - Wait and retry (if retryable)
- [x] LOADING error - Wait and retry (if retryable)
- [x] errConnExpired - Retry immediately
- [x] Max redirect limit - Stop after N redirects
- [x] Slot-based connection selection
- [x] Replica preference handling

### Sentinel Mode
- [x] errConnExpired - Retry immediately
- [x] Retry on transient errors
- [x] Node selection (master vs replica)

### Standalone Mode
- [x] REDIRECT (replica to primary) - Retry on primary
- [x] errConnExpired - Retry immediately
- [x] Replica selection

### Single Client
- [x] errConnExpired - Retry immediately
- [x] Retry on transient errors

### Wire Level
- [x] Context cancellation - Check ctx.Err() before operations
- [x] Timeout handling - Set connection deadline
- [x] Push messages - Skip RESP3 push notifications (type `>`)
- [x] NULL responses - Return Nil error without calling callback
- [x] Simple errors (`-`) - Parse as RedisError
- [x] Blob errors (`!`) - Parse as RedisError
- [x] Command recycling - Call cmds.PutCompleted() on success

### Error Response Handling
- [x] User callback NOT called for error responses
- [x] User callback NOT called for NULL responses
- [x] Error returned as RedisError so client layer can check IsMoved(), IsAsk(), etc.

## Testing Plan

### Unit Tests

1. **pipe_test.go**
   - Test successful response parsing callback
   - Test error response returns RedisError
   - Test NULL response returns Nil
   - Test push message skipping
   - Test context cancellation
   - Test timeout handling

2. **cluster_test.go**
   - Test MOVED redirect handling
   - Test ASK redirect handling (with ASKING command)
   - Test TRYAGAIN retry
   - Test CLUSTERDOWN retry
   - Test max redirect limit
   - Test errConnExpired recovery

3. **sentinel_test.go**
   - Test retry logic
   - Test errConnExpired recovery

4. **standalone_test.go**
   - Test REDIRECT handling
   - Test replica selection

5. **client_test.go**
   - Test single client retry logic
   - Test dedicated client

### Integration Tests

1. Test with real Redis cluster
2. Test redirect scenarios with cluster resharding
3. Test with network partitions
4. Benchmark allocation comparison vs Do()

## API Usage Examples

### Basic String Parsing
```go
var result string
err := client.DoWithReader(ctx, client.B().Get().Key("mykey").Build(),
    func(r *bufio.Reader, typ byte) error {
        if typ == '$' {
            length, err := rueidis.ReadInt(r)
            if err != nil {
                return err
            }
            buf := make([]byte, length)
            io.ReadFull(r, buf)
            r.Discard(2) // \r\n
            result = string(buf)
        }
        return nil
    })
```

### Zero-Copy Streaming to Writer
```go
err := client.DoWithReader(ctx, client.B().Get().Key("large-value").Build(),
    func(r *bufio.Reader, typ byte) error {
        if typ == '$' {
            length, _ := rueidis.ReadInt(r)
            io.CopyN(outputWriter, r, length) // Zero-copy streaming!
            r.Discard(2)
        }
        return nil
    })
```

### Array Parsing
```go
var items []string
err := client.DoWithReader(ctx, client.B().Lrange().Key("list").Start(0).Stop(-1).Build(),
    func(r *bufio.Reader, typ byte) error {
        if typ != '*' {
            return fmt.Errorf("expected array, got %c", typ)
        }
        count, _ := rueidis.ReadInt(r)
        items = make([]string, 0, count)

        for i := int64(0); i < count; i++ {
            elemTyp, _ := r.ReadByte()
            if elemTyp == '$' {
                length, _ := rueidis.ReadInt(r)
                buf := make([]byte, length)
                io.ReadFull(r, buf)
                r.Discard(2)
                items = append(items, string(buf))
            }
        }
        return nil
    })
```

## Implementation Order

1. **Phase 1: Core Implementation**
   - [ ] Add ReaderFunc type to rueidis.go
   - [ ] Add DoWithReader to Client interface
   - [ ] Implement pipe.DoWithReader (wire level)
   - [ ] Export helper functions in resp.go
   - [ ] Add tests for pipe.DoWithReader

2. **Phase 2: Client Implementations**
   - [ ] Implement mux.DoWithReader
   - [ ] Implement singleClient.DoWithReader
   - [ ] Implement clusterClient.DoWithReader
   - [ ] Implement sentinelClient.DoWithReader
   - [ ] Implement standalone.DoWithReader
   - [ ] Add tests for each client type

3. **Phase 3: Dedicated Clients**
   - [ ] Implement dedicatedSingleClient.DoWithReader
   - [ ] Implement dedicatedClusterClient.DoWithReader
   - [ ] Add tests

4. **Phase 4: Wrappers (Optional)**
   - [ ] Add to rueidishook.Hook interface
   - [ ] Implement in rueidishook
   - [ ] Implement in rueidisotel

5. **Phase 5: Documentation & Benchmarks**
   - [ ] Add documentation
   - [ ] Add benchmark comparing DoWithReader vs Do
   - [ ] Add examples

## Mock Implementations (for tests)

The following mock types need DoWithReader implementations:

### `mux_test.go` - mockWire
```go
func (m *mockWire) DoWithReader(ctx context.Context, pool *pool, cmd Completed, fn ReaderFunc) error {
    // Mock implementation for testing
    return nil
}
```

### `client_test.go` - mockConn
```go
func (m *mockConn) DoWithReader(ctx context.Context, cmd Completed, fn ReaderFunc) error {
    // Mock implementation for testing
    return nil
}
```

## Important Implementation Notes

### Callback Error Handling

**CRITICAL IMPLEMENTATION DETAIL:**

If the user callback returns an error, the connection state may be corrupted because the callback failed to fully consume the response. To prevent subsequent commands from reading corrupted data:

```go
// In pipe.DoWithReader:
err = fn(p.r, typ)

// If callback returned an error, close the connection
if err != nil {
    p.error.CompareAndSwap(nil, &errs{error: err})
    p.conn.Close()
    p.background()
}
```

**User Responsibility:**
- The callback MUST check ALL errors from io operations (io.ReadFull, r.Discard, etc.)
- Ignoring errors will corrupt the connection state
- Returning an error will close the connection (safe but costly)
- This is documented in both ReaderFunc type and DoWithReader method

### RESP2 vs RESP3 Compatibility

- RESP2 NULL: `$-1\r\n` (blob string with length -1)
- RESP3 NULL: `_\r\n` (explicit null type)
- Both should be handled and return `Nil` error

```go
// Handle RESP2 null in readB() - returns errOldNull
// Handle RESP3 null with type '_'
```

### Connection State After Callback Error

If the user callback returns an error but hasn't fully consumed the response:
- The connection state may be corrupted
- Consider closing the connection to be safe
- OR document that user MUST fully consume the response

Recommendation: Document that user must fully consume the response, and trust the user to do so (same assumption as DoStream).

### Blocking Commands

`DoWithReader` should work with blocking commands (BLPOP, BRPOP, etc.):
- Uses the dedicated blocking pool (dpool) if command IsBlock()
- Same behavior as DoStream

Actually, looking at how DoStream works - it always uses spool (streaming pool), not dpool. Need to verify if DoWithReader should use dpool for blocking commands or follow DoStream's pattern.

## Open Questions

1. **Should DoWithReader be added to DedicatedClient interface?**
   - Currently DoStream is NOT on DedicatedClient
   - Recommendation: Add it, since DoWithReader is more generally useful

2. **Should we add DoMultiWithReader?**
   - More complex to implement correctly
   - Recommendation: Start with single command, add multi later if needed

3. **Should DoWithReader use blocking pool for blocking commands?**
   - DoStream uses spool (always)
   - Do uses dpool for blocking commands
   - Need to decide which pattern to follow
