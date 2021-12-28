package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"micro_etcd/hello/conf"
	pb "micro_etcd/hello/proto"
)

func main() {
	// init engine
	router := gin.Default()

	// get method
	router.GET("/", func(c *gin.Context) {
		// 初始化服务发现etcd
		reg := etcd.NewRegistry()
		// 初始化micro服务对象，指定etcd为服务发现
		service := micro.NewService(micro.Registry(reg))
		// 初始化客户端
		cli := pb.NewHelloService(conf.Service, service.Client())
		// 调用远程服务
		rsp, err := cli.Call(context.TODO(), &pb.CallRequest{Name: "John"})
		if err != nil {
			log.Fatal(err)
		}
		// 将结构返回给浏览器
		_, err = c.Writer.WriteString(rsp.Msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rsp.Msg)
	})

	// run
	if err := router.Run(conf.Host + conf.GinPort); err != nil {
		log.Fatal(err)
	}
}
