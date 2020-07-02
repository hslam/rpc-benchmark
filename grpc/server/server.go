package main

import (
	"flag"
	"google.golang.org/grpc"
	"github.com/hslam/rpc-benchmark/grpc/service"
	"log"
	"net"
)

var network string
var addr string

func init() {
	flag.StringVar(&network, "network", "tcp", "-network=tcp")
	flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
	flag.Parse()
}

func main() {
	grpcServer := grpc.NewServer()
	service.RegisterArithServiceServer(grpcServer, new(service.Arith))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
