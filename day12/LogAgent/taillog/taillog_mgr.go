package taillog

import (
	"LogAgent/etcd"
	"fmt"
	"time"
)

var tskMgr *tailLogMgr

// tailTask 管理者
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntries []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntries, // 把当前的日志收集项的配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区的通道
	}
	for _, logEntry := range logEntries {
		// conf: *etcd.LogEntry
		// logEntry.Path: 要收集的日志文件的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
	go tskMgr.run()
}

// 监听自己的newConfChan，有了新的配置过来之后就做对应的处理
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			// 1.配置新增
			// 2.配置删除
			// 3.配置变更
			fmt.Println("new config is here~", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan 向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
