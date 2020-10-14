package main

import (
	"fmt"
)

func main() {
	//数组 值引用
	a := [3]int{1, 2, 3}
	var b [3]int
	b = a
	fmt.Println(a)
	fmt.Println(b)

	a[0] = 10
	fmt.Println(a)
	fmt.Println(b)

	//切片 地址引用
	a1 := []int{1, 2, 3}
	var b1 []int
	b1 = a1
	fmt.Println(a1)
	fmt.Println(b1)

	a1[0] = 10
	fmt.Println(a1)
	fmt.Println(b1)

}
