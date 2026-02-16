# Experimental RDMA Connection

This is the experimental support for the [Valkey RDMA connection type](https://valkey.io/topics/RDMA/).

⚠️ **Known Issue**
RDMA connections may hang when handling **large payloads** or operating under **high concurrency**.
This feature is still under active development, and **contributions are very welcome**.

## Benchmark

The following benchmark compares end-to-end TCP vs. RDMA performance:

```sh
$ go test -run=NONE -bench=. -benchtime=2s
goos: linux
goarch: arm64
pkg: github.com/redis/rueidis/rueidisrdma
BenchmarkE2E/TCP-4                 10000            226352 ns/op              32 B/op          2 allocs/op
BenchmarkE2E/RDMA-4                25057             97860 ns/op              32 B/op          2 allocs/op
PASS
ok      github.com/redis/rueidis/rueidisrdma       5.759s
```

## Building and Running Valkey with the RDMA Module

Building Valkey with RDMA support is straightforward:

```sh
sudo apt install -y build-essential pkg-config libjemalloc-dev
sudo apt install -y librdmacm-dev rdma-core rdmacm-utils
wget https://github.com/valkey-io/valkey/archive/refs/tags/9.0.1.tar.gz
tar -zxvf 9.0.1.tar.gz && rm 9.0.1.tar.gz
cd valkey-9.0.1

BUILD_RDMA=module make
```

Running it is a bit more complex. Here is an example of how you can test it with one client and server virtual machines.

First, make sure that RDMA works between BOTH machines. Here we use RoCEv2 as an example:

Setup RoCEv2 on both machines:
```sh
sudo rdma link add rxe0 type rxe netdev <eth0 or enp2s0>
```

Test the setup with `rping` first. On the server:
```sh
rping -s -a <server ip on the rdma netdev> -v
```

Then on the client machine:
```sh
rping -c -a <server ip on the rdma netdev> -v
```

If you see data is transmitting, then you are ready for starting Valkey with RDMA. On the server:

```sh
./src/valkey-server --loadmodule ./src/valkey-rdma.so --rdma-bind <server ip on the rdma netdev> --rdma-port 6378
```

Test it with valkey-cli. On the client machine:

```sh
./src/valkey-cli -h <server ip on the rdma netdev> -p 6378 --rdma hello 3
```

## Connecting with the Go Client

```go
package main

import (
	"context"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisrdma"
)

func main() {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"<server ip on the rdma netdev>:6378"},
		DialCtxFn:   rueidisrdma.DialCtxFn,
	})
	...
}
```