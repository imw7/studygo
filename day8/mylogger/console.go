package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewConsoleLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return c.Level <= logLevel
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(DEBUG) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n",
			now.Format("2006-01-02 15:04:05"),
			getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}
