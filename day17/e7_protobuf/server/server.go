package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"imw7.com/protobuf/pb"
	"log"
	"net"
)

type HelloService struct{}

func (p *HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + req.GetName() + "!"}, nil
}

func main() {
	// 监听本地的额1234端口
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	// 创建gRPC服务器
	server := grpc.NewServer()
	// 在gRPC服务端注册服务
	pb.RegisterHelloServiceServer(server, new(HelloService))
	// 在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(server)

	// Serve方法在listener上接收传入连接，为每个连接
	// 创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
