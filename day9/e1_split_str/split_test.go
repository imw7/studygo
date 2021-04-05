package split_str

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	type testCase struct {
// 		got      string
// 		sep      string
// 		excepted []string
// 	}
//
// 	testGroup := []testCase{
// 		{"a:b:c", ":", []string{"a", "b", "c"}},
// 		{"abcbebdbf", "b", []string{"a", "c", "e", "d", "f"}},
// 		{"about", "bo", []string{"a", "ut"}},
// 		{"上海自来水来自海上", "自", []string{"上海", "来水来", "海上"}},
// 	}
//
// 	for _, tc := range testGroup {
// 		got := Split(tc.got, tc.sep)
// 		if !reflect.DeepEqual(got, tc.excepted) {
// 			t.Fatalf("excepted:%v, got:%v\n", tc.excepted, got)
// 		}
// 	}
// }

// 子测试
func TestSplit(t *testing.T) {
	type testCase struct {
		got      string
		sep      string
		excepted []string
	}

	testGroup := map[string]testCase{
		"first": {
			got:      "a:b:c",
			sep:      ":",
			excepted: []string{"a", "b", "c"},
		},
		"second": {
			got:      "abcbebdbf",
			sep:      "b",
			excepted: []string{"a", "c", "e", "d", "f"},
		},
		"third": {
			got:      "about",
			sep:      "bo",
			excepted: []string{"a", "ut"},
		},
		"fourth": {
			got:      "上海自来水来自海上",
			sep:      "自",
			excepted: []string{"上海", "来水来", "海上"},
		},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.got, tc.sep)
			if !reflect.DeepEqual(got, tc.excepted) {
				t.Fatalf("excepted:%v, got:%v\n", tc.excepted, got)
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
