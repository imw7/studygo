package main

import (
	"fmt"
	"sync"
)

var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBuffChannel() {
	fmt.Println(b)     // <nil>
	b = make(chan int) // 不带缓冲区的通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道中取到了", x)
	}()
	b <- 10 // 死锁
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
}

func buffChannel() {
	fmt.Println(b)         // <nil>
	b = make(chan int, 16) // 带缓冲区的通道的初始化
	b <- 10
	fmt.Println("10发送到通道b中了...")
	b <- 20
	fmt.Println("20发送到通道b中了...")
	x := <-b
	fmt.Println("从通道b中取到了", x)
	close(b)
}

func main() {
	noBuffChannel()
	buffChannel()
}
