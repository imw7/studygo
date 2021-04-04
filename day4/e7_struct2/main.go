package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string
}

// Go语言中函数传参数永远传的是拷贝
func f(x person) {
	x.gender = "Female" // 修改的是副本的gender
}

func f2(x *person) {
	(*x).gender = "Female" // 根据内存地址找到那个原变量，修改的就是原来的变量
	x.gender = "Female"    // 语法糖，自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "Phil"
	p.gender = "Male"
	f(p)
	fmt.Println(p.gender) // Male
	f2(&p)
	fmt.Println(p.gender) // Female

	// 结构体指针1
	var p2 = new(person)
	(*p2).name = "edward"
	p2.gender = "Secret"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  // p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) // 求p2的内存地址

	// 结构体指针2
	// 1> key-value初始化
	var p3 = &person{
		name:   "bob",
		gender: "Male",
	}
	fmt.Printf("%#v\n", p3)
	// 2> 使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一致
	p4 := &person{
		name:   "jack",
		gender: "Male",
	}
	fmt.Printf("%#v\n", p4)
}
