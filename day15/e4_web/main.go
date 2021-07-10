package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, _ *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	// 渲染模板
	u1 := User{
		Name:   "梅西",
		Gender: "男",
		Age:    34,
	}
	m1 := map[string]interface{}{
		"name":   "杰克",
		"gender": "男",
		"age":    18,
	}
	hobbies := []string{
		"抽烟",
		"喝酒",
		"烫头",
	}
	err = t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbies,
	})
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:", err)
		return
	}
}
