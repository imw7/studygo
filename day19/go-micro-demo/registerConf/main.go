package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro-demo/registerConf/config"
	"go-micro-demo/registerConf/handler"
	"go-micro-demo/registerConf/pb"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

func main() {
	// 1.注册consul
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	// 完成前4步之后，在consul管理页面的Key/Value创建micro/config目录，
	// 然后在config目录创建mysql键，以json格式输入mysql配置的值：
	// {
	//	"host": "127.0.0.1",
	//  "user": "root",
	//  "pwd": "password",
	//  "database": "sql_test",
	//  "port": 3306
	// }
	// 具体操作可以在我的博客中找到：
	// https://imw7.github.io/post/Go/go_micro/#%E6%B3%A8%E5%86%8C%E5%92%8C%E9%85%8D%E7%BD%AE%E4%B8%AD%E5%BF%83
	// 5.配置中心
	consulConfig, err := config.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		logger.Fatal(err)
	}

	// 6.获取MySQL配置信息
	info, err := config.GetMySQLFromConsul(consulConfig, "mysql")
	if err != nil {
		logger.Fatal(err)
	}
	// 必须使用logger才行，否则无法显示info信息
	logger.Info("MySQL配置信息:", *info)

	// 2.创建service
	srv := micro.NewService(
		micro.Name("registerConf"),
		micro.Version("latest"),
		// 注册consul中心
		micro.Registry(reg),
	)

	// 3.注册handler
	if err := pb.RegisterRegisterConfHandler(srv.Server(), &handler.RegisterConf{}); err != nil {
		logger.Fatal(err)
	}

	// 4.运行service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
