package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-micro-demo/hello/pb"
	"go-micro.dev/v4"
	"go-micro.dev/v4/util/log"
)

// 使用gin框架实现micro web客户端

func main() {
	// 初始化路由
	router := gin.Default()
	// 发送GET请求
	router.GET("/", func(c *gin.Context) {
		// create a new service
		service := micro.NewService()

		// parse command line flags
		service.Init()

		// user the generated client stub
		cl := pb.NewHelloService("hello", service.Client())

		// make request
		rsp, err := cl.Call(context.Background(), &pb.Request{Name: "John"})
		if err != nil {
			log.Fatal(err)
		}
		// 向网页写数据
		_, err = c.Writer.WriteString(rsp.Msg)
		if err != nil {
			log.Fatal(err)
		}
		// 向终端打印数据
		fmt.Println(rsp.Msg)
	})
	// 启动路由
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
