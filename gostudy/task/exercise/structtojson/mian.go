package main

import (
	"encoding/json"
	"fmt"
)

//json序列化

// Person 定义Person结构体
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Student 定义Ptudent结构体
type Student struct {
	School string `json:"school"`
	Person
}

func main() {
	var happy = Student{
		School: "longhuaerxiao",
		Person: Person{
			Name: "haochenxi",
			Age:  12,
		},
	}

	fmt.Println(happy)
	v, err := json.Marshal(happy)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	fmt.Println(string(v))
	fmt.Printf("%#v \n", string(v))

	str := "{\"School\":\"longhuaerxiao\",\"Name\":\"haochenxi\",\"Age\":12}"
	var Hp = new(Student)
	json.Unmarshal([]byte(str), Hp)
	fmt.Println("struct:", *Hp)
	fmt.Printf("%T\n", *Hp)
}
