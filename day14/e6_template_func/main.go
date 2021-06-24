package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 自定义模板函数

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://imw7.github.io'>魏奇的博客</a>")
	})
	_ = r.Run(":8080")
}
