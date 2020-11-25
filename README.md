# rpc-benchmark

**Comparison to other packages**

|Package| [netrpc](https://github.com/golang/go/tree/master/src/net/rpc "netrpc")| [jsonrpc](https://github.com/golang/go/tree/master/src/net/rpc/jsonrpc "jsonrpc")|[rpc](https://github.com/hslam/rpc "rpc")|[grpc](https://google.golang.org/grpc "grpc")|[rpcx](https://github.com/smallnest/rpcx "rpcx")|
|:--:|:--|:--|:--|:--|:--|
|Epoll/Kqueue|No|No|Yes|No|No|
|Multiplexing|Yes|Yes|Yes|Yes|Yes|
|Pipelining|No|No|Yes|No|No|
|Auto Batching|No|No|Yes|No|No|
|Client|Yes|Yes|Yes|Yes|Yes|
|Transport|No|No|Yes|No|No|
|Server Push|No|No|Yes|Yes|Yes|

##### Low Concurrency

<img src="https://raw.githubusercontent.com/hslam/rpc-benchmark/master/rpc-bar-qps.png" width = "400" height = "300" alt="rpc" align=center><img src="https://raw.githubusercontent.com/hslam/rpc-benchmark/master/rpc-bar-p99.png" width = "400" height = "300" alt="rpc" align=center>

##### High Concurrency

<img src="https://raw.githubusercontent.com/hslam/rpc-benchmark/master/rpc-curve-qps.png" width = "400" height = "300" alt="rpc" align=center><img src="https://raw.githubusercontent.com/hslam/rpc-benchmark/master/rpc-curve-p99.png" width = "400" height = "300" alt="rpc" align=center>
