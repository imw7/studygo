package main

import "fmt"

// 如何判断一个链表有没有闭环
// x一次走1步，y一次走2步。让它们去走，如果某一时刻
// 它们能在同一节点相遇，那么就说明这个链表有闭环。

// 一个节点
type a struct {
	value int
	next  *a
}

// 有n级台阶，一次可以迈1级台阶或者2级台阶，问从底部到顶部共有多少种走法？
func f(n int) int {
	if n < 2 {
		return n
	}
	return f(n-1) + f(n-2)
}

func main() {
	fmt.Println(f(10))
}
