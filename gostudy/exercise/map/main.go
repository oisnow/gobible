package main

import "fmt"

func main() {
	var m1 map[string]int

	m1 = make(map[string]int, 100)
	m1["gaolian"] = 100
	fmt.Println(m1)
	fmt.Println(len(m1))
	// fmt.Println(cap(m1))

}
