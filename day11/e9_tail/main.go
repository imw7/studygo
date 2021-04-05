package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tailf的用法示例

func main() {
	fileName := "./my.log"
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个位置开始读
		ReOpen:    true,                                 // 重新打开
		MustExist: false,                                // 文件不存在不报错
		Follow:    true,                                 // 是否跟随
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen, filename:", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
