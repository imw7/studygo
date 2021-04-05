package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// UDP 客户端

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("connect server failed, err:", err)
		return
	}
	defer func() { _ = socket.Close() }()
	for {
		fmt.Print("Input-> ")
		reader := bufio.NewReader(os.Stdin)
		msg, err := reader.ReadString('\n')
		msg = strings.Trim(msg, "\r\n")
		if strings.ToUpper(msg) == "EXIT" {
			break
		}
		if err != nil {
			fmt.Println("input error, err:", err)
			return
		}
		_, err = socket.Write([]byte(msg)) // 发送数据
		if err != nil {
			fmt.Println("send data failed, err:", err)
			return
		}
		data := make([]byte, 4096)
		n, _, err := socket.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("recv data failed, err:", err)
			return
		}
		fmt.Printf("Output-> %v\n", string(data[:n]))
	}
}
