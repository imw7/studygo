package main

import "fmt"


// 函数可变参数

func f(a ...interface{}) {
	fmt.Printf("type:%T value:%#v\n", a, a)
}

func main() {
	f()
	f(1)
	f(1, false, "hello", struct{}{}, []int{1, 2}, [...]int{3, 4, 5}, map[string]int{"张三": 90})

	var s = []interface{}{1, 3, 5, 7, 9}
	f(s)
	f(s...)
}
