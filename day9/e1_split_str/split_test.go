package split_str

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	// 定义一个测试用例类型
// 	type test struct {
// 		got      string
// 		sep      string
// 		excepted []string
// 	}
// 	// 定义一个存储测试用例的切片
// 	tests := []test{
// 		{got: "a:b:c", sep: ":", excepted: []string{"a", "b", "c"}},
// 		{got: "a:b:c", sep: ",", excepted: []string{"a:b:c"}},
// 		{got: "about", sep: "bo", excepted: []string{"a", "ut"}},
// 		{got: "上海水海上来", sep: "上", excepted: []string{"海水海", "来"}},
// 	}
// 	// 遍历切片，逐一执行测试用例
// 	for _, tc := range tests {
// 		got := Split(tc.got, tc.sep)
// 		if !reflect.DeepEqual(got, tc.excepted) {
// 			t.Fatalf("excepted:%v, got:%v\n", tc.excepted, got)
// 		}
// 	}
// }

// 子测试
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		got      string
		sep      string
		excepted []string
	}

	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {got: "a:b:c", sep: ":", excepted: []string{"a", "b", "c"}},
		"wrong sep":   {got: "a:b:c", sep: ",", excepted: []string{"a:b:c"}},
		"more sep":    {got: "about", sep: "bo", excepted: []string{"a", "ut"}},
		"leading sep": {got: "上海水海上来", sep: "上", excepted: []string{"", "海水海", "来"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.got, tc.sep)
			if !reflect.DeepEqual(got, tc.excepted) {
				t.Errorf("name:%s excepted:%#v, got:%#v", name, tc.excepted, got) // 将测试用例的name格式化输出
			}
		})
	}
}

// BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

// 错误的示范
// func BenchmarkSplit2(b *testing.B) {
// 	Fib(b.N)
// }

// 性能比较测试
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}
