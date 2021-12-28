package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro-demo/hello/pb"
	"go-micro.dev/v4"
	"net/http"
)

type demo struct{}

func NewDemo() *demo {
	return &demo{}
}

func (d *demo) InitRouter(router *gin.Engine) {
	router.POST("/demo", d.demo)
	router.GET("/demo", d.demo)
}

func (d *demo) demo(c *gin.Context) {
	// create a service
	service := micro.NewService()
	service.Init()

	client := pb.NewHelloService("hello", service.Client())

	rsp, err := client.Call(context.Background(), &pb.Request{Name: "John"})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  rsp.Msg,
	})
}
