package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟实现权限验证中间件
// 1.有2个路由，login和home
// 2.login用于设置cookie
// 3.home是访问查看信息的请求
// 4.在请求home之前，先跑中间件代码，检验是否存在cookie

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("cookie"); err == nil {
			if cookie == "login" {
				c.Next()
				return
			}
		}

		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "StatusUnauthorized"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("cookie", "login", 60, "/", "localhost", false, true)
		// 返回信息
		c.JSON(http.StatusOK, gin.H{
			"message": "Login succeed!",
		})
	})
	r.GET("/home", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})
	_ = r.Run(":8080")
}
