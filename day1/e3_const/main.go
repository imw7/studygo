package main

import "fmt"

// 常量
// 定义了常量之后不能修改
// 在程序运行期间不会改变的值
const pi = 3.1415926

// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一样
const (
	n1 = 100
	n2
	n3
)

// iota 是go语言的常量计数器，只能在常量的表达式中使用。
// iota 在const关键字出现时将被重置为0。const中每新增一
// 行常量声明将使 iota 计数一次（iota可理解为const语句块
// 中的行索引）。使用iota能简化定义，在定义枚举时很有用。
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

const (
	b1 = iota // 0
	b2        // 1
	_         // 2
	b3        // 3
)

// 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4        // 3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1:1 d2:2
	d3, d4 = iota + 1, iota + 2 // d3:2 d4:3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 1左移10位（二进制） // KB: 1024
	MB = 1 << (10 * iota) // MB: 1048576
	GB = 1 << (10 * iota) // GB: 1073741824
	TB = 1 << (10 * iota) // TB: 1099511627776
	PB = 1 << (10 * iota) // PB: 1125899906842624
)

func main() {
	fmt.Println("pi:", pi)
	fmt.Println("statusOK:", statusOK)
	fmt.Println("notFound:", notFound)

	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)
}
