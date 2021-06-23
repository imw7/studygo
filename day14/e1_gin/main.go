package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin的HelloWorld

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger()，Recover()
	r := gin.Default()
	// 也可以创建不带中间件的路由
	// r := gin.New()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})
	// 3.监听端口，默认在8080
	_ = r.Run(":8000")
}
