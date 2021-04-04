package main

import "fmt"

// fmt占位符

func main() {
	var n = 100
	fmt.Printf("%T\n", n) // 查看类型 // int
	fmt.Printf("%v\n", n) // 100
	fmt.Printf("%b\n", n) // 1100100
	fmt.Printf("%d\n", n) // 100
	fmt.Printf("%o\n", n) // 144
	fmt.Printf("%x\n", n) // 64
	var s = "Hello go!"
	fmt.Printf("字符串：%s\n", s)  // 字符串：Hello go!
	fmt.Printf("字符串：%v\n", s)  // 字符串：Hello go!
	fmt.Printf("字符串：%#v\n", s) // 字符串："Hello go!"
}
