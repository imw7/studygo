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
		// 初始化的时候起了多少个tailTask都要记下来，为了后续判断方便
		tailTask := NewTailTask(logEntry.Path, logEntry.Topic)
		pt := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[pt] = tailTask
	}
	go tskMgr.run()
}

// 监听自己的newConfChan，有了新的配置过来之后就做对应的处理
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				pt := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[pt]
				if ok {
					// 如果存在，不需要操作
					continue
				} else {
					// 如果不存在的，当做新增的
					tails := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[pt] = tails
				}
			}
			// TODO 找出原来t.logEntry有，但是newConf中没有的，要删掉

			// 配置删除
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
