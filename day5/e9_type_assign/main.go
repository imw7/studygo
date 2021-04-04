package main

import "fmt"

// 类型断言：想知道空接口中接收的值到底是什么类型

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串：", str)
	}
	fmt.Println(str)
}

func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("it's type string:", t)
	case int:
		fmt.Println("it's type int:", t)
	case int64:
		fmt.Println("it's type int64:", t)
	case bool:
		fmt.Println("it's type bool:", t)
	case []int:
		fmt.Println("it's type []int:", t) // slice
	case map[string]int:
		fmt.Println("it's type map[string]int:", t)
	case func():
		fmt.Println("it's type func():", t)
	}
}

func f() {
}

func main() {
	assign(100)
	assign2(true)
	assign2("hello")
	assign2(int64(100))
	assign2([]int{1, 2, 3})
	assign2(map[string]int{"a": 1})
	assign2(f)
}
