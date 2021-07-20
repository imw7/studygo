package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"imw7.com/token/conf"
	"imw7.com/token/pb"
	"log"
)

func main() {
	auth := conf.Authentication{
		User:     "gopher",
		Password: "password",
	}
	// 在每次请求gRPC服务时将Token信息(auth)作为参数传入
	conn, err := grpc.Dial(conf.Port, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	client := pb.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &pb.HelloRequest{Name: "Bruce"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetMessage())
}
