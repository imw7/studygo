package main

import "fmt"

// 数组

// 存放元素的容器
// 数组是同一种数据类型元素的集合。
// 必须指定存放的元素的类型和容量（长度）
// 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。

func main() {
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true ture false false]

	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 数组的初始化
	// 如果不初始化：默认元素都是零值（布尔值：false，整数和浮点数：0，字符串：""）
	fmt.Println(a1, a2)
	// 1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2.初始化方式2
	// a10 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a10)
	// 3.初始化方式3：根据索引来初始化
	a3 := [5]int{0: 1, 4: 2} // [1 0 0 0 2]
	fmt.Println(a3)

	// 数组的遍历
	cities := [...]string{"北京", "上海", "广州", "深圳"}
	// 1.根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	// 2. for range遍历
	for _, city := range cities {
		fmt.Println(city)
	}

	// 多维数组
	// [[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)

	// 遍历
	for _, v1 := range a11 {
		fmt.Println()
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3} // [1 2 3]
	b2 := b1              // [1 2 3] Ctrl+C Ctrl+V
	b2[0] = 100           // b2:[100 2 3]
	fmt.Println(b1, b2)   // b1:? => [1 2 3]

	// 数组指针
	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a                // b 是指向数组的指针
	fmt.Println(a[0], a[1])   // 打印数组的前2个元素
	fmt.Println(b[0], b[1])   // 通过数组指针访问数组元素的方式和数组类似
	fmt.Printf("b:%T\n", b)   // 打印类型
	fmt.Printf("b:%#v\n", b)  // 打印详细信息

	for i, v := range b { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}
	// for range 忽略迭代时的下标
	var times [5][0]int // 内存大小为0
	for range times {
		fmt.Println("hello")
	}

	// 长度为0的数组在内存中并不占用空间
	// 我们并不关心管道中传输数据的真实类型
	// 其中管道接收和发送操作只是用于消息的同步
	// 用空数组来作为管道类型可以减少管道元素赋值时的开销
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1
	// 用无类型的匿名结构体代替长度为0的数组
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{} // struct{}部分是类型，{}表示对应的结构体值
	}()
	<-c2
}
