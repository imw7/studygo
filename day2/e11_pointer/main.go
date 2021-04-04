package main

import "fmt"

// 指针
// 1.make和new都是用来申请内存的
// 2.new很少用，一般用来给基本数据类型申请内存（string、int），返回的是对应类型的指针（*string、*int）
// 3.make是用来给slice、map、chan申请内存的，make函数返回的是对应的这三个类型本身

func main() {
	// 1. &:取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p) // *int: int类型的指针

	// 2. *:根据地址取值
	m := *p
	fmt.Println(m)        // 18
	fmt.Printf("%T\n", m) // int

	var a1 *int       // nil pointer
	fmt.Println(a1)   // <nil>
	var a2 = new(int) // new函数申请一个内存地址
	fmt.Println(a2)
	fmt.Println(*a2) // 0
	*a2 = 100
	fmt.Println(*a2) // 100
}
