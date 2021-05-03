package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

// 专门从日志文件收集日志的模块

var tails *tail.Tail

func Init(filename string) (err error) {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个位置开始读
		ReOpen:    true,                                 // 重新打开
		MustExist: false,                                // 文件不存在不报错
		Follow:    true,                                 // 是否跟随
		Poll:      true,
	}
	tails, err = tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	return
}

func ReadChan() <-chan *tail.Line {
	return tails.Lines
}
