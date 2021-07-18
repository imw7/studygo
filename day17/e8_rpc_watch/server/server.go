package main

import (
	"imw7.com/rpc/db"
	"log"
	"net"
	"net/rpc"
)

// 启动服务

func main() {
	if err := rpc.RegisterName("KVStoreService", db.NewKVStoreService()); err != nil {
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

		go rpc.ServeConn(conn)
	}
}
