package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller() 获取调用函数行号和文件名

func f1() {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file) // /home/imw7/Go/src/github.com/imw7/studygo/day6/e6_runtime/main.go
	fmt.Println(path.Base(file))
	fmt.Println(line) // 11
}

func f() {
	f1()
}

func main() {
	f()
}
