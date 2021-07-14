package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Params 传的参数
type Params struct {
	Width, Height int
}

func main() {
	// 1.连接远程rpc服务
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// 2.调用方法
	var ret int
	// 面积
	err = client.Call("RectService.Area", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("面积:", ret)
	// 周长
	err = client.Call("RectService.Perimeter", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("周长:", ret)
}
