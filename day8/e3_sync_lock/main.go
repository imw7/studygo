package main

import (
	"fmt"
	"sync"
)

// 互斥锁

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x++
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
