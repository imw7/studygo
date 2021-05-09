package taillog

import (
	"LogAgent/kafka"
	"context"
	"fmt"
	"github.com/hpcloud/tail"
)

// 专门从日志文件收集日志的模块

// TailTask 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail // tail打开的文件
	// 为了能实现退出t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tails *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tails = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	err := tails.init() // 根据路径去打开对应的日志
	if err != nil {
		fmt.Println("init tail failed, err:", err)
		return
	}
	return
}

func (t *TailTask) init() (err error) {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个位置开始读
		ReOpen:    true,                                 // 重新打开
		MustExist: false,                                // 文件不存在不报错
		Follow:    true,                                 // 是否跟随
		Poll:      true,
	}
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	// 当goroutine执行的函数退出的时候，goroutine就结束了
	go t.run() // 直接去采集日志发往kafka
	return
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task:%s_%s exit...\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines: // 从tails的通道中一行一行的读取日志数据
			// 发往Kafka
			// kafka.SendToKafka(t.topic, line.Text) // 函数调用函数

			// 先把日志数据发送到一个通道中
			fmt.Printf("get log from %s succeed, data: %v\n", t.path, line.Text)
			kafka.SendToChan(t.topic, line.Text)
			// kafka那个包中有单独的goroutine去取日志数据发送到kafka
		}
	}
}
