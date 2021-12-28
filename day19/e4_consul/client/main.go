package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-consul/pb"
	"log"
)

func main() {
	// 连接服务
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	// 初始化 gRPC 客户端
	client := pb.NewHelloServiceClient(conn)

	// 调用远程函数
	rsp, err := client.Hello(context.TODO(), &pb.Request{Name: "John"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.Msg)
}
