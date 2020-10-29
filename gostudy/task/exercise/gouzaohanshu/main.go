package main

import "fmt"

// Person 定义后可以被其他类引用
type Person struct {
	name  string
	age   int16
	sex   string
	hobby []string
}

//NewPerson 实现构造函数
func NewPerson(name, sex string, age int16, hobby []string) *Person {
	return &Person{
		name:  name,
		age:   age,
		sex:   sex,
		hobby: hobby,
	}
}

func main() {
	nd := NewPerson("gaolian", "woman", 18, []string{"man", "baby"})
	fmt.Println(*nd)
}
