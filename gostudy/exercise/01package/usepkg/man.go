package main

import (
	"fmt"

	"gobible/gostudy/exercise/01package/mypkg"
)

func init() {
	var x int
	fmt.Println(x)
}

func main() {
	fmt.Println(mypkg.Adds(1, 2))
}
