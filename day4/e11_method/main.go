package main

import (
	"fmt"
	"strings"
)

// 方法

// 标识符：变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母大写，表示对外可见（暴露的，公有的）

type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) bark() {
	fmt.Printf("%s: raf raf raf~\n", d.name)
}

// 使用值接收者：传拷贝进去
func (p person) addAge() {
	p.age++
}

// 使用指针接收者：传内存地址进去
func (p *person) trueAddAge() {
	p.age++
}

func (p *person) dream() {
	fmt.Printf("%s's dream is to be a great programmer.\n", strings.Title(p.name))
}

func main() {
	d1 := newDog("peter")
	d1.bark()

	p1 := newPerson("eric", 19)
	fmt.Println(p1.age) // 19
	p1.addAge()
	fmt.Println(p1.age) // ?19
	p1.trueAddAge()
	fmt.Println(p1.age) // ?20
	p1.dream()
}
