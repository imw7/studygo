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
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, new(HelloService))
	reflection.Register(server)

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
