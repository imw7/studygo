package main

import (
	"context"
	"google.golang.org/grpc"
	"imw7.com/grpc/pb"
	"log"
	"net"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "Hello, " + args.GetValue() + "!"}
	return reply, nil
}

func main() {
	// 构造一个gRPC服务对象
	grpcServer := grpc.NewServer()
	// 通过gRPC插件生成的RegisterHelloServiceServer函数注册HelloServiceImpl服务
	pb.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	// 在一个监听端口提供gRPC服务
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
