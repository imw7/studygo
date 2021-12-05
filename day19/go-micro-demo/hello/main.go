package main

import (
	"go-micro-demo/hello/handler"
	"go-micro-demo/hello/pb"
	"go-micro.dev/v4"
	"log"
)

/*
micro new service hello
cd hello
make proto tidy
micro run
micro call go.micro.api.Hello HelloService.Call '{"name": "John"}'
*/

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("go.micro.api.Hello"), // server name
		micro.Version("latest"),
	)

	// initialise flags
	service.Init()

	if err := pb.RegisterHelloServiceHandler(service.Server(), new(handler.HelloService)); err != nil {
		log.Fatal(err)
	}

	// start the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
