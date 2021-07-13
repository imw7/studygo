package main

import "fmt"

// append() 为切片追加元素

func main() {
	s1 := []string{"北京", "上海", "广州", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// s1[4] = "成都" // 错误的写法：会导致编译错误，索引越界
	// fmt.Println(s1)

	// 调用append函数必须用原来的切片变量接收返回值
	// append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	// 必须使用变量接收append的返回值
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "武汉")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"西安", "苏州", "厦门"}
	s1 = append(s1, ss...) // ...表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	// 在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。
	// 因此，从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。
	a := []int{1, 2, 3}
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))
	a = append([]int{0}, a...) // 在开头添加1个元素
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	fmt.Printf("a=%v len(a)=%d cap(a)=%d\n", a, len(a), cap(a))
}
