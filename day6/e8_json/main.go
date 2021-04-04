package main

import (
	"encoding/json"
	"fmt"
)

// json

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name": "Eric", "age": 18}`
	var p person
	_ = json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)
}
