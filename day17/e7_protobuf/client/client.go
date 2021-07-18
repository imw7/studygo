package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"imw7.com/protobuf/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer func() { _ = conn.Close() }()

	client := pb.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &pb.HelloRequest{Name: "Eric"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetMessage())
}
