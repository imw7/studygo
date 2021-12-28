package main

import (
	"context"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"grpc-consul/pb"
	"log"
	"net"
)

// 把grpc服务注册到consul上面

type HelloService struct{}

// Hello 绑定方法，实现接口
func (h *HelloService) Hello(_ context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Msg: "Hello, " + req.Name + "!"}, nil
}

func main() {
	// 初始化consul配置，客户端服务器需要一致
	consulConfig := api.DefaultConfig()

	// 获取consul操作对象
	registry, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 注册服务，服务的常规配置
	registerService := api.AgentServiceRegistration{
		ID:      "1",
		Name:    "HelloService",
		Tags:    []string{"grpc", "consul"},
		Port:    1234,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "127.0.0.1:1234",
			Timeout:  "5s",
			Interval: "5s",
		},
	}

	// 将服务注册到consul上
	if err := registry.Agent().ServiceRegister(&registerService); err != nil {
		log.Fatal(err)
	}

	// 设置监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen failed, err:", err)
	}

	// 构造gRPC对象
	grpcServer := grpc.NewServer()
	// 注册服务
	pb.RegisterHelloServiceServer(grpcServer, &HelloService{})

	// 启动服务
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
