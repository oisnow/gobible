package main

import "fmt"

func makeInt(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	var ch1 = make(chan int, 100)
	go makeInt(ch1)
	for res := range ch1 {
		fmt.Println(res)
	}
}
