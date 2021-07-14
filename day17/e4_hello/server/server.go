package main

import (
	"log"
	"net"
	"net/rpc"
)

// RPC版HelloWorld

// HelloService 注册用服务名
type HelloService struct{}

// Hello 方法必须满足Go语言的RPC规则：
// 方法只能有两个可序列化的参数，其中第一个参数是接收参数，
// 第二个参数是返回给客户端的参数，且必须是指针类型，
// 并且返回一个error类型，同时必须是公开的方法。
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	// 将HelloService类型的对象注册为一个RPC服务
	// rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数
	// 所有注册的方法会放在"HelloService"服务空间之下
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		log.Fatal("RegisterName error:", err)
	}
	// 建立唯一的TCP链接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	// 在该TCP链接上为对方提供RPC服务
	rpc.ServeConn(conn)
}
