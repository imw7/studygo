package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"imw7.com/grpc_stream/pb"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	client := pb.NewHelloServiceClient(conn)

	stream, err := client.Channel(context.Background()) // 获取返回的流对象
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for { // 向服务端发送数据
			if err = stream.Send(&pb.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv() // 接收服务端返回的数据
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
