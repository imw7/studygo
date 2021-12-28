package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-consul/pb"
	"log"
	"net"
)

type HelloService struct{}

// Hello 绑定方法，实现接口
func (h *HelloService) Hello(_ context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Msg: "Hello, " + req.Name + "!"}, nil
}

func main() {
	// 设置监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen failed, err:", err)
	}

	// 构造gRPC对象
	grpcServer := grpc.NewServer()
	// 注册服务
	pb.RegisterHelloServiceServer(grpcServer, &HelloService{})

	// 启动服务
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
