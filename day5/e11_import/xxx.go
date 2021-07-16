package main

import (
	"fmt"
	"imw7.com/studygo/calc"
)

var x = 100

const pi = 3.14

// init 该函数在被导入包的init函数执行完成后执行
func init() {
	fmt.Println("自动执行！")
	fmt.Println(x, pi)
}

// main 该函数在所有init函数执行完成后执行
func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
