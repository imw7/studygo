package main

import (
	"bufio"
	"fmt"
	"imw7.com/safer_hello/model"
	"log"
	"os"
	"strings"
)

// 获取用户输入的名字
func getName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What's your name? ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	name = strings.Title(name)
	return name
}

func main() {
	// 拨号RPC服务
	client, err := model.DialHelloService("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello(getName(), &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
