package main

import "github.com/imw7/studygo/day8/mylogger"

var log mylogger.Logger // 声明一个全局的接口变量

// 测试自己写的日志库
func main() {
	log = mylogger.NewConsoleLog("Info")                               // 终端日志实例
	log = mylogger.NewFileLogger("Info", "./", "hi.log", 10*1024*1024) // 文件日志实例

	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		id := 10010
		name := "Eric"
		log.Error("这是一条Error日志, id:%d, name:%s.", id, name)
		log.Fatal("这是一条Fatal日志")
		// time.Sleep(time.Second)
	}
}
