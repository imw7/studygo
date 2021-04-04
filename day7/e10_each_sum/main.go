package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序。

1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3. 主goroutine从resultChan取出结果并打印到终端输出
*/

// job ...
type job struct {
	value int64
}

// result ...
type result struct {
	job *job
	sum int64
}

var wg sync.WaitGroup

func randomInt64(n chan<- *job) {
	defer wg.Done()
	// 循环生成 int64 类型的随机数，发送到 jobChan
	for {
		rand.Seed(time.Now().UnixNano())
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		n <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func eachSum(n <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	// 从 jobChan 中取出随机数计算各位数的和，将结果发送到 resultChan
	for {
		job := <-n
		sum := int64(0)
		x := job.value
		for x > 0 {
			sum += x % 10
			x /= 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	jobChan := make(chan *job, 100)
	resultChan := make(chan *result, 100)
	wg.Add(1)
	go randomInt64(jobChan)
	// 开启24个 goroutine 执行 eachSum()
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go eachSum(jobChan, resultChan)
	}
	// 主 goroutine 从 resultChan 取出结果并打印到终端输出
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()
}
