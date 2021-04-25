package main

import (
	"fmt"
	split "github.com/imw7/studygo/day9/e1_split_str"
)

func main() {
	ret1 := split.Split("a:b:c", ":")
	fmt.Printf("%#v\n", ret1)

	ret2 := split.Split("a:b:c", ",")
	fmt.Printf("%#v\n", ret2)

	ret3 := split.Split("about", "bo")
	fmt.Printf("%#v\n", ret3)

	ret4 := split.Split("上海水海上来", "上")
	fmt.Printf("%#v\n", ret4)
}
