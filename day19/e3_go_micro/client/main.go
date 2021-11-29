package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"go-micro-demo/pb"
	"log"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	client := pb.NewGreeterService("go.micro.api.hello", service.Client())

	// Make request
	rsp, err := client.Hello(context.Background(), &pb.Request{Name: "John"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.Msg)
}
