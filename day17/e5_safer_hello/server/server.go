package main

import (
	"imw7.com/safer_hello/model"
	"log"
	"net"
	"net/rpc"
)

// HelloService 注册RPC服务用的名字
type HelloService struct{}

// Hello 方法必须满足Go语言的RPC规则：
// 方法只能有两个可序列化的参数，其中第一个参数是接收参数，
// 第二个参数是返回给客户端的参数，且必须是指针类型，
// 并且返回一个error类型，同时必须是公开的方法。
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello, " + request + "!"
	return nil
}

func main() {
	err := model.RegisterHelloService(new(HelloService))
	if err != nil {
		log.Fatal("RegisterName error:", err)
	}
	// 建立唯一的TCP链接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		// 在该TCP链接上为对方提供RPC服务
		go rpc.ServeConn(conn)
	}
}
