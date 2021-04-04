package main

import "fmt"

// switch
// 简化大量的判断（一个变量和具体的值作比较）

func main() {

	var age int
	fmt.Print("How old are you? ")
	_, err := fmt.Scanln(&age)
	if err != nil {
		fmt.Println("Get age failed, err:", err)
		return
	}
	var stage string
	switch {
	case age < 2:
		stage = "baby"
	case age < 4:
		stage = "toddler"
	case age < 13:
		stage = "children"
	case age < 20:
		stage = "teenager"
	case age < 65:
		stage = "adult"
	default:
		stage = "senior"
	}

	fmt.Printf("You are at the stage of %s.\n", stage)

	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}
}
