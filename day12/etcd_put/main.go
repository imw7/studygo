package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

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
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// value := `[{"path":"/home/tmp/nginx.log","topic":"web_log"},{"path":"/usr/xxx/redis.log","topic":"redis_log"}]`
	value := `[{"path":"/home/tmp/nginx.log","topic":"web_log"},{"path":"/usr/xxx/nginx.log","topic":"redis_log"},{"path":"/usr/xxx/mysql.log","topic":"mysql_log"}]`
	_, err = cli.Put(ctx, "/LogAgent/collect_config", value)
	cancel()
	if err != nil {
		fmt.Println("put to etcd failed, err:", err)
		return
	}
}
