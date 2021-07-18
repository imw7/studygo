package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	if err = client.Call("HelloService.Login", "user:password", &reply); err != nil {
		log.Fatal(err)
	}

	if err = client.Call("HelloService.Hello", "hello", &reply); err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
