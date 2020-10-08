package main

import "fmt"

func main() {
	s := []string{"a", "b", "c", "d"}
	defer fmt.Println(s) // [a x y d]
	// defer append(s[:1], "x", "y") // 编译错误
	defer func() {
		_ = append(s[:1], "x", "y")
	}()
}
