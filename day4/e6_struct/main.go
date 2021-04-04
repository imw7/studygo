package main

import "fmt"

// 结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个 person 类型的变量 eric
	var eric person
	// 通过字段赋值
	eric.name = "eric"
	eric.age = 29
	eric.gender = "Male"
	eric.hobby = []string{"抽烟", "喝酒", "烫头"}
	// 访问变量 eric 的字段
	fmt.Printf("%T\n", eric)
	fmt.Println(eric.name)

	var sarah person
	sarah.name = "sarah"
	sarah.age = 18
	sarah.gender = "Female"
	sarah.hobby = []string{"读书", "电影"}
	fmt.Printf("type:%T value:%v\n", sarah, sarah)

	// 匿名结构体：多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "hello"
	s.y = 10086
	fmt.Printf("type:%T value:%v\n", s, s)
}
