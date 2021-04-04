package main

import "fmt"

// map

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        // 还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10) // 要估算好该map容量，避免在程序运行期间再动态扩容
	m1["eric"] = 18
	m1["sarah"] = 21

	fmt.Println(m1)
	fmt.Println(m1["eric"])
	// 约定用 ok 接收返回的布尔值
	fmt.Println(m1["edward"]) // 如果不存在这个key拿到对应值类型的零值
	value, ok := m1["jackson"]
	if !ok {
		fmt.Println("no such key")
	} else {
		fmt.Println(value)
	}

	// map遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除
	delete(m1, "sarah")
	fmt.Println(m1)
	// 删除不存在的值
	delete(m1, "hello") // 什么也不做
	fmt.Println(m1)
}
