package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	// 建立基于json编解码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "Eric", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

// 通过nc命令 nc -l 1234 在同样的端口启动一个TCP服务，然后再次执行RPC调用将会发现nc输出了以下的信息:
// {"method":"HelloService.Hello","params":["Eric"],"id":0}
// 这是一个json编码的数据，其中method部分对应要调用的rpc服务和方法组合成的名字，params部分的第一个元素为参数，
// id是由调用端维护的一个唯一的调用编号。

// 在获取到RPC调用对应的json数据后，我们可以通过直接向架设了RPC服务的TCP服务器发送json数据模拟RPC方法调用：
// $ echo -e '{"method":"HelloService.Hello","params":["Eric"],"id":1}' | nc localhost 1234
// 返回的结果也是一个json格式的数据:
// {"id":0,"result":"Hello, Eric!","error":null}
// 其中id对应输入的id参数，result为返回的结果，error部分在出问题时表示错误信息。
// 对于顺序调用来说，id不是必须的。但是Go语言的RPC框架支持异步调用，
// 当返回结果的顺序和调用的顺序不一致时，可以通过id来识别对应的调用。
