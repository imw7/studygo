package main

import (
	"context"
	"google.golang.org/grpc"
	"imw7.com/token/conf"
	"imw7.com/token/pb"
	"log"
	"net"
)

type grpcServer struct {
	auth *conf.Authentication
}

func (g *grpcServer) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 进行Auth方法认证
	if err := g.auth.Auth(ctx); err != nil {
		return nil, err
	}
	reply := &pb.HelloReply{Message: "Hello, " + in.GetName() + "!"}
	return reply, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, new(grpcServer))

	listener, err := net.Listen("tcp", conf.Port)
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
