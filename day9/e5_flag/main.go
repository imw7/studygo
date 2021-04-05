package main

import (
	"flag"
	"fmt"
	"time"
)

// flag 获取命令行参数
func main() {
	// 创建一个标志位参数
	// name := flag.String("name", "Eric", "What's your name?")
	// age := flag.Int("age", 18, "How old are you?")
	// married := flag.Bool("married", false, "Are you married?")
	// duration := flag.Duration("duration", 0, "How long have you been married?")
	// 使用flag
	// flag.Parse()
	// fmt.Println(*name)
	// fmt.Println(*age)
	// fmt.Println(*married)
	// fmt.Println(*duration)
	// fmt.Printf("%T\n", *duration)

	// 定义命令行参数方式
	var name string
	var age int
	var married bool
	var duration time.Duration
	flag.StringVar(&name, "name", "Eric", "What's your name?")
	flag.IntVar(&age, "age", 18, "How old are you?")
	flag.BoolVar(&married, "married", false, "Are you married?")
	flag.DurationVar(&duration, "duration", 0, "How long have you been married?")

	// 解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, duration)
	// 返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	// 返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	// 返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
