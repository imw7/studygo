package main

import "fmt"

// copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 // 赋值
	// var a3 []int // nil
	var a3 = make([]int, 3, 3) // 开辟新内存
	copy(a3, a1)               // copy
	fmt.Println(a1, a2, a3)    // [1 3 5] [1 3 5] [1 3 5]
	a1[0] = 100                // 修改底层数组
	fmt.Println(a1, a2, a3)    // [100 3 5] [100 3 5] [1 3 5]

	// 将a1中的索引为1的3这个元素删掉
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)      // [100 5]
	fmt.Println(cap(a1)) // 3

	x1 := []int{1, 3, 5}              // 数组
	s1 := x1[:]                       // 切片
	fmt.Println(s1, len(s1), cap(s1)) // [1 3 5] 3 3
	// 1.切片不保存具体的值
	// 2.切片对应一个底层数组
	// 3.底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", &s1[0])
	s1 = append(s1[:1], s1[2:]...) // 修改了底层数组！！！
	fmt.Printf("%p\n", &s1[0])
	fmt.Println(s1, len(s1), cap(s1)) // [1 5] 2 3

	fmt.Println(x1) // [1 5 5]
	s1[0] = 100     // 修改了底层数组
	fmt.Println(x1) // [100 5 5]

	// 可以用copy和append组合可以避免创建中间的临时切片
	a := []int{1, 2, 3}
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))
	a = append(a, 0) // 切片扩展1个空间
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))
	copy(a[1+1:], a[1:]) // a[1:]向后移动1个位置
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))
	a[1] = 8 // 设置新添加的元素
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))

	// 用copy和append组合实现在中间位置插入多个元素(也就是插入一个切片)
	a = []int{1, 2, 3}
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))
	x := []int{9, 10, 11}
	a = append(a, x...) // 为x切片扩展足够的空间
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))
	copy(a[1+len(x):], a[1:]) // a[i:]向后移动len(x)个位置
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))
	copy(a[1:], x) // 复制新添加的切片
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))
}
