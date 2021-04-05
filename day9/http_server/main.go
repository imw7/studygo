package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, _ *http.Request) {
	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf("%v\n", err)))
	}
	_, _ = w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，参数都放在URL上（query param），请求体中是没有数据的。
	queryParam := r.URL.Query() // 自动识别URL中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) // 在服务端打印客户端发来的请求的body
	_, _ = w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/get/", f2)
	_ = http.ListenAndServe("0.0.0.0:9090", nil)
}
