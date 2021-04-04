package main

import (
	"fmt"
	"math/big"
)

// 递归：函数自己调用自己！
// 递归适合处理那种问题相同/问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件

// 3! = 3 * 2 * 1         = 3 * 2!
// 4! = 4 * 3 * 2 * 1     = 4 * 3!
// 5! = 5 * 4 * 3 * 2 * 1 = 5 * 4!

// 计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 上台阶的面试题
// n个台阶，一次可以走1步，也可以走2步，有多少种走法。
func step(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return step(n-2) + step(n-1)
}

func s(n int64) *big.Int {
	if n <= 2 {
		return big.NewInt(n)
	}
	a := big.NewInt(1)
	b := big.NewInt(2)
	c := big.NewInt(0)
	for i := int64(3); i <= n; i++ {
		c = big.NewInt(0).Add(a, b)
		a = b
		b = c
	}
	return c
}

func main() {
	ret := f(10)
	fmt.Println(ret)

	re := step(45)
	fmt.Println(re)

	r := s(100000)
	fmt.Println(r)
}
