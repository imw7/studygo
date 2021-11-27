package main

import "fmt"

// Options 结构体
type Options struct {
	strOp1 string
	strOp2 string
	strOp3 string
	intOp1 int
	intOp2 int
	intOp3 int
}

// InitOptions1 初始化结构体
func InitOptions1(strOp1, strOp2, strOp3 string, intOp1, intOp2, intOp3 int) {
	options := Options{
		strOp1: strOp1,
		strOp2: strOp2,
		strOp3: strOp3,
		intOp1: intOp1,
		intOp2: intOp2,
		intOp3: intOp3,
	}
	fmt.Printf("init option %#v\n", options)
}

// InitOptions2 初始化结构体
func InitOptions2(opts ...interface{}) {
	options := Options{}
	// 遍历参数
	for index, opt := range opts {
		switch index {
		case 0:
			str, ok := opt.(string)
			if !ok {
				return
			}
			options.strOp1 = str
		case 1:
			str, ok := opt.(string)
			if !ok {
				return
			}
			options.strOp2 = str
		case 2:
			str, ok := opt.(string)
			if !ok {
				return
			}
			options.strOp3 = str
		case 3:
			intOp1, ok := opt.(int)
			if !ok {
				return
			}
			options.intOp1 = intOp1
		case 4:
			intOp2, ok := opt.(int)
			if !ok {
				return
			}
			options.intOp2 = intOp2
		case 5:
			intOp3, ok := opt.(int)
			if !ok {
				return
			}
			options.intOp3 = intOp3
		}
	}
	fmt.Printf("init option %#v\n", options)
}

// 选项设计模式

// Option 声明一个函数类型的变量，用于传参
type Option func(opts *Options)

// InitOptions3 初始化结构体
func InitOptions3(opts ...Option) {
	options := &Options{}
	// 遍历opts，得到每一个函数
	for _, opt := range opts {
		// 调用函数，在函数里，给传进去的对象赋值
		opt(options)
	}
	fmt.Printf("init options %#v\n", options)
}

// WithStrOptions1 定义具体给某个字段赋值的方法
func WithStrOptions1(str string) Option {
	return func(opts *Options) {
		opts.strOp1 = str
	}
}

func WithStrOptions2(str string) Option {
	return func(opts *Options) {
		opts.strOp2 = str
	}
}

func WithStrOptions3(str string) Option {
	return func(opts *Options) {
		opts.strOp3 = str
	}
}

func WithIntOptions1(i int) Option {
	return func(opts *Options) {
		opts.intOp1 = i
	}
}

func WithIntOptions2(i int) Option {
	return func(opts *Options) {
		opts.intOp2 = i
	}
}

func WithIntOptions3(i int) Option {
	return func(opts *Options) {
		opts.intOp3 = i
	}
}

func main() {
	InitOptions1("str1", "str2", "str3", 1, 2, 3)
	InitOptions2("str1", "str2", "str3", 1, 2, 3)
	InitOptions3(WithStrOptions1("str1"), WithStrOptions2("str2"), WithIntOptions1(100))
	InitOptions3(WithStrOptions3("hello"), WithIntOptions2(1), WithIntOptions3(2))
}
