package main

import "fmt"

type person struct {
	name   string
	gender string
}

// dream  method of person
func (p person) dream() string {
	return p.name + "跳舞挺性感~~" + "人也挺性感！"
}

func main() {
	var gaolianzi = person{
		name:   "gaolian",
		gender: "man",
	}
	fmt.Println("我是在夸你.....")
	fmt.Println(gaolianzi.dream())
}
