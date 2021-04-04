package main

import (
	"fmt"
	"strings"
)

// 闭包

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")

	fmt.Println(jpgFunc("hello")) // hello.jpg
	fmt.Println(txtFunc("hi.txt"))
	fmt.Println(txtFunc("hello")) // hello.txt
	fmt.Println(jpgFunc("hi.jpg"))
}
