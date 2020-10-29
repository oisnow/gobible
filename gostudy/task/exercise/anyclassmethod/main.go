package main

import "fmt"

//可以给任意类型追加方法
//但是不能给其他包定义的类添加方法

// MyInt 自定义的类型
type MyInt int

func (m *MyInt) sayHi() {
	fmt.Println("hello")
}

func main() {

	var a MyInt
	a.sayHi()
}
