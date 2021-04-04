package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// 从字符串中解析出对应的数字
	str := "10000"
	// ret1 := int64(str)
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parse int failed, err:", err)
		return
	}
	fmt.Printf("value:%#v type:%T\n", result, result)

	// Atoi: 把string类型直接转换为int类型
	ret, _ := strconv.Atoi(str)
	fmt.Printf("value:%#v type:%T\n", ret, ret)

	// 把数字转换成字符串类型
	i := 97
	// ret2 := string(i) // "a"

	ret2 := fmt.Sprintf("%d", i) // "97"
	fmt.Printf("%#v\n", ret2)

	// Itoa: 把int类型直接转换为string类型
	ret3 := strconv.Itoa(i)
	fmt.Printf("value:%#v type:%T\n", ret3, ret3)

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("value:%#v type:%T\n", boolValue, boolValue)

	// 从字符串中解析出浮点数
	floatStr := "3.1415926"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("value:%#v type:%T\n", floatValue, floatValue)
}
