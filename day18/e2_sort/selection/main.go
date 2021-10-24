package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 选择排序

func selectionSort(arr []int) {
	if len(arr) < 2 || arr == nil {
		return
	}
	for i := 0; i < len(arr)-1; i++ { // i ~ N-1
		minIndex := i
		for j := i + 1; j < len(arr); j++ { // i ~ N-1上找最小值的下标
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
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
	selectionSort(arr)
	fmt.Println(arr)
}
