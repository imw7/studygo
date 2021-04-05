package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

var (
	x      int64
	wg     sync.WaitGroup
	rwLock sync.RWMutex
)

func write() {
	// lock.Lock() // 加互斥锁
	rwLock.Lock() // 加写锁
	x++
	time.Sleep(time.Millisecond * 10) // 假设写操作耗时10毫秒
	rwLock.Unlock()                   // 解写锁
	// lock.Unlock() // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock() // 加互斥锁
	rwLock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwLock.RUnlock()             // 解读锁
	// lock.Unlock() // 解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	duration := time.Now().Sub(start)
	fmt.Printf("result:%v duration:%v\n", x, duration)
}
