package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/hslam/rpc-benchmark/rpcx/service"
	"github.com/hslam/stats"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"math/rand"
)

var network string
var addr string
var codec string
var clients int
var total int
var parallel int
var bar bool

func init() {
	flag.StringVar(&network, "network", "tcp", "-network=tcp")
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
	flag.StringVar(&codec, "codec", "pb", "-codec=code")
	flag.IntVar(&total, "total", 100000, "-total=100000")
	flag.IntVar(&parallel, "parallel", 1, "-parallel=1")
	flag.IntVar(&clients, "clients", 1, "-clients=1")
	flag.BoolVar(&bar, "bar", true, "-bar=true")
	flag.Parse()
	stats.SetBar(bar)
	fmt.Printf("./client -network=%s -addr=%s -total=%d -parallel=%d -clients=%d\n", network, addr, total, parallel, clients)
}

func main() {
	if clients < 1 || parallel < 1 || total < 1 {
		return
	}
	var wrkClients []stats.Client
	for i := 0; i < clients; i++ {
		d := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
		option := client.DefaultOption
		option.SerializeType = protocol.ProtoBuffer
		xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
		defer xclient.Close()
		wrkClients = append(wrkClients, &WrkClient{xclient})
	}
	stats.StartPrint(parallel, total, wrkClients)
}

type WrkClient struct {
	client.XClient
}

func (c *WrkClient) Call() (int64, int64, bool) {
	A := rand.Int31n(100)
	B := rand.Int31n(100)
	req := &service.ArithRequest{A: A, B: B}
	var res service.ArithResponse
	err := c.XClient.Call(context.Background(), "Multiply", req, &res)
	if err != nil {
		fmt.Println(err)
		return 0, 0, false
	}
	if res.Pro == A*B {
		return 0, 0, true
	}
	fmt.Printf("err %d * %d = %d\n", A, B, res.Pro)
	return 0, 0, false
}
