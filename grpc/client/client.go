package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"github.com/hslam/rpc-benchmark/grpc/service"
	"github.com/hslam/stats"
	"log"
	"math/rand"
)

var network string
var addr string
var clients int
var total int
var parallel int
var bar bool

func init() {
	flag.StringVar(&network, "network", "tcp", "-network=tcp")
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
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
		if conn, err := grpc.Dial(addr, grpc.WithInsecure()); err != nil {
			log.Fatalln("dailing error: ", err)
		} else {
			wrkClients = append(wrkClients, &WrkClient{conn})
		}
	}
	stats.StartPrint(parallel, total, wrkClients)
}

type WrkClient struct {
	*grpc.ClientConn
}

func (c *WrkClient) Call() (int64, int64, bool) {
	A := rand.Int31n(100)
	B := rand.Int31n(100)
	req := &service.ArithRequest{A: A, B: B}
	var res *service.ArithResponse
	client := service.NewArithServiceClient(c.ClientConn)
	res, err := client.Multiply(context.Background(), req)
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
