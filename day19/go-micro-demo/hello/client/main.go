package main

import (
	"context"
	"fmt"
	"go-micro-demo/hello/pb"
	"go-micro.dev/v4"
	"log"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// user the generated client stub
	cl := pb.NewHelloService("go.micro.api.Hello", service.Client())

	// make request
	rsp, err := cl.Call(context.Background(), &pb.Request{Name: "John"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.Msg)
}
