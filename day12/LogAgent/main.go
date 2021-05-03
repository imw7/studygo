package main

import (
	"LogAgent/conf"
	"LogAgent/etcd"
	"LogAgent/kafka"
	"LogAgent/taillog"
	"fmt"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

// LogAgent入口程序

var cfg = new(conf.AppConf)

func main() {
	// 0.加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed, err:", err)
		return
	}

	// 1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Println("init Kafka failed, err:", err)
		return
	}
	fmt.Println("init kafka succeed.")

	// 2.初始化ETCD
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("init etcd failed, err:", err)
		return
	}
	fmt.Println("init etcd succeed.")

	// 2.1 从etcd中获取日志收集项目的配置信息
	logEntries, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Println("etcd.GetConf failed, err:", err)
		return
	}
	fmt.Printf("get conf from etcd succeed, %v.\n", logEntries)
	// 2.2 派一个哨兵去监视日志收集项的变化（有变化及时通知我的LogAgent实现热加载配置）

	for index, value := range logEntries {
		fmt.Printf("index:%v value:%v\n", index, value)
	}

	// 3.收集日志发往Kafka
	taillog.Init(logEntries)
	// 因为NewConfChan访问了tskMgr的newConfChan，这个channel是在taillog.Init(logEntries)执行的初始化
	newConfChan := taillog.NewConfChan() // 从taillog包中获取对外暴露的通道

	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfChan) // 哨兵发现最新的配置信息会通知上面的newConfChan通道
	wg.Wait()
}
