package main

import "fmt"

// 匿名函数

func main() {
	// 函数内部没有办法声明带名字的函数
	// 匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(1, 2)

	// 如果只是调用一次的函数，还可以简写成
	// 立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
	}(3, 4)
}
