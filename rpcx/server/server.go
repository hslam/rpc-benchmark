package main

import (
	"flag"
	"github.com/hslam/rpc-benchmark/rpcx/service"
	"github.com/smallnest/rpcx/server"
)

var network string
var addr string
var codec string

func init() {
	flag.StringVar(&network, "network", "tcp", "-network=tcp")
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
	flag.StringVar(&codec, "codec", "pb", "-codec=code")
	flag.Parse()
}

func main() {
	server := server.NewServer()
	server.Register(new(service.Arith), "")
	server.Serve(network, addr)
}
