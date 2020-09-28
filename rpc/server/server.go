package main

import (
	"flag"
	"github.com/hslam/rpc"
	"github.com/hslam/rpc-benchmark/service"
)

var network string
var addr string
var codec string
var poll bool

func init() {
	flag.StringVar(&network, "network", "tcp", "-network=tcp")
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
	flag.StringVar(&codec, "codec", "pb", "-codec=code")
	flag.BoolVar(&poll, "poll", false, "-poll=false")
	flag.Parse()
}

func main() {
	rpc.Register(new(service.Arith))
	rpc.SetPoll(poll)
	rpc.Listen(network, addr, codec)
}
