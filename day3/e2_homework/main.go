package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1.判断字符串中汉字的数量
	// 难点是判断一个字符是汉字
	s1 := "Hello中国こんにちは안녕하세요"
	count := 0
	// 1> 依次拿到字符串中的字符
	for _, c := range s1 {
		// 2> 判断当前这个字符是不是汉字
		if unicode.Is(unicode.Han, c) { // 韩文：Hangul 平假名：Hiragana
			// 3> 把汉字出现的次数累加得到最终结果
			count++
		}
	}
	fmt.Printf("%#v中汉字的个数为：%d\n", s1, count)

	// 2.how do you do 单词出现的次数
	s2 := "how do you do"
	// 1>把字符串按照空格切割得到切片
	s3 := strings.Split(s2, " ")
	// 2>遍历切片存储到一个map
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		// 2.1> 如果原来map中不存在w这个key，那么出现次数=1
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			// 2.2> 如果map中存在w这个key，那么出现次数+1
			m1[w]++
		}
	}
	// 3> 累加出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	// 回文判断
	// 字符串从左往右和从右往左是一样的，那么就是回文
	// 上海自来水来自海上 s[0] s[len(s)-1]
	ss := "上海自来水来自海上"
	// 解体思路：
	// 把字符串中的字符拿出来放到一个[]rune中
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		// 上 r[0] r[len(r)-1]
		// 海 r[1] r[len(r)-1-1]
		// 自 r[2] r[len(r)-1-2]
		// ...
		// c r[i] r[len(r)-1-i]
		if r[i] != r[len(r)-1-i] {
			fmt.Printf("%#v不是回文\n", ss)
			return
		}
	}
	fmt.Printf("%#v是回文\n", ss)
}
