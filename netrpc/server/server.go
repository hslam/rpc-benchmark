package main

import (
	"flag"
	"github.com/hslam/rpc-benchmark/service"
	"log"
	"net"
	"net/rpc"
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
	rpc.Register(new(service.Arith))
	lis, err := net.Listen(network, addr)
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
