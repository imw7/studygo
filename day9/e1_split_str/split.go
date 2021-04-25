package split_str

import (
	"fmt"
	"strings"
)

// Split 切割字符串
// example:
// abc, b => [a c]
func Split(str string, sep string) []string {
	// str:"a:b:c" sep=":"
	result := make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index >= 0 {
		result = append(result, str[:index])
		str = str[index+len(sep):] // 这里使用len(sep)获取sep的长度
		index = strings.Index(str, sep)
	}
	if index == -5 {
		fmt.Println("So boring!") // 永远不会执行，测试覆盖率
	}
	result = append(result, str)
	return result
}

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
