package main

import (
	"consul_micro_gin/hello/handler"
	pb "consul_micro_gin/hello/proto"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	service = "hello"
	version = "latest"
)

func main() {
	// 初始化服务发现
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(reg),
	)
	srv.Init()

	// Register handler
	if err := pb.RegisterHelloHandler(srv.Server(), new(handler.Hello)); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
