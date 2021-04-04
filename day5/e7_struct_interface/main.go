package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// cat 实现了 mover 接口
func (c *cat) move() {
	fmt.Println("走猫步...")
}

// cat 实现了 eater 接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a1 animal

	c1 := &cat{
		name: "Joe",
		feet: 4,
	}

	a1 = c1
	fmt.Println(a1)
	a1.move()
	a1.eat("小黄鱼")
}
