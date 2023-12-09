package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Go语言中推荐使用驼峰式命名
var studentName string

var (
	name   string
	age    int
	gender bool // male: true  female: false
)

func main() {
	studentName = "eric"
	age = 19
	fmt.Printf("%s is %d years old.\n",
		cases.Title(language.English).String(studentName), age)

	name = "sarah"
	age = 10
	var strG string
	gender = false
	switch gender {
	case false:
		strG = "girl"
	case true:
	default:
		strG = "boy"
	}
	fmt.Printf("%s is a %d years old %s.\n",
		cases.Title(language.English).String(name), age, strG)

	s1 := "你好，世界！"
	fmt.Println(s1)
}
