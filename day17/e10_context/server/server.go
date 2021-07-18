package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// 上下文信息

// 基于上下文我们可以针对不同客户端提供定制化的RPC服务。
// 我们可以通过为每个链接提供独立的RPC服务来实现对上下文特性的支持。

type HelloService struct {
	conn    net.Conn // 对应链接
	isLogin bool     // 登陆状态
}

func (p *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogin = true
	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
	// 链接RPC服务前，需要先执行登陆操作，登陆成功后才能正常执行其他的服务
	if !p.isLogin {
		return fmt.Errorf("please login")
	}
	*reply = "hello:" + request + ", from " + p.conn.RemoteAddr().String()
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go func() {
			defer func(conn net.Conn) {
				err = conn.Close()
				if err != nil {
					return
				}
			}(conn)

			p := rpc.NewServer()
			err = p.Register(&HelloService{conn: conn})
			if err != nil {
				log.Fatal("Register error:", err)
			}
			p.ServeConn(conn)
		}()
	}
}
