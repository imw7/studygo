package main

import "fmt"

// 切片（slice）

func main() {
	// 切片的定义
	var s1 []int    // 定义一个存放int类型元素的切片
	var s2 []string // 定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true 没有开辟内存空间
	fmt.Println(s2 == nil) // true
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"go", "java", "c"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false
	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 基于一个数组切割，左包含右不包含 [1, 3, 5, 7]
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4] // => [0:4] [1 3 5 7]
	s6 := a1[3:] // => [3:len(a1)] [7 9 11 13]
	s7 := a1[:]  // => [0:len(a1)] [1 3 5 7 9 11 13]
	fmt.Println(s5, s6, s7)
	// 切片的容量是指底层数组的容量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	// 底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))
	// 切片再切割
	s8 := s6[3:] // [13]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))
	// 切片是引用类型，都指向了底层的一个数组。
	fmt.Println("s6:", s6)
	a1[6] = 1300 // 修改了底层数组的值
	fmt.Println("s6:", s6)
	fmt.Println("s8:", s8)

	var a []int // nil切片，和 nil 相等，一般用来表示一个不存在的切片
	fmt.Printf("a:%#v is(a==nil):%v\n", a, a == nil)
	b := []int{} // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
	fmt.Printf("b:%#v is(b==nil):%v\n", b, b == nil)
	c := []int{1, 2, 3} // 有3个元素的切片, len和cap都为3
	fmt.Printf("c:%#v len(c):%d cap(c):%d\n", c, len(c), cap(c))
	d := c[:2] // 有2个元素的切片，len为2，cap为3
	fmt.Printf("d:%#v len(d):%d cap(d):%d\n", d, len(d), cap(d))
	e := c[0:2:cap(c)] // 有2个元素的切片，len为2，cap为3
	fmt.Printf("e:%#v len(e):%d cap(e):%d\n", e, len(e), cap(e))
	f := c[:0] // 有0个元素的切片, len为0, cap为3
	fmt.Printf("f:%#v len(f):%d cap(f):%d\n", f, len(f), cap(f))
	g := make([]int, 3) // 有3个元素的切片, len和cap都为3
	fmt.Printf("g:%#v len(g):%d cap(g):%d\n", g, len(g), cap(g))
	h := make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
	fmt.Printf("h:%#v len(h):%d cap(h):%d\n", h, len(h), cap(h))
	i := make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
	fmt.Printf("i:%#v len(i):%d cap(i):%d\n", i, len(i), cap(i))
}
