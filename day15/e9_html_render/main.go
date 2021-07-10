package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*") // 模板解析
	// r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // 模板渲染
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users/index",
		})
	})

	_ = r.Run(":8080") // 启动server
}
