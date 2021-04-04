package main

import "fmt"

// 接口示例2
// 不管什么牌子的车，都能跑

// 定义一个 car 接口类型
// 不管是什么结构体 只要有 run() 方法都是 car 类型
type car interface {
	run()
}

type ferrari struct {
	brand string
}

func (f ferrari) run() {
	fmt.Printf("%s速度490迈~\n", f.brand)
}

type porsche struct {
	brand string
}

func (p porsche) run() {
	fmt.Printf("%s速度500迈~\n", p.brand)
}

// drive函数接收一个 car 类型的变量
func drive(c car) {
	c.run()
}

func main() {
	f1 := ferrari{
		brand: "法拉利",
	}
	b1 := porsche{
		brand: "保时捷",
	}

	drive(f1)
	drive(b1)
}
