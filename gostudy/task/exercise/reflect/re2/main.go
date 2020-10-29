package main

import (
	"fmt"
	"reflect"
)

type tst struct {
	name string `config:"name"`
	age  int    `config:"age"`
	sex  string `config:"sex"`
}

func main() {

	var ts = &tst{}
	t := reflect.TypeOf(ts)
	v := reflect.ValueOf(ts).Elem()

	fmt.Println("ts type :", t)
	fmt.Println("ts value :", v)

	n := v.NumField()
	fmt.Println("struct字段项个数", n)
	for i := 0; i < v.NumField(); i++ {
		filed := v.Type().Field(i)
		tag := filed.Tag
		name := tag.Get("config")
		fmt.Println(filed, "||", tag, "||", name)
	}
}
