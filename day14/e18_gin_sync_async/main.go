package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 同步异步
// goroutine机制可以方便地实现异步处理
// 另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本

func main() {
	r := gin.Default()
	// 1.异步
	r.GET("/async", func(c *gin.Context) {
		// 需要context的副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})
	// 2.同步
	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})

	_ = r.Run(":8080")
}
