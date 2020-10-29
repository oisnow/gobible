package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello(i int) {
	defer wg.Done()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic error,and recover", err)
		}

	}()

	fmt.Println("Im a thread", i)
	// if i == 10 {
	// 	panic("panic error")
	// }
	var a [10]int
	a[i] = 111 //当x为20时候，导致数组越界，产生一个panic，导致程序崩溃

}

var wg sync.WaitGroup

func main() {

	defer fmt.Println("all done!")

	runtime.GOMAXPROCS(10)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
	fmt.Println("I'm main process")

}
