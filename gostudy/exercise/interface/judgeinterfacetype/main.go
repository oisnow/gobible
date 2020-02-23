package main

import (
	"fmt"
)

type cat struct {
	name  string
	color string
}

func (c cat) speak() {
	fmt.Println("我是cat，我喵喵叫")
}

type dog struct {
	name  string
	color string
}

func (d *dog) speak() {
	fmt.Println("我是dog，我汪汪叫")
}

type animals interface {
	speak()
}

// JudgeType 判断interface类型
func JudgeType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("x is int type,value:", v)
	case string:
		fmt.Println("x is string type,value:", v)
	case bool:
		fmt.Println("x is bool type,value:", v)
	case cat:
		fmt.Println("x is cat sturt{} type,value:", v)

	}

}

func main() {
	var c = cat{
		name:  "花花",
		color: "黄色",
	}
	JudgeType(c)
	JudgeType(10)

	var x animals

	// var d = dog{
	// 	name:  "哈利",
	// 	color: "黄黑花色",
	// }

	x = c
	fmt.Println(x)

}
