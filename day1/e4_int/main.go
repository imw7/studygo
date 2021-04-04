package main

import "fmt"

// 整型

func main() {
	// 十进制转二进制：除2余数从下往上写
	// 二进制转八进制：3位二进制为1位八进制

	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1) // 101
	fmt.Printf("%b\n", i1) // 把十进制转化为二进制 // 1100101
	fmt.Printf("%o\n", i1) // 把十进制转化为八进制 // 145
	fmt.Printf("%x\n", i1) // 把十进制转化为十六进制 // 65

	// 八进制
	i3 := 0x1234567        // 7*16^0 + 6*16^1 + 5*16^2 + ...
	fmt.Printf("%d\n", i3) // 19088743

	// 查看变量的类型
	fmt.Printf("%T\n", i3) // int

	// 声明int8类型的变量
	i4 := int8(9)          // 明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4) // int8
}
