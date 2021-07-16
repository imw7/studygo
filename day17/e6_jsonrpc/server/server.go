package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 跨语言的RPC

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello, " + request + "!"
	return nil
}

func main() {
	if err := rpc.RegisterName("HelloService", new(HelloService)); err != nil {
		log.Fatal("Register error:", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 给当前连接提供针对json格式的rpc服务
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
