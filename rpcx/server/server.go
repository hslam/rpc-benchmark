package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"github.com/hslam/rpc-benchmark/rpcx/service"
)

var network string
var addr string

func init() {
	flag.StringVar(&network, "network", "tcp", "-network=tcp")
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
	flag.Parse()
}

func main() {
	server := server.NewServer()
	server.Register(new(service.Arith), "")
	server.Serve("tcp", addr)
}
