package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// 服务端要给客户端cookie
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie") // 获取Cookie
		if err != nil {
			cookie = "NotSet"
			// 设置Cookie
			// maxAge 单位为秒
			// secure 是否只能通过https访问
			// httpOnly 是否允许别人通过js获取自己的cookie
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})

	_ = r.Run(":8080")
}
