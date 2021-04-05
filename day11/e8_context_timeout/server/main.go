package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// 服务端，随机出现慢响应

func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10) // 耗时10秒的慢响应
		_, _ = fmt.Fprintf(w, "slow response")
		return
	}
	_, _ = fmt.Fprintf(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
