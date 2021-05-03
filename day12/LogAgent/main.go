package main

import (
	"LogAgent/conf"
	"LogAgent/etcd"
	"LogAgent/kafka"
	"LogAgent/taillog"
	"fmt"
	"gopkg.in/ini.v1"
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
	// 2.2 派一个哨兵去监视日志收集项的变化（有变化及时通知我的LogAgent实现热加载配置）
	if err != nil {
		fmt.Println("etcd.GetConf failed, err:", err)
		return
	}
	fmt.Printf("get conf from etcd succeed, %v.\n", logEntries)
	for index, value := range logEntries {
		fmt.Printf("index:%v value:%v\n", index, value)
	}

	// 3.收集日志发往Kafka
	taillog.Init(logEntries)
}
