package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 为路由组注册中间件

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小明") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	shopGroup := r.Group("/shop", StatCost())
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			time.Sleep(5 * time.Second)
		})
		shopGroup.GET("/home", func(c *gin.Context) {
			time.Sleep(3 * time.Second)
		})
	}
	_ = r.Run(":8080")
}
