package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro-demo/registerConf/config"
	"go-micro-demo/registerConf/handler"
	"go-micro-demo/registerConf/pb"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
)

func main() {
	// 注册consul
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	// 配置中心
	consulConfig, err := config.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Fatal(err)
	}

	// MySQL配置信息
	info, err := config.GetMySQLFromConsul(consulConfig, "mysql")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("MySQL配置信息:", info)

	// 创建service
	srv := micro.NewService(
		micro.Name("registerConf"),
		micro.Version("latest"),
		// 注册consul中心
		micro.Registry(reg),
	)

	// 注册handler
	if err := pb.RegisterRegisterConfHandler(srv.Server(), &handler.RegisterConf{}); err != nil {
		log.Fatal(err)
	}

	// 运行service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
