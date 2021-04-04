package main

import (
	"fmt"
	"strings"
)

// 流程控制之跳出for循环

func main() {
	// 当i=5时，跳出for循环
	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// 当i=5时，跳过此次for循环（不执行for循环内部的语句），继续下一次循环
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	var name string
	fmt.Println("What's your name? ")
	_, _ = fmt.Scanln(&name)
	fmt.Printf("Welcome, %s!\n", strings.Title(name))
}
