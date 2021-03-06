package main

import (
	"LogAgent/conf"
	"LogAgent/kafka"
	"LogAgent/taillog"
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

// LogAgent入口程序

var cfg = new(conf.AppConf)

func run() {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2.发送到Kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 0.加载配置文件
	// cfg, err := ini.Load("./conf/config.ini")
	// fmt.Println(cfg.Section("kafka").Key("address"))
	// fmt.Println(cfg.Section("kafka").Key("topic"))
	// fmt.Println(cfg.Section("taillog").Key("filename"))
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed, err:", err)
		return
	}

	// 1.初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Println("init Kafka failed, err:", err)
		return
	}
	fmt.Println("init kafka succeed.")

	// 2.打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.Filename)
	if err != nil {
		fmt.Println("init taillog failed, err:", err)
		return
	}
	fmt.Println("init taillog succeed.")

	// 3.具体业务操作
	run()
}
