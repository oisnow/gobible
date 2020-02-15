package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	func() {
		fmt.Println("匿名函数")
	}()
	r := makeSuffixFunc(".txt")
	s := "sj.txt"
	v := "tq.avi"
	fmt.Println(r(s))
	fmt.Println(r(v))
}
