package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"imw7.com/pubsub/pb"
	"io"
	"log"
)

// 客户端: 向另一个客户端进行订阅信息

func main() {
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

	client := pb.NewPubsubServiceClient(conn)

	var stream pb.PubsubService_SubscribeClient
	var streams []pb.PubsubService_SubscribeClient

	stream, err = client.Subscribe(context.Background(), &pb.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}
	streams = append(streams, stream)

	// stream, err = client.Subscribe(context.Background(), &pb.String{Value: "docker:"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// streams = append(streams, stream)

	for {
		for _, stream := range streams {
			reply, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}

			fmt.Println(reply.GetValue())
		}
	}
}
