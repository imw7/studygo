package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, _ *http.Request) {
	// 定义一个函数kua
	kua := func(name string) (string, error) {
		return name + "年轻又有朝气！", nil
	}
	// 定义模板
	t := template.New("f.tmpl") // 创建一个名字是f.tmpl的模板对象，名字一定要与模板的名字能对应上

	// 告诉模板引擎，现在多了一个自定义的函数kua
	t.Funcs(template.FuncMap{
		"kua99": kua,
	})
	// 解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	name := "索隆"
	// 渲染模板
	_ = t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	name := "路飞"
	// 渲染模板
	_ = t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println("HTTP server start failed, err:", err)
		return
	}
}
