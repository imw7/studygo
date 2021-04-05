package main

import (
	"fmt"
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
		msg := `Hello, hello. How are you?`
		_, _ = conn.Write([]byte(msg))
	}
}
