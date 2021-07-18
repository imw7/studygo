package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"imw7.com/grpc/pb"
	"log"
)

func main() {
	// 和gRPC服务建立链接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	// 基于已经建立的链接构造HelloServiceClient对象
	client := pb.NewHelloServiceClient(conn)
	// 返回的client其实是一个HelloServiceClient接口对象
	// 通过接口定义的方法就可以调用服务端对应的gRPC服务提供的方法
	reply, err := client.Hello(context.Background(), &pb.String{Value: "Eric"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
