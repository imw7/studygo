package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

// 反向RPC

// 通常的RPC是基于C/S结构，RPC的服务端对应网络的服务器，RPC的客户端也对应网络客户端。
// 但是对于一些特殊场景，比如在公司内网提供一个RPC服务，但是在外网无法链接到内网的服务器。
// 这种时候我们可以参考类似反向代理的技术，首先从内网主动链接到外网的TCP服务器，
// 然后基于TCP链接向外网提供RPC服务。

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello, " + request + "!"
	return nil
}

func main() {
	err := rpc.Register(new(HelloService))
	if err != nil {
		log.Fatal("Register error:", err)
	}

	for {
		// 反向RPC的内网服务将不再主动提供TCP监听服务，而是首先主动链接到对方的TCP服务。
		// 然后基于每个建立的TCP链接想对方提供RPC服务。
		conn, err := net.Dial("tcp", ":1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		if err = conn.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
