package main

import "fmt"

//定义结构体匿名字段

type person struct {
	name string
	int
}

func main() {
	var gao = person{
		name: "gaolian",
	}

	fmt.Println(gao.name)

	gao.int = 10
	fmt.Println(gao.int)
}
