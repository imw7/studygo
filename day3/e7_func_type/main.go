package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("Hello, World!")
}

func f2() int {
	return 10
}

// 函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(x, y int) int {
	return x + y
}

func ff(a, b int) int {
	return a * b
}

// 函数还可以作为返回值
func f5() func(int, int) int {
	return ff
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)
	f3(b)
	fmt.Printf("%T\n", f4)

	f6 := f5()
	fmt.Printf("%T\n", f6)
}
