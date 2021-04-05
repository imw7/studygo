package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// TCP 服务端

// 处理函数
func process(conn net.Conn) {
	defer func() { _ = conn.Close() }() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var data [128]byte
		n, err := reader.Read(data[:]) // 读取数据
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			return
		}
		recvStr := string(data[:n])
		fmt.Println("收到client端发来的数据:", recvStr)
		_, _ = conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	// 1.本地端口启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer func() { _ = listen.Close() }()
	// 2.等待别人来跟我建立连接
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		// 3.与客户端通信
		go process(conn) // 启动一个goroutine处理连接
	}
}
