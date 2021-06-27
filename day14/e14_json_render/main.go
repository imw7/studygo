package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JSON渲染

func main() {
	r := gin.Default()

	// gin.H 是map[string]interface{}的缩写
	r.GET("/someJSON", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		// 方法二：使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小明"
		msg.Message = "Hello, world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
	_ = r.Run(":8080")
}
