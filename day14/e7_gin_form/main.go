package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 表单参数

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		// 表单参数设置默认值
		type1 := c.DefaultPostForm("type", "alert")
		// 接收其他的
		username := c.PostForm("username")
		password := c.PostForm("password")
		// 多选框
		hobbies := c.PostFormArray("hobby")
		c.String(http.StatusOK,
			fmt.Sprintf("type is %s, username is %s, password is %s, hobbies are %v",
				type1, username, password, hobbies))
	})
	_ = r.Run(":8000")
}
