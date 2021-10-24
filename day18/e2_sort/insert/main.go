package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 插入排序

func insertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 0~0有序的
	// 0~i想有序
	for i := 1; i < len(arr); i++ { // 0 ~ i 做到有序
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

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
	insertSort(arr)
	fmt.Println(arr)
}
