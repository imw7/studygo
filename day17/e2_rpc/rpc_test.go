package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"rpc_demo/client"
	"rpc_demo/server"
	"testing"
)

// 给服务端注册一个查询用户的方法，客户端使用RPC方式调用

// 定义用户对象
type User struct {
	Name string
	Age  int
}

// 用于测试用户查询的方法
func queryUser(uid int) (User, error) {
	user := make(map[int]User)
	// 假数据
	user[0] = User{"张三", 20}
	user[1] = User{"李四", 19}
	user[2] = User{"王五", 24}
	// 模拟查询用户
	if u, ok := user[uid]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("%d err", uid)
}

func TestRPC(t *testing.T) {
	// 编码中有一个字段是interface{}时，要注册一下
	gob.Register(User{})
	addr := "127.0.0.1:8080"
	// 创建服务端
	srv := server.NewServer(addr)
	// 将服务端方法，注册一下
	srv.Register("queryUser", queryUser)
	// 服务端等待调用
	go srv.Run()
	// 客户端获取连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("dial failed, error:", err)
	}
	// 创建客户端对象
	cli := client.NewClient(conn)
	// 需要声明函数原型
	var query func(int) (User, error)
	cli.CallRPC("queryUser", &query)
	// 得到查询结果
	u, err := query(1)
	if err != nil {
		log.Fatal("query failed, error:", err)
	}
	fmt.Println(u)
}
