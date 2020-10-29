package main

import "fmt"

func a() {
	fmt.Println("a")
}
func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	fmt.Println("b")
	panic("panic in b")
}
func c() {
	fmt.Println("c")
}

func main() {
	a()
	b()
	c()
	var c []int
	a := []int{1, 2, 3}
	fmt.Printf("%T", a)
	copy(a, c)
	fmt.Println(a, c)

}
