package main

/*// 1.在A和B处填入代码，使输出为foo

type S struct {
	m string
}

func f() *S {
	return &S{m: "foo"} // A
}

func main() {
	p := f() // B
	print(p.m)
}
*/

/*// 2.下面代码输出是什么，若想输出012，怎么改

const N = 3

func main() {
	m := make(map[int]*int) // m := make(map[int]int)
	for i := 0; i < N; i++ { // i = 3
		m[i] = &i // m[i] = i
	}
	for _, v := range m {
		print(*v) // 333 // print(v) // 012
	}
}
*/

/*// 3.代码输出什么？为什么？如何改会使得len(m)为10
import "sync"

const N = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		// go func() {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
		// }()
	}
	wg.Wait()
	println(len(m)) // 输出不一定，我的电脑ubuntu系统是 1 或 2 或 3
}
*/

// 4.请描述golang语言的初始化顺序：包，变量，常量，init函数
// main包，import，常量，全局变量，init函数，main()函数

/*// 5.描述下面代码输出
import "fmt"

type S1 struct {}

func (s1 S1) f() {
	fmt.Println("S1.f()")
}

func (s1 S1) g() {
	fmt.Println("S1.g()")
}

type S2 struct {
	S1
}

func (s2 S2) f() {
	fmt.Println("S2.f()")
}

type I interface {
	f()
}

func printType(i I) {
	if s1, ok := i.(S1); ok {
		s1.f()
		s1.g()
	}
	if s2, ok := i.(S2); ok{
		s2.f()
		s2.g()
	}
}

func main() {
	printType(S1{}) // S1.f() S1.g()
	printType(S2{}) // S2.f() S1.g()
}
*/

/*// 6.下面代码有什么问题，怎么修改
//  fatal error: concurrent map writes
//  加锁即可解决
import (
	"math/rand"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{} // 1
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			lock.Lock() // 2
			m[rand.Int()] = rand.Int()
			lock.Unlock() // 3
		}()
	}
	wg.Wait()
	println(len(m))
}
*/

// 7.请描述make和new的区别
//  ① 二者都是用来做内存分配的。
//  ② make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
//  ③ 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

/*// 8.下面代码输出什么，如何让输出为true
import "fmt"

type S struct {
	a, b, c string
}

func main() {
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})
	fmt.Println(x == y) // false

	m := interface{}(S{"a", "b", "c"})
	n := interface{}(S{"a", "b", "c"})
	fmt.Println(m == n) // true
}
*/

/*// 9.下面代码问题是什么，如何修改

type S struct {
	name string
}

func main() {
	// map里结构体无法直接寻址，必须取址
	m := map[string]*S{"x": &S{"one"}}
	m["x"].name = "two"
	fmt.Println(m["x"].name)
}
*/

/*// 10.修改代码，使得status输出为200
import (
	"encoding/json"
	"fmt"
)

type Result struct {
	// status int // 反序列化必须是大写字母开头
	Status int
}

func main() {
	var data = []byte(`{"status":200}`)
	result := &Result{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("result=%+v", result)
}
*/

// 11.描述golang中的stack和heap的区别，分别在什么情况下会分配stack？又在何时会分配到heap中？
//  区别：
// 		栈（stack）：由编译器自动分配和释放，存变量名、各种名；
// 		堆（heap）：在C里由程序员分配和释放内存，go自动了，存栈变量的数据值。
//  a := 3  a就在栈，3在堆。
