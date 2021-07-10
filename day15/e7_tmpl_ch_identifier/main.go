package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.New("index.tmpl").
		Delims("{[", "]}"). // 修改模板引擎的标识符
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	// 渲染模板
	name := "Edward"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("execute template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP server start failed, err:", err)
		return
	}
}
