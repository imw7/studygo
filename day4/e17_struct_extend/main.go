package main

import "fmt"

// 结构体模拟实现其他语言中的“继承”

type animal struct {
	name string
}

// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s can move!\n", a.name)
}

// 狗类
type dog struct {
	feet   int8
	animal // animal拥有的方法，dog此时也有了
}

// 给dog实现一个叫的方法
func (d dog) bark() {
	fmt.Printf("%s is barking: raf raf raf~\n", d.name)
}

func main() {
	d1 := dog{
		feet:   4,
		animal: animal{name: "Bob"},
	}
	fmt.Println(d1)
	d1.bark()
	d1.move()
}
