package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"micro_etcd/hello/conf"
	"micro_etcd/hello/handler"
	pb "micro_etcd/hello/proto"
)

func main() {
	// 设置etcd为注册中心，配置etcd路径，默认端口2379
	reg := etcd.NewRegistry(
		registry.Addrs(conf.Host + conf.EtcdPort),
	)

	// Create service
	srv := micro.NewService(
		micro.Name(conf.Service),
		micro.Version(conf.Version),
		// 把etcd注册为注册中心
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
