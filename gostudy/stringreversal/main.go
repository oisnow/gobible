package main

import "fmt"

var s1 string = "hello"
var s2 string

func main() {
	// fmt.Println(s1)
	// var s2 string
	byteArryS1 := []byte(s1)
	for i := len(byteArryS1) - 1; i >= 0; i-- {
		// fmt.Println(i)
		// fmt.Println(byteArryS1[i])

		s2 = fmt.Sprint(s2, string(byteArryS1[i]))
	}
	fmt.Println(s2)
}
