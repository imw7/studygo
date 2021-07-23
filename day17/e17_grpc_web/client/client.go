package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-web/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
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

	reply, err := client.Hello(context.Background(), &pb.String{Value: "Edward"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
