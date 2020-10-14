package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	ty := reflect.TypeOf(x)
	fmt.Printf("type:%v,kind:%v \n", ty.Name(), ty.Kind())
}

func reflectValue(x interface{}) {
	vl := reflect.ValueOf(x)
	fmt.Printf("value:%v \n", vl)

}

type cat struct {
	name string
}

type person struct {
	name string
	age  int
}

func main() {
	reflectType(100)
	reflectValue(100)

	c1 := cat{name: "咪咪"}
	p1 := person{
		name: "happy",
		age:  4,
	}

	reflectType(&c1)
	reflectType(p1)

	reflectValue(&c1)
	reflectValue(p1)
}
