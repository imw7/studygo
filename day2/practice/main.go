package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

/*1. 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，
使用`fmt.Printf()`搭配`%T`分别打印出上述变量的值和类型。*/
func dataType() {
	d1 := 12
	d2 := 3.14159265354
	d3 := true
	d4 := "hello go world"
	fmt.Printf("d1: {value: %d, type: %T}", d1, d1)
	fmt.Println()
	fmt.Printf("d2: {value: %f, type: %T}", d2, d2)
	fmt.Println()
	fmt.Printf("d3: {value: %v, type: %T}", d3, d3)
	fmt.Println()
	fmt.Printf("d4: {value: %s, type: %T}", d4, d4)
	fmt.Println()
}

// 2. 编写代码统计出字符串`"Hello你好こんにちはสวัสดีค่ะ안녕하세요"`中汉字的数量。
func characterCount(sentence string, unicodeName *unicode.RangeTable) {
	count := 0
	for _, v := range sentence {
		// Han：汉字
		// Hiragana：平假名， Katakana：片假名
		if unicode.Is(unicodeName, v) {
			count++
		}
	}
	str := ""
	switch unicodeName {
	case unicode.Han:
		str = "汉字"
	case unicode.Hiragana:
		str = "平假名"
	case unicode.Katakana:
		str = "片假名"
	case unicode.Hangul:
		str = "韩文"
	case unicode.Khmer:
		str = "柬埔寨文"
	case unicode.Thai:
		str = "泰文"
	}
	fmt.Printf("%#v中%s的个数为：%d\n", sentence, str, count)
}

// 3. 打印99乘法表。
func multiTable() {
	for x := 1; x < 10; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%dx%d=%d\t", y, x, x*y)
		}
		fmt.Println()
	}
}

// 4.有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
func pickNumber(numbers string) {
	res := 0
	for _, v := range numbers {
		// ^: 两位不一样则为1
		res ^= int(v) // 0异或任何数都是原来的数: x ^ 0 = x
	}
	fmt.Println(string(rune(res)))
}

// 5. 求数组`[1, 3, 5, 7, 8]`所有元素的和
func sumArray(numbers []int) {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println(sum)
}

// 6. 找出数组中和为指定值的两个元素的下标，比如从数组`[1, 3, 5, 7, 8]`中找出和为8的两个元素的下标分别为`(0,3)`和`(1,2)`。
func findIndex(numbers []int, sum int) {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == sum {
				fmt.Printf("[%d, %d]\n", i, j)
			}
		}
	}
}

// 7.请写出下面代码的输出结果。
// func main() {
// 	var a = make([]string, 5, 10) // 创建切片，长度为5，容量为10
// 	for i := 0; i < 10; i++ {
// 		a = append(a, fmt.Sprintf("%v", i))
// 	}
// 	fmt.Println(a) // 输出结果： [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
// }

// 8.请使用内置的`sort`包对数组`var a = [...]int{3, 7, 8, 9, 1}`进行排序。
func sortArray() {
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	fmt.Println(a)
}

// 9.写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
func countWords(sentence string) {

	ss := strings.Split(sentence, " ")
	wordsMap := make(map[string]int, len(ss))

	for _, value := range ss {
		wordsMap[value]++
	}

	fmt.Println(wordsMap)
}

// 10.观察下面代码，写出最终的打印结果。
/* func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) // [1 2 3]
	m["john"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s) // [1 3]
	fmt.Printf("%+v\n", m["john"]) // [1 3 3]
} */

func main() {
	dataType()

	s := "Hello你好こんにちはสวัสดีค่ะ안녕하세요"
	characterCount(s, unicode.Han)
	characterCount(s, unicode.Hiragana)
	characterCount(s, unicode.Hangul)
	characterCount(s, unicode.Thai)

	multiTable()

	numbers := "7328917825931"
	pickNumber(numbers) // 5

	numbers1 := []int{1, 3, 5, 7, 8}
	sumArray(numbers1)

	numbers2 := []int{1, 3, 5, 7, 8, 9, -1, 0, 2, 6}
	sum := 8
	findIndex(numbers2, sum)

	sortArray()

	sentence := "how do you do"
	countWords(sentence)
}
