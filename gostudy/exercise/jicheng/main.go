package main

import "fmt"

//使用结构体的嵌套实现集成
type person struct {
	name string
	age  int
}

func (p *person) whatname() {
	fmt.Println("名字为", p.name)
}

func (s *student) whichsch() {
	fmt.Println("就学为", s.school)
}

type student struct {
	school string
	*person
}

func main() {
	var happy = student{
		school: "龙华二小",
		person: &person{
			name: "haochenxi",
		},
	}

	var tudou = person{
		name: "gaolian",
	}
	tudou.whatname()
	happy.whatname()
	happy.whichsch()
}
