package main

import (
	"encoding/json"
	"fmt"
)

// 反射

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name": "Eric", "age": 19}`
	var stu student
	_ = json.Unmarshal([]byte(str), &stu)
	fmt.Printf("%#v\n", stu)
}
