package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var cli *clientv3.Client

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// Init 初始化ETCD的函数
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		// handle error!
		fmt.Println("connect to etcd failed, err:", err)
		return
	}
	return
}

// GetConf 从ETCD中根据key获取配置项
func GetConf(key string) (logEntries []*LogEntry, err error) {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Println("get from etcd failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		// fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logEntries)
		if err != nil {
			fmt.Println("unmarshal etcd value failed, err:", err)
			return
		}
	}
	return
}
