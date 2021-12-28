package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-consul/pb"
	"log"
	"strconv"
)

func main() {
	// 初始化consul配置，客户端服务端需要一致
	consulConfig := api.DefaultConfig()

	// 获取consul操作对象
	registerClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 服务注销
	// if err := registerClient.Agent().ServiceDeregister("1"); err != nil {
	// 	log.Fatal(err)
	// }

	// 服务发现，从consul上获取健康的服务
	// params:
	// 	@service: 服务名。注册服务时指定该string
	// 	@tag: 外号/别名。如果有多个，任选一个
	// 	@passingOnly: 是否通过健康检查
	// 	@q: 查询参数，通常为nil
	// returns:
	// 	@ServiceEntry: 存储服务的切片
	//  @QueryMeta: 额外查询返回值，通常为nil
	//  @error: 错误信息
	serviceEntry, _, err := registerClient.Health().Service(
		"HelloService", "consul", false, &api.QueryOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// 和grpc服务建立连接
	// conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(serviceEntry[0].Service.Address+":"+strconv.Itoa(serviceEntry[0].Service.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	// 初始化 gRPC 客户端
	client := pb.NewHelloServiceClient(conn)

	// 调用远程函数
	rsp, err := client.Hello(context.TODO(), &pb.Request{Name: "John"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.Msg)
}
