package main

import "fmt"

type cat struct{}

func (c cat) cell() string {
	return "喵喵喵"
}

type dog struct{}

func (d dog) cell() string {
	return "汪汪汪"
}

type pig struct{}

func (p pig) cell() string {
	return "哼哼哼"
}

type animals interface {
	cell() string
}

func main() {

	var c = cat{}
	var d = dog{}
	var p = pig{}

	var animalslist []animals
	animalslist = append(animalslist, c, d, p)
	for _, ani := range animalslist {
		ret := ani.cell()
		fmt.Println(ret)
	}
}
