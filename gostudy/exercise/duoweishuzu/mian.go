package main

import "fmt"

func main() {
	// var a [3]int
	a := [3]int{1, 2, 3}

	// var b [3][2]int
	//多维数组只能有第一层用... 其他层不可以
	b := [3][2]int{
		[2]int{1, 1},
		[2]int{2, 2},
	}
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("---------------------------------")
	//多为数组遍历
	for key, val := range b {

		fmt.Println(key, val)
		for _, v := range val {
			fmt.Println(key, v)
		}
	}
}
