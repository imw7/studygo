package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

// etcd 续期

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func(cli *clientv3.Client) {
		err = cli.Close()
		if err != nil {
			return
		}
	}(cli)
	// 设置续期5秒
	response, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	// 将k-v设置到etcd
	_, err = cli.Put(context.TODO(), "root", "password", clientv3.WithLease(response.ID))
	if err != nil {
		log.Fatal(err)
	}
	// 若想一直有效，设置自动续期
	ch, err := cli.KeepAlive(context.TODO(), response.ID)
	if err != nil {
		log.Fatal(err)
	}
	for true {
		c := <-ch
		fmt.Println("chan:", c)
	}
}
