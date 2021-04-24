package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// watch demo

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Println("connect to etcd failed, err:", err)
		return
	}
	fmt.Println("connect to etcd succeed.")
	defer func() {
		err := cli.Close()
		if err != nil {
			return
		}
	}()
	// watch key:eric change
	rch := cli.Watch(context.Background(), "eric") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type:%s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
