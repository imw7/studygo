package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var cli *clientv3.Client

// LogEntry 需要收集的日志的配置信息
type LogEntry struct {
	Path  string `json:"path"`  // 日志存放的路径
	Topic string `json:"topic"` // 日志要发往Kafka中的哪个Topic
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

// WatchConf etcd watch
func WatchConf(key string, ch chan<- []*LogEntry) {
	rch := cli.Watch(context.Background(), key) // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type:%s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			// 通知taillog.tskMgr
			// 1.先判断操作的类型
			var newConf []*LogEntry
			// 如果是删除操作，手动传递一个空的配置项
			if ev.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(ev.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("unmarshal failed, err:", err)
					continue
				}
			}
			fmt.Println("get new conf:", newConf)
			ch <- newConf
		}
	}
}
