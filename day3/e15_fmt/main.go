package main

import "fmt"

// fmt

func main() {
	fmt.Print("sarah")
	fmt.Print("edward")
	fmt.Println()
	fmt.Println("sarah")
	fmt.Println("edward")
	// Printf("格式化字符串", 值)
	// %T: 查看类型
	// %d: 十进制数
	// %b: 二进制数
	// %o: 八进制数
	// %x: 十六进制数
	// %c: 字符，该值对应的Unicode码值
	// %s: 字符串
	// %p: 指针
	// %v: 值
	// %f: 浮点数
	// %t: 布尔值

	var m1 = make(map[string]int, 1)
	m1["eric"] = 100
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)
	printPercent(90)

	fmt.Printf("%v\n", 100)
	// 整数 -> 字符
	fmt.Printf("%q\n", 65)
	// 浮点数和复数
	fmt.Printf("%b\n", 3.14159265354697)
	// 字符串
	fmt.Printf("%q\n", "Hello world!")
	fmt.Printf("%7.3s\n", "Hello world!") // 宽度7，精度3

	// 获取用户输入
	var s string
	fmt.Print("随便输入点什么: ")
	_, _ = fmt.Scan(&s)
	fmt.Print("用户输入的内容是: ", s+"\n")

	var (
		name  string
		age   int
		class string
	)
	fmt.Print("依次输入姓名、年龄和班级（空格隔开）:")
	_, _ = fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Printf("姓名：%s 年龄：%d 班级：%s\n", name, age, class)

	fmt.Print("依次输入姓名、年龄和班级（空格隔开）:")
	_, _ = fmt.Scanln(&name, &age, &class)
	fmt.Printf("%s %d %s\n", name, age, class)

	fmt.Printf("%b\n", 1024)
}

func printPercent(num int) {
	fmt.Printf("%d%%\n", num)
}
