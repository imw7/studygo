package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http Client

// ** 方法1> 共用一个client适用于请求比较频繁
// var (
// 	client = http.Client{
// 		Transport: &http.Transport{
// 			DisableKeepAlives: false,
// 		},
// 	}
// )

func main() {
	// req, err := http.Get("http://0.0.0.0:9090/get/?name=eric&age=19")
	// if err != nil {
	// 	fmt.Println("get url failed, err:", err)
	// 	return
	// }
	apiUrl := "http://0.0.0.0:9090/get/"
	// URL param
	data := url.Values{} // url values
	data.Set("name", "张三")
	data.Set("age", "21")
	u, err := url.Parse(apiUrl)
	if err != nil {
		fmt.Println("parse url requestUrl failed, err:", err)
		return
	}
	u.RawQuery = data.Encode() // URL encode之后的URL
	fmt.Println(u.String())
	// req, err := http.Get(u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println("post failed, err:", err)
		return
	}
	// ** 方法2> 请求不是特别频繁，用完就关闭该连接
	// 禁用KeepAlive的client
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("post failed, err:", err)
		return
	}
	defer func() { _ = resp.Body.Close() }() // 一定要记得关闭resp.Body
	// 发请求
	// 从resp中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(resp.Body) // 在客户端读出服务端返回的响应的body
	if err != nil {
		fmt.Println("get req failed, err:", err)
		return
	}
	fmt.Println(string(b))
}
