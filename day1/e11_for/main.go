package main

import "fmt"

// for 循环

func main() {
	// 基本格式
	for x := 1; x < 10; x++ {
		for y := 1; y < x+1; y++ {
			fmt.Printf("%dx%d=%d\t", y, x, x*y)
		}
		fmt.Println()
	}

	// 变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// 变种2
	var j = 5
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// for range 循环
	s := "Hello中国"
	for k, v := range s {
		fmt.Printf("%d %c\n", k, v) // %c:该值对应的unicode码值
	}

	// 无限循环
	// for {
	// 	fmt.Println("I love Go.")
	// }
}
