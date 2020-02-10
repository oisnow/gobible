package main

import (
	"fmt"
)

func main() {
	//求数组的和
	a := [...]int{1, 3, 4}
	var sum int
	for _, val := range a {
		// fmt.Println(val)
		sum = sum + val
	}
	fmt.Println(sum)

	fmt.Println("---------------------------------")
	//找出数组中为指定值的两个元素的小标，比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	b := []int{1, 3, 5, 7}

	for key, val := range b {
		for k, v := range b {
			if v+val == 8 {
				fmt.Printf("%d--%d \n", key, k)
			}
		}

	}
	fmt.Println("---------------------------------")
	//找出数组中为指定值的两个元素的小标，比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	c := []int{1, 3, 5, 7}
	for i := 0; i < len(c); i++ {
		for j := i + 1; j < len(c); j++ {
			if c[i]+c[j] == 8 {
				fmt.Printf("%d--%d \n", i, j)
			}
		}
	}
}
