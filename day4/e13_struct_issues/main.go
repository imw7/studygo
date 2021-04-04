package main

import "fmt"

// 结构体遇到的问题

type myInt int

type person struct {
	name string
	age  int
}

func (m myInt) hello() {
	fmt.Println("I'm an int.")
}

func main() {
	// 问题1：myInt(100)是什么?
	// 声明一个int32类型的变量x，它的值是10
	// 方法1：
	// var x int32
	// x = 10
	// 方法2：
	// var x int32 = 10
	// 方法3：
	// var x = int32(10)
	// 方法4：
	// x := int32(10)
	// fmt.Println(x)

	// 声明一个myInt类型的变量m，它的值是100
	// 方法1
	// var m myInt
	// m = 100
	// 方法2
	// var m myInt = 100
	// 方法3
	// var m = myInt(100)
	// 方法4
	// m := myInt(100) // 强制类型转换
	// fmt.Println(m)
	m := myInt(100)
	m.hello()

	// 问题2：结构体初始化
	// 方法1：
	var p person // 声明一个person类型的变量p
	p.name = "john"
	p.age = 14
	fmt.Println(p)
	var p1 person
	p1.name = "charlie"
	p1.age = 19
	fmt.Println(p1)
	// 方法2：
	s1 := []int{1, 2, 3, 4}
	m1 := map[string]int{
		"stu1": 100,
		"stu2": 79,
		"stu3": 89,
	}
	fmt.Println(s1, m1)
	// 键值对初始化
	var p2 = person{
		name: "peter",
		age:  49,
	}
	fmt.Println(p2)
	// 值列表初始化
	var p3 = person{
		"ada",
		20,
	}
	fmt.Println(p3)

	jack := newPerson("jack", 18)
	fmt.Printf("%#v\n", jack)
}

// 问题3：为什么要有构造函数？
func newPerson(name string, age int) person {
	// 别人调用我，我能给他一个person类型的变量
	return person{
		name: name,
		age:  age,
	}
}

// func newPerson(name string, age int) *person {
// 	// 别人调用我，我能给他一个person类型的变量
// 	return &person{
// 		name: name,
// 		age:  age,
// 	}
// }
