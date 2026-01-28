package rueidisrdma

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	randv2 "math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/redis/rueidis"
)

func BenchmarkE2E(b *testing.B) {
	f := func(o rueidis.ClientOption) func(b *testing.B) {
		return func(b *testing.B) {
			c, err := rueidis.NewClient(o)
			if err != nil {
				b.Fatal(err)
			}
			defer c.Close()
			b.ResetTimer()
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					i := randv2.Uint64N(100000000000)
					s := strconv.FormatUint(i, 10)
					r, err := c.Do(context.Background(), c.B().Echo().Message(s).Build()).ToString()
					if err != nil || r != s {
						b.Fatal(err)
					}
				}
			})
			b.StopTimer()
		}
	}
	b.Run("TCP", f(rueidis.ClientOption{
		InitAddress: []string{"172.16.255.128:6379"},
	}))
	b.Run("RDMA", f(rueidis.ClientOption{
		InitAddress: []string{"172.16.255.128:6378"},
		DialCtxFn:   DialCtxFn,
	}))
}

func TestLocalRoCECM(t *testing.T) {
	if os.Getenv("CI") == "true" {
		t.Skip()
	}
	needCmd(t, "ip")
	needCmd(t, "rdma")
	needSudo(t)

	addr := "10.200.2.1"
	port := "6378"

	ctx, cancel := context.WithTimeout(t.Context(), 300*time.Second)
	defer cancel()

	defer startValkey901WithRdmaModule(t, ctx, addr, port)()

	c, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{addr + ":" + port},
		DialCtxFn:   DialCtxFn,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			buf := make([]byte, 512)
			for j := 1; j <= len(buf); j++ {
				_, err := rand.Read(buf[:j])
				if err != nil {
					t.Error(err)
					return
				}
				v, err := c.Do(ctx, c.B().Echo().Message(rueidis.BinaryString(buf[:j])).Build()).ToString()
				if err != nil {
					t.Error(err)
					return
				}
				if v != rueidis.BinaryString(buf[:j]) {
					t.Errorf("got %q, want %q", v, string(buf[:j]))
					return
				}
			}
		}()
	}
	wg.Wait()
}

type writer struct {
	fn func(p []byte) (n int, err error)
}

func (i *writer) Write(p []byte) (n int, err error) {
	return i.fn(p)
}

func startValkey901WithRdmaModule(t testing.TB, ctx context.Context, addr, port string) func() {
	t.Helper()

	needCmd(t, "ip")
	needCmd(t, "rdma")
	needSudo(t)

	netdev := "dummy0"
	rxedev := "rxe_dummy0"

	if err := sudoRun(ctx, "modprobe", "rdma_rxe"); err != nil {
		t.Logf("failed to load rdma_rxe: %v", err)
	}
	if err := sudoRun(ctx, "ip", "link", "add", netdev, "type", "dummy"); err != nil {
		t.Logf("failed to add dummy link: %v", err)
	}
	if err := sudoRun(ctx, "ip", "addr", "add", addr+"/24", "dev", netdev); err != nil {
		t.Logf("failed to add dummy addr: %v", err)
	}
	if err := sudoRun(ctx, "ip", "link", "set", netdev, "up"); err != nil {
		t.Logf("failed to set dummy up: %v", err)
	}
	if err := sudoRun(ctx, "rdma", "link", "add", rxedev, "type", "rxe", "netdev", netdev); err != nil {
		t.Logf("failed to add dummy rxe: %v", err)
	}

	valkeybin, rdmaso := buildValkey901WithRdmaModule(t, ctx)

	r := make(chan struct{}, 1)
	b := bytes.NewBuffer(nil)
	w := &writer{fn: func(p []byte) (n int, err error) {
		if len(r) == 0 {
			b.Write(p)
			if bytes.Contains(b.Bytes(), []byte("Ready to accept connections rdma")) {
				r <- struct{}{}
			}
		}
		return os.Stdout.Write(p)
	}}
	srv := exec.CommandContext(ctx, valkeybin, "--loadmodule", rdmaso, "--rdma-bind", addr, "--rdma-port", port)
	srv.Stderr = os.Stderr
	srv.Stdout = w

	if err := srv.Start(); err != nil {
		t.Fatalf("failed to start server: %v", err)
	}

	select {
	case <-r:
	case <-time.After(30 * time.Second):
		_ = srv.Process.Kill()
		_ = srv.Wait()
		t.Fatal("timeout waiting for server to start")
	}
	return func() {
		_ = srv.Process.Kill()
		_ = srv.Wait()
	}
}

func buildValkey901WithRdmaModule(t testing.TB, ctx context.Context) (valkeyServer, rdmaSo string) {
	t.Helper()

	const (
		tarURL    = "https://github.com/valkey-io/valkey/archive/refs/tags/9.0.1.tar.gz"
		tarName   = "9.0.1.tar.gz"
		srcDir    = "valkey-9.0.1"
		serverRel = "src/valkey-server"
		soRel     = "src/valkey-rdma.so"
	)

	cacheRoot := valkeyCacheDir(t)
	t.Logf("using valkey cache dir: %s", cacheRoot)

	tarPath := filepath.Join(cacheRoot, tarName)
	srcPath := filepath.Join(cacheRoot, srcDir)

	valkeyServer = filepath.Join(srcPath, serverRel)
	rdmaSo = filepath.Join(srcPath, soRel)

	// Fast path: already built
	if fileExists(valkeyServer) && fileExists(rdmaSo) {
		return valkeyServer, rdmaSo
	}

	_ = os.MkdirAll(cacheRoot, 0o755)

	// 1) wget https://github.com/.../9.0.1.tar.gz
	if !fileExists(tarPath) && !dirExists(srcPath) {
		mustRun(t, ctx, cacheRoot, "wget", "-q", "-O", tarName, tarURL)
	}

	// 2) tar -zxvf 9.0.1.tar.gz && rm 9.0.1.tar.gz
	if !dirExists(srcPath) {
		mustRun(t, ctx, cacheRoot, "tar", "-zxvf", tarName)
		_ = os.Remove(tarPath)
	}

	// 3) BUILD_RDMA=module make
	if !fileExists(valkeyServer) || !fileExists(rdmaSo) {
		mustRunEnv(t, ctx, srcPath, []string{"BUILD_RDMA=module"}, "make")
	}

	return valkeyServer, rdmaSo
}

func valkeyCacheDir(t testing.TB) string {
	t.Helper()
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("cannot determine home dir: %v", err)
	}
	return filepath.Join(home, ".cache", "valkey")
}

func fileExists(p string) bool {
	st, err := os.Stat(p)
	return err == nil && st.Mode().IsRegular()
}

func dirExists(p string) bool {
	st, err := os.Stat(p)
	return err == nil && st.IsDir()
}

func mustRun(t testing.TB, ctx context.Context, dir, name string, args ...string) {
	t.Helper()
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("run %s %v (dir=%s): %v\n%s", name, args, dir, err, string(out))
	}
}

func mustRunEnv(t testing.TB, ctx context.Context, dir string, env []string, name string, args ...string) {
	t.Helper()
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	// Inherit environment + overrides.
	cmd.Env = append(os.Environ(), env...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("run %s %v env=%v (dir=%s): %v\n%s", name, args, env, dir, err, string(out))
	}
}

func needCmd(t testing.TB, name ...string) {
	t.Helper()
	for _, n := range name {
		if !hasCmd(n) {
			t.Skipf("missing required command %q in PATH", n)
		}
	}
}

func hasCmd(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func needSudo(t testing.TB) {
	t.Helper()
	if !hasCmd("sudo") {
		t.Skip("sudo not found in PATH")
	}
	// Require non-interactive sudo so tests don't hang.
	cmd := exec.Command("sudo", "-n", "true")
	if err := cmd.Run(); err != nil {
		t.Skip("sudo not available without password (run sudo once or configure NOPASSWD)")
	}
}

func sudoRun(ctx context.Context, args ...string) error {
	_, err := sudoOutput(ctx, args...)
	return err
}

func sudoOutput(ctx context.Context, args ...string) (string, error) {
	all := append([]string{"sudo", "-n"}, args...)
	return output(ctx, all...)
}

func output(ctx context.Context, args ...string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("no command")
	}
	out, err := exec.CommandContext(ctx, args[0], args[1:]...).CombinedOutput()
	return string(out), err
}
