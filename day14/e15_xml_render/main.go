package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// XML渲染

func main() {
	r := gin.Default()
	// gin.H 是map[string]interface{}的缩写
	r.GET("/someXML", func(c *gin.Context) {
		// 方式一：自己拼接
		c.XML(http.StatusOK, gin.H{"message": "Hello, world!"})
	})
	r.GET("/moreXML", func(c *gin.Context) {
		// 方式二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小明"
		msg.Message = "Hello, world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	})
	_ = r.Run(":8080")
}
