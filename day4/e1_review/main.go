package main

import "fmt"

func hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 函数作为参数
func eric(f func(string), name string) {
	f(name)
}

// 函数作为返回值
func sarah() func(int, int) int {
	return func(x int, y int) int {
		return x + y
	}
}

func low(f func()) {
	f()
}

// 闭包 = 函数 + 外部变量的引用
func closure(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func main() {
	eric(hello, "Eric")
	ret := sarah()
	fmt.Printf("%T\n", ret)
	sum := ret(10, 20)
	fmt.Println(sum)
	// 闭包示例
	f := closure(hello, "Peter")
	low(f)
}
