package main

import "fmt"

func main() {
	var x interface{}

	x = 10
	fmt.Println(x)

	s := "string"
	x = s
	fmt.Println(x)

}
