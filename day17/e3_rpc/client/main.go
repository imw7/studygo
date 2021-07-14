package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type ArithRequest struct {
	A, B int
}

// ArithResponse 返回给客户端的结果
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	req := ArithRequest{A: 9, B: 2}
	var res ArithResponse
	// 乘法
	err = client.Call("ArithService.Multiply", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
	// 商
	err = client.Call("ArithService.Divide", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d / %d = %d\n", req.A, req.B, res.Quo)
	// 余数
	err = client.Call("ArithService.Remainder", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d %% %d = %d\n", req.A, req.B, res.Rem)
}
