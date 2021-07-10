package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	msg := "小王子"
	// 渲染模板
	_ = t.Execute(w, msg)
}

func home(w http.ResponseWriter, _ *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	msg := "小王子"
	// 渲染模板
	_ = t.Execute(w, msg)
}

func index2(w http.ResponseWriter, _ *http.Request) {
	// 定义模板（模板继承的方式）
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	// 渲染模板
	name := "张三"
	_ = t.ExecuteTemplate(w, "index2.tmpl", name)
}

func home2(w http.ResponseWriter, _ *http.Request) {
	// 定义模板（模板继承的方式）
	// 解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	// 渲染模板
	name := "李四"
	_ = t.ExecuteTemplate(w, "home2.tmpl", name)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println("HTTP server start failed, err:", err)
		return
	}
}
