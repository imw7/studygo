package main

import (
	"log"
	"net"
	"net/rpc"
)

// Go实现RPC程序，求矩形的面积和周长

// Params 定义长宽
type Params struct {
	Width, Height int
}

type RectService struct{}

// Area RPC服务端方法，求矩形面积
// 方法必须满足Go语言的RPC规则：
// 方法只能有两个可序列化的参数，其中第一个参数是接收参数(request)，
// 第二个参数是返回给客户端的参数(reply)，必须是指针类型，
// 并且返回一个error类型，同时必须是公开的方法。
func (r *RectService) Area(request Params, reply *int) error {
	*reply = request.Height * request.Width
	return nil
}

// Perimeter 周长
func (r *RectService) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

func main() {
	// 1.注册服务
	// 将RectService类型的对象注册为一个RPC服务
	// rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数
	// 所有注册的方法会放在"RectService"服务空间之下
	rect := new(RectService)
	// 注册一个rect的服务
	err := rpc.RegisterName("RectService", rect)
	if err != nil {
		log.Fatal("RegisterName error:", err)
	}
	// 2.建立一个唯一的TCP链接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	// 3.在该TCP链接上为对方提供RPC服务
	rpc.ServeConn(conn)
}
