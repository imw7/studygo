package main

import (
	"fmt"
	"unicode"
)

// byte和rune类型

// Go语言中为了处理非ASCII码类型的字符 定义了新的rune类型

func main() {
	s := "Hello go 你好世界 こんにちは せかい 안녕하세요 세계"
	n := len(s)    // len()求的是byte字节的数量；求字符串s的长度，把长度保存到变量n中。
	fmt.Println(n) // 70

	countHan, countHiragana, countHangul := 0, 0, 0
	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			countHan++
		}
		if unicode.Is(unicode.Hiragana, v) {
			countHiragana++
		}
		if unicode.Is(unicode.Hangul, v) {
			countHangul++
		}
	}
	fmt.Printf("汉字：%d个\n平假名：%d个\n韩文：%d个\n", countHan, countHiragana, countHangul)

	for _, c := range s { // 从字符串中拿出具体的字符
		fmt.Printf("%c\n", c) // %c：字符
	}

	// "Hello" => 'H' 'e' 'l' 'l' 'o'
	// 字符串修改
	s2 := "白萝卜"      // '白' '萝' '卜'
	s3 := []rune(s2) // 把字符串强制转换成了一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 把rune切片强制转换成字符串 // 红萝卜

	c1 := "红"
	c2 := '红'                           // rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2) // c1:string c2:int32

	c3 := "H"                           // string
	c4 := 'H'                           // int32
	fmt.Printf("c3:%T c4:%T\n", c3, c4) // c3:string c4:int32
	fmt.Printf("%d\n", c4)              // 72

	// 类型转换
	n1 := 10 // int
	var f float64
	f = float64(n1)
	fmt.Println(f)        // 10
	fmt.Printf("%T\n", f) // float64
}
