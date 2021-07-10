package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()

	// 静态文件：html页面上用到的样式文件 .css js文件 图片

	// 加载静态文件
	r.Static("/xxx", "./statics")

	// gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*") // 模板解析
	// r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl", "templates/link.tmpl")
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

	r.GET("links/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "links/index.tmpl", gin.H{
			"link": "<a href='https://www.baidu.com'>百度一下，你就知道</a>",
		})
	})

	_ = r.Run(":8080") // 启动server
}
