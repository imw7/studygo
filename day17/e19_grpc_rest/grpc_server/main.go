package main

import (
	"context"
	"google.golang.org/grpc"
	"imw7.com/grpc_rest/pb"
	"log"
	"net"
)

type RestServiceImpl struct{}

func (r *RestServiceImpl) Get(ctx context.Context, message *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value: "Get:" + message.Value}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value: "Post:" + message.Value}, nil
}

func main() {
	// 启动grpc服务
	grpcServer := grpc.NewServer()
	pb.RegisterRestServiceServer(grpcServer, new(RestServiceImpl))
	listener, _ := net.Listen("tcp", ":5000")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
