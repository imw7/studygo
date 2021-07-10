package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件

// handlerFunc
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name") // 在上下文c中取值（跨中间件存取值）
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件m1:统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	// 计时
	start := time.Now()
	// go funcXX(c.Copy()) // 在funcXX中只能使用c的拷贝
	c.Next() // 调用后续的处理函数
	// c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "eric") // 在上下文c中设置值
	// c.Abort() // 阻止调用后续的处理函数
	// return
	fmt.Println("m2 out...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 连接数据库
	// 或者一些其他准备工作

	return func(c *gin.Context) {
		// 存放其他的逻辑
		if doCheck {
			// 是否登录的判断
			// if 是登录用户
			c.Next()
			// else
			// c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	// r := gin.Default() // 默认使用Logger中间件和Recovery中间件
	r := gin.New() // 不包含任何中间件

	r.Use(m1, m2, authMiddleware(true)) // 全局注册中间件函数m1,m2

	// GET(relativePath string, handlers ...HandlerFunc) IRoutes
	// r.GET("/index", m1, indexHandler) // 不使用注册中间件函数就该这么写
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		fmt.Println("shop")
		c.JSON(http.StatusOK, gin.H{"msg": "shop"})
	})
	r.GET("/user", func(c *gin.Context) {
		fmt.Println("user")
		c.JSON(http.StatusOK, gin.H{"msg": "user"})
	})

	// // 路由组注册中间件方法1:
	// xxGroup := r.Group("/xx", authMiddleware(true))
	// {
	// 	xxGroup.GET("/index", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
	// 	})
	// }
	// // 路由组注册中间件方法2:
	// xx2Group := r.Group("/xx2")
	// xx2Group.Use(authMiddleware(true))
	// {
	// 	xxGroup.GET("/index", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
	// 	})
	// }

	_ = r.Run(":8080")
}
