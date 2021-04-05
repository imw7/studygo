package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TCP 客户端

func main() {
	// 1.与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer func() { _ = conn.Close() }() // 关闭连接
	// 2.发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Input-> ")
		msg, _ := reader.ReadString('\n') // 读取用户输入
		msg = strings.Trim(msg, "\r\n")
		if strings.ToUpper(msg) == "EXIT" { // 如果输入exit就退出
			break
		}
		_, err = conn.Write([]byte(msg)) // 发送数据
		if err != nil {
			fmt.Println("send data failed, err:", err)
			return
		}
		var data [512]byte
		n, err := conn.Read(data[:])
		if err != nil {
			fmt.Println("recv data failed, err:", err)
			return
		}
		fmt.Println("Output->", strings.ToUpper(string(data[:n])))
	}
}
