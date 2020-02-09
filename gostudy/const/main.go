package main

import (
	"fmt"
)

const (
	a = iota
	b = 100
	// b = iota
	c = iota
	d = iota
)
const e = iota

//定义存储大小
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Printf("a:%d,b:%d,c:%d,d:%d,e:%d", a, b, c, d, e)
	fmt.Println(KB)
}
