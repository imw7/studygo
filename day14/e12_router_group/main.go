package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由组

func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "welcome"})
		})
		userGroup.GET("/login", func(c *gin.Context) {
			name := c.DefaultQuery("name", "Eric")
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Hello, %s.", name),
			})
		})
		userGroup.POST("/login", func(c *gin.Context) {
			name := c.DefaultQuery("name", "Sarah")
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Hello, %s.", name),
			})
		})
	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "welcome"})
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			cart := c.DefaultQuery("cart", "food")
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("There is %s in your cart.", cart),
			})
		})
		shopGroup.POST("/checkout", func(c *gin.Context) {
			cost := c.DefaultQuery("cost", "10")
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("It costs $%s.", cost),
			})
		})
		// 嵌套路由组
		xx := shopGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "world",
			})
		})
	}
	_ = r.Run(":8080")
}
