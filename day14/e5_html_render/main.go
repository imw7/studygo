package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HTML渲染

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	// r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})
	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	_ = r.Run(":8080")
}
