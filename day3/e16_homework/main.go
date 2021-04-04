package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	left = coins
	// 1.依次拿到每个人的名字
	for _, name := range users {
		// 2.拿到一个人名根据分金币的规则去分金币
		e := strings.Count(name, "e") + strings.Count(name, "E")
		i := strings.Count(name, "i") + strings.Count(name, "I")
		o := strings.Count(name, "o") + strings.Count(name, "O")
		u := strings.Count(name, "u") + strings.Count(name, "U")
		each := e*1 + i*2 + o*3 + u*4
		// 	1>每个人分的金币数应该保存到 distribution 中
		distribution[name] = each
		// 	2>还要记录下剩余的金币数
		left -= each
	}
	// 3.整个第2步执行完就能得到最终每个人分的金币数和剩余金币数
	// fmt.Println(distribution)
	for name, value := range distribution {
		if value > 1 {
			fmt.Printf("%s got %d coins.\n", name, value)
		} else {
			fmt.Printf("%s got %d coin.\n", name, value)
		}
	}
	return left
}

func main() {
	left := dispatchCoin()
	fmt.Printf("剩余%d个金币\n", left)
}
