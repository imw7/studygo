package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关代码

var (
	// MaxSize 日志通道缓冲区大小
	MaxSize = 50000
)

// FileLogger 文件日志结构体
type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	line      int
	fileName  string
	timeStamp string
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, MaxSize),
	}
	err = fl.initFile() // 按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

// 根据指定的日志文件路径和文件名打开日志文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Open log file failed, err:", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Open err log file failed, err:", err)
		return err
	}
	// 日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	go f.writeLogBackground() // 开启一个goroutine执行写日志文件
	return nil
}

// 判断是否需要记录日志
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

// 判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Print("Open file failed, err:", err)
		return false
	}
	// 如果当前文件大小 >= 日志文件的最大值，就应该返回true
	return fileInfo.Size() >= f.maxFileSize
}

// 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Get file information failed, err:", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())      // 拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) // 拼接一个日志文件备份的名字
	// 1. 关闭当前的日志文件
	_ = file.Close()
	// 2. 备份一下 rename  xx.log -> xx.log.bak202006072123
	_ = os.Rename(logName, newLogName)
	// 3. 打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Open new log file failed, err:", err)
		return nil, err
	}
	// 4. 将打开的新日志文件对象赋值给 f.fileObj
	return fileObj, nil
}

// 后台写日志文件
func (f *FileLogger) writeLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj) // 日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			// 把日志拼出来
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n",
				logTmp.timeStamp,
				getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)

			_, _ = fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.splitFile(f.errFileObj) // 日志文件
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				// 如果要记录的日志大于ERROR级别，还要在 err 日志文件中再记录一次
				_, _ = fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 取不到日志先休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		// 先把日志发送到通道中
		// 1.造一个logMsg对象
		logTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			line:      lineNo,
			fileName:  fileName,
			timeStamp: now.Format("2006-01-02 15:04:05"),
		}
		select {
		case f.logChan <- logTmp:
		default:
			// 把日志就丢掉保证不出现阻塞
		}
	}
}

// Debug ...
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Info ...
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning ...
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error ...
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal ...
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

// 关闭文件
func (f *FileLogger) Close() {
	_ = f.fileObj.Close()
	_ = f.errFileObj.Close()
}
