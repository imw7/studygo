package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	"go-micro-demo/pb"
	"log"
)

type Greeter struct{}

func (g *Greeter) Hello(_ context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Msg = "Hello, " + req.Name + "!"
	return nil
}

func main() {
	// Create service
	service := micro.NewService(micro.Name("go.micro.srv.HelloWorld"))

	// Register handler
	if err := pb.RegisterGreeterHandler(service.Server(), new(Greeter)); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
