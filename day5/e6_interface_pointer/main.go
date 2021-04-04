package main

import "fmt"

// 使用值接收者和指针接收者实现接口的区别?
// 使用值接收者实现接口可以存指针类型和值类型的变量
// 使用指针接收者实现接口只能存指针类型的变量

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 使用值接收者实现了接口的所有方法
func (c cat) move() {
	fmt.Println("走猫步...")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

type chicken struct {
	feet int8
}

// 使用指针接收者实现了接口的所有方法
func (c *chicken) move() {
	fmt.Println("鸡动...")
}

func (c *chicken) eat(food string) {
	fmt.Printf("鸡吃%s...\n", food)
}

func main() {
	var a1 animal

	c1 := cat{"tom", 4}  // cat
	c2 := &cat{"假老练", 4} // *cat

	a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)

	c3 := chicken{feet: 4}
	c4 := chicken{feet: 6}

	a1 = &c3 // 实现 animal 这个接口的是 chicken 的指针类型
	fmt.Println(a1)
	a1 = &c4
	fmt.Println(a1)
}
