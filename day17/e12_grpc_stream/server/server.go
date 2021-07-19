package main

import (
	"context"
	"google.golang.org/grpc"
	"imw7.com/grpc_stream/pb"
	"io"
	"log"
	"net"
)

// gRPC流

// RPC是远程函数调用，因此每次调用的函数参数和返回值不能太大，否则将严重影响每次调用的响应时间。
// 因此传统的RPC方法调用对于上传和下载较大数据量场景并不适合。同时传统RPC模式也不适用于对时间不
// 确定的订阅和发布模式。为此，gRPC框架针对服务器端和客户端分别提供了流特性。

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

// Channel 可以用于和客户端双向通信
// Send和Recv方法用于流数据的双向通信
func (p *HelloServiceImpl) Channel(stream pb.HelloService_ChannelServer) error {
	for { // 在循环中接收客户端发来的数据
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF { // 表示客户端流被关闭
				return nil
			}
			return nil
		}
		reply := &pb.String{Value: "hello:" + args.GetValue()}

		if err = stream.Send(reply); err != nil {
			return err
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
