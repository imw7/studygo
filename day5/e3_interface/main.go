package main

import "fmt"

// 引出接口的实例

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了 speak 方法的变量都是 speaker 类型，方法签名
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p person) speak() {
	fmt.Println("啊啊啊~")
}

func hit(x speaker) {
	// 接收一个参数，传进来什么就打什么
	x.speak() // 被打到了要叫
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	hit(c1)
	hit(d1)
	hit(p1)

	var sp speaker // 定义一个接口类型：speaker的变量：sp
	sp = c1
	sp = d1
	sp = p1
	fmt.Println(sp)
}
