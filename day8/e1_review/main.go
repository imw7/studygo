package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel

func sendNum(ch chan<- int) {
	for {
		// rand.Seed(time.Now().UnixNano())
		rand.New(rand.NewSource(time.Now().UnixNano()))
		num := rand.Intn(10)
		ch <- num
		time.Sleep(5 * time.Second)
	}
}

func main() {
	ch := make(chan int, 1)
	// ch <- 100     // 把一个值发送到通道中
	// <-ch          // 把通道中100取出来
	// x, ok := <-ch // 再取阻塞
	go sendNum(ch)
	for {
		x, ok := <-ch // 阻塞等4秒钟
		fmt.Println(x, ok)
		time.Sleep(time.Second)
	}
}
