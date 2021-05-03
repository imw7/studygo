package taillog

import "LogAgent/etcd"

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	// tskMap map[string]*TailTask
}

func Init(logEntries []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry: logEntries, // 把当前的日志收集项的配置信息保存起来
	}
	for _, logEntry := range logEntries {
		// conf: *etcd.LogEntry
		// logEntry.Path: 要收集的日志文件的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
