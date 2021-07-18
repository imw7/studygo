package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func doClientWork(client *rpc.Client) {
	go func() { // 用于监控key的变化
		var keyChanged string
		if err := client.Call("KVStoreService.Watch", 30, &keyChanged); err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	}()
	// 通过Set方法修改KV值，服务器会将变化的key通过Watch方法返回
	if err := client.Call("KVStoreService.Set", [2]string{"abc", "abc-value"}, new(struct{})); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)
}

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	doClientWork(client)

	var value string
	if err = client.Call("KVStoreService.Get", "abc", &value); err != nil {
		log.Fatal(err)
	}
	fmt.Println("abc:", value)
}
