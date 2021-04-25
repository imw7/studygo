package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	// \ 本来是具有特殊含义的，我应该告诉程序我写的\就是一个单纯的\
	path := "\"E:\\Go\\src\\github.com\\imw7\\day01\""
	fmt.Println(path) // "E:\Go\src\github.com\imw7\day01"

	s := "I'm OK!"
	fmt.Println(s) // I'm OK!

	// 多行的字符串
	s2 := `
		世情薄
		人情恶
		雨送黄昏花易落
	`
	fmt.Println(s2) // 原样输出

	s3 := `E:\Go\src\github.com\imw7\day01`
	fmt.Println(s3) // E:\Go\src\github.com\imw7\day01

	// 字符串相关操作
	fmt.Println(len(s3)) // 40

	// 字符串拼接
	name := "小明"
	word := "大帅比"

	ss := name + word
	fmt.Println(ss) // 小明大帅比

	ss1 := fmt.Sprintf("%s%s\n", name, word)
	// fmt.Printf("%s%s\n", name, word)
	fmt.Println(ss1) // 小明大帅比

	// 分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret) // [E: Go src github.com imw7 day01]

	// 包含
	fmt.Println(strings.Contains(ss, "小花")) // false
	fmt.Println(strings.Contains(ss, "小明")) // true

	// 前缀
	fmt.Println(strings.HasPrefix(ss, "小明")) // true

	// 后缀
	fmt.Println(strings.HasSuffix(ss, "小明")) // false

	// 索引
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))     // 2
	fmt.Println(strings.LastIndex(s4, "b")) // 5

	// 拼接
	fmt.Println(strings.Join(ret, "+")) // E:+Go+src+github.com+imw7+day01

	s5 := "eric"
	fmt.Println(strings.Title(s5)) // Eric
}
