package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 取api参数
	r.GET("/usr/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, name) // 浏览器输入localhost:8000/usr/eric 输出eric
	})
	_ = r.Run(":8000")
}
