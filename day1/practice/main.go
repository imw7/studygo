package main

import (
	"fmt"
	"unicode"
)

// 1. 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用`fmt.Printf()`搭配`%T`分别打印出上述变量的值和类型。
func dataType() {
	n1 := 10
	n2 := 3.1415926
	n3 := true
	n4 := "Hello world!"

	fmt.Printf("n1 = {值: %d, 类型: %T}\n", n1, n1)
	fmt.Printf("n2 = {值: %f, 类型: %T}\n", n2, n2)
	fmt.Printf("n3 = {值: %v, 类型: %T}\n", n3, n3)
	fmt.Printf("n4 = {值: %s, 类型: %T}\n", n4, n4)
}

// 2. 编写代码统计出字符串`"hello中国"`中汉字的数量。
func countHan() {
	s := "hello中国"
	count := 0
	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Printf("字符串\"%s\"中汉字的数量为：%d\n", s, count)
}

// 3. 99乘法表
func multiTable() {
	for x := 1; x < 10; x++ {
		for y := 1; y < x+1; y++ {
			fmt.Printf("%dx%d=%d\t", y, x, x*y)
		}
		fmt.Println()
	}
}

func main() {
	dataType()
	countHan()
	multiTable()
}
