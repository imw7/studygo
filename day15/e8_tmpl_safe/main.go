package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func xss(w http.ResponseWriter, _ *http.Request) {
	// 定义模板
	// 解析模板
	// 解析模板之前定义一个自定义的函数safe
	tmpl, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 渲染模板
	jsStr := `<script>alert('嘿嘿嘿')</script>`
	err = tmpl.Execute(w, jsStr)
	if err != nil {
		fmt.Println("execute template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/xss", xss)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP server start failed, err:", err)
		return
	}
}
