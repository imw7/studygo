package rpc

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestSession_ReadWriter(t *testing.T) {
	// 定义地址
	addr := "127.0.0.1:8080"
	myData := "hello"
	// 等待组定义
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 写数据的协程
	go func() {
		defer wg.Done()
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			t.Error(err)
			return
		}
		conn, _ := listener.Accept()
		s := Session{conn: conn}
		err = s.Write([]byte(myData))
		if err != nil {
			t.Error(err)
			return
		}
	}()

	// 读数据的协程
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Error(err)
			return
		}
		s := Session{conn: conn}
		data, err := s.Read()
		if err != nil {
			t.Error(err)
			return
		}
		// 最后一层校验
		if string(data) != myData {
			t.Error(err)
			return
		}
		fmt.Println(string(data))
	}()
	wg.Wait()
}
