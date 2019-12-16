package main

import (
	"log"
	"net"
	"github.com/reminance/reminance-go/xinliangnote-Go/02-Go gRPC/codes/01-gRPC Hello World/go_server/proto/hello"
	"github.com/reminance/reminance-go/xinliangnote-Go/02-Go gRPC/codes/01-gRPC Hello World/go_server/controller/hello_controller"
	"google.golang.org/grpc"
)

const (
	Address = "0.0.0.0:9090"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// 服务注册
	hello.RegisterHelloServer(s, &hello_controller.HelloController{})

	log.Println("Listen on " + Address)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
