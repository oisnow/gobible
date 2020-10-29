package main

import "fmt"

//结构体的嵌套
type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	address
}

func main() {
	var tudou = person{
		name: "gaolian",
		age:  13,
		address: address{
			province: "guangdong",
			city:     "shenzhen",
		},
	}
	// fmt.Println(tudou.addr.city)
	fmt.Println(tudou.city)

}
