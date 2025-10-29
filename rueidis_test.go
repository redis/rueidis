package rueidis

import (
	"bufio"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gleak"
)

func SetupLeakDetection() (gomega.Gomega, []gleak.Goroutine) {
	return gomega.NewGomega(func(message string, callerSkip ...int) {
		panic(message)
	}), gleak.Goroutines()
}

func ShouldNotLeak(g gomega.Gomega, snapshot []gleak.Goroutine) {
	g.Eventually(gleak.Goroutines).WithTimeout(time.Minute).ShouldNot(gleak.HaveLeaked(snapshot))
}

func TestMain(m *testing.M) {
	g, snap := SetupLeakDetection()
	code := m.Run()
	ShouldNotLeak(g, snap)
	os.Exit(code)
}

func accept(t *testing.T, ln net.Listener) (*redisMock, error) {
	conn, err := ln.Accept()
	if err != nil {
		t.Error(err)
		return nil, err
	}
	mock := &redisMock{
		t:    t,
		buf:  bufio.NewReader(conn),
		conn: conn,
	}
	mock.Expect("HELLO", "3").
		Reply(slicemsg(
			'%',
			[]RedisMessage{
				strmsg('+', "proto"),
				{typ: ':', intlen: 3},
			},
		))
	mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
		ReplyString("OK")
	return mock, nil
}

func TestNewClusterClient(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	done := make(chan struct{})
	go func() {
		mock, err := accept(t, ln)
		if err != nil {
			return
		}
		slots, _ := slotsResp.ToMessage()
		mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLUSTER", "SLOTS").Reply(slots)
		mock.Close()
		close(done)
	}()

	_, port, _ := net.SplitHostPort(ln.Addr().String())
	client, err := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:" + port},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	if _, ok := client.(*clusterClient); !ok {
		t.Fatal("client should be a clusterClient")
	}
	<-done
}

func TestNewClusterClientError(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	t.Run("cluster slots command error", func(t *testing.T) {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()
		done := make(chan struct{})
		go func() {
			mock, err := accept(t, ln)
			if err != nil {
				return
			}
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLUSTER", "SLOTS").Reply(strmsg('-', "other error"))
			mock.Expect("PING").ReplyString("OK")
			mock.Close()
			close(done)
		}()
		_, port, _ := net.SplitHostPort(ln.Addr().String())
		client, err := NewClient(ClientOption{
			InitAddress: []string{"127.0.0.1:" + port},
		})
		if client != nil || err == nil {
			t.Errorf("unexpected return %v %v", client, err)
		}
		<-done
	})

	t.Run("replica only and send to replicas option conflict", func(t *testing.T) {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, port, _ := net.SplitHostPort(ln.Addr().String())
		client, err := NewClient(ClientOption{
			InitAddress: []string{"127.0.0.1:" + port},
			ReplicaOnly: true,
			SendToReplicas: func(cmd Completed) bool {
				return true
			},
		})
		if client != nil || err == nil {
			t.Errorf("unexpected return %v %v", client, err)
		}

		if !strings.Contains(err.Error(), ErrReplicaOnlyConflict.Error()) {
			t.Errorf("unexpected error %v", err)
		}
	})

	t.Run("replica only and replica selector option conflict", func(t *testing.T) {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, port, _ := net.SplitHostPort(ln.Addr().String())
		client, err := NewClient(ClientOption{
			InitAddress: []string{"127.0.0.1:" + port},
			ReplicaOnly: true,
			ReplicaSelector: func(slot uint16, replicas []NodeInfo) int {
				return 0
			},
		})
		if client != nil || err == nil {
			t.Errorf("unexpected return %v %v", client, err)
		}

		if !strings.Contains(err.Error(), ErrReplicaOnlyConflictWithReplicaSelector.Error()) {
			t.Errorf("unexpected error %v", err)
		}
	})

	t.Run("send to replicas should be set when replica selector is set", func(t *testing.T) {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, port, _ := net.SplitHostPort(ln.Addr().String())
		client, err := NewClient(ClientOption{
			InitAddress: []string{"127.0.0.1:" + port},
			ReplicaSelector: func(slot uint16, replicas []NodeInfo) int {
				return 0
			},
		})
		if client != nil || err == nil {
			t.Errorf("unexpected return %v %v", client, err)
		}

		if !strings.Contains(err.Error(), ErrSendToReplicasNotSet.Error()) {
			t.Errorf("unexpected error %v", err)
		}
	})

	t.Run("replica only and read node selector option conflict", func(t *testing.T) {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, port, _ := net.SplitHostPort(ln.Addr().String())
		client, err := NewClient(ClientOption{
			InitAddress: []string{"127.0.0.1:" + port},
			ReplicaOnly: true,
			ReadNodeSelector: func(slot uint16, nodes []NodeInfo) int {
				return 0
			},
		})
		if client != nil || err == nil {
			t.Errorf("unexpected return %v %v", client, err)
		}

		if !strings.Contains(err.Error(), ErrReplicaOnlyConflictWithReadNodeSelector.Error()) {
			t.Errorf("unexpected error %v", err)
		}
	})

	t.Run("replica selector and read node selector option conflict", func(t *testing.T) {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, port, _ := net.SplitHostPort(ln.Addr().String())
		client, err := NewClient(ClientOption{
			InitAddress: []string{"127.0.0.1:" + port},
			ReplicaSelector: func(slot uint16, replicas []NodeInfo) int {
				return 0
			},
			ReadNodeSelector: func(slot uint16, nodes []NodeInfo) int {
				return 0
			},
		})
		if client != nil || err == nil {
			t.Errorf("unexpected return %v %v", client, err)
		}

		if !strings.Contains(err.Error(), ErrReplicaSelectorConflictWithReadNodeSelector.Error()) {
			t.Errorf("unexpected error %v", err)
		}
	})
}

func TestFallBackSingleClient(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	done := make(chan struct{})
	go func() {
		mock, err := accept(t, ln)
		if err != nil {
			return
		}
		mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLUSTER", "SLOTS").Reply(strmsg('-', "ERR This instance has cluster support disabled"))
		mock.Expect("PING").ReplyString("OK")
		mock.Close()
		close(done)
	}()

	_, port, _ := net.SplitHostPort(ln.Addr().String())
	client, err := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:" + port},
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := client.(*singleClient); !ok {
		t.Fatal("client should be a singleClient")
	}
	client.Close()
	<-done
}

func TestForceSingleClientInitialDialError(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	_, port, _ := net.SplitHostPort(ln.Addr().String())
	initialDialErr := errors.New("initial Dial error")
	var dials atomic.Uint32
	client, err := NewClient(ClientOption{
		InitAddress:       []string{"127.0.0.1:" + port},
		ForceSingleClient: true,
		DialCtxFn: func(ctx context.Context, addr string, dialer *net.Dialer, _ *tls.Config) (net.Conn, error) {
			if dials.Add(1) == 1 {
				return nil, initialDialErr
			}
			return dialer.DialContext(ctx, "tcp", addr)
		},
	})
	if client == nil {
		t.Fatal("NewClient returned a nil client with ForceSingleClient")
	}
	if !errors.Is(err, initialDialErr) {
		t.Fatalf("unexpected NewClient error: %v", err)
	}

	done := make(chan struct{})
	go func() {
		mock, err := accept(t, ln)
		if err != nil {
			return
		}
		mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("PING").ReplyString("OK")
		mock.Close()
		close(done)
	}()

	val, err := client.Do(context.Background(), client.B().Ping().Build()).ToMessage()
	if err != nil {
		t.Errorf("unexpected error result: %v", err)
	}
	if str, _ := val.ToString(); str != "OK" {
		t.Errorf("unexpected result: %v", str)
	}
	client.Close()
	<-done

	if n := dials.Load(); n != 2 {
		t.Errorf("expected 2 Dial calls, got %d", n)
	}
}

func TestForceSingleClient(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	done := make(chan struct{})
	go func() {
		mock, err := accept(t, ln)
		if err != nil {
			return
		}
		mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("PING").ReplyString("OK")
		mock.Close()
		close(done)
	}()
	_, port, _ := net.SplitHostPort(ln.Addr().String())
	client, err := NewClient(ClientOption{
		InitAddress:       []string{"127.0.0.1:" + port},
		ForceSingleClient: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := client.(*singleClient); !ok {
		t.Fatal("client should be a singleClient")
	}
	client.Close()
	<-done
}

func TestStandaloneClientWithNoSendToReplicas(t *testing.T) {
	_, err := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
		Standalone: StandaloneOption{
			ReplicaAddress: []string{"127.0.0.1:6378"},
		},
	})
	if err != ErrNoSendToReplicas {
		t.Fatal(err)
	}
}

func TestStandaloneClient(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	pln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer pln.Close()
	rln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer rln.Close()

	var wg sync.WaitGroup
	mockServer := func(ln net.Listener) {
		defer wg.Done()
		mock, err := accept(t, ln)
		if err != nil {
			return
		}
		mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("PING").ReplyString("OK")
		mock.Close()
	}
	wg.Add(2)
	go mockServer(pln)
	go mockServer(rln)

	_, pport, _ := net.SplitHostPort(pln.Addr().String())
	_, rport, _ := net.SplitHostPort(rln.Addr().String())
	client, err := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:" + pport},
		Standalone: StandaloneOption{
			ReplicaAddress: []string{"127.0.0.1:" + rport},
		},
		SendToReplicas: func(cmd Completed) bool {
			return cmd.IsReadOnly()
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := client.(*standalone); !ok {
		t.Fatal("client should be a standalone")
	}
	client.Close()
	wg.Wait()
}

func TestTLSClient(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	template := x509.Certificate{
		SerialNumber:          serialNumber,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost"},
		Subject:               pkix.Name{Organization: []string{"Acme Co"}},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, pub, priv)
	if err != nil {
		t.Fatalf("Failed to create certificate: %v", err)
	}

	certPem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		t.Fatalf("Unable to marshal private key: %v", err)
	}

	privPem := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})

	cert, err := tls.X509KeyPair(certPem, privPem)
	if err != nil {
		t.Fatalf("Fail to load X509KeyPair: %v", err)
	}

	config := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	ln, err := tls.Listen("tcp", "127.0.0.1:0", config)
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	done := make(chan struct{})
	go func() {
		mock, err := accept(t, ln)
		if err != nil {
			return
		}
		mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLUSTER", "SLOTS").Reply(strmsg('-', "ERR This instance has cluster support disabled"))
		mock.Expect("PING").ReplyString("OK")
		mock.Close()
		close(done)
	}()

	_, port, _ := net.SplitHostPort(ln.Addr().String())
	client, err := NewClient(ClientOption{
		InitAddress:       []string{"127.0.0.1:" + port},
		TLSConfig:         config,
		DisableTCPNoDelay: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := client.(*singleClient); !ok {
		t.Fatal("client should be a singleClient")
	}
	client.Close()
	<-done
}

func TestNewClientMaxMultiplex(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	_, err := NewClient(ClientOption{
		InitAddress:       []string{"127.0.0.1:6379"},
		PipelineMultiplex: MaxPipelineMultiplex + 1,
	})
	if err != ErrWrongPipelineMultiplex {
		t.Fatalf("unexpected error %v", err)
	}
}

func TestNewClientWithEnableRedirectPriority(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	// Test that EnableRedirect has priority over other options
	// We'll test by checking error conditions rather than actually connecting

	// First, test the error case to demonstrate the path
	_, err := NewClient(ClientOption{
		InitAddress: []string{}, // Empty address should cause an error
		Standalone: StandaloneOption{
			EnableRedirect: true,
		},
		ForceSingleClient: true, // This should be ignored when EnableRedirect is set
	})

	if err != ErrNoAddr {
		t.Errorf("expected ErrNoAddr, got %v", err)
	}
}

func TestNewClientWithMultipleReplicaAddresses(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	// Test that ReplicaAddress requires SendToReplicas
	_, err := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
		Standalone: StandaloneOption{
			ReplicaAddress: []string{"127.0.0.1:6380", "127.0.0.1:6381"},
		},
		// Missing SendToReplicas should cause an error
	})

	if err != ErrNoSendToReplicas {
		t.Errorf("expected ErrNoSendToReplicas, got %v", err)
	}
}

func TestNewClientEnableRedirectAndReplicaAddressConflict(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	// Test that EnableRedirect and ReplicaAddress cannot be used together
	_, err := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
		Standalone: StandaloneOption{
			EnableRedirect: true,
			ReplicaAddress: []string{"127.0.0.1:6380"},
		},
	})

	if err == nil {
		t.Error("expected error when EnableRedirect and ReplicaAddress are both used")
	}

	expectedMsg := "EnableRedirect and ReplicaAddress cannot be used together"
	if err.Error() != expectedMsg {
		t.Errorf("expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestSingleClientMultiplex(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	option := ClientOption{}
	if v := singleClientMultiplex(option.PipelineMultiplex); v != 2 {
		t.Fatalf("unexpected value %v", v)
	}
	option.PipelineMultiplex = -1
	if v := singleClientMultiplex(option.PipelineMultiplex); v != 0 {
		t.Fatalf("unexpected value %v", v)
	}
}

func TestCustomDialFnIsCalled(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	isFnCalled := false
	option := ClientOption{
		InitAddress: []string{"127.0.0.1:0"},
		DialFn: func(s string, dialer *net.Dialer, config *tls.Config) (conn net.Conn, err error) {
			isFnCalled = true
			return nil, errors.New("dial error")
		},
	}

	_, err := NewClient(option)

	if !isFnCalled {
		t.Fatalf("excepted ClientOption.DialFn to be called")
	}
	if err == nil {
		t.Fatalf("expected dial error")
	}
}

func TestCustomDialCtxFnIsCalled(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	isFnCalled := false
	option := ClientOption{
		InitAddress: []string{"127.0.0.1:0"},
		DialCtxFn: func(ctx context.Context, s string, dialer *net.Dialer, config *tls.Config) (conn net.Conn, err error) {
			isFnCalled = true
			return nil, errors.New("dial error")
		},
	}

	_, err := NewClient(option)

	if !isFnCalled {
		t.Fatalf("excepted ClientOption.DialFn to be called")
	}
	if err == nil {
		t.Fatalf("expected dial error")
	}
}

func ExampleIsRedisNil() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	_, err = client.Do(context.Background(), client.B().Get().Key("not_exists").Build()).ToString()
	if err != nil && IsRedisNil(err) {
		fmt.Printf("it is a nil response")
	}
}

func ExampleClient_do() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	client.Do(ctx, client.B().Set().Key("k").Value("1").Build()).Error()

	client.Do(ctx, client.B().Get().Key("k").Build()).ToString()

	client.Do(ctx, client.B().Get().Key("k").Build()).AsInt64()

	client.Do(ctx, client.B().Hmget().Key("h").Field("a", "b").Build()).ToArray()

	client.Do(ctx, client.B().Scard().Key("s").Build()).ToInt64()

	client.Do(ctx, client.B().Smembers().Key("s").Build()).AsStrSlice()
}

func ExampleClient_doCache() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	client.DoCache(ctx, client.B().Get().Key("k").Cache(), time.Minute).ToString()

	client.DoCache(ctx, client.B().Get().Key("k").Cache(), time.Minute).AsInt64()

	client.DoCache(ctx, client.B().Hmget().Key("h").Field("a", "b").Cache(), time.Minute).ToArray()

	client.DoCache(ctx, client.B().Scard().Key("s").Cache(), time.Minute).ToInt64()

	client.DoCache(ctx, client.B().Smembers().Key("s").Cache(), time.Minute).AsStrSlice()
}

func ExampleClient_scan() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for _, c := range client.Nodes() { // loop over all your redis nodes
		var scan ScanEntry
		for more := true; more; more = scan.Cursor != 0 {
			if scan, err = c.Do(context.Background(), c.B().Scan().Cursor(scan.Cursor).Build()).AsScanEntry(); err != nil {
				panic(err)
			}
			fmt.Println(scan.Elements)
		}
	}
}

func ExampleClient_dedicatedCAS() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	client.Dedicated(func(client DedicatedClient) error {
		// watch keys first
		if err := client.Do(ctx, client.B().Watch().Key("k1", "k2").Build()).Error(); err != nil {
			return err
		}
		// perform read here
		values, err := client.Do(ctx, client.B().Mget().Key("k1", "k2").Build()).ToArray()
		if err != nil {
			return err
		}
		v1, _ := values[0].ToString()
		v2, _ := values[1].ToString()
		// perform write with MULTI EXEC
		for _, resp := range client.DoMulti(
			ctx,
			client.B().Multi().Build(),
			client.B().Set().Key("k1").Value(v1+"1").Build(),
			client.B().Set().Key("k2").Value(v2+"2").Build(),
			client.B().Exec().Build(),
		) {
			if err := resp.Error(); err != nil {
				return err
			}
		}
		return nil
	})
}

func ExampleClient_dedicateCAS() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	c, cancel := client.Dedicate()
	defer cancel()

	ctx := context.Background()

	// watch keys first
	if err := c.Do(ctx, c.B().Watch().Key("k1", "k2").Build()).Error(); err != nil {
		panic(err)
	}
	// perform read here
	values, err := c.Do(ctx, c.B().Mget().Key("k1", "k2").Build()).ToArray()
	if err != nil {
		panic(err)
	}
	v1, _ := values[0].ToString()
	v2, _ := values[1].ToString()
	// perform write with MULTI EXEC
	for _, resp := range c.DoMulti(
		ctx,
		c.B().Multi().Build(),
		c.B().Set().Key("k1").Value(v1+"1").Build(),
		c.B().Set().Key("k2").Value(v2+"2").Build(),
		c.B().Exec().Build(),
	) {
		if err := resp.Error(); err != nil {
			panic(err)
		}
	}
}

func ExampleNewClient_cluster() {
	client, _ := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"},
		ShuffleInit: true,
	})
	defer client.Close()
}

func ExampleNewClient_single() {
	client, _ := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
	})
	defer client.Close()
}

func ExampleNewClient_sentinel() {
	client, _ := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"},
		Sentinel: SentinelOption{
			MasterSet: "my_master",
		},
	})
	defer client.Close()
}
