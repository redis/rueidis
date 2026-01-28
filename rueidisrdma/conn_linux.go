package rueidisrdma

/*
#cgo LDFLAGS: -libverbs -lrdmacm
#include <errno.h>
#include <stdlib.h>
#include "conn_linux.h"
int rdmaConnect(RdmaContext *ctx, const char *addr, int port, long timeout_msec);
ssize_t rdmaRead(RdmaContext *ctx, char *buf, size_t bufcap, long timeout_msec);
ssize_t rdmaWrite(RdmaContext *ctx, const char *obuf, size_t data_len, long timeout_msec);
void rdmaClose(RdmaContext *ctx);
void rdmaDisconnect(RdmaContext *ctx);
*/
import "C"

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var _ net.Conn = (*conn)(nil)

func DialCtxFn(ctx context.Context, dst string, _ *net.Dialer, _ *tls.Config) (net.Conn, error) {
	host, portstr, err := net.SplitHostPort(dst)
	if err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(portstr)
	if err != nil {
		return nil, err
	}
	c := &conn{
		ctx:   (*C.RdmaContext)(C.malloc(C.sizeof_struct_RdmaContext)),
		timed: -1,
	}
	chost := C.CString(host)
	defer C.free(unsafe.Pointer(chost))

	timeout := int64(10000)
	if dl, ok := ctx.Deadline(); ok {
		timeout = time.Until(dl).Milliseconds()
	}

	if ret := C.rdmaConnect(c.ctx, chost, C.int(port), C.long(timeout)); ret != 0 {
		defer C.free(unsafe.Pointer(c.ctx))
		return nil, c.err()
	}
	return c, nil
}

type conn struct {
	ctx   *C.RdmaContext
	mu    sync.RWMutex
	timed int64
	once  int32
}

func (c *conn) timeout() int64 {
	if c.timed < 0 {
		return 10000
	}
	return c.timed
}

func (c *conn) Read(b []byte) (n int, err error) {
	var ret C.ssize_t
	if len(b) != 0 {
		c.mu.RLock()
		if c.ctx != nil {
			ret = C.rdmaRead(c.ctx, (*C.char)(unsafe.Pointer(&b[0])), C.size_t(len(b)), C.long(c.timeout()))
		}
		c.mu.RUnlock()
		if ret <= 0 {
			return 0, c.err()
		}
	}
	return int(ret), nil
}

func (c *conn) Write(b []byte) (n int, err error) {
	var ret C.ssize_t
	if len(b) != 0 {
		c.mu.RLock()
		if c.ctx != nil {
			ret = C.rdmaWrite(c.ctx, (*C.char)(unsafe.Pointer(&b[0])), C.size_t(len(b)), C.long(c.timeout()))
		}
		c.mu.RUnlock()
		if ret <= 0 {
			return 0, c.err()
		}
	}
	return int(ret), nil
}

func (c *conn) Close() error {
	if atomic.CompareAndSwapInt32(&c.once, 0, 1) {
		C.rdmaDisconnect(c.ctx)
		c.mu.Lock()
		if c.ctx != nil {
			C.rdmaClose(c.ctx)
			C.free(unsafe.Pointer(c.ctx))
			c.ctx = nil
		}
		c.mu.Unlock()
	}
	return nil
}

func (c *conn) err() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.ctx == nil {
		return io.ErrClosedPipe
	}
	return fmt.Errorf("%s: %d", C.GoString(&c.ctx.errstr[0]), int(c.ctx.err))
}

func (c *conn) SetDeadline(t time.Time) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if t.IsZero() {
		c.timed = -1
	} else {
		if d := time.Until(t); d <= 0 {
			c.timed = 0 // Deadline already passed; use immediate timeout.
		} else {
			c.timed = d.Milliseconds()
		}
	}
	return nil
}

func (c *conn) SetReadDeadline(t time.Time) error {
	panic("not implemented")
}

func (c *conn) SetWriteDeadline(t time.Time) error {
	panic("not implemented")
}

func (c *conn) LocalAddr() net.Addr {
	panic("not implemented")
}

func (c *conn) RemoteAddr() net.Addr {
	panic("not implemented")
}
