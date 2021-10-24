package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 冒泡排序

func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for e := len(arr) - 1; e > 0; e-- { // 0 ~ e
		for i := 0; i < e; i++ {
			if arr[i+1] < arr[i] {
				swap(arr, i, i+1)
			}
		}
	}
}

// 交换arr的i和j位置上的值
func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	arr := make([]int, 0, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)
}
