package main

import (
	"flag"
	"github.com/hslam/rpc-benchmark/grpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
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
	grpcServer := grpc.NewServer()
	service.RegisterArithServiceServer(grpcServer, new(service.Arith))
	lis, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
