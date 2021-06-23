package main

import (
	"fmt"
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
	r.GET("/usr/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})
	r.GET("/welcome", func(c *gin.Context) {
		// DefaultQuery第二个参数是默认值
		// 想要取得不同的值需要在浏览器中输入localhost:8000/welcome?name=Sarah
		name := c.DefaultQuery("name", "Eric")
		c.String(http.StatusOK, fmt.Sprintf("Hello, %s.", name))
	})
	_ = r.Run(":8000")
}
