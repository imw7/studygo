package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Default返回一个默认的路由引擎
	r := gin.Default()
	// 获取querystring参数
	// querystring指的是URL中?后面携带的参数，例如：/usr/search?username=小明&address=成都
	r.GET("/usr/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小明")
		// username := c.Query("username")
		address := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	// 获取form参数
	r.POST("/usr/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		// username := c.DefaultPostForm("username", "Eric")
		username := c.PostForm("username")
		address := c.PostForm("address")
		// 输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	// 获取path参数
	r.GET("/usr/search/:username/:address", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		// 输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	_ = r.Run()
}
