package main

import "fmt"

func main() {
	// var f = func() {
	// 	fmt.Println(false)
	// }
	// defer f()
	// f = func() {
	// 	fmt.Println(true)
	// }

	// a := 10
	// defer fmt.Println(a)
	// a = a + 1
	// return

	var t int
	r := 1
	t = r
	r = r + 1
	fmt.Println(t)
	fmt.Println(r)

}
