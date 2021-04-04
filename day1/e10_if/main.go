package main

import (
	"fmt"
)

// if 条件判断

func main() {
	// var age string
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("请输入年龄：")
	// age, _ = reader.ReadString('\n')

	var age int
	fmt.Print("请输入年龄：")
	_, _ = fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	} else {
		fmt.Println("该写暑假作业啦！")
	}

	// 多个判断条件
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习！")
	}

	// 作用域
	// a变量此时只在if条件判断语句中生效
	if a := 19; a > 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	} else {
		fmt.Println("该写暑假作业啦！")
	}

	// fmt.Println(a) // 在这里找不到 a 变量
}
