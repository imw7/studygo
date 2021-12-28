package main

import (
	pb "consul_micro_gin/hello/proto"
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

func main() {
	// 初始化路由
	router := gin.Default()

	// 做路由匹配
	router.GET("/", func(c *gin.Context) {
		// 初始化服务发现consul
		reg := consul.NewRegistry()
		// 初始化micro服务对象，指定consul为服务发现
		service := micro.NewService(micro.Registry(reg))
		// 初始化客户端
		cli := pb.NewHelloService("hello", service.Client())
		// 调用远程服务
		rsp, err := cli.Call(context.TODO(), &pb.CallRequest{Name: "John"})
		if err != nil {
			log.Fatal("cli.Call failed, err:", err)
		}
		// 为了方便查看，在打印之前将结果返回给浏览器
		_, err = c.Writer.WriteString(rsp.Msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rsp.Msg)
	})

	// 启动运行
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
