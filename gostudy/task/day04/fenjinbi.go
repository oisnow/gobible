package main

import (
	"fmt"
)

/*
50枚金币分给以下几个人"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"
分配规则如下
a.名字中包含e或者E的：1枚金币
b.名字中包含i或者I的：2枚金币
c.名字中包含o或者O的：3枚金币
d.名字中包含u或者U的：4枚金币
写一个程序，计算每个用户可以分配到多少硬币，已经剩下的金币数量
*/

var (
	coins = 50
	users = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth", "sss"}
	dist  = make(map[string]int, len(users))
)

//FjbFunc is a function
func FjbFunc(user []string) map[string]int {
	res := make(map[string]int, len(user))
	// lefecoin := 0
	if len(user) > 0 {
		for _, name := range user {
			for _, char := range name {
				switch char {
				case 'e', 'E':
					res[name] = res[name] + 1
				case 'i', 'I':
					res[name] = res[name] + 2
				case 'o', 'O':
					res[name] = res[name] + 3
				case 'u', 'U':
					res[name] = res[name] + 4
					// default:
					// 	res[name] = 0
				}

			}
		}
	}

	return res

}

func main() {

	var costcion int
	// fmt.Println(coins)
	dist = FjbFunc(users)
	fmt.Println(dist)
	for key, val := range dist {
		fmt.Printf("%s分的%d金币\n", key, val)
		costcion += val
	}
	fmt.Println("剩余金币", coins-costcion)
}
