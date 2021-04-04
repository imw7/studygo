package main

import (
	"fmt"
	calc "github.com/imw7/studygo/day5/e10_calc"
)

var x = 100

const pi = 3.14

func init() {
	fmt.Println("自动执行！")
	fmt.Println(x, pi)
}

func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
