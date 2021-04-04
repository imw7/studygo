package main

import "github.com/imw7/studygo/day6/mylogger"

// 测试自己写的日志库

var log mylogger.Logger

func main() {
	log = mylogger.NewConsoleLogger("Info")                                // 终端日志实例
	log = mylogger.NewFileLogger("Info", "./", "my_log.log", 10*1024*1024) // 文件日志实例
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		id := 10086
		name := "中国移动"
		log.Error("这是一条Error日志，id：%d，name：%s", id, name)
		log.Fatal("这是一条Fatal日志")
		// time.Sleep(time.Second)
	}
}
