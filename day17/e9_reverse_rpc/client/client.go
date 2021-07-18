package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// doClientWork 执行RPC调用的操作
func doClientWork(clientChan <-chan *rpc.Client) {
	client := <-clientChan // 从管道中去取一个RPC客户端对象
	defer func() {         // 通过defer语句指定在函数退出前关闭客户端
		if err := client.Close(); err != nil {
			return
		}
	}()

	var reply string
	if err := client.Call("HelloService.Hello", "Jesse", &reply); err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 基于网络链接构造RPC客户端对象并发送到clientChan管道
		clientChan <- rpc.NewClient(conn)
	}()

	doClientWork(clientChan)
}
