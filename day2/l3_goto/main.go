package main

import "fmt"

// goto

func main() {
	// 跳出多层for循环
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'F' {
				flag = true
				break // 跳出内层的for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			fmt.Println("over")
			break // 跳出外层for循环
		}
	}

	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'E' {
				goto XX // 跳到指定的标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
XX: // label标签
	fmt.Println("over")
}
