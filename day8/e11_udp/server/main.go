package main

import (
	"fmt"
	"net"
	"strings"
)

// UDP 服务端

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer func() { _ = listen.Close() }()
	// 不需要建立连接，直接收发数据
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		rpy := strings.ToUpper(string(data[:n]))
		_, err = listen.WriteToUDP([]byte(rpy), addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
