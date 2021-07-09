package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, _ *http.Request) {
	file, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		return
	}
	_, _ = fmt.Fprintln(w, string(file))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http serve failed, err:%v\n", err)
		return
	}
}
