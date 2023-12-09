package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-user/pb"
	"log"
)

func main() {
	// 连接服务端
	conn, err := grpc.Dial(":1234",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			return
		}
	}(conn)

	// 实例化grpc客户端
	client := pb.NewUserInfoServiceClient(conn)
	// 调用
	userInfo, err := client.GetUserInfo(context.Background(), &pb.UserRequest{Name: "Eric"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("id:%d name:%s age:%d hobbies:%v\n", userInfo.Id, userInfo.Name, userInfo.Age, userInfo.Hobbies)
}
