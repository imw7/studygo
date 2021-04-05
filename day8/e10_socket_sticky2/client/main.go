package main

import (
	"fmt"
	"github.com/imw7/studygo/day8/e10_socket_sticky2/proto"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer func() { _ = conn.Close() }()
	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf(`[%02d]Hello, hello. How are you?`, i+1)
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		_, _ = conn.Write(data)
	}
}
