package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

// etcd lease 租约

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to etcd succeed.")
	defer func() {
		err := cli.Close()
		if err != nil {
			return
		}
	}()

	// 构建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	// 5秒钟之后，/eric/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "/eric/", "male", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
}
