package main

import (
	"encoding/json"
	"fmt"
)

//json序列化

// Person 定义Person结构体
type Person struct {
	name string
	age  int
}

// Student 定义Ptudent结构体
type Student struct {
	school string
	Person
}

func main() {
	var happy = Student{
		school: "longhuaerxiao",
		Person: Person{
			name: "haochenxi",
			age:  12,
		},
	}

	var hao = Person{
		name: "haoyunpeng",
		age:  32,
	}

	fmt.Println(happy)
	v, err := json.Marshal(happy)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	fmt.Println(string(v))

	fmt.Println(hao)
	v1, err1 := json.Marshal(hao)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(v1)
	fmt.Println(string(v1))
}
