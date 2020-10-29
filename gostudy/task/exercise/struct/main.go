package main

import "fmt"

type test struct {
	a int16
	b int16
	c int16
}

type student struct {
	name string
	age  int8
}

func main() {
	var t = test{
		1,
		2,
		3,
	}
	fmt.Println(t)
	fmt.Println(&(t.a))
	fmt.Println(&(t.b))
	fmt.Println(&(t.c))

	var stu1 = student{
		"haoyunpeng",
		18,
	}
	stu2 := stu1
	stu2.name = "gaolian"

	fmt.Println(stu1.name)
	fmt.Println(stu2.name)

	stu3 := &stu1

	stu3.name = "gaoxing"

	fmt.Println(stu3.name)
	fmt.Println(stu1.name)

}
